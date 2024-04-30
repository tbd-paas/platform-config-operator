/*
Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package capabilities

import (
	"context"
	"fmt"

	"github.com/go-logr/logr"
	"github.com/nukleros/operator-builder-tools/pkg/controller/phases"
	"github.com/nukleros/operator-builder-tools/pkg/controller/predicates"
	"github.com/nukleros/operator-builder-tools/pkg/controller/workload"
	apierrs "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/manager"

	capabilitiesv1alpha1 "github.com/tbd-paas/platform-config-operator/apis/capabilities/v1alpha1"
	"github.com/tbd-paas/platform-config-operator/apis/capabilities/v1alpha1/platformconfig"
	"github.com/tbd-paas/platform-config-operator/internal/dependencies"
	"github.com/tbd-paas/platform-config-operator/internal/mutate"
)

// PlatformConfigReconciler reconciles a PlatformConfig object.
type PlatformConfigReconciler struct {
	client.Client
	Name         string
	Log          logr.Logger
	Controller   controller.Controller
	Events       record.EventRecorder
	FieldManager string
	Watches      []client.Object
	Phases       *phases.Registry
	Manager      manager.Manager
}

func NewPlatformConfigReconciler(mgr ctrl.Manager) *PlatformConfigReconciler {
	return &PlatformConfigReconciler{
		Name:         "PlatformConfig",
		Client:       mgr.GetClient(),
		Events:       mgr.GetEventRecorderFor("PlatformConfig-Controller"),
		FieldManager: "PlatformConfig-reconciler",
		Log:          ctrl.Log.WithName("controllers").WithName("capabilities").WithName("PlatformConfig"),
		Watches:      []client.Object{},
		Phases:       &phases.Registry{},
		Manager:      mgr,
	}
}

// +kubebuilder:rbac:groups=capabilities.platform.tbd.io,resources=platformconfigs,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=capabilities.platform.tbd.io,resources=platformconfigs/status,verbs=get;update;patch

// +kubebuilder:rbac:groups=core,resources=namespaces,verbs=list;watch

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.7.2/pkg/reconcile
func (r *PlatformConfigReconciler) Reconcile(ctx context.Context, request ctrl.Request) (ctrl.Result, error) {
	req, err := r.NewRequest(ctx, request)
	if err != nil {

		if !apierrs.IsNotFound(err) {
			return ctrl.Result{}, err
		}

		return ctrl.Result{}, nil
	}

	if err := phases.RegisterDeleteHooks(r, req); err != nil {
		return ctrl.Result{}, err
	}

	// execute the phases
	return r.Phases.HandleExecution(r, req)
}

func (r *PlatformConfigReconciler) NewRequest(ctx context.Context, request ctrl.Request) (*workload.Request, error) {
	component := &capabilitiesv1alpha1.PlatformConfig{}

	log := r.Log.WithValues(
		"kind", component.GetWorkloadGVK().Kind,
		"name", request.Name,
		"namespace", request.Namespace,
	)

	// get the component from the cluster
	if err := r.Get(ctx, request.NamespacedName, component); err != nil {
		if !apierrs.IsNotFound(err) {
			log.Error(err, "unable to fetch workload")

			return nil, fmt.Errorf("unable to fetch workload, %w", err)
		}

		return nil, err
	}

	// create the workload request
	workloadRequest := &workload.Request{
		Context:  ctx,
		Workload: component,
		Log:      log,
	}

	return workloadRequest, nil
}

// GetResources resources runs the methods to properly construct the resources in memory.
func (r *PlatformConfigReconciler) GetResources(req *workload.Request) ([]client.Object, error) {
	component, err := platformconfig.ConvertWorkload(req.Workload)
	if err != nil {
		return nil, err
	}

	return platformconfig.Generate(*component, r, req)
}

// GetEventRecorder returns the event recorder for writing kubernetes events.
func (r *PlatformConfigReconciler) GetEventRecorder() record.EventRecorder {
	return r.Events
}

// GetFieldManager returns the name of the field manager for the controller.
func (r *PlatformConfigReconciler) GetFieldManager() string {
	return r.FieldManager
}

// GetLogger returns the logger from the reconciler.
func (r *PlatformConfigReconciler) GetLogger() logr.Logger {
	return r.Log
}

// GetName returns the name of the reconciler.
func (r *PlatformConfigReconciler) GetName() string {
	return r.Name
}

// GetController returns the controller object associated with the reconciler.
func (r *PlatformConfigReconciler) GetController() controller.Controller {
	return r.Controller
}

// GetManager returns the manager object assocated with the reconciler.
func (r *PlatformConfigReconciler) GetManager() manager.Manager {
	return r.Manager
}

// GetWatches returns the objects which are current being watched by the reconciler.
func (r *PlatformConfigReconciler) GetWatches() []client.Object {
	return r.Watches
}

// SetWatch appends a watch to the list of currently watched objects.
func (r *PlatformConfigReconciler) SetWatch(watch client.Object) {
	r.Watches = append(r.Watches, watch)
}

// CheckReady will return whether a component is ready.
func (r *PlatformConfigReconciler) CheckReady(req *workload.Request) (bool, error) {
	return dependencies.PlatformConfigCheckReady(r, req)
}

// Mutate will run the mutate function for the workload.
// WARN: this will be deprecated in the future.  See apis/group/version/kind/mutate*
func (r *PlatformConfigReconciler) Mutate(
	req *workload.Request,
	object client.Object,
) ([]client.Object, bool, error) {
	return mutate.PlatformConfigMutate(r, req, object)
}

func (r *PlatformConfigReconciler) SetupWithManager(mgr ctrl.Manager) error {
	r.InitializePhases()

	baseController, err := ctrl.NewControllerManagedBy(mgr).
		WithEventFilter(predicates.WorkloadPredicates()).
		For(&capabilitiesv1alpha1.PlatformConfig{}).
		Build(r)
	if err != nil {
		return fmt.Errorf("unable to setup controller, %w", err)
	}

	r.Controller = baseController

	return nil
}

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

package constants

// this package includes the constants which include the resource names.  it is a standalone
// package to prevent import cycle errors when attempting to reference the names from other
// packages (e.g. mutate).
const (
	CRDCertmanagersCertificatesPlatformTbdIo                            = "certmanagers.certificates.platform.tbd.io"
	CRDTrustmanagersCertificatesPlatformTbdIo                           = "trustmanagers.certificates.platform.tbd.io"
	ServiceAccountNamespaceCertificatesOperatorControllerManager        = "certificates-operator-controller-manager"
	RoleNamespaceCertificatesOperatorLeaderElectionRole                 = "certificates-operator-leader-election-role"
	ClusterRoleCertificatesOperatorManagerRole                          = "certificates-operator-manager-role"
	ClusterRoleCertificatesOperatorMetricsReader                        = "certificates-operator-metrics-reader"
	ClusterRoleCertificatesOperatorProxyRole                            = "certificates-operator-proxy-role"
	RoleBindingNamespaceCertificatesOperatorLeaderElectionRolebinding   = "certificates-operator-leader-election-rolebinding"
	ClusterRoleBindingCertificatesOperatorManagerRolebinding            = "certificates-operator-manager-rolebinding"
	ClusterRoleBindingCertificatesOperatorProxyRolebinding              = "certificates-operator-proxy-rolebinding"
	ServiceNamespaceCertificatesOperatorControllerManagerMetricsService = "certificates-operator-controller-manager-metrics-service"
	DeploymentNamespaceCertificatesOperatorControllerManager            = "certificates-operator-controller-manager"
	CRDAwspodidentitywebhooksIdentityPlatformTbdIo                      = "awspodidentitywebhooks.identity.platform.tbd.io"
	ServiceAccountNamespaceIdentityOperatorControllerManager            = "identity-operator-controller-manager"
	RoleNamespaceIdentityOperatorLeaderElectionRole                     = "identity-operator-leader-election-role"
	ClusterRoleIdentityOperatorManagerRole                              = "identity-operator-manager-role"
	ClusterRoleIdentityOperatorMetricsReader                            = "identity-operator-metrics-reader"
	ClusterRoleIdentityOperatorProxyRole                                = "identity-operator-proxy-role"
	RoleBindingNamespaceIdentityOperatorLeaderElectionRolebinding       = "identity-operator-leader-election-rolebinding"
	ClusterRoleBindingIdentityOperatorManagerRolebinding                = "identity-operator-manager-rolebinding"
	ClusterRoleBindingIdentityOperatorProxyRolebinding                  = "identity-operator-proxy-rolebinding"
	ServiceNamespaceIdentityOperatorControllerManagerMetricsService     = "identity-operator-controller-manager-metrics-service"
	DeploymentNamespaceIdentityOperatorControllerManager                = "identity-operator-controller-manager"
)

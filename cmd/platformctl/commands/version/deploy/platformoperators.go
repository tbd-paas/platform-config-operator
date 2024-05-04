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

package deploy

import (
	"github.com/spf13/cobra"

	cmdversion "github.com/tbd-paas/platform-config-operator/cmd/platformctl/commands/version"

	"github.com/tbd-paas/platform-config-operator/apis/deploy"
)

// NewPlatformOperatorsSubCommand creates a new command and adds it to its
// parent command.
func NewPlatformOperatorsSubCommand(parentCommand *cobra.Command) {
	versionCmd := &cmdversion.VersionSubCommand{
		Name:         "version",
		Description:  "display the version information",
		VersionFunc:  VersionPlatformOperators,
		SubCommandOf: parentCommand,
	}

	versionCmd.Setup()
}

func VersionPlatformOperators(v *cmdversion.VersionSubCommand) error {
	apiVersions := make([]string, len(deploy.PlatformOperatorsGroupVersions()))

	for i, groupVersion := range deploy.PlatformOperatorsGroupVersions() {
		apiVersions[i] = groupVersion.Version
	}

	versionInfo := cmdversion.VersionInfo{
		CLIVersion:  cmdversion.CLIVersion,
		APIVersions: apiVersions,
	}

	return versionInfo.Display()
}

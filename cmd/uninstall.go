// Copyright © 2017 Karl Hepworth Karl.Hepworth@gmail.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/spf13/cobra"

	"github.com/gregfedorov/dvm/version"
)

// uninstallCmd represents the uninstall command
var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstall a specific version of Drush",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if flagVersion != "" {
			this := version.NewDrushVersion(flagVersion)
			this.SetVersionIdentifier(flagVersion)
			this.Uninstall()
		} else {
			RootCmd.Help()
		}
	},
}

func init() {
	RootCmd.AddCommand(uninstallCmd)
	uninstallCmd.Flags().StringVarP(&flagVersion, "version", "v", "", "Version to target, it does not have a default value.")
}

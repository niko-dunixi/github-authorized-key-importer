/*
Copyright © 2022 Paul FREAKN Baker

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/paul-nelson-baker/github-authorized-key-importer/core"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "github-authorized-key-importer",
	Short: "A brief description of your application",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		if noUsers := len(args) == 0; noUsers {
			return
		}
		allKeys := make([]core.Key, 0, 30)
		for _, username := range args {
			keys, err := core.GetKeys(username)
			if err != nil {
				panic(err)
			}
			allKeys = append(allKeys, keys...)
		}
		filename, err := homedir.Expand("~/.ssh/authorized_keys")
		if err != nil {
			panic(err)
		}
		authKeyFile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0400)
		if err != nil {
			panic(err)
		}
		defer func() {
			err := authKeyFile.Close()
			if err != nil {
				panic(err)
			}
		}()
		for _, currentKey := range allKeys {
			fmt.Fprintln(authKeyFile, currentKey.Key)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.github-authorized-key-importer.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("daemon", "d", false, "Execute in daemon mode, to periodically poll for updated keys")
}

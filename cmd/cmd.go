package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// VERSION must be incremented with each release
// see https://semver.org/
const VERSION = "0.0.0"

var rootCmd = &cobra.Command{
	Use:   "query-builder",
	Short: "query-builder is an example of mysql databases in golang",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

// Execute the cmd
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

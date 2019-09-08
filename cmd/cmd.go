package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/Demitroi/query-builders/models"
	"github.com/Demitroi/query-builders/models/gendry"
	"github.com/Demitroi/query-builders/models/goqu"
	"github.com/Demitroi/query-builders/models/ozzo-dbx"
)

// VERSION must be incremented with each release
// see https://semver.org/
const VERSION = "0.0.0"

var port int

var availableBuilders = [...]string{"gendry", "goqu", "dbx"}

var rootCmd = &cobra.Command{
	Use:   "query-builder [flags] [builder]",
	Short: "query-builder is an example of mysql databases in golang",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("builder required")
			os.Exit(1)
		}
		builder := args[0]
		// Search the builder
		var found bool
		for _, availableBuilder := range availableBuilders {
			if builder == availableBuilder {
				found = true
				break
			}
		}
		if !found {
			fmt.Printf("Builder %s not found\nAvailable builders: %v\n", builder, availableBuilders)
			os.Exit(1)
		}
		switch builder {
			case "gendry":
				models.SelectedQueryBuilder = gendry.New()
			case "goqu":
				models.SelectedQueryBuilder = goqu.New()
			case  "dbx":
				models.SelectedQueryBuilder = dbx.New()
		}
		// TODO: start http listener
	},
}

func init() {
	rootCmd.PersistentFlags().IntVarP(&port, "port", "p", 1081, "Port to serve")
}

// Execute the cmd
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

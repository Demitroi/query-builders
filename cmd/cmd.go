package cmd

import (
	"fmt"
	"os"

	"github.com/Demitroi/query-builders/models"
	"github.com/Demitroi/query-builders/models/gendry"
	"github.com/Demitroi/query-builders/models/goqu"
	dbx "github.com/Demitroi/query-builders/models/ozzo-dbx"
	"github.com/spf13/cobra"
)

// VERSION must be incremented with each release
// see https://semver.org/
const VERSION = "0.0.0"

var (
	port              int
	connectionConfig  models.ConnectionConfig
	availableBuilders = [...]string{"gendry", "goqu", "dbx"}
)

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
		// Open database connection
		err := models.OpenConnection(connectionConfig)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		// Select the builder
		switch builder {
		case "gendry":
			models.SelectedQueryBuilder = gendry.New(models.DB)
		case "goqu":
			models.SelectedQueryBuilder = goqu.New()
		case "dbx":
			models.SelectedQueryBuilder = dbx.New()
		}
		// TODO: start http listener
	},
}

func init() {
	rootCmd.PersistentFlags().IntVarP(&port, "port", "p", 1081, "Port to serve")
	rootCmd.PersistentFlags().StringVarP(&connectionConfig.User, "database-user", "", "root", "Database user")
	rootCmd.PersistentFlags().StringVarP(&connectionConfig.Password, "database-password", "", "", "Database user password")
	rootCmd.PersistentFlags().StringVarP(&connectionConfig.Protocol, "database-protocol", "", "tcp", "Database connection protocol")
	rootCmd.PersistentFlags().StringVarP(&connectionConfig.Address, "database-address", "", "localhost", "Database connection address")
	rootCmd.PersistentFlags().IntVarP(&connectionConfig.Port, "database-port", "", 3306, "Database connection port")
	rootCmd.PersistentFlags().StringVarP(&connectionConfig.DbName, "database-name", "", "query_builders", "Database name")
}

// Execute the cmd
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

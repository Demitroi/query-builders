package cmd

import (
	"fmt"
	"os"

	"github.com/Demitroi/query-builders/handlers"
	"github.com/Demitroi/query-builders/models"
	"github.com/Demitroi/query-builders/models/gendry"
	"github.com/Demitroi/query-builders/models/goqu"
	dbx "github.com/Demitroi/query-builders/models/ozzo-dbx"
	"github.com/kataras/iris/v12"
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
	Use:          "query-builder [flags] [builder]",
	Short:        "query-builder is an example of mysql databases in golang",
	Args:         cobra.MinimumNArgs(1),
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		builder := args[0]
		// Open database connection
		connection, err := models.OpenConnection(connectionConfig)
		if err != nil {
			return err
		}
		// Select the builder
		switch builder {
		case "gendry":
			handlers.QueryBuilder = gendry.New(connection)
		case "goqu":
			handlers.QueryBuilder = goqu.New(connection)
		case "dbx":
			handlers.QueryBuilder = dbx.New(connection)
		default:
			return fmt.Errorf("builder %s not found, available builders: %v", builder, availableBuilders)
		}
		// start http listener
		irisApp := iris.New()
		// register routes
		apiv1 := irisApp.Party("/api/v1")
		handlers.RegisterPersons(apiv1)
		// Iniciar el router
		addr := fmt.Sprintf(":%v", port)
		return irisApp.Run(iris.Addr(addr))
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

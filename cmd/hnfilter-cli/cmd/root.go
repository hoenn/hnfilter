package cmd

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/hoenn/hnfilter/pkg/store"

	//pg driver
	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
)

var (
	dbUser string
	dbPass string
	dbHost string
	dbPort string
	dbName string

	ds *store.DataStore
)

func init() {
	rootCmd.Flags().StringVarP(&dbUser, "dbuser", "u", "", "postgres database user")
	rootCmd.MarkFlagRequired("dbuser")
	rootCmd.Flags().StringVarP(&dbPass, "dbpass", "p", "", "postgres database password")
	rootCmd.MarkFlagRequired("dbpass")
	rootCmd.Flags().StringVarP(&dbHost, "dbhost", "z", "", "postgres database host")
	rootCmd.MarkFlagRequired("dbhost")
	rootCmd.Flags().StringVarP(&dbPort, "dbport", "P", "", "postgres database port")
	rootCmd.MarkFlagRequired("dbport")
	rootCmd.Flags().StringVarP(&dbName, "dbname", "n", "", "postgres database name")
}

var rootCmd = &cobra.Command{
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		dbConn := &store.DBConn{
			Username: dbUser,
			Password: dbPass,
			Host:     dbHost,
			Port:     dbPort,
			Name:     dbName,
		}
		db, err := sql.Open("postgres", dbConn.Format())
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
		err = db.Ping()
		if err != nil {
			log.Fatal(err)
		}

		ds = store.NewDataStore(db)

	},
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

//Execute the command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

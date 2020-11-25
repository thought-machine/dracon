package main

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/thought-machine/dracon/common/jira/config"
	"github.com/thought-machine/dracon/common/jira/jira"
	"github.com/thought-machine/dracon/pkg/enrichment/db"
	"github.com/thought-machine/dracon/producers/jira_producer/sync"
)

// TODO: make all these optionally env vars
var (
	AuthUser         string
	AuthToken        string
	JiraURL          string
	JQL              string
	ConfigPath       string
	ConnectionString string
	DryRun           bool
)

var rootCmd = &cobra.Command{
	Use:   "sync",
	Short: "sync",
	Long:  "tool to sync jira issues against a database",
	RunE: func(cmd *cobra.Command, args []string) error {
		AuthUser = viper.GetString("user")
		AuthToken = viper.GetString("token")
		JiraURL = viper.GetString("jira")
		JQL = viper.GetString("query")
		ConfigPath = viper.GetString("config")
		ConnectionString = viper.GetString("dbcon")
		DryRun = viper.GetBool("dryRun")

		var enrichDB db.EnrichDatabase
		// Parse config.yaml
		file, err := os.Open(ConfigPath)
		if err != nil {
			log.Fatalf("Could not open config file:   #%v ", err)
		}
		jiraConfig, err := config.New(file)
		if err != nil {
			log.Fatalf("Could not parse config file: %v", err)
		}
		// Setup Jira
		client := jira.NewClient(AuthUser, AuthToken, JiraURL, true, jiraConfig)
		issues, err := client.SearchByJQL(JQL)
		if err != nil {
			log.Fatal("Could not retrieve data from Jira ", err)
		}

		// allow dry run without database
		if !DryRun {
			enrichDB, err = db.NewDB(ConnectionString)
			if err != nil {
				log.Fatalf("Cannot connect to db:", err)
			}
		}
		for _, issue := range issues {
			hash := ""
			if hash, err = sync.GetHash(issue, jiraConfig); err != nil {
				log.Print("Ticket doesn't have a hash field or one has not been configured: ", err)
				hash = sync.CalculateHash(issue)
			}
			sync.UpdateDB(hash, jiraConfig, issue, DryRun, enrichDB)
		}
		return nil
	},
}

func init() {
	rootCmd.Flags().StringVar(&AuthToken, "token", "", "jira token")
	rootCmd.Flags().StringVar(&AuthUser, "user", "", "jira username")
	rootCmd.Flags().StringVar(&JiraURL, "jira", "", "the url of your jira instance")
	rootCmd.Flags().StringVar(&ConnectionString, "dbcon", "", "the enrichment db connection string")
	// TODO: the above need env vars

	rootCmd.Flags().StringVar(&JQL, "query", "", "the query to search for")
	rootCmd.Flags().StringVar(&ConfigPath, "config", "", "the path to the config file")
	rootCmd.Flags().BoolVar(&DryRun, "dryRun", false, "only print changes that would be made")

	viper.BindPFlag("token", rootCmd.Flags().Lookup("token"))
	viper.BindPFlag("user", rootCmd.Flags().Lookup("user"))

	viper.BindPFlag("jira", rootCmd.Flags().Lookup("jira"))
	viper.BindPFlag("query", rootCmd.Flags().Lookup("query"))
	viper.BindPFlag("config", rootCmd.Flags().Lookup("config"))
	viper.BindPFlag("dryRun", rootCmd.Flags().Lookup("dryRun"))
	viper.BindPFlag("dbcon", rootCmd.Flags().Lookup("dbcon"))
	viper.SetEnvPrefix("DRACON_SYNC")
	viper.AutomaticEnv()
}

// flags don't work with please, todo: try viper
func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}

}

package main

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/thought-machine/dracon/pkg/enrichment"
	"github.com/thought-machine/dracon/pkg/enrichment/db"
	v1 "github.com/thought-machine/dracon/pkg/genproto/v1"
	"github.com/thought-machine/dracon/pkg/putil"
)

var (
	connStr   string
	readPath  string
	writePath string
)

var rootCmd = &cobra.Command{
	Use:   "enricher",
	Short: "enricher",
	Long:  "tool to enrich issues against a database",
	RunE: func(cmd *cobra.Command, args []string) error {
		connStr = viper.GetString("db_connection")
		db, err := db.NewDB(connStr)
		if err != nil {
			return err
		}
		readPath = viper.GetString("read_path")
		res, err := putil.LoadToolResponse(readPath)
		writePath = viper.GetString("write_path")
		for _, r := range res {
			enrichedIssues := []*v1.EnrichedIssue{}
			for _, i := range r.GetIssues() {
				eI, err := enrichment.EnrichIssue(db, i)
				if err != nil {
					log.Println(err)
					continue
				}
				enrichedIssues = append(enrichedIssues, eI)
			}
			if err := putil.WriteEnrichedResults(r, enrichedIssues,
				filepath.Join(writePath, fmt.Sprintf("%s.enriched.pb", r.GetToolName())),
			); err != nil {
				return err
			}
			putil.WriteResults(r.GetToolName(), r.GetIssues(), filepath.Join(writePath, fmt.Sprintf("%s.raw.pb", r.GetToolName())))
		}

		return nil
	},
}

func init() {
	rootCmd.Flags().StringVar(&connStr, "db_connection", "", "the database connection DSN")
	rootCmd.Flags().StringVar(&readPath, "read_path", "", "the path to read LaunchToolResponses from")
	rootCmd.Flags().StringVar(&writePath, "write_path", "", "the path to write enriched results to")
	viper.BindPFlag("db_connection", rootCmd.Flags().Lookup("db_connection"))
	viper.BindPFlag("read_path", rootCmd.Flags().Lookup("read_path"))
	viper.BindPFlag("write_path", rootCmd.Flags().Lookup("write_path"))
	viper.SetEnvPrefix("enricher")
	viper.AutomaticEnv()
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

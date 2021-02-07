package main

import (
	"os"

	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"xorm.io/xorm"

	"github.com/octo-5/karrot-api/config"
	"github.com/octo-5/karrot-api/datastore"
	"github.com/octo-5/karrot-api/server"
	"github.com/octo-5/karrot-api/util"
)

func init() {
	cmd := &cobra.Command{
		Use:     "karrot-api",
		Version: "0.1.0",
		Run: func(cmd *cobra.Command, args []string) {
			envFile, _ := cmd.Flags().GetString("env-file")
			if envFile != "" {
				util.ReadEnvFile(envFile)
			}
		},
	}

	cmd.Flags().String("env-file", "", "the path of environments file")
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func main() {
	cfg := config.Load()

	errs := util.ValidateStruct(cfg)
	if len(errs) > 0 {
		log.WithFields(
			log.Fields{"errors": errs},
		).Fatalln("unable to start api server: invalid configurations detected")
	}

	level, err := log.ParseLevel(cfg.LogLevel)
	if err != nil {
		level = log.InfoLevel
	}

	log.SetLevel(level)

	db, err := xorm.NewEngine("postgres", cfg.DBDSN)
	if err != nil {
		log.WithError(err).Fatalln("unable to connect the datastore")
	}

	datastore.Migrate(db)
	if err != nil {
		log.WithError(err).Fatalln("falied to migrate the database")
	}

	server.RunServer(cfg, db)
}

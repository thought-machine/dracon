package db

import (
	"log"
	"os"
	"path/filepath"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	bindata "github.com/golang-migrate/migrate/source/go_bindata"
	"github.com/rakyll/statik/fs"

	// Statik bindata for migrations
	_ "github.com/thought-machine/dracon/pkg/enrichment/db/migrations"

	"github.com/jmoiron/sqlx"
)

// DB represents the db methods that are used for the enricher
type DB struct {
	*sqlx.DB
}

// NewDB returns a new DB for the enricher
func NewDB(connStr string) (*DB, error) {
	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		return nil, err
	}

	statikFS, err := fs.New()
	if err != nil {
		return nil, err
	}

	var assetNames []string
	fs.Walk(statikFS, "/", func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			assetNames = append(assetNames, info.Name())
		}
		return nil
	})

	s := bindata.Resource(assetNames,
		func(name string) ([]byte, error) {
			return fs.ReadFile(statikFS, filepath.Join("/", name))
		})

	d, err := bindata.WithInstance(s)
	m, err := migrate.NewWithInstance("go-bindata", d, "postgres", driver)
	if err != nil {
		return nil, err
	}
	log.Println(m.Version())
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return nil, err
	}

	return &DB{db}, nil
}

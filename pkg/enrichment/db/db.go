package db

import (
	"errors"
	"fmt"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	bindata "github.com/golang-migrate/migrate/source/go_bindata"
	"github.com/lib/pq"
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
	searchPath, err := getSchemaSearchPathFromConnStr(connStr)
	if err != nil {
		return nil, err
	}

	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	migrationsConfig := &postgres.Config{}
	if searchPath != "" {
		_, err := db.Exec(fmt.Sprintf(`CREATE SCHEMA IF NOT EXISTS %s`,
			pq.QuoteIdentifier(searchPath)))

		if err != nil {
			return nil, err
		}

		migrationsConfig.SchemaName = searchPath
	}

	driver, err := postgres.WithInstance(db.DB, migrationsConfig)
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

// getSchemaSearchPathFromConnStr extracts the database schema component from a
// PostgreSQL connection string; if no schema was specified, the empty string is
// returned
func getSchemaSearchPathFromConnStr(connStr string) (string, error) {
	url, err := url.Parse(connStr)

	if err == nil && url.Scheme == "postgres" {
		return getSchemaSearchPathFromURL(url)
	} else {
		return getSchemaSearchPathFromKV(connStr)
	}
}

// getSchemaSearchPathFromURL extracts the schema search path component from a
// PostgreSQL connection URL; if no search path is specified, the empty string
// is returned
func getSchemaSearchPathFromURL(connURL *url.URL) (string, error) {
	path, found := connURL.Query()["search_path"]
	if !found {
		return "", nil
	}

	if len(path) == 0 {
		return "", nil
	} else if len(path) == 1 {
		return path[0], nil
	} else {
		return "", errors.New("Multiple search_paths defined in database connection DSN")
	}
}

// getSchemaSearchPathFromKV extracts the schema search path component from a
// PostgreSQL keyword/value connection string; if no search path is specified,
// the empty string is returned
func getSchemaSearchPathFromKV(kvStr string) (string, error) {
	var path string

	for _, pair := range strings.Fields(kvStr) {
		elems := strings.SplitN(pair, "=", 2)
		if elems[0] == "search_path" {
			if path != "" {
				return "", errors.New("Multiple search_paths defined in database connection DSN")
			}

			path = elems[1]
		}
	}

	return path, nil
}

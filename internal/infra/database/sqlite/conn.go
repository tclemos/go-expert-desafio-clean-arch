package sqlite

import (
	"database/sql"
	"embed"
	"errors"
	"log"
	"os"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	_ "github.com/golang-migrate/migrate/v4/source/pkger"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"github.com/tclemos/go-expert-desafio-clean-arch/config"
)

const (
	driverName = "sqlite3"
)

//go:embed migrations/*.sql
var migrations embed.FS

func initDB(cfg config.DBConfig) error {
	dbFilePath, err := filepath.Abs(cfg.Path)
	if err != nil {
		return err
	}
	if _, err := os.Stat(dbFilePath); errors.Is(err, os.ErrNotExist) {
		f, err := os.Create(dbFilePath)
		if err != nil {
			return err
		}
		defer f.Close()
	}
	return nil
}

func initTables(db *sql.DB) error {
	migrationSourceDriver, err := iofs.New(migrations, "migrations")
	if err != nil {
		log.Fatal(err)
	}

	databaseDriver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithInstance("go-bindata", migrationSourceDriver, "sqlite3", databaseDriver)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}
	return nil
}

func MustOpenConn(cfg config.DBConfig) *sql.DB {
	err := initDB(cfg)
	if err != nil {
		panic(err)
	}

	dbFilePath, err := filepath.Abs(cfg.Path)
	if err != nil {
		panic(err)
	}

	db, err := sql.Open(driverName, dbFilePath)
	if err != nil {
		panic(err)
	}

	err = initTables(db)
	if err != nil {
		panic(err)
	}

	return db
}

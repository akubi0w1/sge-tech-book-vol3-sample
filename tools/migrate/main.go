package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/akubi0w1/sge-tech-book-vol3-sample/internal/infra/mysql"
	"github.com/akubi0w1/sge-tech-book-vol3-sample/internal/log"
	"github.com/akubi0w1/sge-tech-book-vol3-sample/internal/terror"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	mmysql "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const (
	migrateDir = "db"
)

func main() {
	// master
	err := migrateMysqlDB(&mysql.Config{
		Addr:     os.Getenv("MYSQL_MASTER_ADDR"),
		Protocol: os.Getenv("MYSQL_MASTER_PROTOCOL"),
		User:     os.Getenv("MYSQL_MASTER_USER"),
		Password: os.Getenv("MYSQL_MASTER_PASSWORD"),
		DB:       os.Getenv("MYSQL_MASTER_DB"),
	})
	if err != nil {
		log.Errorf("failed to migrate master: %v", err)
		os.Exit(1)
	}

	// user
	err = migrateMysqlDB(&mysql.Config{
		Addr:     os.Getenv("MYSQL_SHARD_ADDR"),
		Protocol: os.Getenv("MYSQL_SHARD_PROTOCOL"),
		User:     os.Getenv("MYSQL_SHARD_USER"),
		Password: os.Getenv("MYSQL_SHARD_PASSWORD"),
		DB:       os.Getenv("MYSQL_SHARD_DB"),
	})
	if err != nil {
		log.Errorf("failed to migrate user: %v", err)
		os.Exit(1)
	}
}

func migrateMysqlDB(conf *mysql.Config) error {
	// migration
	db, err := mysql.New(conf)
	if err != nil {
		return err
	}

	driver, err := mmysql.WithInstance(db, &mmysql.Config{})
	if err != nil {
		return terror.Wrapf(terror.CodeInternal, err, "failed to new instance")
	}

	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://./%s/%s", migrateDir, conf.DB),
		"mysql",
		driver,
	)
	if err != nil {
		return terror.Wrapf(terror.CodeInternal, err, "failed to open migrate")
	}

	version, _, err := m.Version()
	if err != nil {
		if !errors.Is(err, migrate.ErrNilVersion) {
			return terror.Wrapf(terror.CodeInternal, err, "failed to get migrate version")
		}
	}

	log.Infof("current version is %d", version)

	if err := m.Up(); err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			return terror.Wrapf(terror.CodeInternal, err, "failed to up")
		}
		log.Infof("schema is not change")
	}

	log.Infof("finish migrate. db=%s", conf.DB)

	return nil
}

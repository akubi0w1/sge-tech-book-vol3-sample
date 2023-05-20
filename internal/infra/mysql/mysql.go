package mysql

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/akubi0w1/sge-tech-book-vol3-sample/internal/log"
	"github.com/akubi0w1/sge-tech-book-vol3-sample/internal/terror"
)

// Config
type Config struct {
	Addr     string
	Protocol string
	User     string
	Password string
	DB       string
}

type MasterDB struct {
	*sql.DB
}

func NewMasterDB(conf *Config) (*MasterDB, error) {
	db, err := New(conf)
	if err != nil {
		return nil, err
	}

	return &MasterDB{
		DB: db,
	}, nil
}

type ShardDB struct {
	*sql.DB
}

func NewShardDB(conf *Config) (*ShardDB, error) {
	db, err := New(conf)
	if err != nil {
		return nil, err
	}

	return &ShardDB{
		DB: db,
	}, nil
}

func New(conf *Config) (*sql.DB, error) {
	driver := "mysql"
	dsn := fmt.Sprintf(
		"%s:%s@%s(%s)/%s?parseTime=true&multiStatements=true",
		conf.User, conf.Password, conf.Protocol, conf.Addr, conf.DB,
	)

	db, err := sql.Open(driver, dsn)
	if err != nil {
		return nil, terror.Wrapf(terror.CodeInternal, err, "failed to open db. dsn=%s", dsn)
	}

	if err := tryPing(db, conf.DB, 5); err != nil {
		return nil, err
	}

	log.Infof("connect mysql database. dsn=%s", dsn)

	return db, nil
}

// tryPing DBに対して　pingを送信して疎通を確認する
func tryPing(db *sql.DB, target string, retryCount int) error {
	var err error
	for i := retryCount; 0 < i; i-- {
		log.Debugf("try ping to %s db...", target)

		err = db.Ping()
		if err == nil {
			log.Debugf("success ping for %s db", target)
			return nil
		}

		time.Sleep(1 * time.Second)
		log.Debugf("failed to ping %s db instance. retry more %d times", target, i-1)
	}

	return terror.Wrapf(terror.CodeInternal, err, "failed to ping to db instance. database=%s", target)
}

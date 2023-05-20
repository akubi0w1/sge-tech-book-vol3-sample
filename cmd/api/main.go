package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/akubi0w1/sge-tech-book-vol3-sample/internal/server"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kelseyhightower/envconfig"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var conf server.Config
	err := envconfig.Process("", &conf)
	if err != nil {
		log.Fatal(err)
	}
	if err := server.ListenAndServe(&conf); err != nil {
		log.Panicf("failed to listen and serve: %+v", err)
	}
}

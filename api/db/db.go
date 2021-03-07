package db

import (
	"api/config"
	"log"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

var db neo4j.Driver

func Init() {
	c := config.GetConfig()
	driver, err := neo4j.NewDriver(c.GetString("db.url"),
		neo4j.BasicAuth(c.GetString("db.user"), c.GetString("db.pass"), ""))
	if err != nil {
		log.Fatal("error connecting to database")
		panic(err)
	}
	db = driver
}

func GetDB() neo4j.Driver {
	return db
}

package db

import (
	"api/config"
	"log"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type db struct {
	driver neo4j.Driver
}

var instance *db

func Init() {
	c := config.GetConfig()
	driver, err := neo4j.NewDriver(c.GetString("db.url"),
		neo4j.BasicAuth(c.GetString("db.user"), c.GetString("db.pass"), ""))
	if err != nil {
		log.Fatal("error connecting to database")
		panic(err)
	}

	instance = &db{driver: driver}
}

func GetDriver() neo4j.Driver {
	return instance.driver
}

func Close() {
	instance.driver.Close()
}

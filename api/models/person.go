package models

import (
	"api/db"
	"log"
	"time"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

//Person model
type Person struct {
	ID          string    `json:"id"`
	FullName    string    `json:"fullname"`
	FirstName   string    `json:"firstname"`
	MiddleNames string    `json:"middlenames"`
	LastName    string    `json:"lastname"`
	Birthday    time.Time `json:"birthday"`
}

func (p Person) GetByID(id string) (*Person, error) {
	//do db call
	driver := db.GetDriver()
	session := driver.NewSession(neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeRead,
		DatabaseName: "ancestree",
	})
	defer session.Close()

	result, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		return tx.Run(
			"MATCH (p:Person {id: $id}) RETURN p.name AS name, p.birthday AS birthday",
			map[string]interface{}{
				"id": id,
			},
		)
	})

	log.Println(result)

	return nil, err
}

func (p Person) Create() (*Person, error) {
	driver := db.GetDriver()
	session := driver.NewSession(neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeWrite,
		DatabaseName: "ancestree",
	})
	defer session.Close()

	result, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		return tx.Run(
			"CREATE (p:Person { id: apoc.create.uuid() name: $name, birthday: $birthday })",
			map[string]interface{}{
				"name":     p.FullName,
				"birthday": p.Birthday,
			},
		)
	})

	log.Println(result)

	return nil, err
}

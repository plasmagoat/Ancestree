package models

import (
	"api/db"
	"log"
	"time"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

//Person model
type Person struct {
	ID          uint64    `json:"id" gorm:"primary_key"`
	FullName    string    `json:"fullname"`
	FirstName   string    `json:"firstname"`
	MiddleNames string    `json:"middlenames"`
	LastName    string    `json:"lastname"`
	Birthday    time.Time `json:"birthday"`
}

func (p Person) GetByID(id uint64) (*Person, error) {
	//do db call
	driver := db.GetDB()
	session := driver.NewSession(neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeRead,
		DatabaseName: "familytree",
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

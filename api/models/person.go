package models

import (
	"api/db"
	"encoding/json"
	"log"
	"time"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

//Person model
type Person struct {
	ID        string    `json:"id,omitempty" swaggerignore:"true"`
	FullName  string    `json:"fullname"`
	BirthName string    `json:"birthname"`
	Birthday  time.Time `json:"birthday"`
	Deathday  time.Time `json:"deathday"`
	Gender    Gender    `json:"gender"`
}

//Gender enum
type Gender int

const (
	//Male gender
	Male Gender = iota + 1
	//Female gender
	Female
)

func (p *Person) getData() (map[string]interface{}, error) {
	obj, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	var personData map[string]interface{}
	err = json.Unmarshal(obj, &personData)
	if err != nil {
		return nil, err
	}
	return personData, nil
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
	personData, err := p.getData()
	if err != nil {
		return nil, err
	}

	session := db.GetNewSession()
	defer session.Close()

	result, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		return tx.Run(
			`CREATE (p:Person { id: apoc.create.uuid() })
			SET p += $personData
			`,
			map[string]interface{}{
				"personData": personData,
			},
		)
	})

	log.Println(result)

	return nil, err
}

func (p Person) CreateParent(childID string) (*Person, error) {
	personData, err := p.getData()
	if err != nil {
		return nil, err
	}
	session := db.GetNewSession()
	defer session.Close()

	result, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		// maybe make merge to make this function more usable
		return tx.Run(
			`MATCH (c:Person { id: $childID })
			WITH c
			OPTIONAL MATCH (c)-[:Parent]->(q:Person {gender: $gender})
			WITH q, c WHERE q IS NULL
			CREATE (c)-[:Parent]->(p:Person { id: apoc.create.uuid() })
			SET p += $personData
			RETURN p
			`,
			map[string]interface{}{
				"childID":    childID,
				"gender":     p.Gender,
				"personData": personData,
			},
		)
	})

	log.Println(json.Marshal(result))

	return nil, err
}

func (p Person) CreateChild(parentID string) (*Person, error) {
	personData, err := p.getData()
	if err != nil {
		return nil, err
	}
	session := db.GetNewSession()
	defer session.Close()

	result, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		// maybe make merge to make this function more usable
		return tx.Run(
			`MATCH (p:Person { id: $parentID })
			WITH p
			CREATE (c:Person { id: apoc.create.uuid() })-[:Parent]->(p)
			SET c += $personData
			RETURN c
			`,
			map[string]interface{}{
				"parentID":   parentID,
				"personData": personData,
			},
		)
	})

	log.Println(json.Marshal(result))

	return nil, err
}

func (p Person) CreateLink(childID string, parentID string) (*Person, error) {
	session := db.GetNewSession()
	defer session.Close()

	result, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		// maybe make merge to make this function more usable
		return tx.Run(
			`MATCH (c:Person {id: $childID}), (p:Person { id: $parentID })
			WITH c, p
			OPTIONAL MATCH (c)-[:Parent]->(q:Person {gender: p.gender})
			WITH c, p, q WHERE q IS NULL
			CREATE (c)-[:Parent]->(p)
			`,
			map[string]interface{}{
				"childID":  childID,
				"parentID": parentID,
			},
		)
	})

	log.Println(json.Marshal(result))

	return nil, err
}

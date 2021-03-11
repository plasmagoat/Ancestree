package models

import (
	"api/db"
	"encoding/json"
	"log"
	"time"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j/dbtype"
)

//Person model
type Person struct {
	ID        string    `json:"id,omitempty" swaggerignore:"true"`
	FullName  string    `json:"fullname,omitempty"`
	BirthName string    `json:"birthname,omitempty"`
	Birthday  time.Time `json:"birthday,omitempty"`
	Deathday  time.Time `json:"deathday,omitempty"`
	Gender    Gender    `json:"gender,omitempty"`
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

func createFromNode(n dbtype.Node) (Person, error) {
	person := Person{}

	jsonPerson, err := json.Marshal(n.Props)
	if err != nil {
		return person, err
	}
	err = json.Unmarshal(jsonPerson, &person)
	if err != nil {
		return person, err
	}
	return person, err
}

//GetByID gets a Person with a given ID
func (p Person) GetByID(id string) (*Person, error) {
	session := db.GetNewSession()
	defer session.Close()

	result, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		records, err := tx.Run(
			`MATCH (p:Person {id: $id}) 
			RETURN p`,
			map[string]interface{}{
				"id": id,
			},
		)
		// In face of driver native errors, make sure to return them directly.
		// Depending on the error, the driver may try to execute the function again.
		if err != nil {
			return nil, err
		}
		record, err := records.Single()
		if err != nil {
			return nil, err
		}

		person, err := createFromNode(record.Values[0].(dbtype.Node))
		return &person, nil
	})

	if err != nil {
		return nil, err
	}
	return result.(*Person), nil
}

//Create creates a Person node and returns the new Person
func (p Person) Create() (*Person, error) {
	personData, err := p.getData()
	if err != nil {
		return nil, err
	}

	session := db.GetNewSession()
	defer session.Close()

	result, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		records, err := tx.Run(
			`CREATE (p:Person { id: apoc.create.uuid() })
			SET p += $personData
			RETURN p
			`,
			map[string]interface{}{
				"personData": personData,
			},
		)
		// In face of driver native errors, make sure to return them directly.
		// Depending on the error, the driver may try to execute the function again.
		if err != nil {
			return nil, err
		}
		record, err := records.Single()
		if err != nil {
			return nil, err
		}

		person, err := createFromNode(record.Values[0].(dbtype.Node))
		return &person, nil
	})

	if err != nil {
		return nil, err
	}
	return result.(*Person), nil
}

//CreateParent creates a Person node with a parent relation to given ID and returns the new parent Person
func (p *Person) CreateParent(childID string) (*Person, error) {
	personData, err := p.getData()
	if err != nil {
		return nil, err
	}
	session := db.GetNewSession()
	defer session.Close()

	result, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		// maybe make merge to make this function more usable
		records, err := tx.Run(
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

		// In face of driver native errors, make sure to return them directly.
		// Depending on the error, the driver may try to execute the function again.
		if err != nil {
			return nil, err
		}
		record, err := records.Single()
		if err != nil {
			return nil, err
		}

		person, err := createFromNode(record.Values[0].(dbtype.Node))
		return &person, nil
	})

	if err != nil {
		return nil, err
	}
	return result.(*Person), nil
}

//CreateChild creates a Person node with a inverce parent (aka child) relation to given Person ID and returns the new child Person
func (p Person) CreateChild(parentID string) (*Person, error) {
	personData, err := p.getData()
	if err != nil {
		return nil, err
	}
	session := db.GetNewSession()
	defer session.Close()

	result, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		records, err := tx.Run(
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
		// In face of driver native errors, make sure to return them directly.
		// Depending on the error, the driver may try to execute the function again.
		if err != nil {
			return nil, err
		}
		record, err := records.Single()
		if err != nil {
			return nil, err
		}

		person, err := createFromNode(record.Values[0].(dbtype.Node))
		return &person, nil
	})

	if err != nil {
		return nil, err
	}
	return result.(*Person), nil
}

//CreateLink creates a Parent relation between a Person with childID and a Person with parentID
//what should this return?
func (p Person) CreateLink(childID string, parentID string) (*Person, error) {
	session := db.GetNewSession()
	defer session.Close()

	result, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
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

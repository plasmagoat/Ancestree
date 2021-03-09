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
	Gender      Gender    `json:"gender"`
}

type Gender int

const (
	Male Gender = iota + 1
	Female
)

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
			"CREATE (p:Person { id: apoc.create.uuid(), name: $name, birthday: $birthday })",
			map[string]interface{}{
				"name":     p.FullName,
				"birthday": p.Birthday,
			},
		)
	})

	log.Println(result)

	return nil, err
}

func (p Person) CreateParent(childID string) (*Person, error) {
	session := db.GetNewSession()
	defer session.Close()

	// first check if parent exists with the given gender...
	hasParent, err := hasParent(session, childID, p.Gender)
	if err != nil || hasParent {
		return nil, err
	}

	// ... then if not create and link
	// maybe make merge to make this function more usable

	result, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		// maybe make merge to make this function more usable
		return tx.Run(
			`MATCH (c:Person { id:$childId})
			CREATE (c)-[r:Parent]->(p:Person { 
				id: apoc.create.uuid(), 
				name: $name, 
				birthday: $birthday,
				gender: $gender,
			})
			`,
			// find a way to make this easier...
			map[string]interface{}{
				"childId":  childID,
				"name":     p.FullName,
				"birthday": p.Birthday,
				"gender":   p.Gender,
			},
		)
	})

	log.Println(result)

	return nil, err
}

func hasParent(session neo4j.Session, childID string, parentGender Gender) (bool, error) {
	result, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		records, err := tx.Run(
			`MATCH (:Person {id:$childID})-[:Parent]->(:Person {gender:$gender}) 
			WITH count(*) as count
			CALL apoc.when (count > 0, // a person can only have one dad
			"RETURN true AS bool",     // one or more parent of given gender
			"RETURN false AS bool",    // no parent with given gender
			{count:count}
			) YIELD value
			return value.bool
			`,
			map[string]interface{}{
				"childId": childID,
				"gender":  parentGender,
			},
		)
		if err != nil {
			return nil, err
		}
		record, err := records.Single()
		if err != nil {
			return nil, err
		}
		return record.Values[0].(bool), nil
	})
	if err != nil {
		return false, err
	}
	return result.(bool), nil
}

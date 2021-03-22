package models

import (
	"api/db"
	"encoding/json"
	"log"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j/dbtype"
)

//FamilyTree with Person array and Relations array
type FamilyTree struct {
	Nodes []Person   `json:"nodes"`
	Edges []Relation `json:"edges"`
}

//GetForID gets a FamilyTree for a Person with a given ID
func (p FamilyTree) GetForID(id string, depth int) (*FamilyTree, error) {
	session := db.GetNewSession()
	defer session.Close()

	result, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		records, err := tx.Run(
			`MATCH (p:Person {id: $id})
			CALL apoc.path.subgraphAll(p, {maxLevel: $depth}) YIELD nodes, relationships
			WITH p,
				nodes as nodes,
				[rel in relationships | rel {.*, type: type(rel),
					child: startNode(rel).id, 
					type: endNode(rel).gender,
					parent: endNode(rel).id}] as rels
			RETURN nodes, rels`,
			map[string]interface{}{
				"id":    id,
				"depth": depth,
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

		tree := FamilyTree{}
		nodes := record.Values[0].([]interface{})
		rels := record.Values[1].([]interface{})
		log.Println(nodes)
		log.Println(rels)

		for _, node := range nodes {
			person, _ := createFromNode(node.(dbtype.Node))
			tree.Nodes = append(tree.Nodes, person)
		}

		for _, rel := range rels {
			rrr, _ := json.Marshal(rel)
			var relation Relation
			json.Unmarshal(rrr, &relation)
			tree.Edges = append(tree.Edges, relation)
		}

		if err != nil {
			return nil, err
		}
		return &tree, nil
	})

	if err != nil {
		return nil, err
	}
	return result.(*FamilyTree), nil
}

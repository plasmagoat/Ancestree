// Create person
// input Person object
// return Person
CREATE (p:Person { 
    id: apoc.create.uuid(), 
    name: $name, 
    birthday: $birthday,
    gender: $gender
})
RETURN p

// Create parent (only one father and one mother)
// input childID and Person object
// return parent Person
MATCH (c:Person { id: $childID })
WITH c
OPTIONAL MATCH (c)-[:Parent]->(q:Person {gender: $gender})
WITH q, c WHERE q IS NULL
CREATE (c)-[:Parent]->(p:Person { id: apoc.create.uuid() })
SET p += $personData
RETURN p

// Create child
// input parentID and Person object
// return child Person 
MATCH (p:Person { id: $parentID })
WITH p
CREATE (c:Person { id: apoc.create.uuid() })-[:Parent]->(p)
SET c += $personData
RETURN c


// Create parent link
// input childID and parentID
MATCH (c:Person {id: $childID}), (p:Person { id: $parentID })
WITH c, p
OPTIONAL MATCH (c)-[:Parent]->(q:Person {gender: p.gender})
WITH c, p, q WHERE q IS NULL
CREATE (c)-[:Parent]->(p)

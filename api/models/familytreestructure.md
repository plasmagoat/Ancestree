# Defining the structure of the data 

## Requirements
* Person properties
  * Name (fullname)
  * Gender
  * Born
  * Died*
  * Image
* Relation type
  * Mother
  * Father
  * Sibling (can be inferred)
  * The rest can be inferred


## Model

### Nodes
* (p:Person)

### Edges
* [f:Father]
* [m:Mother]
* OR
* (child:Person)-[r:Parent]->(parent:Person)
* (n:Person)-[r:Partner]-(m:Person)

### Indexes..
Person id

### Constraints... 
Person can only have 1 father
Person can only have 1 mother

Constraints are mainly only for enterprise... 😟

### Initializer script
* Should import basic data
* Create Indexes
* Create Contraints

## Methods
* Get Person ✔
* Create Person ✔
* Edit Person data 💡 use merge (make sure unique index!!)
* Add Mother to Person by ID ⏰ almost
  * Can't if Person already has Parent relation to female Person
* Add Father to Person by ID ⏰ almost
  * Can't if Person already has Parent relation to male Person
* Add Child to Person by ID ⏰ almost
* Add Partner


## Queries 

Create Person
```
CREATE (p:Person { id: apoc.create.uuid(), name: $name, birthday: $birthday })
```
Find Person by ID
```
MATCH (p:Person {id: $id}) RETURN p.name AS name, p.birthday AS birthday
```

Create Link Between existing
```
MATCH (c:Person), (m:Person)
WHERE c.id = '81d97662-12ef-429c-86c8-5ae3c79fe8b7' AND m.id = '8316a732-bd89-474c-b72e-e4383f5ceaf2'
CREATE (c)-[:RELATION {type:'mother'}]->(m)
```

Find Persons parents
Find Persons siblings
Find Persons grandparents
Find Persons aunts and uncles
Find Persons cousins
Find Persons nephews and neices
Find Persons brother/sister inlaw

Find Persons full blown family tree (specify depth)

```
MATCH (n:Person)
WHERE  n.id IN ['28ec8a56-26e6-4d6b-8e20-7836abf7de70']
CALL apoc.path.subgraphAll(n, {maxLevel:2}) YIELD nodes, relationships
WITH n, 
          [node in nodes | node {.*}] as nodes,
          [rel in relationships | rel {.*, type: type(rel),
           source: startNode(rel).id, 
           target: endNode(rel).id}] as rels
RETURN nodes, rels
```

Find Persons trimmed family tree (only parents) (specify depth)
```
MATCH (n:Person)
WHERE  n.id IN ['28ec8a56-26e6-4d6b-8e20-7836abf7de70']
CALL apoc.path.subgraphAll(n, {
    relationshipFilter: "RELATION>",
    maxLevel:2
}) 
YIELD nodes, relationships
WITH n, 
          [node in nodes | node {.*}] as nodes,
          [rel in relationships | rel {.*, type: type(rel),
           source: startNode(rel).id, 
           target: endNode(rel).id}] as rels
RETURN nodes, rels

```



```
MATCH (c:Person { id:$childId})
CREATE (c)-[r:RELATION { type:'mother' }]->(p:Person { 
    id: apoc.create.uuid(), 
    name: $name, 
    birthday: $birthday 
})
```


```cypher
MATCH (n:Person)
WHERE  n.id IN ['28ec8a56-26e6-4d6b-8e20-7836abf7de70']
CALL apoc.path.subgraphAll(n, {maxLevel:1}) YIELD nodes, relationships
WITH n, 
          [node in nodes | node {.*, label:labels(node)[0]}] as nodes,
          [rel in relationships | rel {.*, type: type(rel),
           fromNode:{label:labels(startNode(rel))[0], key:startNode(rel).id}, 
           toNode:  {label:labels(endNode(rel))[0], key:endNode(rel).id}}] as rels
RETURN apoc.convert.toJson({nodes:nodes, relationships:rels})
```

```
MATCH (p:Person {id: "28ec8a56-26e6-4d6b-8e20-7836abf7de70"})
CALL apoc.path.spanningTree(p, {
	relationshipFilter: "RELATION>",
    minLevel: 1,
    maxLevel: 2
})
YIELD path
RETURN path;
```

```
MATCH (n:Person)
WHERE  n.id IN ['28ec8a56-26e6-4d6b-8e20-7836abf7de70']
CALL apoc.path.subgraphNodes(n, {
	relationshipFilter: "RELATION",
    minLevel: 1,
    maxLevel: 2
}) YIELD node
WITH n, 
          [nod in node | nod {.*, rels:relationships(nod)[0]}] as nodes
RETURN nodes
```




## JSON to the frontend
v1:
```json
{
    "person": {
        "id": "xxx",
        "name": "",
        "birthdate": "",
        "mother": {
            "id": "xxx",
            "name": "",
            "birthdate": "",
            "mother": {
                ...
            },
            "father": {
                ...
            }
        },
        "father": {
            "id": "xxx",
            "name": "",
            "birthdate": "",
            "mother": {
                ...
            },
            "father": {
                ...
            }
        }
    }
}
```

v2:
```json
{
    "xxx": {
        "id": "xxx",
        "name": "",
        "birthdate": "",
        "mother": "mid",
        "father": "fid"
    },
    "mid": {
        "id": "mid",
        "name": "",
        "birthdate": "",
        "mother": "someid",
        "father": "someotherid"
    },
    "fid": {
        "id": "fid",
        "name": "",
        "birthdate": "",
        "mother": "someid",
        "father": "someotherid"
    }
}
```

v3
```json
{
    "nodes":[
        {
            "id": "28ec8a56-26e6-4d6b-8e20-7836abf7de70",
            "birthday": "2021-02-02T00:00:00Z",
            "name": "test"
        },
        {
            "id": "8316a732-bd89-474c-b72e-e4383f5ceaf2",
            "birthday": "2021-02-02T00:00:00[UTC]",
            "name": "mom"
        }
    ],
    "rels":[
        {
            "source": "28ec8a56-26e6-4d6b-8e20-7836abf7de70",
            "type": "RELATION",
            "target": "8316a732-bd89-474c-b72e-e4383f5ceaf2"
        },{
            "source": "8316a732-bd89-474c-b72e-e4383f5ceaf2",
            "type": "RELATION",
            "target": "b6547ae5-2a26-4585-b788-c254b2c6a9d6"
        }
    ]
}
```


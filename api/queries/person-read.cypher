// EVERYTHING?
// ...

// Full Graph
MATCH (n:Person)
WHERE  n.id IN ['28ec8a56-26e6-4d6b-8e20-7836abf7de70']
CALL apoc.path.subgraphAll(n, {maxLevel:2}) YIELD nodes, relationships
WITH n, 
          [node in nodes | node {.*}] as nodes,
          [rel in relationships | rel {.*, type: type(rel),
           source: startNode(rel).id, 
           target: endNode(rel).id}] as rels
RETURN nodes, rels

// Thin Tree
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

// Person 
MATCH (p:Person {id: $id}) RETURN p.name AS name, p.birthday AS birthday
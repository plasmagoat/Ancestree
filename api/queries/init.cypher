
// ID Index
CREATE INDEX FOR (p:Person) ON (p.id);

// ID Contraint
// CREATE CONSTRAINT ON (p:Person) ASSERT p.id IS UNIQUE; // doesnt make sense... already index

CREATE (a:Person { id: apoc.create.uuid(), fullname: "Jakob Kristian Blum Samuelsen", birthday: date("2003-04-17"), gender: 1 })
CREATE (b:Person { id: apoc.create.uuid(), fullname: "Joakim Filip Blum Samuelsen", birthday: date("1997-11-28"), gender: 1 })
CREATE (c:Person { id: apoc.create.uuid(), fullname: "Elise Ida Blum Samuelsen", birthday: date("1995-08-24"), gender: 2 })
CREATE (d:Person { id: apoc.create.uuid(), fullname: "David Mikael Blum Samuelsen", birthday: date("1993-12-14"), gender: 1 })

CREATE (e:Person { id: apoc.create.uuid(), fullname: "Cecilie Morgenstjerne Lindholt", birthname: "Cecilie Marie Blum Samuelsen", birthday: date("1991-07-15"), gender: 2 })
CREATE (e:Person { id: apoc.create.uuid(), fullname: "Christian Morgenstjerne Lindholt", birthname: "Christian Lindholt Nissen", birthday: date("1990-12-17"), gender: 1 })

CREATE (m:Person { id: apoc.create.uuid(), fullname: "Aline Elisabeth Blum Samuelsen", birthname: "Aline Elisabeth Morgenstjerne", birthday: date("1966-11-24"), gender: 2 })
CREATE (f:Person { id: apoc.create.uuid(), fullname: "Peter Blum Samuelsen", birthname: "Peter Samuelsen", birthday: date("1964-07-01"), gender: 1 })

CREATE (mm:Person { id: apoc.create.uuid(), fullname: "Aline Marianne Margrethe Morgenstjerne", birthname: "Aline Marianne Margrethe Blum", birthday: date("1934-12-16"), gender: 2 })
CREATE (mf:Person { id: apoc.create.uuid(), fullname: "Heinrich Peter Procopius Morgenstjerne", birthday: date("1928-06-26"), deathday: date("2011-03-07"), gender: 1 })

CREATE (fm:Person { id: apoc.create.uuid(), fullname: "Hanne Vibeke Samuelsen", birthname: "Hanne Vibeke Plesner", birthday: date("1936-09-06"), gender: 2 })
CREATE (ff:Person { id: apoc.create.uuid(), fullname: "Per Samuelsen", birthday: date("1936-02-05"), gender: 1 })

CREATE (a)-[:Parent]->(m)
CREATE (b)-[:Parent]->(m)
CREATE (c)-[:Parent]->(m)
CREATE (d)-[:Parent]->(m)
CREATE (e)-[:Parent]->(m)

CREATE (a)-[:Parent]->(f)
CREATE (b)-[:Parent]->(f)
CREATE (c)-[:Parent]->(f)
CREATE (d)-[:Parent]->(f)
CREATE (e)-[:Parent]->(f)

CREATE (f)-[:Parent]->(fm)
CREATE (f)-[:Parent]->(ff)

CREATE (m)-[:Parent]->(mm)
CREATE (m)-[:Parent]->(mf)
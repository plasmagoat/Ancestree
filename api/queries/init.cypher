
// ID Index
CREATE INDEX FOR (p:Person) ON (p.id)

// ID Contraint
CREATE CONSTRAINT ON (p:Person)
       ASSERT p.id IS UNIQUE

CREATE (p:Person { id: "14f72cbf-ba72-4702-830a-5a3cdb197218", name: "Jakob Kristian Blum Samuelsen", birthday: date("2003-04-17"), gender: 1 })
CREATE (p:Person { id: "8c7f2afc-b009-41d8-852a-bef9a429202e", name: "Joakim Filip Blum Samuelsen", birthday: date("1997-11-28"), gender: 1 })
CREATE (p:Person { id: "3f5a37ae-18c3-4dba-9e6c-0ad4f2a33de0", name: "Elise Ida Blum Samuelsen", birthday: date("1995-08-24"), gender: 2 })
CREATE (p:Person { id: "d11857ef-349f-45d0-a006-0e3eed8e8a77", name: "David Mikael Blum Samuelsen", birthday: date("1993-12-14"), gender: 1 })
CREATE (p:Person { id: "1e4db026-b341-46c2-b430-d69fd9af1c3d", name: "Cecilie Marie Blum Samuelsen", birthday: date("1991-07-15"), gender: 2 })
CREATE (p:Person { id: "5cdcb985-1f6c-49ef-956b-384e5ab63d1a", name: "Aline Elisabeth Blum Samuelsen", birthday: date("1966-11-24"), gender: 2 })
CREATE (p:Person { id: "295f405f-04f7-458f-aa91-4d45e3cad8fb", name: "Peter Blum Samuelsen", birthday: date("1964-07-01"), gender: 1 })
CREATE (p:Person { id: "2a2927fe-93cf-461d-8194-d2d1285268a0", name: "Christian Morgenstjerne Lindholdt", birthday: date("?"), gender: 1 })


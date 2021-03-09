# Ancestree - a family 🌳

## Mission
* Replicate my 3rd Grade family tree assignment in a graph database availible on the web. 
* Ability to add more nodes to expand the family tree with my siblings, cousins, aunts & uncles and so on...
* Posibility to add new familiy trees 

# WEB
Vue.js frontend ... tbd

# BACKEND API
golang

# DATABASE
Neo4j
(Postgresql + Redis)

# TODO

* [X] Family tree database structure
* [ ] CRUD queries in the backend and CQL database queries
  * [X] CREATE (Person, Parent, Child, Link) 
  * [X] READ (Person, Full-Tree, Thin-Tree) 
  * [ ] UPDATE (Person, ?) 
  * [ ] DELETE (Person, ?) 
* [ ] frontend (using swagger for now)

# Further ideas

* Person profile
  * Name, birthdate, link to nearest relations
  * link to full family tree
* Full family tree view
  * Logged in people can edit the profile and add family members to their tree


## Possible resourses

* https://www.geni.com/
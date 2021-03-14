## D3js together with Vue3

This is tricky

Both D3js and Vue does rendering... 

I want D3's simulation, but i want Vue to render it for me. 

I managed this at first, but nodes and edges didn't seem to be reactive so the simulation was static. 

Divide up into components

* ForceSimulation component
  * Props
    * Layout - height and width and such
    * Nodes from vuex store
    * Edges from vuex store 
  * Computed properties
    * Nodes -> from simulation
    * Links -> from simulation
* Node component
  * Props
    * Node data
* Edge component
  * Props
    * Edge data

Hopefully when divided up into components the node and edge component will react to the simulation changes

Maybe its not even worth it... just let d3 render that shit
https://gist.github.com/jwickens/0c4082c9724c552efef86a45b4f6d087

try something like this.. 
https://stackoverflow.com/a/40028002

or don't and try to do the same as
https://github.com/emiliorizzo/vue-d3-network
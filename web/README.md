# Web Frontend

## Basic idea

Page split in two. One side shows graph the other side shows profile. 

Profile has information about current selected person
    Links to parents, siblings, married/has kids with, children's profile
    Profile can be edited (by famiiy member) (💡 https://editorjs.io/ maybe... always wanted to use this) (this other validation library..? 🤔)
    Information like wiki page.. 
Graph shows full familiy free for seleced person
    Option to add child/parent to current selected node
    Basic form for adding - quick add.. 

Login (with... OAuth2) and link to person

View/Browse other OAuth user's family tree

## Next Step

Wireframes
    @DannyDannyDanny save me!

Tailwindcss - https://tailwindcss.com/ 
    Its time...
    https://medium.com/@FlorianWoelki/vue-3-and-tailwindcss-2041fea3bcd2

### Current Issues

Graph 
    Currently more of just a graph and not a tree... can we get it more tree looking?
    It is impossible to see which direction the relation is going...

    Refactor Network into Vue Composition API

## Project setup
```
npm install
```

### Compiles and hot-reloads for development
```
npm run serve
```

### Compiles and minifies for production
```
npm run build
```

### Lints and fixes files
```
npm run lint
```

### Customize configuration
See [Configuration Reference](https://cli.vuejs.org/config/).

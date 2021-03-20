<template>
  <input type="text" v-model="id" placeholder="Person ID" />
  <button @click="getGraph(id)">Get It</button>
  <network
    :node-list="nodesarray"
    :link-list="links"
    node-text-key="fullname"
  />
</template>
<script>
// @ is an alias to /src
import { useStore } from 'vuex'
import { computed, ref } from 'vue'
import network from '../components/Graph/Network.vue'

export default {
  name: 'force-directed-tree',
  components: {
    network,
  },
  setup() {
    const id = ref('')
    id.value = 'bdcaa46e-554c-4c6e-addf-45b010f33f1b'
    const store = useStore()
    store.dispatch('graph/getGraph', id.value)
    const storenodes = computed(() => store.state.graph.nodes)
    const storeedges = computed(() => store.state.graph.edges)

    const links = computed(() => {
      return store.state.graph.edges.map(function(t) {
        return Object.create({ source: t.child, target: t.parent })
      })
    })

    const nodesarray = computed(() => {
      return store.state.graph.nodes.map(function(t) {
        return Object.create(t)
      })
    })
    const getGraph = id => store.dispatch('graph/getGraph', id)
    const doSomething = () => {
      console.log()
      console.log()
    }
    const layout = { height: 600, width: 1000 }

    const options = {
      force: 3000,
      nodeSize: 20,
      nodeLabels: true,
      linkWidth: 5,
    }

    return {
      id,
      nodesarray,
      links,
      storeedges,
      storenodes,
      getGraph,
      doSomething,
      layout,
      options,
      w: 1000,
      h: 600,
    }
  },
}
</script>

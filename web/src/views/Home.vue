<template>
  <div class="flex h-full">
    <network
      ref="network"
      class="bg-gray-100 dark:bg-gray-800 mx-5 rounded-xl sm:rounded-xl flex-grow h-full"
      :node-list="nodelist"
      :link-list="linklist"
      node-text-key="fullname"
      :svg-size="svgSize"
      v-on:clickNode="onSelect"
    />
    <profile
      v-if="selectedNode"
      class="mx-5 rounded-xl sm:rounded-xl flex-grow-0 w-1/3"
      :node="selectedNode"
    />
  </div>
</template>

<script>
// @ is an alias to /src
// import HelloWorld from "@/components/HelloWorld.vue";
import { computed, ref } from 'vue'
import Profile from '../components/Profile/Profile.vue'
import Network from '../components/Graph/Network.vue'
import { useStore } from 'vuex'
export default {
  name: 'home',
  components: {
    Profile,
    Network,
    // HelloWorld,
  },
  setup() {
    const store = useStore()

    // Current User
    const id = ref('')

    // This is temporary...
    id.value = 'bdcaa46e-554c-4c6e-addf-45b010f33f1b'
    store.dispatch('graph/getGraph', id.value)

    // Open/Close menu.. doesn't go here..
    const open = ref(true)

    const svgSize = ref({
      width: '100%',
      height: '100%',
    })

    const linklist = computed(() => {
      return store.state.graph.edges.map(function(t) {
        return Object.create({ source: t.child, target: t.parent })
      })
    })
    const nodelist = computed(() => {
      return store.state.graph.nodes.map(function(t) {
        return Object.create(t)
      })
    })

    const selectedNode = computed(() => store.getters['graph/selectedNode'])

    const onSelect = (e, target) => {
      store.commit('graph/setSelectedNode', target.id)
    }

    const toggleMenu = () => {
      open.value = !open.value
    }
    return {
      open,
      nodelist,
      linklist,
      svgSize,
      toggleMenu,
      onSelect,
      selectedNode,
    }
  },
}
</script>

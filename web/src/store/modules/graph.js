import tree from '../../api/tree'

// initial state
const state = () => ({
  nodes: [],
  edges: [],
  selectedNode: null,
})

// getters
const getters = {
  selectedNode: state => {
    return state.nodes.find(n => n.id == state.selectedNode)
  },
}

// actions
const actions = {
  getGraph({ commit }, id) {
    tree.getGraph(id, graph => {
      commit('setGraph', graph)
      commit('setSelectedNode', id)
    })
  },
}

// mutations
const mutations = {
  setGraph(state, graph) {
    state.nodes = graph.nodes
    state.edges = graph.edges
  },
  setSelectedNode(state, id) {
    state.selectedNode = id
  },
}

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations,
}

import tree from "../../api/tree";

// initial state
const state = () => ({
  nodes: [],
  edges: []
});

// getters
const getters = {};

// actions
const actions = {
  getGraph({ commit }, id) {
    tree.getGraph(id, graph => {
      commit("setGraph", graph);
    });
  }
};

// mutations
const mutations = {
  setGraph(state, graph) {
    state.nodes = graph.nodes;
    state.edges = graph.edges;
  }

  // decrementProductInventory(state, { id }) {
  //     const product = state.all.find(product => product.id === id)
  //     product.inventory--
  // }
};

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations
};

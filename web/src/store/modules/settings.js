// initial state
const state = () => ({
  darkmode: false,
})

// getters
const getters = {}

// actions
const actions = {}

// mutations
const mutations = {
  toggleDarkmode(state) {
    state.darkmode = !state.darkmode
  },
}

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations,
}

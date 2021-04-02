import auth from '../../api/auth'

// initial state
const user = JSON.parse(localStorage.getItem('user'))
const initialUserState = user
  ? { status: { loggedIn: true }, user }
  : { status: { loggedIn: false }, user: null }
const state = () => ({
  ...initialUserState,
})

// getters
const getters = {}

// actions
const actions = {
  login({ commit }, user) {
    return auth.login(user).then(
      user => {
        commit('loginSuccess', user)
        return Promise.resolve(user)
      },
      error => {
        commit('loginFailure')
        return Promise.reject(error)
      },
    )
  },
  logout({ commit }) {
    auth.logout()
    commit('logout')
  },
  register({ commit }, user) {
    return auth.register(user).then(
      response => {
        commit('registerSuccess')
        return Promise.resolve(response.data)
      },
      error => {
        commit('registerFailure')
        return Promise.reject(error)
      },
    )
  },
}

// mutations
const mutations = {
  loginSuccess(state, user) {
    state.status.loggedIn = true
    state.user = user
  },
  loginFailure(state) {
    state.status.loggedIn = false
    state.user = null
  },
  logout(state) {
    state.status.loggedIn = false
    state.user = null
  },
  registerSuccess(state) {
    state.status.loggedIn = false
  },
  registerFailure(state) {
    state.status.loggedIn = false
  },
}

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations,
}

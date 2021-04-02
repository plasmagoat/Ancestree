import { createStore, createLogger } from 'vuex'
import auth from './modules/auth'
import graph from './modules/graph'
import settings from './modules/settings'

const debug = process.env.NODE_ENV !== 'production'

export const store = createStore({
  modules: {
    auth,
    graph,
    settings,
  },
  strict: debug,
  plugins: debug ? [createLogger()] : [],
})

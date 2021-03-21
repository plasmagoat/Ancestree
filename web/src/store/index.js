import { createStore, createLogger } from 'vuex'
//import profile from './modules/profile'
import graph from './modules/graph'
import settings from './modules/settings'

const debug = process.env.NODE_ENV !== 'production'

export const store = createStore({
  modules: {
    //profile,
    graph,
    settings,
  },
  strict: debug,
  plugins: debug ? [createLogger()] : [],
})

import { createStore, createLogger } from "vuex";
//import profile from './modules/profile'
import graph from "./modules/graph";

const debug = process.env.NODE_ENV !== "production";

export const store = createStore({
  modules: {
    //profile,
    graph
  },
  strict: debug,
  plugins: debug ? [createLogger()] : []
});

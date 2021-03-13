import { PersistentStore } from "./main";
import { TREE_STORE_NAME } from "./store-names";
import tree from "../api/tree";

class TreeStore extends PersistentStore {
  data() {
    return {
      nodes: [],
      edges: []
    };
  }
  getGraph(id) {
    tree.getGraph(id, tree => {
      this.state.nodes = tree.nodes;
      this.state.edges = tree.edges;
    });
  }
}
export const treeStore = new TreeStore(TREE_STORE_NAME);

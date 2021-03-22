//import json from "@/assets/test.json";

export default {
  getGraph(id, cb) {
    // const root = json.tree;
    // cb(root)
    fetch(`http://localhost:9000/api/tree/${id}`)
      .then(response => response.json())
      .then(data => cb(data.tree))
  },
}

//import json from "@/assets/test.json";

export default {
  getGraph(id, cb) {
    // const root = json.tree;
    // cb(root)
    fetch(`https://ancestree-api.procopius.dk/api/tree/${id}`)
      .then(response => response.json())
      .then(data => cb(data.tree))
  },
}

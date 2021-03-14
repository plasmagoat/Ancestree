<template>
  <svg id="forcedDirectedTree">
    <g id="edgecontainer" stroke="#999" stroke-opacity="1"></g>
    <g id="nodecontainer" fill="#fff" stroke="#000" stroke-width="1.5"></g>
  </svg>
</template>
<script>
// @ is an alias to /src
import { mapState } from "vuex";
import { onBeforeMount } from "vue";
import { treeStore } from "../store/tree-store";
import {
  forceSimulation,
  forceLink,
  forceManyBody,
  forceX,
  forceY,
} from "d3-force";
import { select } from "d3-selection";
import { drag } from "d3-drag";

// import { hierarchy } from "d3-hierarchy";
// import json from "@/assets/flare-2.json";
// const root = hierarchy(json);
// //nodes
// const nodes = root.descendants();
// //edges
// const edges = root.links();

import json from "@/assets/test.json";

const root = json.tree;
//nodes
const nodes = root.nodes.map(function(t) {
  return Object.create(t);
});
//edges
const edges = root.edges.map(function(t) {
  return Object.create(t);
});

function dragger(simulation) {
  function dragstarted(event, d) {
    if (!event.active) simulation.alphaTarget(0.3).restart();
    d.fx = d.x;
    d.fy = d.y;
    console.log(d.fullname);
  }

  function dragged(event, d) {
    d.fx = event.x;
    d.fy = event.y;
  }

  function dragended(event, d) {
    if (!event.active) simulation.alphaTarget(0);
    d.fx = null;
    d.fy = null;
  }

  return drag()
    .on("start", dragstarted)
    .on("drag", dragged)
    .on("end", dragended);
}

export default {
  name: "ForceDirectedTree",
  setup() {
    onBeforeMount(async () => await treeStore.init());
    treeStore.getGraph("bdcaa46e-554c-4c6e-addf-45b010f33f1b");

    return {
      treeState: treeStore.getState(),
      isInitialized: treeStore.getIsInitialized(),
    };
  },
  computed: {
    ...mapState("graph", {
      nodesstate: (state) => state.nodes,
      edgesstate: (state) => state.edges,
    }),
  },
  data() {
    return {
      w: 1000,
      h: 600,
    };
  },
  mounted() {
    this.generateGraph();
  },
  watch: {
    nodesstate() {
      this.generateGraph();
    },
  },
  created() {
    this.$store.dispatch(
      "graph/getGraph",
      "bdcaa46e-554c-4c6e-addf-45b010f33f1b"
    );
  },
  methods: {
    generateGraph() {
      //initialize simulation
      const edgeobjects = this.edgesstate.map(function(t) {
        return Object.create({ source: t.child, target: t.parent });
      });
      const nodeObjects = this.nodesstate.map(function(t) {
        return Object.create(t);
      });

      const simulation = forceSimulation(nodeObjects)
        .force(
          "link",
          forceLink(edgeobjects)
            .id((d) => d.id)
            .distance(10)
            .strength(0.2)
        )
        .force("charge", forceManyBody().strength(-50))
        .force("x", forceX())
        .force("y", forceY());

      const svg = select("#forcedDirectedTree")
        .attr("viewBox", [-this.w / 2, -this.h / 2, this.w, this.h])
        .attr("width", this.w)
        .attr("height", this.h);
      //const svg = create("svg")

      const link = select("#edgecontainer")
        .attr("stroke", "#999")
        .attr("stroke-opacity", 0.4)
        .selectAll("line")
        .data(edgeobjects)
        .join("line");

      const node = select("#nodecontainer")
        .attr("fill", "#fff")
        .attr("stroke", "#000")
        .attr("stroke-width", 1.5)
        .selectAll("circle")
        .data(nodeObjects)
        .join("circle")
        .attr("fill", (d) => (d.gender === 1 ? null : "#000"))
        .attr("stroke", (d) => (d.gender === 1 ? null : "#fff"))
        .attr("r", 7)
        .call(dragger(simulation));

      // this comes in double...
      node.append("title").text((d) => d.fullname);

      simulation.on("tick", () => {
        link
          .attr("x1", (d) => d.source.x)
          .attr("y1", (d) => d.source.y)
          .attr("x2", (d) => d.target.x)
          .attr("y2", (d) => d.target.y);

        node.attr("cx", (d) => d.x).attr("cy", (d) => d.y);
      });
    },
    updateGraph() {
      //create whay of updating current simulation based on vue data..
    },
  },
};
</script>

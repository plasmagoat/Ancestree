<template>
  <div id="forcedDirectedTree"></div>

  <svg
    :viewBox="[-this.w / 2, -this.h / 2, this.w, this.h]"
    :width="this.w"
    :height="this.h"
  >
    <g stroke="#999" stroke-opacity="0.4">
      <line :key="e" v-for="e in nodesstate">
        {{ e.child }}
      </line>
      test
    </g>
    <g fill="#fff" stroke="#000" stroke-width="1.5">
      <circle
        v-for="n in treeState.nodes"
        v-on:click="alert"
        :key="n.id"
        :r="5"
        :cx="n.x"
        :cy="n.y"
        fill="#fff"
        stroke="#000"
      >
        {{ n.id }}
      </circle>
    </g>
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
    const alert = () => {
                window.alert("hey")
            }
    return {
      treeState: treeStore.getState(),
      isInitialized: treeStore.getIsInitialized(),
      alert
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
  created() {
    this.$store.dispatch(
      "graph/getGraph",
      "bdcaa46e-554c-4c6e-addf-45b010f33f1b"
    );
  },
  mounted() {
    this.generateGraph();
  },
  methods: {
    generateGraph() {
      const simulation = forceSimulation(nodes)
        .force(
          "link",
          forceLink(edges)
            .id((d) => d.id)
            .distance(10)
            .strength(0.2)
        )
        .force("charge", forceManyBody().strength(-50))
        .force("x", forceX())
        .force("y", forceY());

      const svg = select("#forcedDirectedTree")
        .append("svg")
        .attr("viewBox", [-this.w / 2, -this.h / 2, this.w, this.h])
        .attr("width", this.w)
        .attr("height", this.h);
      //const svg = create("svg")

      const link = svg
        .append("g")
        .attr("stroke", "#999")
        .attr("stroke-opacity", 0.4)
        .selectAll("line")
        .data(edges)
        .join("line");

      const node = svg
        .append("g")
        .attr("fill", "#fff")
        .attr("stroke", "#000")
        .attr("stroke-width", 1.5)
        .selectAll("circle")
        .data(nodes)
        .join("circle")
        .attr("fill", (d) => (d.gender === 1 ? null : "#000"))
        .attr("stroke", (d) => (d.gender === 1 ? null : "#fff"))
        .attr("r", 7)
        .call(dragger(simulation));

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
  },
};
</script>

<template>
  <input type="text" v-model="id" placeholder="Person ID" />
  <button>{{ id }}</button>
  <button @click="getGraph(id)">Get It</button>
  <button @click="doSomething()">DO IT</button>
  <!-- <div id="forcedDirectedTree">{{ treedata.nodes[0].x  }}</div> -->

  <!-- <svg
    v-if="simulation"
    :viewBox="[-this.w / 2, -this.h / 2, this.w, this.h]"
    :width="this.w"
    :height="this.h"
  >
    <g id="edgecontainer" stroke="#999" stroke-opacity="1">
      <line
        v-for="e in treedata.edges"
        :key="e"
        v-bind:x1="e.source.x"
        v-bind:y1="e.source.y"
        v-bind:x2="e.target.x"
        v-bind:y2="e.target.y"
      >
        {{ e.child }}
      </line>
    </g>
    <g id="nodecontainer" fill="#fff" stroke="#000" stroke-width="1.5">
      <circle
        v-for="n in treedata.nodes"
        :key="n.x"
        :r="5"
        :cx="n.x"
        v-bind:cy="n.y"
        fill="#fff"
        stroke="#000"
      >
        {{ n.fullname }}
      </circle>
    </g>
  </svg> -->
  <d3ForceSimulation
    :layout="layout"
    :nodes="nodesarray"
    :edges="links"
  ></d3ForceSimulation>
  <d3-network :net-nodes="nodesarray" :net-links="links" :options="options" />
</template>
<script>
// @ is an alias to /src
import { mapState, useStore } from "vuex";
import { computed, onBeforeMount, reactive, ref, watch } from "vue";
import d3ForceSimulation from "./d3js/d3ForceSimulation.vue";
import D3Network from "vue-d3-network";

import { treeStore } from "../store/tree-store";
import {
  forceSimulation,
  forceLink,
  forceManyBody,
  forceX,
  forceY,
} from "d3-force";
import { select, selectAll } from "d3-selection";
import { drag } from "d3-drag";

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
  components: {
    d3ForceSimulation,
    D3Network,
  },
  setup() {
    const id = ref("");
    id.value = "bdcaa46e-554c-4c6e-addf-45b010f33f1b";
    const store = useStore();
    store.dispatch("graph/getGraph", id.value);
    const storenodes = computed(() => store.state.graph.nodes);
    const storeedges = computed(() => store.state.graph.edges);

    const links = computed(() => {
      return store.state.graph.edges.map(function(t) {
        return { sid: t.child, tid: t.parent };
      });
    });

    const nodesarray = computed(() => {
      return store.state.graph.nodes.map(function(t) {
        return t;
      });
    });
    const simulation = computed(() => {
      const edgeobjects = store.state.graph.edges.map(function(t) {
        return Object.create({ source: t.child, target: t.parent });
      });
      const nodeObjects = store.state.graph.nodes.map(function(t) {
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

      return simulation;
    });
    const nodes = computed(() => simulation.value.nodes());
    const edges = computed(() => simulation.value.force("link").links());

    const treedata = reactive({
      simulation: simulation,
      nodes: computed(() => treedata.simulation.nodes()),
      edges: computed(() => treedata.simulation.force("link").links()),
      wordCount: computed(() => id.value.length),
    });
    const getGraph = (id) => store.dispatch("graph/getGraph", id);

    simulation.value.restart();
    const doSomething = () => {
      treedata.nodes.splice(0, 1, treedata.nodes[0]);
      console.log(treedata.nodes[0]);
      console.log(treedata.nodes[0].fullname);
      simulation.value.stop();
      simulation.value.restart();
      simulation.value.alpha(0.5);
    };

    const layout = { height: 600, width: 1000 };

    const options = {
      force: 3000,
      nodeSize: 20,
      nodeLabels: true,
      linkWidth: 5,
    };

    return {
      id,
      nodesarray,
      links,
      storeedges,
      storenodes,
      treedata,
      getGraph,
      doSomething,
      simulation,
      layout,
      options,
      w: 1000,
      h: 600,
    };
  },
};
</script>

<template>
  <!-- replace with viewbox prop -->
  <svg :viewBox="viewBox" :width="layout.width" :height="layout.height">
    <g id="edgecontainer" stroke="#999" stroke-opacity="1"></g>
    <g id="nodecontainer" fill="#fff" stroke="#000" stroke-width="1.5">
      <d-3-node v-for="n in simnodes" :key="n.x" :node="n"></d-3-node>
    </g>
  </svg>
</template>
<script>
import {
  computed,
  onBeforeMount,
  onBeforeUpdate,
  onUpdated,
  reactive,
  ref,
  watch,
} from "vue";
import {
  forceSimulation,
  forceLink,
  forceManyBody,
  forceX,
  forceY,
} from "d3-force";
import { select, selectAll } from "d3-selection";
import { drag } from "d3-drag";
import d3Node from "./d3Node.vue";

export default {
  name: "d3ForceSimulation",
  components: {
    d3Node,
  },
  props: {
    nodes: Array,
    edges: Array,
    layout: Object,
  },
  setup(props) {
    const simulation = forceSimulation(props.nodes)
      .force(
        "link",
        forceLink(props.edges)
          .id((d) => d.id)
          .distance(10)
          .strength(0.2)
      )
      .force("charge", forceManyBody().strength(-50))
      .force("x", forceX())
      .force("y", forceY());

    const simnodes = computed(() => simulation.nodes());
    const links = computed(() => simulation.force("link").links());

    // var start = new Date();
    // var time = 0;
    // var ticks = 0;
    // var renderStart = new Date();
    // onBeforeUpdate(() => {
    //   renderStart = new Date();
    //   simulation.stop();
    // });
    // onUpdated(() => {
    //   time += new Date() - renderStart;
    //   if (ticks > 298) {
    //     var totalTime = new Date() - start;
    //     console.log("Total Time:", totalTime);
    //     console.log("Render Time:", time);
    //     console.log("Ticks:", ticks);
    //     console.log("Average Time:", totalTime / ticks);
    //   } else {
    //     simulation.tick();
    //   }
    // });

    const viewBox = computed(() => {
      return [
        -props.layout.width / 2,
        -props.layout.height / 2,
        props.layout.width,
        props.layout.height,
      ];
    });

    return {
      viewBox,
      simnodes,
      links,
      simulation
    };
  },
};
</script>

<template>
  <line
    v-bind:x1="e.source.x"
    v-bind:y1="e.source.y"
    v-bind:x2="e.target.x"
    v-bind:y2="e.target.y"
  >
    {{ e.child }}
  </line>
</template>
<script>
export default {
  name: "d3Link",
  props: {
    link: Object,
  },
  setup(props) {
    return {};
  },
};
</script>

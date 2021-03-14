<template>
  <circle
    r="5"
    :cx="node.x"
    :cy="node.y"
    fill="#fff"
    stroke="#000"
  >
    {{ node.fullname }}
  </circle>
</template>
<script>
export default {
  name: "d3Node",
  props: {
    node: Object,
  },
  setup(props) {
    return;
  },
};
</script>

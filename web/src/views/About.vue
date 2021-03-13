<template>
  <div class="about">
    <h1>This is an about page</h1>
  </div>
  <force-directed-tree></force-directed-tree>
</template>
<script>
// @ is an alias to /src
import ForceDirectedTree from "@/components/ForceDirectedTree.vue";

export default {
  name: "About",
  components: {
    ForceDirectedTree
  }
};
</script>

<template>
  <button
    class="bg-purple-600 hover:bg-purple-700 focus:outline-none focus:ring-2 focus:ring-purple-600 focus:ring-opacity-50"
  >
    {{ text }}
  </button>
</template>
<script>
export default {
  props: {
    text: String,
  },
  setup() {},
}
</script>

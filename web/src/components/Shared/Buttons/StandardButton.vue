<template>
  <button
    class="btn bg-green-700 hover:bg-green-600 focus:ring-green-600 font-bold text-gray-100 shadow-lg"
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
<style scoped>
.btn {
  @apply focus:outline-none focus:ring-2  focus:ring-opacity-50 w-auto rounded-2xl p-2;
}
</style>

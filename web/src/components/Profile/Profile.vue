<template>
  <div
    class="bg-white dark:bg-gray-800 rounded-xl sm:rounded-xl p-4 pb-6 sm:p-8 lg:p-4 lg:pb-6 xl:p-8 space-y-6 sm:space-y-8 lg:space-y-6 xl:space-y-8"
  >
    <h1>{{ node.fullname }}</h1>
    <h2>{{ node.birthday }}</h2>
    <standard-button text="Add Parent" />
    <standard-button text="Add Child" />
    <standard-button text="Edit" />
  </div>
</template>
<script>
import StandardButton from '../Shared/Buttons/StandardButton.vue'
export default {
  name: 'profile',
  components: { StandardButton },
  props: {
    node: Object,
  },
  setup() {
    return {}
  },
}
</script>

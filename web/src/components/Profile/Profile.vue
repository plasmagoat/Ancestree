<template>
  <div class="py-6 flex flex-col justify-start">
    <div class="relative p-3">
      <div
        class="absolute inset-0 bg-gradient-to-r from-green-800 to-green-600 shadow-lg transform skew-y-0 -rotate-6 rounded-3xl"
      />
      <div
        class="relative px-4 py-10 bg-gray-100 dark:bg-gray-700 shadow-lg rounded-3xl text"
      >
        <h1 class="text-lg font-bold">
          {{ node.fullname }}
        </h1>
        <h2>
          {{
            new Date(Date.parse(node.birthday)).toLocaleDateString('en-US', {
              weekday: 'long',
              year: 'numeric',
              month: 'long',
              day: 'numeric',
            })
          }}
        </h2>
      </div>
    </div>
    <div class="p-3 flex flex-row mt-5">
      <standard-button class="flex-1" text="Add Parent" />
      <standard-button class="flex-1 mx-2" text="Add Child" />
      <standard-button class="flex-1" text="Edit" />
    </div>
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
<style>
.text {
  @apply dark:text-gray-100;
}
</style>

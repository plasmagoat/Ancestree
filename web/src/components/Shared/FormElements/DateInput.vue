<template>
  <label class="block">
    <span class="text-gray-700">{{ text }}</span>
    <input
      type="date"
      v-model="datelocal"
      class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-300 focus:ring focus:ring-indigo-200 focus:ring-opacity-50"
    />
  </label>
</template>
<script>
import { computed } from 'vue'
export default {
  name: 'date-selector',
  props: {
    date: Number,
    text: String,
  },
  setup(props, { emit }) {
    const datelocal = computed({
      get() {
        // from ISO Format
        return props.date ? props.date.split('T')[0] : ''
      },
      set(newvalue) {
        // To ISO Format
        emit('update:date', new Date(Date.parse(newvalue)).toISOString())
      },
    })
    return {
      datelocal,
    }
  },
}
</script>

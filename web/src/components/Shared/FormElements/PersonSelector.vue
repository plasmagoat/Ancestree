<template>
  <select
    class="block w-full mt-1 rounded-md border-gray-300 shadow-sm focus:border-indigo-300 focus:ring focus:ring-indigo-200 focus:ring-opacity-50"
    v-model="selectedLocal"
  >
    <option :key="person.id" v-for="person in nodes" :value="person.id">
      {{ person.fullname }}
    </option>
  </select>
</template>
<script>
import { computed } from 'vue'
import { useStore } from 'vuex'
export default {
  name: 'person-selector',
  props: {
    selected: Object,
  },
  setup(props, { emit }) {
    const store = useStore()
    const nodes = computed(() => store.state.graph.nodes)

    const selectedLocal = computed({
      get() {
        return props.selected
      },
      set(newvalue) {
        // Radio buttons values can't map to numbers
        emit('update:selected', newvalue)
      },
    })

    return {
      nodes,
      selectedLocal,
    }
  },
}
</script>

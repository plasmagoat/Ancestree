<template>
  <div class="flex justify-between items-center" @click="toggleDarkmode()">
    <div
      class="w-16 h-10 flex items-center bg-gray-300 rounded-full p-1 duration-300 ease-in-out"
      :class="{ 'bg-gray-700': darkmode }"
    >
      <div
        class="bg-white w-8 h-8 rounded-full shadow-md transform duration-300 ease-in-out"
        :class="{ 'translate-x-6': darkmode }"
      />
    </div>
  </div>
</template>
<script>
import { computed } from 'vue'
import { useStore } from 'vuex'
export default {
  name: 'dark-mode-switcher',
  setup() {
    const store = useStore()
    const darkmode = computed(() => store.state.settings.darkmode)
    const toggleDarkmode = () => {
      store.commit('settings/toggleDarkmode')
    }

    return {
      toggleDarkmode,
      darkmode,
    }
  },
}
</script>

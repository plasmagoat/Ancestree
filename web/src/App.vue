<template>
  <div :class="darkmode ? 'dark' : ''">
    <div class="h-screen bg-grey-100 flex flex-col dark:bg-gray-800">
      <header
        class="flex flex-wrap flex-row rounded-xl justify-between items-center md:space-x-4 bg-gradient-to-r from-green-800 to-green-600 py-1 px-4 relative mx-5 my-5"
      >
        <a href="#" class="block">
          <span class="sr-only">themes.dev</span>
          <img
            class="h-8"
            src="./assets/logo.png"
            alt="Acestree Logo"
            title="Themes.dev Logo"
          />
        </a>
        <nav
          class="flex flex-row space-x-6 font-semibold w-auto bg-transparent p-0 "
        >
          <router-link class="nav-link" to="/">
            Home
          </router-link>
          <router-link class="nav-link" to="/about">
            About
          </router-link>
          <router-link class="nav-link" to="/settings">
            Settings
          </router-link>
          <dark-mode-switcher />
        </nav>
      </header>
      <router-view class="mb-5" />
    </div>
  </div>
</template>
<script>
import { computed } from 'vue'
import { useStore } from 'vuex'
import DarkModeSwitcher from './components/Shared/Menu/DarkModeSwitcher.vue'
export default {
  components: {
    DarkModeSwitcher,
  },
  setup() {
    const store = useStore()

    const darkmode = computed(() => store.state.settings.darkmode)

    return { darkmode }
  },
}
</script>
<style>
.nav-link {
  @apply text-gray-100 hover:underline m-auto;
}
</style>

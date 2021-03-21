<template>
  <div
    :class="darkmode ? 'dark' : ''"
    class="h-screen bg-green-700 flex flex-col"
  >
    <header
      class="flex flex-wrap flex-row rounded-xl justify-between items-center md:space-x-4 bg-gray-100 py-1 px-4 relative mx-5 my-5"
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
        class="flex flex-col md:flex-row md:space-x-6 font-semibold w-full md:w-auto bg-gray-100 md:bg-transparent p-6 md:p-0 justify-evenly"
      >
        <dark-mode-switcher />
        <router-link class="text-green-600 hover:underline" to="/">
          Home
        </router-link>
        <router-link class="text-green-600 hover:underline" to="/about">
          About
        </router-link>
      </nav>
    </header>
    <router-view class="mb-5" />
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
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

#nav {
  padding: 30px;
}

#nav a {
  font-weight: bold;
  color: #2c3e50;
}

#nav a.router-link-exact-active {
  color: #42b983;
}
</style>

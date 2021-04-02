<template>
  <div :class="darkmode ? 'dark' : ''">
    <div class="h-screen bg-grey-100 flex flex-col dark:bg-gray-800">
      <header
        class="flex flex-wrap flex-row rounded-xl justify-between items-center md:space-x-4 bg-gradient-to-r from-green-800 to-green-600 py-1 px-4 relative mx-5 my-5 z-10"
      >
        <a href="#" class="flex flex-row">
          <span class="sr-only">ancestree</span>
          <img
            class="h-8 mr-2"
            src="./assets/logo.png"
            alt="Ancestree Logo"
            title="Ancestree Logo"
          />
          <span class="text-gray-100 m-auto font-semibold font-mono text-xl">
            ANCESTREE
          </span>
        </a>

        <nav class="flex flex-row space-x-6 font-semibold w-auto p-0">
          <span v-if="currentUser" class="text-gray-100 m-auto">
            ID: {{ currentUser.id }}
          </span>
          <router-link class="nav-link" to="/">
            Home
          </router-link>
          <router-link class="nav-link" to="/about">
            About
          </router-link>
          <router-link v-show="!currentUser" class="nav-link" to="/login">
            Login
          </router-link>
          <router-link
            v-show="currentUser"
            class="nav-link"
            to="/"
            @click="logout"
          >
            Logout
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
    const currentUser = computed(() => store.state.auth.user)
    const logout = () => {
      store.dispatch('auth/logout')
    }

    return { darkmode, logout, currentUser }
  },
}
</script>
<style>
.nav-link {
  @apply text-gray-100 hover:underline m-auto;
}
</style>

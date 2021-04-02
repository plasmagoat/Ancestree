<template>
  <div class="flex h-full ">
    <div class="py-3 flex-grow flex-col justify-start mx-5 h-full">
      <div class="relative p-3 h-full">
        <div
          class="absolute m-2 inset-0 bg-gradient-to-r from-green-800 to-green-600 shadow-lg transform skew-y-0 -rotate-1 rounded-3xl"
        />
        <div
          class="relative h-full px-4 py-10 bg-gray-100 dark:bg-gray-700 shadow-lg rounded-3xl border-green-800 border-2"
        >
          <h1 class="text-lg font-bold">
            Login
          </h1>
          <h2>
            something
          </h2>
          <input type="text" v-model="user.id" />
          <standard-button text="login" @click="handleLogin" />
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import { ref } from 'vue'
import User from '@/models/user'
import { useStore } from 'vuex'
import { useRouter } from 'vue-router'
import StandardButton from '@/components/Shared/FormElements/Buttons/StandardButton.vue'
// @ is an alias to /src

export default {
  name: 'login',
  components: { StandardButton },
  setup() {
    const store = useStore()
    const router = useRouter()
    const user = ref(new User('0adcef56-c7f3-4f1a-a585-1c67a2282e88', ''))
    //const user = ref(new User('bdcaa46e-554c-4c6e-addf-45b010f33f1b', ''))
    const handleLogin = () => {
      store.dispatch('auth/login', user.value)
      router.push({ name: 'Home' })
    }
    return {
      handleLogin,
      user,
    }
  },
}
</script>

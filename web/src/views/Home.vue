<template>
  <div class="" />
</template>

<script>
// @ is an alias to /src
// import HelloWorld from "@/components/HelloWorld.vue";
import { ref } from 'vue'
export default {
  name: 'home',
  components: {
    // HelloWorld,
  },
  setup() {
    const open = ref(true)
    const toggleMenu = () => {
      open.value = !open.value
      //setTimeout(() => (open.value = true), 1000)
    }
    return { open, toggleMenu }
  },
}
</script>

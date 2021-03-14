<template>
  <div class="home">
    <may-the-force-be-with-you></may-the-force-be-with-you>
    <img alt="Vue logo" src="../assets/logo.png" />
    <HelloWorld msg="Welcome to Your Vue.js App" />
  </div>
</template>

<script>
// @ is an alias to /src
import HelloWorld from "@/components/HelloWorld.vue";
import MayTheForceBeWithYou from "@/components/MayTheForceBeWithYou.vue";
export default {
  name: "Home",
  components: {
    HelloWorld,
    MayTheForceBeWithYou,
  },
};
</script>

<template>
  <div class="app">
    <ExoFlowHome />
  </div>
</template>

<script lang="ts">
import ExoFlowHome from "./components/ExoFlowHome.vue";

export default {
  name: "App",
  components: {
    ExoFlowHome,
  },
};
</script>

<style>
body {
  margin: 0;
  padding: 0;
  background-color: #1a1a1a;
  color: #e0e0e0;
}

.app {
  font-family: Arial, sans-serif;
}
</style>

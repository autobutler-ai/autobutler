<template>
  <div class="app">
    <ExoHeader />
    <ExoFlowHome />
    <ExoFooter />
  </div>
</template>

<script lang="ts">
import ExoFlowHome from "./components/ExoFlowHome.vue";
import ExoFooter from "./components/ExoFooter.vue";
import ExoHeader from "./components/ExoHeader.vue";

export default {
  name: "App",
  components: {
    ExoFlowHome,
    ExoFooter,
    ExoHeader,
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
  min-height: 100vh;
  display: flex;
  flex-direction: column;
}
</style>

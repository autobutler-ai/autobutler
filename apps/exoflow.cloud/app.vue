<template>
  <div class="app">
    <ExoHeader />
    <main class="main-content">
      <NuxtPage />
    </main>
    <ExoFooter />
  </div>
</template>

<script setup lang="ts">
import ExoHeader from "./components/ExoHeader.vue";
import ExoFooter from "./components/ExoFooter.vue";
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

.main-content {
  flex: 1;
  width: 100%;
}
</style>

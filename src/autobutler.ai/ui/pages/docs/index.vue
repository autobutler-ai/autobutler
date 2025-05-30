<template>
  <div>
    <!-- This will redirect before rendering -->
  </div>
</template>

<script setup lang="ts">
// Redirect to the welcome page which is handled by [...slug].vue
await navigateTo("/docs/welcome");
</script>

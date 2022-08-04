<script setup lang="ts">
import { ref } from 'vue';
import Login from './pages/Login.vue';
</script>

<template>
  <Login />
</template>

<style>
html, body {
  width: 100%;
  height: 100%;
}

#app {
  width: 100%;
  height: 100%;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}

* {
  font-family: "Source Sans Pro"
}
</style>

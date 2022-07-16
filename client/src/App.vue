<script setup lang="ts">
import { ref } from 'vue';
import Login from './pages/Login.vue';
import Register from './pages/Register.vue';

const currentRoute = ref("login");

const changeRoute = (path: string) => {
  currentRoute.value = path;
}
</script>

<template>
  <button @click="changeRoute('login')">Login</button>
  <button @click="changeRoute('register')">Register</button>
  <Login v-if="currentRoute === 'login'"/>
  <Register v-else/>
</template>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}
</style>

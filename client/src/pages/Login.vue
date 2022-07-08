<script setup lang="ts">
  import { ref } from 'vue';

  const email = ref("");
  const password = ref("");

  const logEmail = () => console.log(email)
  const logPassword = () => console.log(password)
</script>

<template>
  <input v-model="email" type="email" />
  <input v-model="password" type="password" />
</template>
<script setup lang="ts">
import { onMounted, ref } from "vue";
import axios from "axios";

const email = ref("");
const password = ref("");
const authenticated = ref(false);

onMounted(() => {
  const token = localStorage.getItem('token');
  authenticated.value = token != null;
})

const onSubmit = async () => {
  const response = await axios
    .post("http://localhost:5000/api/login", {
      email: email.value,
      password: password.value,
    })
    .then((res) => res.data)
    .catch((error) => console.error(error));

  localStorage.setItem('token', response.data);
  authenticated.value = true;
};
</script>

<template>
  <input v-model="email" type="email" placeholder="Email"/>
  <input v-model="password" type="password" placeholder="Password"/>
  <button @click="onSubmit">Submit</button>
  <span v-if="authenticated">Authenticated</span>
  <span v-else>Not authenticated</span>
</template>

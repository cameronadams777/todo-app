<script setup lang="ts">
import { onMounted, ref } from "vue";
import axios from "axios";

const email = ref("");
const password = ref("");

const onSubmit = async () => {
  const response = await axios
    .post("http://localhost:5000/api/login", {
      email: email.value,
      password: password.value,
    })
    .then((res) => res.data)
    .catch((error) => console.error(error));

  localStorage.setItem('token', response.data);
};
</script>

<template>
  <div class="w-full h-full flex flex-col justify-center items-center">
    <h2>Welcome Back!</h2>
    <div class="w-1/3 flex flex-col">
      <div class="text-sm font-bold mb-2 flex flex-col">
        <label for="email">Email</label>
        <input v-model="email" id="email" name="email" type="email" class="p-1"/>
      </div>
      <div class="text-sm font-bold mb-3 flex flex-col">
        <label for="password">Password</label>
        <input v-model="password" id="password" name="password" type="password" class="p-1"/>
      </div>
      <div class="flex justify-end mb-3">
        <button class="font-bold border-none text-indigo-600 rounded-md">Forgot Password?</button>
      </div>
      <div class="flex justify-center">
        <button class="w-1/2 p-2 border-none bg-indigo-600 text-white rounded-md" @click="onSubmit">Submit</button>
      </div>
    </div>
  </div>
</template>

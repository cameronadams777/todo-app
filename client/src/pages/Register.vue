<script setup lang="ts">
import { onMounted, ref } from "vue";
import axios from "axios";

const firstName = ref("");
const lastName = ref("");
const email = ref("");
const password = ref("");
const passwordConfirmation = ref("");
const registrationData = ref({});

const onSubmit = async () => {
  const response = await axios
    .post("http://localhost:5000/api/register", {
      firstName: firstName.value,
      lastName: lastName.value,
      email: email.value,
      password: password.value,
      passwordConfirmation: passwordConfirmation.value
    })
    .then((res) => res.data)
    .catch((error) => console.error(error));

  registrationData.value = response;

  localStorage.setItem('token', response.data.token);
};
</script>

<template>
  <input v-model="firstName" type="text" placeholder="First Name" />
  <input v-model="lastName" type="text" placeholder="Last Name" />
  <input v-model="email" type="email" placeholder="Email"/>
  <input v-model="password" type="password" placeholder="Password"/>
  <input v-model="passwordConfirmation" type="password" placeholder="Confirm Password"/>
  <button @click="onSubmit">Submit</button>
  <span>{{registrationData}}</span>
</template>

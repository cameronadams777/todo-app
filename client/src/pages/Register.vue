<script setup lang="ts">
import { ref } from "vue";
import { useRouter } from "vue-router";
import axios from "axios";

const router = useRouter();

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
      passwordConfirmation: passwordConfirmation.value,
    })
    .then((res) => res.data)
    .catch((error) => console.error(error));

  registrationData.value = response;

  localStorage.setItem("token", response.data.token);
  router.push("/");
};
</script>

<template>
  <div class="w-full h-full flex flex-col justify-center items-center">
    <h2>Nice to meet you!</h2>
    <div class="w-1/3 flex flex-col">
      <div class="text-sm font-bold mb-2 flex flex-col">
        <label for="firstName">First Name</label>
        <input
          v-model="firstName"
          id="firstName"
          name="firstName"
          type="text"
          class="p-1"
        />
      </div>
      <div class="text-sm font-bold mb-2 flex flex-col">
        <label for="lastName">Last Name</label>
        <input
          v-model="lastName"
          id="lastName"
          name="lastName"
          type="text"
          class="p-1"
        />
      </div>
      <div class="text-sm font-bold mb-2 flex flex-col">
        <label for="email">Email</label>
        <input
          v-model="email"
          id="email"
          name="email"
          type="email"
          class="p-1"
        />
      </div>
      <div class="text-sm font-bold mb-2 flex flex-col">
        <label for="password">Password</label>
        <input
          v-model="password"
          id="password"
          name="password"
          type="password"
          class="p-1"
        />
      </div>
      <div class="text-sm font-bold mb-3 flex flex-col">
        <label for="passwordConfirmation">Confirm Password</label>
        <input
          v-model="passwordConfirmation"
          id="passwordConfirmation"
          name="passwordConfirmation"
          type="password"
          class="p-1"
        />
      </div>
      <div class="flex justify-center">
        <button
          class="w-1/2 p-2 border-none bg-indigo-600 hover:bg-indigo-800 text-white rounded-md cursor-pointer"
          @click="onSubmit"
        >
          Sign Up
        </button>
      </div>
    </div>
  </div>
</template>

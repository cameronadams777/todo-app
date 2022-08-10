<script lang="ts" setup>
import axios from "axios";
import { onMounted, ref } from "vue";

interface ITodoItem {
  title: string;
}

const todoItems = ref<ITodoItem[]>([]);
const newTodoTitle = ref("");

onMounted(async () => {
  const token = localStorage.getItem("token");
  const response = await axios
    .get("http://localhost:5000/api/todos", {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    })
    .then((res) => res.data)
    .catch((err) => console.error(err));

  const todos = response.map((item: any) => ({
    title: item.Title,
  }));

  todoItems.value = todos;
});

const onSubmit = async () => {
  const token = localStorage.getItem("token");
  const response = await axios
    .post(
      "http://localhost:5000/api/todos",
      { title: newTodoTitle.value },
      {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      }
    )
    .then((res) => res.data)
    .catch((err) => console.error(err));

  todoItems.value.push({
    title: response.Title,
  });
};
</script>

<template>
  <div>
    <input
      v-model="newTodoTitle"
      id="new-todo-item"
      name="new-todo-item"
      type="text"
    />
    <button @click="onSubmit">Create</button>
    <div>
      <p v-for="item in todoItems">{{ item.title }}</p>
    </div>
  </div>
</template>

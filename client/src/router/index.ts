import Login from "../pages/Login.vue";
import Register from "../pages/Register.vue";
import Home from "../pages/Home.vue";
import { createRouter, createWebHistory, RouteRecordRaw } from "vue-router";

const routes: RouteRecordRaw[] = [
  { path: "/", component: Home },
  { path: "/login", component: Login },
  { path: "/register", component: Register },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

router.beforeEach((to) => {
  const nonAuthenticatedRoutes = ["/login", "/register"];
  const isAuthenticated = localStorage.getItem("token")?.length;
  if (!isAuthenticated && !nonAuthenticatedRoutes.includes(to.path)) {
    return { path: "/login" };
  } else if (isAuthenticated && nonAuthenticatedRoutes.includes(to.path)) {
    return { path: "/" };
  }
});

export default router;

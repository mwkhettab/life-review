import { createWebHistory, createRouter } from "vue-router";

import Home from "./views/Home.vue";
import Reviews from "./views/Reviews.vue";
import NotFound from "./views/NotFound.vue";
import Admin from "./views/Admin.vue";

const routes = [
  {
    path: "/",
    component: Home,
  },
  {
    path: "/reviews",
    component: Reviews,
  },
  {
    path: "/:catchAll(.*)",
    component: NotFound,
  },
  {
    path: "/admin",
    component: Admin,
  }
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;

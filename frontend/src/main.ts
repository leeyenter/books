import { createApp } from "vue";
import { createRouter, createWebHistory } from "vue-router";
import { VueQueryPlugin } from "@tanstack/vue-query";
import "./style.css";
import App from "./App.vue";
import MainView from "./views/MainView.vue";
import NewBookView from "./views/NewBookView.vue";
import EditBookView from "./views/EditBookView.vue";
import LoginView from "./views/LoginView.vue";

export const router = createRouter({
  history: createWebHistory(),
  routes: [
    { path: "/", component: MainView },
    { path: "/login", component: LoginView },
    { path: "/new", component: NewBookView },
    { path: "/edit/:id", component: EditBookView },
  ],
});

const app = createApp(App);
app.use(router);
app.use(VueQueryPlugin);
app.mount("#app");

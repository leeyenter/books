import { createApp } from "vue";
import { createRouter, createWebHistory } from "vue-router";
import { VueQueryPlugin } from "@tanstack/vue-query";
import "./style.css";
import App from "./App.vue";
import MainView from "./views/MainView.vue";

const router = createRouter({
  history: createWebHistory(),
  routes: [{ path: "/", component: MainView }],
});

const app = createApp(App);
app.use(router);
app.use(VueQueryPlugin);
app.mount("#app");

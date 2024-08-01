import { createApp } from "vue";
import "./style.css";
import App from "./App.vue";
import Gorik from "./path/Gorik.vue";

import { createMemoryHistory, createRouter } from "vue-router";

const routes = [{ path: "/", component: Gorik }];
const router = createRouter({
    history: createMemoryHistory(),
    routes,
});

createApp(App).use(router).mount("#app");
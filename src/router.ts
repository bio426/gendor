import { createRouter, createWebHistory, RouteRecordRaw } from "vue-router"

import itemRouter from "./views/item/router"

const routes: RouteRecordRaw[] = [
	{ path: "/", name: "index", component: () => import("./views/index.vue") },
	{ path: "/dashboard", name: "dashboard", component: () => import("./views/dashboard.vue") },
	{ path: "/pos", name: "pos", component: () => import("./views/pos.vue") },
	...itemRouter,
]

export default createRouter({
	routes,
	history: createWebHistory()
})

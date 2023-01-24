import { createRouter, createWebHistory, RouteRecordRaw } from "vue-router"

import itemRouter from "./views/item/router"
import posRouter from "./views/pos/router"

const routes: RouteRecordRaw[] = [
	{ path: "/", name: "index", component: () => import("./views/index.vue") },
	...itemRouter,
	...posRouter,
]

export default createRouter({
	routes,
	history: createWebHistory(),
})

import { createRouter, createWebHistory, RouteRecordRaw } from "vue-router"

import itemRouter from "./views/item/router"
import posRouter from "./views/pos/router"
import adminRouter from "./views/admin/router"

const routes: RouteRecordRaw[] = [
	{ path: "/", name: "index", component: () => import("./views/index.vue") },
	...itemRouter,
	...posRouter,
	...adminRouter,
]

export default createRouter({
	routes,
	history: createWebHistory(),
})

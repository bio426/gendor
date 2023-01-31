import { createRouter, createWebHistory, RouteRecordRaw } from "vue-router"

import itemRouter from "./views/item/router"
import posRouter from "./views/pos/router"
import adminRouter from "./views/admin/router"
import saleRouter from "./views/sale/router"

const routes: RouteRecordRaw[] = [
	{ path: "/", name: "index", component: () => import("./views/index.vue") },
	{ path: "/dev", name: "dev", component: () => import("./views/dev.vue") },
	...itemRouter,
	...posRouter,
	...adminRouter,
	...saleRouter,
]

export default createRouter({
	routes,
	history: createWebHistory(),
})

import { createRouter, createWebHistory } from "vue-router"
import type { RouteRecordRaw } from "vue-router"

import products from "./products"

const routes: RouteRecordRaw[] = [
	{
		path: "/",
		name: "index",
		component: () => import("@/views/index.vue"),
	},
	{
		path: "/login",
		name: "login",
		component: () => import("@/views/login.vue"),
	},
	{
		path: "/dev",
		name: "dev",
		component: () => import("@/views/dev.vue"),
	},
	{
		path: "/:pathMatch(.*)*",
		name: "not-found",
		component: () => import("@/views/not-found.vue"),
	},
	...products,
]

const router = createRouter({
	routes,
	history: createWebHistory(import.meta.env.BASE_URL),
})

export default router

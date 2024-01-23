import type { RouteRecordRaw } from "vue-router"

const prefix = "/products"

const routes: RouteRecordRaw[] = [
	{
		path: prefix + "/list",
		name: "products-list",
		component: () => import("@/views/products/list.vue"),
	},
	{
		path: prefix + "/create",
		name: "products-create",
		component: () => import("@/views/products/create.vue"),
	},
]

export default routes

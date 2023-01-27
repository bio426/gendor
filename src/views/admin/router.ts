import { RouteRecordRaw } from "vue-router"

const prefix = "/admin"

const routesItem: RouteRecordRaw[] = [
	{ path: prefix + "/", name: "admin", component: () => import("./index.vue") },
	{
		path: prefix + "/tag",
		name: "adminTag",
		component: () => import("./tag.vue"),
	},
]

export default routesItem

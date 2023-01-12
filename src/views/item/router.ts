import { RouteRecordRaw } from "vue-router"

const prefix = "/item"

const routesItem: RouteRecordRaw[] = [
	{ path: prefix + "/", name: "itemList", component: () => import("./index.vue") },
	{ path: prefix + "/create", name: "itemCreate", component: () => import("./create.vue") }
]

export default routesItem

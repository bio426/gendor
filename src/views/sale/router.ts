import { RouteRecordRaw } from "vue-router"

const prefix = "/sale"

const routesItem: RouteRecordRaw[] = [
	{ path: prefix + "/", name: "sale", component: () => import("./index.vue") },
]

export default routesItem

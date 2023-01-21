import { RouteRecordRaw } from "vue-router"

const prefix = "/pos"

const routesItem: RouteRecordRaw[] = [
	{ path: prefix + "/", name: "pos", component: () => import("./index.vue") },
	{ path: prefix + "/cart", name: "posCart", component: () => import("./cart.vue") },
	{ path: prefix + "/client", name: "posClient", component: () => import("./client.vue") },
]

export default routesItem

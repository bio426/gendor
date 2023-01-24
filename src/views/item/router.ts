import { RouteRecordRaw } from "vue-router"

const prefix = "/item"

const routesItem: RouteRecordRaw[] = [
	{
		path: prefix + "/",
		name: "itemList",
		component: () => import("./index.vue"),
	},
	{
		path: prefix + "/create",
		name: "itemCreate",
		component: () => import("./create.vue"),
	},
	{
		path: prefix + "/:id",
		name: "itemId",
		component: () => import("./id.vue"),
	},
]

export default routesItem

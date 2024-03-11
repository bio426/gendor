import type { RouteRecordRaw } from "vue-router"

const prefix = "/products"

const route: RouteRecordRaw = {
    path: prefix + "/list",
    name: "products-list",
    component: () => import("@/view/products/list.vue"),
}

export default route

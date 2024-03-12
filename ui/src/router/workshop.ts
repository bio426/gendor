import type { RouteRecordRaw } from "vue-router"

const route: RouteRecordRaw = {
    path: "/workshop",
    name: "workshop",
    redirect: { name: "workshop-list" },
    meta: { permittedRoles: ["admin"] },
    children: [
        {
            path: "list",
            name: "workshop-list",
            component: () => import("@/view/workshop/list.vue")
        },
        {
            path: "order",
            name: "workshop-create",
            component: () => import("@/view/workshop/create.vue")
        },
    ]
}

export default route

import type { RouteRecordRaw } from "vue-router"

const route: RouteRecordRaw = {
    path: "/user",
    name: "user",
    redirect: { name: "user-list" },
    meta: { permittedRoles: ["super"] },
    children: [
        {
            path: "list",
            name: "user-list",
            component: () => import("@/view/user/list.vue")
        },
        {
            path: "create",
            name: "user-create",
            component: () => import("@/view/user/create.vue")
        },
    ]
}

export default route

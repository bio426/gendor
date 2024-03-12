import base from "./base"

import * as tWorkshop from "@/type/workshop"

const path = "workshop"

export default {
    async list() {
        const res = await base.get(path)

        return res.json<{ rows: any[] }>()
    },
    async create(body: { username: string, password: string }) {
        const res = await base.post(path, { json: body })

        return res.status
    },
}

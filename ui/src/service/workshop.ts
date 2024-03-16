import base from "./base"

import * as tWorkshop from "@/type/workshop"

const path = "workshop"

export default {
    async list(query: { search: string, page: number }) {
        const res = await base.get(path, {
            searchParams: query
        })

        return res.json<{
            total: number,
            from: number,
            to: number,
            rows: tWorkshop.OrderSimple[]
        }>()
    },
    async detail(id: number) {
        const res = await base.get(`${path}/${id}`)

        return res.json<tWorkshop.OrderDetail>()
    },
    async create(body: tWorkshop.OrderCreate) {
        const res = await base.post(path, { json: body })

        return res.status
    },
    async searchPlate(query: { plate: string }) {
        const res = await base.get(path + "/search-plate", {
            searchParams: query
        })

        return res.json<tWorkshop.OrderFound>()
    },
}

import base from "./_base"
import { IItem } from "../interfaces/item"

const prefix = "item"

async function create(body: IItem) {
	const res = await base.post(prefix, {
		json: body,
	})

	return res
}

async function list(query: { page: string; count: string; search: string }) {
	const params = new URLSearchParams(query)

	const res = await base.get(prefix + "?" + params.toString())

	return res.json<IItem[]>()
}

async function read(id: string) {
	const res = await base.get("/item/" + id)

	return res.json<any>()
}

async function getTags() {
	const res = await base.get(prefix + "/tag")

	return res.json<string[]>()
}

export default {
	create,
	list,
	read,
	getTags,
}

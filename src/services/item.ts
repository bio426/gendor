import base from "./_base"

const prefix = "item"

async function create(body: { name: string; price: number; tags: string[] }) {
	const res = await base.post(prefix, {
		json: body,
	})

	return res
}

async function list(query: { page: string; count: string; search: string }) {
	const params = new URLSearchParams(query)

	const res = await base.get(prefix + "?" + params.toString())

	return res.json<{
		total: number
		rows: {
			name: string
			price: number
			tags: number[]
		}[]
	}>()
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

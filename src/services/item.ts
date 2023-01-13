import base from "./_base"

const prefix = "item"

async function create(body: {
	name: string
	price: number
	tags: number[]
}) {
	const res = await base.post(prefix, {
		json: body
	})

	return res
}

async function list(query: {
	page: string
	count: string
}) {
	const params = new URLSearchParams(query)

	const res = await base.get(prefix + "?" + params.toString())

	return res
}

async function read(query: {
	id: string
}) {
	const params = new URLSearchParams(query)

	const res = await base.get("/item?" + params.toString())
}

export default {
	create,
	list,
	read,
}

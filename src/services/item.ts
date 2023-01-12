import base from "./_base"

async function create(body: {
	name: string
	price: number
}) {
	const res = await base.post("/", {})
}

async function read(query: {
	id: string
}) {
	const params = new URLSearchParams(query)

	const res = await base.get("/item?" + params.toString())
}

export default {
	create,
	read,
}

import base from "./_base"

const prefix = "pos"

async function getItems(query: { search: string }) {
	const params = new URLSearchParams(query)
	const res = await base.get(prefix + "/items" + "?" + params.toString())

	return res.json<
		{
			id: string
			name: string
			price: number
		}[]
	>()
}

export default {
	getItems,
}

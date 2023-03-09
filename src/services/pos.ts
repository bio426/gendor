import pb from "./_base"

const productsColl = pb.collection("products")

async function getItems(query: { search: string }) {
	const res = await productsColl.getList(1, 20)

	return res.items.map((item) => ({
		id: item["id"],
		name: item["name"],
		price: item["price"],
	}))
}

export default {
	getItems,
}

import pb from "./_base"
import { IProduct, ICreateProduct } from "../interfaces/product"

const prefix = "item"

const productsColl = pb.collection("products")

async function create(body: ICreateProduct) {
	const res = await productsColl.create(body)

	return res
}

async function list(query: { page: number; count: number; search: string }) {
	const res = await productsColl.getList(query.page, query.count, {
		expand: "category",
	})
	console.log(res.items)

	return {
		total: res.totalItems,
		rows: res.items.map<IProduct>((item) => {
			const category = item.expand.category as any
			return {
				id: item["id"],
				name: item["name"],
				price: item["price"],
				category: {
					id: category["id"],
					name: category["name"],
					color: category["color"],
				},
			}
		}),
	}
}

async function read(id: string) {
	const res = await productsColl.getOne(id)

	return res
}

export default {
	create,
	list,
	read,
}

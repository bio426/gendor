import base from "./base"
import type * as tProduct from "@/type/product"

const prefix = "product"

const productService = {
	async create(body: any) {
		const formData = new FormData()
		formData.append("name", body.name)
		formData.append("price", body.price.toString())
		formData.append("image", body.image)

		const res = await base.post(prefix, { body: formData })

		return res.status
	},
	async list(query: { page: number; count: number; search: string }) {
		const res = await base.get(prefix, {
			searchParams: query,
		})

		return res.json<{
			total: number
			rows: tProduct.Product[]
		}>()
	},
	async createProductCategory(body: { name: string }) {
		const res = await base.post(prefix + "/category", { json: body })

		return res.status
	},
	async listProductCategories() {
		const res = await base.get(prefix + "/category")

		return res.json<{
			rows: any[]
		}>()
	},
}

export default productService

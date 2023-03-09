import pb from "./_base"
import { ICategory, ICreateCategory } from "../interfaces/admin"

const CategoriesColl = pb.collection("product_categories")

async function getTags() {
	const res = await CategoriesColl.getList(1, 20)

	return res.items.map<ICategory>((item) => ({
		id: item["id"],
		name: item["name"],
		color: item["color"],
	}))
}
async function createTag(body: ICreateCategory) {
	await CategoriesColl.create(body)
}
async function deleteTag(id: string) {
	await CategoriesColl.delete(id)
}

export default {
	getTags,
	createTag,
	deleteTag,
}

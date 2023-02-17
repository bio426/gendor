import base from "./_base"
import { IAdminTag } from "../interfaces/admin"

const prefix = "admin"

async function getTags() {
	const res = await base.get(prefix + "/tag")

	return res.json<IAdminTag[]>()
}
async function createTag(body: IAdminTag) {
	const res = await base.post(prefix + "/tag", { json: body })

	return res.status
}
async function deleteTag(body: { name: string }) {
	const res = await base.delete(prefix + "/tag", { json: body })

	return res.status
}

export default {
	getTags,
	createTag,
	deleteTag,
}

import base from "./_base"

const prefix = "admin"

async function getTags() {
	const res = await base.get(prefix + "/tag")

	return res.json<string[]>()
}
async function createTag(body: { value: string }) {
	const res = await base.post(prefix + "/tag", { json: body })

	return res.status
}
async function deleteTag(body: { value: string }) {
	const res = await base.delete(prefix + "/tag", { json: body })

	return res.status
}

export default {
	getTags,
	createTag,
	deleteTag,
}

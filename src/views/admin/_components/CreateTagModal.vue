<script setup lang="ts">
import { } from "vue"
import { useField } from "vee-validate"
import * as yup from "yup"

import adminService from "../../../services/admin"
import useToast from "../../../composables/useToast"
import Modal from "../../../components/Modal.vue"

const { showToast } = useToast()
const props = defineProps<{
	show: boolean
}>()
const emit = defineEmits(["closing"])
function handleClose() {
	tagR()
	tag.value = ""
	emit("closing")
}

const {
	value: tag,
	errorMessage: tagE,
	validate: tagC,
	resetField: tagR,
} = useField<string>("name", yup.string().required(), { initialValue: "" })

async function save() {
	const check = await tagC()
	if (!check.valid) return
	const body = {
		value: tag.value
	}
	await adminService.createTag(body)
	showToast("New tag succesfully created", 2000)
	handleClose()
}
</script>

<template>
	<Modal title="Create Tag" :show="show" @closing="handleClose">
		<div class="form-control w-full mb-8">
			<label class="label">
				<span class="label-text">Tag</span>
			</label>
			<input type="text" class="input input-bordered w-full" :class="{ 'input-error': !!tagE }" v-model="tag" />
			<label class="label">
				<span class="label-text-alt" v-if="tagE">{{ tagE }}</span>
			</label>
		</div>
		<button class="btn btn-success btn-block" @click="save">Save</button>
	</Modal>
</template>

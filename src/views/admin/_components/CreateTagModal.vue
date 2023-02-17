<script setup lang="ts">
import {} from "vue"
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
	nameR()
	colorR()
	emit("closing")
}

const {
	value: name,
	errorMessage: nameE,
	validate: nameC,
	resetField: nameR,
} = useField<string>("name", yup.string().required(), { initialValue: "" })

const {
	value: color,
	errorMessage: colorE,
	validate: colorC,
	resetField: colorR,
} = useField<string>("color", yup.string().required(), { initialValue: "" })

async function allValid(): Promise<boolean> {
	const results = await Promise.all([nameC(), colorC()])
	if (results.some((res) => !res.valid)) return false
	return true
}

async function save() {
	if (!(await allValid())) return
	const body = {
		name: name.value,
		color: color.value,
	}
	await adminService.createTag(body)
	showToast("New tag succesfully created", 2000)
	handleClose()
}
</script>

<template>
	<Modal title="Create Tag" :show="show" @closing="handleClose">
		<div class="form-control w-full mb-2">
			<label class="label">
				<span class="label-text">Tag</span>
			</label>
			<input
				type="text"
				class="input input-bordered w-full"
				:class="{ 'input-error': !!nameE }"
				v-model="name"
			/>
			<label class="label">
				<span class="label-text-alt" v-if="nameE">{{ nameE }}</span>
			</label>
		</div>
		<div class="form-control w-full mb-8">
			<label class="label">
				<span class="label-text">Color</span>
			</label>
			<input
				type="color"
				class="input input-bordered w-full"
				:class="{ 'input-error': !!colorE }"
				v-model="color"
			/>
			<label class="label">
				<span class="label-text-alt" v-if="colorE">{{ colorE }}</span>
			</label>
		</div>
		<button class="btn btn-success btn-block" @click="save">Save</button>
	</Modal>
</template>

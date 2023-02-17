<script setup lang="ts">
import { ref } from "vue"
import { useRouter } from "vue-router"
import { useField } from "vee-validate"
import * as yup from "yup"
import { XMarkIcon, ArrowUturnLeftIcon } from "@heroicons/vue/24/solid"

import itemService from "../../services/item"
import adminService from "../../services/admin"
import useToast from "../../composables/useToast"
import {IAdminTag} from "../../interfaces/admin"

const router = useRouter()
const { showToast } = useToast()

const tagOpts = ref<IAdminTag[]>([])

async function getTagOpts() {
	const data = await adminService.getTags()
	tagOpts.value = data
}
getTagOpts()

const {
	value: name,
	errorMessage: nameE,
	validate: nameC,
} = useField<string>("name", yup.string().required(), { initialValue: "" })
const {
	value: price,
	errorMessage: priceE,
	validate: priceC,
} = useField<number>("price", yup.number().min(1).required(), {
	initialValue: 0,
})
const {
	value: tags,
	errorMessage: tagsE,
	validate: tagsC,
} = useField<string[]>("tags", yup.array().min(1).required(), {
	initialValue: [],
})

async function allValid(): Promise<boolean> {
	const results = await Promise.all([nameC(), priceC(), tagsC()])
	if (results.some((res) => !res.valid)) return false
	return true
}

async function createProduct() {
	if (!(await allValid())) return
	const body = {
		name: name.value,
		price: price.value,
		tags: tags.value,
	}
	await itemService.create(body)
	showToast("New product created", 2000)
	router.push({ name: "itemList" })
}

// Tags select
function onTagSelect(ev: Event) {
	const el = ev.target as HTMLSelectElement
	const tag = el.options[el.selectedIndex].value
	const selectedOpt = tagOpts.value.find((opt) => opt.name == tag)
	if (
		selectedOpt == undefined ||
		tags.value.some((tag) => tag == selectedOpt.name)
	)
		return
	tags.value.push(selectedOpt.name)
}
function removeTag(idx: number) {
	tags.value.splice(idx, 1)
}
</script>

<template>
	<div class="relative w-full min-h-screen">
		<div class="w-11/12 mx-auto">
			<h1 class="py-4 text-2xl font-bold text-center">
				<a
					class="inline float-left cursor-pointer"
					@click="$router.back()"
				>
					<ArrowUturnLeftIcon class="w-6 h-6 mt-1" />
				</a>
				Create Product
			</h1>
			<div class="form-control w-full">
				<label class="label">
					<span class="label-text">Name</span>
				</label>
				<input
					class="input input-bordered w-full"
					type="text"
					v-model="name"
					:class="{ 'input-error': !!nameE }"
				/>
				<label class="label">
					<span class="label-text-alt" v-if="nameE">{{ nameE }}</span>
				</label>
			</div>
			<div class="form-control w-full">
				<label class="label">
					<span class="label-text">Price</span>
				</label>
				<input
					class="input input-bordered w-full"
					type="number"
					v-model.number="price"
					:class="{ 'input-error': !!priceE }"
				/>
				<label class="label">
					<span class="label-text-alt" v-if="priceE">{{
						priceE
					}}</span>
				</label>
			</div>
			<div class="form-control w-full">
				<label class="label">
					<span class="label-text">Tags</span>
				</label>
				<select
					class="select select-bordered w-full"
					:class="{ 'select-error': !!tagsE }"
					@change="onTagSelect"
				>
					<option hidden selected></option>
					<option v-for="opt in tagOpts" :value="opt.name">
						{{ opt.name }}
					</option>
				</select>
				<label class="label">
					<span class="label-text-alt" v-if="tagsE">{{ tagsE }}</span>
				</label>
			</div>
			<div class="flex flex-wrap gap-4">
				<div
					class="badge badge-success gap-2 cursor-pointer"
					v-for="(tag, i) in tags"
					@click="removeTag(i)"
				>
					{{ tag }}
					<XMarkIcon class="h-4 w-6" />
				</div>
			</div>
			<!-- spacer -->
			<div class="w-full h-24"></div>
		</div>
		<!-- bottom actions -->
		<div class="fixed left-0 bottom-0 w-full bg-base-100">
			<div class="w-11/12 mx-auto py-4">
				<button
					class="btn btn-success btn-block"
					@click="createProduct"
				>
					Create
				</button>
			</div>
		</div>
	</div>
</template>

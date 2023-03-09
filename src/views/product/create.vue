<script setup lang="ts">
import { ref } from "vue"
import { useRouter } from "vue-router"
import { useField } from "vee-validate"
import * as yup from "yup"
import { XMarkIcon, ArrowUturnLeftIcon } from "@heroicons/vue/24/solid"

import itemService from "../../services/product"
import adminService from "../../services/admin"
import useToast from "../../composables/useToast"
import { ICategory } from "../../interfaces/admin"

const router = useRouter()
const { showToast } = useToast()

const tagOpts = ref<ICategory[]>([])

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
	value: category,
	errorMessage: categoryE,
	validate: categoryC,
} = useField<string>("category", yup.string().required(), {
	initialValue: "",
})

async function allValid(): Promise<boolean> {
	const results = await Promise.all([nameC(), priceC(), categoryC()])
	if (results.some((res) => !res.valid)) return false
	return true
}

async function createProduct() {
	if (!(await allValid())) return
	const body = {
		name: name.value,
		price: price.value,
		category: category.value,
	}
	await itemService.create(body)
	showToast("New product created", 2000)
	router.push({ name: "itemList" })
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
					:class="{ 'select-error': !!categoryE }"
					v-model="category"
				>
					<option hidden selected></option>
					<option v-for="opt in tagOpts" :value="opt.id">
						{{ opt.name }}
					</option>
				</select>
				<label class="label">
					<span class="label-text-alt" v-if="categoryE">{{
						categoryE
					}}</span>
				</label>
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

<script setup lang="ts">
import { ref, reactive, computed } from "vue"
import { CurrencyDollarIcon, PlusIcon } from "@heroicons/vue/24/solid"

import useConfirmation from "@/composable/useConfirmation"
import useToast from "@/composable/useToast"
import productService from "@/service/product"
import Header from "@/component/Header.vue"
import Overlay from "@/component/Overlay.vue"
import CreateProductCategory from "@/component/products/CreateProductCategory.vue"
import ManageProductVehicles from "@/component/products/ManageProductVehicles.vue"
import type * as tProduct from "@/type/product"

const confirmation = useConfirmation()
const toasts = useToast()

const loading = ref(false)

const form = reactive<{
	name: string
	price: number
	image: File | undefined
	category: number
}>({
	name: "",
	price: 0,
	image: undefined,
	category: 0,
})

function setFormImage(e: Event) {
	const target = e.target as HTMLInputElement
	if (target.files && target.files.length != 0) {
		form.image = target.files.item(0) as File
	}
}

const categoryOpts = ref<any[]>([])
async function getCategoryOpts() {
	loading.value = true
	const data = await productService.listProductCategories()
	categoryOpts.value = data.rows
	loading.value = false
}
getCategoryOpts()

async function create() {
	if (!(await confirmation.display({}))) return
	if (form.image == undefined) return
	loading.value = true
	await productService.create({
		name: form.name,
		price: form.price,
		image: form.image,
	})
	loading.value = false
	toasts.display({ message: "adasd" })
}

const vehicles = ref<{ id: number; model: string; brand: string }[]>([])
const vehiclesResume = computed(() => {
	if (vehicles.value.length == 0) return "No vehicles selected"
	return `${vehicles.value.length} vehicles selected`
})

// Modals
const categoryModal = reactive({ show: false })
const productVehicles = reactive({ show: false })
</script>

<template>
	<div class="relative w-screen min-h-screen">
		<div class="w-11/12 mx-auto">
			<Header title="Create Product" />
			<Overlay :show="loading">
				<label class="form-control w-full">
					<div class="label">
						<span class="label-text">Name</span>
					</div>
					<input
						class="input input-bordered w-full"
						type="text"
						v-model="form.name"
					/>
				</label>
				<label class="form-control w-full">
					<div class="label">
						<span class="label-text">Price</span>
					</div>
					<div class="join w-full">
						<button
							class="join-item btn btn-square btn-outline"
							disabled
						>
							<CurrencyDollarIcon class="w-6 h-6 fill-white" />
						</button>
						<input
							class="join-item input input-bordered w-full"
							type="number"
							v-model="form.price"
						/>
					</div>
				</label>
				<label class="form-control w-full">
					<div class="label">
						<span class="label-text">File</span>
					</div>
					<input
						class="file-input file-input-bordered w-full"
						type="file"
						@input="setFormImage"
					/>
				</label>
				<hr class="my-4" />
				<div class="form-control w-full">
					<div class="label">
						<span class="label-text">Category</span>
						<span class="label-text-alt">
							<button
								class="btn btn-xs"
								@click="
									categoryModal.show = !categoryModal.show
								"
							>
								<PlusIcon class="w-4 fill-white" />
							</button>
						</span>
					</div>
					<select
						class="select select-bordered w-full"
						v-model="form.category"
					>
						<option
							v-for="category in categoryOpts"
							:key="category.id"
							:value="category.id"
						>
							{{ category.name }}
						</option>
					</select>
				</div>
				<div class="form-control w-full">
					<div class="label">
						<span class="label-text">Vehicles</span>
						<span class="label-text-alt">
							<button
								class="btn btn-xs"
								@click="
									categoryModal.show = !categoryModal.show
								"
							>
								<PlusIcon class="w-4 fill-white" />
							</button>
						</span>
					</div>
					<div class="join">
						<button
							class="btn join-item"
							@click="
								productVehicles.show = !productVehicles.show
							"
						>
							Add
						</button>
						<input
							class="input input-bordered join-item w-full"
							disabled
							:value="vehiclesResume"
						/>
					</div>
				</div>
				<button class="btn btn-block mt-8" @click="create">
					Create
				</button>
			</Overlay>
		</div>
		<CreateProductCategory
			:show="categoryModal.show"
			@close="categoryModal.show = false"
			@refresh="getCategoryOpts"
		/>
		<ManageProductVehicles
			:show="productVehicles.show"
			:vehicles="vehicles"
			@close="productVehicles.show = false"
			@update="vehicles = $event"
		/>
	</div>
</template>

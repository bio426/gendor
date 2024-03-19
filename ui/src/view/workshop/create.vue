<script setup lang="ts">
import { ref, reactive } from "vue"
import { useRouter } from "vue-router"
import { MagnifyingGlassIcon, PencilSquareIcon } from "@heroicons/vue/24/solid"

import * as tWorkshop from "@/type/workshop"
import useToast from "@/composable/useToast"
import workshopService from "@/service/workshop"
import Overlay from "@/component/Overlay.vue"
import OrderItemsModal from "@/component/workshop/OrderItemsModal.vue"

const router = useRouter()
const toast = useToast()

const loading = ref(false)

const form = reactive({
	name: "",
	address: "",
	dni: "",
	ruc: "",
	brand: "",
	model: "",
	color: "",
	plate: "",
	mileage: 0,
	observation: "",
	items: [],
})

const itemsModal = reactive({ show: false })

async function searchPlate() {
	loading.value = true
	const found = await workshopService.searchPlate({ plate: form.plate })
	loading.value = false
	if (found.name == undefined) {
		toast.display({ message: "Plate not found", variant: "error" })
		return
	}
	form.name = found.name
	form.address = found.address
	form.dni = found.dni
	form.ruc = found.ruc
	form.brand = found.brand
	form.model = found.model
	form.color = found.color
	form.mileage = found.mileage
}

async function create() {
	// validate items
	if (form.items.length == 0) {
		toast.display({
			message: "At leas 1 item has to be included",
			variant: "error",
		})
		return
	}
	loading.value = true
	await workshopService.create({
		name: form.name,
		address: form.address,
		dni: form.dni,
		ruc: form.ruc,
		brand: form.brand,
		model: form.model,
		color: form.plate,
		plate: form.plate,
		mileage: form.mileage,
		observation: form.observation,
		items: form.items,
	})
	loading.value = false
	router.push({ name: "workshop-list" })
}
</script>

<template>
	<div class="relative min-h-screen">
		<div class="w-11/12 mx-auto">
			<h1 class="py-8 text-2xl font-bold text-center">Create Order</h1>
			<Overlay :show="loading">
				<form @submit.prevent="create" autocomplete="off">
					<div
						class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4"
					>
						<div class="form-control w-full">
							<div class="label">
								<label class="label-text" for="i1">Placa</label>
								<span class="label-text-alt">
									<button
										class="btn btn-sm btn-square"
										@click.prevent="searchPlate"
									>
										<MagnifyingGlassIcon class="w-4 h-4" />
									</button>
								</span>
							</div>
							<input
								class="input input-bordered w-full"
								type="text"
								id="i1"
								required
								v-model="form.plate"
							/>
						</div>
						<label class="form-control w-full">
							<div class="label">
								<span class="label-text">Nombre</span>
							</div>
							<input
								class="input input-bordered w-full"
								type="text"
								v-model="form.name"
							/>
						</label>
						<label class="form-control w-full">
							<div class="label">
								<span class="label-text">Direccion</span>
							</div>
							<input
								class="input input-bordered w-full"
								type="text"
								v-model="form.address"
							/>
						</label>
						<label class="form-control w-full">
							<div class="label">
								<span class="label-text">DNI</span>
							</div>
							<input
								class="input input-bordered w-full"
								type="text"
								v-model="form.dni"
							/>
						</label>
						<label class="form-control w-full">
							<div class="label">
								<span class="label-text">RUC</span>
							</div>
							<input
								class="input input-bordered w-full"
								type="text"
								v-model="form.ruc"
							/>
						</label>
						<label class="form-control w-full">
							<div class="label">
								<span class="label-text">Marca</span>
							</div>
							<input
								class="input input-bordered w-full"
								type="text"
								v-model="form.brand"
							/>
						</label>
						<label class="form-control w-full">
							<div class="label">
								<span class="label-text">Modelo</span>
							</div>
							<input
								class="input input-bordered w-full"
								type="text"
								v-model="form.model"
							/>
						</label>
						<label class="form-control w-full">
							<div class="label">
								<span class="label-text">Color</span>
							</div>
							<input
								class="input input-bordered w-full"
								type="text"
								v-model="form.color"
							/>
						</label>
						<label class="form-control w-full">
							<div class="label">
								<span class="label-text">Kilometraje</span>
							</div>
							<input
								class="input input-bordered w-full"
								type="number"
								min="0"
								step="1"
								v-model="form.mileage"
							/>
						</label>
						<label class="form-control w-full">
							<div class="label">
								<span class="label-text">Observacion</span>
							</div>
							<textarea
								class="textarea textarea-bordered w-full"
								v-model="form.observation"
							/>
						</label>
						<div class="form-control w-full">
							<div class="label">
								<label class="label-text" for="i2">Items</label>
							</div>
							<div class="join">
								<input
									class="input input-bordered join-item w-full"
									type="number"
									id="i2"
									readonly
									:value="form.items.length"
								/>
								<button
									class="btn btn-square join-item"
									role="button"
									@click.prevent="
										itemsModal.show = !itemsModal.show
									"
								>
									<PencilSquareIcon class="w-6 h-6" />
								</button>
							</div>
						</div>
					</div>

					<button
						class="btn btn-block btn-primary mt-8"
						type="submit"
					>
						Crear
					</button>
				</form>
			</Overlay>
		</div>
		<OrderItemsModal
			:show="itemsModal.show"
			v-model="form.items"
			@closing="itemsModal.show = false"
		/>
	</div>
</template>

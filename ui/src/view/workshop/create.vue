<script setup lang="ts">
import { ref, reactive } from "vue"
import { useRouter } from "vue-router"
import { MagnifyingGlassIcon, PencilSquareIcon } from "@heroicons/vue/24/solid"

import * as tWorkshop from "@/type/workshop"
import useToast from "@/composable/useToast"
import workshopService from "@/service/workshop"
import Overlay from "@/component/Overlay.vue"
import OrderItems from "@/component/workshop/OrderItems.vue"

const router = useRouter()
const toast = useToast()

const loading = ref(false)

const form = reactive({
	propietary: "",
	brand: "",
	model: "",
	year: 0,
	plate: "",
	mileage: 0,
	observation: "",
	discount: 0,
	items: [],
})

async function searchPlate() {
	loading.value = true
	const found = await workshopService.searchPlate({ plate: form.plate })
	loading.value = false
	if (found.propietary == undefined) {
		toast.display({ message: "Plate not found", variant: "error" })
		return
	}
	form.propietary = found.propietary
	form.brand = found.brand
	form.model = found.model
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
		propietary: form.propietary,
		brand: form.brand,
		model: form.model,
		year: form.year,
		plate: form.plate,
		mileage: form.mileage,
		discount: 0,
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
					<div class="grid grid-cols-1 md:grid-cols-2 gap-4">
						<div class="grid grid-cols-1 md:grid-cols-2 gap-4">
							<!-- datos iniciales -->
							<div class="form-control w-full">
								<div class="label">
									<label class="label-text" for="i1">
										Placa
									</label>
									<span class="label-text-alt">
										<button
											class="btn btn-sm btn-square"
											@click.prevent="searchPlate"
										>
											<MagnifyingGlassIcon
												class="w-4 h-4"
											/>
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
									<span class="label-text">Propietario</span>
								</div>
								<input
									class="input input-bordered w-full"
									type="text"
									v-model="form.propietary"
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
									<span class="label-text">Año</span>
								</div>
								<input
									class="input input-bordered w-full"
									type="number"
									min="0"
									step="1"
									v-model="form.year"
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
							<label class="form-control w-full">
								<div class="label">
									<span class="label-text">Descuento</span>
								</div>
								<input
									class="input input-bordered w-full"
									type="number"
									min="0"
									step="1"
									v-model="form.discount"
								/>
							</label>
						</div>
						<div>
							<div class="col-span-3 p-4 border-white border">
								<OrderItems v-model="form.items" />
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
	</div>
</template>

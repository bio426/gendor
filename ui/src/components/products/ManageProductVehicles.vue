<script setup lang="ts">
import { computed, ref, watch } from "vue"

import vehiclesData from "@/assets/vehicles.json"
import productService from "@/services/product"
import Modal from "../Modal.vue"
import Overlay from "../Overlay.vue"
import type * as tProduct from "@/types/product"

const props = defineProps<{ show: boolean; vehicles: tProduct.Vehicle[] }>()
const emit = defineEmits<{ close: []; update: [tProduct.Vehicle[]] }>()

const active = ref(false)
watch(
	() => props.show,
	(v) => {
		if (v) {
			active.value = v
			getVehicleOpts()
		}
	}
)
function onClose() {
	active.value = false
	name.value = ""
	const selectedVehicles = vehicleOpts.value.filter((v) => v.check)
	emit("update", selectedVehicles)
	emit("close")
}

const loading = ref(false)

const name = ref("")

type CheckVehicle = tProduct.Vehicle & { check: boolean }
const vehicleOpts = ref<CheckVehicle[]>([])
async function getVehicleOpts() {
	loading.value = true
	const alreadyMarked = props.vehicles.map((v) => v.id)
	vehicleOpts.value = vehiclesData.map((v) => ({
		...v,
		check: alreadyMarked.includes(v.id),
	}))
	loading.value = false
}

function toogleOptCheck(id: number) {
	const v = vehicleOpts.value.find((v) => v.id == id)
	if (v == undefined) return
	v.check = !v.check
}

const onlySelected = ref(false)
const selectedCount = computed(
	() => vehicleOpts.value.filter((v) => v.check).length
)
const filteredVehicleOpts = computed(() => {
	// const res: CheckVehicle[] = [...vehicleOpts.value]
	const res: CheckVehicle[] = Array.from(vehicleOpts.value)
	if (onlySelected.value) {
		return res.filter((v) => v.check)
	}
	return res
})
</script>

<template>
	<Modal title="Product Vehicles" :show="active" @closing="onClose">
		<Overlay :show="loading">
			<div class="join mb-4">
				<div>
					<div>
						<input
							class="input input-bordered join-item"
							placeholder="Search"
						/>
					</div>
				</div>
				<div class="indicator">
					<button class="btn join-item">Search</button>
				</div>
			</div>
			<div class="flex justify-between items-center gap-4">
				<label class="label cursor-pointer">
					<span class="label-text mr-2">Only selected</span>
					<input
						type="checkbox"
						class="checkbox"
						v-model="onlySelected"
					/>
				</label>
				<small>{{ selectedCount }} vehicles selected</small>
			</div>
			<hr class="my-4" />
			<div class="overflow-x-auto max-h-[60vh]">
				<table class="table">
					<thead>
						<tr>
							<th></th>
							<th>Model</th>
							<th>Brand</th>
						</tr>
					</thead>
					<tbody>
						<tr
							v-for="vehicle in filteredVehicleOpts"
							:key="vehicle.id"
						>
							<td>
								<label>
									<input
										type="checkbox"
										class="checkbox"
										:value="vehicle.check"
										@change="toogleOptCheck(vehicle.id)"
									/>
								</label>
							</td>
							<td>{{ vehicle.model }}</td>
							<td>{{ vehicle.brand }}</td>
						</tr>
					</tbody>
				</table>
			</div>
		</Overlay>
	</Modal>
</template>

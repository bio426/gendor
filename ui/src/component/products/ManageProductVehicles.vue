<script setup lang="ts">
import { computed, ref, watch } from "vue"

import Modal from "../Modal.vue"
import Overlay from "../Overlay.vue"

const props = defineProps<{ show: boolean; vehicles: any[] }>()
const emit = defineEmits<{ close: []; update: [any[]] }>()

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

const vehicleOpts = ref<any[]>([])
async function getVehicleOpts() {
	loading.value = true
	const alreadyMarked = props.vehicles.map((v) => v.id)
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
					<tbody></tbody>
				</table>
			</div>
		</Overlay>
	</Modal>
</template>

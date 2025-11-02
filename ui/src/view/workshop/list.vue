<script setup lang="ts">
import { reactive, ref, watch } from "vue"
import { PrinterIcon } from "@heroicons/vue/24/solid"

import * as tWorkshop from "@/type/workshop"
import workshopService from "@/service/workshop"
import printUtil from "@/util/print"
import Header from "@/component/Header.vue"
import Overlay from "@/component/Overlay.vue"
import Pagination from "@/component/Pagination.vue"

const loading = ref(false)

const pagination = reactive({ total: 0, page: 1, from: 0, to: 0 })
watch(
	() => pagination.page,
	() => getRows()
)
const search = ref("")
const rows = ref<tWorkshop.OrderSimple[]>([])

async function getRows() {
	loading.value = true
	const res = await workshopService.list({
		search: search.value,
		page: pagination.page,
	})
	pagination.total = res.total
	pagination.from = res.from
	pagination.to = res.to
	rows.value = res.rows
	loading.value = false
}
getRows()

async function printOrderPdf(id: number) {
	loading.value = true
	const res = await workshopService.detail(id)
	console.log(res)
	const subtotal = res.items.reduce((acc, item) => item.price + acc, 0)
	const total = subtotal - res.discount
	const strItems = res.items.map((i) => ({
		description: i.description,
		price: i.price.toFixed(2),
	}))
	loading.value = false

	await printUtil.printWithPdf({
		id: res.id,
		propietary: res.propietary,
		date: new Date(res.createdAt).toLocaleDateString("en-GB"),
		brand: res.brand,
		model: res.model,
		year: res.year ? res.year.toString() : "",
		plate: res.plate,
		mileage: `${res.mileage} KM`,
		observation: res.observation,
		discount: res.discount.toFixed(2),
		subtotal: subtotal.toFixed(2),
		total: total.toFixed(2),
		items: strItems,
	})
}
</script>

<template>
	<div class="relative min-h-screen w-full">
		<div class="w-11/12 mx-auto">
			<Header title="Workshop" />
			<div class="mb-4">
				<router-link
					class="btn btn-primary"
					:to="{ name: 'workshop-create' }"
				>
					Crear Orden
				</router-link>
			</div>
			<div class="mb-4">
				<div class="join w-full">
					<input
						class="input input-bordered join-item flex-grow"
						placeholder="Buscar Placa"
						v-model="search"
						@keypress.enter="getRows"
					/>
					<button class="btn join-item" @click="getRows">
						Buscar
					</button>
				</div>
			</div>
			<Overlay :show="loading">
				<div class="mb-2 text-secondary text-sm">
					Mostrando {{ pagination.from }} a {{ pagination.to }} de
					{{ pagination.total }} resultados
				</div>
				<div class="overflow-x-auto mb-4">
					<table class="table">
						<thead>
							<tr>
								<th>Placa</th>
								<th>Propietario</th>
								<th>Creado</th>
								<th>Accion</th>
							</tr>
						</thead>
						<tbody>
							<tr v-for="row in rows" :key="row.id">
								<td>{{ row.plate }}</td>
								<td>{{ row.propietary }}</td>
								<td>
									{{
										new Date(
											row.createdAt
										).toLocaleDateString("en-GB")
									}}
								</td>
								<td>
									<div class="flex gap-2">
										<button
											class="btn btn-xs btn-square"
											title="Imprimir"
											@click="printOrderPdf(row.id)"
										>
											<PrinterIcon class="w-4 y-4" />
										</button>
									</div>
								</td>
							</tr>
							<tr v-if="rows.length == 0">
								<td class="text-center" colspan="4">
									No data available
								</td>
							</tr>
						</tbody>
					</table>
				</div>
				<div class="mb-8">
					<Pagination
						:total="pagination.total"
						:count="20"
						v-model="pagination.page"
					/>
				</div>
			</Overlay>
		</div>
	</div>
</template>

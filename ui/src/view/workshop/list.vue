<script setup lang="ts">
import { reactive, ref, watch } from "vue"
import { PrinterIcon, ClipboardIcon } from "@heroicons/vue/24/solid"

import * as tWorkshop from "@/type/workshop"
import workshopService from "@/service/workshop"
import useToast from "@/composable/useToast"
import printUtil from "@/util/print"
import Header from "@/component/Header.vue"
import Overlay from "@/component/Overlay.vue"
import Pagination from "@/component/Pagination.vue"

const toast = useToast()
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

async function printOrder(id: number) {
	loading.value = true
	const res = await workshopService.detail(id)
	const total = res.items.reduce(
		(acc, item) => item.price * item.quantity + acc,
		0
	)
	const strItems = res.items.map((i) => ({
		code: i.code,
		quantity: i.quantity.toString(),
		description: i.description,
		punit: i.price.toFixed(2),
		ptotal: (i.price * i.quantity).toFixed(2),
	}))

	const printableHtml = printUtil.printableString({
		id: `OT${res.id}`,
		name: res.name,
		address: res.address,
		dni: res.dni,
		ruc: res.ruc,
		date: new Date(res.createdAt).toLocaleDateString(),
		brand: res.brand,
		model: res.model,
		color: res.color,
		plate: res.plate,
		mileage: `${res.mileage} KM`,
		observation: res.observation,
		subtotal: total.toFixed(2),
		total: total.toFixed(2),
		items: strItems,
	})
	loading.value = false

	const win = window.open("", "PrintWindow")
	if (win == null) return
	win.document.write(printableHtml)
	win.document.close
	win.focus()
	win.print()
	win.close()
}

async function orderToClipboard(id: number) {
	loading.value = true
	const res = await workshopService.detail(id)
	const total = res.items.reduce(
		(acc, item) => item.price * item.quantity + acc,
		0
	)
	const strItems = res.items.map((i) => ({
		code: i.code,
		quantity: i.quantity.toString(),
		description: i.description,
		punit: i.price.toFixed(2),
		ptotal: (i.price * i.quantity).toFixed(2),
	}))
	const payload = {
		id: `OT${res.id}`,
		name: res.name,
		address: res.address,
		dni: res.dni,
		ruc: res.ruc,
		date: new Date(res.createdAt).toLocaleDateString(),
		brand: res.brand,
		model: res.model,
		color: res.color,
		plate: res.plate,
		mileage: res.mileage.toString(),
		observation: res.observation,
		subtotal: total.toFixed(2),
		total: total.toFixed(2),
		items: strItems,
	}
	const blob = await printUtil.testClipboard(payload)
	loading.value = false
	await navigator.clipboard.write([
		new ClipboardItem({
			[blob.type]: blob,
		}),
	])
	toast.display({ message: "Copied" })
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
								<th>Nombre</th>
								<th>Creado</th>
								<th>Accion</th>
							</tr>
						</thead>
						<tbody>
							<tr v-for="row in rows" :key="row.id">
								<td>{{ row.plate }}</td>
								<td>{{ row.name }}</td>
								<td>
									{{
										new Date(
											row.createdAt
										).toLocaleDateString()
									}}
								</td>
								<td>
									<div class="flex gap-2">
										<button
											class="btn btn-xs btn-square"
											title="Imprimir"
											@click="printOrder(row.id)"
										>
											<PrinterIcon class="w-4 y-4" />
										</button>
										<button
											class="btn btn-xs btn-square"
											title="Copiar"
											@click="orderToClipboard(row.id)"
										>
											<ClipboardIcon class="w-4 y-4" />
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

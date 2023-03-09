<script setup lang="ts">
import { ref, watch } from "vue"
import { MagnifyingGlassIcon, FunnelIcon } from "@heroicons/vue/24/solid"

import itemService from "../../services/product"
import Navigation from "../../components/Navigation.vue"
import Pagination from "../../components/Pagination.vue"
import Sidebar from "../../components/Sidebar.vue"
import IndexSC from "./_components/IndexSC.vue"
import { IProduct } from "../../interfaces/product"

const rows = ref<IProduct[]>([])

// pagination
const search = ref("")
const count = ref(10)
const page = ref(1)
const total = ref(0)
async function getRows() {
	const body = {
		page: page.value,
		count: count.value,
		search: search.value,
	}
	const data = await itemService.list(body)
	total.value = data.total
	rows.value = data.rows
}
watch(page, () => getRows())
getRows()

// test Sidebar
let showSidebar = ref(false)

// dinamic color
function isDarkText(hexcolor: string) {
	var r = parseInt(hexcolor.substring(1, 3), 16)
	var g = parseInt(hexcolor.substring(3, 5), 16)
	var b = parseInt(hexcolor.substring(5, 7), 16)
	var yiq = (r * 299 + g * 587 + b * 114) / 1000
	return yiq >= 128
}
</script>

<template>
	<Sidebar :show="showSidebar" @closing="showSidebar = false">
		<div class="relative w-full min-h-screen">
			<div class="w-11/12 mx-auto">
				<h1 class="py-4 text-2xl font-bold text-center">Items</h1>
				<form @submit.prevent="getRows">
					<div class="form-control mb-4">
						<div class="input-group">
							<input
								class="input input-bordered flex-grow"
								type="text"
								placeholder="Search…"
								v-model="search"
							/>
							<button
								class="btn btn-square"
								type="button"
								@click="showSidebar = true"
							>
								<FunnelIcon class="w-6 h-6" />
							</button>
							<button
								class="btn btn-square"
								type="submit"
								@click="getRows"
							>
								<MagnifyingGlassIcon class="w-6 h-6" />
							</button>
						</div>
					</div>
				</form>
				<div class="overflow-x-auto mb-4">
					<table class="table table-zebra w-full">
						<thead>
							<tr>
								<th>Name</th>
								<th>Price (S/.)</th>
								<th>Category</th>
							</tr>
						</thead>
						<tbody>
							<tr v-for="(row, i) in rows">
								<td>
									<router-link
										class="link"
										:to="{
											name: 'itemId',
											params: { id: row.name },
										}"
									>
										{{ row.name }}
									</router-link>
								</td>
								<td>{{ row.price }}</td>
								<td class="flex gap-2">
									<div
										class="badge font-medium"
										:class="[
											isDarkText(row.category.color)
												? 'text-neutral'
												: 'text-base-content',
										]"
										:style="{
											backgroundColor: row.category.color,
										}"
									>
										{{ row.category.name }}
									</div>
								</td>
							</tr>
							<tr class="text-center" v-if="rows.length == 0">
								<td colspan="4">No items</td>
							</tr>
						</tbody>
					</table>
				</div>
				<Pagination v-model="page" :total="total" :count="count" />
				<div class="w-full h-36"></div>
			</div>
			<div class="fixed left-0 bottom-16 w-full py-1 bg-base-100">
				<div class="grid grid-cols-1 gap-4 w-11/12 mx-auto">
					<router-link
						class="btn btn-sm btn-success btn-block"
						:to="{ name: 'itemCreate' }"
						>Create</router-link
					>
				</div>
			</div>
			<Navigation />
		</div>
		<template v-slot:content>
			<IndexSC />
		</template>
	</Sidebar>
</template>

<script setup lang="ts">
import { ref, watch } from "vue"
import { MagnifyingGlassIcon, FunnelIcon } from "@heroicons/vue/24/solid"

import itemService from "../../services/item"
import Navigation from "../../components/Navigation.vue"
import Pagination from "../../components/Pagination.vue"
import Sidebar from "../../components/Sidebar.vue"
import IndexSC from "./_components/IndexSC.vue"

const rows = ref<{
	name: string
	price: number
	tags: number[]
}[]>([])

// pagination
const search = ref("")
const count = ref(10)
const page = ref(1)
const total = ref(0)
async function getRows() {
	const body = {
		page: String(page.value),
		count: String(count.value),
		search: search.value,
	}
	const data = await itemService.list(body)
	total.value = data.total
	rows.value = data.rows
	// page.value = 1
}
watch(page, () => getRows())
getRows()

// test Sidebar
let showSidebar = ref(false)
</script>

<template>
	<Sidebar :show="showSidebar" @closing="showSidebar = false">
		<div class="relative w-full min-h-screen">
			<div class="w-11/12 mx-auto">
				<h1 class="py-4 text-2xl font-bold text-center">Items</h1>
				<form @submit.prevent="getRows">
					<div class="form-control mb-4">
						<div class="input-group">
							<input class="input input-bordered flex-grow" type="text" placeholder="Search…"
								v-model="search" />
							<button class="btn btn-square" type="button" @click="showSidebar = true">
								<FunnelIcon class="w-6 h-6" />
							</button>
							<button class="btn btn-square" type="submit" @click="getRows">
								<MagnifyingGlassIcon class="w-6 h-6" />
							</button>
						</div>
					</div>
				</form>
				<div class="overflow-x-auto mb-4">
					<table class="table table-zebra w-full">
						<thead>
							<tr>
								<th></th>
								<th>Name</th>
								<th>Price (S/.)</th>
								<th>Tags</th>
							</tr>
						</thead>
						<tbody>
							<tr v-for="(row, i) in rows">
								<th>{{ ((page - 1) * 10) + (i + 1) }}</th>
								<td>
									<router-link class="link" :to="{ name: 'itemId', params: { id: row.name } }">
										{{ row.name }}
									</router-link>
								</td>
								<td>{{ row.price }}</td>
								<td class="flex gap-2">
									<div class="badge" v-for="tag in row.tags">
										{{ tag }}
									</div>
								</td>
							</tr>
						</tbody>
					</table>
				</div>
				<Pagination v-model="page" :total="total" :count="count" />
				<div class="w-full h-36"></div>
			</div>
			<div class="fixed left-0 bottom-16 w-full py-1 bg-base-100">
				<div class="grid grid-cols-1 gap-4 w-11/12 mx-auto">
					<router-link class="btn btn-sm btn-success btn-block"
						:to="{ name: 'itemCreate' }">Create</router-link>
				</div>
			</div>
			<Navigation />
		</div>
		<template v-slot:content>
			<IndexSC />
		</template>
	</Sidebar>
</template>

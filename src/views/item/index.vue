<script setup lang="ts">
import { ref } from "vue"
import { MagnifyingGlassIcon } from "@heroicons/vue/24/solid"

import itemService from "../../services/item"
import Pagination from "../../components/Pagination.vue"

const rows = ref<{
	name: string
	price: number
	tags: number[]
}[]>([])

async function getRows() {
	const body = {
		page: String(1),
		count: String(10)
	}

	const res = await itemService.list(body)
	rows.value = await res.json()
}
getRows()

// pagination
const total = ref(30)
const page = ref(5)
</script>

<template>
	<div class="relative w-full min-h-screen">
		<div class="w-11/12 mx-auto">
			<h1 class="py-4 text-2xl font-bold text-center">Items</h1>
			<div class="form-control mb-4">
				<div class="input-group">
					<input class="input input-bordered flex-grow" type="text" placeholder="Search…" />
					<button class="btn btn-square">
						<MagnifyingGlassIcon class="w-6 h-6" />
					</button>
				</div>
			</div>
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
						<tr v-for="row in rows">
							<th>1</th>
							<td>{{ row.name }}</td>
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
			<Pagination v-model="page" :total="total" />
			<div class="w-full h-24"></div>
		</div>
		<div class="fixed left-0 bottom-0 w-full bg-base-100">
			<div class="grid grid-cols-2 gap-4 w-11/12 mx-auto">
				<router-link class="btn btn-success btn-block" :to="{ name: 'itemCreate' }">Create</router-link>
				<button class="btn btn-warning btn-block">Mass</button>
			</div>
		</div>
	</div>
</template>

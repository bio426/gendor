<script setup lang="ts">
import { ref, reactive } from "vue"

import productService from "@/service/product"
import Header from "@/component/Header.vue"
import Overlay from "@/component/Overlay.vue"
import Pagination from "@/component/Pagination.vue"

const loading = ref(false)
const rows = ref<any[]>([])

async function getRows() {
	loading.value = true
	const data = await productService.list({ page: 1, count: 1, search: "" })
	rows.value = data.rows
	loading.value = false
}
getRows()

const pagination = reactive({ total: 30, count: 2, page: 1 })
</script>

<template>
	<div class="relative w-screen min-h-screen">
		<div class="w-11/12 mx-auto">
			<Header title="Products" />
			<Overlay :show="loading">
				<pre>{{ pagination }}</pre>
				<Pagination
					:total="pagination.total"
					:count="pagination.count"
					v-model="pagination.page"
				/>
				<div class="overflow-x-auto">
					<table class="table">
						<thead>
							<tr>
								<th>Image</th>
								<th>Name</th>
								<th>Price</th>
							</tr>
						</thead>
						<tbody>
							<tr v-for="row in rows" :key="row.id">
								<td>
									<div class="avatar">
										<div class="w-12 h-12 rounded">
											<img :src="row.image" />
										</div>
									</div>
								</td>
								<td>{{ row.name }}</td>
								<td>{{ row.price }}</td>
							</tr>
							<tr v-if="rows.length == 0">
								<td colspan="3">No data available</td>
							</tr>
						</tbody>
					</table>
				</div>
			</Overlay>
		</div>
	</div>
</template>

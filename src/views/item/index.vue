<script setup lang="ts">
import { ref } from "vue"

import itemService from "../../services/item"

const rows = ref<{

}>([])

async function getRows() {
	const body = {
		page: String(1),
		count: String(10)
	}

	const res = await itemService.list(body)
	console.log(res)
}
getRows()
</script>

<template>
	<div class="relative w-full min-h-screen">
		<div class="w-11/12 mx-auto">
			<h1 class="py-4 text-2xl font-bold text-center">Create Product</h1>
			<div class="btn-group mb-4">
				<button class="btn">1</button>
				<button class="btn">2</button>
				<button class="btn btn-disabled">...</button>
				<button class="btn">99</button>
				<button class="btn">100</button>
			</div>
			<div class="overflow-x-auto">
				<table class="table table-zebra w-full">
					<thead>
						<tr>
							<th></th>
							<th>Name</th>
							<th>Job</th>
							<th>Favorite Color</th>
						</tr>
					</thead>
					<tbody>
						<tr>
							<th>1</th>
							<td>Cy Ganderton</td>
							<td>Quality Control Specialist</td>
							<td>Blue</td>
						</tr>
						<tr>
							<th>2</th>
							<td>Hart Hagerty</td>
							<td>Desktop Support Technician</td>
							<td>Purple</td>
						</tr>
						<tr>
							<th>3</th>
							<td>Brice Swyre</td>
							<td>Tax Accountant</td>
							<td>Red</td>
						</tr>
					</tbody>
				</table>
			</div>
		</div>
	</div>
</template>

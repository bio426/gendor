<script setup lang="ts">
import { ref } from "vue"
import { ArrowUturnLeftIcon, XCircleIcon } from "@heroicons/vue/24/solid"

const method = ref("")
const ruc = null
const rucE = null
const rows = [1, 2, 3, 4, 5]
</script>

<template>
	<div class="relative w-full min-h-screen">
		<div class="w-11/12 mx-auto">
			<h1 class="py-4 text-2xl font-bold text-center">
				<button class="inline float-left cursor-pointer" @click="$router.back()">
					<ArrowUturnLeftIcon class="w-6 mt-1" />
				</button>
				Finish
				<button class="inline float-right cursor-pointer" @click="$router.push({ name: 'pos' })">
					<XCircleIcon class="w-6 mt-1" />
				</button>
			</h1>
			<div class="mb-2">
				<span class="block mb-1 text-sm underline">Client Name</span>
				<span class="block">Some Client Name</span>
			</div>
			<div class="mb-2">
				<span class="block mb-1 text-sm underline">Client Ruc</span>
				<span class="block">123879817238</span>
			</div>
			<div class="grid grid-cols-2 gap-4 mb-4">
				<div class="mb-2">
					<span class="block mb-1 text-sm underline">Client Ruc</span>
					<span class="block">123879817238</span>
				</div>
				<div class="mb-2">
					<span class="block mb-1 text-sm underline">Client Ruc</span>
					<span class="block">123879817238</span>
				</div>
			</div>
			<div class="mb-4 overflow-y-auto h-[60vh]">
				<table class="table table-zebra w-full">
					<thead>
						<tr>
							<th>Info</th>
							<th>Quantity</th>
							<th>Total (S/.)</th>
						</tr>
					</thead>
					<tbody>
						<tr v-for="i in 10">
							<td>
								<span class="text-lg font-medium">ItemName</span>
								<br />
								<span class="text-sm">(S/. 99.99)</span>
							</td>
							<td>99</td>
							<td>
								99.99
							</td>
						</tr>
					</tbody>
				</table>
			</div>
			<div class="w-full h-24"></div>
		</div>
		<div class="fixed left-0 bottom-0 w-full py-1 bg-base-100 border border-white">
			<div class="grid grid-cols-1 gap-4 w-11/12 mx-auto">
				<button class="btn btn-sm btn-success btn-block">Save</button>
			</div>
		</div>
	</div>
</template>

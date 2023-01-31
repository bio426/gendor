<script setup lang="ts">
import { ref } from "vue"
import { ArrowUturnLeftIcon, XCircleIcon } from "@heroicons/vue/24/solid"

import FinishItemsModal from "./_components/FinishItemsModal.vue"

const method = ref("")
const ruc = null
const rucE = null
const rows = [1, 2, 3, 4, 5]

const showItemsModal = ref(false)
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
				<span class="block mb-1 text-sm font-bold">Client Name</span>
				<span class="block">Some Client Name</span>
			</div>
			<div class="mb-2">
				<span class="block mb-1 text-sm font-bold">Client Ruc</span>
				<span class="block">123879817238</span>
			</div>
			<div class="grid grid-cols-2 gap-4 mb-4">
				<div class="mb-2">
					<span class="block mb-1 text-sm font-bold">Client Ruc</span>
					<span class="block">123879817238</span>
				</div>
				<div class="mb-2">
					<span class="block mb-1 text-sm font-bold">Client Ruc</span>
					<span class="block">123879817238</span>
				</div>
			</div>
			<button class="btn btn-sm btn-block mb-4" @click="showItemsModal = true">View Items</button>
			<table class="w-full text-lg">
				<tr>
					<td>Subtotal</td>
					<td>:</td>
					<td class="text-right">99.99</td>
				</tr>
				<tr>
					<td>I.G.V</td>
					<td>:</td>
					<td class="text-right">99.99</td>
				</tr>
				<tr class="font-bold">
					<td>Total</td>
					<td>:</td>
					<td class="text-right">99.99</td>
				</tr>
			</table>
			<div class="w-full h-24"></div>
		</div>
		<div class="fixed left-0 bottom-0 w-full py-1 bg-base-100">
			<div class="grid grid-cols-1 gap-4 w-11/12 mx-auto">
				<button class="btn btn-sm btn-success btn-block">Save</button>
			</div>
		</div>
		<FinishItemsModal :show="showItemsModal" @closing="showItemsModal = false" />
	</div>
</template>

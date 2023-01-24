<script setup lang="ts">
import {} from "vue"
import { useRoute } from "vue-router"
import { ArrowUturnLeftIcon } from "@heroicons/vue/24/solid"

const route = useRoute()
</script>

<template>
	<div class="relative w-screen min-h-screen">
		<div class="w-11/12 mx-auto">
			<h1 class="py-4 text-2xl font-bold text-center">
				<a
					class="inline float-left cursor-pointer"
					@click="$router.back()"
				>
					<ArrowUturnLeftIcon class="w-6 h-6 mt-1" />
				</a>
				{{ route.params.id.toString().toUpperCase() }}
			</h1>
			<img
				class="block w-full mb-8"
				src="https://via.placeholder.com/400"
			/>
			<h4 class="font-semibold">Description:</h4>
			<p class="mb-4">We are using cookies for no reason.</p>
			<h4 class="font-semibold">Price:</h4>
			<span class="block mb-4 text-xl text-center">S/. 99.99</span>
			<h4 class="font-semibold">Specifications:</h4>
			<table class="w-full">
				<tr v-for="i in 4">
					<td>Key</td>
					<td>:</td>
					<td>Value {{ i }}</td>
				</tr>
			</table>
		</div>
	</div>
</template>

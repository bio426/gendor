<script setup lang="ts">
import { ref } from "vue"
import { ArrowUturnLeftIcon } from "@heroicons/vue/24/solid"

const method = ref("")
const ruc = null
const rucE = null
</script>

<template>
	<div class="relative w-full min-h-screen">
		<div class="w-11/12 mx-auto">
			<h1 class="py-4 text-2xl font-bold text-center">
				<a class="inline float-left cursor-pointer" @click="$router.back()">
					<ArrowUturnLeftIcon class="w-6 h-6 mt-1" />
				</a>
				Payment
			</h1>
			<div class="form-control w-full">
				<span class="label">
					<span class="label-text">Method</span>
				</span>
				<!---->
				<div class="form-control">
					<label class="label cursor-pointer">
						<span class="label-text">Cash</span>
						<input type="radio" name="method" class="radio" />
					</label>
				</div>
				<div class="form-control">
					<label class="label cursor-pointer">
						<span class="label-text">Bank transfer</span>
						<input type="radio" name="method" class="radio" />
					</label>
				</div>
				<div class="form-control">
					<label class="label cursor-pointer">
						<span class="label-text">Digital wallet</span>
						<input type="radio" name="method" class="radio" />
					</label>
				</div>
				<!---->
				<label class="label">
					<span class="label-text-alt" v-if="rucE">{{ rucE }}</span>
				</label>
			</div>
			<div class="form-control w-full">
				<label class="label">
					<span class="label-text">Ruc</span>
				</label>
				<input type="file" class="file-input w-full max-w-xs" :class="{ 'input-error': !!rucE }" />
				<label class="label">
					<span class="label-text-alt" v-if="rucE">{{ rucE }}</span>
				</label>
			</div>
			<div class="w-full h-24"></div>
		</div>
		<div class="fixed left-0 bottom-0 w-full py-1 bg-base-100">
			<div class="grid grid-cols-1 gap-4 w-11/12 mx-auto">
				<router-link class="btn btn-sm btn-success btn-block" :to="{ name: 'posCart' }">Finish</router-link>
			</div>
		</div>
	</div>
</template>

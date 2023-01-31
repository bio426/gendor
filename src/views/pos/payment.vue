<script setup lang="ts">
import { ref } from "vue"
import { ArrowUturnLeftIcon, XCircleIcon } from "@heroicons/vue/24/solid"

const method = ref("")
const ruc = null
const rucE = null
</script>

<template>
	<div class="relative w-full min-h-screen">
		<div class="w-11/12 mx-auto">
			<h1 class="py-4 text-2xl font-bold text-center">
				<button class="inline float-left cursor-pointer" @click="$router.back()">
					<ArrowUturnLeftIcon class="w-6 mt-1" />
				</button>
				Payment
				<button class="inline float-right cursor-pointer" @click="$router.push({ name: 'pos' })">
					<XCircleIcon class="w-6 mt-1" />
				</button>
			</h1>
			<div class="form-control w-full mb-4">
				<span class="label">
					<span class="label-text">Method</span>
				</span>
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
			</div>
			<div class="form-control w-full">
				<label class="label">
					<span class="label-text">Voucher</span>
				</label>
				<input type="file" class="file-input w-full" :class="{ 'input-error': !!rucE }" />
				<label class="label">
					<span class="label-text-alt" v-if="rucE">{{ rucE }}</span>
				</label>
			</div>
			<div class="form-control w-full">
				<label class="label">
					<span class="label-text">Pays With</span>
				</label>
				<input type="number" class="input input-bordered w-full" :class="{ 'input-error': !!rucE }" />
				<label class="label">
					<span class="label-text-alt" v-if="rucE">{{ rucE }}</span>
				</label>
			</div>
			<div class="form-control w-full">
				<label class="label">
					<span class="label-text">Transfer Code</span>
				</label>
				<input type="text" class="input input-bordered w-full" :class="{ 'input-error': !!rucE }" />
				<label class="label">
					<span class="label-text-alt" v-if="rucE">{{ rucE }}</span>
				</label>
			</div>
			<div class="w-full h-24"></div>
		</div>
		<div class="fixed left-0 bottom-0 w-full py-1 bg-base-100">
			<div class="grid grid-cols-1 gap-4 w-11/12 mx-auto">
				<router-link class="btn btn-sm btn-success btn-block" :to="{ name: 'posFinish' }">Finish</router-link>
			</div>
		</div>
	</div>
</template>

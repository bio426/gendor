<script setup lang="ts">
import { } from "vue"
import { ArrowUturnLeftIcon, XCircleIcon } from "@heroicons/vue/24/solid"

const name = ""
const nameE = ""
</script>

<template>
	<div class="relative w-full min-h-screen">
		<div class="w-11/12 mx-auto">
			<h1 class="py-4 text-2xl font-bold text-center">
				<button class="inline float-left cursor-pointer" @click="$router.back()">
					<ArrowUturnLeftIcon class="w-6 mt-1" />
				</button>
				Client
				<button class="inline float-right cursor-pointer" @click="$router.push({ name: 'pos' })">
					<XCircleIcon class="w-6 mt-1" />
				</button>
			</h1>
			<div class="form-control w-full">
				<label class="label">
					<span class="label-text">Name</span>
				</label>
				<input class="input input-bordered w-full" type="text" v-model="name"
					:class="{ 'input-error': !!nameE }" />
				<label class="label">
					<span class="label-text-alt" v-if="nameE">{{ nameE }}</span>
				</label>
			</div>
			<div class="form-control w-full">
				<label class="label">
					<span class="label-text">Ruc</span>
				</label>
				<input class="input input-bordered w-full" type="text" v-model="name"
					:class="{ 'input-error': !!nameE }" />
				<label class="label">
					<span class="label-text-alt" v-if="nameE">{{ nameE }}</span>
				</label>
			</div>
			<div class="form-control w-full">
				<label class="label">
					<span class="label-text">Email</span>
				</label>
				<input class="input input-bordered w-full" type="text" v-model="name"
					:class="{ 'input-error': !!nameE }" />
				<label class="label">
					<span class="label-text-alt" v-if="nameE">{{ nameE }}</span>
				</label>
			</div>
			<div class="w-full h-24"></div>
		</div>
		<div class="fixed left-0 bottom-0 w-full py-1 bg-base-100">
			<div class="grid grid-cols-1 gap-4 w-11/12 mx-auto">
				<router-link class="btn btn-sm btn-success btn-block" :to="{ name: 'posPayment' }">Payment</router-link>
			</div>
		</div>
	</div>
</template>

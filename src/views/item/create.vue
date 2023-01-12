<script setup lang="ts">
import { ref } from "vue"
import { useField } from "vee-validate"
import * as yup from "yup"

import itemService from "../../services/item"

const tagOpts = ref(["oil", "filter", "additive"])

const name = useField("", yup.string())

async function createProduct() {
	console.log(name.value)
	const body = {
		name: name.value.value,
		price: 0
	}
	// let res = await itemService.create(body)
}
</script>

<template>
	<div class="relative w-full min-h-screen">
		<div class="w-11/12 mx-auto">
			<h1 class="py-4 text-2xl font-bold text-center">Create Product</h1>
			<div class="form-control w-full">
				<label class="label">
					<span class="label-text">Name</span>
				</label>
				<input class="input input-bordered w-full" type="text" placeholder="ex: plastic bride"
					v-model="name.value" />
				<label class="label">
					<span class="label-text-alt">Error message</span>
				</label>
			</div>
			<div class="form-control w-full">
				<label class="label">
					<span class="label-text">Price</span>
				</label>
				<input class="input input-bordered w-full" type="text" placeholder="ex: 99.99" />
				<label class="label">
					<span class="label-text-alt">Error me</span>
				</label>
			</div>
			<div class="form-control w-full">
				<label class="label">
					<span class="label-text">Tags</span>
				</label>
				<select class="select select-error select-bordered w-full">
					<option hidden selected>Select one...</option>
					<option v-for="opt in tagOpts">Han Solo</option>
				</select>
				<label class="label">
					<span class="label-text-alt">Error me</span>
				</label>
			</div>
			<button class="btn btn-success btn-block mt-8" @click="createProduct">Create</button>
		</div>
	</div>
</template>

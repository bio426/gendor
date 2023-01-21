<script setup lang="ts">
import { ref } from "vue"
import { useDebounceFn } from "@vueuse/core"
import { ArrowUturnLeftIcon } from "@heroicons/vue/24/solid"

import usePosStore from "./_store"
import Navigation from "../../components/Navigation.vue"
import CartItem from "./_components/CartItem.vue"

const { cart } = usePosStore()

const itemOpts = ref<any[]>([])
const search = ref("")

async function getItemOpts() {

}
const handleSearch = useDebounceFn((ev: Event) => {
	console.log(ev)
	if (ev instanceof InputEvent) {
		// typing
		return
	}
	// select
}, 500)
</script>

<template>
	<div class="relative w-full min-h-screen">
		<div class="w-11/12 mx-auto">
			<h1 class="py-4 text-2xl font-bold text-center">
				<a class="inline float-left cursor-pointer" @click="$router.back()">
					<ArrowUturnLeftIcon class="w-6 mt-1" />
				</a>
				Cart
			</h1>
			<input class="input input-bordered w-full mb-8" type="text" placeholder="Search…" list="pos-search"
				v-model="search" @input="handleSearch" />
			<datalist id="pos-search">
				<option value="Chocolate" v-for="opt in itemOpts" />
				<option value="Chocolate" />
				<option value="Coconut" />
			</datalist>
			<div class="grid grid-cols-1 gap-4 max-h-[70vh] overflow-y-auto">
				<CartItem :item="item" v-for="item in cart" />
			</div>
			<div class="w-full h-36"></div>
		</div>
		<div class="fixed left-0 bottom-16 w-full py-1 bg-base-100">
			<div class="grid grid-cols-1 gap-4 w-11/12 mx-auto">
				<router-link class="btn btn-sm btn-success btn-block" :to="{ name: 'posClient' }">Client</router-link>
			</div>
		</div>
		<Navigation />
	</div>
</template>

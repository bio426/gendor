<script setup lang="ts">
import { ref } from "vue"
import { useDebounceFn } from "@vueuse/core"
import { ArrowUturnLeftIcon, XCircleIcon } from "@heroicons/vue/24/solid"

import type { IItem } from "./_types"
import usePosStore from "./_store"
import posService from "../../services/pos"
import CartItem from "./_components/CartItem.vue"

const { cart, addToCart, removeFromCart } = usePosStore()

const itemOpts = ref<IItem[]>([])
const search = ref("")
async function getItemOpts() {
	const query = {
		search: search.value,
	}
	const res = await posService.getItems(query)
	itemOpts.value = res
}
const handleSearch = useDebounceFn((ev: Event) => {
	if (ev instanceof InputEvent) {
		// typing
		getItemOpts()
		return
	}
	const match = itemOpts.value.find((item) => item.name == search.value)
	if (match == undefined) return
	search.value = ""
	itemOpts.value = []
	addToCart(match)
}, 500)
</script>

<template>
	<div class="relative w-full min-h-screen">
		<div class="w-11/12 mx-auto">
			<h1 class="py-4 text-2xl font-bold text-center">
				<button class="inline float-left cursor-pointer" @click="$router.back()">
					<ArrowUturnLeftIcon class="w-6 mt-1" />
				</button>
				Cart
				<button class="inline float-right cursor-pointer" @click="$router.push({ name: 'pos' })">
					<XCircleIcon class="w-6 mt-1" />
				</button>
			</h1>
			<input class="input input-bordered w-full mb-8" type="text" placeholder="Search…" list="pos-search"
				v-model="search" @input="handleSearch" />
			<datalist id="pos-search">
				<option :value="opt.name" :label="opt.id" v-for="opt in itemOpts" />
			</datalist>
			<div class="grid grid-cols-1 gap-4 max-h-[70vh] overflow-y-auto">
				<CartItem :item="item" v-for="(item, i) in cart" :key="item.id" @remove="removeFromCart(i)" />
			</div>
			<div class="w-full h-20"></div>
		</div>
		<div class="fixed left-0 bottom-0 w-full py-1 bg-base-100">
			<div class="grid grid-cols-1 gap-4 w-11/12 mx-auto">
				<router-link class="btn btn-sm btn-success btn-block" :to="{ name: 'posClient' }">Client</router-link>
			</div>
		</div>
	</div>
</template>

<script setup lang="ts">
import {} from "vue"
import { PlusIcon } from "@heroicons/vue/24/solid"

import type { IProduct } from "../_types"

const props = defineProps<{
	item: IProduct
}>()

const emit = defineEmits(["to-cart"])

const isCartItem = !!props.item.quantity
</script>

<template>
	<div class="flex gap-4 justify-between items-center">
		<div class="avatar">
			<div class="w-10 rounded-full">
				<img src="https://via.placeholder.com/400" />
			</div>
		</div>
		<div class="flex-grow">
			<span class="text-lg font-medium">{{ item.name }}</span>
			<br />
			<span
				>S/. {{ item.price }} <span class="text-sm">(99.99)</span></span
			>
		</div>
		<div class="flex gap-4">
			<button class="btn btn-sm" @click="$emit('to-cart')">
				<PlusIcon class="w-4" />
			</button>
		</div>
	</div>
</template>

<script setup lang="ts">
import { } from "vue"

import type { ICartItem } from "../_types"

const props = defineProps<{
	item: ICartItem
}>()

const emit = defineEmits(["to-cart"])

</script>

<template>
	<div class="flex gap-4 justify-between items-center">
		<div class="avatar indicator">
			<span class="indicator-item badge">{{ item.quantity }}</span>
			<div class="w-10 rounded-full">
				<img src="https://via.placeholder.com/400" />
			</div>
		</div>
		<div class="flex-grow">
			<span class="text-lg font-medium">{{ item.name }}</span>
			<br>
			<span>S/. {{ item.price * item.quantity }} <span class="text-sm">({{ item.price }})</span></span>
		</div>
		<div class="text-xl font-bold">
			S/. 99.99
		</div>
	</div>
</template>

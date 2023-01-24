<script setup lang="ts">
import {} from "vue"
import { TrashIcon } from "@heroicons/vue/24/solid"

import type { ICartItem } from "../_types"

const props = defineProps<{
	item: ICartItem
}>()

const emit = defineEmits(["remove"])
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
			<br />
			<span
				>S/. {{ item.price * item.quantity }}
				<span class="text-sm">({{ item.price }})</span></span
			>
		</div>
		<div class="flex gap-4">
			<button class="btn btn-error btn-sm" @click="$emit('remove')">
				<TrashIcon class="w-4" />
			</button>
		</div>
	</div>
</template>

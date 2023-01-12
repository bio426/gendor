<script setup lang="ts">
import { } from "vue"
import { MagnifyingGlassIcon } from "@heroicons/vue/24/solid"
</script>

<template>
	<div class="relative w-full min-h-screen">
		<div class="w-11/12 mx-auto">
			<h1 class="py-4 text-2xl font-bold text-center">POS</h1>
			<div class="flex gap-4">
				<input class="input input-bordered w-full" placeholder="Search Items..." list="pos-search" />
				<datalist id="pos-search">
					<option value="Chocolate" />
					<option value="Coconut" />
					<option value="Mint" />
					<option value="Strawberry" />
					<option value="Vanilla" />
				</datalist>
				<button class="btn">
					<MagnifyingGlassIcon class="h-4 w-4" />
				</button>
			</div>
		</div>
	</div>
</template>

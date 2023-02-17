<script setup lang="ts">
import { ref } from "vue"
import { AdjustmentsHorizontalIcon } from "@heroicons/vue/24/solid"

import Navigation from "../../components/Navigation.vue"
</script>

<template>
	<div class="relative w-full min-h-screen">
		<div class="w-11/12 mx-auto">
			<h1 class="py-4 text-2xl font-bold text-center">Admin</h1>
			<div class="grid grid-cols-1 gap-4">
				<div class="flex justify-between p-4 bg-base-200 rounded-lg">
					<span>Tags</span>
					<router-link :to="{ name: 'adminTag' }">
						<AdjustmentsHorizontalIcon class="w-6" />
					</router-link>
				</div>
			</div>
			<div class="w-full h-36"></div>
		</div>
		<Navigation />
	</div>
</template>

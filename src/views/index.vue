<script setup lang="ts">
import { } from "vue"

import useToast from "../composables/useToast"
import Navigation from "../components/Navigation.vue"

const { showToast } = useToast()

function handle() {
	showToast("asdasdasd", 2000)
}

</script>

<template>
	<div class="relative w-screen min-h-screen">
		<div class="w-11/12 mx-auto">
			<h1 class="py-4 text-2xl font-bold text-center">Dashboard</h1>
			<button class="btn" @click="handle">Show Toast</button>
		</div>
		<Navigation />
	</div>
</template>

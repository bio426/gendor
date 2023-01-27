<script setup lang="ts">
import { } from "vue"

import useToast from "../composables/useToast"
import useConfirmation from "../composables/useConfirmation"
import Navigation from "../components/Navigation.vue"

const { showToast } = useToast()
const { confirm } = useConfirmation()

function handle() {
	showToast("asdasdasd", 2000)
}

async function testConfirm() {
	const res = await confirm()
	console.log("gaaa", res)
}

</script>

<template>
	<div class="relative w-screen min-h-screen">
		<div class="w-11/12 mx-auto">
			<h1 class="py-4 text-2xl font-bold text-center">Dashboard</h1>
			<button class="btn" @click="handle">Show Toast</button>
			<button class="btn" @click="testConfirm">confirm</button>
		</div>
		<Navigation />
	</div>
</template>

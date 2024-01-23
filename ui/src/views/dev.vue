<script setup lang="ts">
import {} from "vue"

import useToast from "@/composables/useToast"
import useConfirmation from "@/composables/useConfirmation"
import Header from "@/components/Header.vue"

const toasts = useToast()
const confirmations = useConfirmation()

function testToast() {
	toasts.display({ message: "Klk", variant: "success" })
}

async function testConfirmation() {
	const res = await confirmations.display({})
	console.log(res)
}
</script>

<template>
	<div class="relative w-screen min-h-screen">
		<div class="w-11/12 mx-auto">
			<Header title="Dashboard" />
			<button class="btn" @click="testToast">Toast</button>
			<button class="btn" @click="testConfirmation">Confirmation</button>
		</div>
	</div>
</template>

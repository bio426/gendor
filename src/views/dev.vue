<script setup lang="ts">
import { ref } from "vue"

import useToast from "../composables/useToast"
import useConfirmation from "../composables/useConfirmation"
import Navigation from "../components/Navigation.vue"
import Overlay from "../components/Overlay.vue"

const { showToast } = useToast()
const { confirm } = useConfirmation()

function handle() {
	showToast("asdasdasd", 2000)
}

async function testConfirm() {
	const res = await confirm()
	console.log("gaaa", res)
}

const isLoading = ref(false)

</script>

<template>
	<div class="relative w-screen min-h-screen">
		<div class="w-11/12 mx-auto">
			<h1 class="py-4 text-2xl font-bold text-center">Development Page</h1>
			<button class="btn" @click="handle">Show Toast</button>
			<button class="btn" @click="testConfirm">confirm</button>
			<button class="btn" @click="isLoading = !isLoading">change loading</button>
			<Overlay :loading="isLoading">
				<p>Lorem ipsum dolor sit amet consectetur, adipisicing elit. Aliquam perspiciatis amet illum, quos omnis
					animi dignissimos fuga architecto voluptas porro voluptate eum assumenda non nam, exercitationem
					delectus impedit modi ipsam.</p>
			</Overlay>
		</div>
		<Navigation />
	</div>
</template>

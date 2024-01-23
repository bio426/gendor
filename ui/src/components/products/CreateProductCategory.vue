<script setup lang="ts">
import { ref, watch } from "vue"

import useConfirmation from "@/composables/useConfirmation"
import productService from "@/services/product"
import Modal from "../Modal.vue"
import Overlay from "../Overlay.vue"

const props = defineProps<{ show: boolean }>()
const emit = defineEmits<{ close: []; refresh: [] }>()

const active = ref(false)
watch(
	() => props.show,
	(v) => {
		if (v) {
			active.value = v
		}
	}
)
function onClose() {
	active.value = false
	name.value = ""
	emit("close")
}

const loading = ref(false)

const confirmations = useConfirmation()
const name = ref("")
async function create() {
	if (!(await confirmations.display({}))) return
	loading.value = true
	await productService.createProductCategory({ name: name.value })
	loading.value = false
	emit("refresh")
	onClose()
}
</script>

<template>
	<Modal title="New Product Category" :show="active" @closing="onClose">
		<Overlay :show="loading">
			<label class="form-control w-full">
				<div class="label">
					<span class="label-text">Name</span>
				</div>
				<input
					class="input input-bordered w-full"
					type="text"
					v-model="name"
				/>
			</label>
			<button class="btn btn-block mt-8" @click="create">Create</button>
		</Overlay>
	</Modal>
</template>

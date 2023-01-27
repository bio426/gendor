<script setup lang="ts">
import { ref, watch } from "vue"
import { XMarkIcon } from "@heroicons/vue/24/solid"

const props = defineProps<{
	show: boolean
	title: string
}>()

const emit = defineEmits(["closing"])

const isActive = ref(false)

watch(() => props.show, val => {
	isActive.value = val
})

watch(isActive, val => {
	if (!val) {
		emit("closing")
	}
})

const id = Math.random().toString(36).substring(2, 7);
const modalId = "cmodal-" + id
</script>

<template>
	<Teleport to="#modals">
		<input type="checkbox" :id="modalId" class="modal-toggle" v-model="isActive" />
		<label :for="modalId" class="modal cursor-pointer">
			<label class="modal-box relative" for="">
				<label :for="modalId" class="btn btn-sm btn-circle absolute right-2 top-2">
					<XMarkIcon class="w-6" />
				</label>
				<h3 class="text-lg font-bold mb-4">{{ title }}</h3>
				<slot />
			</label>
		</label>
	</Teleport>
</template>

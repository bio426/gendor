<script setup lang="ts">
import { ref, watch } from "vue"

const props = defineProps({
	show: {
		type: Boolean,
		required: true,
	},
})
const emit = defineEmits(["closing"])

const isActive = ref(false)

watch(
	() => props.show,
	(val) => {
		if (val) {
			isActive.value = true
		}
	}
)

watch(isActive, (val) => {
	if (!val) {
		emit("closing")
	}
})
</script>

<template>
	<div class="drawer">
		<input
			id="my-drawer"
			type="checkbox"
			class="drawer-toggle"
			v-model="isActive"
		/>
		<div class="drawer-content">
			<slot />
		</div>
		<div class="drawer-side">
			<label for="my-drawer" class="drawer-overlay"></label>
			<ul class="menu p-4 w-80 bg-base-100 text-base-content">
				<slot name="content" />
			</ul>
		</div>
	</div>
</template>

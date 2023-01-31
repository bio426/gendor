<script setup lang="ts">
import { } from "vue"
import { HandRaisedIcon } from "@heroicons/vue/24/solid"

import useToast from "../composables/useToast"

const { toasts } = useToast()
</script>

<template>
	<Teleport to="#toasts">
		<div class="toast toast-top toast-end">
			<TransitionGroup name="list" tag="ul">
				<div class="alert alert-info mb-2" v-for="toast in toasts" :key="toast.key">
					<div>
						<label class="btn btn-circle btn-xs mr-1" v-if="false">
							<HandRaisedIcon class="w-4" />
						</label>
						<span>{{ toast.text }}</span>
					</div>
				</div>
			</TransitionGroup>
		</div>
	</Teleport>
</template>

<style>
.list-move,
.list-enter-active,
.list-leave-active {
	transition: all 0.2s ease;
}

.list-enter-from,
.list-leave-to {
	opacity: 0;
	transform: translateX(2rem);
}

.list-leave-active {
	position: relative;
}
</style>

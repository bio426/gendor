<script setup lang="ts">
import { ref } from "vue"

import useConfirmation from "../composables/useConfirmation"

const { isActive, acceptTrigger, rejectTrigger } = useConfirmation()
const emit = defineEmits(["closing"])

function handleAccept() {
	acceptTrigger.value = Symbol()
}
function handleReject() {
	rejectTrigger.value = Symbol()
}
</script>

<template>
	<Teleport to="#modals">
		<input type="checkbox" id="confirmation-modal" class="modal-toggle" v-model="isActive" />
		<div class="modal">
			<div class="modal-box relative">
				<div>
					modal content
					<div class="h-12"></div>
					<div class="grid grid-cols-2 gap-4">
						<button class="btn btn-success" @click="handleAccept">Accept</button>
						<button class="btn btn-error" @click="handleReject">Reject</button>
					</div>
				</div>
			</div>
		</div>
	</Teleport>
</template>

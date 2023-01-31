<script setup lang="ts">
import { } from "vue"

import Modal from "../../../components/Modal.vue"

const props = defineProps<{
	show: boolean
}>()
const emit = defineEmits(["closing"])
function handleClose() {
	emit("closing")
}
</script>

<template>
	<Modal title="Items" :show="show" @closing="handleClose">
		<div class="mb-4 overflow-y-auto h-[60vh]">
			<table class="table table-zebra w-full">
				<thead>
					<tr>
						<th>Info</th>
						<th>Quantity</th>
						<th>Total (S/.)</th>
					</tr>
				</thead>
				<tbody>
					<tr v-for="i in 10">
						<td>
							<span class="text-lg font-medium">ItemName</span>
							<br />
							<span class="text-sm">(S/. 99.99)</span>
						</td>
						<td>99</td>
						<td>
							99.99
						</td>
					</tr>
				</tbody>
			</table>
		</div>
	</Modal>
</template>

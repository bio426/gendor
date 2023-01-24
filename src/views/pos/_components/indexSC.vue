<script setup lang="ts">
import { ref } from "vue"

const emit = defineEmits<{
	(
		event: "filter",
		info: {
			tag: string
		}
	): void
}>()
</script>

<template>
	<li>
		<div class="form-control">
			<label class="input-group input-group-vertical">
				<span>Tag</span>
				<select class="select select-bordered">
					<option disabled selected>Select a tag...</option>
					<option>Star Wars</option>
					<option>Harry Potter</option>
				</select>
			</label>
		</div>
	</li>
	<div class="w-full h-8"></div>
	<div class="absolute left-0 bottom-0 w-full py-1 bg-base-100">
		<div class="grid grid-cols-2 gap-4 w-11/12 mx-auto">
			<button class="btn btn-sm btn-block">Clear</button>
			<button class="btn btn-sm btn-success btn-block">Done</button>
		</div>
	</div>
</template>

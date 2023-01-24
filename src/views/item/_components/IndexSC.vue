<script setup lang="ts">
import { ref } from "vue"

const emit = defineEmits(["input"])
</script>

<template>
	<li>
		<div class="form-control">
			<label class="input-group input-group-vertical">
				<span>Email</span>
				<input
					type="text"
					placeholder="info@site.com"
					class="input input-bordered"
				/>
			</label>
		</div>
	</li>
	<li>
		<div class="form-control">
			<label class="input-group input-group-vertical">
				<span>Email</span>
				<select class="select select-bordered">
					<option disabled selected>Pick one</option>
					<option>Star Wars</option>
					<option>Harry Potter</option>
					<option>Lord of the Rings</option>
					<option>Planet of the Apes</option>
					<option>Star Trek</option>
				</select>
			</label>
		</div>
	</li>
	<li>
		<select class="select select-bordered">
			<option disabled selected>Pick one</option>
			<option>Star Wars</option>
			<option>Harry Potter</option>
			<option>Lord of the Rings</option>
			<option>Planet of the Apes</option>
			<option>Star Trek</option>
		</select>
	</li>
	<li>
		<input
			type="text"
			placeholder="info@site.com"
			class="input input-bordered"
		/>
	</li>
</template>

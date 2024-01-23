<script setup lang="ts">
import { reactive } from "vue"

const form = reactive({ username: "", password: "" })

async function login() {
	console.log(form)
}
</script>

<template>
	<div class="relative w-screen min-h-screen">
		<div class="w-11/12 mx-auto">
			<div class="navbar-center">
				<h1 class="py-4 text-2xl font-bold text-center">Login</h1>
			</div>
			<form class="" @submit.prevent="login">
				<label class="form-control w-full mb-4">
					<div class="label">
						<span class="label-text">Username</span>
					</div>
					<input
						class="input input-bordered w-full"
						type="text"
						required
						v-model="form.username"
					/>
				</label>
				<label class="form-control w-full mb-8">
					<div class="label">
						<span class="label-text">Password</span>
					</div>
					<input
						class="input input-bordered w-full"
						type="text"
						required
						v-model="form.password"
					/>
				</label>
				<input type="submit" class="btn btn-block" value="Access" />
			</form>
		</div>
	</div>
</template>

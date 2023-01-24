<script setup lang="ts">
import { } from "vue"
import { useRouter } from "vue-router"

import Toasts from "./components/Toast.vue"

const router = useRouter()
</script>

<template>
	<div>
		<router-view />
		<Toasts />
	</div>
</template>

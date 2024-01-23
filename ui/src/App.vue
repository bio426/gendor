<script setup lang="ts">
import {} from "vue"

import Toast from "@/components/Toast.vue"
import Confirmation from "@/components/Confirmation.vue"
</script>

<template>
	<div>
		<router-view />
		<Toast />
		<Confirmation />
	</div>
</template>

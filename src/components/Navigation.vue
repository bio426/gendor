<script setup lang="ts">
import { } from "vue"
import { useRouter, useRoute } from "vue-router"
import {
	HomeIcon,
	BanknotesIcon,
	CubeIcon,
	Cog6ToothIcon,
	ClipboardIcon,
} from "@heroicons/vue/24/solid"

const router = useRouter()
const route = useRoute()
</script>

<template>
	<div class="btm-nav">
		<router-link active-class="active" :to="{ name: 'admin' }">
			<Cog6ToothIcon class="w-6" />
			<span class="btm-nav-label">Admin</span>
		</router-link>
		<router-link active-class="active" :to="{ name: 'itemList' }">
			<CubeIcon class="w-6" />
			<span class="btm-nav-label">Items</span>
		</router-link>
		<router-link active-class="active" :to="{ name: 'index' }">
			<HomeIcon class="w-6" />
			<span class="btm-nav-label">Home</span>
		</router-link>
		<router-link active-class="active" :to="{ name: 'pos' }">
			<BanknotesIcon class="w-6" />
			<span class="btm-nav-label">Pos</span>
		</router-link>
		<router-link active-class="active" :to="{ name: 'sale' }">
			<ClipboardIcon class="w-6" />
			<span class="btm-nav-label">Sales</span>
		</router-link>
	</div>
</template>

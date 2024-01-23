<script setup lang="ts">
import {} from "vue"

import {
	Bars3Icon,
	HomeIcon,
	CubeIcon,
	BanknotesIcon,
	ClipboardIcon,
	EllipsisHorizontalIcon,
} from "@heroicons/vue/24/solid"

const props = defineProps({
	title: {
		type: String,
		default: "",
	},
})
</script>

<template>
	<div class="navbar bg-base-100">
		<div class="navbar-start">
			<div class="dropdown">
				<label tabindex="0" class="btn btn-ghost btn-circle">
					<Bars3Icon class="w-6" />
				</label>
				<ul
					tabindex="0"
					class="menu menu-sm dropdown-content mt-3 z-[1] p-2 shadow bg-base-100 rounded-box w-52"
				>
					<li>
						<router-link
							active-class="active"
							:to="{ name: 'index' }"
						>
							<HomeIcon class="w-4" />
							Home
						</router-link>
					</li>
					<!-- <li> -->
					<!-- 	<router-link -->
					<!-- 		active-class="active" -->
					<!-- 		:to="{ name: 'productList' }" -->
					<!-- 	> -->
					<!-- 		<CubeIcon class="w-4" /> -->
					<!-- 		Products -->
					<!-- 	</router-link> -->
					<!-- </li> -->
					<!-- <li> -->
					<!-- 	<router-link -->
					<!-- 		active-class="active" -->
					<!-- 		:to="{ name: 'pos' }" -->
					<!-- 	> -->
					<!-- 		<BanknotesIcon class="w-4" /> -->
					<!-- 		Pos -->
					<!-- 	</router-link> -->
					<!-- </li> -->
					<li>
						<details open>
							<summary>
								<CubeIcon class="w-4" />
								Products
							</summary>
							<ul>
								<li>
									<router-link
										active-class="active"
										:to="{ name: 'products-list' }"
									>
										List
									</router-link>
								</li>
								<li>
									<router-link
										active-class="active"
										:to="{ name: 'products-create' }"
									>
										Create
									</router-link>
								</li>
							</ul>
						</details>
					</li>
				</ul>
			</div>
		</div>
		<div class="navbar-center">
			<h1 class="py-4 text-2xl font-bold text-center">
				{{ props.title }}
			</h1>
		</div>
		<div class="navbar-end">
			<button class="btn btn-ghost btn-circle" v-if="false">
				<EllipsisHorizontalIcon class="w-6" />
			</button>
		</div>
	</div>
</template>

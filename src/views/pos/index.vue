<script setup lang="ts">
import { ref } from "vue"
import {
	MagnifyingGlassIcon,
	FunnelIcon,
	Square3Stack3DIcon,
	CurrencyDollarIcon,
} from "@heroicons/vue/24/solid"

import type { IItem } from "./_types"
import useToast from "../../composables/useToast"
import posService from "../../services/pos"
import usePosStore from "./_store"
import Navigation from "../../components/Navigation.vue"
import Sidebar from "../../components/Sidebar.vue"
import IndexSC from "./_components/indexSC.vue"
import Item from "./_components/Item.vue"

const { showToast } = useToast()
const { cart, cartPrice, cartLength, addToCart } = usePosStore()

const search = ref("")
const showSidebar = ref(false)

const items = ref<IItem[]>([])

async function getItems() {
	const query = {
		search: search.value,
	}
	const res = await posService.getItems(query)
	items.value = res
}
getItems()

function toCart(item: IItem) {
	addToCart(item)
	showToast(`${item.name} added to cart`, 2000)
}
</script>

<template>
	<Sidebar :show="showSidebar" @closing="showSidebar = false">
		<div class="relative w-full min-h-screen">
			<div class="w-11/12 mx-auto">
				<h1 class="py-4 text-2xl font-bold text-center">POS</h1>
				<form @submit.prevent="getItems">
					<div class="form-control mb-8">
						<div class="input-group">
							<input
								class="input input-bordered flex-grow"
								type="text"
								placeholder="Search…"
								v-model="search"
							/>
							<button
								class="btn btn-square"
								type="button"
								@click="showSidebar = true"
							>
								<FunnelIcon class="w-6" />
							</button>
							<button
								class="btn btn-square"
								type="submit"
								@click="getItems"
							>
								<MagnifyingGlassIcon class="w-6" />
							</button>
						</div>
					</div>
				</form>
				<div
					class="grid grid-cols-1 gap-4 max-h-[70vh] overflow-y-auto"
				>
					<item
						:item="item"
						v-for="item in items"
						@to-cart="toCart(item)"
					/>
				</div>
				<div class="w-full h-24"></div>
			</div>
			<div class="fixed left-0 bottom-16 w-full py-1 bg-base-100">
				<div class="grid grid-cols-2 gap-4 w-11/12 mx-auto">
					<div class="flex justify-center gap-4">
						<div class="badge badge-info gap-2">
							<Square3Stack3DIcon class="w-4" />
							{{ cartLength }}
						</div>
						<div class="badge badge-info gap-2">
							<CurrencyDollarIcon class="w-4" />
							{{ cartPrice.toFixed(2) }}
						</div>
					</div>
					<router-link
						class="btn btn-sm btn-success btn-block"
						:to="{ name: 'posCart' }"
						>Cart</router-link
					>
				</div>
			</div>
			<Navigation />
		</div>
		<template v-slot:content>
			<IndexSC />
		</template>
	</Sidebar>
</template>

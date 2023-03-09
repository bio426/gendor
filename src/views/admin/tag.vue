<script setup lang="ts">
import { ref } from "vue"
import { ArrowUturnLeftIcon, TrashIcon } from "@heroicons/vue/24/solid"

import adminService from "../../services/admin"
import useToast from "../../composables/useToast"
import CreateTagModal from "./_components/CreateTagModal.vue"
import { ICategory } from "../../interfaces/admin"

const { showToast } = useToast()

const rows = ref<ICategory[]>([])

async function getRows() {
	const data = await adminService.getTags()
	rows.value = data
}
getRows()

async function deleteTag(id: string) {
	await adminService.deleteTag(id)
	getRows()
}

const showCreateTag = ref(false)
</script>

<template>
	<div class="relative w-full min-h-screen">
		<div class="w-11/12 mx-auto">
			<h1 class="py-4 text-2xl font-bold text-center">
				<a
					class="inline float-left cursor-pointer"
					@click="$router.back()"
				>
					<ArrowUturnLeftIcon class="w-6 h-6 mt-1" />
				</a>
				Tags
			</h1>
			<div class="overflow-x-auto mb-4">
				<table class="table table-zebra w-full">
					<thead>
						<tr>
							<th>Name</th>
							<th>Color</th>
							<th>Actions</th>
						</tr>
					</thead>
					<tbody>
						<tr v-for="(row, i) in rows">
							<td>
								{{ row.name }}
							</td>
							<td>
								<span
									class="inline-block w-4 h-4 mx-auto rounded-full bg-white"
									:style="{ 'background-color': row.color }"
								></span>
							</td>
							<td>
								<button
									class="btn btn-sm btn-error"
									@click="deleteTag(row.id)"
									disabled
								>
									<TrashIcon class="w-4" />
								</button>
							</td>
						</tr>
					</tbody>
				</table>
			</div>
			<div class="w-full h-24"></div>
		</div>
		<div class="fixed left-0 bottom-0 w-full py-1 bg-base-100">
			<div class="grid grid-cols-1 gap-4 w-11/12 mx-auto">
				<button
					class="btn btn-sm btn-success btn-block"
					@click="showCreateTag = true"
				>
					Create
				</button>
			</div>
		</div>
		<CreateTagModal
			:show="showCreateTag"
			@closing=";(showCreateTag = false), getRows()"
		/>
	</div>
</template>

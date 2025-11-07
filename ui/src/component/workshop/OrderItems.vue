<script setup lang="ts">
import { computed, reactive, ref } from "vue"
import { TrashIcon } from "@heroicons/vue/24/solid"

import * as tWorkshop from "@/type/workshop"

const modelValue = defineModel<tWorkshop.OrderItem[]>({ required: true })

const form = reactive({ description: "", price: 0 })
function addItem() {
	items.value.push({
		description: form.description.trim(),
		price: form.price,
	})
	clearForm()
	save()
}

function clearForm() {
	form.price = 0
	form.description = ""
}

function removeItem(idx: number) {
	items.value.splice(idx, 1)
	save()
}

const items = ref<tWorkshop.OrderItem[]>([])
const itemsPrice = computed(() =>
	items.value.reduce((acc, item) => acc + item.price, 0)
)
async function save() {
	modelValue.value = [...items.value]
}
</script>

<template>
	<div>
		<form @submit.prevent="addItem" autocomplete="off">
			<div class="grid grid-cols-1 gap-4">
				<label class="form-control w-full">
					<div class="label">
						<span class="label-text">Descripcion</span>
					</div>
					<textarea
						class="textarea textarea-bordered w-full"
						type="text"
						required
						v-model="form.description"
					/>
				</label>
				<label class="form-control w-full">
					<div class="label">
						<span class="label-text">Precio</span>
					</div>
					<input
						class="input input-bordered w-full"
						type="number"
						step=".01"
						min="0.01"
						v-model="form.price"
					/>
				</label>
			</div>
			<button class="btn btn-block btn-primary mt-8" type="submit">
				Agregar
			</button>
		</form>
		<hr class="my-4" />
		<div class="overflow-auto">
			<table class="table">
				<thead>
					<tr>
						<th>Descripcion</th>
						<th>Precio</th>
						<th>Accion</th>
					</tr>
				</thead>
				<tbody>
					<tr v-for="(item, idx) in items">
						<td class="whitespace-normal w-fit">
							{{ item.description }}
						</td>
						<td>{{ item.price.toFixed(2) }}</td>
						<td>
							<button
								class="btn btn-sm btn-square"
								title="Quitar"
								@click="removeItem(idx)"
							>
								<TrashIcon class="w-4 fill-error" />
							</button>
						</td>
					</tr>
					<tr v-if="items.length == 0">
						<td class="text-center" colspan="5">
							Aun no hay items agregados
						</td>
					</tr>
				</tbody>
				<tfoot>
					<tr>
						<th colspan="1"></th>
						<th title="Total price">{{ itemsPrice.toFixed(2) }}</th>
					</tr>
				</tfoot>
			</table>
		</div>
	</div>
</template>

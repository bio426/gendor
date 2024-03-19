<script setup lang="ts">
import { computed, reactive, ref, watch } from "vue"
import { TrashIcon } from "@heroicons/vue/24/solid"

import * as tWorkshop from "@/type/workshop"
import Modal from "../Modal.vue"

const props = defineProps<{ show: boolean }>()
const emit = defineEmits<{ closing: [] }>()
const modelValue = defineModel<tWorkshop.OrderItem[]>({ required: true })

watch(
	() => props.show,
	(v) => {
		if (v) {
			items.value = [...modelValue.value]
		}
	}
)
function handleClose() {
	clearForm()
	items.value = []
	emit("closing")
}

const form = reactive({ code: "", quantity: 1, price: 0, description: "" })
function addVariant() {
	items.value.push({
		code: form.code,
		quantity: form.quantity,
		price: form.price,
		description: form.description,
	})
	clearForm()
}

function clearForm() {
	form.code = ""
	form.quantity = 1
	form.price = 0
	form.description = ""
}

function removeItem(idx: number) {
	items.value.splice(idx, 1)
}

const items = ref<tWorkshop.OrderItem[]>([])
const itemsPrice = computed(() =>
	items.value.reduce((acc, item) => acc + item.price * item.quantity, 0)
)
async function save() {
	modelValue.value = [...items.value]
	handleClose()
}
</script>

<template>
	<Modal
		title="Cambiar items de la orden"
		:show="props.show"
		@closing="handleClose"
	>
		<form @submit.prevent="addVariant" autocomplete="off">
			<div class="grid grid-cols-2 md:grid-cols-3 gap-4">
				<label class="form-control w-full col-span-2 md:col-span-1">
					<div class="label">
						<span class="label-text">Codigo</span>
					</div>
					<input
						class="input input-bordered w-full"
						type="text"
						v-model="form.code"
					/>
				</label>
				<label class="form-control w-full">
					<div class="label">
						<span class="label-text">Cantidad</span>
					</div>
					<input
						class="input input-bordered w-full"
						type="number"
						min="1"
						v-model="form.quantity"
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
				<label class="form-control w-full col-span-2 md:col-span-3">
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
						<th>Codigo</th>
						<th>Descripcion</th>
						<th>Cantidad</th>
						<th>Precio</th>
						<th>Accion</th>
					</tr>
				</thead>
				<tbody>
					<tr v-for="(item, idx) in items">
						<td>{{ item.code }}</td>
						<td class="whitespace-normal w-fit">
							{{ item.description }}
						</td>
						<td>{{ item.quantity }}</td>
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
						<th colspan="3"></th>
						<th title="Total price">{{ itemsPrice.toFixed(2) }}</th>
					</tr>
				</tfoot>
			</table>
		</div>
		<button class="btn btn-block btn-primary mt-8" @click="save">
			Guardar
		</button>
	</Modal>
</template>

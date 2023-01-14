<script setup lang="ts">
import { ref, computed } from "vue"

const props = defineProps({
	total: {
		type: Number,
		required: true,
	},
	modelValue: {
		type: Number,
		required: true,
	}
})

const emits = defineEmits(["update:modelValue"])

function changeValue(val: number) {
	emits("update:modelValue", val)
}

const isComplex = computed(() => props.total > 8)
const inversed = computed(() => {
	let arr: number[] = []
	for (let i = props.total; i > props.total - 5; i--) {
		arr.push(i)
	}
	return arr.reverse()
})
</script>

<template>
	<div class="text-center">
		<div class="btn-group">
			<template v-if="!isComplex">
				<button class="btn" :class="{ 'btn-active': i == modelValue }" v-for="i in total"
					@click="changeValue(i)">{{
						i
					}}</button>
			</template>
			<template v-else-if="isComplex && modelValue < 5">
				<button class="btn" :class="{ 'btn-active': i == modelValue }" v-for="i in 5" @click="changeValue(i)">{{
					i
				}}</button>
				<button class="btn btn-disabled">...</button>
				<button class="btn" @click="changeValue(total)">{{ total }}</button>
			</template>
			<template v-else-if="isComplex && modelValue > total - 4">
				<button class="btn" @click="changeValue(1)">1</button>
				<button class="btn btn-disabled">...</button>
				<button class="btn" :class="{ 'btn-active': i == modelValue }" v-for="i in inversed"
					@click="changeValue(i)">{{ i }}</button>
			</template>
			<template v-else>
				<button class="btn" @click="changeValue(1)">1</button>
				<button class="btn btn-disabled">...</button>
				<button class="btn" @click="changeValue(modelValue - 1)">{{ modelValue - 1}}</button>
				<button class="btn btn-active">{{ modelValue }}</button>
				<button class="btn" @click="changeValue(modelValue + 1)">{{ modelValue + 1}}</button>
				<button class="btn btn-disabled">...</button>
				<button class="btn" @click="changeValue(total)">{{ total }}</button>
			</template>
		</div>
	</div>
</template>

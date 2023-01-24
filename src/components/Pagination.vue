<script setup lang="ts">
import { ref, computed, watch } from "vue"

const props = defineProps({
	total: {
		type: Number,
		required: true,
	},
	count: {
		type: Number,
		required: true,
	},
	modelValue: {
		type: Number,
		required: true,
	},
})

const emits = defineEmits(["update:modelValue"])

function toPage(val: number) {
	emits("update:modelValue", val)
}
const pageCount = computed(() => {
	return Math.ceil(props.total / props.count)
})
const isComplex = computed(() => pageCount.value > 8)
const inversed = computed(() => {
	let arr: number[] = []
	for (let i = pageCount.value; i > pageCount.value - 5; i--) {
		arr.push(i)
	}
	return arr.reverse()
})

watch(
	() => props.total,
	(val) => {
		if (props.modelValue * props.count + 1 > props.total) toPage(1)
	}
)
</script>

<template>
	<div class="text-center">
		<div class="btn-group">
			<template v-if="!isComplex">
				<button
					class="btn"
					:class="{ 'btn-active': i == modelValue }"
					v-for="i in pageCount"
					@click="toPage(i)"
				>
					{{ i }}
				</button>
			</template>
			<template v-else-if="isComplex && modelValue < 5">
				<button
					class="btn"
					:class="{ 'btn-active': i == modelValue }"
					v-for="i in 5"
					@click="toPage(i)"
				>
					{{ i }}
				</button>
				<button class="btn btn-disabled">...</button>
				<button class="btn" @click="toPage(pageCount)">
					{{ pageCount }}
				</button>
			</template>
			<template v-else-if="isComplex && modelValue > pageCount - 4">
				<button class="btn" @click="toPage(1)">1</button>
				<button class="btn btn-disabled">...</button>
				<button
					class="btn"
					:class="{ 'btn-active': i == modelValue }"
					v-for="i in inversed"
					@click="toPage(i)"
				>
					{{ i }}
				</button>
			</template>
			<template v-else>
				<button class="btn" @click="toPage(1)">1</button>
				<button class="btn btn-disabled">...</button>
				<button class="btn" @click="toPage(modelValue - 1)">
					{{ modelValue - 1 }}
				</button>
				<button class="btn btn-active">{{ modelValue }}</button>
				<button class="btn" @click="toPage(modelValue + 1)">
					{{ modelValue + 1 }}
				</button>
				<button class="btn btn-disabled">...</button>
				<button class="btn" @click="toPage(pageCount)">
					{{ pageCount }}
				</button>
			</template>
		</div>
	</div>
</template>

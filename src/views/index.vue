<script setup lang="ts">
import { ref, onMounted } from "vue"
import Chart from "chart.js/auto"

import Navigation from "../components/Navigation.vue"

const lineChartCanvas = ref<HTMLCanvasElement>()
let lineChart: Chart
onMounted(() => {
	if (lineChartCanvas.value == undefined) return
	lineChart = new Chart(lineChartCanvas.value, {
		type: "line",
		data: {
			labels: [1, 2, 3, 4, 5, 6, 7],
			datasets: [{
				label: 'My First Dataset',
				data: [65, 59, 80, 81, 56, 55, 40],
				fill: false,
				borderColor: 'rgb(75, 192, 192)',
				tension: 0.1
			}]
		}
	})
})
</script>

<template>
	<div class="relative w-screen min-h-screen">
		<div class="w-11/12 mx-auto">
			<h1 class="py-4 text-2xl font-bold text-center">Dashboard</h1>
			<div class="p-4 bg-base-200 rounded-lg">
				<h4 class="text-lg font-bold">My Chart</h4>
				<canvas ref="lineChartCanvas" />
			</div>
		</div>
		<Navigation />
	</div>
</template>

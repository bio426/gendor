import { createApp } from "vue"
import Chart from "chart.js/auto"

import "./style.css"
import App from "./App.vue"
import router from "./router"

//config chart colors
if (window.matchMedia("(prefers-color-scheme: dark)").matches) {
	Chart.defaults.color = "#ADBABD";
	Chart.defaults.borderColor = "rgba(255,255,255,0.1)";
	Chart.defaults.backgroundColor = "rgba(255,255,0,0.1)";
	Chart.defaults.elements.line.borderColor = "rgba(255,255,0,0.4)";
}

createApp(App).use(router).mount("#app")

import { defineStore } from "pinia"
import { ref, reactive } from "vue"

const useUserStore = defineStore("user", () => {
	const token = ref<string | undefined>()
	const info = reactive<{ id?: string; username?: string }>({})

	return { token, info }
})

export default useUserStore

import { ref, watch } from "vue"

const isActive = ref(false)
const acceptTrigger = ref(Symbol())
const rejectTrigger = ref(Symbol())

function confirm(): Promise<boolean> {
	isActive.value = true
	return new Promise((resolve) => {
		watch(acceptTrigger, () => {
			resolve(true)
			isActive.value = false
		})
		watch(rejectTrigger, () => {
			resolve(false)
			isActive.value = false
		})
	})
}

export default function() {
	return {
		isActive,
		confirm,
		acceptTrigger,
		rejectTrigger,
	}
}

import { ref } from "vue"

type ToastVariant = "success" | "error"

interface IToast {
	text: string
	duration: number
	key: symbol
}

const toasts = ref<IToast[]>([])

function showToast(content: string, duration: number) {
	let key = Symbol()
	toasts.value.unshift({
		text: content,
		duration,
		key,
	})

	setTimeout(() => {
		let idx = toasts.value.findIndex(toast => toast.key == key)
		toasts.value.splice(idx, 1)
	}, duration)
}

export default function() {
	return {
		toasts,
		showToast,
	}
}

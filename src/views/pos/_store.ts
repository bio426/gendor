import { ref, computed } from "vue"

import type { IItem, ICartItem } from "./_types"

const cart = ref<ICartItem[]>([])

const cartLength = computed(() =>
	cart.value.reduce((prev, cur) => {
		return prev + cur.quantity
	}, 0)
)

const cartPrice = computed(() =>
	cart.value.reduce((prev, cur) => {
		return prev + cur.price * cur.quantity
	}, 0)
)

function addToCart(item: IItem) {
	let idx = cart.value.findIndex((cartItem) => cartItem.id == item.id)
	if (idx == -1) {
		cart.value.push({ ...item, quantity: 1 })
		return
	}
	cart.value[idx].quantity += 1
}

function removeFromCart(idx: number) {
	let item = cart.value[idx]
	if (item.quantity > 1) {
		item.quantity -= 1
		return
	}
	cart.value.splice(idx, 1)
}

export default function() {
	return {
		cart,
		cartLength,
		cartPrice,
		addToCart,
		removeFromCart,
	}
}

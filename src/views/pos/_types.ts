interface IItem {
	id: string
	name: string
	price: number
	quantity?: number
}

interface ICartItem extends IItem {
	quantity: number
}

export type { IItem, ICartItem }

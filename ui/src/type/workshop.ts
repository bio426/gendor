export type OrderCreate = {
	propietary: string
	brand: string
	model: string
	year: number
	plate: string
	mileage: number
	discount: number
	observation: string
	items: OrderItem[]
}

export type OrderItem = {
	price: number
	description: string
}

export type OrderSimple = {
	id: number
	propietary: string
	plate: string
	createdAt: string
}

export type OrderFound = {
	name?: string
	address: string
	dni: string
	ruc: string
	brand: string
	model: string
	color: string
	mileage: number
}

export type OrderDetail = OrderCreate & { id: number; createdAt: string }

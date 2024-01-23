export type Product = {
	id: number
	name: string
	price: number
	image: string
}

export type CreateProduct = {
	name: string
	price: number
	image: File
}

export type ProductCategory = {
	id: number
	name: string
}

export type Vehicle = {
	id: number
	model: string
	brand: string
}

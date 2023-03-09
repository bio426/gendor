import { ICategory } from "./admin"

export interface IProduct {
	id: string
	name: string
	price: number
	category: ICategory
}

export interface ICreateProduct {
	name: string
	price: number
	category: string
}

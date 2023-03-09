interface ICategory {
	id: string
	name: string
	color: string
}

interface ICreateCategory {
	name: string
	color: string
}

export type { ICategory, ICreateCategory }

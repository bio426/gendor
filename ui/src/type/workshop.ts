export type Order = {
    id?: number
    name: string
    address: string
    dni: string
    ruc: string
    brand: string
    model: string
    color: string
    plate: string
    mileage: number
    observation: string
    items: OrderItem[]
}

export type OrderItem = {
    code: string
    quantity: number
    price: number
    description: string
}


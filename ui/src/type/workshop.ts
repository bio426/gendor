export type OrderCreate = {
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

export type OrderSimple = {
    id: number,
    name: string,
    plate: string,
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

export type OrderDetail = OrderCreate & { id: number, createdAt: string }

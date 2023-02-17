import {IAdminTag} from "./admin"

export interface IItem{
    name: string
    price: number
    tags: IAdminTag[]
}
type UserRole = "super" | "admin"

export type User = {
    username: string
    role: UserRole
}

export type ExpirableUser = User & { expiryDate: string }

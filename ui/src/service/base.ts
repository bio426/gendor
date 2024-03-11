import ky from "ky"

import config from "../config"
import useToast from "@/composable/useToast"

const base = ky.create({
    prefixUrl: config.restUrl,
    credentials: "include",
    hooks: {
        beforeError: [
            err => {
                if (err.response.status >= 500) {
                    const toast = useToast()
                    toast.display({ message: "Request Error", variant: "error" })
                }
                return err
            }
        ]
    }
})

export default base

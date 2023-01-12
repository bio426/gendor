import ky from "ky"

import config from "../config"

let instance = ky.create({ prefixUrl: config.apiUrl })

export default instance

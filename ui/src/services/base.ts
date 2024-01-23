import ky from "ky"

import config from "../config"

const base = ky.create({ prefixUrl: config.apiUrl })

export default base

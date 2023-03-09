import ky from "ky"
import Pocketbase from "pocketbase"

import config from "../config"

const pb = new Pocketbase(config.PB_URL)

let instance = ky.create({ prefixUrl: config.apiUrl })

export default instance

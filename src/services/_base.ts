import Pocketbase from "pocketbase"

import config from "../config"

const pb = new Pocketbase(config.PB_URL)


export default pb

const config = {
	restUrl: "http://localhost:1323/rest",
	sseUrl: "http://localhost:1323/event",
}

if (import.meta.env.PROD) {
	config.restUrl = `/rest`
}

export default config

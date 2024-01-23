const config = {
	apiUrl:
		(import.meta.env.VITE_API_URL as string) || "http://localhost:1323/api",
	sseUrl:
		(import.meta.env.VITE_SSE_URL as string) ||
		"http://localhost:1323/event",
}

export default config

package main

import "net/http"

func main() {

	// GET /games
	http.HandleFunc("GET /games", getGamesHandler)

	// GET /games/{id}
	http.HandleFunc("GET /game/{id}", getGameIDHandler)

	// GET /providers
	http.HandleFunc("GET /providers", getProviderHandler)

	// GET /sessions
	http.HandleFunc("GET /sessions", getSessionsHandler)

	// POST /games
	http.HandleFunc("POST /games", createGameHandler)

	// POST /sessions
	http.HandleFunc("POST /sessions", createSessionsHandler)

	// Inicia o servidor HTTP na porta 8080.
	http.ListenAndServe(":8080", nil)
}

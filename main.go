package main

import "net/http"

func main() {

	var err error

	db, err = connectDatabase()

	if err != nil {
		panic(err)
	}

	defer db.Close()

	http.HandleFunc("GET /games", getGamesHandler)
	http.HandleFunc("GET /games/{id}", getGamesIDHandler)
	http.HandleFunc("GET /providers", getProvidersHandler)
	http.HandleFunc("GET /sessions", getSessionsHandler)
	http.HandleFunc("POST /games", createGameHandler)
	http.HandleFunc("POST /sessions", createSessionsHandler)

	handler := loggingMiddleware(http.DefaultServeMux)

	go startSessionWorker()

	if err := http.ListenAndServe(":8080", handler); err != nil {
		panic(err)
}
}

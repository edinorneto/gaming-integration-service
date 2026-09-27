package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/edinorneto/gaming-integration-service/domain"

	"github.com/jackc/pgx/v5"
)

// ============================================================
// GET /games
// ============================================================

// handler retorna todos os jogos do catálogo.
func getGamesHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	games, err := getGamesFromDatabase()

	if err != nil {
		writeJSONError(w, "Erro ao buscar games.", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(games)
}

// ============================================================
// GET /providers
// ============================================================

// handler retorna todos os providers do catálogo.

func getProvidersHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	providers, err := getProvidersFromDatabase()

	if err != nil {
		writeJSONError(
			w,
			"Erro ao buscar providers.",
			http.StatusInternalServerError,
		)
		return
	}

	json.NewEncoder(w).Encode(providers)
}

// ============================================================
// GET /sessions
// ============================================================

// handler retorna as sessions.

func getSessionsHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	sessions, err := getSessionsFromDatabase()

	if err != nil {
		writeJSONError(
			w,
			"Erro ao buscar sessões.",
			http.StatusInternalServerError,
		)
		return
	}

	json.NewEncoder(w).Encode(sessions)
}

// ============================================================
// GET /games/{id}
// ============================================================

// getGamesIDHandler procura um jogo pelo ID.
func getGamesIDHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		writeJSONError(
			w,
			"ID inválido.",
			http.StatusBadRequest,
		)
		return
	}

	game, err := getGameByIDFromDatabase(id)

	if err != nil {

		if errors.Is(err, pgx.ErrNoRows) {
			writeJSONError(
				w,
				"Game não encontrado.",
				http.StatusNotFound,
			)
			return
		}

		writeJSONError(
			w,
			"Erro ao buscar game.",
			http.StatusInternalServerError,
		)
		return
	}

	json.NewEncoder(w).Encode(game)
}

// ============================================================
// POST /games
// ============================================================

// createGameHandler recebe um novo jogo em JSON
// e persiste o jogo no PostgreSQL.

func createGameHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var game domain.Game

	err := json.NewDecoder(r.Body).Decode(&game)

	if err != nil {
		writeJSONError(
			w,
			"JSON inválido.",
			http.StatusBadRequest,
		)
		return
	}

	exists, err := providerExistsInDatabase(game.ProviderID)

	if err != nil {
		writeJSONError(
			w,
			"Erro ao verificar provider.",
			http.StatusInternalServerError,
		)
		return
	}

	if !exists {
		writeJSONError(
			w,
			"Provider não encontrado.",
			http.StatusBadRequest,
		)
		return
	}

	createdGame, err := createGameInDatabase(game)

	if err != nil {
		writeJSONError(
			w,
			"Erro ao criar game.",
			http.StatusInternalServerError,
		)
		return
	}

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(createdGame)
}

// ============================================================
// POST /sessions
// ============================================================

// createSessionsHandler recebe uma requisição de sessão
// e inicia uma sessão caso o game exista.

func createSessionsHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var session domain.GameSession

	err := json.NewDecoder(r.Body).Decode(&session)

	if err != nil {
		writeJSONError(
			w,
			"JSON inválido.",
			http.StatusBadRequest,
		)
		return
	}

	exists, err := gameExistsInDatabase(session.GameID)

	if err != nil {
		writeJSONError(
			w,
			"Erro ao verificar game.",
			http.StatusInternalServerError,
		)
		return
	}

	if !exists {
		writeJSONError(
			w,
			"Game não encontrado.",
			http.StatusBadRequest,
		)
		return
	}

	session.Status = "active"

	createdSession, err := createSessionInDatabase(session)

	if err != nil {
		writeJSONError(
			w,
			"Erro ao criar sessão.",
			http.StatusInternalServerError,
		)
		return
	}

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(createdSession)
}

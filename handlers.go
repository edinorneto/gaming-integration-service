package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"projeto-go/domain"
)

// ============================================================
// DADOS EM MEMÓRIA
// ============================================================

// Providers disponíveis no sistema.
// Dados apenas para demonstração e desenvolvimento local.
// Não representam integrações oficiais com essas empresas.
var providers = []domain.Provider{
	{
		ID:     1,
		Nome:   "Jungle Originals",
		Active: true,
	},
	{
		ID:     2,
		Nome:   "Aurora Gaming",
		Active: true,
	},
	{
		ID:     3,
		Nome:   "Emerald Interactive",
		Active: true,
	},
	{
		ID:     4,
		Nome:   "NovaPlay",
		Active: true,
	},
	{
		ID:     5,
		Nome:   "Orbit Gaming",
		Active: true,
	},
}

// Catálogo inicial de jogos.
// ProviderID relaciona cada jogo ao provider correspondente.

var games = []domain.Game{
	// Jungle Originals
	{ID: 1, Nome: "Captain's Treasure", ProviderID: 1},
	{ID: 2, Nome: "Fox the Course Seller", ProviderID: 1},
	{ID: 3, Nome: "Chimp Mines", ProviderID: 1},
	{ID: 4, Nome: "Goblin's Gold", ProviderID: 1},

	// Aurora Gaming
	{ID: 5, Nome: "Emerald Rush", ProviderID: 2},
	{ID: 6, Nome: "Golden Jungle", ProviderID: 2},
	{ID: 7, Nome: "Mystic Fortune", ProviderID: 2},
	{ID: 8, Nome: "Treasure Spins", ProviderID: 2},

	// Emerald Interactive
	{ID: 9, Nome: "Neon Reels", ProviderID: 3},
	{ID: 10, Nome: "Diamond Vault", ProviderID: 3},
	{ID: 11, Nome: "Wild Horizon", ProviderID: 3},
	{ID: 12, Nome: "Lucky Temple", ProviderID: 3},

	// NovaPlay
	{ID: 13, Nome: "Golden Quest", ProviderID: 4},
	{ID: 14, Nome: "Treasure Temple", ProviderID: 4},
	{ID: 15, Nome: "Moonlit Fortune", ProviderID: 4},
	{ID: 16, Nome: "Wild Expedition", ProviderID: 4},

	// Orbit Gaming
	{ID: 17, Nome: "Cyber Fortune", ProviderID: 5},
	{ID: 18, Nome: "Jungle Nights", ProviderID: 5},
	{ID: 19, Nome: "Crystal Reels", ProviderID: 5},
	{ID: 20, Nome: "Lucky Orbit", ProviderID: 5},
}

// Sessões criadas durante a execução da aplicação.
var sessions []domain.GameSession

// ============================================================
// GET /games
// ============================================================

// handler retorna todos os jogos do catálogo.
func getGamesHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(games)
}

// ============================================================
// GET /providers
// ============================================================

// handler retorna todos os providers do catálogo.

func getProvidersHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(providers)

}

// ============================================================
// GET /sessions
// ============================================================

// handler retorna as sessions.

func getSessionsHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(sessions)

}

// ============================================================
// GET /games/{id}
// ============================================================

// gameHandler procura um jogo pelo ID informado na URL.
func getGameIDHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	// Pega o valor do {id} na URL.
	id := r.PathValue("id")

	// Converte o ID de string para int.
	gameID, err := strconv.Atoi(id)

	if err != nil {
		writeJSONError(w, "ID inválido.", http.StatusBadRequest)
		return
	}

	// Procura o jogo no catálogo.
	for _, game := range games {

		if game.ID == gameID {

			json.NewEncoder(w).Encode(game)
			return
		}
	}

	// Se percorreu todos os jogos e não encontrou o ID.
	writeJSONError(w, "Jogo não encontrado.", http.StatusNotFound)

}

// ============================================================
// POST /games
// ============================================================

// createGameHandler recebe um novo jogo em JSON
// e adiciona o jogo ao catálogo em memória.

func createGameHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var game domain.Game

	// Lê o JSON enviado pelo cliente e preenche a struct Game.
	err := json.NewDecoder(r.Body).Decode(&game)

	if err != nil {
		writeJSONError(w, "JSON inválido.", http.StatusBadRequest)
		return
	}

	providerExists := false

	for _, p := range providers {
		if p.ID == game.ProviderID {
			providerExists = true
			break
		}
	}

	if !providerExists {
		writeJSONError(w, "Provider não encontrado.", http.StatusNotFound)
		return
	}

	// Gera um ID simples para o exemplo em memória.
	game.ID = len(games) + 1

	// Adiciona o novo jogo ao catálogo.
	games = append(games, game)

	w.WriteHeader(http.StatusCreated)

	// Retorna o jogo criado.
	json.NewEncoder(w).Encode(game)
}

// ============================================================
// POST /sessions
// ============================================================

// createSessionsHandler recebe uma requisição de sessão
// e inicia uma sessão caso o game exista.

func createSessionHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var session domain.GameSession

	err := json.NewDecoder(r.Body).Decode(&session)

	if err != nil {
		writeJSONError(w, "JSON inválido.", http.StatusBadRequest)
		return
	}

	for _, game := range games {

		if game.ID == session.GameID {

			session.ID = len(sessions) + 1
			session.Status = "active"

			sessions = append(sessions, session)

			w.WriteHeader(http.StatusCreated)

			json.NewEncoder(w).Encode(session)
			return
		}
	}

	writeJSONError(w, "Jogo não encontrado.", http.StatusNotFound)
}

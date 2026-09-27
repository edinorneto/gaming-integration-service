package main

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/edinorneto/gaming-integration-service/domain"
)

func TestGetGamesIDHandler(t *testing.T) {

	var err error

	db, err = connectDatabase()

	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	req := httptest.NewRequest(
		"GET",
		"/games/7",
		nil,
	)

	req.SetPathValue("id", "7")

	rec := httptest.NewRecorder()

	getGamesIDHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf(
			"esperado status 200, recebido %d",
			rec.Code,
		)
	}

	var game domain.Game

	err = json.NewDecoder(rec.Body).Decode(&game)

	if err != nil {
		t.Fatal(err)
	}

	if game.ID != 7 {
		t.Errorf(
			"esperado ID 7, recebido %d",
			game.ID,
		)
	}

	if game.Nome != "Mystic Fortune" {
		t.Errorf(
			"esperado nome Mystic Fortune, recebido %s",
			game.Nome,
		)
	}

	if game.ProviderID != 2 {
		t.Errorf(
			"esperado ProviderID 2, recebido %d",
			game.ProviderID,
		)
	}
}

func TestGetGamesIDHandlerNotFound(t *testing.T) {

	var err error

	db, err = connectDatabase()

	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	req := httptest.NewRequest(
		"GET",
		"/games/999",
		nil,
	)

	req.SetPathValue("id", "999")

	rec := httptest.NewRecorder()

	getGamesIDHandler(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Errorf(
			"esperado status 404, recebido %d",
			rec.Code,
		)
	}
}

func TestGetGamesIDHandlerBadRequest(t *testing.T) {

	var err error

	db, err = connectDatabase()

	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	req := httptest.NewRequest(
		"GET",
		"/games/abc",
		nil,
	)

	req.SetPathValue("id", "abc")

	rec := httptest.NewRecorder()

	getGamesIDHandler(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf(
			"esperado status 400, recebido %d",
			rec.Code,
		)
	}
}

func TestGetProvidersHandler(t *testing.T) {

	var err error

	db, err = connectDatabase()

	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	req := httptest.NewRequest(
		"GET",
		"/providers",
		nil,
	)

	rec := httptest.NewRecorder()

	getProvidersHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf(
			"esperado status 200, recebido %d",
			rec.Code,
		)
	}

	var providers []domain.Provider

	err = json.NewDecoder(rec.Body).Decode(&providers)

	if err != nil {
		t.Fatal(err)
	}

	if len(providers) != 5 {
		t.Errorf(
			"esperado 5 providers, recebido %d",
			len(providers),
		)
	}

	if providers[0].ID != 1 {
		t.Errorf(
			"esperado ID 1, recebido %d",
			providers[0].ID,
		)
	}

	if providers[0].Nome != "Jungle Originals" {
		t.Errorf(
			"esperado nome Jungle Originals, recebido %s",
			providers[0].Nome,
		)
	}

	if !providers[0].Active {
		t.Errorf("esperado provider ativo")
	}
}

func TestCreateGameHandler(t *testing.T) {

	var err error

	db, err = connectDatabase()

	if err != nil {
		t.Fatal(err)
	}

	createdGameID := 0

	t.Cleanup(func() {
		if createdGameID != 0 {
			_, err := db.Exec(
				context.Background(),
				"DELETE FROM games WHERE id = $1",
				createdGameID,
			)

			if err != nil {
				t.Errorf("erro ao limpar game de teste: %v", err)
			}
		}

		db.Close()
	})

	body := strings.NewReader(
		`{"nome":"Test Game","provider_id":2}`,
	)

	req := httptest.NewRequest(
		"POST",
		"/games",
		body,
	)

	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	createGameHandler(rec, req)

	if rec.Code != http.StatusCreated {
		t.Errorf(
			"esperado status 201, recebido %d",
			rec.Code,
		)
	}

	var game domain.Game

	err = json.NewDecoder(rec.Body).Decode(&game)

	if err != nil {
		t.Fatal(err)
	}

	createdGameID = game.ID

	if game.ID <= 0 {
		t.Errorf(
			"esperado criar um ID, recebido %d",
			game.ID,
		)
	}

	if game.Nome != "Test Game" {
		t.Errorf(
			"esperado nome Test Game, recebido %s",
			game.Nome,
		)
	}

	if game.ProviderID != 2 {
		t.Errorf(
			"esperado ProviderID 2, recebido %d",
			game.ProviderID,
		)
	}
}

func TestCreateGameHandlerProviderNotFound(t *testing.T) {

	var err error

	db, err = connectDatabase()

	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	body := strings.NewReader(
		`{"nome":"Test Game Invalid","provider_id":999}`,
	)

	req := httptest.NewRequest(
		"POST",
		"/games",
		body,
	)

	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	createGameHandler(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf(
			"esperado status 400, recebido %d",
			rec.Code,
		)
	}
}

func TestCreateSessionHandler(t *testing.T) {

	var err error

	db, err = connectDatabase()

	if err != nil {
		t.Fatal(err)
	}

	createdSessionID := 0

	t.Cleanup(func() {
		if createdSessionID != 0 {
			_, err := db.Exec(
				context.Background(),
				"DELETE FROM game_sessions WHERE id = $1",
				createdSessionID,
			)

			if err != nil {
				t.Errorf("erro ao limpar sessão de teste: %v", err)
			}
		}

		db.Close()
	})

	body := strings.NewReader(
		`{"player_id":123,"game_id":7}`,
	)

	req := httptest.NewRequest(
		"POST",
		"/sessions",
		body,
	)

	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	createSessionsHandler(rec, req)

	if rec.Code != http.StatusCreated {
		t.Errorf(
			"esperado status 201, recebido %d",
			rec.Code,
		)
	}

	var session domain.GameSession

	err = json.NewDecoder(rec.Body).Decode(&session)

	if err != nil {
		t.Fatal(err)
	}

	createdSessionID = session.ID

	if session.ID <= 0 {
		t.Errorf(
			"esperado criar um ID, recebido %d",
			session.ID,
		)
	}

	if session.PlayerID != 123 {
		t.Errorf(
			"esperado PlayerID 123, recebido %d",
			session.PlayerID,
		)
	}

	if session.GameID != 7 {
		t.Errorf(
			"esperado GameID 7, recebido %d",
			session.GameID,
		)
	}

	if session.Status != "active" {
		t.Errorf(
			"esperado status active, recebido %s",
			session.Status,
		)
	}
}

func TestCreateSessionHandlerGameNotFound(t *testing.T) {

	var err error

	db, err = connectDatabase()

	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	body := strings.NewReader(
		`{"player_id":123,"game_id":999}`,
	)

	req := httptest.NewRequest(
		"POST",
		"/sessions",
		body,
	)

	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	createSessionsHandler(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf(
			"esperado status 400, recebido %d",
			rec.Code,
		)
	}
}

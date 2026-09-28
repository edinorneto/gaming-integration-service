package main

import (
	"context"
	"fmt"
	"os"

	"github.com/edinorneto/gaming-integration-service/domain"

	"github.com/jackc/pgx/v5/pgxpool"
)

var db *pgxpool.Pool

func connectDatabase() (*pgxpool.Pool, error) {

	databaseURL := os.Getenv("DATABASE_URL")

	if databaseURL == "" {
		return nil, fmt.Errorf("DATABASE_URL não definida")
	}

	pool, err := pgxpool.New(
		context.Background(),
		databaseURL,
	)

	if err != nil {
		return nil, err
	}

	err = pool.Ping(context.Background())

	if err != nil {
		pool.Close()
		return nil, err
	}

	return pool, nil
}

func gameExistsInDatabase(id int) (bool, error) {

	var exists bool

	err := db.QueryRow(
		context.Background(),
		"SELECT EXISTS (SELECT 1 FROM games WHERE id = $1)",
		id,
	).Scan(&exists)

	if err != nil {
		return false, err
	}

	return exists, nil
}

func getGamesFromDatabase() ([]domain.Game, error) {

	rows, err := db.Query(
		context.Background(),
		"SELECT id, nome, provider_id FROM games ORDER BY id",
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var games []domain.Game

	for rows.Next() {

		var game domain.Game

		err := rows.Scan(
			&game.ID,
			&game.Nome,
			&game.ProviderID,
		)

		if err != nil {
			return nil, err
		}

		games = append(games, game)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return games, nil
}

func getGameByIDFromDatabase(id int) (domain.Game, error) {

	var game domain.Game

	err := db.QueryRow(
		context.Background(),
		"SELECT id, nome, provider_id FROM games WHERE id = $1",
		id,
	).Scan(
		&game.ID,
		&game.Nome,
		&game.ProviderID,
	)

	if err != nil {
		return domain.Game{}, err
	}

	return game, nil
}

func providerExistsInDatabase(id int) (bool, error) {

	var exists bool

	err := db.QueryRow(
		context.Background(),
		"SELECT EXISTS (SELECT 1 FROM providers WHERE id = $1)",
		id,
	).Scan(&exists)

	if err != nil {
		return false, err
	}

	return exists, nil
}

func getProvidersFromDatabase() ([]domain.Provider, error) {

	rows, err := db.Query(
		context.Background(),
		"SELECT id, nome, active FROM providers ORDER BY id",
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var providers []domain.Provider

	for rows.Next() {

		var provider domain.Provider

		err := rows.Scan(
			&provider.ID,
			&provider.Nome,
			&provider.Active,
		)

		if err != nil {
			return nil, err
		}

		providers = append(providers, provider)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return providers, nil
}

func getSessionsFromDatabase() ([]domain.GameSession, error) {

	rows, err := db.Query(
		context.Background(),
		"SELECT id, player_id, game_id, status FROM game_sessions ORDER BY id",
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var sessions []domain.GameSession

	for rows.Next() {

		var session domain.GameSession

		err := rows.Scan(
			&session.ID,
			&session.PlayerID,
			&session.GameID,
			&session.Status,
		)

		if err != nil {
			return nil, err
		}

		sessions = append(sessions, session)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return sessions, nil
}

func getSessionByIDFromDatabase(id int) (domain.GameSession, error) {

	var session domain.GameSession

	err := db.QueryRow(
		context.Background(),
		`
		SELECT id, player_id, game_id, status
		FROM game_sessions
		WHERE id = $1
		`,
		id,
	).Scan(
		&session.ID,
		&session.PlayerID,
		&session.GameID,
		&session.Status,
	)

	if err != nil {
		return domain.GameSession{}, err
	}

	return session, nil
}

func createGameInDatabase(game domain.Game) (domain.Game, error) {

	var createdGame domain.Game

	err := db.QueryRow(
		context.Background(),
		`
		INSERT INTO games (nome, provider_id)
		VALUES ($1, $2)
		RETURNING id, nome, provider_id
		`,
		game.Nome,
		game.ProviderID,
	).Scan(
		&createdGame.ID,
		&createdGame.Nome,
		&createdGame.ProviderID,
	)

	if err != nil {
		return domain.Game{}, err
	}

	return createdGame, nil
}

func createSessionInDatabase(session domain.GameSession) (domain.GameSession, error) {

	var createdSession domain.GameSession

	err := db.QueryRow(
		context.Background(),
		`
		INSERT INTO game_sessions (player_id, game_id, status)
		VALUES ($1, $2, $3)
		RETURNING id, player_id, game_id, status
		`,
		session.PlayerID,
		session.GameID,
		session.Status,
	).Scan(
		&createdSession.ID,
		&createdSession.PlayerID,
		&createdSession.GameID,
		&createdSession.Status,
	)

	if err != nil {
		return domain.GameSession{}, err
	}

	return createdSession, nil
}

func updateSessionStatusInDatabase(id int, status string) (bool, error) {

	result, err := db.Exec(
		context.Background(),
		`
		UPDATE game_sessions
		SET status = $1
		WHERE id = $2
		  AND status = 'pending'
		`,
		status,
		id,
	)

	if err != nil {
		return false, err
	}

	return result.RowsAffected() > 0, nil
}
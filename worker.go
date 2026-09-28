package main

import "log/slog"

func startSessionWorker() {
	for message := range sessionQueue {

		session, err := getSessionByIDFromDatabase(message.SessionID)

		if err != nil {
			slog.Error(
				"Falha ao processar sessão",
				"session_id", message.SessionID,
				"error", err,
			)

			retryMessage(message)
			continue
		}

		slog.Info(
			"Sessão processada",
			"session_id", session.ID,
			"game_id", session.GameID,
			"player_id", session.PlayerID,
			"status", session.Status,
		)

		updated, err := updateSessionStatusInDatabase(session.ID, "active")

		if err != nil {
			slog.Error(
				"Falha ao atualizar status da sessão",
				"session_id", session.ID,
				"error", err,
			)

			retryMessage(message)
			continue
		}

		if !updated {
			slog.Info(
				"Sessão já processada",
				"session_id", session.ID,
			)
			continue
		}

		slog.Info(
			"Sessão ativada",
			"session_id", session.ID,
		)
	}
}

func retryMessage(message SessionMessage) {
	if message.Attempts >= 3 {
		slog.Error(
			"Limite de tentativas atingido",
			"session_id", message.SessionID,
		)
		return
	}

	message.Attempts++

	slog.Warn(
		"Retry da sessão",
		"session_id", message.SessionID,
		"attempt", message.Attempts,
	)

	sessionQueue <- message
}
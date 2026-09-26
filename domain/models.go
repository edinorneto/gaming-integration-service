package domain

// Game representa um jogo disponibilizado por um provider.
type Game struct {
	ID         int    `json:"id"`
	Nome       string `json:"nome"`
	ProviderID int    `json:"provider_id"`
}

// Provider representa um provedor de jogos.
type Provider struct {
	ID     int    `json:"id"`
	Nome   string `json:"nome"`
	Active bool   `json:"active"`
}

// GameSession representa uma sessão de jogo iniciada por um jogador.
type GameSession struct {
	ID       int    `json:"id"`
	PlayerID int    `json:"player_id"`
	GameID   int    `json:"game_id"`
	Status   string `json:"status"`
}

// IsAvailable verifica se o provider está ativo.
func (p *Provider) IsAvailable() bool {
	return p.Active
}

// IsActive verifica se a sessão está ativa.
func (s *GameSession) IsActive() bool {
	return s.Status == "active"
}

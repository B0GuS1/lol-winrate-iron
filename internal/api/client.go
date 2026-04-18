package api

import (
	"net/http"

	"golang.org/x/time/rate"
)

type RiotClient struct {
	httpClient  *http.Client
	apiKey      string
	limiter     *rate.Limiter
	baseURL     string
	regionalURL string
}

const (
	BaseURLEurope = "https://europe.api.riotgames.com"
	RegionalURLRU = "https://ru.api.riotgames.com"
)

func NewRiotClient(apiKey string) *RiotClient {
	// Rate limits Riot API:
	// 20 запросов в секунду
	// 100 запросов в 2 минуты
	// Используем консервативный лимит: 20 запросов в секунду
	limiter := rate.NewLimiter(rate.Limit(20), 1)

	return &RiotClient{
		httpClient:  http.DefaultClient,
		apiKey:      apiKey,
		limiter:     limiter,
		baseURL:     BaseURLEurope,
		regionalURL: RegionalURLRU,
	}
}

// GetMatchIDsByQueue получает ID матчей конкретного игрока по PUUID с фильтрацией по типу очереди (Ranked Solo/Duo).
// Используется endpoint /lol/match/v5/matches/by-puuid/{puuid}/ids
//func (c *RiotClient) GetMatchIDsByQueue(ctx context.Context, puuid string, queueID int, count int) ([]string, error)

// GetMatchData получает полные данные одного матча по его ID.
//func (c *RiotClient) GetMatchData(ctx context.Context, matchID string) (*MatchDTO, error)

// GetSummonerByRiotID получает PUUID игрока по его Riot ID (никнейм#таг).
//func (c *RiotClient) GetSummonerByRiotID(ctx context.Context, gameName, tagLine string) (*SummonerDTO, error)

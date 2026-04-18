package api

import (
	"context"
	"net/http"

	"golang.org/x/time/rate"
)

type RiotClient struct {
	httpClient *http.Client
	apiKey     string
	limiter    *rate.Limiter
}

// NewRiotClient создает нового клиента с встроенным ограничением запросов.
func NewRiotClient(apiKey string) *RiotClient

// GetMatchIDsByQueue получает ID матчей конкретного игрока по PUUID с фильтрацией по типу очереди (Ranked Solo/Duo).
// Используется endpoint /lol/match/v5/matches/by-puuid/{puuid}/ids
func (c *RiotClient) GetMatchIDsByQueue(ctx context.Context, puuid string, queueID int, count int) ([]string, error)

// GetMatchData получает полные данные одного матча по его ID.
func (c *RiotClient) GetMatchData(ctx context.Context, matchID string) (*MatchDTO, error)

// GetSummonerByRiotID получает PUUID игрока по его Riot ID (никнейм#таг).
func (c *RiotClient) GetSummonerByRiotID(ctx context.Context, gameName, tagLine string) (*SummonerDTO, error)

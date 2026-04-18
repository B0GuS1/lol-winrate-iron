package api

import (
	"context"
	"encoding/json"
	"fmt"
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

func (c *RiotClient) GetIron4Players(ctx context.Context, page int) ([]LeagueEntryDTO, error) {
	url := fmt.Sprintf("%s/lol/league/v4/entries/RANKED_SOLO_5x5/IRON/IV?page=%d&api_key=%s",
		c.regionalURL, page, c.apiKey)
	req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
	c.limiter.Wait(ctx)
	resp, _ := c.httpClient.Do(req)
	var entries []LeagueEntryDTO
	json.NewDecoder(resp.Body).Decode(&entries)
	return entries, nil
}

func (c *RiotClient) GetAllIron4Players(ctx context.Context) ([]LeagueEntryDTO, error) {
	var allEntries []LeagueEntryDTO

	for page := 1; ; page++ {
		entries, err := c.GetIron4Players(ctx, page)
		if err != nil {
			return nil, fmt.Errorf("error fetching page %d: %w", page, err)
		}

		// Если страница пустая - достигли конца
		if len(entries) == 0 {
			break
		}

		allEntries = append(allEntries, entries...)

		// Опционально: логирование прогресса
		fmt.Printf("Fetched page %d, got %d entries", page, len(entries))
	}

	return allEntries, nil
}

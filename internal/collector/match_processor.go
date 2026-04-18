package collector

import (
	"context"
)

// ProcessSummonerMatches получает историю матчей одного игрока, сохраняет только те, что сыграны в Iron IV.
func ProcessSummonerMatches(ctx context.Context, puuid string, client *api.RiotClient, repo *storage.Repository) error {
	// 1. Получить список ID матчей (queue=420) [citation:6].
	ids, _ := client.GetMatchIDsByQueue(ctx, puuid, 420, 100)

	for _, matchID := range ids {
		// 2. Проверить, нет ли уже в БД этого матча (избегаем дублей и экономим запросы).
		if repo.MatchExists(matchID) {
			continue
		}
		match, _ := client.GetMatchData(ctx, matchID)

		// 3. Фильтр: сохраняем только те данные, где участник реально находится в ранге IRON+IV.
		// Это отсеет игры, где Iron IV игрок попал в лобби более высокого ранга (флексы/нормалы).
		repo.SaveMatchData(filterIron4Data(match))
	}
	return nil
}

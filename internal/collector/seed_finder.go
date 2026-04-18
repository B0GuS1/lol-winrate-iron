package collector

// FindIron4Seeds находит пул PUUID игроков, которые находятся в ранге Iron IV.
// Стратегия: запрос лидербордов (если доступны) или сканирование матчей с фильтром по очереди.
func FindIron4Seeds(ctx context.Context, client *api.RiotClient) ([]string, error)

// DiscoverPlayersFromMatch анализирует участников матча и возвращает PUUID тех, у кого Tier=IRON и Rank=IV.
func DiscoverPlayersFromMatch(match *api.MatchDTO) []string

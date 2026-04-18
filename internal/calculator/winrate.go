package calculator

type ChampionWinRate struct {
	ChampionID   int
	ChampionName string
	TotalGames   int
	Wins         int
	WinRate      float64 // Процент побед
	AvgKDA       float64 // Опционально: средний KDA
}

// CalculateIron4Meta рассчитывает итоговую таблицу винрейтов по накопленным данным в БД.
func CalculateIron4Meta(repo *storage.Repository) ([]ChampionWinRate, error) {
	// SQL запрос: SELECT champion_id, COUNT(*), SUM(win) FROM matches WHERE tier='IRON' AND rank='IV' GROUP BY champion_id
	rows := repo.QueryIron4Stats()

	var stats []ChampionWinRate
	for _, row := range rows {
		// Расчет WinRate с плавающей точкой
		wr := float64(row.Wins) / float64(row.TotalGames) * 100
		// ...
	}
	return stats, nil
}

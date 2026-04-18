package api

type LeagueEntryDTO struct {
	LeagueID     string         `json:"leagueId"`
	PUUID        string         `json:"puuid"` // Player's encrypted puuid
	QueueType    string         `json:"queueType"`
	Tier         string         `json:"tier"`
	Rank         string         `json:"rank"` // The player's division within a tier
	LeaguePoints int            `json:"leaguePoints"`
	Wins         int            `json:"wins"`   // Winning team on Summoners Rift
	Losses       int            `json:"losses"` // Losing team on Summoners Rift
	HotStreak    bool           `json:"hotStreak"`
	Veteran      bool           `json:"veteran"`
	FreshBlood   bool           `json:"freshBlood"`
	Inactive     bool           `json:"inactive"`
	MiniSeries   *MiniSeriesDTO `json:"miniSeries,omitempty"` // Указатель и omitempty, т.к. поле не всегда присутствует
}

// Предполагаемая структура для MiniSeriesDTO (основана на стандартном API Riot)
type MiniSeriesDTO struct {
	Losses   int    `json:"losses"`
	Progress string `json:"progress"`
	Target   int    `json:"target"`
	Wins     int    `json:"wins"`
}

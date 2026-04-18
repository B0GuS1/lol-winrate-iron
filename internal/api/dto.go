package api

// ============================================================================
// ACCOUNT API DTOs (Глобальные эндпоинты)
// ============================================================================

// AccountDTO представляет аккаунт Riot (используется в Account-V1 API)
type AccountDTO struct {
	PUUID    string `json:"puuid"`
	GameName string `json:"gameName"`
	TagLine  string `json:"tagLine"`
}

// ActiveShardDTO представляет активный шард игрока
type ActiveShardDTO struct {
	PUUID       string `json:"puuid"`
	Game        string `json:"game"`
	ActiveShard string `json:"activeShard"`
}

// ============================================================================
// SUMMONER API DTOs (Региональные эндпоинты)
// ============================================================================

// SummonerDTO представляет информацию о призывателе
type SummonerDTO struct {
	ID            string `json:"id"`            // Summoner ID (encrypted)
	AccountID     string `json:"accountId"`     // Account ID (encrypted)
	PUUID         string `json:"puuid"`         // PUUID (encrypted)
	Name          string `json:"name"`          // Имя призывателя
	ProfileIconID int    `json:"profileIconId"` // ID иконки профиля
	RevisionDate  int64  `json:"revisionDate"`  // Дата последнего изменения
	SummonerLevel int    `json:"summonerLevel"` // Уровень призывателя
}

// ============================================================================
// LEAGUE API DTOs (Региональные эндпоинты)
// ============================================================================

// LeagueEntryDTO представляет запись о ранге игрока в конкретной очереди
type LeagueEntryDTO struct {
	LeagueID     string         `json:"leagueId"`             // ID лиги
	SummonerID   string         `json:"summonerId"`           // ID призывателя
	SummonerName string         `json:"summonerName"`         // Имя призывателя
	QueueType    string         `json:"queueType"`            // Тип очереди (RANKED_SOLO_5x5, RANKED_FLEX_SR)
	Tier         string         `json:"tier"`                 // Ранг (IRON, BRONZE, SILVER, etc.)
	Rank         string         `json:"rank"`                 // Дивизион (IV, III, II, I)
	LeaguePoints int            `json:"leaguePoints"`         // Очки лиги (LP)
	Wins         int            `json:"wins"`                 // Количество побед
	Losses       int            `json:"losses"`               // Количество поражений
	HotStreak    bool           `json:"hotStreak"`            // Серия побед
	Veteran      bool           `json:"veteran"`              // Ветеран
	FreshBlood   bool           `json:"freshBlood"`           // Новый игрок в лиге
	Inactive     bool           `json:"inactive"`             // Неактивный
	MiniSeries   *MiniSeriesDTO `json:"miniSeries,omitempty"` // Мини-серия (промо)
}

// MiniSeriesDTO представляет информацию о промо-серии
type MiniSeriesDTO struct {
	Losses   int    `json:"losses"`   // Поражения в серии
	Progress string `json:"progress"` // Прогресс (например "WLWNN")
	Target   int    `json:"target"`   // Необходимо побед
	Wins     int    `json:"wins"`     // Победы в серии
}

// LeagueListDTO представляет список игроков в лиге
type LeagueListDTO struct {
	LeagueID string          `json:"leagueId"`
	Entries  []LeagueItemDTO `json:"entries"`
	Tier     string          `json:"tier"`
	Name     string          `json:"name"`
	Queue    string          `json:"queue"`
}

// LeagueItemDTO представляет игрока в списке лиги
type LeagueItemDTO struct {
	SummonerID   string `json:"summonerId"`
	SummonerName string `json:"summonerName"`
	LeaguePoints int    `json:"leaguePoints"`
	Rank         string `json:"rank"`
	Wins         int    `json:"wins"`
	Losses       int    `json:"losses"`
	Veteran      bool   `json:"veteran"`
	Inactive     bool   `json:"inactive"`
	FreshBlood   bool   `json:"freshBlood"`
	HotStreak    bool   `json:"hotStreak"`
}

// ============================================================================
// MATCH API DTOs (Глобальные эндпоинты)
// ============================================================================

// MatchDTO представляет полные данные матча
type MatchDTO struct {
	Metadata MatchMetadataDTO `json:"metadata"`
	Info     MatchInfoDTO     `json:"info"`
}

// MatchMetadataDTO содержит метаданные матча
type MatchMetadataDTO struct {
	DataVersion  string   `json:"dataVersion"`  // Версия данных
	MatchID      string   `json:"matchId"`      // ID матча
	Participants []string `json:"participants"` // Список PUUID участников
}

// MatchInfoDTO содержит основную информацию о матче
type MatchInfoDTO struct {
	GameCreation     int64            `json:"gameCreation"`     // Timestamp создания игры
	GameDuration     int              `json:"gameDuration"`     // Длительность в секундах
	GameEndTimestamp int64            `json:"gameEndTimestamp"` // Timestamp окончания
	GameID           int64            `json:"gameId"`           // ID игры
	GameMode         string           `json:"gameMode"`         // Режим (CLASSIC, ARAM)
	GameName         string           `json:"gameName"`         // Название игры
	GameType         string           `json:"gameType"`         // Тип (MATCHED_GAME)
	GameVersion      string           `json:"gameVersion"`      // Версия патча
	MapID            int              `json:"mapId"`            // ID карты (11 = Summoner's Rift)
	PlatformID       string           `json:"platformId"`       // Платформа (RU, EUW1)
	QueueID          int              `json:"queueId"`          // ID очереди (420 = Ranked Solo)
	TournamentCode   string           `json:"tournamentCode"`   // Код турнира (если есть)
	Participants     []ParticipantDTO `json:"participants"`     // Участники матча
	Teams            []TeamDTO        `json:"teams"`            // Команды
}

// ParticipantDTO представляет участника матча
type ParticipantDTO struct {
	// Основная информация
	PUUID         string `json:"puuid"`
	SummonerID    string `json:"summonerId"`
	SummonerName  string `json:"summonerName"`
	SummonerLevel int    `json:"summonerLevel"`
	ProfileIcon   int    `json:"profileIcon"`

	// Информация о чемпионе
	ChampionID      int    `json:"championId"`
	ChampionName    string `json:"championName"`
	ChampExperience int    `json:"champExperience"`
	ChampLevel      int    `json:"champLevel"`

	// Команда и результат
	TeamID       int    `json:"teamId"`       // 100 = синие, 200 = красные
	TeamPosition string `json:"teamPosition"` // Позиция (TOP, JUNGLE, MIDDLE, BOTTOM, UTILITY)
	Win          bool   `json:"win"`          // Победа/поражение

	// Статистика KDA
	Kills   int     `json:"kills"`
	Deaths  int     `json:"deaths"`
	Assists int     `json:"assists"`
	KDA     float64 `json:"challenges.kda,omitempty"` // Может быть в challenges

	// Мультикиллы
	DoubleKills           int `json:"doubleKills"`
	TripleKills           int `json:"tripleKills"`
	QuadraKills           int `json:"quadraKills"`
	PentaKills            int `json:"pentaKills"`
	LargestMultiKill      int `json:"largestMultiKill"`
	LargestKillingSpree   int `json:"largestKillingSpree"`
	LargestCriticalStrike int `json:"largestCriticalStrike"`

	// Урон
	TotalDamageDealtToChampions int `json:"totalDamageDealtToChampions"`
	TotalDamageTaken            int `json:"totalDamageTaken"`
	TotalDamageDealt            int `json:"totalDamageDealt"`
	DamageDealtToBuildings      int `json:"damageDealtToBuildings"`
	DamageDealtToObjectives     int `json:"damageDealtToObjectives"`
	DamageDealtToTurrets        int `json:"damageDealtToTurrets"`
	PhysicalDamageDealt         int `json:"physicalDamageDealt"`
	MagicDamageDealt            int `json:"magicDamageDealt"`
	TrueDamageDealt             int `json:"trueDamageDealt"`
	PhysicalDamageTaken         int `json:"physicalDamageTaken"`
	MagicDamageTaken            int `json:"magicDamageTaken"`
	TrueDamageTaken             int `json:"trueDamageTaken"`
	DamageSelfMitigated         int `json:"damageSelfMitigated"`

	// Защита и лечение
	TotalHeal                      int `json:"totalHeal"`
	TotalHealsOnTeammates          int `json:"totalHealsOnTeammates"`
	TotalDamageShieldedOnTeammates int `json:"totalDamageShieldedOnTeammates"`
	TotalTimeSpentDead             int `json:"totalTimeSpentDead"`

	// Золото
	GoldEarned int `json:"goldEarned"`
	GoldSpent  int `json:"goldSpent"`

	// Миньоны и лес
	TotalMinionsKilled            int `json:"totalMinionsKilled"`
	NeutralMinionsKilled          int `json:"neutralMinionsKilled"`
	TotalAllyJungleMinionsKilled  int `json:"totalAllyJungleMinionsKilled"`
	TotalEnemyJungleMinionsKilled int `json:"totalEnemyJungleMinionsKilled"`

	// Видение
	VisionScore             int `json:"visionScore"`
	VisionWardsBoughtInGame int `json:"visionWardsBoughtInGame"`
	WardsPlaced             int `json:"wardsPlaced"`
	WardsKilled             int `json:"wardsKilled"`
	SightWardsBoughtInGame  int `json:"sightWardsBoughtInGame"`
	DetectorWardsPlaced     int `json:"detectorWardsPlaced"`

	// Башни и ингибиторы
	TurretKills        int `json:"turretKills"`
	TurretTakedowns    int `json:"turretTakedowns"`
	InhibitorKills     int `json:"inhibitorKills"`
	InhibitorTakedowns int `json:"inhibitorTakedowns"`

	// Драконы и бароны
	DragonKills     int `json:"dragonKills"`
	BaronKills      int `json:"baronKills"`
	RiftHeraldKills int `json:"riftHeraldKills"`
	VoidGrubKills   int `json:"voidGrubKills"` // Новые объекты

	// Первая кровь и башня
	FirstBloodKill   bool `json:"firstBloodKill"`
	FirstBloodAssist bool `json:"firstBloodAssist"`
	FirstTowerKill   bool `json:"firstTowerKill"`
	FirstTowerAssist bool `json:"firstTowerAssist"`

	// Предметы
	Item0 int `json:"item0"`
	Item1 int `json:"item1"`
	Item2 int `json:"item2"`
	Item3 int `json:"item3"`
	Item4 int `json:"item4"`
	Item5 int `json:"item5"`
	Item6 int `json:"item6"`

	// Заклинания призывателя
	Summoner1ID    int `json:"summoner1Id"`
	Summoner2ID    int `json:"summoner2Id"`
	Summoner1Casts int `json:"summoner1Casts"`
	Summoner2Casts int `json:"summoner2Casts"`

	// Руны
	Perks PerksDTO `json:"perks"`

	// Пинг
	Ping int `json:"ping"`

	// Роль и позиция (определенные игрой)
	Role               string `json:"role"`               // SOLO, DUO, CARRY, SUPPORT
	Lane               string `json:"lane"`               // TOP, JUNGLE, MIDDLE, BOTTOM
	IndividualPosition string `json:"individualPosition"` // TOP, JUNGLE, MIDDLE, BOTTOM, UTILITY

	// Прогресс выполнения заданий
	BountyLevel               int  `json:"bountyLevel"`
	ConsumablesPurchased      int  `json:"consumablesPurchased"`
	GameEndedInEarlySurrender bool `json:"gameEndedInEarlySurrender"`
	GameEndedInSurrender      bool `json:"gameEndedInSurrender"`
	LongestTimeSpentLiving    int  `json:"longestTimeSpentLiving"`

	// Challenges (достижения)
	Challenges map[string]interface{} `json:"challenges,omitempty"`
}

// TeamDTO представляет команду в матче
type TeamDTO struct {
	TeamID     int           `json:"teamId"`
	Win        bool          `json:"win"`
	Bans       []BanDTO      `json:"bans"`
	Objectives ObjectivesDTO `json:"objectives"`
}

// BanDTO представляет бан чемпиона
type BanDTO struct {
	ChampionID int `json:"championId"`
	PickTurn   int `json:"pickTurn"`
}

// ObjectivesDTO представляет объекты, взятые командой
type ObjectivesDTO struct {
	Baron      ObjectiveDTO `json:"baron"`
	Champion   ObjectiveDTO `json:"champion"`
	Dragon     ObjectiveDTO `json:"dragon"`
	Horde      ObjectiveDTO `json:"horde"` // Void Grubs
	Inhibitor  ObjectiveDTO `json:"inhibitor"`
	RiftHerald ObjectiveDTO `json:"riftHerald"`
	Tower      ObjectiveDTO `json:"tower"`
}

// ObjectiveDTO представляет конкретный объект
type ObjectiveDTO struct {
	First bool `json:"first"`
	Kills int  `json:"kills"`
}

// PerksDTO представляет руны участника
type PerksDTO struct {
	StatPerks StatPerksDTO   `json:"statPerks"`
	Styles    []PerkStyleDTO `json:"styles"`
}

// StatPerksDTO представляет статы от рун
type StatPerksDTO struct {
	Defense int `json:"defense"`
	Flex    int `json:"flex"`
	Offense int `json:"offense"`
}

// PerkStyleDTO представляет стиль рун (ветку)
type PerkStyleDTO struct {
	Description string             `json:"description"`
	Selections  []PerkSelectionDTO `json:"selections"`
	Style       int                `json:"style"`
}

// PerkSelectionDTO представляет выбранную руну
type PerkSelectionDTO struct {
	Perk int `json:"perk"`
	Var1 int `json:"var1"`
	Var2 int `json:"var2"`
	Var3 int `json:"var3"`
}

// ============================================================================
// MATCH TIMELINE DTOs (Глобальные эндпоинты)
// ============================================================================

// MatchTimelineDTO представляет временную шкалу матча
type MatchTimelineDTO struct {
	Metadata MatchTimelineMetadataDTO `json:"metadata"`
	Info     MatchTimelineInfoDTO     `json:"info"`
}

// MatchTimelineMetadataDTO содержит метаданные таймлайна
type MatchTimelineMetadataDTO struct {
	DataVersion  string   `json:"dataVersion"`
	MatchID      string   `json:"matchId"`
	Participants []string `json:"participants"`
}

// MatchTimelineInfoDTO содержит информацию о таймлайне
type MatchTimelineInfoDTO struct {
	FrameInterval int                `json:"frameInterval"`
	Frames        []TimelineFrameDTO `json:"frames"`
}

// TimelineFrameDTO представляет кадр временной шкалы
type TimelineFrameDTO struct {
	Timestamp         int                            `json:"timestamp"`
	ParticipantFrames map[string]ParticipantFrameDTO `json:"participantFrames"`
	Events            []TimelineEventDTO             `json:"events"`
}

// ParticipantFrameDTO представляет состояние участника на конкретном кадре
type ParticipantFrameDTO struct {
	ChampionStats            map[string]int `json:"championStats"`
	CurrentGold              int            `json:"currentGold"`
	DamageStats              map[string]int `json:"damageStats"`
	GoldPerSecond            int            `json:"goldPerSecond"`
	JungleMinionsKilled      int            `json:"jungleMinionsKilled"`
	Level                    int            `json:"level"`
	MinionsKilled            int            `json:"minionsKilled"`
	ParticipantID            int            `json:"participantId"`
	Position                 PositionDTO    `json:"position"`
	TimeEnemySpentControlled int            `json:"timeEnemySpentControlled"`
	TotalGold                int            `json:"totalGold"`
	XP                       int            `json:"xp"`
}

// PositionDTO представляет позицию на карте
type PositionDTO struct {
	X int `json:"x"`
	Y int `json:"y"`
}

// TimelineEventDTO представляет событие на временной шкале
type TimelineEventDTO struct {
	Type      string `json:"type"`
	Timestamp int    `json:"timestamp"`

	// Для событий убийств
	KillerID                int    `json:"killerId,omitempty"`
	VictimID                int    `json:"victimId,omitempty"`
	AssistingParticipantIDs []int  `json:"assistingParticipantIds,omitempty"`
	KillStreakLength        int    `json:"killStreakLength,omitempty"`
	KillType                string `json:"killType,omitempty"`

	// Для событий покупки предметов
	ItemID        int `json:"itemId,omitempty"`
	ParticipantID int `json:"participantId,omitempty"`

	// Для событий прокачки скиллов
	SkillSlot   int    `json:"skillSlot,omitempty"`
	LevelUpType string `json:"levelUpType,omitempty"`

	// Для событий вардов
	WardType  string `json:"wardType,omitempty"`
	CreatorID int    `json:"creatorId,omitempty"`

	// Для событий уничтожения строений
	BuildingType string      `json:"buildingType,omitempty"`
	LaneType     string      `json:"laneType,omitempty"`
	Position     PositionDTO `json:"position,omitempty"`
	TeamID       int         `json:"teamId,omitempty"`

	// Для событий монстров
	MonsterType    string `json:"monsterType,omitempty"`
	MonsterSubType string `json:"monsterSubType,omitempty"`

	// Для событий игры
	GameID      int64 `json:"gameId,omitempty"`
	WinningTeam int   `json:"winningTeam,omitempty"`

	// Для событий пауз
	RealTimestamp int64 `json:"realTimestamp,omitempty"`
}

// ============================================================================
// SPECTATOR API DTOs (Региональные эндпоинты)
// ============================================================================

// CurrentGameInfoDTO представляет информацию о текущей игре
type CurrentGameInfoDTO struct {
	GameID                  int64                       `json:"gameId"`
	GameType                string                      `json:"gameType"`
	GameStartTime           int64                       `json:"gameStartTime"`
	MapID                   int                         `json:"mapId"`
	GameLength              int                         `json:"gameLength"`
	PlatformID              string                      `json:"platformId"`
	GameMode                string                      `json:"gameMode"`
	BannedChampions         []BannedChampionDTO         `json:"bannedChampions"`
	CurrentGameParticipants []CurrentGameParticipantDTO `json:"participants"`
	Observers               ObserverDTO                 `json:"observers"`
	GameQueueConfigID       int                         `json:"gameQueueConfigId"`
}

// CurrentGameParticipantDTO представляет участника текущей игры
type CurrentGameParticipantDTO struct {
	ChampionID               int                          `json:"championId"`
	Perks                    PerksDTO                     `json:"perks"`
	ProfileIconID            int                          `json:"profileIconId"`
	Bot                      bool                         `json:"bot"`
	TeamID                   int                          `json:"teamId"`
	SummonerName             string                       `json:"summonerName"`
	SummonerID               string                       `json:"summonerId"`
	Spell1ID                 int                          `json:"spell1Id"`
	Spell2ID                 int                          `json:"spell2Id"`
	GameCustomizationObjects []GameCustomizationObjectDTO `json:"gameCustomizationObjects"`
}

// BannedChampionDTO представляет забаненного чемпиона
type BannedChampionDTO struct {
	ChampionID int `json:"championId"`
	TeamID     int `json:"teamId"`
	PickTurn   int `json:"pickTurn"`
}

// ObserverDTO представляет наблюдателей игры
type ObserverDTO struct {
	EncryptionKey string `json:"encryptionKey"`
}

// GameCustomizationObjectDTO представляет кастомизацию в игре
type GameCustomizationObjectDTO struct {
	Category string `json:"category"`
	Content  string `json:"content"`
}

// ============================================================================
// CHAMPION MASTERY DTOs (Региональные эндпоинты)
// ============================================================================

// ChampionMasteryDTO представляет мастерство на чемпионе
type ChampionMasteryDTO struct {
	PUUID                        string `json:"puuid"`
	ChampionID                   int    `json:"championId"`
	ChampionLevel                int    `json:"championLevel"`
	ChampionPoints               int    `json:"championPoints"`
	LastPlayTime                 int64  `json:"lastPlayTime"`
	ChampionPointsSinceLastLevel int    `json:"championPointsSinceLastLevel"`
	ChampionPointsUntilNextLevel int    `json:"championPointsUntilNextLevel"`
	MarkRequiredForNextLevel     int    `json:"markRequiredForNextLevel"`
	TokensEarned                 int    `json:"tokensEarned"`
	ChestGranted                 bool   `json:"chestGranted"`
}

// ============================================================================
// CHAMPION DATA DTOs (Data Dragon)
// ============================================================================

// ChampionDataDTO представляет данные о чемпионах из Data Dragon
type ChampionDataDTO struct {
	Type    string                 `json:"type"`
	Format  string                 `json:"format"`
	Version string                 `json:"version"`
	Data    map[string]ChampionDTO `json:"data"`
}

// ChampionDTO представляет информацию о чемпионе
type ChampionDTO struct {
	ID      string           `json:"id"`
	Key     string           `json:"key"`
	Name    string           `json:"name"`
	Title   string           `json:"title"`
	Blurb   string           `json:"blurb"`
	Image   ImageDTO         `json:"image"`
	Tags    []string         `json:"tags"`
	Partype string           `json:"partype"`
	Stats   ChampionStatsDTO `json:"stats"`
}

// ChampionStatsDTO представляет базовые статы чемпиона
type ChampionStatsDTO struct {
	HP                   float64 `json:"hp"`
	HPPerLevel           float64 `json:"hpperlevel"`
	MP                   float64 `json:"mp"`
	MPPerLevel           float64 `json:"mpperlevel"`
	MoveSpeed            float64 `json:"movespeed"`
	Armor                float64 `json:"armor"`
	ArmorPerLevel        float64 `json:"armorperlevel"`
	SpellBlock           float64 `json:"spellblock"`
	SpellBlockPerLevel   float64 `json:"spellblockperlevel"`
	AttackRange          float64 `json:"attackrange"`
	HPRegen              float64 `json:"hpregen"`
	HPRegenPerLevel      float64 `json:"hpregenperlevel"`
	MPRegen              float64 `json:"mpregen"`
	MPRegenPerLevel      float64 `json:"mpregenperlevel"`
	Crit                 float64 `json:"crit"`
	CritPerLevel         float64 `json:"critperlevel"`
	AttackDamage         float64 `json:"attackdamage"`
	AttackDamagePerLevel float64 `json:"attackdamageperlevel"`
	AttackSpeed          float64 `json:"attackspeed"`
	AttackSpeedPerLevel  float64 `json:"attackspeedperlevel"`
}

// ImageDTO представляет информацию об изображении
type ImageDTO struct {
	Full   string `json:"full"`
	Sprite string `json:"sprite"`
	Group  string `json:"group"`
	X      int    `json:"x"`
	Y      int    `json:"y"`
	W      int    `json:"w"`
	H      int    `json:"h"`
}

// ============================================================================
// Вспомогательные методы
// ============================================================================

// GetWinRate возвращает процент побед
func (l *LeagueEntryDTO) GetWinRate() float64 {
	total := l.Wins + l.Losses
	if total == 0 {
		return 0
	}
	return float64(l.Wins) / float64(total) * 100
}

// GetTotalGames возвращает общее количество игр
func (l *LeagueEntryDTO) GetTotalGames() int {
	return l.Wins + l.Losses
}

// GetKDA возвращает KDA участника
func (p *ParticipantDTO) GetKDA() float64 {
	if p.Deaths == 0 {
		return float64(p.Kills + p.Assists)
	}
	return float64(p.Kills+p.Assists) / float64(p.Deaths)
}

// IsIron4 проверяет, находится ли игрок в ранге Iron IV
func (l *LeagueEntryDTO) IsIron4() bool {
	return l.Tier == "IRON" && l.Rank == "IV"
}

// IsRankedSolo проверяет, является ли очередь ранговой соло/дуо
func (l *LeagueEntryDTO) IsRankedSolo() bool {
	return l.QueueType == "RANKED_SOLO_5x5"
}

// GetRankString возвращает строковое представление ранга
func (l *LeagueEntryDTO) GetRankString() string {
	return l.Tier + " " + l.Rank
}

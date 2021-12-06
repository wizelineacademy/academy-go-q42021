package models

type ChampionsApiResponse struct {
	ChampionsData map[string]ChampionData `json:"data"`
}

type ChampionData struct {
	Version string            `json:"version"`
	ID      string            `json:"id"`
	Key     string            `json:"key"`
	Name    string            `json:"name"`
	Title   string            `json:"title"`
	Blurb   string            `json:"blurb"`
	Info    ChampionDataInfo  `json:"info"`
	Image   ImageData         `json:"image"`
	Tags    []string          `json:"tags"`
	Partype string            `json:"partype"`
	Stats   ChampionDataStats `json:"stats"`
}

type ChampionDataInfo struct {
	Attack     int `json:"attack"`
	Defense    int `json:"defense"`
	Magic      int `json:"magic"`
	Difficulty int `json:"difficulty"`
}

type ImageData struct {
	Full   string `json:"full"`
	Sprite string `json:"sprite"`
	Group  string `json:"group"`
	X      int    `json:"x"`
	Y      int    `json:"y"`
	W      int    `json:"w"`
	H      int    `json:"h"`
}

type ChampionDataStats struct {
	HealthPoints                    float64 `json:"hp"`
	HealthPointsPerLevel            float64 `json:"hpperlevel"`
	ManaPoints                      float64 `json:"mp"`
	ManaPointsPerLevel              float64 `json:"mpperlevel"`
	MovementSpeed                   float64 `json:"movespeed"`
	Armor                           float64 `json:"armor"`
	ArmorPerLevel                   float64 `json:"armorperlevel"`
	SpellBlock                      float64 `json:"spellblock"`
	SpellBlockPerLevel              float64 `json:"spellblockperlevel"`
	AttackRange                     float64 `json:"attackrange"`
	HealthPointRegeneration         float64 `json:"hpregen"`
	HealthPointRegenerationPerLevel float64 `json:"hpregenperlevel"`
	ManaPointRegeneration           float64 `json:"mpregen"`
	ManaPointRegenerationPerLevel   float64 `json:"mpregenperlevel"`
	CriticalStrikeChance            float64 `json:"crit"`
	CriticalStrikeChancePerLevel    float64 `json:"critperlevel"`
	AttackDamage                    float64 `json:"attackdamage"`
	AttackDamagePerLevel            float64 `json:"attackdamageperlevel"`
	AttackSpeedOffset               float64 `json:"attackspeedoffset"`
	AttackSpeedPerLevel             float64 `json:"attackspeedperlevel"`
}
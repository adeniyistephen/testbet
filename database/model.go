package database

type SportKey string

type SportTitle string

type MatchId string

type CommenceTime string

type HomeTeam string

type AwayTeam string

type Sport struct {
	Key           SportKey   `json:"key"`
	Active        bool       `json:"active"`
	Group         string     `json:"group"`
	Description   string     `json:"description"`
	Title         SportTitle `json:"title"`
	Has_Outrights bool       `json:"has_outrights"`
}

type Bookmaker struct {
	Key         string   `json:"key"`
	Title       string   `json:"title"`
	Last_Update string   `json:"last_update"`
	Markets     []Market `json:"markets"`
}

type Market struct {
	Key      string    `json:"key"`
	Outcomes []Outcome `json:"outcomes"`
}

type Outcome struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type UpcomingGames struct {
	GeneratedId   string       `json:"generatedid"`
	Id            MatchId      `json:"id"`
	Sport_Key     SportKey     `json:"sport_key"`
	Sport_Title   SportTitle   `json:"sport_title"`
	Commence_Time CommenceTime `json:"commence_time"`
	Home_Team     HomeTeam     `json:"home_team"`
	Away_Team     AwayTeam     `json:"away_team"`
	Bookmakers    []Bookmaker  `json:"bookmakers"`
}

type In_Play struct {
	InPlay UpcomingGames `json:"in_play"`
}

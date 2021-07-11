package entity

type Stats struct {
	PlayerName           string  `json:"player_name,omitempty"`
	Minutes              float32 `json:"minutes,omitempty"`
	FieldGoals           int32   `json:"field_goals,omitempty"`
	FieldGoalAttempts    int32   `json:"field_goal_attempts,omitempty"`
	ThreePointersMade    int32   `json:"three_pointers_made,omitempty"`
	ThreePointerAttempts int32   `json:"three_pointer_attempts,omitempty"`
	FreeThrowsMade       int32   `json:"free_throws_made,omitempty"`
	FreeThrowAttempts    int32   `json:"free_throw_attempts,omitempty"`
	OffensiveRebounds    int32   `json:"offensive_rebounds,omitempty"`
	DefensiveRebounds    int32   `json:"defensive_rebounds,omitempty"`
	Assists              int32   `json:"assists,omitempty"`
	Steals               int32   `json:"steals,omitempty"`
	Blocks               int32   `json:"blocks,omitempty"`
	Turnovers            int32   `json:"turnovers,omitempty"`
	PersonalFouls        int32   `json:"personal_fouls,omitempty"`
	Points               int32   `json:"points,omitempty"`
	Team                 string  `json:"team,omitempty"`
	Opponent             string  `json:"opponent,omitempty"`
	GameDate             int64   `json:"game_date,omitempty"`
}

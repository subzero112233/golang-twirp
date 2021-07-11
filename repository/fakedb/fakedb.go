package fakedb

import (
	"time"

	"github.com/subzero112233/golang-twirp/entity"
)

type FakeDB struct{}

func NewFakeDB() *FakeDB {
	return &FakeDB{}
}

func (f *FakeDB) Get(player string) ([]entity.Stats, error) {
	var s entity.Stats
	s.PlayerName = "Reggie Jackson"
	s.Minutes = 32.9
	s.FieldGoals = 14
	s.FieldGoalAttempts = 20
	s.ThreePointersMade = 4
	s.ThreePointerAttempts = 7
	s.FreeThrowsMade = 6
	s.FreeThrowAttempts = 6
	s.OffensiveRebounds = 1
	s.DefensiveRebounds = 3
	s.Assists = 6
	s.Steals = 1
	s.Blocks = 0
	s.Turnovers = 2
	s.PersonalFouls = 3
	s.Points = 38
	s.Team = "Los Angeles Clippers"
	s.Opponent = "Utah Jazz"
	s.GameDate = time.Now().Unix()

	st := make([]entity.Stats, 1)
	st[0] = s
	return st, nil
}

func (f *FakeDB) Add(st []entity.Stats) (err error) {
	return nil
}

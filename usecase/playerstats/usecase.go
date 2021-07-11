package playerstats

import "github.com/subzero112233/golang-twirp/entity"

type UseCase interface {
	GetStats(player string) ([]entity.Stats, error)
	AddStats(stat []entity.Stats) (err error)
}

package playerstats

import "github.com/subzero112233/golang-twirp/entity"

type Repository interface {
	Get(player string) ([]entity.Stats, error)
	Add(stat []entity.Stats) error
}

package playerstats

import "github.com/subzero112233/golang-twirp/entity"

func (ser *Service) GetStats(player string) (output []entity.Stats, err error) {
	output, err = ser.Repository.Get(player)
	if err != nil {
		ser.Logger.Error().
			Msgf("failed to get stats with error: %v", err)
		return output, err
	}
	return output, nil
}

func (ser *Service) AddStats(input []entity.Stats) error {
	err := ser.Repository.Add(input)
	if err != nil {
		ser.Logger.Error().
			Msgf("failed to add stats with error: %v", err)
		return err
	}
	return nil
}

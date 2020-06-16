package stats

import "mutant/pkg/storage"

type Stats interface {
	MutantStats() (*MutantStats, error)
}

func NewStats(mutantStorage storage.MutantStorage) Stats {
	return &stats{mutantStorage}
}

type stats struct {
	mutantStorage storage.MutantStorage
}

type MutantStats struct {
	CountMutant int64
	CountHuman  int64
	Ratio       float64
}

func (s *stats) MutantStats() (*MutantStats, error) {
	var mutantCount int64
	var humanCount int64
	var err error

	if mutantCount, err = s.mutantStorage.Count(true); err == nil {
		if humanCount, err = s.mutantStorage.Count(false); err == nil {
			return &MutantStats{
				CountMutant: mutantCount,
				CountHuman:  humanCount,
				Ratio:       s.safeDivideFloat(float64(mutantCount), float64(humanCount)),
			}, nil
		}
	}

	return nil, err
}

func (s *stats) safeDivideFloat(a, b float64) float64 {
	if b == 0 {
		return 0
	}

	return a / b
}

package api

// MutantStats is an API struct used to show stats
type MutantStats struct {
	CountMutantDna int64   `json:"count_mutant_dna"`
	CountHumanDna  int64   `json:"count_human_dna"`
	Ratio          float64 `json:"ratio"`
}

package featurelab

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
)

type TreatmentAssigner struct {
}

func NewTreatmentAssigner() TreatmentAssigner {
	return TreatmentAssigner{}
}

type TreatmentAssignment struct {
	Treatment string `json:"treatment"`
}

func (ta *TreatmentAssigner) GetTreatmentAssignment(feature Feature, criteria string) (TreatmentAssignment, error) {
	hashInput := feature.App() + feature.Name() + criteria

	hashBytes := sha256.Sum256([]byte(hashInput))
	hash := binary.LittleEndian.Uint64(hashBytes[:])

	score := ta.calculateTreatmentAssignmentScore(feature, hash)
	log.Println(fmt.Sprintf("Calculated score for feature %s:%s and criteria %s is: %d",
		feature.App(), feature.Name(), criteria, score))

	for _, allocation := range feature.Allocations() {
		if score < allocation.Weight() {
			return TreatmentAssignment{allocation.Name()}, nil
		}

		score -= allocation.Weight()
	}

	return TreatmentAssignment{}, fmt.Errorf("unable to determine treatment for feature: %s, criteria: %s", feature.Name(), criteria)
}

func (ta *TreatmentAssigner) calculateTreatmentAssignmentScore(f Feature, hash uint64) uint32 {
	// discard higher 32 bits to get a uniformly distributed random number between 0 and 2^32 - 1
	hashLower := hash & 0xFFFFFFFF

	return uint32((hashLower * uint64(f.TotalAllocationsWeight())) / (1 << 32))
}

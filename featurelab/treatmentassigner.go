package featurelab

import (
	"crypto/sha256"
	"encoding/binary"
	"errors"
	"fmt"
	"log"
)

var FeatureOff = TreatmentAssignment{}

var ErrInvalidTreatmentAllocation = errors.New("invalid treatment allocation")

type TreatmentAssigner struct {
}

func NewTreatmentAssigner() TreatmentAssigner {
	return TreatmentAssigner{}
}

type TreatmentAssignment struct {
	Treatment string `json:"treatment"`
}

func (ta *TreatmentAssigner) GetTreatmentAssignment(feature Feature, criteria string) (TreatmentAssignment, error) {
	if feature.Allocations == nil || len(feature.Allocations) == 0 {
		return FeatureOff, nil
	}

	hashInput := feature.App + feature.Name + criteria

	hashBytes := sha256.Sum256([]byte(hashInput))
	hash := binary.LittleEndian.Uint64(hashBytes[:])

	score := ta.calculateTreatmentAssignmentScore(feature, hash)
	log.Println(fmt.Sprintf("Calculated score for Name %s:%s and criteria %s is: %d",
		feature.App, feature.Name, criteria, score))

	for _, allocation := range feature.Allocations {
		if score < allocation.Weight {
			return TreatmentAssignment{allocation.Treatment}, nil
		}

		score -= allocation.Weight
	}

	log.Printf("Invalid treatment allocation for Name %s:%s and criteria %s\n", feature.App, feature.Name, criteria)

	return TreatmentAssignment{}, ErrInvalidTreatmentAllocation
}

func (ta *TreatmentAssigner) calculateTreatmentAssignmentScore(f Feature, hash uint64) uint32 {
	// discard higher 32 bits to get a uniformly distributed random number between 0 and 2^32 - 1
	hashLower := hash & 0xFFFFFFFF

	return uint32((hashLower * uint64(f.TotalAllocationWeight())) / (1 << 32))
}

package main

import (
	"fmt"
	"github.com/torresjeff/go-feature-lab/pkg/featurelab"
	"log"
	"os"
)

func main() {
	allocations := []featurelab.FeatureAllocation{
		featurelab.NewFeatureAllocation("C", 50),
		featurelab.NewFeatureAllocation("T1", 50),
	}
	feature := featurelab.NewFeature("Feature1", allocations)

	treatmentAssigner := featurelab.NewTreatmentAssigner()
	userIds := []string{"123456", "123457", "654321"}

	for _, cid := range userIds {
		treatment, err := treatmentAssigner.GetTreatment(feature, cid)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}

		log.Println(fmt.Sprintf("Treatment for feature %s using key %s is: %s", feature.Name(), cid, treatment))
	}
}

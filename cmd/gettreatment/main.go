package main

import (
	"fmt"
	"github.com/torresjeff/go-feature-lab/featurelab"
	"log"
)

func main() {
	featureLab := featurelab.New("localhost:3000")

	userIds := []string{"123456", "456789", "789123", "123789", "987654", "654321"}
	featureName := "ShowRecommendations"

	for _, userId := range userIds {
		treatment, err := featureLab.GetTreatment("FeatureLab", featureName, userId)
		if err != nil {
			log.Printf(err.Error())
		} else {
			log.Println(fmt.Sprintf("TreatmentName for feature %s using criteria %s is: %+v", featureName, userId, treatment))
		}

	}
}

package main

import (
	"fmt"
	"github.com/torresjeff/go-feature-lab/featurelab"
	"log"
)

func main() {
	featureLab := featurelab.NewFeatureLabClient("http://localhost:3000")

	feature, flError := featureLab.FetchFeature("FeatureLab", "ShowRecommendations")
	if flError != nil {
		log.Printf("error fetching feature: %s", flError)
	} else {
		log.Printf("got features features: %+v", feature)
	}

	features, fetchFeatureError := featureLab.FetchFeatures("Test")
	if fetchFeatureError != nil {
		log.Printf("error fetching features: %s", fetchFeatureError)
	} else {
		log.Printf("got features features: %+v", features)
	}

	userIds := []string{"123456", "456789", "789123", "123789", "987654", "654321"}
	featureName := "ChangeBuyButtonColor"

	for _, userId := range userIds {
		treatment, err := featureLab.GetTreatment("FeatureLab", featureName, userId)
		if err != nil {
			log.Printf(err.Error())
		} else {
			log.Println(fmt.Sprintf("Treatment for feature %s using criteria %s is: %+v", featureName, userId, treatment))
		}

	}
}

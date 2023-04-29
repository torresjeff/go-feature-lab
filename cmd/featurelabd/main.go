package main

import (
	"fmt"
	"github.com/torresjeff/go-feature-lab/featurelab"
	"log"
	"time"
)

func main() {
	featureLab := featurelab.NewCacheableFeatureLab(10*time.Minute, 30*time.Minute)

	// Initial fetch to cache features
	_, err := featureLab.FetchFeatures()
	// TODO: error handling, retries
	if err != nil {
		log.Fatal(err)
	}

	userIds := []string{"123456", "123457", "654321"}
	featureName := "ChangeBuyButtonColor"

	for _, userId := range userIds {
		treatment, err := featureLab.GetTreatment(featureName, userId)
		if err != nil {
			log.Printf(err.Error())
		} else {
			log.Println(fmt.Sprintf("Treatment for feature %s using criteria %s is: %s", featureName, userId, treatment))
		}

	}
}

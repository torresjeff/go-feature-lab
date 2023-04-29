package main

import (
	"fmt"
	"github.com/torresjeff/go-feature-lab/featurelab"
	"log"
)

func main() {
	featureLab := featurelab.New()

	userIds := []string{"123456", "123457", "654321"}
	featureName := "ShowRecommendations"

	for _, userId := range userIds {
		treatment, err := featureLab.GetTreatment(featureName, userId)
		if err != nil {
			log.Printf(err.Error())
		} else {
			log.Println(fmt.Sprintf("Treatment for feature %s using criteria %s is: %+v", featureName, userId, treatment))
		}

	}
}

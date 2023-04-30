package main

import (
	"flag"
	"fmt"
	"github.com/torresjeff/go-feature-lab/featurelab"
	"log"
	"time"
)

func main() {
	var featureLabHost string
	flag.StringVar(&featureLabHost,
		"featurelab-host",
		"featurelab.com:3000",
		"URL where Feature Lab server is located. Eg: localhost:3000")
	flag.StringVar(&featureLabHost,
		"f",
		"featurelab.com:3000",
		"URL where Feature Lab server is located. Eg: localhost:3000")
	flag.Parse()

	featureLab := featurelab.NewCacheableFeatureLab(featureLabHost, 10*time.Minute, 30*time.Minute)

	app := "FeatureLab"
	// Initial fetch to cache features
	_, err := featureLab.FetchFeatures(app)
	// TODO: error handling, retries
	if err != nil {
		log.Fatal(err)
	}

	userIds := []string{"123456", "123457", "654321", "456789", "987654", "123321", "741852", "852963", "369147", "258963"}
	featureName := "ChangeBuyButtonColor"

	for _, userId := range userIds {
		treatment, err := featureLab.GetTreatment(app, featureName, userId)
		if err != nil {
			log.Printf(err.Error())
		} else {
			log.Println(fmt.Sprintf("Treatment for feature %s using criteria %s is: %+v", featureName, userId, treatment))
		}

	}
}

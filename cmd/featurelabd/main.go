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
		"featurelab.com:3000", // requires modification of /etc/hosts file: 127.0.0.1 featurelab.com
		"URL where Feature Lab server is located. Eg: localhost:3000")
	flag.StringVar(&featureLabHost,
		"f",
		"featurelab.com:3000", // requires modification of /etc/hosts file: 127.0.0.1 featurelab.com
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

	userIds := []string{"123456", "456789", "789123", "789456", "987654", "654321", "321987", "123789", "741852", "852963"}
	featureName := "ChangeBuyButtonColor"

	for _, userId := range userIds {
		treatment, err := featureLab.GetTreatment(app, featureName, userId)
		if err != nil {
			log.Printf(err.Error())
		} else {
			log.Println(fmt.Sprintf("TreatmentName for feature %s using criteria %s is: %+v", featureName, userId, treatment))
		}

	}
}

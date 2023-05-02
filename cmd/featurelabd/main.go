package main

import (
	"flag"
	"fmt"
	"github.com/torresjeff/go-feature-lab/featurelab"
	"google.golang.org/grpc"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	var featureLabHost string
	flag.StringVar(&featureLabHost,
		"featurelab-host",
		"featurelab.com:3000", // requires modification of /etc/hosts file: 127.0.0.1 featurelab.com
		"URL where Name Lab server is located. Eg: localhost:3000")
	flag.StringVar(&featureLabHost,
		"f",
		"featurelab.com:3000", // requires modification of /etc/hosts file: 127.0.0.1 featurelab.com
		"URL where Name Lab server is located. Eg: localhost:3000")
	flag.Parse()

	featureLab, conn, err := featurelab.NewFeatureLabDaemonClient(43743, "FeatureLab")
	if err != nil {
		log.Fatal(err)
	}

	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(conn)

	app := "FeatureLab"

	userIds := []string{"123456", "456789", "789123", "789456", "987654", "654321", "321987", "123789", "741852", "852963"}
	featureName := "ChangeBuyButtonColor"

	for _, userId := range userIds {
		treatment, flError := featureLab.GetTreatment(app, featureName, userId)
		log.Println(fmt.Sprintf("treatment %+v", treatment))
		if flError != nil {
			log.Printf(flError.Error())
		} else {
			log.Println(fmt.Sprintf("Treatment for feature %s using criteria %s is: %+v", featureName, userId, treatment))
		}

	}
}

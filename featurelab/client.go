package featurelab

import (
	"fmt"
)

type Client struct {
	featureLabHost string
}

func (c *Client) GetTreatment(app string, featureName string, criteria string) (TreatmentAssignment, error) {
	// TODO: make call to FeatureLab server

	feature, err := c.FetchFeature(app, featureName)
	if err != nil {
		return TreatmentAssignment{}, err
	}

	assigner := NewTreatmentAssigner()
	return assigner.GetTreatmentAssignment(feature, criteria)
}

func (c *Client) FetchFeature(app string, featureName string) (Feature, error) {
	// TODO: make call to FeatureLab server

	if app == "FeatureLab" && featureName == "ShowRecommendations" {
		return Feature{
			App:  "FeatureLab",
			Name: "ShowRecommendations",
			Allocations: []FeatureAllocation{
				NewFeatureAllocation("C", 10),
				NewFeatureAllocation("T1", 10),
				NewFeatureAllocation("T2", 10),
			}}, nil
	} else if app == "FeatureLab" && featureName == "ChangeBuyButtonColor" {
		return Feature{
			App:  "FeatureLab",
			Name: "ChangeBuyButtonColor",
			Allocations: []FeatureAllocation{
				NewFeatureAllocation("C", 32),
				NewFeatureAllocation("T1", 68),
			}}, nil
	}
	return Feature{}, fmt.Errorf("Name %s or app %s doesn't exist", featureName, app)
}

func (c *Client) FetchFeatures(app string) ([]Feature, error) {
	// TODO: make call to FeatureLab server
	if app != "FeatureLab" {
		return nil, fmt.Errorf("app %s doesn't exist", app)
	}

	return []Feature{
		Feature{
			App:  "FeatureLab",
			Name: "ShowRecommendations",
			Allocations: []FeatureAllocation{
				NewFeatureAllocation("C", 10),
				NewFeatureAllocation("T1", 10),
				NewFeatureAllocation("T2", 10),
			}},
		Feature{
			App:  "FeatureLab",
			Name: "ChangeBuyButtonColor",
			Allocations: []FeatureAllocation{
				NewFeatureAllocation("C", 32),
				NewFeatureAllocation("T1", 68),
			}},
	}, nil
}

func NewClient(featureLabHost string) *Client {
	return &Client{
		featureLabHost: featureLabHost,
	}
}

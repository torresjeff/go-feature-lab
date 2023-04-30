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
		return &feature{
			AppName:     "FeatureLab",
			FeatureName: "ShowRecommendations",
			TreatmentAllocations: []FeatureAllocation{
				NewFeatureAllocation("C", 10),
				NewFeatureAllocation("T1", 10),
				NewFeatureAllocation("T2", 10),
			}}, nil
	} else if app == "FeatureLab" && featureName == "ChangeBuyButtonColor" {
		return &feature{
			AppName:     "FeatureLab",
			FeatureName: "ChangeBuyButtonColor",
			TreatmentAllocations: []FeatureAllocation{
				NewFeatureAllocation("C", 32),
				NewFeatureAllocation("T1", 68),
			}}, nil
	}
	return nil, fmt.Errorf("feature %s or app %s doesn't exist", featureName, app)
}

func (c *Client) FetchFeatures(app string) ([]Feature, error) {
	// TODO: make call to FeatureLab server
	if app != "FeatureLab" {
		return nil, fmt.Errorf("app %s doesn't exist", app)
	}

	return []Feature{
		&feature{
			AppName:     "FeatureLab",
			FeatureName: "ShowRecommendations",
			TreatmentAllocations: []FeatureAllocation{
				NewFeatureAllocation("C", 10),
				NewFeatureAllocation("T1", 10),
				NewFeatureAllocation("T2", 10),
			}},
		&feature{
			AppName:     "FeatureLab",
			FeatureName: "ChangeBuyButtonColor",
			TreatmentAllocations: []FeatureAllocation{
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

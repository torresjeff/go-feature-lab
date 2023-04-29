package featurelab

import (
	"fmt"
)

type Client struct {
}

func (c *Client) GetTreatment(featureName string, criteria string) (TreatmentAssignment, error) {
	// TODO: make call to FeatureLab server

	feature, err := c.FetchFeature(featureName)
	if err != nil {
		return TreatmentAssignment{}, err
	}

	assigner := NewTreatmentAssigner()
	return assigner.GetTreatmentAssignment(feature, criteria)
}

func (c *Client) FetchFeature(featureName string) (Feature, error) {
	// TODO: make call to FeatureLab server

	if featureName == "ShowRecommendations" {
		return &feature{
			FeatureName: "ShowRecommendations",
			TreatmentAllocations: []FeatureAllocation{
				NewFeatureAllocation("C", 10),
				NewFeatureAllocation("T1", 10),
				NewFeatureAllocation("T2", 10),
			}}, nil
	} else if featureName == "ChangeBuyButtonColor" {
		return &feature{
			FeatureName: "ChangeBuyButtonColor",
			TreatmentAllocations: []FeatureAllocation{
				NewFeatureAllocation("C", 32),
				NewFeatureAllocation("T1", 68),
			}}, nil
	}
	return nil, fmt.Errorf("feature %s doesn't exist", featureName)
}

func (c *Client) FetchFeatures() ([]Feature, error) {
	// TODO: make call to FeatureLab server

	return []Feature{
		&feature{
			FeatureName: "ShowRecommendations",
			TreatmentAllocations: []FeatureAllocation{
				NewFeatureAllocation("C", 10),
				NewFeatureAllocation("T1", 10),
				NewFeatureAllocation("T2", 10),
			}},
		&feature{
			FeatureName: "ChangeBuyButtonColor",
			TreatmentAllocations: []FeatureAllocation{
				NewFeatureAllocation("C", 32),
				NewFeatureAllocation("T1", 68),
			}},
	}, nil
}

func NewClient() *Client {
	return &Client{}
}

package featurelab

type Client struct {
}

func (c *Client) FetchFeatures() ([]Feature, error) {
	// TODO: make call to FeatureLab server

	return []Feature{
		&feature{
			name: "ShowRecommendations",
			treatmentAllocations: []FeatureAllocation{
				NewFeatureAllocation("C", 10),
				NewFeatureAllocation("T1", 10),
				NewFeatureAllocation("T2", 10),
			}},
		&feature{
			name: "ChangeBuyButtonColor",
			treatmentAllocations: []FeatureAllocation{
				NewFeatureAllocation("C", 32),
				NewFeatureAllocation("T1", 68),
			}},
	}, nil
}

func NewClient() *Client {
	return &Client{}
}

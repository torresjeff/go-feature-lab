package featurelab

type FeatureLab interface {
	GetTreatment(app string, featureName string, criteria string) (TreatmentAssignment, *Error)
	FetchFeatures(app string) ([]Feature, *Error)
	FetchFeature(app string, featureName string) (Feature, *Error)
}

type FeatureLabClient struct {
	client *Client
}

// NewFeatureLabClient creates an instance of FeatureLab which can make calls to the Feature Lab backend service.
func NewFeatureLabClient(featureLabHost string) FeatureLab {
	// TODO: URL validation
	fl := &FeatureLabClient{
		client: NewClient(featureLabHost),
	}

	return fl
}

// GetTreatment fetches the treatment that is assigned for a criteria in a particular Feature.
func (f *FeatureLabClient) GetTreatment(app string, featureName string, criteria string) (TreatmentAssignment, *Error) {
	return f.client.GetTreatment(app, featureName, criteria)
}

// FetchFeature fetches the Feature information of a Feature from the Feature Lab backend service.
func (f *FeatureLabClient) FetchFeature(app string, featureName string) (Feature, *Error) {
	return f.client.FetchFeature(app, featureName)
}

// FetchFeatures fetches all features from the Feature Lab backend service.
func (f *FeatureLabClient) FetchFeatures(app string) ([]Feature, *Error) {
	return f.client.FetchFeatures(app)
}

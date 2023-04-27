package featurelab

type FeatureLab interface {
	GetTreatment(featureName string, criteria string) (string, error)
	FetchFeatures() ([]Feature, error)
}

type featureLab struct {
	ta               TreatmentAssigner
	featureLabClient *Client
	featureCache     FeatureCache
}

func (f *featureLab) GetTreatment(featureName string, criteria string) (string, error) {
	feat, err := f.featureCache.GetFeature(featureName)
	if err != nil {
		return "", err
	}

	return f.ta.GetTreatment(feat, criteria)
}

// FetchFeatures fetches all features from the Feature Lab backend service.
// The results will be cached, overwriting any existing features with the same name.
func (f *featureLab) FetchFeatures() ([]Feature, error) {
	features, err := f.featureLabClient.FetchFeatures()
	if err != nil {
		return nil, err
	}

	for _, feature := range features {
		f.featureCache.PutFeature(feature.Name(), feature)
	}

	return features, nil
}

// New creates an instance of FeatureLab and makes the initial fetch of all features, caching them.
func New() FeatureLab {
	fl := &featureLab{
		ta:               TreatmentAssigner{},
		featureLabClient: NewClient(),
		featureCache:     NewDefaultFeatureCache(),
	}

	fl.FetchFeatures()

	return fl
}

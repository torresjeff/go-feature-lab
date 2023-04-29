package featurelab

import "log"

type FeatureLab interface {
	GetTreatment(feature string, criteria string) (string, error)
	FetchFeatures() ([]Feature, error)
	FetchFeature(featureName string) (Feature, error)
}

type featureLab struct {
	featureLabClient *Client
}

// New creates an instance of FeatureLab which can make calls to the Feature Lab backend service.
func New() FeatureLab {
	fl := &featureLab{
		featureLabClient: NewClient(),
	}

	return fl
}

// GetTreatment fetches the treatment that is assigned for a criteria in a particular feature.
func (f *featureLab) GetTreatment(featureName string, criteria string) (string, error) {
	return f.featureLabClient.GetTreatment(featureName, criteria)
}

// FetchFeatures fetches all features from the Feature Lab backend service.
func (f *featureLab) FetchFeatures() ([]Feature, error) {
	return f.featureLabClient.FetchFeatures()
}

// FetchFeature fetches the feature information of a feature from the Feature Lab backend service.
func (f *featureLab) FetchFeature(featureName string) (Feature, error) {
	return f.featureLabClient.FetchFeature(featureName)
}

type cacheableFeatureLab struct {
	featureLab
	treatmentAssigner TreatmentAssigner
	featureCache      FeatureCache
}

// GetTreatment calculates the treatment that is assigned for a criteria in a particular feature stored in the cache.
// If the feature is not in the cache, then an error is returned.
// This method won't make calls to the backend service, it will only operate on features stored in cache.
// To force a cache update, you must call FetchFeatures (to cache all features) or FetchFeature (for a specific feature).
func (f *cacheableFeatureLab) GetTreatment(featureName string, criteria string) (string, error) {
	feature, err := f.featureCache.GetFeature(featureName)
	if err != nil {
		return "", err
	}
	log.Printf("Found feature %s in cache, calculating treatment\n", featureName)

	return f.treatmentAssigner.GetTreatment(feature, criteria)
}

// FetchFeatures fetches all features from the Feature Lab backend service and stores them in cache, overwriting any values that already
// exist in the cache with the same feature names.
func (f *cacheableFeatureLab) FetchFeatures() ([]Feature, error) {
	features, err := f.featureLab.FetchFeatures()
	if err != nil {
		return nil, err
	}

	f.featureCache.PutFeatures(features)

	log.Println("Finished fetching features and cached features")

	return features, nil
}

// FetchFeature fetches the feature information of a feature from the Feature Lab backend service and stores it in cache, overwriting any value
// that already exists in the cache with the same feature name.
func (f *cacheableFeatureLab) FetchFeature(featureName string) (Feature, error) {
	feature, err := f.featureLabClient.FetchFeature(featureName)
	if err != nil {
		return nil, err
	}

	f.featureCache.PutFeature(featureName, feature)

	log.Printf("Fetched feature %s and cached it\n", featureName)

	return feature, nil
}

func NewCacheableFeatureLab() FeatureLab {
	return &cacheableFeatureLab{
		featureLab:        featureLab{},
		treatmentAssigner: TreatmentAssigner{},
		featureCache:      NewDefaultFeatureCache(),
	}
}

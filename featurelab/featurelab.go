package featurelab

import (
	"log"
	"time"
)

type FeatureLab interface {
	GetTreatment(app string, featureName string, criteria string) (TreatmentAssignment, error)
	FetchFeatures(app string) ([]Feature, error)
	FetchFeature(app string, featureName string) (Feature, error)
}

type featureLab struct {
	featureLabClient *Client
}

// New creates an instance of FeatureLab which can make calls to the Feature Lab backend service.
func New(featureLabHost string) FeatureLab {
	fl := &featureLab{
		featureLabClient: NewClient(featureLabHost),
	}

	return fl
}

// GetTreatment fetches the treatment that is assigned for a criteria in a particular Feature.
func (f *featureLab) GetTreatment(app string, featureName string, criteria string) (TreatmentAssignment, error) {
	return f.featureLabClient.GetTreatment(app, featureName, criteria)
}

// FetchFeatures fetches all features from the Feature Lab backend service.
func (f *featureLab) FetchFeatures(app string) ([]Feature, error) {
	return f.featureLabClient.FetchFeatures(app)
}

// FetchFeature fetches the Feature information of a Feature from the Feature Lab backend service.
func (f *featureLab) FetchFeature(app string, featureName string) (Feature, error) {
	return f.featureLabClient.FetchFeature(app, featureName)
}

type cacheableFeatureLab struct {
	featureLab
	treatmentAssigner TreatmentAssigner
	featureCache      FeatureCache
}

// GetTreatment calculates the treatment that is assigned for a criteria in a particular Feature stored in the cache.
// If the Feature is not in the cache, then an error is returned.
// This method won't make calls to the backend service, it will only operate on features stored in cache.
// To force a cache update, you must call FetchFeatures (to cache all features) or FetchFeature (for a specific Feature).
func (f *cacheableFeatureLab) GetTreatment(app string, featureName string, criteria string) (TreatmentAssignment, error) {
	feature, err := f.featureCache.GetFeature(app, featureName)
	if err != nil {
		return TreatmentAssignment{}, err
	}

	return f.treatmentAssigner.GetTreatmentAssignment(feature, criteria)
}

// FetchFeatures fetches all features from the Feature Lab backend service and stores them in cache, overwriting any values that already
// exist in the cache with the same Feature names.
func (f *cacheableFeatureLab) FetchFeatures(app string) ([]Feature, error) {
	features, err := f.featureLab.FetchFeatures(app)
	if err != nil {
		return nil, err
	}

	f.featureCache.PutFeatures(features)

	log.Printf("Finished fetching features and cached features for app: %s\n", app)

	return features, nil
}

// FetchFeature fetches the Feature information of a Feature from the Feature Lab backend service and stores it in cache, overwriting any value
// that already exists in the cache with the same Feature Treatment.
func (f *cacheableFeatureLab) FetchFeature(app string, featureName string) (Feature, error) {
	feature, err := f.featureLabClient.FetchFeature(app, featureName)
	if err != nil {
		return Feature{}, err
	}

	f.featureCache.PutFeature(app, featureName, feature)

	return feature, nil
}

func NewCacheableFeatureLab(featureLabHost string, ttl, cleanUpInterval time.Duration) FeatureLab {
	return &cacheableFeatureLab{
		featureLab: featureLab{
			featureLabClient: NewClient(featureLabHost),
		},
		treatmentAssigner: TreatmentAssigner{},
		featureCache:      NewDefaultFeatureCache(ttl, cleanUpInterval),
	}
}

package featurelab

import (
	"fmt"
	"log"
	"time"
)

type FeatureLab interface {
	GetTreatment(app string, featureName string, criteria string) (TreatmentAssignment, *Error)
	FetchFeatures(app string) ([]Feature, *Error)
	FetchFeature(app string, featureName string) (Feature, *Error)
}

type featureLab struct {
	featureLabClient *Client
}

// New creates an instance of FeatureLab which can make calls to the Feature Lab backend service.
func New(featureLabHost string) FeatureLab {
	// TODO: URL validation
	fl := &featureLab{
		featureLabClient: NewClient(featureLabHost),
	}

	return fl
}

// GetTreatment fetches the treatment that is assigned for a criteria in a particular Feature.
func (f *featureLab) GetTreatment(app string, featureName string, criteria string) (TreatmentAssignment, *Error) {
	return f.featureLabClient.GetTreatment(app, featureName, criteria)
}

// FetchFeatures fetches all features from the Feature Lab backend service.
func (f *featureLab) FetchFeatures(app string) ([]Feature, *Error) {
	return f.featureLabClient.FetchFeatures(app)
}

// FetchFeature fetches the Feature information of a Feature from the Feature Lab backend service.
func (f *featureLab) FetchFeature(app string, featureName string) (Feature, *Error) {
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
func (f *cacheableFeatureLab) GetTreatment(app string, featureName string, criteria string) (TreatmentAssignment, *Error) {
	feature, found := f.featureCache.GetFeature(app, featureName)
	if !found {
		return TreatmentAssignment{}, NewError(ErrNotFound, fmt.Sprintf("Feature %s not found in cache", getCacheKey(app, featureName)))
	}

	assignment, err := f.treatmentAssigner.GetTreatmentAssignment(feature, criteria)
	if err == ErrInvalidTreatmentAllocation {
		return TreatmentAssignment{}, NewError(ErrInvalidTreatment,
			fmt.Sprintf("Feature %s has invalid treatment allocation", getCacheKey(app, featureName)))
	}

	return assignment, nil
}

// FetchFeatures fetches all features from the Feature Lab backend service and stores them in cache, overwriting any values that already
// exist in the cache with the same Feature names.
func (f *cacheableFeatureLab) FetchFeatures(app string) ([]Feature, *Error) {
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
func (f *cacheableFeatureLab) FetchFeature(app string, featureName string) (Feature, *Error) {
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

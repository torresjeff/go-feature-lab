package featurelab

import (
	"fmt"
	"github.com/patrickmn/go-cache"
)

type FeatureCache interface {
	GetFeature(name string) (Feature, error)
	PutFeature(name string, feature Feature)
}

type defaultFeatureCache struct {
	cache *cache.Cache
}

func (d *defaultFeatureCache) GetFeature(name string) (Feature, error) {
	if feature, found := d.cache.Get(name); found {
		f, ok := feature.(Feature)
		if !ok {
			return nil, fmt.Errorf("expected to find a Feature in cache, but instead found %v", f)
		}
		return f, nil
	}

	return nil, fmt.Errorf("unable to find Feature %s in cache", name)
}

func (d *defaultFeatureCache) PutFeature(name string, feature Feature) {
	d.cache.Set(name, feature, DefaultConfig.cacheTTL)
}

func NewDefaultFeatureCache() FeatureCache {
	return &defaultFeatureCache{
		cache: cache.New(DefaultConfig.cacheTTL, DefaultConfig.cacheTTL),
	}
}

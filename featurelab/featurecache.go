package featurelab

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"log"
	"time"
)

type FeatureCache interface {
	GetFeature(app, name string) (Feature, bool)
	PutFeature(app, name string, feature Feature)
	PutFeatures(features []Feature)
}

type defaultFeatureCache struct {
	cache *cache.Cache
}

func (d *defaultFeatureCache) GetFeature(app, name string) (Feature, bool) {
	feature, found := d.cache.Get(getCacheKey(app, name))
	if !found {
		return Feature{}, false
	}

	f, ok := feature.(Feature)
	if !ok {
		panic(fmt.Sprintf("expected to find a Feature in cache, but instead found %+v", f))
	}

	log.Printf("Found Name %s in cache\n", getCacheKey(app, name))
	return f, true
}

func (d *defaultFeatureCache) PutFeature(app, name string, feature Feature) {
	d.cache.Set(getCacheKey(app, name), feature, cache.DefaultExpiration)

	log.Printf("Cached Name: %s\n", getCacheKey(app, name))
}

func (d *defaultFeatureCache) PutFeatures(features []Feature) {
	for _, f := range features {
		d.PutFeature(f.App, f.Name, f)
	}
}

func NewDefaultFeatureCache(ttl time.Duration, cleanUpInterval time.Duration) FeatureCache {
	return &defaultFeatureCache{
		cache: cache.New(ttl, cleanUpInterval),
	}
}

func getCacheKey(app, featureName string) string {
	return fmt.Sprintf("%s:%s", app, featureName)
}

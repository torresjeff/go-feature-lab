package featurelab

import (
	"reflect"
	"testing"
	"time"
)

func TestNewDefaultFeatureCache(t *testing.T) {
	duration := 10 * time.Minute

	got := NewDefaultFeatureCache(duration, duration)

	if got == nil {
		t.Errorf("got nil, wantTreatment not nil")
	}

	app := "FeatureLab"
	showRecommendations := "ShowRecommendations"
	allocations := []FeatureAllocation{
		NewFeatureAllocation("C", 10),
		NewFeatureAllocation("T1", 10),
	}

	feature := NewFeature(app, showRecommendations, allocations)
	got.PutFeature(app, showRecommendations, feature)

	if f, err := got.GetFeature(app, showRecommendations); err != nil {
		t.Errorf("got error %s, want feature %+v", err, feature)
	} else if !reflect.DeepEqual(f, feature) {
		t.Errorf("got %+v, want %+v", f, feature)
	}

	nonExistingFeatureName := "NonExistingFeature"
	if f, err := got.GetFeature(app, nonExistingFeatureName); err == nil {
		t.Errorf("got %+v, want error", f)
	}

}

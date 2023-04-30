package featurelab

import (
	"reflect"
	"testing"
)

func TestFeature(t *testing.T) {
	app := "FeatureLab"
	name := "ShowRecommendations"
	allocations := []FeatureAllocation{
		NewFeatureAllocation("C", 10),
		NewFeatureAllocation("T1", 10),
	}

	got := NewFeature(app, name, allocations)

	if got.App() != app {
		t.Errorf("got %s, wantTreatment %s", got.App(), app)
	}
	if got.Name() != name {
		t.Errorf("got %s, wantTreatment %s", got.Name(), name)
	}

	if !reflect.DeepEqual(got.Allocations(), allocations) {
		t.Errorf("got %+v, wantTreatment %+v", got.Allocations(), allocations)
	}

	if got.TotalAllocationWeight() != 20 {
		t.Errorf("got %d, wantTreatment %d", got.TotalAllocationWeight(), 20)
	}
}

func TestFeatureAllocation(t *testing.T) {
	treatment := "C"
	weight := uint32(10)

	got := NewFeatureAllocation(treatment, weight)

	if got.Treatment() != treatment {
		t.Errorf("got %s, wantTreatment %s", got.Treatment(), treatment)
	}
	if got.Weight() != weight {
		t.Errorf("got %d, wantTreatment %d", got.Weight(), weight)
	}
}

package featurelab

type Feature struct {
	AppName              string              `json:"app"`
	FeatureName          string              `json:"featureName"`
	TreatmentAllocations []FeatureAllocation `json:"treatmentAllocations"`
}

func NewFeature(app string, name string, allocations []FeatureAllocation) Feature {
	return Feature{
		AppName:              app,
		FeatureName:          name,
		TreatmentAllocations: allocations,
	}
}

func (f *Feature) Name() string {
	return f.FeatureName
}

func (f *Feature) App() string {
	return f.AppName
}

func (f *Feature) Allocations() []FeatureAllocation {
	return f.TreatmentAllocations
}

func (f *Feature) TotalAllocationWeight() uint32 {
	var sum uint32 = 0

	for _, ta := range f.TreatmentAllocations {
		sum += ta.Weight()
	}

	return sum
}

type FeatureAllocation struct {
	TreatmentName   string `json:"treatment"`
	TreatmentWeight uint32 `json:"weight"`
}

func NewFeatureAllocation(name string, weight uint32) FeatureAllocation {
	return FeatureAllocation{
		TreatmentName:   name,
		TreatmentWeight: weight,
	}
}

func (f *FeatureAllocation) Treatment() string {
	return f.TreatmentName
}

func (f *FeatureAllocation) Weight() uint32 {
	return f.TreatmentWeight
}

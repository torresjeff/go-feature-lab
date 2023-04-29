package featurelab

type Feature interface {
	Name() string
	Allocations() []FeatureAllocation
	TotalAllocationsWeight() uint32
}

type feature struct {
	FeatureName          string              `json:"featureName"`
	TreatmentAllocations []FeatureAllocation `json:"treatmentAllocations"`
}

func NewFeature(name string, allocations []FeatureAllocation) Feature {
	return &feature{
		FeatureName:          name,
		TreatmentAllocations: allocations,
	}
}

func (f *feature) Name() string {
	return f.FeatureName
}

func (f *feature) Allocations() []FeatureAllocation {
	return f.TreatmentAllocations
}

func (f *feature) TotalAllocationsWeight() uint32 {
	var sum uint32 = 0

	for _, ta := range f.TreatmentAllocations {
		sum += ta.Weight()
	}

	return sum
}

type FeatureAllocation interface {
	Name() string
	Weight() uint32
}

type featureAllocation struct {
	Treatment        string `json:"treatment"`
	AllocationWeight uint32 `json:"weight"`
}

func NewFeatureAllocation(name string, weight uint32) FeatureAllocation {
	return &featureAllocation{
		Treatment:        name,
		AllocationWeight: weight,
	}
}

func (f *featureAllocation) Name() string {
	return f.Treatment
}

func (f *featureAllocation) Weight() uint32 {
	return f.AllocationWeight
}

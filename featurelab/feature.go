package featurelab

type Feature interface {
	Name() string
	App() string
	Allocations() []FeatureAllocation
	TotalAllocationsWeight() uint32
}

type feature struct {
	AppName              string              `json:"app"`
	FeatureName          string              `json:"featureName"`
	TreatmentAllocations []FeatureAllocation `json:"treatmentAllocations"`
}

func NewFeature(app string, name string, allocations []FeatureAllocation) Feature {
	return &feature{
		AppName:              app,
		FeatureName:          name,
		TreatmentAllocations: allocations,
	}
}

func (f *feature) Name() string {
	return f.FeatureName
}

func (f *feature) App() string {
	return f.AppName
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

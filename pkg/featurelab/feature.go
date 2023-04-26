package featurelab

type Feature interface {
	Name() string
	Allocations() []FeatureAllocation
	TotalAllocationsWeight() uint32
}

type feature struct {
	name                 string
	treatmentAllocations []FeatureAllocation
}

func NewFeature(name string, allocations []FeatureAllocation) Feature {
	return &feature{
		name:                 name,
		treatmentAllocations: allocations,
	}
}

func (f *feature) Name() string {
	return f.name
}

func (f *feature) Allocations() []FeatureAllocation {
	return f.treatmentAllocations
}

func (f *feature) TotalAllocationsWeight() uint32 {
	var sum uint32 = 0

	for _, ta := range f.treatmentAllocations {
		sum += ta.Weight()
	}

	return sum
}

type FeatureAllocation interface {
	Name() string
	Weight() uint32
}

type featureAllocation struct {
	name   string
	weight uint32
}

func NewFeatureAllocation(name string, weight uint32) FeatureAllocation {
	return &featureAllocation{
		name:   name,
		weight: weight,
	}
}

func (f *featureAllocation) Name() string {
	return f.name
}

func (f *featureAllocation) Weight() uint32 {
	return f.weight
}

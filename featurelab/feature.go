package featurelab

type Feature struct {
	App         string              `json:"app"`
	Name        string              `json:"name"`
	Allocations []FeatureAllocation `json:"allocations"`
}

func NewFeature(app string, name string, allocations []FeatureAllocation) Feature {
	return Feature{
		App:         app,
		Name:        name,
		Allocations: allocations,
	}
}

func (f *Feature) TotalAllocationWeight() uint32 {
	var sum uint32 = 0

	for _, ta := range f.Allocations {
		sum += ta.Weight
	}

	return sum
}

type FeatureAllocation struct {
	Treatment string `json:"treatment"`
	Weight    uint32 `json:"weight"`
}

func NewFeatureAllocation(treatment string, weight uint32) FeatureAllocation {
	return FeatureAllocation{
		Treatment: treatment,
		Weight:    weight,
	}
}

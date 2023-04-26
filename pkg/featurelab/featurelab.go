package featurelab

type FeatureLab interface {
	GetTreatment(featureName string, key string) (string, error)
}

type featureLab struct {
	ta TreatmentAssigner
}

func (f *featureLab) GetTreatment(featureName string, key string) (string, error) {
	// TODO: fetch feature allocation for featureName in the cache (probably a map featureName -> Feature)
	return f.ta.GetTreatment(NewFeature(featureName, nil), key)
}

func GetDefault() FeatureLab {
	return &featureLab{ta: TreatmentAssigner{}}
}

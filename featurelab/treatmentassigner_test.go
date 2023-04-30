package featurelab

import "testing"

func TestTreatmentAssigner_GetTreatmentAssignment(t *testing.T) {

	showRecommendationFeature := NewFeature("FeatureLab",
		"ShowRecommendations",
		[]FeatureAllocation{
			NewFeatureAllocation("C", 10),
			NewFeatureAllocation("T1", 10),
			NewFeatureAllocation("T2", 10),
		})
	changeBuyButtonColorFeature := NewFeature("FeatureLab",
		"ChangeBuyButtonColor",
		[]FeatureAllocation{
			NewFeatureAllocation("C", 32),
			NewFeatureAllocation("T1", 68),
		})
	featureOff := NewFeature("FeatureLab",
		"FeatureOff",
		[]FeatureAllocation{})
	invalidTreatmentFeature := NewFeature("FeatureLab",
		"InvalidTreatments",
		[]FeatureAllocation{
			NewFeatureAllocation("C", 0),
			NewFeatureAllocation("T1", 0),
		})

	tests := []struct {
		feature       Feature
		criteria      string
		wantTreatment TreatmentAssignment
		wantError     error
	}{
		{feature: changeBuyButtonColorFeature, criteria: "123456", wantTreatment: TreatmentAssignment{Treatment: "C"}, wantError: nil},
		{feature: changeBuyButtonColorFeature, criteria: "456789", wantTreatment: TreatmentAssignment{Treatment: "T1"}, wantError: nil},
		{feature: showRecommendationFeature, criteria: "987654", wantTreatment: TreatmentAssignment{Treatment: "C"}, wantError: nil},
		{feature: showRecommendationFeature, criteria: "789123", wantTreatment: TreatmentAssignment{Treatment: "T1"}, wantError: nil},
		{feature: showRecommendationFeature, criteria: "123456", wantTreatment: TreatmentAssignment{Treatment: "T2"}, wantError: nil},
		{feature: featureOff, criteria: "123456", wantTreatment: TreatmentAssignment{}, wantError: nil},
		{feature: invalidTreatmentFeature, criteria: "123456", wantTreatment: TreatmentAssignment{}, wantError: ErrInvalidTreatmentAllocation},
	}

	treatmentAssigner := NewTreatmentAssigner()

	for _, s := range tests {
		gotTreatment, gotError := treatmentAssigner.GetTreatmentAssignment(s.feature, s.criteria)
		if gotTreatment != s.wantTreatment {
			t.Errorf("got %+v, want %+v", gotTreatment, s.wantTreatment)
		}
		if gotError != s.wantError {
			t.Errorf("got %+v, want %+v", gotError, s.wantError)
		}
	}
}

package featurelab

import (
	"context"
	"fmt"
	"github.com/torresjeff/featurelabd/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type FeatureLabDaemonClient struct {
	featureLabClient pb.FeatureLabClient
}

func NewFeatureLabDaemonClient(port uint, apps ...string) (FeatureLab, *grpc.ClientConn, error) {
	// TODO: use apps for initial fetch of features

	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return &FeatureLabDaemonClient{}, nil, err
	}

	return &FeatureLabDaemonClient{
		featureLabClient: pb.NewFeatureLabClient(conn),
	}, conn, nil
}

// GetTreatment fetches the treatment that is assigned for a criteria in a particular Feature.
func (f *FeatureLabDaemonClient) GetTreatment(app string, featureName string, criteria string) (TreatmentAssignment, *Error) {
	treatment, err := f.featureLabClient.GetTreatment(context.Background(), &pb.GetTreatmentRequest{
		App:      app,
		Feature:  featureName,
		Criteria: criteria,
	})

	if err != nil {
		// TODO: translate to featurelab error
		return TreatmentAssignment{}, NewError(ErrInternalServerError, err.Error())
	}

	return TreatmentAssignment{
		Treatment: treatment.Treatment,
	}, nil
}

// FetchFeature fetches the Feature information of a Feature from the Feature Lab backend service.
func (f *FeatureLabDaemonClient) FetchFeature(app string, featureName string) (Feature, *Error) {
	feature, err := f.featureLabClient.FetchFeature(context.Background(), &pb.FetchFeatureRequest{
		App:     app,
		Feature: featureName,
	})

	if err != nil {
		// TODO: translate to featurelab error
		return Feature{}, NewError(ErrInternalServerError, err.Error())
	}

	return toFeature(feature.Feature), nil
}

// FetchFeatures fetches all features from the Feature Lab backend service.
func (f *FeatureLabDaemonClient) FetchFeatures(app string) ([]Feature, *Error) {
	features, err := f.featureLabClient.FetchFeatures(context.Background(), &pb.FetchFeaturesRequest{
		App: app,
	})

	if err != nil {
		// TODO: translate to featurelab error
		return nil, NewError(ErrInternalServerError, err.Error())
	}

	result := make([]Feature, len(features.Features))
	for i, feature := range features.Features {
		result[i] = toFeature(feature)
	}

	return result, nil
}

func toFeature(feature *pb.Feature) Feature {
	return Feature{
		App:         feature.App,
		Name:        feature.Name,
		Allocations: toFeatureAllocations(feature.Allocations),
	}
}

func toFeatureAllocations(allocations []*pb.TreatmentAllocation) []FeatureAllocation {
	featureAllocations := make([]FeatureAllocation, len(allocations))
	for i, allocation := range allocations {
		featureAllocations[i] = FeatureAllocation{
			Treatment: allocation.Treatment,
			Weight:    allocation.Weight,
		}
	}
	return featureAllocations
}

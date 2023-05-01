package featurelab

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"log"
)

type Client struct {
	featureLabHost string
	restClient     *resty.Client
}

func (c *Client) GetTreatment(app string, feature string, criteria string) (TreatmentAssignment, *Error) {
	log.Printf("Fetching treatment for app: %s, criteria: %s\n", getCacheKey(app, feature), criteria)

	resp, err := c.restClient.R().SetPathParams(map[string]string{
		"app":      app,
		"feature":  feature,
		"criteria": criteria,
	}).
		Get("/app/{app}/features/{feature}/treatment/{criteria}")

	if err != nil {
		// TODO: what errors can resty return here?
		return TreatmentAssignment{}, NewError(ErrInternalServerError, err.Error())
	}

	if resp.IsError() {
		var error *Error
		err = json.Unmarshal(resp.Body(), error)
		if err != nil {
			return TreatmentAssignment{}, NewError(ErrInternalServerError, err.Error())
		}
		return TreatmentAssignment{}, error
	}

	var treatment TreatmentAssignment
	err = json.Unmarshal(resp.Body(), &treatment)
	if err != nil {
		return TreatmentAssignment{}, NewError(ErrInternalServerError, err.Error())
	}

	return treatment, nil
}

func (c *Client) FetchFeature(app string, featureName string) (Feature, *Error) {
	log.Printf("Fetching feature %s\n", getCacheKey(app, featureName))
	resp, err := c.restClient.R().SetPathParams(map[string]string{
		"app":     app,
		"feature": featureName,
	}).
		Get("/app/{app}/features/{feature}")

	if err != nil {
		// TODO: what errors can resty return here?
		return Feature{}, NewError(ErrInternalServerError, err.Error())
	}

	if resp.IsError() {
		var error *Error
		err = json.Unmarshal(resp.Body(), error)
		if err != nil {
			return Feature{}, NewError(ErrInternalServerError, err.Error())
		}
		return Feature{}, error
	}

	var feature Feature
	err = json.Unmarshal(resp.Body(), &feature)
	if err != nil {
		return Feature{}, NewError(ErrInternalServerError, err.Error())
	}

	return feature, nil
}

func (c *Client) FetchFeatures(app string) ([]Feature, *Error) {
	log.Printf("Fetching features for app %s\n", app)
	resp, err := c.restClient.R().SetPathParams(map[string]string{
		"app": app,
	}).
		Get("/app/{app}/features")

	if err != nil {
		// TODO: what errors can resty return here?
		return nil, NewError(ErrInternalServerError, err.Error())
	}

	if resp.IsError() {
		var error *Error
		err = json.Unmarshal(resp.Body(), error)
		if err != nil {
			return nil, NewError(ErrInternalServerError, err.Error())
		}
		return nil, error
	}

	var features []Feature
	err = json.Unmarshal(resp.Body(), &features)
	if err != nil {
		return nil, NewError(ErrInternalServerError, err.Error())
	}

	return features, nil
}

func NewClient(featureLabHost string) *Client {
	// TODO: featureLabHost URL validation
	return &Client{
		featureLabHost: featureLabHost,
		restClient:     resty.New().SetBaseURL(featureLabHost + "/api/v1"),
	}
}

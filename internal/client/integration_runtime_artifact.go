package client

import (
	"context"
	"fmt"
)

func IntegrationRuntimeArtifacts(ctx context.Context) ([]IntegrationRuntimeArtifact, error) {
	type responseBody struct {
		Data struct {
			Results []IntegrationRuntimeArtifact `json:"results"`
		} `json:"d"`
	}

	request := GetInstance(ctx).R().
		SetURL("IntegrationRuntimeArtifacts").
		SetResult(&responseBody{})

	response, err := Get(request)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve integration packages: %w", err)
	}

	return response.Result().(*responseBody).Data.Results, nil
}

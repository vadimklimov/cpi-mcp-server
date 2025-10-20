package client

import (
	"context"
	"fmt"
)

func IntegrationPackages(ctx context.Context) ([]IntegrationPackage, error) {
	type responseBody struct {
		Data struct {
			Results []IntegrationPackage `json:"results"`
		} `json:"d"`
	}

	request := GetInstance(ctx).R().
		SetURL("IntegrationPackages").
		SetResult(&responseBody{})

	response, err := Get(request)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve integration packages: %w", err)
	}

	return response.Result().(*responseBody).Data.Results, nil
}

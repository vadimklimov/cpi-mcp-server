package client

import (
	"context"
	"fmt"
)

func ValueMappings(ctx context.Context) ([]ValueMapping, error) {
	type responseBody struct {
		Data struct {
			Results []ValueMapping `json:"results"`
		} `json:"d"`
	}

	request := GetInstance(ctx).R().
		SetURL("ValueMappingDesigntimeArtifacts").
		SetResult(&responseBody{})

	response, err := Get(request)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve value mappings: %w", err)
	}

	return response.Result().(*responseBody).Data.Results, nil
}

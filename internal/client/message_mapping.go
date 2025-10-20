package client

import (
	"context"
	"fmt"
)

func MessageMappings(ctx context.Context) ([]MessageMapping, error) {
	type responseBody struct {
		Data struct {
			Results []MessageMapping `json:"results"`
		} `json:"d"`
	}

	request := GetInstance(ctx).R().
		SetURL("MessageMappingDesigntimeArtifacts").
		SetResult(&responseBody{})

	response, err := Get(request)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve message mappings: %w", err)
	}

	return response.Result().(*responseBody).Data.Results, nil
}

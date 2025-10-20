package client

import (
	"context"
	"errors"
	"fmt"
)

func IntegrationFlows(ctx context.Context) ([]IntegrationFlow, error) {
	integrationPackages, _ := IntegrationPackages(ctx)

	integrationPackageIDs := make([]string, len(integrationPackages))
	for i, integrationPackage := range integrationPackages {
		integrationPackageIDs[i] = integrationPackage.ID
	}

	integrationFlows, errs, sysErr := Run(ctx, integrationPackageIDs, IntegrationFlowsByPackageID)

	if len(errs) > 0 {
		return nil, fmt.Errorf("failed to retrieve integration flows: %w", errors.Join(errs...))
	}

	if sysErr != nil {
		return nil, fmt.Errorf("failed to retrieve integration flows due to system error: %w", sysErr)
	}

	return integrationFlows, nil
}

func IntegrationFlowsByPackageID(ctx context.Context, packageID string) ([]IntegrationFlow, error) {
	type responseBody struct {
		Data struct {
			Results []IntegrationFlow `json:"results"`
		} `json:"d"`
	}

	request := GetInstance(ctx).R().
		SetPathParam("packageID", packageID).
		SetURL("IntegrationPackages('{packageID}')/IntegrationDesigntimeArtifacts").
		SetResult(&responseBody{})

	response, err := Get(request)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve integration flows: %w", err)
	}

	return response.Result().(*responseBody).Data.Results, nil
}

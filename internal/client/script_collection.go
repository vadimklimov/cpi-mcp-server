package client

import (
	"context"
	"errors"
	"fmt"
)

func ScriptCollections(ctx context.Context) ([]ScriptCollection, error) {
	integrationPackages, _ := IntegrationPackages(ctx)

	integrationPackageIDs := make([]string, len(integrationPackages))
	for i, integrationPackage := range integrationPackages {
		integrationPackageIDs[i] = integrationPackage.ID
	}

	scriptCollections, errs, sysErr := Run(ctx, integrationPackageIDs, ScriptCollectionsByPackageID)

	if len(errs) > 0 {
		return nil, fmt.Errorf("failed to retrieve script collections: %w", errors.Join(errs...))
	}

	if sysErr != nil {
		return nil, fmt.Errorf("failed to retrieve script collections due to system error: %w", sysErr)
	}

	return scriptCollections, nil
}

func ScriptCollectionsByPackageID(ctx context.Context, packageID string) ([]ScriptCollection, error) {
	type responseBody struct {
		Data struct {
			Results []ScriptCollection `json:"results"`
		} `json:"d"`
	}

	request := GetInstance(ctx).R().
		SetPathParam("packageID", packageID).
		SetURL("IntegrationPackages('{packageID}')/ScriptCollectionDesigntimeArtifacts").
		SetResult(&responseBody{})

	response, err := Get(request)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve script collections: %w", err)
	}

	return response.Result().(*responseBody).Data.Results, nil
}

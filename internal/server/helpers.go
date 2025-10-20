package server

import (
	"github.com/vadimklimov/cpi-mcp-server/internal/client"
)

func convert[I any, O any](input []I, converter func(I) O) []O {
	output := make([]O, len(input))
	for i, value := range input {
		output[i] = converter(value)
	}

	return output
}

func convertIntegrationPackage(integrationPackage client.IntegrationPackage) IntegrationPackage {
	return IntegrationPackage{
		ID:         integrationPackage.ID,
		Version:    integrationPackage.Version,
		Name:       integrationPackage.Name,
		ShortText:  integrationPackage.ShortText,
		Vendor:     integrationPackage.Vendor,
		Mode:       integrationPackage.Mode,
		CreatedBy:  integrationPackage.CreatedBy,
		CreatedAt:  integrationPackage.CreationDate.String(),
		ModifiedBy: integrationPackage.ModifiedBy,
		ModifiedAt: integrationPackage.ModifiedDate.String(),
	}
}

func convertIntegrationFlow(integrationFlow client.IntegrationFlow) IntegrationFlow {
	return IntegrationFlow{
		ID:          integrationFlow.ID,
		Version:     integrationFlow.Version,
		PackageID:   integrationFlow.PackageID,
		Name:        integrationFlow.Name,
		Description: integrationFlow.Description,
		CreatedBy:   integrationFlow.CreatedBy,
		CreatedAt:   integrationFlow.CreatedAt.String(),
		ModifiedBy:  integrationFlow.ModifiedBy,
		ModifiedAt:  integrationFlow.ModifiedAt.String(),
	}
}

func convertValueMapping(valueMapping client.ValueMapping) ValueMapping {
	return ValueMapping{
		ID:          valueMapping.ID,
		Version:     valueMapping.Version,
		PackageID:   valueMapping.PackageID,
		Name:        valueMapping.Name,
		Description: valueMapping.Description,
	}
}

func convertMessageMapping(messageMapping client.MessageMapping) MessageMapping {
	return MessageMapping{
		ID:          messageMapping.ID,
		Version:     messageMapping.Version,
		PackageID:   messageMapping.PackageID,
		Name:        messageMapping.Name,
		Description: messageMapping.Description,
	}
}

func convertScriptCollection(scriptCollection client.ScriptCollection) ScriptCollection {
	return ScriptCollection{
		ID:          scriptCollection.ID,
		Version:     scriptCollection.Version,
		PackageID:   scriptCollection.PackageID,
		Name:        scriptCollection.Name,
		Description: scriptCollection.Description,
	}
}

func convertIntegrationRuntimeArtifact(integrationRuntimeArtifact client.IntegrationRuntimeArtifact) IntegrationRuntimeArtifact {
	return IntegrationRuntimeArtifact{
		ID:         integrationRuntimeArtifact.ID,
		Version:    integrationRuntimeArtifact.Version,
		Name:       integrationRuntimeArtifact.Name,
		Type:       integrationRuntimeArtifact.Type,
		Status:     integrationRuntimeArtifact.Status,
		DeployedBy: integrationRuntimeArtifact.DeployedBy,
		DeployedAt: integrationRuntimeArtifact.DeployedOn.String(),
	}
}

package server

type IntegrationPackage struct {
	ID         string `json:"id"`
	Version    string `json:"version"`
	Name       string `json:"name"`
	ShortText  string `json:"shortText"`
	Vendor     string `json:"vendor"`
	Mode       string `json:"mode"`
	CreatedBy  string `json:"createdBy"`
	CreatedAt  string `json:"createdAt"`
	ModifiedBy string `json:"modifiedBy"`
	ModifiedAt string `json:"modifiedAt"`
}

type IntegrationFlow struct {
	ID          string `json:"id"`
	Version     string `json:"version"`
	PackageID   string `json:"packageId"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedBy   string `json:"createdBy"`
	CreatedAt   string `json:"createdAt"`
	ModifiedBy  string `json:"modifiedBy"`
	ModifiedAt  string `json:"modifiedAt"`
}

type ValueMapping struct {
	ID          string `json:"id"`
	Version     string `json:"version"`
	PackageID   string `json:"packageId"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type MessageMapping struct {
	ID          string `json:"id"`
	Version     string `json:"version"`
	PackageID   string `json:"packageId"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ScriptCollection struct {
	ID          string `json:"id"`
	Version     string `json:"version"`
	PackageID   string `json:"packageId"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type IntegrationRuntimeArtifact struct {
	ID         string `json:"id"`
	Version    string `json:"version"`
	Name       string `json:"name"`
	Type       string `json:"type"`
	DeployedBy string `json:"deployedBy"`
	DeployedAt string `json:"deployedAt"`
	Status     string `json:"status"`
}

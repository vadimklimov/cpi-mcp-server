package client

import "github.com/vadimklimov/cpi-mcp-server/internal/util"

type IntegrationPackage struct {
	ID                string         `json:"Id"`
	Version           string         `json:"Version"`
	Name              string         `json:"Name"`
	ShortText         string         `json:"ShortText"`
	Description       string         `json:"Description"`
	Vendor            string         `json:"Vendor"`
	PartnerContent    bool           `json:"PartnerContent"`
	Mode              string         `json:"Mode"`
	UpdateAvailable   bool           `json:"UpdateAvailable"`
	SupportedPlatform string         `json:"SupportedPlatform"`
	Products          string         `json:"Products"`
	Keywords          string         `json:"Keywords"`
	Countries         string         `json:"Countries"`
	Industries        string         `json:"Industries"`
	LineOfBusiness    string         `json:"LineOfBusiness"`
	ResourceID        string         `json:"ResourceId"`
	CreatedBy         string         `json:"CreatedBy"`
	CreationDate      util.Timestamp `json:"CreationDate"`
	ModifiedBy        string         `json:"ModifiedBy"`
	ModifiedDate      util.Timestamp `json:"ModifiedDate"`
}

type IntegrationFlow struct {
	ID          string         `json:"Id"`
	Version     string         `json:"Version"`
	PackageID   string         `json:"PackageId"`
	Name        string         `json:"Name"`
	Description string         `json:"Description"`
	Sender      string         `json:"Sender"`
	Receiver    string         `json:"Receiver"`
	CreatedBy   string         `json:"CreatedBy"`
	CreatedAt   util.Timestamp `json:"CreatedAt"`
	ModifiedBy  string         `json:"ModifiedBy"`
	ModifiedAt  util.Timestamp `json:"ModifiedAt"`
}

type ValueMapping struct {
	ID          string `json:"Id"`
	Version     string `json:"Version"`
	PackageID   string `json:"PackageId"`
	Name        string `json:"Name"`
	Description string `json:"Description"`
}

type MessageMapping struct {
	ID          string `json:"Id"`
	Version     string `json:"Version"`
	PackageID   string `json:"PackageId"`
	Name        string `json:"Name"`
	Description string `json:"Description"`
}

type ScriptCollection struct {
	ID          string `json:"Id"`
	Version     string `json:"Version"`
	PackageID   string `json:"PackageId"`
	Name        string `json:"Name"`
	Description string `json:"Description"`
}

type IntegrationRuntimeArtifact struct {
	ID         string         `json:"Id"`
	Version    string         `json:"Version"`
	Name       string         `json:"Name"`
	Type       string         `json:"Type"`
	DeployedBy string         `json:"DeployedBy"`
	DeployedOn util.Timestamp `json:"DeployedOn"`
	Status     string         `json:"Status"`
}

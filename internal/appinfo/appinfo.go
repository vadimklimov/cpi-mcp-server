package appinfo

import (
	"cmp"
	"sync"
)

type AppInfo struct {
	ID       string
	Name     string
	FullName string
	Version  string
}

var (
	instance *AppInfo
	once     sync.Once
)

// Set using ldflags during build.
var version string

func GetInstance() *AppInfo {
	once.Do(func() {
		instance = &AppInfo{
			ID:       "cpi-mcp-server",
			Name:     "CPI MCP server",
			FullName: "MCP server for SAP Cloud Integration",
			Version:  cmp.Or(version, "unknown"),
		}
	})

	return instance
}

func ID() string {
	return GetInstance().ID
}

func Name() string {
	return GetInstance().Name
}

func FullName() string {
	return GetInstance().FullName
}

func Version() string {
	return GetInstance().Version
}

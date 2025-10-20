variable "VERSION" {
  type        = string
  description = "MCP server version"
  default     = "dev"
}

variable "COMMIT_SHA" {
  type        = string
  description = "Commit hash for build"
  default     = "unknown"
}

group "default" {
  targets = ["dev"]
}

target "dev" {
  context    = "."
  dockerfile = "Dockerfile"
  tags       = ["vadimklimov/cpi-mcp-server:dev"]
  args = {
    VERSION = VERSION
  }
}

target "snapshot" {
  inherits = ["dev"]
  tags     = ["vadimklimov/cpi-mcp-server:${VERSION}-snapshot"]
  platforms = [
    "linux/amd64",
    "linux/arm64"
  ]
  labels = {
    "org.opencontainers.image.title"         = "CPI MCP server"
    "org.opencontainers.image.description"   = "MCP server for SAP Cloud Integration"
    "org.opencontainers.image.vendor"        = "Vadim Klimov"
    "org.opencontainers.image.authors"       = "Vadim Klimov"
    "org.opencontainers.image.licenses"      = "MIT"
    "org.opencontainers.image.url"           = "https://github.com/vadimklimov/cpi-mcp-server"
    "org.opencontainers.image.documentation" = "https://github.com/vadimklimov/cpi-mcp-server"
    "org.opencontainers.image.source"        = "https://github.com/vadimklimov/cpi-mcp-server"
    "org.opencontainers.image.version"       = "${VERSION}-snapshot"
    "org.opencontainers.image.revision"      = "${COMMIT_SHA}"
    "org.opencontainers.image.created"       = "${timestamp()}"
    "org.opencontainers.image.ref.name"      = "docker.io/vadimklimov/cpi-mcp-server:${VERSION}-snapshot"
    "org.opencontainers.image.base.name"     = "docker.io/library/alpine:3.22"
  }
}

target "release" {
  inherits = ["snapshot"]
  tags = [
    "vadimklimov/cpi-mcp-server:${VERSION}",
    "vadimklimov/cpi-mcp-server:latest",
  ]
  output = ["type=registry"]
  labels = {
    "org.opencontainers.image.version"  = "${VERSION}"
    "org.opencontainers.image.ref.name" = "docker.io/vadimklimov/cpi-mcp-server:${VERSION}"
  }
}
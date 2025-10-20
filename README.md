# MCP server for SAP Cloud Integration

The MCP server for SAP Cloud Integration provides AI-powered applications with access to integration packages and integration artifacts within an SAP Cloud Integration (a part of SAP Integration Suite) tenant.

## Tools

The MCP server exposes the following tools:

| Tool                                 | Description                                                               |
| ------------------------------------ | ------------------------------------------------------------------------- |
| get_integration_packages             | Get all integration packages                                              |
| search_integration_packages          | Search integration packages by ID, version, name, vendor or mode          |
| get_integration_flows                | Get all integration flows                                                 |
| search_integration_flows             | Search integration flows by ID, version or name                           |
| get_value_mappings                   | Get all value mappings                                                    |
| search_value_mappings                | Search value mappings by ID, version or name                              |
| get_message_mappings                 | Get all message mappings                                                  |
| search_message_mappings              | Search message mappings by ID, version or name                            |
| get_script_collections               | Get all script collections                                                |
| search_script_collections            | Search script collections by ID, version or name                          |
| get_integration_runtime_artifacts    | Get all integration runtime artifacts                                     |
| search_integration_runtime_artifacts | Search integration runtime artifacts by ID, version, name, type or status |

## Installation

The MCP server can be installed as a standalone binary or run inside a Docker container.

### Standalone binary

#### Option 1: Download binary

Download a binary compatible with your operating system / architecture from the [Releases](https://github.com/vadimklimov/cpi-mcp-server/releases) page.

#### Option 2: Install from source

Use the `go install` command to compile and install the MCP server:

```bash
go install github.com/vadimklimov/cpi-mcp-server@latest
```

#### Option 3: Build from source

1. Clone the repository and navigate into the created directory:

```bash
git clone https://github.com/vadimklimov/cpi-mcp-server.git
cd cpi-mcp-server
```

2. Optionally, download dependencies (otherwise, dependencies will be downloaded automatically during the next step, build):

```bash
go mod download
go mod tidy
```

3. Use the `go build` command to compile the MCP server.

Example of basic build:

```bash
go build
```

Example of build with version information:

```bash
go build \
  -ldflags "-s -w -X github.com/vadimklimov/cpi-mcp-server/internal/appinfo.version=1.2.3" \
  -o cpi-mcp-server \
  .
```

Alternatively, use [GoReleaser](https://goreleaser.com/) to build the MCP server using the provided GoReleaser configuration file.

### Docker container

#### Option 1: Pull Docker image

1. Use the `docker image pull` command (or `docker pull` shorthand) to pull the MCP server image:

```bash
docker pull vadimklimov/cpi-mcp-server
```

#### Option 2: Build Docker image

1. Clone the repository and navigate into the created directory:

```bash
git clone https://github.com/vadimklimov/cpi-mcp-server.git
cd cpi-mcp-server
```

2. Use the `docker build` command to build the Docker image using the provided Dockerfile.

The following build arguments are supported:

| Argument | Description        |
| -------- | ------------------ |
| VERSION  | MCP server version |

Example of basic build:

```bash
docker build .
```

Example of build with version information:

```bash
docker build \
  --tag vadimklimov/cpi-mcp-server:1.2.3 \
  --build-arg VERSION=1.2.3 \
  .
```

Alternatively, use the `docker bake` command (a part of Docker Buildx) to build the Docker image using the provided Bake definition file.

The Bake definition file includes the following build targets:

| Target   | Description       | Default | Multi-platform build | Image tag          |
| -------- | ----------------- | ------- | -------------------- | ------------------ |
| dev      | Development build | yes     | no                   | dev                |
| snapshot | Snapshot build    | no      | yes                  | {VERSION}-snapshot |
| release  | Release build     | no      | yes                  | {VERSION}, latest  |

The following Bake variables are supported:

| Variable   | Description                            | Default value | Used by target         |
| ---------- | -------------------------------------- | ------------- | ---------------------- |
| VERSION    | MCP server version                     | dev           | dev, snapshot, release |
| COMMIT_SHA | Build revision (commit hash for build) | unknown       | snapshot, release      |

> [!NOTE]
> For `snapshot` and `release` builds, using Docker Bake is the recommended approach. The Bake definition file provides sensible defaults and adds additional image metadata.

> [!WARNING]
> Before running `snapshot` or `release` builds, ensure that multi-platform build support is enabled in the Docker environment.

> [!WARNING]
> The `release` target automatically pushes the built image to the registry.

Example of basic build (uses the `dev` target by default):

```bash
docker bake
```

Example of build with version information:

```bash
export VERSION=1.2.3
docker bake
```

or

```bash
docker bake \
  --set dev.tags=vadimklimov/cpi-mcp-server:1.2.3 \
  --set dev.args.VERSION=1.2.3
```

> [!TIP]
> Given that the project is managed with Git, it is recommended to use the current commit hash (an SHA-1 hash) as the build revision.
>
> To retrieve the long (full) form of the current commit hash, run the `git rev-parse HEAD` command.
>
> To retrieve the short form of the current commit hash, run the `git rev-parse --short HEAD` command.

> [!TIP]
> Since the project already uses Git tags for MCP server release versioning, it is recommended to use Git tags for build versioning as well.
>
> To retrieve the most recent tag, run the `git describe --tags $(git rev-list --tags --max-count=1)` command.

## Configuration

For the MCP server to function correctly, configuration must be provided in environment variables.

### Environment variables

The following environment variables are supported:

| Environment variable    | Description                    | Required | Valid values                   | Default value               |
| ----------------------- | ------------------------------ | -------- | ------------------------------ | --------------------------- |
| MCP_CPI_BASE_URL        | Base URL                       | yes      |                                |                             |
| MCP_CPI_TOKEN_URL       | Token URL                      | yes      |                                |                             |
| MCP_CPI_CLIENT_ID       | Client ID                      | yes      |                                |                             |
| MCP_CPI_CLIENT_SECRET   | Client secret                  | yes      |                                |                             |
| MCP_CPI_TRANSPORT       | MCP transport                  | no       | stdio, http                    | stdio                       |
| MCP_CPI_PORT            | Port to use for HTTP transport | no       | value in range 1024-65535      | 8080                        |
| MCP_CPI_LOG_LEVEL       | Log level                      | no       | none, error, warn, info, debug | none                        |
| MCP_CPI_LOG_FILE        | Path to log file               | no       | /path/to/log/file.log          | /var/log/cpi-mcp-server.log |
| MCP_CPI_MAX_CONCURRENCY | Maximum concurrency            | no       |                                | Number of logical CPUs      |
| MCP_CPI_TIMEOUT         | Timeout (in seconds)           | no       |                                | 60                          |

### Access to SAP Cloud Integration APIs

The MCP server uses the public APIs of SAP Cloud Integration to retrieve the required information. These APIs are OAuth-protected and support the client credentials flow. This authentication mechanism is employed by the MCP server to authenticate and authorize API calls made to the SAP Cloud Integration tenant.

To enable the MCP server to access the necessary APIs of the SAP Cloud Integration tenant, an OAuth client for it must be created in the SAP Business Technology Platform subaccount where the corresponding subscription for SAP Integration Suite has been created.

In a Cloud Foundry environment, a service instance represents an OAuth client - hence, a service instance and a service key for it must be created.

**Step 1: Create a service instance**

Create a service instance with the following configuration:

- Service: `SAP Process Integration Runtime`
- Plan: `api`
- Grant type: `Client Credentials`

The following roles are required by the corresponding MCP server tools:

| Tool                                 | Role                  |
| ------------------------------------ | --------------------- |
| get_integration_packages             | WorkspacePackagesRead |
| search_integration_packages          | WorkspacePackagesRead |
| get_integration_flows                | WorkspacePackagesRead |
| search_integration_flows             | WorkspacePackagesRead |
| get_value_mappings                   | WorkspacePackagesRead |
| search_value_mappings                | WorkspacePackagesRead |
| get_message_mappings                 | WorkspacePackagesRead |
| search_message_mappings              | WorkspacePackagesRead |
| get_script_collections               | WorkspacePackagesRead |
| search_script_collections            | WorkspacePackagesRead |
| get_integration_runtime_artifacts    | MonitoringDataRead    |
| search_integration_runtime_artifacts | MonitoringDataRead    |

**Step 2: Create a service key**

Create a service key with the following configuration for the service instance created in the previous step:

- Key type: `ClientId/Secret`

OAuth client credentials and endpoints can be found in the `oauth` section of the service key:

| Configuration parameter | Service key attribute |
| ----------------------- | --------------------- |
| Base URL                | url                   |
| Token URL               | tokenurl              |
| Client ID               | clientid              |
| Client secret           | clientsecret          |

### Examples

Minimum configuration (required environment variables only):

```env
MCP_CPI_BASE_URL=https://{subdomain}.it-cpi{xxxxx}.cfapps.{region}.hana.ondemand.com/api/v1
MCP_CPI_TOKEN_URL=https://{subdomain}.authentication.{region}.hana.ondemand.com/oauth/token
MCP_CPI_CLIENT_ID=xxxxxxxxxx
MCP_CPI_CLIENT_SECRET=xxxxxxxxxx
```

Full configuration (all supported environment variables):

```env
MCP_CPI_BASE_URL=https://{subdomain}.it-cpi{xxxxx}.cfapps.{region}.hana.ondemand.com/api/v1
MCP_CPI_TOKEN_URL=https://{subdomain}.authentication.{region}.hana.ondemand.com/oauth/token
MCP_CPI_CLIENT_ID=xxxxxxxxxx
MCP_CPI_CLIENT_SECRET=xxxxxxxxxx
MCP_CPI_MAX_CONCURRENCY=3
MCP_CPI_TIMEOUT=120
MCP_CPI_TRANSPORT=http
MCP_CPI_PORT=8080
MCP_CPI_LOG_LEVEL=debug
MCP_CPI_LOG_FILE=/var/log/cpi-mcp-server.log
```

## Agentic tools configuration

Refer to the documentation of the IDE, code editor, AI assistant or other agentic tool you use that supports MCP for configuration details.

### Local MCP server

The examples below illustrate the configuration required to add a local MCP server to the **Claude Desktop**. Refer to the documentation of the agentic tool relevant to your setup to adapt these examples.

#### Option 1: Standalone binary

Assuming the `cpi-mcp-server` binary is located in the `/Users/demo/go/bin` directory, add the following configuration to the `claude_desktop_config.json` file:

```json
{
  "mcpServers": {
    "cpi": {
      "command": "/Users/demo/go/bin/cpi-mcp-server",
      "env": {
        "MCP_CPI_BASE_URL": "https://{subdomain}.it-cpi{xxxxx}.cfapps.{region}.hana.ondemand.com/api/v1",
        "MCP_CPI_TOKEN_URL": "https://{subdomain}.authentication.{region}.hana.ondemand.com/oauth/token",
        "MCP_CPI_CLIENT_ID": "xxxxxxxxxx",
        "MCP_CPI_CLIENT_SECRET": "xxxxxxxxxx"
      }
    }
  }
}
```

#### Option 2: Docker container

Assuming the MCP server configuration (see [Configuration](#configuration)) is defined in the `cpi-mcp-server.env` located in the `/Users/demo/.config/mcp` directory, add the following configuration to the `claude_desktop_config.json` file:

```json
{
  "mcpServers": {
    "cpi": {
      "command": "docker",
      "args": [
        "run",
        "-i",
        "--rm",
        "--env-file",
        "/Users/demo/.config/mcp/cpi-mcp-server.env",
        "vadimklimov/cpi-mcp-server"
      ]
    }
  }
}
```

### Remote MCP server

The examples below illustrate the configuration required to add a remote MCP server to the **Claude Code**. Refer to the documentation of the agentic tool relevant to your setup to adapt these examples.

**Step 1: Start the MCP server**

#### Option 1: Standalone binary

Assuming the `cpi-mcp-server` binary is located in the `/Users/demo/go/bin` directory, use the following command to start the MCP server:

```bash
export MCP_CPI_BASE_URL=https://{subdomain}.it-cpi{xxxxx}.cfapps.{region}.hana.ondemand.com/api/v1
export MCP_CPI_TOKEN_URL=https://{subdomain}.authentication.{region}.hana.ondemand.com/oauth/token
export MCP_CPI_CLIENT_ID=xxxxxxxxxx
export MCP_CPI_CLIENT_SECRET=xxxxxxxxxx
export MCP_CPI_TRANSPORT=http
/Users/demo/go/bin/cpi-mcp-server
```

#### Option 2: Docker container

Assuming the MCP server configuration (see [Configuration](#configuration)) is defined in the `cpi-mcp-server.env` located in the `/Users/demo/.config/mcp` directory, use the `docker container run` command (or `docker run` shorthand) to start the MCP server:

```bash
docker run \
  --detach \
  --name cpi-mcp-server \
  --publish 8080:8080 \
  --env-file /Users/demo/.config/mcp/cpi-mcp-server.env \
  vadimklimov/cpi-mcp-server
```

> [!TIP]
> When using the HTTP transport in the MCP server, the container listens on port 8080.
> The `-p` (`--publish`) option of the `docker run` command can be used to publish the container's port 8080 to a different host port (for example, `-p 8888:8080`).

**Step 2: Add a remote MCP server to Claude Code**

#### Option 1: Use Claude Code CLI

Use the `claude mcp add` command to add the remote MCP server:

```bash
claude mcp add --transport http cpi http://localhost:8080/mcp
```

#### Option 2: Manually edit configuration file

Add the following configuration to the `.mcp.json` file:

```json
{
  "mcpServers": {
    "cpi": {
      "type": "http",
      "url": "http://localhost:8080/mcp"
    }
  }
}
```

> [!NOTE]
> When running the MCP server on a remote host, ensure that the host name is set appropriately when adding it to the agentic tool.

> [!WARNING]
> MCP servers operating in remote mode must be configured to use HTTPS (encrypted transport protocol) to ensure secure communication between the client and the server. Use of HTTP is strongly discouraged, as it does not provide transport-level security.

> [!WARNING]
> The current implementation of the MCP server does not provide authorization capabilities. When operating the server in remote mode, appropriate access restriction mechanisms must be implemented to prevent unauthorized access to the server.

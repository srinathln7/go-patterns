# Configurable Server in Go

## `pattern.go`

This file demonstrates a configurable server implementation in Go using functional options. The code includes:

- `Config` type: A function type for configuring server options.
- `Opts` type: Represents the configuration options for the server.
- `Server` type: Contains the server configuration.
- `defaultOpts` function: Returns default configuration options.
- `withTLS` function: A configuration function to enable TLS.
- `maxConnections` function: A configuration function to set the maximum number of connections.
- `newServer` function: Creates a new server instance with specified configurations.
- The `main` function showcases creating servers with different configurations.

## `simple.go`

This file provides a simpler server implementation without using functional options. It includes:

- `Server` type: Represents the server with basic configurations.
- `newServer` function: Creates a new server instance with specified parameters.
- The `main` function demonstrates creating a server with specific configurations.

The `pattern.go` file introduces a more flexible and extensible approach to configuring servers, allowing dynamic adjustments without modifying the original code. In contrast, the `simple.go` file shows a less flexible approach where modifications require updating the source code.

Feel free to explore and run the code to understand the benefits of using functional options for configuring servers.

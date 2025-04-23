# IP Registry Lookup

This Go application performs reverse DNS lookups for a list of IP addresses. It maintains a local registry of IP-to-hostname mappings and dynamically updates the registry with new hostnames resolved during runtime.

## Features

- Maintains a local registry of IP addresses and their associated hostnames.
- Performs reverse DNS lookups using Go's `net.LookupAddr` function.
- Dynamically updates the registry with resolved hostnames.
- Outputs the IP-to-hostname mappings to the console.
- Includes unit tests for the `resolve` function with mock DNS lookups.


## Prerequisites

- Go 1.18 or later installed on your system.

## Usage

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd ipregistry
   ```

2. Run the application:
   ```bash
   go run main.go
   ```

3. The output will display the resolved hostnames for the given IP addresses.

## Running Tests

Unit tests are included for the `resolve` function. The tests use a mock implementation of the `search` function to simulate DNS lookups.

To run the tests, use the following command:
```bash
go test
```

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.
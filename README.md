# IP Registry Lookup

This Go application performs reverse DNS lookups for a list of IP addresses. It maintains a local registry of IP-to-hostname mappings and dynamically updates the registry with new hostnames resolved during runtime.

## Features

- Maintains a local registry of IP addresses and their associated hostnames.
- Performs reverse DNS lookups using Go's `net.LookupAddr` function.
- Dynamically updates the registry with resolved hostnames.
- Outputs the IP-to-hostname mappings to the console.

## How It Works

1. The application initializes a local registry (`ipregistry`) with predefined IP-to-hostname mappings.
2. It iterates over a list of IP addresses and attempts to resolve their hostnames:
   - If the IP exists in the local registry, it uses the cached hostnames.
   - If the IP is not in the registry, it performs a reverse DNS lookup and updates the registry with the resolved hostnames.
3. The results are printed to the console in the format: `IP = hostname`.

## Prerequisites

- Go 1.18 or later installed on your system.

## Usage

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd ipregistry
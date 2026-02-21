<p align="center">
  <a href="https://thepilotclub.org" target="_blank" rel="noreferrer">
    <img
      src="https://static1.squarespace.com/static/614689d3918044012d2ac1b4/t/616ff36761fabc72642806e3/1634726781251/TPC_FullColor_TransparentBg_1280x1024_72dpi.png"
      width="280"
      alt="The Pilot Club Logo"
    />
  </a>
</p>

# tpcgo

Go client for The Pilot Club (TPC) APIs (Core API + Flight Crew Portal) plus the VATSIM endpoints used by TPC tooling.

[![Go Reference](https://pkg.go.dev/badge/github.com/the-pilot-club/tpcgo.svg)](https://pkg.go.dev/github.com/the-pilot-club/tpcgo)

## Features

- **Single session** that groups credentials/config for:
  - **FCP** (Flight Crew Portal) API
  - **Core API** (suggestions + quiz)
  - **VATSIM** (datafeed + member lookups)
- Typed request/response structs
- Common HTTP error mapping via exported sentinel errors (e.g. `ErrUnAuthorized`, `ErrNotFound`)

## Requirements

- Go 1.24+ (see `go.mod`)

## Install

```bash
go get github.com/the-pilot-club/tpcgo
```

## Quick start

Create a session by providing the keys you plan to use. Use empty strings for services you don’t need.

```go
package main

import (
	"fmt"
	"log"
	"github.com/the-pilot-club/tpcgo"
)

func main() { 
	s, err := tpcgo.NewSession(tpcgo.SessionConfig{
		FCPKey: "<FCP_API_KEY>", 
		FCPEnv: "production", // "production" or "beta" 
		VATSIMAPIKey: "<VATSIM_API_KEY>", 
		CoreApiKey: "<CORE_API_KEY>", 
	}) 
	
	if err != nil { 
		log.Fatal(err) 
	}
	
	df, err := s.GetVatsimDataFeed()
	
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Connected clients: %d\n", df.General.ConnectedClients)
}
```

## Configuration

### FCP environment

FCP requests require an environment string:

- `production`
- `beta`

If an unsupported value is used, endpoint construction may produce an empty base URL and requests will fail.

### Authentication

Keys are sent as HTTP headers. **Never** hardcode secrets in source code.

Suggested approaches:

- environment variables
- your preferred secrets manager / CI secret store
- local `.env` loaded by your application (not committed)


## Notes / limitations

- Requests are performed with `net/http` using a default client configuration.
  If you need timeouts, retries, custom transports, or tracing, you may want to wrap/fork the request layer to accept an `*http.Client`.
- APIs evolve. If you encounter a mismatched field or endpoint change, please open an issue or PR.

## Versioning

This project follows Go module versioning. Please pin to a release/tag in production usage.

## Contributing

Issues and PRs are welcome. Keep changes focused and include tests where possible.

## License

See [LICENSE](LICENSE).

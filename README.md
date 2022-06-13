# Netflix Conductor Client SDK

`conductor-go` repository provides the client SDKs to build Task Workers in Go

## Quick Start

1. [Setup conductor-go package](#Setup-conductor-go-package)
2. [Create and run Task Workers](workers_sdk.md)
3. [Create workflows using Code](workflow_sdk.md)

### Setup conductor go package

Create a folder to build your package:
```shell
mkdir conductor-go/
cd conductor-go/
```

Create a go.mod file for dependencies
```go
module conductor_test

go 1.18

require (
	github.com/conductor-sdk/conductor-go
)   
```

Install dependencies.  This will download all the required dependencies
```shell
go get
```

## Configuration

### Authentication settings (optional)
Use if your conductor server requires authentication
* keyId: Key
* keySecret: Secret for the Key

```go
authenticationSettings := settings.NewAuthenticationSettings(
  "keyId",
  "keySecret",
)
```


### Configure API Client
```go

apiClient := conductor_http_client.NewAPIClient(
    settings.NewAuthenticationSettings(
        KEY,
        SECRET,
    ),
    settings.NewHttpSettings(
        "https://play.orkes.io",
    ),
)
	
```

### Next: [Create and run Task Workers](workers_sdk.md)
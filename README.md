<p align="center">
  <a href="https://developer.rizefs.com/" target="_blank" align="center">
    <img src="https://cdn.rizefs.com/web-content/logos/rize-github.png" width="200">
  </a>
  <br />
</p>

# Rize Go SDK

Rize makes financial services simple and accessible by enabling fintechs, financial institutions and brands to build across multiple account types with one API. The Rize Go SDK enables access to all platform services in our sandbox, integration and production environments.

When you're ready to build, [open a sandbox environment](https://rizefs.com/get-access/). If you have questions, feedback, or a use case you want to discuss with us, contact us at [hello@rizemoney.com](mailto:hello@rizemoney.com).

For more information, check out our [Platform API Documentation](https://developer.rizefs.com/).

### Supported Go Versions

The Rize Go SDK is compatible with the two most recent, major Go releases. 

We currently support **Go v1.18 and higher**.

## Getting Started

### Installation

To install the latest Rize Go SDK, add the module as a dependency using go mod:
```
go get github.com/rizefinance/rize-go-sdk@latest
```

To install a specific release version:
```
go get github.com/rizefinance/rize-go-sdk@v1.0.0
```

### Configuration

The SDK requires program configuration credentials in order to interact with the API.

You can find these in the [**Rize Admin Portal**](https://admin-sandbox.rizefs.com/).

| Parameter   | Description                                                  | Default   |
| ----------- | ------------------------------------------------------------ | --------- |
| HMACKey     | HMAC key for the target environment | "" |
| ProgramUID  | Program UID for the target environment | "" |
| Environment | The Rize environment to be used:<br> `"sandbox"`, `"integration"` or `"production"` | "sandbox" |
| Debug  | Enable debug logging | false |

### Import the SDK

Import the SDK module into your code:

```go
import rize "github.com/rizefinance/rize-go-sdk/platform"
```

### Start Making API Calls

```go
import rize "github.com/rizefinance/rize-go-sdk/platform"

func main() {
	config := rize.Config{
		HMACKey:     hmac,
		ProgramUID:  programUID,
		Environment: environment,
		Debug:       false,
	}
	rc, err := rize.NewRizeClient(&config)
	if err != nil {
		log.Fatal("Error building RizeClient\n", err)
	}
	
	// Create a new Compliance Worflow
	wcp := rize.WorkflowCreateParams{
		CustomerUID:              "h9MzupcjtA3LPW2e",
		ProductCompliancePlanUID: "25NQX3GGXpAtpUmP",
	}
	c, err := rc.ComplianceWorkflows.Create(context.Background(), &wcp)
	if err != nil {
		log.Fatal("Error creating new Compliance Workflow\n", err)
	}
}
```

### Configure `http.Client`

You have the option to supply your own `http.Client`. By default, the SDK uses `DefaultClient` with a 30s timeout.

To set a proxy for all requests, configure the Transport for the HTTPClient:

```go
config := rize.Config{
	HMACKey:     hmac,
	ProgramUID:  programUID,
	Environment: environment,
	HTTPClient: &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyUrl),
		},
	},
	Debug:       false,
}
```

Similarly, to configure the timeout, set it on the HTTPClient:

```go
config := rize.Config{
	HMACKey:     hmac,
	ProgramUID:  programUID,
	Environment: environment,
	HTTPClient: &http.Client{
		Timeout: time.Minute,
	},
	Debug:       false,
}
```

## Examples

The [examples](examples/) directory provides basic implementation examples for each API endpoint. They require configuration credentials to be set as environment variables. 

Use the provided `dotenv` file to set those credentials for local testing. 

```sh
# Generate a local configuration file
$ cp .env-example .env
```

## Unit Tests

```go
$ go test ./platform -v
```

## Documentation

* [Platform API Documentation](https://developer.rizefs.com/)
* [Rize Postman](https://www.postman.com/rizemoney/)
* [Rize GitHub](https://github.com/RizeFinance)
* [Rize JS SDK](https://github.com/RizeFinance/rize-js)
* [Rize Website](https://www.rizemoney.com/)

## License
MIT License. Copyright 2021-Present Rize Money, Inc. All rights reserved.

See [LICENSE](LICENSE)

<p align="center">
  <a href="https://rizefs.com" target="_blank" align="center">
    <img src="https://rizefs.com/wp-content/uploads/2021/01/rizelogo-grey.svg" width="200">
  </a>
  <br />
</p>

*Make financial services simple and accessible. Rize enables fintechs, financial institutions and brands to build across multiple account types with one API. If you want to join us [**Check out our open positions**](https://rizefs.com/careers/).*

# Official Rize SDKs for GO

### Usage

```go
import rize "github.com/rizefinance/rize-go-sdk/v1"

func main() {
	config := rize.Config{
		HMACKey:     hmac,
		ProgramUID:  programUID,
		Environment: environment,
	}
	
	// Start making API calls
}
```

| Parameter   | Description                                                  | Default   |
| ----------- | ------------------------------------------------------------ | --------- |
| HMACKey     | HMAC key for the target environment | nil |
| ProgramUID  | Program UID for the target environment | nil |
| Environment | The Rize environment to be used: `'sandbox'`, `'integration'` or `'production'` | 'sandbox' |

### API Docs

Go to [https://developer.rizefs.com/](https://developer.rizefs.com/)

### Testing Locally

Run the SDK from the command line.

```sh
# Generate a local configuration file
$ cp .env-example .env
```

```sh
# Test an SDK command
$ go run cmd/rize/main.go [COMMAND]
```

```sh
# Open the help menu
$ go run cmd/rize/main.go -h
```

### License
MIT License. Copyright 2021-Present Rize Money, Inc. All rights reserved.

See [LICENSE](LICENSE)

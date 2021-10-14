<p align="center">
  <a href="https://rizefs.com" target="_blank" align="center">
    <img src="https://rizefs.com/wp-content/uploads/2021/01/rizelogo-grey.svg" width="200">
  </a>
  <br />
</p>

*Make financial services simple and accessible. Rize enables fintechs, financial institutions and brands to build across multiple account types with one API. If you want to join us [**Check out our open positions**](https://rizefs.com/careers/).*

# Official Rize SDKs for GO

## Changelogs

Checkout the changelogs per release [here]()

## Usage

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

## API Docs

Go to [https://developer.rizefs.com/](https://developer.rizefs.com/)

## Testing Locally

Run the SDK from the command line.

```sh
$ cd cmd/rize
$ go run main.go -k <HMAC_KEY> -p <PROGRAM_UID> -e <ENVIRONMENT>
```

```sh
# Open the help menu
$ go run main.go -h
```

## License
MIT License. Copyright 2021 Rize Money, Inc. All rights reserved.

See [LICENSE](LICENSE)

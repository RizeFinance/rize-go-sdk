<p align="center">
  <a href="https://rizefs.com" target="_blank" align="center">
    <img src="https://rizefs.com/wp-content/uploads/2021/01/rizelogo-grey.svg" width="200">
  </a>
  <br />
</p>

*Making financial services simple and accessible. Rize enables fintechs, financial institutions and brands to build across multiple account types with one API.*

# Official Rize SDKs for Go

### Go Version

+ Go >= 1.18

### Getting Started

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
	
	// Start making API calls
}
```

| Parameter   | Description                                                  | Default   |
| ----------- | ------------------------------------------------------------ | --------- |
| HMACKey     | HMAC key for the target environment | nil |
| ProgramUID  | Program UID for the target environment | nil |
| Environment | The Rize environment to be used: `'sandbox'`, `'integration'` or `'production'` | 'sandbox' |
| Debug  | Enable debug logging | false |

### Examples

The API examples require availability of certain Rize environment variables. You can set them in the provided dotenv file. 

```sh
# Generate a local configuration file
$ cp .env-example .env
```

### API Docs

Full API documentation is available at [https://developer.rizefs.com/](https://developer.rizefs.com/).

### License
MIT License. Copyright 2021-Present Rize Money, Inc. All rights reserved.

See [LICENSE](LICENSE)

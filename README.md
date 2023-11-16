# Ding Go SDK

The Ding Golang library provides convenient access to the Ding API from applications written in the Golang language.

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/ding-live/ding-golang
```
<!-- End SDK Installation -->

## SDK Example Usage

<!-- Start SDK Example Usage -->
### Send a code

Send an OTP code to a user's phone number.


```go
package main

import (
	"context"
	dinggolang "github.com/ding-live/ding-golang"
	"github.com/ding-live/ding-golang/models/components"
	"log"
)

func main() {
	s := dinggolang.New(
		dinggolang.WithSecurity("YOUR_API_KEY"),
	)

	ctx := context.Background()
	res, err := s.Otp.CreateAutentication(ctx, &components.CreateAuthenticationRequest{
		CustomerUUID: "eae192ab-9e1e-4b21-b5b1-bfcb79a32fcc",
		PhoneNumber:  "+1234567890",
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.CreateAuthenticationResponse != nil {
		// handle response
	}
}

```

### Check a code

Check that a code entered by a user is valid.


```go
package main

import (
	"context"
	dinggolang "github.com/ding-live/ding-golang"
	"github.com/ding-live/ding-golang/models/components"
	"log"
)

func main() {
	s := dinggolang.New(
		dinggolang.WithSecurity("YOUR_API_KEY"),
	)

	ctx := context.Background()
	res, err := s.Otp.Check(ctx, &components.CreateCheckRequest{
		AuthenticationUUID: "e0e7b0e9-739d-424b-922f-1c2cb48ab077",
		CheckCode:          "123456",
		CustomerUUID:       "8f1196d5-806e-4b71-9b24-5f96ec052808",
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.CreateCheckResponse != nil {
		// handle response
	}
}

```

### Perform a retry

Perform a retry if a user has not received the code.


```go
package main

import (
	"context"
	dinggolang "github.com/ding-live/ding-golang"
	"github.com/ding-live/ding-golang/models/components"
	"log"
)

func main() {
	s := dinggolang.New(
		dinggolang.WithSecurity("YOUR_API_KEY"),
	)

	ctx := context.Background()
	res, err := s.Otp.Retry(ctx, &components.RetryAuthenticationRequest{
		AuthenticationUUID: "a74ee547-564d-487a-91df-37fb25413a91",
		CustomerUUID:       "3c8b3a46-881e-4cdd-93a6-f7f238bf020a",
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.RetryAuthenticationResponse != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## Available Resources and Operations


### [Otp](docs/sdks/otp/README.md)

* [Check](docs/sdks/otp/README.md#check) - Check a code
* [CreateAutentication](docs/sdks/otp/README.md#createautentication) - Send a code
* [Retry](docs/sdks/otp/README.md#retry) - Perform a retry

### [Lookup](docs/sdks/lookup/README.md)

* [Lookup](docs/sdks/lookup/README.md#lookup) - Perform a phone number lookup
<!-- End SDK Available Operations -->

<!-- Start Dev Containers -->

<!-- End Dev Containers -->

<!-- Start Error Handling -->
## Error Handling

Handling errors in this SDK should largely match your expectations.  All operations return a response object or an error, they will never return both.  When specified by the OpenAPI spec document, the SDK will return the appropriate subclass.

| Error Object            | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.ErrorResponse | 400                     | application/json        |
| sdkerrors.SDKError      | 400-600                 | */*                     |

### Example

```go
package main

import (
	"context"
	"errors"
	dinggolang "github.com/ding-live/ding-golang"
	"github.com/ding-live/ding-golang/models/components"
	"github.com/ding-live/ding-golang/models/sdkerrors"
	"log"
)

func main() {
	s := dinggolang.New(
		dinggolang.WithSecurity("YOUR_API_KEY"),
	)

	ctx := context.Background()
	res, err := s.Otp.Check(ctx, &components.CreateCheckRequest{
		AuthenticationUUID: "e0e7b0e9-739d-424b-922f-1c2cb48ab077",
		CheckCode:          "123456",
		CustomerUUID:       "8f1196d5-806e-4b71-9b24-5f96ec052808",
	})
	if err != nil {

		var e *sdkerrors.ErrorResponse
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.SDKError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling -->

<!-- Start Custom HTTP Client -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client -->

<!-- Start Authentication -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security scheme globally:

| Name     | Type     | Scheme   |
| -------- | -------- | -------- |
| `APIKey` | apiKey   | API key  |

You can configure it using the `WithSecurity` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	dinggolang "github.com/ding-live/ding-golang"
	"github.com/ding-live/ding-golang/models/components"
	"log"
)

func main() {
	s := dinggolang.New(
		dinggolang.WithSecurity("YOUR_API_KEY"),
	)

	ctx := context.Background()
	res, err := s.Otp.Check(ctx, &components.CreateCheckRequest{
		AuthenticationUUID: "e0e7b0e9-739d-424b-922f-1c2cb48ab077",
		CheckCode:          "123456",
		CustomerUUID:       "8f1196d5-806e-4b71-9b24-5f96ec052808",
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.CreateCheckResponse != nil {
		// handle response
	}
}

```
<!-- End Authentication -->

<!-- Start Go Types -->

<!-- End Go Types -->



<!-- Start Server Selection -->
## Server Selection

### Select Server by Name

You can override the default server globally using the `WithServer` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the names associated with the available servers:

| Name | Server | Variables |
| ----- | ------ | --------- |
| `production` | `https://api.ding.live/v1` | None |
#### Example

```go
package main

import (
	"context"
	dinggolang "github.com/ding-live/ding-golang"
	"github.com/ding-live/ding-golang/models/components"
	"log"
)

func main() {
	s := dinggolang.New(
		dinggolang.WithServer("production"),
		dinggolang.WithSecurity("YOUR_API_KEY"),
	)

	ctx := context.Background()
	res, err := s.Otp.Check(ctx, &components.CreateCheckRequest{
		AuthenticationUUID: "e0e7b0e9-739d-424b-922f-1c2cb48ab077",
		CheckCode:          "123456",
		CustomerUUID:       "8f1196d5-806e-4b71-9b24-5f96ec052808",
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.CreateCheckResponse != nil {
		// handle response
	}
}

```


### Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	dinggolang "github.com/ding-live/ding-golang"
	"github.com/ding-live/ding-golang/models/components"
	"log"
)

func main() {
	s := dinggolang.New(
		dinggolang.WithServerURL("https://api.ding.live/v1"),
		dinggolang.WithSecurity("YOUR_API_KEY"),
	)

	ctx := context.Background()
	res, err := s.Otp.Check(ctx, &components.CreateCheckRequest{
		AuthenticationUUID: "e0e7b0e9-739d-424b-922f-1c2cb48ab077",
		CheckCode:          "123456",
		CustomerUUID:       "8f1196d5-806e-4b71-9b24-5f96ec052808",
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.CreateCheckResponse != nil {
		// handle response
	}
}

```
<!-- End Server Selection -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

# Development

## Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

## Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release!

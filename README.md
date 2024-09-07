# Ding Go SDK

The Ding Golang library provides convenient access to the Ding API from applications written in the Golang language.

<!-- Start Summary [summary] -->
## Summary

Ding: The OTP API allows you to send authentication codes to your users using their phone numbers.
<!-- End Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents

* [SDK Installation](#sdk-installation)
* [SDK Example Usage](#sdk-example-usage)
* [Available Resources and Operations](#available-resources-and-operations)
* [Retries](#retries)
* [Error Handling](#error-handling)
* [Server Selection](#server-selection)
* [Custom HTTP Client](#custom-http-client)
* [Authentication](#authentication)
* [Special Types](#special-types)
<!-- End Table of Contents [toc] -->

<!-- Start SDK Installation [installation] -->
## SDK Installation

To add the SDK as a dependency to your project:
```bash
go get github.com/ding-live/ding-golang
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

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
	res, err := s.Otp.CreateAuthentication(ctx, &components.CreateAuthenticationRequest{
		CustomerUUID: "c9f826e0-deca-41ec-871f-ecd6e8efeb46",
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
	"log"
)

func main() {
	s := dinggolang.New(
		dinggolang.WithSecurity("YOUR_API_KEY"),
	)

	ctx := context.Background()
	res, err := s.Otp.Retry(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	if res.RetryAuthenticationResponse != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

### [Otp](docs/sdks/otp/README.md)

* [Check](docs/sdks/otp/README.md#check) - Check a code
* [CreateAuthentication](docs/sdks/otp/README.md#createauthentication) - Send a code
* [Feedback](docs/sdks/otp/README.md#feedback) - Send feedback
* [Retry](docs/sdks/otp/README.md#retry) - Perform a retry

### [Lookup](docs/sdks/lookup/README.md)

* [Lookup](docs/sdks/lookup/README.md#lookup) - Perform a phone number lookup
<!-- End Available Resources and Operations [operations] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations.  All operations return a response object or an error, they will never return both.  When specified by the OpenAPI spec document, the SDK will return the appropriate subclass.

| Error Object            | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.ErrorResponse | 400                     | application/json        |
| sdkerrors.SDKError      | 4xx-5xx                 | */*                     |

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
<!-- End Error Handling [errors] -->

<!-- Start Custom HTTP Client [http-client] -->
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
<!-- End Custom HTTP Client [http-client] -->

<!-- Start Authentication [security] -->
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
<!-- End Authentication [security] -->

<!-- Start Special Types [types] -->
## Special Types


<!-- End Special Types [types] -->

<!-- Start Server Selection [server] -->
## Server Selection

### Select Server by Index

You can override the default server globally using the `WithServerIndex` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the indexes associated with the available servers:

| # | Server | Variables |
| - | ------ | --------- |
| 0 | `https://api.ding.live/v1` | None |

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
		dinggolang.WithServerIndex(0),
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
<!-- End Server Selection [server] -->

<!-- Start Retries [retries] -->
## Retries

Some of the endpoints in this SDK support retries. If you use the SDK without any configuration, it will fall back to the default retry strategy provided by the API. However, the default retry strategy can be overridden on a per-operation basis, or across the entire SDK.

To change the default retry strategy for a single API call, simply provide a `retry.Config` object to the call by using the `WithRetries` option:
```go
package main

import (
	"context"
	dinggolang "github.com/ding-live/ding-golang"
	"github.com/ding-live/ding-golang/models/components"
	"github.com/ding-live/ding-golang/retry"
	"log"
	"models/operations"
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
	}, operations.WithRetries(
		retry.Config{
			Strategy: "backoff",
			Backoff: &retry.BackoffStrategy{
				InitialInterval: 1,
				MaxInterval:     50,
				Exponent:        1.1,
				MaxElapsedTime:  100,
			},
			RetryConnectionErrors: false,
		}))
	if err != nil {
		log.Fatal(err)
	}
	if res.CreateCheckResponse != nil {
		// handle response
	}
}

```

If you'd like to override the default retry strategy for all operations that support retries, you can use the `WithRetryConfig` option at SDK initialization:
```go
package main

import (
	"context"
	dinggolang "github.com/ding-live/ding-golang"
	"github.com/ding-live/ding-golang/models/components"
	"github.com/ding-live/ding-golang/retry"
	"log"
)

func main() {
	s := dinggolang.New(
		dinggolang.WithRetryConfig(
			retry.Config{
				Strategy: "backoff",
				Backoff: &retry.BackoffStrategy{
					InitialInterval: 1,
					MaxInterval:     50,
					Exponent:        1.1,
					MaxElapsedTime:  100,
				},
				RetryConnectionErrors: false,
			}),
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
<!-- End Retries [retries] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

# Development

## Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

## Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release!

# Ding Go SDK

The Ding Golang library provides convenient access to the Ding API from applications written in the Golang language.

<!-- Start Summary [summary] -->
## Summary

Ding: The OTP API allows you to send authentication codes to your users using their phone numbers.
<!-- End Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents
<!-- $toc-max-depth=2 -->
* [Ding Go SDK](#ding-go-sdk)
  * [SDK Installation](#sdk-installation)
  * [SDK Example Usage](#sdk-example-usage)
  * [Available Resources and Operations](#available-resources-and-operations)
  * [Error Handling](#error-handling)
  * [Custom HTTP Client](#custom-http-client)
  * [Authentication](#authentication)
  * [Server Selection](#server-selection)
  * [Retries](#retries)
* [Development](#development)
  * [Maturity](#maturity)
  * [Contributions](#contributions)

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
	ctx := context.Background()

	s := dinggolang.New(
		dinggolang.WithSecurity("YOUR_API_KEY"),
	)

	res, err := s.Otp.CreateAuthentication(ctx, &components.CreateAuthenticationRequest{
		CustomerUUID: "cf2edc1c-7fc6-48fb-86da-b7508c6b7b71",
		Locale:       dinggolang.String("fr-FR"),
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
	ctx := context.Background()

	s := dinggolang.New(
		dinggolang.WithSecurity("YOUR_API_KEY"),
	)

	res, err := s.Otp.Check(ctx, &components.CreateCheckRequest{
		AuthenticationUUID: "eebe792b-2fcc-44a0-87f1-650e79259e02",
		CheckCode:          "123456",
		CustomerUUID:       "64f66a7c-4b2c-4131-a8ff-d5b954cca05f",
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
	ctx := context.Background()

	s := dinggolang.New(
		dinggolang.WithSecurity("YOUR_API_KEY"),
	)

	res, err := s.Otp.Retry(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	if res.RetryAuthenticationResponse != nil {
		// handle response
	}
}

```

### Send feedback

Send feedback about the authentication process.


```go
package main

import (
	"context"
	dinggolang "github.com/ding-live/ding-golang"
	"github.com/ding-live/ding-golang/models/components"
	"log"
)

func main() {
	ctx := context.Background()

	s := dinggolang.New(
		dinggolang.WithSecurity("YOUR_API_KEY"),
	)

	res, err := s.Otp.Feedback(ctx, &components.FeedbackRequest{
		CustomerUUID: "cc0f6c04-40de-448f-8301-3cb0e6565dff",
		PhoneNumber:  "+1234567890",
		Status:       components.FeedbackRequestStatusOnboarded,
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.FeedbackResponse != nil {
		// handle response
	}
}

```

### Get authentication status

Get the status of an authentication.

```go
package main

import (
	"context"
	dinggolang "github.com/ding-live/ding-golang"
	"log"
)

func main() {
	ctx := context.Background()

	s := dinggolang.New(
		dinggolang.WithSecurity("YOUR_API_KEY"),
	)

	res, err := s.Otp.GetAuthenticationStatus(ctx, "d8446450-f2fa-4dd9-806b-df5b8c661f23")
	if err != nil {
		log.Fatal(err)
	}
	if res.AuthenticationStatusResponse != nil {
		// handle response
	}
}

```

### Look up for phone number

Perform a phone number lookup.

```go
package main

import (
	"context"
	dinggolang "github.com/ding-live/ding-golang"
	"log"
)

func main() {
	ctx := context.Background()

	s := dinggolang.New(
		dinggolang.WithSecurity("YOUR_API_KEY"),
	)

	res, err := s.Lookup.Lookup(ctx, "69a197d9-356c-45d1-a807-41874e16b555", "<value>", nil)
	if err != nil {
		log.Fatal(err)
	}
	if res.LookupResponse != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

<details open>
<summary>Available methods</summary>


### [Lookup](docs/sdks/lookup/README.md)

* [Lookup](docs/sdks/lookup/README.md#lookup) - Look up for phone number

### [Otp](docs/sdks/otp/README.md)

* [Check](docs/sdks/otp/README.md#check) - Check a code
* [CreateAuthentication](docs/sdks/otp/README.md#createauthentication) - Send a code
* [Feedback](docs/sdks/otp/README.md#feedback) - Send feedback
* [GetAuthenticationStatus](docs/sdks/otp/README.md#getauthenticationstatus) - Get authentication status
* [Retry](docs/sdks/otp/README.md#retry) - Perform a retry

</details>
<!-- End Available Resources and Operations [operations] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations. All operations return a response object or an error, they will never return both.

By Default, an API error will return `sdkerrors.SDKError`. When custom error responses are specified for an operation, the SDK may also return their associated error. You can refer to respective *Errors* tables in SDK docs for more details on possible error types for each operation.

For example, the `Check` function may return the following errors:

| Error Type              | Status Code | Content Type     |
| ----------------------- | ----------- | ---------------- |
| sdkerrors.ErrorResponse | 400         | application/json |
| sdkerrors.SDKError      | 4XX, 5XX    | \*/\*            |

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
	ctx := context.Background()

	s := dinggolang.New(
		dinggolang.WithSecurity("YOUR_API_KEY"),
	)

	res, err := s.Otp.Check(ctx, &components.CreateCheckRequest{
		AuthenticationUUID: "eebe792b-2fcc-44a0-87f1-650e79259e02",
		CheckCode:          "123456",
		CustomerUUID:       "64f66a7c-4b2c-4131-a8ff-d5b954cca05f",
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

| Name     | Type   | Scheme  |
| -------- | ------ | ------- |
| `APIKey` | apiKey | API key |

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
	ctx := context.Background()

	s := dinggolang.New(
		dinggolang.WithSecurity("YOUR_API_KEY"),
	)

	res, err := s.Otp.Check(ctx, &components.CreateCheckRequest{
		AuthenticationUUID: "eebe792b-2fcc-44a0-87f1-650e79259e02",
		CheckCode:          "123456",
		CustomerUUID:       "64f66a7c-4b2c-4131-a8ff-d5b954cca05f",
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

<!-- Start Server Selection [server] -->
## Server Selection

### Override Server URL Per-Client

The default server can be overridden globally using the `WithServerURL(serverURL string)` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	dinggolang "github.com/ding-live/ding-golang"
	"github.com/ding-live/ding-golang/models/components"
	"log"
)

func main() {
	ctx := context.Background()

	s := dinggolang.New(
		dinggolang.WithServerURL("https://api.ding.live/v1"),
		dinggolang.WithSecurity("YOUR_API_KEY"),
	)

	res, err := s.Otp.Check(ctx, &components.CreateCheckRequest{
		AuthenticationUUID: "eebe792b-2fcc-44a0-87f1-650e79259e02",
		CheckCode:          "123456",
		CustomerUUID:       "64f66a7c-4b2c-4131-a8ff-d5b954cca05f",
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
	ctx := context.Background()

	s := dinggolang.New(
		dinggolang.WithSecurity("YOUR_API_KEY"),
	)

	res, err := s.Otp.Check(ctx, &components.CreateCheckRequest{
		AuthenticationUUID: "eebe792b-2fcc-44a0-87f1-650e79259e02",
		CheckCode:          "123456",
		CustomerUUID:       "64f66a7c-4b2c-4131-a8ff-d5b954cca05f",
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
	ctx := context.Background()

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

	res, err := s.Otp.Check(ctx, &components.CreateCheckRequest{
		AuthenticationUUID: "eebe792b-2fcc-44a0-87f1-650e79259e02",
		CheckCode:          "123456",
		CustomerUUID:       "64f66a7c-4b2c-4131-a8ff-d5b954cca05f",
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

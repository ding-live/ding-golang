<!-- Start SDK Example Usage [usage] -->
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
		CustomerUUID: "c9f826e0-deca-41ec-871f-ecd6e8efeb46",
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
		CustomerUUID: "c0c405fa-6bcb-4094-9430-7d6e2428ff23",
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

	res, err := s.Lookup.Lookup(ctx, "6e93aa15-9177-4d09-8395-b69ce50db1c8", "<value>", nil)
	if err != nil {
		log.Fatal(err)
	}
	if res.LookupResponse != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->
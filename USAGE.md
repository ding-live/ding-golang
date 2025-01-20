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
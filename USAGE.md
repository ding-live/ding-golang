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
	s := dinggolang.New(
		dinggolang.WithSecurity("YOUR_API_KEY"),
	)

	var request *components.CreateAuthenticationRequest = &components.CreateAuthenticationRequest{
		CustomerUUID: "c9f826e0-deca-41ec-871f-ecd6e8efeb46",
		PhoneNumber:  "+1234567890",
	}

	ctx := context.Background()
	res, err := s.Otp.CreateAuthentication(ctx, request)
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

	var request *components.CreateCheckRequest = &components.CreateCheckRequest{
		AuthenticationUUID: "e0e7b0e9-739d-424b-922f-1c2cb48ab077",
		CheckCode:          "123456",
		CustomerUUID:       "8f1196d5-806e-4b71-9b24-5f96ec052808",
	}

	ctx := context.Background()
	res, err := s.Otp.Check(ctx, request)
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

	var request *components.RetryAuthenticationRequest = &components.RetryAuthenticationRequest{
		AuthenticationUUID: "a74ee547-564d-487a-91df-37fb25413a91",
		CustomerUUID:       "3c8b3a46-881e-4cdd-93a6-f7f238bf020a",
	}

	ctx := context.Background()
	res, err := s.Otp.Retry(ctx, request)
	if err != nil {
		log.Fatal(err)
	}
	if res.RetryAuthenticationResponse != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->
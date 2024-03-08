# Otp
(*Otp*)

## Overview

Send OTP codes to your users using their phone numbers.

### Available Operations

* [Check](#check) - Check a code
* [CreateAutentication](#createautentication) - Send a code
* [Retry](#retry) - Perform a retry

## Check

Check a code

### Example Usage

```go
package main

import(
	"github.com/ding-live/ding-golang/models/components"
	dinggolang "github.com/ding-live/ding-golang"
	"context"
	"log"
)

func main() {
    s := dinggolang.New(
        dinggolang.WithSecurity("YOUR_API_KEY"),
    )

    ctx := context.Background()
    res, err := s.Otp.Check(ctx, &components.CreateCheckRequest{
        AuthenticationUUID: "e0e7b0e9-739d-424b-922f-1c2cb48ab077",
        CheckCode: "123456",
        CustomerUUID: "8f1196d5-806e-4b71-9b24-5f96ec052808",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateCheckResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [components.CreateCheckRequest](../../models/components/createcheckrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.CheckResponse](../../models/operations/checkresponse.md), error**
| Error Object            | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.ErrorResponse | 400                     | application/json        |
| sdkerrors.SDKError      | 4xx-5xx                 | */*                     |

## CreateAutentication

Send a code

### Example Usage

```go
package main

import(
	"github.com/ding-live/ding-golang/models/components"
	dinggolang "github.com/ding-live/ding-golang"
	"context"
	"log"
)

func main() {
    s := dinggolang.New(
        dinggolang.WithSecurity("YOUR_API_KEY"),
    )

    ctx := context.Background()
    res, err := s.Otp.CreateAutentication(ctx, &components.CreateAuthenticationRequest{
        CustomerUUID: "eae192ab-9e1e-4b21-b5b1-bfcb79a32fcc",
        PhoneNumber: "+1234567890",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateAuthenticationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [components.CreateAuthenticationRequest](../../models/components/createauthenticationrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.CreateAutenticationResponse](../../models/operations/createautenticationresponse.md), error**
| Error Object            | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.ErrorResponse | 400                     | application/json        |
| sdkerrors.SDKError      | 4xx-5xx                 | */*                     |

## Retry

Perform a retry

### Example Usage

```go
package main

import(
	"github.com/ding-live/ding-golang/models/components"
	dinggolang "github.com/ding-live/ding-golang"
	"context"
	"log"
)

func main() {
    s := dinggolang.New(
        dinggolang.WithSecurity("YOUR_API_KEY"),
    )

    ctx := context.Background()
    res, err := s.Otp.Retry(ctx, &components.RetryAuthenticationRequest{
        AuthenticationUUID: "a74ee547-564d-487a-91df-37fb25413a91",
        CustomerUUID: "3c8b3a46-881e-4cdd-93a6-f7f238bf020a",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.RetryAuthenticationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [components.RetryAuthenticationRequest](../../models/components/retryauthenticationrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.RetryResponse](../../models/operations/retryresponse.md), error**
| Error Object            | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.ErrorResponse | 400                     | application/json        |
| sdkerrors.SDKError      | 4xx-5xx                 | */*                     |

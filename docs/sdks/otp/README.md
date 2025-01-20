# Otp
(*Otp*)

## Overview

Send OTP codes to your users using their phone numbers.

### Available Operations

* [Check](#check) - Check a code
* [CreateAuthentication](#createauthentication) - Send a code
* [Feedback](#feedback) - Send feedback
* [GetAuthenticationStatus](#getauthenticationstatus) - Get authentication status
* [Retry](#retry) - Perform a retry

## Check

Check a code

### Example Usage

```go
package main

import(
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
        CheckCode: "123456",
        CustomerUUID: "64f66a7c-4b2c-4131-a8ff-d5b954cca05f",
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
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.CheckResponse](../../models/operations/checkresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.ErrorResponse | 400                     | application/json        |
| sdkerrors.SDKError      | 4XX, 5XX                | \*/\*                   |

## CreateAuthentication

Send a code

### Example Usage

```go
package main

import(
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
        Locale: dinggolang.String("fr-FR"),
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
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.CreateAuthenticationResponse](../../models/operations/createauthenticationresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.ErrorResponse | 400                     | application/json        |
| sdkerrors.SDKError      | 4XX, 5XX                | \*/\*                   |

## Feedback

Send feedback

### Example Usage

```go
package main

import(
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
        PhoneNumber: "+1234567890",
        Status: components.FeedbackRequestStatusOnboarded,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.FeedbackResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [components.FeedbackRequest](../../models/components/feedbackrequest.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |
| `opts`                                                                   | [][operations.Option](../../models/operations/option.md)                 | :heavy_minus_sign:                                                       | The options for this request.                                            |

### Response

**[*operations.FeedbackResponse](../../models/operations/feedbackresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.ErrorResponse | 400                     | application/json        |
| sdkerrors.SDKError      | 4XX, 5XX                | \*/\*                   |

## GetAuthenticationStatus

Get authentication status

### Example Usage

```go
package main

import(
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

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `authUUID`                                               | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetAuthenticationStatusResponse](../../models/operations/getauthenticationstatusresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.ErrorResponse | 400                     | application/json        |
| sdkerrors.SDKError      | 4XX, 5XX                | \*/\*                   |

## Retry

Perform a retry

### Example Usage

```go
package main

import(
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

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [components.RetryAuthenticationRequest](../../models/components/retryauthenticationrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.RetryResponse](../../models/operations/retryresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.ErrorResponse | 400                     | application/json        |
| sdkerrors.SDKError      | 4XX, 5XX                | \*/\*                   |
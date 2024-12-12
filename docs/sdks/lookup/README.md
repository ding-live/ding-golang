# Lookup
(*Lookup*)

## Overview

Retrieve up-to-date metadata about a specific phone number

### Available Operations

* [Lookup](#lookup) - Look up for phone number

## Lookup

Look up for phone number

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

    res, err := s.Lookup.Lookup(ctx, "6e93aa15-9177-4d09-8395-b69ce50db1c8", "<value>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.LookupResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `customerUUID`                                           | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `phoneNumber`                                            | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `type_`                                                  | [][operations.Type](../../models/operations/type.md)     | :heavy_minus_sign:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.LookupResponse](../../models/operations/lookupresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.ErrorResponse | 400                     | application/json        |
| sdkerrors.SDKError      | 4XX, 5XX                | \*/\*                   |
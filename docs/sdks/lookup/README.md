# Lookup
(*Lookup*)

## Overview

Retrieve up-to-date metadata about a specific phone number

### Available Operations

* [Lookup](#lookup) - Perform a phone number lookup

## Lookup

Perform a phone number lookup

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


    var customerUUID string = "6e93aa15-9177-4d09-8395-b69ce50db1c8"

    var phoneNumber string = "<value>"

    ctx := context.Background()
    res, err := s.Lookup.Lookup(ctx, customerUUID, phoneNumber)
    if err != nil {
        log.Fatal(err)
    }

    if res.LookupResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `customerUUID`                                        | *string*                                              | :heavy_check_mark:                                    | N/A                                                   |
| `phoneNumber`                                         | *string*                                              | :heavy_check_mark:                                    | N/A                                                   |


### Response

**[*operations.LookupResponse](../../models/operations/lookupresponse.md), error**
| Error Object            | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.ErrorResponse | 400                     | application/json        |
| sdkerrors.SDKError      | 4xx-5xx                 | */*                     |

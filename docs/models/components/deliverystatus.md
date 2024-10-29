# DeliveryStatus


## Fields

| Field                                                                                                                                                                                                                                                                                                                                                                                  | Type                                                                                                                                                                                                                                                                                                                                                                                   | Required                                                                                                                                                                                                                                                                                                                                                                               | Description                                                                                                                                                                                                                                                                                                                                                                            |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `AttemptID`                                                                                                                                                                                                                                                                                                                                                                            | **string*                                                                                                                                                                                                                                                                                                                                                                              | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                     | The ID of the attempt.                                                                                                                                                                                                                                                                                                                                                                 |
| `AttemptNumber`                                                                                                                                                                                                                                                                                                                                                                        | **int64*                                                                                                                                                                                                                                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                     | The attempt number.                                                                                                                                                                                                                                                                                                                                                                    |
| `CreatedAt`                                                                                                                                                                                                                                                                                                                                                                            | [*time.Time](https://pkg.go.dev/time#Time)                                                                                                                                                                                                                                                                                                                                             | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                     | N/A                                                                                                                                                                                                                                                                                                                                                                                    |
| `Status`                                                                                                                                                                                                                                                                                                                                                                               | [*components.AuthenticationStatusResponseStatus](../../models/components/authenticationstatusresponsestatus.md)                                                                                                                                                                                                                                                                        | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                     | The status of the delivery. Possible values are:<br/>  * `unknown` - The status of the delivery is unknown.<br/>  * `submitted` - The message has been submitted to the carrier.<br/>  * `in_transit` - The message is in transit to the recipient.<br/>  * `delivered` - The message has been delivered to the recipient.<br/>  * `undeliverable` - The message could not be delivered to the recipient.<br/> |
| `Type`                                                                                                                                                                                                                                                                                                                                                                                 | [*components.AuthenticationStatusResponseSchemasType](../../models/components/authenticationstatusresponseschemastype.md)                                                                                                                                                                                                                                                              | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                     | The type of the event.                                                                                                                                                                                                                                                                                                                                                                 |
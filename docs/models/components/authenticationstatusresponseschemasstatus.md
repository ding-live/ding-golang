# AuthenticationStatusResponseSchemasStatus

The status of the attempt. Possible values are:
  * `pending` - The attempt is pending.
  * `delivered` - The attempt was delivered.
  * `failed` - The attempt failed.
  * `rate_limited` - The authentication was rate limited and cannot be retried.
  * `expired` - The authentication has expired and cannot be retried.



## Values

| Name                                                   | Value                                                  |
| ------------------------------------------------------ | ------------------------------------------------------ |
| `AuthenticationStatusResponseSchemasStatusPending`     | pending                                                |
| `AuthenticationStatusResponseSchemasStatusDelivered`   | delivered                                              |
| `AuthenticationStatusResponseSchemasStatusFailed`      | failed                                                 |
| `AuthenticationStatusResponseSchemasStatusRateLimited` | rate_limited                                           |
| `AuthenticationStatusResponseSchemasStatusExpired`     | expired                                                |
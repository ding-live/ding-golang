# RetryAuthenticationResponseStatus

The status of the retry. Possible values are:
  * `approved` - The retry was approved and a new code was sent.
  * `denied` - The retry was denied.
  * `no_attempt` - No attempt was sent yet, so a retry cannot be completed.
  * `rate_limited` - The authentication was rate limited and cannot be retried.
  * `expired_auth` - The authentication has expired and cannot be retried.
  * `already_validated` - The authentication has already been validated.



## Values

| Name                                                | Value                                               |
| --------------------------------------------------- | --------------------------------------------------- |
| `RetryAuthenticationResponseStatusApproved`         | approved                                            |
| `RetryAuthenticationResponseStatusDenied`           | denied                                              |
| `RetryAuthenticationResponseStatusNoAttempt`        | no_attempt                                          |
| `RetryAuthenticationResponseStatusRateLimited`      | rate_limited                                        |
| `RetryAuthenticationResponseStatusExpiredAuth`      | expired_auth                                        |
| `RetryAuthenticationResponseStatusAlreadyValidated` | already_validated                                   |
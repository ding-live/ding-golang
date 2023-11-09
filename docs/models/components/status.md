# Status

The status of the authentication. Possible values are:
  * `pending` - The OTP code is being sent.
  * `rate_limited` - This user is rate-limited and cannot receive another code.
  * `spam_detected` - This attempt is flagged as spam. Go to the dashboard for more details.



## Values

| Name                 | Value                |
| -------------------- | -------------------- |
| `StatusPending`      | pending              |
| `StatusRateLimited`  | rate_limited         |
| `StatusSpamDetected` | spam_detected        |
# Code

A machine-readable code that describes the error. Possible values are:
  * `invalid_phone_number` - This is not a valid E.164 number.
  * `internal_server_error` - An internal server error occurred.
  * `bad_request` - The request was malformed.
  * `account_invalid` - The provided customer UUID is invalid.
  * `negative_balance` - You have a negative balance.
  * `invalid_line` - Ding does not support this type of phone number.
  * `unsupported_region` - Ding does not support this region yet.
  * `invalid_auth_uuid` - The provided authentication UUID is invalid.
  * `blocked_number` - The phone number is in the blocklist.
  * `invalid_app_version` - The provided application version is invalid.
  * `invalid_os_version` - The provided OS version is invalid.
  * `invalid_device_model` - The provided device model is invalid.
  * `invalid_device_id` - The provided device ID is invalid.
  * `no_associated_auth_found` - The associated authentication was not found.
  * `duplicated_feedback_status` - Duplicated feedback status has found.
  * `invalid_feedback_status` - The provided feedback status is invalid.
  * `invalid_template_id` - The provided template ID is invalid.



## Values

| Name                                | Value                               |
| ----------------------------------- | ----------------------------------- |
| `CodeInvalidPhoneNumber`            | invalid_phone_number                |
| `CodeInternalServerError`           | internal_server_error               |
| `CodeBadRequest`                    | bad_request                         |
| `CodeAccountInvalid`                | account_invalid                     |
| `CodeNegativeBalance`               | negative_balance                    |
| `CodeInvalidLine`                   | invalid_line                        |
| `CodeUnsupportedRegion`             | unsupported_region                  |
| `CodeInvalidAuthUUID`               | invalid_auth_uuid                   |
| `CodeInvalidAppRealm`               | invalid_app_realm                   |
| `CodeUnsupportedAppRealmDeviceType` | unsupported_app_realm_device_type   |
| `CodeAppRealmRequireDeviceType`     | app_realm_require_device_type       |
| `CodeBlockedNumber`                 | blocked_number                      |
| `CodeInvalidAppVersion`             | invalid_app_version                 |
| `CodeInvalidOsVersion`              | invalid_os_version                  |
| `CodeInvalidDeviceModel`            | invalid_device_model                |
| `CodeInvalidDeviceID`               | invalid_device_id                   |
| `CodeNoAssociatedAuthFound`         | no_associated_auth_found            |
| `CodeDuplicatedFeedbackStatus`      | duplicated_feedback_status          |
| `CodeInvalidFeedbackStatus`         | invalid_feedback_status             |
| `CodeInvalidTemplateID`             | invalid_template_id                 |
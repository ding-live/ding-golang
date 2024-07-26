// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"fmt"
)

// Code - A machine-readable code that describes the error.
type Code string

const (
	CodeInvalidPhoneNumber            Code = "invalid_phone_number"
	CodeInternalServerError           Code = "internal_server_error"
	CodeBadRequest                    Code = "bad_request"
	CodeAccountInvalid                Code = "account_invalid"
	CodeNegativeBalance               Code = "negative_balance"
	CodeInvalidLine                   Code = "invalid_line"
	CodeUnsupportedRegion             Code = "unsupported_region"
	CodeInvalidAuthUUID               Code = "invalid_auth_uuid"
	CodeInvalidAppRealm               Code = "invalid_app_realm"
	CodeUnsupportedAppRealmDeviceType Code = "unsupported_app_realm_device_type"
	CodeAppRealmRequireDeviceType     Code = "app_realm_require_device_type"
	CodeBlockedNumber                 Code = "blocked_number"
	CodeInvalidAppVersion             Code = "invalid_app_version"
	CodeInvalidOsVersion              Code = "invalid_os_version"
	CodeInvalidDeviceModel            Code = "invalid_device_model"
	CodeInvalidDeviceID               Code = "invalid_device_id"
	CodeNoAssociatedAuthFound         Code = "no_associated_auth_found"
	CodeDuplicatedFeedbackStatus      Code = "duplicated_feedback_status"
	CodeInvalidFeedbackStatus         Code = "invalid_feedback_status"
	CodeInvalidTemplateID             Code = "invalid_template_id"
	CodeSuspendedAccount              Code = "suspended_account"
)

func (e Code) ToPointer() *Code {
	return &e
}
func (e *Code) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "invalid_phone_number":
		fallthrough
	case "internal_server_error":
		fallthrough
	case "bad_request":
		fallthrough
	case "account_invalid":
		fallthrough
	case "negative_balance":
		fallthrough
	case "invalid_line":
		fallthrough
	case "unsupported_region":
		fallthrough
	case "invalid_auth_uuid":
		fallthrough
	case "invalid_app_realm":
		fallthrough
	case "unsupported_app_realm_device_type":
		fallthrough
	case "app_realm_require_device_type":
		fallthrough
	case "blocked_number":
		fallthrough
	case "invalid_app_version":
		fallthrough
	case "invalid_os_version":
		fallthrough
	case "invalid_device_model":
		fallthrough
	case "invalid_device_id":
		fallthrough
	case "no_associated_auth_found":
		fallthrough
	case "duplicated_feedback_status":
		fallthrough
	case "invalid_feedback_status":
		fallthrough
	case "invalid_template_id":
		fallthrough
	case "suspended_account":
		*e = Code(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Code: %v", v)
	}
}

// ErrorResponse - Bad Request
type ErrorResponse struct {
	// A machine-readable code that describes the error.
	Code *Code `json:"code,omitempty"`
	// A link to the documentation that describes the error.
	DocURL *string `json:"doc_url,omitempty"`
	// A human-readable message that describes the error.
	Message *string `json:"message,omitempty"`
}

var _ error = &ErrorResponse{}

func (e *ErrorResponse) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

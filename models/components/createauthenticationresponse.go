// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
	"github.com/ding-live/ding-golang/internal/utils"
	"time"
)

// Status - The status of the authentication. Possible values are:
//   - `pending` - The OTP code is being sent.
//   - `rate_limited` - This user is rate-limited and cannot receive another code.
//   - `spam_detected` - This attempt is flagged as spam. Go to the dashboard for more details.
type Status string

const (
	StatusPending      Status = "pending"
	StatusRateLimited  Status = "rate_limited"
	StatusSpamDetected Status = "spam_detected"
)

func (e Status) ToPointer() *Status {
	return &e
}
func (e *Status) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "pending":
		fallthrough
	case "rate_limited":
		fallthrough
	case "spam_detected":
		*e = Status(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Status: %v", v)
	}
}

// CreateAuthenticationResponse - A successful response to an authentication creation request.
type CreateAuthenticationResponse struct {
	// A unique identifier for the authentication that you can use on the /check and /retry endpoints.
	AuthenticationUUID *string    `json:"authentication_uuid,omitempty"`
	CreatedAt          *time.Time `json:"created_at,omitempty"`
	// The time at which the authentication expires and can no longer be checked or retried.
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
	// The status of the authentication. Possible values are:
	//   * `pending` - The OTP code is being sent.
	//   * `rate_limited` - This user is rate-limited and cannot receive another code.
	//   * `spam_detected` - This attempt is flagged as spam. Go to the dashboard for more details.
	//
	Status *Status `json:"status,omitempty"`
}

func (c CreateAuthenticationResponse) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateAuthenticationResponse) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateAuthenticationResponse) GetAuthenticationUUID() *string {
	if o == nil {
		return nil
	}
	return o.AuthenticationUUID
}

func (o *CreateAuthenticationResponse) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *CreateAuthenticationResponse) GetExpiresAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.ExpiresAt
}

func (o *CreateAuthenticationResponse) GetStatus() *Status {
	if o == nil {
		return nil
	}
	return o.Status
}

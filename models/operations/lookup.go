// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/ding-live/ding-golang/models/components"
	"net/http"
)

type LookupRequest struct {
	CustomerUUID string `header:"style=simple,explode=false,name=customer-uuid"`
	PhoneNumber  string `pathParam:"style=simple,explode=false,name=phone_number"`
}

func (o *LookupRequest) GetCustomerUUID() string {
	if o == nil {
		return ""
	}
	return o.CustomerUUID
}

func (o *LookupRequest) GetPhoneNumber() string {
	if o == nil {
		return ""
	}
	return o.PhoneNumber
}

type LookupResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// OK
	LookupResponse *components.LookupResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *LookupResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *LookupResponse) GetLookupResponse() *components.LookupResponse {
	if o == nil {
		return nil
	}
	return o.LookupResponse
}

func (o *LookupResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *LookupResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

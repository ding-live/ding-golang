// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/ding-live/ding-golang/models/components"
	"net/http"
)

type GetAuthenticationStatusRequest struct {
	AuthUUID string `pathParam:"style=simple,explode=false,name=auth_uuid"`
}

func (o *GetAuthenticationStatusRequest) GetAuthUUID() string {
	if o == nil {
		return ""
	}
	return o.AuthUUID
}

type GetAuthenticationStatusResponse struct {
	// OK
	AuthenticationStatusResponse *components.AuthenticationStatusResponse
	// HTTP response content type for this operation
	ContentType string
	// Bad Request
	ErrorResponse *components.ErrorResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetAuthenticationStatusResponse) GetAuthenticationStatusResponse() *components.AuthenticationStatusResponse {
	if o == nil {
		return nil
	}
	return o.AuthenticationStatusResponse
}

func (o *GetAuthenticationStatusResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAuthenticationStatusResponse) GetErrorResponse() *components.ErrorResponse {
	if o == nil {
		return nil
	}
	return o.ErrorResponse
}

func (o *GetAuthenticationStatusResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAuthenticationStatusResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

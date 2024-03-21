// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/ding-live/ding-golang/models/components"
	"net/http"
)

type CreateAuthenticationResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// OK
	CreateAuthenticationResponse *components.CreateAuthenticationResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreateAuthenticationResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateAuthenticationResponse) GetCreateAuthenticationResponse() *components.CreateAuthenticationResponse {
	if o == nil {
		return nil
	}
	return o.CreateAuthenticationResponse
}

func (o *CreateAuthenticationResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateAuthenticationResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
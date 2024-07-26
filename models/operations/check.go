// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package operations

import (
	"github.com/ding-live/ding-golang/models/components"
	"net/http"
)

type CheckResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// OK
	CreateCheckResponse *components.CreateCheckResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CheckResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CheckResponse) GetCreateCheckResponse() *components.CreateCheckResponse {
	if o == nil {
		return nil
	}
	return o.CreateCheckResponse
}

func (o *CheckResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CheckResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

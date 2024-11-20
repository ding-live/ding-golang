// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package dinggolang

import (
	"context"
	"fmt"
	"github.com/ding-live/ding-golang/internal/hooks"
	"github.com/ding-live/ding-golang/internal/utils"
	"github.com/ding-live/ding-golang/models/components"
	"github.com/ding-live/ding-golang/retry"
	"net/http"
	"time"
)

// ServerList contains the list of servers available to the SDK
var ServerList = []string{
	// The production Ding API server
	"https://api.ding.live/v1",
}

// HTTPClient provides an interface for suplying the SDK with a custom HTTP client
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// String provides a helper function to return a pointer to a string
func String(s string) *string { return &s }

// Bool provides a helper function to return a pointer to a bool
func Bool(b bool) *bool { return &b }

// Int provides a helper function to return a pointer to an int
func Int(i int) *int { return &i }

// Int64 provides a helper function to return a pointer to an int64
func Int64(i int64) *int64 { return &i }

// Float32 provides a helper function to return a pointer to a float32
func Float32(f float32) *float32 { return &f }

// Float64 provides a helper function to return a pointer to a float64
func Float64(f float64) *float64 { return &f }

// Pointer provides a helper function to return a pointer to a type
func Pointer[T any](v T) *T { return &v }

type sdkConfiguration struct {
	Client            HTTPClient
	Security          func(context.Context) (interface{}, error)
	ServerURL         string
	ServerIndex       int
	Language          string
	OpenAPIDocVersion string
	SDKVersion        string
	GenVersion        string
	UserAgent         string
	RetryConfig       *retry.Config
	Hooks             *hooks.Hooks
	Timeout           *time.Duration
}

func (c *sdkConfiguration) GetServerDetails() (string, map[string]string) {
	if c.ServerURL != "" {
		return c.ServerURL, nil
	}

	return ServerList[c.ServerIndex], nil
}

// Ding - Ding: The OTP API allows you to send authentication codes to your users using their phone numbers.
type Ding struct {
	// Send OTP codes to your users using their phone numbers.
	Otp *Otp
	// Retrieve up-to-date metadata about a specific phone number
	Lookup *Lookup

	sdkConfiguration sdkConfiguration
}

type SDKOption func(*Ding)

// WithServerURL allows the overriding of the default server URL
func WithServerURL(serverURL string) SDKOption {
	return func(sdk *Ding) {
		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithTemplatedServerURL allows the overriding of the default server URL with a templated URL populated with the provided parameters
func WithTemplatedServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *Ding) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithServerIndex allows the overriding of the default server by index
func WithServerIndex(serverIndex int) SDKOption {
	return func(sdk *Ding) {
		if serverIndex < 0 || serverIndex >= len(ServerList) {
			panic(fmt.Errorf("server index %d out of range", serverIndex))
		}

		sdk.sdkConfiguration.ServerIndex = serverIndex
	}
}

// WithClient allows the overriding of the default HTTP client used by the SDK
func WithClient(client HTTPClient) SDKOption {
	return func(sdk *Ding) {
		sdk.sdkConfiguration.Client = client
	}
}

// WithSecurity configures the SDK to use the provided security details
func WithSecurity(apiKey string) SDKOption {
	return func(sdk *Ding) {
		security := components.Security{APIKey: apiKey}
		sdk.sdkConfiguration.Security = utils.AsSecuritySource(&security)
	}
}

// WithSecuritySource configures the SDK to invoke the Security Source function on each method call to determine authentication
func WithSecuritySource(security func(context.Context) (components.Security, error)) SDKOption {
	return func(sdk *Ding) {
		sdk.sdkConfiguration.Security = func(ctx context.Context) (interface{}, error) {
			return security(ctx)
		}
	}
}

func WithRetryConfig(retryConfig retry.Config) SDKOption {
	return func(sdk *Ding) {
		sdk.sdkConfiguration.RetryConfig = &retryConfig
	}
}

// WithTimeout Optional request timeout applied to each operation
func WithTimeout(timeout time.Duration) SDKOption {
	return func(sdk *Ding) {
		sdk.sdkConfiguration.Timeout = &timeout
	}
}

// New creates a new instance of the SDK with the provided options
func New(opts ...SDKOption) *Ding {
	sdk := &Ding{
		sdkConfiguration: sdkConfiguration{
			Language:          "go",
			OpenAPIDocVersion: "1.0.0",
			SDKVersion:        "0.14.23",
			GenVersion:        "2.460.1",
			UserAgent:         "speakeasy-sdk/go 0.14.23 2.460.1 1.0.0 github.com/ding-live/ding-golang",
			Hooks:             hooks.New(),
		},
	}
	for _, opt := range opts {
		opt(sdk)
	}

	// Use WithClient to override the default client if you would like to customize the timeout
	if sdk.sdkConfiguration.Client == nil {
		sdk.sdkConfiguration.Client = &http.Client{Timeout: 60 * time.Second}
	}

	currentServerURL, _ := sdk.sdkConfiguration.GetServerDetails()
	serverURL := currentServerURL
	serverURL, sdk.sdkConfiguration.Client = sdk.sdkConfiguration.Hooks.SDKInit(currentServerURL, sdk.sdkConfiguration.Client)
	if serverURL != currentServerURL {
		sdk.sdkConfiguration.ServerURL = serverURL
	}

	sdk.Otp = newOtp(sdk.sdkConfiguration)

	sdk.Lookup = newLookup(sdk.sdkConfiguration)

	return sdk
}

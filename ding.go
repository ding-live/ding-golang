// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package dinggolang

import (
	"context"
	"fmt"
	"github.com/ding-live/ding-golang/internal/utils"
	"github.com/ding-live/ding-golang/models/components"
	"net/http"
	"time"
)

const (
	// The production Ding API server
	ServerProduction string = "production"
)

// ServerList contains the list of servers available to the SDK
var ServerList = map[string]string{
	ServerProduction: "https://api.ding.live/v1",
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

type sdkConfiguration struct {
	DefaultClient     HTTPClient
	SecurityClient    HTTPClient
	Security          func(context.Context) (interface{}, error)
	ServerURL         string
	Server            string
	Language          string
	OpenAPIDocVersion string
	SDKVersion        string
	GenVersion        string
	UserAgent         string
	RetryConfig       *utils.RetryConfig
}

func (c *sdkConfiguration) GetServerDetails() (string, map[string]string) {
	if c.ServerURL != "" {
		return c.ServerURL, nil
	}

	if c.Server == "" {
		c.Server = "production"
	}

	return ServerList[c.Server], nil
}

// Ding - Ding: The OTP API allows you to send authentication codes to your users using their phone numbers
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

// WithServer allows the overriding of the default server by name
func WithServer(server string) SDKOption {
	return func(sdk *Ding) {
		_, ok := ServerList[server]
		if !ok {
			panic(fmt.Errorf("server %s not found", server))
		}

		sdk.sdkConfiguration.Server = server
	}
}

// WithClient allows the overriding of the default HTTP client used by the SDK
func WithClient(client HTTPClient) SDKOption {
	return func(sdk *Ding) {
		sdk.sdkConfiguration.DefaultClient = client
	}
}

func withSecurity(security interface{}) func(context.Context) (interface{}, error) {
	return func(context.Context) (interface{}, error) {
		return &security, nil
	}
}

// WithSecurity configures the SDK to use the provided security details

func WithSecurity(apiKey string) SDKOption {
	return func(sdk *Ding) {
		security := components.Security{APIKey: apiKey}
		sdk.sdkConfiguration.Security = withSecurity(&security)
	}
}

func WithRetryConfig(retryConfig utils.RetryConfig) SDKOption {
	return func(sdk *Ding) {
		sdk.sdkConfiguration.RetryConfig = &retryConfig
	}
}

// New creates a new instance of the SDK with the provided options
func New(opts ...SDKOption) *Ding {
	sdk := &Ding{
		sdkConfiguration: sdkConfiguration{
			Language:          "go",
			OpenAPIDocVersion: "1.0.0",
			SDKVersion:        "0.1.2",
			GenVersion:        "2.187.2",
			UserAgent:         "speakeasy-sdk/go 0.1.2 2.187.2 1.0.0 github.com/ding-live/ding-golang",
		},
	}
	for _, opt := range opts {
		opt(sdk)
	}

	// Use WithClient to override the default client if you would like to customize the timeout
	if sdk.sdkConfiguration.DefaultClient == nil {
		sdk.sdkConfiguration.DefaultClient = &http.Client{Timeout: 60 * time.Second}
	}
	if sdk.sdkConfiguration.SecurityClient == nil {
		if sdk.sdkConfiguration.Security != nil {
			sdk.sdkConfiguration.SecurityClient = utils.ConfigureSecurityClient(sdk.sdkConfiguration.DefaultClient, sdk.sdkConfiguration.Security)
		} else {
			sdk.sdkConfiguration.SecurityClient = sdk.sdkConfiguration.DefaultClient
		}
	}

	sdk.Otp = newOtp(sdk.sdkConfiguration)

	sdk.Lookup = newLookup(sdk.sdkConfiguration)

	return sdk
}

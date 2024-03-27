// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package dinggolang

import (
	"bytes"
	"context"
	"fmt"
	"github.com/ding-live/ding-golang/internal/hooks"
	"github.com/ding-live/ding-golang/internal/utils"
	"github.com/ding-live/ding-golang/models/components"
	"github.com/ding-live/ding-golang/models/operations"
	"github.com/ding-live/ding-golang/models/sdkerrors"
	"io"
	"net/http"
)

// Lookup - Retrieve up-to-date metadata about a specific phone number
type Lookup struct {
	sdkConfiguration sdkConfiguration
}

func newLookup(sdkConfig sdkConfiguration) *Lookup {
	return &Lookup{
		sdkConfiguration: sdkConfig,
	}
}

// Lookup - Perform a phone number lookup
func (s *Lookup) Lookup(ctx context.Context, customerUUID string, phoneNumber string) (*operations.LookupResponse, error) {
	hookCtx := hooks.HookContext{
		Context:        ctx,
		OperationID:    "lookup",
		SecuritySource: s.sdkConfiguration.Security,
	}

	request := operations.LookupRequest{
		CustomerUUID: customerUUID,
		PhoneNumber:  phoneNumber,
	}

	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	opURL, err := utils.GenerateURL(ctx, baseURL, "/lookup/{phone_number}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", opURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", s.sdkConfiguration.UserAgent)

	utils.PopulateHeaders(ctx, req, request, nil)

	if err := utils.PopulateSecurity(ctx, req, s.sdkConfiguration.Security); err != nil {
		return nil, err
	}

	req, err = s.sdkConfiguration.Hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
	if err != nil {
		return nil, err
	}

	httpRes, err := s.sdkConfiguration.Client.Do(req)
	if err != nil || httpRes == nil {
		if err != nil {
			err = fmt.Errorf("error sending request: %w", err)
		} else {
			err = fmt.Errorf("error sending request: no response")
		}

		_, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
		return nil, err
	} else if utils.MatchStatusCodes([]string{"400", "4XX", "5XX"}, httpRes.StatusCode) {
		httpRes, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, httpRes, nil)
		if err != nil {
			return nil, err
		}
	} else {
		httpRes, err = s.sdkConfiguration.Hooks.AfterSuccess(hooks.AfterSuccessContext{HookContext: hookCtx}, httpRes)
		if err != nil {
			return nil, err
		}
	}

	res := &operations.LookupResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: httpRes.Header.Get("Content-Type"),
		RawResponse: httpRes,
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/json`):
			var out components.LookupResponse
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.LookupResponse = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 400:
		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/json`):
			var out sdkerrors.ErrorResponse
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			return nil, &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode >= 400 && httpRes.StatusCode < 500:
		fallthrough
	case httpRes.StatusCode >= 500 && httpRes.StatusCode < 600:
		return nil, sdkerrors.NewSDKError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	default:
		return nil, sdkerrors.NewSDKError("unknown status code returned", httpRes.StatusCode, string(rawBody), httpRes)
	}

	return res, nil
}

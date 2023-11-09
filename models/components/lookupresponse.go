// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// LineType - The type of phone line.
type LineType string

const (
	LineTypeFixedLine       LineType = "FixedLine"
	LineTypeMobile          LineType = "Mobile"
	LineTypeTollFree        LineType = "TollFree"
	LineTypePremiumRate     LineType = "PremiumRate"
	LineTypeSharedCost      LineType = "SharedCost"
	LineTypeVoip            LineType = "Voip"
	LineTypePager           LineType = "Pager"
	LineTypeVoiceMail       LineType = "VoiceMail"
	LineTypeUniversalAccess LineType = "UniversalAccess"
	LineTypeService         LineType = "Service"
	LineTypeUnknown         LineType = "Unknown"
)

func (e LineType) ToPointer() *LineType {
	return &e
}

func (e *LineType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "FixedLine":
		fallthrough
	case "Mobile":
		fallthrough
	case "TollFree":
		fallthrough
	case "PremiumRate":
		fallthrough
	case "SharedCost":
		fallthrough
	case "Voip":
		fallthrough
	case "Pager":
		fallthrough
	case "VoiceMail":
		fallthrough
	case "UniversalAccess":
		fallthrough
	case "Service":
		fallthrough
	case "Unknown":
		*e = LineType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LineType: %v", v)
	}
}

type LookupResponse struct {
	// The carrier of the phone number.
	Carrier *string `json:"carrier,omitempty"`
	// The ISO 3166-1 alpha-2 country code of the phone number.
	CountryCode *string `json:"country_code,omitempty"`
	// The type of phone line.
	LineType *LineType `json:"line_type,omitempty"`
	// The mobile country code of the phone number.
	Mcc *string `json:"mcc,omitempty"`
	// The mobile network code of the phone number.
	Mnc *string `json:"mnc,omitempty"`
	// Whether the phone number has been ported.
	NumberPorted *bool `json:"number_ported,omitempty"`
	// An E.164 formatted phone number.
	PhoneNumber *string `json:"phone_number,omitempty"`
}

func (o *LookupResponse) GetCarrier() *string {
	if o == nil {
		return nil
	}
	return o.Carrier
}

func (o *LookupResponse) GetCountryCode() *string {
	if o == nil {
		return nil
	}
	return o.CountryCode
}

func (o *LookupResponse) GetLineType() *LineType {
	if o == nil {
		return nil
	}
	return o.LineType
}

func (o *LookupResponse) GetMcc() *string {
	if o == nil {
		return nil
	}
	return o.Mcc
}

func (o *LookupResponse) GetMnc() *string {
	if o == nil {
		return nil
	}
	return o.Mnc
}

func (o *LookupResponse) GetNumberPorted() *bool {
	if o == nil {
		return nil
	}
	return o.NumberPorted
}

func (o *LookupResponse) GetPhoneNumber() *string {
	if o == nil {
		return nil
	}
	return o.PhoneNumber
}

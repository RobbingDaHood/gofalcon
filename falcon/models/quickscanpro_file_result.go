// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// QuickscanproFileResult quickscanpro file result
//
// swagger:model quickscanpro.FileResult
type QuickscanproFileResult struct {

	// filename
	Filename string `json:"filename,omitempty"`

	// sha256
	// Required: true
	Sha256 *string `json:"sha256"`

	// verdict
	// Required: true
	// Enum: [clean suspicious malicious unknown]
	Verdict *string `json:"verdict"`

	// verdict reason
	VerdictReason string `json:"verdict_reason,omitempty"`
}

// Validate validates this quickscanpro file result
func (m *QuickscanproFileResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSha256(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVerdict(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QuickscanproFileResult) validateSha256(formats strfmt.Registry) error {

	if err := validate.Required("sha256", "body", m.Sha256); err != nil {
		return err
	}

	return nil
}

var quickscanproFileResultTypeVerdictPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["clean","suspicious","malicious","unknown"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		quickscanproFileResultTypeVerdictPropEnum = append(quickscanproFileResultTypeVerdictPropEnum, v)
	}
}

const (

	// QuickscanproFileResultVerdictClean captures enum value "clean"
	QuickscanproFileResultVerdictClean string = "clean"

	// QuickscanproFileResultVerdictSuspicious captures enum value "suspicious"
	QuickscanproFileResultVerdictSuspicious string = "suspicious"

	// QuickscanproFileResultVerdictMalicious captures enum value "malicious"
	QuickscanproFileResultVerdictMalicious string = "malicious"

	// QuickscanproFileResultVerdictUnknown captures enum value "unknown"
	QuickscanproFileResultVerdictUnknown string = "unknown"
)

// prop value enum
func (m *QuickscanproFileResult) validateVerdictEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, quickscanproFileResultTypeVerdictPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *QuickscanproFileResult) validateVerdict(formats strfmt.Registry) error {

	if err := validate.Required("verdict", "body", m.Verdict); err != nil {
		return err
	}

	// value enum
	if err := m.validateVerdictEnum("verdict", "body", *m.Verdict); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this quickscanpro file result based on context it is used
func (m *QuickscanproFileResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *QuickscanproFileResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QuickscanproFileResult) UnmarshalBinary(b []byte) error {
	var res QuickscanproFileResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

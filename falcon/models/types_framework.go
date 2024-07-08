// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TypesFramework types framework
//
// swagger:model types.Framework
type TypesFramework struct {

	// accessing elements
	AccessingElements []string `json:"accessingElements"`

	// domain
	Domain string `json:"domain,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this types framework
func (m *TypesFramework) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this types framework based on context it is used
func (m *TypesFramework) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TypesFramework) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TypesFramework) UnmarshalBinary(b []byte) error {
	var res TypesFramework
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

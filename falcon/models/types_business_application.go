// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TypesBusinessApplication types business application
//
// swagger:model types.BusinessApplication
type TypesBusinessApplication struct {

	// id
	ID int64 `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this types business application
func (m *TypesBusinessApplication) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this types business application based on context it is used
func (m *TypesBusinessApplication) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TypesBusinessApplication) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TypesBusinessApplication) UnmarshalBinary(b []byte) error {
	var res TypesBusinessApplication
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TypesIntegrationTaskType types integration task type
//
// swagger:model types.IntegrationTaskType
type TypesIntegrationTaskType struct {

	// category
	Category string `json:"category,omitempty"`

	// display name
	DisplayName string `json:"display_name,omitempty"`

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// id
	ID int32 `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// required integration types
	RequiredIntegrationTypes []*TypesIntegrationType `json:"required_integration_types"`
}

// Validate validates this types integration task type
func (m *TypesIntegrationTaskType) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRequiredIntegrationTypes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TypesIntegrationTaskType) validateRequiredIntegrationTypes(formats strfmt.Registry) error {
	if swag.IsZero(m.RequiredIntegrationTypes) { // not required
		return nil
	}

	for i := 0; i < len(m.RequiredIntegrationTypes); i++ {
		if swag.IsZero(m.RequiredIntegrationTypes[i]) { // not required
			continue
		}

		if m.RequiredIntegrationTypes[i] != nil {
			if err := m.RequiredIntegrationTypes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("required_integration_types" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("required_integration_types" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this types integration task type based on the context it is used
func (m *TypesIntegrationTaskType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRequiredIntegrationTypes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TypesIntegrationTaskType) contextValidateRequiredIntegrationTypes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.RequiredIntegrationTypes); i++ {

		if m.RequiredIntegrationTypes[i] != nil {

			if swag.IsZero(m.RequiredIntegrationTypes[i]) { // not required
				return nil
			}

			if err := m.RequiredIntegrationTypes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("required_integration_types" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("required_integration_types" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TypesIntegrationTaskType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TypesIntegrationTaskType) UnmarshalBinary(b []byte) error {
	var res TypesIntegrationTaskType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

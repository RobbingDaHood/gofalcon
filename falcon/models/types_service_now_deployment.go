// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TypesServiceNowDeployment types service now deployment
//
// swagger:model types.ServiceNowDeployment
type TypesServiceNowDeployment struct {

	// deployment unit descriptor
	// Required: true
	DeploymentUnitDescriptor *TypesDeploymentUnitDescriptor `json:"DeploymentUnitDescriptor"`

	// type name
	// Required: true
	TypeName *string `json:"type_name"`
}

// Validate validates this types service now deployment
func (m *TypesServiceNowDeployment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeploymentUnitDescriptor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTypeName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TypesServiceNowDeployment) validateDeploymentUnitDescriptor(formats strfmt.Registry) error {

	if err := validate.Required("DeploymentUnitDescriptor", "body", m.DeploymentUnitDescriptor); err != nil {
		return err
	}

	if m.DeploymentUnitDescriptor != nil {
		if err := m.DeploymentUnitDescriptor.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("DeploymentUnitDescriptor")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("DeploymentUnitDescriptor")
			}
			return err
		}
	}

	return nil
}

func (m *TypesServiceNowDeployment) validateTypeName(formats strfmt.Registry) error {

	if err := validate.Required("type_name", "body", m.TypeName); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this types service now deployment based on the context it is used
func (m *TypesServiceNowDeployment) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDeploymentUnitDescriptor(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TypesServiceNowDeployment) contextValidateDeploymentUnitDescriptor(ctx context.Context, formats strfmt.Registry) error {

	if m.DeploymentUnitDescriptor != nil {

		if err := m.DeploymentUnitDescriptor.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("DeploymentUnitDescriptor")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("DeploymentUnitDescriptor")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TypesServiceNowDeployment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TypesServiceNowDeployment) UnmarshalBinary(b []byte) error {
	var res TypesServiceNowDeployment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

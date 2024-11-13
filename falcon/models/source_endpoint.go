// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SourceEndpoint source endpoint
//
// swagger:model .sourceEndpoint
type SourceEndpoint struct {

	// entity Id
	EntityID *TypesPolicyRulesCondition `json:"entityId,omitempty"`

	// group membership
	GroupMembership *TypesPolicyRulesCondition `json:"groupMembership,omitempty"`
}

// Validate validates this source endpoint
func (m *SourceEndpoint) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEntityID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroupMembership(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SourceEndpoint) validateEntityID(formats strfmt.Registry) error {
	if swag.IsZero(m.EntityID) { // not required
		return nil
	}

	if m.EntityID != nil {
		if err := m.EntityID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityId")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityId")
			}
			return err
		}
	}

	return nil
}

func (m *SourceEndpoint) validateGroupMembership(formats strfmt.Registry) error {
	if swag.IsZero(m.GroupMembership) { // not required
		return nil
	}

	if m.GroupMembership != nil {
		if err := m.GroupMembership.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("groupMembership")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("groupMembership")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this source endpoint based on the context it is used
func (m *SourceEndpoint) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEntityID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGroupMembership(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SourceEndpoint) contextValidateEntityID(ctx context.Context, formats strfmt.Registry) error {

	if m.EntityID != nil {

		if swag.IsZero(m.EntityID) { // not required
			return nil
		}

		if err := m.EntityID.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityId")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityId")
			}
			return err
		}
	}

	return nil
}

func (m *SourceEndpoint) contextValidateGroupMembership(ctx context.Context, formats strfmt.Registry) error {

	if m.GroupMembership != nil {

		if swag.IsZero(m.GroupMembership) { // not required
			return nil
		}

		if err := m.GroupMembership.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("groupMembership")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("groupMembership")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SourceEndpoint) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SourceEndpoint) UnmarshalBinary(b []byte) error {
	var res SourceEndpoint
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

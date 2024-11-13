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

// QuickscanproMITREAttack quickscanpro m i t r e attack
//
// swagger:model quickscanpro.MITREAttack
type QuickscanproMITREAttack struct {

	// attack id
	AttackID string `json:"attack_id,omitempty"`

	// attack id wiki
	AttackIDWiki string `json:"attack_id_wiki,omitempty"`

	// capec id
	CapecID string `json:"capec_id,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// parent
	Parent *QuickscanproMITREAttackParent `json:"parent,omitempty"`

	// tactic
	Tactic string `json:"tactic,omitempty"`

	// technique
	Technique string `json:"technique,omitempty"`
}

// Validate validates this quickscanpro m i t r e attack
func (m *QuickscanproMITREAttack) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateParent(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QuickscanproMITREAttack) validateParent(formats strfmt.Registry) error {
	if swag.IsZero(m.Parent) { // not required
		return nil
	}

	if m.Parent != nil {
		if err := m.Parent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("parent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("parent")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this quickscanpro m i t r e attack based on the context it is used
func (m *QuickscanproMITREAttack) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateParent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QuickscanproMITREAttack) contextValidateParent(ctx context.Context, formats strfmt.Registry) error {

	if m.Parent != nil {

		if swag.IsZero(m.Parent) { // not required
			return nil
		}

		if err := m.Parent.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("parent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("parent")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *QuickscanproMITREAttack) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QuickscanproMITREAttack) UnmarshalBinary(b []byte) error {
	var res QuickscanproMITREAttack
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

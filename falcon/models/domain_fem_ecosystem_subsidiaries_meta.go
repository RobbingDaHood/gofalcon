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

// DomainFemEcosystemSubsidiariesMeta domain fem ecosystem subsidiaries meta
//
// swagger:model domain.FemEcosystemSubsidiariesMeta
type DomainFemEcosystemSubsidiariesMeta struct {

	// meta info
	// Required: true
	MetaInfo *MsaMetaInfo `json:"MetaInfo"`

	// The version ID of an ecosystem
	// Required: true
	VersionID *string `json:"version_id"`
}

// Validate validates this domain fem ecosystem subsidiaries meta
func (m *DomainFemEcosystemSubsidiariesMeta) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMetaInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersionID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainFemEcosystemSubsidiariesMeta) validateMetaInfo(formats strfmt.Registry) error {

	if err := validate.Required("MetaInfo", "body", m.MetaInfo); err != nil {
		return err
	}

	if m.MetaInfo != nil {
		if err := m.MetaInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("MetaInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("MetaInfo")
			}
			return err
		}
	}

	return nil
}

func (m *DomainFemEcosystemSubsidiariesMeta) validateVersionID(formats strfmt.Registry) error {

	if err := validate.Required("version_id", "body", m.VersionID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this domain fem ecosystem subsidiaries meta based on the context it is used
func (m *DomainFemEcosystemSubsidiariesMeta) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMetaInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainFemEcosystemSubsidiariesMeta) contextValidateMetaInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.MetaInfo != nil {

		if err := m.MetaInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("MetaInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("MetaInfo")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DomainFemEcosystemSubsidiariesMeta) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainFemEcosystemSubsidiariesMeta) UnmarshalBinary(b []byte) error {
	var res DomainFemEcosystemSubsidiariesMeta
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

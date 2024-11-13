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

// DomainSpotlightParams domain spotlight params
//
// swagger:model domain.SpotlightParams
type DomainSpotlightParams struct {

	// template fields
	// Required: true
	TemplateFields []string `json:"template_fields"`

	// template name
	// Required: true
	TemplateName *string `json:"template_name"`

	// top n results
	// Required: true
	TopnResults *int32 `json:"top_n_results"`
}

// Validate validates this domain spotlight params
func (m *DomainSpotlightParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTemplateFields(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTemplateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTopnResults(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainSpotlightParams) validateTemplateFields(formats strfmt.Registry) error {

	if err := validate.Required("template_fields", "body", m.TemplateFields); err != nil {
		return err
	}

	return nil
}

func (m *DomainSpotlightParams) validateTemplateName(formats strfmt.Registry) error {

	if err := validate.Required("template_name", "body", m.TemplateName); err != nil {
		return err
	}

	return nil
}

func (m *DomainSpotlightParams) validateTopnResults(formats strfmt.Registry) error {

	if err := validate.Required("top_n_results", "body", m.TopnResults); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this domain spotlight params based on context it is used
func (m *DomainSpotlightParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DomainSpotlightParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainSpotlightParams) UnmarshalBinary(b []byte) error {
	var res DomainSpotlightParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

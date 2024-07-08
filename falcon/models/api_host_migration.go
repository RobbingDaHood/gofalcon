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
	"github.com/go-openapi/validate"
)

// APIHostMigration api host migration
//
// swagger:model api.HostMigration
type APIHostMigration struct {

	// created time
	// Required: true
	// Format: date-time
	CreatedTime *strfmt.DateTime `json:"created_time"`

	// events
	Events []*APIEvent `json:"events"`

	// host migration id
	// Required: true
	HostMigrationID *string `json:"host_migration_id"`

	// assigned static hostgroups, may need more details here
	// Required: true
	Hostgroups []string `json:"hostgroups"`

	// hostname at the time of migration
	// Required: true
	Hostname *string `json:"hostname"`

	// migration id
	// Required: true
	MigrationID *string `json:"migration_id"`

	// platform at the time of migration
	// Required: true
	Platform *string `json:"platform"`

	// source cid
	// Required: true
	SourceCid *string `json:"source_cid"`

	// device_id in the source cid
	// Required: true
	SourceDeviceID *string `json:"source_device_id"`

	// status
	// Required: true
	Status *string `json:"status"`

	// status details
	// Required: true
	StatusDetails *string `json:"status_details"`

	// not sure if this is necessary since it's common
	// Required: true
	TargetCid *string `json:"target_cid"`

	// device_id in the target cid. This may change while the migration is incomplete.
	// Required: true
	TargetDeviceID *string `json:"target_device_id"`

	// updated time
	// Required: true
	// Format: date-time
	UpdatedTime *strfmt.DateTime `json:"updated_time"`
}

// Validate validates this api host migration
func (m *APIHostMigration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEvents(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHostMigrationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHostgroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHostname(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMigrationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlatform(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceCid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceDeviceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatusDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetCid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetDeviceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIHostMigration) validateCreatedTime(formats strfmt.Registry) error {

	if err := validate.Required("created_time", "body", m.CreatedTime); err != nil {
		return err
	}

	if err := validate.FormatOf("created_time", "body", "date-time", m.CreatedTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *APIHostMigration) validateEvents(formats strfmt.Registry) error {
	if swag.IsZero(m.Events) { // not required
		return nil
	}

	for i := 0; i < len(m.Events); i++ {
		if swag.IsZero(m.Events[i]) { // not required
			continue
		}

		if m.Events[i] != nil {
			if err := m.Events[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("events" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("events" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *APIHostMigration) validateHostMigrationID(formats strfmt.Registry) error {

	if err := validate.Required("host_migration_id", "body", m.HostMigrationID); err != nil {
		return err
	}

	return nil
}

func (m *APIHostMigration) validateHostgroups(formats strfmt.Registry) error {

	if err := validate.Required("hostgroups", "body", m.Hostgroups); err != nil {
		return err
	}

	return nil
}

func (m *APIHostMigration) validateHostname(formats strfmt.Registry) error {

	if err := validate.Required("hostname", "body", m.Hostname); err != nil {
		return err
	}

	return nil
}

func (m *APIHostMigration) validateMigrationID(formats strfmt.Registry) error {

	if err := validate.Required("migration_id", "body", m.MigrationID); err != nil {
		return err
	}

	return nil
}

func (m *APIHostMigration) validatePlatform(formats strfmt.Registry) error {

	if err := validate.Required("platform", "body", m.Platform); err != nil {
		return err
	}

	return nil
}

func (m *APIHostMigration) validateSourceCid(formats strfmt.Registry) error {

	if err := validate.Required("source_cid", "body", m.SourceCid); err != nil {
		return err
	}

	return nil
}

func (m *APIHostMigration) validateSourceDeviceID(formats strfmt.Registry) error {

	if err := validate.Required("source_device_id", "body", m.SourceDeviceID); err != nil {
		return err
	}

	return nil
}

func (m *APIHostMigration) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *APIHostMigration) validateStatusDetails(formats strfmt.Registry) error {

	if err := validate.Required("status_details", "body", m.StatusDetails); err != nil {
		return err
	}

	return nil
}

func (m *APIHostMigration) validateTargetCid(formats strfmt.Registry) error {

	if err := validate.Required("target_cid", "body", m.TargetCid); err != nil {
		return err
	}

	return nil
}

func (m *APIHostMigration) validateTargetDeviceID(formats strfmt.Registry) error {

	if err := validate.Required("target_device_id", "body", m.TargetDeviceID); err != nil {
		return err
	}

	return nil
}

func (m *APIHostMigration) validateUpdatedTime(formats strfmt.Registry) error {

	if err := validate.Required("updated_time", "body", m.UpdatedTime); err != nil {
		return err
	}

	if err := validate.FormatOf("updated_time", "body", "date-time", m.UpdatedTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this api host migration based on the context it is used
func (m *APIHostMigration) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEvents(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIHostMigration) contextValidateEvents(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Events); i++ {

		if m.Events[i] != nil {

			if swag.IsZero(m.Events[i]) { // not required
				return nil
			}

			if err := m.Events[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("events" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("events" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *APIHostMigration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIHostMigration) UnmarshalBinary(b []byte) error {
	var res APIHostMigration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

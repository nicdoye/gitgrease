// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Member member
// swagger:model Member
type Member struct {

	// avatar
	Avatar *Avatar `json:"avatar,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// repositories
	Repositories []string `json:"repositories"`

	// teams
	Teams []*Team `json:"teams"`
}

// Validate validates this member
func (m *Member) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAvatar(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTeams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Member) validateAvatar(formats strfmt.Registry) error {

	if swag.IsZero(m.Avatar) { // not required
		return nil
	}

	if m.Avatar != nil {
		if err := m.Avatar.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("avatar")
			}
			return err
		}
	}

	return nil
}

func (m *Member) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *Member) validateTeams(formats strfmt.Registry) error {

	if swag.IsZero(m.Teams) { // not required
		return nil
	}

	for i := 0; i < len(m.Teams); i++ {
		if swag.IsZero(m.Teams[i]) { // not required
			continue
		}

		if m.Teams[i] != nil {
			if err := m.Teams[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("teams" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Member) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Member) UnmarshalBinary(b []byte) error {
	var res Member
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
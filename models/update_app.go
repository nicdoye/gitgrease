// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UpdateApp Description of an updated application.
// swagger:model UpdateApp
type UpdateApp struct {

	// The URI for the application's homepage
	// Required: true
	ApplicationURI *string `json:"application_uri"`

	// The e-mail address of the avatar to use for the application
	AvatarEmail string `json:"avatar_email,omitempty"`

	// The human-readable description for the application
	Description string `json:"description,omitempty"`

	// The name of the application
	// Required: true
	Name *string `json:"name"`

	// The URI for the application's OAuth redirect
	// Required: true
	RedirectURI *string `json:"redirect_uri"`
}

// Validate validates this update app
func (m *UpdateApp) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApplicationURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRedirectURI(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateApp) validateApplicationURI(formats strfmt.Registry) error {

	if err := validate.Required("application_uri", "body", m.ApplicationURI); err != nil {
		return err
	}

	return nil
}

func (m *UpdateApp) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *UpdateApp) validateRedirectURI(formats strfmt.Registry) error {

	if err := validate.Required("redirect_uri", "body", m.RedirectURI); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateApp) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateApp) UnmarshalBinary(b []byte) error {
	var res UpdateApp
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

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

// RepositoryBuildRequest Description of a new repository build.
// swagger:model RepositoryBuildRequest
type RepositoryBuildRequest struct {

	// The URL of the .tar.gz to build. Must start with "http" or "https".
	ArchiveURL string `json:"archive_url,omitempty"`

	// The tags to which the built images will be pushed. If none specified, "latest" is used.
	// Min Items: 1
	// Unique: true
	DockerTags []string `json:"docker_tags"`

	// The file id that was generated when the build spec was uploaded
	FileID string `json:"file_id,omitempty"`

	// Username of a Quay robot account to use as pull credentials
	PullRobot string `json:"pull_robot,omitempty"`

	// Subdirectory in which the Dockerfile can be found
	Subdirectory string `json:"subdirectory,omitempty"`
}

// Validate validates this repository build request
func (m *RepositoryBuildRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDockerTags(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RepositoryBuildRequest) validateDockerTags(formats strfmt.Registry) error {

	if swag.IsZero(m.DockerTags) { // not required
		return nil
	}

	iDockerTagsSize := int64(len(m.DockerTags))

	if err := validate.MinItems("docker_tags", "body", iDockerTagsSize, 1); err != nil {
		return err
	}

	if err := validate.UniqueItems("docker_tags", "body", m.DockerTags); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RepositoryBuildRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RepositoryBuildRequest) UnmarshalBinary(b []byte) error {
	var res RepositoryBuildRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

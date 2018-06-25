// Code generated by go-swagger; DO NOT EDIT.

package organization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/nicdoye/quaytermass/models"
)

// GetOrganizationApplicationReader is a Reader for the GetOrganizationApplication structure.
type GetOrganizationApplicationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrganizationApplicationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetOrganizationApplicationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetOrganizationApplicationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetOrganizationApplicationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetOrganizationApplicationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetOrganizationApplicationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetOrganizationApplicationOK creates a GetOrganizationApplicationOK with default headers values
func NewGetOrganizationApplicationOK() *GetOrganizationApplicationOK {
	return &GetOrganizationApplicationOK{}
}

/*GetOrganizationApplicationOK handles this case with default header values.

Successful invocation
*/
type GetOrganizationApplicationOK struct {
}

func (o *GetOrganizationApplicationOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/organization/{orgname}/applications/{client_id}][%d] getOrganizationApplicationOK ", 200)
}

func (o *GetOrganizationApplicationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetOrganizationApplicationBadRequest creates a GetOrganizationApplicationBadRequest with default headers values
func NewGetOrganizationApplicationBadRequest() *GetOrganizationApplicationBadRequest {
	return &GetOrganizationApplicationBadRequest{}
}

/*GetOrganizationApplicationBadRequest handles this case with default header values.

Bad Request
*/
type GetOrganizationApplicationBadRequest struct {
	Payload *models.APIError
}

func (o *GetOrganizationApplicationBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v1/organization/{orgname}/applications/{client_id}][%d] getOrganizationApplicationBadRequest  %+v", 400, o.Payload)
}

func (o *GetOrganizationApplicationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationApplicationUnauthorized creates a GetOrganizationApplicationUnauthorized with default headers values
func NewGetOrganizationApplicationUnauthorized() *GetOrganizationApplicationUnauthorized {
	return &GetOrganizationApplicationUnauthorized{}
}

/*GetOrganizationApplicationUnauthorized handles this case with default header values.

Session required
*/
type GetOrganizationApplicationUnauthorized struct {
	Payload *models.APIError
}

func (o *GetOrganizationApplicationUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/organization/{orgname}/applications/{client_id}][%d] getOrganizationApplicationUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOrganizationApplicationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationApplicationForbidden creates a GetOrganizationApplicationForbidden with default headers values
func NewGetOrganizationApplicationForbidden() *GetOrganizationApplicationForbidden {
	return &GetOrganizationApplicationForbidden{}
}

/*GetOrganizationApplicationForbidden handles this case with default header values.

Unauthorized access
*/
type GetOrganizationApplicationForbidden struct {
	Payload *models.APIError
}

func (o *GetOrganizationApplicationForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/organization/{orgname}/applications/{client_id}][%d] getOrganizationApplicationForbidden  %+v", 403, o.Payload)
}

func (o *GetOrganizationApplicationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationApplicationNotFound creates a GetOrganizationApplicationNotFound with default headers values
func NewGetOrganizationApplicationNotFound() *GetOrganizationApplicationNotFound {
	return &GetOrganizationApplicationNotFound{}
}

/*GetOrganizationApplicationNotFound handles this case with default header values.

Not found
*/
type GetOrganizationApplicationNotFound struct {
	Payload *models.APIError
}

func (o *GetOrganizationApplicationNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/organization/{orgname}/applications/{client_id}][%d] getOrganizationApplicationNotFound  %+v", 404, o.Payload)
}

func (o *GetOrganizationApplicationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

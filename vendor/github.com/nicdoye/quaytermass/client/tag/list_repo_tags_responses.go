// Code generated by go-swagger; DO NOT EDIT.

package tag

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/nicdoye/quaytermass/models"
)

// ListRepoTagsReader is a Reader for the ListRepoTags structure.
type ListRepoTagsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListRepoTagsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListRepoTagsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewListRepoTagsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewListRepoTagsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewListRepoTagsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewListRepoTagsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListRepoTagsOK creates a ListRepoTagsOK with default headers values
func NewListRepoTagsOK() *ListRepoTagsOK {
	return &ListRepoTagsOK{}
}

/*ListRepoTagsOK handles this case with default header values.

Successful invocation
*/
type ListRepoTagsOK struct {
}

func (o *ListRepoTagsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/repository/{repository}/tag/][%d] listRepoTagsOK ", 200)
}

func (o *ListRepoTagsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListRepoTagsBadRequest creates a ListRepoTagsBadRequest with default headers values
func NewListRepoTagsBadRequest() *ListRepoTagsBadRequest {
	return &ListRepoTagsBadRequest{}
}

/*ListRepoTagsBadRequest handles this case with default header values.

Bad Request
*/
type ListRepoTagsBadRequest struct {
	Payload *models.APIError
}

func (o *ListRepoTagsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v1/repository/{repository}/tag/][%d] listRepoTagsBadRequest  %+v", 400, o.Payload)
}

func (o *ListRepoTagsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListRepoTagsUnauthorized creates a ListRepoTagsUnauthorized with default headers values
func NewListRepoTagsUnauthorized() *ListRepoTagsUnauthorized {
	return &ListRepoTagsUnauthorized{}
}

/*ListRepoTagsUnauthorized handles this case with default header values.

Session required
*/
type ListRepoTagsUnauthorized struct {
	Payload *models.APIError
}

func (o *ListRepoTagsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/repository/{repository}/tag/][%d] listRepoTagsUnauthorized  %+v", 401, o.Payload)
}

func (o *ListRepoTagsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListRepoTagsForbidden creates a ListRepoTagsForbidden with default headers values
func NewListRepoTagsForbidden() *ListRepoTagsForbidden {
	return &ListRepoTagsForbidden{}
}

/*ListRepoTagsForbidden handles this case with default header values.

Unauthorized access
*/
type ListRepoTagsForbidden struct {
	Payload *models.APIError
}

func (o *ListRepoTagsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/repository/{repository}/tag/][%d] listRepoTagsForbidden  %+v", 403, o.Payload)
}

func (o *ListRepoTagsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListRepoTagsNotFound creates a ListRepoTagsNotFound with default headers values
func NewListRepoTagsNotFound() *ListRepoTagsNotFound {
	return &ListRepoTagsNotFound{}
}

/*ListRepoTagsNotFound handles this case with default header values.

Not found
*/
type ListRepoTagsNotFound struct {
	Payload *models.APIError
}

func (o *ListRepoTagsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/repository/{repository}/tag/][%d] listRepoTagsNotFound  %+v", 404, o.Payload)
}

func (o *ListRepoTagsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
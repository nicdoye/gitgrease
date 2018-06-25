// Code generated by go-swagger; DO NOT EDIT.

package billing

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/nicdoye/quaytermass/models"
)

// ListOrgInvoicesReader is a Reader for the ListOrgInvoices structure.
type ListOrgInvoicesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListOrgInvoicesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListOrgInvoicesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewListOrgInvoicesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewListOrgInvoicesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewListOrgInvoicesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewListOrgInvoicesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListOrgInvoicesOK creates a ListOrgInvoicesOK with default headers values
func NewListOrgInvoicesOK() *ListOrgInvoicesOK {
	return &ListOrgInvoicesOK{}
}

/*ListOrgInvoicesOK handles this case with default header values.

Successful invocation
*/
type ListOrgInvoicesOK struct {
}

func (o *ListOrgInvoicesOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/organization/{orgname}/invoices][%d] listOrgInvoicesOK ", 200)
}

func (o *ListOrgInvoicesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListOrgInvoicesBadRequest creates a ListOrgInvoicesBadRequest with default headers values
func NewListOrgInvoicesBadRequest() *ListOrgInvoicesBadRequest {
	return &ListOrgInvoicesBadRequest{}
}

/*ListOrgInvoicesBadRequest handles this case with default header values.

Bad Request
*/
type ListOrgInvoicesBadRequest struct {
	Payload *models.APIError
}

func (o *ListOrgInvoicesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v1/organization/{orgname}/invoices][%d] listOrgInvoicesBadRequest  %+v", 400, o.Payload)
}

func (o *ListOrgInvoicesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListOrgInvoicesUnauthorized creates a ListOrgInvoicesUnauthorized with default headers values
func NewListOrgInvoicesUnauthorized() *ListOrgInvoicesUnauthorized {
	return &ListOrgInvoicesUnauthorized{}
}

/*ListOrgInvoicesUnauthorized handles this case with default header values.

Session required
*/
type ListOrgInvoicesUnauthorized struct {
	Payload *models.APIError
}

func (o *ListOrgInvoicesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/organization/{orgname}/invoices][%d] listOrgInvoicesUnauthorized  %+v", 401, o.Payload)
}

func (o *ListOrgInvoicesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListOrgInvoicesForbidden creates a ListOrgInvoicesForbidden with default headers values
func NewListOrgInvoicesForbidden() *ListOrgInvoicesForbidden {
	return &ListOrgInvoicesForbidden{}
}

/*ListOrgInvoicesForbidden handles this case with default header values.

Unauthorized access
*/
type ListOrgInvoicesForbidden struct {
	Payload *models.APIError
}

func (o *ListOrgInvoicesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/organization/{orgname}/invoices][%d] listOrgInvoicesForbidden  %+v", 403, o.Payload)
}

func (o *ListOrgInvoicesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListOrgInvoicesNotFound creates a ListOrgInvoicesNotFound with default headers values
func NewListOrgInvoicesNotFound() *ListOrgInvoicesNotFound {
	return &ListOrgInvoicesNotFound{}
}

/*ListOrgInvoicesNotFound handles this case with default header values.

Not found
*/
type ListOrgInvoicesNotFound struct {
	Payload *models.APIError
}

func (o *ListOrgInvoicesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/organization/{orgname}/invoices][%d] listOrgInvoicesNotFound  %+v", 404, o.Payload)
}

func (o *ListOrgInvoicesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

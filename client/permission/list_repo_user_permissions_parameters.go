// Code generated by go-swagger; DO NOT EDIT.

package permission

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewListRepoUserPermissionsParams creates a new ListRepoUserPermissionsParams object
// with the default values initialized.
func NewListRepoUserPermissionsParams() *ListRepoUserPermissionsParams {
	var ()
	return &ListRepoUserPermissionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListRepoUserPermissionsParamsWithTimeout creates a new ListRepoUserPermissionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListRepoUserPermissionsParamsWithTimeout(timeout time.Duration) *ListRepoUserPermissionsParams {
	var ()
	return &ListRepoUserPermissionsParams{

		timeout: timeout,
	}
}

// NewListRepoUserPermissionsParamsWithContext creates a new ListRepoUserPermissionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListRepoUserPermissionsParamsWithContext(ctx context.Context) *ListRepoUserPermissionsParams {
	var ()
	return &ListRepoUserPermissionsParams{

		Context: ctx,
	}
}

// NewListRepoUserPermissionsParamsWithHTTPClient creates a new ListRepoUserPermissionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListRepoUserPermissionsParamsWithHTTPClient(client *http.Client) *ListRepoUserPermissionsParams {
	var ()
	return &ListRepoUserPermissionsParams{
		HTTPClient: client,
	}
}

/*ListRepoUserPermissionsParams contains all the parameters to send to the API endpoint
for the list repo user permissions operation typically these are written to a http.Request
*/
type ListRepoUserPermissionsParams struct {

	/*Repository
	  The full path of the repository. e.g. namespace/name

	*/
	Repository string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list repo user permissions params
func (o *ListRepoUserPermissionsParams) WithTimeout(timeout time.Duration) *ListRepoUserPermissionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list repo user permissions params
func (o *ListRepoUserPermissionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list repo user permissions params
func (o *ListRepoUserPermissionsParams) WithContext(ctx context.Context) *ListRepoUserPermissionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list repo user permissions params
func (o *ListRepoUserPermissionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list repo user permissions params
func (o *ListRepoUserPermissionsParams) WithHTTPClient(client *http.Client) *ListRepoUserPermissionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list repo user permissions params
func (o *ListRepoUserPermissionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRepository adds the repository to the list repo user permissions params
func (o *ListRepoUserPermissionsParams) WithRepository(repository string) *ListRepoUserPermissionsParams {
	o.SetRepository(repository)
	return o
}

// SetRepository adds the repository to the list repo user permissions params
func (o *ListRepoUserPermissionsParams) SetRepository(repository string) {
	o.Repository = repository
}

// WriteToRequest writes these params to a swagger request
func (o *ListRepoUserPermissionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param repository
	if err := r.SetPathParam("repository", o.Repository); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
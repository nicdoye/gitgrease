// Code generated by go-swagger; DO NOT EDIT.

package repotoken

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

	models "github.com/nicdoye/quaytermass/models"
)

// NewCreateTokenParams creates a new CreateTokenParams object
// with the default values initialized.
func NewCreateTokenParams() *CreateTokenParams {
	var ()
	return &CreateTokenParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateTokenParamsWithTimeout creates a new CreateTokenParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateTokenParamsWithTimeout(timeout time.Duration) *CreateTokenParams {
	var ()
	return &CreateTokenParams{

		timeout: timeout,
	}
}

// NewCreateTokenParamsWithContext creates a new CreateTokenParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateTokenParamsWithContext(ctx context.Context) *CreateTokenParams {
	var ()
	return &CreateTokenParams{

		Context: ctx,
	}
}

// NewCreateTokenParamsWithHTTPClient creates a new CreateTokenParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateTokenParamsWithHTTPClient(client *http.Client) *CreateTokenParams {
	var ()
	return &CreateTokenParams{
		HTTPClient: client,
	}
}

/*CreateTokenParams contains all the parameters to send to the API endpoint
for the create token operation typically these are written to a http.Request
*/
type CreateTokenParams struct {

	/*Body
	  Request body contents.

	*/
	Body *models.NewToken
	/*Repository
	  The full path of the repository. e.g. namespace/name

	*/
	Repository string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create token params
func (o *CreateTokenParams) WithTimeout(timeout time.Duration) *CreateTokenParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create token params
func (o *CreateTokenParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create token params
func (o *CreateTokenParams) WithContext(ctx context.Context) *CreateTokenParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create token params
func (o *CreateTokenParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create token params
func (o *CreateTokenParams) WithHTTPClient(client *http.Client) *CreateTokenParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create token params
func (o *CreateTokenParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create token params
func (o *CreateTokenParams) WithBody(body *models.NewToken) *CreateTokenParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create token params
func (o *CreateTokenParams) SetBody(body *models.NewToken) {
	o.Body = body
}

// WithRepository adds the repository to the create token params
func (o *CreateTokenParams) WithRepository(repository string) *CreateTokenParams {
	o.SetRepository(repository)
	return o
}

// SetRepository adds the repository to the create token params
func (o *CreateTokenParams) SetRepository(repository string) {
	o.Repository = repository
}

// WriteToRequest writes these params to a swagger request
func (o *CreateTokenParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param repository
	if err := r.SetPathParam("repository", o.Repository); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

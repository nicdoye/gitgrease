// Code generated by go-swagger; DO NOT EDIT.

package tag

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

// NewRevertTagParams creates a new RevertTagParams object
// with the default values initialized.
func NewRevertTagParams() *RevertTagParams {
	var ()
	return &RevertTagParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRevertTagParamsWithTimeout creates a new RevertTagParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRevertTagParamsWithTimeout(timeout time.Duration) *RevertTagParams {
	var ()
	return &RevertTagParams{

		timeout: timeout,
	}
}

// NewRevertTagParamsWithContext creates a new RevertTagParams object
// with the default values initialized, and the ability to set a context for a request
func NewRevertTagParamsWithContext(ctx context.Context) *RevertTagParams {
	var ()
	return &RevertTagParams{

		Context: ctx,
	}
}

// NewRevertTagParamsWithHTTPClient creates a new RevertTagParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRevertTagParamsWithHTTPClient(client *http.Client) *RevertTagParams {
	var ()
	return &RevertTagParams{
		HTTPClient: client,
	}
}

/*RevertTagParams contains all the parameters to send to the API endpoint
for the revert tag operation typically these are written to a http.Request
*/
type RevertTagParams struct {

	/*Body
	  Request body contents.

	*/
	Body *models.RevertTag
	/*Repository
	  The full path of the repository. e.g. namespace/name

	*/
	Repository string
	/*Tag
	  The name of the tag

	*/
	Tag string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the revert tag params
func (o *RevertTagParams) WithTimeout(timeout time.Duration) *RevertTagParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the revert tag params
func (o *RevertTagParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the revert tag params
func (o *RevertTagParams) WithContext(ctx context.Context) *RevertTagParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the revert tag params
func (o *RevertTagParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the revert tag params
func (o *RevertTagParams) WithHTTPClient(client *http.Client) *RevertTagParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the revert tag params
func (o *RevertTagParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the revert tag params
func (o *RevertTagParams) WithBody(body *models.RevertTag) *RevertTagParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the revert tag params
func (o *RevertTagParams) SetBody(body *models.RevertTag) {
	o.Body = body
}

// WithRepository adds the repository to the revert tag params
func (o *RevertTagParams) WithRepository(repository string) *RevertTagParams {
	o.SetRepository(repository)
	return o
}

// SetRepository adds the repository to the revert tag params
func (o *RevertTagParams) SetRepository(repository string) {
	o.Repository = repository
}

// WithTag adds the tag to the revert tag params
func (o *RevertTagParams) WithTag(tag string) *RevertTagParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the revert tag params
func (o *RevertTagParams) SetTag(tag string) {
	o.Tag = tag
}

// WriteToRequest writes these params to a swagger request
func (o *RevertTagParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param tag
	if err := r.SetPathParam("tag", o.Tag); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
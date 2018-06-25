// Code generated by go-swagger; DO NOT EDIT.

package discovery

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDiscoveryParams creates a new DiscoveryParams object
// with the default values initialized.
func NewDiscoveryParams() *DiscoveryParams {
	var ()
	return &DiscoveryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDiscoveryParamsWithTimeout creates a new DiscoveryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDiscoveryParamsWithTimeout(timeout time.Duration) *DiscoveryParams {
	var ()
	return &DiscoveryParams{

		timeout: timeout,
	}
}

// NewDiscoveryParamsWithContext creates a new DiscoveryParams object
// with the default values initialized, and the ability to set a context for a request
func NewDiscoveryParamsWithContext(ctx context.Context) *DiscoveryParams {
	var ()
	return &DiscoveryParams{

		Context: ctx,
	}
}

// NewDiscoveryParamsWithHTTPClient creates a new DiscoveryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDiscoveryParamsWithHTTPClient(client *http.Client) *DiscoveryParams {
	var ()
	return &DiscoveryParams{
		HTTPClient: client,
	}
}

/*DiscoveryParams contains all the parameters to send to the API endpoint
for the discovery operation typically these are written to a http.Request
*/
type DiscoveryParams struct {

	/*Internal
	  Whether to include internal APIs.

	*/
	Internal *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the discovery params
func (o *DiscoveryParams) WithTimeout(timeout time.Duration) *DiscoveryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the discovery params
func (o *DiscoveryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the discovery params
func (o *DiscoveryParams) WithContext(ctx context.Context) *DiscoveryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the discovery params
func (o *DiscoveryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the discovery params
func (o *DiscoveryParams) WithHTTPClient(client *http.Client) *DiscoveryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the discovery params
func (o *DiscoveryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInternal adds the internal to the discovery params
func (o *DiscoveryParams) WithInternal(internal *bool) *DiscoveryParams {
	o.SetInternal(internal)
	return o
}

// SetInternal adds the internal to the discovery params
func (o *DiscoveryParams) SetInternal(internal *bool) {
	o.Internal = internal
}

// WriteToRequest writes these params to a swagger request
func (o *DiscoveryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Internal != nil {

		// query param internal
		var qrInternal bool
		if o.Internal != nil {
			qrInternal = *o.Internal
		}
		qInternal := swag.FormatBool(qrInternal)
		if qInternal != "" {
			if err := r.SetQueryParam("internal", qInternal); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

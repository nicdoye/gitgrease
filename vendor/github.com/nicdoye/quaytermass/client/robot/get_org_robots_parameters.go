// Code generated by go-swagger; DO NOT EDIT.

package robot

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

// NewGetOrgRobotsParams creates a new GetOrgRobotsParams object
// with the default values initialized.
func NewGetOrgRobotsParams() *GetOrgRobotsParams {
	var ()
	return &GetOrgRobotsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrgRobotsParamsWithTimeout creates a new GetOrgRobotsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetOrgRobotsParamsWithTimeout(timeout time.Duration) *GetOrgRobotsParams {
	var ()
	return &GetOrgRobotsParams{

		timeout: timeout,
	}
}

// NewGetOrgRobotsParamsWithContext creates a new GetOrgRobotsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetOrgRobotsParamsWithContext(ctx context.Context) *GetOrgRobotsParams {
	var ()
	return &GetOrgRobotsParams{

		Context: ctx,
	}
}

// NewGetOrgRobotsParamsWithHTTPClient creates a new GetOrgRobotsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetOrgRobotsParamsWithHTTPClient(client *http.Client) *GetOrgRobotsParams {
	var ()
	return &GetOrgRobotsParams{
		HTTPClient: client,
	}
}

/*GetOrgRobotsParams contains all the parameters to send to the API endpoint
for the get org robots operation typically these are written to a http.Request
*/
type GetOrgRobotsParams struct {

	/*Orgname
	  The name of the organization

	*/
	Orgname string
	/*Permissions
	  Whether to include repostories and teams in which the robots have permission.

	*/
	Permissions *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get org robots params
func (o *GetOrgRobotsParams) WithTimeout(timeout time.Duration) *GetOrgRobotsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get org robots params
func (o *GetOrgRobotsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get org robots params
func (o *GetOrgRobotsParams) WithContext(ctx context.Context) *GetOrgRobotsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get org robots params
func (o *GetOrgRobotsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get org robots params
func (o *GetOrgRobotsParams) WithHTTPClient(client *http.Client) *GetOrgRobotsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get org robots params
func (o *GetOrgRobotsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrgname adds the orgname to the get org robots params
func (o *GetOrgRobotsParams) WithOrgname(orgname string) *GetOrgRobotsParams {
	o.SetOrgname(orgname)
	return o
}

// SetOrgname adds the orgname to the get org robots params
func (o *GetOrgRobotsParams) SetOrgname(orgname string) {
	o.Orgname = orgname
}

// WithPermissions adds the permissions to the get org robots params
func (o *GetOrgRobotsParams) WithPermissions(permissions *bool) *GetOrgRobotsParams {
	o.SetPermissions(permissions)
	return o
}

// SetPermissions adds the permissions to the get org robots params
func (o *GetOrgRobotsParams) SetPermissions(permissions *bool) {
	o.Permissions = permissions
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrgRobotsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param orgname
	if err := r.SetPathParam("orgname", o.Orgname); err != nil {
		return err
	}

	if o.Permissions != nil {

		// query param permissions
		var qrPermissions bool
		if o.Permissions != nil {
			qrPermissions = *o.Permissions
		}
		qPermissions := swag.FormatBool(qrPermissions)
		if qPermissions != "" {
			if err := r.SetQueryParam("permissions", qPermissions); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

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

// NewGetUserRobotsParams creates a new GetUserRobotsParams object
// with the default values initialized.
func NewGetUserRobotsParams() *GetUserRobotsParams {
	var ()
	return &GetUserRobotsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetUserRobotsParamsWithTimeout creates a new GetUserRobotsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetUserRobotsParamsWithTimeout(timeout time.Duration) *GetUserRobotsParams {
	var ()
	return &GetUserRobotsParams{

		timeout: timeout,
	}
}

// NewGetUserRobotsParamsWithContext creates a new GetUserRobotsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetUserRobotsParamsWithContext(ctx context.Context) *GetUserRobotsParams {
	var ()
	return &GetUserRobotsParams{

		Context: ctx,
	}
}

// NewGetUserRobotsParamsWithHTTPClient creates a new GetUserRobotsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetUserRobotsParamsWithHTTPClient(client *http.Client) *GetUserRobotsParams {
	var ()
	return &GetUserRobotsParams{
		HTTPClient: client,
	}
}

/*GetUserRobotsParams contains all the parameters to send to the API endpoint
for the get user robots operation typically these are written to a http.Request
*/
type GetUserRobotsParams struct {

	/*Permissions
	  Whether to include repostories and teams in which the robots have permission.

	*/
	Permissions *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get user robots params
func (o *GetUserRobotsParams) WithTimeout(timeout time.Duration) *GetUserRobotsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get user robots params
func (o *GetUserRobotsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get user robots params
func (o *GetUserRobotsParams) WithContext(ctx context.Context) *GetUserRobotsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get user robots params
func (o *GetUserRobotsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get user robots params
func (o *GetUserRobotsParams) WithHTTPClient(client *http.Client) *GetUserRobotsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get user robots params
func (o *GetUserRobotsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPermissions adds the permissions to the get user robots params
func (o *GetUserRobotsParams) WithPermissions(permissions *bool) *GetUserRobotsParams {
	o.SetPermissions(permissions)
	return o
}

// SetPermissions adds the permissions to the get user robots params
func (o *GetUserRobotsParams) SetPermissions(permissions *bool) {
	o.Permissions = permissions
}

// WriteToRequest writes these params to a swagger request
func (o *GetUserRobotsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
// Code generated by go-swagger; DO NOT EDIT.

package s_a_n

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewLunDeleteParams creates a new LunDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLunDeleteParams() *LunDeleteParams {
	return &LunDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLunDeleteParamsWithTimeout creates a new LunDeleteParams object
// with the ability to set a timeout on a request.
func NewLunDeleteParamsWithTimeout(timeout time.Duration) *LunDeleteParams {
	return &LunDeleteParams{
		timeout: timeout,
	}
}

// NewLunDeleteParamsWithContext creates a new LunDeleteParams object
// with the ability to set a context for a request.
func NewLunDeleteParamsWithContext(ctx context.Context) *LunDeleteParams {
	return &LunDeleteParams{
		Context: ctx,
	}
}

// NewLunDeleteParamsWithHTTPClient creates a new LunDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewLunDeleteParamsWithHTTPClient(client *http.Client) *LunDeleteParams {
	return &LunDeleteParams{
		HTTPClient: client,
	}
}

/*
LunDeleteParams contains all the parameters to send to the API endpoint

	for the lun delete operation.

	Typically these are written to a http.Request.
*/
type LunDeleteParams struct {

	/* AllowDeleteWhileMapped.

	     Allows deletion of a mapped LUN.</br>
	A mapped LUN might be in use. Deleting a mapped LUN also deletes the LUN map and makes the data no longer available. This might cause a disruption in the availability of data.</br>
	**This parameter should be used with caution.**

	*/
	AllowDeleteWhileMapped *bool

	/* UUID.

	   The unique identifier of the LUN to retrieve.

	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the lun delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LunDeleteParams) WithDefaults() *LunDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the lun delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LunDeleteParams) SetDefaults() {
	var (
		allowDeleteWhileMappedDefault = bool(false)
	)

	val := LunDeleteParams{
		AllowDeleteWhileMapped: &allowDeleteWhileMappedDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the lun delete params
func (o *LunDeleteParams) WithTimeout(timeout time.Duration) *LunDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the lun delete params
func (o *LunDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the lun delete params
func (o *LunDeleteParams) WithContext(ctx context.Context) *LunDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the lun delete params
func (o *LunDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the lun delete params
func (o *LunDeleteParams) WithHTTPClient(client *http.Client) *LunDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the lun delete params
func (o *LunDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAllowDeleteWhileMapped adds the allowDeleteWhileMapped to the lun delete params
func (o *LunDeleteParams) WithAllowDeleteWhileMapped(allowDeleteWhileMapped *bool) *LunDeleteParams {
	o.SetAllowDeleteWhileMapped(allowDeleteWhileMapped)
	return o
}

// SetAllowDeleteWhileMapped adds the allowDeleteWhileMapped to the lun delete params
func (o *LunDeleteParams) SetAllowDeleteWhileMapped(allowDeleteWhileMapped *bool) {
	o.AllowDeleteWhileMapped = allowDeleteWhileMapped
}

// WithUUID adds the uuid to the lun delete params
func (o *LunDeleteParams) WithUUID(uuid string) *LunDeleteParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the lun delete params
func (o *LunDeleteParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *LunDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AllowDeleteWhileMapped != nil {

		// query param allow_delete_while_mapped
		var qrAllowDeleteWhileMapped bool

		if o.AllowDeleteWhileMapped != nil {
			qrAllowDeleteWhileMapped = *o.AllowDeleteWhileMapped
		}
		qAllowDeleteWhileMapped := swag.FormatBool(qrAllowDeleteWhileMapped)
		if qAllowDeleteWhileMapped != "" {

			if err := r.SetQueryParam("allow_delete_while_mapped", qAllowDeleteWhileMapped); err != nil {
				return err
			}
		}
	}

	// path param uuid
	if err := r.SetPathParam("uuid", o.UUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

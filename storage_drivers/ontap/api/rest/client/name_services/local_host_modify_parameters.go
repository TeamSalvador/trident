// Code generated by go-swagger; DO NOT EDIT.

package name_services

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

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// NewLocalHostModifyParams creates a new LocalHostModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLocalHostModifyParams() *LocalHostModifyParams {
	return &LocalHostModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLocalHostModifyParamsWithTimeout creates a new LocalHostModifyParams object
// with the ability to set a timeout on a request.
func NewLocalHostModifyParamsWithTimeout(timeout time.Duration) *LocalHostModifyParams {
	return &LocalHostModifyParams{
		timeout: timeout,
	}
}

// NewLocalHostModifyParamsWithContext creates a new LocalHostModifyParams object
// with the ability to set a context for a request.
func NewLocalHostModifyParamsWithContext(ctx context.Context) *LocalHostModifyParams {
	return &LocalHostModifyParams{
		Context: ctx,
	}
}

// NewLocalHostModifyParamsWithHTTPClient creates a new LocalHostModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewLocalHostModifyParamsWithHTTPClient(client *http.Client) *LocalHostModifyParams {
	return &LocalHostModifyParams{
		HTTPClient: client,
	}
}

/* LocalHostModifyParams contains all the parameters to send to the API endpoint
   for the local host modify operation.

   Typically these are written to a http.Request.
*/
type LocalHostModifyParams struct {

	/* Address.

	   The IP address.
	*/
	AddressPathParameter string

	/* Info.

	   Info specification
	*/
	Info *models.LocalHost

	/* OwnerUUID.

	   UUID of the owner to which this object belongs.
	*/
	OwnerUUIDPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the local host modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LocalHostModifyParams) WithDefaults() *LocalHostModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the local host modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LocalHostModifyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the local host modify params
func (o *LocalHostModifyParams) WithTimeout(timeout time.Duration) *LocalHostModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the local host modify params
func (o *LocalHostModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the local host modify params
func (o *LocalHostModifyParams) WithContext(ctx context.Context) *LocalHostModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the local host modify params
func (o *LocalHostModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the local host modify params
func (o *LocalHostModifyParams) WithHTTPClient(client *http.Client) *LocalHostModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the local host modify params
func (o *LocalHostModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAddressPathParameter adds the address to the local host modify params
func (o *LocalHostModifyParams) WithAddressPathParameter(address string) *LocalHostModifyParams {
	o.SetAddressPathParameter(address)
	return o
}

// SetAddressPathParameter adds the address to the local host modify params
func (o *LocalHostModifyParams) SetAddressPathParameter(address string) {
	o.AddressPathParameter = address
}

// WithInfo adds the info to the local host modify params
func (o *LocalHostModifyParams) WithInfo(info *models.LocalHost) *LocalHostModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the local host modify params
func (o *LocalHostModifyParams) SetInfo(info *models.LocalHost) {
	o.Info = info
}

// WithOwnerUUIDPathParameter adds the ownerUUID to the local host modify params
func (o *LocalHostModifyParams) WithOwnerUUIDPathParameter(ownerUUID string) *LocalHostModifyParams {
	o.SetOwnerUUIDPathParameter(ownerUUID)
	return o
}

// SetOwnerUUIDPathParameter adds the ownerUuid to the local host modify params
func (o *LocalHostModifyParams) SetOwnerUUIDPathParameter(ownerUUID string) {
	o.OwnerUUIDPathParameter = ownerUUID
}

// WriteToRequest writes these params to a swagger request
func (o *LocalHostModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param address
	if err := r.SetPathParam("address", o.AddressPathParameter); err != nil {
		return err
	}
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
			return err
		}
	}

	// path param owner.uuid
	if err := r.SetPathParam("owner.uuid", o.OwnerUUIDPathParameter); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
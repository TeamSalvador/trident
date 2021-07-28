// Code generated by go-swagger; DO NOT EDIT.

package support

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

// NewEmsFilterRuleCollectionGetParams creates a new EmsFilterRuleCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEmsFilterRuleCollectionGetParams() *EmsFilterRuleCollectionGetParams {
	return &EmsFilterRuleCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEmsFilterRuleCollectionGetParamsWithTimeout creates a new EmsFilterRuleCollectionGetParams object
// with the ability to set a timeout on a request.
func NewEmsFilterRuleCollectionGetParamsWithTimeout(timeout time.Duration) *EmsFilterRuleCollectionGetParams {
	return &EmsFilterRuleCollectionGetParams{
		timeout: timeout,
	}
}

// NewEmsFilterRuleCollectionGetParamsWithContext creates a new EmsFilterRuleCollectionGetParams object
// with the ability to set a context for a request.
func NewEmsFilterRuleCollectionGetParamsWithContext(ctx context.Context) *EmsFilterRuleCollectionGetParams {
	return &EmsFilterRuleCollectionGetParams{
		Context: ctx,
	}
}

// NewEmsFilterRuleCollectionGetParamsWithHTTPClient creates a new EmsFilterRuleCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewEmsFilterRuleCollectionGetParamsWithHTTPClient(client *http.Client) *EmsFilterRuleCollectionGetParams {
	return &EmsFilterRuleCollectionGetParams{
		HTTPClient: client,
	}
}

/* EmsFilterRuleCollectionGetParams contains all the parameters to send to the API endpoint
   for the ems filter rule collection get operation.

   Typically these are written to a http.Request.
*/
type EmsFilterRuleCollectionGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* Index.

	   Filter by index
	*/
	IndexQueryParameter *int64

	/* MaxRecords.

	   Limit the number of records returned.
	*/
	MaxRecords *int64

	/* MessageCriteriaNamePattern.

	   Filter by message_criteria.name_pattern
	*/
	MessageCriteriaNamePatternQueryParameter *string

	/* MessageCriteriaSeverities.

	   Filter by message_criteria.severities
	*/
	MessageCriteriaSeveritiesQueryParameter *string

	/* MessageCriteriaSnmpTrapTypes.

	   Filter by message_criteria.snmp_trap_types
	*/
	MessageCriteriaSnmpTrapTypesQueryParameter *string

	/* Name.

	   Filter Name
	*/
	NamePathParameter string

	/* OrderBy.

	   Order results by specified fields and optional [asc|desc] direction. Default direction is 'asc' for ascending.
	*/
	OrderBy []string

	/* ReturnRecords.

	   The default is true for GET calls.  When set to false, only the number of records is returned.

	   Default: true
	*/
	ReturnRecords *bool

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached.

	   Default: 15
	*/
	ReturnTimeout *int64

	/* Type.

	   Filter by type
	*/
	TypeQueryParameter *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ems filter rule collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EmsFilterRuleCollectionGetParams) WithDefaults() *EmsFilterRuleCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ems filter rule collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EmsFilterRuleCollectionGetParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)
	)

	val := EmsFilterRuleCollectionGetParams{
		ReturnRecords: &returnRecordsDefault,
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the ems filter rule collection get params
func (o *EmsFilterRuleCollectionGetParams) WithTimeout(timeout time.Duration) *EmsFilterRuleCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ems filter rule collection get params
func (o *EmsFilterRuleCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ems filter rule collection get params
func (o *EmsFilterRuleCollectionGetParams) WithContext(ctx context.Context) *EmsFilterRuleCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ems filter rule collection get params
func (o *EmsFilterRuleCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ems filter rule collection get params
func (o *EmsFilterRuleCollectionGetParams) WithHTTPClient(client *http.Client) *EmsFilterRuleCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ems filter rule collection get params
func (o *EmsFilterRuleCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the ems filter rule collection get params
func (o *EmsFilterRuleCollectionGetParams) WithFields(fields []string) *EmsFilterRuleCollectionGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the ems filter rule collection get params
func (o *EmsFilterRuleCollectionGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithIndexQueryParameter adds the index to the ems filter rule collection get params
func (o *EmsFilterRuleCollectionGetParams) WithIndexQueryParameter(index *int64) *EmsFilterRuleCollectionGetParams {
	o.SetIndexQueryParameter(index)
	return o
}

// SetIndexQueryParameter adds the index to the ems filter rule collection get params
func (o *EmsFilterRuleCollectionGetParams) SetIndexQueryParameter(index *int64) {
	o.IndexQueryParameter = index
}

// WithMaxRecords adds the maxRecords to the ems filter rule collection get params
func (o *EmsFilterRuleCollectionGetParams) WithMaxRecords(maxRecords *int64) *EmsFilterRuleCollectionGetParams {
	o.SetMaxRecords(maxRecords)
	return o
}

// SetMaxRecords adds the maxRecords to the ems filter rule collection get params
func (o *EmsFilterRuleCollectionGetParams) SetMaxRecords(maxRecords *int64) {
	o.MaxRecords = maxRecords
}

// WithMessageCriteriaNamePatternQueryParameter adds the messageCriteriaNamePattern to the ems filter rule collection get params
func (o *EmsFilterRuleCollectionGetParams) WithMessageCriteriaNamePatternQueryParameter(messageCriteriaNamePattern *string) *EmsFilterRuleCollectionGetParams {
	o.SetMessageCriteriaNamePatternQueryParameter(messageCriteriaNamePattern)
	return o
}

// SetMessageCriteriaNamePatternQueryParameter adds the messageCriteriaNamePattern to the ems filter rule collection get params
func (o *EmsFilterRuleCollectionGetParams) SetMessageCriteriaNamePatternQueryParameter(messageCriteriaNamePattern *string) {
	o.MessageCriteriaNamePatternQueryParameter = messageCriteriaNamePattern
}

// WithMessageCriteriaSeveritiesQueryParameter adds the messageCriteriaSeverities to the ems filter rule collection get params
func (o *EmsFilterRuleCollectionGetParams) WithMessageCriteriaSeveritiesQueryParameter(messageCriteriaSeverities *string) *EmsFilterRuleCollectionGetParams {
	o.SetMessageCriteriaSeveritiesQueryParameter(messageCriteriaSeverities)
	return o
}

// SetMessageCriteriaSeveritiesQueryParameter adds the messageCriteriaSeverities to the ems filter rule collection get params
func (o *EmsFilterRuleCollectionGetParams) SetMessageCriteriaSeveritiesQueryParameter(messageCriteriaSeverities *string) {
	o.MessageCriteriaSeveritiesQueryParameter = messageCriteriaSeverities
}

// WithMessageCriteriaSnmpTrapTypesQueryParameter adds the messageCriteriaSnmpTrapTypes to the ems filter rule collection get params
func (o *EmsFilterRuleCollectionGetParams) WithMessageCriteriaSnmpTrapTypesQueryParameter(messageCriteriaSnmpTrapTypes *string) *EmsFilterRuleCollectionGetParams {
	o.SetMessageCriteriaSnmpTrapTypesQueryParameter(messageCriteriaSnmpTrapTypes)
	return o
}

// SetMessageCriteriaSnmpTrapTypesQueryParameter adds the messageCriteriaSnmpTrapTypes to the ems filter rule collection get params
func (o *EmsFilterRuleCollectionGetParams) SetMessageCriteriaSnmpTrapTypesQueryParameter(messageCriteriaSnmpTrapTypes *string) {
	o.MessageCriteriaSnmpTrapTypesQueryParameter = messageCriteriaSnmpTrapTypes
}

// WithNamePathParameter adds the name to the ems filter rule collection get params
func (o *EmsFilterRuleCollectionGetParams) WithNamePathParameter(name string) *EmsFilterRuleCollectionGetParams {
	o.SetNamePathParameter(name)
	return o
}

// SetNamePathParameter adds the name to the ems filter rule collection get params
func (o *EmsFilterRuleCollectionGetParams) SetNamePathParameter(name string) {
	o.NamePathParameter = name
}

// WithOrderBy adds the orderBy to the ems filter rule collection get params
func (o *EmsFilterRuleCollectionGetParams) WithOrderBy(orderBy []string) *EmsFilterRuleCollectionGetParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the ems filter rule collection get params
func (o *EmsFilterRuleCollectionGetParams) SetOrderBy(orderBy []string) {
	o.OrderBy = orderBy
}

// WithReturnRecords adds the returnRecords to the ems filter rule collection get params
func (o *EmsFilterRuleCollectionGetParams) WithReturnRecords(returnRecords *bool) *EmsFilterRuleCollectionGetParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the ems filter rule collection get params
func (o *EmsFilterRuleCollectionGetParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the ems filter rule collection get params
func (o *EmsFilterRuleCollectionGetParams) WithReturnTimeout(returnTimeout *int64) *EmsFilterRuleCollectionGetParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the ems filter rule collection get params
func (o *EmsFilterRuleCollectionGetParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithTypeQueryParameter adds the typeVar to the ems filter rule collection get params
func (o *EmsFilterRuleCollectionGetParams) WithTypeQueryParameter(typeVar *string) *EmsFilterRuleCollectionGetParams {
	o.SetTypeQueryParameter(typeVar)
	return o
}

// SetTypeQueryParameter adds the type to the ems filter rule collection get params
func (o *EmsFilterRuleCollectionGetParams) SetTypeQueryParameter(typeVar *string) {
	o.TypeQueryParameter = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *EmsFilterRuleCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Fields != nil {

		// binding items for fields
		joinedFields := o.bindParamFields(reg)

		// query array param fields
		if err := r.SetQueryParam("fields", joinedFields...); err != nil {
			return err
		}
	}

	if o.IndexQueryParameter != nil {

		// query param index
		var qrIndex int64

		if o.IndexQueryParameter != nil {
			qrIndex = *o.IndexQueryParameter
		}
		qIndex := swag.FormatInt64(qrIndex)
		if qIndex != "" {

			if err := r.SetQueryParam("index", qIndex); err != nil {
				return err
			}
		}
	}

	if o.MaxRecords != nil {

		// query param max_records
		var qrMaxRecords int64

		if o.MaxRecords != nil {
			qrMaxRecords = *o.MaxRecords
		}
		qMaxRecords := swag.FormatInt64(qrMaxRecords)
		if qMaxRecords != "" {

			if err := r.SetQueryParam("max_records", qMaxRecords); err != nil {
				return err
			}
		}
	}

	if o.MessageCriteriaNamePatternQueryParameter != nil {

		// query param message_criteria.name_pattern
		var qrMessageCriteriaNamePattern string

		if o.MessageCriteriaNamePatternQueryParameter != nil {
			qrMessageCriteriaNamePattern = *o.MessageCriteriaNamePatternQueryParameter
		}
		qMessageCriteriaNamePattern := qrMessageCriteriaNamePattern
		if qMessageCriteriaNamePattern != "" {

			if err := r.SetQueryParam("message_criteria.name_pattern", qMessageCriteriaNamePattern); err != nil {
				return err
			}
		}
	}

	if o.MessageCriteriaSeveritiesQueryParameter != nil {

		// query param message_criteria.severities
		var qrMessageCriteriaSeverities string

		if o.MessageCriteriaSeveritiesQueryParameter != nil {
			qrMessageCriteriaSeverities = *o.MessageCriteriaSeveritiesQueryParameter
		}
		qMessageCriteriaSeverities := qrMessageCriteriaSeverities
		if qMessageCriteriaSeverities != "" {

			if err := r.SetQueryParam("message_criteria.severities", qMessageCriteriaSeverities); err != nil {
				return err
			}
		}
	}

	if o.MessageCriteriaSnmpTrapTypesQueryParameter != nil {

		// query param message_criteria.snmp_trap_types
		var qrMessageCriteriaSnmpTrapTypes string

		if o.MessageCriteriaSnmpTrapTypesQueryParameter != nil {
			qrMessageCriteriaSnmpTrapTypes = *o.MessageCriteriaSnmpTrapTypesQueryParameter
		}
		qMessageCriteriaSnmpTrapTypes := qrMessageCriteriaSnmpTrapTypes
		if qMessageCriteriaSnmpTrapTypes != "" {

			if err := r.SetQueryParam("message_criteria.snmp_trap_types", qMessageCriteriaSnmpTrapTypes); err != nil {
				return err
			}
		}
	}

	// path param name
	if err := r.SetPathParam("name", o.NamePathParameter); err != nil {
		return err
	}

	if o.OrderBy != nil {

		// binding items for order_by
		joinedOrderBy := o.bindParamOrderBy(reg)

		// query array param order_by
		if err := r.SetQueryParam("order_by", joinedOrderBy...); err != nil {
			return err
		}
	}

	if o.ReturnRecords != nil {

		// query param return_records
		var qrReturnRecords bool

		if o.ReturnRecords != nil {
			qrReturnRecords = *o.ReturnRecords
		}
		qReturnRecords := swag.FormatBool(qrReturnRecords)
		if qReturnRecords != "" {

			if err := r.SetQueryParam("return_records", qReturnRecords); err != nil {
				return err
			}
		}
	}

	if o.ReturnTimeout != nil {

		// query param return_timeout
		var qrReturnTimeout int64

		if o.ReturnTimeout != nil {
			qrReturnTimeout = *o.ReturnTimeout
		}
		qReturnTimeout := swag.FormatInt64(qrReturnTimeout)
		if qReturnTimeout != "" {

			if err := r.SetQueryParam("return_timeout", qReturnTimeout); err != nil {
				return err
			}
		}
	}

	if o.TypeQueryParameter != nil {

		// query param type
		var qrType string

		if o.TypeQueryParameter != nil {
			qrType = *o.TypeQueryParameter
		}
		qType := qrType
		if qType != "" {

			if err := r.SetQueryParam("type", qType); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamEmsFilterRuleCollectionGet binds the parameter fields
func (o *EmsFilterRuleCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
	fieldsIR := o.Fields

	var fieldsIC []string
	for _, fieldsIIR := range fieldsIR { // explode []string

		fieldsIIV := fieldsIIR // string as string
		fieldsIC = append(fieldsIC, fieldsIIV)
	}

	// items.CollectionFormat: "csv"
	fieldsIS := swag.JoinByFormat(fieldsIC, "csv")

	return fieldsIS
}

// bindParamEmsFilterRuleCollectionGet binds the parameter order_by
func (o *EmsFilterRuleCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
	orderByIR := o.OrderBy

	var orderByIC []string
	for _, orderByIIR := range orderByIR { // explode []string

		orderByIIV := orderByIIR // string as string
		orderByIC = append(orderByIC, orderByIIV)
	}

	// items.CollectionFormat: "csv"
	orderByIS := swag.JoinByFormat(orderByIC, "csv")

	return orderByIS
}
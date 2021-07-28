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

// NewLunCollectionPerformanceMetricsGetParams creates a new LunCollectionPerformanceMetricsGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLunCollectionPerformanceMetricsGetParams() *LunCollectionPerformanceMetricsGetParams {
	return &LunCollectionPerformanceMetricsGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLunCollectionPerformanceMetricsGetParamsWithTimeout creates a new LunCollectionPerformanceMetricsGetParams object
// with the ability to set a timeout on a request.
func NewLunCollectionPerformanceMetricsGetParamsWithTimeout(timeout time.Duration) *LunCollectionPerformanceMetricsGetParams {
	return &LunCollectionPerformanceMetricsGetParams{
		timeout: timeout,
	}
}

// NewLunCollectionPerformanceMetricsGetParamsWithContext creates a new LunCollectionPerformanceMetricsGetParams object
// with the ability to set a context for a request.
func NewLunCollectionPerformanceMetricsGetParamsWithContext(ctx context.Context) *LunCollectionPerformanceMetricsGetParams {
	return &LunCollectionPerformanceMetricsGetParams{
		Context: ctx,
	}
}

// NewLunCollectionPerformanceMetricsGetParamsWithHTTPClient creates a new LunCollectionPerformanceMetricsGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewLunCollectionPerformanceMetricsGetParamsWithHTTPClient(client *http.Client) *LunCollectionPerformanceMetricsGetParams {
	return &LunCollectionPerformanceMetricsGetParams{
		HTTPClient: client,
	}
}

/* LunCollectionPerformanceMetricsGetParams contains all the parameters to send to the API endpoint
   for the lun collection performance metrics get operation.

   Typically these are written to a http.Request.
*/
type LunCollectionPerformanceMetricsGetParams struct {

	/* Duration.

	   Filter by duration
	*/
	DurationQueryParameter *string

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* Interval.

	     The time range for the data. Examples can be 1h, 1d, 1m, 1w, 1y.
	The period for each time range is as follows:
	* 1h: Metrics over the most recent hour sampled over 15 seconds.
	* 1d: Metrics over the most recent day sampled over 5 minutes.
	* 1w: Metrics over the most recent week sampled over 30 minutes.
	* 1m: Metrics over the most recent month sampled over 2 hours.
	* 1y: Metrics over the most recent year sampled over a day.


	     Default: "1h"
	*/
	IntervalQueryParameter *string

	/* IopsOther.

	   Filter by iops.other
	*/
	IopsOtherQueryParameter *int64

	/* IopsRead.

	   Filter by iops.read
	*/
	IopsReadQueryParameter *int64

	/* IopsTotal.

	   Filter by iops.total
	*/
	IopsTotalQueryParameter *int64

	/* IopsWrite.

	   Filter by iops.write
	*/
	IopsWriteQueryParameter *int64

	/* LatencyOther.

	   Filter by latency.other
	*/
	LatencyOtherQueryParameter *int64

	/* LatencyRead.

	   Filter by latency.read
	*/
	LatencyReadQueryParameter *int64

	/* LatencyTotal.

	   Filter by latency.total
	*/
	LatencyTotalQueryParameter *int64

	/* LatencyWrite.

	   Filter by latency.write
	*/
	LatencyWriteQueryParameter *int64

	/* MaxRecords.

	   Limit the number of records returned.
	*/
	MaxRecords *int64

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

	/* Status.

	   Filter by status
	*/
	StatusQueryParameter *string

	/* ThroughputOther.

	   Filter by throughput.other
	*/
	ThroughputOtherQueryParameter *int64

	/* ThroughputRead.

	   Filter by throughput.read
	*/
	ThroughputReadQueryParameter *int64

	/* ThroughputTotal.

	   Filter by throughput.total
	*/
	ThroughputTotalQueryParameter *int64

	/* ThroughputWrite.

	   Filter by throughput.write
	*/
	ThroughputWriteQueryParameter *int64

	/* Timestamp.

	   Filter by timestamp
	*/
	TimestampQueryParameter *string

	/* UUID.

	   Unique identifier of the LUN.
	*/
	UUIDPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the lun collection performance metrics get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LunCollectionPerformanceMetricsGetParams) WithDefaults() *LunCollectionPerformanceMetricsGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the lun collection performance metrics get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LunCollectionPerformanceMetricsGetParams) SetDefaults() {
	var (
		intervalQueryParameterDefault = string("1h")

		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)
	)

	val := LunCollectionPerformanceMetricsGetParams{
		IntervalQueryParameter: &intervalQueryParameterDefault,
		ReturnRecords:          &returnRecordsDefault,
		ReturnTimeout:          &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the lun collection performance metrics get params
func (o *LunCollectionPerformanceMetricsGetParams) WithTimeout(timeout time.Duration) *LunCollectionPerformanceMetricsGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the lun collection performance metrics get params
func (o *LunCollectionPerformanceMetricsGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the lun collection performance metrics get params
func (o *LunCollectionPerformanceMetricsGetParams) WithContext(ctx context.Context) *LunCollectionPerformanceMetricsGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the lun collection performance metrics get params
func (o *LunCollectionPerformanceMetricsGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the lun collection performance metrics get params
func (o *LunCollectionPerformanceMetricsGetParams) WithHTTPClient(client *http.Client) *LunCollectionPerformanceMetricsGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the lun collection performance metrics get params
func (o *LunCollectionPerformanceMetricsGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDurationQueryParameter adds the duration to the lun collection performance metrics get params
func (o *LunCollectionPerformanceMetricsGetParams) WithDurationQueryParameter(duration *string) *LunCollectionPerformanceMetricsGetParams {
	o.SetDurationQueryParameter(duration)
	return o
}

// SetDurationQueryParameter adds the duration to the lun collection performance metrics get params
func (o *LunCollectionPerformanceMetricsGetParams) SetDurationQueryParameter(duration *string) {
	o.DurationQueryParameter = duration
}

// WithFields adds the fields to the lun collection performance metrics get params
func (o *LunCollectionPerformanceMetricsGetParams) WithFields(fields []string) *LunCollectionPerformanceMetricsGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the lun collection performance metrics get params
func (o *LunCollectionPerformanceMetricsGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithIntervalQueryParameter adds the interval to the lun collection performance metrics get params
func (o *LunCollectionPerformanceMetricsGetParams) WithIntervalQueryParameter(interval *string) *LunCollectionPerformanceMetricsGetParams {
	o.SetIntervalQueryParameter(interval)
	return o
}

// SetIntervalQueryParameter adds the interval to the lun collection performance metrics get params
func (o *LunCollectionPerformanceMetricsGetParams) SetIntervalQueryParameter(interval *string) {
	o.IntervalQueryParameter = interval
}

// WithIopsOtherQueryParameter adds the iopsOther to the lun collection performance metrics get params
func (o *LunCollectionPerformanceMetricsGetParams) WithIopsOtherQueryParameter(iopsOther *int64) *LunCollectionPerformanceMetricsGetParams {
	o.SetIopsOtherQueryParameter(iopsOther)
	return o
}

// SetIopsOtherQueryParameter adds the iopsOther to the lun collection performance metrics get params
func (o *LunCollectionPerformanceMetricsGetParams) SetIopsOtherQueryParameter(iopsOther *int64) {
	o.IopsOtherQueryParameter = iopsOther
}

// WithIopsReadQueryParameter adds the iopsRead to the lun collection performance metrics get params
func (o *LunCollectionPerformanceMetricsGetParams) WithIopsReadQueryParameter(iopsRead *int64) *LunCollectionPerformanceMetricsGetParams {
	o.SetIopsReadQueryParameter(iopsRead)
	return o
}

// SetIopsReadQueryParameter adds the iopsRead to the lun collection performance metrics get params
func (o *LunCollectionPerformanceMetricsGetParams) SetIopsReadQueryParameter(iopsRead *int64) {
	o.IopsReadQueryParameter = iopsRead
}

// WithIopsTotalQueryParameter adds the iopsTotal to the lun collection performance metrics get params
func (o *LunCollectionPerformanceMetricsGetParams) WithIopsTotalQueryParameter(iopsTotal *int64) *LunCollectionPerformanceMetricsGetParams {
	o.SetIopsTotalQueryParameter(iopsTotal)
	return o
}

// SetIopsTotalQueryParameter adds the iopsTotal to the lun collection performance metrics get params
func (o *LunCollectionPerformanceMetricsGetParams) SetIopsTotalQueryParameter(iopsTotal *int64) {
	o.IopsTotalQueryParameter = iopsTotal
}

// WithIopsWriteQueryParameter adds the iopsWrite to the lun collection performance metrics get params
func (o *LunCollectionPerformanceMetricsGetParams) WithIopsWriteQueryParameter(iopsWrite *int64) *LunCollectionPerformanceMetricsGetParams {
	o.SetIopsWriteQueryParameter(iopsWrite)
	return o
}

// SetIopsWriteQueryParameter adds the iopsWrite to the lun collection performance metrics get params
func (o *LunCollectionPerformanceMetricsGetParams) SetIopsWriteQueryParameter(iopsWrite *int64) {
	o.IopsWriteQueryParameter = iopsWrite
}

// WithLatencyOtherQueryParameter adds the latencyOther to the lun collection performance metrics get params
func (o *LunCollectionPerformanceMetricsGetParams) WithLatencyOtherQueryParameter(latencyOther *int64) *LunCollectionPerformanceMetricsGetParams {
	o.SetLatencyOtherQueryParameter(latencyOther)
	return o
}

// SetLatencyOtherQueryParameter adds the latencyOther to the lun collection performance metrics get params
func (o *LunCollectionPerformanceMetricsGetParams) SetLatencyOtherQueryParameter(latencyOther *int64) {
	o.LatencyOtherQueryParameter = latencyOther
}

// WithLatencyReadQueryParameter adds the latencyRead to the lun collection performance metrics get params
func (o *LunCollectionPerformanceMetricsGetParams) WithLatencyReadQueryParameter(latencyRead *int64) *LunCollectionPerformanceMetricsGetParams {
	o.SetLatencyReadQueryParameter(latencyRead)
	return o
}

// SetLatencyReadQueryParameter adds the latencyRead to the lun collection performance metrics get params
func (o *LunCollectionPerformanceMetricsGetParams) SetLatencyReadQueryParameter(latencyRead *int64) {
	o.LatencyReadQueryParameter = latencyRead
}

// WithLatencyTotalQueryParameter adds the latencyTotal to the lun collection performance metrics get params
func (o *LunCollectionPerformanceMetricsGetParams) WithLatencyTotalQueryParameter(latencyTotal *int64) *LunCollectionPerformanceMetricsGetParams {
	o.SetLatencyTotalQueryParameter(latencyTotal)
	return o
}

// SetLatencyTotalQueryParameter adds the latencyTotal to the lun collection performance metrics get params
func (o *LunCollectionPerformanceMetricsGetParams) SetLatencyTotalQueryParameter(latencyTotal *int64) {
	o.LatencyTotalQueryParameter = latencyTotal
}

// WithLatencyWriteQueryParameter adds the latencyWrite to the lun collection performance metrics get params
func (o *LunCollectionPerformanceMetricsGetParams) WithLatencyWriteQueryParameter(latencyWrite *int64) *LunCollectionPerformanceMetricsGetParams {
	o.SetLatencyWriteQueryParameter(latencyWrite)
	return o
}

// SetLatencyWriteQueryParameter adds the latencyWrite to the lun collection performance metrics get params
func (o *LunCollectionPerformanceMetricsGetParams) SetLatencyWriteQueryParameter(latencyWrite *int64) {
	o.LatencyWriteQueryParameter = latencyWrite
}

// WithMaxRecords adds the maxRecords to the lun collection performance metrics get params
func (o *LunCollectionPerformanceMetricsGetParams) WithMaxRecords(maxRecords *int64) *LunCollectionPerformanceMetricsGetParams {
	o.SetMaxRecords(maxRecords)
	return o
}

// SetMaxRecords adds the maxRecords to the lun collection performance metrics get params
func (o *LunCollectionPerformanceMetricsGetParams) SetMaxRecords(maxRecords *int64) {
	o.MaxRecords = maxRecords
}

// WithOrderBy adds the orderBy to the lun collection performance metrics get params
func (o *LunCollectionPerformanceMetricsGetParams) WithOrderBy(orderBy []string) *LunCollectionPerformanceMetricsGetParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the lun collection performance metrics get params
func (o *LunCollectionPerformanceMetricsGetParams) SetOrderBy(orderBy []string) {
	o.OrderBy = orderBy
}

// WithReturnRecords adds the returnRecords to the lun collection performance metrics get params
func (o *LunCollectionPerformanceMetricsGetParams) WithReturnRecords(returnRecords *bool) *LunCollectionPerformanceMetricsGetParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the lun collection performance metrics get params
func (o *LunCollectionPerformanceMetricsGetParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the lun collection performance metrics get params
func (o *LunCollectionPerformanceMetricsGetParams) WithReturnTimeout(returnTimeout *int64) *LunCollectionPerformanceMetricsGetParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the lun collection performance metrics get params
func (o *LunCollectionPerformanceMetricsGetParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithStatusQueryParameter adds the status to the lun collection performance metrics get params
func (o *LunCollectionPerformanceMetricsGetParams) WithStatusQueryParameter(status *string) *LunCollectionPerformanceMetricsGetParams {
	o.SetStatusQueryParameter(status)
	return o
}

// SetStatusQueryParameter adds the status to the lun collection performance metrics get params
func (o *LunCollectionPerformanceMetricsGetParams) SetStatusQueryParameter(status *string) {
	o.StatusQueryParameter = status
}

// WithThroughputOtherQueryParameter adds the throughputOther to the lun collection performance metrics get params
func (o *LunCollectionPerformanceMetricsGetParams) WithThroughputOtherQueryParameter(throughputOther *int64) *LunCollectionPerformanceMetricsGetParams {
	o.SetThroughputOtherQueryParameter(throughputOther)
	return o
}

// SetThroughputOtherQueryParameter adds the throughputOther to the lun collection performance metrics get params
func (o *LunCollectionPerformanceMetricsGetParams) SetThroughputOtherQueryParameter(throughputOther *int64) {
	o.ThroughputOtherQueryParameter = throughputOther
}

// WithThroughputReadQueryParameter adds the throughputRead to the lun collection performance metrics get params
func (o *LunCollectionPerformanceMetricsGetParams) WithThroughputReadQueryParameter(throughputRead *int64) *LunCollectionPerformanceMetricsGetParams {
	o.SetThroughputReadQueryParameter(throughputRead)
	return o
}

// SetThroughputReadQueryParameter adds the throughputRead to the lun collection performance metrics get params
func (o *LunCollectionPerformanceMetricsGetParams) SetThroughputReadQueryParameter(throughputRead *int64) {
	o.ThroughputReadQueryParameter = throughputRead
}

// WithThroughputTotalQueryParameter adds the throughputTotal to the lun collection performance metrics get params
func (o *LunCollectionPerformanceMetricsGetParams) WithThroughputTotalQueryParameter(throughputTotal *int64) *LunCollectionPerformanceMetricsGetParams {
	o.SetThroughputTotalQueryParameter(throughputTotal)
	return o
}

// SetThroughputTotalQueryParameter adds the throughputTotal to the lun collection performance metrics get params
func (o *LunCollectionPerformanceMetricsGetParams) SetThroughputTotalQueryParameter(throughputTotal *int64) {
	o.ThroughputTotalQueryParameter = throughputTotal
}

// WithThroughputWriteQueryParameter adds the throughputWrite to the lun collection performance metrics get params
func (o *LunCollectionPerformanceMetricsGetParams) WithThroughputWriteQueryParameter(throughputWrite *int64) *LunCollectionPerformanceMetricsGetParams {
	o.SetThroughputWriteQueryParameter(throughputWrite)
	return o
}

// SetThroughputWriteQueryParameter adds the throughputWrite to the lun collection performance metrics get params
func (o *LunCollectionPerformanceMetricsGetParams) SetThroughputWriteQueryParameter(throughputWrite *int64) {
	o.ThroughputWriteQueryParameter = throughputWrite
}

// WithTimestampQueryParameter adds the timestamp to the lun collection performance metrics get params
func (o *LunCollectionPerformanceMetricsGetParams) WithTimestampQueryParameter(timestamp *string) *LunCollectionPerformanceMetricsGetParams {
	o.SetTimestampQueryParameter(timestamp)
	return o
}

// SetTimestampQueryParameter adds the timestamp to the lun collection performance metrics get params
func (o *LunCollectionPerformanceMetricsGetParams) SetTimestampQueryParameter(timestamp *string) {
	o.TimestampQueryParameter = timestamp
}

// WithUUIDPathParameter adds the uuid to the lun collection performance metrics get params
func (o *LunCollectionPerformanceMetricsGetParams) WithUUIDPathParameter(uuid string) *LunCollectionPerformanceMetricsGetParams {
	o.SetUUIDPathParameter(uuid)
	return o
}

// SetUUIDPathParameter adds the uuid to the lun collection performance metrics get params
func (o *LunCollectionPerformanceMetricsGetParams) SetUUIDPathParameter(uuid string) {
	o.UUIDPathParameter = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *LunCollectionPerformanceMetricsGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DurationQueryParameter != nil {

		// query param duration
		var qrDuration string

		if o.DurationQueryParameter != nil {
			qrDuration = *o.DurationQueryParameter
		}
		qDuration := qrDuration
		if qDuration != "" {

			if err := r.SetQueryParam("duration", qDuration); err != nil {
				return err
			}
		}
	}

	if o.Fields != nil {

		// binding items for fields
		joinedFields := o.bindParamFields(reg)

		// query array param fields
		if err := r.SetQueryParam("fields", joinedFields...); err != nil {
			return err
		}
	}

	if o.IntervalQueryParameter != nil {

		// query param interval
		var qrInterval string

		if o.IntervalQueryParameter != nil {
			qrInterval = *o.IntervalQueryParameter
		}
		qInterval := qrInterval
		if qInterval != "" {

			if err := r.SetQueryParam("interval", qInterval); err != nil {
				return err
			}
		}
	}

	if o.IopsOtherQueryParameter != nil {

		// query param iops.other
		var qrIopsOther int64

		if o.IopsOtherQueryParameter != nil {
			qrIopsOther = *o.IopsOtherQueryParameter
		}
		qIopsOther := swag.FormatInt64(qrIopsOther)
		if qIopsOther != "" {

			if err := r.SetQueryParam("iops.other", qIopsOther); err != nil {
				return err
			}
		}
	}

	if o.IopsReadQueryParameter != nil {

		// query param iops.read
		var qrIopsRead int64

		if o.IopsReadQueryParameter != nil {
			qrIopsRead = *o.IopsReadQueryParameter
		}
		qIopsRead := swag.FormatInt64(qrIopsRead)
		if qIopsRead != "" {

			if err := r.SetQueryParam("iops.read", qIopsRead); err != nil {
				return err
			}
		}
	}

	if o.IopsTotalQueryParameter != nil {

		// query param iops.total
		var qrIopsTotal int64

		if o.IopsTotalQueryParameter != nil {
			qrIopsTotal = *o.IopsTotalQueryParameter
		}
		qIopsTotal := swag.FormatInt64(qrIopsTotal)
		if qIopsTotal != "" {

			if err := r.SetQueryParam("iops.total", qIopsTotal); err != nil {
				return err
			}
		}
	}

	if o.IopsWriteQueryParameter != nil {

		// query param iops.write
		var qrIopsWrite int64

		if o.IopsWriteQueryParameter != nil {
			qrIopsWrite = *o.IopsWriteQueryParameter
		}
		qIopsWrite := swag.FormatInt64(qrIopsWrite)
		if qIopsWrite != "" {

			if err := r.SetQueryParam("iops.write", qIopsWrite); err != nil {
				return err
			}
		}
	}

	if o.LatencyOtherQueryParameter != nil {

		// query param latency.other
		var qrLatencyOther int64

		if o.LatencyOtherQueryParameter != nil {
			qrLatencyOther = *o.LatencyOtherQueryParameter
		}
		qLatencyOther := swag.FormatInt64(qrLatencyOther)
		if qLatencyOther != "" {

			if err := r.SetQueryParam("latency.other", qLatencyOther); err != nil {
				return err
			}
		}
	}

	if o.LatencyReadQueryParameter != nil {

		// query param latency.read
		var qrLatencyRead int64

		if o.LatencyReadQueryParameter != nil {
			qrLatencyRead = *o.LatencyReadQueryParameter
		}
		qLatencyRead := swag.FormatInt64(qrLatencyRead)
		if qLatencyRead != "" {

			if err := r.SetQueryParam("latency.read", qLatencyRead); err != nil {
				return err
			}
		}
	}

	if o.LatencyTotalQueryParameter != nil {

		// query param latency.total
		var qrLatencyTotal int64

		if o.LatencyTotalQueryParameter != nil {
			qrLatencyTotal = *o.LatencyTotalQueryParameter
		}
		qLatencyTotal := swag.FormatInt64(qrLatencyTotal)
		if qLatencyTotal != "" {

			if err := r.SetQueryParam("latency.total", qLatencyTotal); err != nil {
				return err
			}
		}
	}

	if o.LatencyWriteQueryParameter != nil {

		// query param latency.write
		var qrLatencyWrite int64

		if o.LatencyWriteQueryParameter != nil {
			qrLatencyWrite = *o.LatencyWriteQueryParameter
		}
		qLatencyWrite := swag.FormatInt64(qrLatencyWrite)
		if qLatencyWrite != "" {

			if err := r.SetQueryParam("latency.write", qLatencyWrite); err != nil {
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

	if o.StatusQueryParameter != nil {

		// query param status
		var qrStatus string

		if o.StatusQueryParameter != nil {
			qrStatus = *o.StatusQueryParameter
		}
		qStatus := qrStatus
		if qStatus != "" {

			if err := r.SetQueryParam("status", qStatus); err != nil {
				return err
			}
		}
	}

	if o.ThroughputOtherQueryParameter != nil {

		// query param throughput.other
		var qrThroughputOther int64

		if o.ThroughputOtherQueryParameter != nil {
			qrThroughputOther = *o.ThroughputOtherQueryParameter
		}
		qThroughputOther := swag.FormatInt64(qrThroughputOther)
		if qThroughputOther != "" {

			if err := r.SetQueryParam("throughput.other", qThroughputOther); err != nil {
				return err
			}
		}
	}

	if o.ThroughputReadQueryParameter != nil {

		// query param throughput.read
		var qrThroughputRead int64

		if o.ThroughputReadQueryParameter != nil {
			qrThroughputRead = *o.ThroughputReadQueryParameter
		}
		qThroughputRead := swag.FormatInt64(qrThroughputRead)
		if qThroughputRead != "" {

			if err := r.SetQueryParam("throughput.read", qThroughputRead); err != nil {
				return err
			}
		}
	}

	if o.ThroughputTotalQueryParameter != nil {

		// query param throughput.total
		var qrThroughputTotal int64

		if o.ThroughputTotalQueryParameter != nil {
			qrThroughputTotal = *o.ThroughputTotalQueryParameter
		}
		qThroughputTotal := swag.FormatInt64(qrThroughputTotal)
		if qThroughputTotal != "" {

			if err := r.SetQueryParam("throughput.total", qThroughputTotal); err != nil {
				return err
			}
		}
	}

	if o.ThroughputWriteQueryParameter != nil {

		// query param throughput.write
		var qrThroughputWrite int64

		if o.ThroughputWriteQueryParameter != nil {
			qrThroughputWrite = *o.ThroughputWriteQueryParameter
		}
		qThroughputWrite := swag.FormatInt64(qrThroughputWrite)
		if qThroughputWrite != "" {

			if err := r.SetQueryParam("throughput.write", qThroughputWrite); err != nil {
				return err
			}
		}
	}

	if o.TimestampQueryParameter != nil {

		// query param timestamp
		var qrTimestamp string

		if o.TimestampQueryParameter != nil {
			qrTimestamp = *o.TimestampQueryParameter
		}
		qTimestamp := qrTimestamp
		if qTimestamp != "" {

			if err := r.SetQueryParam("timestamp", qTimestamp); err != nil {
				return err
			}
		}
	}

	// path param uuid
	if err := r.SetPathParam("uuid", o.UUIDPathParameter); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamLunCollectionPerformanceMetricsGet binds the parameter fields
func (o *LunCollectionPerformanceMetricsGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamLunCollectionPerformanceMetricsGet binds the parameter order_by
func (o *LunCollectionPerformanceMetricsGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
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
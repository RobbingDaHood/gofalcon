// Code generated by go-swagger; DO NOT EDIT.

package container_images

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

// NewGetCombinedImagesParams creates a new GetCombinedImagesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCombinedImagesParams() *GetCombinedImagesParams {
	return &GetCombinedImagesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCombinedImagesParamsWithTimeout creates a new GetCombinedImagesParams object
// with the ability to set a timeout on a request.
func NewGetCombinedImagesParamsWithTimeout(timeout time.Duration) *GetCombinedImagesParams {
	return &GetCombinedImagesParams{
		timeout: timeout,
	}
}

// NewGetCombinedImagesParamsWithContext creates a new GetCombinedImagesParams object
// with the ability to set a context for a request.
func NewGetCombinedImagesParamsWithContext(ctx context.Context) *GetCombinedImagesParams {
	return &GetCombinedImagesParams{
		Context: ctx,
	}
}

// NewGetCombinedImagesParamsWithHTTPClient creates a new GetCombinedImagesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCombinedImagesParamsWithHTTPClient(client *http.Client) *GetCombinedImagesParams {
	return &GetCombinedImagesParams{
		HTTPClient: client,
	}
}

/*
GetCombinedImagesParams contains all the parameters to send to the API endpoint

	for the get combined images operation.

	Typically these are written to a http.Request.
*/
type GetCombinedImagesParams struct {

	/* Filter.

	   Filter images using a query in Falcon Query Language (FQL). Supported filters:  container_id, container_running_status, cve_id, detection_name, detection_severity, first_seen, image_digest, image_id, registry, repository, tag, vulnerability_severity
	*/
	Filter *string

	/* Limit.

	   The upper-bound on the number of records to retrieve [1-100]
	*/
	Limit *int64

	/* Offset.

	   The offset from where to begin.
	*/
	Offset *int64

	/* Sort.

	   The fields to sort the records on. Supported columns:  [first_seen highest_detection_severity highest_vulnerability_severity image_digest image_id registry repository source tag]
	*/
	Sort *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get combined images params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCombinedImagesParams) WithDefaults() *GetCombinedImagesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get combined images params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCombinedImagesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get combined images params
func (o *GetCombinedImagesParams) WithTimeout(timeout time.Duration) *GetCombinedImagesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get combined images params
func (o *GetCombinedImagesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get combined images params
func (o *GetCombinedImagesParams) WithContext(ctx context.Context) *GetCombinedImagesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get combined images params
func (o *GetCombinedImagesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get combined images params
func (o *GetCombinedImagesParams) WithHTTPClient(client *http.Client) *GetCombinedImagesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get combined images params
func (o *GetCombinedImagesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilter adds the filter to the get combined images params
func (o *GetCombinedImagesParams) WithFilter(filter *string) *GetCombinedImagesParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the get combined images params
func (o *GetCombinedImagesParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithLimit adds the limit to the get combined images params
func (o *GetCombinedImagesParams) WithLimit(limit *int64) *GetCombinedImagesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get combined images params
func (o *GetCombinedImagesParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the get combined images params
func (o *GetCombinedImagesParams) WithOffset(offset *int64) *GetCombinedImagesParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get combined images params
func (o *GetCombinedImagesParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithSort adds the sort to the get combined images params
func (o *GetCombinedImagesParams) WithSort(sort *string) *GetCombinedImagesParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the get combined images params
func (o *GetCombinedImagesParams) SetSort(sort *string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *GetCombinedImagesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Filter != nil {

		// query param filter
		var qrFilter string

		if o.Filter != nil {
			qrFilter = *o.Filter
		}
		qFilter := qrFilter
		if qFilter != "" {

			if err := r.SetQueryParam("filter", qFilter); err != nil {
				return err
			}
		}
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {

			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}
	}

	if o.Sort != nil {

		// query param sort
		var qrSort string

		if o.Sort != nil {
			qrSort = *o.Sort
		}
		qSort := qrSort
		if qSort != "" {

			if err := r.SetQueryParam("sort", qSort); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

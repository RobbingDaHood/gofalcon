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

// NewDeleteBaseImagesParams creates a new DeleteBaseImagesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteBaseImagesParams() *DeleteBaseImagesParams {
	return &DeleteBaseImagesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteBaseImagesParamsWithTimeout creates a new DeleteBaseImagesParams object
// with the ability to set a timeout on a request.
func NewDeleteBaseImagesParamsWithTimeout(timeout time.Duration) *DeleteBaseImagesParams {
	return &DeleteBaseImagesParams{
		timeout: timeout,
	}
}

// NewDeleteBaseImagesParamsWithContext creates a new DeleteBaseImagesParams object
// with the ability to set a context for a request.
func NewDeleteBaseImagesParamsWithContext(ctx context.Context) *DeleteBaseImagesParams {
	return &DeleteBaseImagesParams{
		Context: ctx,
	}
}

// NewDeleteBaseImagesParamsWithHTTPClient creates a new DeleteBaseImagesParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteBaseImagesParamsWithHTTPClient(client *http.Client) *DeleteBaseImagesParams {
	return &DeleteBaseImagesParams{
		HTTPClient: client,
	}
}

/*
DeleteBaseImagesParams contains all the parameters to send to the API endpoint

	for the delete base images operation.

	Typically these are written to a http.Request.
*/
type DeleteBaseImagesParams struct {

	/* Ids.

	   BaseImageIDs
	*/
	Ids []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete base images params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteBaseImagesParams) WithDefaults() *DeleteBaseImagesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete base images params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteBaseImagesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete base images params
func (o *DeleteBaseImagesParams) WithTimeout(timeout time.Duration) *DeleteBaseImagesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete base images params
func (o *DeleteBaseImagesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete base images params
func (o *DeleteBaseImagesParams) WithContext(ctx context.Context) *DeleteBaseImagesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete base images params
func (o *DeleteBaseImagesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete base images params
func (o *DeleteBaseImagesParams) WithHTTPClient(client *http.Client) *DeleteBaseImagesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete base images params
func (o *DeleteBaseImagesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIds adds the ids to the delete base images params
func (o *DeleteBaseImagesParams) WithIds(ids []string) *DeleteBaseImagesParams {
	o.SetIds(ids)
	return o
}

// SetIds adds the ids to the delete base images params
func (o *DeleteBaseImagesParams) SetIds(ids []string) {
	o.Ids = ids
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteBaseImagesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Ids != nil {

		// binding items for ids
		joinedIds := o.bindParamIds(reg)

		// query array param ids
		if err := r.SetQueryParam("ids", joinedIds...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamDeleteBaseImages binds the parameter ids
func (o *DeleteBaseImagesParams) bindParamIds(formats strfmt.Registry) []string {
	idsIR := o.Ids

	var idsIC []string
	for _, idsIIR := range idsIR { // explode []string

		idsIIV := idsIIR // string as string
		idsIC = append(idsIC, idsIIV)
	}

	// items.CollectionFormat: "csv"
	idsIS := swag.JoinByFormat(idsIC, "csv")

	return idsIS
}

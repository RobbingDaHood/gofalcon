// Code generated by go-swagger; DO NOT EDIT.

package certificate_based_exclusions

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

	"github.com/crowdstrike/gofalcon/falcon/models"
)

// NewCbExclusionsCreateV1Params creates a new CbExclusionsCreateV1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCbExclusionsCreateV1Params() *CbExclusionsCreateV1Params {
	return &CbExclusionsCreateV1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewCbExclusionsCreateV1ParamsWithTimeout creates a new CbExclusionsCreateV1Params object
// with the ability to set a timeout on a request.
func NewCbExclusionsCreateV1ParamsWithTimeout(timeout time.Duration) *CbExclusionsCreateV1Params {
	return &CbExclusionsCreateV1Params{
		timeout: timeout,
	}
}

// NewCbExclusionsCreateV1ParamsWithContext creates a new CbExclusionsCreateV1Params object
// with the ability to set a context for a request.
func NewCbExclusionsCreateV1ParamsWithContext(ctx context.Context) *CbExclusionsCreateV1Params {
	return &CbExclusionsCreateV1Params{
		Context: ctx,
	}
}

// NewCbExclusionsCreateV1ParamsWithHTTPClient creates a new CbExclusionsCreateV1Params object
// with the ability to set a custom HTTPClient for a request.
func NewCbExclusionsCreateV1ParamsWithHTTPClient(client *http.Client) *CbExclusionsCreateV1Params {
	return &CbExclusionsCreateV1Params{
		HTTPClient: client,
	}
}

/*
CbExclusionsCreateV1Params contains all the parameters to send to the API endpoint

	for the cb exclusions create v1 operation.

	Typically these are written to a http.Request.
*/
type CbExclusionsCreateV1Params struct {

	// Body.
	Body *models.APICertBasedExclusionsCreateReqV1

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cb exclusions create v1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CbExclusionsCreateV1Params) WithDefaults() *CbExclusionsCreateV1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cb exclusions create v1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CbExclusionsCreateV1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the cb exclusions create v1 params
func (o *CbExclusionsCreateV1Params) WithTimeout(timeout time.Duration) *CbExclusionsCreateV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cb exclusions create v1 params
func (o *CbExclusionsCreateV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cb exclusions create v1 params
func (o *CbExclusionsCreateV1Params) WithContext(ctx context.Context) *CbExclusionsCreateV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cb exclusions create v1 params
func (o *CbExclusionsCreateV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cb exclusions create v1 params
func (o *CbExclusionsCreateV1Params) WithHTTPClient(client *http.Client) *CbExclusionsCreateV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cb exclusions create v1 params
func (o *CbExclusionsCreateV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the cb exclusions create v1 params
func (o *CbExclusionsCreateV1Params) WithBody(body *models.APICertBasedExclusionsCreateReqV1) *CbExclusionsCreateV1Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the cb exclusions create v1 params
func (o *CbExclusionsCreateV1Params) SetBody(body *models.APICertBasedExclusionsCreateReqV1) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CbExclusionsCreateV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

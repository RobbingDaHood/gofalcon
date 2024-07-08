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

// NewCbExclusionsUpdateV1Params creates a new CbExclusionsUpdateV1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCbExclusionsUpdateV1Params() *CbExclusionsUpdateV1Params {
	return &CbExclusionsUpdateV1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewCbExclusionsUpdateV1ParamsWithTimeout creates a new CbExclusionsUpdateV1Params object
// with the ability to set a timeout on a request.
func NewCbExclusionsUpdateV1ParamsWithTimeout(timeout time.Duration) *CbExclusionsUpdateV1Params {
	return &CbExclusionsUpdateV1Params{
		timeout: timeout,
	}
}

// NewCbExclusionsUpdateV1ParamsWithContext creates a new CbExclusionsUpdateV1Params object
// with the ability to set a context for a request.
func NewCbExclusionsUpdateV1ParamsWithContext(ctx context.Context) *CbExclusionsUpdateV1Params {
	return &CbExclusionsUpdateV1Params{
		Context: ctx,
	}
}

// NewCbExclusionsUpdateV1ParamsWithHTTPClient creates a new CbExclusionsUpdateV1Params object
// with the ability to set a custom HTTPClient for a request.
func NewCbExclusionsUpdateV1ParamsWithHTTPClient(client *http.Client) *CbExclusionsUpdateV1Params {
	return &CbExclusionsUpdateV1Params{
		HTTPClient: client,
	}
}

/*
CbExclusionsUpdateV1Params contains all the parameters to send to the API endpoint

	for the cb exclusions update v1 operation.

	Typically these are written to a http.Request.
*/
type CbExclusionsUpdateV1Params struct {

	// Body.
	Body *models.APICertBasedExclusionsUpdateReqV1

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cb exclusions update v1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CbExclusionsUpdateV1Params) WithDefaults() *CbExclusionsUpdateV1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cb exclusions update v1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CbExclusionsUpdateV1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the cb exclusions update v1 params
func (o *CbExclusionsUpdateV1Params) WithTimeout(timeout time.Duration) *CbExclusionsUpdateV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cb exclusions update v1 params
func (o *CbExclusionsUpdateV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cb exclusions update v1 params
func (o *CbExclusionsUpdateV1Params) WithContext(ctx context.Context) *CbExclusionsUpdateV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cb exclusions update v1 params
func (o *CbExclusionsUpdateV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cb exclusions update v1 params
func (o *CbExclusionsUpdateV1Params) WithHTTPClient(client *http.Client) *CbExclusionsUpdateV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cb exclusions update v1 params
func (o *CbExclusionsUpdateV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the cb exclusions update v1 params
func (o *CbExclusionsUpdateV1Params) WithBody(body *models.APICertBasedExclusionsUpdateReqV1) *CbExclusionsUpdateV1Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the cb exclusions update v1 params
func (o *CbExclusionsUpdateV1Params) SetBody(body *models.APICertBasedExclusionsUpdateReqV1) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CbExclusionsUpdateV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

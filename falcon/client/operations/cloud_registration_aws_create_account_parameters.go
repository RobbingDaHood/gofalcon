// Code generated by go-swagger; DO NOT EDIT.

package operations

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

// NewCloudRegistrationAwsCreateAccountParams creates a new CloudRegistrationAwsCreateAccountParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCloudRegistrationAwsCreateAccountParams() *CloudRegistrationAwsCreateAccountParams {
	return &CloudRegistrationAwsCreateAccountParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCloudRegistrationAwsCreateAccountParamsWithTimeout creates a new CloudRegistrationAwsCreateAccountParams object
// with the ability to set a timeout on a request.
func NewCloudRegistrationAwsCreateAccountParamsWithTimeout(timeout time.Duration) *CloudRegistrationAwsCreateAccountParams {
	return &CloudRegistrationAwsCreateAccountParams{
		timeout: timeout,
	}
}

// NewCloudRegistrationAwsCreateAccountParamsWithContext creates a new CloudRegistrationAwsCreateAccountParams object
// with the ability to set a context for a request.
func NewCloudRegistrationAwsCreateAccountParamsWithContext(ctx context.Context) *CloudRegistrationAwsCreateAccountParams {
	return &CloudRegistrationAwsCreateAccountParams{
		Context: ctx,
	}
}

// NewCloudRegistrationAwsCreateAccountParamsWithHTTPClient creates a new CloudRegistrationAwsCreateAccountParams object
// with the ability to set a custom HTTPClient for a request.
func NewCloudRegistrationAwsCreateAccountParamsWithHTTPClient(client *http.Client) *CloudRegistrationAwsCreateAccountParams {
	return &CloudRegistrationAwsCreateAccountParams{
		HTTPClient: client,
	}
}

/*
CloudRegistrationAwsCreateAccountParams contains all the parameters to send to the API endpoint

	for the cloud registration aws create account operation.

	Typically these are written to a http.Request.
*/
type CloudRegistrationAwsCreateAccountParams struct {

	// Body.
	Body *models.RestAWSAccountCreateRequestExtv1

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cloud registration aws create account params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CloudRegistrationAwsCreateAccountParams) WithDefaults() *CloudRegistrationAwsCreateAccountParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cloud registration aws create account params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CloudRegistrationAwsCreateAccountParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the cloud registration aws create account params
func (o *CloudRegistrationAwsCreateAccountParams) WithTimeout(timeout time.Duration) *CloudRegistrationAwsCreateAccountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cloud registration aws create account params
func (o *CloudRegistrationAwsCreateAccountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cloud registration aws create account params
func (o *CloudRegistrationAwsCreateAccountParams) WithContext(ctx context.Context) *CloudRegistrationAwsCreateAccountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cloud registration aws create account params
func (o *CloudRegistrationAwsCreateAccountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cloud registration aws create account params
func (o *CloudRegistrationAwsCreateAccountParams) WithHTTPClient(client *http.Client) *CloudRegistrationAwsCreateAccountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cloud registration aws create account params
func (o *CloudRegistrationAwsCreateAccountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the cloud registration aws create account params
func (o *CloudRegistrationAwsCreateAccountParams) WithBody(body *models.RestAWSAccountCreateRequestExtv1) *CloudRegistrationAwsCreateAccountParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the cloud registration aws create account params
func (o *CloudRegistrationAwsCreateAccountParams) SetBody(body *models.RestAWSAccountCreateRequestExtv1) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CloudRegistrationAwsCreateAccountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

// Code generated by go-swagger; DO NOT EDIT.

package kubernetes_protection

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
)

// NewDeploymentCountParams creates a new DeploymentCountParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeploymentCountParams() *DeploymentCountParams {
	return &DeploymentCountParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeploymentCountParamsWithTimeout creates a new DeploymentCountParams object
// with the ability to set a timeout on a request.
func NewDeploymentCountParamsWithTimeout(timeout time.Duration) *DeploymentCountParams {
	return &DeploymentCountParams{
		timeout: timeout,
	}
}

// NewDeploymentCountParamsWithContext creates a new DeploymentCountParams object
// with the ability to set a context for a request.
func NewDeploymentCountParamsWithContext(ctx context.Context) *DeploymentCountParams {
	return &DeploymentCountParams{
		Context: ctx,
	}
}

// NewDeploymentCountParamsWithHTTPClient creates a new DeploymentCountParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeploymentCountParamsWithHTTPClient(client *http.Client) *DeploymentCountParams {
	return &DeploymentCountParams{
		HTTPClient: client,
	}
}

/*
DeploymentCountParams contains all the parameters to send to the API endpoint

	for the deployment count operation.

	Typically these are written to a http.Request.
*/
type DeploymentCountParams struct {

	/* Filter.

	   Retrieve count of Kubernetes deployments that match a query in Falcon Query Language (FQL). Supported filters:  agent_type,annotations_list,cid,cloud_account_id,cloud_name,cloud_region,cluster_id,cluster_name,deployment_id,deployment_name,first_seen,last_seen,namespace,pod_count,resource_status
	*/
	Filter *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the deployment count params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeploymentCountParams) WithDefaults() *DeploymentCountParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the deployment count params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeploymentCountParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the deployment count params
func (o *DeploymentCountParams) WithTimeout(timeout time.Duration) *DeploymentCountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the deployment count params
func (o *DeploymentCountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the deployment count params
func (o *DeploymentCountParams) WithContext(ctx context.Context) *DeploymentCountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the deployment count params
func (o *DeploymentCountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the deployment count params
func (o *DeploymentCountParams) WithHTTPClient(client *http.Client) *DeploymentCountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the deployment count params
func (o *DeploymentCountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilter adds the filter to the deployment count params
func (o *DeploymentCountParams) WithFilter(filter *string) *DeploymentCountParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the deployment count params
func (o *DeploymentCountParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WriteToRequest writes these params to a swagger request
func (o *DeploymentCountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

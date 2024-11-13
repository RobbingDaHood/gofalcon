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

// NewClusterCountParams creates a new ClusterCountParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewClusterCountParams() *ClusterCountParams {
	return &ClusterCountParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewClusterCountParamsWithTimeout creates a new ClusterCountParams object
// with the ability to set a timeout on a request.
func NewClusterCountParamsWithTimeout(timeout time.Duration) *ClusterCountParams {
	return &ClusterCountParams{
		timeout: timeout,
	}
}

// NewClusterCountParamsWithContext creates a new ClusterCountParams object
// with the ability to set a context for a request.
func NewClusterCountParamsWithContext(ctx context.Context) *ClusterCountParams {
	return &ClusterCountParams{
		Context: ctx,
	}
}

// NewClusterCountParamsWithHTTPClient creates a new ClusterCountParams object
// with the ability to set a custom HTTPClient for a request.
func NewClusterCountParamsWithHTTPClient(client *http.Client) *ClusterCountParams {
	return &ClusterCountParams{
		HTTPClient: client,
	}
}

/*
ClusterCountParams contains all the parameters to send to the API endpoint

	for the cluster count operation.

	Typically these are written to a http.Request.
*/
type ClusterCountParams struct {

	/* Filter.

	   Retrieve count of Kubernetes clusters that match a query in Falcon Query Language (FQL). Supported filters:  access,agent_id,agent_status,agent_type,cid,cloud_account_id,cloud_name,cloud_region,cloud_service,cluster_id,cluster_name,cluster_status,container_count,iar_coverage,kac_agent_id,kubernetes_version,last_seen,management_status,node_count,pod_count,tags
	*/
	Filter *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cluster count params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ClusterCountParams) WithDefaults() *ClusterCountParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cluster count params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ClusterCountParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the cluster count params
func (o *ClusterCountParams) WithTimeout(timeout time.Duration) *ClusterCountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cluster count params
func (o *ClusterCountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cluster count params
func (o *ClusterCountParams) WithContext(ctx context.Context) *ClusterCountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cluster count params
func (o *ClusterCountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cluster count params
func (o *ClusterCountParams) WithHTTPClient(client *http.Client) *ClusterCountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cluster count params
func (o *ClusterCountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilter adds the filter to the cluster count params
func (o *ClusterCountParams) WithFilter(filter *string) *ClusterCountParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the cluster count params
func (o *ClusterCountParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WriteToRequest writes these params to a swagger request
func (o *ClusterCountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

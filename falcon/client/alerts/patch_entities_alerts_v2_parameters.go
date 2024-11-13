// Code generated by go-swagger; DO NOT EDIT.

package alerts

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

// NewPatchEntitiesAlertsV2Params creates a new PatchEntitiesAlertsV2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchEntitiesAlertsV2Params() *PatchEntitiesAlertsV2Params {
	return &PatchEntitiesAlertsV2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchEntitiesAlertsV2ParamsWithTimeout creates a new PatchEntitiesAlertsV2Params object
// with the ability to set a timeout on a request.
func NewPatchEntitiesAlertsV2ParamsWithTimeout(timeout time.Duration) *PatchEntitiesAlertsV2Params {
	return &PatchEntitiesAlertsV2Params{
		timeout: timeout,
	}
}

// NewPatchEntitiesAlertsV2ParamsWithContext creates a new PatchEntitiesAlertsV2Params object
// with the ability to set a context for a request.
func NewPatchEntitiesAlertsV2ParamsWithContext(ctx context.Context) *PatchEntitiesAlertsV2Params {
	return &PatchEntitiesAlertsV2Params{
		Context: ctx,
	}
}

// NewPatchEntitiesAlertsV2ParamsWithHTTPClient creates a new PatchEntitiesAlertsV2Params object
// with the ability to set a custom HTTPClient for a request.
func NewPatchEntitiesAlertsV2ParamsWithHTTPClient(client *http.Client) *PatchEntitiesAlertsV2Params {
	return &PatchEntitiesAlertsV2Params{
		HTTPClient: client,
	}
}

/*
PatchEntitiesAlertsV2Params contains all the parameters to send to the API endpoint

	for the patch entities alerts v2 operation.

	Typically these are written to a http.Request.
*/
type PatchEntitiesAlertsV2Params struct {

	/* Body.

	     `ids`
	- IDs of Alerts to modify.

	`action_parameters` values
	- `assign_to_uuid`
		- Assign Alert to user UUID, such as `00000000-0000-0000-0000-000000000000`
	- `assign_to_user_id`
		- Assign Alert to user ID, such as `user@example.com`
	- `assign_to_name`
		- Assign Alert to username, such as `John Doe`
	- `unassign`
		- Unassign Alert clears out the assigned user UUID, user ID, and username.
	- `add_tag`
	 	- Add a tag to the Alert.
	- `remove_tag`
		- Remove a tag from the Alert.
	- `remove_tags_by_prefix`
		- Remove tags from the Alert based on the prefix.
	- `append_comment`
		- Comments are displayed with the Alert in Falcon and are usually used to provide context or notes for other Falcon users. An Alert can have multiple comments over time.
	- `update_status` values
		- `new`
		- `in_progress`
		- `reopened`
		- `closed`
	- `show_in_ui` values
		- `true`: This alert is displayed in Falcon
		- `false`: This alert is not displayed in Falcon.

	*/
	Body *models.DetectsapiPatchEntitiesAlertsV2Request

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch entities alerts v2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchEntitiesAlertsV2Params) WithDefaults() *PatchEntitiesAlertsV2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch entities alerts v2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchEntitiesAlertsV2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch entities alerts v2 params
func (o *PatchEntitiesAlertsV2Params) WithTimeout(timeout time.Duration) *PatchEntitiesAlertsV2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch entities alerts v2 params
func (o *PatchEntitiesAlertsV2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch entities alerts v2 params
func (o *PatchEntitiesAlertsV2Params) WithContext(ctx context.Context) *PatchEntitiesAlertsV2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch entities alerts v2 params
func (o *PatchEntitiesAlertsV2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch entities alerts v2 params
func (o *PatchEntitiesAlertsV2Params) WithHTTPClient(client *http.Client) *PatchEntitiesAlertsV2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch entities alerts v2 params
func (o *PatchEntitiesAlertsV2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch entities alerts v2 params
func (o *PatchEntitiesAlertsV2Params) WithBody(body *models.DetectsapiPatchEntitiesAlertsV2Request) *PatchEntitiesAlertsV2Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch entities alerts v2 params
func (o *PatchEntitiesAlertsV2Params) SetBody(body *models.DetectsapiPatchEntitiesAlertsV2Request) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PatchEntitiesAlertsV2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

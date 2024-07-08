// Code generated by go-swagger; DO NOT EDIT.

package certificate_based_exclusions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new certificate based exclusions API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for certificate based exclusions API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CbExclusionsCreateV1(params *CbExclusionsCreateV1Params, opts ...ClientOption) (*CbExclusionsCreateV1Created, error)

	CbExclusionsDeleteV1(params *CbExclusionsDeleteV1Params, opts ...ClientOption) (*CbExclusionsDeleteV1OK, error)

	CbExclusionsGetV1(params *CbExclusionsGetV1Params, opts ...ClientOption) (*CbExclusionsGetV1OK, error)

	CbExclusionsQueryV1(params *CbExclusionsQueryV1Params, opts ...ClientOption) (*CbExclusionsQueryV1OK, error)

	CbExclusionsUpdateV1(params *CbExclusionsUpdateV1Params, opts ...ClientOption) (*CbExclusionsUpdateV1OK, error)

	CertificatesGetV1(params *CertificatesGetV1Params, opts ...ClientOption) (*CertificatesGetV1OK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CbExclusionsCreateV1 creates new certificate based exclusions
*/
func (a *Client) CbExclusionsCreateV1(params *CbExclusionsCreateV1Params, opts ...ClientOption) (*CbExclusionsCreateV1Created, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCbExclusionsCreateV1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "cb-exclusions.create.v1",
		Method:             "POST",
		PathPattern:        "/exclusions/entities/cert-based-exclusions/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CbExclusionsCreateV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CbExclusionsCreateV1Created)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for cb-exclusions.create.v1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CbExclusionsDeleteV1 deletes the exclusions by id
*/
func (a *Client) CbExclusionsDeleteV1(params *CbExclusionsDeleteV1Params, opts ...ClientOption) (*CbExclusionsDeleteV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCbExclusionsDeleteV1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "cb-exclusions.delete.v1",
		Method:             "DELETE",
		PathPattern:        "/exclusions/entities/cert-based-exclusions/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CbExclusionsDeleteV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CbExclusionsDeleteV1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for cb-exclusions.delete.v1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CbExclusionsGetV1 finds all exclusion i ds matching the query with filter
*/
func (a *Client) CbExclusionsGetV1(params *CbExclusionsGetV1Params, opts ...ClientOption) (*CbExclusionsGetV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCbExclusionsGetV1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "cb-exclusions.get.v1",
		Method:             "GET",
		PathPattern:        "/exclusions/entities/cert-based-exclusions/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CbExclusionsGetV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CbExclusionsGetV1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for cb-exclusions.get.v1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CbExclusionsQueryV1 searches for cert based exclusions
*/
func (a *Client) CbExclusionsQueryV1(params *CbExclusionsQueryV1Params, opts ...ClientOption) (*CbExclusionsQueryV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCbExclusionsQueryV1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "cb-exclusions.query.v1",
		Method:             "GET",
		PathPattern:        "/exclusions/queries/cert-based-exclusions/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CbExclusionsQueryV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CbExclusionsQueryV1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for cb-exclusions.query.v1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CbExclusionsUpdateV1 updates existing certificate based exclusions
*/
func (a *Client) CbExclusionsUpdateV1(params *CbExclusionsUpdateV1Params, opts ...ClientOption) (*CbExclusionsUpdateV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCbExclusionsUpdateV1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "cb-exclusions.update.v1",
		Method:             "PATCH",
		PathPattern:        "/exclusions/entities/cert-based-exclusions/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CbExclusionsUpdateV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CbExclusionsUpdateV1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for cb-exclusions.update.v1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CertificatesGetV1 retrieves certificate signing information for a file
*/
func (a *Client) CertificatesGetV1(params *CertificatesGetV1Params, opts ...ClientOption) (*CertificatesGetV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCertificatesGetV1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "certificates.get.v1",
		Method:             "GET",
		PathPattern:        "/exclusions/entities/certificates/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CertificatesGetV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CertificatesGetV1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for certificates.get.v1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

// Code generated by go-swagger; DO NOT EDIT.

package custom_storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new custom storage API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for custom storage API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteVersionedObject(params *DeleteVersionedObjectParams, opts ...ClientOption) (*DeleteVersionedObjectOK, error)

	GetVersionedObject(params *GetVersionedObjectParams, writer io.Writer, opts ...ClientOption) (*GetVersionedObjectOK, error)

	GetVersionedObjectMetadata(params *GetVersionedObjectMetadataParams, opts ...ClientOption) (*GetVersionedObjectMetadataOK, error)

	ListObjectsByVersion(params *ListObjectsByVersionParams, opts ...ClientOption) (*ListObjectsByVersionOK, error)

	PutObjectByVersion(params *PutObjectByVersionParams, opts ...ClientOption) (*PutObjectByVersionOK, error)

	SearchObjectsByVersion(params *SearchObjectsByVersionParams, opts ...ClientOption) (*SearchObjectsByVersionOK, error)

	Delete(params *DeleteParams, opts ...ClientOption) (*DeleteOK, error)

	Get(params *GetParams, writer io.Writer, opts ...ClientOption) (*GetOK, error)

	List(params *ListParams, opts ...ClientOption) (*ListOK, error)

	Metadata(params *MetadataParams, opts ...ClientOption) (*MetadataOK, error)

	Search(params *SearchParams, opts ...ClientOption) (*SearchOK, error)

	Upload(params *UploadParams, opts ...ClientOption) (*UploadOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
DeleteVersionedObject deletes the specified versioned object
*/
func (a *Client) DeleteVersionedObject(params *DeleteVersionedObjectParams, opts ...ClientOption) (*DeleteVersionedObjectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteVersionedObjectParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteVersionedObject",
		Method:             "DELETE",
		PathPattern:        "/customobjects/v1/collections/{collection_name}/{collection_version}/objects/{object_key}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteVersionedObjectReader{formats: a.formats},
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
	success, ok := result.(*DeleteVersionedObjectOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteVersionedObject: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetVersionedObject gets the bytes for the specified object
*/
func (a *Client) GetVersionedObject(params *GetVersionedObjectParams, writer io.Writer, opts ...ClientOption) (*GetVersionedObjectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVersionedObjectParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetVersionedObject",
		Method:             "GET",
		PathPattern:        "/customobjects/v1/collections/{collection_name}/{collection_version}/objects/{object_key}",
		ProducesMediaTypes: []string{"application/json", "application/octet-stream"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetVersionedObjectReader{formats: a.formats, writer: writer},
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
	success, ok := result.(*GetVersionedObjectOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetVersionedObject: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetVersionedObjectMetadata gets the metadata for the specified object
*/
func (a *Client) GetVersionedObjectMetadata(params *GetVersionedObjectMetadataParams, opts ...ClientOption) (*GetVersionedObjectMetadataOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVersionedObjectMetadataParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetVersionedObjectMetadata",
		Method:             "GET",
		PathPattern:        "/customobjects/v1/collections/{collection_name}/{collection_version}/objects/{object_key}/metadata",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetVersionedObjectMetadataReader{formats: a.formats},
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
	success, ok := result.(*GetVersionedObjectMetadataOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetVersionedObjectMetadata: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ListObjectsByVersion lists the object keys in the specified collection in alphabetical order
*/
func (a *Client) ListObjectsByVersion(params *ListObjectsByVersionParams, opts ...ClientOption) (*ListObjectsByVersionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListObjectsByVersionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ListObjectsByVersion",
		Method:             "GET",
		PathPattern:        "/customobjects/v1/collections/{collection_name}/{collection_version}/objects",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListObjectsByVersionReader{formats: a.formats},
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
	success, ok := result.(*ListObjectsByVersionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ListObjectsByVersion: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PutObjectByVersion puts the specified new object at the given key or overwrite an existing object at the given key
*/
func (a *Client) PutObjectByVersion(params *PutObjectByVersionParams, opts ...ClientOption) (*PutObjectByVersionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutObjectByVersionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PutObjectByVersion",
		Method:             "PUT",
		PathPattern:        "/customobjects/v1/collections/{collection_name}/{collection_version}/objects/{object_key}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/octet-stream"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutObjectByVersionReader{formats: a.formats},
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
	success, ok := result.(*PutObjectByVersionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PutObjectByVersion: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
SearchObjectsByVersion searches for objects that match the specified filter criteria returns metadata not actual objects
*/
func (a *Client) SearchObjectsByVersion(params *SearchObjectsByVersionParams, opts ...ClientOption) (*SearchObjectsByVersionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSearchObjectsByVersionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SearchObjectsByVersion",
		Method:             "POST",
		PathPattern:        "/customobjects/v1/collections/{collection_name}/{collection_version}/objects",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/octet-stream"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SearchObjectsByVersionReader{formats: a.formats},
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
	success, ok := result.(*SearchObjectsByVersionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for SearchObjectsByVersion: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
Delete deletes the specified object
*/
func (a *Client) Delete(params *DeleteParams, opts ...ClientOption) (*DeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "delete",
		Method:             "DELETE",
		PathPattern:        "/customobjects/v1/collections/{collection_name}/objects/{object_key}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteReader{formats: a.formats},
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
	success, ok := result.(*DeleteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
Get gets the bytes for the specified object
*/
func (a *Client) Get(params *GetParams, writer io.Writer, opts ...ClientOption) (*GetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get",
		Method:             "GET",
		PathPattern:        "/customobjects/v1/collections/{collection_name}/objects/{object_key}",
		ProducesMediaTypes: []string{"application/json", "application/octet-stream"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetReader{formats: a.formats, writer: writer},
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
	success, ok := result.(*GetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
List lists the object keys in the specified collection in alphabetical order
*/
func (a *Client) List(params *ListParams, opts ...ClientOption) (*ListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "list",
		Method:             "GET",
		PathPattern:        "/customobjects/v1/collections/{collection_name}/objects",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListReader{formats: a.formats},
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
	success, ok := result.(*ListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
Metadata gets the metadata for the specified object
*/
func (a *Client) Metadata(params *MetadataParams, opts ...ClientOption) (*MetadataOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMetadataParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "metadata",
		Method:             "GET",
		PathPattern:        "/customobjects/v1/collections/{collection_name}/objects/{object_key}/metadata",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &MetadataReader{formats: a.formats},
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
	success, ok := result.(*MetadataOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for metadata: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
Search searches for objects that match the specified filter criteria returns metadata not actual objects
*/
func (a *Client) Search(params *SearchParams, opts ...ClientOption) (*SearchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSearchParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "search",
		Method:             "POST",
		PathPattern:        "/customobjects/v1/collections/{collection_name}/objects",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/octet-stream"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SearchReader{formats: a.formats},
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
	success, ok := result.(*SearchOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for search: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
Upload puts the specified new object at the given key or overwrite an existing object at the given key
*/
func (a *Client) Upload(params *UploadParams, opts ...ClientOption) (*UploadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUploadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "upload",
		Method:             "PUT",
		PathPattern:        "/customobjects/v1/collections/{collection_name}/objects/{object_key}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/octet-stream"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UploadReader{formats: a.formats},
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
	success, ok := result.(*UploadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for upload: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

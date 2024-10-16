// Code generated by go-swagger; DO NOT EDIT.

package application

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new application API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new application API client with basic auth credentials.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - user: user for basic authentication header.
// - password: password for basic authentication header.
func NewClientWithBasicAuth(host, basePath, scheme, user, password string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BasicAuth(user, password)
	return &Client{transport: transport, formats: strfmt.Default}
}

// New creates a new application API client with a bearer token for authentication.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - bearerToken: bearer token for Bearer authentication header.
func NewClientWithBearerToken(host, basePath, scheme, bearerToken string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BearerToken(bearerToken)
	return &Client{transport: transport, formats: strfmt.Default}
}

/*
Client for application API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// This client is generated with a few options you might find useful for your swagger spec.
//
// Feel free to add you own set of options.

// WithContentType allows the client to force the Content-Type header
// to negotiate a specific Consumer from the server.
//
// You may use this option to set arbitrary extensions to your MIME media type.
func WithContentType(mime string) ClientOption {
	return func(r *runtime.ClientOperation) {
		r.ConsumesMediaTypes = []string{mime}
	}
}

// WithContentTypeApplicationJSON sets the Content-Type header to "application/json".
func WithContentTypeApplicationJSON(r *runtime.ClientOperation) {
	r.ConsumesMediaTypes = []string{"application/json"}
}

// WithContentTypeMultipartFormData sets the Content-Type header to "multipart/form-data".
func WithContentTypeMultipartFormData(r *runtime.ClientOperation) {
	r.ConsumesMediaTypes = []string{"multipart/form-data"}
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateApp(params *CreateAppParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateAppOK, error)

	DeleteApp(params *DeleteAppParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteAppOK, error)

	GetApps(params *GetAppsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAppsOK, error)

	RemoveAppImage(params *RemoveAppImageParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RemoveAppImageOK, error)

	UpdateApplication(params *UpdateApplicationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateApplicationOK, error)

	UploadAppImage(params *UploadAppImageParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UploadAppImageOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateApp creates an application
*/
func (a *Client) CreateApp(params *CreateAppParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateAppOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateAppParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createApp",
		Method:             "POST",
		PathPattern:        "/application",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateAppReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*CreateAppOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createApp: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteApp deletes an application
*/
func (a *Client) DeleteApp(params *DeleteAppParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteAppOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteAppParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteApp",
		Method:             "DELETE",
		PathPattern:        "/application/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteAppReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*DeleteAppOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteApp: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetApps returns all applications
*/
func (a *Client) GetApps(params *GetAppsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAppsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAppsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getApps",
		Method:             "GET",
		PathPattern:        "/application",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetAppsReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*GetAppsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getApps: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
RemoveAppImage deletes an image of an application
*/
func (a *Client) RemoveAppImage(params *RemoveAppImageParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RemoveAppImageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRemoveAppImageParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "removeAppImage",
		Method:             "DELETE",
		PathPattern:        "/application/{id}/image",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RemoveAppImageReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*RemoveAppImageOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for removeAppImage: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateApplication updates an application
*/
func (a *Client) UpdateApplication(params *UpdateApplicationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateApplicationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateApplicationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateApplication",
		Method:             "PUT",
		PathPattern:        "/application/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateApplicationReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*UpdateApplicationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateApplication: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UploadAppImage uploads an image for an application
*/
func (a *Client) UploadAppImage(params *UploadAppImageParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UploadAppImageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUploadAppImageParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "uploadAppImage",
		Method:             "POST",
		PathPattern:        "/application/{id}/image",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UploadAppImageReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*UploadAppImageOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for uploadAppImage: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

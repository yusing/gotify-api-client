// Code generated by go-swagger; DO NOT EDIT.

package message

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new message API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new message API client with basic auth credentials.
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

// New creates a new message API client with a bearer token for authentication.
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
Client for message API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateMessage(params *CreateMessageParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateMessageOK, error)

	DeleteAppMessages(params *DeleteAppMessagesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteAppMessagesOK, error)

	DeleteMessage(params *DeleteMessageParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteMessageOK, error)

	DeleteMessages(params *DeleteMessagesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteMessagesOK, error)

	GetAppMessages(params *GetAppMessagesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAppMessagesOK, error)

	GetMessages(params *GetMessagesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetMessagesOK, error)

	StreamMessages(params *StreamMessagesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StreamMessagesOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateMessage creates a message

__NOTE__: This API ONLY accepts an application token as authentication.
*/
func (a *Client) CreateMessage(params *CreateMessageParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateMessageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateMessageParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createMessage",
		Method:             "POST",
		PathPattern:        "/message",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateMessageReader{formats: a.formats},
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
	success, ok := result.(*CreateMessageOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createMessage: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteAppMessages deletes all messages from a specific application
*/
func (a *Client) DeleteAppMessages(params *DeleteAppMessagesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteAppMessagesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteAppMessagesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteAppMessages",
		Method:             "DELETE",
		PathPattern:        "/application/{id}/message",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteAppMessagesReader{formats: a.formats},
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
	success, ok := result.(*DeleteAppMessagesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteAppMessages: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteMessage deletes a message with an id
*/
func (a *Client) DeleteMessage(params *DeleteMessageParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteMessageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteMessageParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteMessage",
		Method:             "DELETE",
		PathPattern:        "/message/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteMessageReader{formats: a.formats},
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
	success, ok := result.(*DeleteMessageOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteMessage: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteMessages deletes all messages
*/
func (a *Client) DeleteMessages(params *DeleteMessagesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteMessagesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteMessagesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteMessages",
		Method:             "DELETE",
		PathPattern:        "/message",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteMessagesReader{formats: a.formats},
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
	success, ok := result.(*DeleteMessagesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteMessages: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAppMessages returns all messages from a specific application
*/
func (a *Client) GetAppMessages(params *GetAppMessagesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAppMessagesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAppMessagesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAppMessages",
		Method:             "GET",
		PathPattern:        "/application/{id}/message",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetAppMessagesReader{formats: a.formats},
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
	success, ok := result.(*GetAppMessagesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAppMessages: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetMessages returns all messages
*/
func (a *Client) GetMessages(params *GetMessagesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetMessagesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMessagesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getMessages",
		Method:             "GET",
		PathPattern:        "/message",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetMessagesReader{formats: a.formats},
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
	success, ok := result.(*GetMessagesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getMessages: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
StreamMessages websockets return newly created messages
*/
func (a *Client) StreamMessages(params *StreamMessagesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StreamMessagesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStreamMessagesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "streamMessages",
		Method:             "GET",
		PathPattern:        "/stream",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StreamMessagesReader{formats: a.formats},
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
	success, ok := result.(*StreamMessagesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for streamMessages: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

// Code generated by go-swagger; DO NOT EDIT.

package message

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

// NewStreamMessagesParams creates a new StreamMessagesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStreamMessagesParams() *StreamMessagesParams {
	return &StreamMessagesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStreamMessagesParamsWithTimeout creates a new StreamMessagesParams object
// with the ability to set a timeout on a request.
func NewStreamMessagesParamsWithTimeout(timeout time.Duration) *StreamMessagesParams {
	return &StreamMessagesParams{
		timeout: timeout,
	}
}

// NewStreamMessagesParamsWithContext creates a new StreamMessagesParams object
// with the ability to set a context for a request.
func NewStreamMessagesParamsWithContext(ctx context.Context) *StreamMessagesParams {
	return &StreamMessagesParams{
		Context: ctx,
	}
}

// NewStreamMessagesParamsWithHTTPClient creates a new StreamMessagesParams object
// with the ability to set a custom HTTPClient for a request.
func NewStreamMessagesParamsWithHTTPClient(client *http.Client) *StreamMessagesParams {
	return &StreamMessagesParams{
		HTTPClient: client,
	}
}

/*
StreamMessagesParams contains all the parameters to send to the API endpoint

	for the stream messages operation.

	Typically these are written to a http.Request.
*/
type StreamMessagesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the stream messages params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StreamMessagesParams) WithDefaults() *StreamMessagesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the stream messages params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StreamMessagesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the stream messages params
func (o *StreamMessagesParams) WithTimeout(timeout time.Duration) *StreamMessagesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the stream messages params
func (o *StreamMessagesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the stream messages params
func (o *StreamMessagesParams) WithContext(ctx context.Context) *StreamMessagesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the stream messages params
func (o *StreamMessagesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the stream messages params
func (o *StreamMessagesParams) WithHTTPClient(client *http.Client) *StreamMessagesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the stream messages params
func (o *StreamMessagesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *StreamMessagesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

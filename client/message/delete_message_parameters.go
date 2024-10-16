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
	"github.com/go-openapi/swag"
)

// NewDeleteMessageParams creates a new DeleteMessageParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteMessageParams() *DeleteMessageParams {
	return &DeleteMessageParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteMessageParamsWithTimeout creates a new DeleteMessageParams object
// with the ability to set a timeout on a request.
func NewDeleteMessageParamsWithTimeout(timeout time.Duration) *DeleteMessageParams {
	return &DeleteMessageParams{
		timeout: timeout,
	}
}

// NewDeleteMessageParamsWithContext creates a new DeleteMessageParams object
// with the ability to set a context for a request.
func NewDeleteMessageParamsWithContext(ctx context.Context) *DeleteMessageParams {
	return &DeleteMessageParams{
		Context: ctx,
	}
}

// NewDeleteMessageParamsWithHTTPClient creates a new DeleteMessageParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteMessageParamsWithHTTPClient(client *http.Client) *DeleteMessageParams {
	return &DeleteMessageParams{
		HTTPClient: client,
	}
}

/*
DeleteMessageParams contains all the parameters to send to the API endpoint

	for the delete message operation.

	Typically these are written to a http.Request.
*/
type DeleteMessageParams struct {

	/* ID.

	   the message id

	   Format: int64
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete message params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteMessageParams) WithDefaults() *DeleteMessageParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete message params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteMessageParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete message params
func (o *DeleteMessageParams) WithTimeout(timeout time.Duration) *DeleteMessageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete message params
func (o *DeleteMessageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete message params
func (o *DeleteMessageParams) WithContext(ctx context.Context) *DeleteMessageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete message params
func (o *DeleteMessageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete message params
func (o *DeleteMessageParams) WithHTTPClient(client *http.Client) *DeleteMessageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete message params
func (o *DeleteMessageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete message params
func (o *DeleteMessageParams) WithID(id int64) *DeleteMessageParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete message params
func (o *DeleteMessageParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteMessageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

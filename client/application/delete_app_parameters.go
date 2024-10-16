// Code generated by go-swagger; DO NOT EDIT.

package application

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

// NewDeleteAppParams creates a new DeleteAppParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteAppParams() *DeleteAppParams {
	return &DeleteAppParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAppParamsWithTimeout creates a new DeleteAppParams object
// with the ability to set a timeout on a request.
func NewDeleteAppParamsWithTimeout(timeout time.Duration) *DeleteAppParams {
	return &DeleteAppParams{
		timeout: timeout,
	}
}

// NewDeleteAppParamsWithContext creates a new DeleteAppParams object
// with the ability to set a context for a request.
func NewDeleteAppParamsWithContext(ctx context.Context) *DeleteAppParams {
	return &DeleteAppParams{
		Context: ctx,
	}
}

// NewDeleteAppParamsWithHTTPClient creates a new DeleteAppParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteAppParamsWithHTTPClient(client *http.Client) *DeleteAppParams {
	return &DeleteAppParams{
		HTTPClient: client,
	}
}

/*
DeleteAppParams contains all the parameters to send to the API endpoint

	for the delete app operation.

	Typically these are written to a http.Request.
*/
type DeleteAppParams struct {

	/* ID.

	   the application id

	   Format: int64
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete app params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAppParams) WithDefaults() *DeleteAppParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete app params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAppParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete app params
func (o *DeleteAppParams) WithTimeout(timeout time.Duration) *DeleteAppParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete app params
func (o *DeleteAppParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete app params
func (o *DeleteAppParams) WithContext(ctx context.Context) *DeleteAppParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete app params
func (o *DeleteAppParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete app params
func (o *DeleteAppParams) WithHTTPClient(client *http.Client) *DeleteAppParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete app params
func (o *DeleteAppParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete app params
func (o *DeleteAppParams) WithID(id int64) *DeleteAppParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete app params
func (o *DeleteAppParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAppParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

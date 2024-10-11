// Code generated by go-swagger; DO NOT EDIT.

package user

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

// NewCurrentUserParams creates a new CurrentUserParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCurrentUserParams() *CurrentUserParams {
	return &CurrentUserParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCurrentUserParamsWithTimeout creates a new CurrentUserParams object
// with the ability to set a timeout on a request.
func NewCurrentUserParamsWithTimeout(timeout time.Duration) *CurrentUserParams {
	return &CurrentUserParams{
		timeout: timeout,
	}
}

// NewCurrentUserParamsWithContext creates a new CurrentUserParams object
// with the ability to set a context for a request.
func NewCurrentUserParamsWithContext(ctx context.Context) *CurrentUserParams {
	return &CurrentUserParams{
		Context: ctx,
	}
}

// NewCurrentUserParamsWithHTTPClient creates a new CurrentUserParams object
// with the ability to set a custom HTTPClient for a request.
func NewCurrentUserParamsWithHTTPClient(client *http.Client) *CurrentUserParams {
	return &CurrentUserParams{
		HTTPClient: client,
	}
}

/*
CurrentUserParams contains all the parameters to send to the API endpoint

	for the current user operation.

	Typically these are written to a http.Request.
*/
type CurrentUserParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the current user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CurrentUserParams) WithDefaults() *CurrentUserParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the current user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CurrentUserParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the current user params
func (o *CurrentUserParams) WithTimeout(timeout time.Duration) *CurrentUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the current user params
func (o *CurrentUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the current user params
func (o *CurrentUserParams) WithContext(ctx context.Context) *CurrentUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the current user params
func (o *CurrentUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the current user params
func (o *CurrentUserParams) WithHTTPClient(client *http.Client) *CurrentUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the current user params
func (o *CurrentUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *CurrentUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

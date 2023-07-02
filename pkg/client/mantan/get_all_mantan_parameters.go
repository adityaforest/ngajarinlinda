// Code generated by go-swagger; DO NOT EDIT.

package mantan

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

// NewGetAllMantanParams creates a new GetAllMantanParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAllMantanParams() *GetAllMantanParams {
	return &GetAllMantanParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAllMantanParamsWithTimeout creates a new GetAllMantanParams object
// with the ability to set a timeout on a request.
func NewGetAllMantanParamsWithTimeout(timeout time.Duration) *GetAllMantanParams {
	return &GetAllMantanParams{
		timeout: timeout,
	}
}

// NewGetAllMantanParamsWithContext creates a new GetAllMantanParams object
// with the ability to set a context for a request.
func NewGetAllMantanParamsWithContext(ctx context.Context) *GetAllMantanParams {
	return &GetAllMantanParams{
		Context: ctx,
	}
}

// NewGetAllMantanParamsWithHTTPClient creates a new GetAllMantanParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAllMantanParamsWithHTTPClient(client *http.Client) *GetAllMantanParams {
	return &GetAllMantanParams{
		HTTPClient: client,
	}
}

/*
GetAllMantanParams contains all the parameters to send to the API endpoint

	for the get all mantan operation.

	Typically these are written to a http.Request.
*/
type GetAllMantanParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get all mantan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAllMantanParams) WithDefaults() *GetAllMantanParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get all mantan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAllMantanParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get all mantan params
func (o *GetAllMantanParams) WithTimeout(timeout time.Duration) *GetAllMantanParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get all mantan params
func (o *GetAllMantanParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get all mantan params
func (o *GetAllMantanParams) WithContext(ctx context.Context) *GetAllMantanParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get all mantan params
func (o *GetAllMantanParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get all mantan params
func (o *GetAllMantanParams) WithHTTPClient(client *http.Client) *GetAllMantanParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get all mantan params
func (o *GetAllMantanParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetAllMantanParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
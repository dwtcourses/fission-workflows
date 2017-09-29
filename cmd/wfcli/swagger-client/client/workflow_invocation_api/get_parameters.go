package workflow_invocation_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetParams creates a new GetParams object
// with the default values initialized.
func NewGetParams() *GetParams {
	var ()
	return &GetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetParamsWithTimeout creates a new GetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetParamsWithTimeout(timeout time.Duration) *GetParams {
	var ()
	return &GetParams{

		timeout: timeout,
	}
}

// NewGetParamsWithContext creates a new GetParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetParamsWithContext(ctx context.Context) *GetParams {
	var ()
	return &GetParams{

		Context: ctx,
	}
}

/*GetParams contains all the parameters to send to the API endpoint
for the get operation typically these are written to a http.Request
*/
type GetParams struct {

	/*ID*/
	ID string

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the get params
func (o *GetParams) WithTimeout(timeout time.Duration) *GetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get params
func (o *GetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get params
func (o *GetParams) WithContext(ctx context.Context) *GetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get params
func (o *GetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithID adds the id to the get params
func (o *GetParams) WithID(id string) *GetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get params
func (o *GetParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

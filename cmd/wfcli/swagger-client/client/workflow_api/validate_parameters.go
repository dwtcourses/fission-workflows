package workflow_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/fission/fission-workflows/cmd/wfcli/swagger-client/models"
)

// NewValidateParams creates a new ValidateParams object
// with the default values initialized.
func NewValidateParams() *ValidateParams {
	var ()
	return &ValidateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewValidateParamsWithTimeout creates a new ValidateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewValidateParamsWithTimeout(timeout time.Duration) *ValidateParams {
	var ()
	return &ValidateParams{

		timeout: timeout,
	}
}

// NewValidateParamsWithContext creates a new ValidateParams object
// with the default values initialized, and the ability to set a context for a request
func NewValidateParamsWithContext(ctx context.Context) *ValidateParams {
	var ()
	return &ValidateParams{

		Context: ctx,
	}
}

/*ValidateParams contains all the parameters to send to the API endpoint
for the validate operation typically these are written to a http.Request
*/
type ValidateParams struct {

	/*Body*/
	Body *models.WorkflowSpec

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the validate params
func (o *ValidateParams) WithTimeout(timeout time.Duration) *ValidateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the validate params
func (o *ValidateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the validate params
func (o *ValidateParams) WithContext(ctx context.Context) *ValidateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the validate params
func (o *ValidateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the validate params
func (o *ValidateParams) WithBody(body *models.WorkflowSpec) *ValidateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the validate params
func (o *ValidateParams) SetBody(body *models.WorkflowSpec) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ValidateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Body == nil {
		o.Body = new(models.WorkflowSpec)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/treeverse/lakefs/api/gen/models"
)

// NewCommitParams creates a new CommitParams object
// with the default values initialized.
func NewCommitParams() *CommitParams {
	var ()
	return &CommitParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCommitParamsWithTimeout creates a new CommitParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCommitParamsWithTimeout(timeout time.Duration) *CommitParams {
	var ()
	return &CommitParams{

		timeout: timeout,
	}
}

// NewCommitParamsWithContext creates a new CommitParams object
// with the default values initialized, and the ability to set a context for a request
func NewCommitParamsWithContext(ctx context.Context) *CommitParams {
	var ()
	return &CommitParams{

		Context: ctx,
	}
}

// NewCommitParamsWithHTTPClient creates a new CommitParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCommitParamsWithHTTPClient(client *http.Client) *CommitParams {
	var ()
	return &CommitParams{
		HTTPClient: client,
	}
}

/*CommitParams contains all the parameters to send to the API endpoint
for the commit operation typically these are written to a http.Request
*/
type CommitParams struct {

	/*BranchID*/
	BranchID string
	/*Commit*/
	Commit *models.CommitCreation
	/*RepositoryID*/
	RepositoryID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the commit params
func (o *CommitParams) WithTimeout(timeout time.Duration) *CommitParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the commit params
func (o *CommitParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the commit params
func (o *CommitParams) WithContext(ctx context.Context) *CommitParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the commit params
func (o *CommitParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the commit params
func (o *CommitParams) WithHTTPClient(client *http.Client) *CommitParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the commit params
func (o *CommitParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBranchID adds the branchID to the commit params
func (o *CommitParams) WithBranchID(branchID string) *CommitParams {
	o.SetBranchID(branchID)
	return o
}

// SetBranchID adds the branchId to the commit params
func (o *CommitParams) SetBranchID(branchID string) {
	o.BranchID = branchID
}

// WithCommit adds the commit to the commit params
func (o *CommitParams) WithCommit(commit *models.CommitCreation) *CommitParams {
	o.SetCommit(commit)
	return o
}

// SetCommit adds the commit to the commit params
func (o *CommitParams) SetCommit(commit *models.CommitCreation) {
	o.Commit = commit
}

// WithRepositoryID adds the repositoryID to the commit params
func (o *CommitParams) WithRepositoryID(repositoryID string) *CommitParams {
	o.SetRepositoryID(repositoryID)
	return o
}

// SetRepositoryID adds the repositoryId to the commit params
func (o *CommitParams) SetRepositoryID(repositoryID string) {
	o.RepositoryID = repositoryID
}

// WriteToRequest writes these params to a swagger request
func (o *CommitParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param branchId
	if err := r.SetPathParam("branchId", o.BranchID); err != nil {
		return err
	}

	if o.Commit != nil {
		if err := r.SetBodyParam(o.Commit); err != nil {
			return err
		}
	}

	// path param repositoryId
	if err := r.SetPathParam("repositoryId", o.RepositoryID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
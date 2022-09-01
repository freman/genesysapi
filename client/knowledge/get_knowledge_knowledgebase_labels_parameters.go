// Code generated by go-swagger; DO NOT EDIT.

package knowledge

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

// NewGetKnowledgeKnowledgebaseLabelsParams creates a new GetKnowledgeKnowledgebaseLabelsParams object
// with the default values initialized.
func NewGetKnowledgeKnowledgebaseLabelsParams() *GetKnowledgeKnowledgebaseLabelsParams {
	var ()
	return &GetKnowledgeKnowledgebaseLabelsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetKnowledgeKnowledgebaseLabelsParamsWithTimeout creates a new GetKnowledgeKnowledgebaseLabelsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetKnowledgeKnowledgebaseLabelsParamsWithTimeout(timeout time.Duration) *GetKnowledgeKnowledgebaseLabelsParams {
	var ()
	return &GetKnowledgeKnowledgebaseLabelsParams{

		timeout: timeout,
	}
}

// NewGetKnowledgeKnowledgebaseLabelsParamsWithContext creates a new GetKnowledgeKnowledgebaseLabelsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetKnowledgeKnowledgebaseLabelsParamsWithContext(ctx context.Context) *GetKnowledgeKnowledgebaseLabelsParams {
	var ()
	return &GetKnowledgeKnowledgebaseLabelsParams{

		Context: ctx,
	}
}

// NewGetKnowledgeKnowledgebaseLabelsParamsWithHTTPClient creates a new GetKnowledgeKnowledgebaseLabelsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetKnowledgeKnowledgebaseLabelsParamsWithHTTPClient(client *http.Client) *GetKnowledgeKnowledgebaseLabelsParams {
	var ()
	return &GetKnowledgeKnowledgebaseLabelsParams{
		HTTPClient: client,
	}
}

/*GetKnowledgeKnowledgebaseLabelsParams contains all the parameters to send to the API endpoint
for the get knowledge knowledgebase labels operation typically these are written to a http.Request
*/
type GetKnowledgeKnowledgebaseLabelsParams struct {

	/*After
	  The cursor that points to the end of the set of entities that has been returned.

	*/
	After *string
	/*Before
	  The cursor that points to the start of the set of entities that has been returned.

	*/
	Before *string
	/*IncludeDocumentCount
	  If specified, retrieves the number of documents related to label.

	*/
	IncludeDocumentCount *bool
	/*KnowledgeBaseID
	  Knowledge base ID

	*/
	KnowledgeBaseID string
	/*Name
	  Filter to return the labels that contains the given phrase in the name.

	*/
	Name *string
	/*PageSize
	  Number of entities to return. Maximum of 200.

	*/
	PageSize *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get knowledge knowledgebase labels params
func (o *GetKnowledgeKnowledgebaseLabelsParams) WithTimeout(timeout time.Duration) *GetKnowledgeKnowledgebaseLabelsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get knowledge knowledgebase labels params
func (o *GetKnowledgeKnowledgebaseLabelsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get knowledge knowledgebase labels params
func (o *GetKnowledgeKnowledgebaseLabelsParams) WithContext(ctx context.Context) *GetKnowledgeKnowledgebaseLabelsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get knowledge knowledgebase labels params
func (o *GetKnowledgeKnowledgebaseLabelsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get knowledge knowledgebase labels params
func (o *GetKnowledgeKnowledgebaseLabelsParams) WithHTTPClient(client *http.Client) *GetKnowledgeKnowledgebaseLabelsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get knowledge knowledgebase labels params
func (o *GetKnowledgeKnowledgebaseLabelsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAfter adds the after to the get knowledge knowledgebase labels params
func (o *GetKnowledgeKnowledgebaseLabelsParams) WithAfter(after *string) *GetKnowledgeKnowledgebaseLabelsParams {
	o.SetAfter(after)
	return o
}

// SetAfter adds the after to the get knowledge knowledgebase labels params
func (o *GetKnowledgeKnowledgebaseLabelsParams) SetAfter(after *string) {
	o.After = after
}

// WithBefore adds the before to the get knowledge knowledgebase labels params
func (o *GetKnowledgeKnowledgebaseLabelsParams) WithBefore(before *string) *GetKnowledgeKnowledgebaseLabelsParams {
	o.SetBefore(before)
	return o
}

// SetBefore adds the before to the get knowledge knowledgebase labels params
func (o *GetKnowledgeKnowledgebaseLabelsParams) SetBefore(before *string) {
	o.Before = before
}

// WithIncludeDocumentCount adds the includeDocumentCount to the get knowledge knowledgebase labels params
func (o *GetKnowledgeKnowledgebaseLabelsParams) WithIncludeDocumentCount(includeDocumentCount *bool) *GetKnowledgeKnowledgebaseLabelsParams {
	o.SetIncludeDocumentCount(includeDocumentCount)
	return o
}

// SetIncludeDocumentCount adds the includeDocumentCount to the get knowledge knowledgebase labels params
func (o *GetKnowledgeKnowledgebaseLabelsParams) SetIncludeDocumentCount(includeDocumentCount *bool) {
	o.IncludeDocumentCount = includeDocumentCount
}

// WithKnowledgeBaseID adds the knowledgeBaseID to the get knowledge knowledgebase labels params
func (o *GetKnowledgeKnowledgebaseLabelsParams) WithKnowledgeBaseID(knowledgeBaseID string) *GetKnowledgeKnowledgebaseLabelsParams {
	o.SetKnowledgeBaseID(knowledgeBaseID)
	return o
}

// SetKnowledgeBaseID adds the knowledgeBaseId to the get knowledge knowledgebase labels params
func (o *GetKnowledgeKnowledgebaseLabelsParams) SetKnowledgeBaseID(knowledgeBaseID string) {
	o.KnowledgeBaseID = knowledgeBaseID
}

// WithName adds the name to the get knowledge knowledgebase labels params
func (o *GetKnowledgeKnowledgebaseLabelsParams) WithName(name *string) *GetKnowledgeKnowledgebaseLabelsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get knowledge knowledgebase labels params
func (o *GetKnowledgeKnowledgebaseLabelsParams) SetName(name *string) {
	o.Name = name
}

// WithPageSize adds the pageSize to the get knowledge knowledgebase labels params
func (o *GetKnowledgeKnowledgebaseLabelsParams) WithPageSize(pageSize *string) *GetKnowledgeKnowledgebaseLabelsParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get knowledge knowledgebase labels params
func (o *GetKnowledgeKnowledgebaseLabelsParams) SetPageSize(pageSize *string) {
	o.PageSize = pageSize
}

// WriteToRequest writes these params to a swagger request
func (o *GetKnowledgeKnowledgebaseLabelsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.After != nil {

		// query param after
		var qrAfter string
		if o.After != nil {
			qrAfter = *o.After
		}
		qAfter := qrAfter
		if qAfter != "" {
			if err := r.SetQueryParam("after", qAfter); err != nil {
				return err
			}
		}

	}

	if o.Before != nil {

		// query param before
		var qrBefore string
		if o.Before != nil {
			qrBefore = *o.Before
		}
		qBefore := qrBefore
		if qBefore != "" {
			if err := r.SetQueryParam("before", qBefore); err != nil {
				return err
			}
		}

	}

	if o.IncludeDocumentCount != nil {

		// query param includeDocumentCount
		var qrIncludeDocumentCount bool
		if o.IncludeDocumentCount != nil {
			qrIncludeDocumentCount = *o.IncludeDocumentCount
		}
		qIncludeDocumentCount := swag.FormatBool(qrIncludeDocumentCount)
		if qIncludeDocumentCount != "" {
			if err := r.SetQueryParam("includeDocumentCount", qIncludeDocumentCount); err != nil {
				return err
			}
		}

	}

	// path param knowledgeBaseId
	if err := r.SetPathParam("knowledgeBaseId", o.KnowledgeBaseID); err != nil {
		return err
	}

	if o.Name != nil {

		// query param name
		var qrName string
		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {
			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}

	}

	if o.PageSize != nil {

		// query param pageSize
		var qrPageSize string
		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := qrPageSize
		if qPageSize != "" {
			if err := r.SetQueryParam("pageSize", qPageSize); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

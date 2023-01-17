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

// NewGetKnowledgeKnowledgebaseDocumentParams creates a new GetKnowledgeKnowledgebaseDocumentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetKnowledgeKnowledgebaseDocumentParams() *GetKnowledgeKnowledgebaseDocumentParams {
	return &GetKnowledgeKnowledgebaseDocumentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetKnowledgeKnowledgebaseDocumentParamsWithTimeout creates a new GetKnowledgeKnowledgebaseDocumentParams object
// with the ability to set a timeout on a request.
func NewGetKnowledgeKnowledgebaseDocumentParamsWithTimeout(timeout time.Duration) *GetKnowledgeKnowledgebaseDocumentParams {
	return &GetKnowledgeKnowledgebaseDocumentParams{
		timeout: timeout,
	}
}

// NewGetKnowledgeKnowledgebaseDocumentParamsWithContext creates a new GetKnowledgeKnowledgebaseDocumentParams object
// with the ability to set a context for a request.
func NewGetKnowledgeKnowledgebaseDocumentParamsWithContext(ctx context.Context) *GetKnowledgeKnowledgebaseDocumentParams {
	return &GetKnowledgeKnowledgebaseDocumentParams{
		Context: ctx,
	}
}

// NewGetKnowledgeKnowledgebaseDocumentParamsWithHTTPClient creates a new GetKnowledgeKnowledgebaseDocumentParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetKnowledgeKnowledgebaseDocumentParamsWithHTTPClient(client *http.Client) *GetKnowledgeKnowledgebaseDocumentParams {
	return &GetKnowledgeKnowledgebaseDocumentParams{
		HTTPClient: client,
	}
}

/*
GetKnowledgeKnowledgebaseDocumentParams contains all the parameters to send to the API endpoint

	for the get knowledge knowledgebase document operation.

	Typically these are written to a http.Request.
*/
type GetKnowledgeKnowledgebaseDocumentParams struct {

	/* DocumentID.

	   Document ID.
	*/
	DocumentID string

	/* Expand.

	   The specified entity attributes will be filled. Comma separated values expected. Max No. of variations that can be returned on expand is 20.
	*/
	Expand []string

	/* KnowledgeBaseID.

	   Knowledge base ID.
	*/
	KnowledgeBaseID string

	/* State.

	   "when state is "Draft", draft version of the document is returned,otherwise by default published version is returned in the response.
	*/
	State *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get knowledge knowledgebase document params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetKnowledgeKnowledgebaseDocumentParams) WithDefaults() *GetKnowledgeKnowledgebaseDocumentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get knowledge knowledgebase document params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetKnowledgeKnowledgebaseDocumentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get knowledge knowledgebase document params
func (o *GetKnowledgeKnowledgebaseDocumentParams) WithTimeout(timeout time.Duration) *GetKnowledgeKnowledgebaseDocumentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get knowledge knowledgebase document params
func (o *GetKnowledgeKnowledgebaseDocumentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get knowledge knowledgebase document params
func (o *GetKnowledgeKnowledgebaseDocumentParams) WithContext(ctx context.Context) *GetKnowledgeKnowledgebaseDocumentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get knowledge knowledgebase document params
func (o *GetKnowledgeKnowledgebaseDocumentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get knowledge knowledgebase document params
func (o *GetKnowledgeKnowledgebaseDocumentParams) WithHTTPClient(client *http.Client) *GetKnowledgeKnowledgebaseDocumentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get knowledge knowledgebase document params
func (o *GetKnowledgeKnowledgebaseDocumentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDocumentID adds the documentID to the get knowledge knowledgebase document params
func (o *GetKnowledgeKnowledgebaseDocumentParams) WithDocumentID(documentID string) *GetKnowledgeKnowledgebaseDocumentParams {
	o.SetDocumentID(documentID)
	return o
}

// SetDocumentID adds the documentId to the get knowledge knowledgebase document params
func (o *GetKnowledgeKnowledgebaseDocumentParams) SetDocumentID(documentID string) {
	o.DocumentID = documentID
}

// WithExpand adds the expand to the get knowledge knowledgebase document params
func (o *GetKnowledgeKnowledgebaseDocumentParams) WithExpand(expand []string) *GetKnowledgeKnowledgebaseDocumentParams {
	o.SetExpand(expand)
	return o
}

// SetExpand adds the expand to the get knowledge knowledgebase document params
func (o *GetKnowledgeKnowledgebaseDocumentParams) SetExpand(expand []string) {
	o.Expand = expand
}

// WithKnowledgeBaseID adds the knowledgeBaseID to the get knowledge knowledgebase document params
func (o *GetKnowledgeKnowledgebaseDocumentParams) WithKnowledgeBaseID(knowledgeBaseID string) *GetKnowledgeKnowledgebaseDocumentParams {
	o.SetKnowledgeBaseID(knowledgeBaseID)
	return o
}

// SetKnowledgeBaseID adds the knowledgeBaseId to the get knowledge knowledgebase document params
func (o *GetKnowledgeKnowledgebaseDocumentParams) SetKnowledgeBaseID(knowledgeBaseID string) {
	o.KnowledgeBaseID = knowledgeBaseID
}

// WithState adds the state to the get knowledge knowledgebase document params
func (o *GetKnowledgeKnowledgebaseDocumentParams) WithState(state *string) *GetKnowledgeKnowledgebaseDocumentParams {
	o.SetState(state)
	return o
}

// SetState adds the state to the get knowledge knowledgebase document params
func (o *GetKnowledgeKnowledgebaseDocumentParams) SetState(state *string) {
	o.State = state
}

// WriteToRequest writes these params to a swagger request
func (o *GetKnowledgeKnowledgebaseDocumentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param documentId
	if err := r.SetPathParam("documentId", o.DocumentID); err != nil {
		return err
	}

	if o.Expand != nil {

		// binding items for expand
		joinedExpand := o.bindParamExpand(reg)

		// query array param expand
		if err := r.SetQueryParam("expand", joinedExpand...); err != nil {
			return err
		}
	}

	// path param knowledgeBaseId
	if err := r.SetPathParam("knowledgeBaseId", o.KnowledgeBaseID); err != nil {
		return err
	}

	if o.State != nil {

		// query param state
		var qrState string

		if o.State != nil {
			qrState = *o.State
		}
		qState := qrState
		if qState != "" {

			if err := r.SetQueryParam("state", qState); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetKnowledgeKnowledgebaseDocument binds the parameter expand
func (o *GetKnowledgeKnowledgebaseDocumentParams) bindParamExpand(formats strfmt.Registry) []string {
	expandIR := o.Expand

	var expandIC []string
	for _, expandIIR := range expandIR { // explode []string

		expandIIV := expandIIR // string as string
		expandIC = append(expandIC, expandIIV)
	}

	// items.CollectionFormat: "multi"
	expandIS := swag.JoinByFormat(expandIC, "multi")

	return expandIS
}

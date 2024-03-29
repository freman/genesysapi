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

	"github.com/freman/genesysapi/models"
)

// NewPatchKnowledgeKnowledgebaseUnansweredGroupPhrasegroupParams creates a new PatchKnowledgeKnowledgebaseUnansweredGroupPhrasegroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchKnowledgeKnowledgebaseUnansweredGroupPhrasegroupParams() *PatchKnowledgeKnowledgebaseUnansweredGroupPhrasegroupParams {
	return &PatchKnowledgeKnowledgebaseUnansweredGroupPhrasegroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchKnowledgeKnowledgebaseUnansweredGroupPhrasegroupParamsWithTimeout creates a new PatchKnowledgeKnowledgebaseUnansweredGroupPhrasegroupParams object
// with the ability to set a timeout on a request.
func NewPatchKnowledgeKnowledgebaseUnansweredGroupPhrasegroupParamsWithTimeout(timeout time.Duration) *PatchKnowledgeKnowledgebaseUnansweredGroupPhrasegroupParams {
	return &PatchKnowledgeKnowledgebaseUnansweredGroupPhrasegroupParams{
		timeout: timeout,
	}
}

// NewPatchKnowledgeKnowledgebaseUnansweredGroupPhrasegroupParamsWithContext creates a new PatchKnowledgeKnowledgebaseUnansweredGroupPhrasegroupParams object
// with the ability to set a context for a request.
func NewPatchKnowledgeKnowledgebaseUnansweredGroupPhrasegroupParamsWithContext(ctx context.Context) *PatchKnowledgeKnowledgebaseUnansweredGroupPhrasegroupParams {
	return &PatchKnowledgeKnowledgebaseUnansweredGroupPhrasegroupParams{
		Context: ctx,
	}
}

// NewPatchKnowledgeKnowledgebaseUnansweredGroupPhrasegroupParamsWithHTTPClient creates a new PatchKnowledgeKnowledgebaseUnansweredGroupPhrasegroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchKnowledgeKnowledgebaseUnansweredGroupPhrasegroupParamsWithHTTPClient(client *http.Client) *PatchKnowledgeKnowledgebaseUnansweredGroupPhrasegroupParams {
	return &PatchKnowledgeKnowledgebaseUnansweredGroupPhrasegroupParams{
		HTTPClient: client,
	}
}

/*
PatchKnowledgeKnowledgebaseUnansweredGroupPhrasegroupParams contains all the parameters to send to the API endpoint

	for the patch knowledge knowledgebase unanswered group phrasegroup operation.

	Typically these are written to a http.Request.
*/
type PatchKnowledgeKnowledgebaseUnansweredGroupPhrasegroupParams struct {

	/* Body.

	   Request body of the update unanswered group endpoint.
	*/
	Body *models.UnansweredPhraseGroupPatchRequestBody

	/* GroupID.

	   The ID of the group to be updated.
	*/
	GroupID string

	/* KnowledgeBaseID.

	   Knowledge base ID
	*/
	KnowledgeBaseID string

	/* PhraseGroupID.

	   The ID of the phraseGroup to be updated.
	*/
	PhraseGroupID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch knowledge knowledgebase unanswered group phrasegroup params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchKnowledgeKnowledgebaseUnansweredGroupPhrasegroupParams) WithDefaults() *PatchKnowledgeKnowledgebaseUnansweredGroupPhrasegroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch knowledge knowledgebase unanswered group phrasegroup params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchKnowledgeKnowledgebaseUnansweredGroupPhrasegroupParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch knowledge knowledgebase unanswered group phrasegroup params
func (o *PatchKnowledgeKnowledgebaseUnansweredGroupPhrasegroupParams) WithTimeout(timeout time.Duration) *PatchKnowledgeKnowledgebaseUnansweredGroupPhrasegroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch knowledge knowledgebase unanswered group phrasegroup params
func (o *PatchKnowledgeKnowledgebaseUnansweredGroupPhrasegroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch knowledge knowledgebase unanswered group phrasegroup params
func (o *PatchKnowledgeKnowledgebaseUnansweredGroupPhrasegroupParams) WithContext(ctx context.Context) *PatchKnowledgeKnowledgebaseUnansweredGroupPhrasegroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch knowledge knowledgebase unanswered group phrasegroup params
func (o *PatchKnowledgeKnowledgebaseUnansweredGroupPhrasegroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch knowledge knowledgebase unanswered group phrasegroup params
func (o *PatchKnowledgeKnowledgebaseUnansweredGroupPhrasegroupParams) WithHTTPClient(client *http.Client) *PatchKnowledgeKnowledgebaseUnansweredGroupPhrasegroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch knowledge knowledgebase unanswered group phrasegroup params
func (o *PatchKnowledgeKnowledgebaseUnansweredGroupPhrasegroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch knowledge knowledgebase unanswered group phrasegroup params
func (o *PatchKnowledgeKnowledgebaseUnansweredGroupPhrasegroupParams) WithBody(body *models.UnansweredPhraseGroupPatchRequestBody) *PatchKnowledgeKnowledgebaseUnansweredGroupPhrasegroupParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch knowledge knowledgebase unanswered group phrasegroup params
func (o *PatchKnowledgeKnowledgebaseUnansweredGroupPhrasegroupParams) SetBody(body *models.UnansweredPhraseGroupPatchRequestBody) {
	o.Body = body
}

// WithGroupID adds the groupID to the patch knowledge knowledgebase unanswered group phrasegroup params
func (o *PatchKnowledgeKnowledgebaseUnansweredGroupPhrasegroupParams) WithGroupID(groupID string) *PatchKnowledgeKnowledgebaseUnansweredGroupPhrasegroupParams {
	o.SetGroupID(groupID)
	return o
}

// SetGroupID adds the groupId to the patch knowledge knowledgebase unanswered group phrasegroup params
func (o *PatchKnowledgeKnowledgebaseUnansweredGroupPhrasegroupParams) SetGroupID(groupID string) {
	o.GroupID = groupID
}

// WithKnowledgeBaseID adds the knowledgeBaseID to the patch knowledge knowledgebase unanswered group phrasegroup params
func (o *PatchKnowledgeKnowledgebaseUnansweredGroupPhrasegroupParams) WithKnowledgeBaseID(knowledgeBaseID string) *PatchKnowledgeKnowledgebaseUnansweredGroupPhrasegroupParams {
	o.SetKnowledgeBaseID(knowledgeBaseID)
	return o
}

// SetKnowledgeBaseID adds the knowledgeBaseId to the patch knowledge knowledgebase unanswered group phrasegroup params
func (o *PatchKnowledgeKnowledgebaseUnansweredGroupPhrasegroupParams) SetKnowledgeBaseID(knowledgeBaseID string) {
	o.KnowledgeBaseID = knowledgeBaseID
}

// WithPhraseGroupID adds the phraseGroupID to the patch knowledge knowledgebase unanswered group phrasegroup params
func (o *PatchKnowledgeKnowledgebaseUnansweredGroupPhrasegroupParams) WithPhraseGroupID(phraseGroupID string) *PatchKnowledgeKnowledgebaseUnansweredGroupPhrasegroupParams {
	o.SetPhraseGroupID(phraseGroupID)
	return o
}

// SetPhraseGroupID adds the phraseGroupId to the patch knowledge knowledgebase unanswered group phrasegroup params
func (o *PatchKnowledgeKnowledgebaseUnansweredGroupPhrasegroupParams) SetPhraseGroupID(phraseGroupID string) {
	o.PhraseGroupID = phraseGroupID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchKnowledgeKnowledgebaseUnansweredGroupPhrasegroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param groupId
	if err := r.SetPathParam("groupId", o.GroupID); err != nil {
		return err
	}

	// path param knowledgeBaseId
	if err := r.SetPathParam("knowledgeBaseId", o.KnowledgeBaseID); err != nil {
		return err
	}

	// path param phraseGroupId
	if err := r.SetPathParam("phraseGroupId", o.PhraseGroupID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

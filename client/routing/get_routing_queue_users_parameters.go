// Code generated by go-swagger; DO NOT EDIT.

package routing

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

// NewGetRoutingQueueUsersParams creates a new GetRoutingQueueUsersParams object
// with the default values initialized.
func NewGetRoutingQueueUsersParams() *GetRoutingQueueUsersParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
		sortByDefault     = string("name")
	)
	return &GetRoutingQueueUsersParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		SortBy:     &sortByDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRoutingQueueUsersParamsWithTimeout creates a new GetRoutingQueueUsersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRoutingQueueUsersParamsWithTimeout(timeout time.Duration) *GetRoutingQueueUsersParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
		sortByDefault     = string("name")
	)
	return &GetRoutingQueueUsersParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		SortBy:     &sortByDefault,

		timeout: timeout,
	}
}

// NewGetRoutingQueueUsersParamsWithContext creates a new GetRoutingQueueUsersParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRoutingQueueUsersParamsWithContext(ctx context.Context) *GetRoutingQueueUsersParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
		sortByDefault     = string("name")
	)
	return &GetRoutingQueueUsersParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		SortBy:     &sortByDefault,

		Context: ctx,
	}
}

// NewGetRoutingQueueUsersParamsWithHTTPClient creates a new GetRoutingQueueUsersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRoutingQueueUsersParamsWithHTTPClient(client *http.Client) *GetRoutingQueueUsersParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
		sortByDefault     = string("name")
	)
	return &GetRoutingQueueUsersParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		SortBy:     &sortByDefault,
		HTTPClient: client,
	}
}

/*GetRoutingQueueUsersParams contains all the parameters to send to the API endpoint
for the get routing queue users operation typically these are written to a http.Request
*/
type GetRoutingQueueUsersParams struct {

	/*Expand
	  Which fields, if any, to expand.

	*/
	Expand []string
	/*Joined
	  Filter by joined status

	*/
	Joined *bool
	/*Languages
	  Filter by language

	*/
	Languages []string
	/*Name
	  Filter by queue member name

	*/
	Name *string
	/*PageNumber
	  Page number

	*/
	PageNumber *int32
	/*PageSize
	  Page size

	*/
	PageSize *int32
	/*Presence
	  Filter by presence

	*/
	Presence []string
	/*ProfileSkills
	  Filter by profile skill

	*/
	ProfileSkills []string
	/*QueueID
	  Queue ID

	*/
	QueueID string
	/*RoutingStatus
	  Filter by routing status

	*/
	RoutingStatus []string
	/*Skills
	  Filter by skill

	*/
	Skills []string
	/*SortBy
	  Sort by

	*/
	SortBy *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get routing queue users params
func (o *GetRoutingQueueUsersParams) WithTimeout(timeout time.Duration) *GetRoutingQueueUsersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get routing queue users params
func (o *GetRoutingQueueUsersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get routing queue users params
func (o *GetRoutingQueueUsersParams) WithContext(ctx context.Context) *GetRoutingQueueUsersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get routing queue users params
func (o *GetRoutingQueueUsersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get routing queue users params
func (o *GetRoutingQueueUsersParams) WithHTTPClient(client *http.Client) *GetRoutingQueueUsersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get routing queue users params
func (o *GetRoutingQueueUsersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExpand adds the expand to the get routing queue users params
func (o *GetRoutingQueueUsersParams) WithExpand(expand []string) *GetRoutingQueueUsersParams {
	o.SetExpand(expand)
	return o
}

// SetExpand adds the expand to the get routing queue users params
func (o *GetRoutingQueueUsersParams) SetExpand(expand []string) {
	o.Expand = expand
}

// WithJoined adds the joined to the get routing queue users params
func (o *GetRoutingQueueUsersParams) WithJoined(joined *bool) *GetRoutingQueueUsersParams {
	o.SetJoined(joined)
	return o
}

// SetJoined adds the joined to the get routing queue users params
func (o *GetRoutingQueueUsersParams) SetJoined(joined *bool) {
	o.Joined = joined
}

// WithLanguages adds the languages to the get routing queue users params
func (o *GetRoutingQueueUsersParams) WithLanguages(languages []string) *GetRoutingQueueUsersParams {
	o.SetLanguages(languages)
	return o
}

// SetLanguages adds the languages to the get routing queue users params
func (o *GetRoutingQueueUsersParams) SetLanguages(languages []string) {
	o.Languages = languages
}

// WithName adds the name to the get routing queue users params
func (o *GetRoutingQueueUsersParams) WithName(name *string) *GetRoutingQueueUsersParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get routing queue users params
func (o *GetRoutingQueueUsersParams) SetName(name *string) {
	o.Name = name
}

// WithPageNumber adds the pageNumber to the get routing queue users params
func (o *GetRoutingQueueUsersParams) WithPageNumber(pageNumber *int32) *GetRoutingQueueUsersParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get routing queue users params
func (o *GetRoutingQueueUsersParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get routing queue users params
func (o *GetRoutingQueueUsersParams) WithPageSize(pageSize *int32) *GetRoutingQueueUsersParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get routing queue users params
func (o *GetRoutingQueueUsersParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithPresence adds the presence to the get routing queue users params
func (o *GetRoutingQueueUsersParams) WithPresence(presence []string) *GetRoutingQueueUsersParams {
	o.SetPresence(presence)
	return o
}

// SetPresence adds the presence to the get routing queue users params
func (o *GetRoutingQueueUsersParams) SetPresence(presence []string) {
	o.Presence = presence
}

// WithProfileSkills adds the profileSkills to the get routing queue users params
func (o *GetRoutingQueueUsersParams) WithProfileSkills(profileSkills []string) *GetRoutingQueueUsersParams {
	o.SetProfileSkills(profileSkills)
	return o
}

// SetProfileSkills adds the profileSkills to the get routing queue users params
func (o *GetRoutingQueueUsersParams) SetProfileSkills(profileSkills []string) {
	o.ProfileSkills = profileSkills
}

// WithQueueID adds the queueID to the get routing queue users params
func (o *GetRoutingQueueUsersParams) WithQueueID(queueID string) *GetRoutingQueueUsersParams {
	o.SetQueueID(queueID)
	return o
}

// SetQueueID adds the queueId to the get routing queue users params
func (o *GetRoutingQueueUsersParams) SetQueueID(queueID string) {
	o.QueueID = queueID
}

// WithRoutingStatus adds the routingStatus to the get routing queue users params
func (o *GetRoutingQueueUsersParams) WithRoutingStatus(routingStatus []string) *GetRoutingQueueUsersParams {
	o.SetRoutingStatus(routingStatus)
	return o
}

// SetRoutingStatus adds the routingStatus to the get routing queue users params
func (o *GetRoutingQueueUsersParams) SetRoutingStatus(routingStatus []string) {
	o.RoutingStatus = routingStatus
}

// WithSkills adds the skills to the get routing queue users params
func (o *GetRoutingQueueUsersParams) WithSkills(skills []string) *GetRoutingQueueUsersParams {
	o.SetSkills(skills)
	return o
}

// SetSkills adds the skills to the get routing queue users params
func (o *GetRoutingQueueUsersParams) SetSkills(skills []string) {
	o.Skills = skills
}

// WithSortBy adds the sortBy to the get routing queue users params
func (o *GetRoutingQueueUsersParams) WithSortBy(sortBy *string) *GetRoutingQueueUsersParams {
	o.SetSortBy(sortBy)
	return o
}

// SetSortBy adds the sortBy to the get routing queue users params
func (o *GetRoutingQueueUsersParams) SetSortBy(sortBy *string) {
	o.SortBy = sortBy
}

// WriteToRequest writes these params to a swagger request
func (o *GetRoutingQueueUsersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	valuesExpand := o.Expand

	joinedExpand := swag.JoinByFormat(valuesExpand, "multi")
	// query array param expand
	if err := r.SetQueryParam("expand", joinedExpand...); err != nil {
		return err
	}

	if o.Joined != nil {

		// query param joined
		var qrJoined bool
		if o.Joined != nil {
			qrJoined = *o.Joined
		}
		qJoined := swag.FormatBool(qrJoined)
		if qJoined != "" {
			if err := r.SetQueryParam("joined", qJoined); err != nil {
				return err
			}
		}

	}

	valuesLanguages := o.Languages

	joinedLanguages := swag.JoinByFormat(valuesLanguages, "multi")
	// query array param languages
	if err := r.SetQueryParam("languages", joinedLanguages...); err != nil {
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

	if o.PageNumber != nil {

		// query param pageNumber
		var qrPageNumber int32
		if o.PageNumber != nil {
			qrPageNumber = *o.PageNumber
		}
		qPageNumber := swag.FormatInt32(qrPageNumber)
		if qPageNumber != "" {
			if err := r.SetQueryParam("pageNumber", qPageNumber); err != nil {
				return err
			}
		}

	}

	if o.PageSize != nil {

		// query param pageSize
		var qrPageSize int32
		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatInt32(qrPageSize)
		if qPageSize != "" {
			if err := r.SetQueryParam("pageSize", qPageSize); err != nil {
				return err
			}
		}

	}

	valuesPresence := o.Presence

	joinedPresence := swag.JoinByFormat(valuesPresence, "multi")
	// query array param presence
	if err := r.SetQueryParam("presence", joinedPresence...); err != nil {
		return err
	}

	valuesProfileSkills := o.ProfileSkills

	joinedProfileSkills := swag.JoinByFormat(valuesProfileSkills, "multi")
	// query array param profileSkills
	if err := r.SetQueryParam("profileSkills", joinedProfileSkills...); err != nil {
		return err
	}

	// path param queueId
	if err := r.SetPathParam("queueId", o.QueueID); err != nil {
		return err
	}

	valuesRoutingStatus := o.RoutingStatus

	joinedRoutingStatus := swag.JoinByFormat(valuesRoutingStatus, "multi")
	// query array param routingStatus
	if err := r.SetQueryParam("routingStatus", joinedRoutingStatus...); err != nil {
		return err
	}

	valuesSkills := o.Skills

	joinedSkills := swag.JoinByFormat(valuesSkills, "multi")
	// query array param skills
	if err := r.SetQueryParam("skills", joinedSkills...); err != nil {
		return err
	}

	if o.SortBy != nil {

		// query param sortBy
		var qrSortBy string
		if o.SortBy != nil {
			qrSortBy = *o.SortBy
		}
		qSortBy := qrSortBy
		if qSortBy != "" {
			if err := r.SetQueryParam("sortBy", qSortBy); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
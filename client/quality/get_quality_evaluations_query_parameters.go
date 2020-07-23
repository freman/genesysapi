// Code generated by go-swagger; DO NOT EDIT.

package quality

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

// NewGetQualityEvaluationsQueryParams creates a new GetQualityEvaluationsQueryParams object
// with the default values initialized.
func NewGetQualityEvaluationsQueryParams() *GetQualityEvaluationsQueryParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
	)
	return &GetQualityEvaluationsQueryParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetQualityEvaluationsQueryParamsWithTimeout creates a new GetQualityEvaluationsQueryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetQualityEvaluationsQueryParamsWithTimeout(timeout time.Duration) *GetQualityEvaluationsQueryParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
	)
	return &GetQualityEvaluationsQueryParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,

		timeout: timeout,
	}
}

// NewGetQualityEvaluationsQueryParamsWithContext creates a new GetQualityEvaluationsQueryParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetQualityEvaluationsQueryParamsWithContext(ctx context.Context) *GetQualityEvaluationsQueryParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
	)
	return &GetQualityEvaluationsQueryParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,

		Context: ctx,
	}
}

// NewGetQualityEvaluationsQueryParamsWithHTTPClient creates a new GetQualityEvaluationsQueryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetQualityEvaluationsQueryParamsWithHTTPClient(client *http.Client) *GetQualityEvaluationsQueryParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
	)
	return &GetQualityEvaluationsQueryParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		HTTPClient: client,
	}
}

/*GetQualityEvaluationsQueryParams contains all the parameters to send to the API endpoint
for the get quality evaluations query operation typically these are written to a http.Request
*/
type GetQualityEvaluationsQueryParams struct {

	/*AgentHasRead
	  agent has the evaluation

	*/
	AgentHasRead *bool
	/*AgentUserID
	  user id of the agent

	*/
	AgentUserID *string
	/*ConversationID
	  conversationId specified

	*/
	ConversationID *string
	/*EndTime
	  end time of the evaluation query

	*/
	EndTime *string
	/*EvaluationState*/
	EvaluationState []string
	/*EvaluatorUserID
	  evaluator user id

	*/
	EvaluatorUserID *string
	/*Expand
	  variable name requested by expand list

	*/
	Expand []string
	/*ExpandAnswerTotalScores
	  get the total scores for evaluations

	*/
	ExpandAnswerTotalScores *bool
	/*IsReleased
	  the evaluation has been released

	*/
	IsReleased *bool
	/*Maximum
	  maximum

	*/
	Maximum *int32
	/*NextPage
	  next page token

	*/
	NextPage *string
	/*PageNumber
	  The page number requested

	*/
	PageNumber *int32
	/*PageSize
	  The total page size requested

	*/
	PageSize *int32
	/*PreviousPage
	  Previous page token

	*/
	PreviousPage *string
	/*QueueID
	  queue id

	*/
	QueueID *string
	/*SortBy
	  variable name requested to sort by

	*/
	SortBy *string
	/*SortOrder
	  sort order options for agentUserId or evaluatorUserId query. Valid options are 'a', 'asc', 'ascending', 'd', 'desc', 'descending'

	*/
	SortOrder *string
	/*StartTime
	  start time of the evaluation query

	*/
	StartTime *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get quality evaluations query params
func (o *GetQualityEvaluationsQueryParams) WithTimeout(timeout time.Duration) *GetQualityEvaluationsQueryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get quality evaluations query params
func (o *GetQualityEvaluationsQueryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get quality evaluations query params
func (o *GetQualityEvaluationsQueryParams) WithContext(ctx context.Context) *GetQualityEvaluationsQueryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get quality evaluations query params
func (o *GetQualityEvaluationsQueryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get quality evaluations query params
func (o *GetQualityEvaluationsQueryParams) WithHTTPClient(client *http.Client) *GetQualityEvaluationsQueryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get quality evaluations query params
func (o *GetQualityEvaluationsQueryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAgentHasRead adds the agentHasRead to the get quality evaluations query params
func (o *GetQualityEvaluationsQueryParams) WithAgentHasRead(agentHasRead *bool) *GetQualityEvaluationsQueryParams {
	o.SetAgentHasRead(agentHasRead)
	return o
}

// SetAgentHasRead adds the agentHasRead to the get quality evaluations query params
func (o *GetQualityEvaluationsQueryParams) SetAgentHasRead(agentHasRead *bool) {
	o.AgentHasRead = agentHasRead
}

// WithAgentUserID adds the agentUserID to the get quality evaluations query params
func (o *GetQualityEvaluationsQueryParams) WithAgentUserID(agentUserID *string) *GetQualityEvaluationsQueryParams {
	o.SetAgentUserID(agentUserID)
	return o
}

// SetAgentUserID adds the agentUserId to the get quality evaluations query params
func (o *GetQualityEvaluationsQueryParams) SetAgentUserID(agentUserID *string) {
	o.AgentUserID = agentUserID
}

// WithConversationID adds the conversationID to the get quality evaluations query params
func (o *GetQualityEvaluationsQueryParams) WithConversationID(conversationID *string) *GetQualityEvaluationsQueryParams {
	o.SetConversationID(conversationID)
	return o
}

// SetConversationID adds the conversationId to the get quality evaluations query params
func (o *GetQualityEvaluationsQueryParams) SetConversationID(conversationID *string) {
	o.ConversationID = conversationID
}

// WithEndTime adds the endTime to the get quality evaluations query params
func (o *GetQualityEvaluationsQueryParams) WithEndTime(endTime *string) *GetQualityEvaluationsQueryParams {
	o.SetEndTime(endTime)
	return o
}

// SetEndTime adds the endTime to the get quality evaluations query params
func (o *GetQualityEvaluationsQueryParams) SetEndTime(endTime *string) {
	o.EndTime = endTime
}

// WithEvaluationState adds the evaluationState to the get quality evaluations query params
func (o *GetQualityEvaluationsQueryParams) WithEvaluationState(evaluationState []string) *GetQualityEvaluationsQueryParams {
	o.SetEvaluationState(evaluationState)
	return o
}

// SetEvaluationState adds the evaluationState to the get quality evaluations query params
func (o *GetQualityEvaluationsQueryParams) SetEvaluationState(evaluationState []string) {
	o.EvaluationState = evaluationState
}

// WithEvaluatorUserID adds the evaluatorUserID to the get quality evaluations query params
func (o *GetQualityEvaluationsQueryParams) WithEvaluatorUserID(evaluatorUserID *string) *GetQualityEvaluationsQueryParams {
	o.SetEvaluatorUserID(evaluatorUserID)
	return o
}

// SetEvaluatorUserID adds the evaluatorUserId to the get quality evaluations query params
func (o *GetQualityEvaluationsQueryParams) SetEvaluatorUserID(evaluatorUserID *string) {
	o.EvaluatorUserID = evaluatorUserID
}

// WithExpand adds the expand to the get quality evaluations query params
func (o *GetQualityEvaluationsQueryParams) WithExpand(expand []string) *GetQualityEvaluationsQueryParams {
	o.SetExpand(expand)
	return o
}

// SetExpand adds the expand to the get quality evaluations query params
func (o *GetQualityEvaluationsQueryParams) SetExpand(expand []string) {
	o.Expand = expand
}

// WithExpandAnswerTotalScores adds the expandAnswerTotalScores to the get quality evaluations query params
func (o *GetQualityEvaluationsQueryParams) WithExpandAnswerTotalScores(expandAnswerTotalScores *bool) *GetQualityEvaluationsQueryParams {
	o.SetExpandAnswerTotalScores(expandAnswerTotalScores)
	return o
}

// SetExpandAnswerTotalScores adds the expandAnswerTotalScores to the get quality evaluations query params
func (o *GetQualityEvaluationsQueryParams) SetExpandAnswerTotalScores(expandAnswerTotalScores *bool) {
	o.ExpandAnswerTotalScores = expandAnswerTotalScores
}

// WithIsReleased adds the isReleased to the get quality evaluations query params
func (o *GetQualityEvaluationsQueryParams) WithIsReleased(isReleased *bool) *GetQualityEvaluationsQueryParams {
	o.SetIsReleased(isReleased)
	return o
}

// SetIsReleased adds the isReleased to the get quality evaluations query params
func (o *GetQualityEvaluationsQueryParams) SetIsReleased(isReleased *bool) {
	o.IsReleased = isReleased
}

// WithMaximum adds the maximum to the get quality evaluations query params
func (o *GetQualityEvaluationsQueryParams) WithMaximum(maximum *int32) *GetQualityEvaluationsQueryParams {
	o.SetMaximum(maximum)
	return o
}

// SetMaximum adds the maximum to the get quality evaluations query params
func (o *GetQualityEvaluationsQueryParams) SetMaximum(maximum *int32) {
	o.Maximum = maximum
}

// WithNextPage adds the nextPage to the get quality evaluations query params
func (o *GetQualityEvaluationsQueryParams) WithNextPage(nextPage *string) *GetQualityEvaluationsQueryParams {
	o.SetNextPage(nextPage)
	return o
}

// SetNextPage adds the nextPage to the get quality evaluations query params
func (o *GetQualityEvaluationsQueryParams) SetNextPage(nextPage *string) {
	o.NextPage = nextPage
}

// WithPageNumber adds the pageNumber to the get quality evaluations query params
func (o *GetQualityEvaluationsQueryParams) WithPageNumber(pageNumber *int32) *GetQualityEvaluationsQueryParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get quality evaluations query params
func (o *GetQualityEvaluationsQueryParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get quality evaluations query params
func (o *GetQualityEvaluationsQueryParams) WithPageSize(pageSize *int32) *GetQualityEvaluationsQueryParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get quality evaluations query params
func (o *GetQualityEvaluationsQueryParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithPreviousPage adds the previousPage to the get quality evaluations query params
func (o *GetQualityEvaluationsQueryParams) WithPreviousPage(previousPage *string) *GetQualityEvaluationsQueryParams {
	o.SetPreviousPage(previousPage)
	return o
}

// SetPreviousPage adds the previousPage to the get quality evaluations query params
func (o *GetQualityEvaluationsQueryParams) SetPreviousPage(previousPage *string) {
	o.PreviousPage = previousPage
}

// WithQueueID adds the queueID to the get quality evaluations query params
func (o *GetQualityEvaluationsQueryParams) WithQueueID(queueID *string) *GetQualityEvaluationsQueryParams {
	o.SetQueueID(queueID)
	return o
}

// SetQueueID adds the queueId to the get quality evaluations query params
func (o *GetQualityEvaluationsQueryParams) SetQueueID(queueID *string) {
	o.QueueID = queueID
}

// WithSortBy adds the sortBy to the get quality evaluations query params
func (o *GetQualityEvaluationsQueryParams) WithSortBy(sortBy *string) *GetQualityEvaluationsQueryParams {
	o.SetSortBy(sortBy)
	return o
}

// SetSortBy adds the sortBy to the get quality evaluations query params
func (o *GetQualityEvaluationsQueryParams) SetSortBy(sortBy *string) {
	o.SortBy = sortBy
}

// WithSortOrder adds the sortOrder to the get quality evaluations query params
func (o *GetQualityEvaluationsQueryParams) WithSortOrder(sortOrder *string) *GetQualityEvaluationsQueryParams {
	o.SetSortOrder(sortOrder)
	return o
}

// SetSortOrder adds the sortOrder to the get quality evaluations query params
func (o *GetQualityEvaluationsQueryParams) SetSortOrder(sortOrder *string) {
	o.SortOrder = sortOrder
}

// WithStartTime adds the startTime to the get quality evaluations query params
func (o *GetQualityEvaluationsQueryParams) WithStartTime(startTime *string) *GetQualityEvaluationsQueryParams {
	o.SetStartTime(startTime)
	return o
}

// SetStartTime adds the startTime to the get quality evaluations query params
func (o *GetQualityEvaluationsQueryParams) SetStartTime(startTime *string) {
	o.StartTime = startTime
}

// WriteToRequest writes these params to a swagger request
func (o *GetQualityEvaluationsQueryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AgentHasRead != nil {

		// query param agentHasRead
		var qrAgentHasRead bool
		if o.AgentHasRead != nil {
			qrAgentHasRead = *o.AgentHasRead
		}
		qAgentHasRead := swag.FormatBool(qrAgentHasRead)
		if qAgentHasRead != "" {
			if err := r.SetQueryParam("agentHasRead", qAgentHasRead); err != nil {
				return err
			}
		}

	}

	if o.AgentUserID != nil {

		// query param agentUserId
		var qrAgentUserID string
		if o.AgentUserID != nil {
			qrAgentUserID = *o.AgentUserID
		}
		qAgentUserID := qrAgentUserID
		if qAgentUserID != "" {
			if err := r.SetQueryParam("agentUserId", qAgentUserID); err != nil {
				return err
			}
		}

	}

	if o.ConversationID != nil {

		// query param conversationId
		var qrConversationID string
		if o.ConversationID != nil {
			qrConversationID = *o.ConversationID
		}
		qConversationID := qrConversationID
		if qConversationID != "" {
			if err := r.SetQueryParam("conversationId", qConversationID); err != nil {
				return err
			}
		}

	}

	if o.EndTime != nil {

		// query param endTime
		var qrEndTime string
		if o.EndTime != nil {
			qrEndTime = *o.EndTime
		}
		qEndTime := qrEndTime
		if qEndTime != "" {
			if err := r.SetQueryParam("endTime", qEndTime); err != nil {
				return err
			}
		}

	}

	valuesEvaluationState := o.EvaluationState

	joinedEvaluationState := swag.JoinByFormat(valuesEvaluationState, "multi")
	// query array param evaluationState
	if err := r.SetQueryParam("evaluationState", joinedEvaluationState...); err != nil {
		return err
	}

	if o.EvaluatorUserID != nil {

		// query param evaluatorUserId
		var qrEvaluatorUserID string
		if o.EvaluatorUserID != nil {
			qrEvaluatorUserID = *o.EvaluatorUserID
		}
		qEvaluatorUserID := qrEvaluatorUserID
		if qEvaluatorUserID != "" {
			if err := r.SetQueryParam("evaluatorUserId", qEvaluatorUserID); err != nil {
				return err
			}
		}

	}

	valuesExpand := o.Expand

	joinedExpand := swag.JoinByFormat(valuesExpand, "multi")
	// query array param expand
	if err := r.SetQueryParam("expand", joinedExpand...); err != nil {
		return err
	}

	if o.ExpandAnswerTotalScores != nil {

		// query param expandAnswerTotalScores
		var qrExpandAnswerTotalScores bool
		if o.ExpandAnswerTotalScores != nil {
			qrExpandAnswerTotalScores = *o.ExpandAnswerTotalScores
		}
		qExpandAnswerTotalScores := swag.FormatBool(qrExpandAnswerTotalScores)
		if qExpandAnswerTotalScores != "" {
			if err := r.SetQueryParam("expandAnswerTotalScores", qExpandAnswerTotalScores); err != nil {
				return err
			}
		}

	}

	if o.IsReleased != nil {

		// query param isReleased
		var qrIsReleased bool
		if o.IsReleased != nil {
			qrIsReleased = *o.IsReleased
		}
		qIsReleased := swag.FormatBool(qrIsReleased)
		if qIsReleased != "" {
			if err := r.SetQueryParam("isReleased", qIsReleased); err != nil {
				return err
			}
		}

	}

	if o.Maximum != nil {

		// query param maximum
		var qrMaximum int32
		if o.Maximum != nil {
			qrMaximum = *o.Maximum
		}
		qMaximum := swag.FormatInt32(qrMaximum)
		if qMaximum != "" {
			if err := r.SetQueryParam("maximum", qMaximum); err != nil {
				return err
			}
		}

	}

	if o.NextPage != nil {

		// query param nextPage
		var qrNextPage string
		if o.NextPage != nil {
			qrNextPage = *o.NextPage
		}
		qNextPage := qrNextPage
		if qNextPage != "" {
			if err := r.SetQueryParam("nextPage", qNextPage); err != nil {
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

	if o.PreviousPage != nil {

		// query param previousPage
		var qrPreviousPage string
		if o.PreviousPage != nil {
			qrPreviousPage = *o.PreviousPage
		}
		qPreviousPage := qrPreviousPage
		if qPreviousPage != "" {
			if err := r.SetQueryParam("previousPage", qPreviousPage); err != nil {
				return err
			}
		}

	}

	if o.QueueID != nil {

		// query param queueId
		var qrQueueID string
		if o.QueueID != nil {
			qrQueueID = *o.QueueID
		}
		qQueueID := qrQueueID
		if qQueueID != "" {
			if err := r.SetQueryParam("queueId", qQueueID); err != nil {
				return err
			}
		}

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

	if o.SortOrder != nil {

		// query param sortOrder
		var qrSortOrder string
		if o.SortOrder != nil {
			qrSortOrder = *o.SortOrder
		}
		qSortOrder := qrSortOrder
		if qSortOrder != "" {
			if err := r.SetQueryParam("sortOrder", qSortOrder); err != nil {
				return err
			}
		}

	}

	if o.StartTime != nil {

		// query param startTime
		var qrStartTime string
		if o.StartTime != nil {
			qrStartTime = *o.StartTime
		}
		qStartTime := qrStartTime
		if qStartTime != "" {
			if err := r.SetQueryParam("startTime", qStartTime); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

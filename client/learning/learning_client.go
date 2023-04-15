// Code generated by go-swagger; DO NOT EDIT.

package learning

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

//go:generate mockery -name API -inpkg

// API is the interface of the learning client
type API interface {
	/*
	   DeleteLearningAssignment deletes a learning assignment
	*/
	DeleteLearningAssignment(ctx context.Context, params *DeleteLearningAssignmentParams) (*DeleteLearningAssignmentNoContent, error)
	/*
	   DeleteLearningModule deletes a learning module
	   This will delete a learning module if it is unpublished or it will delete a published and archived learning module
	*/
	DeleteLearningModule(ctx context.Context, params *DeleteLearningModuleParams) (*DeleteLearningModuleNoContent, error)
	/*
	   GetLearningAssignment gets learning assignment
	   Permission not required if you are the assigned user of the learning assignment
	*/
	GetLearningAssignment(ctx context.Context, params *GetLearningAssignmentParams) (*GetLearningAssignmentOK, error)
	/*
	   GetLearningAssignments lists of learning module assignments
	   Either moduleId or user value is required
	*/
	GetLearningAssignments(ctx context.Context, params *GetLearningAssignmentsParams) (*GetLearningAssignmentsOK, error)
	/*
	   GetLearningAssignmentsMe lists of learning assignments assigned to current user
	*/
	GetLearningAssignmentsMe(ctx context.Context, params *GetLearningAssignmentsMeParams) (*GetLearningAssignmentsMeOK, error)
	/*
	   GetLearningModule gets a learning module
	*/
	GetLearningModule(ctx context.Context, params *GetLearningModuleParams) (*GetLearningModuleOK, error)
	/*
	   GetLearningModuleJob gets a specific learning module job status
	*/
	GetLearningModuleJob(ctx context.Context, params *GetLearningModuleJobParams) (*GetLearningModuleJobOK, error)
	/*
	   GetLearningModuleRule gets a learning module rule
	*/
	GetLearningModuleRule(ctx context.Context, params *GetLearningModuleRuleParams) (*GetLearningModuleRuleOK, error)
	/*
	   GetLearningModuleVersion gets specific version of a published module
	*/
	GetLearningModuleVersion(ctx context.Context, params *GetLearningModuleVersionParams) (*GetLearningModuleVersionOK, error)
	/*
	   GetLearningModules gets all learning modules of an organization
	*/
	GetLearningModules(ctx context.Context, params *GetLearningModulesParams) (*GetLearningModulesOK, error)
	/*
	   GetLearningModulesAssignments gets all learning modules of an organization including assignments for a specific user
	*/
	GetLearningModulesAssignments(ctx context.Context, params *GetLearningModulesAssignmentsParams) (*GetLearningModulesAssignmentsOK, error)
	/*
	   GetLearningModulesCoverartCoverArtID gets a specific learning module cover art using ID
	*/
	GetLearningModulesCoverartCoverArtID(ctx context.Context, params *GetLearningModulesCoverartCoverArtIDParams) (*GetLearningModulesCoverartCoverArtIDOK, error)
	/*
	   PatchLearningAssignment updates learning assignment
	*/
	PatchLearningAssignment(ctx context.Context, params *PatchLearningAssignmentParams) (*PatchLearningAssignmentOK, error)
	/*
	   PatchLearningAssignmentReschedule reschedules learning assignment
	*/
	PatchLearningAssignmentReschedule(ctx context.Context, params *PatchLearningAssignmentRescheduleParams) (*PatchLearningAssignmentRescheduleOK, error)
	/*
	   PostLearningAssessmentsScoring scores learning assessment for preview
	*/
	PostLearningAssessmentsScoring(ctx context.Context, params *PostLearningAssessmentsScoringParams) (*PostLearningAssessmentsScoringOK, error)
	/*
	   PostLearningAssignmentReassign reassigns learning assignment
	   This will reassign the state of the assignment to 'Assigned' and update the assignment to the latest version of the module
	*/
	PostLearningAssignmentReassign(ctx context.Context, params *PostLearningAssignmentReassignParams) (*PostLearningAssignmentReassignOK, error)
	/*
	   PostLearningAssignmentReset resets learning assignment
	   This will reset the state of the assignment to 'Assigned' and remove the version of Learning module associated with the assignment
	*/
	PostLearningAssignmentReset(ctx context.Context, params *PostLearningAssignmentResetParams) (*PostLearningAssignmentResetOK, error)
	/*
	   PostLearningAssignments creates learning assignment
	*/
	PostLearningAssignments(ctx context.Context, params *PostLearningAssignmentsParams) (*PostLearningAssignmentsOK, error)
	/*
	   PostLearningAssignmentsAggregatesQuery retrieves aggregated assignment data
	*/
	PostLearningAssignmentsAggregatesQuery(ctx context.Context, params *PostLearningAssignmentsAggregatesQueryParams) (*PostLearningAssignmentsAggregatesQueryOK, error)
	/*
	   PostLearningAssignmentsBulkadd adds multiple learning assignments
	*/
	PostLearningAssignmentsBulkadd(ctx context.Context, params *PostLearningAssignmentsBulkaddParams) (*PostLearningAssignmentsBulkaddOK, error)
	/*
	   PostLearningAssignmentsBulkremove removes multiple learning assignments
	*/
	PostLearningAssignmentsBulkremove(ctx context.Context, params *PostLearningAssignmentsBulkremoveParams) (*PostLearningAssignmentsBulkremoveOK, error)
	/*
	   PostLearningModuleJobs starts a specified operation on learning module
	   This will initiate operation specified in the request body for a learning module
	*/
	PostLearningModuleJobs(ctx context.Context, params *PostLearningModuleJobsParams) (*PostLearningModuleJobsAccepted, error)
	/*
	   PostLearningModulePublish publishes a learning module
	*/
	PostLearningModulePublish(ctx context.Context, params *PostLearningModulePublishParams) (*PostLearningModulePublishOK, error)
	/*
	   PostLearningModules creates a new learning module
	   This will create a new unpublished learning module with the specified fields.
	*/
	PostLearningModules(ctx context.Context, params *PostLearningModulesParams) (*PostLearningModulesOK, error)
	/*
	   PostLearningRulesQuery gets users for learning module rule
	   This will get the users who matches the given rule.
	*/
	PostLearningRulesQuery(ctx context.Context, params *PostLearningRulesQueryParams) (*PostLearningRulesQueryOK, error)
	/*
	   PostLearningScheduleslotsQuery gets list of possible slots where a learning activity can be scheduled
	*/
	PostLearningScheduleslotsQuery(ctx context.Context, params *PostLearningScheduleslotsQueryParams) (*PostLearningScheduleslotsQueryOK, error)
	/*
	   PutLearningModule updates a learning module
	   This will update the name, description, completion time in days and inform steps for a learning module
	*/
	PutLearningModule(ctx context.Context, params *PutLearningModuleParams) (*PutLearningModuleOK, error)
	/*
	   PutLearningModuleRule updates a learning module rule
	   This will update a learning module rule with the specified fields.
	*/
	PutLearningModuleRule(ctx context.Context, params *PutLearningModuleRuleParams) (*PutLearningModuleRuleOK, error)
}

// New creates a new learning API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, authInfo runtime.ClientAuthInfoWriter) *Client {
	return &Client{
		transport: transport,
		formats:   formats,
		authInfo:  authInfo,
	}
}

/*
Client for learning API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	authInfo  runtime.ClientAuthInfoWriter
}

/*
DeleteLearningAssignment deletes a learning assignment
*/
func (a *Client) DeleteLearningAssignment(ctx context.Context, params *DeleteLearningAssignmentParams) (*DeleteLearningAssignmentNoContent, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteLearningAssignment",
		Method:             "DELETE",
		PathPattern:        "/api/v2/learning/assignments/{assignmentId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteLearningAssignmentReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteLearningAssignmentNoContent), nil

}

/*
DeleteLearningModule deletes a learning module

This will delete a learning module if it is unpublished or it will delete a published and archived learning module
*/
func (a *Client) DeleteLearningModule(ctx context.Context, params *DeleteLearningModuleParams) (*DeleteLearningModuleNoContent, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteLearningModule",
		Method:             "DELETE",
		PathPattern:        "/api/v2/learning/modules/{moduleId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteLearningModuleReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteLearningModuleNoContent), nil

}

/*
GetLearningAssignment gets learning assignment

Permission not required if you are the assigned user of the learning assignment
*/
func (a *Client) GetLearningAssignment(ctx context.Context, params *GetLearningAssignmentParams) (*GetLearningAssignmentOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getLearningAssignment",
		Method:             "GET",
		PathPattern:        "/api/v2/learning/assignments/{assignmentId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetLearningAssignmentReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetLearningAssignmentOK), nil

}

/*
GetLearningAssignments lists of learning module assignments

Either moduleId or user value is required
*/
func (a *Client) GetLearningAssignments(ctx context.Context, params *GetLearningAssignmentsParams) (*GetLearningAssignmentsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getLearningAssignments",
		Method:             "GET",
		PathPattern:        "/api/v2/learning/assignments",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetLearningAssignmentsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetLearningAssignmentsOK), nil

}

/*
GetLearningAssignmentsMe lists of learning assignments assigned to current user
*/
func (a *Client) GetLearningAssignmentsMe(ctx context.Context, params *GetLearningAssignmentsMeParams) (*GetLearningAssignmentsMeOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getLearningAssignmentsMe",
		Method:             "GET",
		PathPattern:        "/api/v2/learning/assignments/me",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetLearningAssignmentsMeReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetLearningAssignmentsMeOK), nil

}

/*
GetLearningModule gets a learning module
*/
func (a *Client) GetLearningModule(ctx context.Context, params *GetLearningModuleParams) (*GetLearningModuleOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getLearningModule",
		Method:             "GET",
		PathPattern:        "/api/v2/learning/modules/{moduleId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetLearningModuleReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetLearningModuleOK), nil

}

/*
GetLearningModuleJob gets a specific learning module job status
*/
func (a *Client) GetLearningModuleJob(ctx context.Context, params *GetLearningModuleJobParams) (*GetLearningModuleJobOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getLearningModuleJob",
		Method:             "GET",
		PathPattern:        "/api/v2/learning/modules/{moduleId}/jobs/{jobId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetLearningModuleJobReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetLearningModuleJobOK), nil

}

/*
GetLearningModuleRule gets a learning module rule
*/
func (a *Client) GetLearningModuleRule(ctx context.Context, params *GetLearningModuleRuleParams) (*GetLearningModuleRuleOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getLearningModuleRule",
		Method:             "GET",
		PathPattern:        "/api/v2/learning/modules/{moduleId}/rule",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetLearningModuleRuleReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetLearningModuleRuleOK), nil

}

/*
GetLearningModuleVersion gets specific version of a published module
*/
func (a *Client) GetLearningModuleVersion(ctx context.Context, params *GetLearningModuleVersionParams) (*GetLearningModuleVersionOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getLearningModuleVersion",
		Method:             "GET",
		PathPattern:        "/api/v2/learning/modules/{moduleId}/versions/{versionId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetLearningModuleVersionReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetLearningModuleVersionOK), nil

}

/*
GetLearningModules gets all learning modules of an organization
*/
func (a *Client) GetLearningModules(ctx context.Context, params *GetLearningModulesParams) (*GetLearningModulesOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getLearningModules",
		Method:             "GET",
		PathPattern:        "/api/v2/learning/modules",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetLearningModulesReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetLearningModulesOK), nil

}

/*
GetLearningModulesAssignments gets all learning modules of an organization including assignments for a specific user
*/
func (a *Client) GetLearningModulesAssignments(ctx context.Context, params *GetLearningModulesAssignmentsParams) (*GetLearningModulesAssignmentsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getLearningModulesAssignments",
		Method:             "GET",
		PathPattern:        "/api/v2/learning/modules/assignments",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetLearningModulesAssignmentsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetLearningModulesAssignmentsOK), nil

}

/*
GetLearningModulesCoverartCoverArtID gets a specific learning module cover art using ID
*/
func (a *Client) GetLearningModulesCoverartCoverArtID(ctx context.Context, params *GetLearningModulesCoverartCoverArtIDParams) (*GetLearningModulesCoverartCoverArtIDOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getLearningModulesCoverartCoverArtId",
		Method:             "GET",
		PathPattern:        "/api/v2/learning/modules/coverart/{coverArtId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetLearningModulesCoverartCoverArtIDReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetLearningModulesCoverartCoverArtIDOK), nil

}

/*
PatchLearningAssignment updates learning assignment
*/
func (a *Client) PatchLearningAssignment(ctx context.Context, params *PatchLearningAssignmentParams) (*PatchLearningAssignmentOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "patchLearningAssignment",
		Method:             "PATCH",
		PathPattern:        "/api/v2/learning/assignments/{assignmentId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchLearningAssignmentReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PatchLearningAssignmentOK), nil

}

/*
PatchLearningAssignmentReschedule reschedules learning assignment
*/
func (a *Client) PatchLearningAssignmentReschedule(ctx context.Context, params *PatchLearningAssignmentRescheduleParams) (*PatchLearningAssignmentRescheduleOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "patchLearningAssignmentReschedule",
		Method:             "PATCH",
		PathPattern:        "/api/v2/learning/assignments/{assignmentId}/reschedule",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchLearningAssignmentRescheduleReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PatchLearningAssignmentRescheduleOK), nil

}

/*
PostLearningAssessmentsScoring scores learning assessment for preview
*/
func (a *Client) PostLearningAssessmentsScoring(ctx context.Context, params *PostLearningAssessmentsScoringParams) (*PostLearningAssessmentsScoringOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postLearningAssessmentsScoring",
		Method:             "POST",
		PathPattern:        "/api/v2/learning/assessments/scoring",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostLearningAssessmentsScoringReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostLearningAssessmentsScoringOK), nil

}

/*
PostLearningAssignmentReassign reassigns learning assignment

This will reassign the state of the assignment to 'Assigned' and update the assignment to the latest version of the module
*/
func (a *Client) PostLearningAssignmentReassign(ctx context.Context, params *PostLearningAssignmentReassignParams) (*PostLearningAssignmentReassignOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postLearningAssignmentReassign",
		Method:             "POST",
		PathPattern:        "/api/v2/learning/assignments/{assignmentId}/reassign",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostLearningAssignmentReassignReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostLearningAssignmentReassignOK), nil

}

/*
PostLearningAssignmentReset resets learning assignment

This will reset the state of the assignment to 'Assigned' and remove the version of Learning module associated with the assignment
*/
func (a *Client) PostLearningAssignmentReset(ctx context.Context, params *PostLearningAssignmentResetParams) (*PostLearningAssignmentResetOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postLearningAssignmentReset",
		Method:             "POST",
		PathPattern:        "/api/v2/learning/assignments/{assignmentId}/reset",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostLearningAssignmentResetReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostLearningAssignmentResetOK), nil

}

/*
PostLearningAssignments creates learning assignment
*/
func (a *Client) PostLearningAssignments(ctx context.Context, params *PostLearningAssignmentsParams) (*PostLearningAssignmentsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postLearningAssignments",
		Method:             "POST",
		PathPattern:        "/api/v2/learning/assignments",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostLearningAssignmentsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostLearningAssignmentsOK), nil

}

/*
PostLearningAssignmentsAggregatesQuery retrieves aggregated assignment data
*/
func (a *Client) PostLearningAssignmentsAggregatesQuery(ctx context.Context, params *PostLearningAssignmentsAggregatesQueryParams) (*PostLearningAssignmentsAggregatesQueryOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postLearningAssignmentsAggregatesQuery",
		Method:             "POST",
		PathPattern:        "/api/v2/learning/assignments/aggregates/query",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostLearningAssignmentsAggregatesQueryReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostLearningAssignmentsAggregatesQueryOK), nil

}

/*
PostLearningAssignmentsBulkadd adds multiple learning assignments
*/
func (a *Client) PostLearningAssignmentsBulkadd(ctx context.Context, params *PostLearningAssignmentsBulkaddParams) (*PostLearningAssignmentsBulkaddOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postLearningAssignmentsBulkadd",
		Method:             "POST",
		PathPattern:        "/api/v2/learning/assignments/bulkadd",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostLearningAssignmentsBulkaddReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostLearningAssignmentsBulkaddOK), nil

}

/*
PostLearningAssignmentsBulkremove removes multiple learning assignments
*/
func (a *Client) PostLearningAssignmentsBulkremove(ctx context.Context, params *PostLearningAssignmentsBulkremoveParams) (*PostLearningAssignmentsBulkremoveOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postLearningAssignmentsBulkremove",
		Method:             "POST",
		PathPattern:        "/api/v2/learning/assignments/bulkremove",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostLearningAssignmentsBulkremoveReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostLearningAssignmentsBulkremoveOK), nil

}

/*
PostLearningModuleJobs starts a specified operation on learning module

This will initiate operation specified in the request body for a learning module
*/
func (a *Client) PostLearningModuleJobs(ctx context.Context, params *PostLearningModuleJobsParams) (*PostLearningModuleJobsAccepted, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postLearningModuleJobs",
		Method:             "POST",
		PathPattern:        "/api/v2/learning/modules/{moduleId}/jobs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostLearningModuleJobsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostLearningModuleJobsAccepted), nil

}

/*
PostLearningModulePublish publishes a learning module
*/
func (a *Client) PostLearningModulePublish(ctx context.Context, params *PostLearningModulePublishParams) (*PostLearningModulePublishOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postLearningModulePublish",
		Method:             "POST",
		PathPattern:        "/api/v2/learning/modules/{moduleId}/publish",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostLearningModulePublishReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostLearningModulePublishOK), nil

}

/*
PostLearningModules creates a new learning module

This will create a new unpublished learning module with the specified fields.
*/
func (a *Client) PostLearningModules(ctx context.Context, params *PostLearningModulesParams) (*PostLearningModulesOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postLearningModules",
		Method:             "POST",
		PathPattern:        "/api/v2/learning/modules",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostLearningModulesReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostLearningModulesOK), nil

}

/*
PostLearningRulesQuery gets users for learning module rule

This will get the users who matches the given rule.
*/
func (a *Client) PostLearningRulesQuery(ctx context.Context, params *PostLearningRulesQueryParams) (*PostLearningRulesQueryOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postLearningRulesQuery",
		Method:             "POST",
		PathPattern:        "/api/v2/learning/rules/query",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostLearningRulesQueryReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostLearningRulesQueryOK), nil

}

/*
PostLearningScheduleslotsQuery gets list of possible slots where a learning activity can be scheduled
*/
func (a *Client) PostLearningScheduleslotsQuery(ctx context.Context, params *PostLearningScheduleslotsQueryParams) (*PostLearningScheduleslotsQueryOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postLearningScheduleslotsQuery",
		Method:             "POST",
		PathPattern:        "/api/v2/learning/scheduleslots/query",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostLearningScheduleslotsQueryReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostLearningScheduleslotsQueryOK), nil

}

/*
PutLearningModule updates a learning module

This will update the name, description, completion time in days and inform steps for a learning module
*/
func (a *Client) PutLearningModule(ctx context.Context, params *PutLearningModuleParams) (*PutLearningModuleOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "putLearningModule",
		Method:             "PUT",
		PathPattern:        "/api/v2/learning/modules/{moduleId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutLearningModuleReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutLearningModuleOK), nil

}

/*
PutLearningModuleRule updates a learning module rule

This will update a learning module rule with the specified fields.
*/
func (a *Client) PutLearningModuleRule(ctx context.Context, params *PutLearningModuleRuleParams) (*PutLearningModuleRuleOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "putLearningModuleRule",
		Method:             "PUT",
		PathPattern:        "/api/v2/learning/modules/{moduleId}/rule",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutLearningModuleRuleReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutLearningModuleRuleOK), nil

}

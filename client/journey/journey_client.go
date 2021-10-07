// Code generated by go-swagger; DO NOT EDIT.

package journey

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

//go:generate mockery -name API -inpkg

// API is the interface of the journey client
type API interface {
	/*
	   DeleteJourneyActionmap deletes single action map
	*/
	DeleteJourneyActionmap(ctx context.Context, params *DeleteJourneyActionmapParams) (*DeleteJourneyActionmapNoContent, error)
	/*
	   DeleteJourneyActiontemplate deletes a single action template
	*/
	DeleteJourneyActiontemplate(ctx context.Context, params *DeleteJourneyActiontemplateParams) (*DeleteJourneyActiontemplateNoContent, error)
	/*
	   DeleteJourneyOutcome deletes an outcome
	*/
	DeleteJourneyOutcome(ctx context.Context, params *DeleteJourneyOutcomeParams) (*DeleteJourneyOutcomeNoContent, error)
	/*
	   DeleteJourneySegment deletes a segment
	*/
	DeleteJourneySegment(ctx context.Context, params *DeleteJourneySegmentParams) (*DeleteJourneySegmentNoContent, error)
	/*
	   GetJourneyActionmap retrieves a single action map
	*/
	GetJourneyActionmap(ctx context.Context, params *GetJourneyActionmapParams) (*GetJourneyActionmapOK, error)
	/*
	   GetJourneyActionmaps retrieves all action maps
	*/
	GetJourneyActionmaps(ctx context.Context, params *GetJourneyActionmapsParams) (*GetJourneyActionmapsOK, error)
	/*
	   GetJourneyActiontarget retrieves a single action target
	*/
	GetJourneyActiontarget(ctx context.Context, params *GetJourneyActiontargetParams) (*GetJourneyActiontargetOK, error)
	/*
	   GetJourneyActiontargets retrieves all action targets
	*/
	GetJourneyActiontargets(ctx context.Context, params *GetJourneyActiontargetsParams) (*GetJourneyActiontargetsOK, error)
	/*
	   GetJourneyActiontemplate retrieves a single action template
	*/
	GetJourneyActiontemplate(ctx context.Context, params *GetJourneyActiontemplateParams) (*GetJourneyActiontemplateOK, error)
	/*
	   GetJourneyActiontemplates retrieves all action templates
	*/
	GetJourneyActiontemplates(ctx context.Context, params *GetJourneyActiontemplatesParams) (*GetJourneyActiontemplatesOK, error)
	/*
	   GetJourneyOutcome retrieves a single outcome
	*/
	GetJourneyOutcome(ctx context.Context, params *GetJourneyOutcomeParams) (*GetJourneyOutcomeOK, error)
	/*
	   GetJourneyOutcomes retrieves all outcomes
	*/
	GetJourneyOutcomes(ctx context.Context, params *GetJourneyOutcomesParams) (*GetJourneyOutcomesOK, error)
	/*
	   GetJourneySegment retrieves a single segment
	*/
	GetJourneySegment(ctx context.Context, params *GetJourneySegmentParams) (*GetJourneySegmentOK, error)
	/*
	   GetJourneySegments retrieves all segments
	*/
	GetJourneySegments(ctx context.Context, params *GetJourneySegmentsParams) (*GetJourneySegmentsOK, error)
	/*
	   GetJourneySession retrieves a single session
	*/
	GetJourneySession(ctx context.Context, params *GetJourneySessionParams) (*GetJourneySessionOK, error)
	/*
	   GetJourneySessionOutcomescores retrieves latest outcome score associated with a session for all outcomes
	*/
	GetJourneySessionOutcomescores(ctx context.Context, params *GetJourneySessionOutcomescoresParams) (*GetJourneySessionOutcomescoresOK, error)
	/*
	   PatchJourneyActionmap updates single action map
	*/
	PatchJourneyActionmap(ctx context.Context, params *PatchJourneyActionmapParams) (*PatchJourneyActionmapOK, error)
	/*
	   PatchJourneyActiontarget updates a single action target
	*/
	PatchJourneyActiontarget(ctx context.Context, params *PatchJourneyActiontargetParams) (*PatchJourneyActiontargetOK, error)
	/*
	   PatchJourneyActiontemplate updates a single action template
	*/
	PatchJourneyActiontemplate(ctx context.Context, params *PatchJourneyActiontemplateParams) (*PatchJourneyActiontemplateOK, error)
	/*
	   PatchJourneyOutcome updates an outcome
	*/
	PatchJourneyOutcome(ctx context.Context, params *PatchJourneyOutcomeParams) (*PatchJourneyOutcomeOK, error)
	/*
	   PatchJourneySegment updates a segment
	*/
	PatchJourneySegment(ctx context.Context, params *PatchJourneySegmentParams) (*PatchJourneySegmentOK, error)
	/*
	   PostJourneyActionmaps creates an action map
	*/
	PostJourneyActionmaps(ctx context.Context, params *PostJourneyActionmapsParams) (*PostJourneyActionmapsOK, *PostJourneyActionmapsCreated, error)
	/*
	   PostJourneyActiontemplates creates a single action template
	*/
	PostJourneyActiontemplates(ctx context.Context, params *PostJourneyActiontemplatesParams) (*PostJourneyActiontemplatesOK, *PostJourneyActiontemplatesCreated, error)
	/*
	   PostJourneyOutcomes creates an outcome
	*/
	PostJourneyOutcomes(ctx context.Context, params *PostJourneyOutcomesParams) (*PostJourneyOutcomesOK, *PostJourneyOutcomesCreated, error)
	/*
	   PostJourneySegments creates a segment
	*/
	PostJourneySegments(ctx context.Context, params *PostJourneySegmentsParams) (*PostJourneySegmentsOK, *PostJourneySegmentsCreated, error)
}

// New creates a new journey API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, authInfo runtime.ClientAuthInfoWriter) *Client {
	return &Client{
		transport: transport,
		formats:   formats,
		authInfo:  authInfo,
	}
}

/*
Client for journey API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	authInfo  runtime.ClientAuthInfoWriter
}

/*
DeleteJourneyActionmap deletes single action map
*/
func (a *Client) DeleteJourneyActionmap(ctx context.Context, params *DeleteJourneyActionmapParams) (*DeleteJourneyActionmapNoContent, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteJourneyActionmap",
		Method:             "DELETE",
		PathPattern:        "/api/v2/journey/actionmaps/{actionMapId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteJourneyActionmapReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteJourneyActionmapNoContent), nil

}

/*
DeleteJourneyActiontemplate deletes a single action template
*/
func (a *Client) DeleteJourneyActiontemplate(ctx context.Context, params *DeleteJourneyActiontemplateParams) (*DeleteJourneyActiontemplateNoContent, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteJourneyActiontemplate",
		Method:             "DELETE",
		PathPattern:        "/api/v2/journey/actiontemplates/{actionTemplateId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteJourneyActiontemplateReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteJourneyActiontemplateNoContent), nil

}

/*
DeleteJourneyOutcome deletes an outcome
*/
func (a *Client) DeleteJourneyOutcome(ctx context.Context, params *DeleteJourneyOutcomeParams) (*DeleteJourneyOutcomeNoContent, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteJourneyOutcome",
		Method:             "DELETE",
		PathPattern:        "/api/v2/journey/outcomes/{outcomeId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteJourneyOutcomeReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteJourneyOutcomeNoContent), nil

}

/*
DeleteJourneySegment deletes a segment
*/
func (a *Client) DeleteJourneySegment(ctx context.Context, params *DeleteJourneySegmentParams) (*DeleteJourneySegmentNoContent, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteJourneySegment",
		Method:             "DELETE",
		PathPattern:        "/api/v2/journey/segments/{segmentId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteJourneySegmentReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteJourneySegmentNoContent), nil

}

/*
GetJourneyActionmap retrieves a single action map
*/
func (a *Client) GetJourneyActionmap(ctx context.Context, params *GetJourneyActionmapParams) (*GetJourneyActionmapOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getJourneyActionmap",
		Method:             "GET",
		PathPattern:        "/api/v2/journey/actionmaps/{actionMapId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetJourneyActionmapReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetJourneyActionmapOK), nil

}

/*
GetJourneyActionmaps retrieves all action maps
*/
func (a *Client) GetJourneyActionmaps(ctx context.Context, params *GetJourneyActionmapsParams) (*GetJourneyActionmapsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getJourneyActionmaps",
		Method:             "GET",
		PathPattern:        "/api/v2/journey/actionmaps",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetJourneyActionmapsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetJourneyActionmapsOK), nil

}

/*
GetJourneyActiontarget retrieves a single action target
*/
func (a *Client) GetJourneyActiontarget(ctx context.Context, params *GetJourneyActiontargetParams) (*GetJourneyActiontargetOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getJourneyActiontarget",
		Method:             "GET",
		PathPattern:        "/api/v2/journey/actiontargets/{actionTargetId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetJourneyActiontargetReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetJourneyActiontargetOK), nil

}

/*
GetJourneyActiontargets retrieves all action targets
*/
func (a *Client) GetJourneyActiontargets(ctx context.Context, params *GetJourneyActiontargetsParams) (*GetJourneyActiontargetsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getJourneyActiontargets",
		Method:             "GET",
		PathPattern:        "/api/v2/journey/actiontargets",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetJourneyActiontargetsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetJourneyActiontargetsOK), nil

}

/*
GetJourneyActiontemplate retrieves a single action template
*/
func (a *Client) GetJourneyActiontemplate(ctx context.Context, params *GetJourneyActiontemplateParams) (*GetJourneyActiontemplateOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getJourneyActiontemplate",
		Method:             "GET",
		PathPattern:        "/api/v2/journey/actiontemplates/{actionTemplateId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetJourneyActiontemplateReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetJourneyActiontemplateOK), nil

}

/*
GetJourneyActiontemplates retrieves all action templates
*/
func (a *Client) GetJourneyActiontemplates(ctx context.Context, params *GetJourneyActiontemplatesParams) (*GetJourneyActiontemplatesOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getJourneyActiontemplates",
		Method:             "GET",
		PathPattern:        "/api/v2/journey/actiontemplates",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetJourneyActiontemplatesReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetJourneyActiontemplatesOK), nil

}

/*
GetJourneyOutcome retrieves a single outcome
*/
func (a *Client) GetJourneyOutcome(ctx context.Context, params *GetJourneyOutcomeParams) (*GetJourneyOutcomeOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getJourneyOutcome",
		Method:             "GET",
		PathPattern:        "/api/v2/journey/outcomes/{outcomeId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetJourneyOutcomeReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetJourneyOutcomeOK), nil

}

/*
GetJourneyOutcomes retrieves all outcomes
*/
func (a *Client) GetJourneyOutcomes(ctx context.Context, params *GetJourneyOutcomesParams) (*GetJourneyOutcomesOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getJourneyOutcomes",
		Method:             "GET",
		PathPattern:        "/api/v2/journey/outcomes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetJourneyOutcomesReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetJourneyOutcomesOK), nil

}

/*
GetJourneySegment retrieves a single segment
*/
func (a *Client) GetJourneySegment(ctx context.Context, params *GetJourneySegmentParams) (*GetJourneySegmentOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getJourneySegment",
		Method:             "GET",
		PathPattern:        "/api/v2/journey/segments/{segmentId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetJourneySegmentReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetJourneySegmentOK), nil

}

/*
GetJourneySegments retrieves all segments
*/
func (a *Client) GetJourneySegments(ctx context.Context, params *GetJourneySegmentsParams) (*GetJourneySegmentsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getJourneySegments",
		Method:             "GET",
		PathPattern:        "/api/v2/journey/segments",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetJourneySegmentsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetJourneySegmentsOK), nil

}

/*
GetJourneySession retrieves a single session
*/
func (a *Client) GetJourneySession(ctx context.Context, params *GetJourneySessionParams) (*GetJourneySessionOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getJourneySession",
		Method:             "GET",
		PathPattern:        "/api/v2/journey/sessions/{sessionId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetJourneySessionReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetJourneySessionOK), nil

}

/*
GetJourneySessionOutcomescores retrieves latest outcome score associated with a session for all outcomes
*/
func (a *Client) GetJourneySessionOutcomescores(ctx context.Context, params *GetJourneySessionOutcomescoresParams) (*GetJourneySessionOutcomescoresOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getJourneySessionOutcomescores",
		Method:             "GET",
		PathPattern:        "/api/v2/journey/sessions/{sessionId}/outcomescores",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetJourneySessionOutcomescoresReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetJourneySessionOutcomescoresOK), nil

}

/*
PatchJourneyActionmap updates single action map
*/
func (a *Client) PatchJourneyActionmap(ctx context.Context, params *PatchJourneyActionmapParams) (*PatchJourneyActionmapOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "patchJourneyActionmap",
		Method:             "PATCH",
		PathPattern:        "/api/v2/journey/actionmaps/{actionMapId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchJourneyActionmapReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PatchJourneyActionmapOK), nil

}

/*
PatchJourneyActiontarget updates a single action target
*/
func (a *Client) PatchJourneyActiontarget(ctx context.Context, params *PatchJourneyActiontargetParams) (*PatchJourneyActiontargetOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "patchJourneyActiontarget",
		Method:             "PATCH",
		PathPattern:        "/api/v2/journey/actiontargets/{actionTargetId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchJourneyActiontargetReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PatchJourneyActiontargetOK), nil

}

/*
PatchJourneyActiontemplate updates a single action template
*/
func (a *Client) PatchJourneyActiontemplate(ctx context.Context, params *PatchJourneyActiontemplateParams) (*PatchJourneyActiontemplateOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "patchJourneyActiontemplate",
		Method:             "PATCH",
		PathPattern:        "/api/v2/journey/actiontemplates/{actionTemplateId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchJourneyActiontemplateReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PatchJourneyActiontemplateOK), nil

}

/*
PatchJourneyOutcome updates an outcome
*/
func (a *Client) PatchJourneyOutcome(ctx context.Context, params *PatchJourneyOutcomeParams) (*PatchJourneyOutcomeOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "patchJourneyOutcome",
		Method:             "PATCH",
		PathPattern:        "/api/v2/journey/outcomes/{outcomeId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchJourneyOutcomeReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PatchJourneyOutcomeOK), nil

}

/*
PatchJourneySegment updates a segment
*/
func (a *Client) PatchJourneySegment(ctx context.Context, params *PatchJourneySegmentParams) (*PatchJourneySegmentOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "patchJourneySegment",
		Method:             "PATCH",
		PathPattern:        "/api/v2/journey/segments/{segmentId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchJourneySegmentReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PatchJourneySegmentOK), nil

}

/*
PostJourneyActionmaps creates an action map
*/
func (a *Client) PostJourneyActionmaps(ctx context.Context, params *PostJourneyActionmapsParams) (*PostJourneyActionmapsOK, *PostJourneyActionmapsCreated, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postJourneyActionmaps",
		Method:             "POST",
		PathPattern:        "/api/v2/journey/actionmaps",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostJourneyActionmapsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *PostJourneyActionmapsOK:
		return value, nil, nil
	case *PostJourneyActionmapsCreated:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
PostJourneyActiontemplates creates a single action template
*/
func (a *Client) PostJourneyActiontemplates(ctx context.Context, params *PostJourneyActiontemplatesParams) (*PostJourneyActiontemplatesOK, *PostJourneyActiontemplatesCreated, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postJourneyActiontemplates",
		Method:             "POST",
		PathPattern:        "/api/v2/journey/actiontemplates",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostJourneyActiontemplatesReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *PostJourneyActiontemplatesOK:
		return value, nil, nil
	case *PostJourneyActiontemplatesCreated:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
PostJourneyOutcomes creates an outcome
*/
func (a *Client) PostJourneyOutcomes(ctx context.Context, params *PostJourneyOutcomesParams) (*PostJourneyOutcomesOK, *PostJourneyOutcomesCreated, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postJourneyOutcomes",
		Method:             "POST",
		PathPattern:        "/api/v2/journey/outcomes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostJourneyOutcomesReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *PostJourneyOutcomesOK:
		return value, nil, nil
	case *PostJourneyOutcomesCreated:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
PostJourneySegments creates a segment
*/
func (a *Client) PostJourneySegments(ctx context.Context, params *PostJourneySegmentsParams) (*PostJourneySegmentsOK, *PostJourneySegmentsCreated, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postJourneySegments",
		Method:             "POST",
		PathPattern:        "/api/v2/journey/segments",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostJourneySegmentsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *PostJourneySegmentsOK:
		return value, nil, nil
	case *PostJourneySegmentsCreated:
		return nil, value, nil
	}
	return nil, nil, nil

}

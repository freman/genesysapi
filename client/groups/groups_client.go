// Code generated by go-swagger; DO NOT EDIT.

package groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

//go:generate mockery -name API -inpkg

// API is the interface of the groups client
type API interface {
	/*
	   DeleteGroup deletes group
	*/
	DeleteGroup(ctx context.Context, params *DeleteGroupParams) error
	/*
	   DeleteGroupMembers removes members
	*/
	DeleteGroupMembers(ctx context.Context, params *DeleteGroupMembersParams) (*DeleteGroupMembersAccepted, error)
	/*
	   GetGroup gets group
	*/
	GetGroup(ctx context.Context, params *GetGroupParams) (*GetGroupOK, error)
	/*
	   GetGroupIndividuals gets all individuals associated with the group
	*/
	GetGroupIndividuals(ctx context.Context, params *GetGroupIndividualsParams) (*GetGroupIndividualsOK, error)
	/*
	   GetGroupMembers gets group members includes individuals owners and dynamically included people
	*/
	GetGroupMembers(ctx context.Context, params *GetGroupMembersParams) (*GetGroupMembersOK, error)
	/*
	   GetGroupProfile gets group profile
	   This api is deprecated. Use /api/v2/groups instead
	*/
	GetGroupProfile(ctx context.Context, params *GetGroupProfileParams) (*GetGroupProfileOK, error)
	/*
	   GetGroups gets a group list
	*/
	GetGroups(ctx context.Context, params *GetGroupsParams) (*GetGroupsOK, error)
	/*
	   GetGroupsSearch searches groups using the q64 value returned from a previous search
	*/
	GetGroupsSearch(ctx context.Context, params *GetGroupsSearchParams) (*GetGroupsSearchOK, error)
	/*
	   GetProfilesGroups gets group profile listing
	   This api is deprecated. Use /api/v2/groups instead.
	*/
	GetProfilesGroups(ctx context.Context, params *GetProfilesGroupsParams) (*GetProfilesGroupsOK, error)
	/*
	   PostGroupMembers adds members
	*/
	PostGroupMembers(ctx context.Context, params *PostGroupMembersParams) (*PostGroupMembersAccepted, error)
	/*
	   PostGroups creates a group
	*/
	PostGroups(ctx context.Context, params *PostGroupsParams) (*PostGroupsOK, error)
	/*
	   PostGroupsSearch searches groups
	*/
	PostGroupsSearch(ctx context.Context, params *PostGroupsSearchParams) (*PostGroupsSearchOK, error)
	/*
	   PutGroup updates group
	*/
	PutGroup(ctx context.Context, params *PutGroupParams) (*PutGroupOK, error)
}

// New creates a new groups API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, authInfo runtime.ClientAuthInfoWriter) *Client {
	return &Client{
		transport: transport,
		formats:   formats,
		authInfo:  authInfo,
	}
}

/*
Client for groups API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	authInfo  runtime.ClientAuthInfoWriter
}

/*
DeleteGroup deletes group
*/
func (a *Client) DeleteGroup(ctx context.Context, params *DeleteGroupParams) error {

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteGroup",
		Method:             "DELETE",
		PathPattern:        "/api/v2/groups/{groupId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteGroupReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
DeleteGroupMembers removes members
*/
func (a *Client) DeleteGroupMembers(ctx context.Context, params *DeleteGroupMembersParams) (*DeleteGroupMembersAccepted, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteGroupMembers",
		Method:             "DELETE",
		PathPattern:        "/api/v2/groups/{groupId}/members",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteGroupMembersReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteGroupMembersAccepted), nil

}

/*
GetGroup gets group
*/
func (a *Client) GetGroup(ctx context.Context, params *GetGroupParams) (*GetGroupOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getGroup",
		Method:             "GET",
		PathPattern:        "/api/v2/groups/{groupId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetGroupReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetGroupOK), nil

}

/*
GetGroupIndividuals gets all individuals associated with the group
*/
func (a *Client) GetGroupIndividuals(ctx context.Context, params *GetGroupIndividualsParams) (*GetGroupIndividualsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getGroupIndividuals",
		Method:             "GET",
		PathPattern:        "/api/v2/groups/{groupId}/individuals",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetGroupIndividualsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetGroupIndividualsOK), nil

}

/*
GetGroupMembers gets group members includes individuals owners and dynamically included people
*/
func (a *Client) GetGroupMembers(ctx context.Context, params *GetGroupMembersParams) (*GetGroupMembersOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getGroupMembers",
		Method:             "GET",
		PathPattern:        "/api/v2/groups/{groupId}/members",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetGroupMembersReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetGroupMembersOK), nil

}

/*
GetGroupProfile gets group profile

This api is deprecated. Use /api/v2/groups instead
*/
func (a *Client) GetGroupProfile(ctx context.Context, params *GetGroupProfileParams) (*GetGroupProfileOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getGroupProfile",
		Method:             "GET",
		PathPattern:        "/api/v2/groups/{groupId}/profile",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetGroupProfileReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetGroupProfileOK), nil

}

/*
GetGroups gets a group list
*/
func (a *Client) GetGroups(ctx context.Context, params *GetGroupsParams) (*GetGroupsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getGroups",
		Method:             "GET",
		PathPattern:        "/api/v2/groups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetGroupsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetGroupsOK), nil

}

/*
GetGroupsSearch searches groups using the q64 value returned from a previous search
*/
func (a *Client) GetGroupsSearch(ctx context.Context, params *GetGroupsSearchParams) (*GetGroupsSearchOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getGroupsSearch",
		Method:             "GET",
		PathPattern:        "/api/v2/groups/search",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetGroupsSearchReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetGroupsSearchOK), nil

}

/*
GetProfilesGroups gets group profile listing

This api is deprecated. Use /api/v2/groups instead.
*/
func (a *Client) GetProfilesGroups(ctx context.Context, params *GetProfilesGroupsParams) (*GetProfilesGroupsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getProfilesGroups",
		Method:             "GET",
		PathPattern:        "/api/v2/profiles/groups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetProfilesGroupsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetProfilesGroupsOK), nil

}

/*
PostGroupMembers adds members
*/
func (a *Client) PostGroupMembers(ctx context.Context, params *PostGroupMembersParams) (*PostGroupMembersAccepted, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postGroupMembers",
		Method:             "POST",
		PathPattern:        "/api/v2/groups/{groupId}/members",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostGroupMembersReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostGroupMembersAccepted), nil

}

/*
PostGroups creates a group
*/
func (a *Client) PostGroups(ctx context.Context, params *PostGroupsParams) (*PostGroupsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postGroups",
		Method:             "POST",
		PathPattern:        "/api/v2/groups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostGroupsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostGroupsOK), nil

}

/*
PostGroupsSearch searches groups
*/
func (a *Client) PostGroupsSearch(ctx context.Context, params *PostGroupsSearchParams) (*PostGroupsSearchOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postGroupsSearch",
		Method:             "POST",
		PathPattern:        "/api/v2/groups/search",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostGroupsSearchReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostGroupsSearchOK), nil

}

/*
PutGroup updates group
*/
func (a *Client) PutGroup(ctx context.Context, params *PutGroupParams) (*PutGroupOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "putGroup",
		Method:             "PUT",
		PathPattern:        "/api/v2/groups/{groupId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutGroupReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutGroupOK), nil

}
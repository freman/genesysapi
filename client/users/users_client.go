// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

//go:generate mockery -name API -inpkg

// API is the interface of the users client
type API interface {
	/*
	   DeleteAnalyticsUsersDetailsJob deletes cancel an async request
	*/
	DeleteAnalyticsUsersDetailsJob(ctx context.Context, params *DeleteAnalyticsUsersDetailsJobParams) (*DeleteAnalyticsUsersDetailsJobNoContent, error)
	/*
	   DeleteUser deletes user
	*/
	DeleteUser(ctx context.Context, params *DeleteUserParams) (*DeleteUserOK, error)
	/*
	   DeleteUserStationAssociatedstation clears associated station
	*/
	DeleteUserStationAssociatedstation(ctx context.Context, params *DeleteUserStationAssociatedstationParams) (*DeleteUserStationAssociatedstationAccepted, error)
	/*
	   DeleteUserStationDefaultstation clears default station
	*/
	DeleteUserStationDefaultstation(ctx context.Context, params *DeleteUserStationDefaultstationParams) (*DeleteUserStationDefaultstationAccepted, error)
	/*
	   GetAnalyticsUsersDetailsJob gets status for async query for user details
	*/
	GetAnalyticsUsersDetailsJob(ctx context.Context, params *GetAnalyticsUsersDetailsJobParams) (*GetAnalyticsUsersDetailsJobOK, *GetAnalyticsUsersDetailsJobAccepted, error)
	/*
	   GetAnalyticsUsersDetailsJobResults fetches a page of results for an async query
	*/
	GetAnalyticsUsersDetailsJobResults(ctx context.Context, params *GetAnalyticsUsersDetailsJobResultsParams) (*GetAnalyticsUsersDetailsJobResultsOK, error)
	/*
	   GetProfilesUsers gets a user profile listing
	   This api is deprecated. User /api/v2/users
	*/
	GetProfilesUsers(ctx context.Context, params *GetProfilesUsersParams) (*GetProfilesUsersOK, error)
	/*
	   GetUser gets user
	*/
	GetUser(ctx context.Context, params *GetUserParams) (*GetUserOK, error)
	/*
	   GetUserAdjacents gets adjacents
	*/
	GetUserAdjacents(ctx context.Context, params *GetUserAdjacentsParams) (*GetUserAdjacentsOK, error)
	/*
	   GetUserCallforwarding gets a user s call forwarding
	*/
	GetUserCallforwarding(ctx context.Context, params *GetUserCallforwardingParams) (*GetUserCallforwardingOK, error)
	/*
	   GetUserDirectreports gets direct reports
	*/
	GetUserDirectreports(ctx context.Context, params *GetUserDirectreportsParams) (*GetUserDirectreportsOK, error)
	/*
	   GetUserFavorites gets favorites
	*/
	GetUserFavorites(ctx context.Context, params *GetUserFavoritesParams) (*GetUserFavoritesOK, error)
	/*
	   GetUserOutofoffice gets a out of office
	*/
	GetUserOutofoffice(ctx context.Context, params *GetUserOutofofficeParams) (*GetUserOutofofficeOK, error)
	/*
	   GetUserProfile gets user profile
	   This api has been deprecated. Use api/v2/users instead
	*/
	GetUserProfile(ctx context.Context, params *GetUserProfileParams) (*GetUserProfileOK, error)
	/*
	   GetUserProfileskills lists profile skills for a user
	*/
	GetUserProfileskills(ctx context.Context, params *GetUserProfileskillsParams) (*GetUserProfileskillsOK, error)
	/*
	   GetUserRoutingstatus fetches the routing status of a user
	*/
	GetUserRoutingstatus(ctx context.Context, params *GetUserRoutingstatusParams) (*GetUserRoutingstatusOK, error)
	/*
	   GetUserStation gets station information for user
	*/
	GetUserStation(ctx context.Context, params *GetUserStationParams) (*GetUserStationOK, error)
	/*
	   GetUserSuperiors gets superiors
	*/
	GetUserSuperiors(ctx context.Context, params *GetUserSuperiorsParams) (*GetUserSuperiorsOK, error)
	/*
	   GetUserTrustors lists the organizations that have authorized trusted the user
	*/
	GetUserTrustors(ctx context.Context, params *GetUserTrustorsParams) (*GetUserTrustorsOK, error)
	/*
	   GetUsers gets the list of available users
	*/
	GetUsers(ctx context.Context, params *GetUsersParams) (*GetUsersOK, error)
	/*
	   GetUsersMe gets current user details
	   This request is not valid when using the Client Credentials OAuth grant.
	*/
	GetUsersMe(ctx context.Context, params *GetUsersMeParams) (*GetUsersMeOK, error)
	/*
	   PatchUser updates user
	*/
	PatchUser(ctx context.Context, params *PatchUserParams) (*PatchUserOK, error)
	/*
	   PatchUserCallforwarding patches a user s call forwarding
	*/
	PatchUserCallforwarding(ctx context.Context, params *PatchUserCallforwardingParams) (*PatchUserCallforwardingOK, error)
	/*
	   PatchUsersBulk updates bulk acd autoanswer on users
	*/
	PatchUsersBulk(ctx context.Context, params *PatchUsersBulkParams) (*PatchUsersBulkOK, error)
	/*
	   PostAnalyticsUsersAggregatesQuery queries for user aggregates
	*/
	PostAnalyticsUsersAggregatesQuery(ctx context.Context, params *PostAnalyticsUsersAggregatesQueryParams) (*PostAnalyticsUsersAggregatesQueryOK, error)
	/*
	   PostAnalyticsUsersDetailsJobs queries for user details asynchronously
	*/
	PostAnalyticsUsersDetailsJobs(ctx context.Context, params *PostAnalyticsUsersDetailsJobsParams) (*PostAnalyticsUsersDetailsJobsAccepted, error)
	/*
	   PostAnalyticsUsersDetailsQuery queries for user details
	*/
	PostAnalyticsUsersDetailsQuery(ctx context.Context, params *PostAnalyticsUsersDetailsQueryParams) (*PostAnalyticsUsersDetailsQueryOK, error)
	/*
	   PostAnalyticsUsersObservationsQuery queries for user observations
	*/
	PostAnalyticsUsersObservationsQuery(ctx context.Context, params *PostAnalyticsUsersObservationsQueryParams) (*PostAnalyticsUsersObservationsQueryOK, error)
	/*
	   PostUserInvite sends an activation email to the user
	*/
	PostUserInvite(ctx context.Context, params *PostUserInviteParams) (*PostUserInviteNoContent, error)
	/*
	   PostUserPassword changes a users password
	*/
	PostUserPassword(ctx context.Context, params *PostUserPasswordParams) (*PostUserPasswordNoContent, error)
	/*
	   PostUsers creates user
	*/
	PostUsers(ctx context.Context, params *PostUsersParams) (*PostUsersOK, error)
	/*
	   PostUsersMePassword changes your password
	*/
	PostUsersMePassword(ctx context.Context, params *PostUsersMePasswordParams) (*PostUsersMePasswordNoContent, error)
	/*
	   PutUserCallforwarding updates a user s call forwarding
	*/
	PutUserCallforwarding(ctx context.Context, params *PutUserCallforwardingParams) (*PutUserCallforwardingOK, error)
	/*
	   PutUserOutofoffice updates an out of office
	*/
	PutUserOutofoffice(ctx context.Context, params *PutUserOutofofficeParams) (*PutUserOutofofficeOK, error)
	/*
	   PutUserProfileskills updates profile skills for a user
	*/
	PutUserProfileskills(ctx context.Context, params *PutUserProfileskillsParams) (*PutUserProfileskillsOK, error)
	/*
	   PutUserRoutingstatus updates the routing status of a user
	*/
	PutUserRoutingstatus(ctx context.Context, params *PutUserRoutingstatusParams) (*PutUserRoutingstatusOK, error)
	/*
	   PutUserStationAssociatedstationStationID sets associated station
	*/
	PutUserStationAssociatedstationStationID(ctx context.Context, params *PutUserStationAssociatedstationStationIDParams) (*PutUserStationAssociatedstationStationIDAccepted, error)
	/*
	   PutUserStationDefaultstationStationID sets default station
	*/
	PutUserStationDefaultstationStationID(ctx context.Context, params *PutUserStationDefaultstationStationIDParams) (*PutUserStationDefaultstationStationIDAccepted, error)
}

// New creates a new users API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, authInfo runtime.ClientAuthInfoWriter) *Client {
	return &Client{
		transport: transport,
		formats:   formats,
		authInfo:  authInfo,
	}
}

/*
Client for users API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	authInfo  runtime.ClientAuthInfoWriter
}

/*
DeleteAnalyticsUsersDetailsJob deletes cancel an async request
*/
func (a *Client) DeleteAnalyticsUsersDetailsJob(ctx context.Context, params *DeleteAnalyticsUsersDetailsJobParams) (*DeleteAnalyticsUsersDetailsJobNoContent, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteAnalyticsUsersDetailsJob",
		Method:             "DELETE",
		PathPattern:        "/api/v2/analytics/users/details/jobs/{jobId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteAnalyticsUsersDetailsJobReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteAnalyticsUsersDetailsJobNoContent), nil

}

/*
DeleteUser deletes user
*/
func (a *Client) DeleteUser(ctx context.Context, params *DeleteUserParams) (*DeleteUserOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteUser",
		Method:             "DELETE",
		PathPattern:        "/api/v2/users/{userId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteUserReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteUserOK), nil

}

/*
DeleteUserStationAssociatedstation clears associated station
*/
func (a *Client) DeleteUserStationAssociatedstation(ctx context.Context, params *DeleteUserStationAssociatedstationParams) (*DeleteUserStationAssociatedstationAccepted, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteUserStationAssociatedstation",
		Method:             "DELETE",
		PathPattern:        "/api/v2/users/{userId}/station/associatedstation",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteUserStationAssociatedstationReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteUserStationAssociatedstationAccepted), nil

}

/*
DeleteUserStationDefaultstation clears default station
*/
func (a *Client) DeleteUserStationDefaultstation(ctx context.Context, params *DeleteUserStationDefaultstationParams) (*DeleteUserStationDefaultstationAccepted, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteUserStationDefaultstation",
		Method:             "DELETE",
		PathPattern:        "/api/v2/users/{userId}/station/defaultstation",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteUserStationDefaultstationReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteUserStationDefaultstationAccepted), nil

}

/*
GetAnalyticsUsersDetailsJob gets status for async query for user details
*/
func (a *Client) GetAnalyticsUsersDetailsJob(ctx context.Context, params *GetAnalyticsUsersDetailsJobParams) (*GetAnalyticsUsersDetailsJobOK, *GetAnalyticsUsersDetailsJobAccepted, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAnalyticsUsersDetailsJob",
		Method:             "GET",
		PathPattern:        "/api/v2/analytics/users/details/jobs/{jobId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAnalyticsUsersDetailsJobReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *GetAnalyticsUsersDetailsJobOK:
		return value, nil, nil
	case *GetAnalyticsUsersDetailsJobAccepted:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
GetAnalyticsUsersDetailsJobResults fetches a page of results for an async query
*/
func (a *Client) GetAnalyticsUsersDetailsJobResults(ctx context.Context, params *GetAnalyticsUsersDetailsJobResultsParams) (*GetAnalyticsUsersDetailsJobResultsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAnalyticsUsersDetailsJobResults",
		Method:             "GET",
		PathPattern:        "/api/v2/analytics/users/details/jobs/{jobId}/results",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAnalyticsUsersDetailsJobResultsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAnalyticsUsersDetailsJobResultsOK), nil

}

/*
GetProfilesUsers gets a user profile listing

This api is deprecated. User /api/v2/users
*/
func (a *Client) GetProfilesUsers(ctx context.Context, params *GetProfilesUsersParams) (*GetProfilesUsersOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getProfilesUsers",
		Method:             "GET",
		PathPattern:        "/api/v2/profiles/users",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetProfilesUsersReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetProfilesUsersOK), nil

}

/*
GetUser gets user
*/
func (a *Client) GetUser(ctx context.Context, params *GetUserParams) (*GetUserOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getUser",
		Method:             "GET",
		PathPattern:        "/api/v2/users/{userId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetUserReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetUserOK), nil

}

/*
GetUserAdjacents gets adjacents
*/
func (a *Client) GetUserAdjacents(ctx context.Context, params *GetUserAdjacentsParams) (*GetUserAdjacentsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getUserAdjacents",
		Method:             "GET",
		PathPattern:        "/api/v2/users/{userId}/adjacents",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetUserAdjacentsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetUserAdjacentsOK), nil

}

/*
GetUserCallforwarding gets a user s call forwarding
*/
func (a *Client) GetUserCallforwarding(ctx context.Context, params *GetUserCallforwardingParams) (*GetUserCallforwardingOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getUserCallforwarding",
		Method:             "GET",
		PathPattern:        "/api/v2/users/{userId}/callforwarding",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetUserCallforwardingReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetUserCallforwardingOK), nil

}

/*
GetUserDirectreports gets direct reports
*/
func (a *Client) GetUserDirectreports(ctx context.Context, params *GetUserDirectreportsParams) (*GetUserDirectreportsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getUserDirectreports",
		Method:             "GET",
		PathPattern:        "/api/v2/users/{userId}/directreports",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetUserDirectreportsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetUserDirectreportsOK), nil

}

/*
GetUserFavorites gets favorites
*/
func (a *Client) GetUserFavorites(ctx context.Context, params *GetUserFavoritesParams) (*GetUserFavoritesOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getUserFavorites",
		Method:             "GET",
		PathPattern:        "/api/v2/users/{userId}/favorites",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetUserFavoritesReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetUserFavoritesOK), nil

}

/*
GetUserOutofoffice gets a out of office
*/
func (a *Client) GetUserOutofoffice(ctx context.Context, params *GetUserOutofofficeParams) (*GetUserOutofofficeOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getUserOutofoffice",
		Method:             "GET",
		PathPattern:        "/api/v2/users/{userId}/outofoffice",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetUserOutofofficeReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetUserOutofofficeOK), nil

}

/*
GetUserProfile gets user profile

This api has been deprecated. Use api/v2/users instead
*/
func (a *Client) GetUserProfile(ctx context.Context, params *GetUserProfileParams) (*GetUserProfileOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getUserProfile",
		Method:             "GET",
		PathPattern:        "/api/v2/users/{userId}/profile",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetUserProfileReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetUserProfileOK), nil

}

/*
GetUserProfileskills lists profile skills for a user
*/
func (a *Client) GetUserProfileskills(ctx context.Context, params *GetUserProfileskillsParams) (*GetUserProfileskillsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getUserProfileskills",
		Method:             "GET",
		PathPattern:        "/api/v2/users/{userId}/profileskills",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetUserProfileskillsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetUserProfileskillsOK), nil

}

/*
GetUserRoutingstatus fetches the routing status of a user
*/
func (a *Client) GetUserRoutingstatus(ctx context.Context, params *GetUserRoutingstatusParams) (*GetUserRoutingstatusOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getUserRoutingstatus",
		Method:             "GET",
		PathPattern:        "/api/v2/users/{userId}/routingstatus",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetUserRoutingstatusReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetUserRoutingstatusOK), nil

}

/*
GetUserStation gets station information for user
*/
func (a *Client) GetUserStation(ctx context.Context, params *GetUserStationParams) (*GetUserStationOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getUserStation",
		Method:             "GET",
		PathPattern:        "/api/v2/users/{userId}/station",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetUserStationReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetUserStationOK), nil

}

/*
GetUserSuperiors gets superiors
*/
func (a *Client) GetUserSuperiors(ctx context.Context, params *GetUserSuperiorsParams) (*GetUserSuperiorsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getUserSuperiors",
		Method:             "GET",
		PathPattern:        "/api/v2/users/{userId}/superiors",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetUserSuperiorsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetUserSuperiorsOK), nil

}

/*
GetUserTrustors lists the organizations that have authorized trusted the user
*/
func (a *Client) GetUserTrustors(ctx context.Context, params *GetUserTrustorsParams) (*GetUserTrustorsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getUserTrustors",
		Method:             "GET",
		PathPattern:        "/api/v2/users/{userId}/trustors",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetUserTrustorsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetUserTrustorsOK), nil

}

/*
GetUsers gets the list of available users
*/
func (a *Client) GetUsers(ctx context.Context, params *GetUsersParams) (*GetUsersOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getUsers",
		Method:             "GET",
		PathPattern:        "/api/v2/users",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetUsersReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetUsersOK), nil

}

/*
GetUsersMe gets current user details

This request is not valid when using the Client Credentials OAuth grant.
*/
func (a *Client) GetUsersMe(ctx context.Context, params *GetUsersMeParams) (*GetUsersMeOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getUsersMe",
		Method:             "GET",
		PathPattern:        "/api/v2/users/me",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetUsersMeReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetUsersMeOK), nil

}

/*
PatchUser updates user
*/
func (a *Client) PatchUser(ctx context.Context, params *PatchUserParams) (*PatchUserOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "patchUser",
		Method:             "PATCH",
		PathPattern:        "/api/v2/users/{userId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchUserReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PatchUserOK), nil

}

/*
PatchUserCallforwarding patches a user s call forwarding
*/
func (a *Client) PatchUserCallforwarding(ctx context.Context, params *PatchUserCallforwardingParams) (*PatchUserCallforwardingOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "patchUserCallforwarding",
		Method:             "PATCH",
		PathPattern:        "/api/v2/users/{userId}/callforwarding",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchUserCallforwardingReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PatchUserCallforwardingOK), nil

}

/*
PatchUsersBulk updates bulk acd autoanswer on users
*/
func (a *Client) PatchUsersBulk(ctx context.Context, params *PatchUsersBulkParams) (*PatchUsersBulkOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "patchUsersBulk",
		Method:             "PATCH",
		PathPattern:        "/api/v2/users/bulk",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchUsersBulkReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PatchUsersBulkOK), nil

}

/*
PostAnalyticsUsersAggregatesQuery queries for user aggregates
*/
func (a *Client) PostAnalyticsUsersAggregatesQuery(ctx context.Context, params *PostAnalyticsUsersAggregatesQueryParams) (*PostAnalyticsUsersAggregatesQueryOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postAnalyticsUsersAggregatesQuery",
		Method:             "POST",
		PathPattern:        "/api/v2/analytics/users/aggregates/query",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostAnalyticsUsersAggregatesQueryReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostAnalyticsUsersAggregatesQueryOK), nil

}

/*
PostAnalyticsUsersDetailsJobs queries for user details asynchronously
*/
func (a *Client) PostAnalyticsUsersDetailsJobs(ctx context.Context, params *PostAnalyticsUsersDetailsJobsParams) (*PostAnalyticsUsersDetailsJobsAccepted, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postAnalyticsUsersDetailsJobs",
		Method:             "POST",
		PathPattern:        "/api/v2/analytics/users/details/jobs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostAnalyticsUsersDetailsJobsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostAnalyticsUsersDetailsJobsAccepted), nil

}

/*
PostAnalyticsUsersDetailsQuery queries for user details
*/
func (a *Client) PostAnalyticsUsersDetailsQuery(ctx context.Context, params *PostAnalyticsUsersDetailsQueryParams) (*PostAnalyticsUsersDetailsQueryOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postAnalyticsUsersDetailsQuery",
		Method:             "POST",
		PathPattern:        "/api/v2/analytics/users/details/query",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostAnalyticsUsersDetailsQueryReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostAnalyticsUsersDetailsQueryOK), nil

}

/*
PostAnalyticsUsersObservationsQuery queries for user observations
*/
func (a *Client) PostAnalyticsUsersObservationsQuery(ctx context.Context, params *PostAnalyticsUsersObservationsQueryParams) (*PostAnalyticsUsersObservationsQueryOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postAnalyticsUsersObservationsQuery",
		Method:             "POST",
		PathPattern:        "/api/v2/analytics/users/observations/query",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostAnalyticsUsersObservationsQueryReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostAnalyticsUsersObservationsQueryOK), nil

}

/*
PostUserInvite sends an activation email to the user
*/
func (a *Client) PostUserInvite(ctx context.Context, params *PostUserInviteParams) (*PostUserInviteNoContent, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postUserInvite",
		Method:             "POST",
		PathPattern:        "/api/v2/users/{userId}/invite",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostUserInviteReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostUserInviteNoContent), nil

}

/*
PostUserPassword changes a users password
*/
func (a *Client) PostUserPassword(ctx context.Context, params *PostUserPasswordParams) (*PostUserPasswordNoContent, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postUserPassword",
		Method:             "POST",
		PathPattern:        "/api/v2/users/{userId}/password",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostUserPasswordReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostUserPasswordNoContent), nil

}

/*
PostUsers creates user
*/
func (a *Client) PostUsers(ctx context.Context, params *PostUsersParams) (*PostUsersOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postUsers",
		Method:             "POST",
		PathPattern:        "/api/v2/users",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostUsersReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostUsersOK), nil

}

/*
PostUsersMePassword changes your password
*/
func (a *Client) PostUsersMePassword(ctx context.Context, params *PostUsersMePasswordParams) (*PostUsersMePasswordNoContent, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postUsersMePassword",
		Method:             "POST",
		PathPattern:        "/api/v2/users/me/password",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostUsersMePasswordReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostUsersMePasswordNoContent), nil

}

/*
PutUserCallforwarding updates a user s call forwarding
*/
func (a *Client) PutUserCallforwarding(ctx context.Context, params *PutUserCallforwardingParams) (*PutUserCallforwardingOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "putUserCallforwarding",
		Method:             "PUT",
		PathPattern:        "/api/v2/users/{userId}/callforwarding",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutUserCallforwardingReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutUserCallforwardingOK), nil

}

/*
PutUserOutofoffice updates an out of office
*/
func (a *Client) PutUserOutofoffice(ctx context.Context, params *PutUserOutofofficeParams) (*PutUserOutofofficeOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "putUserOutofoffice",
		Method:             "PUT",
		PathPattern:        "/api/v2/users/{userId}/outofoffice",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutUserOutofofficeReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutUserOutofofficeOK), nil

}

/*
PutUserProfileskills updates profile skills for a user
*/
func (a *Client) PutUserProfileskills(ctx context.Context, params *PutUserProfileskillsParams) (*PutUserProfileskillsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "putUserProfileskills",
		Method:             "PUT",
		PathPattern:        "/api/v2/users/{userId}/profileskills",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutUserProfileskillsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutUserProfileskillsOK), nil

}

/*
PutUserRoutingstatus updates the routing status of a user
*/
func (a *Client) PutUserRoutingstatus(ctx context.Context, params *PutUserRoutingstatusParams) (*PutUserRoutingstatusOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "putUserRoutingstatus",
		Method:             "PUT",
		PathPattern:        "/api/v2/users/{userId}/routingstatus",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutUserRoutingstatusReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutUserRoutingstatusOK), nil

}

/*
PutUserStationAssociatedstationStationID sets associated station
*/
func (a *Client) PutUserStationAssociatedstationStationID(ctx context.Context, params *PutUserStationAssociatedstationStationIDParams) (*PutUserStationAssociatedstationStationIDAccepted, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "putUserStationAssociatedstationStationId",
		Method:             "PUT",
		PathPattern:        "/api/v2/users/{userId}/station/associatedstation/{stationId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutUserStationAssociatedstationStationIDReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutUserStationAssociatedstationStationIDAccepted), nil

}

/*
PutUserStationDefaultstationStationID sets default station
*/
func (a *Client) PutUserStationDefaultstationStationID(ctx context.Context, params *PutUserStationDefaultstationStationIDParams) (*PutUserStationDefaultstationStationIDAccepted, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "putUserStationDefaultstationStationId",
		Method:             "PUT",
		PathPattern:        "/api/v2/users/{userId}/station/defaultstation/{stationId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutUserStationDefaultstationStationIDReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutUserStationDefaultstationStationIDAccepted), nil

}

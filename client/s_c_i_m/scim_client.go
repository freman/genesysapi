// Code generated by go-swagger; DO NOT EDIT.

package s_c_i_m

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

//go:generate mockery -name API -inpkg

// API is the interface of the s c i m client
type API interface {
	/*
	   DeleteScimGroup deletes a group
	*/
	DeleteScimGroup(ctx context.Context, params *DeleteScimGroupParams) (*DeleteScimGroupNoContent, error)
	/*
	   DeleteScimUser deletes a user
	*/
	DeleteScimUser(ctx context.Context, params *DeleteScimUserParams) (*DeleteScimUserOK, *DeleteScimUserNoContent, error)
	/*
	   DeleteScimV2Group deletes a group
	*/
	DeleteScimV2Group(ctx context.Context, params *DeleteScimV2GroupParams) (*DeleteScimV2GroupNoContent, error)
	/*
	   DeleteScimV2User deletes a user
	*/
	DeleteScimV2User(ctx context.Context, params *DeleteScimV2UserParams) (*DeleteScimV2UserOK, *DeleteScimV2UserNoContent, error)
	/*
	   GetScimGroup gets a group
	*/
	GetScimGroup(ctx context.Context, params *GetScimGroupParams) (*GetScimGroupOK, error)
	/*
	   GetScimGroups gets a list of groups
	*/
	GetScimGroups(ctx context.Context, params *GetScimGroupsParams) (*GetScimGroupsOK, error)
	/*
	   GetScimResourcetype gets a resource type
	*/
	GetScimResourcetype(ctx context.Context, params *GetScimResourcetypeParams) (*GetScimResourcetypeOK, error)
	/*
	   GetScimResourcetypes gets a list of resource types
	*/
	GetScimResourcetypes(ctx context.Context, params *GetScimResourcetypesParams) (*GetScimResourcetypesOK, error)
	/*
	   GetScimSchema gets a s c i m schema
	*/
	GetScimSchema(ctx context.Context, params *GetScimSchemaParams) (*GetScimSchemaOK, error)
	/*
	   GetScimSchemas gets a list of s c i m schemas
	*/
	GetScimSchemas(ctx context.Context, params *GetScimSchemasParams) (*GetScimSchemasOK, error)
	/*
	   GetScimServiceproviderconfig gets a service provider s configuration
	*/
	GetScimServiceproviderconfig(ctx context.Context, params *GetScimServiceproviderconfigParams) (*GetScimServiceproviderconfigOK, error)
	/*
	   GetScimUser gets a user
	*/
	GetScimUser(ctx context.Context, params *GetScimUserParams) (*GetScimUserOK, error)
	/*
	   GetScimUsers gets a list of users
	   To return all active users, do not use the filter parameter. To return inactive users, set the filter parameter to "active eq false". By default, returns SCIM attributes "externalId", "enterprise-user:manager", and "roles". To exclude these attributes, set the attributes parameter to "id,active" or the excludeAttributes parameter to "externalId,roles,urn:ietf:params:scim:schemas:extension:enterprise:2.0:User:division".
	*/
	GetScimUsers(ctx context.Context, params *GetScimUsersParams) (*GetScimUsersOK, error)
	/*
	   GetScimV2Group gets a group
	*/
	GetScimV2Group(ctx context.Context, params *GetScimV2GroupParams) (*GetScimV2GroupOK, error)
	/*
	   GetScimV2Groups gets a list of groups
	*/
	GetScimV2Groups(ctx context.Context, params *GetScimV2GroupsParams) (*GetScimV2GroupsOK, error)
	/*
	   GetScimV2Resourcetype gets a resource type
	*/
	GetScimV2Resourcetype(ctx context.Context, params *GetScimV2ResourcetypeParams) (*GetScimV2ResourcetypeOK, error)
	/*
	   GetScimV2Resourcetypes gets a list of resource types
	*/
	GetScimV2Resourcetypes(ctx context.Context, params *GetScimV2ResourcetypesParams) (*GetScimV2ResourcetypesOK, error)
	/*
	   GetScimV2Schema gets a s c i m schema
	*/
	GetScimV2Schema(ctx context.Context, params *GetScimV2SchemaParams) (*GetScimV2SchemaOK, error)
	/*
	   GetScimV2Schemas gets a list of s c i m schemas
	*/
	GetScimV2Schemas(ctx context.Context, params *GetScimV2SchemasParams) (*GetScimV2SchemasOK, error)
	/*
	   GetScimV2Serviceproviderconfig gets a service provider s configuration
	*/
	GetScimV2Serviceproviderconfig(ctx context.Context, params *GetScimV2ServiceproviderconfigParams) (*GetScimV2ServiceproviderconfigOK, error)
	/*
	   GetScimV2User gets a user
	*/
	GetScimV2User(ctx context.Context, params *GetScimV2UserParams) (*GetScimV2UserOK, error)
	/*
	   GetScimV2Users gets a list of users
	   To return all active users, do not use the filter parameter. To return inactive users, set the filter parameter to "active eq false". By default, returns SCIM attributes "externalId", "enterprise-user:manager", and "roles". To exclude these attributes, set the attributes parameter to "id,active" or the excludeAttributes parameter to "externalId,roles,urn:ietf:params:scim:schemas:extension:enterprise:2.0:User:division".
	*/
	GetScimV2Users(ctx context.Context, params *GetScimV2UsersParams) (*GetScimV2UsersOK, error)
	/*
	   PatchScimGroup modifies a group
	*/
	PatchScimGroup(ctx context.Context, params *PatchScimGroupParams) (*PatchScimGroupOK, error)
	/*
	   PatchScimUser modifies a user
	*/
	PatchScimUser(ctx context.Context, params *PatchScimUserParams) (*PatchScimUserOK, error)
	/*
	   PatchScimV2Group modifies a group
	*/
	PatchScimV2Group(ctx context.Context, params *PatchScimV2GroupParams) (*PatchScimV2GroupOK, error)
	/*
	   PatchScimV2User modifies a user
	*/
	PatchScimV2User(ctx context.Context, params *PatchScimV2UserParams) (*PatchScimV2UserOK, error)
	/*
	   PostScimGroups creates a group
	   Creates a Genesys Cloud group with group visibility set to "public" and rules visibility set to "true". Auto-creates an "externalId". "externalId" is used to determine if DELETE /api/v2/scim/groups/{groupId} or DELETE /api/v2/scim/v2/groups/{groupId} is allowed.
	*/
	PostScimGroups(ctx context.Context, params *PostScimGroupsParams) (*PostScimGroupsOK, error)
	/*
	   PostScimUsers creates a user
	*/
	PostScimUsers(ctx context.Context, params *PostScimUsersParams) (*PostScimUsersOK, *PostScimUsersCreated, error)
	/*
	   PostScimV2Groups creates a group
	   Creates an "official" Genesys Cloud group with group visibility set to "public" and rules visibility set to "true". Auto-creates an "externalId". "externalId" is used to determine if DELETE /api/v2/scim/groups/{groupId} or DELETE /api/v2/scim/v2/groups/{groupId} should be allowed.
	*/
	PostScimV2Groups(ctx context.Context, params *PostScimV2GroupsParams) (*PostScimV2GroupsOK, error)
	/*
	   PostScimV2Users creates a user
	*/
	PostScimV2Users(ctx context.Context, params *PostScimV2UsersParams) (*PostScimV2UsersOK, *PostScimV2UsersCreated, error)
	/*
	   PutScimGroup replaces a group
	*/
	PutScimGroup(ctx context.Context, params *PutScimGroupParams) (*PutScimGroupOK, error)
	/*
	   PutScimUser replaces a user
	*/
	PutScimUser(ctx context.Context, params *PutScimUserParams) (*PutScimUserOK, error)
	/*
	   PutScimV2Group replaces a group
	*/
	PutScimV2Group(ctx context.Context, params *PutScimV2GroupParams) (*PutScimV2GroupOK, error)
	/*
	   PutScimV2User replaces a user
	*/
	PutScimV2User(ctx context.Context, params *PutScimV2UserParams) (*PutScimV2UserOK, error)
}

// New creates a new s c i m API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, authInfo runtime.ClientAuthInfoWriter) *Client {
	return &Client{
		transport: transport,
		formats:   formats,
		authInfo:  authInfo,
	}
}

/*
Client for s c i m API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	authInfo  runtime.ClientAuthInfoWriter
}

/*
DeleteScimGroup deletes a group
*/
func (a *Client) DeleteScimGroup(ctx context.Context, params *DeleteScimGroupParams) (*DeleteScimGroupNoContent, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteScimGroup",
		Method:             "DELETE",
		PathPattern:        "/api/v2/scim/groups/{groupId}",
		ProducesMediaTypes: []string{"application/json", "application/scim+json"},
		ConsumesMediaTypes: []string{"application/json", "application/scim+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteScimGroupReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteScimGroupNoContent), nil

}

/*
DeleteScimUser deletes a user
*/
func (a *Client) DeleteScimUser(ctx context.Context, params *DeleteScimUserParams) (*DeleteScimUserOK, *DeleteScimUserNoContent, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteScimUser",
		Method:             "DELETE",
		PathPattern:        "/api/v2/scim/users/{userId}",
		ProducesMediaTypes: []string{"application/json", "application/scim+json"},
		ConsumesMediaTypes: []string{"application/json", "application/scim+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteScimUserReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *DeleteScimUserOK:
		return value, nil, nil
	case *DeleteScimUserNoContent:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
DeleteScimV2Group deletes a group
*/
func (a *Client) DeleteScimV2Group(ctx context.Context, params *DeleteScimV2GroupParams) (*DeleteScimV2GroupNoContent, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteScimV2Group",
		Method:             "DELETE",
		PathPattern:        "/api/v2/scim/v2/groups/{groupId}",
		ProducesMediaTypes: []string{"application/json", "application/scim+json"},
		ConsumesMediaTypes: []string{"application/json", "application/scim+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteScimV2GroupReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteScimV2GroupNoContent), nil

}

/*
DeleteScimV2User deletes a user
*/
func (a *Client) DeleteScimV2User(ctx context.Context, params *DeleteScimV2UserParams) (*DeleteScimV2UserOK, *DeleteScimV2UserNoContent, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteScimV2User",
		Method:             "DELETE",
		PathPattern:        "/api/v2/scim/v2/users/{userId}",
		ProducesMediaTypes: []string{"application/json", "application/scim+json"},
		ConsumesMediaTypes: []string{"application/json", "application/scim+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteScimV2UserReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *DeleteScimV2UserOK:
		return value, nil, nil
	case *DeleteScimV2UserNoContent:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
GetScimGroup gets a group
*/
func (a *Client) GetScimGroup(ctx context.Context, params *GetScimGroupParams) (*GetScimGroupOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getScimGroup",
		Method:             "GET",
		PathPattern:        "/api/v2/scim/groups/{groupId}",
		ProducesMediaTypes: []string{"application/json", "application/scim+json"},
		ConsumesMediaTypes: []string{"application/json", "application/scim+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetScimGroupReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetScimGroupOK), nil

}

/*
GetScimGroups gets a list of groups
*/
func (a *Client) GetScimGroups(ctx context.Context, params *GetScimGroupsParams) (*GetScimGroupsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getScimGroups",
		Method:             "GET",
		PathPattern:        "/api/v2/scim/groups",
		ProducesMediaTypes: []string{"application/json", "application/scim+json"},
		ConsumesMediaTypes: []string{"application/json", "application/scim+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetScimGroupsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetScimGroupsOK), nil

}

/*
GetScimResourcetype gets a resource type
*/
func (a *Client) GetScimResourcetype(ctx context.Context, params *GetScimResourcetypeParams) (*GetScimResourcetypeOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getScimResourcetype",
		Method:             "GET",
		PathPattern:        "/api/v2/scim/resourcetypes/{resourceType}",
		ProducesMediaTypes: []string{"application/json", "application/scim+json"},
		ConsumesMediaTypes: []string{"application/json", "application/scim+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetScimResourcetypeReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetScimResourcetypeOK), nil

}

/*
GetScimResourcetypes gets a list of resource types
*/
func (a *Client) GetScimResourcetypes(ctx context.Context, params *GetScimResourcetypesParams) (*GetScimResourcetypesOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getScimResourcetypes",
		Method:             "GET",
		PathPattern:        "/api/v2/scim/resourcetypes",
		ProducesMediaTypes: []string{"application/json", "application/scim+json"},
		ConsumesMediaTypes: []string{"application/json", "application/scim+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetScimResourcetypesReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetScimResourcetypesOK), nil

}

/*
GetScimSchema gets a s c i m schema
*/
func (a *Client) GetScimSchema(ctx context.Context, params *GetScimSchemaParams) (*GetScimSchemaOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getScimSchema",
		Method:             "GET",
		PathPattern:        "/api/v2/scim/schemas/{schemaId}",
		ProducesMediaTypes: []string{"application/json", "application/scim+json"},
		ConsumesMediaTypes: []string{"application/json", "application/scim+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetScimSchemaReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetScimSchemaOK), nil

}

/*
GetScimSchemas gets a list of s c i m schemas
*/
func (a *Client) GetScimSchemas(ctx context.Context, params *GetScimSchemasParams) (*GetScimSchemasOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getScimSchemas",
		Method:             "GET",
		PathPattern:        "/api/v2/scim/schemas",
		ProducesMediaTypes: []string{"application/json", "application/scim+json"},
		ConsumesMediaTypes: []string{"application/json", "application/scim+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetScimSchemasReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetScimSchemasOK), nil

}

/*
GetScimServiceproviderconfig gets a service provider s configuration
*/
func (a *Client) GetScimServiceproviderconfig(ctx context.Context, params *GetScimServiceproviderconfigParams) (*GetScimServiceproviderconfigOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getScimServiceproviderconfig",
		Method:             "GET",
		PathPattern:        "/api/v2/scim/serviceproviderconfig",
		ProducesMediaTypes: []string{"application/json", "application/scim+json"},
		ConsumesMediaTypes: []string{"application/json", "application/scim+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetScimServiceproviderconfigReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetScimServiceproviderconfigOK), nil

}

/*
GetScimUser gets a user
*/
func (a *Client) GetScimUser(ctx context.Context, params *GetScimUserParams) (*GetScimUserOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getScimUser",
		Method:             "GET",
		PathPattern:        "/api/v2/scim/users/{userId}",
		ProducesMediaTypes: []string{"application/json", "application/scim+json"},
		ConsumesMediaTypes: []string{"application/json", "application/scim+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetScimUserReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetScimUserOK), nil

}

/*
GetScimUsers gets a list of users

To return all active users, do not use the filter parameter. To return inactive users, set the filter parameter to "active eq false". By default, returns SCIM attributes "externalId", "enterprise-user:manager", and "roles". To exclude these attributes, set the attributes parameter to "id,active" or the excludeAttributes parameter to "externalId,roles,urn:ietf:params:scim:schemas:extension:enterprise:2.0:User:division".
*/
func (a *Client) GetScimUsers(ctx context.Context, params *GetScimUsersParams) (*GetScimUsersOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getScimUsers",
		Method:             "GET",
		PathPattern:        "/api/v2/scim/users",
		ProducesMediaTypes: []string{"application/json", "application/scim+json"},
		ConsumesMediaTypes: []string{"application/json", "application/scim+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetScimUsersReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetScimUsersOK), nil

}

/*
GetScimV2Group gets a group
*/
func (a *Client) GetScimV2Group(ctx context.Context, params *GetScimV2GroupParams) (*GetScimV2GroupOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getScimV2Group",
		Method:             "GET",
		PathPattern:        "/api/v2/scim/v2/groups/{groupId}",
		ProducesMediaTypes: []string{"application/json", "application/scim+json"},
		ConsumesMediaTypes: []string{"application/json", "application/scim+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetScimV2GroupReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetScimV2GroupOK), nil

}

/*
GetScimV2Groups gets a list of groups
*/
func (a *Client) GetScimV2Groups(ctx context.Context, params *GetScimV2GroupsParams) (*GetScimV2GroupsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getScimV2Groups",
		Method:             "GET",
		PathPattern:        "/api/v2/scim/v2/groups",
		ProducesMediaTypes: []string{"application/json", "application/scim+json"},
		ConsumesMediaTypes: []string{"application/json", "application/scim+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetScimV2GroupsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetScimV2GroupsOK), nil

}

/*
GetScimV2Resourcetype gets a resource type
*/
func (a *Client) GetScimV2Resourcetype(ctx context.Context, params *GetScimV2ResourcetypeParams) (*GetScimV2ResourcetypeOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getScimV2Resourcetype",
		Method:             "GET",
		PathPattern:        "/api/v2/scim/v2/resourcetypes/{resourceType}",
		ProducesMediaTypes: []string{"application/json", "application/scim+json"},
		ConsumesMediaTypes: []string{"application/json", "application/scim+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetScimV2ResourcetypeReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetScimV2ResourcetypeOK), nil

}

/*
GetScimV2Resourcetypes gets a list of resource types
*/
func (a *Client) GetScimV2Resourcetypes(ctx context.Context, params *GetScimV2ResourcetypesParams) (*GetScimV2ResourcetypesOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getScimV2Resourcetypes",
		Method:             "GET",
		PathPattern:        "/api/v2/scim/v2/resourcetypes",
		ProducesMediaTypes: []string{"application/json", "application/scim+json"},
		ConsumesMediaTypes: []string{"application/json", "application/scim+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetScimV2ResourcetypesReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetScimV2ResourcetypesOK), nil

}

/*
GetScimV2Schema gets a s c i m schema
*/
func (a *Client) GetScimV2Schema(ctx context.Context, params *GetScimV2SchemaParams) (*GetScimV2SchemaOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getScimV2Schema",
		Method:             "GET",
		PathPattern:        "/api/v2/scim/v2/schemas/{schemaId}",
		ProducesMediaTypes: []string{"application/json", "application/scim+json"},
		ConsumesMediaTypes: []string{"application/json", "application/scim+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetScimV2SchemaReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetScimV2SchemaOK), nil

}

/*
GetScimV2Schemas gets a list of s c i m schemas
*/
func (a *Client) GetScimV2Schemas(ctx context.Context, params *GetScimV2SchemasParams) (*GetScimV2SchemasOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getScimV2Schemas",
		Method:             "GET",
		PathPattern:        "/api/v2/scim/v2/schemas",
		ProducesMediaTypes: []string{"application/json", "application/scim+json"},
		ConsumesMediaTypes: []string{"application/json", "application/scim+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetScimV2SchemasReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetScimV2SchemasOK), nil

}

/*
GetScimV2Serviceproviderconfig gets a service provider s configuration
*/
func (a *Client) GetScimV2Serviceproviderconfig(ctx context.Context, params *GetScimV2ServiceproviderconfigParams) (*GetScimV2ServiceproviderconfigOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getScimV2Serviceproviderconfig",
		Method:             "GET",
		PathPattern:        "/api/v2/scim/v2/serviceproviderconfig",
		ProducesMediaTypes: []string{"application/json", "application/scim+json"},
		ConsumesMediaTypes: []string{"application/json", "application/scim+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetScimV2ServiceproviderconfigReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetScimV2ServiceproviderconfigOK), nil

}

/*
GetScimV2User gets a user
*/
func (a *Client) GetScimV2User(ctx context.Context, params *GetScimV2UserParams) (*GetScimV2UserOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getScimV2User",
		Method:             "GET",
		PathPattern:        "/api/v2/scim/v2/users/{userId}",
		ProducesMediaTypes: []string{"application/json", "application/scim+json"},
		ConsumesMediaTypes: []string{"application/json", "application/scim+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetScimV2UserReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetScimV2UserOK), nil

}

/*
GetScimV2Users gets a list of users

To return all active users, do not use the filter parameter. To return inactive users, set the filter parameter to "active eq false". By default, returns SCIM attributes "externalId", "enterprise-user:manager", and "roles". To exclude these attributes, set the attributes parameter to "id,active" or the excludeAttributes parameter to "externalId,roles,urn:ietf:params:scim:schemas:extension:enterprise:2.0:User:division".
*/
func (a *Client) GetScimV2Users(ctx context.Context, params *GetScimV2UsersParams) (*GetScimV2UsersOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getScimV2Users",
		Method:             "GET",
		PathPattern:        "/api/v2/scim/v2/users",
		ProducesMediaTypes: []string{"application/json", "application/scim+json"},
		ConsumesMediaTypes: []string{"application/json", "application/scim+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetScimV2UsersReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetScimV2UsersOK), nil

}

/*
PatchScimGroup modifies a group
*/
func (a *Client) PatchScimGroup(ctx context.Context, params *PatchScimGroupParams) (*PatchScimGroupOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "patchScimGroup",
		Method:             "PATCH",
		PathPattern:        "/api/v2/scim/groups/{groupId}",
		ProducesMediaTypes: []string{"application/json", "application/scim+json"},
		ConsumesMediaTypes: []string{"application/json", "application/scim+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchScimGroupReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PatchScimGroupOK), nil

}

/*
PatchScimUser modifies a user
*/
func (a *Client) PatchScimUser(ctx context.Context, params *PatchScimUserParams) (*PatchScimUserOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "patchScimUser",
		Method:             "PATCH",
		PathPattern:        "/api/v2/scim/users/{userId}",
		ProducesMediaTypes: []string{"application/json", "application/scim+json"},
		ConsumesMediaTypes: []string{"application/json", "application/scim+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchScimUserReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PatchScimUserOK), nil

}

/*
PatchScimV2Group modifies a group
*/
func (a *Client) PatchScimV2Group(ctx context.Context, params *PatchScimV2GroupParams) (*PatchScimV2GroupOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "patchScimV2Group",
		Method:             "PATCH",
		PathPattern:        "/api/v2/scim/v2/groups/{groupId}",
		ProducesMediaTypes: []string{"application/json", "application/scim+json"},
		ConsumesMediaTypes: []string{"application/json", "application/scim+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchScimV2GroupReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PatchScimV2GroupOK), nil

}

/*
PatchScimV2User modifies a user
*/
func (a *Client) PatchScimV2User(ctx context.Context, params *PatchScimV2UserParams) (*PatchScimV2UserOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "patchScimV2User",
		Method:             "PATCH",
		PathPattern:        "/api/v2/scim/v2/users/{userId}",
		ProducesMediaTypes: []string{"application/json", "application/scim+json"},
		ConsumesMediaTypes: []string{"application/json", "application/scim+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchScimV2UserReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PatchScimV2UserOK), nil

}

/*
PostScimGroups creates a group

Creates a Genesys Cloud group with group visibility set to "public" and rules visibility set to "true". Auto-creates an "externalId". "externalId" is used to determine if DELETE /api/v2/scim/groups/{groupId} or DELETE /api/v2/scim/v2/groups/{groupId} is allowed.
*/
func (a *Client) PostScimGroups(ctx context.Context, params *PostScimGroupsParams) (*PostScimGroupsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postScimGroups",
		Method:             "POST",
		PathPattern:        "/api/v2/scim/groups",
		ProducesMediaTypes: []string{"application/json", "application/scim+json"},
		ConsumesMediaTypes: []string{"application/json", "application/scim+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostScimGroupsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostScimGroupsOK), nil

}

/*
PostScimUsers creates a user
*/
func (a *Client) PostScimUsers(ctx context.Context, params *PostScimUsersParams) (*PostScimUsersOK, *PostScimUsersCreated, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postScimUsers",
		Method:             "POST",
		PathPattern:        "/api/v2/scim/users",
		ProducesMediaTypes: []string{"application/json", "application/scim+json"},
		ConsumesMediaTypes: []string{"application/json", "application/scim+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostScimUsersReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *PostScimUsersOK:
		return value, nil, nil
	case *PostScimUsersCreated:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
PostScimV2Groups creates a group

Creates an "official" Genesys Cloud group with group visibility set to "public" and rules visibility set to "true". Auto-creates an "externalId". "externalId" is used to determine if DELETE /api/v2/scim/groups/{groupId} or DELETE /api/v2/scim/v2/groups/{groupId} should be allowed.
*/
func (a *Client) PostScimV2Groups(ctx context.Context, params *PostScimV2GroupsParams) (*PostScimV2GroupsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postScimV2Groups",
		Method:             "POST",
		PathPattern:        "/api/v2/scim/v2/groups",
		ProducesMediaTypes: []string{"application/json", "application/scim+json"},
		ConsumesMediaTypes: []string{"application/json", "application/scim+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostScimV2GroupsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostScimV2GroupsOK), nil

}

/*
PostScimV2Users creates a user
*/
func (a *Client) PostScimV2Users(ctx context.Context, params *PostScimV2UsersParams) (*PostScimV2UsersOK, *PostScimV2UsersCreated, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postScimV2Users",
		Method:             "POST",
		PathPattern:        "/api/v2/scim/v2/users",
		ProducesMediaTypes: []string{"application/json", "application/scim+json"},
		ConsumesMediaTypes: []string{"application/json", "application/scim+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostScimV2UsersReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *PostScimV2UsersOK:
		return value, nil, nil
	case *PostScimV2UsersCreated:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
PutScimGroup replaces a group
*/
func (a *Client) PutScimGroup(ctx context.Context, params *PutScimGroupParams) (*PutScimGroupOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "putScimGroup",
		Method:             "PUT",
		PathPattern:        "/api/v2/scim/groups/{groupId}",
		ProducesMediaTypes: []string{"application/json", "application/scim+json"},
		ConsumesMediaTypes: []string{"application/json", "application/scim+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutScimGroupReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutScimGroupOK), nil

}

/*
PutScimUser replaces a user
*/
func (a *Client) PutScimUser(ctx context.Context, params *PutScimUserParams) (*PutScimUserOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "putScimUser",
		Method:             "PUT",
		PathPattern:        "/api/v2/scim/users/{userId}",
		ProducesMediaTypes: []string{"application/json", "application/scim+json"},
		ConsumesMediaTypes: []string{"application/json", "application/scim+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutScimUserReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutScimUserOK), nil

}

/*
PutScimV2Group replaces a group
*/
func (a *Client) PutScimV2Group(ctx context.Context, params *PutScimV2GroupParams) (*PutScimV2GroupOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "putScimV2Group",
		Method:             "PUT",
		PathPattern:        "/api/v2/scim/v2/groups/{groupId}",
		ProducesMediaTypes: []string{"application/json", "application/scim+json"},
		ConsumesMediaTypes: []string{"application/json", "application/scim+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutScimV2GroupReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutScimV2GroupOK), nil

}

/*
PutScimV2User replaces a user
*/
func (a *Client) PutScimV2User(ctx context.Context, params *PutScimV2UserParams) (*PutScimV2UserOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "putScimV2User",
		Method:             "PUT",
		PathPattern:        "/api/v2/scim/v2/users/{userId}",
		ProducesMediaTypes: []string{"application/json", "application/scim+json"},
		ConsumesMediaTypes: []string{"application/json", "application/scim+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutScimV2UserReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutScimV2UserOK), nil

}

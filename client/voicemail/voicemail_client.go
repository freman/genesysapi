// Code generated by go-swagger; DO NOT EDIT.

package voicemail

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

//go:generate mockery -name API -inpkg

// API is the interface of the voicemail client
type API interface {
	/*
	   DeleteVoicemailMessage deletes a voicemail message
	   A user voicemail can only be deleted by its associated user. A group voicemail can only be deleted by a user that is a member of the group. A queue voicemail can only be deleted by a user with the acd voicemail delete permission.
	*/
	DeleteVoicemailMessage(ctx context.Context, params *DeleteVoicemailMessageParams) (*DeleteVoicemailMessageOK, error)
	/*
	   DeleteVoicemailMessages deletes all voicemail messages
	*/
	DeleteVoicemailMessages(ctx context.Context, params *DeleteVoicemailMessagesParams) (*DeleteVoicemailMessagesOK, error)
	/*
	   GetVoicemailGroupMailbox gets the group s mailbox information
	*/
	GetVoicemailGroupMailbox(ctx context.Context, params *GetVoicemailGroupMailboxParams) (*GetVoicemailGroupMailboxOK, error)
	/*
	   GetVoicemailGroupMessages lists voicemail messages
	*/
	GetVoicemailGroupMessages(ctx context.Context, params *GetVoicemailGroupMessagesParams) (*GetVoicemailGroupMessagesOK, error)
	/*
	   GetVoicemailGroupPolicy gets a group s voicemail policy
	*/
	GetVoicemailGroupPolicy(ctx context.Context, params *GetVoicemailGroupPolicyParams) (*GetVoicemailGroupPolicyOK, error)
	/*
	   GetVoicemailMailbox gets the current user s mailbox information
	*/
	GetVoicemailMailbox(ctx context.Context, params *GetVoicemailMailboxParams) (*GetVoicemailMailboxOK, error)
	/*
	   GetVoicemailMeMailbox gets the current user s mailbox information
	*/
	GetVoicemailMeMailbox(ctx context.Context, params *GetVoicemailMeMailboxParams) (*GetVoicemailMeMailboxOK, error)
	/*
	   GetVoicemailMeMessages lists voicemail messages
	*/
	GetVoicemailMeMessages(ctx context.Context, params *GetVoicemailMeMessagesParams) (*GetVoicemailMeMessagesOK, error)
	/*
	   GetVoicemailMePolicy gets the current user s voicemail policy
	*/
	GetVoicemailMePolicy(ctx context.Context, params *GetVoicemailMePolicyParams) (*GetVoicemailMePolicyOK, error)
	/*
	   GetVoicemailMessage gets a voicemail message
	*/
	GetVoicemailMessage(ctx context.Context, params *GetVoicemailMessageParams) (*GetVoicemailMessageOK, error)
	/*
	   GetVoicemailMessageMedia gets media playback URI for this voicemail message
	*/
	GetVoicemailMessageMedia(ctx context.Context, params *GetVoicemailMessageMediaParams) (*GetVoicemailMessageMediaOK, error)
	/*
	   GetVoicemailMessages lists voicemail messages
	*/
	GetVoicemailMessages(ctx context.Context, params *GetVoicemailMessagesParams) (*GetVoicemailMessagesOK, error)
	/*
	   GetVoicemailPolicy gets a policy
	*/
	GetVoicemailPolicy(ctx context.Context, params *GetVoicemailPolicyParams) (*GetVoicemailPolicyOK, error)
	/*
	   GetVoicemailQueueMessages lists voicemail messages
	*/
	GetVoicemailQueueMessages(ctx context.Context, params *GetVoicemailQueueMessagesParams) (*GetVoicemailQueueMessagesOK, error)
	/*
	   GetVoicemailUserpolicy gets a user s voicemail policy
	*/
	GetVoicemailUserpolicy(ctx context.Context, params *GetVoicemailUserpolicyParams) (*GetVoicemailUserpolicyOK, error)
	/*
	   PatchVoicemailGroupPolicy updates a group s voicemail policy
	*/
	PatchVoicemailGroupPolicy(ctx context.Context, params *PatchVoicemailGroupPolicyParams) (*PatchVoicemailGroupPolicyOK, error)
	/*
	   PatchVoicemailMePolicy updates the current user s voicemail policy
	*/
	PatchVoicemailMePolicy(ctx context.Context, params *PatchVoicemailMePolicyParams) (*PatchVoicemailMePolicyOK, error)
	/*
	   PatchVoicemailMessage updates a voicemail message
	   A user voicemail can only be modified by its associated user. A group voicemail can only be modified by a user that is a member of the group. A queue voicemail can only be modified by a participant of the conversation the voicemail is associated with.
	*/
	PatchVoicemailMessage(ctx context.Context, params *PatchVoicemailMessageParams) (*PatchVoicemailMessageOK, error)
	/*
	   PatchVoicemailUserpolicy updates a user s voicemail policy
	*/
	PatchVoicemailUserpolicy(ctx context.Context, params *PatchVoicemailUserpolicyParams) (*PatchVoicemailUserpolicyOK, error)
	/*
	   PostVoicemailMessages copies a voicemail message to a user or group
	*/
	PostVoicemailMessages(ctx context.Context, params *PostVoicemailMessagesParams) (*PostVoicemailMessagesOK, error)
	/*
	   PutVoicemailMessage updates a voicemail message
	   A user voicemail can only be modified by its associated user. A group voicemail can only be modified by a user that is a member of the group. A queue voicemail can only be modified by a participant of the conversation the voicemail is associated with.
	*/
	PutVoicemailMessage(ctx context.Context, params *PutVoicemailMessageParams) (*PutVoicemailMessageOK, error)
	/*
	   PutVoicemailPolicy updates a policy
	*/
	PutVoicemailPolicy(ctx context.Context, params *PutVoicemailPolicyParams) (*PutVoicemailPolicyOK, error)
}

// New creates a new voicemail API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, authInfo runtime.ClientAuthInfoWriter) *Client {
	return &Client{
		transport: transport,
		formats:   formats,
		authInfo:  authInfo,
	}
}

/*
Client for voicemail API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	authInfo  runtime.ClientAuthInfoWriter
}

/*
DeleteVoicemailMessage deletes a voicemail message

A user voicemail can only be deleted by its associated user. A group voicemail can only be deleted by a user that is a member of the group. A queue voicemail can only be deleted by a user with the acd voicemail delete permission.
*/
func (a *Client) DeleteVoicemailMessage(ctx context.Context, params *DeleteVoicemailMessageParams) (*DeleteVoicemailMessageOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteVoicemailMessage",
		Method:             "DELETE",
		PathPattern:        "/api/v2/voicemail/messages/{messageId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteVoicemailMessageReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteVoicemailMessageOK), nil

}

/*
DeleteVoicemailMessages deletes all voicemail messages
*/
func (a *Client) DeleteVoicemailMessages(ctx context.Context, params *DeleteVoicemailMessagesParams) (*DeleteVoicemailMessagesOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteVoicemailMessages",
		Method:             "DELETE",
		PathPattern:        "/api/v2/voicemail/messages",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteVoicemailMessagesReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteVoicemailMessagesOK), nil

}

/*
GetVoicemailGroupMailbox gets the group s mailbox information
*/
func (a *Client) GetVoicemailGroupMailbox(ctx context.Context, params *GetVoicemailGroupMailboxParams) (*GetVoicemailGroupMailboxOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getVoicemailGroupMailbox",
		Method:             "GET",
		PathPattern:        "/api/v2/voicemail/groups/{groupId}/mailbox",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetVoicemailGroupMailboxReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetVoicemailGroupMailboxOK), nil

}

/*
GetVoicemailGroupMessages lists voicemail messages
*/
func (a *Client) GetVoicemailGroupMessages(ctx context.Context, params *GetVoicemailGroupMessagesParams) (*GetVoicemailGroupMessagesOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getVoicemailGroupMessages",
		Method:             "GET",
		PathPattern:        "/api/v2/voicemail/groups/{groupId}/messages",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetVoicemailGroupMessagesReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetVoicemailGroupMessagesOK), nil

}

/*
GetVoicemailGroupPolicy gets a group s voicemail policy
*/
func (a *Client) GetVoicemailGroupPolicy(ctx context.Context, params *GetVoicemailGroupPolicyParams) (*GetVoicemailGroupPolicyOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getVoicemailGroupPolicy",
		Method:             "GET",
		PathPattern:        "/api/v2/voicemail/groups/{groupId}/policy",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetVoicemailGroupPolicyReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetVoicemailGroupPolicyOK), nil

}

/*
GetVoicemailMailbox gets the current user s mailbox information
*/
func (a *Client) GetVoicemailMailbox(ctx context.Context, params *GetVoicemailMailboxParams) (*GetVoicemailMailboxOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getVoicemailMailbox",
		Method:             "GET",
		PathPattern:        "/api/v2/voicemail/mailbox",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetVoicemailMailboxReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetVoicemailMailboxOK), nil

}

/*
GetVoicemailMeMailbox gets the current user s mailbox information
*/
func (a *Client) GetVoicemailMeMailbox(ctx context.Context, params *GetVoicemailMeMailboxParams) (*GetVoicemailMeMailboxOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getVoicemailMeMailbox",
		Method:             "GET",
		PathPattern:        "/api/v2/voicemail/me/mailbox",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetVoicemailMeMailboxReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetVoicemailMeMailboxOK), nil

}

/*
GetVoicemailMeMessages lists voicemail messages
*/
func (a *Client) GetVoicemailMeMessages(ctx context.Context, params *GetVoicemailMeMessagesParams) (*GetVoicemailMeMessagesOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getVoicemailMeMessages",
		Method:             "GET",
		PathPattern:        "/api/v2/voicemail/me/messages",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetVoicemailMeMessagesReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetVoicemailMeMessagesOK), nil

}

/*
GetVoicemailMePolicy gets the current user s voicemail policy
*/
func (a *Client) GetVoicemailMePolicy(ctx context.Context, params *GetVoicemailMePolicyParams) (*GetVoicemailMePolicyOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getVoicemailMePolicy",
		Method:             "GET",
		PathPattern:        "/api/v2/voicemail/me/policy",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetVoicemailMePolicyReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetVoicemailMePolicyOK), nil

}

/*
GetVoicemailMessage gets a voicemail message
*/
func (a *Client) GetVoicemailMessage(ctx context.Context, params *GetVoicemailMessageParams) (*GetVoicemailMessageOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getVoicemailMessage",
		Method:             "GET",
		PathPattern:        "/api/v2/voicemail/messages/{messageId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetVoicemailMessageReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetVoicemailMessageOK), nil

}

/*
GetVoicemailMessageMedia gets media playback URI for this voicemail message
*/
func (a *Client) GetVoicemailMessageMedia(ctx context.Context, params *GetVoicemailMessageMediaParams) (*GetVoicemailMessageMediaOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getVoicemailMessageMedia",
		Method:             "GET",
		PathPattern:        "/api/v2/voicemail/messages/{messageId}/media",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetVoicemailMessageMediaReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetVoicemailMessageMediaOK), nil

}

/*
GetVoicemailMessages lists voicemail messages
*/
func (a *Client) GetVoicemailMessages(ctx context.Context, params *GetVoicemailMessagesParams) (*GetVoicemailMessagesOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getVoicemailMessages",
		Method:             "GET",
		PathPattern:        "/api/v2/voicemail/messages",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetVoicemailMessagesReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetVoicemailMessagesOK), nil

}

/*
GetVoicemailPolicy gets a policy
*/
func (a *Client) GetVoicemailPolicy(ctx context.Context, params *GetVoicemailPolicyParams) (*GetVoicemailPolicyOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getVoicemailPolicy",
		Method:             "GET",
		PathPattern:        "/api/v2/voicemail/policy",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetVoicemailPolicyReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetVoicemailPolicyOK), nil

}

/*
GetVoicemailQueueMessages lists voicemail messages
*/
func (a *Client) GetVoicemailQueueMessages(ctx context.Context, params *GetVoicemailQueueMessagesParams) (*GetVoicemailQueueMessagesOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getVoicemailQueueMessages",
		Method:             "GET",
		PathPattern:        "/api/v2/voicemail/queues/{queueId}/messages",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetVoicemailQueueMessagesReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetVoicemailQueueMessagesOK), nil

}

/*
GetVoicemailUserpolicy gets a user s voicemail policy
*/
func (a *Client) GetVoicemailUserpolicy(ctx context.Context, params *GetVoicemailUserpolicyParams) (*GetVoicemailUserpolicyOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getVoicemailUserpolicy",
		Method:             "GET",
		PathPattern:        "/api/v2/voicemail/userpolicies/{userId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetVoicemailUserpolicyReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetVoicemailUserpolicyOK), nil

}

/*
PatchVoicemailGroupPolicy updates a group s voicemail policy
*/
func (a *Client) PatchVoicemailGroupPolicy(ctx context.Context, params *PatchVoicemailGroupPolicyParams) (*PatchVoicemailGroupPolicyOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "patchVoicemailGroupPolicy",
		Method:             "PATCH",
		PathPattern:        "/api/v2/voicemail/groups/{groupId}/policy",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchVoicemailGroupPolicyReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PatchVoicemailGroupPolicyOK), nil

}

/*
PatchVoicemailMePolicy updates the current user s voicemail policy
*/
func (a *Client) PatchVoicemailMePolicy(ctx context.Context, params *PatchVoicemailMePolicyParams) (*PatchVoicemailMePolicyOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "patchVoicemailMePolicy",
		Method:             "PATCH",
		PathPattern:        "/api/v2/voicemail/me/policy",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchVoicemailMePolicyReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PatchVoicemailMePolicyOK), nil

}

/*
PatchVoicemailMessage updates a voicemail message

A user voicemail can only be modified by its associated user. A group voicemail can only be modified by a user that is a member of the group. A queue voicemail can only be modified by a participant of the conversation the voicemail is associated with.
*/
func (a *Client) PatchVoicemailMessage(ctx context.Context, params *PatchVoicemailMessageParams) (*PatchVoicemailMessageOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "patchVoicemailMessage",
		Method:             "PATCH",
		PathPattern:        "/api/v2/voicemail/messages/{messageId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchVoicemailMessageReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PatchVoicemailMessageOK), nil

}

/*
PatchVoicemailUserpolicy updates a user s voicemail policy
*/
func (a *Client) PatchVoicemailUserpolicy(ctx context.Context, params *PatchVoicemailUserpolicyParams) (*PatchVoicemailUserpolicyOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "patchVoicemailUserpolicy",
		Method:             "PATCH",
		PathPattern:        "/api/v2/voicemail/userpolicies/{userId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchVoicemailUserpolicyReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PatchVoicemailUserpolicyOK), nil

}

/*
PostVoicemailMessages copies a voicemail message to a user or group
*/
func (a *Client) PostVoicemailMessages(ctx context.Context, params *PostVoicemailMessagesParams) (*PostVoicemailMessagesOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postVoicemailMessages",
		Method:             "POST",
		PathPattern:        "/api/v2/voicemail/messages",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostVoicemailMessagesReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostVoicemailMessagesOK), nil

}

/*
PutVoicemailMessage updates a voicemail message

A user voicemail can only be modified by its associated user. A group voicemail can only be modified by a user that is a member of the group. A queue voicemail can only be modified by a participant of the conversation the voicemail is associated with.
*/
func (a *Client) PutVoicemailMessage(ctx context.Context, params *PutVoicemailMessageParams) (*PutVoicemailMessageOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "putVoicemailMessage",
		Method:             "PUT",
		PathPattern:        "/api/v2/voicemail/messages/{messageId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutVoicemailMessageReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutVoicemailMessageOK), nil

}

/*
PutVoicemailPolicy updates a policy
*/
func (a *Client) PutVoicemailPolicy(ctx context.Context, params *PutVoicemailPolicyParams) (*PutVoicemailPolicyOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "putVoicemailPolicy",
		Method:             "PUT",
		PathPattern:        "/api/v2/voicemail/policy",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutVoicemailPolicyReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutVoicemailPolicyOK), nil

}

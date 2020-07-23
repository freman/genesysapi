// Code generated by go-swagger; DO NOT EDIT.

package coaching

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

//go:generate mockery -name API -inpkg

// API is the interface of the coaching client
type API interface {
	/*
	   DeleteCoachingAppointment deletes an existing appointment
	   Permission not required if you are the creator of the appointment
	*/
	DeleteCoachingAppointment(ctx context.Context, params *DeleteCoachingAppointmentParams) (*DeleteCoachingAppointmentNoContent, error)
	/*
	   DeleteCoachingAppointmentAnnotation deletes an existing annotation
	   You must have the appropriate permission for the type of annotation you are updating. Permission not required if you are the creator or facilitator of the appointment
	*/
	DeleteCoachingAppointmentAnnotation(ctx context.Context, params *DeleteCoachingAppointmentAnnotationParams) (*DeleteCoachingAppointmentAnnotationNoContent, error)
	/*
	   GetCoachingAppointment retrieves an appointment
	   Permission not required if you are the attendee, creator or facilitator of the appointment
	*/
	GetCoachingAppointment(ctx context.Context, params *GetCoachingAppointmentParams) (*GetCoachingAppointmentOK, error)
	/*
	   GetCoachingAppointmentAnnotation retrieves an annotation
	   You must have the appropriate permission for the type of annotation you are creating. Permission not required if you are related to the appointment (only the creator or facilitator can view private annotations).
	*/
	GetCoachingAppointmentAnnotation(ctx context.Context, params *GetCoachingAppointmentAnnotationParams) (*GetCoachingAppointmentAnnotationOK, error)
	/*
	   GetCoachingAppointmentAnnotations gets a list of annotations
	   You must have the appropriate permission for the type of annotation you are creating. Permission not required if you are related to the appointment (only the creator or facilitator can view private annotations).
	*/
	GetCoachingAppointmentAnnotations(ctx context.Context, params *GetCoachingAppointmentAnnotationsParams) (*GetCoachingAppointmentAnnotationsOK, error)
	/*
	   GetCoachingAppointmentStatuses gets the list of status changes for a coaching appointment
	   Permission not required if you are an attendee, creator or facilitator of the appointment
	*/
	GetCoachingAppointmentStatuses(ctx context.Context, params *GetCoachingAppointmentStatusesParams) (*GetCoachingAppointmentStatusesOK, error)
	/*
	   GetCoachingAppointments gets appointments for users and optional date range
	*/
	GetCoachingAppointments(ctx context.Context, params *GetCoachingAppointmentsParams) (*GetCoachingAppointmentsOK, error)
	/*
	   GetCoachingAppointmentsMe gets my appointments for a given date range
	*/
	GetCoachingAppointmentsMe(ctx context.Context, params *GetCoachingAppointmentsMeParams) (*GetCoachingAppointmentsMeOK, error)
	/*
	   GetCoachingNotification gets an existing notification
	   Permission not required if you are the owner of the notification.
	*/
	GetCoachingNotification(ctx context.Context, params *GetCoachingNotificationParams) (*GetCoachingNotificationOK, error)
	/*
	   GetCoachingNotifications retrieves the list of your notifications
	*/
	GetCoachingNotifications(ctx context.Context, params *GetCoachingNotificationsParams) (*GetCoachingNotificationsOK, error)
	/*
	   PatchCoachingAppointment updates an existing appointment
	   Permission not required if you are the creator or facilitator of the appointment
	*/
	PatchCoachingAppointment(ctx context.Context, params *PatchCoachingAppointmentParams) (*PatchCoachingAppointmentOK, error)
	/*
	   PatchCoachingAppointmentAnnotation updates an existing annotation
	   You must have the appropriate permission for the type of annotation you are updating. Permission not required if you are the creator or facilitator of the appointment
	*/
	PatchCoachingAppointmentAnnotation(ctx context.Context, params *PatchCoachingAppointmentAnnotationParams) (*PatchCoachingAppointmentAnnotationOK, error)
	/*
	   PatchCoachingAppointmentStatus updates the status of a coaching appointment
	   Permission not required if you are an attendee, creator or facilitator of the appointment
	*/
	PatchCoachingAppointmentStatus(ctx context.Context, params *PatchCoachingAppointmentStatusParams) (*PatchCoachingAppointmentStatusOK, error)
	/*
	   PatchCoachingNotification updates an existing notification
	   Can only update your own notifications.
	*/
	PatchCoachingNotification(ctx context.Context, params *PatchCoachingNotificationParams) (*PatchCoachingNotificationOK, error)
	/*
	   PostCoachingAppointmentAnnotations creates a new annotation
	   You must have the appropriate permission for the type of annotation you are creating. Permission not required if you are related to the appointment (only the creator or facilitator can create private annotations).
	*/
	PostCoachingAppointmentAnnotations(ctx context.Context, params *PostCoachingAppointmentAnnotationsParams) (*PostCoachingAppointmentAnnotationsCreated, error)
	/*
	   PostCoachingAppointments creates a new appointment
	*/
	PostCoachingAppointments(ctx context.Context, params *PostCoachingAppointmentsParams) (*PostCoachingAppointmentsCreated, error)
}

// New creates a new coaching API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, authInfo runtime.ClientAuthInfoWriter) *Client {
	return &Client{
		transport: transport,
		formats:   formats,
		authInfo:  authInfo,
	}
}

/*
Client for coaching API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	authInfo  runtime.ClientAuthInfoWriter
}

/*
DeleteCoachingAppointment deletes an existing appointment

Permission not required if you are the creator of the appointment
*/
func (a *Client) DeleteCoachingAppointment(ctx context.Context, params *DeleteCoachingAppointmentParams) (*DeleteCoachingAppointmentNoContent, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteCoachingAppointment",
		Method:             "DELETE",
		PathPattern:        "/api/v2/coaching/appointments/{appointmentId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteCoachingAppointmentReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteCoachingAppointmentNoContent), nil

}

/*
DeleteCoachingAppointmentAnnotation deletes an existing annotation

You must have the appropriate permission for the type of annotation you are updating. Permission not required if you are the creator or facilitator of the appointment
*/
func (a *Client) DeleteCoachingAppointmentAnnotation(ctx context.Context, params *DeleteCoachingAppointmentAnnotationParams) (*DeleteCoachingAppointmentAnnotationNoContent, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteCoachingAppointmentAnnotation",
		Method:             "DELETE",
		PathPattern:        "/api/v2/coaching/appointments/{appointmentId}/annotations/{annotationId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteCoachingAppointmentAnnotationReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteCoachingAppointmentAnnotationNoContent), nil

}

/*
GetCoachingAppointment retrieves an appointment

Permission not required if you are the attendee, creator or facilitator of the appointment
*/
func (a *Client) GetCoachingAppointment(ctx context.Context, params *GetCoachingAppointmentParams) (*GetCoachingAppointmentOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getCoachingAppointment",
		Method:             "GET",
		PathPattern:        "/api/v2/coaching/appointments/{appointmentId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCoachingAppointmentReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCoachingAppointmentOK), nil

}

/*
GetCoachingAppointmentAnnotation retrieves an annotation

You must have the appropriate permission for the type of annotation you are creating. Permission not required if you are related to the appointment (only the creator or facilitator can view private annotations).
*/
func (a *Client) GetCoachingAppointmentAnnotation(ctx context.Context, params *GetCoachingAppointmentAnnotationParams) (*GetCoachingAppointmentAnnotationOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getCoachingAppointmentAnnotation",
		Method:             "GET",
		PathPattern:        "/api/v2/coaching/appointments/{appointmentId}/annotations/{annotationId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCoachingAppointmentAnnotationReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCoachingAppointmentAnnotationOK), nil

}

/*
GetCoachingAppointmentAnnotations gets a list of annotations

You must have the appropriate permission for the type of annotation you are creating. Permission not required if you are related to the appointment (only the creator or facilitator can view private annotations).
*/
func (a *Client) GetCoachingAppointmentAnnotations(ctx context.Context, params *GetCoachingAppointmentAnnotationsParams) (*GetCoachingAppointmentAnnotationsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getCoachingAppointmentAnnotations",
		Method:             "GET",
		PathPattern:        "/api/v2/coaching/appointments/{appointmentId}/annotations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCoachingAppointmentAnnotationsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCoachingAppointmentAnnotationsOK), nil

}

/*
GetCoachingAppointmentStatuses gets the list of status changes for a coaching appointment

Permission not required if you are an attendee, creator or facilitator of the appointment
*/
func (a *Client) GetCoachingAppointmentStatuses(ctx context.Context, params *GetCoachingAppointmentStatusesParams) (*GetCoachingAppointmentStatusesOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getCoachingAppointmentStatuses",
		Method:             "GET",
		PathPattern:        "/api/v2/coaching/appointments/{appointmentId}/statuses",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCoachingAppointmentStatusesReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCoachingAppointmentStatusesOK), nil

}

/*
GetCoachingAppointments gets appointments for users and optional date range
*/
func (a *Client) GetCoachingAppointments(ctx context.Context, params *GetCoachingAppointmentsParams) (*GetCoachingAppointmentsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getCoachingAppointments",
		Method:             "GET",
		PathPattern:        "/api/v2/coaching/appointments",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCoachingAppointmentsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCoachingAppointmentsOK), nil

}

/*
GetCoachingAppointmentsMe gets my appointments for a given date range
*/
func (a *Client) GetCoachingAppointmentsMe(ctx context.Context, params *GetCoachingAppointmentsMeParams) (*GetCoachingAppointmentsMeOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getCoachingAppointmentsMe",
		Method:             "GET",
		PathPattern:        "/api/v2/coaching/appointments/me",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCoachingAppointmentsMeReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCoachingAppointmentsMeOK), nil

}

/*
GetCoachingNotification gets an existing notification

Permission not required if you are the owner of the notification.
*/
func (a *Client) GetCoachingNotification(ctx context.Context, params *GetCoachingNotificationParams) (*GetCoachingNotificationOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getCoachingNotification",
		Method:             "GET",
		PathPattern:        "/api/v2/coaching/notifications/{notificationId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCoachingNotificationReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCoachingNotificationOK), nil

}

/*
GetCoachingNotifications retrieves the list of your notifications
*/
func (a *Client) GetCoachingNotifications(ctx context.Context, params *GetCoachingNotificationsParams) (*GetCoachingNotificationsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getCoachingNotifications",
		Method:             "GET",
		PathPattern:        "/api/v2/coaching/notifications",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCoachingNotificationsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCoachingNotificationsOK), nil

}

/*
PatchCoachingAppointment updates an existing appointment

Permission not required if you are the creator or facilitator of the appointment
*/
func (a *Client) PatchCoachingAppointment(ctx context.Context, params *PatchCoachingAppointmentParams) (*PatchCoachingAppointmentOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "patchCoachingAppointment",
		Method:             "PATCH",
		PathPattern:        "/api/v2/coaching/appointments/{appointmentId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchCoachingAppointmentReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PatchCoachingAppointmentOK), nil

}

/*
PatchCoachingAppointmentAnnotation updates an existing annotation

You must have the appropriate permission for the type of annotation you are updating. Permission not required if you are the creator or facilitator of the appointment
*/
func (a *Client) PatchCoachingAppointmentAnnotation(ctx context.Context, params *PatchCoachingAppointmentAnnotationParams) (*PatchCoachingAppointmentAnnotationOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "patchCoachingAppointmentAnnotation",
		Method:             "PATCH",
		PathPattern:        "/api/v2/coaching/appointments/{appointmentId}/annotations/{annotationId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchCoachingAppointmentAnnotationReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PatchCoachingAppointmentAnnotationOK), nil

}

/*
PatchCoachingAppointmentStatus updates the status of a coaching appointment

Permission not required if you are an attendee, creator or facilitator of the appointment
*/
func (a *Client) PatchCoachingAppointmentStatus(ctx context.Context, params *PatchCoachingAppointmentStatusParams) (*PatchCoachingAppointmentStatusOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "patchCoachingAppointmentStatus",
		Method:             "PATCH",
		PathPattern:        "/api/v2/coaching/appointments/{appointmentId}/status",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchCoachingAppointmentStatusReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PatchCoachingAppointmentStatusOK), nil

}

/*
PatchCoachingNotification updates an existing notification

Can only update your own notifications.
*/
func (a *Client) PatchCoachingNotification(ctx context.Context, params *PatchCoachingNotificationParams) (*PatchCoachingNotificationOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "patchCoachingNotification",
		Method:             "PATCH",
		PathPattern:        "/api/v2/coaching/notifications/{notificationId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchCoachingNotificationReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PatchCoachingNotificationOK), nil

}

/*
PostCoachingAppointmentAnnotations creates a new annotation

You must have the appropriate permission for the type of annotation you are creating. Permission not required if you are related to the appointment (only the creator or facilitator can create private annotations).
*/
func (a *Client) PostCoachingAppointmentAnnotations(ctx context.Context, params *PostCoachingAppointmentAnnotationsParams) (*PostCoachingAppointmentAnnotationsCreated, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postCoachingAppointmentAnnotations",
		Method:             "POST",
		PathPattern:        "/api/v2/coaching/appointments/{appointmentId}/annotations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostCoachingAppointmentAnnotationsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostCoachingAppointmentAnnotationsCreated), nil

}

/*
PostCoachingAppointments creates a new appointment
*/
func (a *Client) PostCoachingAppointments(ctx context.Context, params *PostCoachingAppointmentsParams) (*PostCoachingAppointmentsCreated, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postCoachingAppointments",
		Method:             "POST",
		PathPattern:        "/api/v2/coaching/appointments",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostCoachingAppointmentsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostCoachingAppointmentsCreated), nil

}

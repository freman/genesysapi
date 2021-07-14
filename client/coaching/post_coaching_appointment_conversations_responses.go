// Code generated by go-swagger; DO NOT EDIT.

package coaching

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostCoachingAppointmentConversationsReader is a Reader for the PostCoachingAppointmentConversations structure.
type PostCoachingAppointmentConversationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostCoachingAppointmentConversationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostCoachingAppointmentConversationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostCoachingAppointmentConversationsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostCoachingAppointmentConversationsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostCoachingAppointmentConversationsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostCoachingAppointmentConversationsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostCoachingAppointmentConversationsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPostCoachingAppointmentConversationsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostCoachingAppointmentConversationsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostCoachingAppointmentConversationsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostCoachingAppointmentConversationsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostCoachingAppointmentConversationsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostCoachingAppointmentConversationsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostCoachingAppointmentConversationsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostCoachingAppointmentConversationsOK creates a PostCoachingAppointmentConversationsOK with default headers values
func NewPostCoachingAppointmentConversationsOK() *PostCoachingAppointmentConversationsOK {
	return &PostCoachingAppointmentConversationsOK{}
}

/*PostCoachingAppointmentConversationsOK handles this case with default header values.

Conversation Added
*/
type PostCoachingAppointmentConversationsOK struct {
	Payload *models.AddConversationResponse
}

func (o *PostCoachingAppointmentConversationsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/coaching/appointments/{appointmentId}/conversations][%d] postCoachingAppointmentConversationsOK  %+v", 200, o.Payload)
}

func (o *PostCoachingAppointmentConversationsOK) GetPayload() *models.AddConversationResponse {
	return o.Payload
}

func (o *PostCoachingAppointmentConversationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AddConversationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCoachingAppointmentConversationsBadRequest creates a PostCoachingAppointmentConversationsBadRequest with default headers values
func NewPostCoachingAppointmentConversationsBadRequest() *PostCoachingAppointmentConversationsBadRequest {
	return &PostCoachingAppointmentConversationsBadRequest{}
}

/*PostCoachingAppointmentConversationsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostCoachingAppointmentConversationsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostCoachingAppointmentConversationsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/coaching/appointments/{appointmentId}/conversations][%d] postCoachingAppointmentConversationsBadRequest  %+v", 400, o.Payload)
}

func (o *PostCoachingAppointmentConversationsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostCoachingAppointmentConversationsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCoachingAppointmentConversationsUnauthorized creates a PostCoachingAppointmentConversationsUnauthorized with default headers values
func NewPostCoachingAppointmentConversationsUnauthorized() *PostCoachingAppointmentConversationsUnauthorized {
	return &PostCoachingAppointmentConversationsUnauthorized{}
}

/*PostCoachingAppointmentConversationsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostCoachingAppointmentConversationsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostCoachingAppointmentConversationsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/coaching/appointments/{appointmentId}/conversations][%d] postCoachingAppointmentConversationsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostCoachingAppointmentConversationsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostCoachingAppointmentConversationsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCoachingAppointmentConversationsForbidden creates a PostCoachingAppointmentConversationsForbidden with default headers values
func NewPostCoachingAppointmentConversationsForbidden() *PostCoachingAppointmentConversationsForbidden {
	return &PostCoachingAppointmentConversationsForbidden{}
}

/*PostCoachingAppointmentConversationsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostCoachingAppointmentConversationsForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostCoachingAppointmentConversationsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/coaching/appointments/{appointmentId}/conversations][%d] postCoachingAppointmentConversationsForbidden  %+v", 403, o.Payload)
}

func (o *PostCoachingAppointmentConversationsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostCoachingAppointmentConversationsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCoachingAppointmentConversationsNotFound creates a PostCoachingAppointmentConversationsNotFound with default headers values
func NewPostCoachingAppointmentConversationsNotFound() *PostCoachingAppointmentConversationsNotFound {
	return &PostCoachingAppointmentConversationsNotFound{}
}

/*PostCoachingAppointmentConversationsNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostCoachingAppointmentConversationsNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostCoachingAppointmentConversationsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/coaching/appointments/{appointmentId}/conversations][%d] postCoachingAppointmentConversationsNotFound  %+v", 404, o.Payload)
}

func (o *PostCoachingAppointmentConversationsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostCoachingAppointmentConversationsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCoachingAppointmentConversationsRequestTimeout creates a PostCoachingAppointmentConversationsRequestTimeout with default headers values
func NewPostCoachingAppointmentConversationsRequestTimeout() *PostCoachingAppointmentConversationsRequestTimeout {
	return &PostCoachingAppointmentConversationsRequestTimeout{}
}

/*PostCoachingAppointmentConversationsRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostCoachingAppointmentConversationsRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostCoachingAppointmentConversationsRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/coaching/appointments/{appointmentId}/conversations][%d] postCoachingAppointmentConversationsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostCoachingAppointmentConversationsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostCoachingAppointmentConversationsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCoachingAppointmentConversationsConflict creates a PostCoachingAppointmentConversationsConflict with default headers values
func NewPostCoachingAppointmentConversationsConflict() *PostCoachingAppointmentConversationsConflict {
	return &PostCoachingAppointmentConversationsConflict{}
}

/*PostCoachingAppointmentConversationsConflict handles this case with default header values.

Conflict
*/
type PostCoachingAppointmentConversationsConflict struct {
	Payload *models.ErrorBody
}

func (o *PostCoachingAppointmentConversationsConflict) Error() string {
	return fmt.Sprintf("[POST /api/v2/coaching/appointments/{appointmentId}/conversations][%d] postCoachingAppointmentConversationsConflict  %+v", 409, o.Payload)
}

func (o *PostCoachingAppointmentConversationsConflict) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostCoachingAppointmentConversationsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCoachingAppointmentConversationsRequestEntityTooLarge creates a PostCoachingAppointmentConversationsRequestEntityTooLarge with default headers values
func NewPostCoachingAppointmentConversationsRequestEntityTooLarge() *PostCoachingAppointmentConversationsRequestEntityTooLarge {
	return &PostCoachingAppointmentConversationsRequestEntityTooLarge{}
}

/*PostCoachingAppointmentConversationsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostCoachingAppointmentConversationsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostCoachingAppointmentConversationsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/coaching/appointments/{appointmentId}/conversations][%d] postCoachingAppointmentConversationsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostCoachingAppointmentConversationsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostCoachingAppointmentConversationsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCoachingAppointmentConversationsUnsupportedMediaType creates a PostCoachingAppointmentConversationsUnsupportedMediaType with default headers values
func NewPostCoachingAppointmentConversationsUnsupportedMediaType() *PostCoachingAppointmentConversationsUnsupportedMediaType {
	return &PostCoachingAppointmentConversationsUnsupportedMediaType{}
}

/*PostCoachingAppointmentConversationsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostCoachingAppointmentConversationsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostCoachingAppointmentConversationsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/coaching/appointments/{appointmentId}/conversations][%d] postCoachingAppointmentConversationsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostCoachingAppointmentConversationsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostCoachingAppointmentConversationsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCoachingAppointmentConversationsTooManyRequests creates a PostCoachingAppointmentConversationsTooManyRequests with default headers values
func NewPostCoachingAppointmentConversationsTooManyRequests() *PostCoachingAppointmentConversationsTooManyRequests {
	return &PostCoachingAppointmentConversationsTooManyRequests{}
}

/*PostCoachingAppointmentConversationsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostCoachingAppointmentConversationsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostCoachingAppointmentConversationsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/coaching/appointments/{appointmentId}/conversations][%d] postCoachingAppointmentConversationsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostCoachingAppointmentConversationsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostCoachingAppointmentConversationsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCoachingAppointmentConversationsInternalServerError creates a PostCoachingAppointmentConversationsInternalServerError with default headers values
func NewPostCoachingAppointmentConversationsInternalServerError() *PostCoachingAppointmentConversationsInternalServerError {
	return &PostCoachingAppointmentConversationsInternalServerError{}
}

/*PostCoachingAppointmentConversationsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostCoachingAppointmentConversationsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostCoachingAppointmentConversationsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/coaching/appointments/{appointmentId}/conversations][%d] postCoachingAppointmentConversationsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostCoachingAppointmentConversationsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostCoachingAppointmentConversationsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCoachingAppointmentConversationsServiceUnavailable creates a PostCoachingAppointmentConversationsServiceUnavailable with default headers values
func NewPostCoachingAppointmentConversationsServiceUnavailable() *PostCoachingAppointmentConversationsServiceUnavailable {
	return &PostCoachingAppointmentConversationsServiceUnavailable{}
}

/*PostCoachingAppointmentConversationsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostCoachingAppointmentConversationsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostCoachingAppointmentConversationsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/coaching/appointments/{appointmentId}/conversations][%d] postCoachingAppointmentConversationsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostCoachingAppointmentConversationsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostCoachingAppointmentConversationsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCoachingAppointmentConversationsGatewayTimeout creates a PostCoachingAppointmentConversationsGatewayTimeout with default headers values
func NewPostCoachingAppointmentConversationsGatewayTimeout() *PostCoachingAppointmentConversationsGatewayTimeout {
	return &PostCoachingAppointmentConversationsGatewayTimeout{}
}

/*PostCoachingAppointmentConversationsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostCoachingAppointmentConversationsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostCoachingAppointmentConversationsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/coaching/appointments/{appointmentId}/conversations][%d] postCoachingAppointmentConversationsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostCoachingAppointmentConversationsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostCoachingAppointmentConversationsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

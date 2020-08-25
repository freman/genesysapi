// Code generated by go-swagger; DO NOT EDIT.

package conversations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PatchConversationsMessageParticipantReader is a Reader for the PatchConversationsMessageParticipant structure.
type PatchConversationsMessageParticipantReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchConversationsMessageParticipantReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewPatchConversationsMessageParticipantAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchConversationsMessageParticipantBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchConversationsMessageParticipantUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchConversationsMessageParticipantForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchConversationsMessageParticipantNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPatchConversationsMessageParticipantRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPatchConversationsMessageParticipantUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPatchConversationsMessageParticipantTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchConversationsMessageParticipantInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPatchConversationsMessageParticipantServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPatchConversationsMessageParticipantGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchConversationsMessageParticipantAccepted creates a PatchConversationsMessageParticipantAccepted with default headers values
func NewPatchConversationsMessageParticipantAccepted() *PatchConversationsMessageParticipantAccepted {
	return &PatchConversationsMessageParticipantAccepted{}
}

/*PatchConversationsMessageParticipantAccepted handles this case with default header values.

Accepted
*/
type PatchConversationsMessageParticipantAccepted struct {
}

func (o *PatchConversationsMessageParticipantAccepted) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messages/{conversationId}/participants/{participantId}][%d] patchConversationsMessageParticipantAccepted ", 202)
}

func (o *PatchConversationsMessageParticipantAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchConversationsMessageParticipantBadRequest creates a PatchConversationsMessageParticipantBadRequest with default headers values
func NewPatchConversationsMessageParticipantBadRequest() *PatchConversationsMessageParticipantBadRequest {
	return &PatchConversationsMessageParticipantBadRequest{}
}

/*PatchConversationsMessageParticipantBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PatchConversationsMessageParticipantBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsMessageParticipantBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messages/{conversationId}/participants/{participantId}][%d] patchConversationsMessageParticipantBadRequest  %+v", 400, o.Payload)
}

func (o *PatchConversationsMessageParticipantBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsMessageParticipantBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsMessageParticipantUnauthorized creates a PatchConversationsMessageParticipantUnauthorized with default headers values
func NewPatchConversationsMessageParticipantUnauthorized() *PatchConversationsMessageParticipantUnauthorized {
	return &PatchConversationsMessageParticipantUnauthorized{}
}

/*PatchConversationsMessageParticipantUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PatchConversationsMessageParticipantUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsMessageParticipantUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messages/{conversationId}/participants/{participantId}][%d] patchConversationsMessageParticipantUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchConversationsMessageParticipantUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsMessageParticipantUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsMessageParticipantForbidden creates a PatchConversationsMessageParticipantForbidden with default headers values
func NewPatchConversationsMessageParticipantForbidden() *PatchConversationsMessageParticipantForbidden {
	return &PatchConversationsMessageParticipantForbidden{}
}

/*PatchConversationsMessageParticipantForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PatchConversationsMessageParticipantForbidden struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsMessageParticipantForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messages/{conversationId}/participants/{participantId}][%d] patchConversationsMessageParticipantForbidden  %+v", 403, o.Payload)
}

func (o *PatchConversationsMessageParticipantForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsMessageParticipantForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsMessageParticipantNotFound creates a PatchConversationsMessageParticipantNotFound with default headers values
func NewPatchConversationsMessageParticipantNotFound() *PatchConversationsMessageParticipantNotFound {
	return &PatchConversationsMessageParticipantNotFound{}
}

/*PatchConversationsMessageParticipantNotFound handles this case with default header values.

The requested resource was not found.
*/
type PatchConversationsMessageParticipantNotFound struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsMessageParticipantNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messages/{conversationId}/participants/{participantId}][%d] patchConversationsMessageParticipantNotFound  %+v", 404, o.Payload)
}

func (o *PatchConversationsMessageParticipantNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsMessageParticipantNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsMessageParticipantRequestEntityTooLarge creates a PatchConversationsMessageParticipantRequestEntityTooLarge with default headers values
func NewPatchConversationsMessageParticipantRequestEntityTooLarge() *PatchConversationsMessageParticipantRequestEntityTooLarge {
	return &PatchConversationsMessageParticipantRequestEntityTooLarge{}
}

/*PatchConversationsMessageParticipantRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PatchConversationsMessageParticipantRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsMessageParticipantRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messages/{conversationId}/participants/{participantId}][%d] patchConversationsMessageParticipantRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchConversationsMessageParticipantRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsMessageParticipantRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsMessageParticipantUnsupportedMediaType creates a PatchConversationsMessageParticipantUnsupportedMediaType with default headers values
func NewPatchConversationsMessageParticipantUnsupportedMediaType() *PatchConversationsMessageParticipantUnsupportedMediaType {
	return &PatchConversationsMessageParticipantUnsupportedMediaType{}
}

/*PatchConversationsMessageParticipantUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PatchConversationsMessageParticipantUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsMessageParticipantUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messages/{conversationId}/participants/{participantId}][%d] patchConversationsMessageParticipantUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchConversationsMessageParticipantUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsMessageParticipantUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsMessageParticipantTooManyRequests creates a PatchConversationsMessageParticipantTooManyRequests with default headers values
func NewPatchConversationsMessageParticipantTooManyRequests() *PatchConversationsMessageParticipantTooManyRequests {
	return &PatchConversationsMessageParticipantTooManyRequests{}
}

/*PatchConversationsMessageParticipantTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PatchConversationsMessageParticipantTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsMessageParticipantTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messages/{conversationId}/participants/{participantId}][%d] patchConversationsMessageParticipantTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchConversationsMessageParticipantTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsMessageParticipantTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsMessageParticipantInternalServerError creates a PatchConversationsMessageParticipantInternalServerError with default headers values
func NewPatchConversationsMessageParticipantInternalServerError() *PatchConversationsMessageParticipantInternalServerError {
	return &PatchConversationsMessageParticipantInternalServerError{}
}

/*PatchConversationsMessageParticipantInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PatchConversationsMessageParticipantInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsMessageParticipantInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messages/{conversationId}/participants/{participantId}][%d] patchConversationsMessageParticipantInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchConversationsMessageParticipantInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsMessageParticipantInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsMessageParticipantServiceUnavailable creates a PatchConversationsMessageParticipantServiceUnavailable with default headers values
func NewPatchConversationsMessageParticipantServiceUnavailable() *PatchConversationsMessageParticipantServiceUnavailable {
	return &PatchConversationsMessageParticipantServiceUnavailable{}
}

/*PatchConversationsMessageParticipantServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PatchConversationsMessageParticipantServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsMessageParticipantServiceUnavailable) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messages/{conversationId}/participants/{participantId}][%d] patchConversationsMessageParticipantServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchConversationsMessageParticipantServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsMessageParticipantServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsMessageParticipantGatewayTimeout creates a PatchConversationsMessageParticipantGatewayTimeout with default headers values
func NewPatchConversationsMessageParticipantGatewayTimeout() *PatchConversationsMessageParticipantGatewayTimeout {
	return &PatchConversationsMessageParticipantGatewayTimeout{}
}

/*PatchConversationsMessageParticipantGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PatchConversationsMessageParticipantGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsMessageParticipantGatewayTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messages/{conversationId}/participants/{participantId}][%d] patchConversationsMessageParticipantGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchConversationsMessageParticipantGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsMessageParticipantGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
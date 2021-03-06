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

// PatchConversationsChatParticipantCommunicationReader is a Reader for the PatchConversationsChatParticipantCommunication structure.
type PatchConversationsChatParticipantCommunicationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchConversationsChatParticipantCommunicationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchConversationsChatParticipantCommunicationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchConversationsChatParticipantCommunicationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchConversationsChatParticipantCommunicationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchConversationsChatParticipantCommunicationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchConversationsChatParticipantCommunicationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPatchConversationsChatParticipantCommunicationRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPatchConversationsChatParticipantCommunicationUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPatchConversationsChatParticipantCommunicationTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchConversationsChatParticipantCommunicationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPatchConversationsChatParticipantCommunicationServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPatchConversationsChatParticipantCommunicationGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchConversationsChatParticipantCommunicationOK creates a PatchConversationsChatParticipantCommunicationOK with default headers values
func NewPatchConversationsChatParticipantCommunicationOK() *PatchConversationsChatParticipantCommunicationOK {
	return &PatchConversationsChatParticipantCommunicationOK{}
}

/*PatchConversationsChatParticipantCommunicationOK handles this case with default header values.

successful operation
*/
type PatchConversationsChatParticipantCommunicationOK struct {
	Payload models.Empty
}

func (o *PatchConversationsChatParticipantCommunicationOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/chats/{conversationId}/participants/{participantId}/communications/{communicationId}][%d] patchConversationsChatParticipantCommunicationOK  %+v", 200, o.Payload)
}

func (o *PatchConversationsChatParticipantCommunicationOK) GetPayload() models.Empty {
	return o.Payload
}

func (o *PatchConversationsChatParticipantCommunicationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsChatParticipantCommunicationBadRequest creates a PatchConversationsChatParticipantCommunicationBadRequest with default headers values
func NewPatchConversationsChatParticipantCommunicationBadRequest() *PatchConversationsChatParticipantCommunicationBadRequest {
	return &PatchConversationsChatParticipantCommunicationBadRequest{}
}

/*PatchConversationsChatParticipantCommunicationBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PatchConversationsChatParticipantCommunicationBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsChatParticipantCommunicationBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/chats/{conversationId}/participants/{participantId}/communications/{communicationId}][%d] patchConversationsChatParticipantCommunicationBadRequest  %+v", 400, o.Payload)
}

func (o *PatchConversationsChatParticipantCommunicationBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsChatParticipantCommunicationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsChatParticipantCommunicationUnauthorized creates a PatchConversationsChatParticipantCommunicationUnauthorized with default headers values
func NewPatchConversationsChatParticipantCommunicationUnauthorized() *PatchConversationsChatParticipantCommunicationUnauthorized {
	return &PatchConversationsChatParticipantCommunicationUnauthorized{}
}

/*PatchConversationsChatParticipantCommunicationUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PatchConversationsChatParticipantCommunicationUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsChatParticipantCommunicationUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/chats/{conversationId}/participants/{participantId}/communications/{communicationId}][%d] patchConversationsChatParticipantCommunicationUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchConversationsChatParticipantCommunicationUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsChatParticipantCommunicationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsChatParticipantCommunicationForbidden creates a PatchConversationsChatParticipantCommunicationForbidden with default headers values
func NewPatchConversationsChatParticipantCommunicationForbidden() *PatchConversationsChatParticipantCommunicationForbidden {
	return &PatchConversationsChatParticipantCommunicationForbidden{}
}

/*PatchConversationsChatParticipantCommunicationForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PatchConversationsChatParticipantCommunicationForbidden struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsChatParticipantCommunicationForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/chats/{conversationId}/participants/{participantId}/communications/{communicationId}][%d] patchConversationsChatParticipantCommunicationForbidden  %+v", 403, o.Payload)
}

func (o *PatchConversationsChatParticipantCommunicationForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsChatParticipantCommunicationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsChatParticipantCommunicationNotFound creates a PatchConversationsChatParticipantCommunicationNotFound with default headers values
func NewPatchConversationsChatParticipantCommunicationNotFound() *PatchConversationsChatParticipantCommunicationNotFound {
	return &PatchConversationsChatParticipantCommunicationNotFound{}
}

/*PatchConversationsChatParticipantCommunicationNotFound handles this case with default header values.

The requested resource was not found.
*/
type PatchConversationsChatParticipantCommunicationNotFound struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsChatParticipantCommunicationNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/chats/{conversationId}/participants/{participantId}/communications/{communicationId}][%d] patchConversationsChatParticipantCommunicationNotFound  %+v", 404, o.Payload)
}

func (o *PatchConversationsChatParticipantCommunicationNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsChatParticipantCommunicationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsChatParticipantCommunicationRequestEntityTooLarge creates a PatchConversationsChatParticipantCommunicationRequestEntityTooLarge with default headers values
func NewPatchConversationsChatParticipantCommunicationRequestEntityTooLarge() *PatchConversationsChatParticipantCommunicationRequestEntityTooLarge {
	return &PatchConversationsChatParticipantCommunicationRequestEntityTooLarge{}
}

/*PatchConversationsChatParticipantCommunicationRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PatchConversationsChatParticipantCommunicationRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsChatParticipantCommunicationRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/chats/{conversationId}/participants/{participantId}/communications/{communicationId}][%d] patchConversationsChatParticipantCommunicationRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchConversationsChatParticipantCommunicationRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsChatParticipantCommunicationRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsChatParticipantCommunicationUnsupportedMediaType creates a PatchConversationsChatParticipantCommunicationUnsupportedMediaType with default headers values
func NewPatchConversationsChatParticipantCommunicationUnsupportedMediaType() *PatchConversationsChatParticipantCommunicationUnsupportedMediaType {
	return &PatchConversationsChatParticipantCommunicationUnsupportedMediaType{}
}

/*PatchConversationsChatParticipantCommunicationUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PatchConversationsChatParticipantCommunicationUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsChatParticipantCommunicationUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/chats/{conversationId}/participants/{participantId}/communications/{communicationId}][%d] patchConversationsChatParticipantCommunicationUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchConversationsChatParticipantCommunicationUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsChatParticipantCommunicationUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsChatParticipantCommunicationTooManyRequests creates a PatchConversationsChatParticipantCommunicationTooManyRequests with default headers values
func NewPatchConversationsChatParticipantCommunicationTooManyRequests() *PatchConversationsChatParticipantCommunicationTooManyRequests {
	return &PatchConversationsChatParticipantCommunicationTooManyRequests{}
}

/*PatchConversationsChatParticipantCommunicationTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PatchConversationsChatParticipantCommunicationTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsChatParticipantCommunicationTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/chats/{conversationId}/participants/{participantId}/communications/{communicationId}][%d] patchConversationsChatParticipantCommunicationTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchConversationsChatParticipantCommunicationTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsChatParticipantCommunicationTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsChatParticipantCommunicationInternalServerError creates a PatchConversationsChatParticipantCommunicationInternalServerError with default headers values
func NewPatchConversationsChatParticipantCommunicationInternalServerError() *PatchConversationsChatParticipantCommunicationInternalServerError {
	return &PatchConversationsChatParticipantCommunicationInternalServerError{}
}

/*PatchConversationsChatParticipantCommunicationInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PatchConversationsChatParticipantCommunicationInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsChatParticipantCommunicationInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/chats/{conversationId}/participants/{participantId}/communications/{communicationId}][%d] patchConversationsChatParticipantCommunicationInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchConversationsChatParticipantCommunicationInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsChatParticipantCommunicationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsChatParticipantCommunicationServiceUnavailable creates a PatchConversationsChatParticipantCommunicationServiceUnavailable with default headers values
func NewPatchConversationsChatParticipantCommunicationServiceUnavailable() *PatchConversationsChatParticipantCommunicationServiceUnavailable {
	return &PatchConversationsChatParticipantCommunicationServiceUnavailable{}
}

/*PatchConversationsChatParticipantCommunicationServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PatchConversationsChatParticipantCommunicationServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsChatParticipantCommunicationServiceUnavailable) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/chats/{conversationId}/participants/{participantId}/communications/{communicationId}][%d] patchConversationsChatParticipantCommunicationServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchConversationsChatParticipantCommunicationServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsChatParticipantCommunicationServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsChatParticipantCommunicationGatewayTimeout creates a PatchConversationsChatParticipantCommunicationGatewayTimeout with default headers values
func NewPatchConversationsChatParticipantCommunicationGatewayTimeout() *PatchConversationsChatParticipantCommunicationGatewayTimeout {
	return &PatchConversationsChatParticipantCommunicationGatewayTimeout{}
}

/*PatchConversationsChatParticipantCommunicationGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PatchConversationsChatParticipantCommunicationGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsChatParticipantCommunicationGatewayTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/chats/{conversationId}/participants/{participantId}/communications/{communicationId}][%d] patchConversationsChatParticipantCommunicationGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchConversationsChatParticipantCommunicationGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsChatParticipantCommunicationGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

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

// PatchConversationsEmailParticipantReader is a Reader for the PatchConversationsEmailParticipant structure.
type PatchConversationsEmailParticipantReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchConversationsEmailParticipantReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewPatchConversationsEmailParticipantAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchConversationsEmailParticipantBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchConversationsEmailParticipantUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchConversationsEmailParticipantForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchConversationsEmailParticipantNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPatchConversationsEmailParticipantRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPatchConversationsEmailParticipantUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPatchConversationsEmailParticipantTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchConversationsEmailParticipantInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPatchConversationsEmailParticipantServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPatchConversationsEmailParticipantGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchConversationsEmailParticipantAccepted creates a PatchConversationsEmailParticipantAccepted with default headers values
func NewPatchConversationsEmailParticipantAccepted() *PatchConversationsEmailParticipantAccepted {
	return &PatchConversationsEmailParticipantAccepted{}
}

/*PatchConversationsEmailParticipantAccepted handles this case with default header values.

Accepted
*/
type PatchConversationsEmailParticipantAccepted struct {
}

func (o *PatchConversationsEmailParticipantAccepted) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/emails/{conversationId}/participants/{participantId}][%d] patchConversationsEmailParticipantAccepted ", 202)
}

func (o *PatchConversationsEmailParticipantAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchConversationsEmailParticipantBadRequest creates a PatchConversationsEmailParticipantBadRequest with default headers values
func NewPatchConversationsEmailParticipantBadRequest() *PatchConversationsEmailParticipantBadRequest {
	return &PatchConversationsEmailParticipantBadRequest{}
}

/*PatchConversationsEmailParticipantBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PatchConversationsEmailParticipantBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsEmailParticipantBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/emails/{conversationId}/participants/{participantId}][%d] patchConversationsEmailParticipantBadRequest  %+v", 400, o.Payload)
}

func (o *PatchConversationsEmailParticipantBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsEmailParticipantBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsEmailParticipantUnauthorized creates a PatchConversationsEmailParticipantUnauthorized with default headers values
func NewPatchConversationsEmailParticipantUnauthorized() *PatchConversationsEmailParticipantUnauthorized {
	return &PatchConversationsEmailParticipantUnauthorized{}
}

/*PatchConversationsEmailParticipantUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PatchConversationsEmailParticipantUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsEmailParticipantUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/emails/{conversationId}/participants/{participantId}][%d] patchConversationsEmailParticipantUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchConversationsEmailParticipantUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsEmailParticipantUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsEmailParticipantForbidden creates a PatchConversationsEmailParticipantForbidden with default headers values
func NewPatchConversationsEmailParticipantForbidden() *PatchConversationsEmailParticipantForbidden {
	return &PatchConversationsEmailParticipantForbidden{}
}

/*PatchConversationsEmailParticipantForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PatchConversationsEmailParticipantForbidden struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsEmailParticipantForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/emails/{conversationId}/participants/{participantId}][%d] patchConversationsEmailParticipantForbidden  %+v", 403, o.Payload)
}

func (o *PatchConversationsEmailParticipantForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsEmailParticipantForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsEmailParticipantNotFound creates a PatchConversationsEmailParticipantNotFound with default headers values
func NewPatchConversationsEmailParticipantNotFound() *PatchConversationsEmailParticipantNotFound {
	return &PatchConversationsEmailParticipantNotFound{}
}

/*PatchConversationsEmailParticipantNotFound handles this case with default header values.

The requested resource was not found.
*/
type PatchConversationsEmailParticipantNotFound struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsEmailParticipantNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/emails/{conversationId}/participants/{participantId}][%d] patchConversationsEmailParticipantNotFound  %+v", 404, o.Payload)
}

func (o *PatchConversationsEmailParticipantNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsEmailParticipantNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsEmailParticipantRequestEntityTooLarge creates a PatchConversationsEmailParticipantRequestEntityTooLarge with default headers values
func NewPatchConversationsEmailParticipantRequestEntityTooLarge() *PatchConversationsEmailParticipantRequestEntityTooLarge {
	return &PatchConversationsEmailParticipantRequestEntityTooLarge{}
}

/*PatchConversationsEmailParticipantRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PatchConversationsEmailParticipantRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsEmailParticipantRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/emails/{conversationId}/participants/{participantId}][%d] patchConversationsEmailParticipantRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchConversationsEmailParticipantRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsEmailParticipantRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsEmailParticipantUnsupportedMediaType creates a PatchConversationsEmailParticipantUnsupportedMediaType with default headers values
func NewPatchConversationsEmailParticipantUnsupportedMediaType() *PatchConversationsEmailParticipantUnsupportedMediaType {
	return &PatchConversationsEmailParticipantUnsupportedMediaType{}
}

/*PatchConversationsEmailParticipantUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PatchConversationsEmailParticipantUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsEmailParticipantUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/emails/{conversationId}/participants/{participantId}][%d] patchConversationsEmailParticipantUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchConversationsEmailParticipantUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsEmailParticipantUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsEmailParticipantTooManyRequests creates a PatchConversationsEmailParticipantTooManyRequests with default headers values
func NewPatchConversationsEmailParticipantTooManyRequests() *PatchConversationsEmailParticipantTooManyRequests {
	return &PatchConversationsEmailParticipantTooManyRequests{}
}

/*PatchConversationsEmailParticipantTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PatchConversationsEmailParticipantTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsEmailParticipantTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/emails/{conversationId}/participants/{participantId}][%d] patchConversationsEmailParticipantTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchConversationsEmailParticipantTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsEmailParticipantTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsEmailParticipantInternalServerError creates a PatchConversationsEmailParticipantInternalServerError with default headers values
func NewPatchConversationsEmailParticipantInternalServerError() *PatchConversationsEmailParticipantInternalServerError {
	return &PatchConversationsEmailParticipantInternalServerError{}
}

/*PatchConversationsEmailParticipantInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PatchConversationsEmailParticipantInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsEmailParticipantInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/emails/{conversationId}/participants/{participantId}][%d] patchConversationsEmailParticipantInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchConversationsEmailParticipantInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsEmailParticipantInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsEmailParticipantServiceUnavailable creates a PatchConversationsEmailParticipantServiceUnavailable with default headers values
func NewPatchConversationsEmailParticipantServiceUnavailable() *PatchConversationsEmailParticipantServiceUnavailable {
	return &PatchConversationsEmailParticipantServiceUnavailable{}
}

/*PatchConversationsEmailParticipantServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PatchConversationsEmailParticipantServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsEmailParticipantServiceUnavailable) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/emails/{conversationId}/participants/{participantId}][%d] patchConversationsEmailParticipantServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchConversationsEmailParticipantServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsEmailParticipantServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsEmailParticipantGatewayTimeout creates a PatchConversationsEmailParticipantGatewayTimeout with default headers values
func NewPatchConversationsEmailParticipantGatewayTimeout() *PatchConversationsEmailParticipantGatewayTimeout {
	return &PatchConversationsEmailParticipantGatewayTimeout{}
}

/*PatchConversationsEmailParticipantGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PatchConversationsEmailParticipantGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsEmailParticipantGatewayTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/emails/{conversationId}/participants/{participantId}][%d] patchConversationsEmailParticipantGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchConversationsEmailParticipantGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsEmailParticipantGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

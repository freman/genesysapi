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

// PatchConversationsCallParticipantReader is a Reader for the PatchConversationsCallParticipant structure.
type PatchConversationsCallParticipantReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchConversationsCallParticipantReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewPatchConversationsCallParticipantAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchConversationsCallParticipantBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchConversationsCallParticipantUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchConversationsCallParticipantForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchConversationsCallParticipantNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPatchConversationsCallParticipantRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPatchConversationsCallParticipantRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPatchConversationsCallParticipantUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPatchConversationsCallParticipantTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchConversationsCallParticipantInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPatchConversationsCallParticipantServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPatchConversationsCallParticipantGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchConversationsCallParticipantAccepted creates a PatchConversationsCallParticipantAccepted with default headers values
func NewPatchConversationsCallParticipantAccepted() *PatchConversationsCallParticipantAccepted {
	return &PatchConversationsCallParticipantAccepted{}
}

/*PatchConversationsCallParticipantAccepted handles this case with default header values.

Accepted
*/
type PatchConversationsCallParticipantAccepted struct {
}

func (o *PatchConversationsCallParticipantAccepted) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/calls/{conversationId}/participants/{participantId}][%d] patchConversationsCallParticipantAccepted ", 202)
}

func (o *PatchConversationsCallParticipantAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchConversationsCallParticipantBadRequest creates a PatchConversationsCallParticipantBadRequest with default headers values
func NewPatchConversationsCallParticipantBadRequest() *PatchConversationsCallParticipantBadRequest {
	return &PatchConversationsCallParticipantBadRequest{}
}

/*PatchConversationsCallParticipantBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PatchConversationsCallParticipantBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsCallParticipantBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/calls/{conversationId}/participants/{participantId}][%d] patchConversationsCallParticipantBadRequest  %+v", 400, o.Payload)
}

func (o *PatchConversationsCallParticipantBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsCallParticipantBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsCallParticipantUnauthorized creates a PatchConversationsCallParticipantUnauthorized with default headers values
func NewPatchConversationsCallParticipantUnauthorized() *PatchConversationsCallParticipantUnauthorized {
	return &PatchConversationsCallParticipantUnauthorized{}
}

/*PatchConversationsCallParticipantUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PatchConversationsCallParticipantUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsCallParticipantUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/calls/{conversationId}/participants/{participantId}][%d] patchConversationsCallParticipantUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchConversationsCallParticipantUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsCallParticipantUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsCallParticipantForbidden creates a PatchConversationsCallParticipantForbidden with default headers values
func NewPatchConversationsCallParticipantForbidden() *PatchConversationsCallParticipantForbidden {
	return &PatchConversationsCallParticipantForbidden{}
}

/*PatchConversationsCallParticipantForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PatchConversationsCallParticipantForbidden struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsCallParticipantForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/calls/{conversationId}/participants/{participantId}][%d] patchConversationsCallParticipantForbidden  %+v", 403, o.Payload)
}

func (o *PatchConversationsCallParticipantForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsCallParticipantForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsCallParticipantNotFound creates a PatchConversationsCallParticipantNotFound with default headers values
func NewPatchConversationsCallParticipantNotFound() *PatchConversationsCallParticipantNotFound {
	return &PatchConversationsCallParticipantNotFound{}
}

/*PatchConversationsCallParticipantNotFound handles this case with default header values.

The requested resource was not found.
*/
type PatchConversationsCallParticipantNotFound struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsCallParticipantNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/calls/{conversationId}/participants/{participantId}][%d] patchConversationsCallParticipantNotFound  %+v", 404, o.Payload)
}

func (o *PatchConversationsCallParticipantNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsCallParticipantNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsCallParticipantRequestTimeout creates a PatchConversationsCallParticipantRequestTimeout with default headers values
func NewPatchConversationsCallParticipantRequestTimeout() *PatchConversationsCallParticipantRequestTimeout {
	return &PatchConversationsCallParticipantRequestTimeout{}
}

/*PatchConversationsCallParticipantRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PatchConversationsCallParticipantRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsCallParticipantRequestTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/calls/{conversationId}/participants/{participantId}][%d] patchConversationsCallParticipantRequestTimeout  %+v", 408, o.Payload)
}

func (o *PatchConversationsCallParticipantRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsCallParticipantRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsCallParticipantRequestEntityTooLarge creates a PatchConversationsCallParticipantRequestEntityTooLarge with default headers values
func NewPatchConversationsCallParticipantRequestEntityTooLarge() *PatchConversationsCallParticipantRequestEntityTooLarge {
	return &PatchConversationsCallParticipantRequestEntityTooLarge{}
}

/*PatchConversationsCallParticipantRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PatchConversationsCallParticipantRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsCallParticipantRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/calls/{conversationId}/participants/{participantId}][%d] patchConversationsCallParticipantRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchConversationsCallParticipantRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsCallParticipantRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsCallParticipantUnsupportedMediaType creates a PatchConversationsCallParticipantUnsupportedMediaType with default headers values
func NewPatchConversationsCallParticipantUnsupportedMediaType() *PatchConversationsCallParticipantUnsupportedMediaType {
	return &PatchConversationsCallParticipantUnsupportedMediaType{}
}

/*PatchConversationsCallParticipantUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PatchConversationsCallParticipantUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsCallParticipantUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/calls/{conversationId}/participants/{participantId}][%d] patchConversationsCallParticipantUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchConversationsCallParticipantUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsCallParticipantUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsCallParticipantTooManyRequests creates a PatchConversationsCallParticipantTooManyRequests with default headers values
func NewPatchConversationsCallParticipantTooManyRequests() *PatchConversationsCallParticipantTooManyRequests {
	return &PatchConversationsCallParticipantTooManyRequests{}
}

/*PatchConversationsCallParticipantTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PatchConversationsCallParticipantTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsCallParticipantTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/calls/{conversationId}/participants/{participantId}][%d] patchConversationsCallParticipantTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchConversationsCallParticipantTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsCallParticipantTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsCallParticipantInternalServerError creates a PatchConversationsCallParticipantInternalServerError with default headers values
func NewPatchConversationsCallParticipantInternalServerError() *PatchConversationsCallParticipantInternalServerError {
	return &PatchConversationsCallParticipantInternalServerError{}
}

/*PatchConversationsCallParticipantInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PatchConversationsCallParticipantInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsCallParticipantInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/calls/{conversationId}/participants/{participantId}][%d] patchConversationsCallParticipantInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchConversationsCallParticipantInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsCallParticipantInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsCallParticipantServiceUnavailable creates a PatchConversationsCallParticipantServiceUnavailable with default headers values
func NewPatchConversationsCallParticipantServiceUnavailable() *PatchConversationsCallParticipantServiceUnavailable {
	return &PatchConversationsCallParticipantServiceUnavailable{}
}

/*PatchConversationsCallParticipantServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PatchConversationsCallParticipantServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsCallParticipantServiceUnavailable) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/calls/{conversationId}/participants/{participantId}][%d] patchConversationsCallParticipantServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchConversationsCallParticipantServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsCallParticipantServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsCallParticipantGatewayTimeout creates a PatchConversationsCallParticipantGatewayTimeout with default headers values
func NewPatchConversationsCallParticipantGatewayTimeout() *PatchConversationsCallParticipantGatewayTimeout {
	return &PatchConversationsCallParticipantGatewayTimeout{}
}

/*PatchConversationsCallParticipantGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PatchConversationsCallParticipantGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsCallParticipantGatewayTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/calls/{conversationId}/participants/{participantId}][%d] patchConversationsCallParticipantGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchConversationsCallParticipantGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsCallParticipantGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

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

// PatchConversationsCallParticipantAttributesReader is a Reader for the PatchConversationsCallParticipantAttributes structure.
type PatchConversationsCallParticipantAttributesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchConversationsCallParticipantAttributesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewPatchConversationsCallParticipantAttributesAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchConversationsCallParticipantAttributesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchConversationsCallParticipantAttributesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchConversationsCallParticipantAttributesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchConversationsCallParticipantAttributesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPatchConversationsCallParticipantAttributesRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPatchConversationsCallParticipantAttributesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPatchConversationsCallParticipantAttributesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPatchConversationsCallParticipantAttributesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchConversationsCallParticipantAttributesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPatchConversationsCallParticipantAttributesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPatchConversationsCallParticipantAttributesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchConversationsCallParticipantAttributesAccepted creates a PatchConversationsCallParticipantAttributesAccepted with default headers values
func NewPatchConversationsCallParticipantAttributesAccepted() *PatchConversationsCallParticipantAttributesAccepted {
	return &PatchConversationsCallParticipantAttributesAccepted{}
}

/*PatchConversationsCallParticipantAttributesAccepted handles this case with default header values.

Accepted
*/
type PatchConversationsCallParticipantAttributesAccepted struct {
}

func (o *PatchConversationsCallParticipantAttributesAccepted) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/calls/{conversationId}/participants/{participantId}/attributes][%d] patchConversationsCallParticipantAttributesAccepted ", 202)
}

func (o *PatchConversationsCallParticipantAttributesAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchConversationsCallParticipantAttributesBadRequest creates a PatchConversationsCallParticipantAttributesBadRequest with default headers values
func NewPatchConversationsCallParticipantAttributesBadRequest() *PatchConversationsCallParticipantAttributesBadRequest {
	return &PatchConversationsCallParticipantAttributesBadRequest{}
}

/*PatchConversationsCallParticipantAttributesBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PatchConversationsCallParticipantAttributesBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsCallParticipantAttributesBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/calls/{conversationId}/participants/{participantId}/attributes][%d] patchConversationsCallParticipantAttributesBadRequest  %+v", 400, o.Payload)
}

func (o *PatchConversationsCallParticipantAttributesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsCallParticipantAttributesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsCallParticipantAttributesUnauthorized creates a PatchConversationsCallParticipantAttributesUnauthorized with default headers values
func NewPatchConversationsCallParticipantAttributesUnauthorized() *PatchConversationsCallParticipantAttributesUnauthorized {
	return &PatchConversationsCallParticipantAttributesUnauthorized{}
}

/*PatchConversationsCallParticipantAttributesUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PatchConversationsCallParticipantAttributesUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsCallParticipantAttributesUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/calls/{conversationId}/participants/{participantId}/attributes][%d] patchConversationsCallParticipantAttributesUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchConversationsCallParticipantAttributesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsCallParticipantAttributesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsCallParticipantAttributesForbidden creates a PatchConversationsCallParticipantAttributesForbidden with default headers values
func NewPatchConversationsCallParticipantAttributesForbidden() *PatchConversationsCallParticipantAttributesForbidden {
	return &PatchConversationsCallParticipantAttributesForbidden{}
}

/*PatchConversationsCallParticipantAttributesForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PatchConversationsCallParticipantAttributesForbidden struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsCallParticipantAttributesForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/calls/{conversationId}/participants/{participantId}/attributes][%d] patchConversationsCallParticipantAttributesForbidden  %+v", 403, o.Payload)
}

func (o *PatchConversationsCallParticipantAttributesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsCallParticipantAttributesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsCallParticipantAttributesNotFound creates a PatchConversationsCallParticipantAttributesNotFound with default headers values
func NewPatchConversationsCallParticipantAttributesNotFound() *PatchConversationsCallParticipantAttributesNotFound {
	return &PatchConversationsCallParticipantAttributesNotFound{}
}

/*PatchConversationsCallParticipantAttributesNotFound handles this case with default header values.

The requested resource was not found.
*/
type PatchConversationsCallParticipantAttributesNotFound struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsCallParticipantAttributesNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/calls/{conversationId}/participants/{participantId}/attributes][%d] patchConversationsCallParticipantAttributesNotFound  %+v", 404, o.Payload)
}

func (o *PatchConversationsCallParticipantAttributesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsCallParticipantAttributesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsCallParticipantAttributesRequestTimeout creates a PatchConversationsCallParticipantAttributesRequestTimeout with default headers values
func NewPatchConversationsCallParticipantAttributesRequestTimeout() *PatchConversationsCallParticipantAttributesRequestTimeout {
	return &PatchConversationsCallParticipantAttributesRequestTimeout{}
}

/*PatchConversationsCallParticipantAttributesRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PatchConversationsCallParticipantAttributesRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsCallParticipantAttributesRequestTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/calls/{conversationId}/participants/{participantId}/attributes][%d] patchConversationsCallParticipantAttributesRequestTimeout  %+v", 408, o.Payload)
}

func (o *PatchConversationsCallParticipantAttributesRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsCallParticipantAttributesRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsCallParticipantAttributesRequestEntityTooLarge creates a PatchConversationsCallParticipantAttributesRequestEntityTooLarge with default headers values
func NewPatchConversationsCallParticipantAttributesRequestEntityTooLarge() *PatchConversationsCallParticipantAttributesRequestEntityTooLarge {
	return &PatchConversationsCallParticipantAttributesRequestEntityTooLarge{}
}

/*PatchConversationsCallParticipantAttributesRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PatchConversationsCallParticipantAttributesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsCallParticipantAttributesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/calls/{conversationId}/participants/{participantId}/attributes][%d] patchConversationsCallParticipantAttributesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchConversationsCallParticipantAttributesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsCallParticipantAttributesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsCallParticipantAttributesUnsupportedMediaType creates a PatchConversationsCallParticipantAttributesUnsupportedMediaType with default headers values
func NewPatchConversationsCallParticipantAttributesUnsupportedMediaType() *PatchConversationsCallParticipantAttributesUnsupportedMediaType {
	return &PatchConversationsCallParticipantAttributesUnsupportedMediaType{}
}

/*PatchConversationsCallParticipantAttributesUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PatchConversationsCallParticipantAttributesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsCallParticipantAttributesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/calls/{conversationId}/participants/{participantId}/attributes][%d] patchConversationsCallParticipantAttributesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchConversationsCallParticipantAttributesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsCallParticipantAttributesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsCallParticipantAttributesTooManyRequests creates a PatchConversationsCallParticipantAttributesTooManyRequests with default headers values
func NewPatchConversationsCallParticipantAttributesTooManyRequests() *PatchConversationsCallParticipantAttributesTooManyRequests {
	return &PatchConversationsCallParticipantAttributesTooManyRequests{}
}

/*PatchConversationsCallParticipantAttributesTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PatchConversationsCallParticipantAttributesTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsCallParticipantAttributesTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/calls/{conversationId}/participants/{participantId}/attributes][%d] patchConversationsCallParticipantAttributesTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchConversationsCallParticipantAttributesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsCallParticipantAttributesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsCallParticipantAttributesInternalServerError creates a PatchConversationsCallParticipantAttributesInternalServerError with default headers values
func NewPatchConversationsCallParticipantAttributesInternalServerError() *PatchConversationsCallParticipantAttributesInternalServerError {
	return &PatchConversationsCallParticipantAttributesInternalServerError{}
}

/*PatchConversationsCallParticipantAttributesInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PatchConversationsCallParticipantAttributesInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsCallParticipantAttributesInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/calls/{conversationId}/participants/{participantId}/attributes][%d] patchConversationsCallParticipantAttributesInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchConversationsCallParticipantAttributesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsCallParticipantAttributesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsCallParticipantAttributesServiceUnavailable creates a PatchConversationsCallParticipantAttributesServiceUnavailable with default headers values
func NewPatchConversationsCallParticipantAttributesServiceUnavailable() *PatchConversationsCallParticipantAttributesServiceUnavailable {
	return &PatchConversationsCallParticipantAttributesServiceUnavailable{}
}

/*PatchConversationsCallParticipantAttributesServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PatchConversationsCallParticipantAttributesServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsCallParticipantAttributesServiceUnavailable) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/calls/{conversationId}/participants/{participantId}/attributes][%d] patchConversationsCallParticipantAttributesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchConversationsCallParticipantAttributesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsCallParticipantAttributesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsCallParticipantAttributesGatewayTimeout creates a PatchConversationsCallParticipantAttributesGatewayTimeout with default headers values
func NewPatchConversationsCallParticipantAttributesGatewayTimeout() *PatchConversationsCallParticipantAttributesGatewayTimeout {
	return &PatchConversationsCallParticipantAttributesGatewayTimeout{}
}

/*PatchConversationsCallParticipantAttributesGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PatchConversationsCallParticipantAttributesGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsCallParticipantAttributesGatewayTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/calls/{conversationId}/participants/{participantId}/attributes][%d] patchConversationsCallParticipantAttributesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchConversationsCallParticipantAttributesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsCallParticipantAttributesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

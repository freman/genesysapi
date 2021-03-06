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

// PatchConversationsCobrowsesessionParticipantAttributesReader is a Reader for the PatchConversationsCobrowsesessionParticipantAttributes structure.
type PatchConversationsCobrowsesessionParticipantAttributesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchConversationsCobrowsesessionParticipantAttributesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewPatchConversationsCobrowsesessionParticipantAttributesAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchConversationsCobrowsesessionParticipantAttributesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchConversationsCobrowsesessionParticipantAttributesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchConversationsCobrowsesessionParticipantAttributesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchConversationsCobrowsesessionParticipantAttributesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPatchConversationsCobrowsesessionParticipantAttributesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPatchConversationsCobrowsesessionParticipantAttributesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPatchConversationsCobrowsesessionParticipantAttributesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchConversationsCobrowsesessionParticipantAttributesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPatchConversationsCobrowsesessionParticipantAttributesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPatchConversationsCobrowsesessionParticipantAttributesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchConversationsCobrowsesessionParticipantAttributesAccepted creates a PatchConversationsCobrowsesessionParticipantAttributesAccepted with default headers values
func NewPatchConversationsCobrowsesessionParticipantAttributesAccepted() *PatchConversationsCobrowsesessionParticipantAttributesAccepted {
	return &PatchConversationsCobrowsesessionParticipantAttributesAccepted{}
}

/*PatchConversationsCobrowsesessionParticipantAttributesAccepted handles this case with default header values.

Accepted
*/
type PatchConversationsCobrowsesessionParticipantAttributesAccepted struct {
}

func (o *PatchConversationsCobrowsesessionParticipantAttributesAccepted) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/cobrowsesessions/{conversationId}/participants/{participantId}/attributes][%d] patchConversationsCobrowsesessionParticipantAttributesAccepted ", 202)
}

func (o *PatchConversationsCobrowsesessionParticipantAttributesAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchConversationsCobrowsesessionParticipantAttributesBadRequest creates a PatchConversationsCobrowsesessionParticipantAttributesBadRequest with default headers values
func NewPatchConversationsCobrowsesessionParticipantAttributesBadRequest() *PatchConversationsCobrowsesessionParticipantAttributesBadRequest {
	return &PatchConversationsCobrowsesessionParticipantAttributesBadRequest{}
}

/*PatchConversationsCobrowsesessionParticipantAttributesBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PatchConversationsCobrowsesessionParticipantAttributesBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsCobrowsesessionParticipantAttributesBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/cobrowsesessions/{conversationId}/participants/{participantId}/attributes][%d] patchConversationsCobrowsesessionParticipantAttributesBadRequest  %+v", 400, o.Payload)
}

func (o *PatchConversationsCobrowsesessionParticipantAttributesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsCobrowsesessionParticipantAttributesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsCobrowsesessionParticipantAttributesUnauthorized creates a PatchConversationsCobrowsesessionParticipantAttributesUnauthorized with default headers values
func NewPatchConversationsCobrowsesessionParticipantAttributesUnauthorized() *PatchConversationsCobrowsesessionParticipantAttributesUnauthorized {
	return &PatchConversationsCobrowsesessionParticipantAttributesUnauthorized{}
}

/*PatchConversationsCobrowsesessionParticipantAttributesUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PatchConversationsCobrowsesessionParticipantAttributesUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsCobrowsesessionParticipantAttributesUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/cobrowsesessions/{conversationId}/participants/{participantId}/attributes][%d] patchConversationsCobrowsesessionParticipantAttributesUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchConversationsCobrowsesessionParticipantAttributesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsCobrowsesessionParticipantAttributesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsCobrowsesessionParticipantAttributesForbidden creates a PatchConversationsCobrowsesessionParticipantAttributesForbidden with default headers values
func NewPatchConversationsCobrowsesessionParticipantAttributesForbidden() *PatchConversationsCobrowsesessionParticipantAttributesForbidden {
	return &PatchConversationsCobrowsesessionParticipantAttributesForbidden{}
}

/*PatchConversationsCobrowsesessionParticipantAttributesForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PatchConversationsCobrowsesessionParticipantAttributesForbidden struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsCobrowsesessionParticipantAttributesForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/cobrowsesessions/{conversationId}/participants/{participantId}/attributes][%d] patchConversationsCobrowsesessionParticipantAttributesForbidden  %+v", 403, o.Payload)
}

func (o *PatchConversationsCobrowsesessionParticipantAttributesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsCobrowsesessionParticipantAttributesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsCobrowsesessionParticipantAttributesNotFound creates a PatchConversationsCobrowsesessionParticipantAttributesNotFound with default headers values
func NewPatchConversationsCobrowsesessionParticipantAttributesNotFound() *PatchConversationsCobrowsesessionParticipantAttributesNotFound {
	return &PatchConversationsCobrowsesessionParticipantAttributesNotFound{}
}

/*PatchConversationsCobrowsesessionParticipantAttributesNotFound handles this case with default header values.

The requested resource was not found.
*/
type PatchConversationsCobrowsesessionParticipantAttributesNotFound struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsCobrowsesessionParticipantAttributesNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/cobrowsesessions/{conversationId}/participants/{participantId}/attributes][%d] patchConversationsCobrowsesessionParticipantAttributesNotFound  %+v", 404, o.Payload)
}

func (o *PatchConversationsCobrowsesessionParticipantAttributesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsCobrowsesessionParticipantAttributesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsCobrowsesessionParticipantAttributesRequestEntityTooLarge creates a PatchConversationsCobrowsesessionParticipantAttributesRequestEntityTooLarge with default headers values
func NewPatchConversationsCobrowsesessionParticipantAttributesRequestEntityTooLarge() *PatchConversationsCobrowsesessionParticipantAttributesRequestEntityTooLarge {
	return &PatchConversationsCobrowsesessionParticipantAttributesRequestEntityTooLarge{}
}

/*PatchConversationsCobrowsesessionParticipantAttributesRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PatchConversationsCobrowsesessionParticipantAttributesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsCobrowsesessionParticipantAttributesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/cobrowsesessions/{conversationId}/participants/{participantId}/attributes][%d] patchConversationsCobrowsesessionParticipantAttributesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchConversationsCobrowsesessionParticipantAttributesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsCobrowsesessionParticipantAttributesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsCobrowsesessionParticipantAttributesUnsupportedMediaType creates a PatchConversationsCobrowsesessionParticipantAttributesUnsupportedMediaType with default headers values
func NewPatchConversationsCobrowsesessionParticipantAttributesUnsupportedMediaType() *PatchConversationsCobrowsesessionParticipantAttributesUnsupportedMediaType {
	return &PatchConversationsCobrowsesessionParticipantAttributesUnsupportedMediaType{}
}

/*PatchConversationsCobrowsesessionParticipantAttributesUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PatchConversationsCobrowsesessionParticipantAttributesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsCobrowsesessionParticipantAttributesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/cobrowsesessions/{conversationId}/participants/{participantId}/attributes][%d] patchConversationsCobrowsesessionParticipantAttributesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchConversationsCobrowsesessionParticipantAttributesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsCobrowsesessionParticipantAttributesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsCobrowsesessionParticipantAttributesTooManyRequests creates a PatchConversationsCobrowsesessionParticipantAttributesTooManyRequests with default headers values
func NewPatchConversationsCobrowsesessionParticipantAttributesTooManyRequests() *PatchConversationsCobrowsesessionParticipantAttributesTooManyRequests {
	return &PatchConversationsCobrowsesessionParticipantAttributesTooManyRequests{}
}

/*PatchConversationsCobrowsesessionParticipantAttributesTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PatchConversationsCobrowsesessionParticipantAttributesTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsCobrowsesessionParticipantAttributesTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/cobrowsesessions/{conversationId}/participants/{participantId}/attributes][%d] patchConversationsCobrowsesessionParticipantAttributesTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchConversationsCobrowsesessionParticipantAttributesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsCobrowsesessionParticipantAttributesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsCobrowsesessionParticipantAttributesInternalServerError creates a PatchConversationsCobrowsesessionParticipantAttributesInternalServerError with default headers values
func NewPatchConversationsCobrowsesessionParticipantAttributesInternalServerError() *PatchConversationsCobrowsesessionParticipantAttributesInternalServerError {
	return &PatchConversationsCobrowsesessionParticipantAttributesInternalServerError{}
}

/*PatchConversationsCobrowsesessionParticipantAttributesInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PatchConversationsCobrowsesessionParticipantAttributesInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsCobrowsesessionParticipantAttributesInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/cobrowsesessions/{conversationId}/participants/{participantId}/attributes][%d] patchConversationsCobrowsesessionParticipantAttributesInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchConversationsCobrowsesessionParticipantAttributesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsCobrowsesessionParticipantAttributesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsCobrowsesessionParticipantAttributesServiceUnavailable creates a PatchConversationsCobrowsesessionParticipantAttributesServiceUnavailable with default headers values
func NewPatchConversationsCobrowsesessionParticipantAttributesServiceUnavailable() *PatchConversationsCobrowsesessionParticipantAttributesServiceUnavailable {
	return &PatchConversationsCobrowsesessionParticipantAttributesServiceUnavailable{}
}

/*PatchConversationsCobrowsesessionParticipantAttributesServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PatchConversationsCobrowsesessionParticipantAttributesServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsCobrowsesessionParticipantAttributesServiceUnavailable) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/cobrowsesessions/{conversationId}/participants/{participantId}/attributes][%d] patchConversationsCobrowsesessionParticipantAttributesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchConversationsCobrowsesessionParticipantAttributesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsCobrowsesessionParticipantAttributesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsCobrowsesessionParticipantAttributesGatewayTimeout creates a PatchConversationsCobrowsesessionParticipantAttributesGatewayTimeout with default headers values
func NewPatchConversationsCobrowsesessionParticipantAttributesGatewayTimeout() *PatchConversationsCobrowsesessionParticipantAttributesGatewayTimeout {
	return &PatchConversationsCobrowsesessionParticipantAttributesGatewayTimeout{}
}

/*PatchConversationsCobrowsesessionParticipantAttributesGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PatchConversationsCobrowsesessionParticipantAttributesGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsCobrowsesessionParticipantAttributesGatewayTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/cobrowsesessions/{conversationId}/participants/{participantId}/attributes][%d] patchConversationsCobrowsesessionParticipantAttributesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchConversationsCobrowsesessionParticipantAttributesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsCobrowsesessionParticipantAttributesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

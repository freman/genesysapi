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

// PatchConversationsEmailParticipantAttributesReader is a Reader for the PatchConversationsEmailParticipantAttributes structure.
type PatchConversationsEmailParticipantAttributesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchConversationsEmailParticipantAttributesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewPatchConversationsEmailParticipantAttributesAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchConversationsEmailParticipantAttributesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchConversationsEmailParticipantAttributesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchConversationsEmailParticipantAttributesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchConversationsEmailParticipantAttributesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPatchConversationsEmailParticipantAttributesRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPatchConversationsEmailParticipantAttributesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPatchConversationsEmailParticipantAttributesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPatchConversationsEmailParticipantAttributesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchConversationsEmailParticipantAttributesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPatchConversationsEmailParticipantAttributesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPatchConversationsEmailParticipantAttributesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchConversationsEmailParticipantAttributesAccepted creates a PatchConversationsEmailParticipantAttributesAccepted with default headers values
func NewPatchConversationsEmailParticipantAttributesAccepted() *PatchConversationsEmailParticipantAttributesAccepted {
	return &PatchConversationsEmailParticipantAttributesAccepted{}
}

/*PatchConversationsEmailParticipantAttributesAccepted handles this case with default header values.

Accepted
*/
type PatchConversationsEmailParticipantAttributesAccepted struct {
}

func (o *PatchConversationsEmailParticipantAttributesAccepted) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/emails/{conversationId}/participants/{participantId}/attributes][%d] patchConversationsEmailParticipantAttributesAccepted ", 202)
}

func (o *PatchConversationsEmailParticipantAttributesAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchConversationsEmailParticipantAttributesBadRequest creates a PatchConversationsEmailParticipantAttributesBadRequest with default headers values
func NewPatchConversationsEmailParticipantAttributesBadRequest() *PatchConversationsEmailParticipantAttributesBadRequest {
	return &PatchConversationsEmailParticipantAttributesBadRequest{}
}

/*PatchConversationsEmailParticipantAttributesBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PatchConversationsEmailParticipantAttributesBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsEmailParticipantAttributesBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/emails/{conversationId}/participants/{participantId}/attributes][%d] patchConversationsEmailParticipantAttributesBadRequest  %+v", 400, o.Payload)
}

func (o *PatchConversationsEmailParticipantAttributesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsEmailParticipantAttributesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsEmailParticipantAttributesUnauthorized creates a PatchConversationsEmailParticipantAttributesUnauthorized with default headers values
func NewPatchConversationsEmailParticipantAttributesUnauthorized() *PatchConversationsEmailParticipantAttributesUnauthorized {
	return &PatchConversationsEmailParticipantAttributesUnauthorized{}
}

/*PatchConversationsEmailParticipantAttributesUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PatchConversationsEmailParticipantAttributesUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsEmailParticipantAttributesUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/emails/{conversationId}/participants/{participantId}/attributes][%d] patchConversationsEmailParticipantAttributesUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchConversationsEmailParticipantAttributesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsEmailParticipantAttributesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsEmailParticipantAttributesForbidden creates a PatchConversationsEmailParticipantAttributesForbidden with default headers values
func NewPatchConversationsEmailParticipantAttributesForbidden() *PatchConversationsEmailParticipantAttributesForbidden {
	return &PatchConversationsEmailParticipantAttributesForbidden{}
}

/*PatchConversationsEmailParticipantAttributesForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PatchConversationsEmailParticipantAttributesForbidden struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsEmailParticipantAttributesForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/emails/{conversationId}/participants/{participantId}/attributes][%d] patchConversationsEmailParticipantAttributesForbidden  %+v", 403, o.Payload)
}

func (o *PatchConversationsEmailParticipantAttributesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsEmailParticipantAttributesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsEmailParticipantAttributesNotFound creates a PatchConversationsEmailParticipantAttributesNotFound with default headers values
func NewPatchConversationsEmailParticipantAttributesNotFound() *PatchConversationsEmailParticipantAttributesNotFound {
	return &PatchConversationsEmailParticipantAttributesNotFound{}
}

/*PatchConversationsEmailParticipantAttributesNotFound handles this case with default header values.

The requested resource was not found.
*/
type PatchConversationsEmailParticipantAttributesNotFound struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsEmailParticipantAttributesNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/emails/{conversationId}/participants/{participantId}/attributes][%d] patchConversationsEmailParticipantAttributesNotFound  %+v", 404, o.Payload)
}

func (o *PatchConversationsEmailParticipantAttributesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsEmailParticipantAttributesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsEmailParticipantAttributesRequestTimeout creates a PatchConversationsEmailParticipantAttributesRequestTimeout with default headers values
func NewPatchConversationsEmailParticipantAttributesRequestTimeout() *PatchConversationsEmailParticipantAttributesRequestTimeout {
	return &PatchConversationsEmailParticipantAttributesRequestTimeout{}
}

/*PatchConversationsEmailParticipantAttributesRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PatchConversationsEmailParticipantAttributesRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsEmailParticipantAttributesRequestTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/emails/{conversationId}/participants/{participantId}/attributes][%d] patchConversationsEmailParticipantAttributesRequestTimeout  %+v", 408, o.Payload)
}

func (o *PatchConversationsEmailParticipantAttributesRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsEmailParticipantAttributesRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsEmailParticipantAttributesRequestEntityTooLarge creates a PatchConversationsEmailParticipantAttributesRequestEntityTooLarge with default headers values
func NewPatchConversationsEmailParticipantAttributesRequestEntityTooLarge() *PatchConversationsEmailParticipantAttributesRequestEntityTooLarge {
	return &PatchConversationsEmailParticipantAttributesRequestEntityTooLarge{}
}

/*PatchConversationsEmailParticipantAttributesRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PatchConversationsEmailParticipantAttributesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsEmailParticipantAttributesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/emails/{conversationId}/participants/{participantId}/attributes][%d] patchConversationsEmailParticipantAttributesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchConversationsEmailParticipantAttributesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsEmailParticipantAttributesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsEmailParticipantAttributesUnsupportedMediaType creates a PatchConversationsEmailParticipantAttributesUnsupportedMediaType with default headers values
func NewPatchConversationsEmailParticipantAttributesUnsupportedMediaType() *PatchConversationsEmailParticipantAttributesUnsupportedMediaType {
	return &PatchConversationsEmailParticipantAttributesUnsupportedMediaType{}
}

/*PatchConversationsEmailParticipantAttributesUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PatchConversationsEmailParticipantAttributesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsEmailParticipantAttributesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/emails/{conversationId}/participants/{participantId}/attributes][%d] patchConversationsEmailParticipantAttributesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchConversationsEmailParticipantAttributesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsEmailParticipantAttributesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsEmailParticipantAttributesTooManyRequests creates a PatchConversationsEmailParticipantAttributesTooManyRequests with default headers values
func NewPatchConversationsEmailParticipantAttributesTooManyRequests() *PatchConversationsEmailParticipantAttributesTooManyRequests {
	return &PatchConversationsEmailParticipantAttributesTooManyRequests{}
}

/*PatchConversationsEmailParticipantAttributesTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PatchConversationsEmailParticipantAttributesTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsEmailParticipantAttributesTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/emails/{conversationId}/participants/{participantId}/attributes][%d] patchConversationsEmailParticipantAttributesTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchConversationsEmailParticipantAttributesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsEmailParticipantAttributesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsEmailParticipantAttributesInternalServerError creates a PatchConversationsEmailParticipantAttributesInternalServerError with default headers values
func NewPatchConversationsEmailParticipantAttributesInternalServerError() *PatchConversationsEmailParticipantAttributesInternalServerError {
	return &PatchConversationsEmailParticipantAttributesInternalServerError{}
}

/*PatchConversationsEmailParticipantAttributesInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PatchConversationsEmailParticipantAttributesInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsEmailParticipantAttributesInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/emails/{conversationId}/participants/{participantId}/attributes][%d] patchConversationsEmailParticipantAttributesInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchConversationsEmailParticipantAttributesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsEmailParticipantAttributesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsEmailParticipantAttributesServiceUnavailable creates a PatchConversationsEmailParticipantAttributesServiceUnavailable with default headers values
func NewPatchConversationsEmailParticipantAttributesServiceUnavailable() *PatchConversationsEmailParticipantAttributesServiceUnavailable {
	return &PatchConversationsEmailParticipantAttributesServiceUnavailable{}
}

/*PatchConversationsEmailParticipantAttributesServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PatchConversationsEmailParticipantAttributesServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsEmailParticipantAttributesServiceUnavailable) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/emails/{conversationId}/participants/{participantId}/attributes][%d] patchConversationsEmailParticipantAttributesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchConversationsEmailParticipantAttributesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsEmailParticipantAttributesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsEmailParticipantAttributesGatewayTimeout creates a PatchConversationsEmailParticipantAttributesGatewayTimeout with default headers values
func NewPatchConversationsEmailParticipantAttributesGatewayTimeout() *PatchConversationsEmailParticipantAttributesGatewayTimeout {
	return &PatchConversationsEmailParticipantAttributesGatewayTimeout{}
}

/*PatchConversationsEmailParticipantAttributesGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PatchConversationsEmailParticipantAttributesGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsEmailParticipantAttributesGatewayTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/emails/{conversationId}/participants/{participantId}/attributes][%d] patchConversationsEmailParticipantAttributesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchConversationsEmailParticipantAttributesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsEmailParticipantAttributesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

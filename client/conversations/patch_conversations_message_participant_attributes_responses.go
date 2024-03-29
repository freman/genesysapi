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

// PatchConversationsMessageParticipantAttributesReader is a Reader for the PatchConversationsMessageParticipantAttributes structure.
type PatchConversationsMessageParticipantAttributesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchConversationsMessageParticipantAttributesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewPatchConversationsMessageParticipantAttributesAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchConversationsMessageParticipantAttributesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchConversationsMessageParticipantAttributesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchConversationsMessageParticipantAttributesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchConversationsMessageParticipantAttributesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPatchConversationsMessageParticipantAttributesRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPatchConversationsMessageParticipantAttributesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPatchConversationsMessageParticipantAttributesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPatchConversationsMessageParticipantAttributesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchConversationsMessageParticipantAttributesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPatchConversationsMessageParticipantAttributesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPatchConversationsMessageParticipantAttributesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchConversationsMessageParticipantAttributesAccepted creates a PatchConversationsMessageParticipantAttributesAccepted with default headers values
func NewPatchConversationsMessageParticipantAttributesAccepted() *PatchConversationsMessageParticipantAttributesAccepted {
	return &PatchConversationsMessageParticipantAttributesAccepted{}
}

/*
PatchConversationsMessageParticipantAttributesAccepted describes a response with status code 202, with default header values.

Accepted
*/
type PatchConversationsMessageParticipantAttributesAccepted struct {
}

// IsSuccess returns true when this patch conversations message participant attributes accepted response has a 2xx status code
func (o *PatchConversationsMessageParticipantAttributesAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch conversations message participant attributes accepted response has a 3xx status code
func (o *PatchConversationsMessageParticipantAttributesAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversations message participant attributes accepted response has a 4xx status code
func (o *PatchConversationsMessageParticipantAttributesAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch conversations message participant attributes accepted response has a 5xx status code
func (o *PatchConversationsMessageParticipantAttributesAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this patch conversations message participant attributes accepted response a status code equal to that given
func (o *PatchConversationsMessageParticipantAttributesAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *PatchConversationsMessageParticipantAttributesAccepted) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messages/{conversationId}/participants/{participantId}/attributes][%d] patchConversationsMessageParticipantAttributesAccepted ", 202)
}

func (o *PatchConversationsMessageParticipantAttributesAccepted) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messages/{conversationId}/participants/{participantId}/attributes][%d] patchConversationsMessageParticipantAttributesAccepted ", 202)
}

func (o *PatchConversationsMessageParticipantAttributesAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchConversationsMessageParticipantAttributesBadRequest creates a PatchConversationsMessageParticipantAttributesBadRequest with default headers values
func NewPatchConversationsMessageParticipantAttributesBadRequest() *PatchConversationsMessageParticipantAttributesBadRequest {
	return &PatchConversationsMessageParticipantAttributesBadRequest{}
}

/*
PatchConversationsMessageParticipantAttributesBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PatchConversationsMessageParticipantAttributesBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversations message participant attributes bad request response has a 2xx status code
func (o *PatchConversationsMessageParticipantAttributesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversations message participant attributes bad request response has a 3xx status code
func (o *PatchConversationsMessageParticipantAttributesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversations message participant attributes bad request response has a 4xx status code
func (o *PatchConversationsMessageParticipantAttributesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch conversations message participant attributes bad request response has a 5xx status code
func (o *PatchConversationsMessageParticipantAttributesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this patch conversations message participant attributes bad request response a status code equal to that given
func (o *PatchConversationsMessageParticipantAttributesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PatchConversationsMessageParticipantAttributesBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messages/{conversationId}/participants/{participantId}/attributes][%d] patchConversationsMessageParticipantAttributesBadRequest  %+v", 400, o.Payload)
}

func (o *PatchConversationsMessageParticipantAttributesBadRequest) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messages/{conversationId}/participants/{participantId}/attributes][%d] patchConversationsMessageParticipantAttributesBadRequest  %+v", 400, o.Payload)
}

func (o *PatchConversationsMessageParticipantAttributesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsMessageParticipantAttributesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsMessageParticipantAttributesUnauthorized creates a PatchConversationsMessageParticipantAttributesUnauthorized with default headers values
func NewPatchConversationsMessageParticipantAttributesUnauthorized() *PatchConversationsMessageParticipantAttributesUnauthorized {
	return &PatchConversationsMessageParticipantAttributesUnauthorized{}
}

/*
PatchConversationsMessageParticipantAttributesUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PatchConversationsMessageParticipantAttributesUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversations message participant attributes unauthorized response has a 2xx status code
func (o *PatchConversationsMessageParticipantAttributesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversations message participant attributes unauthorized response has a 3xx status code
func (o *PatchConversationsMessageParticipantAttributesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversations message participant attributes unauthorized response has a 4xx status code
func (o *PatchConversationsMessageParticipantAttributesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch conversations message participant attributes unauthorized response has a 5xx status code
func (o *PatchConversationsMessageParticipantAttributesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this patch conversations message participant attributes unauthorized response a status code equal to that given
func (o *PatchConversationsMessageParticipantAttributesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PatchConversationsMessageParticipantAttributesUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messages/{conversationId}/participants/{participantId}/attributes][%d] patchConversationsMessageParticipantAttributesUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchConversationsMessageParticipantAttributesUnauthorized) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messages/{conversationId}/participants/{participantId}/attributes][%d] patchConversationsMessageParticipantAttributesUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchConversationsMessageParticipantAttributesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsMessageParticipantAttributesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsMessageParticipantAttributesForbidden creates a PatchConversationsMessageParticipantAttributesForbidden with default headers values
func NewPatchConversationsMessageParticipantAttributesForbidden() *PatchConversationsMessageParticipantAttributesForbidden {
	return &PatchConversationsMessageParticipantAttributesForbidden{}
}

/*
PatchConversationsMessageParticipantAttributesForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PatchConversationsMessageParticipantAttributesForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversations message participant attributes forbidden response has a 2xx status code
func (o *PatchConversationsMessageParticipantAttributesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversations message participant attributes forbidden response has a 3xx status code
func (o *PatchConversationsMessageParticipantAttributesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversations message participant attributes forbidden response has a 4xx status code
func (o *PatchConversationsMessageParticipantAttributesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch conversations message participant attributes forbidden response has a 5xx status code
func (o *PatchConversationsMessageParticipantAttributesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this patch conversations message participant attributes forbidden response a status code equal to that given
func (o *PatchConversationsMessageParticipantAttributesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PatchConversationsMessageParticipantAttributesForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messages/{conversationId}/participants/{participantId}/attributes][%d] patchConversationsMessageParticipantAttributesForbidden  %+v", 403, o.Payload)
}

func (o *PatchConversationsMessageParticipantAttributesForbidden) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messages/{conversationId}/participants/{participantId}/attributes][%d] patchConversationsMessageParticipantAttributesForbidden  %+v", 403, o.Payload)
}

func (o *PatchConversationsMessageParticipantAttributesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsMessageParticipantAttributesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsMessageParticipantAttributesNotFound creates a PatchConversationsMessageParticipantAttributesNotFound with default headers values
func NewPatchConversationsMessageParticipantAttributesNotFound() *PatchConversationsMessageParticipantAttributesNotFound {
	return &PatchConversationsMessageParticipantAttributesNotFound{}
}

/*
PatchConversationsMessageParticipantAttributesNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PatchConversationsMessageParticipantAttributesNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversations message participant attributes not found response has a 2xx status code
func (o *PatchConversationsMessageParticipantAttributesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversations message participant attributes not found response has a 3xx status code
func (o *PatchConversationsMessageParticipantAttributesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversations message participant attributes not found response has a 4xx status code
func (o *PatchConversationsMessageParticipantAttributesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch conversations message participant attributes not found response has a 5xx status code
func (o *PatchConversationsMessageParticipantAttributesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this patch conversations message participant attributes not found response a status code equal to that given
func (o *PatchConversationsMessageParticipantAttributesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PatchConversationsMessageParticipantAttributesNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messages/{conversationId}/participants/{participantId}/attributes][%d] patchConversationsMessageParticipantAttributesNotFound  %+v", 404, o.Payload)
}

func (o *PatchConversationsMessageParticipantAttributesNotFound) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messages/{conversationId}/participants/{participantId}/attributes][%d] patchConversationsMessageParticipantAttributesNotFound  %+v", 404, o.Payload)
}

func (o *PatchConversationsMessageParticipantAttributesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsMessageParticipantAttributesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsMessageParticipantAttributesRequestTimeout creates a PatchConversationsMessageParticipantAttributesRequestTimeout with default headers values
func NewPatchConversationsMessageParticipantAttributesRequestTimeout() *PatchConversationsMessageParticipantAttributesRequestTimeout {
	return &PatchConversationsMessageParticipantAttributesRequestTimeout{}
}

/*
PatchConversationsMessageParticipantAttributesRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PatchConversationsMessageParticipantAttributesRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversations message participant attributes request timeout response has a 2xx status code
func (o *PatchConversationsMessageParticipantAttributesRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversations message participant attributes request timeout response has a 3xx status code
func (o *PatchConversationsMessageParticipantAttributesRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversations message participant attributes request timeout response has a 4xx status code
func (o *PatchConversationsMessageParticipantAttributesRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch conversations message participant attributes request timeout response has a 5xx status code
func (o *PatchConversationsMessageParticipantAttributesRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this patch conversations message participant attributes request timeout response a status code equal to that given
func (o *PatchConversationsMessageParticipantAttributesRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PatchConversationsMessageParticipantAttributesRequestTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messages/{conversationId}/participants/{participantId}/attributes][%d] patchConversationsMessageParticipantAttributesRequestTimeout  %+v", 408, o.Payload)
}

func (o *PatchConversationsMessageParticipantAttributesRequestTimeout) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messages/{conversationId}/participants/{participantId}/attributes][%d] patchConversationsMessageParticipantAttributesRequestTimeout  %+v", 408, o.Payload)
}

func (o *PatchConversationsMessageParticipantAttributesRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsMessageParticipantAttributesRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsMessageParticipantAttributesRequestEntityTooLarge creates a PatchConversationsMessageParticipantAttributesRequestEntityTooLarge with default headers values
func NewPatchConversationsMessageParticipantAttributesRequestEntityTooLarge() *PatchConversationsMessageParticipantAttributesRequestEntityTooLarge {
	return &PatchConversationsMessageParticipantAttributesRequestEntityTooLarge{}
}

/*
PatchConversationsMessageParticipantAttributesRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PatchConversationsMessageParticipantAttributesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversations message participant attributes request entity too large response has a 2xx status code
func (o *PatchConversationsMessageParticipantAttributesRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversations message participant attributes request entity too large response has a 3xx status code
func (o *PatchConversationsMessageParticipantAttributesRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversations message participant attributes request entity too large response has a 4xx status code
func (o *PatchConversationsMessageParticipantAttributesRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch conversations message participant attributes request entity too large response has a 5xx status code
func (o *PatchConversationsMessageParticipantAttributesRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this patch conversations message participant attributes request entity too large response a status code equal to that given
func (o *PatchConversationsMessageParticipantAttributesRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PatchConversationsMessageParticipantAttributesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messages/{conversationId}/participants/{participantId}/attributes][%d] patchConversationsMessageParticipantAttributesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchConversationsMessageParticipantAttributesRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messages/{conversationId}/participants/{participantId}/attributes][%d] patchConversationsMessageParticipantAttributesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchConversationsMessageParticipantAttributesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsMessageParticipantAttributesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsMessageParticipantAttributesUnsupportedMediaType creates a PatchConversationsMessageParticipantAttributesUnsupportedMediaType with default headers values
func NewPatchConversationsMessageParticipantAttributesUnsupportedMediaType() *PatchConversationsMessageParticipantAttributesUnsupportedMediaType {
	return &PatchConversationsMessageParticipantAttributesUnsupportedMediaType{}
}

/*
PatchConversationsMessageParticipantAttributesUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PatchConversationsMessageParticipantAttributesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversations message participant attributes unsupported media type response has a 2xx status code
func (o *PatchConversationsMessageParticipantAttributesUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversations message participant attributes unsupported media type response has a 3xx status code
func (o *PatchConversationsMessageParticipantAttributesUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversations message participant attributes unsupported media type response has a 4xx status code
func (o *PatchConversationsMessageParticipantAttributesUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch conversations message participant attributes unsupported media type response has a 5xx status code
func (o *PatchConversationsMessageParticipantAttributesUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this patch conversations message participant attributes unsupported media type response a status code equal to that given
func (o *PatchConversationsMessageParticipantAttributesUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PatchConversationsMessageParticipantAttributesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messages/{conversationId}/participants/{participantId}/attributes][%d] patchConversationsMessageParticipantAttributesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchConversationsMessageParticipantAttributesUnsupportedMediaType) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messages/{conversationId}/participants/{participantId}/attributes][%d] patchConversationsMessageParticipantAttributesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchConversationsMessageParticipantAttributesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsMessageParticipantAttributesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsMessageParticipantAttributesTooManyRequests creates a PatchConversationsMessageParticipantAttributesTooManyRequests with default headers values
func NewPatchConversationsMessageParticipantAttributesTooManyRequests() *PatchConversationsMessageParticipantAttributesTooManyRequests {
	return &PatchConversationsMessageParticipantAttributesTooManyRequests{}
}

/*
PatchConversationsMessageParticipantAttributesTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PatchConversationsMessageParticipantAttributesTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversations message participant attributes too many requests response has a 2xx status code
func (o *PatchConversationsMessageParticipantAttributesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversations message participant attributes too many requests response has a 3xx status code
func (o *PatchConversationsMessageParticipantAttributesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversations message participant attributes too many requests response has a 4xx status code
func (o *PatchConversationsMessageParticipantAttributesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch conversations message participant attributes too many requests response has a 5xx status code
func (o *PatchConversationsMessageParticipantAttributesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this patch conversations message participant attributes too many requests response a status code equal to that given
func (o *PatchConversationsMessageParticipantAttributesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PatchConversationsMessageParticipantAttributesTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messages/{conversationId}/participants/{participantId}/attributes][%d] patchConversationsMessageParticipantAttributesTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchConversationsMessageParticipantAttributesTooManyRequests) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messages/{conversationId}/participants/{participantId}/attributes][%d] patchConversationsMessageParticipantAttributesTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchConversationsMessageParticipantAttributesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsMessageParticipantAttributesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsMessageParticipantAttributesInternalServerError creates a PatchConversationsMessageParticipantAttributesInternalServerError with default headers values
func NewPatchConversationsMessageParticipantAttributesInternalServerError() *PatchConversationsMessageParticipantAttributesInternalServerError {
	return &PatchConversationsMessageParticipantAttributesInternalServerError{}
}

/*
PatchConversationsMessageParticipantAttributesInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PatchConversationsMessageParticipantAttributesInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversations message participant attributes internal server error response has a 2xx status code
func (o *PatchConversationsMessageParticipantAttributesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversations message participant attributes internal server error response has a 3xx status code
func (o *PatchConversationsMessageParticipantAttributesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversations message participant attributes internal server error response has a 4xx status code
func (o *PatchConversationsMessageParticipantAttributesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch conversations message participant attributes internal server error response has a 5xx status code
func (o *PatchConversationsMessageParticipantAttributesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this patch conversations message participant attributes internal server error response a status code equal to that given
func (o *PatchConversationsMessageParticipantAttributesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PatchConversationsMessageParticipantAttributesInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messages/{conversationId}/participants/{participantId}/attributes][%d] patchConversationsMessageParticipantAttributesInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchConversationsMessageParticipantAttributesInternalServerError) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messages/{conversationId}/participants/{participantId}/attributes][%d] patchConversationsMessageParticipantAttributesInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchConversationsMessageParticipantAttributesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsMessageParticipantAttributesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsMessageParticipantAttributesServiceUnavailable creates a PatchConversationsMessageParticipantAttributesServiceUnavailable with default headers values
func NewPatchConversationsMessageParticipantAttributesServiceUnavailable() *PatchConversationsMessageParticipantAttributesServiceUnavailable {
	return &PatchConversationsMessageParticipantAttributesServiceUnavailable{}
}

/*
PatchConversationsMessageParticipantAttributesServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PatchConversationsMessageParticipantAttributesServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversations message participant attributes service unavailable response has a 2xx status code
func (o *PatchConversationsMessageParticipantAttributesServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversations message participant attributes service unavailable response has a 3xx status code
func (o *PatchConversationsMessageParticipantAttributesServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversations message participant attributes service unavailable response has a 4xx status code
func (o *PatchConversationsMessageParticipantAttributesServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch conversations message participant attributes service unavailable response has a 5xx status code
func (o *PatchConversationsMessageParticipantAttributesServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this patch conversations message participant attributes service unavailable response a status code equal to that given
func (o *PatchConversationsMessageParticipantAttributesServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PatchConversationsMessageParticipantAttributesServiceUnavailable) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messages/{conversationId}/participants/{participantId}/attributes][%d] patchConversationsMessageParticipantAttributesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchConversationsMessageParticipantAttributesServiceUnavailable) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messages/{conversationId}/participants/{participantId}/attributes][%d] patchConversationsMessageParticipantAttributesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchConversationsMessageParticipantAttributesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsMessageParticipantAttributesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsMessageParticipantAttributesGatewayTimeout creates a PatchConversationsMessageParticipantAttributesGatewayTimeout with default headers values
func NewPatchConversationsMessageParticipantAttributesGatewayTimeout() *PatchConversationsMessageParticipantAttributesGatewayTimeout {
	return &PatchConversationsMessageParticipantAttributesGatewayTimeout{}
}

/*
PatchConversationsMessageParticipantAttributesGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PatchConversationsMessageParticipantAttributesGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversations message participant attributes gateway timeout response has a 2xx status code
func (o *PatchConversationsMessageParticipantAttributesGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversations message participant attributes gateway timeout response has a 3xx status code
func (o *PatchConversationsMessageParticipantAttributesGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversations message participant attributes gateway timeout response has a 4xx status code
func (o *PatchConversationsMessageParticipantAttributesGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch conversations message participant attributes gateway timeout response has a 5xx status code
func (o *PatchConversationsMessageParticipantAttributesGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this patch conversations message participant attributes gateway timeout response a status code equal to that given
func (o *PatchConversationsMessageParticipantAttributesGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PatchConversationsMessageParticipantAttributesGatewayTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messages/{conversationId}/participants/{participantId}/attributes][%d] patchConversationsMessageParticipantAttributesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchConversationsMessageParticipantAttributesGatewayTimeout) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messages/{conversationId}/participants/{participantId}/attributes][%d] patchConversationsMessageParticipantAttributesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchConversationsMessageParticipantAttributesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsMessageParticipantAttributesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

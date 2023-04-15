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
	case 408:
		result := NewPatchConversationsMessageParticipantRequestTimeout()
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

/*
PatchConversationsMessageParticipantAccepted describes a response with status code 202, with default header values.

Accepted
*/
type PatchConversationsMessageParticipantAccepted struct {
}

// IsSuccess returns true when this patch conversations message participant accepted response has a 2xx status code
func (o *PatchConversationsMessageParticipantAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch conversations message participant accepted response has a 3xx status code
func (o *PatchConversationsMessageParticipantAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversations message participant accepted response has a 4xx status code
func (o *PatchConversationsMessageParticipantAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch conversations message participant accepted response has a 5xx status code
func (o *PatchConversationsMessageParticipantAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this patch conversations message participant accepted response a status code equal to that given
func (o *PatchConversationsMessageParticipantAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *PatchConversationsMessageParticipantAccepted) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messages/{conversationId}/participants/{participantId}][%d] patchConversationsMessageParticipantAccepted ", 202)
}

func (o *PatchConversationsMessageParticipantAccepted) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messages/{conversationId}/participants/{participantId}][%d] patchConversationsMessageParticipantAccepted ", 202)
}

func (o *PatchConversationsMessageParticipantAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchConversationsMessageParticipantBadRequest creates a PatchConversationsMessageParticipantBadRequest with default headers values
func NewPatchConversationsMessageParticipantBadRequest() *PatchConversationsMessageParticipantBadRequest {
	return &PatchConversationsMessageParticipantBadRequest{}
}

/*
PatchConversationsMessageParticipantBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PatchConversationsMessageParticipantBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversations message participant bad request response has a 2xx status code
func (o *PatchConversationsMessageParticipantBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversations message participant bad request response has a 3xx status code
func (o *PatchConversationsMessageParticipantBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversations message participant bad request response has a 4xx status code
func (o *PatchConversationsMessageParticipantBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch conversations message participant bad request response has a 5xx status code
func (o *PatchConversationsMessageParticipantBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this patch conversations message participant bad request response a status code equal to that given
func (o *PatchConversationsMessageParticipantBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PatchConversationsMessageParticipantBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messages/{conversationId}/participants/{participantId}][%d] patchConversationsMessageParticipantBadRequest  %+v", 400, o.Payload)
}

func (o *PatchConversationsMessageParticipantBadRequest) String() string {
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

/*
PatchConversationsMessageParticipantUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PatchConversationsMessageParticipantUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversations message participant unauthorized response has a 2xx status code
func (o *PatchConversationsMessageParticipantUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversations message participant unauthorized response has a 3xx status code
func (o *PatchConversationsMessageParticipantUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversations message participant unauthorized response has a 4xx status code
func (o *PatchConversationsMessageParticipantUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch conversations message participant unauthorized response has a 5xx status code
func (o *PatchConversationsMessageParticipantUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this patch conversations message participant unauthorized response a status code equal to that given
func (o *PatchConversationsMessageParticipantUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PatchConversationsMessageParticipantUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messages/{conversationId}/participants/{participantId}][%d] patchConversationsMessageParticipantUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchConversationsMessageParticipantUnauthorized) String() string {
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

/*
PatchConversationsMessageParticipantForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PatchConversationsMessageParticipantForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversations message participant forbidden response has a 2xx status code
func (o *PatchConversationsMessageParticipantForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversations message participant forbidden response has a 3xx status code
func (o *PatchConversationsMessageParticipantForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversations message participant forbidden response has a 4xx status code
func (o *PatchConversationsMessageParticipantForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch conversations message participant forbidden response has a 5xx status code
func (o *PatchConversationsMessageParticipantForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this patch conversations message participant forbidden response a status code equal to that given
func (o *PatchConversationsMessageParticipantForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PatchConversationsMessageParticipantForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messages/{conversationId}/participants/{participantId}][%d] patchConversationsMessageParticipantForbidden  %+v", 403, o.Payload)
}

func (o *PatchConversationsMessageParticipantForbidden) String() string {
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

/*
PatchConversationsMessageParticipantNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PatchConversationsMessageParticipantNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversations message participant not found response has a 2xx status code
func (o *PatchConversationsMessageParticipantNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversations message participant not found response has a 3xx status code
func (o *PatchConversationsMessageParticipantNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversations message participant not found response has a 4xx status code
func (o *PatchConversationsMessageParticipantNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch conversations message participant not found response has a 5xx status code
func (o *PatchConversationsMessageParticipantNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this patch conversations message participant not found response a status code equal to that given
func (o *PatchConversationsMessageParticipantNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PatchConversationsMessageParticipantNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messages/{conversationId}/participants/{participantId}][%d] patchConversationsMessageParticipantNotFound  %+v", 404, o.Payload)
}

func (o *PatchConversationsMessageParticipantNotFound) String() string {
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

// NewPatchConversationsMessageParticipantRequestTimeout creates a PatchConversationsMessageParticipantRequestTimeout with default headers values
func NewPatchConversationsMessageParticipantRequestTimeout() *PatchConversationsMessageParticipantRequestTimeout {
	return &PatchConversationsMessageParticipantRequestTimeout{}
}

/*
PatchConversationsMessageParticipantRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PatchConversationsMessageParticipantRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversations message participant request timeout response has a 2xx status code
func (o *PatchConversationsMessageParticipantRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversations message participant request timeout response has a 3xx status code
func (o *PatchConversationsMessageParticipantRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversations message participant request timeout response has a 4xx status code
func (o *PatchConversationsMessageParticipantRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch conversations message participant request timeout response has a 5xx status code
func (o *PatchConversationsMessageParticipantRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this patch conversations message participant request timeout response a status code equal to that given
func (o *PatchConversationsMessageParticipantRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PatchConversationsMessageParticipantRequestTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messages/{conversationId}/participants/{participantId}][%d] patchConversationsMessageParticipantRequestTimeout  %+v", 408, o.Payload)
}

func (o *PatchConversationsMessageParticipantRequestTimeout) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messages/{conversationId}/participants/{participantId}][%d] patchConversationsMessageParticipantRequestTimeout  %+v", 408, o.Payload)
}

func (o *PatchConversationsMessageParticipantRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsMessageParticipantRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*
PatchConversationsMessageParticipantRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PatchConversationsMessageParticipantRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversations message participant request entity too large response has a 2xx status code
func (o *PatchConversationsMessageParticipantRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversations message participant request entity too large response has a 3xx status code
func (o *PatchConversationsMessageParticipantRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversations message participant request entity too large response has a 4xx status code
func (o *PatchConversationsMessageParticipantRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch conversations message participant request entity too large response has a 5xx status code
func (o *PatchConversationsMessageParticipantRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this patch conversations message participant request entity too large response a status code equal to that given
func (o *PatchConversationsMessageParticipantRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PatchConversationsMessageParticipantRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messages/{conversationId}/participants/{participantId}][%d] patchConversationsMessageParticipantRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchConversationsMessageParticipantRequestEntityTooLarge) String() string {
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

/*
PatchConversationsMessageParticipantUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PatchConversationsMessageParticipantUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversations message participant unsupported media type response has a 2xx status code
func (o *PatchConversationsMessageParticipantUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversations message participant unsupported media type response has a 3xx status code
func (o *PatchConversationsMessageParticipantUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversations message participant unsupported media type response has a 4xx status code
func (o *PatchConversationsMessageParticipantUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch conversations message participant unsupported media type response has a 5xx status code
func (o *PatchConversationsMessageParticipantUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this patch conversations message participant unsupported media type response a status code equal to that given
func (o *PatchConversationsMessageParticipantUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PatchConversationsMessageParticipantUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messages/{conversationId}/participants/{participantId}][%d] patchConversationsMessageParticipantUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchConversationsMessageParticipantUnsupportedMediaType) String() string {
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

/*
PatchConversationsMessageParticipantTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PatchConversationsMessageParticipantTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversations message participant too many requests response has a 2xx status code
func (o *PatchConversationsMessageParticipantTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversations message participant too many requests response has a 3xx status code
func (o *PatchConversationsMessageParticipantTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversations message participant too many requests response has a 4xx status code
func (o *PatchConversationsMessageParticipantTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch conversations message participant too many requests response has a 5xx status code
func (o *PatchConversationsMessageParticipantTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this patch conversations message participant too many requests response a status code equal to that given
func (o *PatchConversationsMessageParticipantTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PatchConversationsMessageParticipantTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messages/{conversationId}/participants/{participantId}][%d] patchConversationsMessageParticipantTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchConversationsMessageParticipantTooManyRequests) String() string {
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

/*
PatchConversationsMessageParticipantInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PatchConversationsMessageParticipantInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversations message participant internal server error response has a 2xx status code
func (o *PatchConversationsMessageParticipantInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversations message participant internal server error response has a 3xx status code
func (o *PatchConversationsMessageParticipantInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversations message participant internal server error response has a 4xx status code
func (o *PatchConversationsMessageParticipantInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch conversations message participant internal server error response has a 5xx status code
func (o *PatchConversationsMessageParticipantInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this patch conversations message participant internal server error response a status code equal to that given
func (o *PatchConversationsMessageParticipantInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PatchConversationsMessageParticipantInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messages/{conversationId}/participants/{participantId}][%d] patchConversationsMessageParticipantInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchConversationsMessageParticipantInternalServerError) String() string {
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

/*
PatchConversationsMessageParticipantServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PatchConversationsMessageParticipantServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversations message participant service unavailable response has a 2xx status code
func (o *PatchConversationsMessageParticipantServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversations message participant service unavailable response has a 3xx status code
func (o *PatchConversationsMessageParticipantServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversations message participant service unavailable response has a 4xx status code
func (o *PatchConversationsMessageParticipantServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch conversations message participant service unavailable response has a 5xx status code
func (o *PatchConversationsMessageParticipantServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this patch conversations message participant service unavailable response a status code equal to that given
func (o *PatchConversationsMessageParticipantServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PatchConversationsMessageParticipantServiceUnavailable) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messages/{conversationId}/participants/{participantId}][%d] patchConversationsMessageParticipantServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchConversationsMessageParticipantServiceUnavailable) String() string {
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

/*
PatchConversationsMessageParticipantGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PatchConversationsMessageParticipantGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversations message participant gateway timeout response has a 2xx status code
func (o *PatchConversationsMessageParticipantGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversations message participant gateway timeout response has a 3xx status code
func (o *PatchConversationsMessageParticipantGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversations message participant gateway timeout response has a 4xx status code
func (o *PatchConversationsMessageParticipantGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch conversations message participant gateway timeout response has a 5xx status code
func (o *PatchConversationsMessageParticipantGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this patch conversations message participant gateway timeout response a status code equal to that given
func (o *PatchConversationsMessageParticipantGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PatchConversationsMessageParticipantGatewayTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messages/{conversationId}/participants/{participantId}][%d] patchConversationsMessageParticipantGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchConversationsMessageParticipantGatewayTimeout) String() string {
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

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

// PatchConversationParticipantReader is a Reader for the PatchConversationParticipant structure.
type PatchConversationParticipantReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchConversationParticipantReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewPatchConversationParticipantAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchConversationParticipantBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchConversationParticipantUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchConversationParticipantForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchConversationParticipantNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPatchConversationParticipantRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPatchConversationParticipantRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPatchConversationParticipantUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPatchConversationParticipantTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchConversationParticipantInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPatchConversationParticipantServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPatchConversationParticipantGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchConversationParticipantAccepted creates a PatchConversationParticipantAccepted with default headers values
func NewPatchConversationParticipantAccepted() *PatchConversationParticipantAccepted {
	return &PatchConversationParticipantAccepted{}
}

/*
PatchConversationParticipantAccepted describes a response with status code 202, with default header values.

Accepted
*/
type PatchConversationParticipantAccepted struct {
}

// IsSuccess returns true when this patch conversation participant accepted response has a 2xx status code
func (o *PatchConversationParticipantAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch conversation participant accepted response has a 3xx status code
func (o *PatchConversationParticipantAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversation participant accepted response has a 4xx status code
func (o *PatchConversationParticipantAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch conversation participant accepted response has a 5xx status code
func (o *PatchConversationParticipantAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this patch conversation participant accepted response a status code equal to that given
func (o *PatchConversationParticipantAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *PatchConversationParticipantAccepted) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/{conversationId}/participants/{participantId}][%d] patchConversationParticipantAccepted ", 202)
}

func (o *PatchConversationParticipantAccepted) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/{conversationId}/participants/{participantId}][%d] patchConversationParticipantAccepted ", 202)
}

func (o *PatchConversationParticipantAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchConversationParticipantBadRequest creates a PatchConversationParticipantBadRequest with default headers values
func NewPatchConversationParticipantBadRequest() *PatchConversationParticipantBadRequest {
	return &PatchConversationParticipantBadRequest{}
}

/*
PatchConversationParticipantBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PatchConversationParticipantBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversation participant bad request response has a 2xx status code
func (o *PatchConversationParticipantBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversation participant bad request response has a 3xx status code
func (o *PatchConversationParticipantBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversation participant bad request response has a 4xx status code
func (o *PatchConversationParticipantBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch conversation participant bad request response has a 5xx status code
func (o *PatchConversationParticipantBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this patch conversation participant bad request response a status code equal to that given
func (o *PatchConversationParticipantBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PatchConversationParticipantBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/{conversationId}/participants/{participantId}][%d] patchConversationParticipantBadRequest  %+v", 400, o.Payload)
}

func (o *PatchConversationParticipantBadRequest) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/{conversationId}/participants/{participantId}][%d] patchConversationParticipantBadRequest  %+v", 400, o.Payload)
}

func (o *PatchConversationParticipantBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationParticipantBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationParticipantUnauthorized creates a PatchConversationParticipantUnauthorized with default headers values
func NewPatchConversationParticipantUnauthorized() *PatchConversationParticipantUnauthorized {
	return &PatchConversationParticipantUnauthorized{}
}

/*
PatchConversationParticipantUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PatchConversationParticipantUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversation participant unauthorized response has a 2xx status code
func (o *PatchConversationParticipantUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversation participant unauthorized response has a 3xx status code
func (o *PatchConversationParticipantUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversation participant unauthorized response has a 4xx status code
func (o *PatchConversationParticipantUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch conversation participant unauthorized response has a 5xx status code
func (o *PatchConversationParticipantUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this patch conversation participant unauthorized response a status code equal to that given
func (o *PatchConversationParticipantUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PatchConversationParticipantUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/{conversationId}/participants/{participantId}][%d] patchConversationParticipantUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchConversationParticipantUnauthorized) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/{conversationId}/participants/{participantId}][%d] patchConversationParticipantUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchConversationParticipantUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationParticipantUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationParticipantForbidden creates a PatchConversationParticipantForbidden with default headers values
func NewPatchConversationParticipantForbidden() *PatchConversationParticipantForbidden {
	return &PatchConversationParticipantForbidden{}
}

/*
PatchConversationParticipantForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PatchConversationParticipantForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversation participant forbidden response has a 2xx status code
func (o *PatchConversationParticipantForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversation participant forbidden response has a 3xx status code
func (o *PatchConversationParticipantForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversation participant forbidden response has a 4xx status code
func (o *PatchConversationParticipantForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch conversation participant forbidden response has a 5xx status code
func (o *PatchConversationParticipantForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this patch conversation participant forbidden response a status code equal to that given
func (o *PatchConversationParticipantForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PatchConversationParticipantForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/{conversationId}/participants/{participantId}][%d] patchConversationParticipantForbidden  %+v", 403, o.Payload)
}

func (o *PatchConversationParticipantForbidden) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/{conversationId}/participants/{participantId}][%d] patchConversationParticipantForbidden  %+v", 403, o.Payload)
}

func (o *PatchConversationParticipantForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationParticipantForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationParticipantNotFound creates a PatchConversationParticipantNotFound with default headers values
func NewPatchConversationParticipantNotFound() *PatchConversationParticipantNotFound {
	return &PatchConversationParticipantNotFound{}
}

/*
PatchConversationParticipantNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PatchConversationParticipantNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversation participant not found response has a 2xx status code
func (o *PatchConversationParticipantNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversation participant not found response has a 3xx status code
func (o *PatchConversationParticipantNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversation participant not found response has a 4xx status code
func (o *PatchConversationParticipantNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch conversation participant not found response has a 5xx status code
func (o *PatchConversationParticipantNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this patch conversation participant not found response a status code equal to that given
func (o *PatchConversationParticipantNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PatchConversationParticipantNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/{conversationId}/participants/{participantId}][%d] patchConversationParticipantNotFound  %+v", 404, o.Payload)
}

func (o *PatchConversationParticipantNotFound) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/{conversationId}/participants/{participantId}][%d] patchConversationParticipantNotFound  %+v", 404, o.Payload)
}

func (o *PatchConversationParticipantNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationParticipantNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationParticipantRequestTimeout creates a PatchConversationParticipantRequestTimeout with default headers values
func NewPatchConversationParticipantRequestTimeout() *PatchConversationParticipantRequestTimeout {
	return &PatchConversationParticipantRequestTimeout{}
}

/*
PatchConversationParticipantRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PatchConversationParticipantRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversation participant request timeout response has a 2xx status code
func (o *PatchConversationParticipantRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversation participant request timeout response has a 3xx status code
func (o *PatchConversationParticipantRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversation participant request timeout response has a 4xx status code
func (o *PatchConversationParticipantRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch conversation participant request timeout response has a 5xx status code
func (o *PatchConversationParticipantRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this patch conversation participant request timeout response a status code equal to that given
func (o *PatchConversationParticipantRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PatchConversationParticipantRequestTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/{conversationId}/participants/{participantId}][%d] patchConversationParticipantRequestTimeout  %+v", 408, o.Payload)
}

func (o *PatchConversationParticipantRequestTimeout) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/{conversationId}/participants/{participantId}][%d] patchConversationParticipantRequestTimeout  %+v", 408, o.Payload)
}

func (o *PatchConversationParticipantRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationParticipantRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationParticipantRequestEntityTooLarge creates a PatchConversationParticipantRequestEntityTooLarge with default headers values
func NewPatchConversationParticipantRequestEntityTooLarge() *PatchConversationParticipantRequestEntityTooLarge {
	return &PatchConversationParticipantRequestEntityTooLarge{}
}

/*
PatchConversationParticipantRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PatchConversationParticipantRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversation participant request entity too large response has a 2xx status code
func (o *PatchConversationParticipantRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversation participant request entity too large response has a 3xx status code
func (o *PatchConversationParticipantRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversation participant request entity too large response has a 4xx status code
func (o *PatchConversationParticipantRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch conversation participant request entity too large response has a 5xx status code
func (o *PatchConversationParticipantRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this patch conversation participant request entity too large response a status code equal to that given
func (o *PatchConversationParticipantRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PatchConversationParticipantRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/{conversationId}/participants/{participantId}][%d] patchConversationParticipantRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchConversationParticipantRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/{conversationId}/participants/{participantId}][%d] patchConversationParticipantRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchConversationParticipantRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationParticipantRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationParticipantUnsupportedMediaType creates a PatchConversationParticipantUnsupportedMediaType with default headers values
func NewPatchConversationParticipantUnsupportedMediaType() *PatchConversationParticipantUnsupportedMediaType {
	return &PatchConversationParticipantUnsupportedMediaType{}
}

/*
PatchConversationParticipantUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PatchConversationParticipantUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversation participant unsupported media type response has a 2xx status code
func (o *PatchConversationParticipantUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversation participant unsupported media type response has a 3xx status code
func (o *PatchConversationParticipantUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversation participant unsupported media type response has a 4xx status code
func (o *PatchConversationParticipantUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch conversation participant unsupported media type response has a 5xx status code
func (o *PatchConversationParticipantUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this patch conversation participant unsupported media type response a status code equal to that given
func (o *PatchConversationParticipantUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PatchConversationParticipantUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/{conversationId}/participants/{participantId}][%d] patchConversationParticipantUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchConversationParticipantUnsupportedMediaType) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/{conversationId}/participants/{participantId}][%d] patchConversationParticipantUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchConversationParticipantUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationParticipantUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationParticipantTooManyRequests creates a PatchConversationParticipantTooManyRequests with default headers values
func NewPatchConversationParticipantTooManyRequests() *PatchConversationParticipantTooManyRequests {
	return &PatchConversationParticipantTooManyRequests{}
}

/*
PatchConversationParticipantTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PatchConversationParticipantTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversation participant too many requests response has a 2xx status code
func (o *PatchConversationParticipantTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversation participant too many requests response has a 3xx status code
func (o *PatchConversationParticipantTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversation participant too many requests response has a 4xx status code
func (o *PatchConversationParticipantTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch conversation participant too many requests response has a 5xx status code
func (o *PatchConversationParticipantTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this patch conversation participant too many requests response a status code equal to that given
func (o *PatchConversationParticipantTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PatchConversationParticipantTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/{conversationId}/participants/{participantId}][%d] patchConversationParticipantTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchConversationParticipantTooManyRequests) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/{conversationId}/participants/{participantId}][%d] patchConversationParticipantTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchConversationParticipantTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationParticipantTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationParticipantInternalServerError creates a PatchConversationParticipantInternalServerError with default headers values
func NewPatchConversationParticipantInternalServerError() *PatchConversationParticipantInternalServerError {
	return &PatchConversationParticipantInternalServerError{}
}

/*
PatchConversationParticipantInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PatchConversationParticipantInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversation participant internal server error response has a 2xx status code
func (o *PatchConversationParticipantInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversation participant internal server error response has a 3xx status code
func (o *PatchConversationParticipantInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversation participant internal server error response has a 4xx status code
func (o *PatchConversationParticipantInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch conversation participant internal server error response has a 5xx status code
func (o *PatchConversationParticipantInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this patch conversation participant internal server error response a status code equal to that given
func (o *PatchConversationParticipantInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PatchConversationParticipantInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/{conversationId}/participants/{participantId}][%d] patchConversationParticipantInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchConversationParticipantInternalServerError) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/{conversationId}/participants/{participantId}][%d] patchConversationParticipantInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchConversationParticipantInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationParticipantInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationParticipantServiceUnavailable creates a PatchConversationParticipantServiceUnavailable with default headers values
func NewPatchConversationParticipantServiceUnavailable() *PatchConversationParticipantServiceUnavailable {
	return &PatchConversationParticipantServiceUnavailable{}
}

/*
PatchConversationParticipantServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PatchConversationParticipantServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversation participant service unavailable response has a 2xx status code
func (o *PatchConversationParticipantServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversation participant service unavailable response has a 3xx status code
func (o *PatchConversationParticipantServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversation participant service unavailable response has a 4xx status code
func (o *PatchConversationParticipantServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch conversation participant service unavailable response has a 5xx status code
func (o *PatchConversationParticipantServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this patch conversation participant service unavailable response a status code equal to that given
func (o *PatchConversationParticipantServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PatchConversationParticipantServiceUnavailable) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/{conversationId}/participants/{participantId}][%d] patchConversationParticipantServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchConversationParticipantServiceUnavailable) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/{conversationId}/participants/{participantId}][%d] patchConversationParticipantServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchConversationParticipantServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationParticipantServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationParticipantGatewayTimeout creates a PatchConversationParticipantGatewayTimeout with default headers values
func NewPatchConversationParticipantGatewayTimeout() *PatchConversationParticipantGatewayTimeout {
	return &PatchConversationParticipantGatewayTimeout{}
}

/*
PatchConversationParticipantGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PatchConversationParticipantGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversation participant gateway timeout response has a 2xx status code
func (o *PatchConversationParticipantGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversation participant gateway timeout response has a 3xx status code
func (o *PatchConversationParticipantGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversation participant gateway timeout response has a 4xx status code
func (o *PatchConversationParticipantGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch conversation participant gateway timeout response has a 5xx status code
func (o *PatchConversationParticipantGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this patch conversation participant gateway timeout response a status code equal to that given
func (o *PatchConversationParticipantGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PatchConversationParticipantGatewayTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/{conversationId}/participants/{participantId}][%d] patchConversationParticipantGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchConversationParticipantGatewayTimeout) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/{conversationId}/participants/{participantId}][%d] patchConversationParticipantGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchConversationParticipantGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationParticipantGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

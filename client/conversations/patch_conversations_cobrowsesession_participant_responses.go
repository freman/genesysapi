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

// PatchConversationsCobrowsesessionParticipantReader is a Reader for the PatchConversationsCobrowsesessionParticipant structure.
type PatchConversationsCobrowsesessionParticipantReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchConversationsCobrowsesessionParticipantReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewPatchConversationsCobrowsesessionParticipantAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchConversationsCobrowsesessionParticipantBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchConversationsCobrowsesessionParticipantUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchConversationsCobrowsesessionParticipantForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchConversationsCobrowsesessionParticipantNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPatchConversationsCobrowsesessionParticipantRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPatchConversationsCobrowsesessionParticipantRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPatchConversationsCobrowsesessionParticipantUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPatchConversationsCobrowsesessionParticipantTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchConversationsCobrowsesessionParticipantInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPatchConversationsCobrowsesessionParticipantServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPatchConversationsCobrowsesessionParticipantGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchConversationsCobrowsesessionParticipantAccepted creates a PatchConversationsCobrowsesessionParticipantAccepted with default headers values
func NewPatchConversationsCobrowsesessionParticipantAccepted() *PatchConversationsCobrowsesessionParticipantAccepted {
	return &PatchConversationsCobrowsesessionParticipantAccepted{}
}

/*
PatchConversationsCobrowsesessionParticipantAccepted describes a response with status code 202, with default header values.

Accepted
*/
type PatchConversationsCobrowsesessionParticipantAccepted struct {
}

// IsSuccess returns true when this patch conversations cobrowsesession participant accepted response has a 2xx status code
func (o *PatchConversationsCobrowsesessionParticipantAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch conversations cobrowsesession participant accepted response has a 3xx status code
func (o *PatchConversationsCobrowsesessionParticipantAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversations cobrowsesession participant accepted response has a 4xx status code
func (o *PatchConversationsCobrowsesessionParticipantAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch conversations cobrowsesession participant accepted response has a 5xx status code
func (o *PatchConversationsCobrowsesessionParticipantAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this patch conversations cobrowsesession participant accepted response a status code equal to that given
func (o *PatchConversationsCobrowsesessionParticipantAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *PatchConversationsCobrowsesessionParticipantAccepted) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/cobrowsesessions/{conversationId}/participants/{participantId}][%d] patchConversationsCobrowsesessionParticipantAccepted ", 202)
}

func (o *PatchConversationsCobrowsesessionParticipantAccepted) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/cobrowsesessions/{conversationId}/participants/{participantId}][%d] patchConversationsCobrowsesessionParticipantAccepted ", 202)
}

func (o *PatchConversationsCobrowsesessionParticipantAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchConversationsCobrowsesessionParticipantBadRequest creates a PatchConversationsCobrowsesessionParticipantBadRequest with default headers values
func NewPatchConversationsCobrowsesessionParticipantBadRequest() *PatchConversationsCobrowsesessionParticipantBadRequest {
	return &PatchConversationsCobrowsesessionParticipantBadRequest{}
}

/*
PatchConversationsCobrowsesessionParticipantBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PatchConversationsCobrowsesessionParticipantBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversations cobrowsesession participant bad request response has a 2xx status code
func (o *PatchConversationsCobrowsesessionParticipantBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversations cobrowsesession participant bad request response has a 3xx status code
func (o *PatchConversationsCobrowsesessionParticipantBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversations cobrowsesession participant bad request response has a 4xx status code
func (o *PatchConversationsCobrowsesessionParticipantBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch conversations cobrowsesession participant bad request response has a 5xx status code
func (o *PatchConversationsCobrowsesessionParticipantBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this patch conversations cobrowsesession participant bad request response a status code equal to that given
func (o *PatchConversationsCobrowsesessionParticipantBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PatchConversationsCobrowsesessionParticipantBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/cobrowsesessions/{conversationId}/participants/{participantId}][%d] patchConversationsCobrowsesessionParticipantBadRequest  %+v", 400, o.Payload)
}

func (o *PatchConversationsCobrowsesessionParticipantBadRequest) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/cobrowsesessions/{conversationId}/participants/{participantId}][%d] patchConversationsCobrowsesessionParticipantBadRequest  %+v", 400, o.Payload)
}

func (o *PatchConversationsCobrowsesessionParticipantBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsCobrowsesessionParticipantBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsCobrowsesessionParticipantUnauthorized creates a PatchConversationsCobrowsesessionParticipantUnauthorized with default headers values
func NewPatchConversationsCobrowsesessionParticipantUnauthorized() *PatchConversationsCobrowsesessionParticipantUnauthorized {
	return &PatchConversationsCobrowsesessionParticipantUnauthorized{}
}

/*
PatchConversationsCobrowsesessionParticipantUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PatchConversationsCobrowsesessionParticipantUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversations cobrowsesession participant unauthorized response has a 2xx status code
func (o *PatchConversationsCobrowsesessionParticipantUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversations cobrowsesession participant unauthorized response has a 3xx status code
func (o *PatchConversationsCobrowsesessionParticipantUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversations cobrowsesession participant unauthorized response has a 4xx status code
func (o *PatchConversationsCobrowsesessionParticipantUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch conversations cobrowsesession participant unauthorized response has a 5xx status code
func (o *PatchConversationsCobrowsesessionParticipantUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this patch conversations cobrowsesession participant unauthorized response a status code equal to that given
func (o *PatchConversationsCobrowsesessionParticipantUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PatchConversationsCobrowsesessionParticipantUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/cobrowsesessions/{conversationId}/participants/{participantId}][%d] patchConversationsCobrowsesessionParticipantUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchConversationsCobrowsesessionParticipantUnauthorized) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/cobrowsesessions/{conversationId}/participants/{participantId}][%d] patchConversationsCobrowsesessionParticipantUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchConversationsCobrowsesessionParticipantUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsCobrowsesessionParticipantUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsCobrowsesessionParticipantForbidden creates a PatchConversationsCobrowsesessionParticipantForbidden with default headers values
func NewPatchConversationsCobrowsesessionParticipantForbidden() *PatchConversationsCobrowsesessionParticipantForbidden {
	return &PatchConversationsCobrowsesessionParticipantForbidden{}
}

/*
PatchConversationsCobrowsesessionParticipantForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PatchConversationsCobrowsesessionParticipantForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversations cobrowsesession participant forbidden response has a 2xx status code
func (o *PatchConversationsCobrowsesessionParticipantForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversations cobrowsesession participant forbidden response has a 3xx status code
func (o *PatchConversationsCobrowsesessionParticipantForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversations cobrowsesession participant forbidden response has a 4xx status code
func (o *PatchConversationsCobrowsesessionParticipantForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch conversations cobrowsesession participant forbidden response has a 5xx status code
func (o *PatchConversationsCobrowsesessionParticipantForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this patch conversations cobrowsesession participant forbidden response a status code equal to that given
func (o *PatchConversationsCobrowsesessionParticipantForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PatchConversationsCobrowsesessionParticipantForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/cobrowsesessions/{conversationId}/participants/{participantId}][%d] patchConversationsCobrowsesessionParticipantForbidden  %+v", 403, o.Payload)
}

func (o *PatchConversationsCobrowsesessionParticipantForbidden) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/cobrowsesessions/{conversationId}/participants/{participantId}][%d] patchConversationsCobrowsesessionParticipantForbidden  %+v", 403, o.Payload)
}

func (o *PatchConversationsCobrowsesessionParticipantForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsCobrowsesessionParticipantForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsCobrowsesessionParticipantNotFound creates a PatchConversationsCobrowsesessionParticipantNotFound with default headers values
func NewPatchConversationsCobrowsesessionParticipantNotFound() *PatchConversationsCobrowsesessionParticipantNotFound {
	return &PatchConversationsCobrowsesessionParticipantNotFound{}
}

/*
PatchConversationsCobrowsesessionParticipantNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PatchConversationsCobrowsesessionParticipantNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversations cobrowsesession participant not found response has a 2xx status code
func (o *PatchConversationsCobrowsesessionParticipantNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversations cobrowsesession participant not found response has a 3xx status code
func (o *PatchConversationsCobrowsesessionParticipantNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversations cobrowsesession participant not found response has a 4xx status code
func (o *PatchConversationsCobrowsesessionParticipantNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch conversations cobrowsesession participant not found response has a 5xx status code
func (o *PatchConversationsCobrowsesessionParticipantNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this patch conversations cobrowsesession participant not found response a status code equal to that given
func (o *PatchConversationsCobrowsesessionParticipantNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PatchConversationsCobrowsesessionParticipantNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/cobrowsesessions/{conversationId}/participants/{participantId}][%d] patchConversationsCobrowsesessionParticipantNotFound  %+v", 404, o.Payload)
}

func (o *PatchConversationsCobrowsesessionParticipantNotFound) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/cobrowsesessions/{conversationId}/participants/{participantId}][%d] patchConversationsCobrowsesessionParticipantNotFound  %+v", 404, o.Payload)
}

func (o *PatchConversationsCobrowsesessionParticipantNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsCobrowsesessionParticipantNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsCobrowsesessionParticipantRequestTimeout creates a PatchConversationsCobrowsesessionParticipantRequestTimeout with default headers values
func NewPatchConversationsCobrowsesessionParticipantRequestTimeout() *PatchConversationsCobrowsesessionParticipantRequestTimeout {
	return &PatchConversationsCobrowsesessionParticipantRequestTimeout{}
}

/*
PatchConversationsCobrowsesessionParticipantRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PatchConversationsCobrowsesessionParticipantRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversations cobrowsesession participant request timeout response has a 2xx status code
func (o *PatchConversationsCobrowsesessionParticipantRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversations cobrowsesession participant request timeout response has a 3xx status code
func (o *PatchConversationsCobrowsesessionParticipantRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversations cobrowsesession participant request timeout response has a 4xx status code
func (o *PatchConversationsCobrowsesessionParticipantRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch conversations cobrowsesession participant request timeout response has a 5xx status code
func (o *PatchConversationsCobrowsesessionParticipantRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this patch conversations cobrowsesession participant request timeout response a status code equal to that given
func (o *PatchConversationsCobrowsesessionParticipantRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PatchConversationsCobrowsesessionParticipantRequestTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/cobrowsesessions/{conversationId}/participants/{participantId}][%d] patchConversationsCobrowsesessionParticipantRequestTimeout  %+v", 408, o.Payload)
}

func (o *PatchConversationsCobrowsesessionParticipantRequestTimeout) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/cobrowsesessions/{conversationId}/participants/{participantId}][%d] patchConversationsCobrowsesessionParticipantRequestTimeout  %+v", 408, o.Payload)
}

func (o *PatchConversationsCobrowsesessionParticipantRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsCobrowsesessionParticipantRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsCobrowsesessionParticipantRequestEntityTooLarge creates a PatchConversationsCobrowsesessionParticipantRequestEntityTooLarge with default headers values
func NewPatchConversationsCobrowsesessionParticipantRequestEntityTooLarge() *PatchConversationsCobrowsesessionParticipantRequestEntityTooLarge {
	return &PatchConversationsCobrowsesessionParticipantRequestEntityTooLarge{}
}

/*
PatchConversationsCobrowsesessionParticipantRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PatchConversationsCobrowsesessionParticipantRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversations cobrowsesession participant request entity too large response has a 2xx status code
func (o *PatchConversationsCobrowsesessionParticipantRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversations cobrowsesession participant request entity too large response has a 3xx status code
func (o *PatchConversationsCobrowsesessionParticipantRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversations cobrowsesession participant request entity too large response has a 4xx status code
func (o *PatchConversationsCobrowsesessionParticipantRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch conversations cobrowsesession participant request entity too large response has a 5xx status code
func (o *PatchConversationsCobrowsesessionParticipantRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this patch conversations cobrowsesession participant request entity too large response a status code equal to that given
func (o *PatchConversationsCobrowsesessionParticipantRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PatchConversationsCobrowsesessionParticipantRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/cobrowsesessions/{conversationId}/participants/{participantId}][%d] patchConversationsCobrowsesessionParticipantRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchConversationsCobrowsesessionParticipantRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/cobrowsesessions/{conversationId}/participants/{participantId}][%d] patchConversationsCobrowsesessionParticipantRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchConversationsCobrowsesessionParticipantRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsCobrowsesessionParticipantRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsCobrowsesessionParticipantUnsupportedMediaType creates a PatchConversationsCobrowsesessionParticipantUnsupportedMediaType with default headers values
func NewPatchConversationsCobrowsesessionParticipantUnsupportedMediaType() *PatchConversationsCobrowsesessionParticipantUnsupportedMediaType {
	return &PatchConversationsCobrowsesessionParticipantUnsupportedMediaType{}
}

/*
PatchConversationsCobrowsesessionParticipantUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PatchConversationsCobrowsesessionParticipantUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversations cobrowsesession participant unsupported media type response has a 2xx status code
func (o *PatchConversationsCobrowsesessionParticipantUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversations cobrowsesession participant unsupported media type response has a 3xx status code
func (o *PatchConversationsCobrowsesessionParticipantUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversations cobrowsesession participant unsupported media type response has a 4xx status code
func (o *PatchConversationsCobrowsesessionParticipantUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch conversations cobrowsesession participant unsupported media type response has a 5xx status code
func (o *PatchConversationsCobrowsesessionParticipantUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this patch conversations cobrowsesession participant unsupported media type response a status code equal to that given
func (o *PatchConversationsCobrowsesessionParticipantUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PatchConversationsCobrowsesessionParticipantUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/cobrowsesessions/{conversationId}/participants/{participantId}][%d] patchConversationsCobrowsesessionParticipantUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchConversationsCobrowsesessionParticipantUnsupportedMediaType) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/cobrowsesessions/{conversationId}/participants/{participantId}][%d] patchConversationsCobrowsesessionParticipantUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchConversationsCobrowsesessionParticipantUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsCobrowsesessionParticipantUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsCobrowsesessionParticipantTooManyRequests creates a PatchConversationsCobrowsesessionParticipantTooManyRequests with default headers values
func NewPatchConversationsCobrowsesessionParticipantTooManyRequests() *PatchConversationsCobrowsesessionParticipantTooManyRequests {
	return &PatchConversationsCobrowsesessionParticipantTooManyRequests{}
}

/*
PatchConversationsCobrowsesessionParticipantTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PatchConversationsCobrowsesessionParticipantTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversations cobrowsesession participant too many requests response has a 2xx status code
func (o *PatchConversationsCobrowsesessionParticipantTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversations cobrowsesession participant too many requests response has a 3xx status code
func (o *PatchConversationsCobrowsesessionParticipantTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversations cobrowsesession participant too many requests response has a 4xx status code
func (o *PatchConversationsCobrowsesessionParticipantTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch conversations cobrowsesession participant too many requests response has a 5xx status code
func (o *PatchConversationsCobrowsesessionParticipantTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this patch conversations cobrowsesession participant too many requests response a status code equal to that given
func (o *PatchConversationsCobrowsesessionParticipantTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PatchConversationsCobrowsesessionParticipantTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/cobrowsesessions/{conversationId}/participants/{participantId}][%d] patchConversationsCobrowsesessionParticipantTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchConversationsCobrowsesessionParticipantTooManyRequests) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/cobrowsesessions/{conversationId}/participants/{participantId}][%d] patchConversationsCobrowsesessionParticipantTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchConversationsCobrowsesessionParticipantTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsCobrowsesessionParticipantTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsCobrowsesessionParticipantInternalServerError creates a PatchConversationsCobrowsesessionParticipantInternalServerError with default headers values
func NewPatchConversationsCobrowsesessionParticipantInternalServerError() *PatchConversationsCobrowsesessionParticipantInternalServerError {
	return &PatchConversationsCobrowsesessionParticipantInternalServerError{}
}

/*
PatchConversationsCobrowsesessionParticipantInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PatchConversationsCobrowsesessionParticipantInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversations cobrowsesession participant internal server error response has a 2xx status code
func (o *PatchConversationsCobrowsesessionParticipantInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversations cobrowsesession participant internal server error response has a 3xx status code
func (o *PatchConversationsCobrowsesessionParticipantInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversations cobrowsesession participant internal server error response has a 4xx status code
func (o *PatchConversationsCobrowsesessionParticipantInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch conversations cobrowsesession participant internal server error response has a 5xx status code
func (o *PatchConversationsCobrowsesessionParticipantInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this patch conversations cobrowsesession participant internal server error response a status code equal to that given
func (o *PatchConversationsCobrowsesessionParticipantInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PatchConversationsCobrowsesessionParticipantInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/cobrowsesessions/{conversationId}/participants/{participantId}][%d] patchConversationsCobrowsesessionParticipantInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchConversationsCobrowsesessionParticipantInternalServerError) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/cobrowsesessions/{conversationId}/participants/{participantId}][%d] patchConversationsCobrowsesessionParticipantInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchConversationsCobrowsesessionParticipantInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsCobrowsesessionParticipantInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsCobrowsesessionParticipantServiceUnavailable creates a PatchConversationsCobrowsesessionParticipantServiceUnavailable with default headers values
func NewPatchConversationsCobrowsesessionParticipantServiceUnavailable() *PatchConversationsCobrowsesessionParticipantServiceUnavailable {
	return &PatchConversationsCobrowsesessionParticipantServiceUnavailable{}
}

/*
PatchConversationsCobrowsesessionParticipantServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PatchConversationsCobrowsesessionParticipantServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversations cobrowsesession participant service unavailable response has a 2xx status code
func (o *PatchConversationsCobrowsesessionParticipantServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversations cobrowsesession participant service unavailable response has a 3xx status code
func (o *PatchConversationsCobrowsesessionParticipantServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversations cobrowsesession participant service unavailable response has a 4xx status code
func (o *PatchConversationsCobrowsesessionParticipantServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch conversations cobrowsesession participant service unavailable response has a 5xx status code
func (o *PatchConversationsCobrowsesessionParticipantServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this patch conversations cobrowsesession participant service unavailable response a status code equal to that given
func (o *PatchConversationsCobrowsesessionParticipantServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PatchConversationsCobrowsesessionParticipantServiceUnavailable) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/cobrowsesessions/{conversationId}/participants/{participantId}][%d] patchConversationsCobrowsesessionParticipantServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchConversationsCobrowsesessionParticipantServiceUnavailable) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/cobrowsesessions/{conversationId}/participants/{participantId}][%d] patchConversationsCobrowsesessionParticipantServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchConversationsCobrowsesessionParticipantServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsCobrowsesessionParticipantServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsCobrowsesessionParticipantGatewayTimeout creates a PatchConversationsCobrowsesessionParticipantGatewayTimeout with default headers values
func NewPatchConversationsCobrowsesessionParticipantGatewayTimeout() *PatchConversationsCobrowsesessionParticipantGatewayTimeout {
	return &PatchConversationsCobrowsesessionParticipantGatewayTimeout{}
}

/*
PatchConversationsCobrowsesessionParticipantGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PatchConversationsCobrowsesessionParticipantGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversations cobrowsesession participant gateway timeout response has a 2xx status code
func (o *PatchConversationsCobrowsesessionParticipantGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversations cobrowsesession participant gateway timeout response has a 3xx status code
func (o *PatchConversationsCobrowsesessionParticipantGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversations cobrowsesession participant gateway timeout response has a 4xx status code
func (o *PatchConversationsCobrowsesessionParticipantGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch conversations cobrowsesession participant gateway timeout response has a 5xx status code
func (o *PatchConversationsCobrowsesessionParticipantGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this patch conversations cobrowsesession participant gateway timeout response a status code equal to that given
func (o *PatchConversationsCobrowsesessionParticipantGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PatchConversationsCobrowsesessionParticipantGatewayTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/cobrowsesessions/{conversationId}/participants/{participantId}][%d] patchConversationsCobrowsesessionParticipantGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchConversationsCobrowsesessionParticipantGatewayTimeout) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/cobrowsesessions/{conversationId}/participants/{participantId}][%d] patchConversationsCobrowsesessionParticipantGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchConversationsCobrowsesessionParticipantGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsCobrowsesessionParticipantGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

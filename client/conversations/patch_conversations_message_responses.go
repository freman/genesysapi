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

// PatchConversationsMessageReader is a Reader for the PatchConversationsMessage structure.
type PatchConversationsMessageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchConversationsMessageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchConversationsMessageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchConversationsMessageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchConversationsMessageUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchConversationsMessageForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchConversationsMessageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPatchConversationsMessageRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPatchConversationsMessageRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPatchConversationsMessageUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPatchConversationsMessageTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchConversationsMessageInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPatchConversationsMessageServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPatchConversationsMessageGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchConversationsMessageOK creates a PatchConversationsMessageOK with default headers values
func NewPatchConversationsMessageOK() *PatchConversationsMessageOK {
	return &PatchConversationsMessageOK{}
}

/*
PatchConversationsMessageOK describes a response with status code 200, with default header values.

successful operation
*/
type PatchConversationsMessageOK struct {
	Payload *models.Conversation
}

// IsSuccess returns true when this patch conversations message o k response has a 2xx status code
func (o *PatchConversationsMessageOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch conversations message o k response has a 3xx status code
func (o *PatchConversationsMessageOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversations message o k response has a 4xx status code
func (o *PatchConversationsMessageOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch conversations message o k response has a 5xx status code
func (o *PatchConversationsMessageOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch conversations message o k response a status code equal to that given
func (o *PatchConversationsMessageOK) IsCode(code int) bool {
	return code == 200
}

func (o *PatchConversationsMessageOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messages/{conversationId}][%d] patchConversationsMessageOK  %+v", 200, o.Payload)
}

func (o *PatchConversationsMessageOK) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messages/{conversationId}][%d] patchConversationsMessageOK  %+v", 200, o.Payload)
}

func (o *PatchConversationsMessageOK) GetPayload() *models.Conversation {
	return o.Payload
}

func (o *PatchConversationsMessageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Conversation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsMessageBadRequest creates a PatchConversationsMessageBadRequest with default headers values
func NewPatchConversationsMessageBadRequest() *PatchConversationsMessageBadRequest {
	return &PatchConversationsMessageBadRequest{}
}

/*
PatchConversationsMessageBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PatchConversationsMessageBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversations message bad request response has a 2xx status code
func (o *PatchConversationsMessageBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversations message bad request response has a 3xx status code
func (o *PatchConversationsMessageBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversations message bad request response has a 4xx status code
func (o *PatchConversationsMessageBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch conversations message bad request response has a 5xx status code
func (o *PatchConversationsMessageBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this patch conversations message bad request response a status code equal to that given
func (o *PatchConversationsMessageBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PatchConversationsMessageBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messages/{conversationId}][%d] patchConversationsMessageBadRequest  %+v", 400, o.Payload)
}

func (o *PatchConversationsMessageBadRequest) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messages/{conversationId}][%d] patchConversationsMessageBadRequest  %+v", 400, o.Payload)
}

func (o *PatchConversationsMessageBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsMessageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsMessageUnauthorized creates a PatchConversationsMessageUnauthorized with default headers values
func NewPatchConversationsMessageUnauthorized() *PatchConversationsMessageUnauthorized {
	return &PatchConversationsMessageUnauthorized{}
}

/*
PatchConversationsMessageUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PatchConversationsMessageUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversations message unauthorized response has a 2xx status code
func (o *PatchConversationsMessageUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversations message unauthorized response has a 3xx status code
func (o *PatchConversationsMessageUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversations message unauthorized response has a 4xx status code
func (o *PatchConversationsMessageUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch conversations message unauthorized response has a 5xx status code
func (o *PatchConversationsMessageUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this patch conversations message unauthorized response a status code equal to that given
func (o *PatchConversationsMessageUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PatchConversationsMessageUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messages/{conversationId}][%d] patchConversationsMessageUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchConversationsMessageUnauthorized) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messages/{conversationId}][%d] patchConversationsMessageUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchConversationsMessageUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsMessageUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsMessageForbidden creates a PatchConversationsMessageForbidden with default headers values
func NewPatchConversationsMessageForbidden() *PatchConversationsMessageForbidden {
	return &PatchConversationsMessageForbidden{}
}

/*
PatchConversationsMessageForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PatchConversationsMessageForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversations message forbidden response has a 2xx status code
func (o *PatchConversationsMessageForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversations message forbidden response has a 3xx status code
func (o *PatchConversationsMessageForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversations message forbidden response has a 4xx status code
func (o *PatchConversationsMessageForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch conversations message forbidden response has a 5xx status code
func (o *PatchConversationsMessageForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this patch conversations message forbidden response a status code equal to that given
func (o *PatchConversationsMessageForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PatchConversationsMessageForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messages/{conversationId}][%d] patchConversationsMessageForbidden  %+v", 403, o.Payload)
}

func (o *PatchConversationsMessageForbidden) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messages/{conversationId}][%d] patchConversationsMessageForbidden  %+v", 403, o.Payload)
}

func (o *PatchConversationsMessageForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsMessageForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsMessageNotFound creates a PatchConversationsMessageNotFound with default headers values
func NewPatchConversationsMessageNotFound() *PatchConversationsMessageNotFound {
	return &PatchConversationsMessageNotFound{}
}

/*
PatchConversationsMessageNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PatchConversationsMessageNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversations message not found response has a 2xx status code
func (o *PatchConversationsMessageNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversations message not found response has a 3xx status code
func (o *PatchConversationsMessageNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversations message not found response has a 4xx status code
func (o *PatchConversationsMessageNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch conversations message not found response has a 5xx status code
func (o *PatchConversationsMessageNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this patch conversations message not found response a status code equal to that given
func (o *PatchConversationsMessageNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PatchConversationsMessageNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messages/{conversationId}][%d] patchConversationsMessageNotFound  %+v", 404, o.Payload)
}

func (o *PatchConversationsMessageNotFound) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messages/{conversationId}][%d] patchConversationsMessageNotFound  %+v", 404, o.Payload)
}

func (o *PatchConversationsMessageNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsMessageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsMessageRequestTimeout creates a PatchConversationsMessageRequestTimeout with default headers values
func NewPatchConversationsMessageRequestTimeout() *PatchConversationsMessageRequestTimeout {
	return &PatchConversationsMessageRequestTimeout{}
}

/*
PatchConversationsMessageRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PatchConversationsMessageRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversations message request timeout response has a 2xx status code
func (o *PatchConversationsMessageRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversations message request timeout response has a 3xx status code
func (o *PatchConversationsMessageRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversations message request timeout response has a 4xx status code
func (o *PatchConversationsMessageRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch conversations message request timeout response has a 5xx status code
func (o *PatchConversationsMessageRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this patch conversations message request timeout response a status code equal to that given
func (o *PatchConversationsMessageRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PatchConversationsMessageRequestTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messages/{conversationId}][%d] patchConversationsMessageRequestTimeout  %+v", 408, o.Payload)
}

func (o *PatchConversationsMessageRequestTimeout) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messages/{conversationId}][%d] patchConversationsMessageRequestTimeout  %+v", 408, o.Payload)
}

func (o *PatchConversationsMessageRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsMessageRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsMessageRequestEntityTooLarge creates a PatchConversationsMessageRequestEntityTooLarge with default headers values
func NewPatchConversationsMessageRequestEntityTooLarge() *PatchConversationsMessageRequestEntityTooLarge {
	return &PatchConversationsMessageRequestEntityTooLarge{}
}

/*
PatchConversationsMessageRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PatchConversationsMessageRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversations message request entity too large response has a 2xx status code
func (o *PatchConversationsMessageRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversations message request entity too large response has a 3xx status code
func (o *PatchConversationsMessageRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversations message request entity too large response has a 4xx status code
func (o *PatchConversationsMessageRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch conversations message request entity too large response has a 5xx status code
func (o *PatchConversationsMessageRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this patch conversations message request entity too large response a status code equal to that given
func (o *PatchConversationsMessageRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PatchConversationsMessageRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messages/{conversationId}][%d] patchConversationsMessageRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchConversationsMessageRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messages/{conversationId}][%d] patchConversationsMessageRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchConversationsMessageRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsMessageRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsMessageUnsupportedMediaType creates a PatchConversationsMessageUnsupportedMediaType with default headers values
func NewPatchConversationsMessageUnsupportedMediaType() *PatchConversationsMessageUnsupportedMediaType {
	return &PatchConversationsMessageUnsupportedMediaType{}
}

/*
PatchConversationsMessageUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PatchConversationsMessageUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversations message unsupported media type response has a 2xx status code
func (o *PatchConversationsMessageUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversations message unsupported media type response has a 3xx status code
func (o *PatchConversationsMessageUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversations message unsupported media type response has a 4xx status code
func (o *PatchConversationsMessageUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch conversations message unsupported media type response has a 5xx status code
func (o *PatchConversationsMessageUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this patch conversations message unsupported media type response a status code equal to that given
func (o *PatchConversationsMessageUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PatchConversationsMessageUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messages/{conversationId}][%d] patchConversationsMessageUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchConversationsMessageUnsupportedMediaType) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messages/{conversationId}][%d] patchConversationsMessageUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchConversationsMessageUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsMessageUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsMessageTooManyRequests creates a PatchConversationsMessageTooManyRequests with default headers values
func NewPatchConversationsMessageTooManyRequests() *PatchConversationsMessageTooManyRequests {
	return &PatchConversationsMessageTooManyRequests{}
}

/*
PatchConversationsMessageTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PatchConversationsMessageTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversations message too many requests response has a 2xx status code
func (o *PatchConversationsMessageTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversations message too many requests response has a 3xx status code
func (o *PatchConversationsMessageTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversations message too many requests response has a 4xx status code
func (o *PatchConversationsMessageTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch conversations message too many requests response has a 5xx status code
func (o *PatchConversationsMessageTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this patch conversations message too many requests response a status code equal to that given
func (o *PatchConversationsMessageTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PatchConversationsMessageTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messages/{conversationId}][%d] patchConversationsMessageTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchConversationsMessageTooManyRequests) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messages/{conversationId}][%d] patchConversationsMessageTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchConversationsMessageTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsMessageTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsMessageInternalServerError creates a PatchConversationsMessageInternalServerError with default headers values
func NewPatchConversationsMessageInternalServerError() *PatchConversationsMessageInternalServerError {
	return &PatchConversationsMessageInternalServerError{}
}

/*
PatchConversationsMessageInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PatchConversationsMessageInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversations message internal server error response has a 2xx status code
func (o *PatchConversationsMessageInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversations message internal server error response has a 3xx status code
func (o *PatchConversationsMessageInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversations message internal server error response has a 4xx status code
func (o *PatchConversationsMessageInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch conversations message internal server error response has a 5xx status code
func (o *PatchConversationsMessageInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this patch conversations message internal server error response a status code equal to that given
func (o *PatchConversationsMessageInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PatchConversationsMessageInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messages/{conversationId}][%d] patchConversationsMessageInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchConversationsMessageInternalServerError) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messages/{conversationId}][%d] patchConversationsMessageInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchConversationsMessageInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsMessageInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsMessageServiceUnavailable creates a PatchConversationsMessageServiceUnavailable with default headers values
func NewPatchConversationsMessageServiceUnavailable() *PatchConversationsMessageServiceUnavailable {
	return &PatchConversationsMessageServiceUnavailable{}
}

/*
PatchConversationsMessageServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PatchConversationsMessageServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversations message service unavailable response has a 2xx status code
func (o *PatchConversationsMessageServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversations message service unavailable response has a 3xx status code
func (o *PatchConversationsMessageServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversations message service unavailable response has a 4xx status code
func (o *PatchConversationsMessageServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch conversations message service unavailable response has a 5xx status code
func (o *PatchConversationsMessageServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this patch conversations message service unavailable response a status code equal to that given
func (o *PatchConversationsMessageServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PatchConversationsMessageServiceUnavailable) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messages/{conversationId}][%d] patchConversationsMessageServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchConversationsMessageServiceUnavailable) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messages/{conversationId}][%d] patchConversationsMessageServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchConversationsMessageServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsMessageServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsMessageGatewayTimeout creates a PatchConversationsMessageGatewayTimeout with default headers values
func NewPatchConversationsMessageGatewayTimeout() *PatchConversationsMessageGatewayTimeout {
	return &PatchConversationsMessageGatewayTimeout{}
}

/*
PatchConversationsMessageGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PatchConversationsMessageGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversations message gateway timeout response has a 2xx status code
func (o *PatchConversationsMessageGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversations message gateway timeout response has a 3xx status code
func (o *PatchConversationsMessageGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversations message gateway timeout response has a 4xx status code
func (o *PatchConversationsMessageGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch conversations message gateway timeout response has a 5xx status code
func (o *PatchConversationsMessageGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this patch conversations message gateway timeout response a status code equal to that given
func (o *PatchConversationsMessageGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PatchConversationsMessageGatewayTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messages/{conversationId}][%d] patchConversationsMessageGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchConversationsMessageGatewayTimeout) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messages/{conversationId}][%d] patchConversationsMessageGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchConversationsMessageGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsMessageGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

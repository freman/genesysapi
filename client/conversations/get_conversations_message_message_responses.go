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

// GetConversationsMessageMessageReader is a Reader for the GetConversationsMessageMessage structure.
type GetConversationsMessageMessageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConversationsMessageMessageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetConversationsMessageMessageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetConversationsMessageMessageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetConversationsMessageMessageUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetConversationsMessageMessageForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetConversationsMessageMessageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetConversationsMessageMessageRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetConversationsMessageMessageRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetConversationsMessageMessageUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetConversationsMessageMessageTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetConversationsMessageMessageInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetConversationsMessageMessageServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetConversationsMessageMessageGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetConversationsMessageMessageOK creates a GetConversationsMessageMessageOK with default headers values
func NewGetConversationsMessageMessageOK() *GetConversationsMessageMessageOK {
	return &GetConversationsMessageMessageOK{}
}

/*
GetConversationsMessageMessageOK describes a response with status code 200, with default header values.

successful operation
*/
type GetConversationsMessageMessageOK struct {
	Payload *models.MessageData
}

// IsSuccess returns true when this get conversations message message o k response has a 2xx status code
func (o *GetConversationsMessageMessageOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get conversations message message o k response has a 3xx status code
func (o *GetConversationsMessageMessageOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations message message o k response has a 4xx status code
func (o *GetConversationsMessageMessageOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversations message message o k response has a 5xx status code
func (o *GetConversationsMessageMessageOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations message message o k response a status code equal to that given
func (o *GetConversationsMessageMessageOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetConversationsMessageMessageOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages/{conversationId}/messages/{messageId}][%d] getConversationsMessageMessageOK  %+v", 200, o.Payload)
}

func (o *GetConversationsMessageMessageOK) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages/{conversationId}/messages/{messageId}][%d] getConversationsMessageMessageOK  %+v", 200, o.Payload)
}

func (o *GetConversationsMessageMessageOK) GetPayload() *models.MessageData {
	return o.Payload
}

func (o *GetConversationsMessageMessageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MessageData)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessageMessageBadRequest creates a GetConversationsMessageMessageBadRequest with default headers values
func NewGetConversationsMessageMessageBadRequest() *GetConversationsMessageMessageBadRequest {
	return &GetConversationsMessageMessageBadRequest{}
}

/*
GetConversationsMessageMessageBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetConversationsMessageMessageBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations message message bad request response has a 2xx status code
func (o *GetConversationsMessageMessageBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations message message bad request response has a 3xx status code
func (o *GetConversationsMessageMessageBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations message message bad request response has a 4xx status code
func (o *GetConversationsMessageMessageBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations message message bad request response has a 5xx status code
func (o *GetConversationsMessageMessageBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations message message bad request response a status code equal to that given
func (o *GetConversationsMessageMessageBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetConversationsMessageMessageBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages/{conversationId}/messages/{messageId}][%d] getConversationsMessageMessageBadRequest  %+v", 400, o.Payload)
}

func (o *GetConversationsMessageMessageBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages/{conversationId}/messages/{messageId}][%d] getConversationsMessageMessageBadRequest  %+v", 400, o.Payload)
}

func (o *GetConversationsMessageMessageBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessageMessageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessageMessageUnauthorized creates a GetConversationsMessageMessageUnauthorized with default headers values
func NewGetConversationsMessageMessageUnauthorized() *GetConversationsMessageMessageUnauthorized {
	return &GetConversationsMessageMessageUnauthorized{}
}

/*
GetConversationsMessageMessageUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetConversationsMessageMessageUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations message message unauthorized response has a 2xx status code
func (o *GetConversationsMessageMessageUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations message message unauthorized response has a 3xx status code
func (o *GetConversationsMessageMessageUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations message message unauthorized response has a 4xx status code
func (o *GetConversationsMessageMessageUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations message message unauthorized response has a 5xx status code
func (o *GetConversationsMessageMessageUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations message message unauthorized response a status code equal to that given
func (o *GetConversationsMessageMessageUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetConversationsMessageMessageUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages/{conversationId}/messages/{messageId}][%d] getConversationsMessageMessageUnauthorized  %+v", 401, o.Payload)
}

func (o *GetConversationsMessageMessageUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages/{conversationId}/messages/{messageId}][%d] getConversationsMessageMessageUnauthorized  %+v", 401, o.Payload)
}

func (o *GetConversationsMessageMessageUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessageMessageUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessageMessageForbidden creates a GetConversationsMessageMessageForbidden with default headers values
func NewGetConversationsMessageMessageForbidden() *GetConversationsMessageMessageForbidden {
	return &GetConversationsMessageMessageForbidden{}
}

/*
GetConversationsMessageMessageForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetConversationsMessageMessageForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations message message forbidden response has a 2xx status code
func (o *GetConversationsMessageMessageForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations message message forbidden response has a 3xx status code
func (o *GetConversationsMessageMessageForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations message message forbidden response has a 4xx status code
func (o *GetConversationsMessageMessageForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations message message forbidden response has a 5xx status code
func (o *GetConversationsMessageMessageForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations message message forbidden response a status code equal to that given
func (o *GetConversationsMessageMessageForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetConversationsMessageMessageForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages/{conversationId}/messages/{messageId}][%d] getConversationsMessageMessageForbidden  %+v", 403, o.Payload)
}

func (o *GetConversationsMessageMessageForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages/{conversationId}/messages/{messageId}][%d] getConversationsMessageMessageForbidden  %+v", 403, o.Payload)
}

func (o *GetConversationsMessageMessageForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessageMessageForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessageMessageNotFound creates a GetConversationsMessageMessageNotFound with default headers values
func NewGetConversationsMessageMessageNotFound() *GetConversationsMessageMessageNotFound {
	return &GetConversationsMessageMessageNotFound{}
}

/*
GetConversationsMessageMessageNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetConversationsMessageMessageNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations message message not found response has a 2xx status code
func (o *GetConversationsMessageMessageNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations message message not found response has a 3xx status code
func (o *GetConversationsMessageMessageNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations message message not found response has a 4xx status code
func (o *GetConversationsMessageMessageNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations message message not found response has a 5xx status code
func (o *GetConversationsMessageMessageNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations message message not found response a status code equal to that given
func (o *GetConversationsMessageMessageNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetConversationsMessageMessageNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages/{conversationId}/messages/{messageId}][%d] getConversationsMessageMessageNotFound  %+v", 404, o.Payload)
}

func (o *GetConversationsMessageMessageNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages/{conversationId}/messages/{messageId}][%d] getConversationsMessageMessageNotFound  %+v", 404, o.Payload)
}

func (o *GetConversationsMessageMessageNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessageMessageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessageMessageRequestTimeout creates a GetConversationsMessageMessageRequestTimeout with default headers values
func NewGetConversationsMessageMessageRequestTimeout() *GetConversationsMessageMessageRequestTimeout {
	return &GetConversationsMessageMessageRequestTimeout{}
}

/*
GetConversationsMessageMessageRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetConversationsMessageMessageRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations message message request timeout response has a 2xx status code
func (o *GetConversationsMessageMessageRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations message message request timeout response has a 3xx status code
func (o *GetConversationsMessageMessageRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations message message request timeout response has a 4xx status code
func (o *GetConversationsMessageMessageRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations message message request timeout response has a 5xx status code
func (o *GetConversationsMessageMessageRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations message message request timeout response a status code equal to that given
func (o *GetConversationsMessageMessageRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetConversationsMessageMessageRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages/{conversationId}/messages/{messageId}][%d] getConversationsMessageMessageRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetConversationsMessageMessageRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages/{conversationId}/messages/{messageId}][%d] getConversationsMessageMessageRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetConversationsMessageMessageRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessageMessageRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessageMessageRequestEntityTooLarge creates a GetConversationsMessageMessageRequestEntityTooLarge with default headers values
func NewGetConversationsMessageMessageRequestEntityTooLarge() *GetConversationsMessageMessageRequestEntityTooLarge {
	return &GetConversationsMessageMessageRequestEntityTooLarge{}
}

/*
GetConversationsMessageMessageRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetConversationsMessageMessageRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations message message request entity too large response has a 2xx status code
func (o *GetConversationsMessageMessageRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations message message request entity too large response has a 3xx status code
func (o *GetConversationsMessageMessageRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations message message request entity too large response has a 4xx status code
func (o *GetConversationsMessageMessageRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations message message request entity too large response has a 5xx status code
func (o *GetConversationsMessageMessageRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations message message request entity too large response a status code equal to that given
func (o *GetConversationsMessageMessageRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetConversationsMessageMessageRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages/{conversationId}/messages/{messageId}][%d] getConversationsMessageMessageRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetConversationsMessageMessageRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages/{conversationId}/messages/{messageId}][%d] getConversationsMessageMessageRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetConversationsMessageMessageRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessageMessageRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessageMessageUnsupportedMediaType creates a GetConversationsMessageMessageUnsupportedMediaType with default headers values
func NewGetConversationsMessageMessageUnsupportedMediaType() *GetConversationsMessageMessageUnsupportedMediaType {
	return &GetConversationsMessageMessageUnsupportedMediaType{}
}

/*
GetConversationsMessageMessageUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetConversationsMessageMessageUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations message message unsupported media type response has a 2xx status code
func (o *GetConversationsMessageMessageUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations message message unsupported media type response has a 3xx status code
func (o *GetConversationsMessageMessageUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations message message unsupported media type response has a 4xx status code
func (o *GetConversationsMessageMessageUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations message message unsupported media type response has a 5xx status code
func (o *GetConversationsMessageMessageUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations message message unsupported media type response a status code equal to that given
func (o *GetConversationsMessageMessageUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetConversationsMessageMessageUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages/{conversationId}/messages/{messageId}][%d] getConversationsMessageMessageUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetConversationsMessageMessageUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages/{conversationId}/messages/{messageId}][%d] getConversationsMessageMessageUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetConversationsMessageMessageUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessageMessageUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessageMessageTooManyRequests creates a GetConversationsMessageMessageTooManyRequests with default headers values
func NewGetConversationsMessageMessageTooManyRequests() *GetConversationsMessageMessageTooManyRequests {
	return &GetConversationsMessageMessageTooManyRequests{}
}

/*
GetConversationsMessageMessageTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetConversationsMessageMessageTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations message message too many requests response has a 2xx status code
func (o *GetConversationsMessageMessageTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations message message too many requests response has a 3xx status code
func (o *GetConversationsMessageMessageTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations message message too many requests response has a 4xx status code
func (o *GetConversationsMessageMessageTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations message message too many requests response has a 5xx status code
func (o *GetConversationsMessageMessageTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations message message too many requests response a status code equal to that given
func (o *GetConversationsMessageMessageTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetConversationsMessageMessageTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages/{conversationId}/messages/{messageId}][%d] getConversationsMessageMessageTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetConversationsMessageMessageTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages/{conversationId}/messages/{messageId}][%d] getConversationsMessageMessageTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetConversationsMessageMessageTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessageMessageTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessageMessageInternalServerError creates a GetConversationsMessageMessageInternalServerError with default headers values
func NewGetConversationsMessageMessageInternalServerError() *GetConversationsMessageMessageInternalServerError {
	return &GetConversationsMessageMessageInternalServerError{}
}

/*
GetConversationsMessageMessageInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetConversationsMessageMessageInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations message message internal server error response has a 2xx status code
func (o *GetConversationsMessageMessageInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations message message internal server error response has a 3xx status code
func (o *GetConversationsMessageMessageInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations message message internal server error response has a 4xx status code
func (o *GetConversationsMessageMessageInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversations message message internal server error response has a 5xx status code
func (o *GetConversationsMessageMessageInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get conversations message message internal server error response a status code equal to that given
func (o *GetConversationsMessageMessageInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetConversationsMessageMessageInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages/{conversationId}/messages/{messageId}][%d] getConversationsMessageMessageInternalServerError  %+v", 500, o.Payload)
}

func (o *GetConversationsMessageMessageInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages/{conversationId}/messages/{messageId}][%d] getConversationsMessageMessageInternalServerError  %+v", 500, o.Payload)
}

func (o *GetConversationsMessageMessageInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessageMessageInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessageMessageServiceUnavailable creates a GetConversationsMessageMessageServiceUnavailable with default headers values
func NewGetConversationsMessageMessageServiceUnavailable() *GetConversationsMessageMessageServiceUnavailable {
	return &GetConversationsMessageMessageServiceUnavailable{}
}

/*
GetConversationsMessageMessageServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetConversationsMessageMessageServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations message message service unavailable response has a 2xx status code
func (o *GetConversationsMessageMessageServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations message message service unavailable response has a 3xx status code
func (o *GetConversationsMessageMessageServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations message message service unavailable response has a 4xx status code
func (o *GetConversationsMessageMessageServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversations message message service unavailable response has a 5xx status code
func (o *GetConversationsMessageMessageServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get conversations message message service unavailable response a status code equal to that given
func (o *GetConversationsMessageMessageServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetConversationsMessageMessageServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages/{conversationId}/messages/{messageId}][%d] getConversationsMessageMessageServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetConversationsMessageMessageServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages/{conversationId}/messages/{messageId}][%d] getConversationsMessageMessageServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetConversationsMessageMessageServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessageMessageServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessageMessageGatewayTimeout creates a GetConversationsMessageMessageGatewayTimeout with default headers values
func NewGetConversationsMessageMessageGatewayTimeout() *GetConversationsMessageMessageGatewayTimeout {
	return &GetConversationsMessageMessageGatewayTimeout{}
}

/*
GetConversationsMessageMessageGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetConversationsMessageMessageGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations message message gateway timeout response has a 2xx status code
func (o *GetConversationsMessageMessageGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations message message gateway timeout response has a 3xx status code
func (o *GetConversationsMessageMessageGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations message message gateway timeout response has a 4xx status code
func (o *GetConversationsMessageMessageGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversations message message gateway timeout response has a 5xx status code
func (o *GetConversationsMessageMessageGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get conversations message message gateway timeout response a status code equal to that given
func (o *GetConversationsMessageMessageGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetConversationsMessageMessageGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages/{conversationId}/messages/{messageId}][%d] getConversationsMessageMessageGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetConversationsMessageMessageGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages/{conversationId}/messages/{messageId}][%d] getConversationsMessageMessageGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetConversationsMessageMessageGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessageMessageGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

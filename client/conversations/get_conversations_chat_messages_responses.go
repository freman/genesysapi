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

// GetConversationsChatMessagesReader is a Reader for the GetConversationsChatMessages structure.
type GetConversationsChatMessagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConversationsChatMessagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetConversationsChatMessagesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetConversationsChatMessagesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetConversationsChatMessagesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetConversationsChatMessagesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetConversationsChatMessagesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetConversationsChatMessagesRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetConversationsChatMessagesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetConversationsChatMessagesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetConversationsChatMessagesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetConversationsChatMessagesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetConversationsChatMessagesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetConversationsChatMessagesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetConversationsChatMessagesOK creates a GetConversationsChatMessagesOK with default headers values
func NewGetConversationsChatMessagesOK() *GetConversationsChatMessagesOK {
	return &GetConversationsChatMessagesOK{}
}

/*
GetConversationsChatMessagesOK describes a response with status code 200, with default header values.

successful operation
*/
type GetConversationsChatMessagesOK struct {
	Payload *models.WebChatMessageEntityList
}

// IsSuccess returns true when this get conversations chat messages o k response has a 2xx status code
func (o *GetConversationsChatMessagesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get conversations chat messages o k response has a 3xx status code
func (o *GetConversationsChatMessagesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations chat messages o k response has a 4xx status code
func (o *GetConversationsChatMessagesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversations chat messages o k response has a 5xx status code
func (o *GetConversationsChatMessagesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations chat messages o k response a status code equal to that given
func (o *GetConversationsChatMessagesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetConversationsChatMessagesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}/messages][%d] getConversationsChatMessagesOK  %+v", 200, o.Payload)
}

func (o *GetConversationsChatMessagesOK) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}/messages][%d] getConversationsChatMessagesOK  %+v", 200, o.Payload)
}

func (o *GetConversationsChatMessagesOK) GetPayload() *models.WebChatMessageEntityList {
	return o.Payload
}

func (o *GetConversationsChatMessagesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WebChatMessageEntityList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsChatMessagesBadRequest creates a GetConversationsChatMessagesBadRequest with default headers values
func NewGetConversationsChatMessagesBadRequest() *GetConversationsChatMessagesBadRequest {
	return &GetConversationsChatMessagesBadRequest{}
}

/*
GetConversationsChatMessagesBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetConversationsChatMessagesBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations chat messages bad request response has a 2xx status code
func (o *GetConversationsChatMessagesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations chat messages bad request response has a 3xx status code
func (o *GetConversationsChatMessagesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations chat messages bad request response has a 4xx status code
func (o *GetConversationsChatMessagesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations chat messages bad request response has a 5xx status code
func (o *GetConversationsChatMessagesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations chat messages bad request response a status code equal to that given
func (o *GetConversationsChatMessagesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetConversationsChatMessagesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}/messages][%d] getConversationsChatMessagesBadRequest  %+v", 400, o.Payload)
}

func (o *GetConversationsChatMessagesBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}/messages][%d] getConversationsChatMessagesBadRequest  %+v", 400, o.Payload)
}

func (o *GetConversationsChatMessagesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsChatMessagesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsChatMessagesUnauthorized creates a GetConversationsChatMessagesUnauthorized with default headers values
func NewGetConversationsChatMessagesUnauthorized() *GetConversationsChatMessagesUnauthorized {
	return &GetConversationsChatMessagesUnauthorized{}
}

/*
GetConversationsChatMessagesUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetConversationsChatMessagesUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations chat messages unauthorized response has a 2xx status code
func (o *GetConversationsChatMessagesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations chat messages unauthorized response has a 3xx status code
func (o *GetConversationsChatMessagesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations chat messages unauthorized response has a 4xx status code
func (o *GetConversationsChatMessagesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations chat messages unauthorized response has a 5xx status code
func (o *GetConversationsChatMessagesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations chat messages unauthorized response a status code equal to that given
func (o *GetConversationsChatMessagesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetConversationsChatMessagesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}/messages][%d] getConversationsChatMessagesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetConversationsChatMessagesUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}/messages][%d] getConversationsChatMessagesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetConversationsChatMessagesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsChatMessagesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsChatMessagesForbidden creates a GetConversationsChatMessagesForbidden with default headers values
func NewGetConversationsChatMessagesForbidden() *GetConversationsChatMessagesForbidden {
	return &GetConversationsChatMessagesForbidden{}
}

/*
GetConversationsChatMessagesForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetConversationsChatMessagesForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations chat messages forbidden response has a 2xx status code
func (o *GetConversationsChatMessagesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations chat messages forbidden response has a 3xx status code
func (o *GetConversationsChatMessagesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations chat messages forbidden response has a 4xx status code
func (o *GetConversationsChatMessagesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations chat messages forbidden response has a 5xx status code
func (o *GetConversationsChatMessagesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations chat messages forbidden response a status code equal to that given
func (o *GetConversationsChatMessagesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetConversationsChatMessagesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}/messages][%d] getConversationsChatMessagesForbidden  %+v", 403, o.Payload)
}

func (o *GetConversationsChatMessagesForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}/messages][%d] getConversationsChatMessagesForbidden  %+v", 403, o.Payload)
}

func (o *GetConversationsChatMessagesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsChatMessagesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsChatMessagesNotFound creates a GetConversationsChatMessagesNotFound with default headers values
func NewGetConversationsChatMessagesNotFound() *GetConversationsChatMessagesNotFound {
	return &GetConversationsChatMessagesNotFound{}
}

/*
GetConversationsChatMessagesNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetConversationsChatMessagesNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations chat messages not found response has a 2xx status code
func (o *GetConversationsChatMessagesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations chat messages not found response has a 3xx status code
func (o *GetConversationsChatMessagesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations chat messages not found response has a 4xx status code
func (o *GetConversationsChatMessagesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations chat messages not found response has a 5xx status code
func (o *GetConversationsChatMessagesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations chat messages not found response a status code equal to that given
func (o *GetConversationsChatMessagesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetConversationsChatMessagesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}/messages][%d] getConversationsChatMessagesNotFound  %+v", 404, o.Payload)
}

func (o *GetConversationsChatMessagesNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}/messages][%d] getConversationsChatMessagesNotFound  %+v", 404, o.Payload)
}

func (o *GetConversationsChatMessagesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsChatMessagesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsChatMessagesRequestTimeout creates a GetConversationsChatMessagesRequestTimeout with default headers values
func NewGetConversationsChatMessagesRequestTimeout() *GetConversationsChatMessagesRequestTimeout {
	return &GetConversationsChatMessagesRequestTimeout{}
}

/*
GetConversationsChatMessagesRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetConversationsChatMessagesRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations chat messages request timeout response has a 2xx status code
func (o *GetConversationsChatMessagesRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations chat messages request timeout response has a 3xx status code
func (o *GetConversationsChatMessagesRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations chat messages request timeout response has a 4xx status code
func (o *GetConversationsChatMessagesRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations chat messages request timeout response has a 5xx status code
func (o *GetConversationsChatMessagesRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations chat messages request timeout response a status code equal to that given
func (o *GetConversationsChatMessagesRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetConversationsChatMessagesRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}/messages][%d] getConversationsChatMessagesRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetConversationsChatMessagesRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}/messages][%d] getConversationsChatMessagesRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetConversationsChatMessagesRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsChatMessagesRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsChatMessagesRequestEntityTooLarge creates a GetConversationsChatMessagesRequestEntityTooLarge with default headers values
func NewGetConversationsChatMessagesRequestEntityTooLarge() *GetConversationsChatMessagesRequestEntityTooLarge {
	return &GetConversationsChatMessagesRequestEntityTooLarge{}
}

/*
GetConversationsChatMessagesRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetConversationsChatMessagesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations chat messages request entity too large response has a 2xx status code
func (o *GetConversationsChatMessagesRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations chat messages request entity too large response has a 3xx status code
func (o *GetConversationsChatMessagesRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations chat messages request entity too large response has a 4xx status code
func (o *GetConversationsChatMessagesRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations chat messages request entity too large response has a 5xx status code
func (o *GetConversationsChatMessagesRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations chat messages request entity too large response a status code equal to that given
func (o *GetConversationsChatMessagesRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetConversationsChatMessagesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}/messages][%d] getConversationsChatMessagesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetConversationsChatMessagesRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}/messages][%d] getConversationsChatMessagesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetConversationsChatMessagesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsChatMessagesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsChatMessagesUnsupportedMediaType creates a GetConversationsChatMessagesUnsupportedMediaType with default headers values
func NewGetConversationsChatMessagesUnsupportedMediaType() *GetConversationsChatMessagesUnsupportedMediaType {
	return &GetConversationsChatMessagesUnsupportedMediaType{}
}

/*
GetConversationsChatMessagesUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetConversationsChatMessagesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations chat messages unsupported media type response has a 2xx status code
func (o *GetConversationsChatMessagesUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations chat messages unsupported media type response has a 3xx status code
func (o *GetConversationsChatMessagesUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations chat messages unsupported media type response has a 4xx status code
func (o *GetConversationsChatMessagesUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations chat messages unsupported media type response has a 5xx status code
func (o *GetConversationsChatMessagesUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations chat messages unsupported media type response a status code equal to that given
func (o *GetConversationsChatMessagesUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetConversationsChatMessagesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}/messages][%d] getConversationsChatMessagesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetConversationsChatMessagesUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}/messages][%d] getConversationsChatMessagesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetConversationsChatMessagesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsChatMessagesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsChatMessagesTooManyRequests creates a GetConversationsChatMessagesTooManyRequests with default headers values
func NewGetConversationsChatMessagesTooManyRequests() *GetConversationsChatMessagesTooManyRequests {
	return &GetConversationsChatMessagesTooManyRequests{}
}

/*
GetConversationsChatMessagesTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetConversationsChatMessagesTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations chat messages too many requests response has a 2xx status code
func (o *GetConversationsChatMessagesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations chat messages too many requests response has a 3xx status code
func (o *GetConversationsChatMessagesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations chat messages too many requests response has a 4xx status code
func (o *GetConversationsChatMessagesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations chat messages too many requests response has a 5xx status code
func (o *GetConversationsChatMessagesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations chat messages too many requests response a status code equal to that given
func (o *GetConversationsChatMessagesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetConversationsChatMessagesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}/messages][%d] getConversationsChatMessagesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetConversationsChatMessagesTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}/messages][%d] getConversationsChatMessagesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetConversationsChatMessagesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsChatMessagesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsChatMessagesInternalServerError creates a GetConversationsChatMessagesInternalServerError with default headers values
func NewGetConversationsChatMessagesInternalServerError() *GetConversationsChatMessagesInternalServerError {
	return &GetConversationsChatMessagesInternalServerError{}
}

/*
GetConversationsChatMessagesInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetConversationsChatMessagesInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations chat messages internal server error response has a 2xx status code
func (o *GetConversationsChatMessagesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations chat messages internal server error response has a 3xx status code
func (o *GetConversationsChatMessagesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations chat messages internal server error response has a 4xx status code
func (o *GetConversationsChatMessagesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversations chat messages internal server error response has a 5xx status code
func (o *GetConversationsChatMessagesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get conversations chat messages internal server error response a status code equal to that given
func (o *GetConversationsChatMessagesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetConversationsChatMessagesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}/messages][%d] getConversationsChatMessagesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetConversationsChatMessagesInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}/messages][%d] getConversationsChatMessagesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetConversationsChatMessagesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsChatMessagesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsChatMessagesServiceUnavailable creates a GetConversationsChatMessagesServiceUnavailable with default headers values
func NewGetConversationsChatMessagesServiceUnavailable() *GetConversationsChatMessagesServiceUnavailable {
	return &GetConversationsChatMessagesServiceUnavailable{}
}

/*
GetConversationsChatMessagesServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetConversationsChatMessagesServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations chat messages service unavailable response has a 2xx status code
func (o *GetConversationsChatMessagesServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations chat messages service unavailable response has a 3xx status code
func (o *GetConversationsChatMessagesServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations chat messages service unavailable response has a 4xx status code
func (o *GetConversationsChatMessagesServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversations chat messages service unavailable response has a 5xx status code
func (o *GetConversationsChatMessagesServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get conversations chat messages service unavailable response a status code equal to that given
func (o *GetConversationsChatMessagesServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetConversationsChatMessagesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}/messages][%d] getConversationsChatMessagesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetConversationsChatMessagesServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}/messages][%d] getConversationsChatMessagesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetConversationsChatMessagesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsChatMessagesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsChatMessagesGatewayTimeout creates a GetConversationsChatMessagesGatewayTimeout with default headers values
func NewGetConversationsChatMessagesGatewayTimeout() *GetConversationsChatMessagesGatewayTimeout {
	return &GetConversationsChatMessagesGatewayTimeout{}
}

/*
GetConversationsChatMessagesGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetConversationsChatMessagesGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations chat messages gateway timeout response has a 2xx status code
func (o *GetConversationsChatMessagesGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations chat messages gateway timeout response has a 3xx status code
func (o *GetConversationsChatMessagesGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations chat messages gateway timeout response has a 4xx status code
func (o *GetConversationsChatMessagesGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversations chat messages gateway timeout response has a 5xx status code
func (o *GetConversationsChatMessagesGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get conversations chat messages gateway timeout response a status code equal to that given
func (o *GetConversationsChatMessagesGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetConversationsChatMessagesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}/messages][%d] getConversationsChatMessagesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetConversationsChatMessagesGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}/messages][%d] getConversationsChatMessagesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetConversationsChatMessagesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsChatMessagesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

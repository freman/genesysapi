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

// GetConversationsChatsReader is a Reader for the GetConversationsChats structure.
type GetConversationsChatsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConversationsChatsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetConversationsChatsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetConversationsChatsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetConversationsChatsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetConversationsChatsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetConversationsChatsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetConversationsChatsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetConversationsChatsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetConversationsChatsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetConversationsChatsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetConversationsChatsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetConversationsChatsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetConversationsChatsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetConversationsChatsOK creates a GetConversationsChatsOK with default headers values
func NewGetConversationsChatsOK() *GetConversationsChatsOK {
	return &GetConversationsChatsOK{}
}

/*
GetConversationsChatsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetConversationsChatsOK struct {
	Payload *models.ChatConversationEntityListing
}

// IsSuccess returns true when this get conversations chats o k response has a 2xx status code
func (o *GetConversationsChatsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get conversations chats o k response has a 3xx status code
func (o *GetConversationsChatsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations chats o k response has a 4xx status code
func (o *GetConversationsChatsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversations chats o k response has a 5xx status code
func (o *GetConversationsChatsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations chats o k response a status code equal to that given
func (o *GetConversationsChatsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetConversationsChatsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats][%d] getConversationsChatsOK  %+v", 200, o.Payload)
}

func (o *GetConversationsChatsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats][%d] getConversationsChatsOK  %+v", 200, o.Payload)
}

func (o *GetConversationsChatsOK) GetPayload() *models.ChatConversationEntityListing {
	return o.Payload
}

func (o *GetConversationsChatsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ChatConversationEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsChatsBadRequest creates a GetConversationsChatsBadRequest with default headers values
func NewGetConversationsChatsBadRequest() *GetConversationsChatsBadRequest {
	return &GetConversationsChatsBadRequest{}
}

/*
GetConversationsChatsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetConversationsChatsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations chats bad request response has a 2xx status code
func (o *GetConversationsChatsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations chats bad request response has a 3xx status code
func (o *GetConversationsChatsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations chats bad request response has a 4xx status code
func (o *GetConversationsChatsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations chats bad request response has a 5xx status code
func (o *GetConversationsChatsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations chats bad request response a status code equal to that given
func (o *GetConversationsChatsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetConversationsChatsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats][%d] getConversationsChatsBadRequest  %+v", 400, o.Payload)
}

func (o *GetConversationsChatsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats][%d] getConversationsChatsBadRequest  %+v", 400, o.Payload)
}

func (o *GetConversationsChatsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsChatsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsChatsUnauthorized creates a GetConversationsChatsUnauthorized with default headers values
func NewGetConversationsChatsUnauthorized() *GetConversationsChatsUnauthorized {
	return &GetConversationsChatsUnauthorized{}
}

/*
GetConversationsChatsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetConversationsChatsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations chats unauthorized response has a 2xx status code
func (o *GetConversationsChatsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations chats unauthorized response has a 3xx status code
func (o *GetConversationsChatsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations chats unauthorized response has a 4xx status code
func (o *GetConversationsChatsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations chats unauthorized response has a 5xx status code
func (o *GetConversationsChatsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations chats unauthorized response a status code equal to that given
func (o *GetConversationsChatsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetConversationsChatsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats][%d] getConversationsChatsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetConversationsChatsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats][%d] getConversationsChatsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetConversationsChatsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsChatsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsChatsForbidden creates a GetConversationsChatsForbidden with default headers values
func NewGetConversationsChatsForbidden() *GetConversationsChatsForbidden {
	return &GetConversationsChatsForbidden{}
}

/*
GetConversationsChatsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetConversationsChatsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations chats forbidden response has a 2xx status code
func (o *GetConversationsChatsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations chats forbidden response has a 3xx status code
func (o *GetConversationsChatsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations chats forbidden response has a 4xx status code
func (o *GetConversationsChatsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations chats forbidden response has a 5xx status code
func (o *GetConversationsChatsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations chats forbidden response a status code equal to that given
func (o *GetConversationsChatsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetConversationsChatsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats][%d] getConversationsChatsForbidden  %+v", 403, o.Payload)
}

func (o *GetConversationsChatsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats][%d] getConversationsChatsForbidden  %+v", 403, o.Payload)
}

func (o *GetConversationsChatsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsChatsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsChatsNotFound creates a GetConversationsChatsNotFound with default headers values
func NewGetConversationsChatsNotFound() *GetConversationsChatsNotFound {
	return &GetConversationsChatsNotFound{}
}

/*
GetConversationsChatsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetConversationsChatsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations chats not found response has a 2xx status code
func (o *GetConversationsChatsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations chats not found response has a 3xx status code
func (o *GetConversationsChatsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations chats not found response has a 4xx status code
func (o *GetConversationsChatsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations chats not found response has a 5xx status code
func (o *GetConversationsChatsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations chats not found response a status code equal to that given
func (o *GetConversationsChatsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetConversationsChatsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats][%d] getConversationsChatsNotFound  %+v", 404, o.Payload)
}

func (o *GetConversationsChatsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats][%d] getConversationsChatsNotFound  %+v", 404, o.Payload)
}

func (o *GetConversationsChatsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsChatsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsChatsRequestTimeout creates a GetConversationsChatsRequestTimeout with default headers values
func NewGetConversationsChatsRequestTimeout() *GetConversationsChatsRequestTimeout {
	return &GetConversationsChatsRequestTimeout{}
}

/*
GetConversationsChatsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetConversationsChatsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations chats request timeout response has a 2xx status code
func (o *GetConversationsChatsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations chats request timeout response has a 3xx status code
func (o *GetConversationsChatsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations chats request timeout response has a 4xx status code
func (o *GetConversationsChatsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations chats request timeout response has a 5xx status code
func (o *GetConversationsChatsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations chats request timeout response a status code equal to that given
func (o *GetConversationsChatsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetConversationsChatsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats][%d] getConversationsChatsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetConversationsChatsRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats][%d] getConversationsChatsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetConversationsChatsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsChatsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsChatsRequestEntityTooLarge creates a GetConversationsChatsRequestEntityTooLarge with default headers values
func NewGetConversationsChatsRequestEntityTooLarge() *GetConversationsChatsRequestEntityTooLarge {
	return &GetConversationsChatsRequestEntityTooLarge{}
}

/*
GetConversationsChatsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetConversationsChatsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations chats request entity too large response has a 2xx status code
func (o *GetConversationsChatsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations chats request entity too large response has a 3xx status code
func (o *GetConversationsChatsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations chats request entity too large response has a 4xx status code
func (o *GetConversationsChatsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations chats request entity too large response has a 5xx status code
func (o *GetConversationsChatsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations chats request entity too large response a status code equal to that given
func (o *GetConversationsChatsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetConversationsChatsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats][%d] getConversationsChatsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetConversationsChatsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats][%d] getConversationsChatsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetConversationsChatsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsChatsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsChatsUnsupportedMediaType creates a GetConversationsChatsUnsupportedMediaType with default headers values
func NewGetConversationsChatsUnsupportedMediaType() *GetConversationsChatsUnsupportedMediaType {
	return &GetConversationsChatsUnsupportedMediaType{}
}

/*
GetConversationsChatsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetConversationsChatsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations chats unsupported media type response has a 2xx status code
func (o *GetConversationsChatsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations chats unsupported media type response has a 3xx status code
func (o *GetConversationsChatsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations chats unsupported media type response has a 4xx status code
func (o *GetConversationsChatsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations chats unsupported media type response has a 5xx status code
func (o *GetConversationsChatsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations chats unsupported media type response a status code equal to that given
func (o *GetConversationsChatsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetConversationsChatsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats][%d] getConversationsChatsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetConversationsChatsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats][%d] getConversationsChatsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetConversationsChatsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsChatsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsChatsTooManyRequests creates a GetConversationsChatsTooManyRequests with default headers values
func NewGetConversationsChatsTooManyRequests() *GetConversationsChatsTooManyRequests {
	return &GetConversationsChatsTooManyRequests{}
}

/*
GetConversationsChatsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetConversationsChatsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations chats too many requests response has a 2xx status code
func (o *GetConversationsChatsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations chats too many requests response has a 3xx status code
func (o *GetConversationsChatsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations chats too many requests response has a 4xx status code
func (o *GetConversationsChatsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations chats too many requests response has a 5xx status code
func (o *GetConversationsChatsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations chats too many requests response a status code equal to that given
func (o *GetConversationsChatsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetConversationsChatsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats][%d] getConversationsChatsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetConversationsChatsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats][%d] getConversationsChatsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetConversationsChatsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsChatsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsChatsInternalServerError creates a GetConversationsChatsInternalServerError with default headers values
func NewGetConversationsChatsInternalServerError() *GetConversationsChatsInternalServerError {
	return &GetConversationsChatsInternalServerError{}
}

/*
GetConversationsChatsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetConversationsChatsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations chats internal server error response has a 2xx status code
func (o *GetConversationsChatsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations chats internal server error response has a 3xx status code
func (o *GetConversationsChatsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations chats internal server error response has a 4xx status code
func (o *GetConversationsChatsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversations chats internal server error response has a 5xx status code
func (o *GetConversationsChatsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get conversations chats internal server error response a status code equal to that given
func (o *GetConversationsChatsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetConversationsChatsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats][%d] getConversationsChatsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetConversationsChatsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats][%d] getConversationsChatsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetConversationsChatsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsChatsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsChatsServiceUnavailable creates a GetConversationsChatsServiceUnavailable with default headers values
func NewGetConversationsChatsServiceUnavailable() *GetConversationsChatsServiceUnavailable {
	return &GetConversationsChatsServiceUnavailable{}
}

/*
GetConversationsChatsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetConversationsChatsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations chats service unavailable response has a 2xx status code
func (o *GetConversationsChatsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations chats service unavailable response has a 3xx status code
func (o *GetConversationsChatsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations chats service unavailable response has a 4xx status code
func (o *GetConversationsChatsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversations chats service unavailable response has a 5xx status code
func (o *GetConversationsChatsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get conversations chats service unavailable response a status code equal to that given
func (o *GetConversationsChatsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetConversationsChatsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats][%d] getConversationsChatsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetConversationsChatsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats][%d] getConversationsChatsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetConversationsChatsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsChatsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsChatsGatewayTimeout creates a GetConversationsChatsGatewayTimeout with default headers values
func NewGetConversationsChatsGatewayTimeout() *GetConversationsChatsGatewayTimeout {
	return &GetConversationsChatsGatewayTimeout{}
}

/*
GetConversationsChatsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetConversationsChatsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations chats gateway timeout response has a 2xx status code
func (o *GetConversationsChatsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations chats gateway timeout response has a 3xx status code
func (o *GetConversationsChatsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations chats gateway timeout response has a 4xx status code
func (o *GetConversationsChatsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversations chats gateway timeout response has a 5xx status code
func (o *GetConversationsChatsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get conversations chats gateway timeout response a status code equal to that given
func (o *GetConversationsChatsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetConversationsChatsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats][%d] getConversationsChatsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetConversationsChatsGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats][%d] getConversationsChatsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetConversationsChatsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsChatsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

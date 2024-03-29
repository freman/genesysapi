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

// PostConversationsChatsReader is a Reader for the PostConversationsChats structure.
type PostConversationsChatsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostConversationsChatsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostConversationsChatsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostConversationsChatsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostConversationsChatsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostConversationsChatsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostConversationsChatsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostConversationsChatsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostConversationsChatsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostConversationsChatsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostConversationsChatsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostConversationsChatsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostConversationsChatsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostConversationsChatsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostConversationsChatsOK creates a PostConversationsChatsOK with default headers values
func NewPostConversationsChatsOK() *PostConversationsChatsOK {
	return &PostConversationsChatsOK{}
}

/*
PostConversationsChatsOK describes a response with status code 200, with default header values.

successful operation
*/
type PostConversationsChatsOK struct {
	Payload *models.ChatConversation
}

// IsSuccess returns true when this post conversations chats o k response has a 2xx status code
func (o *PostConversationsChatsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post conversations chats o k response has a 3xx status code
func (o *PostConversationsChatsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversations chats o k response has a 4xx status code
func (o *PostConversationsChatsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post conversations chats o k response has a 5xx status code
func (o *PostConversationsChatsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post conversations chats o k response a status code equal to that given
func (o *PostConversationsChatsOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostConversationsChatsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/chats][%d] postConversationsChatsOK  %+v", 200, o.Payload)
}

func (o *PostConversationsChatsOK) String() string {
	return fmt.Sprintf("[POST /api/v2/conversations/chats][%d] postConversationsChatsOK  %+v", 200, o.Payload)
}

func (o *PostConversationsChatsOK) GetPayload() *models.ChatConversation {
	return o.Payload
}

func (o *PostConversationsChatsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ChatConversation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsChatsBadRequest creates a PostConversationsChatsBadRequest with default headers values
func NewPostConversationsChatsBadRequest() *PostConversationsChatsBadRequest {
	return &PostConversationsChatsBadRequest{}
}

/*
PostConversationsChatsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostConversationsChatsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post conversations chats bad request response has a 2xx status code
func (o *PostConversationsChatsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post conversations chats bad request response has a 3xx status code
func (o *PostConversationsChatsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversations chats bad request response has a 4xx status code
func (o *PostConversationsChatsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post conversations chats bad request response has a 5xx status code
func (o *PostConversationsChatsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post conversations chats bad request response a status code equal to that given
func (o *PostConversationsChatsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostConversationsChatsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/chats][%d] postConversationsChatsBadRequest  %+v", 400, o.Payload)
}

func (o *PostConversationsChatsBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/conversations/chats][%d] postConversationsChatsBadRequest  %+v", 400, o.Payload)
}

func (o *PostConversationsChatsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsChatsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsChatsUnauthorized creates a PostConversationsChatsUnauthorized with default headers values
func NewPostConversationsChatsUnauthorized() *PostConversationsChatsUnauthorized {
	return &PostConversationsChatsUnauthorized{}
}

/*
PostConversationsChatsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostConversationsChatsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post conversations chats unauthorized response has a 2xx status code
func (o *PostConversationsChatsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post conversations chats unauthorized response has a 3xx status code
func (o *PostConversationsChatsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversations chats unauthorized response has a 4xx status code
func (o *PostConversationsChatsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post conversations chats unauthorized response has a 5xx status code
func (o *PostConversationsChatsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post conversations chats unauthorized response a status code equal to that given
func (o *PostConversationsChatsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostConversationsChatsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/chats][%d] postConversationsChatsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostConversationsChatsUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/conversations/chats][%d] postConversationsChatsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostConversationsChatsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsChatsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsChatsForbidden creates a PostConversationsChatsForbidden with default headers values
func NewPostConversationsChatsForbidden() *PostConversationsChatsForbidden {
	return &PostConversationsChatsForbidden{}
}

/*
PostConversationsChatsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostConversationsChatsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post conversations chats forbidden response has a 2xx status code
func (o *PostConversationsChatsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post conversations chats forbidden response has a 3xx status code
func (o *PostConversationsChatsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversations chats forbidden response has a 4xx status code
func (o *PostConversationsChatsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post conversations chats forbidden response has a 5xx status code
func (o *PostConversationsChatsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post conversations chats forbidden response a status code equal to that given
func (o *PostConversationsChatsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostConversationsChatsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/chats][%d] postConversationsChatsForbidden  %+v", 403, o.Payload)
}

func (o *PostConversationsChatsForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/conversations/chats][%d] postConversationsChatsForbidden  %+v", 403, o.Payload)
}

func (o *PostConversationsChatsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsChatsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsChatsNotFound creates a PostConversationsChatsNotFound with default headers values
func NewPostConversationsChatsNotFound() *PostConversationsChatsNotFound {
	return &PostConversationsChatsNotFound{}
}

/*
PostConversationsChatsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostConversationsChatsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post conversations chats not found response has a 2xx status code
func (o *PostConversationsChatsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post conversations chats not found response has a 3xx status code
func (o *PostConversationsChatsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversations chats not found response has a 4xx status code
func (o *PostConversationsChatsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post conversations chats not found response has a 5xx status code
func (o *PostConversationsChatsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post conversations chats not found response a status code equal to that given
func (o *PostConversationsChatsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostConversationsChatsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/chats][%d] postConversationsChatsNotFound  %+v", 404, o.Payload)
}

func (o *PostConversationsChatsNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/conversations/chats][%d] postConversationsChatsNotFound  %+v", 404, o.Payload)
}

func (o *PostConversationsChatsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsChatsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsChatsRequestTimeout creates a PostConversationsChatsRequestTimeout with default headers values
func NewPostConversationsChatsRequestTimeout() *PostConversationsChatsRequestTimeout {
	return &PostConversationsChatsRequestTimeout{}
}

/*
PostConversationsChatsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostConversationsChatsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post conversations chats request timeout response has a 2xx status code
func (o *PostConversationsChatsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post conversations chats request timeout response has a 3xx status code
func (o *PostConversationsChatsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversations chats request timeout response has a 4xx status code
func (o *PostConversationsChatsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post conversations chats request timeout response has a 5xx status code
func (o *PostConversationsChatsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post conversations chats request timeout response a status code equal to that given
func (o *PostConversationsChatsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostConversationsChatsRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/chats][%d] postConversationsChatsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostConversationsChatsRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/conversations/chats][%d] postConversationsChatsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostConversationsChatsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsChatsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsChatsRequestEntityTooLarge creates a PostConversationsChatsRequestEntityTooLarge with default headers values
func NewPostConversationsChatsRequestEntityTooLarge() *PostConversationsChatsRequestEntityTooLarge {
	return &PostConversationsChatsRequestEntityTooLarge{}
}

/*
PostConversationsChatsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PostConversationsChatsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post conversations chats request entity too large response has a 2xx status code
func (o *PostConversationsChatsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post conversations chats request entity too large response has a 3xx status code
func (o *PostConversationsChatsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversations chats request entity too large response has a 4xx status code
func (o *PostConversationsChatsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post conversations chats request entity too large response has a 5xx status code
func (o *PostConversationsChatsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post conversations chats request entity too large response a status code equal to that given
func (o *PostConversationsChatsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostConversationsChatsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/chats][%d] postConversationsChatsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostConversationsChatsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/conversations/chats][%d] postConversationsChatsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostConversationsChatsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsChatsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsChatsUnsupportedMediaType creates a PostConversationsChatsUnsupportedMediaType with default headers values
func NewPostConversationsChatsUnsupportedMediaType() *PostConversationsChatsUnsupportedMediaType {
	return &PostConversationsChatsUnsupportedMediaType{}
}

/*
PostConversationsChatsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostConversationsChatsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post conversations chats unsupported media type response has a 2xx status code
func (o *PostConversationsChatsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post conversations chats unsupported media type response has a 3xx status code
func (o *PostConversationsChatsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversations chats unsupported media type response has a 4xx status code
func (o *PostConversationsChatsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post conversations chats unsupported media type response has a 5xx status code
func (o *PostConversationsChatsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post conversations chats unsupported media type response a status code equal to that given
func (o *PostConversationsChatsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostConversationsChatsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/chats][%d] postConversationsChatsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostConversationsChatsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/conversations/chats][%d] postConversationsChatsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostConversationsChatsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsChatsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsChatsTooManyRequests creates a PostConversationsChatsTooManyRequests with default headers values
func NewPostConversationsChatsTooManyRequests() *PostConversationsChatsTooManyRequests {
	return &PostConversationsChatsTooManyRequests{}
}

/*
PostConversationsChatsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostConversationsChatsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post conversations chats too many requests response has a 2xx status code
func (o *PostConversationsChatsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post conversations chats too many requests response has a 3xx status code
func (o *PostConversationsChatsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversations chats too many requests response has a 4xx status code
func (o *PostConversationsChatsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post conversations chats too many requests response has a 5xx status code
func (o *PostConversationsChatsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post conversations chats too many requests response a status code equal to that given
func (o *PostConversationsChatsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostConversationsChatsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/chats][%d] postConversationsChatsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostConversationsChatsTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/conversations/chats][%d] postConversationsChatsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostConversationsChatsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsChatsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsChatsInternalServerError creates a PostConversationsChatsInternalServerError with default headers values
func NewPostConversationsChatsInternalServerError() *PostConversationsChatsInternalServerError {
	return &PostConversationsChatsInternalServerError{}
}

/*
PostConversationsChatsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostConversationsChatsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post conversations chats internal server error response has a 2xx status code
func (o *PostConversationsChatsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post conversations chats internal server error response has a 3xx status code
func (o *PostConversationsChatsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversations chats internal server error response has a 4xx status code
func (o *PostConversationsChatsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post conversations chats internal server error response has a 5xx status code
func (o *PostConversationsChatsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post conversations chats internal server error response a status code equal to that given
func (o *PostConversationsChatsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostConversationsChatsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/chats][%d] postConversationsChatsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostConversationsChatsInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/conversations/chats][%d] postConversationsChatsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostConversationsChatsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsChatsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsChatsServiceUnavailable creates a PostConversationsChatsServiceUnavailable with default headers values
func NewPostConversationsChatsServiceUnavailable() *PostConversationsChatsServiceUnavailable {
	return &PostConversationsChatsServiceUnavailable{}
}

/*
PostConversationsChatsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostConversationsChatsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post conversations chats service unavailable response has a 2xx status code
func (o *PostConversationsChatsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post conversations chats service unavailable response has a 3xx status code
func (o *PostConversationsChatsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversations chats service unavailable response has a 4xx status code
func (o *PostConversationsChatsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post conversations chats service unavailable response has a 5xx status code
func (o *PostConversationsChatsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post conversations chats service unavailable response a status code equal to that given
func (o *PostConversationsChatsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostConversationsChatsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/chats][%d] postConversationsChatsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostConversationsChatsServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/conversations/chats][%d] postConversationsChatsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostConversationsChatsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsChatsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsChatsGatewayTimeout creates a PostConversationsChatsGatewayTimeout with default headers values
func NewPostConversationsChatsGatewayTimeout() *PostConversationsChatsGatewayTimeout {
	return &PostConversationsChatsGatewayTimeout{}
}

/*
PostConversationsChatsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostConversationsChatsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post conversations chats gateway timeout response has a 2xx status code
func (o *PostConversationsChatsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post conversations chats gateway timeout response has a 3xx status code
func (o *PostConversationsChatsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversations chats gateway timeout response has a 4xx status code
func (o *PostConversationsChatsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post conversations chats gateway timeout response has a 5xx status code
func (o *PostConversationsChatsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post conversations chats gateway timeout response a status code equal to that given
func (o *PostConversationsChatsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostConversationsChatsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/chats][%d] postConversationsChatsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostConversationsChatsGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/conversations/chats][%d] postConversationsChatsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostConversationsChatsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsChatsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

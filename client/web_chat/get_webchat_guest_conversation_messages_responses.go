// Code generated by go-swagger; DO NOT EDIT.

package web_chat

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetWebchatGuestConversationMessagesReader is a Reader for the GetWebchatGuestConversationMessages structure.
type GetWebchatGuestConversationMessagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWebchatGuestConversationMessagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWebchatGuestConversationMessagesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetWebchatGuestConversationMessagesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetWebchatGuestConversationMessagesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetWebchatGuestConversationMessagesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetWebchatGuestConversationMessagesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetWebchatGuestConversationMessagesRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetWebchatGuestConversationMessagesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetWebchatGuestConversationMessagesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetWebchatGuestConversationMessagesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetWebchatGuestConversationMessagesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetWebchatGuestConversationMessagesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetWebchatGuestConversationMessagesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetWebchatGuestConversationMessagesOK creates a GetWebchatGuestConversationMessagesOK with default headers values
func NewGetWebchatGuestConversationMessagesOK() *GetWebchatGuestConversationMessagesOK {
	return &GetWebchatGuestConversationMessagesOK{}
}

/*
GetWebchatGuestConversationMessagesOK describes a response with status code 200, with default header values.

successful operation
*/
type GetWebchatGuestConversationMessagesOK struct {
	Payload *models.WebChatMessageEntityList
}

// IsSuccess returns true when this get webchat guest conversation messages o k response has a 2xx status code
func (o *GetWebchatGuestConversationMessagesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get webchat guest conversation messages o k response has a 3xx status code
func (o *GetWebchatGuestConversationMessagesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get webchat guest conversation messages o k response has a 4xx status code
func (o *GetWebchatGuestConversationMessagesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get webchat guest conversation messages o k response has a 5xx status code
func (o *GetWebchatGuestConversationMessagesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get webchat guest conversation messages o k response a status code equal to that given
func (o *GetWebchatGuestConversationMessagesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetWebchatGuestConversationMessagesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/webchat/guest/conversations/{conversationId}/messages][%d] getWebchatGuestConversationMessagesOK  %+v", 200, o.Payload)
}

func (o *GetWebchatGuestConversationMessagesOK) String() string {
	return fmt.Sprintf("[GET /api/v2/webchat/guest/conversations/{conversationId}/messages][%d] getWebchatGuestConversationMessagesOK  %+v", 200, o.Payload)
}

func (o *GetWebchatGuestConversationMessagesOK) GetPayload() *models.WebChatMessageEntityList {
	return o.Payload
}

func (o *GetWebchatGuestConversationMessagesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WebChatMessageEntityList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWebchatGuestConversationMessagesBadRequest creates a GetWebchatGuestConversationMessagesBadRequest with default headers values
func NewGetWebchatGuestConversationMessagesBadRequest() *GetWebchatGuestConversationMessagesBadRequest {
	return &GetWebchatGuestConversationMessagesBadRequest{}
}

/*
GetWebchatGuestConversationMessagesBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetWebchatGuestConversationMessagesBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get webchat guest conversation messages bad request response has a 2xx status code
func (o *GetWebchatGuestConversationMessagesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get webchat guest conversation messages bad request response has a 3xx status code
func (o *GetWebchatGuestConversationMessagesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get webchat guest conversation messages bad request response has a 4xx status code
func (o *GetWebchatGuestConversationMessagesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get webchat guest conversation messages bad request response has a 5xx status code
func (o *GetWebchatGuestConversationMessagesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get webchat guest conversation messages bad request response a status code equal to that given
func (o *GetWebchatGuestConversationMessagesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetWebchatGuestConversationMessagesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/webchat/guest/conversations/{conversationId}/messages][%d] getWebchatGuestConversationMessagesBadRequest  %+v", 400, o.Payload)
}

func (o *GetWebchatGuestConversationMessagesBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/webchat/guest/conversations/{conversationId}/messages][%d] getWebchatGuestConversationMessagesBadRequest  %+v", 400, o.Payload)
}

func (o *GetWebchatGuestConversationMessagesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWebchatGuestConversationMessagesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWebchatGuestConversationMessagesUnauthorized creates a GetWebchatGuestConversationMessagesUnauthorized with default headers values
func NewGetWebchatGuestConversationMessagesUnauthorized() *GetWebchatGuestConversationMessagesUnauthorized {
	return &GetWebchatGuestConversationMessagesUnauthorized{}
}

/*
GetWebchatGuestConversationMessagesUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetWebchatGuestConversationMessagesUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get webchat guest conversation messages unauthorized response has a 2xx status code
func (o *GetWebchatGuestConversationMessagesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get webchat guest conversation messages unauthorized response has a 3xx status code
func (o *GetWebchatGuestConversationMessagesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get webchat guest conversation messages unauthorized response has a 4xx status code
func (o *GetWebchatGuestConversationMessagesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get webchat guest conversation messages unauthorized response has a 5xx status code
func (o *GetWebchatGuestConversationMessagesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get webchat guest conversation messages unauthorized response a status code equal to that given
func (o *GetWebchatGuestConversationMessagesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetWebchatGuestConversationMessagesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/webchat/guest/conversations/{conversationId}/messages][%d] getWebchatGuestConversationMessagesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetWebchatGuestConversationMessagesUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/webchat/guest/conversations/{conversationId}/messages][%d] getWebchatGuestConversationMessagesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetWebchatGuestConversationMessagesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWebchatGuestConversationMessagesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWebchatGuestConversationMessagesForbidden creates a GetWebchatGuestConversationMessagesForbidden with default headers values
func NewGetWebchatGuestConversationMessagesForbidden() *GetWebchatGuestConversationMessagesForbidden {
	return &GetWebchatGuestConversationMessagesForbidden{}
}

/*
GetWebchatGuestConversationMessagesForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetWebchatGuestConversationMessagesForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get webchat guest conversation messages forbidden response has a 2xx status code
func (o *GetWebchatGuestConversationMessagesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get webchat guest conversation messages forbidden response has a 3xx status code
func (o *GetWebchatGuestConversationMessagesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get webchat guest conversation messages forbidden response has a 4xx status code
func (o *GetWebchatGuestConversationMessagesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get webchat guest conversation messages forbidden response has a 5xx status code
func (o *GetWebchatGuestConversationMessagesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get webchat guest conversation messages forbidden response a status code equal to that given
func (o *GetWebchatGuestConversationMessagesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetWebchatGuestConversationMessagesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/webchat/guest/conversations/{conversationId}/messages][%d] getWebchatGuestConversationMessagesForbidden  %+v", 403, o.Payload)
}

func (o *GetWebchatGuestConversationMessagesForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/webchat/guest/conversations/{conversationId}/messages][%d] getWebchatGuestConversationMessagesForbidden  %+v", 403, o.Payload)
}

func (o *GetWebchatGuestConversationMessagesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWebchatGuestConversationMessagesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWebchatGuestConversationMessagesNotFound creates a GetWebchatGuestConversationMessagesNotFound with default headers values
func NewGetWebchatGuestConversationMessagesNotFound() *GetWebchatGuestConversationMessagesNotFound {
	return &GetWebchatGuestConversationMessagesNotFound{}
}

/*
GetWebchatGuestConversationMessagesNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetWebchatGuestConversationMessagesNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get webchat guest conversation messages not found response has a 2xx status code
func (o *GetWebchatGuestConversationMessagesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get webchat guest conversation messages not found response has a 3xx status code
func (o *GetWebchatGuestConversationMessagesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get webchat guest conversation messages not found response has a 4xx status code
func (o *GetWebchatGuestConversationMessagesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get webchat guest conversation messages not found response has a 5xx status code
func (o *GetWebchatGuestConversationMessagesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get webchat guest conversation messages not found response a status code equal to that given
func (o *GetWebchatGuestConversationMessagesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetWebchatGuestConversationMessagesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/webchat/guest/conversations/{conversationId}/messages][%d] getWebchatGuestConversationMessagesNotFound  %+v", 404, o.Payload)
}

func (o *GetWebchatGuestConversationMessagesNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/webchat/guest/conversations/{conversationId}/messages][%d] getWebchatGuestConversationMessagesNotFound  %+v", 404, o.Payload)
}

func (o *GetWebchatGuestConversationMessagesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWebchatGuestConversationMessagesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWebchatGuestConversationMessagesRequestTimeout creates a GetWebchatGuestConversationMessagesRequestTimeout with default headers values
func NewGetWebchatGuestConversationMessagesRequestTimeout() *GetWebchatGuestConversationMessagesRequestTimeout {
	return &GetWebchatGuestConversationMessagesRequestTimeout{}
}

/*
GetWebchatGuestConversationMessagesRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetWebchatGuestConversationMessagesRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get webchat guest conversation messages request timeout response has a 2xx status code
func (o *GetWebchatGuestConversationMessagesRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get webchat guest conversation messages request timeout response has a 3xx status code
func (o *GetWebchatGuestConversationMessagesRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get webchat guest conversation messages request timeout response has a 4xx status code
func (o *GetWebchatGuestConversationMessagesRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get webchat guest conversation messages request timeout response has a 5xx status code
func (o *GetWebchatGuestConversationMessagesRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get webchat guest conversation messages request timeout response a status code equal to that given
func (o *GetWebchatGuestConversationMessagesRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetWebchatGuestConversationMessagesRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/webchat/guest/conversations/{conversationId}/messages][%d] getWebchatGuestConversationMessagesRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetWebchatGuestConversationMessagesRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/webchat/guest/conversations/{conversationId}/messages][%d] getWebchatGuestConversationMessagesRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetWebchatGuestConversationMessagesRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWebchatGuestConversationMessagesRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWebchatGuestConversationMessagesRequestEntityTooLarge creates a GetWebchatGuestConversationMessagesRequestEntityTooLarge with default headers values
func NewGetWebchatGuestConversationMessagesRequestEntityTooLarge() *GetWebchatGuestConversationMessagesRequestEntityTooLarge {
	return &GetWebchatGuestConversationMessagesRequestEntityTooLarge{}
}

/*
GetWebchatGuestConversationMessagesRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetWebchatGuestConversationMessagesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get webchat guest conversation messages request entity too large response has a 2xx status code
func (o *GetWebchatGuestConversationMessagesRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get webchat guest conversation messages request entity too large response has a 3xx status code
func (o *GetWebchatGuestConversationMessagesRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get webchat guest conversation messages request entity too large response has a 4xx status code
func (o *GetWebchatGuestConversationMessagesRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get webchat guest conversation messages request entity too large response has a 5xx status code
func (o *GetWebchatGuestConversationMessagesRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get webchat guest conversation messages request entity too large response a status code equal to that given
func (o *GetWebchatGuestConversationMessagesRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetWebchatGuestConversationMessagesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/webchat/guest/conversations/{conversationId}/messages][%d] getWebchatGuestConversationMessagesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetWebchatGuestConversationMessagesRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/webchat/guest/conversations/{conversationId}/messages][%d] getWebchatGuestConversationMessagesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetWebchatGuestConversationMessagesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWebchatGuestConversationMessagesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWebchatGuestConversationMessagesUnsupportedMediaType creates a GetWebchatGuestConversationMessagesUnsupportedMediaType with default headers values
func NewGetWebchatGuestConversationMessagesUnsupportedMediaType() *GetWebchatGuestConversationMessagesUnsupportedMediaType {
	return &GetWebchatGuestConversationMessagesUnsupportedMediaType{}
}

/*
GetWebchatGuestConversationMessagesUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetWebchatGuestConversationMessagesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get webchat guest conversation messages unsupported media type response has a 2xx status code
func (o *GetWebchatGuestConversationMessagesUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get webchat guest conversation messages unsupported media type response has a 3xx status code
func (o *GetWebchatGuestConversationMessagesUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get webchat guest conversation messages unsupported media type response has a 4xx status code
func (o *GetWebchatGuestConversationMessagesUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get webchat guest conversation messages unsupported media type response has a 5xx status code
func (o *GetWebchatGuestConversationMessagesUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get webchat guest conversation messages unsupported media type response a status code equal to that given
func (o *GetWebchatGuestConversationMessagesUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetWebchatGuestConversationMessagesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/webchat/guest/conversations/{conversationId}/messages][%d] getWebchatGuestConversationMessagesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetWebchatGuestConversationMessagesUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/webchat/guest/conversations/{conversationId}/messages][%d] getWebchatGuestConversationMessagesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetWebchatGuestConversationMessagesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWebchatGuestConversationMessagesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWebchatGuestConversationMessagesTooManyRequests creates a GetWebchatGuestConversationMessagesTooManyRequests with default headers values
func NewGetWebchatGuestConversationMessagesTooManyRequests() *GetWebchatGuestConversationMessagesTooManyRequests {
	return &GetWebchatGuestConversationMessagesTooManyRequests{}
}

/*
GetWebchatGuestConversationMessagesTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetWebchatGuestConversationMessagesTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get webchat guest conversation messages too many requests response has a 2xx status code
func (o *GetWebchatGuestConversationMessagesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get webchat guest conversation messages too many requests response has a 3xx status code
func (o *GetWebchatGuestConversationMessagesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get webchat guest conversation messages too many requests response has a 4xx status code
func (o *GetWebchatGuestConversationMessagesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get webchat guest conversation messages too many requests response has a 5xx status code
func (o *GetWebchatGuestConversationMessagesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get webchat guest conversation messages too many requests response a status code equal to that given
func (o *GetWebchatGuestConversationMessagesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetWebchatGuestConversationMessagesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/webchat/guest/conversations/{conversationId}/messages][%d] getWebchatGuestConversationMessagesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetWebchatGuestConversationMessagesTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/webchat/guest/conversations/{conversationId}/messages][%d] getWebchatGuestConversationMessagesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetWebchatGuestConversationMessagesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWebchatGuestConversationMessagesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWebchatGuestConversationMessagesInternalServerError creates a GetWebchatGuestConversationMessagesInternalServerError with default headers values
func NewGetWebchatGuestConversationMessagesInternalServerError() *GetWebchatGuestConversationMessagesInternalServerError {
	return &GetWebchatGuestConversationMessagesInternalServerError{}
}

/*
GetWebchatGuestConversationMessagesInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetWebchatGuestConversationMessagesInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get webchat guest conversation messages internal server error response has a 2xx status code
func (o *GetWebchatGuestConversationMessagesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get webchat guest conversation messages internal server error response has a 3xx status code
func (o *GetWebchatGuestConversationMessagesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get webchat guest conversation messages internal server error response has a 4xx status code
func (o *GetWebchatGuestConversationMessagesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get webchat guest conversation messages internal server error response has a 5xx status code
func (o *GetWebchatGuestConversationMessagesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get webchat guest conversation messages internal server error response a status code equal to that given
func (o *GetWebchatGuestConversationMessagesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetWebchatGuestConversationMessagesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/webchat/guest/conversations/{conversationId}/messages][%d] getWebchatGuestConversationMessagesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetWebchatGuestConversationMessagesInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/webchat/guest/conversations/{conversationId}/messages][%d] getWebchatGuestConversationMessagesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetWebchatGuestConversationMessagesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWebchatGuestConversationMessagesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWebchatGuestConversationMessagesServiceUnavailable creates a GetWebchatGuestConversationMessagesServiceUnavailable with default headers values
func NewGetWebchatGuestConversationMessagesServiceUnavailable() *GetWebchatGuestConversationMessagesServiceUnavailable {
	return &GetWebchatGuestConversationMessagesServiceUnavailable{}
}

/*
GetWebchatGuestConversationMessagesServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetWebchatGuestConversationMessagesServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get webchat guest conversation messages service unavailable response has a 2xx status code
func (o *GetWebchatGuestConversationMessagesServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get webchat guest conversation messages service unavailable response has a 3xx status code
func (o *GetWebchatGuestConversationMessagesServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get webchat guest conversation messages service unavailable response has a 4xx status code
func (o *GetWebchatGuestConversationMessagesServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get webchat guest conversation messages service unavailable response has a 5xx status code
func (o *GetWebchatGuestConversationMessagesServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get webchat guest conversation messages service unavailable response a status code equal to that given
func (o *GetWebchatGuestConversationMessagesServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetWebchatGuestConversationMessagesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/webchat/guest/conversations/{conversationId}/messages][%d] getWebchatGuestConversationMessagesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetWebchatGuestConversationMessagesServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/webchat/guest/conversations/{conversationId}/messages][%d] getWebchatGuestConversationMessagesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetWebchatGuestConversationMessagesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWebchatGuestConversationMessagesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWebchatGuestConversationMessagesGatewayTimeout creates a GetWebchatGuestConversationMessagesGatewayTimeout with default headers values
func NewGetWebchatGuestConversationMessagesGatewayTimeout() *GetWebchatGuestConversationMessagesGatewayTimeout {
	return &GetWebchatGuestConversationMessagesGatewayTimeout{}
}

/*
GetWebchatGuestConversationMessagesGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetWebchatGuestConversationMessagesGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get webchat guest conversation messages gateway timeout response has a 2xx status code
func (o *GetWebchatGuestConversationMessagesGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get webchat guest conversation messages gateway timeout response has a 3xx status code
func (o *GetWebchatGuestConversationMessagesGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get webchat guest conversation messages gateway timeout response has a 4xx status code
func (o *GetWebchatGuestConversationMessagesGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get webchat guest conversation messages gateway timeout response has a 5xx status code
func (o *GetWebchatGuestConversationMessagesGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get webchat guest conversation messages gateway timeout response a status code equal to that given
func (o *GetWebchatGuestConversationMessagesGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetWebchatGuestConversationMessagesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/webchat/guest/conversations/{conversationId}/messages][%d] getWebchatGuestConversationMessagesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetWebchatGuestConversationMessagesGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/webchat/guest/conversations/{conversationId}/messages][%d] getWebchatGuestConversationMessagesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetWebchatGuestConversationMessagesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWebchatGuestConversationMessagesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

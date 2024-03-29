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

// GetConversationsCallbackReader is a Reader for the GetConversationsCallback structure.
type GetConversationsCallbackReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConversationsCallbackReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetConversationsCallbackOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetConversationsCallbackBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetConversationsCallbackUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetConversationsCallbackForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetConversationsCallbackNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetConversationsCallbackRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetConversationsCallbackRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetConversationsCallbackUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetConversationsCallbackTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetConversationsCallbackInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetConversationsCallbackServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetConversationsCallbackGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetConversationsCallbackOK creates a GetConversationsCallbackOK with default headers values
func NewGetConversationsCallbackOK() *GetConversationsCallbackOK {
	return &GetConversationsCallbackOK{}
}

/*
GetConversationsCallbackOK describes a response with status code 200, with default header values.

successful operation
*/
type GetConversationsCallbackOK struct {
	Payload *models.CallbackConversation
}

// IsSuccess returns true when this get conversations callback o k response has a 2xx status code
func (o *GetConversationsCallbackOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get conversations callback o k response has a 3xx status code
func (o *GetConversationsCallbackOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations callback o k response has a 4xx status code
func (o *GetConversationsCallbackOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversations callback o k response has a 5xx status code
func (o *GetConversationsCallbackOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations callback o k response a status code equal to that given
func (o *GetConversationsCallbackOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetConversationsCallbackOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/callbacks/{conversationId}][%d] getConversationsCallbackOK  %+v", 200, o.Payload)
}

func (o *GetConversationsCallbackOK) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/callbacks/{conversationId}][%d] getConversationsCallbackOK  %+v", 200, o.Payload)
}

func (o *GetConversationsCallbackOK) GetPayload() *models.CallbackConversation {
	return o.Payload
}

func (o *GetConversationsCallbackOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CallbackConversation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCallbackBadRequest creates a GetConversationsCallbackBadRequest with default headers values
func NewGetConversationsCallbackBadRequest() *GetConversationsCallbackBadRequest {
	return &GetConversationsCallbackBadRequest{}
}

/*
GetConversationsCallbackBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetConversationsCallbackBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations callback bad request response has a 2xx status code
func (o *GetConversationsCallbackBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations callback bad request response has a 3xx status code
func (o *GetConversationsCallbackBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations callback bad request response has a 4xx status code
func (o *GetConversationsCallbackBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations callback bad request response has a 5xx status code
func (o *GetConversationsCallbackBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations callback bad request response a status code equal to that given
func (o *GetConversationsCallbackBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetConversationsCallbackBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/callbacks/{conversationId}][%d] getConversationsCallbackBadRequest  %+v", 400, o.Payload)
}

func (o *GetConversationsCallbackBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/callbacks/{conversationId}][%d] getConversationsCallbackBadRequest  %+v", 400, o.Payload)
}

func (o *GetConversationsCallbackBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCallbackBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCallbackUnauthorized creates a GetConversationsCallbackUnauthorized with default headers values
func NewGetConversationsCallbackUnauthorized() *GetConversationsCallbackUnauthorized {
	return &GetConversationsCallbackUnauthorized{}
}

/*
GetConversationsCallbackUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetConversationsCallbackUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations callback unauthorized response has a 2xx status code
func (o *GetConversationsCallbackUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations callback unauthorized response has a 3xx status code
func (o *GetConversationsCallbackUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations callback unauthorized response has a 4xx status code
func (o *GetConversationsCallbackUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations callback unauthorized response has a 5xx status code
func (o *GetConversationsCallbackUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations callback unauthorized response a status code equal to that given
func (o *GetConversationsCallbackUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetConversationsCallbackUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/callbacks/{conversationId}][%d] getConversationsCallbackUnauthorized  %+v", 401, o.Payload)
}

func (o *GetConversationsCallbackUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/callbacks/{conversationId}][%d] getConversationsCallbackUnauthorized  %+v", 401, o.Payload)
}

func (o *GetConversationsCallbackUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCallbackUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCallbackForbidden creates a GetConversationsCallbackForbidden with default headers values
func NewGetConversationsCallbackForbidden() *GetConversationsCallbackForbidden {
	return &GetConversationsCallbackForbidden{}
}

/*
GetConversationsCallbackForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetConversationsCallbackForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations callback forbidden response has a 2xx status code
func (o *GetConversationsCallbackForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations callback forbidden response has a 3xx status code
func (o *GetConversationsCallbackForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations callback forbidden response has a 4xx status code
func (o *GetConversationsCallbackForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations callback forbidden response has a 5xx status code
func (o *GetConversationsCallbackForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations callback forbidden response a status code equal to that given
func (o *GetConversationsCallbackForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetConversationsCallbackForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/callbacks/{conversationId}][%d] getConversationsCallbackForbidden  %+v", 403, o.Payload)
}

func (o *GetConversationsCallbackForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/callbacks/{conversationId}][%d] getConversationsCallbackForbidden  %+v", 403, o.Payload)
}

func (o *GetConversationsCallbackForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCallbackForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCallbackNotFound creates a GetConversationsCallbackNotFound with default headers values
func NewGetConversationsCallbackNotFound() *GetConversationsCallbackNotFound {
	return &GetConversationsCallbackNotFound{}
}

/*
GetConversationsCallbackNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetConversationsCallbackNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations callback not found response has a 2xx status code
func (o *GetConversationsCallbackNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations callback not found response has a 3xx status code
func (o *GetConversationsCallbackNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations callback not found response has a 4xx status code
func (o *GetConversationsCallbackNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations callback not found response has a 5xx status code
func (o *GetConversationsCallbackNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations callback not found response a status code equal to that given
func (o *GetConversationsCallbackNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetConversationsCallbackNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/callbacks/{conversationId}][%d] getConversationsCallbackNotFound  %+v", 404, o.Payload)
}

func (o *GetConversationsCallbackNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/callbacks/{conversationId}][%d] getConversationsCallbackNotFound  %+v", 404, o.Payload)
}

func (o *GetConversationsCallbackNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCallbackNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCallbackRequestTimeout creates a GetConversationsCallbackRequestTimeout with default headers values
func NewGetConversationsCallbackRequestTimeout() *GetConversationsCallbackRequestTimeout {
	return &GetConversationsCallbackRequestTimeout{}
}

/*
GetConversationsCallbackRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetConversationsCallbackRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations callback request timeout response has a 2xx status code
func (o *GetConversationsCallbackRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations callback request timeout response has a 3xx status code
func (o *GetConversationsCallbackRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations callback request timeout response has a 4xx status code
func (o *GetConversationsCallbackRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations callback request timeout response has a 5xx status code
func (o *GetConversationsCallbackRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations callback request timeout response a status code equal to that given
func (o *GetConversationsCallbackRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetConversationsCallbackRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/callbacks/{conversationId}][%d] getConversationsCallbackRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetConversationsCallbackRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/callbacks/{conversationId}][%d] getConversationsCallbackRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetConversationsCallbackRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCallbackRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCallbackRequestEntityTooLarge creates a GetConversationsCallbackRequestEntityTooLarge with default headers values
func NewGetConversationsCallbackRequestEntityTooLarge() *GetConversationsCallbackRequestEntityTooLarge {
	return &GetConversationsCallbackRequestEntityTooLarge{}
}

/*
GetConversationsCallbackRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetConversationsCallbackRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations callback request entity too large response has a 2xx status code
func (o *GetConversationsCallbackRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations callback request entity too large response has a 3xx status code
func (o *GetConversationsCallbackRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations callback request entity too large response has a 4xx status code
func (o *GetConversationsCallbackRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations callback request entity too large response has a 5xx status code
func (o *GetConversationsCallbackRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations callback request entity too large response a status code equal to that given
func (o *GetConversationsCallbackRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetConversationsCallbackRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/callbacks/{conversationId}][%d] getConversationsCallbackRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetConversationsCallbackRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/callbacks/{conversationId}][%d] getConversationsCallbackRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetConversationsCallbackRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCallbackRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCallbackUnsupportedMediaType creates a GetConversationsCallbackUnsupportedMediaType with default headers values
func NewGetConversationsCallbackUnsupportedMediaType() *GetConversationsCallbackUnsupportedMediaType {
	return &GetConversationsCallbackUnsupportedMediaType{}
}

/*
GetConversationsCallbackUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetConversationsCallbackUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations callback unsupported media type response has a 2xx status code
func (o *GetConversationsCallbackUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations callback unsupported media type response has a 3xx status code
func (o *GetConversationsCallbackUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations callback unsupported media type response has a 4xx status code
func (o *GetConversationsCallbackUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations callback unsupported media type response has a 5xx status code
func (o *GetConversationsCallbackUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations callback unsupported media type response a status code equal to that given
func (o *GetConversationsCallbackUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetConversationsCallbackUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/callbacks/{conversationId}][%d] getConversationsCallbackUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetConversationsCallbackUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/callbacks/{conversationId}][%d] getConversationsCallbackUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetConversationsCallbackUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCallbackUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCallbackTooManyRequests creates a GetConversationsCallbackTooManyRequests with default headers values
func NewGetConversationsCallbackTooManyRequests() *GetConversationsCallbackTooManyRequests {
	return &GetConversationsCallbackTooManyRequests{}
}

/*
GetConversationsCallbackTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetConversationsCallbackTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations callback too many requests response has a 2xx status code
func (o *GetConversationsCallbackTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations callback too many requests response has a 3xx status code
func (o *GetConversationsCallbackTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations callback too many requests response has a 4xx status code
func (o *GetConversationsCallbackTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations callback too many requests response has a 5xx status code
func (o *GetConversationsCallbackTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations callback too many requests response a status code equal to that given
func (o *GetConversationsCallbackTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetConversationsCallbackTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/callbacks/{conversationId}][%d] getConversationsCallbackTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetConversationsCallbackTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/callbacks/{conversationId}][%d] getConversationsCallbackTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetConversationsCallbackTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCallbackTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCallbackInternalServerError creates a GetConversationsCallbackInternalServerError with default headers values
func NewGetConversationsCallbackInternalServerError() *GetConversationsCallbackInternalServerError {
	return &GetConversationsCallbackInternalServerError{}
}

/*
GetConversationsCallbackInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetConversationsCallbackInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations callback internal server error response has a 2xx status code
func (o *GetConversationsCallbackInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations callback internal server error response has a 3xx status code
func (o *GetConversationsCallbackInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations callback internal server error response has a 4xx status code
func (o *GetConversationsCallbackInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversations callback internal server error response has a 5xx status code
func (o *GetConversationsCallbackInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get conversations callback internal server error response a status code equal to that given
func (o *GetConversationsCallbackInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetConversationsCallbackInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/callbacks/{conversationId}][%d] getConversationsCallbackInternalServerError  %+v", 500, o.Payload)
}

func (o *GetConversationsCallbackInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/callbacks/{conversationId}][%d] getConversationsCallbackInternalServerError  %+v", 500, o.Payload)
}

func (o *GetConversationsCallbackInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCallbackInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCallbackServiceUnavailable creates a GetConversationsCallbackServiceUnavailable with default headers values
func NewGetConversationsCallbackServiceUnavailable() *GetConversationsCallbackServiceUnavailable {
	return &GetConversationsCallbackServiceUnavailable{}
}

/*
GetConversationsCallbackServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetConversationsCallbackServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations callback service unavailable response has a 2xx status code
func (o *GetConversationsCallbackServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations callback service unavailable response has a 3xx status code
func (o *GetConversationsCallbackServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations callback service unavailable response has a 4xx status code
func (o *GetConversationsCallbackServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversations callback service unavailable response has a 5xx status code
func (o *GetConversationsCallbackServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get conversations callback service unavailable response a status code equal to that given
func (o *GetConversationsCallbackServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetConversationsCallbackServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/callbacks/{conversationId}][%d] getConversationsCallbackServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetConversationsCallbackServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/callbacks/{conversationId}][%d] getConversationsCallbackServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetConversationsCallbackServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCallbackServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCallbackGatewayTimeout creates a GetConversationsCallbackGatewayTimeout with default headers values
func NewGetConversationsCallbackGatewayTimeout() *GetConversationsCallbackGatewayTimeout {
	return &GetConversationsCallbackGatewayTimeout{}
}

/*
GetConversationsCallbackGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetConversationsCallbackGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations callback gateway timeout response has a 2xx status code
func (o *GetConversationsCallbackGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations callback gateway timeout response has a 3xx status code
func (o *GetConversationsCallbackGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations callback gateway timeout response has a 4xx status code
func (o *GetConversationsCallbackGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversations callback gateway timeout response has a 5xx status code
func (o *GetConversationsCallbackGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get conversations callback gateway timeout response a status code equal to that given
func (o *GetConversationsCallbackGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetConversationsCallbackGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/callbacks/{conversationId}][%d] getConversationsCallbackGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetConversationsCallbackGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/callbacks/{conversationId}][%d] getConversationsCallbackGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetConversationsCallbackGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCallbackGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

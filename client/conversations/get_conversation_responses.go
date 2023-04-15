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

// GetConversationReader is a Reader for the GetConversation structure.
type GetConversationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConversationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetConversationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetConversationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetConversationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetConversationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetConversationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetConversationRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetConversationRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetConversationUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetConversationTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetConversationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetConversationServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetConversationGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetConversationOK creates a GetConversationOK with default headers values
func NewGetConversationOK() *GetConversationOK {
	return &GetConversationOK{}
}

/*
GetConversationOK describes a response with status code 200, with default header values.

successful operation
*/
type GetConversationOK struct {
	Payload *models.Conversation
}

// IsSuccess returns true when this get conversation o k response has a 2xx status code
func (o *GetConversationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get conversation o k response has a 3xx status code
func (o *GetConversationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversation o k response has a 4xx status code
func (o *GetConversationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversation o k response has a 5xx status code
func (o *GetConversationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversation o k response a status code equal to that given
func (o *GetConversationOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetConversationOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}][%d] getConversationOK  %+v", 200, o.Payload)
}

func (o *GetConversationOK) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}][%d] getConversationOK  %+v", 200, o.Payload)
}

func (o *GetConversationOK) GetPayload() *models.Conversation {
	return o.Payload
}

func (o *GetConversationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Conversation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationBadRequest creates a GetConversationBadRequest with default headers values
func NewGetConversationBadRequest() *GetConversationBadRequest {
	return &GetConversationBadRequest{}
}

/*
GetConversationBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetConversationBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversation bad request response has a 2xx status code
func (o *GetConversationBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversation bad request response has a 3xx status code
func (o *GetConversationBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversation bad request response has a 4xx status code
func (o *GetConversationBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversation bad request response has a 5xx status code
func (o *GetConversationBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversation bad request response a status code equal to that given
func (o *GetConversationBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetConversationBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}][%d] getConversationBadRequest  %+v", 400, o.Payload)
}

func (o *GetConversationBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}][%d] getConversationBadRequest  %+v", 400, o.Payload)
}

func (o *GetConversationBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationUnauthorized creates a GetConversationUnauthorized with default headers values
func NewGetConversationUnauthorized() *GetConversationUnauthorized {
	return &GetConversationUnauthorized{}
}

/*
GetConversationUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetConversationUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversation unauthorized response has a 2xx status code
func (o *GetConversationUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversation unauthorized response has a 3xx status code
func (o *GetConversationUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversation unauthorized response has a 4xx status code
func (o *GetConversationUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversation unauthorized response has a 5xx status code
func (o *GetConversationUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversation unauthorized response a status code equal to that given
func (o *GetConversationUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetConversationUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}][%d] getConversationUnauthorized  %+v", 401, o.Payload)
}

func (o *GetConversationUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}][%d] getConversationUnauthorized  %+v", 401, o.Payload)
}

func (o *GetConversationUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationForbidden creates a GetConversationForbidden with default headers values
func NewGetConversationForbidden() *GetConversationForbidden {
	return &GetConversationForbidden{}
}

/*
GetConversationForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetConversationForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversation forbidden response has a 2xx status code
func (o *GetConversationForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversation forbidden response has a 3xx status code
func (o *GetConversationForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversation forbidden response has a 4xx status code
func (o *GetConversationForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversation forbidden response has a 5xx status code
func (o *GetConversationForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversation forbidden response a status code equal to that given
func (o *GetConversationForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetConversationForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}][%d] getConversationForbidden  %+v", 403, o.Payload)
}

func (o *GetConversationForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}][%d] getConversationForbidden  %+v", 403, o.Payload)
}

func (o *GetConversationForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationNotFound creates a GetConversationNotFound with default headers values
func NewGetConversationNotFound() *GetConversationNotFound {
	return &GetConversationNotFound{}
}

/*
GetConversationNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetConversationNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversation not found response has a 2xx status code
func (o *GetConversationNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversation not found response has a 3xx status code
func (o *GetConversationNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversation not found response has a 4xx status code
func (o *GetConversationNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversation not found response has a 5xx status code
func (o *GetConversationNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversation not found response a status code equal to that given
func (o *GetConversationNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetConversationNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}][%d] getConversationNotFound  %+v", 404, o.Payload)
}

func (o *GetConversationNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}][%d] getConversationNotFound  %+v", 404, o.Payload)
}

func (o *GetConversationNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationRequestTimeout creates a GetConversationRequestTimeout with default headers values
func NewGetConversationRequestTimeout() *GetConversationRequestTimeout {
	return &GetConversationRequestTimeout{}
}

/*
GetConversationRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetConversationRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversation request timeout response has a 2xx status code
func (o *GetConversationRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversation request timeout response has a 3xx status code
func (o *GetConversationRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversation request timeout response has a 4xx status code
func (o *GetConversationRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversation request timeout response has a 5xx status code
func (o *GetConversationRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversation request timeout response a status code equal to that given
func (o *GetConversationRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetConversationRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}][%d] getConversationRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetConversationRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}][%d] getConversationRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetConversationRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationRequestEntityTooLarge creates a GetConversationRequestEntityTooLarge with default headers values
func NewGetConversationRequestEntityTooLarge() *GetConversationRequestEntityTooLarge {
	return &GetConversationRequestEntityTooLarge{}
}

/*
GetConversationRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetConversationRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversation request entity too large response has a 2xx status code
func (o *GetConversationRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversation request entity too large response has a 3xx status code
func (o *GetConversationRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversation request entity too large response has a 4xx status code
func (o *GetConversationRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversation request entity too large response has a 5xx status code
func (o *GetConversationRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversation request entity too large response a status code equal to that given
func (o *GetConversationRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetConversationRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}][%d] getConversationRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetConversationRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}][%d] getConversationRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetConversationRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationUnsupportedMediaType creates a GetConversationUnsupportedMediaType with default headers values
func NewGetConversationUnsupportedMediaType() *GetConversationUnsupportedMediaType {
	return &GetConversationUnsupportedMediaType{}
}

/*
GetConversationUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetConversationUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversation unsupported media type response has a 2xx status code
func (o *GetConversationUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversation unsupported media type response has a 3xx status code
func (o *GetConversationUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversation unsupported media type response has a 4xx status code
func (o *GetConversationUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversation unsupported media type response has a 5xx status code
func (o *GetConversationUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversation unsupported media type response a status code equal to that given
func (o *GetConversationUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetConversationUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}][%d] getConversationUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetConversationUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}][%d] getConversationUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetConversationUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationTooManyRequests creates a GetConversationTooManyRequests with default headers values
func NewGetConversationTooManyRequests() *GetConversationTooManyRequests {
	return &GetConversationTooManyRequests{}
}

/*
GetConversationTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetConversationTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversation too many requests response has a 2xx status code
func (o *GetConversationTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversation too many requests response has a 3xx status code
func (o *GetConversationTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversation too many requests response has a 4xx status code
func (o *GetConversationTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversation too many requests response has a 5xx status code
func (o *GetConversationTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversation too many requests response a status code equal to that given
func (o *GetConversationTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetConversationTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}][%d] getConversationTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetConversationTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}][%d] getConversationTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetConversationTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationInternalServerError creates a GetConversationInternalServerError with default headers values
func NewGetConversationInternalServerError() *GetConversationInternalServerError {
	return &GetConversationInternalServerError{}
}

/*
GetConversationInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetConversationInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversation internal server error response has a 2xx status code
func (o *GetConversationInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversation internal server error response has a 3xx status code
func (o *GetConversationInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversation internal server error response has a 4xx status code
func (o *GetConversationInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversation internal server error response has a 5xx status code
func (o *GetConversationInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get conversation internal server error response a status code equal to that given
func (o *GetConversationInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetConversationInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}][%d] getConversationInternalServerError  %+v", 500, o.Payload)
}

func (o *GetConversationInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}][%d] getConversationInternalServerError  %+v", 500, o.Payload)
}

func (o *GetConversationInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationServiceUnavailable creates a GetConversationServiceUnavailable with default headers values
func NewGetConversationServiceUnavailable() *GetConversationServiceUnavailable {
	return &GetConversationServiceUnavailable{}
}

/*
GetConversationServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetConversationServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversation service unavailable response has a 2xx status code
func (o *GetConversationServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversation service unavailable response has a 3xx status code
func (o *GetConversationServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversation service unavailable response has a 4xx status code
func (o *GetConversationServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversation service unavailable response has a 5xx status code
func (o *GetConversationServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get conversation service unavailable response a status code equal to that given
func (o *GetConversationServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetConversationServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}][%d] getConversationServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetConversationServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}][%d] getConversationServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetConversationServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationGatewayTimeout creates a GetConversationGatewayTimeout with default headers values
func NewGetConversationGatewayTimeout() *GetConversationGatewayTimeout {
	return &GetConversationGatewayTimeout{}
}

/*
GetConversationGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetConversationGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversation gateway timeout response has a 2xx status code
func (o *GetConversationGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversation gateway timeout response has a 3xx status code
func (o *GetConversationGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversation gateway timeout response has a 4xx status code
func (o *GetConversationGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversation gateway timeout response has a 5xx status code
func (o *GetConversationGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get conversation gateway timeout response a status code equal to that given
func (o *GetConversationGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetConversationGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}][%d] getConversationGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetConversationGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}][%d] getConversationGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetConversationGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

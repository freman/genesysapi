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

// GetConversationsMessagesReader is a Reader for the GetConversationsMessages structure.
type GetConversationsMessagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConversationsMessagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetConversationsMessagesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetConversationsMessagesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetConversationsMessagesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetConversationsMessagesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetConversationsMessagesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetConversationsMessagesRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetConversationsMessagesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetConversationsMessagesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetConversationsMessagesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetConversationsMessagesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetConversationsMessagesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetConversationsMessagesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetConversationsMessagesOK creates a GetConversationsMessagesOK with default headers values
func NewGetConversationsMessagesOK() *GetConversationsMessagesOK {
	return &GetConversationsMessagesOK{}
}

/*
GetConversationsMessagesOK describes a response with status code 200, with default header values.

successful operation
*/
type GetConversationsMessagesOK struct {
	Payload *models.MessageConversationEntityListing
}

// IsSuccess returns true when this get conversations messages o k response has a 2xx status code
func (o *GetConversationsMessagesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get conversations messages o k response has a 3xx status code
func (o *GetConversationsMessagesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations messages o k response has a 4xx status code
func (o *GetConversationsMessagesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversations messages o k response has a 5xx status code
func (o *GetConversationsMessagesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations messages o k response a status code equal to that given
func (o *GetConversationsMessagesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetConversationsMessagesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages][%d] getConversationsMessagesOK  %+v", 200, o.Payload)
}

func (o *GetConversationsMessagesOK) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages][%d] getConversationsMessagesOK  %+v", 200, o.Payload)
}

func (o *GetConversationsMessagesOK) GetPayload() *models.MessageConversationEntityListing {
	return o.Payload
}

func (o *GetConversationsMessagesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MessageConversationEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagesBadRequest creates a GetConversationsMessagesBadRequest with default headers values
func NewGetConversationsMessagesBadRequest() *GetConversationsMessagesBadRequest {
	return &GetConversationsMessagesBadRequest{}
}

/*
GetConversationsMessagesBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetConversationsMessagesBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations messages bad request response has a 2xx status code
func (o *GetConversationsMessagesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations messages bad request response has a 3xx status code
func (o *GetConversationsMessagesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations messages bad request response has a 4xx status code
func (o *GetConversationsMessagesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations messages bad request response has a 5xx status code
func (o *GetConversationsMessagesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations messages bad request response a status code equal to that given
func (o *GetConversationsMessagesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetConversationsMessagesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages][%d] getConversationsMessagesBadRequest  %+v", 400, o.Payload)
}

func (o *GetConversationsMessagesBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages][%d] getConversationsMessagesBadRequest  %+v", 400, o.Payload)
}

func (o *GetConversationsMessagesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagesUnauthorized creates a GetConversationsMessagesUnauthorized with default headers values
func NewGetConversationsMessagesUnauthorized() *GetConversationsMessagesUnauthorized {
	return &GetConversationsMessagesUnauthorized{}
}

/*
GetConversationsMessagesUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetConversationsMessagesUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations messages unauthorized response has a 2xx status code
func (o *GetConversationsMessagesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations messages unauthorized response has a 3xx status code
func (o *GetConversationsMessagesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations messages unauthorized response has a 4xx status code
func (o *GetConversationsMessagesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations messages unauthorized response has a 5xx status code
func (o *GetConversationsMessagesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations messages unauthorized response a status code equal to that given
func (o *GetConversationsMessagesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetConversationsMessagesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages][%d] getConversationsMessagesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetConversationsMessagesUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages][%d] getConversationsMessagesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetConversationsMessagesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagesForbidden creates a GetConversationsMessagesForbidden with default headers values
func NewGetConversationsMessagesForbidden() *GetConversationsMessagesForbidden {
	return &GetConversationsMessagesForbidden{}
}

/*
GetConversationsMessagesForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetConversationsMessagesForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations messages forbidden response has a 2xx status code
func (o *GetConversationsMessagesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations messages forbidden response has a 3xx status code
func (o *GetConversationsMessagesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations messages forbidden response has a 4xx status code
func (o *GetConversationsMessagesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations messages forbidden response has a 5xx status code
func (o *GetConversationsMessagesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations messages forbidden response a status code equal to that given
func (o *GetConversationsMessagesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetConversationsMessagesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages][%d] getConversationsMessagesForbidden  %+v", 403, o.Payload)
}

func (o *GetConversationsMessagesForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages][%d] getConversationsMessagesForbidden  %+v", 403, o.Payload)
}

func (o *GetConversationsMessagesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagesNotFound creates a GetConversationsMessagesNotFound with default headers values
func NewGetConversationsMessagesNotFound() *GetConversationsMessagesNotFound {
	return &GetConversationsMessagesNotFound{}
}

/*
GetConversationsMessagesNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetConversationsMessagesNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations messages not found response has a 2xx status code
func (o *GetConversationsMessagesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations messages not found response has a 3xx status code
func (o *GetConversationsMessagesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations messages not found response has a 4xx status code
func (o *GetConversationsMessagesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations messages not found response has a 5xx status code
func (o *GetConversationsMessagesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations messages not found response a status code equal to that given
func (o *GetConversationsMessagesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetConversationsMessagesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages][%d] getConversationsMessagesNotFound  %+v", 404, o.Payload)
}

func (o *GetConversationsMessagesNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages][%d] getConversationsMessagesNotFound  %+v", 404, o.Payload)
}

func (o *GetConversationsMessagesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagesRequestTimeout creates a GetConversationsMessagesRequestTimeout with default headers values
func NewGetConversationsMessagesRequestTimeout() *GetConversationsMessagesRequestTimeout {
	return &GetConversationsMessagesRequestTimeout{}
}

/*
GetConversationsMessagesRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetConversationsMessagesRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations messages request timeout response has a 2xx status code
func (o *GetConversationsMessagesRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations messages request timeout response has a 3xx status code
func (o *GetConversationsMessagesRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations messages request timeout response has a 4xx status code
func (o *GetConversationsMessagesRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations messages request timeout response has a 5xx status code
func (o *GetConversationsMessagesRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations messages request timeout response a status code equal to that given
func (o *GetConversationsMessagesRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetConversationsMessagesRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages][%d] getConversationsMessagesRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetConversationsMessagesRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages][%d] getConversationsMessagesRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetConversationsMessagesRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagesRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagesRequestEntityTooLarge creates a GetConversationsMessagesRequestEntityTooLarge with default headers values
func NewGetConversationsMessagesRequestEntityTooLarge() *GetConversationsMessagesRequestEntityTooLarge {
	return &GetConversationsMessagesRequestEntityTooLarge{}
}

/*
GetConversationsMessagesRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetConversationsMessagesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations messages request entity too large response has a 2xx status code
func (o *GetConversationsMessagesRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations messages request entity too large response has a 3xx status code
func (o *GetConversationsMessagesRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations messages request entity too large response has a 4xx status code
func (o *GetConversationsMessagesRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations messages request entity too large response has a 5xx status code
func (o *GetConversationsMessagesRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations messages request entity too large response a status code equal to that given
func (o *GetConversationsMessagesRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetConversationsMessagesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages][%d] getConversationsMessagesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetConversationsMessagesRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages][%d] getConversationsMessagesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetConversationsMessagesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagesUnsupportedMediaType creates a GetConversationsMessagesUnsupportedMediaType with default headers values
func NewGetConversationsMessagesUnsupportedMediaType() *GetConversationsMessagesUnsupportedMediaType {
	return &GetConversationsMessagesUnsupportedMediaType{}
}

/*
GetConversationsMessagesUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetConversationsMessagesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations messages unsupported media type response has a 2xx status code
func (o *GetConversationsMessagesUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations messages unsupported media type response has a 3xx status code
func (o *GetConversationsMessagesUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations messages unsupported media type response has a 4xx status code
func (o *GetConversationsMessagesUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations messages unsupported media type response has a 5xx status code
func (o *GetConversationsMessagesUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations messages unsupported media type response a status code equal to that given
func (o *GetConversationsMessagesUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetConversationsMessagesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages][%d] getConversationsMessagesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetConversationsMessagesUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages][%d] getConversationsMessagesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetConversationsMessagesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagesTooManyRequests creates a GetConversationsMessagesTooManyRequests with default headers values
func NewGetConversationsMessagesTooManyRequests() *GetConversationsMessagesTooManyRequests {
	return &GetConversationsMessagesTooManyRequests{}
}

/*
GetConversationsMessagesTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetConversationsMessagesTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations messages too many requests response has a 2xx status code
func (o *GetConversationsMessagesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations messages too many requests response has a 3xx status code
func (o *GetConversationsMessagesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations messages too many requests response has a 4xx status code
func (o *GetConversationsMessagesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations messages too many requests response has a 5xx status code
func (o *GetConversationsMessagesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations messages too many requests response a status code equal to that given
func (o *GetConversationsMessagesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetConversationsMessagesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages][%d] getConversationsMessagesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetConversationsMessagesTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages][%d] getConversationsMessagesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetConversationsMessagesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagesInternalServerError creates a GetConversationsMessagesInternalServerError with default headers values
func NewGetConversationsMessagesInternalServerError() *GetConversationsMessagesInternalServerError {
	return &GetConversationsMessagesInternalServerError{}
}

/*
GetConversationsMessagesInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetConversationsMessagesInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations messages internal server error response has a 2xx status code
func (o *GetConversationsMessagesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations messages internal server error response has a 3xx status code
func (o *GetConversationsMessagesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations messages internal server error response has a 4xx status code
func (o *GetConversationsMessagesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversations messages internal server error response has a 5xx status code
func (o *GetConversationsMessagesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get conversations messages internal server error response a status code equal to that given
func (o *GetConversationsMessagesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetConversationsMessagesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages][%d] getConversationsMessagesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetConversationsMessagesInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages][%d] getConversationsMessagesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetConversationsMessagesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagesServiceUnavailable creates a GetConversationsMessagesServiceUnavailable with default headers values
func NewGetConversationsMessagesServiceUnavailable() *GetConversationsMessagesServiceUnavailable {
	return &GetConversationsMessagesServiceUnavailable{}
}

/*
GetConversationsMessagesServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetConversationsMessagesServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations messages service unavailable response has a 2xx status code
func (o *GetConversationsMessagesServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations messages service unavailable response has a 3xx status code
func (o *GetConversationsMessagesServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations messages service unavailable response has a 4xx status code
func (o *GetConversationsMessagesServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversations messages service unavailable response has a 5xx status code
func (o *GetConversationsMessagesServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get conversations messages service unavailable response a status code equal to that given
func (o *GetConversationsMessagesServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetConversationsMessagesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages][%d] getConversationsMessagesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetConversationsMessagesServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages][%d] getConversationsMessagesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetConversationsMessagesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagesGatewayTimeout creates a GetConversationsMessagesGatewayTimeout with default headers values
func NewGetConversationsMessagesGatewayTimeout() *GetConversationsMessagesGatewayTimeout {
	return &GetConversationsMessagesGatewayTimeout{}
}

/*
GetConversationsMessagesGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetConversationsMessagesGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations messages gateway timeout response has a 2xx status code
func (o *GetConversationsMessagesGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations messages gateway timeout response has a 3xx status code
func (o *GetConversationsMessagesGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations messages gateway timeout response has a 4xx status code
func (o *GetConversationsMessagesGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversations messages gateway timeout response has a 5xx status code
func (o *GetConversationsMessagesGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get conversations messages gateway timeout response a status code equal to that given
func (o *GetConversationsMessagesGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetConversationsMessagesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages][%d] getConversationsMessagesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetConversationsMessagesGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages][%d] getConversationsMessagesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetConversationsMessagesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

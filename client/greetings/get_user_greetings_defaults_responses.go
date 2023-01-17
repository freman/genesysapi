// Code generated by go-swagger; DO NOT EDIT.

package greetings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetUserGreetingsDefaultsReader is a Reader for the GetUserGreetingsDefaults structure.
type GetUserGreetingsDefaultsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserGreetingsDefaultsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUserGreetingsDefaultsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetUserGreetingsDefaultsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetUserGreetingsDefaultsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetUserGreetingsDefaultsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetUserGreetingsDefaultsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetUserGreetingsDefaultsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetUserGreetingsDefaultsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetUserGreetingsDefaultsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetUserGreetingsDefaultsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUserGreetingsDefaultsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetUserGreetingsDefaultsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetUserGreetingsDefaultsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUserGreetingsDefaultsOK creates a GetUserGreetingsDefaultsOK with default headers values
func NewGetUserGreetingsDefaultsOK() *GetUserGreetingsDefaultsOK {
	return &GetUserGreetingsDefaultsOK{}
}

/*
GetUserGreetingsDefaultsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetUserGreetingsDefaultsOK struct {
	Payload *models.DefaultGreetingList
}

// IsSuccess returns true when this get user greetings defaults o k response has a 2xx status code
func (o *GetUserGreetingsDefaultsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get user greetings defaults o k response has a 3xx status code
func (o *GetUserGreetingsDefaultsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user greetings defaults o k response has a 4xx status code
func (o *GetUserGreetingsDefaultsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get user greetings defaults o k response has a 5xx status code
func (o *GetUserGreetingsDefaultsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get user greetings defaults o k response a status code equal to that given
func (o *GetUserGreetingsDefaultsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetUserGreetingsDefaultsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/greetings/defaults][%d] getUserGreetingsDefaultsOK  %+v", 200, o.Payload)
}

func (o *GetUserGreetingsDefaultsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/greetings/defaults][%d] getUserGreetingsDefaultsOK  %+v", 200, o.Payload)
}

func (o *GetUserGreetingsDefaultsOK) GetPayload() *models.DefaultGreetingList {
	return o.Payload
}

func (o *GetUserGreetingsDefaultsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DefaultGreetingList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserGreetingsDefaultsBadRequest creates a GetUserGreetingsDefaultsBadRequest with default headers values
func NewGetUserGreetingsDefaultsBadRequest() *GetUserGreetingsDefaultsBadRequest {
	return &GetUserGreetingsDefaultsBadRequest{}
}

/*
GetUserGreetingsDefaultsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetUserGreetingsDefaultsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user greetings defaults bad request response has a 2xx status code
func (o *GetUserGreetingsDefaultsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user greetings defaults bad request response has a 3xx status code
func (o *GetUserGreetingsDefaultsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user greetings defaults bad request response has a 4xx status code
func (o *GetUserGreetingsDefaultsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user greetings defaults bad request response has a 5xx status code
func (o *GetUserGreetingsDefaultsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get user greetings defaults bad request response a status code equal to that given
func (o *GetUserGreetingsDefaultsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetUserGreetingsDefaultsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/greetings/defaults][%d] getUserGreetingsDefaultsBadRequest  %+v", 400, o.Payload)
}

func (o *GetUserGreetingsDefaultsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/greetings/defaults][%d] getUserGreetingsDefaultsBadRequest  %+v", 400, o.Payload)
}

func (o *GetUserGreetingsDefaultsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserGreetingsDefaultsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserGreetingsDefaultsUnauthorized creates a GetUserGreetingsDefaultsUnauthorized with default headers values
func NewGetUserGreetingsDefaultsUnauthorized() *GetUserGreetingsDefaultsUnauthorized {
	return &GetUserGreetingsDefaultsUnauthorized{}
}

/*
GetUserGreetingsDefaultsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetUserGreetingsDefaultsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user greetings defaults unauthorized response has a 2xx status code
func (o *GetUserGreetingsDefaultsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user greetings defaults unauthorized response has a 3xx status code
func (o *GetUserGreetingsDefaultsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user greetings defaults unauthorized response has a 4xx status code
func (o *GetUserGreetingsDefaultsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user greetings defaults unauthorized response has a 5xx status code
func (o *GetUserGreetingsDefaultsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get user greetings defaults unauthorized response a status code equal to that given
func (o *GetUserGreetingsDefaultsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetUserGreetingsDefaultsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/greetings/defaults][%d] getUserGreetingsDefaultsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetUserGreetingsDefaultsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/greetings/defaults][%d] getUserGreetingsDefaultsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetUserGreetingsDefaultsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserGreetingsDefaultsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserGreetingsDefaultsForbidden creates a GetUserGreetingsDefaultsForbidden with default headers values
func NewGetUserGreetingsDefaultsForbidden() *GetUserGreetingsDefaultsForbidden {
	return &GetUserGreetingsDefaultsForbidden{}
}

/*
GetUserGreetingsDefaultsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetUserGreetingsDefaultsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user greetings defaults forbidden response has a 2xx status code
func (o *GetUserGreetingsDefaultsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user greetings defaults forbidden response has a 3xx status code
func (o *GetUserGreetingsDefaultsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user greetings defaults forbidden response has a 4xx status code
func (o *GetUserGreetingsDefaultsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user greetings defaults forbidden response has a 5xx status code
func (o *GetUserGreetingsDefaultsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get user greetings defaults forbidden response a status code equal to that given
func (o *GetUserGreetingsDefaultsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetUserGreetingsDefaultsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/greetings/defaults][%d] getUserGreetingsDefaultsForbidden  %+v", 403, o.Payload)
}

func (o *GetUserGreetingsDefaultsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/greetings/defaults][%d] getUserGreetingsDefaultsForbidden  %+v", 403, o.Payload)
}

func (o *GetUserGreetingsDefaultsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserGreetingsDefaultsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserGreetingsDefaultsNotFound creates a GetUserGreetingsDefaultsNotFound with default headers values
func NewGetUserGreetingsDefaultsNotFound() *GetUserGreetingsDefaultsNotFound {
	return &GetUserGreetingsDefaultsNotFound{}
}

/*
GetUserGreetingsDefaultsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetUserGreetingsDefaultsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user greetings defaults not found response has a 2xx status code
func (o *GetUserGreetingsDefaultsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user greetings defaults not found response has a 3xx status code
func (o *GetUserGreetingsDefaultsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user greetings defaults not found response has a 4xx status code
func (o *GetUserGreetingsDefaultsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user greetings defaults not found response has a 5xx status code
func (o *GetUserGreetingsDefaultsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get user greetings defaults not found response a status code equal to that given
func (o *GetUserGreetingsDefaultsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetUserGreetingsDefaultsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/greetings/defaults][%d] getUserGreetingsDefaultsNotFound  %+v", 404, o.Payload)
}

func (o *GetUserGreetingsDefaultsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/greetings/defaults][%d] getUserGreetingsDefaultsNotFound  %+v", 404, o.Payload)
}

func (o *GetUserGreetingsDefaultsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserGreetingsDefaultsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserGreetingsDefaultsRequestTimeout creates a GetUserGreetingsDefaultsRequestTimeout with default headers values
func NewGetUserGreetingsDefaultsRequestTimeout() *GetUserGreetingsDefaultsRequestTimeout {
	return &GetUserGreetingsDefaultsRequestTimeout{}
}

/*
GetUserGreetingsDefaultsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetUserGreetingsDefaultsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user greetings defaults request timeout response has a 2xx status code
func (o *GetUserGreetingsDefaultsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user greetings defaults request timeout response has a 3xx status code
func (o *GetUserGreetingsDefaultsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user greetings defaults request timeout response has a 4xx status code
func (o *GetUserGreetingsDefaultsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user greetings defaults request timeout response has a 5xx status code
func (o *GetUserGreetingsDefaultsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get user greetings defaults request timeout response a status code equal to that given
func (o *GetUserGreetingsDefaultsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetUserGreetingsDefaultsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/greetings/defaults][%d] getUserGreetingsDefaultsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetUserGreetingsDefaultsRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/greetings/defaults][%d] getUserGreetingsDefaultsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetUserGreetingsDefaultsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserGreetingsDefaultsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserGreetingsDefaultsRequestEntityTooLarge creates a GetUserGreetingsDefaultsRequestEntityTooLarge with default headers values
func NewGetUserGreetingsDefaultsRequestEntityTooLarge() *GetUserGreetingsDefaultsRequestEntityTooLarge {
	return &GetUserGreetingsDefaultsRequestEntityTooLarge{}
}

/*
GetUserGreetingsDefaultsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetUserGreetingsDefaultsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user greetings defaults request entity too large response has a 2xx status code
func (o *GetUserGreetingsDefaultsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user greetings defaults request entity too large response has a 3xx status code
func (o *GetUserGreetingsDefaultsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user greetings defaults request entity too large response has a 4xx status code
func (o *GetUserGreetingsDefaultsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user greetings defaults request entity too large response has a 5xx status code
func (o *GetUserGreetingsDefaultsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get user greetings defaults request entity too large response a status code equal to that given
func (o *GetUserGreetingsDefaultsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetUserGreetingsDefaultsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/greetings/defaults][%d] getUserGreetingsDefaultsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetUserGreetingsDefaultsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/greetings/defaults][%d] getUserGreetingsDefaultsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetUserGreetingsDefaultsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserGreetingsDefaultsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserGreetingsDefaultsUnsupportedMediaType creates a GetUserGreetingsDefaultsUnsupportedMediaType with default headers values
func NewGetUserGreetingsDefaultsUnsupportedMediaType() *GetUserGreetingsDefaultsUnsupportedMediaType {
	return &GetUserGreetingsDefaultsUnsupportedMediaType{}
}

/*
GetUserGreetingsDefaultsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetUserGreetingsDefaultsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user greetings defaults unsupported media type response has a 2xx status code
func (o *GetUserGreetingsDefaultsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user greetings defaults unsupported media type response has a 3xx status code
func (o *GetUserGreetingsDefaultsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user greetings defaults unsupported media type response has a 4xx status code
func (o *GetUserGreetingsDefaultsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user greetings defaults unsupported media type response has a 5xx status code
func (o *GetUserGreetingsDefaultsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get user greetings defaults unsupported media type response a status code equal to that given
func (o *GetUserGreetingsDefaultsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetUserGreetingsDefaultsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/greetings/defaults][%d] getUserGreetingsDefaultsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetUserGreetingsDefaultsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/greetings/defaults][%d] getUserGreetingsDefaultsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetUserGreetingsDefaultsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserGreetingsDefaultsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserGreetingsDefaultsTooManyRequests creates a GetUserGreetingsDefaultsTooManyRequests with default headers values
func NewGetUserGreetingsDefaultsTooManyRequests() *GetUserGreetingsDefaultsTooManyRequests {
	return &GetUserGreetingsDefaultsTooManyRequests{}
}

/*
GetUserGreetingsDefaultsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetUserGreetingsDefaultsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user greetings defaults too many requests response has a 2xx status code
func (o *GetUserGreetingsDefaultsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user greetings defaults too many requests response has a 3xx status code
func (o *GetUserGreetingsDefaultsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user greetings defaults too many requests response has a 4xx status code
func (o *GetUserGreetingsDefaultsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user greetings defaults too many requests response has a 5xx status code
func (o *GetUserGreetingsDefaultsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get user greetings defaults too many requests response a status code equal to that given
func (o *GetUserGreetingsDefaultsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetUserGreetingsDefaultsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/greetings/defaults][%d] getUserGreetingsDefaultsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetUserGreetingsDefaultsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/greetings/defaults][%d] getUserGreetingsDefaultsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetUserGreetingsDefaultsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserGreetingsDefaultsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserGreetingsDefaultsInternalServerError creates a GetUserGreetingsDefaultsInternalServerError with default headers values
func NewGetUserGreetingsDefaultsInternalServerError() *GetUserGreetingsDefaultsInternalServerError {
	return &GetUserGreetingsDefaultsInternalServerError{}
}

/*
GetUserGreetingsDefaultsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetUserGreetingsDefaultsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user greetings defaults internal server error response has a 2xx status code
func (o *GetUserGreetingsDefaultsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user greetings defaults internal server error response has a 3xx status code
func (o *GetUserGreetingsDefaultsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user greetings defaults internal server error response has a 4xx status code
func (o *GetUserGreetingsDefaultsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get user greetings defaults internal server error response has a 5xx status code
func (o *GetUserGreetingsDefaultsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get user greetings defaults internal server error response a status code equal to that given
func (o *GetUserGreetingsDefaultsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetUserGreetingsDefaultsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/greetings/defaults][%d] getUserGreetingsDefaultsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUserGreetingsDefaultsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/greetings/defaults][%d] getUserGreetingsDefaultsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUserGreetingsDefaultsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserGreetingsDefaultsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserGreetingsDefaultsServiceUnavailable creates a GetUserGreetingsDefaultsServiceUnavailable with default headers values
func NewGetUserGreetingsDefaultsServiceUnavailable() *GetUserGreetingsDefaultsServiceUnavailable {
	return &GetUserGreetingsDefaultsServiceUnavailable{}
}

/*
GetUserGreetingsDefaultsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetUserGreetingsDefaultsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user greetings defaults service unavailable response has a 2xx status code
func (o *GetUserGreetingsDefaultsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user greetings defaults service unavailable response has a 3xx status code
func (o *GetUserGreetingsDefaultsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user greetings defaults service unavailable response has a 4xx status code
func (o *GetUserGreetingsDefaultsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get user greetings defaults service unavailable response has a 5xx status code
func (o *GetUserGreetingsDefaultsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get user greetings defaults service unavailable response a status code equal to that given
func (o *GetUserGreetingsDefaultsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetUserGreetingsDefaultsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/greetings/defaults][%d] getUserGreetingsDefaultsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetUserGreetingsDefaultsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/greetings/defaults][%d] getUserGreetingsDefaultsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetUserGreetingsDefaultsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserGreetingsDefaultsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserGreetingsDefaultsGatewayTimeout creates a GetUserGreetingsDefaultsGatewayTimeout with default headers values
func NewGetUserGreetingsDefaultsGatewayTimeout() *GetUserGreetingsDefaultsGatewayTimeout {
	return &GetUserGreetingsDefaultsGatewayTimeout{}
}

/*
GetUserGreetingsDefaultsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetUserGreetingsDefaultsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user greetings defaults gateway timeout response has a 2xx status code
func (o *GetUserGreetingsDefaultsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user greetings defaults gateway timeout response has a 3xx status code
func (o *GetUserGreetingsDefaultsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user greetings defaults gateway timeout response has a 4xx status code
func (o *GetUserGreetingsDefaultsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get user greetings defaults gateway timeout response has a 5xx status code
func (o *GetUserGreetingsDefaultsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get user greetings defaults gateway timeout response a status code equal to that given
func (o *GetUserGreetingsDefaultsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetUserGreetingsDefaultsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/greetings/defaults][%d] getUserGreetingsDefaultsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetUserGreetingsDefaultsGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/greetings/defaults][%d] getUserGreetingsDefaultsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetUserGreetingsDefaultsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserGreetingsDefaultsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

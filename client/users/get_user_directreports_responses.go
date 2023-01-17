// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetUserDirectreportsReader is a Reader for the GetUserDirectreports structure.
type GetUserDirectreportsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserDirectreportsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUserDirectreportsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetUserDirectreportsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetUserDirectreportsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetUserDirectreportsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetUserDirectreportsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetUserDirectreportsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetUserDirectreportsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetUserDirectreportsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetUserDirectreportsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUserDirectreportsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetUserDirectreportsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetUserDirectreportsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUserDirectreportsOK creates a GetUserDirectreportsOK with default headers values
func NewGetUserDirectreportsOK() *GetUserDirectreportsOK {
	return &GetUserDirectreportsOK{}
}

/*
GetUserDirectreportsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetUserDirectreportsOK struct {
	Payload []*models.User
}

// IsSuccess returns true when this get user directreports o k response has a 2xx status code
func (o *GetUserDirectreportsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get user directreports o k response has a 3xx status code
func (o *GetUserDirectreportsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user directreports o k response has a 4xx status code
func (o *GetUserDirectreportsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get user directreports o k response has a 5xx status code
func (o *GetUserDirectreportsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get user directreports o k response a status code equal to that given
func (o *GetUserDirectreportsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetUserDirectreportsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/directreports][%d] getUserDirectreportsOK  %+v", 200, o.Payload)
}

func (o *GetUserDirectreportsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/directreports][%d] getUserDirectreportsOK  %+v", 200, o.Payload)
}

func (o *GetUserDirectreportsOK) GetPayload() []*models.User {
	return o.Payload
}

func (o *GetUserDirectreportsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserDirectreportsBadRequest creates a GetUserDirectreportsBadRequest with default headers values
func NewGetUserDirectreportsBadRequest() *GetUserDirectreportsBadRequest {
	return &GetUserDirectreportsBadRequest{}
}

/*
GetUserDirectreportsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetUserDirectreportsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user directreports bad request response has a 2xx status code
func (o *GetUserDirectreportsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user directreports bad request response has a 3xx status code
func (o *GetUserDirectreportsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user directreports bad request response has a 4xx status code
func (o *GetUserDirectreportsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user directreports bad request response has a 5xx status code
func (o *GetUserDirectreportsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get user directreports bad request response a status code equal to that given
func (o *GetUserDirectreportsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetUserDirectreportsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/directreports][%d] getUserDirectreportsBadRequest  %+v", 400, o.Payload)
}

func (o *GetUserDirectreportsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/directreports][%d] getUserDirectreportsBadRequest  %+v", 400, o.Payload)
}

func (o *GetUserDirectreportsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserDirectreportsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserDirectreportsUnauthorized creates a GetUserDirectreportsUnauthorized with default headers values
func NewGetUserDirectreportsUnauthorized() *GetUserDirectreportsUnauthorized {
	return &GetUserDirectreportsUnauthorized{}
}

/*
GetUserDirectreportsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetUserDirectreportsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user directreports unauthorized response has a 2xx status code
func (o *GetUserDirectreportsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user directreports unauthorized response has a 3xx status code
func (o *GetUserDirectreportsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user directreports unauthorized response has a 4xx status code
func (o *GetUserDirectreportsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user directreports unauthorized response has a 5xx status code
func (o *GetUserDirectreportsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get user directreports unauthorized response a status code equal to that given
func (o *GetUserDirectreportsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetUserDirectreportsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/directreports][%d] getUserDirectreportsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetUserDirectreportsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/directreports][%d] getUserDirectreportsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetUserDirectreportsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserDirectreportsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserDirectreportsForbidden creates a GetUserDirectreportsForbidden with default headers values
func NewGetUserDirectreportsForbidden() *GetUserDirectreportsForbidden {
	return &GetUserDirectreportsForbidden{}
}

/*
GetUserDirectreportsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetUserDirectreportsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user directreports forbidden response has a 2xx status code
func (o *GetUserDirectreportsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user directreports forbidden response has a 3xx status code
func (o *GetUserDirectreportsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user directreports forbidden response has a 4xx status code
func (o *GetUserDirectreportsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user directreports forbidden response has a 5xx status code
func (o *GetUserDirectreportsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get user directreports forbidden response a status code equal to that given
func (o *GetUserDirectreportsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetUserDirectreportsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/directreports][%d] getUserDirectreportsForbidden  %+v", 403, o.Payload)
}

func (o *GetUserDirectreportsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/directreports][%d] getUserDirectreportsForbidden  %+v", 403, o.Payload)
}

func (o *GetUserDirectreportsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserDirectreportsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserDirectreportsNotFound creates a GetUserDirectreportsNotFound with default headers values
func NewGetUserDirectreportsNotFound() *GetUserDirectreportsNotFound {
	return &GetUserDirectreportsNotFound{}
}

/*
GetUserDirectreportsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetUserDirectreportsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user directreports not found response has a 2xx status code
func (o *GetUserDirectreportsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user directreports not found response has a 3xx status code
func (o *GetUserDirectreportsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user directreports not found response has a 4xx status code
func (o *GetUserDirectreportsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user directreports not found response has a 5xx status code
func (o *GetUserDirectreportsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get user directreports not found response a status code equal to that given
func (o *GetUserDirectreportsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetUserDirectreportsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/directreports][%d] getUserDirectreportsNotFound  %+v", 404, o.Payload)
}

func (o *GetUserDirectreportsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/directreports][%d] getUserDirectreportsNotFound  %+v", 404, o.Payload)
}

func (o *GetUserDirectreportsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserDirectreportsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserDirectreportsRequestTimeout creates a GetUserDirectreportsRequestTimeout with default headers values
func NewGetUserDirectreportsRequestTimeout() *GetUserDirectreportsRequestTimeout {
	return &GetUserDirectreportsRequestTimeout{}
}

/*
GetUserDirectreportsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetUserDirectreportsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user directreports request timeout response has a 2xx status code
func (o *GetUserDirectreportsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user directreports request timeout response has a 3xx status code
func (o *GetUserDirectreportsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user directreports request timeout response has a 4xx status code
func (o *GetUserDirectreportsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user directreports request timeout response has a 5xx status code
func (o *GetUserDirectreportsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get user directreports request timeout response a status code equal to that given
func (o *GetUserDirectreportsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetUserDirectreportsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/directreports][%d] getUserDirectreportsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetUserDirectreportsRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/directreports][%d] getUserDirectreportsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetUserDirectreportsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserDirectreportsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserDirectreportsRequestEntityTooLarge creates a GetUserDirectreportsRequestEntityTooLarge with default headers values
func NewGetUserDirectreportsRequestEntityTooLarge() *GetUserDirectreportsRequestEntityTooLarge {
	return &GetUserDirectreportsRequestEntityTooLarge{}
}

/*
GetUserDirectreportsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetUserDirectreportsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user directreports request entity too large response has a 2xx status code
func (o *GetUserDirectreportsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user directreports request entity too large response has a 3xx status code
func (o *GetUserDirectreportsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user directreports request entity too large response has a 4xx status code
func (o *GetUserDirectreportsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user directreports request entity too large response has a 5xx status code
func (o *GetUserDirectreportsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get user directreports request entity too large response a status code equal to that given
func (o *GetUserDirectreportsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetUserDirectreportsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/directreports][%d] getUserDirectreportsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetUserDirectreportsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/directreports][%d] getUserDirectreportsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetUserDirectreportsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserDirectreportsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserDirectreportsUnsupportedMediaType creates a GetUserDirectreportsUnsupportedMediaType with default headers values
func NewGetUserDirectreportsUnsupportedMediaType() *GetUserDirectreportsUnsupportedMediaType {
	return &GetUserDirectreportsUnsupportedMediaType{}
}

/*
GetUserDirectreportsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetUserDirectreportsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user directreports unsupported media type response has a 2xx status code
func (o *GetUserDirectreportsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user directreports unsupported media type response has a 3xx status code
func (o *GetUserDirectreportsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user directreports unsupported media type response has a 4xx status code
func (o *GetUserDirectreportsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user directreports unsupported media type response has a 5xx status code
func (o *GetUserDirectreportsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get user directreports unsupported media type response a status code equal to that given
func (o *GetUserDirectreportsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetUserDirectreportsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/directreports][%d] getUserDirectreportsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetUserDirectreportsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/directreports][%d] getUserDirectreportsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetUserDirectreportsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserDirectreportsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserDirectreportsTooManyRequests creates a GetUserDirectreportsTooManyRequests with default headers values
func NewGetUserDirectreportsTooManyRequests() *GetUserDirectreportsTooManyRequests {
	return &GetUserDirectreportsTooManyRequests{}
}

/*
GetUserDirectreportsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetUserDirectreportsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user directreports too many requests response has a 2xx status code
func (o *GetUserDirectreportsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user directreports too many requests response has a 3xx status code
func (o *GetUserDirectreportsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user directreports too many requests response has a 4xx status code
func (o *GetUserDirectreportsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user directreports too many requests response has a 5xx status code
func (o *GetUserDirectreportsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get user directreports too many requests response a status code equal to that given
func (o *GetUserDirectreportsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetUserDirectreportsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/directreports][%d] getUserDirectreportsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetUserDirectreportsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/directreports][%d] getUserDirectreportsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetUserDirectreportsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserDirectreportsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserDirectreportsInternalServerError creates a GetUserDirectreportsInternalServerError with default headers values
func NewGetUserDirectreportsInternalServerError() *GetUserDirectreportsInternalServerError {
	return &GetUserDirectreportsInternalServerError{}
}

/*
GetUserDirectreportsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetUserDirectreportsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user directreports internal server error response has a 2xx status code
func (o *GetUserDirectreportsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user directreports internal server error response has a 3xx status code
func (o *GetUserDirectreportsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user directreports internal server error response has a 4xx status code
func (o *GetUserDirectreportsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get user directreports internal server error response has a 5xx status code
func (o *GetUserDirectreportsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get user directreports internal server error response a status code equal to that given
func (o *GetUserDirectreportsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetUserDirectreportsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/directreports][%d] getUserDirectreportsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUserDirectreportsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/directreports][%d] getUserDirectreportsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUserDirectreportsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserDirectreportsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserDirectreportsServiceUnavailable creates a GetUserDirectreportsServiceUnavailable with default headers values
func NewGetUserDirectreportsServiceUnavailable() *GetUserDirectreportsServiceUnavailable {
	return &GetUserDirectreportsServiceUnavailable{}
}

/*
GetUserDirectreportsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetUserDirectreportsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user directreports service unavailable response has a 2xx status code
func (o *GetUserDirectreportsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user directreports service unavailable response has a 3xx status code
func (o *GetUserDirectreportsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user directreports service unavailable response has a 4xx status code
func (o *GetUserDirectreportsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get user directreports service unavailable response has a 5xx status code
func (o *GetUserDirectreportsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get user directreports service unavailable response a status code equal to that given
func (o *GetUserDirectreportsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetUserDirectreportsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/directreports][%d] getUserDirectreportsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetUserDirectreportsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/directreports][%d] getUserDirectreportsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetUserDirectreportsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserDirectreportsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserDirectreportsGatewayTimeout creates a GetUserDirectreportsGatewayTimeout with default headers values
func NewGetUserDirectreportsGatewayTimeout() *GetUserDirectreportsGatewayTimeout {
	return &GetUserDirectreportsGatewayTimeout{}
}

/*
GetUserDirectreportsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetUserDirectreportsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user directreports gateway timeout response has a 2xx status code
func (o *GetUserDirectreportsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user directreports gateway timeout response has a 3xx status code
func (o *GetUserDirectreportsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user directreports gateway timeout response has a 4xx status code
func (o *GetUserDirectreportsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get user directreports gateway timeout response has a 5xx status code
func (o *GetUserDirectreportsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get user directreports gateway timeout response a status code equal to that given
func (o *GetUserDirectreportsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetUserDirectreportsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/directreports][%d] getUserDirectreportsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetUserDirectreportsGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/directreports][%d] getUserDirectreportsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetUserDirectreportsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserDirectreportsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

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

// GetUserGreetingsReader is a Reader for the GetUserGreetings structure.
type GetUserGreetingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserGreetingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUserGreetingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetUserGreetingsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetUserGreetingsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetUserGreetingsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetUserGreetingsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetUserGreetingsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetUserGreetingsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetUserGreetingsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetUserGreetingsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUserGreetingsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetUserGreetingsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetUserGreetingsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUserGreetingsOK creates a GetUserGreetingsOK with default headers values
func NewGetUserGreetingsOK() *GetUserGreetingsOK {
	return &GetUserGreetingsOK{}
}

/*
GetUserGreetingsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetUserGreetingsOK struct {
	Payload *models.DomainEntityListing
}

// IsSuccess returns true when this get user greetings o k response has a 2xx status code
func (o *GetUserGreetingsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get user greetings o k response has a 3xx status code
func (o *GetUserGreetingsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user greetings o k response has a 4xx status code
func (o *GetUserGreetingsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get user greetings o k response has a 5xx status code
func (o *GetUserGreetingsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get user greetings o k response a status code equal to that given
func (o *GetUserGreetingsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetUserGreetingsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/greetings][%d] getUserGreetingsOK  %+v", 200, o.Payload)
}

func (o *GetUserGreetingsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/greetings][%d] getUserGreetingsOK  %+v", 200, o.Payload)
}

func (o *GetUserGreetingsOK) GetPayload() *models.DomainEntityListing {
	return o.Payload
}

func (o *GetUserGreetingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DomainEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserGreetingsBadRequest creates a GetUserGreetingsBadRequest with default headers values
func NewGetUserGreetingsBadRequest() *GetUserGreetingsBadRequest {
	return &GetUserGreetingsBadRequest{}
}

/*
GetUserGreetingsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetUserGreetingsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user greetings bad request response has a 2xx status code
func (o *GetUserGreetingsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user greetings bad request response has a 3xx status code
func (o *GetUserGreetingsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user greetings bad request response has a 4xx status code
func (o *GetUserGreetingsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user greetings bad request response has a 5xx status code
func (o *GetUserGreetingsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get user greetings bad request response a status code equal to that given
func (o *GetUserGreetingsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetUserGreetingsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/greetings][%d] getUserGreetingsBadRequest  %+v", 400, o.Payload)
}

func (o *GetUserGreetingsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/greetings][%d] getUserGreetingsBadRequest  %+v", 400, o.Payload)
}

func (o *GetUserGreetingsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserGreetingsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserGreetingsUnauthorized creates a GetUserGreetingsUnauthorized with default headers values
func NewGetUserGreetingsUnauthorized() *GetUserGreetingsUnauthorized {
	return &GetUserGreetingsUnauthorized{}
}

/*
GetUserGreetingsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetUserGreetingsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user greetings unauthorized response has a 2xx status code
func (o *GetUserGreetingsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user greetings unauthorized response has a 3xx status code
func (o *GetUserGreetingsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user greetings unauthorized response has a 4xx status code
func (o *GetUserGreetingsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user greetings unauthorized response has a 5xx status code
func (o *GetUserGreetingsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get user greetings unauthorized response a status code equal to that given
func (o *GetUserGreetingsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetUserGreetingsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/greetings][%d] getUserGreetingsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetUserGreetingsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/greetings][%d] getUserGreetingsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetUserGreetingsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserGreetingsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserGreetingsForbidden creates a GetUserGreetingsForbidden with default headers values
func NewGetUserGreetingsForbidden() *GetUserGreetingsForbidden {
	return &GetUserGreetingsForbidden{}
}

/*
GetUserGreetingsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetUserGreetingsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user greetings forbidden response has a 2xx status code
func (o *GetUserGreetingsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user greetings forbidden response has a 3xx status code
func (o *GetUserGreetingsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user greetings forbidden response has a 4xx status code
func (o *GetUserGreetingsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user greetings forbidden response has a 5xx status code
func (o *GetUserGreetingsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get user greetings forbidden response a status code equal to that given
func (o *GetUserGreetingsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetUserGreetingsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/greetings][%d] getUserGreetingsForbidden  %+v", 403, o.Payload)
}

func (o *GetUserGreetingsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/greetings][%d] getUserGreetingsForbidden  %+v", 403, o.Payload)
}

func (o *GetUserGreetingsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserGreetingsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserGreetingsNotFound creates a GetUserGreetingsNotFound with default headers values
func NewGetUserGreetingsNotFound() *GetUserGreetingsNotFound {
	return &GetUserGreetingsNotFound{}
}

/*
GetUserGreetingsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetUserGreetingsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user greetings not found response has a 2xx status code
func (o *GetUserGreetingsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user greetings not found response has a 3xx status code
func (o *GetUserGreetingsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user greetings not found response has a 4xx status code
func (o *GetUserGreetingsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user greetings not found response has a 5xx status code
func (o *GetUserGreetingsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get user greetings not found response a status code equal to that given
func (o *GetUserGreetingsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetUserGreetingsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/greetings][%d] getUserGreetingsNotFound  %+v", 404, o.Payload)
}

func (o *GetUserGreetingsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/greetings][%d] getUserGreetingsNotFound  %+v", 404, o.Payload)
}

func (o *GetUserGreetingsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserGreetingsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserGreetingsRequestTimeout creates a GetUserGreetingsRequestTimeout with default headers values
func NewGetUserGreetingsRequestTimeout() *GetUserGreetingsRequestTimeout {
	return &GetUserGreetingsRequestTimeout{}
}

/*
GetUserGreetingsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetUserGreetingsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user greetings request timeout response has a 2xx status code
func (o *GetUserGreetingsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user greetings request timeout response has a 3xx status code
func (o *GetUserGreetingsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user greetings request timeout response has a 4xx status code
func (o *GetUserGreetingsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user greetings request timeout response has a 5xx status code
func (o *GetUserGreetingsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get user greetings request timeout response a status code equal to that given
func (o *GetUserGreetingsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetUserGreetingsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/greetings][%d] getUserGreetingsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetUserGreetingsRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/greetings][%d] getUserGreetingsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetUserGreetingsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserGreetingsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserGreetingsRequestEntityTooLarge creates a GetUserGreetingsRequestEntityTooLarge with default headers values
func NewGetUserGreetingsRequestEntityTooLarge() *GetUserGreetingsRequestEntityTooLarge {
	return &GetUserGreetingsRequestEntityTooLarge{}
}

/*
GetUserGreetingsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetUserGreetingsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user greetings request entity too large response has a 2xx status code
func (o *GetUserGreetingsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user greetings request entity too large response has a 3xx status code
func (o *GetUserGreetingsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user greetings request entity too large response has a 4xx status code
func (o *GetUserGreetingsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user greetings request entity too large response has a 5xx status code
func (o *GetUserGreetingsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get user greetings request entity too large response a status code equal to that given
func (o *GetUserGreetingsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetUserGreetingsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/greetings][%d] getUserGreetingsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetUserGreetingsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/greetings][%d] getUserGreetingsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetUserGreetingsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserGreetingsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserGreetingsUnsupportedMediaType creates a GetUserGreetingsUnsupportedMediaType with default headers values
func NewGetUserGreetingsUnsupportedMediaType() *GetUserGreetingsUnsupportedMediaType {
	return &GetUserGreetingsUnsupportedMediaType{}
}

/*
GetUserGreetingsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetUserGreetingsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user greetings unsupported media type response has a 2xx status code
func (o *GetUserGreetingsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user greetings unsupported media type response has a 3xx status code
func (o *GetUserGreetingsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user greetings unsupported media type response has a 4xx status code
func (o *GetUserGreetingsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user greetings unsupported media type response has a 5xx status code
func (o *GetUserGreetingsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get user greetings unsupported media type response a status code equal to that given
func (o *GetUserGreetingsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetUserGreetingsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/greetings][%d] getUserGreetingsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetUserGreetingsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/greetings][%d] getUserGreetingsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetUserGreetingsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserGreetingsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserGreetingsTooManyRequests creates a GetUserGreetingsTooManyRequests with default headers values
func NewGetUserGreetingsTooManyRequests() *GetUserGreetingsTooManyRequests {
	return &GetUserGreetingsTooManyRequests{}
}

/*
GetUserGreetingsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetUserGreetingsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user greetings too many requests response has a 2xx status code
func (o *GetUserGreetingsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user greetings too many requests response has a 3xx status code
func (o *GetUserGreetingsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user greetings too many requests response has a 4xx status code
func (o *GetUserGreetingsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user greetings too many requests response has a 5xx status code
func (o *GetUserGreetingsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get user greetings too many requests response a status code equal to that given
func (o *GetUserGreetingsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetUserGreetingsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/greetings][%d] getUserGreetingsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetUserGreetingsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/greetings][%d] getUserGreetingsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetUserGreetingsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserGreetingsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserGreetingsInternalServerError creates a GetUserGreetingsInternalServerError with default headers values
func NewGetUserGreetingsInternalServerError() *GetUserGreetingsInternalServerError {
	return &GetUserGreetingsInternalServerError{}
}

/*
GetUserGreetingsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetUserGreetingsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user greetings internal server error response has a 2xx status code
func (o *GetUserGreetingsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user greetings internal server error response has a 3xx status code
func (o *GetUserGreetingsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user greetings internal server error response has a 4xx status code
func (o *GetUserGreetingsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get user greetings internal server error response has a 5xx status code
func (o *GetUserGreetingsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get user greetings internal server error response a status code equal to that given
func (o *GetUserGreetingsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetUserGreetingsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/greetings][%d] getUserGreetingsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUserGreetingsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/greetings][%d] getUserGreetingsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUserGreetingsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserGreetingsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserGreetingsServiceUnavailable creates a GetUserGreetingsServiceUnavailable with default headers values
func NewGetUserGreetingsServiceUnavailable() *GetUserGreetingsServiceUnavailable {
	return &GetUserGreetingsServiceUnavailable{}
}

/*
GetUserGreetingsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetUserGreetingsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user greetings service unavailable response has a 2xx status code
func (o *GetUserGreetingsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user greetings service unavailable response has a 3xx status code
func (o *GetUserGreetingsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user greetings service unavailable response has a 4xx status code
func (o *GetUserGreetingsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get user greetings service unavailable response has a 5xx status code
func (o *GetUserGreetingsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get user greetings service unavailable response a status code equal to that given
func (o *GetUserGreetingsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetUserGreetingsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/greetings][%d] getUserGreetingsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetUserGreetingsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/greetings][%d] getUserGreetingsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetUserGreetingsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserGreetingsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserGreetingsGatewayTimeout creates a GetUserGreetingsGatewayTimeout with default headers values
func NewGetUserGreetingsGatewayTimeout() *GetUserGreetingsGatewayTimeout {
	return &GetUserGreetingsGatewayTimeout{}
}

/*
GetUserGreetingsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetUserGreetingsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user greetings gateway timeout response has a 2xx status code
func (o *GetUserGreetingsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user greetings gateway timeout response has a 3xx status code
func (o *GetUserGreetingsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user greetings gateway timeout response has a 4xx status code
func (o *GetUserGreetingsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get user greetings gateway timeout response has a 5xx status code
func (o *GetUserGreetingsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get user greetings gateway timeout response a status code equal to that given
func (o *GetUserGreetingsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetUserGreetingsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/greetings][%d] getUserGreetingsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetUserGreetingsGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/greetings][%d] getUserGreetingsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetUserGreetingsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserGreetingsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

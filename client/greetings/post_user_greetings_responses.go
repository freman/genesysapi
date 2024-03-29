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

// PostUserGreetingsReader is a Reader for the PostUserGreetings structure.
type PostUserGreetingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostUserGreetingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostUserGreetingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostUserGreetingsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostUserGreetingsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostUserGreetingsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostUserGreetingsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostUserGreetingsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostUserGreetingsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostUserGreetingsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostUserGreetingsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostUserGreetingsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostUserGreetingsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostUserGreetingsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostUserGreetingsOK creates a PostUserGreetingsOK with default headers values
func NewPostUserGreetingsOK() *PostUserGreetingsOK {
	return &PostUserGreetingsOK{}
}

/*
PostUserGreetingsOK describes a response with status code 200, with default header values.

successful operation
*/
type PostUserGreetingsOK struct {
	Payload *models.Greeting
}

// IsSuccess returns true when this post user greetings o k response has a 2xx status code
func (o *PostUserGreetingsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post user greetings o k response has a 3xx status code
func (o *PostUserGreetingsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post user greetings o k response has a 4xx status code
func (o *PostUserGreetingsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post user greetings o k response has a 5xx status code
func (o *PostUserGreetingsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post user greetings o k response a status code equal to that given
func (o *PostUserGreetingsOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostUserGreetingsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/greetings][%d] postUserGreetingsOK  %+v", 200, o.Payload)
}

func (o *PostUserGreetingsOK) String() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/greetings][%d] postUserGreetingsOK  %+v", 200, o.Payload)
}

func (o *PostUserGreetingsOK) GetPayload() *models.Greeting {
	return o.Payload
}

func (o *PostUserGreetingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Greeting)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUserGreetingsBadRequest creates a PostUserGreetingsBadRequest with default headers values
func NewPostUserGreetingsBadRequest() *PostUserGreetingsBadRequest {
	return &PostUserGreetingsBadRequest{}
}

/*
PostUserGreetingsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostUserGreetingsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post user greetings bad request response has a 2xx status code
func (o *PostUserGreetingsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post user greetings bad request response has a 3xx status code
func (o *PostUserGreetingsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post user greetings bad request response has a 4xx status code
func (o *PostUserGreetingsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post user greetings bad request response has a 5xx status code
func (o *PostUserGreetingsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post user greetings bad request response a status code equal to that given
func (o *PostUserGreetingsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostUserGreetingsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/greetings][%d] postUserGreetingsBadRequest  %+v", 400, o.Payload)
}

func (o *PostUserGreetingsBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/greetings][%d] postUserGreetingsBadRequest  %+v", 400, o.Payload)
}

func (o *PostUserGreetingsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUserGreetingsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUserGreetingsUnauthorized creates a PostUserGreetingsUnauthorized with default headers values
func NewPostUserGreetingsUnauthorized() *PostUserGreetingsUnauthorized {
	return &PostUserGreetingsUnauthorized{}
}

/*
PostUserGreetingsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostUserGreetingsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post user greetings unauthorized response has a 2xx status code
func (o *PostUserGreetingsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post user greetings unauthorized response has a 3xx status code
func (o *PostUserGreetingsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post user greetings unauthorized response has a 4xx status code
func (o *PostUserGreetingsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post user greetings unauthorized response has a 5xx status code
func (o *PostUserGreetingsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post user greetings unauthorized response a status code equal to that given
func (o *PostUserGreetingsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostUserGreetingsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/greetings][%d] postUserGreetingsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostUserGreetingsUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/greetings][%d] postUserGreetingsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostUserGreetingsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUserGreetingsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUserGreetingsForbidden creates a PostUserGreetingsForbidden with default headers values
func NewPostUserGreetingsForbidden() *PostUserGreetingsForbidden {
	return &PostUserGreetingsForbidden{}
}

/*
PostUserGreetingsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostUserGreetingsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post user greetings forbidden response has a 2xx status code
func (o *PostUserGreetingsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post user greetings forbidden response has a 3xx status code
func (o *PostUserGreetingsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post user greetings forbidden response has a 4xx status code
func (o *PostUserGreetingsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post user greetings forbidden response has a 5xx status code
func (o *PostUserGreetingsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post user greetings forbidden response a status code equal to that given
func (o *PostUserGreetingsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostUserGreetingsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/greetings][%d] postUserGreetingsForbidden  %+v", 403, o.Payload)
}

func (o *PostUserGreetingsForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/greetings][%d] postUserGreetingsForbidden  %+v", 403, o.Payload)
}

func (o *PostUserGreetingsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUserGreetingsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUserGreetingsNotFound creates a PostUserGreetingsNotFound with default headers values
func NewPostUserGreetingsNotFound() *PostUserGreetingsNotFound {
	return &PostUserGreetingsNotFound{}
}

/*
PostUserGreetingsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostUserGreetingsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post user greetings not found response has a 2xx status code
func (o *PostUserGreetingsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post user greetings not found response has a 3xx status code
func (o *PostUserGreetingsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post user greetings not found response has a 4xx status code
func (o *PostUserGreetingsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post user greetings not found response has a 5xx status code
func (o *PostUserGreetingsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post user greetings not found response a status code equal to that given
func (o *PostUserGreetingsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostUserGreetingsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/greetings][%d] postUserGreetingsNotFound  %+v", 404, o.Payload)
}

func (o *PostUserGreetingsNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/greetings][%d] postUserGreetingsNotFound  %+v", 404, o.Payload)
}

func (o *PostUserGreetingsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUserGreetingsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUserGreetingsRequestTimeout creates a PostUserGreetingsRequestTimeout with default headers values
func NewPostUserGreetingsRequestTimeout() *PostUserGreetingsRequestTimeout {
	return &PostUserGreetingsRequestTimeout{}
}

/*
PostUserGreetingsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostUserGreetingsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post user greetings request timeout response has a 2xx status code
func (o *PostUserGreetingsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post user greetings request timeout response has a 3xx status code
func (o *PostUserGreetingsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post user greetings request timeout response has a 4xx status code
func (o *PostUserGreetingsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post user greetings request timeout response has a 5xx status code
func (o *PostUserGreetingsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post user greetings request timeout response a status code equal to that given
func (o *PostUserGreetingsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostUserGreetingsRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/greetings][%d] postUserGreetingsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostUserGreetingsRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/greetings][%d] postUserGreetingsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostUserGreetingsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUserGreetingsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUserGreetingsRequestEntityTooLarge creates a PostUserGreetingsRequestEntityTooLarge with default headers values
func NewPostUserGreetingsRequestEntityTooLarge() *PostUserGreetingsRequestEntityTooLarge {
	return &PostUserGreetingsRequestEntityTooLarge{}
}

/*
PostUserGreetingsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PostUserGreetingsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post user greetings request entity too large response has a 2xx status code
func (o *PostUserGreetingsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post user greetings request entity too large response has a 3xx status code
func (o *PostUserGreetingsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post user greetings request entity too large response has a 4xx status code
func (o *PostUserGreetingsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post user greetings request entity too large response has a 5xx status code
func (o *PostUserGreetingsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post user greetings request entity too large response a status code equal to that given
func (o *PostUserGreetingsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostUserGreetingsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/greetings][%d] postUserGreetingsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostUserGreetingsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/greetings][%d] postUserGreetingsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostUserGreetingsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUserGreetingsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUserGreetingsUnsupportedMediaType creates a PostUserGreetingsUnsupportedMediaType with default headers values
func NewPostUserGreetingsUnsupportedMediaType() *PostUserGreetingsUnsupportedMediaType {
	return &PostUserGreetingsUnsupportedMediaType{}
}

/*
PostUserGreetingsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostUserGreetingsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post user greetings unsupported media type response has a 2xx status code
func (o *PostUserGreetingsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post user greetings unsupported media type response has a 3xx status code
func (o *PostUserGreetingsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post user greetings unsupported media type response has a 4xx status code
func (o *PostUserGreetingsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post user greetings unsupported media type response has a 5xx status code
func (o *PostUserGreetingsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post user greetings unsupported media type response a status code equal to that given
func (o *PostUserGreetingsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostUserGreetingsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/greetings][%d] postUserGreetingsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostUserGreetingsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/greetings][%d] postUserGreetingsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostUserGreetingsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUserGreetingsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUserGreetingsTooManyRequests creates a PostUserGreetingsTooManyRequests with default headers values
func NewPostUserGreetingsTooManyRequests() *PostUserGreetingsTooManyRequests {
	return &PostUserGreetingsTooManyRequests{}
}

/*
PostUserGreetingsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostUserGreetingsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post user greetings too many requests response has a 2xx status code
func (o *PostUserGreetingsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post user greetings too many requests response has a 3xx status code
func (o *PostUserGreetingsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post user greetings too many requests response has a 4xx status code
func (o *PostUserGreetingsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post user greetings too many requests response has a 5xx status code
func (o *PostUserGreetingsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post user greetings too many requests response a status code equal to that given
func (o *PostUserGreetingsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostUserGreetingsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/greetings][%d] postUserGreetingsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostUserGreetingsTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/greetings][%d] postUserGreetingsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostUserGreetingsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUserGreetingsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUserGreetingsInternalServerError creates a PostUserGreetingsInternalServerError with default headers values
func NewPostUserGreetingsInternalServerError() *PostUserGreetingsInternalServerError {
	return &PostUserGreetingsInternalServerError{}
}

/*
PostUserGreetingsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostUserGreetingsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post user greetings internal server error response has a 2xx status code
func (o *PostUserGreetingsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post user greetings internal server error response has a 3xx status code
func (o *PostUserGreetingsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post user greetings internal server error response has a 4xx status code
func (o *PostUserGreetingsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post user greetings internal server error response has a 5xx status code
func (o *PostUserGreetingsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post user greetings internal server error response a status code equal to that given
func (o *PostUserGreetingsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostUserGreetingsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/greetings][%d] postUserGreetingsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostUserGreetingsInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/greetings][%d] postUserGreetingsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostUserGreetingsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUserGreetingsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUserGreetingsServiceUnavailable creates a PostUserGreetingsServiceUnavailable with default headers values
func NewPostUserGreetingsServiceUnavailable() *PostUserGreetingsServiceUnavailable {
	return &PostUserGreetingsServiceUnavailable{}
}

/*
PostUserGreetingsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostUserGreetingsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post user greetings service unavailable response has a 2xx status code
func (o *PostUserGreetingsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post user greetings service unavailable response has a 3xx status code
func (o *PostUserGreetingsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post user greetings service unavailable response has a 4xx status code
func (o *PostUserGreetingsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post user greetings service unavailable response has a 5xx status code
func (o *PostUserGreetingsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post user greetings service unavailable response a status code equal to that given
func (o *PostUserGreetingsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostUserGreetingsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/greetings][%d] postUserGreetingsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostUserGreetingsServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/greetings][%d] postUserGreetingsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostUserGreetingsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUserGreetingsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUserGreetingsGatewayTimeout creates a PostUserGreetingsGatewayTimeout with default headers values
func NewPostUserGreetingsGatewayTimeout() *PostUserGreetingsGatewayTimeout {
	return &PostUserGreetingsGatewayTimeout{}
}

/*
PostUserGreetingsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostUserGreetingsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post user greetings gateway timeout response has a 2xx status code
func (o *PostUserGreetingsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post user greetings gateway timeout response has a 3xx status code
func (o *PostUserGreetingsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post user greetings gateway timeout response has a 4xx status code
func (o *PostUserGreetingsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post user greetings gateway timeout response has a 5xx status code
func (o *PostUserGreetingsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post user greetings gateway timeout response a status code equal to that given
func (o *PostUserGreetingsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostUserGreetingsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/greetings][%d] postUserGreetingsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostUserGreetingsGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/greetings][%d] postUserGreetingsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostUserGreetingsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUserGreetingsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

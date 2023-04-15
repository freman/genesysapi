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

// PostGroupGreetingsReader is a Reader for the PostGroupGreetings structure.
type PostGroupGreetingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostGroupGreetingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostGroupGreetingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostGroupGreetingsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostGroupGreetingsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostGroupGreetingsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostGroupGreetingsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostGroupGreetingsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostGroupGreetingsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostGroupGreetingsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostGroupGreetingsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostGroupGreetingsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostGroupGreetingsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostGroupGreetingsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostGroupGreetingsOK creates a PostGroupGreetingsOK with default headers values
func NewPostGroupGreetingsOK() *PostGroupGreetingsOK {
	return &PostGroupGreetingsOK{}
}

/*
PostGroupGreetingsOK describes a response with status code 200, with default header values.

successful operation
*/
type PostGroupGreetingsOK struct {
	Payload *models.Greeting
}

// IsSuccess returns true when this post group greetings o k response has a 2xx status code
func (o *PostGroupGreetingsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post group greetings o k response has a 3xx status code
func (o *PostGroupGreetingsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post group greetings o k response has a 4xx status code
func (o *PostGroupGreetingsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post group greetings o k response has a 5xx status code
func (o *PostGroupGreetingsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post group greetings o k response a status code equal to that given
func (o *PostGroupGreetingsOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostGroupGreetingsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/groups/{groupId}/greetings][%d] postGroupGreetingsOK  %+v", 200, o.Payload)
}

func (o *PostGroupGreetingsOK) String() string {
	return fmt.Sprintf("[POST /api/v2/groups/{groupId}/greetings][%d] postGroupGreetingsOK  %+v", 200, o.Payload)
}

func (o *PostGroupGreetingsOK) GetPayload() *models.Greeting {
	return o.Payload
}

func (o *PostGroupGreetingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Greeting)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGroupGreetingsBadRequest creates a PostGroupGreetingsBadRequest with default headers values
func NewPostGroupGreetingsBadRequest() *PostGroupGreetingsBadRequest {
	return &PostGroupGreetingsBadRequest{}
}

/*
PostGroupGreetingsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostGroupGreetingsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post group greetings bad request response has a 2xx status code
func (o *PostGroupGreetingsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post group greetings bad request response has a 3xx status code
func (o *PostGroupGreetingsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post group greetings bad request response has a 4xx status code
func (o *PostGroupGreetingsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post group greetings bad request response has a 5xx status code
func (o *PostGroupGreetingsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post group greetings bad request response a status code equal to that given
func (o *PostGroupGreetingsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostGroupGreetingsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/groups/{groupId}/greetings][%d] postGroupGreetingsBadRequest  %+v", 400, o.Payload)
}

func (o *PostGroupGreetingsBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/groups/{groupId}/greetings][%d] postGroupGreetingsBadRequest  %+v", 400, o.Payload)
}

func (o *PostGroupGreetingsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGroupGreetingsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGroupGreetingsUnauthorized creates a PostGroupGreetingsUnauthorized with default headers values
func NewPostGroupGreetingsUnauthorized() *PostGroupGreetingsUnauthorized {
	return &PostGroupGreetingsUnauthorized{}
}

/*
PostGroupGreetingsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostGroupGreetingsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post group greetings unauthorized response has a 2xx status code
func (o *PostGroupGreetingsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post group greetings unauthorized response has a 3xx status code
func (o *PostGroupGreetingsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post group greetings unauthorized response has a 4xx status code
func (o *PostGroupGreetingsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post group greetings unauthorized response has a 5xx status code
func (o *PostGroupGreetingsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post group greetings unauthorized response a status code equal to that given
func (o *PostGroupGreetingsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostGroupGreetingsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/groups/{groupId}/greetings][%d] postGroupGreetingsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostGroupGreetingsUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/groups/{groupId}/greetings][%d] postGroupGreetingsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostGroupGreetingsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGroupGreetingsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGroupGreetingsForbidden creates a PostGroupGreetingsForbidden with default headers values
func NewPostGroupGreetingsForbidden() *PostGroupGreetingsForbidden {
	return &PostGroupGreetingsForbidden{}
}

/*
PostGroupGreetingsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostGroupGreetingsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post group greetings forbidden response has a 2xx status code
func (o *PostGroupGreetingsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post group greetings forbidden response has a 3xx status code
func (o *PostGroupGreetingsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post group greetings forbidden response has a 4xx status code
func (o *PostGroupGreetingsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post group greetings forbidden response has a 5xx status code
func (o *PostGroupGreetingsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post group greetings forbidden response a status code equal to that given
func (o *PostGroupGreetingsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostGroupGreetingsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/groups/{groupId}/greetings][%d] postGroupGreetingsForbidden  %+v", 403, o.Payload)
}

func (o *PostGroupGreetingsForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/groups/{groupId}/greetings][%d] postGroupGreetingsForbidden  %+v", 403, o.Payload)
}

func (o *PostGroupGreetingsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGroupGreetingsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGroupGreetingsNotFound creates a PostGroupGreetingsNotFound with default headers values
func NewPostGroupGreetingsNotFound() *PostGroupGreetingsNotFound {
	return &PostGroupGreetingsNotFound{}
}

/*
PostGroupGreetingsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostGroupGreetingsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post group greetings not found response has a 2xx status code
func (o *PostGroupGreetingsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post group greetings not found response has a 3xx status code
func (o *PostGroupGreetingsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post group greetings not found response has a 4xx status code
func (o *PostGroupGreetingsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post group greetings not found response has a 5xx status code
func (o *PostGroupGreetingsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post group greetings not found response a status code equal to that given
func (o *PostGroupGreetingsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostGroupGreetingsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/groups/{groupId}/greetings][%d] postGroupGreetingsNotFound  %+v", 404, o.Payload)
}

func (o *PostGroupGreetingsNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/groups/{groupId}/greetings][%d] postGroupGreetingsNotFound  %+v", 404, o.Payload)
}

func (o *PostGroupGreetingsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGroupGreetingsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGroupGreetingsRequestTimeout creates a PostGroupGreetingsRequestTimeout with default headers values
func NewPostGroupGreetingsRequestTimeout() *PostGroupGreetingsRequestTimeout {
	return &PostGroupGreetingsRequestTimeout{}
}

/*
PostGroupGreetingsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostGroupGreetingsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post group greetings request timeout response has a 2xx status code
func (o *PostGroupGreetingsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post group greetings request timeout response has a 3xx status code
func (o *PostGroupGreetingsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post group greetings request timeout response has a 4xx status code
func (o *PostGroupGreetingsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post group greetings request timeout response has a 5xx status code
func (o *PostGroupGreetingsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post group greetings request timeout response a status code equal to that given
func (o *PostGroupGreetingsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostGroupGreetingsRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/groups/{groupId}/greetings][%d] postGroupGreetingsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostGroupGreetingsRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/groups/{groupId}/greetings][%d] postGroupGreetingsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostGroupGreetingsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGroupGreetingsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGroupGreetingsRequestEntityTooLarge creates a PostGroupGreetingsRequestEntityTooLarge with default headers values
func NewPostGroupGreetingsRequestEntityTooLarge() *PostGroupGreetingsRequestEntityTooLarge {
	return &PostGroupGreetingsRequestEntityTooLarge{}
}

/*
PostGroupGreetingsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PostGroupGreetingsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post group greetings request entity too large response has a 2xx status code
func (o *PostGroupGreetingsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post group greetings request entity too large response has a 3xx status code
func (o *PostGroupGreetingsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post group greetings request entity too large response has a 4xx status code
func (o *PostGroupGreetingsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post group greetings request entity too large response has a 5xx status code
func (o *PostGroupGreetingsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post group greetings request entity too large response a status code equal to that given
func (o *PostGroupGreetingsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostGroupGreetingsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/groups/{groupId}/greetings][%d] postGroupGreetingsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostGroupGreetingsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/groups/{groupId}/greetings][%d] postGroupGreetingsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostGroupGreetingsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGroupGreetingsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGroupGreetingsUnsupportedMediaType creates a PostGroupGreetingsUnsupportedMediaType with default headers values
func NewPostGroupGreetingsUnsupportedMediaType() *PostGroupGreetingsUnsupportedMediaType {
	return &PostGroupGreetingsUnsupportedMediaType{}
}

/*
PostGroupGreetingsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostGroupGreetingsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post group greetings unsupported media type response has a 2xx status code
func (o *PostGroupGreetingsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post group greetings unsupported media type response has a 3xx status code
func (o *PostGroupGreetingsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post group greetings unsupported media type response has a 4xx status code
func (o *PostGroupGreetingsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post group greetings unsupported media type response has a 5xx status code
func (o *PostGroupGreetingsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post group greetings unsupported media type response a status code equal to that given
func (o *PostGroupGreetingsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostGroupGreetingsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/groups/{groupId}/greetings][%d] postGroupGreetingsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostGroupGreetingsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/groups/{groupId}/greetings][%d] postGroupGreetingsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostGroupGreetingsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGroupGreetingsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGroupGreetingsTooManyRequests creates a PostGroupGreetingsTooManyRequests with default headers values
func NewPostGroupGreetingsTooManyRequests() *PostGroupGreetingsTooManyRequests {
	return &PostGroupGreetingsTooManyRequests{}
}

/*
PostGroupGreetingsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostGroupGreetingsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post group greetings too many requests response has a 2xx status code
func (o *PostGroupGreetingsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post group greetings too many requests response has a 3xx status code
func (o *PostGroupGreetingsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post group greetings too many requests response has a 4xx status code
func (o *PostGroupGreetingsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post group greetings too many requests response has a 5xx status code
func (o *PostGroupGreetingsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post group greetings too many requests response a status code equal to that given
func (o *PostGroupGreetingsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostGroupGreetingsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/groups/{groupId}/greetings][%d] postGroupGreetingsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostGroupGreetingsTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/groups/{groupId}/greetings][%d] postGroupGreetingsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostGroupGreetingsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGroupGreetingsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGroupGreetingsInternalServerError creates a PostGroupGreetingsInternalServerError with default headers values
func NewPostGroupGreetingsInternalServerError() *PostGroupGreetingsInternalServerError {
	return &PostGroupGreetingsInternalServerError{}
}

/*
PostGroupGreetingsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostGroupGreetingsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post group greetings internal server error response has a 2xx status code
func (o *PostGroupGreetingsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post group greetings internal server error response has a 3xx status code
func (o *PostGroupGreetingsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post group greetings internal server error response has a 4xx status code
func (o *PostGroupGreetingsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post group greetings internal server error response has a 5xx status code
func (o *PostGroupGreetingsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post group greetings internal server error response a status code equal to that given
func (o *PostGroupGreetingsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostGroupGreetingsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/groups/{groupId}/greetings][%d] postGroupGreetingsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostGroupGreetingsInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/groups/{groupId}/greetings][%d] postGroupGreetingsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostGroupGreetingsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGroupGreetingsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGroupGreetingsServiceUnavailable creates a PostGroupGreetingsServiceUnavailable with default headers values
func NewPostGroupGreetingsServiceUnavailable() *PostGroupGreetingsServiceUnavailable {
	return &PostGroupGreetingsServiceUnavailable{}
}

/*
PostGroupGreetingsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostGroupGreetingsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post group greetings service unavailable response has a 2xx status code
func (o *PostGroupGreetingsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post group greetings service unavailable response has a 3xx status code
func (o *PostGroupGreetingsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post group greetings service unavailable response has a 4xx status code
func (o *PostGroupGreetingsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post group greetings service unavailable response has a 5xx status code
func (o *PostGroupGreetingsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post group greetings service unavailable response a status code equal to that given
func (o *PostGroupGreetingsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostGroupGreetingsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/groups/{groupId}/greetings][%d] postGroupGreetingsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostGroupGreetingsServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/groups/{groupId}/greetings][%d] postGroupGreetingsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostGroupGreetingsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGroupGreetingsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGroupGreetingsGatewayTimeout creates a PostGroupGreetingsGatewayTimeout with default headers values
func NewPostGroupGreetingsGatewayTimeout() *PostGroupGreetingsGatewayTimeout {
	return &PostGroupGreetingsGatewayTimeout{}
}

/*
PostGroupGreetingsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostGroupGreetingsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post group greetings gateway timeout response has a 2xx status code
func (o *PostGroupGreetingsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post group greetings gateway timeout response has a 3xx status code
func (o *PostGroupGreetingsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post group greetings gateway timeout response has a 4xx status code
func (o *PostGroupGreetingsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post group greetings gateway timeout response has a 5xx status code
func (o *PostGroupGreetingsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post group greetings gateway timeout response a status code equal to that given
func (o *PostGroupGreetingsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostGroupGreetingsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/groups/{groupId}/greetings][%d] postGroupGreetingsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostGroupGreetingsGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/groups/{groupId}/greetings][%d] postGroupGreetingsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostGroupGreetingsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGroupGreetingsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

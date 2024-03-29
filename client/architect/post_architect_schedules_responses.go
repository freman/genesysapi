// Code generated by go-swagger; DO NOT EDIT.

package architect

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostArchitectSchedulesReader is a Reader for the PostArchitectSchedules structure.
type PostArchitectSchedulesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostArchitectSchedulesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostArchitectSchedulesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostArchitectSchedulesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostArchitectSchedulesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostArchitectSchedulesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostArchitectSchedulesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostArchitectSchedulesRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostArchitectSchedulesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostArchitectSchedulesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostArchitectSchedulesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostArchitectSchedulesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostArchitectSchedulesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostArchitectSchedulesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostArchitectSchedulesOK creates a PostArchitectSchedulesOK with default headers values
func NewPostArchitectSchedulesOK() *PostArchitectSchedulesOK {
	return &PostArchitectSchedulesOK{}
}

/*
PostArchitectSchedulesOK describes a response with status code 200, with default header values.

successful operation
*/
type PostArchitectSchedulesOK struct {
	Payload *models.Schedule
}

// IsSuccess returns true when this post architect schedules o k response has a 2xx status code
func (o *PostArchitectSchedulesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post architect schedules o k response has a 3xx status code
func (o *PostArchitectSchedulesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post architect schedules o k response has a 4xx status code
func (o *PostArchitectSchedulesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post architect schedules o k response has a 5xx status code
func (o *PostArchitectSchedulesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post architect schedules o k response a status code equal to that given
func (o *PostArchitectSchedulesOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostArchitectSchedulesOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/schedules][%d] postArchitectSchedulesOK  %+v", 200, o.Payload)
}

func (o *PostArchitectSchedulesOK) String() string {
	return fmt.Sprintf("[POST /api/v2/architect/schedules][%d] postArchitectSchedulesOK  %+v", 200, o.Payload)
}

func (o *PostArchitectSchedulesOK) GetPayload() *models.Schedule {
	return o.Payload
}

func (o *PostArchitectSchedulesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Schedule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostArchitectSchedulesBadRequest creates a PostArchitectSchedulesBadRequest with default headers values
func NewPostArchitectSchedulesBadRequest() *PostArchitectSchedulesBadRequest {
	return &PostArchitectSchedulesBadRequest{}
}

/*
PostArchitectSchedulesBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostArchitectSchedulesBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post architect schedules bad request response has a 2xx status code
func (o *PostArchitectSchedulesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post architect schedules bad request response has a 3xx status code
func (o *PostArchitectSchedulesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post architect schedules bad request response has a 4xx status code
func (o *PostArchitectSchedulesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post architect schedules bad request response has a 5xx status code
func (o *PostArchitectSchedulesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post architect schedules bad request response a status code equal to that given
func (o *PostArchitectSchedulesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostArchitectSchedulesBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/schedules][%d] postArchitectSchedulesBadRequest  %+v", 400, o.Payload)
}

func (o *PostArchitectSchedulesBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/architect/schedules][%d] postArchitectSchedulesBadRequest  %+v", 400, o.Payload)
}

func (o *PostArchitectSchedulesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostArchitectSchedulesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostArchitectSchedulesUnauthorized creates a PostArchitectSchedulesUnauthorized with default headers values
func NewPostArchitectSchedulesUnauthorized() *PostArchitectSchedulesUnauthorized {
	return &PostArchitectSchedulesUnauthorized{}
}

/*
PostArchitectSchedulesUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostArchitectSchedulesUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post architect schedules unauthorized response has a 2xx status code
func (o *PostArchitectSchedulesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post architect schedules unauthorized response has a 3xx status code
func (o *PostArchitectSchedulesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post architect schedules unauthorized response has a 4xx status code
func (o *PostArchitectSchedulesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post architect schedules unauthorized response has a 5xx status code
func (o *PostArchitectSchedulesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post architect schedules unauthorized response a status code equal to that given
func (o *PostArchitectSchedulesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostArchitectSchedulesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/schedules][%d] postArchitectSchedulesUnauthorized  %+v", 401, o.Payload)
}

func (o *PostArchitectSchedulesUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/architect/schedules][%d] postArchitectSchedulesUnauthorized  %+v", 401, o.Payload)
}

func (o *PostArchitectSchedulesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostArchitectSchedulesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostArchitectSchedulesForbidden creates a PostArchitectSchedulesForbidden with default headers values
func NewPostArchitectSchedulesForbidden() *PostArchitectSchedulesForbidden {
	return &PostArchitectSchedulesForbidden{}
}

/*
PostArchitectSchedulesForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostArchitectSchedulesForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post architect schedules forbidden response has a 2xx status code
func (o *PostArchitectSchedulesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post architect schedules forbidden response has a 3xx status code
func (o *PostArchitectSchedulesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post architect schedules forbidden response has a 4xx status code
func (o *PostArchitectSchedulesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post architect schedules forbidden response has a 5xx status code
func (o *PostArchitectSchedulesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post architect schedules forbidden response a status code equal to that given
func (o *PostArchitectSchedulesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostArchitectSchedulesForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/schedules][%d] postArchitectSchedulesForbidden  %+v", 403, o.Payload)
}

func (o *PostArchitectSchedulesForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/architect/schedules][%d] postArchitectSchedulesForbidden  %+v", 403, o.Payload)
}

func (o *PostArchitectSchedulesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostArchitectSchedulesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostArchitectSchedulesNotFound creates a PostArchitectSchedulesNotFound with default headers values
func NewPostArchitectSchedulesNotFound() *PostArchitectSchedulesNotFound {
	return &PostArchitectSchedulesNotFound{}
}

/*
PostArchitectSchedulesNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostArchitectSchedulesNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post architect schedules not found response has a 2xx status code
func (o *PostArchitectSchedulesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post architect schedules not found response has a 3xx status code
func (o *PostArchitectSchedulesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post architect schedules not found response has a 4xx status code
func (o *PostArchitectSchedulesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post architect schedules not found response has a 5xx status code
func (o *PostArchitectSchedulesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post architect schedules not found response a status code equal to that given
func (o *PostArchitectSchedulesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostArchitectSchedulesNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/schedules][%d] postArchitectSchedulesNotFound  %+v", 404, o.Payload)
}

func (o *PostArchitectSchedulesNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/architect/schedules][%d] postArchitectSchedulesNotFound  %+v", 404, o.Payload)
}

func (o *PostArchitectSchedulesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostArchitectSchedulesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostArchitectSchedulesRequestTimeout creates a PostArchitectSchedulesRequestTimeout with default headers values
func NewPostArchitectSchedulesRequestTimeout() *PostArchitectSchedulesRequestTimeout {
	return &PostArchitectSchedulesRequestTimeout{}
}

/*
PostArchitectSchedulesRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostArchitectSchedulesRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post architect schedules request timeout response has a 2xx status code
func (o *PostArchitectSchedulesRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post architect schedules request timeout response has a 3xx status code
func (o *PostArchitectSchedulesRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post architect schedules request timeout response has a 4xx status code
func (o *PostArchitectSchedulesRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post architect schedules request timeout response has a 5xx status code
func (o *PostArchitectSchedulesRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post architect schedules request timeout response a status code equal to that given
func (o *PostArchitectSchedulesRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostArchitectSchedulesRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/schedules][%d] postArchitectSchedulesRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostArchitectSchedulesRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/architect/schedules][%d] postArchitectSchedulesRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostArchitectSchedulesRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostArchitectSchedulesRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostArchitectSchedulesRequestEntityTooLarge creates a PostArchitectSchedulesRequestEntityTooLarge with default headers values
func NewPostArchitectSchedulesRequestEntityTooLarge() *PostArchitectSchedulesRequestEntityTooLarge {
	return &PostArchitectSchedulesRequestEntityTooLarge{}
}

/*
PostArchitectSchedulesRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PostArchitectSchedulesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post architect schedules request entity too large response has a 2xx status code
func (o *PostArchitectSchedulesRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post architect schedules request entity too large response has a 3xx status code
func (o *PostArchitectSchedulesRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post architect schedules request entity too large response has a 4xx status code
func (o *PostArchitectSchedulesRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post architect schedules request entity too large response has a 5xx status code
func (o *PostArchitectSchedulesRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post architect schedules request entity too large response a status code equal to that given
func (o *PostArchitectSchedulesRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostArchitectSchedulesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/schedules][%d] postArchitectSchedulesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostArchitectSchedulesRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/architect/schedules][%d] postArchitectSchedulesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostArchitectSchedulesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostArchitectSchedulesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostArchitectSchedulesUnsupportedMediaType creates a PostArchitectSchedulesUnsupportedMediaType with default headers values
func NewPostArchitectSchedulesUnsupportedMediaType() *PostArchitectSchedulesUnsupportedMediaType {
	return &PostArchitectSchedulesUnsupportedMediaType{}
}

/*
PostArchitectSchedulesUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostArchitectSchedulesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post architect schedules unsupported media type response has a 2xx status code
func (o *PostArchitectSchedulesUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post architect schedules unsupported media type response has a 3xx status code
func (o *PostArchitectSchedulesUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post architect schedules unsupported media type response has a 4xx status code
func (o *PostArchitectSchedulesUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post architect schedules unsupported media type response has a 5xx status code
func (o *PostArchitectSchedulesUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post architect schedules unsupported media type response a status code equal to that given
func (o *PostArchitectSchedulesUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostArchitectSchedulesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/schedules][%d] postArchitectSchedulesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostArchitectSchedulesUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/architect/schedules][%d] postArchitectSchedulesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostArchitectSchedulesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostArchitectSchedulesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostArchitectSchedulesTooManyRequests creates a PostArchitectSchedulesTooManyRequests with default headers values
func NewPostArchitectSchedulesTooManyRequests() *PostArchitectSchedulesTooManyRequests {
	return &PostArchitectSchedulesTooManyRequests{}
}

/*
PostArchitectSchedulesTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostArchitectSchedulesTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post architect schedules too many requests response has a 2xx status code
func (o *PostArchitectSchedulesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post architect schedules too many requests response has a 3xx status code
func (o *PostArchitectSchedulesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post architect schedules too many requests response has a 4xx status code
func (o *PostArchitectSchedulesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post architect schedules too many requests response has a 5xx status code
func (o *PostArchitectSchedulesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post architect schedules too many requests response a status code equal to that given
func (o *PostArchitectSchedulesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostArchitectSchedulesTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/schedules][%d] postArchitectSchedulesTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostArchitectSchedulesTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/architect/schedules][%d] postArchitectSchedulesTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostArchitectSchedulesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostArchitectSchedulesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostArchitectSchedulesInternalServerError creates a PostArchitectSchedulesInternalServerError with default headers values
func NewPostArchitectSchedulesInternalServerError() *PostArchitectSchedulesInternalServerError {
	return &PostArchitectSchedulesInternalServerError{}
}

/*
PostArchitectSchedulesInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostArchitectSchedulesInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post architect schedules internal server error response has a 2xx status code
func (o *PostArchitectSchedulesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post architect schedules internal server error response has a 3xx status code
func (o *PostArchitectSchedulesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post architect schedules internal server error response has a 4xx status code
func (o *PostArchitectSchedulesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post architect schedules internal server error response has a 5xx status code
func (o *PostArchitectSchedulesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post architect schedules internal server error response a status code equal to that given
func (o *PostArchitectSchedulesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostArchitectSchedulesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/schedules][%d] postArchitectSchedulesInternalServerError  %+v", 500, o.Payload)
}

func (o *PostArchitectSchedulesInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/architect/schedules][%d] postArchitectSchedulesInternalServerError  %+v", 500, o.Payload)
}

func (o *PostArchitectSchedulesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostArchitectSchedulesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostArchitectSchedulesServiceUnavailable creates a PostArchitectSchedulesServiceUnavailable with default headers values
func NewPostArchitectSchedulesServiceUnavailable() *PostArchitectSchedulesServiceUnavailable {
	return &PostArchitectSchedulesServiceUnavailable{}
}

/*
PostArchitectSchedulesServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostArchitectSchedulesServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post architect schedules service unavailable response has a 2xx status code
func (o *PostArchitectSchedulesServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post architect schedules service unavailable response has a 3xx status code
func (o *PostArchitectSchedulesServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post architect schedules service unavailable response has a 4xx status code
func (o *PostArchitectSchedulesServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post architect schedules service unavailable response has a 5xx status code
func (o *PostArchitectSchedulesServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post architect schedules service unavailable response a status code equal to that given
func (o *PostArchitectSchedulesServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostArchitectSchedulesServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/schedules][%d] postArchitectSchedulesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostArchitectSchedulesServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/architect/schedules][%d] postArchitectSchedulesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostArchitectSchedulesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostArchitectSchedulesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostArchitectSchedulesGatewayTimeout creates a PostArchitectSchedulesGatewayTimeout with default headers values
func NewPostArchitectSchedulesGatewayTimeout() *PostArchitectSchedulesGatewayTimeout {
	return &PostArchitectSchedulesGatewayTimeout{}
}

/*
PostArchitectSchedulesGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostArchitectSchedulesGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post architect schedules gateway timeout response has a 2xx status code
func (o *PostArchitectSchedulesGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post architect schedules gateway timeout response has a 3xx status code
func (o *PostArchitectSchedulesGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post architect schedules gateway timeout response has a 4xx status code
func (o *PostArchitectSchedulesGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post architect schedules gateway timeout response has a 5xx status code
func (o *PostArchitectSchedulesGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post architect schedules gateway timeout response a status code equal to that given
func (o *PostArchitectSchedulesGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostArchitectSchedulesGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/schedules][%d] postArchitectSchedulesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostArchitectSchedulesGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/architect/schedules][%d] postArchitectSchedulesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostArchitectSchedulesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostArchitectSchedulesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

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

// PostArchitectSchedulegroupsReader is a Reader for the PostArchitectSchedulegroups structure.
type PostArchitectSchedulegroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostArchitectSchedulegroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostArchitectSchedulegroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostArchitectSchedulegroupsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostArchitectSchedulegroupsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostArchitectSchedulegroupsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostArchitectSchedulegroupsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostArchitectSchedulegroupsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostArchitectSchedulegroupsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostArchitectSchedulegroupsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostArchitectSchedulegroupsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostArchitectSchedulegroupsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostArchitectSchedulegroupsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostArchitectSchedulegroupsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostArchitectSchedulegroupsOK creates a PostArchitectSchedulegroupsOK with default headers values
func NewPostArchitectSchedulegroupsOK() *PostArchitectSchedulegroupsOK {
	return &PostArchitectSchedulegroupsOK{}
}

/*
PostArchitectSchedulegroupsOK describes a response with status code 200, with default header values.

successful operation
*/
type PostArchitectSchedulegroupsOK struct {
	Payload *models.ScheduleGroup
}

// IsSuccess returns true when this post architect schedulegroups o k response has a 2xx status code
func (o *PostArchitectSchedulegroupsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post architect schedulegroups o k response has a 3xx status code
func (o *PostArchitectSchedulegroupsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post architect schedulegroups o k response has a 4xx status code
func (o *PostArchitectSchedulegroupsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post architect schedulegroups o k response has a 5xx status code
func (o *PostArchitectSchedulegroupsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post architect schedulegroups o k response a status code equal to that given
func (o *PostArchitectSchedulegroupsOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostArchitectSchedulegroupsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/schedulegroups][%d] postArchitectSchedulegroupsOK  %+v", 200, o.Payload)
}

func (o *PostArchitectSchedulegroupsOK) String() string {
	return fmt.Sprintf("[POST /api/v2/architect/schedulegroups][%d] postArchitectSchedulegroupsOK  %+v", 200, o.Payload)
}

func (o *PostArchitectSchedulegroupsOK) GetPayload() *models.ScheduleGroup {
	return o.Payload
}

func (o *PostArchitectSchedulegroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ScheduleGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostArchitectSchedulegroupsBadRequest creates a PostArchitectSchedulegroupsBadRequest with default headers values
func NewPostArchitectSchedulegroupsBadRequest() *PostArchitectSchedulegroupsBadRequest {
	return &PostArchitectSchedulegroupsBadRequest{}
}

/*
PostArchitectSchedulegroupsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostArchitectSchedulegroupsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post architect schedulegroups bad request response has a 2xx status code
func (o *PostArchitectSchedulegroupsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post architect schedulegroups bad request response has a 3xx status code
func (o *PostArchitectSchedulegroupsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post architect schedulegroups bad request response has a 4xx status code
func (o *PostArchitectSchedulegroupsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post architect schedulegroups bad request response has a 5xx status code
func (o *PostArchitectSchedulegroupsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post architect schedulegroups bad request response a status code equal to that given
func (o *PostArchitectSchedulegroupsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostArchitectSchedulegroupsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/schedulegroups][%d] postArchitectSchedulegroupsBadRequest  %+v", 400, o.Payload)
}

func (o *PostArchitectSchedulegroupsBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/architect/schedulegroups][%d] postArchitectSchedulegroupsBadRequest  %+v", 400, o.Payload)
}

func (o *PostArchitectSchedulegroupsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostArchitectSchedulegroupsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostArchitectSchedulegroupsUnauthorized creates a PostArchitectSchedulegroupsUnauthorized with default headers values
func NewPostArchitectSchedulegroupsUnauthorized() *PostArchitectSchedulegroupsUnauthorized {
	return &PostArchitectSchedulegroupsUnauthorized{}
}

/*
PostArchitectSchedulegroupsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostArchitectSchedulegroupsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post architect schedulegroups unauthorized response has a 2xx status code
func (o *PostArchitectSchedulegroupsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post architect schedulegroups unauthorized response has a 3xx status code
func (o *PostArchitectSchedulegroupsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post architect schedulegroups unauthorized response has a 4xx status code
func (o *PostArchitectSchedulegroupsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post architect schedulegroups unauthorized response has a 5xx status code
func (o *PostArchitectSchedulegroupsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post architect schedulegroups unauthorized response a status code equal to that given
func (o *PostArchitectSchedulegroupsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostArchitectSchedulegroupsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/schedulegroups][%d] postArchitectSchedulegroupsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostArchitectSchedulegroupsUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/architect/schedulegroups][%d] postArchitectSchedulegroupsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostArchitectSchedulegroupsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostArchitectSchedulegroupsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostArchitectSchedulegroupsForbidden creates a PostArchitectSchedulegroupsForbidden with default headers values
func NewPostArchitectSchedulegroupsForbidden() *PostArchitectSchedulegroupsForbidden {
	return &PostArchitectSchedulegroupsForbidden{}
}

/*
PostArchitectSchedulegroupsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostArchitectSchedulegroupsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post architect schedulegroups forbidden response has a 2xx status code
func (o *PostArchitectSchedulegroupsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post architect schedulegroups forbidden response has a 3xx status code
func (o *PostArchitectSchedulegroupsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post architect schedulegroups forbidden response has a 4xx status code
func (o *PostArchitectSchedulegroupsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post architect schedulegroups forbidden response has a 5xx status code
func (o *PostArchitectSchedulegroupsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post architect schedulegroups forbidden response a status code equal to that given
func (o *PostArchitectSchedulegroupsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostArchitectSchedulegroupsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/schedulegroups][%d] postArchitectSchedulegroupsForbidden  %+v", 403, o.Payload)
}

func (o *PostArchitectSchedulegroupsForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/architect/schedulegroups][%d] postArchitectSchedulegroupsForbidden  %+v", 403, o.Payload)
}

func (o *PostArchitectSchedulegroupsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostArchitectSchedulegroupsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostArchitectSchedulegroupsNotFound creates a PostArchitectSchedulegroupsNotFound with default headers values
func NewPostArchitectSchedulegroupsNotFound() *PostArchitectSchedulegroupsNotFound {
	return &PostArchitectSchedulegroupsNotFound{}
}

/*
PostArchitectSchedulegroupsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostArchitectSchedulegroupsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post architect schedulegroups not found response has a 2xx status code
func (o *PostArchitectSchedulegroupsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post architect schedulegroups not found response has a 3xx status code
func (o *PostArchitectSchedulegroupsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post architect schedulegroups not found response has a 4xx status code
func (o *PostArchitectSchedulegroupsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post architect schedulegroups not found response has a 5xx status code
func (o *PostArchitectSchedulegroupsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post architect schedulegroups not found response a status code equal to that given
func (o *PostArchitectSchedulegroupsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostArchitectSchedulegroupsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/schedulegroups][%d] postArchitectSchedulegroupsNotFound  %+v", 404, o.Payload)
}

func (o *PostArchitectSchedulegroupsNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/architect/schedulegroups][%d] postArchitectSchedulegroupsNotFound  %+v", 404, o.Payload)
}

func (o *PostArchitectSchedulegroupsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostArchitectSchedulegroupsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostArchitectSchedulegroupsRequestTimeout creates a PostArchitectSchedulegroupsRequestTimeout with default headers values
func NewPostArchitectSchedulegroupsRequestTimeout() *PostArchitectSchedulegroupsRequestTimeout {
	return &PostArchitectSchedulegroupsRequestTimeout{}
}

/*
PostArchitectSchedulegroupsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostArchitectSchedulegroupsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post architect schedulegroups request timeout response has a 2xx status code
func (o *PostArchitectSchedulegroupsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post architect schedulegroups request timeout response has a 3xx status code
func (o *PostArchitectSchedulegroupsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post architect schedulegroups request timeout response has a 4xx status code
func (o *PostArchitectSchedulegroupsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post architect schedulegroups request timeout response has a 5xx status code
func (o *PostArchitectSchedulegroupsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post architect schedulegroups request timeout response a status code equal to that given
func (o *PostArchitectSchedulegroupsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostArchitectSchedulegroupsRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/schedulegroups][%d] postArchitectSchedulegroupsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostArchitectSchedulegroupsRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/architect/schedulegroups][%d] postArchitectSchedulegroupsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostArchitectSchedulegroupsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostArchitectSchedulegroupsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostArchitectSchedulegroupsRequestEntityTooLarge creates a PostArchitectSchedulegroupsRequestEntityTooLarge with default headers values
func NewPostArchitectSchedulegroupsRequestEntityTooLarge() *PostArchitectSchedulegroupsRequestEntityTooLarge {
	return &PostArchitectSchedulegroupsRequestEntityTooLarge{}
}

/*
PostArchitectSchedulegroupsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PostArchitectSchedulegroupsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post architect schedulegroups request entity too large response has a 2xx status code
func (o *PostArchitectSchedulegroupsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post architect schedulegroups request entity too large response has a 3xx status code
func (o *PostArchitectSchedulegroupsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post architect schedulegroups request entity too large response has a 4xx status code
func (o *PostArchitectSchedulegroupsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post architect schedulegroups request entity too large response has a 5xx status code
func (o *PostArchitectSchedulegroupsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post architect schedulegroups request entity too large response a status code equal to that given
func (o *PostArchitectSchedulegroupsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostArchitectSchedulegroupsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/schedulegroups][%d] postArchitectSchedulegroupsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostArchitectSchedulegroupsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/architect/schedulegroups][%d] postArchitectSchedulegroupsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostArchitectSchedulegroupsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostArchitectSchedulegroupsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostArchitectSchedulegroupsUnsupportedMediaType creates a PostArchitectSchedulegroupsUnsupportedMediaType with default headers values
func NewPostArchitectSchedulegroupsUnsupportedMediaType() *PostArchitectSchedulegroupsUnsupportedMediaType {
	return &PostArchitectSchedulegroupsUnsupportedMediaType{}
}

/*
PostArchitectSchedulegroupsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostArchitectSchedulegroupsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post architect schedulegroups unsupported media type response has a 2xx status code
func (o *PostArchitectSchedulegroupsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post architect schedulegroups unsupported media type response has a 3xx status code
func (o *PostArchitectSchedulegroupsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post architect schedulegroups unsupported media type response has a 4xx status code
func (o *PostArchitectSchedulegroupsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post architect schedulegroups unsupported media type response has a 5xx status code
func (o *PostArchitectSchedulegroupsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post architect schedulegroups unsupported media type response a status code equal to that given
func (o *PostArchitectSchedulegroupsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostArchitectSchedulegroupsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/schedulegroups][%d] postArchitectSchedulegroupsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostArchitectSchedulegroupsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/architect/schedulegroups][%d] postArchitectSchedulegroupsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostArchitectSchedulegroupsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostArchitectSchedulegroupsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostArchitectSchedulegroupsTooManyRequests creates a PostArchitectSchedulegroupsTooManyRequests with default headers values
func NewPostArchitectSchedulegroupsTooManyRequests() *PostArchitectSchedulegroupsTooManyRequests {
	return &PostArchitectSchedulegroupsTooManyRequests{}
}

/*
PostArchitectSchedulegroupsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostArchitectSchedulegroupsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post architect schedulegroups too many requests response has a 2xx status code
func (o *PostArchitectSchedulegroupsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post architect schedulegroups too many requests response has a 3xx status code
func (o *PostArchitectSchedulegroupsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post architect schedulegroups too many requests response has a 4xx status code
func (o *PostArchitectSchedulegroupsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post architect schedulegroups too many requests response has a 5xx status code
func (o *PostArchitectSchedulegroupsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post architect schedulegroups too many requests response a status code equal to that given
func (o *PostArchitectSchedulegroupsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostArchitectSchedulegroupsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/schedulegroups][%d] postArchitectSchedulegroupsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostArchitectSchedulegroupsTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/architect/schedulegroups][%d] postArchitectSchedulegroupsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostArchitectSchedulegroupsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostArchitectSchedulegroupsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostArchitectSchedulegroupsInternalServerError creates a PostArchitectSchedulegroupsInternalServerError with default headers values
func NewPostArchitectSchedulegroupsInternalServerError() *PostArchitectSchedulegroupsInternalServerError {
	return &PostArchitectSchedulegroupsInternalServerError{}
}

/*
PostArchitectSchedulegroupsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostArchitectSchedulegroupsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post architect schedulegroups internal server error response has a 2xx status code
func (o *PostArchitectSchedulegroupsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post architect schedulegroups internal server error response has a 3xx status code
func (o *PostArchitectSchedulegroupsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post architect schedulegroups internal server error response has a 4xx status code
func (o *PostArchitectSchedulegroupsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post architect schedulegroups internal server error response has a 5xx status code
func (o *PostArchitectSchedulegroupsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post architect schedulegroups internal server error response a status code equal to that given
func (o *PostArchitectSchedulegroupsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostArchitectSchedulegroupsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/schedulegroups][%d] postArchitectSchedulegroupsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostArchitectSchedulegroupsInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/architect/schedulegroups][%d] postArchitectSchedulegroupsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostArchitectSchedulegroupsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostArchitectSchedulegroupsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostArchitectSchedulegroupsServiceUnavailable creates a PostArchitectSchedulegroupsServiceUnavailable with default headers values
func NewPostArchitectSchedulegroupsServiceUnavailable() *PostArchitectSchedulegroupsServiceUnavailable {
	return &PostArchitectSchedulegroupsServiceUnavailable{}
}

/*
PostArchitectSchedulegroupsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostArchitectSchedulegroupsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post architect schedulegroups service unavailable response has a 2xx status code
func (o *PostArchitectSchedulegroupsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post architect schedulegroups service unavailable response has a 3xx status code
func (o *PostArchitectSchedulegroupsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post architect schedulegroups service unavailable response has a 4xx status code
func (o *PostArchitectSchedulegroupsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post architect schedulegroups service unavailable response has a 5xx status code
func (o *PostArchitectSchedulegroupsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post architect schedulegroups service unavailable response a status code equal to that given
func (o *PostArchitectSchedulegroupsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostArchitectSchedulegroupsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/schedulegroups][%d] postArchitectSchedulegroupsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostArchitectSchedulegroupsServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/architect/schedulegroups][%d] postArchitectSchedulegroupsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostArchitectSchedulegroupsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostArchitectSchedulegroupsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostArchitectSchedulegroupsGatewayTimeout creates a PostArchitectSchedulegroupsGatewayTimeout with default headers values
func NewPostArchitectSchedulegroupsGatewayTimeout() *PostArchitectSchedulegroupsGatewayTimeout {
	return &PostArchitectSchedulegroupsGatewayTimeout{}
}

/*
PostArchitectSchedulegroupsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostArchitectSchedulegroupsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post architect schedulegroups gateway timeout response has a 2xx status code
func (o *PostArchitectSchedulegroupsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post architect schedulegroups gateway timeout response has a 3xx status code
func (o *PostArchitectSchedulegroupsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post architect schedulegroups gateway timeout response has a 4xx status code
func (o *PostArchitectSchedulegroupsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post architect schedulegroups gateway timeout response has a 5xx status code
func (o *PostArchitectSchedulegroupsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post architect schedulegroups gateway timeout response a status code equal to that given
func (o *PostArchitectSchedulegroupsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostArchitectSchedulegroupsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/schedulegroups][%d] postArchitectSchedulegroupsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostArchitectSchedulegroupsGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/architect/schedulegroups][%d] postArchitectSchedulegroupsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostArchitectSchedulegroupsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostArchitectSchedulegroupsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

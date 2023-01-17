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

// PutArchitectSchedulegroupReader is a Reader for the PutArchitectSchedulegroup structure.
type PutArchitectSchedulegroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutArchitectSchedulegroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutArchitectSchedulegroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutArchitectSchedulegroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutArchitectSchedulegroupUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutArchitectSchedulegroupForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutArchitectSchedulegroupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPutArchitectSchedulegroupRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPutArchitectSchedulegroupRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPutArchitectSchedulegroupUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutArchitectSchedulegroupTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutArchitectSchedulegroupInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutArchitectSchedulegroupServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutArchitectSchedulegroupGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutArchitectSchedulegroupOK creates a PutArchitectSchedulegroupOK with default headers values
func NewPutArchitectSchedulegroupOK() *PutArchitectSchedulegroupOK {
	return &PutArchitectSchedulegroupOK{}
}

/*
PutArchitectSchedulegroupOK describes a response with status code 200, with default header values.

successful operation
*/
type PutArchitectSchedulegroupOK struct {
	Payload *models.ScheduleGroup
}

// IsSuccess returns true when this put architect schedulegroup o k response has a 2xx status code
func (o *PutArchitectSchedulegroupOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put architect schedulegroup o k response has a 3xx status code
func (o *PutArchitectSchedulegroupOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put architect schedulegroup o k response has a 4xx status code
func (o *PutArchitectSchedulegroupOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put architect schedulegroup o k response has a 5xx status code
func (o *PutArchitectSchedulegroupOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put architect schedulegroup o k response a status code equal to that given
func (o *PutArchitectSchedulegroupOK) IsCode(code int) bool {
	return code == 200
}

func (o *PutArchitectSchedulegroupOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/architect/schedulegroups/{scheduleGroupId}][%d] putArchitectSchedulegroupOK  %+v", 200, o.Payload)
}

func (o *PutArchitectSchedulegroupOK) String() string {
	return fmt.Sprintf("[PUT /api/v2/architect/schedulegroups/{scheduleGroupId}][%d] putArchitectSchedulegroupOK  %+v", 200, o.Payload)
}

func (o *PutArchitectSchedulegroupOK) GetPayload() *models.ScheduleGroup {
	return o.Payload
}

func (o *PutArchitectSchedulegroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ScheduleGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutArchitectSchedulegroupBadRequest creates a PutArchitectSchedulegroupBadRequest with default headers values
func NewPutArchitectSchedulegroupBadRequest() *PutArchitectSchedulegroupBadRequest {
	return &PutArchitectSchedulegroupBadRequest{}
}

/*
PutArchitectSchedulegroupBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutArchitectSchedulegroupBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put architect schedulegroup bad request response has a 2xx status code
func (o *PutArchitectSchedulegroupBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put architect schedulegroup bad request response has a 3xx status code
func (o *PutArchitectSchedulegroupBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put architect schedulegroup bad request response has a 4xx status code
func (o *PutArchitectSchedulegroupBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this put architect schedulegroup bad request response has a 5xx status code
func (o *PutArchitectSchedulegroupBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this put architect schedulegroup bad request response a status code equal to that given
func (o *PutArchitectSchedulegroupBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PutArchitectSchedulegroupBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/architect/schedulegroups/{scheduleGroupId}][%d] putArchitectSchedulegroupBadRequest  %+v", 400, o.Payload)
}

func (o *PutArchitectSchedulegroupBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/v2/architect/schedulegroups/{scheduleGroupId}][%d] putArchitectSchedulegroupBadRequest  %+v", 400, o.Payload)
}

func (o *PutArchitectSchedulegroupBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutArchitectSchedulegroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutArchitectSchedulegroupUnauthorized creates a PutArchitectSchedulegroupUnauthorized with default headers values
func NewPutArchitectSchedulegroupUnauthorized() *PutArchitectSchedulegroupUnauthorized {
	return &PutArchitectSchedulegroupUnauthorized{}
}

/*
PutArchitectSchedulegroupUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PutArchitectSchedulegroupUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put architect schedulegroup unauthorized response has a 2xx status code
func (o *PutArchitectSchedulegroupUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put architect schedulegroup unauthorized response has a 3xx status code
func (o *PutArchitectSchedulegroupUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put architect schedulegroup unauthorized response has a 4xx status code
func (o *PutArchitectSchedulegroupUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this put architect schedulegroup unauthorized response has a 5xx status code
func (o *PutArchitectSchedulegroupUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this put architect schedulegroup unauthorized response a status code equal to that given
func (o *PutArchitectSchedulegroupUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PutArchitectSchedulegroupUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/architect/schedulegroups/{scheduleGroupId}][%d] putArchitectSchedulegroupUnauthorized  %+v", 401, o.Payload)
}

func (o *PutArchitectSchedulegroupUnauthorized) String() string {
	return fmt.Sprintf("[PUT /api/v2/architect/schedulegroups/{scheduleGroupId}][%d] putArchitectSchedulegroupUnauthorized  %+v", 401, o.Payload)
}

func (o *PutArchitectSchedulegroupUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutArchitectSchedulegroupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutArchitectSchedulegroupForbidden creates a PutArchitectSchedulegroupForbidden with default headers values
func NewPutArchitectSchedulegroupForbidden() *PutArchitectSchedulegroupForbidden {
	return &PutArchitectSchedulegroupForbidden{}
}

/*
PutArchitectSchedulegroupForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PutArchitectSchedulegroupForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put architect schedulegroup forbidden response has a 2xx status code
func (o *PutArchitectSchedulegroupForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put architect schedulegroup forbidden response has a 3xx status code
func (o *PutArchitectSchedulegroupForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put architect schedulegroup forbidden response has a 4xx status code
func (o *PutArchitectSchedulegroupForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this put architect schedulegroup forbidden response has a 5xx status code
func (o *PutArchitectSchedulegroupForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this put architect schedulegroup forbidden response a status code equal to that given
func (o *PutArchitectSchedulegroupForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PutArchitectSchedulegroupForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/architect/schedulegroups/{scheduleGroupId}][%d] putArchitectSchedulegroupForbidden  %+v", 403, o.Payload)
}

func (o *PutArchitectSchedulegroupForbidden) String() string {
	return fmt.Sprintf("[PUT /api/v2/architect/schedulegroups/{scheduleGroupId}][%d] putArchitectSchedulegroupForbidden  %+v", 403, o.Payload)
}

func (o *PutArchitectSchedulegroupForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutArchitectSchedulegroupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutArchitectSchedulegroupNotFound creates a PutArchitectSchedulegroupNotFound with default headers values
func NewPutArchitectSchedulegroupNotFound() *PutArchitectSchedulegroupNotFound {
	return &PutArchitectSchedulegroupNotFound{}
}

/*
PutArchitectSchedulegroupNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PutArchitectSchedulegroupNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put architect schedulegroup not found response has a 2xx status code
func (o *PutArchitectSchedulegroupNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put architect schedulegroup not found response has a 3xx status code
func (o *PutArchitectSchedulegroupNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put architect schedulegroup not found response has a 4xx status code
func (o *PutArchitectSchedulegroupNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this put architect schedulegroup not found response has a 5xx status code
func (o *PutArchitectSchedulegroupNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this put architect schedulegroup not found response a status code equal to that given
func (o *PutArchitectSchedulegroupNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PutArchitectSchedulegroupNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/architect/schedulegroups/{scheduleGroupId}][%d] putArchitectSchedulegroupNotFound  %+v", 404, o.Payload)
}

func (o *PutArchitectSchedulegroupNotFound) String() string {
	return fmt.Sprintf("[PUT /api/v2/architect/schedulegroups/{scheduleGroupId}][%d] putArchitectSchedulegroupNotFound  %+v", 404, o.Payload)
}

func (o *PutArchitectSchedulegroupNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutArchitectSchedulegroupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutArchitectSchedulegroupRequestTimeout creates a PutArchitectSchedulegroupRequestTimeout with default headers values
func NewPutArchitectSchedulegroupRequestTimeout() *PutArchitectSchedulegroupRequestTimeout {
	return &PutArchitectSchedulegroupRequestTimeout{}
}

/*
PutArchitectSchedulegroupRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PutArchitectSchedulegroupRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put architect schedulegroup request timeout response has a 2xx status code
func (o *PutArchitectSchedulegroupRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put architect schedulegroup request timeout response has a 3xx status code
func (o *PutArchitectSchedulegroupRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put architect schedulegroup request timeout response has a 4xx status code
func (o *PutArchitectSchedulegroupRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this put architect schedulegroup request timeout response has a 5xx status code
func (o *PutArchitectSchedulegroupRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this put architect schedulegroup request timeout response a status code equal to that given
func (o *PutArchitectSchedulegroupRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PutArchitectSchedulegroupRequestTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/architect/schedulegroups/{scheduleGroupId}][%d] putArchitectSchedulegroupRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutArchitectSchedulegroupRequestTimeout) String() string {
	return fmt.Sprintf("[PUT /api/v2/architect/schedulegroups/{scheduleGroupId}][%d] putArchitectSchedulegroupRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutArchitectSchedulegroupRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutArchitectSchedulegroupRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutArchitectSchedulegroupRequestEntityTooLarge creates a PutArchitectSchedulegroupRequestEntityTooLarge with default headers values
func NewPutArchitectSchedulegroupRequestEntityTooLarge() *PutArchitectSchedulegroupRequestEntityTooLarge {
	return &PutArchitectSchedulegroupRequestEntityTooLarge{}
}

/*
PutArchitectSchedulegroupRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PutArchitectSchedulegroupRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put architect schedulegroup request entity too large response has a 2xx status code
func (o *PutArchitectSchedulegroupRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put architect schedulegroup request entity too large response has a 3xx status code
func (o *PutArchitectSchedulegroupRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put architect schedulegroup request entity too large response has a 4xx status code
func (o *PutArchitectSchedulegroupRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this put architect schedulegroup request entity too large response has a 5xx status code
func (o *PutArchitectSchedulegroupRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this put architect schedulegroup request entity too large response a status code equal to that given
func (o *PutArchitectSchedulegroupRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PutArchitectSchedulegroupRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /api/v2/architect/schedulegroups/{scheduleGroupId}][%d] putArchitectSchedulegroupRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutArchitectSchedulegroupRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[PUT /api/v2/architect/schedulegroups/{scheduleGroupId}][%d] putArchitectSchedulegroupRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutArchitectSchedulegroupRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutArchitectSchedulegroupRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutArchitectSchedulegroupUnsupportedMediaType creates a PutArchitectSchedulegroupUnsupportedMediaType with default headers values
func NewPutArchitectSchedulegroupUnsupportedMediaType() *PutArchitectSchedulegroupUnsupportedMediaType {
	return &PutArchitectSchedulegroupUnsupportedMediaType{}
}

/*
PutArchitectSchedulegroupUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutArchitectSchedulegroupUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put architect schedulegroup unsupported media type response has a 2xx status code
func (o *PutArchitectSchedulegroupUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put architect schedulegroup unsupported media type response has a 3xx status code
func (o *PutArchitectSchedulegroupUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put architect schedulegroup unsupported media type response has a 4xx status code
func (o *PutArchitectSchedulegroupUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this put architect schedulegroup unsupported media type response has a 5xx status code
func (o *PutArchitectSchedulegroupUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this put architect schedulegroup unsupported media type response a status code equal to that given
func (o *PutArchitectSchedulegroupUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PutArchitectSchedulegroupUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v2/architect/schedulegroups/{scheduleGroupId}][%d] putArchitectSchedulegroupUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutArchitectSchedulegroupUnsupportedMediaType) String() string {
	return fmt.Sprintf("[PUT /api/v2/architect/schedulegroups/{scheduleGroupId}][%d] putArchitectSchedulegroupUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutArchitectSchedulegroupUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutArchitectSchedulegroupUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutArchitectSchedulegroupTooManyRequests creates a PutArchitectSchedulegroupTooManyRequests with default headers values
func NewPutArchitectSchedulegroupTooManyRequests() *PutArchitectSchedulegroupTooManyRequests {
	return &PutArchitectSchedulegroupTooManyRequests{}
}

/*
PutArchitectSchedulegroupTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PutArchitectSchedulegroupTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put architect schedulegroup too many requests response has a 2xx status code
func (o *PutArchitectSchedulegroupTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put architect schedulegroup too many requests response has a 3xx status code
func (o *PutArchitectSchedulegroupTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put architect schedulegroup too many requests response has a 4xx status code
func (o *PutArchitectSchedulegroupTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this put architect schedulegroup too many requests response has a 5xx status code
func (o *PutArchitectSchedulegroupTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this put architect schedulegroup too many requests response a status code equal to that given
func (o *PutArchitectSchedulegroupTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PutArchitectSchedulegroupTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v2/architect/schedulegroups/{scheduleGroupId}][%d] putArchitectSchedulegroupTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutArchitectSchedulegroupTooManyRequests) String() string {
	return fmt.Sprintf("[PUT /api/v2/architect/schedulegroups/{scheduleGroupId}][%d] putArchitectSchedulegroupTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutArchitectSchedulegroupTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutArchitectSchedulegroupTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutArchitectSchedulegroupInternalServerError creates a PutArchitectSchedulegroupInternalServerError with default headers values
func NewPutArchitectSchedulegroupInternalServerError() *PutArchitectSchedulegroupInternalServerError {
	return &PutArchitectSchedulegroupInternalServerError{}
}

/*
PutArchitectSchedulegroupInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutArchitectSchedulegroupInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put architect schedulegroup internal server error response has a 2xx status code
func (o *PutArchitectSchedulegroupInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put architect schedulegroup internal server error response has a 3xx status code
func (o *PutArchitectSchedulegroupInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put architect schedulegroup internal server error response has a 4xx status code
func (o *PutArchitectSchedulegroupInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this put architect schedulegroup internal server error response has a 5xx status code
func (o *PutArchitectSchedulegroupInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this put architect schedulegroup internal server error response a status code equal to that given
func (o *PutArchitectSchedulegroupInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PutArchitectSchedulegroupInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/architect/schedulegroups/{scheduleGroupId}][%d] putArchitectSchedulegroupInternalServerError  %+v", 500, o.Payload)
}

func (o *PutArchitectSchedulegroupInternalServerError) String() string {
	return fmt.Sprintf("[PUT /api/v2/architect/schedulegroups/{scheduleGroupId}][%d] putArchitectSchedulegroupInternalServerError  %+v", 500, o.Payload)
}

func (o *PutArchitectSchedulegroupInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutArchitectSchedulegroupInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutArchitectSchedulegroupServiceUnavailable creates a PutArchitectSchedulegroupServiceUnavailable with default headers values
func NewPutArchitectSchedulegroupServiceUnavailable() *PutArchitectSchedulegroupServiceUnavailable {
	return &PutArchitectSchedulegroupServiceUnavailable{}
}

/*
PutArchitectSchedulegroupServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutArchitectSchedulegroupServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put architect schedulegroup service unavailable response has a 2xx status code
func (o *PutArchitectSchedulegroupServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put architect schedulegroup service unavailable response has a 3xx status code
func (o *PutArchitectSchedulegroupServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put architect schedulegroup service unavailable response has a 4xx status code
func (o *PutArchitectSchedulegroupServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this put architect schedulegroup service unavailable response has a 5xx status code
func (o *PutArchitectSchedulegroupServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this put architect schedulegroup service unavailable response a status code equal to that given
func (o *PutArchitectSchedulegroupServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PutArchitectSchedulegroupServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/architect/schedulegroups/{scheduleGroupId}][%d] putArchitectSchedulegroupServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutArchitectSchedulegroupServiceUnavailable) String() string {
	return fmt.Sprintf("[PUT /api/v2/architect/schedulegroups/{scheduleGroupId}][%d] putArchitectSchedulegroupServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutArchitectSchedulegroupServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutArchitectSchedulegroupServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutArchitectSchedulegroupGatewayTimeout creates a PutArchitectSchedulegroupGatewayTimeout with default headers values
func NewPutArchitectSchedulegroupGatewayTimeout() *PutArchitectSchedulegroupGatewayTimeout {
	return &PutArchitectSchedulegroupGatewayTimeout{}
}

/*
PutArchitectSchedulegroupGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PutArchitectSchedulegroupGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put architect schedulegroup gateway timeout response has a 2xx status code
func (o *PutArchitectSchedulegroupGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put architect schedulegroup gateway timeout response has a 3xx status code
func (o *PutArchitectSchedulegroupGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put architect schedulegroup gateway timeout response has a 4xx status code
func (o *PutArchitectSchedulegroupGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this put architect schedulegroup gateway timeout response has a 5xx status code
func (o *PutArchitectSchedulegroupGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this put architect schedulegroup gateway timeout response a status code equal to that given
func (o *PutArchitectSchedulegroupGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PutArchitectSchedulegroupGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/architect/schedulegroups/{scheduleGroupId}][%d] putArchitectSchedulegroupGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutArchitectSchedulegroupGatewayTimeout) String() string {
	return fmt.Sprintf("[PUT /api/v2/architect/schedulegroups/{scheduleGroupId}][%d] putArchitectSchedulegroupGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutArchitectSchedulegroupGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutArchitectSchedulegroupGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

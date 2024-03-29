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

// GetArchitectScheduleReader is a Reader for the GetArchitectSchedule structure.
type GetArchitectScheduleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetArchitectScheduleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetArchitectScheduleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetArchitectScheduleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetArchitectScheduleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetArchitectScheduleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetArchitectScheduleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetArchitectScheduleRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetArchitectScheduleRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetArchitectScheduleUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetArchitectScheduleTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetArchitectScheduleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetArchitectScheduleServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetArchitectScheduleGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetArchitectScheduleOK creates a GetArchitectScheduleOK with default headers values
func NewGetArchitectScheduleOK() *GetArchitectScheduleOK {
	return &GetArchitectScheduleOK{}
}

/*
GetArchitectScheduleOK describes a response with status code 200, with default header values.

successful operation
*/
type GetArchitectScheduleOK struct {
	Payload *models.Schedule
}

// IsSuccess returns true when this get architect schedule o k response has a 2xx status code
func (o *GetArchitectScheduleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get architect schedule o k response has a 3xx status code
func (o *GetArchitectScheduleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get architect schedule o k response has a 4xx status code
func (o *GetArchitectScheduleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get architect schedule o k response has a 5xx status code
func (o *GetArchitectScheduleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get architect schedule o k response a status code equal to that given
func (o *GetArchitectScheduleOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetArchitectScheduleOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/schedules/{scheduleId}][%d] getArchitectScheduleOK  %+v", 200, o.Payload)
}

func (o *GetArchitectScheduleOK) String() string {
	return fmt.Sprintf("[GET /api/v2/architect/schedules/{scheduleId}][%d] getArchitectScheduleOK  %+v", 200, o.Payload)
}

func (o *GetArchitectScheduleOK) GetPayload() *models.Schedule {
	return o.Payload
}

func (o *GetArchitectScheduleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Schedule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectScheduleBadRequest creates a GetArchitectScheduleBadRequest with default headers values
func NewGetArchitectScheduleBadRequest() *GetArchitectScheduleBadRequest {
	return &GetArchitectScheduleBadRequest{}
}

/*
GetArchitectScheduleBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetArchitectScheduleBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get architect schedule bad request response has a 2xx status code
func (o *GetArchitectScheduleBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get architect schedule bad request response has a 3xx status code
func (o *GetArchitectScheduleBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get architect schedule bad request response has a 4xx status code
func (o *GetArchitectScheduleBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get architect schedule bad request response has a 5xx status code
func (o *GetArchitectScheduleBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get architect schedule bad request response a status code equal to that given
func (o *GetArchitectScheduleBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetArchitectScheduleBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/schedules/{scheduleId}][%d] getArchitectScheduleBadRequest  %+v", 400, o.Payload)
}

func (o *GetArchitectScheduleBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/architect/schedules/{scheduleId}][%d] getArchitectScheduleBadRequest  %+v", 400, o.Payload)
}

func (o *GetArchitectScheduleBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectScheduleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectScheduleUnauthorized creates a GetArchitectScheduleUnauthorized with default headers values
func NewGetArchitectScheduleUnauthorized() *GetArchitectScheduleUnauthorized {
	return &GetArchitectScheduleUnauthorized{}
}

/*
GetArchitectScheduleUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetArchitectScheduleUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get architect schedule unauthorized response has a 2xx status code
func (o *GetArchitectScheduleUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get architect schedule unauthorized response has a 3xx status code
func (o *GetArchitectScheduleUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get architect schedule unauthorized response has a 4xx status code
func (o *GetArchitectScheduleUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get architect schedule unauthorized response has a 5xx status code
func (o *GetArchitectScheduleUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get architect schedule unauthorized response a status code equal to that given
func (o *GetArchitectScheduleUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetArchitectScheduleUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/schedules/{scheduleId}][%d] getArchitectScheduleUnauthorized  %+v", 401, o.Payload)
}

func (o *GetArchitectScheduleUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/architect/schedules/{scheduleId}][%d] getArchitectScheduleUnauthorized  %+v", 401, o.Payload)
}

func (o *GetArchitectScheduleUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectScheduleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectScheduleForbidden creates a GetArchitectScheduleForbidden with default headers values
func NewGetArchitectScheduleForbidden() *GetArchitectScheduleForbidden {
	return &GetArchitectScheduleForbidden{}
}

/*
GetArchitectScheduleForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetArchitectScheduleForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get architect schedule forbidden response has a 2xx status code
func (o *GetArchitectScheduleForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get architect schedule forbidden response has a 3xx status code
func (o *GetArchitectScheduleForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get architect schedule forbidden response has a 4xx status code
func (o *GetArchitectScheduleForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get architect schedule forbidden response has a 5xx status code
func (o *GetArchitectScheduleForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get architect schedule forbidden response a status code equal to that given
func (o *GetArchitectScheduleForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetArchitectScheduleForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/schedules/{scheduleId}][%d] getArchitectScheduleForbidden  %+v", 403, o.Payload)
}

func (o *GetArchitectScheduleForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/architect/schedules/{scheduleId}][%d] getArchitectScheduleForbidden  %+v", 403, o.Payload)
}

func (o *GetArchitectScheduleForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectScheduleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectScheduleNotFound creates a GetArchitectScheduleNotFound with default headers values
func NewGetArchitectScheduleNotFound() *GetArchitectScheduleNotFound {
	return &GetArchitectScheduleNotFound{}
}

/*
GetArchitectScheduleNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetArchitectScheduleNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get architect schedule not found response has a 2xx status code
func (o *GetArchitectScheduleNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get architect schedule not found response has a 3xx status code
func (o *GetArchitectScheduleNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get architect schedule not found response has a 4xx status code
func (o *GetArchitectScheduleNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get architect schedule not found response has a 5xx status code
func (o *GetArchitectScheduleNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get architect schedule not found response a status code equal to that given
func (o *GetArchitectScheduleNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetArchitectScheduleNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/schedules/{scheduleId}][%d] getArchitectScheduleNotFound  %+v", 404, o.Payload)
}

func (o *GetArchitectScheduleNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/architect/schedules/{scheduleId}][%d] getArchitectScheduleNotFound  %+v", 404, o.Payload)
}

func (o *GetArchitectScheduleNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectScheduleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectScheduleRequestTimeout creates a GetArchitectScheduleRequestTimeout with default headers values
func NewGetArchitectScheduleRequestTimeout() *GetArchitectScheduleRequestTimeout {
	return &GetArchitectScheduleRequestTimeout{}
}

/*
GetArchitectScheduleRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetArchitectScheduleRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get architect schedule request timeout response has a 2xx status code
func (o *GetArchitectScheduleRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get architect schedule request timeout response has a 3xx status code
func (o *GetArchitectScheduleRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get architect schedule request timeout response has a 4xx status code
func (o *GetArchitectScheduleRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get architect schedule request timeout response has a 5xx status code
func (o *GetArchitectScheduleRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get architect schedule request timeout response a status code equal to that given
func (o *GetArchitectScheduleRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetArchitectScheduleRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/schedules/{scheduleId}][%d] getArchitectScheduleRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetArchitectScheduleRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/architect/schedules/{scheduleId}][%d] getArchitectScheduleRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetArchitectScheduleRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectScheduleRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectScheduleRequestEntityTooLarge creates a GetArchitectScheduleRequestEntityTooLarge with default headers values
func NewGetArchitectScheduleRequestEntityTooLarge() *GetArchitectScheduleRequestEntityTooLarge {
	return &GetArchitectScheduleRequestEntityTooLarge{}
}

/*
GetArchitectScheduleRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetArchitectScheduleRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get architect schedule request entity too large response has a 2xx status code
func (o *GetArchitectScheduleRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get architect schedule request entity too large response has a 3xx status code
func (o *GetArchitectScheduleRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get architect schedule request entity too large response has a 4xx status code
func (o *GetArchitectScheduleRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get architect schedule request entity too large response has a 5xx status code
func (o *GetArchitectScheduleRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get architect schedule request entity too large response a status code equal to that given
func (o *GetArchitectScheduleRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetArchitectScheduleRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/schedules/{scheduleId}][%d] getArchitectScheduleRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetArchitectScheduleRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/architect/schedules/{scheduleId}][%d] getArchitectScheduleRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetArchitectScheduleRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectScheduleRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectScheduleUnsupportedMediaType creates a GetArchitectScheduleUnsupportedMediaType with default headers values
func NewGetArchitectScheduleUnsupportedMediaType() *GetArchitectScheduleUnsupportedMediaType {
	return &GetArchitectScheduleUnsupportedMediaType{}
}

/*
GetArchitectScheduleUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetArchitectScheduleUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get architect schedule unsupported media type response has a 2xx status code
func (o *GetArchitectScheduleUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get architect schedule unsupported media type response has a 3xx status code
func (o *GetArchitectScheduleUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get architect schedule unsupported media type response has a 4xx status code
func (o *GetArchitectScheduleUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get architect schedule unsupported media type response has a 5xx status code
func (o *GetArchitectScheduleUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get architect schedule unsupported media type response a status code equal to that given
func (o *GetArchitectScheduleUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetArchitectScheduleUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/schedules/{scheduleId}][%d] getArchitectScheduleUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetArchitectScheduleUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/architect/schedules/{scheduleId}][%d] getArchitectScheduleUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetArchitectScheduleUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectScheduleUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectScheduleTooManyRequests creates a GetArchitectScheduleTooManyRequests with default headers values
func NewGetArchitectScheduleTooManyRequests() *GetArchitectScheduleTooManyRequests {
	return &GetArchitectScheduleTooManyRequests{}
}

/*
GetArchitectScheduleTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetArchitectScheduleTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get architect schedule too many requests response has a 2xx status code
func (o *GetArchitectScheduleTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get architect schedule too many requests response has a 3xx status code
func (o *GetArchitectScheduleTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get architect schedule too many requests response has a 4xx status code
func (o *GetArchitectScheduleTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get architect schedule too many requests response has a 5xx status code
func (o *GetArchitectScheduleTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get architect schedule too many requests response a status code equal to that given
func (o *GetArchitectScheduleTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetArchitectScheduleTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/schedules/{scheduleId}][%d] getArchitectScheduleTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetArchitectScheduleTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/architect/schedules/{scheduleId}][%d] getArchitectScheduleTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetArchitectScheduleTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectScheduleTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectScheduleInternalServerError creates a GetArchitectScheduleInternalServerError with default headers values
func NewGetArchitectScheduleInternalServerError() *GetArchitectScheduleInternalServerError {
	return &GetArchitectScheduleInternalServerError{}
}

/*
GetArchitectScheduleInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetArchitectScheduleInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get architect schedule internal server error response has a 2xx status code
func (o *GetArchitectScheduleInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get architect schedule internal server error response has a 3xx status code
func (o *GetArchitectScheduleInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get architect schedule internal server error response has a 4xx status code
func (o *GetArchitectScheduleInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get architect schedule internal server error response has a 5xx status code
func (o *GetArchitectScheduleInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get architect schedule internal server error response a status code equal to that given
func (o *GetArchitectScheduleInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetArchitectScheduleInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/schedules/{scheduleId}][%d] getArchitectScheduleInternalServerError  %+v", 500, o.Payload)
}

func (o *GetArchitectScheduleInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/architect/schedules/{scheduleId}][%d] getArchitectScheduleInternalServerError  %+v", 500, o.Payload)
}

func (o *GetArchitectScheduleInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectScheduleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectScheduleServiceUnavailable creates a GetArchitectScheduleServiceUnavailable with default headers values
func NewGetArchitectScheduleServiceUnavailable() *GetArchitectScheduleServiceUnavailable {
	return &GetArchitectScheduleServiceUnavailable{}
}

/*
GetArchitectScheduleServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetArchitectScheduleServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get architect schedule service unavailable response has a 2xx status code
func (o *GetArchitectScheduleServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get architect schedule service unavailable response has a 3xx status code
func (o *GetArchitectScheduleServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get architect schedule service unavailable response has a 4xx status code
func (o *GetArchitectScheduleServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get architect schedule service unavailable response has a 5xx status code
func (o *GetArchitectScheduleServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get architect schedule service unavailable response a status code equal to that given
func (o *GetArchitectScheduleServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetArchitectScheduleServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/schedules/{scheduleId}][%d] getArchitectScheduleServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetArchitectScheduleServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/architect/schedules/{scheduleId}][%d] getArchitectScheduleServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetArchitectScheduleServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectScheduleServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectScheduleGatewayTimeout creates a GetArchitectScheduleGatewayTimeout with default headers values
func NewGetArchitectScheduleGatewayTimeout() *GetArchitectScheduleGatewayTimeout {
	return &GetArchitectScheduleGatewayTimeout{}
}

/*
GetArchitectScheduleGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetArchitectScheduleGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get architect schedule gateway timeout response has a 2xx status code
func (o *GetArchitectScheduleGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get architect schedule gateway timeout response has a 3xx status code
func (o *GetArchitectScheduleGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get architect schedule gateway timeout response has a 4xx status code
func (o *GetArchitectScheduleGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get architect schedule gateway timeout response has a 5xx status code
func (o *GetArchitectScheduleGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get architect schedule gateway timeout response a status code equal to that given
func (o *GetArchitectScheduleGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetArchitectScheduleGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/schedules/{scheduleId}][%d] getArchitectScheduleGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetArchitectScheduleGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/architect/schedules/{scheduleId}][%d] getArchitectScheduleGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetArchitectScheduleGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectScheduleGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

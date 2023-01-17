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

// DeleteArchitectScheduleReader is a Reader for the DeleteArchitectSchedule structure.
type DeleteArchitectScheduleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteArchitectScheduleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteArchitectScheduleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteArchitectScheduleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteArchitectScheduleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteArchitectScheduleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteArchitectScheduleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewDeleteArchitectScheduleRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDeleteArchitectScheduleConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteArchitectScheduleRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteArchitectScheduleUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteArchitectScheduleTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteArchitectScheduleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteArchitectScheduleServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteArchitectScheduleGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteArchitectScheduleOK creates a DeleteArchitectScheduleOK with default headers values
func NewDeleteArchitectScheduleOK() *DeleteArchitectScheduleOK {
	return &DeleteArchitectScheduleOK{}
}

/*
DeleteArchitectScheduleOK describes a response with status code 200, with default header values.

Operation was successful.
*/
type DeleteArchitectScheduleOK struct {
}

// IsSuccess returns true when this delete architect schedule o k response has a 2xx status code
func (o *DeleteArchitectScheduleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete architect schedule o k response has a 3xx status code
func (o *DeleteArchitectScheduleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete architect schedule o k response has a 4xx status code
func (o *DeleteArchitectScheduleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete architect schedule o k response has a 5xx status code
func (o *DeleteArchitectScheduleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete architect schedule o k response a status code equal to that given
func (o *DeleteArchitectScheduleOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteArchitectScheduleOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/architect/schedules/{scheduleId}][%d] deleteArchitectScheduleOK ", 200)
}

func (o *DeleteArchitectScheduleOK) String() string {
	return fmt.Sprintf("[DELETE /api/v2/architect/schedules/{scheduleId}][%d] deleteArchitectScheduleOK ", 200)
}

func (o *DeleteArchitectScheduleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteArchitectScheduleBadRequest creates a DeleteArchitectScheduleBadRequest with default headers values
func NewDeleteArchitectScheduleBadRequest() *DeleteArchitectScheduleBadRequest {
	return &DeleteArchitectScheduleBadRequest{}
}

/*
DeleteArchitectScheduleBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteArchitectScheduleBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete architect schedule bad request response has a 2xx status code
func (o *DeleteArchitectScheduleBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete architect schedule bad request response has a 3xx status code
func (o *DeleteArchitectScheduleBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete architect schedule bad request response has a 4xx status code
func (o *DeleteArchitectScheduleBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete architect schedule bad request response has a 5xx status code
func (o *DeleteArchitectScheduleBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete architect schedule bad request response a status code equal to that given
func (o *DeleteArchitectScheduleBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DeleteArchitectScheduleBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/architect/schedules/{scheduleId}][%d] deleteArchitectScheduleBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteArchitectScheduleBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api/v2/architect/schedules/{scheduleId}][%d] deleteArchitectScheduleBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteArchitectScheduleBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteArchitectScheduleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteArchitectScheduleUnauthorized creates a DeleteArchitectScheduleUnauthorized with default headers values
func NewDeleteArchitectScheduleUnauthorized() *DeleteArchitectScheduleUnauthorized {
	return &DeleteArchitectScheduleUnauthorized{}
}

/*
DeleteArchitectScheduleUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteArchitectScheduleUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete architect schedule unauthorized response has a 2xx status code
func (o *DeleteArchitectScheduleUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete architect schedule unauthorized response has a 3xx status code
func (o *DeleteArchitectScheduleUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete architect schedule unauthorized response has a 4xx status code
func (o *DeleteArchitectScheduleUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete architect schedule unauthorized response has a 5xx status code
func (o *DeleteArchitectScheduleUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete architect schedule unauthorized response a status code equal to that given
func (o *DeleteArchitectScheduleUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteArchitectScheduleUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/architect/schedules/{scheduleId}][%d] deleteArchitectScheduleUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteArchitectScheduleUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/v2/architect/schedules/{scheduleId}][%d] deleteArchitectScheduleUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteArchitectScheduleUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteArchitectScheduleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteArchitectScheduleForbidden creates a DeleteArchitectScheduleForbidden with default headers values
func NewDeleteArchitectScheduleForbidden() *DeleteArchitectScheduleForbidden {
	return &DeleteArchitectScheduleForbidden{}
}

/*
DeleteArchitectScheduleForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type DeleteArchitectScheduleForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete architect schedule forbidden response has a 2xx status code
func (o *DeleteArchitectScheduleForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete architect schedule forbidden response has a 3xx status code
func (o *DeleteArchitectScheduleForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete architect schedule forbidden response has a 4xx status code
func (o *DeleteArchitectScheduleForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete architect schedule forbidden response has a 5xx status code
func (o *DeleteArchitectScheduleForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete architect schedule forbidden response a status code equal to that given
func (o *DeleteArchitectScheduleForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteArchitectScheduleForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/architect/schedules/{scheduleId}][%d] deleteArchitectScheduleForbidden  %+v", 403, o.Payload)
}

func (o *DeleteArchitectScheduleForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v2/architect/schedules/{scheduleId}][%d] deleteArchitectScheduleForbidden  %+v", 403, o.Payload)
}

func (o *DeleteArchitectScheduleForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteArchitectScheduleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteArchitectScheduleNotFound creates a DeleteArchitectScheduleNotFound with default headers values
func NewDeleteArchitectScheduleNotFound() *DeleteArchitectScheduleNotFound {
	return &DeleteArchitectScheduleNotFound{}
}

/*
DeleteArchitectScheduleNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type DeleteArchitectScheduleNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete architect schedule not found response has a 2xx status code
func (o *DeleteArchitectScheduleNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete architect schedule not found response has a 3xx status code
func (o *DeleteArchitectScheduleNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete architect schedule not found response has a 4xx status code
func (o *DeleteArchitectScheduleNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete architect schedule not found response has a 5xx status code
func (o *DeleteArchitectScheduleNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete architect schedule not found response a status code equal to that given
func (o *DeleteArchitectScheduleNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteArchitectScheduleNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/architect/schedules/{scheduleId}][%d] deleteArchitectScheduleNotFound  %+v", 404, o.Payload)
}

func (o *DeleteArchitectScheduleNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/v2/architect/schedules/{scheduleId}][%d] deleteArchitectScheduleNotFound  %+v", 404, o.Payload)
}

func (o *DeleteArchitectScheduleNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteArchitectScheduleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteArchitectScheduleRequestTimeout creates a DeleteArchitectScheduleRequestTimeout with default headers values
func NewDeleteArchitectScheduleRequestTimeout() *DeleteArchitectScheduleRequestTimeout {
	return &DeleteArchitectScheduleRequestTimeout{}
}

/*
DeleteArchitectScheduleRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type DeleteArchitectScheduleRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete architect schedule request timeout response has a 2xx status code
func (o *DeleteArchitectScheduleRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete architect schedule request timeout response has a 3xx status code
func (o *DeleteArchitectScheduleRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete architect schedule request timeout response has a 4xx status code
func (o *DeleteArchitectScheduleRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete architect schedule request timeout response has a 5xx status code
func (o *DeleteArchitectScheduleRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this delete architect schedule request timeout response a status code equal to that given
func (o *DeleteArchitectScheduleRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *DeleteArchitectScheduleRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/architect/schedules/{scheduleId}][%d] deleteArchitectScheduleRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteArchitectScheduleRequestTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/architect/schedules/{scheduleId}][%d] deleteArchitectScheduleRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteArchitectScheduleRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteArchitectScheduleRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteArchitectScheduleConflict creates a DeleteArchitectScheduleConflict with default headers values
func NewDeleteArchitectScheduleConflict() *DeleteArchitectScheduleConflict {
	return &DeleteArchitectScheduleConflict{}
}

/*
DeleteArchitectScheduleConflict describes a response with status code 409, with default header values.

Conflict
*/
type DeleteArchitectScheduleConflict struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete architect schedule conflict response has a 2xx status code
func (o *DeleteArchitectScheduleConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete architect schedule conflict response has a 3xx status code
func (o *DeleteArchitectScheduleConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete architect schedule conflict response has a 4xx status code
func (o *DeleteArchitectScheduleConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete architect schedule conflict response has a 5xx status code
func (o *DeleteArchitectScheduleConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this delete architect schedule conflict response a status code equal to that given
func (o *DeleteArchitectScheduleConflict) IsCode(code int) bool {
	return code == 409
}

func (o *DeleteArchitectScheduleConflict) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/architect/schedules/{scheduleId}][%d] deleteArchitectScheduleConflict  %+v", 409, o.Payload)
}

func (o *DeleteArchitectScheduleConflict) String() string {
	return fmt.Sprintf("[DELETE /api/v2/architect/schedules/{scheduleId}][%d] deleteArchitectScheduleConflict  %+v", 409, o.Payload)
}

func (o *DeleteArchitectScheduleConflict) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteArchitectScheduleConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteArchitectScheduleRequestEntityTooLarge creates a DeleteArchitectScheduleRequestEntityTooLarge with default headers values
func NewDeleteArchitectScheduleRequestEntityTooLarge() *DeleteArchitectScheduleRequestEntityTooLarge {
	return &DeleteArchitectScheduleRequestEntityTooLarge{}
}

/*
DeleteArchitectScheduleRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type DeleteArchitectScheduleRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete architect schedule request entity too large response has a 2xx status code
func (o *DeleteArchitectScheduleRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete architect schedule request entity too large response has a 3xx status code
func (o *DeleteArchitectScheduleRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete architect schedule request entity too large response has a 4xx status code
func (o *DeleteArchitectScheduleRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete architect schedule request entity too large response has a 5xx status code
func (o *DeleteArchitectScheduleRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this delete architect schedule request entity too large response a status code equal to that given
func (o *DeleteArchitectScheduleRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *DeleteArchitectScheduleRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/architect/schedules/{scheduleId}][%d] deleteArchitectScheduleRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteArchitectScheduleRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[DELETE /api/v2/architect/schedules/{scheduleId}][%d] deleteArchitectScheduleRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteArchitectScheduleRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteArchitectScheduleRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteArchitectScheduleUnsupportedMediaType creates a DeleteArchitectScheduleUnsupportedMediaType with default headers values
func NewDeleteArchitectScheduleUnsupportedMediaType() *DeleteArchitectScheduleUnsupportedMediaType {
	return &DeleteArchitectScheduleUnsupportedMediaType{}
}

/*
DeleteArchitectScheduleUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteArchitectScheduleUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete architect schedule unsupported media type response has a 2xx status code
func (o *DeleteArchitectScheduleUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete architect schedule unsupported media type response has a 3xx status code
func (o *DeleteArchitectScheduleUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete architect schedule unsupported media type response has a 4xx status code
func (o *DeleteArchitectScheduleUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete architect schedule unsupported media type response has a 5xx status code
func (o *DeleteArchitectScheduleUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this delete architect schedule unsupported media type response a status code equal to that given
func (o *DeleteArchitectScheduleUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *DeleteArchitectScheduleUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/architect/schedules/{scheduleId}][%d] deleteArchitectScheduleUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteArchitectScheduleUnsupportedMediaType) String() string {
	return fmt.Sprintf("[DELETE /api/v2/architect/schedules/{scheduleId}][%d] deleteArchitectScheduleUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteArchitectScheduleUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteArchitectScheduleUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteArchitectScheduleTooManyRequests creates a DeleteArchitectScheduleTooManyRequests with default headers values
func NewDeleteArchitectScheduleTooManyRequests() *DeleteArchitectScheduleTooManyRequests {
	return &DeleteArchitectScheduleTooManyRequests{}
}

/*
DeleteArchitectScheduleTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type DeleteArchitectScheduleTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete architect schedule too many requests response has a 2xx status code
func (o *DeleteArchitectScheduleTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete architect schedule too many requests response has a 3xx status code
func (o *DeleteArchitectScheduleTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete architect schedule too many requests response has a 4xx status code
func (o *DeleteArchitectScheduleTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete architect schedule too many requests response has a 5xx status code
func (o *DeleteArchitectScheduleTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this delete architect schedule too many requests response a status code equal to that given
func (o *DeleteArchitectScheduleTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *DeleteArchitectScheduleTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/architect/schedules/{scheduleId}][%d] deleteArchitectScheduleTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteArchitectScheduleTooManyRequests) String() string {
	return fmt.Sprintf("[DELETE /api/v2/architect/schedules/{scheduleId}][%d] deleteArchitectScheduleTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteArchitectScheduleTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteArchitectScheduleTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteArchitectScheduleInternalServerError creates a DeleteArchitectScheduleInternalServerError with default headers values
func NewDeleteArchitectScheduleInternalServerError() *DeleteArchitectScheduleInternalServerError {
	return &DeleteArchitectScheduleInternalServerError{}
}

/*
DeleteArchitectScheduleInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteArchitectScheduleInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete architect schedule internal server error response has a 2xx status code
func (o *DeleteArchitectScheduleInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete architect schedule internal server error response has a 3xx status code
func (o *DeleteArchitectScheduleInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete architect schedule internal server error response has a 4xx status code
func (o *DeleteArchitectScheduleInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete architect schedule internal server error response has a 5xx status code
func (o *DeleteArchitectScheduleInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete architect schedule internal server error response a status code equal to that given
func (o *DeleteArchitectScheduleInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DeleteArchitectScheduleInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/architect/schedules/{scheduleId}][%d] deleteArchitectScheduleInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteArchitectScheduleInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /api/v2/architect/schedules/{scheduleId}][%d] deleteArchitectScheduleInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteArchitectScheduleInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteArchitectScheduleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteArchitectScheduleServiceUnavailable creates a DeleteArchitectScheduleServiceUnavailable with default headers values
func NewDeleteArchitectScheduleServiceUnavailable() *DeleteArchitectScheduleServiceUnavailable {
	return &DeleteArchitectScheduleServiceUnavailable{}
}

/*
DeleteArchitectScheduleServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteArchitectScheduleServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete architect schedule service unavailable response has a 2xx status code
func (o *DeleteArchitectScheduleServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete architect schedule service unavailable response has a 3xx status code
func (o *DeleteArchitectScheduleServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete architect schedule service unavailable response has a 4xx status code
func (o *DeleteArchitectScheduleServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete architect schedule service unavailable response has a 5xx status code
func (o *DeleteArchitectScheduleServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this delete architect schedule service unavailable response a status code equal to that given
func (o *DeleteArchitectScheduleServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *DeleteArchitectScheduleServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/architect/schedules/{scheduleId}][%d] deleteArchitectScheduleServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteArchitectScheduleServiceUnavailable) String() string {
	return fmt.Sprintf("[DELETE /api/v2/architect/schedules/{scheduleId}][%d] deleteArchitectScheduleServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteArchitectScheduleServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteArchitectScheduleServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteArchitectScheduleGatewayTimeout creates a DeleteArchitectScheduleGatewayTimeout with default headers values
func NewDeleteArchitectScheduleGatewayTimeout() *DeleteArchitectScheduleGatewayTimeout {
	return &DeleteArchitectScheduleGatewayTimeout{}
}

/*
DeleteArchitectScheduleGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type DeleteArchitectScheduleGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete architect schedule gateway timeout response has a 2xx status code
func (o *DeleteArchitectScheduleGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete architect schedule gateway timeout response has a 3xx status code
func (o *DeleteArchitectScheduleGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete architect schedule gateway timeout response has a 4xx status code
func (o *DeleteArchitectScheduleGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete architect schedule gateway timeout response has a 5xx status code
func (o *DeleteArchitectScheduleGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this delete architect schedule gateway timeout response a status code equal to that given
func (o *DeleteArchitectScheduleGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *DeleteArchitectScheduleGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/architect/schedules/{scheduleId}][%d] deleteArchitectScheduleGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteArchitectScheduleGatewayTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/architect/schedules/{scheduleId}][%d] deleteArchitectScheduleGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteArchitectScheduleGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteArchitectScheduleGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

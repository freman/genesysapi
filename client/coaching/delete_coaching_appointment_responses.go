// Code generated by go-swagger; DO NOT EDIT.

package coaching

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// DeleteCoachingAppointmentReader is a Reader for the DeleteCoachingAppointment structure.
type DeleteCoachingAppointmentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCoachingAppointmentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewDeleteCoachingAppointmentAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeleteCoachingAppointmentNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteCoachingAppointmentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteCoachingAppointmentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteCoachingAppointmentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteCoachingAppointmentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewDeleteCoachingAppointmentRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDeleteCoachingAppointmentConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteCoachingAppointmentRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteCoachingAppointmentUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteCoachingAppointmentTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteCoachingAppointmentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteCoachingAppointmentServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteCoachingAppointmentGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteCoachingAppointmentAccepted creates a DeleteCoachingAppointmentAccepted with default headers values
func NewDeleteCoachingAppointmentAccepted() *DeleteCoachingAppointmentAccepted {
	return &DeleteCoachingAppointmentAccepted{}
}

/*
DeleteCoachingAppointmentAccepted describes a response with status code 202, with default header values.

Appointment delete request accepted.
*/
type DeleteCoachingAppointmentAccepted struct {
	Payload *models.CoachingAppointmentReference
}

// IsSuccess returns true when this delete coaching appointment accepted response has a 2xx status code
func (o *DeleteCoachingAppointmentAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete coaching appointment accepted response has a 3xx status code
func (o *DeleteCoachingAppointmentAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete coaching appointment accepted response has a 4xx status code
func (o *DeleteCoachingAppointmentAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete coaching appointment accepted response has a 5xx status code
func (o *DeleteCoachingAppointmentAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this delete coaching appointment accepted response a status code equal to that given
func (o *DeleteCoachingAppointmentAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *DeleteCoachingAppointmentAccepted) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/coaching/appointments/{appointmentId}][%d] deleteCoachingAppointmentAccepted  %+v", 202, o.Payload)
}

func (o *DeleteCoachingAppointmentAccepted) String() string {
	return fmt.Sprintf("[DELETE /api/v2/coaching/appointments/{appointmentId}][%d] deleteCoachingAppointmentAccepted  %+v", 202, o.Payload)
}

func (o *DeleteCoachingAppointmentAccepted) GetPayload() *models.CoachingAppointmentReference {
	return o.Payload
}

func (o *DeleteCoachingAppointmentAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CoachingAppointmentReference)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCoachingAppointmentNoContent creates a DeleteCoachingAppointmentNoContent with default headers values
func NewDeleteCoachingAppointmentNoContent() *DeleteCoachingAppointmentNoContent {
	return &DeleteCoachingAppointmentNoContent{}
}

/*
DeleteCoachingAppointmentNoContent describes a response with status code 204, with default header values.

Appointment deleted
*/
type DeleteCoachingAppointmentNoContent struct {
}

// IsSuccess returns true when this delete coaching appointment no content response has a 2xx status code
func (o *DeleteCoachingAppointmentNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete coaching appointment no content response has a 3xx status code
func (o *DeleteCoachingAppointmentNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete coaching appointment no content response has a 4xx status code
func (o *DeleteCoachingAppointmentNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete coaching appointment no content response has a 5xx status code
func (o *DeleteCoachingAppointmentNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete coaching appointment no content response a status code equal to that given
func (o *DeleteCoachingAppointmentNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DeleteCoachingAppointmentNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/coaching/appointments/{appointmentId}][%d] deleteCoachingAppointmentNoContent ", 204)
}

func (o *DeleteCoachingAppointmentNoContent) String() string {
	return fmt.Sprintf("[DELETE /api/v2/coaching/appointments/{appointmentId}][%d] deleteCoachingAppointmentNoContent ", 204)
}

func (o *DeleteCoachingAppointmentNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteCoachingAppointmentBadRequest creates a DeleteCoachingAppointmentBadRequest with default headers values
func NewDeleteCoachingAppointmentBadRequest() *DeleteCoachingAppointmentBadRequest {
	return &DeleteCoachingAppointmentBadRequest{}
}

/*
DeleteCoachingAppointmentBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteCoachingAppointmentBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete coaching appointment bad request response has a 2xx status code
func (o *DeleteCoachingAppointmentBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete coaching appointment bad request response has a 3xx status code
func (o *DeleteCoachingAppointmentBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete coaching appointment bad request response has a 4xx status code
func (o *DeleteCoachingAppointmentBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete coaching appointment bad request response has a 5xx status code
func (o *DeleteCoachingAppointmentBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete coaching appointment bad request response a status code equal to that given
func (o *DeleteCoachingAppointmentBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DeleteCoachingAppointmentBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/coaching/appointments/{appointmentId}][%d] deleteCoachingAppointmentBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteCoachingAppointmentBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api/v2/coaching/appointments/{appointmentId}][%d] deleteCoachingAppointmentBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteCoachingAppointmentBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteCoachingAppointmentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCoachingAppointmentUnauthorized creates a DeleteCoachingAppointmentUnauthorized with default headers values
func NewDeleteCoachingAppointmentUnauthorized() *DeleteCoachingAppointmentUnauthorized {
	return &DeleteCoachingAppointmentUnauthorized{}
}

/*
DeleteCoachingAppointmentUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteCoachingAppointmentUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete coaching appointment unauthorized response has a 2xx status code
func (o *DeleteCoachingAppointmentUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete coaching appointment unauthorized response has a 3xx status code
func (o *DeleteCoachingAppointmentUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete coaching appointment unauthorized response has a 4xx status code
func (o *DeleteCoachingAppointmentUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete coaching appointment unauthorized response has a 5xx status code
func (o *DeleteCoachingAppointmentUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete coaching appointment unauthorized response a status code equal to that given
func (o *DeleteCoachingAppointmentUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteCoachingAppointmentUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/coaching/appointments/{appointmentId}][%d] deleteCoachingAppointmentUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteCoachingAppointmentUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/v2/coaching/appointments/{appointmentId}][%d] deleteCoachingAppointmentUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteCoachingAppointmentUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteCoachingAppointmentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCoachingAppointmentForbidden creates a DeleteCoachingAppointmentForbidden with default headers values
func NewDeleteCoachingAppointmentForbidden() *DeleteCoachingAppointmentForbidden {
	return &DeleteCoachingAppointmentForbidden{}
}

/*
DeleteCoachingAppointmentForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type DeleteCoachingAppointmentForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete coaching appointment forbidden response has a 2xx status code
func (o *DeleteCoachingAppointmentForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete coaching appointment forbidden response has a 3xx status code
func (o *DeleteCoachingAppointmentForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete coaching appointment forbidden response has a 4xx status code
func (o *DeleteCoachingAppointmentForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete coaching appointment forbidden response has a 5xx status code
func (o *DeleteCoachingAppointmentForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete coaching appointment forbidden response a status code equal to that given
func (o *DeleteCoachingAppointmentForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteCoachingAppointmentForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/coaching/appointments/{appointmentId}][%d] deleteCoachingAppointmentForbidden  %+v", 403, o.Payload)
}

func (o *DeleteCoachingAppointmentForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v2/coaching/appointments/{appointmentId}][%d] deleteCoachingAppointmentForbidden  %+v", 403, o.Payload)
}

func (o *DeleteCoachingAppointmentForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteCoachingAppointmentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCoachingAppointmentNotFound creates a DeleteCoachingAppointmentNotFound with default headers values
func NewDeleteCoachingAppointmentNotFound() *DeleteCoachingAppointmentNotFound {
	return &DeleteCoachingAppointmentNotFound{}
}

/*
DeleteCoachingAppointmentNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type DeleteCoachingAppointmentNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete coaching appointment not found response has a 2xx status code
func (o *DeleteCoachingAppointmentNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete coaching appointment not found response has a 3xx status code
func (o *DeleteCoachingAppointmentNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete coaching appointment not found response has a 4xx status code
func (o *DeleteCoachingAppointmentNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete coaching appointment not found response has a 5xx status code
func (o *DeleteCoachingAppointmentNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete coaching appointment not found response a status code equal to that given
func (o *DeleteCoachingAppointmentNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteCoachingAppointmentNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/coaching/appointments/{appointmentId}][%d] deleteCoachingAppointmentNotFound  %+v", 404, o.Payload)
}

func (o *DeleteCoachingAppointmentNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/v2/coaching/appointments/{appointmentId}][%d] deleteCoachingAppointmentNotFound  %+v", 404, o.Payload)
}

func (o *DeleteCoachingAppointmentNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteCoachingAppointmentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCoachingAppointmentRequestTimeout creates a DeleteCoachingAppointmentRequestTimeout with default headers values
func NewDeleteCoachingAppointmentRequestTimeout() *DeleteCoachingAppointmentRequestTimeout {
	return &DeleteCoachingAppointmentRequestTimeout{}
}

/*
DeleteCoachingAppointmentRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type DeleteCoachingAppointmentRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete coaching appointment request timeout response has a 2xx status code
func (o *DeleteCoachingAppointmentRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete coaching appointment request timeout response has a 3xx status code
func (o *DeleteCoachingAppointmentRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete coaching appointment request timeout response has a 4xx status code
func (o *DeleteCoachingAppointmentRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete coaching appointment request timeout response has a 5xx status code
func (o *DeleteCoachingAppointmentRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this delete coaching appointment request timeout response a status code equal to that given
func (o *DeleteCoachingAppointmentRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *DeleteCoachingAppointmentRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/coaching/appointments/{appointmentId}][%d] deleteCoachingAppointmentRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteCoachingAppointmentRequestTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/coaching/appointments/{appointmentId}][%d] deleteCoachingAppointmentRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteCoachingAppointmentRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteCoachingAppointmentRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCoachingAppointmentConflict creates a DeleteCoachingAppointmentConflict with default headers values
func NewDeleteCoachingAppointmentConflict() *DeleteCoachingAppointmentConflict {
	return &DeleteCoachingAppointmentConflict{}
}

/*
DeleteCoachingAppointmentConflict describes a response with status code 409, with default header values.

Conflict
*/
type DeleteCoachingAppointmentConflict struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete coaching appointment conflict response has a 2xx status code
func (o *DeleteCoachingAppointmentConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete coaching appointment conflict response has a 3xx status code
func (o *DeleteCoachingAppointmentConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete coaching appointment conflict response has a 4xx status code
func (o *DeleteCoachingAppointmentConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete coaching appointment conflict response has a 5xx status code
func (o *DeleteCoachingAppointmentConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this delete coaching appointment conflict response a status code equal to that given
func (o *DeleteCoachingAppointmentConflict) IsCode(code int) bool {
	return code == 409
}

func (o *DeleteCoachingAppointmentConflict) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/coaching/appointments/{appointmentId}][%d] deleteCoachingAppointmentConflict  %+v", 409, o.Payload)
}

func (o *DeleteCoachingAppointmentConflict) String() string {
	return fmt.Sprintf("[DELETE /api/v2/coaching/appointments/{appointmentId}][%d] deleteCoachingAppointmentConflict  %+v", 409, o.Payload)
}

func (o *DeleteCoachingAppointmentConflict) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteCoachingAppointmentConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCoachingAppointmentRequestEntityTooLarge creates a DeleteCoachingAppointmentRequestEntityTooLarge with default headers values
func NewDeleteCoachingAppointmentRequestEntityTooLarge() *DeleteCoachingAppointmentRequestEntityTooLarge {
	return &DeleteCoachingAppointmentRequestEntityTooLarge{}
}

/*
DeleteCoachingAppointmentRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type DeleteCoachingAppointmentRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete coaching appointment request entity too large response has a 2xx status code
func (o *DeleteCoachingAppointmentRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete coaching appointment request entity too large response has a 3xx status code
func (o *DeleteCoachingAppointmentRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete coaching appointment request entity too large response has a 4xx status code
func (o *DeleteCoachingAppointmentRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete coaching appointment request entity too large response has a 5xx status code
func (o *DeleteCoachingAppointmentRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this delete coaching appointment request entity too large response a status code equal to that given
func (o *DeleteCoachingAppointmentRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *DeleteCoachingAppointmentRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/coaching/appointments/{appointmentId}][%d] deleteCoachingAppointmentRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteCoachingAppointmentRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[DELETE /api/v2/coaching/appointments/{appointmentId}][%d] deleteCoachingAppointmentRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteCoachingAppointmentRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteCoachingAppointmentRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCoachingAppointmentUnsupportedMediaType creates a DeleteCoachingAppointmentUnsupportedMediaType with default headers values
func NewDeleteCoachingAppointmentUnsupportedMediaType() *DeleteCoachingAppointmentUnsupportedMediaType {
	return &DeleteCoachingAppointmentUnsupportedMediaType{}
}

/*
DeleteCoachingAppointmentUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteCoachingAppointmentUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete coaching appointment unsupported media type response has a 2xx status code
func (o *DeleteCoachingAppointmentUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete coaching appointment unsupported media type response has a 3xx status code
func (o *DeleteCoachingAppointmentUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete coaching appointment unsupported media type response has a 4xx status code
func (o *DeleteCoachingAppointmentUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete coaching appointment unsupported media type response has a 5xx status code
func (o *DeleteCoachingAppointmentUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this delete coaching appointment unsupported media type response a status code equal to that given
func (o *DeleteCoachingAppointmentUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *DeleteCoachingAppointmentUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/coaching/appointments/{appointmentId}][%d] deleteCoachingAppointmentUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteCoachingAppointmentUnsupportedMediaType) String() string {
	return fmt.Sprintf("[DELETE /api/v2/coaching/appointments/{appointmentId}][%d] deleteCoachingAppointmentUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteCoachingAppointmentUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteCoachingAppointmentUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCoachingAppointmentTooManyRequests creates a DeleteCoachingAppointmentTooManyRequests with default headers values
func NewDeleteCoachingAppointmentTooManyRequests() *DeleteCoachingAppointmentTooManyRequests {
	return &DeleteCoachingAppointmentTooManyRequests{}
}

/*
DeleteCoachingAppointmentTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type DeleteCoachingAppointmentTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete coaching appointment too many requests response has a 2xx status code
func (o *DeleteCoachingAppointmentTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete coaching appointment too many requests response has a 3xx status code
func (o *DeleteCoachingAppointmentTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete coaching appointment too many requests response has a 4xx status code
func (o *DeleteCoachingAppointmentTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete coaching appointment too many requests response has a 5xx status code
func (o *DeleteCoachingAppointmentTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this delete coaching appointment too many requests response a status code equal to that given
func (o *DeleteCoachingAppointmentTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *DeleteCoachingAppointmentTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/coaching/appointments/{appointmentId}][%d] deleteCoachingAppointmentTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteCoachingAppointmentTooManyRequests) String() string {
	return fmt.Sprintf("[DELETE /api/v2/coaching/appointments/{appointmentId}][%d] deleteCoachingAppointmentTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteCoachingAppointmentTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteCoachingAppointmentTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCoachingAppointmentInternalServerError creates a DeleteCoachingAppointmentInternalServerError with default headers values
func NewDeleteCoachingAppointmentInternalServerError() *DeleteCoachingAppointmentInternalServerError {
	return &DeleteCoachingAppointmentInternalServerError{}
}

/*
DeleteCoachingAppointmentInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteCoachingAppointmentInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete coaching appointment internal server error response has a 2xx status code
func (o *DeleteCoachingAppointmentInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete coaching appointment internal server error response has a 3xx status code
func (o *DeleteCoachingAppointmentInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete coaching appointment internal server error response has a 4xx status code
func (o *DeleteCoachingAppointmentInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete coaching appointment internal server error response has a 5xx status code
func (o *DeleteCoachingAppointmentInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete coaching appointment internal server error response a status code equal to that given
func (o *DeleteCoachingAppointmentInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DeleteCoachingAppointmentInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/coaching/appointments/{appointmentId}][%d] deleteCoachingAppointmentInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteCoachingAppointmentInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /api/v2/coaching/appointments/{appointmentId}][%d] deleteCoachingAppointmentInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteCoachingAppointmentInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteCoachingAppointmentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCoachingAppointmentServiceUnavailable creates a DeleteCoachingAppointmentServiceUnavailable with default headers values
func NewDeleteCoachingAppointmentServiceUnavailable() *DeleteCoachingAppointmentServiceUnavailable {
	return &DeleteCoachingAppointmentServiceUnavailable{}
}

/*
DeleteCoachingAppointmentServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteCoachingAppointmentServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete coaching appointment service unavailable response has a 2xx status code
func (o *DeleteCoachingAppointmentServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete coaching appointment service unavailable response has a 3xx status code
func (o *DeleteCoachingAppointmentServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete coaching appointment service unavailable response has a 4xx status code
func (o *DeleteCoachingAppointmentServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete coaching appointment service unavailable response has a 5xx status code
func (o *DeleteCoachingAppointmentServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this delete coaching appointment service unavailable response a status code equal to that given
func (o *DeleteCoachingAppointmentServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *DeleteCoachingAppointmentServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/coaching/appointments/{appointmentId}][%d] deleteCoachingAppointmentServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteCoachingAppointmentServiceUnavailable) String() string {
	return fmt.Sprintf("[DELETE /api/v2/coaching/appointments/{appointmentId}][%d] deleteCoachingAppointmentServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteCoachingAppointmentServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteCoachingAppointmentServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCoachingAppointmentGatewayTimeout creates a DeleteCoachingAppointmentGatewayTimeout with default headers values
func NewDeleteCoachingAppointmentGatewayTimeout() *DeleteCoachingAppointmentGatewayTimeout {
	return &DeleteCoachingAppointmentGatewayTimeout{}
}

/*
DeleteCoachingAppointmentGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type DeleteCoachingAppointmentGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete coaching appointment gateway timeout response has a 2xx status code
func (o *DeleteCoachingAppointmentGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete coaching appointment gateway timeout response has a 3xx status code
func (o *DeleteCoachingAppointmentGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete coaching appointment gateway timeout response has a 4xx status code
func (o *DeleteCoachingAppointmentGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete coaching appointment gateway timeout response has a 5xx status code
func (o *DeleteCoachingAppointmentGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this delete coaching appointment gateway timeout response a status code equal to that given
func (o *DeleteCoachingAppointmentGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *DeleteCoachingAppointmentGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/coaching/appointments/{appointmentId}][%d] deleteCoachingAppointmentGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteCoachingAppointmentGatewayTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/coaching/appointments/{appointmentId}][%d] deleteCoachingAppointmentGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteCoachingAppointmentGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteCoachingAppointmentGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

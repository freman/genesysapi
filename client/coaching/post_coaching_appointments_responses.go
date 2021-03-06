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

// PostCoachingAppointmentsReader is a Reader for the PostCoachingAppointments structure.
type PostCoachingAppointmentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostCoachingAppointmentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostCoachingAppointmentsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewPostCoachingAppointmentsAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostCoachingAppointmentsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostCoachingAppointmentsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostCoachingAppointmentsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostCoachingAppointmentsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPostCoachingAppointmentsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostCoachingAppointmentsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostCoachingAppointmentsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostCoachingAppointmentsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostCoachingAppointmentsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostCoachingAppointmentsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostCoachingAppointmentsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostCoachingAppointmentsCreated creates a PostCoachingAppointmentsCreated with default headers values
func NewPostCoachingAppointmentsCreated() *PostCoachingAppointmentsCreated {
	return &PostCoachingAppointmentsCreated{}
}

/*PostCoachingAppointmentsCreated handles this case with default header values.

Appointment created
*/
type PostCoachingAppointmentsCreated struct {
	Payload *models.CoachingAppointmentResponse
}

func (o *PostCoachingAppointmentsCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/coaching/appointments][%d] postCoachingAppointmentsCreated  %+v", 201, o.Payload)
}

func (o *PostCoachingAppointmentsCreated) GetPayload() *models.CoachingAppointmentResponse {
	return o.Payload
}

func (o *PostCoachingAppointmentsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CoachingAppointmentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCoachingAppointmentsAccepted creates a PostCoachingAppointmentsAccepted with default headers values
func NewPostCoachingAppointmentsAccepted() *PostCoachingAppointmentsAccepted {
	return &PostCoachingAppointmentsAccepted{}
}

/*PostCoachingAppointmentsAccepted handles this case with default header values.

Appointment create request accepted
*/
type PostCoachingAppointmentsAccepted struct {
	Payload *models.CoachingAppointmentReference
}

func (o *PostCoachingAppointmentsAccepted) Error() string {
	return fmt.Sprintf("[POST /api/v2/coaching/appointments][%d] postCoachingAppointmentsAccepted  %+v", 202, o.Payload)
}

func (o *PostCoachingAppointmentsAccepted) GetPayload() *models.CoachingAppointmentReference {
	return o.Payload
}

func (o *PostCoachingAppointmentsAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CoachingAppointmentReference)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCoachingAppointmentsBadRequest creates a PostCoachingAppointmentsBadRequest with default headers values
func NewPostCoachingAppointmentsBadRequest() *PostCoachingAppointmentsBadRequest {
	return &PostCoachingAppointmentsBadRequest{}
}

/*PostCoachingAppointmentsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostCoachingAppointmentsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostCoachingAppointmentsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/coaching/appointments][%d] postCoachingAppointmentsBadRequest  %+v", 400, o.Payload)
}

func (o *PostCoachingAppointmentsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostCoachingAppointmentsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCoachingAppointmentsUnauthorized creates a PostCoachingAppointmentsUnauthorized with default headers values
func NewPostCoachingAppointmentsUnauthorized() *PostCoachingAppointmentsUnauthorized {
	return &PostCoachingAppointmentsUnauthorized{}
}

/*PostCoachingAppointmentsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostCoachingAppointmentsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostCoachingAppointmentsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/coaching/appointments][%d] postCoachingAppointmentsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostCoachingAppointmentsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostCoachingAppointmentsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCoachingAppointmentsForbidden creates a PostCoachingAppointmentsForbidden with default headers values
func NewPostCoachingAppointmentsForbidden() *PostCoachingAppointmentsForbidden {
	return &PostCoachingAppointmentsForbidden{}
}

/*PostCoachingAppointmentsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostCoachingAppointmentsForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostCoachingAppointmentsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/coaching/appointments][%d] postCoachingAppointmentsForbidden  %+v", 403, o.Payload)
}

func (o *PostCoachingAppointmentsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostCoachingAppointmentsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCoachingAppointmentsNotFound creates a PostCoachingAppointmentsNotFound with default headers values
func NewPostCoachingAppointmentsNotFound() *PostCoachingAppointmentsNotFound {
	return &PostCoachingAppointmentsNotFound{}
}

/*PostCoachingAppointmentsNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostCoachingAppointmentsNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostCoachingAppointmentsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/coaching/appointments][%d] postCoachingAppointmentsNotFound  %+v", 404, o.Payload)
}

func (o *PostCoachingAppointmentsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostCoachingAppointmentsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCoachingAppointmentsConflict creates a PostCoachingAppointmentsConflict with default headers values
func NewPostCoachingAppointmentsConflict() *PostCoachingAppointmentsConflict {
	return &PostCoachingAppointmentsConflict{}
}

/*PostCoachingAppointmentsConflict handles this case with default header values.

Conflict
*/
type PostCoachingAppointmentsConflict struct {
	Payload *models.ErrorBody
}

func (o *PostCoachingAppointmentsConflict) Error() string {
	return fmt.Sprintf("[POST /api/v2/coaching/appointments][%d] postCoachingAppointmentsConflict  %+v", 409, o.Payload)
}

func (o *PostCoachingAppointmentsConflict) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostCoachingAppointmentsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCoachingAppointmentsRequestEntityTooLarge creates a PostCoachingAppointmentsRequestEntityTooLarge with default headers values
func NewPostCoachingAppointmentsRequestEntityTooLarge() *PostCoachingAppointmentsRequestEntityTooLarge {
	return &PostCoachingAppointmentsRequestEntityTooLarge{}
}

/*PostCoachingAppointmentsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostCoachingAppointmentsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostCoachingAppointmentsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/coaching/appointments][%d] postCoachingAppointmentsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostCoachingAppointmentsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostCoachingAppointmentsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCoachingAppointmentsUnsupportedMediaType creates a PostCoachingAppointmentsUnsupportedMediaType with default headers values
func NewPostCoachingAppointmentsUnsupportedMediaType() *PostCoachingAppointmentsUnsupportedMediaType {
	return &PostCoachingAppointmentsUnsupportedMediaType{}
}

/*PostCoachingAppointmentsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostCoachingAppointmentsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostCoachingAppointmentsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/coaching/appointments][%d] postCoachingAppointmentsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostCoachingAppointmentsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostCoachingAppointmentsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCoachingAppointmentsTooManyRequests creates a PostCoachingAppointmentsTooManyRequests with default headers values
func NewPostCoachingAppointmentsTooManyRequests() *PostCoachingAppointmentsTooManyRequests {
	return &PostCoachingAppointmentsTooManyRequests{}
}

/*PostCoachingAppointmentsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PostCoachingAppointmentsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostCoachingAppointmentsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/coaching/appointments][%d] postCoachingAppointmentsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostCoachingAppointmentsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostCoachingAppointmentsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCoachingAppointmentsInternalServerError creates a PostCoachingAppointmentsInternalServerError with default headers values
func NewPostCoachingAppointmentsInternalServerError() *PostCoachingAppointmentsInternalServerError {
	return &PostCoachingAppointmentsInternalServerError{}
}

/*PostCoachingAppointmentsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostCoachingAppointmentsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostCoachingAppointmentsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/coaching/appointments][%d] postCoachingAppointmentsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostCoachingAppointmentsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostCoachingAppointmentsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCoachingAppointmentsServiceUnavailable creates a PostCoachingAppointmentsServiceUnavailable with default headers values
func NewPostCoachingAppointmentsServiceUnavailable() *PostCoachingAppointmentsServiceUnavailable {
	return &PostCoachingAppointmentsServiceUnavailable{}
}

/*PostCoachingAppointmentsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostCoachingAppointmentsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostCoachingAppointmentsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/coaching/appointments][%d] postCoachingAppointmentsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostCoachingAppointmentsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostCoachingAppointmentsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCoachingAppointmentsGatewayTimeout creates a PostCoachingAppointmentsGatewayTimeout with default headers values
func NewPostCoachingAppointmentsGatewayTimeout() *PostCoachingAppointmentsGatewayTimeout {
	return &PostCoachingAppointmentsGatewayTimeout{}
}

/*PostCoachingAppointmentsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostCoachingAppointmentsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostCoachingAppointmentsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/coaching/appointments][%d] postCoachingAppointmentsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostCoachingAppointmentsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostCoachingAppointmentsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

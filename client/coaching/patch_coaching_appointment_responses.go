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

// PatchCoachingAppointmentReader is a Reader for the PatchCoachingAppointment structure.
type PatchCoachingAppointmentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchCoachingAppointmentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchCoachingAppointmentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchCoachingAppointmentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchCoachingAppointmentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchCoachingAppointmentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchCoachingAppointmentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPatchCoachingAppointmentConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPatchCoachingAppointmentRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPatchCoachingAppointmentUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPatchCoachingAppointmentTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchCoachingAppointmentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPatchCoachingAppointmentServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPatchCoachingAppointmentGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchCoachingAppointmentOK creates a PatchCoachingAppointmentOK with default headers values
func NewPatchCoachingAppointmentOK() *PatchCoachingAppointmentOK {
	return &PatchCoachingAppointmentOK{}
}

/*PatchCoachingAppointmentOK handles this case with default header values.

Appointment updated
*/
type PatchCoachingAppointmentOK struct {
	Payload *models.CoachingAppointmentResponse
}

func (o *PatchCoachingAppointmentOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/coaching/appointments/{appointmentId}][%d] patchCoachingAppointmentOK  %+v", 200, o.Payload)
}

func (o *PatchCoachingAppointmentOK) GetPayload() *models.CoachingAppointmentResponse {
	return o.Payload
}

func (o *PatchCoachingAppointmentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CoachingAppointmentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchCoachingAppointmentBadRequest creates a PatchCoachingAppointmentBadRequest with default headers values
func NewPatchCoachingAppointmentBadRequest() *PatchCoachingAppointmentBadRequest {
	return &PatchCoachingAppointmentBadRequest{}
}

/*PatchCoachingAppointmentBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PatchCoachingAppointmentBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PatchCoachingAppointmentBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/coaching/appointments/{appointmentId}][%d] patchCoachingAppointmentBadRequest  %+v", 400, o.Payload)
}

func (o *PatchCoachingAppointmentBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchCoachingAppointmentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchCoachingAppointmentUnauthorized creates a PatchCoachingAppointmentUnauthorized with default headers values
func NewPatchCoachingAppointmentUnauthorized() *PatchCoachingAppointmentUnauthorized {
	return &PatchCoachingAppointmentUnauthorized{}
}

/*PatchCoachingAppointmentUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PatchCoachingAppointmentUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PatchCoachingAppointmentUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/coaching/appointments/{appointmentId}][%d] patchCoachingAppointmentUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchCoachingAppointmentUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchCoachingAppointmentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchCoachingAppointmentForbidden creates a PatchCoachingAppointmentForbidden with default headers values
func NewPatchCoachingAppointmentForbidden() *PatchCoachingAppointmentForbidden {
	return &PatchCoachingAppointmentForbidden{}
}

/*PatchCoachingAppointmentForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PatchCoachingAppointmentForbidden struct {
	Payload *models.ErrorBody
}

func (o *PatchCoachingAppointmentForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/coaching/appointments/{appointmentId}][%d] patchCoachingAppointmentForbidden  %+v", 403, o.Payload)
}

func (o *PatchCoachingAppointmentForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchCoachingAppointmentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchCoachingAppointmentNotFound creates a PatchCoachingAppointmentNotFound with default headers values
func NewPatchCoachingAppointmentNotFound() *PatchCoachingAppointmentNotFound {
	return &PatchCoachingAppointmentNotFound{}
}

/*PatchCoachingAppointmentNotFound handles this case with default header values.

The requested resource was not found.
*/
type PatchCoachingAppointmentNotFound struct {
	Payload *models.ErrorBody
}

func (o *PatchCoachingAppointmentNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/coaching/appointments/{appointmentId}][%d] patchCoachingAppointmentNotFound  %+v", 404, o.Payload)
}

func (o *PatchCoachingAppointmentNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchCoachingAppointmentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchCoachingAppointmentConflict creates a PatchCoachingAppointmentConflict with default headers values
func NewPatchCoachingAppointmentConflict() *PatchCoachingAppointmentConflict {
	return &PatchCoachingAppointmentConflict{}
}

/*PatchCoachingAppointmentConflict handles this case with default header values.

Conflict
*/
type PatchCoachingAppointmentConflict struct {
	Payload *models.ErrorBody
}

func (o *PatchCoachingAppointmentConflict) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/coaching/appointments/{appointmentId}][%d] patchCoachingAppointmentConflict  %+v", 409, o.Payload)
}

func (o *PatchCoachingAppointmentConflict) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchCoachingAppointmentConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchCoachingAppointmentRequestEntityTooLarge creates a PatchCoachingAppointmentRequestEntityTooLarge with default headers values
func NewPatchCoachingAppointmentRequestEntityTooLarge() *PatchCoachingAppointmentRequestEntityTooLarge {
	return &PatchCoachingAppointmentRequestEntityTooLarge{}
}

/*PatchCoachingAppointmentRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PatchCoachingAppointmentRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PatchCoachingAppointmentRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/coaching/appointments/{appointmentId}][%d] patchCoachingAppointmentRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchCoachingAppointmentRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchCoachingAppointmentRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchCoachingAppointmentUnsupportedMediaType creates a PatchCoachingAppointmentUnsupportedMediaType with default headers values
func NewPatchCoachingAppointmentUnsupportedMediaType() *PatchCoachingAppointmentUnsupportedMediaType {
	return &PatchCoachingAppointmentUnsupportedMediaType{}
}

/*PatchCoachingAppointmentUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PatchCoachingAppointmentUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PatchCoachingAppointmentUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/coaching/appointments/{appointmentId}][%d] patchCoachingAppointmentUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchCoachingAppointmentUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchCoachingAppointmentUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchCoachingAppointmentTooManyRequests creates a PatchCoachingAppointmentTooManyRequests with default headers values
func NewPatchCoachingAppointmentTooManyRequests() *PatchCoachingAppointmentTooManyRequests {
	return &PatchCoachingAppointmentTooManyRequests{}
}

/*PatchCoachingAppointmentTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PatchCoachingAppointmentTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PatchCoachingAppointmentTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/coaching/appointments/{appointmentId}][%d] patchCoachingAppointmentTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchCoachingAppointmentTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchCoachingAppointmentTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchCoachingAppointmentInternalServerError creates a PatchCoachingAppointmentInternalServerError with default headers values
func NewPatchCoachingAppointmentInternalServerError() *PatchCoachingAppointmentInternalServerError {
	return &PatchCoachingAppointmentInternalServerError{}
}

/*PatchCoachingAppointmentInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PatchCoachingAppointmentInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PatchCoachingAppointmentInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/coaching/appointments/{appointmentId}][%d] patchCoachingAppointmentInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchCoachingAppointmentInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchCoachingAppointmentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchCoachingAppointmentServiceUnavailable creates a PatchCoachingAppointmentServiceUnavailable with default headers values
func NewPatchCoachingAppointmentServiceUnavailable() *PatchCoachingAppointmentServiceUnavailable {
	return &PatchCoachingAppointmentServiceUnavailable{}
}

/*PatchCoachingAppointmentServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PatchCoachingAppointmentServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PatchCoachingAppointmentServiceUnavailable) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/coaching/appointments/{appointmentId}][%d] patchCoachingAppointmentServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchCoachingAppointmentServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchCoachingAppointmentServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchCoachingAppointmentGatewayTimeout creates a PatchCoachingAppointmentGatewayTimeout with default headers values
func NewPatchCoachingAppointmentGatewayTimeout() *PatchCoachingAppointmentGatewayTimeout {
	return &PatchCoachingAppointmentGatewayTimeout{}
}

/*PatchCoachingAppointmentGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PatchCoachingAppointmentGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PatchCoachingAppointmentGatewayTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/coaching/appointments/{appointmentId}][%d] patchCoachingAppointmentGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchCoachingAppointmentGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchCoachingAppointmentGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

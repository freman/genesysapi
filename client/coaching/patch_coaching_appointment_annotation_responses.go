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

// PatchCoachingAppointmentAnnotationReader is a Reader for the PatchCoachingAppointmentAnnotation structure.
type PatchCoachingAppointmentAnnotationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchCoachingAppointmentAnnotationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchCoachingAppointmentAnnotationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchCoachingAppointmentAnnotationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchCoachingAppointmentAnnotationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchCoachingAppointmentAnnotationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchCoachingAppointmentAnnotationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPatchCoachingAppointmentAnnotationRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPatchCoachingAppointmentAnnotationUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPatchCoachingAppointmentAnnotationTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchCoachingAppointmentAnnotationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPatchCoachingAppointmentAnnotationServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPatchCoachingAppointmentAnnotationGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchCoachingAppointmentAnnotationOK creates a PatchCoachingAppointmentAnnotationOK with default headers values
func NewPatchCoachingAppointmentAnnotationOK() *PatchCoachingAppointmentAnnotationOK {
	return &PatchCoachingAppointmentAnnotationOK{}
}

/*PatchCoachingAppointmentAnnotationOK handles this case with default header values.

successful operation
*/
type PatchCoachingAppointmentAnnotationOK struct {
	Payload *models.CoachingAnnotation
}

func (o *PatchCoachingAppointmentAnnotationOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/coaching/appointments/{appointmentId}/annotations/{annotationId}][%d] patchCoachingAppointmentAnnotationOK  %+v", 200, o.Payload)
}

func (o *PatchCoachingAppointmentAnnotationOK) GetPayload() *models.CoachingAnnotation {
	return o.Payload
}

func (o *PatchCoachingAppointmentAnnotationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CoachingAnnotation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchCoachingAppointmentAnnotationBadRequest creates a PatchCoachingAppointmentAnnotationBadRequest with default headers values
func NewPatchCoachingAppointmentAnnotationBadRequest() *PatchCoachingAppointmentAnnotationBadRequest {
	return &PatchCoachingAppointmentAnnotationBadRequest{}
}

/*PatchCoachingAppointmentAnnotationBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PatchCoachingAppointmentAnnotationBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PatchCoachingAppointmentAnnotationBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/coaching/appointments/{appointmentId}/annotations/{annotationId}][%d] patchCoachingAppointmentAnnotationBadRequest  %+v", 400, o.Payload)
}

func (o *PatchCoachingAppointmentAnnotationBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchCoachingAppointmentAnnotationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchCoachingAppointmentAnnotationUnauthorized creates a PatchCoachingAppointmentAnnotationUnauthorized with default headers values
func NewPatchCoachingAppointmentAnnotationUnauthorized() *PatchCoachingAppointmentAnnotationUnauthorized {
	return &PatchCoachingAppointmentAnnotationUnauthorized{}
}

/*PatchCoachingAppointmentAnnotationUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PatchCoachingAppointmentAnnotationUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PatchCoachingAppointmentAnnotationUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/coaching/appointments/{appointmentId}/annotations/{annotationId}][%d] patchCoachingAppointmentAnnotationUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchCoachingAppointmentAnnotationUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchCoachingAppointmentAnnotationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchCoachingAppointmentAnnotationForbidden creates a PatchCoachingAppointmentAnnotationForbidden with default headers values
func NewPatchCoachingAppointmentAnnotationForbidden() *PatchCoachingAppointmentAnnotationForbidden {
	return &PatchCoachingAppointmentAnnotationForbidden{}
}

/*PatchCoachingAppointmentAnnotationForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PatchCoachingAppointmentAnnotationForbidden struct {
	Payload *models.ErrorBody
}

func (o *PatchCoachingAppointmentAnnotationForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/coaching/appointments/{appointmentId}/annotations/{annotationId}][%d] patchCoachingAppointmentAnnotationForbidden  %+v", 403, o.Payload)
}

func (o *PatchCoachingAppointmentAnnotationForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchCoachingAppointmentAnnotationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchCoachingAppointmentAnnotationNotFound creates a PatchCoachingAppointmentAnnotationNotFound with default headers values
func NewPatchCoachingAppointmentAnnotationNotFound() *PatchCoachingAppointmentAnnotationNotFound {
	return &PatchCoachingAppointmentAnnotationNotFound{}
}

/*PatchCoachingAppointmentAnnotationNotFound handles this case with default header values.

The requested resource was not found.
*/
type PatchCoachingAppointmentAnnotationNotFound struct {
	Payload *models.ErrorBody
}

func (o *PatchCoachingAppointmentAnnotationNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/coaching/appointments/{appointmentId}/annotations/{annotationId}][%d] patchCoachingAppointmentAnnotationNotFound  %+v", 404, o.Payload)
}

func (o *PatchCoachingAppointmentAnnotationNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchCoachingAppointmentAnnotationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchCoachingAppointmentAnnotationRequestEntityTooLarge creates a PatchCoachingAppointmentAnnotationRequestEntityTooLarge with default headers values
func NewPatchCoachingAppointmentAnnotationRequestEntityTooLarge() *PatchCoachingAppointmentAnnotationRequestEntityTooLarge {
	return &PatchCoachingAppointmentAnnotationRequestEntityTooLarge{}
}

/*PatchCoachingAppointmentAnnotationRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PatchCoachingAppointmentAnnotationRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PatchCoachingAppointmentAnnotationRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/coaching/appointments/{appointmentId}/annotations/{annotationId}][%d] patchCoachingAppointmentAnnotationRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchCoachingAppointmentAnnotationRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchCoachingAppointmentAnnotationRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchCoachingAppointmentAnnotationUnsupportedMediaType creates a PatchCoachingAppointmentAnnotationUnsupportedMediaType with default headers values
func NewPatchCoachingAppointmentAnnotationUnsupportedMediaType() *PatchCoachingAppointmentAnnotationUnsupportedMediaType {
	return &PatchCoachingAppointmentAnnotationUnsupportedMediaType{}
}

/*PatchCoachingAppointmentAnnotationUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PatchCoachingAppointmentAnnotationUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PatchCoachingAppointmentAnnotationUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/coaching/appointments/{appointmentId}/annotations/{annotationId}][%d] patchCoachingAppointmentAnnotationUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchCoachingAppointmentAnnotationUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchCoachingAppointmentAnnotationUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchCoachingAppointmentAnnotationTooManyRequests creates a PatchCoachingAppointmentAnnotationTooManyRequests with default headers values
func NewPatchCoachingAppointmentAnnotationTooManyRequests() *PatchCoachingAppointmentAnnotationTooManyRequests {
	return &PatchCoachingAppointmentAnnotationTooManyRequests{}
}

/*PatchCoachingAppointmentAnnotationTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PatchCoachingAppointmentAnnotationTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PatchCoachingAppointmentAnnotationTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/coaching/appointments/{appointmentId}/annotations/{annotationId}][%d] patchCoachingAppointmentAnnotationTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchCoachingAppointmentAnnotationTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchCoachingAppointmentAnnotationTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchCoachingAppointmentAnnotationInternalServerError creates a PatchCoachingAppointmentAnnotationInternalServerError with default headers values
func NewPatchCoachingAppointmentAnnotationInternalServerError() *PatchCoachingAppointmentAnnotationInternalServerError {
	return &PatchCoachingAppointmentAnnotationInternalServerError{}
}

/*PatchCoachingAppointmentAnnotationInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PatchCoachingAppointmentAnnotationInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PatchCoachingAppointmentAnnotationInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/coaching/appointments/{appointmentId}/annotations/{annotationId}][%d] patchCoachingAppointmentAnnotationInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchCoachingAppointmentAnnotationInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchCoachingAppointmentAnnotationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchCoachingAppointmentAnnotationServiceUnavailable creates a PatchCoachingAppointmentAnnotationServiceUnavailable with default headers values
func NewPatchCoachingAppointmentAnnotationServiceUnavailable() *PatchCoachingAppointmentAnnotationServiceUnavailable {
	return &PatchCoachingAppointmentAnnotationServiceUnavailable{}
}

/*PatchCoachingAppointmentAnnotationServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PatchCoachingAppointmentAnnotationServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PatchCoachingAppointmentAnnotationServiceUnavailable) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/coaching/appointments/{appointmentId}/annotations/{annotationId}][%d] patchCoachingAppointmentAnnotationServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchCoachingAppointmentAnnotationServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchCoachingAppointmentAnnotationServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchCoachingAppointmentAnnotationGatewayTimeout creates a PatchCoachingAppointmentAnnotationGatewayTimeout with default headers values
func NewPatchCoachingAppointmentAnnotationGatewayTimeout() *PatchCoachingAppointmentAnnotationGatewayTimeout {
	return &PatchCoachingAppointmentAnnotationGatewayTimeout{}
}

/*PatchCoachingAppointmentAnnotationGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PatchCoachingAppointmentAnnotationGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PatchCoachingAppointmentAnnotationGatewayTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/coaching/appointments/{appointmentId}/annotations/{annotationId}][%d] patchCoachingAppointmentAnnotationGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchCoachingAppointmentAnnotationGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchCoachingAppointmentAnnotationGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

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

// GetCoachingAppointmentAnnotationReader is a Reader for the GetCoachingAppointmentAnnotation structure.
type GetCoachingAppointmentAnnotationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCoachingAppointmentAnnotationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCoachingAppointmentAnnotationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetCoachingAppointmentAnnotationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetCoachingAppointmentAnnotationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCoachingAppointmentAnnotationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetCoachingAppointmentAnnotationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetCoachingAppointmentAnnotationRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetCoachingAppointmentAnnotationUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetCoachingAppointmentAnnotationTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCoachingAppointmentAnnotationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetCoachingAppointmentAnnotationServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetCoachingAppointmentAnnotationGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCoachingAppointmentAnnotationOK creates a GetCoachingAppointmentAnnotationOK with default headers values
func NewGetCoachingAppointmentAnnotationOK() *GetCoachingAppointmentAnnotationOK {
	return &GetCoachingAppointmentAnnotationOK{}
}

/*GetCoachingAppointmentAnnotationOK handles this case with default header values.

Annotation retrieved
*/
type GetCoachingAppointmentAnnotationOK struct {
	Payload *models.CoachingAnnotation
}

func (o *GetCoachingAppointmentAnnotationOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/coaching/appointments/{appointmentId}/annotations/{annotationId}][%d] getCoachingAppointmentAnnotationOK  %+v", 200, o.Payload)
}

func (o *GetCoachingAppointmentAnnotationOK) GetPayload() *models.CoachingAnnotation {
	return o.Payload
}

func (o *GetCoachingAppointmentAnnotationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CoachingAnnotation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCoachingAppointmentAnnotationBadRequest creates a GetCoachingAppointmentAnnotationBadRequest with default headers values
func NewGetCoachingAppointmentAnnotationBadRequest() *GetCoachingAppointmentAnnotationBadRequest {
	return &GetCoachingAppointmentAnnotationBadRequest{}
}

/*GetCoachingAppointmentAnnotationBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetCoachingAppointmentAnnotationBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetCoachingAppointmentAnnotationBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/coaching/appointments/{appointmentId}/annotations/{annotationId}][%d] getCoachingAppointmentAnnotationBadRequest  %+v", 400, o.Payload)
}

func (o *GetCoachingAppointmentAnnotationBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetCoachingAppointmentAnnotationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCoachingAppointmentAnnotationUnauthorized creates a GetCoachingAppointmentAnnotationUnauthorized with default headers values
func NewGetCoachingAppointmentAnnotationUnauthorized() *GetCoachingAppointmentAnnotationUnauthorized {
	return &GetCoachingAppointmentAnnotationUnauthorized{}
}

/*GetCoachingAppointmentAnnotationUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetCoachingAppointmentAnnotationUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetCoachingAppointmentAnnotationUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/coaching/appointments/{appointmentId}/annotations/{annotationId}][%d] getCoachingAppointmentAnnotationUnauthorized  %+v", 401, o.Payload)
}

func (o *GetCoachingAppointmentAnnotationUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetCoachingAppointmentAnnotationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCoachingAppointmentAnnotationForbidden creates a GetCoachingAppointmentAnnotationForbidden with default headers values
func NewGetCoachingAppointmentAnnotationForbidden() *GetCoachingAppointmentAnnotationForbidden {
	return &GetCoachingAppointmentAnnotationForbidden{}
}

/*GetCoachingAppointmentAnnotationForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetCoachingAppointmentAnnotationForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetCoachingAppointmentAnnotationForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/coaching/appointments/{appointmentId}/annotations/{annotationId}][%d] getCoachingAppointmentAnnotationForbidden  %+v", 403, o.Payload)
}

func (o *GetCoachingAppointmentAnnotationForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetCoachingAppointmentAnnotationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCoachingAppointmentAnnotationNotFound creates a GetCoachingAppointmentAnnotationNotFound with default headers values
func NewGetCoachingAppointmentAnnotationNotFound() *GetCoachingAppointmentAnnotationNotFound {
	return &GetCoachingAppointmentAnnotationNotFound{}
}

/*GetCoachingAppointmentAnnotationNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetCoachingAppointmentAnnotationNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetCoachingAppointmentAnnotationNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/coaching/appointments/{appointmentId}/annotations/{annotationId}][%d] getCoachingAppointmentAnnotationNotFound  %+v", 404, o.Payload)
}

func (o *GetCoachingAppointmentAnnotationNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetCoachingAppointmentAnnotationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCoachingAppointmentAnnotationRequestEntityTooLarge creates a GetCoachingAppointmentAnnotationRequestEntityTooLarge with default headers values
func NewGetCoachingAppointmentAnnotationRequestEntityTooLarge() *GetCoachingAppointmentAnnotationRequestEntityTooLarge {
	return &GetCoachingAppointmentAnnotationRequestEntityTooLarge{}
}

/*GetCoachingAppointmentAnnotationRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetCoachingAppointmentAnnotationRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetCoachingAppointmentAnnotationRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/coaching/appointments/{appointmentId}/annotations/{annotationId}][%d] getCoachingAppointmentAnnotationRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetCoachingAppointmentAnnotationRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetCoachingAppointmentAnnotationRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCoachingAppointmentAnnotationUnsupportedMediaType creates a GetCoachingAppointmentAnnotationUnsupportedMediaType with default headers values
func NewGetCoachingAppointmentAnnotationUnsupportedMediaType() *GetCoachingAppointmentAnnotationUnsupportedMediaType {
	return &GetCoachingAppointmentAnnotationUnsupportedMediaType{}
}

/*GetCoachingAppointmentAnnotationUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetCoachingAppointmentAnnotationUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetCoachingAppointmentAnnotationUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/coaching/appointments/{appointmentId}/annotations/{annotationId}][%d] getCoachingAppointmentAnnotationUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetCoachingAppointmentAnnotationUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetCoachingAppointmentAnnotationUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCoachingAppointmentAnnotationTooManyRequests creates a GetCoachingAppointmentAnnotationTooManyRequests with default headers values
func NewGetCoachingAppointmentAnnotationTooManyRequests() *GetCoachingAppointmentAnnotationTooManyRequests {
	return &GetCoachingAppointmentAnnotationTooManyRequests{}
}

/*GetCoachingAppointmentAnnotationTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetCoachingAppointmentAnnotationTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetCoachingAppointmentAnnotationTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/coaching/appointments/{appointmentId}/annotations/{annotationId}][%d] getCoachingAppointmentAnnotationTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetCoachingAppointmentAnnotationTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetCoachingAppointmentAnnotationTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCoachingAppointmentAnnotationInternalServerError creates a GetCoachingAppointmentAnnotationInternalServerError with default headers values
func NewGetCoachingAppointmentAnnotationInternalServerError() *GetCoachingAppointmentAnnotationInternalServerError {
	return &GetCoachingAppointmentAnnotationInternalServerError{}
}

/*GetCoachingAppointmentAnnotationInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetCoachingAppointmentAnnotationInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetCoachingAppointmentAnnotationInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/coaching/appointments/{appointmentId}/annotations/{annotationId}][%d] getCoachingAppointmentAnnotationInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCoachingAppointmentAnnotationInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetCoachingAppointmentAnnotationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCoachingAppointmentAnnotationServiceUnavailable creates a GetCoachingAppointmentAnnotationServiceUnavailable with default headers values
func NewGetCoachingAppointmentAnnotationServiceUnavailable() *GetCoachingAppointmentAnnotationServiceUnavailable {
	return &GetCoachingAppointmentAnnotationServiceUnavailable{}
}

/*GetCoachingAppointmentAnnotationServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetCoachingAppointmentAnnotationServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetCoachingAppointmentAnnotationServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/coaching/appointments/{appointmentId}/annotations/{annotationId}][%d] getCoachingAppointmentAnnotationServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetCoachingAppointmentAnnotationServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetCoachingAppointmentAnnotationServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCoachingAppointmentAnnotationGatewayTimeout creates a GetCoachingAppointmentAnnotationGatewayTimeout with default headers values
func NewGetCoachingAppointmentAnnotationGatewayTimeout() *GetCoachingAppointmentAnnotationGatewayTimeout {
	return &GetCoachingAppointmentAnnotationGatewayTimeout{}
}

/*GetCoachingAppointmentAnnotationGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetCoachingAppointmentAnnotationGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetCoachingAppointmentAnnotationGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/coaching/appointments/{appointmentId}/annotations/{annotationId}][%d] getCoachingAppointmentAnnotationGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetCoachingAppointmentAnnotationGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetCoachingAppointmentAnnotationGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

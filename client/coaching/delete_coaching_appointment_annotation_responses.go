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

// DeleteCoachingAppointmentAnnotationReader is a Reader for the DeleteCoachingAppointmentAnnotation structure.
type DeleteCoachingAppointmentAnnotationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCoachingAppointmentAnnotationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteCoachingAppointmentAnnotationNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteCoachingAppointmentAnnotationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteCoachingAppointmentAnnotationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteCoachingAppointmentAnnotationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteCoachingAppointmentAnnotationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewDeleteCoachingAppointmentAnnotationRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteCoachingAppointmentAnnotationRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteCoachingAppointmentAnnotationUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteCoachingAppointmentAnnotationTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteCoachingAppointmentAnnotationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteCoachingAppointmentAnnotationServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteCoachingAppointmentAnnotationGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteCoachingAppointmentAnnotationNoContent creates a DeleteCoachingAppointmentAnnotationNoContent with default headers values
func NewDeleteCoachingAppointmentAnnotationNoContent() *DeleteCoachingAppointmentAnnotationNoContent {
	return &DeleteCoachingAppointmentAnnotationNoContent{}
}

/*DeleteCoachingAppointmentAnnotationNoContent handles this case with default header values.

Annotation deleted
*/
type DeleteCoachingAppointmentAnnotationNoContent struct {
}

func (o *DeleteCoachingAppointmentAnnotationNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/coaching/appointments/{appointmentId}/annotations/{annotationId}][%d] deleteCoachingAppointmentAnnotationNoContent ", 204)
}

func (o *DeleteCoachingAppointmentAnnotationNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteCoachingAppointmentAnnotationBadRequest creates a DeleteCoachingAppointmentAnnotationBadRequest with default headers values
func NewDeleteCoachingAppointmentAnnotationBadRequest() *DeleteCoachingAppointmentAnnotationBadRequest {
	return &DeleteCoachingAppointmentAnnotationBadRequest{}
}

/*DeleteCoachingAppointmentAnnotationBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteCoachingAppointmentAnnotationBadRequest struct {
	Payload *models.ErrorBody
}

func (o *DeleteCoachingAppointmentAnnotationBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/coaching/appointments/{appointmentId}/annotations/{annotationId}][%d] deleteCoachingAppointmentAnnotationBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteCoachingAppointmentAnnotationBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteCoachingAppointmentAnnotationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCoachingAppointmentAnnotationUnauthorized creates a DeleteCoachingAppointmentAnnotationUnauthorized with default headers values
func NewDeleteCoachingAppointmentAnnotationUnauthorized() *DeleteCoachingAppointmentAnnotationUnauthorized {
	return &DeleteCoachingAppointmentAnnotationUnauthorized{}
}

/*DeleteCoachingAppointmentAnnotationUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteCoachingAppointmentAnnotationUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *DeleteCoachingAppointmentAnnotationUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/coaching/appointments/{appointmentId}/annotations/{annotationId}][%d] deleteCoachingAppointmentAnnotationUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteCoachingAppointmentAnnotationUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteCoachingAppointmentAnnotationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCoachingAppointmentAnnotationForbidden creates a DeleteCoachingAppointmentAnnotationForbidden with default headers values
func NewDeleteCoachingAppointmentAnnotationForbidden() *DeleteCoachingAppointmentAnnotationForbidden {
	return &DeleteCoachingAppointmentAnnotationForbidden{}
}

/*DeleteCoachingAppointmentAnnotationForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type DeleteCoachingAppointmentAnnotationForbidden struct {
	Payload *models.ErrorBody
}

func (o *DeleteCoachingAppointmentAnnotationForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/coaching/appointments/{appointmentId}/annotations/{annotationId}][%d] deleteCoachingAppointmentAnnotationForbidden  %+v", 403, o.Payload)
}

func (o *DeleteCoachingAppointmentAnnotationForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteCoachingAppointmentAnnotationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCoachingAppointmentAnnotationNotFound creates a DeleteCoachingAppointmentAnnotationNotFound with default headers values
func NewDeleteCoachingAppointmentAnnotationNotFound() *DeleteCoachingAppointmentAnnotationNotFound {
	return &DeleteCoachingAppointmentAnnotationNotFound{}
}

/*DeleteCoachingAppointmentAnnotationNotFound handles this case with default header values.

The requested resource was not found.
*/
type DeleteCoachingAppointmentAnnotationNotFound struct {
	Payload *models.ErrorBody
}

func (o *DeleteCoachingAppointmentAnnotationNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/coaching/appointments/{appointmentId}/annotations/{annotationId}][%d] deleteCoachingAppointmentAnnotationNotFound  %+v", 404, o.Payload)
}

func (o *DeleteCoachingAppointmentAnnotationNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteCoachingAppointmentAnnotationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCoachingAppointmentAnnotationRequestTimeout creates a DeleteCoachingAppointmentAnnotationRequestTimeout with default headers values
func NewDeleteCoachingAppointmentAnnotationRequestTimeout() *DeleteCoachingAppointmentAnnotationRequestTimeout {
	return &DeleteCoachingAppointmentAnnotationRequestTimeout{}
}

/*DeleteCoachingAppointmentAnnotationRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type DeleteCoachingAppointmentAnnotationRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteCoachingAppointmentAnnotationRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/coaching/appointments/{appointmentId}/annotations/{annotationId}][%d] deleteCoachingAppointmentAnnotationRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteCoachingAppointmentAnnotationRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteCoachingAppointmentAnnotationRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCoachingAppointmentAnnotationRequestEntityTooLarge creates a DeleteCoachingAppointmentAnnotationRequestEntityTooLarge with default headers values
func NewDeleteCoachingAppointmentAnnotationRequestEntityTooLarge() *DeleteCoachingAppointmentAnnotationRequestEntityTooLarge {
	return &DeleteCoachingAppointmentAnnotationRequestEntityTooLarge{}
}

/*DeleteCoachingAppointmentAnnotationRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type DeleteCoachingAppointmentAnnotationRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *DeleteCoachingAppointmentAnnotationRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/coaching/appointments/{appointmentId}/annotations/{annotationId}][%d] deleteCoachingAppointmentAnnotationRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteCoachingAppointmentAnnotationRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteCoachingAppointmentAnnotationRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCoachingAppointmentAnnotationUnsupportedMediaType creates a DeleteCoachingAppointmentAnnotationUnsupportedMediaType with default headers values
func NewDeleteCoachingAppointmentAnnotationUnsupportedMediaType() *DeleteCoachingAppointmentAnnotationUnsupportedMediaType {
	return &DeleteCoachingAppointmentAnnotationUnsupportedMediaType{}
}

/*DeleteCoachingAppointmentAnnotationUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteCoachingAppointmentAnnotationUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *DeleteCoachingAppointmentAnnotationUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/coaching/appointments/{appointmentId}/annotations/{annotationId}][%d] deleteCoachingAppointmentAnnotationUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteCoachingAppointmentAnnotationUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteCoachingAppointmentAnnotationUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCoachingAppointmentAnnotationTooManyRequests creates a DeleteCoachingAppointmentAnnotationTooManyRequests with default headers values
func NewDeleteCoachingAppointmentAnnotationTooManyRequests() *DeleteCoachingAppointmentAnnotationTooManyRequests {
	return &DeleteCoachingAppointmentAnnotationTooManyRequests{}
}

/*DeleteCoachingAppointmentAnnotationTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type DeleteCoachingAppointmentAnnotationTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *DeleteCoachingAppointmentAnnotationTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/coaching/appointments/{appointmentId}/annotations/{annotationId}][%d] deleteCoachingAppointmentAnnotationTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteCoachingAppointmentAnnotationTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteCoachingAppointmentAnnotationTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCoachingAppointmentAnnotationInternalServerError creates a DeleteCoachingAppointmentAnnotationInternalServerError with default headers values
func NewDeleteCoachingAppointmentAnnotationInternalServerError() *DeleteCoachingAppointmentAnnotationInternalServerError {
	return &DeleteCoachingAppointmentAnnotationInternalServerError{}
}

/*DeleteCoachingAppointmentAnnotationInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteCoachingAppointmentAnnotationInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *DeleteCoachingAppointmentAnnotationInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/coaching/appointments/{appointmentId}/annotations/{annotationId}][%d] deleteCoachingAppointmentAnnotationInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteCoachingAppointmentAnnotationInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteCoachingAppointmentAnnotationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCoachingAppointmentAnnotationServiceUnavailable creates a DeleteCoachingAppointmentAnnotationServiceUnavailable with default headers values
func NewDeleteCoachingAppointmentAnnotationServiceUnavailable() *DeleteCoachingAppointmentAnnotationServiceUnavailable {
	return &DeleteCoachingAppointmentAnnotationServiceUnavailable{}
}

/*DeleteCoachingAppointmentAnnotationServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteCoachingAppointmentAnnotationServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *DeleteCoachingAppointmentAnnotationServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/coaching/appointments/{appointmentId}/annotations/{annotationId}][%d] deleteCoachingAppointmentAnnotationServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteCoachingAppointmentAnnotationServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteCoachingAppointmentAnnotationServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCoachingAppointmentAnnotationGatewayTimeout creates a DeleteCoachingAppointmentAnnotationGatewayTimeout with default headers values
func NewDeleteCoachingAppointmentAnnotationGatewayTimeout() *DeleteCoachingAppointmentAnnotationGatewayTimeout {
	return &DeleteCoachingAppointmentAnnotationGatewayTimeout{}
}

/*DeleteCoachingAppointmentAnnotationGatewayTimeout handles this case with default header values.

The request timed out.
*/
type DeleteCoachingAppointmentAnnotationGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteCoachingAppointmentAnnotationGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/coaching/appointments/{appointmentId}/annotations/{annotationId}][%d] deleteCoachingAppointmentAnnotationGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteCoachingAppointmentAnnotationGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteCoachingAppointmentAnnotationGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

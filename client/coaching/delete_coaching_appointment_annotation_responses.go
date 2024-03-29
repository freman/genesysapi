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

/*
DeleteCoachingAppointmentAnnotationNoContent describes a response with status code 204, with default header values.

Annotation deleted
*/
type DeleteCoachingAppointmentAnnotationNoContent struct {
}

// IsSuccess returns true when this delete coaching appointment annotation no content response has a 2xx status code
func (o *DeleteCoachingAppointmentAnnotationNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete coaching appointment annotation no content response has a 3xx status code
func (o *DeleteCoachingAppointmentAnnotationNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete coaching appointment annotation no content response has a 4xx status code
func (o *DeleteCoachingAppointmentAnnotationNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete coaching appointment annotation no content response has a 5xx status code
func (o *DeleteCoachingAppointmentAnnotationNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete coaching appointment annotation no content response a status code equal to that given
func (o *DeleteCoachingAppointmentAnnotationNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DeleteCoachingAppointmentAnnotationNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/coaching/appointments/{appointmentId}/annotations/{annotationId}][%d] deleteCoachingAppointmentAnnotationNoContent ", 204)
}

func (o *DeleteCoachingAppointmentAnnotationNoContent) String() string {
	return fmt.Sprintf("[DELETE /api/v2/coaching/appointments/{appointmentId}/annotations/{annotationId}][%d] deleteCoachingAppointmentAnnotationNoContent ", 204)
}

func (o *DeleteCoachingAppointmentAnnotationNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteCoachingAppointmentAnnotationBadRequest creates a DeleteCoachingAppointmentAnnotationBadRequest with default headers values
func NewDeleteCoachingAppointmentAnnotationBadRequest() *DeleteCoachingAppointmentAnnotationBadRequest {
	return &DeleteCoachingAppointmentAnnotationBadRequest{}
}

/*
DeleteCoachingAppointmentAnnotationBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteCoachingAppointmentAnnotationBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete coaching appointment annotation bad request response has a 2xx status code
func (o *DeleteCoachingAppointmentAnnotationBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete coaching appointment annotation bad request response has a 3xx status code
func (o *DeleteCoachingAppointmentAnnotationBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete coaching appointment annotation bad request response has a 4xx status code
func (o *DeleteCoachingAppointmentAnnotationBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete coaching appointment annotation bad request response has a 5xx status code
func (o *DeleteCoachingAppointmentAnnotationBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete coaching appointment annotation bad request response a status code equal to that given
func (o *DeleteCoachingAppointmentAnnotationBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DeleteCoachingAppointmentAnnotationBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/coaching/appointments/{appointmentId}/annotations/{annotationId}][%d] deleteCoachingAppointmentAnnotationBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteCoachingAppointmentAnnotationBadRequest) String() string {
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

/*
DeleteCoachingAppointmentAnnotationUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteCoachingAppointmentAnnotationUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete coaching appointment annotation unauthorized response has a 2xx status code
func (o *DeleteCoachingAppointmentAnnotationUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete coaching appointment annotation unauthorized response has a 3xx status code
func (o *DeleteCoachingAppointmentAnnotationUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete coaching appointment annotation unauthorized response has a 4xx status code
func (o *DeleteCoachingAppointmentAnnotationUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete coaching appointment annotation unauthorized response has a 5xx status code
func (o *DeleteCoachingAppointmentAnnotationUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete coaching appointment annotation unauthorized response a status code equal to that given
func (o *DeleteCoachingAppointmentAnnotationUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteCoachingAppointmentAnnotationUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/coaching/appointments/{appointmentId}/annotations/{annotationId}][%d] deleteCoachingAppointmentAnnotationUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteCoachingAppointmentAnnotationUnauthorized) String() string {
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

/*
DeleteCoachingAppointmentAnnotationForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type DeleteCoachingAppointmentAnnotationForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete coaching appointment annotation forbidden response has a 2xx status code
func (o *DeleteCoachingAppointmentAnnotationForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete coaching appointment annotation forbidden response has a 3xx status code
func (o *DeleteCoachingAppointmentAnnotationForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete coaching appointment annotation forbidden response has a 4xx status code
func (o *DeleteCoachingAppointmentAnnotationForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete coaching appointment annotation forbidden response has a 5xx status code
func (o *DeleteCoachingAppointmentAnnotationForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete coaching appointment annotation forbidden response a status code equal to that given
func (o *DeleteCoachingAppointmentAnnotationForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteCoachingAppointmentAnnotationForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/coaching/appointments/{appointmentId}/annotations/{annotationId}][%d] deleteCoachingAppointmentAnnotationForbidden  %+v", 403, o.Payload)
}

func (o *DeleteCoachingAppointmentAnnotationForbidden) String() string {
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

/*
DeleteCoachingAppointmentAnnotationNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type DeleteCoachingAppointmentAnnotationNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete coaching appointment annotation not found response has a 2xx status code
func (o *DeleteCoachingAppointmentAnnotationNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete coaching appointment annotation not found response has a 3xx status code
func (o *DeleteCoachingAppointmentAnnotationNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete coaching appointment annotation not found response has a 4xx status code
func (o *DeleteCoachingAppointmentAnnotationNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete coaching appointment annotation not found response has a 5xx status code
func (o *DeleteCoachingAppointmentAnnotationNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete coaching appointment annotation not found response a status code equal to that given
func (o *DeleteCoachingAppointmentAnnotationNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteCoachingAppointmentAnnotationNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/coaching/appointments/{appointmentId}/annotations/{annotationId}][%d] deleteCoachingAppointmentAnnotationNotFound  %+v", 404, o.Payload)
}

func (o *DeleteCoachingAppointmentAnnotationNotFound) String() string {
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

/*
DeleteCoachingAppointmentAnnotationRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type DeleteCoachingAppointmentAnnotationRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete coaching appointment annotation request timeout response has a 2xx status code
func (o *DeleteCoachingAppointmentAnnotationRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete coaching appointment annotation request timeout response has a 3xx status code
func (o *DeleteCoachingAppointmentAnnotationRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete coaching appointment annotation request timeout response has a 4xx status code
func (o *DeleteCoachingAppointmentAnnotationRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete coaching appointment annotation request timeout response has a 5xx status code
func (o *DeleteCoachingAppointmentAnnotationRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this delete coaching appointment annotation request timeout response a status code equal to that given
func (o *DeleteCoachingAppointmentAnnotationRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *DeleteCoachingAppointmentAnnotationRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/coaching/appointments/{appointmentId}/annotations/{annotationId}][%d] deleteCoachingAppointmentAnnotationRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteCoachingAppointmentAnnotationRequestTimeout) String() string {
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

/*
DeleteCoachingAppointmentAnnotationRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type DeleteCoachingAppointmentAnnotationRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete coaching appointment annotation request entity too large response has a 2xx status code
func (o *DeleteCoachingAppointmentAnnotationRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete coaching appointment annotation request entity too large response has a 3xx status code
func (o *DeleteCoachingAppointmentAnnotationRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete coaching appointment annotation request entity too large response has a 4xx status code
func (o *DeleteCoachingAppointmentAnnotationRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete coaching appointment annotation request entity too large response has a 5xx status code
func (o *DeleteCoachingAppointmentAnnotationRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this delete coaching appointment annotation request entity too large response a status code equal to that given
func (o *DeleteCoachingAppointmentAnnotationRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *DeleteCoachingAppointmentAnnotationRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/coaching/appointments/{appointmentId}/annotations/{annotationId}][%d] deleteCoachingAppointmentAnnotationRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteCoachingAppointmentAnnotationRequestEntityTooLarge) String() string {
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

/*
DeleteCoachingAppointmentAnnotationUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteCoachingAppointmentAnnotationUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete coaching appointment annotation unsupported media type response has a 2xx status code
func (o *DeleteCoachingAppointmentAnnotationUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete coaching appointment annotation unsupported media type response has a 3xx status code
func (o *DeleteCoachingAppointmentAnnotationUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete coaching appointment annotation unsupported media type response has a 4xx status code
func (o *DeleteCoachingAppointmentAnnotationUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete coaching appointment annotation unsupported media type response has a 5xx status code
func (o *DeleteCoachingAppointmentAnnotationUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this delete coaching appointment annotation unsupported media type response a status code equal to that given
func (o *DeleteCoachingAppointmentAnnotationUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *DeleteCoachingAppointmentAnnotationUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/coaching/appointments/{appointmentId}/annotations/{annotationId}][%d] deleteCoachingAppointmentAnnotationUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteCoachingAppointmentAnnotationUnsupportedMediaType) String() string {
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

/*
DeleteCoachingAppointmentAnnotationTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type DeleteCoachingAppointmentAnnotationTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete coaching appointment annotation too many requests response has a 2xx status code
func (o *DeleteCoachingAppointmentAnnotationTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete coaching appointment annotation too many requests response has a 3xx status code
func (o *DeleteCoachingAppointmentAnnotationTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete coaching appointment annotation too many requests response has a 4xx status code
func (o *DeleteCoachingAppointmentAnnotationTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete coaching appointment annotation too many requests response has a 5xx status code
func (o *DeleteCoachingAppointmentAnnotationTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this delete coaching appointment annotation too many requests response a status code equal to that given
func (o *DeleteCoachingAppointmentAnnotationTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *DeleteCoachingAppointmentAnnotationTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/coaching/appointments/{appointmentId}/annotations/{annotationId}][%d] deleteCoachingAppointmentAnnotationTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteCoachingAppointmentAnnotationTooManyRequests) String() string {
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

/*
DeleteCoachingAppointmentAnnotationInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteCoachingAppointmentAnnotationInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete coaching appointment annotation internal server error response has a 2xx status code
func (o *DeleteCoachingAppointmentAnnotationInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete coaching appointment annotation internal server error response has a 3xx status code
func (o *DeleteCoachingAppointmentAnnotationInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete coaching appointment annotation internal server error response has a 4xx status code
func (o *DeleteCoachingAppointmentAnnotationInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete coaching appointment annotation internal server error response has a 5xx status code
func (o *DeleteCoachingAppointmentAnnotationInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete coaching appointment annotation internal server error response a status code equal to that given
func (o *DeleteCoachingAppointmentAnnotationInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DeleteCoachingAppointmentAnnotationInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/coaching/appointments/{appointmentId}/annotations/{annotationId}][%d] deleteCoachingAppointmentAnnotationInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteCoachingAppointmentAnnotationInternalServerError) String() string {
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

/*
DeleteCoachingAppointmentAnnotationServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteCoachingAppointmentAnnotationServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete coaching appointment annotation service unavailable response has a 2xx status code
func (o *DeleteCoachingAppointmentAnnotationServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete coaching appointment annotation service unavailable response has a 3xx status code
func (o *DeleteCoachingAppointmentAnnotationServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete coaching appointment annotation service unavailable response has a 4xx status code
func (o *DeleteCoachingAppointmentAnnotationServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete coaching appointment annotation service unavailable response has a 5xx status code
func (o *DeleteCoachingAppointmentAnnotationServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this delete coaching appointment annotation service unavailable response a status code equal to that given
func (o *DeleteCoachingAppointmentAnnotationServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *DeleteCoachingAppointmentAnnotationServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/coaching/appointments/{appointmentId}/annotations/{annotationId}][%d] deleteCoachingAppointmentAnnotationServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteCoachingAppointmentAnnotationServiceUnavailable) String() string {
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

/*
DeleteCoachingAppointmentAnnotationGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type DeleteCoachingAppointmentAnnotationGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete coaching appointment annotation gateway timeout response has a 2xx status code
func (o *DeleteCoachingAppointmentAnnotationGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete coaching appointment annotation gateway timeout response has a 3xx status code
func (o *DeleteCoachingAppointmentAnnotationGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete coaching appointment annotation gateway timeout response has a 4xx status code
func (o *DeleteCoachingAppointmentAnnotationGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete coaching appointment annotation gateway timeout response has a 5xx status code
func (o *DeleteCoachingAppointmentAnnotationGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this delete coaching appointment annotation gateway timeout response a status code equal to that given
func (o *DeleteCoachingAppointmentAnnotationGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *DeleteCoachingAppointmentAnnotationGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/coaching/appointments/{appointmentId}/annotations/{annotationId}][%d] deleteCoachingAppointmentAnnotationGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteCoachingAppointmentAnnotationGatewayTimeout) String() string {
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

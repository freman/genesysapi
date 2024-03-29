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
	case 408:
		result := NewGetCoachingAppointmentAnnotationRequestTimeout()
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

/*
GetCoachingAppointmentAnnotationOK describes a response with status code 200, with default header values.

Annotation retrieved
*/
type GetCoachingAppointmentAnnotationOK struct {
	Payload *models.CoachingAnnotation
}

// IsSuccess returns true when this get coaching appointment annotation o k response has a 2xx status code
func (o *GetCoachingAppointmentAnnotationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get coaching appointment annotation o k response has a 3xx status code
func (o *GetCoachingAppointmentAnnotationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get coaching appointment annotation o k response has a 4xx status code
func (o *GetCoachingAppointmentAnnotationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get coaching appointment annotation o k response has a 5xx status code
func (o *GetCoachingAppointmentAnnotationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get coaching appointment annotation o k response a status code equal to that given
func (o *GetCoachingAppointmentAnnotationOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetCoachingAppointmentAnnotationOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/coaching/appointments/{appointmentId}/annotations/{annotationId}][%d] getCoachingAppointmentAnnotationOK  %+v", 200, o.Payload)
}

func (o *GetCoachingAppointmentAnnotationOK) String() string {
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

/*
GetCoachingAppointmentAnnotationBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetCoachingAppointmentAnnotationBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get coaching appointment annotation bad request response has a 2xx status code
func (o *GetCoachingAppointmentAnnotationBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get coaching appointment annotation bad request response has a 3xx status code
func (o *GetCoachingAppointmentAnnotationBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get coaching appointment annotation bad request response has a 4xx status code
func (o *GetCoachingAppointmentAnnotationBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get coaching appointment annotation bad request response has a 5xx status code
func (o *GetCoachingAppointmentAnnotationBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get coaching appointment annotation bad request response a status code equal to that given
func (o *GetCoachingAppointmentAnnotationBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetCoachingAppointmentAnnotationBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/coaching/appointments/{appointmentId}/annotations/{annotationId}][%d] getCoachingAppointmentAnnotationBadRequest  %+v", 400, o.Payload)
}

func (o *GetCoachingAppointmentAnnotationBadRequest) String() string {
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

/*
GetCoachingAppointmentAnnotationUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetCoachingAppointmentAnnotationUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get coaching appointment annotation unauthorized response has a 2xx status code
func (o *GetCoachingAppointmentAnnotationUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get coaching appointment annotation unauthorized response has a 3xx status code
func (o *GetCoachingAppointmentAnnotationUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get coaching appointment annotation unauthorized response has a 4xx status code
func (o *GetCoachingAppointmentAnnotationUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get coaching appointment annotation unauthorized response has a 5xx status code
func (o *GetCoachingAppointmentAnnotationUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get coaching appointment annotation unauthorized response a status code equal to that given
func (o *GetCoachingAppointmentAnnotationUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetCoachingAppointmentAnnotationUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/coaching/appointments/{appointmentId}/annotations/{annotationId}][%d] getCoachingAppointmentAnnotationUnauthorized  %+v", 401, o.Payload)
}

func (o *GetCoachingAppointmentAnnotationUnauthorized) String() string {
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

/*
GetCoachingAppointmentAnnotationForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetCoachingAppointmentAnnotationForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get coaching appointment annotation forbidden response has a 2xx status code
func (o *GetCoachingAppointmentAnnotationForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get coaching appointment annotation forbidden response has a 3xx status code
func (o *GetCoachingAppointmentAnnotationForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get coaching appointment annotation forbidden response has a 4xx status code
func (o *GetCoachingAppointmentAnnotationForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get coaching appointment annotation forbidden response has a 5xx status code
func (o *GetCoachingAppointmentAnnotationForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get coaching appointment annotation forbidden response a status code equal to that given
func (o *GetCoachingAppointmentAnnotationForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetCoachingAppointmentAnnotationForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/coaching/appointments/{appointmentId}/annotations/{annotationId}][%d] getCoachingAppointmentAnnotationForbidden  %+v", 403, o.Payload)
}

func (o *GetCoachingAppointmentAnnotationForbidden) String() string {
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

/*
GetCoachingAppointmentAnnotationNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetCoachingAppointmentAnnotationNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get coaching appointment annotation not found response has a 2xx status code
func (o *GetCoachingAppointmentAnnotationNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get coaching appointment annotation not found response has a 3xx status code
func (o *GetCoachingAppointmentAnnotationNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get coaching appointment annotation not found response has a 4xx status code
func (o *GetCoachingAppointmentAnnotationNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get coaching appointment annotation not found response has a 5xx status code
func (o *GetCoachingAppointmentAnnotationNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get coaching appointment annotation not found response a status code equal to that given
func (o *GetCoachingAppointmentAnnotationNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetCoachingAppointmentAnnotationNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/coaching/appointments/{appointmentId}/annotations/{annotationId}][%d] getCoachingAppointmentAnnotationNotFound  %+v", 404, o.Payload)
}

func (o *GetCoachingAppointmentAnnotationNotFound) String() string {
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

// NewGetCoachingAppointmentAnnotationRequestTimeout creates a GetCoachingAppointmentAnnotationRequestTimeout with default headers values
func NewGetCoachingAppointmentAnnotationRequestTimeout() *GetCoachingAppointmentAnnotationRequestTimeout {
	return &GetCoachingAppointmentAnnotationRequestTimeout{}
}

/*
GetCoachingAppointmentAnnotationRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetCoachingAppointmentAnnotationRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get coaching appointment annotation request timeout response has a 2xx status code
func (o *GetCoachingAppointmentAnnotationRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get coaching appointment annotation request timeout response has a 3xx status code
func (o *GetCoachingAppointmentAnnotationRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get coaching appointment annotation request timeout response has a 4xx status code
func (o *GetCoachingAppointmentAnnotationRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get coaching appointment annotation request timeout response has a 5xx status code
func (o *GetCoachingAppointmentAnnotationRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get coaching appointment annotation request timeout response a status code equal to that given
func (o *GetCoachingAppointmentAnnotationRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetCoachingAppointmentAnnotationRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/coaching/appointments/{appointmentId}/annotations/{annotationId}][%d] getCoachingAppointmentAnnotationRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetCoachingAppointmentAnnotationRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/coaching/appointments/{appointmentId}/annotations/{annotationId}][%d] getCoachingAppointmentAnnotationRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetCoachingAppointmentAnnotationRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetCoachingAppointmentAnnotationRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*
GetCoachingAppointmentAnnotationRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetCoachingAppointmentAnnotationRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get coaching appointment annotation request entity too large response has a 2xx status code
func (o *GetCoachingAppointmentAnnotationRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get coaching appointment annotation request entity too large response has a 3xx status code
func (o *GetCoachingAppointmentAnnotationRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get coaching appointment annotation request entity too large response has a 4xx status code
func (o *GetCoachingAppointmentAnnotationRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get coaching appointment annotation request entity too large response has a 5xx status code
func (o *GetCoachingAppointmentAnnotationRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get coaching appointment annotation request entity too large response a status code equal to that given
func (o *GetCoachingAppointmentAnnotationRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetCoachingAppointmentAnnotationRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/coaching/appointments/{appointmentId}/annotations/{annotationId}][%d] getCoachingAppointmentAnnotationRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetCoachingAppointmentAnnotationRequestEntityTooLarge) String() string {
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

/*
GetCoachingAppointmentAnnotationUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetCoachingAppointmentAnnotationUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get coaching appointment annotation unsupported media type response has a 2xx status code
func (o *GetCoachingAppointmentAnnotationUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get coaching appointment annotation unsupported media type response has a 3xx status code
func (o *GetCoachingAppointmentAnnotationUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get coaching appointment annotation unsupported media type response has a 4xx status code
func (o *GetCoachingAppointmentAnnotationUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get coaching appointment annotation unsupported media type response has a 5xx status code
func (o *GetCoachingAppointmentAnnotationUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get coaching appointment annotation unsupported media type response a status code equal to that given
func (o *GetCoachingAppointmentAnnotationUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetCoachingAppointmentAnnotationUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/coaching/appointments/{appointmentId}/annotations/{annotationId}][%d] getCoachingAppointmentAnnotationUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetCoachingAppointmentAnnotationUnsupportedMediaType) String() string {
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

/*
GetCoachingAppointmentAnnotationTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetCoachingAppointmentAnnotationTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get coaching appointment annotation too many requests response has a 2xx status code
func (o *GetCoachingAppointmentAnnotationTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get coaching appointment annotation too many requests response has a 3xx status code
func (o *GetCoachingAppointmentAnnotationTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get coaching appointment annotation too many requests response has a 4xx status code
func (o *GetCoachingAppointmentAnnotationTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get coaching appointment annotation too many requests response has a 5xx status code
func (o *GetCoachingAppointmentAnnotationTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get coaching appointment annotation too many requests response a status code equal to that given
func (o *GetCoachingAppointmentAnnotationTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetCoachingAppointmentAnnotationTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/coaching/appointments/{appointmentId}/annotations/{annotationId}][%d] getCoachingAppointmentAnnotationTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetCoachingAppointmentAnnotationTooManyRequests) String() string {
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

/*
GetCoachingAppointmentAnnotationInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetCoachingAppointmentAnnotationInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get coaching appointment annotation internal server error response has a 2xx status code
func (o *GetCoachingAppointmentAnnotationInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get coaching appointment annotation internal server error response has a 3xx status code
func (o *GetCoachingAppointmentAnnotationInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get coaching appointment annotation internal server error response has a 4xx status code
func (o *GetCoachingAppointmentAnnotationInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get coaching appointment annotation internal server error response has a 5xx status code
func (o *GetCoachingAppointmentAnnotationInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get coaching appointment annotation internal server error response a status code equal to that given
func (o *GetCoachingAppointmentAnnotationInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetCoachingAppointmentAnnotationInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/coaching/appointments/{appointmentId}/annotations/{annotationId}][%d] getCoachingAppointmentAnnotationInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCoachingAppointmentAnnotationInternalServerError) String() string {
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

/*
GetCoachingAppointmentAnnotationServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetCoachingAppointmentAnnotationServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get coaching appointment annotation service unavailable response has a 2xx status code
func (o *GetCoachingAppointmentAnnotationServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get coaching appointment annotation service unavailable response has a 3xx status code
func (o *GetCoachingAppointmentAnnotationServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get coaching appointment annotation service unavailable response has a 4xx status code
func (o *GetCoachingAppointmentAnnotationServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get coaching appointment annotation service unavailable response has a 5xx status code
func (o *GetCoachingAppointmentAnnotationServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get coaching appointment annotation service unavailable response a status code equal to that given
func (o *GetCoachingAppointmentAnnotationServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetCoachingAppointmentAnnotationServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/coaching/appointments/{appointmentId}/annotations/{annotationId}][%d] getCoachingAppointmentAnnotationServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetCoachingAppointmentAnnotationServiceUnavailable) String() string {
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

/*
GetCoachingAppointmentAnnotationGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetCoachingAppointmentAnnotationGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get coaching appointment annotation gateway timeout response has a 2xx status code
func (o *GetCoachingAppointmentAnnotationGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get coaching appointment annotation gateway timeout response has a 3xx status code
func (o *GetCoachingAppointmentAnnotationGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get coaching appointment annotation gateway timeout response has a 4xx status code
func (o *GetCoachingAppointmentAnnotationGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get coaching appointment annotation gateway timeout response has a 5xx status code
func (o *GetCoachingAppointmentAnnotationGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get coaching appointment annotation gateway timeout response a status code equal to that given
func (o *GetCoachingAppointmentAnnotationGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetCoachingAppointmentAnnotationGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/coaching/appointments/{appointmentId}/annotations/{annotationId}][%d] getCoachingAppointmentAnnotationGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetCoachingAppointmentAnnotationGatewayTimeout) String() string {
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

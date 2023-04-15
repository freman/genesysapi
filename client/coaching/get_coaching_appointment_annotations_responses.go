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

// GetCoachingAppointmentAnnotationsReader is a Reader for the GetCoachingAppointmentAnnotations structure.
type GetCoachingAppointmentAnnotationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCoachingAppointmentAnnotationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCoachingAppointmentAnnotationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetCoachingAppointmentAnnotationsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetCoachingAppointmentAnnotationsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCoachingAppointmentAnnotationsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetCoachingAppointmentAnnotationsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetCoachingAppointmentAnnotationsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetCoachingAppointmentAnnotationsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetCoachingAppointmentAnnotationsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetCoachingAppointmentAnnotationsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCoachingAppointmentAnnotationsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetCoachingAppointmentAnnotationsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetCoachingAppointmentAnnotationsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCoachingAppointmentAnnotationsOK creates a GetCoachingAppointmentAnnotationsOK with default headers values
func NewGetCoachingAppointmentAnnotationsOK() *GetCoachingAppointmentAnnotationsOK {
	return &GetCoachingAppointmentAnnotationsOK{}
}

/*
GetCoachingAppointmentAnnotationsOK describes a response with status code 200, with default header values.

Annotations retrieved
*/
type GetCoachingAppointmentAnnotationsOK struct {
	Payload *models.CoachingAnnotationList
}

// IsSuccess returns true when this get coaching appointment annotations o k response has a 2xx status code
func (o *GetCoachingAppointmentAnnotationsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get coaching appointment annotations o k response has a 3xx status code
func (o *GetCoachingAppointmentAnnotationsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get coaching appointment annotations o k response has a 4xx status code
func (o *GetCoachingAppointmentAnnotationsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get coaching appointment annotations o k response has a 5xx status code
func (o *GetCoachingAppointmentAnnotationsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get coaching appointment annotations o k response a status code equal to that given
func (o *GetCoachingAppointmentAnnotationsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetCoachingAppointmentAnnotationsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/coaching/appointments/{appointmentId}/annotations][%d] getCoachingAppointmentAnnotationsOK  %+v", 200, o.Payload)
}

func (o *GetCoachingAppointmentAnnotationsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/coaching/appointments/{appointmentId}/annotations][%d] getCoachingAppointmentAnnotationsOK  %+v", 200, o.Payload)
}

func (o *GetCoachingAppointmentAnnotationsOK) GetPayload() *models.CoachingAnnotationList {
	return o.Payload
}

func (o *GetCoachingAppointmentAnnotationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CoachingAnnotationList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCoachingAppointmentAnnotationsBadRequest creates a GetCoachingAppointmentAnnotationsBadRequest with default headers values
func NewGetCoachingAppointmentAnnotationsBadRequest() *GetCoachingAppointmentAnnotationsBadRequest {
	return &GetCoachingAppointmentAnnotationsBadRequest{}
}

/*
GetCoachingAppointmentAnnotationsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetCoachingAppointmentAnnotationsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get coaching appointment annotations bad request response has a 2xx status code
func (o *GetCoachingAppointmentAnnotationsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get coaching appointment annotations bad request response has a 3xx status code
func (o *GetCoachingAppointmentAnnotationsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get coaching appointment annotations bad request response has a 4xx status code
func (o *GetCoachingAppointmentAnnotationsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get coaching appointment annotations bad request response has a 5xx status code
func (o *GetCoachingAppointmentAnnotationsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get coaching appointment annotations bad request response a status code equal to that given
func (o *GetCoachingAppointmentAnnotationsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetCoachingAppointmentAnnotationsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/coaching/appointments/{appointmentId}/annotations][%d] getCoachingAppointmentAnnotationsBadRequest  %+v", 400, o.Payload)
}

func (o *GetCoachingAppointmentAnnotationsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/coaching/appointments/{appointmentId}/annotations][%d] getCoachingAppointmentAnnotationsBadRequest  %+v", 400, o.Payload)
}

func (o *GetCoachingAppointmentAnnotationsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetCoachingAppointmentAnnotationsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCoachingAppointmentAnnotationsUnauthorized creates a GetCoachingAppointmentAnnotationsUnauthorized with default headers values
func NewGetCoachingAppointmentAnnotationsUnauthorized() *GetCoachingAppointmentAnnotationsUnauthorized {
	return &GetCoachingAppointmentAnnotationsUnauthorized{}
}

/*
GetCoachingAppointmentAnnotationsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetCoachingAppointmentAnnotationsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get coaching appointment annotations unauthorized response has a 2xx status code
func (o *GetCoachingAppointmentAnnotationsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get coaching appointment annotations unauthorized response has a 3xx status code
func (o *GetCoachingAppointmentAnnotationsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get coaching appointment annotations unauthorized response has a 4xx status code
func (o *GetCoachingAppointmentAnnotationsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get coaching appointment annotations unauthorized response has a 5xx status code
func (o *GetCoachingAppointmentAnnotationsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get coaching appointment annotations unauthorized response a status code equal to that given
func (o *GetCoachingAppointmentAnnotationsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetCoachingAppointmentAnnotationsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/coaching/appointments/{appointmentId}/annotations][%d] getCoachingAppointmentAnnotationsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetCoachingAppointmentAnnotationsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/coaching/appointments/{appointmentId}/annotations][%d] getCoachingAppointmentAnnotationsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetCoachingAppointmentAnnotationsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetCoachingAppointmentAnnotationsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCoachingAppointmentAnnotationsForbidden creates a GetCoachingAppointmentAnnotationsForbidden with default headers values
func NewGetCoachingAppointmentAnnotationsForbidden() *GetCoachingAppointmentAnnotationsForbidden {
	return &GetCoachingAppointmentAnnotationsForbidden{}
}

/*
GetCoachingAppointmentAnnotationsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetCoachingAppointmentAnnotationsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get coaching appointment annotations forbidden response has a 2xx status code
func (o *GetCoachingAppointmentAnnotationsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get coaching appointment annotations forbidden response has a 3xx status code
func (o *GetCoachingAppointmentAnnotationsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get coaching appointment annotations forbidden response has a 4xx status code
func (o *GetCoachingAppointmentAnnotationsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get coaching appointment annotations forbidden response has a 5xx status code
func (o *GetCoachingAppointmentAnnotationsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get coaching appointment annotations forbidden response a status code equal to that given
func (o *GetCoachingAppointmentAnnotationsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetCoachingAppointmentAnnotationsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/coaching/appointments/{appointmentId}/annotations][%d] getCoachingAppointmentAnnotationsForbidden  %+v", 403, o.Payload)
}

func (o *GetCoachingAppointmentAnnotationsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/coaching/appointments/{appointmentId}/annotations][%d] getCoachingAppointmentAnnotationsForbidden  %+v", 403, o.Payload)
}

func (o *GetCoachingAppointmentAnnotationsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetCoachingAppointmentAnnotationsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCoachingAppointmentAnnotationsNotFound creates a GetCoachingAppointmentAnnotationsNotFound with default headers values
func NewGetCoachingAppointmentAnnotationsNotFound() *GetCoachingAppointmentAnnotationsNotFound {
	return &GetCoachingAppointmentAnnotationsNotFound{}
}

/*
GetCoachingAppointmentAnnotationsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetCoachingAppointmentAnnotationsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get coaching appointment annotations not found response has a 2xx status code
func (o *GetCoachingAppointmentAnnotationsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get coaching appointment annotations not found response has a 3xx status code
func (o *GetCoachingAppointmentAnnotationsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get coaching appointment annotations not found response has a 4xx status code
func (o *GetCoachingAppointmentAnnotationsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get coaching appointment annotations not found response has a 5xx status code
func (o *GetCoachingAppointmentAnnotationsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get coaching appointment annotations not found response a status code equal to that given
func (o *GetCoachingAppointmentAnnotationsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetCoachingAppointmentAnnotationsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/coaching/appointments/{appointmentId}/annotations][%d] getCoachingAppointmentAnnotationsNotFound  %+v", 404, o.Payload)
}

func (o *GetCoachingAppointmentAnnotationsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/coaching/appointments/{appointmentId}/annotations][%d] getCoachingAppointmentAnnotationsNotFound  %+v", 404, o.Payload)
}

func (o *GetCoachingAppointmentAnnotationsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetCoachingAppointmentAnnotationsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCoachingAppointmentAnnotationsRequestTimeout creates a GetCoachingAppointmentAnnotationsRequestTimeout with default headers values
func NewGetCoachingAppointmentAnnotationsRequestTimeout() *GetCoachingAppointmentAnnotationsRequestTimeout {
	return &GetCoachingAppointmentAnnotationsRequestTimeout{}
}

/*
GetCoachingAppointmentAnnotationsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetCoachingAppointmentAnnotationsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get coaching appointment annotations request timeout response has a 2xx status code
func (o *GetCoachingAppointmentAnnotationsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get coaching appointment annotations request timeout response has a 3xx status code
func (o *GetCoachingAppointmentAnnotationsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get coaching appointment annotations request timeout response has a 4xx status code
func (o *GetCoachingAppointmentAnnotationsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get coaching appointment annotations request timeout response has a 5xx status code
func (o *GetCoachingAppointmentAnnotationsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get coaching appointment annotations request timeout response a status code equal to that given
func (o *GetCoachingAppointmentAnnotationsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetCoachingAppointmentAnnotationsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/coaching/appointments/{appointmentId}/annotations][%d] getCoachingAppointmentAnnotationsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetCoachingAppointmentAnnotationsRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/coaching/appointments/{appointmentId}/annotations][%d] getCoachingAppointmentAnnotationsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetCoachingAppointmentAnnotationsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetCoachingAppointmentAnnotationsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCoachingAppointmentAnnotationsRequestEntityTooLarge creates a GetCoachingAppointmentAnnotationsRequestEntityTooLarge with default headers values
func NewGetCoachingAppointmentAnnotationsRequestEntityTooLarge() *GetCoachingAppointmentAnnotationsRequestEntityTooLarge {
	return &GetCoachingAppointmentAnnotationsRequestEntityTooLarge{}
}

/*
GetCoachingAppointmentAnnotationsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetCoachingAppointmentAnnotationsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get coaching appointment annotations request entity too large response has a 2xx status code
func (o *GetCoachingAppointmentAnnotationsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get coaching appointment annotations request entity too large response has a 3xx status code
func (o *GetCoachingAppointmentAnnotationsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get coaching appointment annotations request entity too large response has a 4xx status code
func (o *GetCoachingAppointmentAnnotationsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get coaching appointment annotations request entity too large response has a 5xx status code
func (o *GetCoachingAppointmentAnnotationsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get coaching appointment annotations request entity too large response a status code equal to that given
func (o *GetCoachingAppointmentAnnotationsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetCoachingAppointmentAnnotationsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/coaching/appointments/{appointmentId}/annotations][%d] getCoachingAppointmentAnnotationsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetCoachingAppointmentAnnotationsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/coaching/appointments/{appointmentId}/annotations][%d] getCoachingAppointmentAnnotationsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetCoachingAppointmentAnnotationsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetCoachingAppointmentAnnotationsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCoachingAppointmentAnnotationsUnsupportedMediaType creates a GetCoachingAppointmentAnnotationsUnsupportedMediaType with default headers values
func NewGetCoachingAppointmentAnnotationsUnsupportedMediaType() *GetCoachingAppointmentAnnotationsUnsupportedMediaType {
	return &GetCoachingAppointmentAnnotationsUnsupportedMediaType{}
}

/*
GetCoachingAppointmentAnnotationsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetCoachingAppointmentAnnotationsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get coaching appointment annotations unsupported media type response has a 2xx status code
func (o *GetCoachingAppointmentAnnotationsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get coaching appointment annotations unsupported media type response has a 3xx status code
func (o *GetCoachingAppointmentAnnotationsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get coaching appointment annotations unsupported media type response has a 4xx status code
func (o *GetCoachingAppointmentAnnotationsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get coaching appointment annotations unsupported media type response has a 5xx status code
func (o *GetCoachingAppointmentAnnotationsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get coaching appointment annotations unsupported media type response a status code equal to that given
func (o *GetCoachingAppointmentAnnotationsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetCoachingAppointmentAnnotationsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/coaching/appointments/{appointmentId}/annotations][%d] getCoachingAppointmentAnnotationsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetCoachingAppointmentAnnotationsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/coaching/appointments/{appointmentId}/annotations][%d] getCoachingAppointmentAnnotationsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetCoachingAppointmentAnnotationsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetCoachingAppointmentAnnotationsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCoachingAppointmentAnnotationsTooManyRequests creates a GetCoachingAppointmentAnnotationsTooManyRequests with default headers values
func NewGetCoachingAppointmentAnnotationsTooManyRequests() *GetCoachingAppointmentAnnotationsTooManyRequests {
	return &GetCoachingAppointmentAnnotationsTooManyRequests{}
}

/*
GetCoachingAppointmentAnnotationsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetCoachingAppointmentAnnotationsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get coaching appointment annotations too many requests response has a 2xx status code
func (o *GetCoachingAppointmentAnnotationsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get coaching appointment annotations too many requests response has a 3xx status code
func (o *GetCoachingAppointmentAnnotationsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get coaching appointment annotations too many requests response has a 4xx status code
func (o *GetCoachingAppointmentAnnotationsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get coaching appointment annotations too many requests response has a 5xx status code
func (o *GetCoachingAppointmentAnnotationsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get coaching appointment annotations too many requests response a status code equal to that given
func (o *GetCoachingAppointmentAnnotationsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetCoachingAppointmentAnnotationsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/coaching/appointments/{appointmentId}/annotations][%d] getCoachingAppointmentAnnotationsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetCoachingAppointmentAnnotationsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/coaching/appointments/{appointmentId}/annotations][%d] getCoachingAppointmentAnnotationsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetCoachingAppointmentAnnotationsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetCoachingAppointmentAnnotationsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCoachingAppointmentAnnotationsInternalServerError creates a GetCoachingAppointmentAnnotationsInternalServerError with default headers values
func NewGetCoachingAppointmentAnnotationsInternalServerError() *GetCoachingAppointmentAnnotationsInternalServerError {
	return &GetCoachingAppointmentAnnotationsInternalServerError{}
}

/*
GetCoachingAppointmentAnnotationsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetCoachingAppointmentAnnotationsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get coaching appointment annotations internal server error response has a 2xx status code
func (o *GetCoachingAppointmentAnnotationsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get coaching appointment annotations internal server error response has a 3xx status code
func (o *GetCoachingAppointmentAnnotationsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get coaching appointment annotations internal server error response has a 4xx status code
func (o *GetCoachingAppointmentAnnotationsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get coaching appointment annotations internal server error response has a 5xx status code
func (o *GetCoachingAppointmentAnnotationsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get coaching appointment annotations internal server error response a status code equal to that given
func (o *GetCoachingAppointmentAnnotationsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetCoachingAppointmentAnnotationsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/coaching/appointments/{appointmentId}/annotations][%d] getCoachingAppointmentAnnotationsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCoachingAppointmentAnnotationsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/coaching/appointments/{appointmentId}/annotations][%d] getCoachingAppointmentAnnotationsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCoachingAppointmentAnnotationsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetCoachingAppointmentAnnotationsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCoachingAppointmentAnnotationsServiceUnavailable creates a GetCoachingAppointmentAnnotationsServiceUnavailable with default headers values
func NewGetCoachingAppointmentAnnotationsServiceUnavailable() *GetCoachingAppointmentAnnotationsServiceUnavailable {
	return &GetCoachingAppointmentAnnotationsServiceUnavailable{}
}

/*
GetCoachingAppointmentAnnotationsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetCoachingAppointmentAnnotationsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get coaching appointment annotations service unavailable response has a 2xx status code
func (o *GetCoachingAppointmentAnnotationsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get coaching appointment annotations service unavailable response has a 3xx status code
func (o *GetCoachingAppointmentAnnotationsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get coaching appointment annotations service unavailable response has a 4xx status code
func (o *GetCoachingAppointmentAnnotationsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get coaching appointment annotations service unavailable response has a 5xx status code
func (o *GetCoachingAppointmentAnnotationsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get coaching appointment annotations service unavailable response a status code equal to that given
func (o *GetCoachingAppointmentAnnotationsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetCoachingAppointmentAnnotationsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/coaching/appointments/{appointmentId}/annotations][%d] getCoachingAppointmentAnnotationsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetCoachingAppointmentAnnotationsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/coaching/appointments/{appointmentId}/annotations][%d] getCoachingAppointmentAnnotationsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetCoachingAppointmentAnnotationsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetCoachingAppointmentAnnotationsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCoachingAppointmentAnnotationsGatewayTimeout creates a GetCoachingAppointmentAnnotationsGatewayTimeout with default headers values
func NewGetCoachingAppointmentAnnotationsGatewayTimeout() *GetCoachingAppointmentAnnotationsGatewayTimeout {
	return &GetCoachingAppointmentAnnotationsGatewayTimeout{}
}

/*
GetCoachingAppointmentAnnotationsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetCoachingAppointmentAnnotationsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get coaching appointment annotations gateway timeout response has a 2xx status code
func (o *GetCoachingAppointmentAnnotationsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get coaching appointment annotations gateway timeout response has a 3xx status code
func (o *GetCoachingAppointmentAnnotationsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get coaching appointment annotations gateway timeout response has a 4xx status code
func (o *GetCoachingAppointmentAnnotationsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get coaching appointment annotations gateway timeout response has a 5xx status code
func (o *GetCoachingAppointmentAnnotationsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get coaching appointment annotations gateway timeout response a status code equal to that given
func (o *GetCoachingAppointmentAnnotationsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetCoachingAppointmentAnnotationsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/coaching/appointments/{appointmentId}/annotations][%d] getCoachingAppointmentAnnotationsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetCoachingAppointmentAnnotationsGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/coaching/appointments/{appointmentId}/annotations][%d] getCoachingAppointmentAnnotationsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetCoachingAppointmentAnnotationsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetCoachingAppointmentAnnotationsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

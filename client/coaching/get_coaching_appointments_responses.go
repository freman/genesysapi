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

// GetCoachingAppointmentsReader is a Reader for the GetCoachingAppointments structure.
type GetCoachingAppointmentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCoachingAppointmentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCoachingAppointmentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetCoachingAppointmentsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetCoachingAppointmentsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCoachingAppointmentsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetCoachingAppointmentsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetCoachingAppointmentsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetCoachingAppointmentsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetCoachingAppointmentsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetCoachingAppointmentsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCoachingAppointmentsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetCoachingAppointmentsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetCoachingAppointmentsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCoachingAppointmentsOK creates a GetCoachingAppointmentsOK with default headers values
func NewGetCoachingAppointmentsOK() *GetCoachingAppointmentsOK {
	return &GetCoachingAppointmentsOK{}
}

/*
GetCoachingAppointmentsOK describes a response with status code 200, with default header values.

Get users coaching appointments successful
*/
type GetCoachingAppointmentsOK struct {
	Payload *models.CoachingAppointmentResponseList
}

// IsSuccess returns true when this get coaching appointments o k response has a 2xx status code
func (o *GetCoachingAppointmentsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get coaching appointments o k response has a 3xx status code
func (o *GetCoachingAppointmentsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get coaching appointments o k response has a 4xx status code
func (o *GetCoachingAppointmentsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get coaching appointments o k response has a 5xx status code
func (o *GetCoachingAppointmentsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get coaching appointments o k response a status code equal to that given
func (o *GetCoachingAppointmentsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetCoachingAppointmentsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/coaching/appointments][%d] getCoachingAppointmentsOK  %+v", 200, o.Payload)
}

func (o *GetCoachingAppointmentsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/coaching/appointments][%d] getCoachingAppointmentsOK  %+v", 200, o.Payload)
}

func (o *GetCoachingAppointmentsOK) GetPayload() *models.CoachingAppointmentResponseList {
	return o.Payload
}

func (o *GetCoachingAppointmentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CoachingAppointmentResponseList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCoachingAppointmentsBadRequest creates a GetCoachingAppointmentsBadRequest with default headers values
func NewGetCoachingAppointmentsBadRequest() *GetCoachingAppointmentsBadRequest {
	return &GetCoachingAppointmentsBadRequest{}
}

/*
GetCoachingAppointmentsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetCoachingAppointmentsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get coaching appointments bad request response has a 2xx status code
func (o *GetCoachingAppointmentsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get coaching appointments bad request response has a 3xx status code
func (o *GetCoachingAppointmentsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get coaching appointments bad request response has a 4xx status code
func (o *GetCoachingAppointmentsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get coaching appointments bad request response has a 5xx status code
func (o *GetCoachingAppointmentsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get coaching appointments bad request response a status code equal to that given
func (o *GetCoachingAppointmentsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetCoachingAppointmentsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/coaching/appointments][%d] getCoachingAppointmentsBadRequest  %+v", 400, o.Payload)
}

func (o *GetCoachingAppointmentsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/coaching/appointments][%d] getCoachingAppointmentsBadRequest  %+v", 400, o.Payload)
}

func (o *GetCoachingAppointmentsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetCoachingAppointmentsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCoachingAppointmentsUnauthorized creates a GetCoachingAppointmentsUnauthorized with default headers values
func NewGetCoachingAppointmentsUnauthorized() *GetCoachingAppointmentsUnauthorized {
	return &GetCoachingAppointmentsUnauthorized{}
}

/*
GetCoachingAppointmentsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetCoachingAppointmentsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get coaching appointments unauthorized response has a 2xx status code
func (o *GetCoachingAppointmentsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get coaching appointments unauthorized response has a 3xx status code
func (o *GetCoachingAppointmentsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get coaching appointments unauthorized response has a 4xx status code
func (o *GetCoachingAppointmentsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get coaching appointments unauthorized response has a 5xx status code
func (o *GetCoachingAppointmentsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get coaching appointments unauthorized response a status code equal to that given
func (o *GetCoachingAppointmentsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetCoachingAppointmentsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/coaching/appointments][%d] getCoachingAppointmentsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetCoachingAppointmentsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/coaching/appointments][%d] getCoachingAppointmentsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetCoachingAppointmentsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetCoachingAppointmentsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCoachingAppointmentsForbidden creates a GetCoachingAppointmentsForbidden with default headers values
func NewGetCoachingAppointmentsForbidden() *GetCoachingAppointmentsForbidden {
	return &GetCoachingAppointmentsForbidden{}
}

/*
GetCoachingAppointmentsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetCoachingAppointmentsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get coaching appointments forbidden response has a 2xx status code
func (o *GetCoachingAppointmentsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get coaching appointments forbidden response has a 3xx status code
func (o *GetCoachingAppointmentsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get coaching appointments forbidden response has a 4xx status code
func (o *GetCoachingAppointmentsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get coaching appointments forbidden response has a 5xx status code
func (o *GetCoachingAppointmentsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get coaching appointments forbidden response a status code equal to that given
func (o *GetCoachingAppointmentsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetCoachingAppointmentsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/coaching/appointments][%d] getCoachingAppointmentsForbidden  %+v", 403, o.Payload)
}

func (o *GetCoachingAppointmentsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/coaching/appointments][%d] getCoachingAppointmentsForbidden  %+v", 403, o.Payload)
}

func (o *GetCoachingAppointmentsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetCoachingAppointmentsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCoachingAppointmentsNotFound creates a GetCoachingAppointmentsNotFound with default headers values
func NewGetCoachingAppointmentsNotFound() *GetCoachingAppointmentsNotFound {
	return &GetCoachingAppointmentsNotFound{}
}

/*
GetCoachingAppointmentsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetCoachingAppointmentsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get coaching appointments not found response has a 2xx status code
func (o *GetCoachingAppointmentsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get coaching appointments not found response has a 3xx status code
func (o *GetCoachingAppointmentsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get coaching appointments not found response has a 4xx status code
func (o *GetCoachingAppointmentsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get coaching appointments not found response has a 5xx status code
func (o *GetCoachingAppointmentsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get coaching appointments not found response a status code equal to that given
func (o *GetCoachingAppointmentsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetCoachingAppointmentsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/coaching/appointments][%d] getCoachingAppointmentsNotFound  %+v", 404, o.Payload)
}

func (o *GetCoachingAppointmentsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/coaching/appointments][%d] getCoachingAppointmentsNotFound  %+v", 404, o.Payload)
}

func (o *GetCoachingAppointmentsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetCoachingAppointmentsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCoachingAppointmentsRequestTimeout creates a GetCoachingAppointmentsRequestTimeout with default headers values
func NewGetCoachingAppointmentsRequestTimeout() *GetCoachingAppointmentsRequestTimeout {
	return &GetCoachingAppointmentsRequestTimeout{}
}

/*
GetCoachingAppointmentsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetCoachingAppointmentsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get coaching appointments request timeout response has a 2xx status code
func (o *GetCoachingAppointmentsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get coaching appointments request timeout response has a 3xx status code
func (o *GetCoachingAppointmentsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get coaching appointments request timeout response has a 4xx status code
func (o *GetCoachingAppointmentsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get coaching appointments request timeout response has a 5xx status code
func (o *GetCoachingAppointmentsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get coaching appointments request timeout response a status code equal to that given
func (o *GetCoachingAppointmentsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetCoachingAppointmentsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/coaching/appointments][%d] getCoachingAppointmentsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetCoachingAppointmentsRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/coaching/appointments][%d] getCoachingAppointmentsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetCoachingAppointmentsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetCoachingAppointmentsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCoachingAppointmentsRequestEntityTooLarge creates a GetCoachingAppointmentsRequestEntityTooLarge with default headers values
func NewGetCoachingAppointmentsRequestEntityTooLarge() *GetCoachingAppointmentsRequestEntityTooLarge {
	return &GetCoachingAppointmentsRequestEntityTooLarge{}
}

/*
GetCoachingAppointmentsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetCoachingAppointmentsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get coaching appointments request entity too large response has a 2xx status code
func (o *GetCoachingAppointmentsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get coaching appointments request entity too large response has a 3xx status code
func (o *GetCoachingAppointmentsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get coaching appointments request entity too large response has a 4xx status code
func (o *GetCoachingAppointmentsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get coaching appointments request entity too large response has a 5xx status code
func (o *GetCoachingAppointmentsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get coaching appointments request entity too large response a status code equal to that given
func (o *GetCoachingAppointmentsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetCoachingAppointmentsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/coaching/appointments][%d] getCoachingAppointmentsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetCoachingAppointmentsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/coaching/appointments][%d] getCoachingAppointmentsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetCoachingAppointmentsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetCoachingAppointmentsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCoachingAppointmentsUnsupportedMediaType creates a GetCoachingAppointmentsUnsupportedMediaType with default headers values
func NewGetCoachingAppointmentsUnsupportedMediaType() *GetCoachingAppointmentsUnsupportedMediaType {
	return &GetCoachingAppointmentsUnsupportedMediaType{}
}

/*
GetCoachingAppointmentsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetCoachingAppointmentsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get coaching appointments unsupported media type response has a 2xx status code
func (o *GetCoachingAppointmentsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get coaching appointments unsupported media type response has a 3xx status code
func (o *GetCoachingAppointmentsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get coaching appointments unsupported media type response has a 4xx status code
func (o *GetCoachingAppointmentsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get coaching appointments unsupported media type response has a 5xx status code
func (o *GetCoachingAppointmentsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get coaching appointments unsupported media type response a status code equal to that given
func (o *GetCoachingAppointmentsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetCoachingAppointmentsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/coaching/appointments][%d] getCoachingAppointmentsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetCoachingAppointmentsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/coaching/appointments][%d] getCoachingAppointmentsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetCoachingAppointmentsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetCoachingAppointmentsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCoachingAppointmentsTooManyRequests creates a GetCoachingAppointmentsTooManyRequests with default headers values
func NewGetCoachingAppointmentsTooManyRequests() *GetCoachingAppointmentsTooManyRequests {
	return &GetCoachingAppointmentsTooManyRequests{}
}

/*
GetCoachingAppointmentsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetCoachingAppointmentsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get coaching appointments too many requests response has a 2xx status code
func (o *GetCoachingAppointmentsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get coaching appointments too many requests response has a 3xx status code
func (o *GetCoachingAppointmentsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get coaching appointments too many requests response has a 4xx status code
func (o *GetCoachingAppointmentsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get coaching appointments too many requests response has a 5xx status code
func (o *GetCoachingAppointmentsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get coaching appointments too many requests response a status code equal to that given
func (o *GetCoachingAppointmentsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetCoachingAppointmentsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/coaching/appointments][%d] getCoachingAppointmentsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetCoachingAppointmentsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/coaching/appointments][%d] getCoachingAppointmentsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetCoachingAppointmentsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetCoachingAppointmentsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCoachingAppointmentsInternalServerError creates a GetCoachingAppointmentsInternalServerError with default headers values
func NewGetCoachingAppointmentsInternalServerError() *GetCoachingAppointmentsInternalServerError {
	return &GetCoachingAppointmentsInternalServerError{}
}

/*
GetCoachingAppointmentsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetCoachingAppointmentsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get coaching appointments internal server error response has a 2xx status code
func (o *GetCoachingAppointmentsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get coaching appointments internal server error response has a 3xx status code
func (o *GetCoachingAppointmentsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get coaching appointments internal server error response has a 4xx status code
func (o *GetCoachingAppointmentsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get coaching appointments internal server error response has a 5xx status code
func (o *GetCoachingAppointmentsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get coaching appointments internal server error response a status code equal to that given
func (o *GetCoachingAppointmentsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetCoachingAppointmentsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/coaching/appointments][%d] getCoachingAppointmentsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCoachingAppointmentsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/coaching/appointments][%d] getCoachingAppointmentsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCoachingAppointmentsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetCoachingAppointmentsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCoachingAppointmentsServiceUnavailable creates a GetCoachingAppointmentsServiceUnavailable with default headers values
func NewGetCoachingAppointmentsServiceUnavailable() *GetCoachingAppointmentsServiceUnavailable {
	return &GetCoachingAppointmentsServiceUnavailable{}
}

/*
GetCoachingAppointmentsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetCoachingAppointmentsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get coaching appointments service unavailable response has a 2xx status code
func (o *GetCoachingAppointmentsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get coaching appointments service unavailable response has a 3xx status code
func (o *GetCoachingAppointmentsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get coaching appointments service unavailable response has a 4xx status code
func (o *GetCoachingAppointmentsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get coaching appointments service unavailable response has a 5xx status code
func (o *GetCoachingAppointmentsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get coaching appointments service unavailable response a status code equal to that given
func (o *GetCoachingAppointmentsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetCoachingAppointmentsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/coaching/appointments][%d] getCoachingAppointmentsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetCoachingAppointmentsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/coaching/appointments][%d] getCoachingAppointmentsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetCoachingAppointmentsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetCoachingAppointmentsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCoachingAppointmentsGatewayTimeout creates a GetCoachingAppointmentsGatewayTimeout with default headers values
func NewGetCoachingAppointmentsGatewayTimeout() *GetCoachingAppointmentsGatewayTimeout {
	return &GetCoachingAppointmentsGatewayTimeout{}
}

/*
GetCoachingAppointmentsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetCoachingAppointmentsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get coaching appointments gateway timeout response has a 2xx status code
func (o *GetCoachingAppointmentsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get coaching appointments gateway timeout response has a 3xx status code
func (o *GetCoachingAppointmentsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get coaching appointments gateway timeout response has a 4xx status code
func (o *GetCoachingAppointmentsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get coaching appointments gateway timeout response has a 5xx status code
func (o *GetCoachingAppointmentsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get coaching appointments gateway timeout response a status code equal to that given
func (o *GetCoachingAppointmentsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetCoachingAppointmentsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/coaching/appointments][%d] getCoachingAppointmentsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetCoachingAppointmentsGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/coaching/appointments][%d] getCoachingAppointmentsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetCoachingAppointmentsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetCoachingAppointmentsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package analytics

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetAnalyticsReportingScheduleReader is a Reader for the GetAnalyticsReportingSchedule structure.
type GetAnalyticsReportingScheduleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAnalyticsReportingScheduleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAnalyticsReportingScheduleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAnalyticsReportingScheduleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAnalyticsReportingScheduleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAnalyticsReportingScheduleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAnalyticsReportingScheduleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetAnalyticsReportingScheduleRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetAnalyticsReportingScheduleRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetAnalyticsReportingScheduleUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetAnalyticsReportingScheduleTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAnalyticsReportingScheduleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetAnalyticsReportingScheduleServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetAnalyticsReportingScheduleGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAnalyticsReportingScheduleOK creates a GetAnalyticsReportingScheduleOK with default headers values
func NewGetAnalyticsReportingScheduleOK() *GetAnalyticsReportingScheduleOK {
	return &GetAnalyticsReportingScheduleOK{}
}

/*
GetAnalyticsReportingScheduleOK describes a response with status code 200, with default header values.

successful operation
*/
type GetAnalyticsReportingScheduleOK struct {
	Payload *models.ReportSchedule
}

// IsSuccess returns true when this get analytics reporting schedule o k response has a 2xx status code
func (o *GetAnalyticsReportingScheduleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get analytics reporting schedule o k response has a 3xx status code
func (o *GetAnalyticsReportingScheduleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get analytics reporting schedule o k response has a 4xx status code
func (o *GetAnalyticsReportingScheduleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get analytics reporting schedule o k response has a 5xx status code
func (o *GetAnalyticsReportingScheduleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get analytics reporting schedule o k response a status code equal to that given
func (o *GetAnalyticsReportingScheduleOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetAnalyticsReportingScheduleOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/schedules/{scheduleId}][%d] getAnalyticsReportingScheduleOK  %+v", 200, o.Payload)
}

func (o *GetAnalyticsReportingScheduleOK) String() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/schedules/{scheduleId}][%d] getAnalyticsReportingScheduleOK  %+v", 200, o.Payload)
}

func (o *GetAnalyticsReportingScheduleOK) GetPayload() *models.ReportSchedule {
	return o.Payload
}

func (o *GetAnalyticsReportingScheduleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ReportSchedule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingScheduleBadRequest creates a GetAnalyticsReportingScheduleBadRequest with default headers values
func NewGetAnalyticsReportingScheduleBadRequest() *GetAnalyticsReportingScheduleBadRequest {
	return &GetAnalyticsReportingScheduleBadRequest{}
}

/*
GetAnalyticsReportingScheduleBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetAnalyticsReportingScheduleBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get analytics reporting schedule bad request response has a 2xx status code
func (o *GetAnalyticsReportingScheduleBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get analytics reporting schedule bad request response has a 3xx status code
func (o *GetAnalyticsReportingScheduleBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get analytics reporting schedule bad request response has a 4xx status code
func (o *GetAnalyticsReportingScheduleBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get analytics reporting schedule bad request response has a 5xx status code
func (o *GetAnalyticsReportingScheduleBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get analytics reporting schedule bad request response a status code equal to that given
func (o *GetAnalyticsReportingScheduleBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetAnalyticsReportingScheduleBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/schedules/{scheduleId}][%d] getAnalyticsReportingScheduleBadRequest  %+v", 400, o.Payload)
}

func (o *GetAnalyticsReportingScheduleBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/schedules/{scheduleId}][%d] getAnalyticsReportingScheduleBadRequest  %+v", 400, o.Payload)
}

func (o *GetAnalyticsReportingScheduleBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingScheduleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingScheduleUnauthorized creates a GetAnalyticsReportingScheduleUnauthorized with default headers values
func NewGetAnalyticsReportingScheduleUnauthorized() *GetAnalyticsReportingScheduleUnauthorized {
	return &GetAnalyticsReportingScheduleUnauthorized{}
}

/*
GetAnalyticsReportingScheduleUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetAnalyticsReportingScheduleUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get analytics reporting schedule unauthorized response has a 2xx status code
func (o *GetAnalyticsReportingScheduleUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get analytics reporting schedule unauthorized response has a 3xx status code
func (o *GetAnalyticsReportingScheduleUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get analytics reporting schedule unauthorized response has a 4xx status code
func (o *GetAnalyticsReportingScheduleUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get analytics reporting schedule unauthorized response has a 5xx status code
func (o *GetAnalyticsReportingScheduleUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get analytics reporting schedule unauthorized response a status code equal to that given
func (o *GetAnalyticsReportingScheduleUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetAnalyticsReportingScheduleUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/schedules/{scheduleId}][%d] getAnalyticsReportingScheduleUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAnalyticsReportingScheduleUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/schedules/{scheduleId}][%d] getAnalyticsReportingScheduleUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAnalyticsReportingScheduleUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingScheduleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingScheduleForbidden creates a GetAnalyticsReportingScheduleForbidden with default headers values
func NewGetAnalyticsReportingScheduleForbidden() *GetAnalyticsReportingScheduleForbidden {
	return &GetAnalyticsReportingScheduleForbidden{}
}

/*
GetAnalyticsReportingScheduleForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetAnalyticsReportingScheduleForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get analytics reporting schedule forbidden response has a 2xx status code
func (o *GetAnalyticsReportingScheduleForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get analytics reporting schedule forbidden response has a 3xx status code
func (o *GetAnalyticsReportingScheduleForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get analytics reporting schedule forbidden response has a 4xx status code
func (o *GetAnalyticsReportingScheduleForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get analytics reporting schedule forbidden response has a 5xx status code
func (o *GetAnalyticsReportingScheduleForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get analytics reporting schedule forbidden response a status code equal to that given
func (o *GetAnalyticsReportingScheduleForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetAnalyticsReportingScheduleForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/schedules/{scheduleId}][%d] getAnalyticsReportingScheduleForbidden  %+v", 403, o.Payload)
}

func (o *GetAnalyticsReportingScheduleForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/schedules/{scheduleId}][%d] getAnalyticsReportingScheduleForbidden  %+v", 403, o.Payload)
}

func (o *GetAnalyticsReportingScheduleForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingScheduleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingScheduleNotFound creates a GetAnalyticsReportingScheduleNotFound with default headers values
func NewGetAnalyticsReportingScheduleNotFound() *GetAnalyticsReportingScheduleNotFound {
	return &GetAnalyticsReportingScheduleNotFound{}
}

/*
GetAnalyticsReportingScheduleNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetAnalyticsReportingScheduleNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get analytics reporting schedule not found response has a 2xx status code
func (o *GetAnalyticsReportingScheduleNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get analytics reporting schedule not found response has a 3xx status code
func (o *GetAnalyticsReportingScheduleNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get analytics reporting schedule not found response has a 4xx status code
func (o *GetAnalyticsReportingScheduleNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get analytics reporting schedule not found response has a 5xx status code
func (o *GetAnalyticsReportingScheduleNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get analytics reporting schedule not found response a status code equal to that given
func (o *GetAnalyticsReportingScheduleNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetAnalyticsReportingScheduleNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/schedules/{scheduleId}][%d] getAnalyticsReportingScheduleNotFound  %+v", 404, o.Payload)
}

func (o *GetAnalyticsReportingScheduleNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/schedules/{scheduleId}][%d] getAnalyticsReportingScheduleNotFound  %+v", 404, o.Payload)
}

func (o *GetAnalyticsReportingScheduleNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingScheduleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingScheduleRequestTimeout creates a GetAnalyticsReportingScheduleRequestTimeout with default headers values
func NewGetAnalyticsReportingScheduleRequestTimeout() *GetAnalyticsReportingScheduleRequestTimeout {
	return &GetAnalyticsReportingScheduleRequestTimeout{}
}

/*
GetAnalyticsReportingScheduleRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetAnalyticsReportingScheduleRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get analytics reporting schedule request timeout response has a 2xx status code
func (o *GetAnalyticsReportingScheduleRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get analytics reporting schedule request timeout response has a 3xx status code
func (o *GetAnalyticsReportingScheduleRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get analytics reporting schedule request timeout response has a 4xx status code
func (o *GetAnalyticsReportingScheduleRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get analytics reporting schedule request timeout response has a 5xx status code
func (o *GetAnalyticsReportingScheduleRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get analytics reporting schedule request timeout response a status code equal to that given
func (o *GetAnalyticsReportingScheduleRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetAnalyticsReportingScheduleRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/schedules/{scheduleId}][%d] getAnalyticsReportingScheduleRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetAnalyticsReportingScheduleRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/schedules/{scheduleId}][%d] getAnalyticsReportingScheduleRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetAnalyticsReportingScheduleRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingScheduleRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingScheduleRequestEntityTooLarge creates a GetAnalyticsReportingScheduleRequestEntityTooLarge with default headers values
func NewGetAnalyticsReportingScheduleRequestEntityTooLarge() *GetAnalyticsReportingScheduleRequestEntityTooLarge {
	return &GetAnalyticsReportingScheduleRequestEntityTooLarge{}
}

/*
GetAnalyticsReportingScheduleRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetAnalyticsReportingScheduleRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get analytics reporting schedule request entity too large response has a 2xx status code
func (o *GetAnalyticsReportingScheduleRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get analytics reporting schedule request entity too large response has a 3xx status code
func (o *GetAnalyticsReportingScheduleRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get analytics reporting schedule request entity too large response has a 4xx status code
func (o *GetAnalyticsReportingScheduleRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get analytics reporting schedule request entity too large response has a 5xx status code
func (o *GetAnalyticsReportingScheduleRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get analytics reporting schedule request entity too large response a status code equal to that given
func (o *GetAnalyticsReportingScheduleRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetAnalyticsReportingScheduleRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/schedules/{scheduleId}][%d] getAnalyticsReportingScheduleRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetAnalyticsReportingScheduleRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/schedules/{scheduleId}][%d] getAnalyticsReportingScheduleRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetAnalyticsReportingScheduleRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingScheduleRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingScheduleUnsupportedMediaType creates a GetAnalyticsReportingScheduleUnsupportedMediaType with default headers values
func NewGetAnalyticsReportingScheduleUnsupportedMediaType() *GetAnalyticsReportingScheduleUnsupportedMediaType {
	return &GetAnalyticsReportingScheduleUnsupportedMediaType{}
}

/*
GetAnalyticsReportingScheduleUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetAnalyticsReportingScheduleUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get analytics reporting schedule unsupported media type response has a 2xx status code
func (o *GetAnalyticsReportingScheduleUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get analytics reporting schedule unsupported media type response has a 3xx status code
func (o *GetAnalyticsReportingScheduleUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get analytics reporting schedule unsupported media type response has a 4xx status code
func (o *GetAnalyticsReportingScheduleUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get analytics reporting schedule unsupported media type response has a 5xx status code
func (o *GetAnalyticsReportingScheduleUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get analytics reporting schedule unsupported media type response a status code equal to that given
func (o *GetAnalyticsReportingScheduleUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetAnalyticsReportingScheduleUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/schedules/{scheduleId}][%d] getAnalyticsReportingScheduleUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetAnalyticsReportingScheduleUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/schedules/{scheduleId}][%d] getAnalyticsReportingScheduleUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetAnalyticsReportingScheduleUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingScheduleUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingScheduleTooManyRequests creates a GetAnalyticsReportingScheduleTooManyRequests with default headers values
func NewGetAnalyticsReportingScheduleTooManyRequests() *GetAnalyticsReportingScheduleTooManyRequests {
	return &GetAnalyticsReportingScheduleTooManyRequests{}
}

/*
GetAnalyticsReportingScheduleTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetAnalyticsReportingScheduleTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get analytics reporting schedule too many requests response has a 2xx status code
func (o *GetAnalyticsReportingScheduleTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get analytics reporting schedule too many requests response has a 3xx status code
func (o *GetAnalyticsReportingScheduleTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get analytics reporting schedule too many requests response has a 4xx status code
func (o *GetAnalyticsReportingScheduleTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get analytics reporting schedule too many requests response has a 5xx status code
func (o *GetAnalyticsReportingScheduleTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get analytics reporting schedule too many requests response a status code equal to that given
func (o *GetAnalyticsReportingScheduleTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetAnalyticsReportingScheduleTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/schedules/{scheduleId}][%d] getAnalyticsReportingScheduleTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetAnalyticsReportingScheduleTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/schedules/{scheduleId}][%d] getAnalyticsReportingScheduleTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetAnalyticsReportingScheduleTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingScheduleTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingScheduleInternalServerError creates a GetAnalyticsReportingScheduleInternalServerError with default headers values
func NewGetAnalyticsReportingScheduleInternalServerError() *GetAnalyticsReportingScheduleInternalServerError {
	return &GetAnalyticsReportingScheduleInternalServerError{}
}

/*
GetAnalyticsReportingScheduleInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetAnalyticsReportingScheduleInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get analytics reporting schedule internal server error response has a 2xx status code
func (o *GetAnalyticsReportingScheduleInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get analytics reporting schedule internal server error response has a 3xx status code
func (o *GetAnalyticsReportingScheduleInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get analytics reporting schedule internal server error response has a 4xx status code
func (o *GetAnalyticsReportingScheduleInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get analytics reporting schedule internal server error response has a 5xx status code
func (o *GetAnalyticsReportingScheduleInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get analytics reporting schedule internal server error response a status code equal to that given
func (o *GetAnalyticsReportingScheduleInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetAnalyticsReportingScheduleInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/schedules/{scheduleId}][%d] getAnalyticsReportingScheduleInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAnalyticsReportingScheduleInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/schedules/{scheduleId}][%d] getAnalyticsReportingScheduleInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAnalyticsReportingScheduleInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingScheduleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingScheduleServiceUnavailable creates a GetAnalyticsReportingScheduleServiceUnavailable with default headers values
func NewGetAnalyticsReportingScheduleServiceUnavailable() *GetAnalyticsReportingScheduleServiceUnavailable {
	return &GetAnalyticsReportingScheduleServiceUnavailable{}
}

/*
GetAnalyticsReportingScheduleServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetAnalyticsReportingScheduleServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get analytics reporting schedule service unavailable response has a 2xx status code
func (o *GetAnalyticsReportingScheduleServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get analytics reporting schedule service unavailable response has a 3xx status code
func (o *GetAnalyticsReportingScheduleServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get analytics reporting schedule service unavailable response has a 4xx status code
func (o *GetAnalyticsReportingScheduleServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get analytics reporting schedule service unavailable response has a 5xx status code
func (o *GetAnalyticsReportingScheduleServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get analytics reporting schedule service unavailable response a status code equal to that given
func (o *GetAnalyticsReportingScheduleServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetAnalyticsReportingScheduleServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/schedules/{scheduleId}][%d] getAnalyticsReportingScheduleServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetAnalyticsReportingScheduleServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/schedules/{scheduleId}][%d] getAnalyticsReportingScheduleServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetAnalyticsReportingScheduleServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingScheduleServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingScheduleGatewayTimeout creates a GetAnalyticsReportingScheduleGatewayTimeout with default headers values
func NewGetAnalyticsReportingScheduleGatewayTimeout() *GetAnalyticsReportingScheduleGatewayTimeout {
	return &GetAnalyticsReportingScheduleGatewayTimeout{}
}

/*
GetAnalyticsReportingScheduleGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetAnalyticsReportingScheduleGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get analytics reporting schedule gateway timeout response has a 2xx status code
func (o *GetAnalyticsReportingScheduleGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get analytics reporting schedule gateway timeout response has a 3xx status code
func (o *GetAnalyticsReportingScheduleGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get analytics reporting schedule gateway timeout response has a 4xx status code
func (o *GetAnalyticsReportingScheduleGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get analytics reporting schedule gateway timeout response has a 5xx status code
func (o *GetAnalyticsReportingScheduleGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get analytics reporting schedule gateway timeout response a status code equal to that given
func (o *GetAnalyticsReportingScheduleGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetAnalyticsReportingScheduleGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/schedules/{scheduleId}][%d] getAnalyticsReportingScheduleGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetAnalyticsReportingScheduleGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/schedules/{scheduleId}][%d] getAnalyticsReportingScheduleGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetAnalyticsReportingScheduleGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingScheduleGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

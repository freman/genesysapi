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

// PutAnalyticsReportingScheduleReader is a Reader for the PutAnalyticsReportingSchedule structure.
type PutAnalyticsReportingScheduleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutAnalyticsReportingScheduleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutAnalyticsReportingScheduleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutAnalyticsReportingScheduleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutAnalyticsReportingScheduleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutAnalyticsReportingScheduleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutAnalyticsReportingScheduleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPutAnalyticsReportingScheduleRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPutAnalyticsReportingScheduleRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPutAnalyticsReportingScheduleUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutAnalyticsReportingScheduleTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutAnalyticsReportingScheduleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutAnalyticsReportingScheduleServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutAnalyticsReportingScheduleGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutAnalyticsReportingScheduleOK creates a PutAnalyticsReportingScheduleOK with default headers values
func NewPutAnalyticsReportingScheduleOK() *PutAnalyticsReportingScheduleOK {
	return &PutAnalyticsReportingScheduleOK{}
}

/*PutAnalyticsReportingScheduleOK handles this case with default header values.

successful operation
*/
type PutAnalyticsReportingScheduleOK struct {
	Payload *models.ReportSchedule
}

func (o *PutAnalyticsReportingScheduleOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/analytics/reporting/schedules/{scheduleId}][%d] putAnalyticsReportingScheduleOK  %+v", 200, o.Payload)
}

func (o *PutAnalyticsReportingScheduleOK) GetPayload() *models.ReportSchedule {
	return o.Payload
}

func (o *PutAnalyticsReportingScheduleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ReportSchedule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAnalyticsReportingScheduleBadRequest creates a PutAnalyticsReportingScheduleBadRequest with default headers values
func NewPutAnalyticsReportingScheduleBadRequest() *PutAnalyticsReportingScheduleBadRequest {
	return &PutAnalyticsReportingScheduleBadRequest{}
}

/*PutAnalyticsReportingScheduleBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutAnalyticsReportingScheduleBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PutAnalyticsReportingScheduleBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/analytics/reporting/schedules/{scheduleId}][%d] putAnalyticsReportingScheduleBadRequest  %+v", 400, o.Payload)
}

func (o *PutAnalyticsReportingScheduleBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutAnalyticsReportingScheduleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAnalyticsReportingScheduleUnauthorized creates a PutAnalyticsReportingScheduleUnauthorized with default headers values
func NewPutAnalyticsReportingScheduleUnauthorized() *PutAnalyticsReportingScheduleUnauthorized {
	return &PutAnalyticsReportingScheduleUnauthorized{}
}

/*PutAnalyticsReportingScheduleUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PutAnalyticsReportingScheduleUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PutAnalyticsReportingScheduleUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/analytics/reporting/schedules/{scheduleId}][%d] putAnalyticsReportingScheduleUnauthorized  %+v", 401, o.Payload)
}

func (o *PutAnalyticsReportingScheduleUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutAnalyticsReportingScheduleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAnalyticsReportingScheduleForbidden creates a PutAnalyticsReportingScheduleForbidden with default headers values
func NewPutAnalyticsReportingScheduleForbidden() *PutAnalyticsReportingScheduleForbidden {
	return &PutAnalyticsReportingScheduleForbidden{}
}

/*PutAnalyticsReportingScheduleForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PutAnalyticsReportingScheduleForbidden struct {
	Payload *models.ErrorBody
}

func (o *PutAnalyticsReportingScheduleForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/analytics/reporting/schedules/{scheduleId}][%d] putAnalyticsReportingScheduleForbidden  %+v", 403, o.Payload)
}

func (o *PutAnalyticsReportingScheduleForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutAnalyticsReportingScheduleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAnalyticsReportingScheduleNotFound creates a PutAnalyticsReportingScheduleNotFound with default headers values
func NewPutAnalyticsReportingScheduleNotFound() *PutAnalyticsReportingScheduleNotFound {
	return &PutAnalyticsReportingScheduleNotFound{}
}

/*PutAnalyticsReportingScheduleNotFound handles this case with default header values.

The requested resource was not found.
*/
type PutAnalyticsReportingScheduleNotFound struct {
	Payload *models.ErrorBody
}

func (o *PutAnalyticsReportingScheduleNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/analytics/reporting/schedules/{scheduleId}][%d] putAnalyticsReportingScheduleNotFound  %+v", 404, o.Payload)
}

func (o *PutAnalyticsReportingScheduleNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutAnalyticsReportingScheduleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAnalyticsReportingScheduleRequestTimeout creates a PutAnalyticsReportingScheduleRequestTimeout with default headers values
func NewPutAnalyticsReportingScheduleRequestTimeout() *PutAnalyticsReportingScheduleRequestTimeout {
	return &PutAnalyticsReportingScheduleRequestTimeout{}
}

/*PutAnalyticsReportingScheduleRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PutAnalyticsReportingScheduleRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PutAnalyticsReportingScheduleRequestTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/analytics/reporting/schedules/{scheduleId}][%d] putAnalyticsReportingScheduleRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutAnalyticsReportingScheduleRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutAnalyticsReportingScheduleRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAnalyticsReportingScheduleRequestEntityTooLarge creates a PutAnalyticsReportingScheduleRequestEntityTooLarge with default headers values
func NewPutAnalyticsReportingScheduleRequestEntityTooLarge() *PutAnalyticsReportingScheduleRequestEntityTooLarge {
	return &PutAnalyticsReportingScheduleRequestEntityTooLarge{}
}

/*PutAnalyticsReportingScheduleRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PutAnalyticsReportingScheduleRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PutAnalyticsReportingScheduleRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /api/v2/analytics/reporting/schedules/{scheduleId}][%d] putAnalyticsReportingScheduleRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutAnalyticsReportingScheduleRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutAnalyticsReportingScheduleRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAnalyticsReportingScheduleUnsupportedMediaType creates a PutAnalyticsReportingScheduleUnsupportedMediaType with default headers values
func NewPutAnalyticsReportingScheduleUnsupportedMediaType() *PutAnalyticsReportingScheduleUnsupportedMediaType {
	return &PutAnalyticsReportingScheduleUnsupportedMediaType{}
}

/*PutAnalyticsReportingScheduleUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutAnalyticsReportingScheduleUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PutAnalyticsReportingScheduleUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v2/analytics/reporting/schedules/{scheduleId}][%d] putAnalyticsReportingScheduleUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutAnalyticsReportingScheduleUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutAnalyticsReportingScheduleUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAnalyticsReportingScheduleTooManyRequests creates a PutAnalyticsReportingScheduleTooManyRequests with default headers values
func NewPutAnalyticsReportingScheduleTooManyRequests() *PutAnalyticsReportingScheduleTooManyRequests {
	return &PutAnalyticsReportingScheduleTooManyRequests{}
}

/*PutAnalyticsReportingScheduleTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PutAnalyticsReportingScheduleTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PutAnalyticsReportingScheduleTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v2/analytics/reporting/schedules/{scheduleId}][%d] putAnalyticsReportingScheduleTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutAnalyticsReportingScheduleTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutAnalyticsReportingScheduleTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAnalyticsReportingScheduleInternalServerError creates a PutAnalyticsReportingScheduleInternalServerError with default headers values
func NewPutAnalyticsReportingScheduleInternalServerError() *PutAnalyticsReportingScheduleInternalServerError {
	return &PutAnalyticsReportingScheduleInternalServerError{}
}

/*PutAnalyticsReportingScheduleInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutAnalyticsReportingScheduleInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PutAnalyticsReportingScheduleInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/analytics/reporting/schedules/{scheduleId}][%d] putAnalyticsReportingScheduleInternalServerError  %+v", 500, o.Payload)
}

func (o *PutAnalyticsReportingScheduleInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutAnalyticsReportingScheduleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAnalyticsReportingScheduleServiceUnavailable creates a PutAnalyticsReportingScheduleServiceUnavailable with default headers values
func NewPutAnalyticsReportingScheduleServiceUnavailable() *PutAnalyticsReportingScheduleServiceUnavailable {
	return &PutAnalyticsReportingScheduleServiceUnavailable{}
}

/*PutAnalyticsReportingScheduleServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutAnalyticsReportingScheduleServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PutAnalyticsReportingScheduleServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/analytics/reporting/schedules/{scheduleId}][%d] putAnalyticsReportingScheduleServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutAnalyticsReportingScheduleServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutAnalyticsReportingScheduleServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAnalyticsReportingScheduleGatewayTimeout creates a PutAnalyticsReportingScheduleGatewayTimeout with default headers values
func NewPutAnalyticsReportingScheduleGatewayTimeout() *PutAnalyticsReportingScheduleGatewayTimeout {
	return &PutAnalyticsReportingScheduleGatewayTimeout{}
}

/*PutAnalyticsReportingScheduleGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PutAnalyticsReportingScheduleGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PutAnalyticsReportingScheduleGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/analytics/reporting/schedules/{scheduleId}][%d] putAnalyticsReportingScheduleGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutAnalyticsReportingScheduleGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutAnalyticsReportingScheduleGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

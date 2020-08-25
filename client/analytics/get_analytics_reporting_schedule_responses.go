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

/*GetAnalyticsReportingScheduleOK handles this case with default header values.

successful operation
*/
type GetAnalyticsReportingScheduleOK struct {
	Payload *models.ReportSchedule
}

func (o *GetAnalyticsReportingScheduleOK) Error() string {
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

/*GetAnalyticsReportingScheduleBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetAnalyticsReportingScheduleBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsReportingScheduleBadRequest) Error() string {
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

/*GetAnalyticsReportingScheduleUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetAnalyticsReportingScheduleUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsReportingScheduleUnauthorized) Error() string {
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

/*GetAnalyticsReportingScheduleForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetAnalyticsReportingScheduleForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsReportingScheduleForbidden) Error() string {
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

/*GetAnalyticsReportingScheduleNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetAnalyticsReportingScheduleNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsReportingScheduleNotFound) Error() string {
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

// NewGetAnalyticsReportingScheduleRequestEntityTooLarge creates a GetAnalyticsReportingScheduleRequestEntityTooLarge with default headers values
func NewGetAnalyticsReportingScheduleRequestEntityTooLarge() *GetAnalyticsReportingScheduleRequestEntityTooLarge {
	return &GetAnalyticsReportingScheduleRequestEntityTooLarge{}
}

/*GetAnalyticsReportingScheduleRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetAnalyticsReportingScheduleRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsReportingScheduleRequestEntityTooLarge) Error() string {
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

/*GetAnalyticsReportingScheduleUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetAnalyticsReportingScheduleUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsReportingScheduleUnsupportedMediaType) Error() string {
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

/*GetAnalyticsReportingScheduleTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetAnalyticsReportingScheduleTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsReportingScheduleTooManyRequests) Error() string {
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

/*GetAnalyticsReportingScheduleInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetAnalyticsReportingScheduleInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsReportingScheduleInternalServerError) Error() string {
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

/*GetAnalyticsReportingScheduleServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetAnalyticsReportingScheduleServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsReportingScheduleServiceUnavailable) Error() string {
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

/*GetAnalyticsReportingScheduleGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetAnalyticsReportingScheduleGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsReportingScheduleGatewayTimeout) Error() string {
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
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

// GetAnalyticsReportingScheduleHistoryLatestReader is a Reader for the GetAnalyticsReportingScheduleHistoryLatest structure.
type GetAnalyticsReportingScheduleHistoryLatestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAnalyticsReportingScheduleHistoryLatestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAnalyticsReportingScheduleHistoryLatestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAnalyticsReportingScheduleHistoryLatestBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAnalyticsReportingScheduleHistoryLatestUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAnalyticsReportingScheduleHistoryLatestForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAnalyticsReportingScheduleHistoryLatestNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetAnalyticsReportingScheduleHistoryLatestRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetAnalyticsReportingScheduleHistoryLatestUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetAnalyticsReportingScheduleHistoryLatestTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAnalyticsReportingScheduleHistoryLatestInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetAnalyticsReportingScheduleHistoryLatestServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetAnalyticsReportingScheduleHistoryLatestGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAnalyticsReportingScheduleHistoryLatestOK creates a GetAnalyticsReportingScheduleHistoryLatestOK with default headers values
func NewGetAnalyticsReportingScheduleHistoryLatestOK() *GetAnalyticsReportingScheduleHistoryLatestOK {
	return &GetAnalyticsReportingScheduleHistoryLatestOK{}
}

/*GetAnalyticsReportingScheduleHistoryLatestOK handles this case with default header values.

successful operation
*/
type GetAnalyticsReportingScheduleHistoryLatestOK struct {
	Payload *models.ReportRunEntry
}

func (o *GetAnalyticsReportingScheduleHistoryLatestOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/schedules/{scheduleId}/history/latest][%d] getAnalyticsReportingScheduleHistoryLatestOK  %+v", 200, o.Payload)
}

func (o *GetAnalyticsReportingScheduleHistoryLatestOK) GetPayload() *models.ReportRunEntry {
	return o.Payload
}

func (o *GetAnalyticsReportingScheduleHistoryLatestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ReportRunEntry)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingScheduleHistoryLatestBadRequest creates a GetAnalyticsReportingScheduleHistoryLatestBadRequest with default headers values
func NewGetAnalyticsReportingScheduleHistoryLatestBadRequest() *GetAnalyticsReportingScheduleHistoryLatestBadRequest {
	return &GetAnalyticsReportingScheduleHistoryLatestBadRequest{}
}

/*GetAnalyticsReportingScheduleHistoryLatestBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetAnalyticsReportingScheduleHistoryLatestBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsReportingScheduleHistoryLatestBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/schedules/{scheduleId}/history/latest][%d] getAnalyticsReportingScheduleHistoryLatestBadRequest  %+v", 400, o.Payload)
}

func (o *GetAnalyticsReportingScheduleHistoryLatestBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingScheduleHistoryLatestBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingScheduleHistoryLatestUnauthorized creates a GetAnalyticsReportingScheduleHistoryLatestUnauthorized with default headers values
func NewGetAnalyticsReportingScheduleHistoryLatestUnauthorized() *GetAnalyticsReportingScheduleHistoryLatestUnauthorized {
	return &GetAnalyticsReportingScheduleHistoryLatestUnauthorized{}
}

/*GetAnalyticsReportingScheduleHistoryLatestUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetAnalyticsReportingScheduleHistoryLatestUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsReportingScheduleHistoryLatestUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/schedules/{scheduleId}/history/latest][%d] getAnalyticsReportingScheduleHistoryLatestUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAnalyticsReportingScheduleHistoryLatestUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingScheduleHistoryLatestUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingScheduleHistoryLatestForbidden creates a GetAnalyticsReportingScheduleHistoryLatestForbidden with default headers values
func NewGetAnalyticsReportingScheduleHistoryLatestForbidden() *GetAnalyticsReportingScheduleHistoryLatestForbidden {
	return &GetAnalyticsReportingScheduleHistoryLatestForbidden{}
}

/*GetAnalyticsReportingScheduleHistoryLatestForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetAnalyticsReportingScheduleHistoryLatestForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsReportingScheduleHistoryLatestForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/schedules/{scheduleId}/history/latest][%d] getAnalyticsReportingScheduleHistoryLatestForbidden  %+v", 403, o.Payload)
}

func (o *GetAnalyticsReportingScheduleHistoryLatestForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingScheduleHistoryLatestForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingScheduleHistoryLatestNotFound creates a GetAnalyticsReportingScheduleHistoryLatestNotFound with default headers values
func NewGetAnalyticsReportingScheduleHistoryLatestNotFound() *GetAnalyticsReportingScheduleHistoryLatestNotFound {
	return &GetAnalyticsReportingScheduleHistoryLatestNotFound{}
}

/*GetAnalyticsReportingScheduleHistoryLatestNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetAnalyticsReportingScheduleHistoryLatestNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsReportingScheduleHistoryLatestNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/schedules/{scheduleId}/history/latest][%d] getAnalyticsReportingScheduleHistoryLatestNotFound  %+v", 404, o.Payload)
}

func (o *GetAnalyticsReportingScheduleHistoryLatestNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingScheduleHistoryLatestNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingScheduleHistoryLatestRequestEntityTooLarge creates a GetAnalyticsReportingScheduleHistoryLatestRequestEntityTooLarge with default headers values
func NewGetAnalyticsReportingScheduleHistoryLatestRequestEntityTooLarge() *GetAnalyticsReportingScheduleHistoryLatestRequestEntityTooLarge {
	return &GetAnalyticsReportingScheduleHistoryLatestRequestEntityTooLarge{}
}

/*GetAnalyticsReportingScheduleHistoryLatestRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetAnalyticsReportingScheduleHistoryLatestRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsReportingScheduleHistoryLatestRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/schedules/{scheduleId}/history/latest][%d] getAnalyticsReportingScheduleHistoryLatestRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetAnalyticsReportingScheduleHistoryLatestRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingScheduleHistoryLatestRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingScheduleHistoryLatestUnsupportedMediaType creates a GetAnalyticsReportingScheduleHistoryLatestUnsupportedMediaType with default headers values
func NewGetAnalyticsReportingScheduleHistoryLatestUnsupportedMediaType() *GetAnalyticsReportingScheduleHistoryLatestUnsupportedMediaType {
	return &GetAnalyticsReportingScheduleHistoryLatestUnsupportedMediaType{}
}

/*GetAnalyticsReportingScheduleHistoryLatestUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetAnalyticsReportingScheduleHistoryLatestUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsReportingScheduleHistoryLatestUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/schedules/{scheduleId}/history/latest][%d] getAnalyticsReportingScheduleHistoryLatestUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetAnalyticsReportingScheduleHistoryLatestUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingScheduleHistoryLatestUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingScheduleHistoryLatestTooManyRequests creates a GetAnalyticsReportingScheduleHistoryLatestTooManyRequests with default headers values
func NewGetAnalyticsReportingScheduleHistoryLatestTooManyRequests() *GetAnalyticsReportingScheduleHistoryLatestTooManyRequests {
	return &GetAnalyticsReportingScheduleHistoryLatestTooManyRequests{}
}

/*GetAnalyticsReportingScheduleHistoryLatestTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetAnalyticsReportingScheduleHistoryLatestTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsReportingScheduleHistoryLatestTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/schedules/{scheduleId}/history/latest][%d] getAnalyticsReportingScheduleHistoryLatestTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetAnalyticsReportingScheduleHistoryLatestTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingScheduleHistoryLatestTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingScheduleHistoryLatestInternalServerError creates a GetAnalyticsReportingScheduleHistoryLatestInternalServerError with default headers values
func NewGetAnalyticsReportingScheduleHistoryLatestInternalServerError() *GetAnalyticsReportingScheduleHistoryLatestInternalServerError {
	return &GetAnalyticsReportingScheduleHistoryLatestInternalServerError{}
}

/*GetAnalyticsReportingScheduleHistoryLatestInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetAnalyticsReportingScheduleHistoryLatestInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsReportingScheduleHistoryLatestInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/schedules/{scheduleId}/history/latest][%d] getAnalyticsReportingScheduleHistoryLatestInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAnalyticsReportingScheduleHistoryLatestInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingScheduleHistoryLatestInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingScheduleHistoryLatestServiceUnavailable creates a GetAnalyticsReportingScheduleHistoryLatestServiceUnavailable with default headers values
func NewGetAnalyticsReportingScheduleHistoryLatestServiceUnavailable() *GetAnalyticsReportingScheduleHistoryLatestServiceUnavailable {
	return &GetAnalyticsReportingScheduleHistoryLatestServiceUnavailable{}
}

/*GetAnalyticsReportingScheduleHistoryLatestServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetAnalyticsReportingScheduleHistoryLatestServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsReportingScheduleHistoryLatestServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/schedules/{scheduleId}/history/latest][%d] getAnalyticsReportingScheduleHistoryLatestServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetAnalyticsReportingScheduleHistoryLatestServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingScheduleHistoryLatestServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingScheduleHistoryLatestGatewayTimeout creates a GetAnalyticsReportingScheduleHistoryLatestGatewayTimeout with default headers values
func NewGetAnalyticsReportingScheduleHistoryLatestGatewayTimeout() *GetAnalyticsReportingScheduleHistoryLatestGatewayTimeout {
	return &GetAnalyticsReportingScheduleHistoryLatestGatewayTimeout{}
}

/*GetAnalyticsReportingScheduleHistoryLatestGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetAnalyticsReportingScheduleHistoryLatestGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsReportingScheduleHistoryLatestGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/schedules/{scheduleId}/history/latest][%d] getAnalyticsReportingScheduleHistoryLatestGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetAnalyticsReportingScheduleHistoryLatestGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingScheduleHistoryLatestGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
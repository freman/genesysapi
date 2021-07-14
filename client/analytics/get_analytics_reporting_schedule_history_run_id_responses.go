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

// GetAnalyticsReportingScheduleHistoryRunIDReader is a Reader for the GetAnalyticsReportingScheduleHistoryRunID structure.
type GetAnalyticsReportingScheduleHistoryRunIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAnalyticsReportingScheduleHistoryRunIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAnalyticsReportingScheduleHistoryRunIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAnalyticsReportingScheduleHistoryRunIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAnalyticsReportingScheduleHistoryRunIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAnalyticsReportingScheduleHistoryRunIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAnalyticsReportingScheduleHistoryRunIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetAnalyticsReportingScheduleHistoryRunIDRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetAnalyticsReportingScheduleHistoryRunIDRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetAnalyticsReportingScheduleHistoryRunIDUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetAnalyticsReportingScheduleHistoryRunIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAnalyticsReportingScheduleHistoryRunIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetAnalyticsReportingScheduleHistoryRunIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetAnalyticsReportingScheduleHistoryRunIDGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAnalyticsReportingScheduleHistoryRunIDOK creates a GetAnalyticsReportingScheduleHistoryRunIDOK with default headers values
func NewGetAnalyticsReportingScheduleHistoryRunIDOK() *GetAnalyticsReportingScheduleHistoryRunIDOK {
	return &GetAnalyticsReportingScheduleHistoryRunIDOK{}
}

/*GetAnalyticsReportingScheduleHistoryRunIDOK handles this case with default header values.

successful operation
*/
type GetAnalyticsReportingScheduleHistoryRunIDOK struct {
	Payload *models.ReportRunEntry
}

func (o *GetAnalyticsReportingScheduleHistoryRunIDOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/schedules/{scheduleId}/history/{runId}][%d] getAnalyticsReportingScheduleHistoryRunIdOK  %+v", 200, o.Payload)
}

func (o *GetAnalyticsReportingScheduleHistoryRunIDOK) GetPayload() *models.ReportRunEntry {
	return o.Payload
}

func (o *GetAnalyticsReportingScheduleHistoryRunIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ReportRunEntry)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingScheduleHistoryRunIDBadRequest creates a GetAnalyticsReportingScheduleHistoryRunIDBadRequest with default headers values
func NewGetAnalyticsReportingScheduleHistoryRunIDBadRequest() *GetAnalyticsReportingScheduleHistoryRunIDBadRequest {
	return &GetAnalyticsReportingScheduleHistoryRunIDBadRequest{}
}

/*GetAnalyticsReportingScheduleHistoryRunIDBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetAnalyticsReportingScheduleHistoryRunIDBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsReportingScheduleHistoryRunIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/schedules/{scheduleId}/history/{runId}][%d] getAnalyticsReportingScheduleHistoryRunIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetAnalyticsReportingScheduleHistoryRunIDBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingScheduleHistoryRunIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingScheduleHistoryRunIDUnauthorized creates a GetAnalyticsReportingScheduleHistoryRunIDUnauthorized with default headers values
func NewGetAnalyticsReportingScheduleHistoryRunIDUnauthorized() *GetAnalyticsReportingScheduleHistoryRunIDUnauthorized {
	return &GetAnalyticsReportingScheduleHistoryRunIDUnauthorized{}
}

/*GetAnalyticsReportingScheduleHistoryRunIDUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetAnalyticsReportingScheduleHistoryRunIDUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsReportingScheduleHistoryRunIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/schedules/{scheduleId}/history/{runId}][%d] getAnalyticsReportingScheduleHistoryRunIdUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAnalyticsReportingScheduleHistoryRunIDUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingScheduleHistoryRunIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingScheduleHistoryRunIDForbidden creates a GetAnalyticsReportingScheduleHistoryRunIDForbidden with default headers values
func NewGetAnalyticsReportingScheduleHistoryRunIDForbidden() *GetAnalyticsReportingScheduleHistoryRunIDForbidden {
	return &GetAnalyticsReportingScheduleHistoryRunIDForbidden{}
}

/*GetAnalyticsReportingScheduleHistoryRunIDForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetAnalyticsReportingScheduleHistoryRunIDForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsReportingScheduleHistoryRunIDForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/schedules/{scheduleId}/history/{runId}][%d] getAnalyticsReportingScheduleHistoryRunIdForbidden  %+v", 403, o.Payload)
}

func (o *GetAnalyticsReportingScheduleHistoryRunIDForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingScheduleHistoryRunIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingScheduleHistoryRunIDNotFound creates a GetAnalyticsReportingScheduleHistoryRunIDNotFound with default headers values
func NewGetAnalyticsReportingScheduleHistoryRunIDNotFound() *GetAnalyticsReportingScheduleHistoryRunIDNotFound {
	return &GetAnalyticsReportingScheduleHistoryRunIDNotFound{}
}

/*GetAnalyticsReportingScheduleHistoryRunIDNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetAnalyticsReportingScheduleHistoryRunIDNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsReportingScheduleHistoryRunIDNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/schedules/{scheduleId}/history/{runId}][%d] getAnalyticsReportingScheduleHistoryRunIdNotFound  %+v", 404, o.Payload)
}

func (o *GetAnalyticsReportingScheduleHistoryRunIDNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingScheduleHistoryRunIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingScheduleHistoryRunIDRequestTimeout creates a GetAnalyticsReportingScheduleHistoryRunIDRequestTimeout with default headers values
func NewGetAnalyticsReportingScheduleHistoryRunIDRequestTimeout() *GetAnalyticsReportingScheduleHistoryRunIDRequestTimeout {
	return &GetAnalyticsReportingScheduleHistoryRunIDRequestTimeout{}
}

/*GetAnalyticsReportingScheduleHistoryRunIDRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetAnalyticsReportingScheduleHistoryRunIDRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsReportingScheduleHistoryRunIDRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/schedules/{scheduleId}/history/{runId}][%d] getAnalyticsReportingScheduleHistoryRunIdRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetAnalyticsReportingScheduleHistoryRunIDRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingScheduleHistoryRunIDRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingScheduleHistoryRunIDRequestEntityTooLarge creates a GetAnalyticsReportingScheduleHistoryRunIDRequestEntityTooLarge with default headers values
func NewGetAnalyticsReportingScheduleHistoryRunIDRequestEntityTooLarge() *GetAnalyticsReportingScheduleHistoryRunIDRequestEntityTooLarge {
	return &GetAnalyticsReportingScheduleHistoryRunIDRequestEntityTooLarge{}
}

/*GetAnalyticsReportingScheduleHistoryRunIDRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetAnalyticsReportingScheduleHistoryRunIDRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsReportingScheduleHistoryRunIDRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/schedules/{scheduleId}/history/{runId}][%d] getAnalyticsReportingScheduleHistoryRunIdRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetAnalyticsReportingScheduleHistoryRunIDRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingScheduleHistoryRunIDRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingScheduleHistoryRunIDUnsupportedMediaType creates a GetAnalyticsReportingScheduleHistoryRunIDUnsupportedMediaType with default headers values
func NewGetAnalyticsReportingScheduleHistoryRunIDUnsupportedMediaType() *GetAnalyticsReportingScheduleHistoryRunIDUnsupportedMediaType {
	return &GetAnalyticsReportingScheduleHistoryRunIDUnsupportedMediaType{}
}

/*GetAnalyticsReportingScheduleHistoryRunIDUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetAnalyticsReportingScheduleHistoryRunIDUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsReportingScheduleHistoryRunIDUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/schedules/{scheduleId}/history/{runId}][%d] getAnalyticsReportingScheduleHistoryRunIdUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetAnalyticsReportingScheduleHistoryRunIDUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingScheduleHistoryRunIDUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingScheduleHistoryRunIDTooManyRequests creates a GetAnalyticsReportingScheduleHistoryRunIDTooManyRequests with default headers values
func NewGetAnalyticsReportingScheduleHistoryRunIDTooManyRequests() *GetAnalyticsReportingScheduleHistoryRunIDTooManyRequests {
	return &GetAnalyticsReportingScheduleHistoryRunIDTooManyRequests{}
}

/*GetAnalyticsReportingScheduleHistoryRunIDTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetAnalyticsReportingScheduleHistoryRunIDTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsReportingScheduleHistoryRunIDTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/schedules/{scheduleId}/history/{runId}][%d] getAnalyticsReportingScheduleHistoryRunIdTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetAnalyticsReportingScheduleHistoryRunIDTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingScheduleHistoryRunIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingScheduleHistoryRunIDInternalServerError creates a GetAnalyticsReportingScheduleHistoryRunIDInternalServerError with default headers values
func NewGetAnalyticsReportingScheduleHistoryRunIDInternalServerError() *GetAnalyticsReportingScheduleHistoryRunIDInternalServerError {
	return &GetAnalyticsReportingScheduleHistoryRunIDInternalServerError{}
}

/*GetAnalyticsReportingScheduleHistoryRunIDInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetAnalyticsReportingScheduleHistoryRunIDInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsReportingScheduleHistoryRunIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/schedules/{scheduleId}/history/{runId}][%d] getAnalyticsReportingScheduleHistoryRunIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAnalyticsReportingScheduleHistoryRunIDInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingScheduleHistoryRunIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingScheduleHistoryRunIDServiceUnavailable creates a GetAnalyticsReportingScheduleHistoryRunIDServiceUnavailable with default headers values
func NewGetAnalyticsReportingScheduleHistoryRunIDServiceUnavailable() *GetAnalyticsReportingScheduleHistoryRunIDServiceUnavailable {
	return &GetAnalyticsReportingScheduleHistoryRunIDServiceUnavailable{}
}

/*GetAnalyticsReportingScheduleHistoryRunIDServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetAnalyticsReportingScheduleHistoryRunIDServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsReportingScheduleHistoryRunIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/schedules/{scheduleId}/history/{runId}][%d] getAnalyticsReportingScheduleHistoryRunIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetAnalyticsReportingScheduleHistoryRunIDServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingScheduleHistoryRunIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingScheduleHistoryRunIDGatewayTimeout creates a GetAnalyticsReportingScheduleHistoryRunIDGatewayTimeout with default headers values
func NewGetAnalyticsReportingScheduleHistoryRunIDGatewayTimeout() *GetAnalyticsReportingScheduleHistoryRunIDGatewayTimeout {
	return &GetAnalyticsReportingScheduleHistoryRunIDGatewayTimeout{}
}

/*GetAnalyticsReportingScheduleHistoryRunIDGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetAnalyticsReportingScheduleHistoryRunIDGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsReportingScheduleHistoryRunIDGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/schedules/{scheduleId}/history/{runId}][%d] getAnalyticsReportingScheduleHistoryRunIdGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetAnalyticsReportingScheduleHistoryRunIDGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingScheduleHistoryRunIDGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

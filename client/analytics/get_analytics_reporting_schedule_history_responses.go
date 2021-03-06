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

// GetAnalyticsReportingScheduleHistoryReader is a Reader for the GetAnalyticsReportingScheduleHistory structure.
type GetAnalyticsReportingScheduleHistoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAnalyticsReportingScheduleHistoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAnalyticsReportingScheduleHistoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAnalyticsReportingScheduleHistoryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAnalyticsReportingScheduleHistoryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAnalyticsReportingScheduleHistoryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAnalyticsReportingScheduleHistoryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetAnalyticsReportingScheduleHistoryRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetAnalyticsReportingScheduleHistoryUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetAnalyticsReportingScheduleHistoryTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAnalyticsReportingScheduleHistoryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetAnalyticsReportingScheduleHistoryServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetAnalyticsReportingScheduleHistoryGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAnalyticsReportingScheduleHistoryOK creates a GetAnalyticsReportingScheduleHistoryOK with default headers values
func NewGetAnalyticsReportingScheduleHistoryOK() *GetAnalyticsReportingScheduleHistoryOK {
	return &GetAnalyticsReportingScheduleHistoryOK{}
}

/*GetAnalyticsReportingScheduleHistoryOK handles this case with default header values.

successful operation
*/
type GetAnalyticsReportingScheduleHistoryOK struct {
	Payload *models.ReportRunEntryEntityDomainListing
}

func (o *GetAnalyticsReportingScheduleHistoryOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/schedules/{scheduleId}/history][%d] getAnalyticsReportingScheduleHistoryOK  %+v", 200, o.Payload)
}

func (o *GetAnalyticsReportingScheduleHistoryOK) GetPayload() *models.ReportRunEntryEntityDomainListing {
	return o.Payload
}

func (o *GetAnalyticsReportingScheduleHistoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ReportRunEntryEntityDomainListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingScheduleHistoryBadRequest creates a GetAnalyticsReportingScheduleHistoryBadRequest with default headers values
func NewGetAnalyticsReportingScheduleHistoryBadRequest() *GetAnalyticsReportingScheduleHistoryBadRequest {
	return &GetAnalyticsReportingScheduleHistoryBadRequest{}
}

/*GetAnalyticsReportingScheduleHistoryBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetAnalyticsReportingScheduleHistoryBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsReportingScheduleHistoryBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/schedules/{scheduleId}/history][%d] getAnalyticsReportingScheduleHistoryBadRequest  %+v", 400, o.Payload)
}

func (o *GetAnalyticsReportingScheduleHistoryBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingScheduleHistoryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingScheduleHistoryUnauthorized creates a GetAnalyticsReportingScheduleHistoryUnauthorized with default headers values
func NewGetAnalyticsReportingScheduleHistoryUnauthorized() *GetAnalyticsReportingScheduleHistoryUnauthorized {
	return &GetAnalyticsReportingScheduleHistoryUnauthorized{}
}

/*GetAnalyticsReportingScheduleHistoryUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetAnalyticsReportingScheduleHistoryUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsReportingScheduleHistoryUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/schedules/{scheduleId}/history][%d] getAnalyticsReportingScheduleHistoryUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAnalyticsReportingScheduleHistoryUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingScheduleHistoryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingScheduleHistoryForbidden creates a GetAnalyticsReportingScheduleHistoryForbidden with default headers values
func NewGetAnalyticsReportingScheduleHistoryForbidden() *GetAnalyticsReportingScheduleHistoryForbidden {
	return &GetAnalyticsReportingScheduleHistoryForbidden{}
}

/*GetAnalyticsReportingScheduleHistoryForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetAnalyticsReportingScheduleHistoryForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsReportingScheduleHistoryForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/schedules/{scheduleId}/history][%d] getAnalyticsReportingScheduleHistoryForbidden  %+v", 403, o.Payload)
}

func (o *GetAnalyticsReportingScheduleHistoryForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingScheduleHistoryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingScheduleHistoryNotFound creates a GetAnalyticsReportingScheduleHistoryNotFound with default headers values
func NewGetAnalyticsReportingScheduleHistoryNotFound() *GetAnalyticsReportingScheduleHistoryNotFound {
	return &GetAnalyticsReportingScheduleHistoryNotFound{}
}

/*GetAnalyticsReportingScheduleHistoryNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetAnalyticsReportingScheduleHistoryNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsReportingScheduleHistoryNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/schedules/{scheduleId}/history][%d] getAnalyticsReportingScheduleHistoryNotFound  %+v", 404, o.Payload)
}

func (o *GetAnalyticsReportingScheduleHistoryNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingScheduleHistoryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingScheduleHistoryRequestEntityTooLarge creates a GetAnalyticsReportingScheduleHistoryRequestEntityTooLarge with default headers values
func NewGetAnalyticsReportingScheduleHistoryRequestEntityTooLarge() *GetAnalyticsReportingScheduleHistoryRequestEntityTooLarge {
	return &GetAnalyticsReportingScheduleHistoryRequestEntityTooLarge{}
}

/*GetAnalyticsReportingScheduleHistoryRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetAnalyticsReportingScheduleHistoryRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsReportingScheduleHistoryRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/schedules/{scheduleId}/history][%d] getAnalyticsReportingScheduleHistoryRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetAnalyticsReportingScheduleHistoryRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingScheduleHistoryRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingScheduleHistoryUnsupportedMediaType creates a GetAnalyticsReportingScheduleHistoryUnsupportedMediaType with default headers values
func NewGetAnalyticsReportingScheduleHistoryUnsupportedMediaType() *GetAnalyticsReportingScheduleHistoryUnsupportedMediaType {
	return &GetAnalyticsReportingScheduleHistoryUnsupportedMediaType{}
}

/*GetAnalyticsReportingScheduleHistoryUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetAnalyticsReportingScheduleHistoryUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsReportingScheduleHistoryUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/schedules/{scheduleId}/history][%d] getAnalyticsReportingScheduleHistoryUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetAnalyticsReportingScheduleHistoryUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingScheduleHistoryUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingScheduleHistoryTooManyRequests creates a GetAnalyticsReportingScheduleHistoryTooManyRequests with default headers values
func NewGetAnalyticsReportingScheduleHistoryTooManyRequests() *GetAnalyticsReportingScheduleHistoryTooManyRequests {
	return &GetAnalyticsReportingScheduleHistoryTooManyRequests{}
}

/*GetAnalyticsReportingScheduleHistoryTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetAnalyticsReportingScheduleHistoryTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsReportingScheduleHistoryTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/schedules/{scheduleId}/history][%d] getAnalyticsReportingScheduleHistoryTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetAnalyticsReportingScheduleHistoryTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingScheduleHistoryTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingScheduleHistoryInternalServerError creates a GetAnalyticsReportingScheduleHistoryInternalServerError with default headers values
func NewGetAnalyticsReportingScheduleHistoryInternalServerError() *GetAnalyticsReportingScheduleHistoryInternalServerError {
	return &GetAnalyticsReportingScheduleHistoryInternalServerError{}
}

/*GetAnalyticsReportingScheduleHistoryInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetAnalyticsReportingScheduleHistoryInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsReportingScheduleHistoryInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/schedules/{scheduleId}/history][%d] getAnalyticsReportingScheduleHistoryInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAnalyticsReportingScheduleHistoryInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingScheduleHistoryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingScheduleHistoryServiceUnavailable creates a GetAnalyticsReportingScheduleHistoryServiceUnavailable with default headers values
func NewGetAnalyticsReportingScheduleHistoryServiceUnavailable() *GetAnalyticsReportingScheduleHistoryServiceUnavailable {
	return &GetAnalyticsReportingScheduleHistoryServiceUnavailable{}
}

/*GetAnalyticsReportingScheduleHistoryServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetAnalyticsReportingScheduleHistoryServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsReportingScheduleHistoryServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/schedules/{scheduleId}/history][%d] getAnalyticsReportingScheduleHistoryServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetAnalyticsReportingScheduleHistoryServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingScheduleHistoryServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingScheduleHistoryGatewayTimeout creates a GetAnalyticsReportingScheduleHistoryGatewayTimeout with default headers values
func NewGetAnalyticsReportingScheduleHistoryGatewayTimeout() *GetAnalyticsReportingScheduleHistoryGatewayTimeout {
	return &GetAnalyticsReportingScheduleHistoryGatewayTimeout{}
}

/*GetAnalyticsReportingScheduleHistoryGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetAnalyticsReportingScheduleHistoryGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsReportingScheduleHistoryGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/schedules/{scheduleId}/history][%d] getAnalyticsReportingScheduleHistoryGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetAnalyticsReportingScheduleHistoryGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingScheduleHistoryGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

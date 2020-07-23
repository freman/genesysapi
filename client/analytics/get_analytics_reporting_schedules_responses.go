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

// GetAnalyticsReportingSchedulesReader is a Reader for the GetAnalyticsReportingSchedules structure.
type GetAnalyticsReportingSchedulesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAnalyticsReportingSchedulesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAnalyticsReportingSchedulesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAnalyticsReportingSchedulesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAnalyticsReportingSchedulesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAnalyticsReportingSchedulesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAnalyticsReportingSchedulesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetAnalyticsReportingSchedulesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetAnalyticsReportingSchedulesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetAnalyticsReportingSchedulesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAnalyticsReportingSchedulesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetAnalyticsReportingSchedulesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetAnalyticsReportingSchedulesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAnalyticsReportingSchedulesOK creates a GetAnalyticsReportingSchedulesOK with default headers values
func NewGetAnalyticsReportingSchedulesOK() *GetAnalyticsReportingSchedulesOK {
	return &GetAnalyticsReportingSchedulesOK{}
}

/*GetAnalyticsReportingSchedulesOK handles this case with default header values.

successful operation
*/
type GetAnalyticsReportingSchedulesOK struct {
	Payload *models.ReportScheduleEntityListing
}

func (o *GetAnalyticsReportingSchedulesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/schedules][%d] getAnalyticsReportingSchedulesOK  %+v", 200, o.Payload)
}

func (o *GetAnalyticsReportingSchedulesOK) GetPayload() *models.ReportScheduleEntityListing {
	return o.Payload
}

func (o *GetAnalyticsReportingSchedulesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ReportScheduleEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingSchedulesBadRequest creates a GetAnalyticsReportingSchedulesBadRequest with default headers values
func NewGetAnalyticsReportingSchedulesBadRequest() *GetAnalyticsReportingSchedulesBadRequest {
	return &GetAnalyticsReportingSchedulesBadRequest{}
}

/*GetAnalyticsReportingSchedulesBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetAnalyticsReportingSchedulesBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsReportingSchedulesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/schedules][%d] getAnalyticsReportingSchedulesBadRequest  %+v", 400, o.Payload)
}

func (o *GetAnalyticsReportingSchedulesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingSchedulesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingSchedulesUnauthorized creates a GetAnalyticsReportingSchedulesUnauthorized with default headers values
func NewGetAnalyticsReportingSchedulesUnauthorized() *GetAnalyticsReportingSchedulesUnauthorized {
	return &GetAnalyticsReportingSchedulesUnauthorized{}
}

/*GetAnalyticsReportingSchedulesUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetAnalyticsReportingSchedulesUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsReportingSchedulesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/schedules][%d] getAnalyticsReportingSchedulesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAnalyticsReportingSchedulesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingSchedulesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingSchedulesForbidden creates a GetAnalyticsReportingSchedulesForbidden with default headers values
func NewGetAnalyticsReportingSchedulesForbidden() *GetAnalyticsReportingSchedulesForbidden {
	return &GetAnalyticsReportingSchedulesForbidden{}
}

/*GetAnalyticsReportingSchedulesForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetAnalyticsReportingSchedulesForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsReportingSchedulesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/schedules][%d] getAnalyticsReportingSchedulesForbidden  %+v", 403, o.Payload)
}

func (o *GetAnalyticsReportingSchedulesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingSchedulesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingSchedulesNotFound creates a GetAnalyticsReportingSchedulesNotFound with default headers values
func NewGetAnalyticsReportingSchedulesNotFound() *GetAnalyticsReportingSchedulesNotFound {
	return &GetAnalyticsReportingSchedulesNotFound{}
}

/*GetAnalyticsReportingSchedulesNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetAnalyticsReportingSchedulesNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsReportingSchedulesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/schedules][%d] getAnalyticsReportingSchedulesNotFound  %+v", 404, o.Payload)
}

func (o *GetAnalyticsReportingSchedulesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingSchedulesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingSchedulesRequestEntityTooLarge creates a GetAnalyticsReportingSchedulesRequestEntityTooLarge with default headers values
func NewGetAnalyticsReportingSchedulesRequestEntityTooLarge() *GetAnalyticsReportingSchedulesRequestEntityTooLarge {
	return &GetAnalyticsReportingSchedulesRequestEntityTooLarge{}
}

/*GetAnalyticsReportingSchedulesRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetAnalyticsReportingSchedulesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsReportingSchedulesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/schedules][%d] getAnalyticsReportingSchedulesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetAnalyticsReportingSchedulesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingSchedulesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingSchedulesUnsupportedMediaType creates a GetAnalyticsReportingSchedulesUnsupportedMediaType with default headers values
func NewGetAnalyticsReportingSchedulesUnsupportedMediaType() *GetAnalyticsReportingSchedulesUnsupportedMediaType {
	return &GetAnalyticsReportingSchedulesUnsupportedMediaType{}
}

/*GetAnalyticsReportingSchedulesUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetAnalyticsReportingSchedulesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsReportingSchedulesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/schedules][%d] getAnalyticsReportingSchedulesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetAnalyticsReportingSchedulesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingSchedulesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingSchedulesTooManyRequests creates a GetAnalyticsReportingSchedulesTooManyRequests with default headers values
func NewGetAnalyticsReportingSchedulesTooManyRequests() *GetAnalyticsReportingSchedulesTooManyRequests {
	return &GetAnalyticsReportingSchedulesTooManyRequests{}
}

/*GetAnalyticsReportingSchedulesTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetAnalyticsReportingSchedulesTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsReportingSchedulesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/schedules][%d] getAnalyticsReportingSchedulesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetAnalyticsReportingSchedulesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingSchedulesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingSchedulesInternalServerError creates a GetAnalyticsReportingSchedulesInternalServerError with default headers values
func NewGetAnalyticsReportingSchedulesInternalServerError() *GetAnalyticsReportingSchedulesInternalServerError {
	return &GetAnalyticsReportingSchedulesInternalServerError{}
}

/*GetAnalyticsReportingSchedulesInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetAnalyticsReportingSchedulesInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsReportingSchedulesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/schedules][%d] getAnalyticsReportingSchedulesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAnalyticsReportingSchedulesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingSchedulesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingSchedulesServiceUnavailable creates a GetAnalyticsReportingSchedulesServiceUnavailable with default headers values
func NewGetAnalyticsReportingSchedulesServiceUnavailable() *GetAnalyticsReportingSchedulesServiceUnavailable {
	return &GetAnalyticsReportingSchedulesServiceUnavailable{}
}

/*GetAnalyticsReportingSchedulesServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetAnalyticsReportingSchedulesServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsReportingSchedulesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/schedules][%d] getAnalyticsReportingSchedulesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetAnalyticsReportingSchedulesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingSchedulesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingSchedulesGatewayTimeout creates a GetAnalyticsReportingSchedulesGatewayTimeout with default headers values
func NewGetAnalyticsReportingSchedulesGatewayTimeout() *GetAnalyticsReportingSchedulesGatewayTimeout {
	return &GetAnalyticsReportingSchedulesGatewayTimeout{}
}

/*GetAnalyticsReportingSchedulesGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetAnalyticsReportingSchedulesGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsReportingSchedulesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/schedules][%d] getAnalyticsReportingSchedulesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetAnalyticsReportingSchedulesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingSchedulesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package workforce_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetWorkforcemanagementManagementunitWeekScheduleReader is a Reader for the GetWorkforcemanagementManagementunitWeekSchedule structure.
type GetWorkforcemanagementManagementunitWeekScheduleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWorkforcemanagementManagementunitWeekScheduleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWorkforcemanagementManagementunitWeekScheduleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetWorkforcemanagementManagementunitWeekScheduleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetWorkforcemanagementManagementunitWeekScheduleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetWorkforcemanagementManagementunitWeekScheduleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetWorkforcemanagementManagementunitWeekScheduleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetWorkforcemanagementManagementunitWeekScheduleRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetWorkforcemanagementManagementunitWeekScheduleRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetWorkforcemanagementManagementunitWeekScheduleUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetWorkforcemanagementManagementunitWeekScheduleTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetWorkforcemanagementManagementunitWeekScheduleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetWorkforcemanagementManagementunitWeekScheduleServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetWorkforcemanagementManagementunitWeekScheduleGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetWorkforcemanagementManagementunitWeekScheduleOK creates a GetWorkforcemanagementManagementunitWeekScheduleOK with default headers values
func NewGetWorkforcemanagementManagementunitWeekScheduleOK() *GetWorkforcemanagementManagementunitWeekScheduleOK {
	return &GetWorkforcemanagementManagementunitWeekScheduleOK{}
}

/*GetWorkforcemanagementManagementunitWeekScheduleOK handles this case with default header values.

successful operation
*/
type GetWorkforcemanagementManagementunitWeekScheduleOK struct {
	Payload *models.WeekScheduleResponse
}

func (o *GetWorkforcemanagementManagementunitWeekScheduleOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekId}/schedules/{scheduleId}][%d] getWorkforcemanagementManagementunitWeekScheduleOK  %+v", 200, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitWeekScheduleOK) GetPayload() *models.WeekScheduleResponse {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitWeekScheduleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WeekScheduleResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitWeekScheduleBadRequest creates a GetWorkforcemanagementManagementunitWeekScheduleBadRequest with default headers values
func NewGetWorkforcemanagementManagementunitWeekScheduleBadRequest() *GetWorkforcemanagementManagementunitWeekScheduleBadRequest {
	return &GetWorkforcemanagementManagementunitWeekScheduleBadRequest{}
}

/*GetWorkforcemanagementManagementunitWeekScheduleBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetWorkforcemanagementManagementunitWeekScheduleBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitWeekScheduleBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekId}/schedules/{scheduleId}][%d] getWorkforcemanagementManagementunitWeekScheduleBadRequest  %+v", 400, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitWeekScheduleBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitWeekScheduleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitWeekScheduleUnauthorized creates a GetWorkforcemanagementManagementunitWeekScheduleUnauthorized with default headers values
func NewGetWorkforcemanagementManagementunitWeekScheduleUnauthorized() *GetWorkforcemanagementManagementunitWeekScheduleUnauthorized {
	return &GetWorkforcemanagementManagementunitWeekScheduleUnauthorized{}
}

/*GetWorkforcemanagementManagementunitWeekScheduleUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetWorkforcemanagementManagementunitWeekScheduleUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitWeekScheduleUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekId}/schedules/{scheduleId}][%d] getWorkforcemanagementManagementunitWeekScheduleUnauthorized  %+v", 401, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitWeekScheduleUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitWeekScheduleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitWeekScheduleForbidden creates a GetWorkforcemanagementManagementunitWeekScheduleForbidden with default headers values
func NewGetWorkforcemanagementManagementunitWeekScheduleForbidden() *GetWorkforcemanagementManagementunitWeekScheduleForbidden {
	return &GetWorkforcemanagementManagementunitWeekScheduleForbidden{}
}

/*GetWorkforcemanagementManagementunitWeekScheduleForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetWorkforcemanagementManagementunitWeekScheduleForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitWeekScheduleForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekId}/schedules/{scheduleId}][%d] getWorkforcemanagementManagementunitWeekScheduleForbidden  %+v", 403, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitWeekScheduleForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitWeekScheduleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitWeekScheduleNotFound creates a GetWorkforcemanagementManagementunitWeekScheduleNotFound with default headers values
func NewGetWorkforcemanagementManagementunitWeekScheduleNotFound() *GetWorkforcemanagementManagementunitWeekScheduleNotFound {
	return &GetWorkforcemanagementManagementunitWeekScheduleNotFound{}
}

/*GetWorkforcemanagementManagementunitWeekScheduleNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetWorkforcemanagementManagementunitWeekScheduleNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitWeekScheduleNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekId}/schedules/{scheduleId}][%d] getWorkforcemanagementManagementunitWeekScheduleNotFound  %+v", 404, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitWeekScheduleNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitWeekScheduleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitWeekScheduleRequestTimeout creates a GetWorkforcemanagementManagementunitWeekScheduleRequestTimeout with default headers values
func NewGetWorkforcemanagementManagementunitWeekScheduleRequestTimeout() *GetWorkforcemanagementManagementunitWeekScheduleRequestTimeout {
	return &GetWorkforcemanagementManagementunitWeekScheduleRequestTimeout{}
}

/*GetWorkforcemanagementManagementunitWeekScheduleRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetWorkforcemanagementManagementunitWeekScheduleRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitWeekScheduleRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekId}/schedules/{scheduleId}][%d] getWorkforcemanagementManagementunitWeekScheduleRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitWeekScheduleRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitWeekScheduleRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitWeekScheduleRequestEntityTooLarge creates a GetWorkforcemanagementManagementunitWeekScheduleRequestEntityTooLarge with default headers values
func NewGetWorkforcemanagementManagementunitWeekScheduleRequestEntityTooLarge() *GetWorkforcemanagementManagementunitWeekScheduleRequestEntityTooLarge {
	return &GetWorkforcemanagementManagementunitWeekScheduleRequestEntityTooLarge{}
}

/*GetWorkforcemanagementManagementunitWeekScheduleRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetWorkforcemanagementManagementunitWeekScheduleRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitWeekScheduleRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekId}/schedules/{scheduleId}][%d] getWorkforcemanagementManagementunitWeekScheduleRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitWeekScheduleRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitWeekScheduleRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitWeekScheduleUnsupportedMediaType creates a GetWorkforcemanagementManagementunitWeekScheduleUnsupportedMediaType with default headers values
func NewGetWorkforcemanagementManagementunitWeekScheduleUnsupportedMediaType() *GetWorkforcemanagementManagementunitWeekScheduleUnsupportedMediaType {
	return &GetWorkforcemanagementManagementunitWeekScheduleUnsupportedMediaType{}
}

/*GetWorkforcemanagementManagementunitWeekScheduleUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetWorkforcemanagementManagementunitWeekScheduleUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitWeekScheduleUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekId}/schedules/{scheduleId}][%d] getWorkforcemanagementManagementunitWeekScheduleUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitWeekScheduleUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitWeekScheduleUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitWeekScheduleTooManyRequests creates a GetWorkforcemanagementManagementunitWeekScheduleTooManyRequests with default headers values
func NewGetWorkforcemanagementManagementunitWeekScheduleTooManyRequests() *GetWorkforcemanagementManagementunitWeekScheduleTooManyRequests {
	return &GetWorkforcemanagementManagementunitWeekScheduleTooManyRequests{}
}

/*GetWorkforcemanagementManagementunitWeekScheduleTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetWorkforcemanagementManagementunitWeekScheduleTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitWeekScheduleTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekId}/schedules/{scheduleId}][%d] getWorkforcemanagementManagementunitWeekScheduleTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitWeekScheduleTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitWeekScheduleTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitWeekScheduleInternalServerError creates a GetWorkforcemanagementManagementunitWeekScheduleInternalServerError with default headers values
func NewGetWorkforcemanagementManagementunitWeekScheduleInternalServerError() *GetWorkforcemanagementManagementunitWeekScheduleInternalServerError {
	return &GetWorkforcemanagementManagementunitWeekScheduleInternalServerError{}
}

/*GetWorkforcemanagementManagementunitWeekScheduleInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetWorkforcemanagementManagementunitWeekScheduleInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitWeekScheduleInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekId}/schedules/{scheduleId}][%d] getWorkforcemanagementManagementunitWeekScheduleInternalServerError  %+v", 500, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitWeekScheduleInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitWeekScheduleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitWeekScheduleServiceUnavailable creates a GetWorkforcemanagementManagementunitWeekScheduleServiceUnavailable with default headers values
func NewGetWorkforcemanagementManagementunitWeekScheduleServiceUnavailable() *GetWorkforcemanagementManagementunitWeekScheduleServiceUnavailable {
	return &GetWorkforcemanagementManagementunitWeekScheduleServiceUnavailable{}
}

/*GetWorkforcemanagementManagementunitWeekScheduleServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetWorkforcemanagementManagementunitWeekScheduleServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitWeekScheduleServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekId}/schedules/{scheduleId}][%d] getWorkforcemanagementManagementunitWeekScheduleServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitWeekScheduleServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitWeekScheduleServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitWeekScheduleGatewayTimeout creates a GetWorkforcemanagementManagementunitWeekScheduleGatewayTimeout with default headers values
func NewGetWorkforcemanagementManagementunitWeekScheduleGatewayTimeout() *GetWorkforcemanagementManagementunitWeekScheduleGatewayTimeout {
	return &GetWorkforcemanagementManagementunitWeekScheduleGatewayTimeout{}
}

/*GetWorkforcemanagementManagementunitWeekScheduleGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetWorkforcemanagementManagementunitWeekScheduleGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitWeekScheduleGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekId}/schedules/{scheduleId}][%d] getWorkforcemanagementManagementunitWeekScheduleGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitWeekScheduleGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitWeekScheduleGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

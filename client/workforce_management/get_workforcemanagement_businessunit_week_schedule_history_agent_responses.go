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

// GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentReader is a Reader for the GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgent structure.
type GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentOK creates a GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentOK with default headers values
func NewGetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentOK() *GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentOK {
	return &GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentOK{}
}

/*GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentOK handles this case with default header values.

successful operation
*/
type GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentOK struct {
	Payload *models.BuAgentScheduleHistoryResponse
}

func (o *GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/history/agents/{agentId}][%d] getWorkforcemanagementBusinessunitWeekScheduleHistoryAgentOK  %+v", 200, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentOK) GetPayload() *models.BuAgentScheduleHistoryResponse {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BuAgentScheduleHistoryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentBadRequest creates a GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentBadRequest with default headers values
func NewGetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentBadRequest() *GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentBadRequest {
	return &GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentBadRequest{}
}

/*GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/history/agents/{agentId}][%d] getWorkforcemanagementBusinessunitWeekScheduleHistoryAgentBadRequest  %+v", 400, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentUnauthorized creates a GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentUnauthorized with default headers values
func NewGetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentUnauthorized() *GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentUnauthorized {
	return &GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentUnauthorized{}
}

/*GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/history/agents/{agentId}][%d] getWorkforcemanagementBusinessunitWeekScheduleHistoryAgentUnauthorized  %+v", 401, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentForbidden creates a GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentForbidden with default headers values
func NewGetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentForbidden() *GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentForbidden {
	return &GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentForbidden{}
}

/*GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/history/agents/{agentId}][%d] getWorkforcemanagementBusinessunitWeekScheduleHistoryAgentForbidden  %+v", 403, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentNotFound creates a GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentNotFound with default headers values
func NewGetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentNotFound() *GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentNotFound {
	return &GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentNotFound{}
}

/*GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/history/agents/{agentId}][%d] getWorkforcemanagementBusinessunitWeekScheduleHistoryAgentNotFound  %+v", 404, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentRequestEntityTooLarge creates a GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentRequestEntityTooLarge with default headers values
func NewGetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentRequestEntityTooLarge() *GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentRequestEntityTooLarge {
	return &GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentRequestEntityTooLarge{}
}

/*GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/history/agents/{agentId}][%d] getWorkforcemanagementBusinessunitWeekScheduleHistoryAgentRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentUnsupportedMediaType creates a GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentUnsupportedMediaType with default headers values
func NewGetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentUnsupportedMediaType() *GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentUnsupportedMediaType {
	return &GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentUnsupportedMediaType{}
}

/*GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/history/agents/{agentId}][%d] getWorkforcemanagementBusinessunitWeekScheduleHistoryAgentUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentTooManyRequests creates a GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentTooManyRequests with default headers values
func NewGetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentTooManyRequests() *GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentTooManyRequests {
	return &GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentTooManyRequests{}
}

/*GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/history/agents/{agentId}][%d] getWorkforcemanagementBusinessunitWeekScheduleHistoryAgentTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentInternalServerError creates a GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentInternalServerError with default headers values
func NewGetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentInternalServerError() *GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentInternalServerError {
	return &GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentInternalServerError{}
}

/*GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/history/agents/{agentId}][%d] getWorkforcemanagementBusinessunitWeekScheduleHistoryAgentInternalServerError  %+v", 500, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentServiceUnavailable creates a GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentServiceUnavailable with default headers values
func NewGetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentServiceUnavailable() *GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentServiceUnavailable {
	return &GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentServiceUnavailable{}
}

/*GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/history/agents/{agentId}][%d] getWorkforcemanagementBusinessunitWeekScheduleHistoryAgentServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentGatewayTimeout creates a GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentGatewayTimeout with default headers values
func NewGetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentGatewayTimeout() *GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentGatewayTimeout {
	return &GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentGatewayTimeout{}
}

/*GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/history/agents/{agentId}][%d] getWorkforcemanagementBusinessunitWeekScheduleHistoryAgentGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgentGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
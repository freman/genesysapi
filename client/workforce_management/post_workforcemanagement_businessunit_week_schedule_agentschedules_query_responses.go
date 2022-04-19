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

// PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryReader is a Reader for the PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQuery structure.
type PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryOK creates a PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryOK with default headers values
func NewPostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryOK() *PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryOK {
	return &PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryOK{}
}

/*PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryOK handles this case with default header values.

successful operation
*/
type PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryOK struct {
	Payload *models.BuAsyncAgentSchedulesQueryResponse
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/agentschedules/query][%d] postWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryOK  %+v", 200, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryOK) GetPayload() *models.BuAsyncAgentSchedulesQueryResponse {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BuAsyncAgentSchedulesQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryBadRequest creates a PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryBadRequest with default headers values
func NewPostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryBadRequest() *PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryBadRequest {
	return &PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryBadRequest{}
}

/*PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/agentschedules/query][%d] postWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryBadRequest  %+v", 400, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryUnauthorized creates a PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryUnauthorized with default headers values
func NewPostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryUnauthorized() *PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryUnauthorized {
	return &PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryUnauthorized{}
}

/*PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/agentschedules/query][%d] postWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryUnauthorized  %+v", 401, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryForbidden creates a PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryForbidden with default headers values
func NewPostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryForbidden() *PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryForbidden {
	return &PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryForbidden{}
}

/*PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/agentschedules/query][%d] postWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryForbidden  %+v", 403, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryNotFound creates a PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryNotFound with default headers values
func NewPostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryNotFound() *PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryNotFound {
	return &PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryNotFound{}
}

/*PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/agentschedules/query][%d] postWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryNotFound  %+v", 404, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryRequestTimeout creates a PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryRequestTimeout with default headers values
func NewPostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryRequestTimeout() *PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryRequestTimeout {
	return &PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryRequestTimeout{}
}

/*PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/agentschedules/query][%d] postWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryRequestEntityTooLarge creates a PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryRequestEntityTooLarge with default headers values
func NewPostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryRequestEntityTooLarge() *PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryRequestEntityTooLarge {
	return &PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryRequestEntityTooLarge{}
}

/*PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/agentschedules/query][%d] postWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryUnsupportedMediaType creates a PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryUnsupportedMediaType with default headers values
func NewPostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryUnsupportedMediaType() *PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryUnsupportedMediaType {
	return &PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryUnsupportedMediaType{}
}

/*PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/agentschedules/query][%d] postWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryTooManyRequests creates a PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryTooManyRequests with default headers values
func NewPostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryTooManyRequests() *PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryTooManyRequests {
	return &PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryTooManyRequests{}
}

/*PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/agentschedules/query][%d] postWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryInternalServerError creates a PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryInternalServerError with default headers values
func NewPostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryInternalServerError() *PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryInternalServerError {
	return &PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryInternalServerError{}
}

/*PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/agentschedules/query][%d] postWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryInternalServerError  %+v", 500, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryServiceUnavailable creates a PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryServiceUnavailable with default headers values
func NewPostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryServiceUnavailable() *PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryServiceUnavailable {
	return &PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryServiceUnavailable{}
}

/*PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/agentschedules/query][%d] postWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryGatewayTimeout creates a PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryGatewayTimeout with default headers values
func NewPostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryGatewayTimeout() *PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryGatewayTimeout {
	return &PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryGatewayTimeout{}
}

/*PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/agentschedules/query][%d] postWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQueryGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

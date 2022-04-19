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

// PostWorkforcemanagementBusinessunitWeekScheduleRescheduleReader is a Reader for the PostWorkforcemanagementBusinessunitWeekScheduleReschedule structure.
type PostWorkforcemanagementBusinessunitWeekScheduleRescheduleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostWorkforcemanagementBusinessunitWeekScheduleRescheduleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostWorkforcemanagementBusinessunitWeekScheduleRescheduleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewPostWorkforcemanagementBusinessunitWeekScheduleRescheduleAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostWorkforcemanagementBusinessunitWeekScheduleRescheduleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostWorkforcemanagementBusinessunitWeekScheduleRescheduleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostWorkforcemanagementBusinessunitWeekScheduleRescheduleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostWorkforcemanagementBusinessunitWeekScheduleRescheduleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostWorkforcemanagementBusinessunitWeekScheduleRescheduleRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostWorkforcemanagementBusinessunitWeekScheduleRescheduleRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostWorkforcemanagementBusinessunitWeekScheduleRescheduleUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostWorkforcemanagementBusinessunitWeekScheduleRescheduleTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostWorkforcemanagementBusinessunitWeekScheduleRescheduleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostWorkforcemanagementBusinessunitWeekScheduleRescheduleServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostWorkforcemanagementBusinessunitWeekScheduleRescheduleGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostWorkforcemanagementBusinessunitWeekScheduleRescheduleOK creates a PostWorkforcemanagementBusinessunitWeekScheduleRescheduleOK with default headers values
func NewPostWorkforcemanagementBusinessunitWeekScheduleRescheduleOK() *PostWorkforcemanagementBusinessunitWeekScheduleRescheduleOK {
	return &PostWorkforcemanagementBusinessunitWeekScheduleRescheduleOK{}
}

/*PostWorkforcemanagementBusinessunitWeekScheduleRescheduleOK handles this case with default header values.

successful operation
*/
type PostWorkforcemanagementBusinessunitWeekScheduleRescheduleOK struct {
	Payload *models.BuAsyncScheduleRunResponse
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleRescheduleOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/reschedule][%d] postWorkforcemanagementBusinessunitWeekScheduleRescheduleOK  %+v", 200, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleRescheduleOK) GetPayload() *models.BuAsyncScheduleRunResponse {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleRescheduleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BuAsyncScheduleRunResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitWeekScheduleRescheduleAccepted creates a PostWorkforcemanagementBusinessunitWeekScheduleRescheduleAccepted with default headers values
func NewPostWorkforcemanagementBusinessunitWeekScheduleRescheduleAccepted() *PostWorkforcemanagementBusinessunitWeekScheduleRescheduleAccepted {
	return &PostWorkforcemanagementBusinessunitWeekScheduleRescheduleAccepted{}
}

/*PostWorkforcemanagementBusinessunitWeekScheduleRescheduleAccepted handles this case with default header values.

The rescheduling run was started and updates will be sent by notification
*/
type PostWorkforcemanagementBusinessunitWeekScheduleRescheduleAccepted struct {
	Payload *models.BuAsyncScheduleRunResponse
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleRescheduleAccepted) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/reschedule][%d] postWorkforcemanagementBusinessunitWeekScheduleRescheduleAccepted  %+v", 202, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleRescheduleAccepted) GetPayload() *models.BuAsyncScheduleRunResponse {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleRescheduleAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BuAsyncScheduleRunResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitWeekScheduleRescheduleBadRequest creates a PostWorkforcemanagementBusinessunitWeekScheduleRescheduleBadRequest with default headers values
func NewPostWorkforcemanagementBusinessunitWeekScheduleRescheduleBadRequest() *PostWorkforcemanagementBusinessunitWeekScheduleRescheduleBadRequest {
	return &PostWorkforcemanagementBusinessunitWeekScheduleRescheduleBadRequest{}
}

/*PostWorkforcemanagementBusinessunitWeekScheduleRescheduleBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostWorkforcemanagementBusinessunitWeekScheduleRescheduleBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleRescheduleBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/reschedule][%d] postWorkforcemanagementBusinessunitWeekScheduleRescheduleBadRequest  %+v", 400, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleRescheduleBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleRescheduleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitWeekScheduleRescheduleUnauthorized creates a PostWorkforcemanagementBusinessunitWeekScheduleRescheduleUnauthorized with default headers values
func NewPostWorkforcemanagementBusinessunitWeekScheduleRescheduleUnauthorized() *PostWorkforcemanagementBusinessunitWeekScheduleRescheduleUnauthorized {
	return &PostWorkforcemanagementBusinessunitWeekScheduleRescheduleUnauthorized{}
}

/*PostWorkforcemanagementBusinessunitWeekScheduleRescheduleUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostWorkforcemanagementBusinessunitWeekScheduleRescheduleUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleRescheduleUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/reschedule][%d] postWorkforcemanagementBusinessunitWeekScheduleRescheduleUnauthorized  %+v", 401, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleRescheduleUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleRescheduleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitWeekScheduleRescheduleForbidden creates a PostWorkforcemanagementBusinessunitWeekScheduleRescheduleForbidden with default headers values
func NewPostWorkforcemanagementBusinessunitWeekScheduleRescheduleForbidden() *PostWorkforcemanagementBusinessunitWeekScheduleRescheduleForbidden {
	return &PostWorkforcemanagementBusinessunitWeekScheduleRescheduleForbidden{}
}

/*PostWorkforcemanagementBusinessunitWeekScheduleRescheduleForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostWorkforcemanagementBusinessunitWeekScheduleRescheduleForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleRescheduleForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/reschedule][%d] postWorkforcemanagementBusinessunitWeekScheduleRescheduleForbidden  %+v", 403, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleRescheduleForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleRescheduleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitWeekScheduleRescheduleNotFound creates a PostWorkforcemanagementBusinessunitWeekScheduleRescheduleNotFound with default headers values
func NewPostWorkforcemanagementBusinessunitWeekScheduleRescheduleNotFound() *PostWorkforcemanagementBusinessunitWeekScheduleRescheduleNotFound {
	return &PostWorkforcemanagementBusinessunitWeekScheduleRescheduleNotFound{}
}

/*PostWorkforcemanagementBusinessunitWeekScheduleRescheduleNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostWorkforcemanagementBusinessunitWeekScheduleRescheduleNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleRescheduleNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/reschedule][%d] postWorkforcemanagementBusinessunitWeekScheduleRescheduleNotFound  %+v", 404, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleRescheduleNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleRescheduleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitWeekScheduleRescheduleRequestTimeout creates a PostWorkforcemanagementBusinessunitWeekScheduleRescheduleRequestTimeout with default headers values
func NewPostWorkforcemanagementBusinessunitWeekScheduleRescheduleRequestTimeout() *PostWorkforcemanagementBusinessunitWeekScheduleRescheduleRequestTimeout {
	return &PostWorkforcemanagementBusinessunitWeekScheduleRescheduleRequestTimeout{}
}

/*PostWorkforcemanagementBusinessunitWeekScheduleRescheduleRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostWorkforcemanagementBusinessunitWeekScheduleRescheduleRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleRescheduleRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/reschedule][%d] postWorkforcemanagementBusinessunitWeekScheduleRescheduleRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleRescheduleRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleRescheduleRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitWeekScheduleRescheduleRequestEntityTooLarge creates a PostWorkforcemanagementBusinessunitWeekScheduleRescheduleRequestEntityTooLarge with default headers values
func NewPostWorkforcemanagementBusinessunitWeekScheduleRescheduleRequestEntityTooLarge() *PostWorkforcemanagementBusinessunitWeekScheduleRescheduleRequestEntityTooLarge {
	return &PostWorkforcemanagementBusinessunitWeekScheduleRescheduleRequestEntityTooLarge{}
}

/*PostWorkforcemanagementBusinessunitWeekScheduleRescheduleRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PostWorkforcemanagementBusinessunitWeekScheduleRescheduleRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleRescheduleRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/reschedule][%d] postWorkforcemanagementBusinessunitWeekScheduleRescheduleRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleRescheduleRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleRescheduleRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitWeekScheduleRescheduleUnsupportedMediaType creates a PostWorkforcemanagementBusinessunitWeekScheduleRescheduleUnsupportedMediaType with default headers values
func NewPostWorkforcemanagementBusinessunitWeekScheduleRescheduleUnsupportedMediaType() *PostWorkforcemanagementBusinessunitWeekScheduleRescheduleUnsupportedMediaType {
	return &PostWorkforcemanagementBusinessunitWeekScheduleRescheduleUnsupportedMediaType{}
}

/*PostWorkforcemanagementBusinessunitWeekScheduleRescheduleUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostWorkforcemanagementBusinessunitWeekScheduleRescheduleUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleRescheduleUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/reschedule][%d] postWorkforcemanagementBusinessunitWeekScheduleRescheduleUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleRescheduleUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleRescheduleUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitWeekScheduleRescheduleTooManyRequests creates a PostWorkforcemanagementBusinessunitWeekScheduleRescheduleTooManyRequests with default headers values
func NewPostWorkforcemanagementBusinessunitWeekScheduleRescheduleTooManyRequests() *PostWorkforcemanagementBusinessunitWeekScheduleRescheduleTooManyRequests {
	return &PostWorkforcemanagementBusinessunitWeekScheduleRescheduleTooManyRequests{}
}

/*PostWorkforcemanagementBusinessunitWeekScheduleRescheduleTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostWorkforcemanagementBusinessunitWeekScheduleRescheduleTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleRescheduleTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/reschedule][%d] postWorkforcemanagementBusinessunitWeekScheduleRescheduleTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleRescheduleTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleRescheduleTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitWeekScheduleRescheduleInternalServerError creates a PostWorkforcemanagementBusinessunitWeekScheduleRescheduleInternalServerError with default headers values
func NewPostWorkforcemanagementBusinessunitWeekScheduleRescheduleInternalServerError() *PostWorkforcemanagementBusinessunitWeekScheduleRescheduleInternalServerError {
	return &PostWorkforcemanagementBusinessunitWeekScheduleRescheduleInternalServerError{}
}

/*PostWorkforcemanagementBusinessunitWeekScheduleRescheduleInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostWorkforcemanagementBusinessunitWeekScheduleRescheduleInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleRescheduleInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/reschedule][%d] postWorkforcemanagementBusinessunitWeekScheduleRescheduleInternalServerError  %+v", 500, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleRescheduleInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleRescheduleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitWeekScheduleRescheduleServiceUnavailable creates a PostWorkforcemanagementBusinessunitWeekScheduleRescheduleServiceUnavailable with default headers values
func NewPostWorkforcemanagementBusinessunitWeekScheduleRescheduleServiceUnavailable() *PostWorkforcemanagementBusinessunitWeekScheduleRescheduleServiceUnavailable {
	return &PostWorkforcemanagementBusinessunitWeekScheduleRescheduleServiceUnavailable{}
}

/*PostWorkforcemanagementBusinessunitWeekScheduleRescheduleServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostWorkforcemanagementBusinessunitWeekScheduleRescheduleServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleRescheduleServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/reschedule][%d] postWorkforcemanagementBusinessunitWeekScheduleRescheduleServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleRescheduleServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleRescheduleServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitWeekScheduleRescheduleGatewayTimeout creates a PostWorkforcemanagementBusinessunitWeekScheduleRescheduleGatewayTimeout with default headers values
func NewPostWorkforcemanagementBusinessunitWeekScheduleRescheduleGatewayTimeout() *PostWorkforcemanagementBusinessunitWeekScheduleRescheduleGatewayTimeout {
	return &PostWorkforcemanagementBusinessunitWeekScheduleRescheduleGatewayTimeout{}
}

/*PostWorkforcemanagementBusinessunitWeekScheduleRescheduleGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostWorkforcemanagementBusinessunitWeekScheduleRescheduleGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleRescheduleGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/reschedule][%d] postWorkforcemanagementBusinessunitWeekScheduleRescheduleGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleRescheduleGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleRescheduleGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

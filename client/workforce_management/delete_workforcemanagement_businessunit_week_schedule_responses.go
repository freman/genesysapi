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

// DeleteWorkforcemanagementBusinessunitWeekScheduleReader is a Reader for the DeleteWorkforcemanagementBusinessunitWeekSchedule structure.
type DeleteWorkforcemanagementBusinessunitWeekScheduleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteWorkforcemanagementBusinessunitWeekScheduleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteWorkforcemanagementBusinessunitWeekScheduleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewDeleteWorkforcemanagementBusinessunitWeekScheduleAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteWorkforcemanagementBusinessunitWeekScheduleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteWorkforcemanagementBusinessunitWeekScheduleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteWorkforcemanagementBusinessunitWeekScheduleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteWorkforcemanagementBusinessunitWeekScheduleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteWorkforcemanagementBusinessunitWeekScheduleRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteWorkforcemanagementBusinessunitWeekScheduleUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteWorkforcemanagementBusinessunitWeekScheduleTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteWorkforcemanagementBusinessunitWeekScheduleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteWorkforcemanagementBusinessunitWeekScheduleServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteWorkforcemanagementBusinessunitWeekScheduleGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteWorkforcemanagementBusinessunitWeekScheduleOK creates a DeleteWorkforcemanagementBusinessunitWeekScheduleOK with default headers values
func NewDeleteWorkforcemanagementBusinessunitWeekScheduleOK() *DeleteWorkforcemanagementBusinessunitWeekScheduleOK {
	return &DeleteWorkforcemanagementBusinessunitWeekScheduleOK{}
}

/*DeleteWorkforcemanagementBusinessunitWeekScheduleOK handles this case with default header values.

successful operation
*/
type DeleteWorkforcemanagementBusinessunitWeekScheduleOK struct {
	Payload *models.BuAsyncScheduleResponse
}

func (o *DeleteWorkforcemanagementBusinessunitWeekScheduleOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}][%d] deleteWorkforcemanagementBusinessunitWeekScheduleOK  %+v", 200, o.Payload)
}

func (o *DeleteWorkforcemanagementBusinessunitWeekScheduleOK) GetPayload() *models.BuAsyncScheduleResponse {
	return o.Payload
}

func (o *DeleteWorkforcemanagementBusinessunitWeekScheduleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BuAsyncScheduleResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWorkforcemanagementBusinessunitWeekScheduleAccepted creates a DeleteWorkforcemanagementBusinessunitWeekScheduleAccepted with default headers values
func NewDeleteWorkforcemanagementBusinessunitWeekScheduleAccepted() *DeleteWorkforcemanagementBusinessunitWeekScheduleAccepted {
	return &DeleteWorkforcemanagementBusinessunitWeekScheduleAccepted{}
}

/*DeleteWorkforcemanagementBusinessunitWeekScheduleAccepted handles this case with default header values.

The schedule is being deleted
*/
type DeleteWorkforcemanagementBusinessunitWeekScheduleAccepted struct {
}

func (o *DeleteWorkforcemanagementBusinessunitWeekScheduleAccepted) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}][%d] deleteWorkforcemanagementBusinessunitWeekScheduleAccepted ", 202)
}

func (o *DeleteWorkforcemanagementBusinessunitWeekScheduleAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteWorkforcemanagementBusinessunitWeekScheduleBadRequest creates a DeleteWorkforcemanagementBusinessunitWeekScheduleBadRequest with default headers values
func NewDeleteWorkforcemanagementBusinessunitWeekScheduleBadRequest() *DeleteWorkforcemanagementBusinessunitWeekScheduleBadRequest {
	return &DeleteWorkforcemanagementBusinessunitWeekScheduleBadRequest{}
}

/*DeleteWorkforcemanagementBusinessunitWeekScheduleBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteWorkforcemanagementBusinessunitWeekScheduleBadRequest struct {
	Payload *models.ErrorBody
}

func (o *DeleteWorkforcemanagementBusinessunitWeekScheduleBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}][%d] deleteWorkforcemanagementBusinessunitWeekScheduleBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteWorkforcemanagementBusinessunitWeekScheduleBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWorkforcemanagementBusinessunitWeekScheduleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWorkforcemanagementBusinessunitWeekScheduleUnauthorized creates a DeleteWorkforcemanagementBusinessunitWeekScheduleUnauthorized with default headers values
func NewDeleteWorkforcemanagementBusinessunitWeekScheduleUnauthorized() *DeleteWorkforcemanagementBusinessunitWeekScheduleUnauthorized {
	return &DeleteWorkforcemanagementBusinessunitWeekScheduleUnauthorized{}
}

/*DeleteWorkforcemanagementBusinessunitWeekScheduleUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteWorkforcemanagementBusinessunitWeekScheduleUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *DeleteWorkforcemanagementBusinessunitWeekScheduleUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}][%d] deleteWorkforcemanagementBusinessunitWeekScheduleUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteWorkforcemanagementBusinessunitWeekScheduleUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWorkforcemanagementBusinessunitWeekScheduleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWorkforcemanagementBusinessunitWeekScheduleForbidden creates a DeleteWorkforcemanagementBusinessunitWeekScheduleForbidden with default headers values
func NewDeleteWorkforcemanagementBusinessunitWeekScheduleForbidden() *DeleteWorkforcemanagementBusinessunitWeekScheduleForbidden {
	return &DeleteWorkforcemanagementBusinessunitWeekScheduleForbidden{}
}

/*DeleteWorkforcemanagementBusinessunitWeekScheduleForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type DeleteWorkforcemanagementBusinessunitWeekScheduleForbidden struct {
	Payload *models.ErrorBody
}

func (o *DeleteWorkforcemanagementBusinessunitWeekScheduleForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}][%d] deleteWorkforcemanagementBusinessunitWeekScheduleForbidden  %+v", 403, o.Payload)
}

func (o *DeleteWorkforcemanagementBusinessunitWeekScheduleForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWorkforcemanagementBusinessunitWeekScheduleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWorkforcemanagementBusinessunitWeekScheduleNotFound creates a DeleteWorkforcemanagementBusinessunitWeekScheduleNotFound with default headers values
func NewDeleteWorkforcemanagementBusinessunitWeekScheduleNotFound() *DeleteWorkforcemanagementBusinessunitWeekScheduleNotFound {
	return &DeleteWorkforcemanagementBusinessunitWeekScheduleNotFound{}
}

/*DeleteWorkforcemanagementBusinessunitWeekScheduleNotFound handles this case with default header values.

The requested resource was not found.
*/
type DeleteWorkforcemanagementBusinessunitWeekScheduleNotFound struct {
	Payload *models.ErrorBody
}

func (o *DeleteWorkforcemanagementBusinessunitWeekScheduleNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}][%d] deleteWorkforcemanagementBusinessunitWeekScheduleNotFound  %+v", 404, o.Payload)
}

func (o *DeleteWorkforcemanagementBusinessunitWeekScheduleNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWorkforcemanagementBusinessunitWeekScheduleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWorkforcemanagementBusinessunitWeekScheduleRequestEntityTooLarge creates a DeleteWorkforcemanagementBusinessunitWeekScheduleRequestEntityTooLarge with default headers values
func NewDeleteWorkforcemanagementBusinessunitWeekScheduleRequestEntityTooLarge() *DeleteWorkforcemanagementBusinessunitWeekScheduleRequestEntityTooLarge {
	return &DeleteWorkforcemanagementBusinessunitWeekScheduleRequestEntityTooLarge{}
}

/*DeleteWorkforcemanagementBusinessunitWeekScheduleRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type DeleteWorkforcemanagementBusinessunitWeekScheduleRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *DeleteWorkforcemanagementBusinessunitWeekScheduleRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}][%d] deleteWorkforcemanagementBusinessunitWeekScheduleRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteWorkforcemanagementBusinessunitWeekScheduleRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWorkforcemanagementBusinessunitWeekScheduleRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWorkforcemanagementBusinessunitWeekScheduleUnsupportedMediaType creates a DeleteWorkforcemanagementBusinessunitWeekScheduleUnsupportedMediaType with default headers values
func NewDeleteWorkforcemanagementBusinessunitWeekScheduleUnsupportedMediaType() *DeleteWorkforcemanagementBusinessunitWeekScheduleUnsupportedMediaType {
	return &DeleteWorkforcemanagementBusinessunitWeekScheduleUnsupportedMediaType{}
}

/*DeleteWorkforcemanagementBusinessunitWeekScheduleUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteWorkforcemanagementBusinessunitWeekScheduleUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *DeleteWorkforcemanagementBusinessunitWeekScheduleUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}][%d] deleteWorkforcemanagementBusinessunitWeekScheduleUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteWorkforcemanagementBusinessunitWeekScheduleUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWorkforcemanagementBusinessunitWeekScheduleUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWorkforcemanagementBusinessunitWeekScheduleTooManyRequests creates a DeleteWorkforcemanagementBusinessunitWeekScheduleTooManyRequests with default headers values
func NewDeleteWorkforcemanagementBusinessunitWeekScheduleTooManyRequests() *DeleteWorkforcemanagementBusinessunitWeekScheduleTooManyRequests {
	return &DeleteWorkforcemanagementBusinessunitWeekScheduleTooManyRequests{}
}

/*DeleteWorkforcemanagementBusinessunitWeekScheduleTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type DeleteWorkforcemanagementBusinessunitWeekScheduleTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *DeleteWorkforcemanagementBusinessunitWeekScheduleTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}][%d] deleteWorkforcemanagementBusinessunitWeekScheduleTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteWorkforcemanagementBusinessunitWeekScheduleTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWorkforcemanagementBusinessunitWeekScheduleTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWorkforcemanagementBusinessunitWeekScheduleInternalServerError creates a DeleteWorkforcemanagementBusinessunitWeekScheduleInternalServerError with default headers values
func NewDeleteWorkforcemanagementBusinessunitWeekScheduleInternalServerError() *DeleteWorkforcemanagementBusinessunitWeekScheduleInternalServerError {
	return &DeleteWorkforcemanagementBusinessunitWeekScheduleInternalServerError{}
}

/*DeleteWorkforcemanagementBusinessunitWeekScheduleInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteWorkforcemanagementBusinessunitWeekScheduleInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *DeleteWorkforcemanagementBusinessunitWeekScheduleInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}][%d] deleteWorkforcemanagementBusinessunitWeekScheduleInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteWorkforcemanagementBusinessunitWeekScheduleInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWorkforcemanagementBusinessunitWeekScheduleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWorkforcemanagementBusinessunitWeekScheduleServiceUnavailable creates a DeleteWorkforcemanagementBusinessunitWeekScheduleServiceUnavailable with default headers values
func NewDeleteWorkforcemanagementBusinessunitWeekScheduleServiceUnavailable() *DeleteWorkforcemanagementBusinessunitWeekScheduleServiceUnavailable {
	return &DeleteWorkforcemanagementBusinessunitWeekScheduleServiceUnavailable{}
}

/*DeleteWorkforcemanagementBusinessunitWeekScheduleServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteWorkforcemanagementBusinessunitWeekScheduleServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *DeleteWorkforcemanagementBusinessunitWeekScheduleServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}][%d] deleteWorkforcemanagementBusinessunitWeekScheduleServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteWorkforcemanagementBusinessunitWeekScheduleServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWorkforcemanagementBusinessunitWeekScheduleServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWorkforcemanagementBusinessunitWeekScheduleGatewayTimeout creates a DeleteWorkforcemanagementBusinessunitWeekScheduleGatewayTimeout with default headers values
func NewDeleteWorkforcemanagementBusinessunitWeekScheduleGatewayTimeout() *DeleteWorkforcemanagementBusinessunitWeekScheduleGatewayTimeout {
	return &DeleteWorkforcemanagementBusinessunitWeekScheduleGatewayTimeout{}
}

/*DeleteWorkforcemanagementBusinessunitWeekScheduleGatewayTimeout handles this case with default header values.

The request timed out.
*/
type DeleteWorkforcemanagementBusinessunitWeekScheduleGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteWorkforcemanagementBusinessunitWeekScheduleGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}][%d] deleteWorkforcemanagementBusinessunitWeekScheduleGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteWorkforcemanagementBusinessunitWeekScheduleGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWorkforcemanagementBusinessunitWeekScheduleGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
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

// PostWorkforcemanagementBusinessunitWeekScheduleCopyReader is a Reader for the PostWorkforcemanagementBusinessunitWeekScheduleCopy structure.
type PostWorkforcemanagementBusinessunitWeekScheduleCopyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostWorkforcemanagementBusinessunitWeekScheduleCopyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostWorkforcemanagementBusinessunitWeekScheduleCopyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostWorkforcemanagementBusinessunitWeekScheduleCopyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostWorkforcemanagementBusinessunitWeekScheduleCopyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostWorkforcemanagementBusinessunitWeekScheduleCopyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostWorkforcemanagementBusinessunitWeekScheduleCopyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostWorkforcemanagementBusinessunitWeekScheduleCopyRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostWorkforcemanagementBusinessunitWeekScheduleCopyUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostWorkforcemanagementBusinessunitWeekScheduleCopyTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostWorkforcemanagementBusinessunitWeekScheduleCopyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostWorkforcemanagementBusinessunitWeekScheduleCopyServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostWorkforcemanagementBusinessunitWeekScheduleCopyGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostWorkforcemanagementBusinessunitWeekScheduleCopyOK creates a PostWorkforcemanagementBusinessunitWeekScheduleCopyOK with default headers values
func NewPostWorkforcemanagementBusinessunitWeekScheduleCopyOK() *PostWorkforcemanagementBusinessunitWeekScheduleCopyOK {
	return &PostWorkforcemanagementBusinessunitWeekScheduleCopyOK{}
}

/*PostWorkforcemanagementBusinessunitWeekScheduleCopyOK handles this case with default header values.

successful operation
*/
type PostWorkforcemanagementBusinessunitWeekScheduleCopyOK struct {
	Payload *models.BuAsyncScheduleResponse
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleCopyOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/copy][%d] postWorkforcemanagementBusinessunitWeekScheduleCopyOK  %+v", 200, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleCopyOK) GetPayload() *models.BuAsyncScheduleResponse {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleCopyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BuAsyncScheduleResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitWeekScheduleCopyBadRequest creates a PostWorkforcemanagementBusinessunitWeekScheduleCopyBadRequest with default headers values
func NewPostWorkforcemanagementBusinessunitWeekScheduleCopyBadRequest() *PostWorkforcemanagementBusinessunitWeekScheduleCopyBadRequest {
	return &PostWorkforcemanagementBusinessunitWeekScheduleCopyBadRequest{}
}

/*PostWorkforcemanagementBusinessunitWeekScheduleCopyBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostWorkforcemanagementBusinessunitWeekScheduleCopyBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleCopyBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/copy][%d] postWorkforcemanagementBusinessunitWeekScheduleCopyBadRequest  %+v", 400, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleCopyBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleCopyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitWeekScheduleCopyUnauthorized creates a PostWorkforcemanagementBusinessunitWeekScheduleCopyUnauthorized with default headers values
func NewPostWorkforcemanagementBusinessunitWeekScheduleCopyUnauthorized() *PostWorkforcemanagementBusinessunitWeekScheduleCopyUnauthorized {
	return &PostWorkforcemanagementBusinessunitWeekScheduleCopyUnauthorized{}
}

/*PostWorkforcemanagementBusinessunitWeekScheduleCopyUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostWorkforcemanagementBusinessunitWeekScheduleCopyUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleCopyUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/copy][%d] postWorkforcemanagementBusinessunitWeekScheduleCopyUnauthorized  %+v", 401, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleCopyUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleCopyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitWeekScheduleCopyForbidden creates a PostWorkforcemanagementBusinessunitWeekScheduleCopyForbidden with default headers values
func NewPostWorkforcemanagementBusinessunitWeekScheduleCopyForbidden() *PostWorkforcemanagementBusinessunitWeekScheduleCopyForbidden {
	return &PostWorkforcemanagementBusinessunitWeekScheduleCopyForbidden{}
}

/*PostWorkforcemanagementBusinessunitWeekScheduleCopyForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostWorkforcemanagementBusinessunitWeekScheduleCopyForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleCopyForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/copy][%d] postWorkforcemanagementBusinessunitWeekScheduleCopyForbidden  %+v", 403, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleCopyForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleCopyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitWeekScheduleCopyNotFound creates a PostWorkforcemanagementBusinessunitWeekScheduleCopyNotFound with default headers values
func NewPostWorkforcemanagementBusinessunitWeekScheduleCopyNotFound() *PostWorkforcemanagementBusinessunitWeekScheduleCopyNotFound {
	return &PostWorkforcemanagementBusinessunitWeekScheduleCopyNotFound{}
}

/*PostWorkforcemanagementBusinessunitWeekScheduleCopyNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostWorkforcemanagementBusinessunitWeekScheduleCopyNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleCopyNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/copy][%d] postWorkforcemanagementBusinessunitWeekScheduleCopyNotFound  %+v", 404, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleCopyNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleCopyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitWeekScheduleCopyRequestEntityTooLarge creates a PostWorkforcemanagementBusinessunitWeekScheduleCopyRequestEntityTooLarge with default headers values
func NewPostWorkforcemanagementBusinessunitWeekScheduleCopyRequestEntityTooLarge() *PostWorkforcemanagementBusinessunitWeekScheduleCopyRequestEntityTooLarge {
	return &PostWorkforcemanagementBusinessunitWeekScheduleCopyRequestEntityTooLarge{}
}

/*PostWorkforcemanagementBusinessunitWeekScheduleCopyRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostWorkforcemanagementBusinessunitWeekScheduleCopyRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleCopyRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/copy][%d] postWorkforcemanagementBusinessunitWeekScheduleCopyRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleCopyRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleCopyRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitWeekScheduleCopyUnsupportedMediaType creates a PostWorkforcemanagementBusinessunitWeekScheduleCopyUnsupportedMediaType with default headers values
func NewPostWorkforcemanagementBusinessunitWeekScheduleCopyUnsupportedMediaType() *PostWorkforcemanagementBusinessunitWeekScheduleCopyUnsupportedMediaType {
	return &PostWorkforcemanagementBusinessunitWeekScheduleCopyUnsupportedMediaType{}
}

/*PostWorkforcemanagementBusinessunitWeekScheduleCopyUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostWorkforcemanagementBusinessunitWeekScheduleCopyUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleCopyUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/copy][%d] postWorkforcemanagementBusinessunitWeekScheduleCopyUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleCopyUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleCopyUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitWeekScheduleCopyTooManyRequests creates a PostWorkforcemanagementBusinessunitWeekScheduleCopyTooManyRequests with default headers values
func NewPostWorkforcemanagementBusinessunitWeekScheduleCopyTooManyRequests() *PostWorkforcemanagementBusinessunitWeekScheduleCopyTooManyRequests {
	return &PostWorkforcemanagementBusinessunitWeekScheduleCopyTooManyRequests{}
}

/*PostWorkforcemanagementBusinessunitWeekScheduleCopyTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PostWorkforcemanagementBusinessunitWeekScheduleCopyTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleCopyTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/copy][%d] postWorkforcemanagementBusinessunitWeekScheduleCopyTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleCopyTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleCopyTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitWeekScheduleCopyInternalServerError creates a PostWorkforcemanagementBusinessunitWeekScheduleCopyInternalServerError with default headers values
func NewPostWorkforcemanagementBusinessunitWeekScheduleCopyInternalServerError() *PostWorkforcemanagementBusinessunitWeekScheduleCopyInternalServerError {
	return &PostWorkforcemanagementBusinessunitWeekScheduleCopyInternalServerError{}
}

/*PostWorkforcemanagementBusinessunitWeekScheduleCopyInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostWorkforcemanagementBusinessunitWeekScheduleCopyInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleCopyInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/copy][%d] postWorkforcemanagementBusinessunitWeekScheduleCopyInternalServerError  %+v", 500, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleCopyInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleCopyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitWeekScheduleCopyServiceUnavailable creates a PostWorkforcemanagementBusinessunitWeekScheduleCopyServiceUnavailable with default headers values
func NewPostWorkforcemanagementBusinessunitWeekScheduleCopyServiceUnavailable() *PostWorkforcemanagementBusinessunitWeekScheduleCopyServiceUnavailable {
	return &PostWorkforcemanagementBusinessunitWeekScheduleCopyServiceUnavailable{}
}

/*PostWorkforcemanagementBusinessunitWeekScheduleCopyServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostWorkforcemanagementBusinessunitWeekScheduleCopyServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleCopyServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/copy][%d] postWorkforcemanagementBusinessunitWeekScheduleCopyServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleCopyServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleCopyServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitWeekScheduleCopyGatewayTimeout creates a PostWorkforcemanagementBusinessunitWeekScheduleCopyGatewayTimeout with default headers values
func NewPostWorkforcemanagementBusinessunitWeekScheduleCopyGatewayTimeout() *PostWorkforcemanagementBusinessunitWeekScheduleCopyGatewayTimeout {
	return &PostWorkforcemanagementBusinessunitWeekScheduleCopyGatewayTimeout{}
}

/*PostWorkforcemanagementBusinessunitWeekScheduleCopyGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostWorkforcemanagementBusinessunitWeekScheduleCopyGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleCopyGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/copy][%d] postWorkforcemanagementBusinessunitWeekScheduleCopyGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleCopyGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleCopyGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

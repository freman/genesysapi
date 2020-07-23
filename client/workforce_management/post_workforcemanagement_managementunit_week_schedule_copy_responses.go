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

// PostWorkforcemanagementManagementunitWeekScheduleCopyReader is a Reader for the PostWorkforcemanagementManagementunitWeekScheduleCopy structure.
type PostWorkforcemanagementManagementunitWeekScheduleCopyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostWorkforcemanagementManagementunitWeekScheduleCopyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 400:
		result := NewPostWorkforcemanagementManagementunitWeekScheduleCopyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostWorkforcemanagementManagementunitWeekScheduleCopyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostWorkforcemanagementManagementunitWeekScheduleCopyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostWorkforcemanagementManagementunitWeekScheduleCopyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewPostWorkforcemanagementManagementunitWeekScheduleCopyGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostWorkforcemanagementManagementunitWeekScheduleCopyRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostWorkforcemanagementManagementunitWeekScheduleCopyUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostWorkforcemanagementManagementunitWeekScheduleCopyTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostWorkforcemanagementManagementunitWeekScheduleCopyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostWorkforcemanagementManagementunitWeekScheduleCopyServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostWorkforcemanagementManagementunitWeekScheduleCopyGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPostWorkforcemanagementManagementunitWeekScheduleCopyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostWorkforcemanagementManagementunitWeekScheduleCopyBadRequest creates a PostWorkforcemanagementManagementunitWeekScheduleCopyBadRequest with default headers values
func NewPostWorkforcemanagementManagementunitWeekScheduleCopyBadRequest() *PostWorkforcemanagementManagementunitWeekScheduleCopyBadRequest {
	return &PostWorkforcemanagementManagementunitWeekScheduleCopyBadRequest{}
}

/*PostWorkforcemanagementManagementunitWeekScheduleCopyBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostWorkforcemanagementManagementunitWeekScheduleCopyBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementManagementunitWeekScheduleCopyBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekId}/schedules/{scheduleId}/copy][%d] postWorkforcemanagementManagementunitWeekScheduleCopyBadRequest  %+v", 400, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitWeekScheduleCopyBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitWeekScheduleCopyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitWeekScheduleCopyUnauthorized creates a PostWorkforcemanagementManagementunitWeekScheduleCopyUnauthorized with default headers values
func NewPostWorkforcemanagementManagementunitWeekScheduleCopyUnauthorized() *PostWorkforcemanagementManagementunitWeekScheduleCopyUnauthorized {
	return &PostWorkforcemanagementManagementunitWeekScheduleCopyUnauthorized{}
}

/*PostWorkforcemanagementManagementunitWeekScheduleCopyUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostWorkforcemanagementManagementunitWeekScheduleCopyUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementManagementunitWeekScheduleCopyUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekId}/schedules/{scheduleId}/copy][%d] postWorkforcemanagementManagementunitWeekScheduleCopyUnauthorized  %+v", 401, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitWeekScheduleCopyUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitWeekScheduleCopyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitWeekScheduleCopyForbidden creates a PostWorkforcemanagementManagementunitWeekScheduleCopyForbidden with default headers values
func NewPostWorkforcemanagementManagementunitWeekScheduleCopyForbidden() *PostWorkforcemanagementManagementunitWeekScheduleCopyForbidden {
	return &PostWorkforcemanagementManagementunitWeekScheduleCopyForbidden{}
}

/*PostWorkforcemanagementManagementunitWeekScheduleCopyForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostWorkforcemanagementManagementunitWeekScheduleCopyForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementManagementunitWeekScheduleCopyForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekId}/schedules/{scheduleId}/copy][%d] postWorkforcemanagementManagementunitWeekScheduleCopyForbidden  %+v", 403, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitWeekScheduleCopyForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitWeekScheduleCopyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitWeekScheduleCopyNotFound creates a PostWorkforcemanagementManagementunitWeekScheduleCopyNotFound with default headers values
func NewPostWorkforcemanagementManagementunitWeekScheduleCopyNotFound() *PostWorkforcemanagementManagementunitWeekScheduleCopyNotFound {
	return &PostWorkforcemanagementManagementunitWeekScheduleCopyNotFound{}
}

/*PostWorkforcemanagementManagementunitWeekScheduleCopyNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostWorkforcemanagementManagementunitWeekScheduleCopyNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementManagementunitWeekScheduleCopyNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekId}/schedules/{scheduleId}/copy][%d] postWorkforcemanagementManagementunitWeekScheduleCopyNotFound  %+v", 404, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitWeekScheduleCopyNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitWeekScheduleCopyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitWeekScheduleCopyGone creates a PostWorkforcemanagementManagementunitWeekScheduleCopyGone with default headers values
func NewPostWorkforcemanagementManagementunitWeekScheduleCopyGone() *PostWorkforcemanagementManagementunitWeekScheduleCopyGone {
	return &PostWorkforcemanagementManagementunitWeekScheduleCopyGone{}
}

/*PostWorkforcemanagementManagementunitWeekScheduleCopyGone handles this case with default header values.

Gone
*/
type PostWorkforcemanagementManagementunitWeekScheduleCopyGone struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementManagementunitWeekScheduleCopyGone) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekId}/schedules/{scheduleId}/copy][%d] postWorkforcemanagementManagementunitWeekScheduleCopyGone  %+v", 410, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitWeekScheduleCopyGone) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitWeekScheduleCopyGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitWeekScheduleCopyRequestEntityTooLarge creates a PostWorkforcemanagementManagementunitWeekScheduleCopyRequestEntityTooLarge with default headers values
func NewPostWorkforcemanagementManagementunitWeekScheduleCopyRequestEntityTooLarge() *PostWorkforcemanagementManagementunitWeekScheduleCopyRequestEntityTooLarge {
	return &PostWorkforcemanagementManagementunitWeekScheduleCopyRequestEntityTooLarge{}
}

/*PostWorkforcemanagementManagementunitWeekScheduleCopyRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostWorkforcemanagementManagementunitWeekScheduleCopyRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementManagementunitWeekScheduleCopyRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekId}/schedules/{scheduleId}/copy][%d] postWorkforcemanagementManagementunitWeekScheduleCopyRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitWeekScheduleCopyRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitWeekScheduleCopyRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitWeekScheduleCopyUnsupportedMediaType creates a PostWorkforcemanagementManagementunitWeekScheduleCopyUnsupportedMediaType with default headers values
func NewPostWorkforcemanagementManagementunitWeekScheduleCopyUnsupportedMediaType() *PostWorkforcemanagementManagementunitWeekScheduleCopyUnsupportedMediaType {
	return &PostWorkforcemanagementManagementunitWeekScheduleCopyUnsupportedMediaType{}
}

/*PostWorkforcemanagementManagementunitWeekScheduleCopyUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostWorkforcemanagementManagementunitWeekScheduleCopyUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementManagementunitWeekScheduleCopyUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekId}/schedules/{scheduleId}/copy][%d] postWorkforcemanagementManagementunitWeekScheduleCopyUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitWeekScheduleCopyUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitWeekScheduleCopyUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitWeekScheduleCopyTooManyRequests creates a PostWorkforcemanagementManagementunitWeekScheduleCopyTooManyRequests with default headers values
func NewPostWorkforcemanagementManagementunitWeekScheduleCopyTooManyRequests() *PostWorkforcemanagementManagementunitWeekScheduleCopyTooManyRequests {
	return &PostWorkforcemanagementManagementunitWeekScheduleCopyTooManyRequests{}
}

/*PostWorkforcemanagementManagementunitWeekScheduleCopyTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PostWorkforcemanagementManagementunitWeekScheduleCopyTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementManagementunitWeekScheduleCopyTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekId}/schedules/{scheduleId}/copy][%d] postWorkforcemanagementManagementunitWeekScheduleCopyTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitWeekScheduleCopyTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitWeekScheduleCopyTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitWeekScheduleCopyInternalServerError creates a PostWorkforcemanagementManagementunitWeekScheduleCopyInternalServerError with default headers values
func NewPostWorkforcemanagementManagementunitWeekScheduleCopyInternalServerError() *PostWorkforcemanagementManagementunitWeekScheduleCopyInternalServerError {
	return &PostWorkforcemanagementManagementunitWeekScheduleCopyInternalServerError{}
}

/*PostWorkforcemanagementManagementunitWeekScheduleCopyInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostWorkforcemanagementManagementunitWeekScheduleCopyInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementManagementunitWeekScheduleCopyInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekId}/schedules/{scheduleId}/copy][%d] postWorkforcemanagementManagementunitWeekScheduleCopyInternalServerError  %+v", 500, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitWeekScheduleCopyInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitWeekScheduleCopyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitWeekScheduleCopyServiceUnavailable creates a PostWorkforcemanagementManagementunitWeekScheduleCopyServiceUnavailable with default headers values
func NewPostWorkforcemanagementManagementunitWeekScheduleCopyServiceUnavailable() *PostWorkforcemanagementManagementunitWeekScheduleCopyServiceUnavailable {
	return &PostWorkforcemanagementManagementunitWeekScheduleCopyServiceUnavailable{}
}

/*PostWorkforcemanagementManagementunitWeekScheduleCopyServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostWorkforcemanagementManagementunitWeekScheduleCopyServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementManagementunitWeekScheduleCopyServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekId}/schedules/{scheduleId}/copy][%d] postWorkforcemanagementManagementunitWeekScheduleCopyServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitWeekScheduleCopyServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitWeekScheduleCopyServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitWeekScheduleCopyGatewayTimeout creates a PostWorkforcemanagementManagementunitWeekScheduleCopyGatewayTimeout with default headers values
func NewPostWorkforcemanagementManagementunitWeekScheduleCopyGatewayTimeout() *PostWorkforcemanagementManagementunitWeekScheduleCopyGatewayTimeout {
	return &PostWorkforcemanagementManagementunitWeekScheduleCopyGatewayTimeout{}
}

/*PostWorkforcemanagementManagementunitWeekScheduleCopyGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostWorkforcemanagementManagementunitWeekScheduleCopyGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementManagementunitWeekScheduleCopyGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekId}/schedules/{scheduleId}/copy][%d] postWorkforcemanagementManagementunitWeekScheduleCopyGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitWeekScheduleCopyGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitWeekScheduleCopyGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitWeekScheduleCopyDefault creates a PostWorkforcemanagementManagementunitWeekScheduleCopyDefault with default headers values
func NewPostWorkforcemanagementManagementunitWeekScheduleCopyDefault(code int) *PostWorkforcemanagementManagementunitWeekScheduleCopyDefault {
	return &PostWorkforcemanagementManagementunitWeekScheduleCopyDefault{
		_statusCode: code,
	}
}

/*PostWorkforcemanagementManagementunitWeekScheduleCopyDefault handles this case with default header values.

successful operation
*/
type PostWorkforcemanagementManagementunitWeekScheduleCopyDefault struct {
	_statusCode int
}

// Code gets the status code for the post workforcemanagement managementunit week schedule copy default response
func (o *PostWorkforcemanagementManagementunitWeekScheduleCopyDefault) Code() int {
	return o._statusCode
}

func (o *PostWorkforcemanagementManagementunitWeekScheduleCopyDefault) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekId}/schedules/{scheduleId}/copy][%d] postWorkforcemanagementManagementunitWeekScheduleCopy default ", o._statusCode)
}

func (o *PostWorkforcemanagementManagementunitWeekScheduleCopyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

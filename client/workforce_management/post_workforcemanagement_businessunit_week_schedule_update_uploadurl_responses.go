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

// PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlReader is a Reader for the PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurl structure.
type PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewPostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlOK creates a PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlOK with default headers values
func NewPostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlOK() *PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlOK {
	return &PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlOK{}
}

/*PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlOK handles this case with default header values.

successful operation
*/
type PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlOK struct {
	Payload *models.UpdateScheduleUploadResponse
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/update/uploadurl][%d] postWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlOK  %+v", 200, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlOK) GetPayload() *models.UpdateScheduleUploadResponse {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UpdateScheduleUploadResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlCreated creates a PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlCreated with default headers values
func NewPostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlCreated() *PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlCreated {
	return &PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlCreated{}
}

/*PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlCreated handles this case with default header values.

The upload url for schedule update was successfully created
*/
type PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlCreated struct {
	Payload *models.UpdateScheduleUploadResponse
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/update/uploadurl][%d] postWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlCreated  %+v", 201, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlCreated) GetPayload() *models.UpdateScheduleUploadResponse {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UpdateScheduleUploadResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlBadRequest creates a PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlBadRequest with default headers values
func NewPostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlBadRequest() *PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlBadRequest {
	return &PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlBadRequest{}
}

/*PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/update/uploadurl][%d] postWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlBadRequest  %+v", 400, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlUnauthorized creates a PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlUnauthorized with default headers values
func NewPostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlUnauthorized() *PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlUnauthorized {
	return &PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlUnauthorized{}
}

/*PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/update/uploadurl][%d] postWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlUnauthorized  %+v", 401, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlForbidden creates a PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlForbidden with default headers values
func NewPostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlForbidden() *PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlForbidden {
	return &PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlForbidden{}
}

/*PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/update/uploadurl][%d] postWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlForbidden  %+v", 403, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlNotFound creates a PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlNotFound with default headers values
func NewPostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlNotFound() *PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlNotFound {
	return &PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlNotFound{}
}

/*PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/update/uploadurl][%d] postWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlNotFound  %+v", 404, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlRequestTimeout creates a PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlRequestTimeout with default headers values
func NewPostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlRequestTimeout() *PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlRequestTimeout {
	return &PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlRequestTimeout{}
}

/*PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/update/uploadurl][%d] postWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlRequestEntityTooLarge creates a PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlRequestEntityTooLarge with default headers values
func NewPostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlRequestEntityTooLarge() *PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlRequestEntityTooLarge {
	return &PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlRequestEntityTooLarge{}
}

/*PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/update/uploadurl][%d] postWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlUnsupportedMediaType creates a PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlUnsupportedMediaType with default headers values
func NewPostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlUnsupportedMediaType() *PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlUnsupportedMediaType {
	return &PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlUnsupportedMediaType{}
}

/*PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/update/uploadurl][%d] postWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlTooManyRequests creates a PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlTooManyRequests with default headers values
func NewPostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlTooManyRequests() *PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlTooManyRequests {
	return &PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlTooManyRequests{}
}

/*PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/update/uploadurl][%d] postWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlInternalServerError creates a PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlInternalServerError with default headers values
func NewPostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlInternalServerError() *PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlInternalServerError {
	return &PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlInternalServerError{}
}

/*PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/update/uploadurl][%d] postWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlInternalServerError  %+v", 500, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlServiceUnavailable creates a PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlServiceUnavailable with default headers values
func NewPostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlServiceUnavailable() *PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlServiceUnavailable {
	return &PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlServiceUnavailable{}
}

/*PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/update/uploadurl][%d] postWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlGatewayTimeout creates a PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlGatewayTimeout with default headers values
func NewPostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlGatewayTimeout() *PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlGatewayTimeout {
	return &PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlGatewayTimeout{}
}

/*PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/update/uploadurl][%d] postWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurlGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

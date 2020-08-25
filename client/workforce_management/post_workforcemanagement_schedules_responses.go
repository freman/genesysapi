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

// PostWorkforcemanagementSchedulesReader is a Reader for the PostWorkforcemanagementSchedules structure.
type PostWorkforcemanagementSchedulesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostWorkforcemanagementSchedulesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostWorkforcemanagementSchedulesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostWorkforcemanagementSchedulesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostWorkforcemanagementSchedulesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostWorkforcemanagementSchedulesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostWorkforcemanagementSchedulesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostWorkforcemanagementSchedulesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostWorkforcemanagementSchedulesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostWorkforcemanagementSchedulesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostWorkforcemanagementSchedulesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostWorkforcemanagementSchedulesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostWorkforcemanagementSchedulesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostWorkforcemanagementSchedulesOK creates a PostWorkforcemanagementSchedulesOK with default headers values
func NewPostWorkforcemanagementSchedulesOK() *PostWorkforcemanagementSchedulesOK {
	return &PostWorkforcemanagementSchedulesOK{}
}

/*PostWorkforcemanagementSchedulesOK handles this case with default header values.

successful operation
*/
type PostWorkforcemanagementSchedulesOK struct {
	Payload *models.UserScheduleContainer
}

func (o *PostWorkforcemanagementSchedulesOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/schedules][%d] postWorkforcemanagementSchedulesOK  %+v", 200, o.Payload)
}

func (o *PostWorkforcemanagementSchedulesOK) GetPayload() *models.UserScheduleContainer {
	return o.Payload
}

func (o *PostWorkforcemanagementSchedulesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserScheduleContainer)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementSchedulesBadRequest creates a PostWorkforcemanagementSchedulesBadRequest with default headers values
func NewPostWorkforcemanagementSchedulesBadRequest() *PostWorkforcemanagementSchedulesBadRequest {
	return &PostWorkforcemanagementSchedulesBadRequest{}
}

/*PostWorkforcemanagementSchedulesBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostWorkforcemanagementSchedulesBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementSchedulesBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/schedules][%d] postWorkforcemanagementSchedulesBadRequest  %+v", 400, o.Payload)
}

func (o *PostWorkforcemanagementSchedulesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementSchedulesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementSchedulesUnauthorized creates a PostWorkforcemanagementSchedulesUnauthorized with default headers values
func NewPostWorkforcemanagementSchedulesUnauthorized() *PostWorkforcemanagementSchedulesUnauthorized {
	return &PostWorkforcemanagementSchedulesUnauthorized{}
}

/*PostWorkforcemanagementSchedulesUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostWorkforcemanagementSchedulesUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementSchedulesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/schedules][%d] postWorkforcemanagementSchedulesUnauthorized  %+v", 401, o.Payload)
}

func (o *PostWorkforcemanagementSchedulesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementSchedulesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementSchedulesForbidden creates a PostWorkforcemanagementSchedulesForbidden with default headers values
func NewPostWorkforcemanagementSchedulesForbidden() *PostWorkforcemanagementSchedulesForbidden {
	return &PostWorkforcemanagementSchedulesForbidden{}
}

/*PostWorkforcemanagementSchedulesForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostWorkforcemanagementSchedulesForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementSchedulesForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/schedules][%d] postWorkforcemanagementSchedulesForbidden  %+v", 403, o.Payload)
}

func (o *PostWorkforcemanagementSchedulesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementSchedulesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementSchedulesNotFound creates a PostWorkforcemanagementSchedulesNotFound with default headers values
func NewPostWorkforcemanagementSchedulesNotFound() *PostWorkforcemanagementSchedulesNotFound {
	return &PostWorkforcemanagementSchedulesNotFound{}
}

/*PostWorkforcemanagementSchedulesNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostWorkforcemanagementSchedulesNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementSchedulesNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/schedules][%d] postWorkforcemanagementSchedulesNotFound  %+v", 404, o.Payload)
}

func (o *PostWorkforcemanagementSchedulesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementSchedulesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementSchedulesRequestEntityTooLarge creates a PostWorkforcemanagementSchedulesRequestEntityTooLarge with default headers values
func NewPostWorkforcemanagementSchedulesRequestEntityTooLarge() *PostWorkforcemanagementSchedulesRequestEntityTooLarge {
	return &PostWorkforcemanagementSchedulesRequestEntityTooLarge{}
}

/*PostWorkforcemanagementSchedulesRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostWorkforcemanagementSchedulesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementSchedulesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/schedules][%d] postWorkforcemanagementSchedulesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostWorkforcemanagementSchedulesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementSchedulesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementSchedulesUnsupportedMediaType creates a PostWorkforcemanagementSchedulesUnsupportedMediaType with default headers values
func NewPostWorkforcemanagementSchedulesUnsupportedMediaType() *PostWorkforcemanagementSchedulesUnsupportedMediaType {
	return &PostWorkforcemanagementSchedulesUnsupportedMediaType{}
}

/*PostWorkforcemanagementSchedulesUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostWorkforcemanagementSchedulesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementSchedulesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/schedules][%d] postWorkforcemanagementSchedulesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostWorkforcemanagementSchedulesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementSchedulesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementSchedulesTooManyRequests creates a PostWorkforcemanagementSchedulesTooManyRequests with default headers values
func NewPostWorkforcemanagementSchedulesTooManyRequests() *PostWorkforcemanagementSchedulesTooManyRequests {
	return &PostWorkforcemanagementSchedulesTooManyRequests{}
}

/*PostWorkforcemanagementSchedulesTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PostWorkforcemanagementSchedulesTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementSchedulesTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/schedules][%d] postWorkforcemanagementSchedulesTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostWorkforcemanagementSchedulesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementSchedulesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementSchedulesInternalServerError creates a PostWorkforcemanagementSchedulesInternalServerError with default headers values
func NewPostWorkforcemanagementSchedulesInternalServerError() *PostWorkforcemanagementSchedulesInternalServerError {
	return &PostWorkforcemanagementSchedulesInternalServerError{}
}

/*PostWorkforcemanagementSchedulesInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostWorkforcemanagementSchedulesInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementSchedulesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/schedules][%d] postWorkforcemanagementSchedulesInternalServerError  %+v", 500, o.Payload)
}

func (o *PostWorkforcemanagementSchedulesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementSchedulesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementSchedulesServiceUnavailable creates a PostWorkforcemanagementSchedulesServiceUnavailable with default headers values
func NewPostWorkforcemanagementSchedulesServiceUnavailable() *PostWorkforcemanagementSchedulesServiceUnavailable {
	return &PostWorkforcemanagementSchedulesServiceUnavailable{}
}

/*PostWorkforcemanagementSchedulesServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostWorkforcemanagementSchedulesServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementSchedulesServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/schedules][%d] postWorkforcemanagementSchedulesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostWorkforcemanagementSchedulesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementSchedulesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementSchedulesGatewayTimeout creates a PostWorkforcemanagementSchedulesGatewayTimeout with default headers values
func NewPostWorkforcemanagementSchedulesGatewayTimeout() *PostWorkforcemanagementSchedulesGatewayTimeout {
	return &PostWorkforcemanagementSchedulesGatewayTimeout{}
}

/*PostWorkforcemanagementSchedulesGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostWorkforcemanagementSchedulesGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementSchedulesGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/schedules][%d] postWorkforcemanagementSchedulesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostWorkforcemanagementSchedulesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementSchedulesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
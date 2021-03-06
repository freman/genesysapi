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

// PostWorkforcemanagementBusinessunitWeekSchedulesGenerateReader is a Reader for the PostWorkforcemanagementBusinessunitWeekSchedulesGenerate structure.
type PostWorkforcemanagementBusinessunitWeekSchedulesGenerateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostWorkforcemanagementBusinessunitWeekSchedulesGenerateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostWorkforcemanagementBusinessunitWeekSchedulesGenerateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostWorkforcemanagementBusinessunitWeekSchedulesGenerateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostWorkforcemanagementBusinessunitWeekSchedulesGenerateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostWorkforcemanagementBusinessunitWeekSchedulesGenerateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostWorkforcemanagementBusinessunitWeekSchedulesGenerateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostWorkforcemanagementBusinessunitWeekSchedulesGenerateRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostWorkforcemanagementBusinessunitWeekSchedulesGenerateUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostWorkforcemanagementBusinessunitWeekSchedulesGenerateTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostWorkforcemanagementBusinessunitWeekSchedulesGenerateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostWorkforcemanagementBusinessunitWeekSchedulesGenerateServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostWorkforcemanagementBusinessunitWeekSchedulesGenerateGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostWorkforcemanagementBusinessunitWeekSchedulesGenerateOK creates a PostWorkforcemanagementBusinessunitWeekSchedulesGenerateOK with default headers values
func NewPostWorkforcemanagementBusinessunitWeekSchedulesGenerateOK() *PostWorkforcemanagementBusinessunitWeekSchedulesGenerateOK {
	return &PostWorkforcemanagementBusinessunitWeekSchedulesGenerateOK{}
}

/*PostWorkforcemanagementBusinessunitWeekSchedulesGenerateOK handles this case with default header values.

successful operation
*/
type PostWorkforcemanagementBusinessunitWeekSchedulesGenerateOK struct {
	Payload *models.BuAsyncScheduleRunResponse
}

func (o *PostWorkforcemanagementBusinessunitWeekSchedulesGenerateOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/generate][%d] postWorkforcemanagementBusinessunitWeekSchedulesGenerateOK  %+v", 200, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitWeekSchedulesGenerateOK) GetPayload() *models.BuAsyncScheduleRunResponse {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitWeekSchedulesGenerateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BuAsyncScheduleRunResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitWeekSchedulesGenerateBadRequest creates a PostWorkforcemanagementBusinessunitWeekSchedulesGenerateBadRequest with default headers values
func NewPostWorkforcemanagementBusinessunitWeekSchedulesGenerateBadRequest() *PostWorkforcemanagementBusinessunitWeekSchedulesGenerateBadRequest {
	return &PostWorkforcemanagementBusinessunitWeekSchedulesGenerateBadRequest{}
}

/*PostWorkforcemanagementBusinessunitWeekSchedulesGenerateBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostWorkforcemanagementBusinessunitWeekSchedulesGenerateBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementBusinessunitWeekSchedulesGenerateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/generate][%d] postWorkforcemanagementBusinessunitWeekSchedulesGenerateBadRequest  %+v", 400, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitWeekSchedulesGenerateBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitWeekSchedulesGenerateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitWeekSchedulesGenerateUnauthorized creates a PostWorkforcemanagementBusinessunitWeekSchedulesGenerateUnauthorized with default headers values
func NewPostWorkforcemanagementBusinessunitWeekSchedulesGenerateUnauthorized() *PostWorkforcemanagementBusinessunitWeekSchedulesGenerateUnauthorized {
	return &PostWorkforcemanagementBusinessunitWeekSchedulesGenerateUnauthorized{}
}

/*PostWorkforcemanagementBusinessunitWeekSchedulesGenerateUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostWorkforcemanagementBusinessunitWeekSchedulesGenerateUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementBusinessunitWeekSchedulesGenerateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/generate][%d] postWorkforcemanagementBusinessunitWeekSchedulesGenerateUnauthorized  %+v", 401, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitWeekSchedulesGenerateUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitWeekSchedulesGenerateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitWeekSchedulesGenerateForbidden creates a PostWorkforcemanagementBusinessunitWeekSchedulesGenerateForbidden with default headers values
func NewPostWorkforcemanagementBusinessunitWeekSchedulesGenerateForbidden() *PostWorkforcemanagementBusinessunitWeekSchedulesGenerateForbidden {
	return &PostWorkforcemanagementBusinessunitWeekSchedulesGenerateForbidden{}
}

/*PostWorkforcemanagementBusinessunitWeekSchedulesGenerateForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostWorkforcemanagementBusinessunitWeekSchedulesGenerateForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementBusinessunitWeekSchedulesGenerateForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/generate][%d] postWorkforcemanagementBusinessunitWeekSchedulesGenerateForbidden  %+v", 403, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitWeekSchedulesGenerateForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitWeekSchedulesGenerateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitWeekSchedulesGenerateNotFound creates a PostWorkforcemanagementBusinessunitWeekSchedulesGenerateNotFound with default headers values
func NewPostWorkforcemanagementBusinessunitWeekSchedulesGenerateNotFound() *PostWorkforcemanagementBusinessunitWeekSchedulesGenerateNotFound {
	return &PostWorkforcemanagementBusinessunitWeekSchedulesGenerateNotFound{}
}

/*PostWorkforcemanagementBusinessunitWeekSchedulesGenerateNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostWorkforcemanagementBusinessunitWeekSchedulesGenerateNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementBusinessunitWeekSchedulesGenerateNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/generate][%d] postWorkforcemanagementBusinessunitWeekSchedulesGenerateNotFound  %+v", 404, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitWeekSchedulesGenerateNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitWeekSchedulesGenerateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitWeekSchedulesGenerateRequestEntityTooLarge creates a PostWorkforcemanagementBusinessunitWeekSchedulesGenerateRequestEntityTooLarge with default headers values
func NewPostWorkforcemanagementBusinessunitWeekSchedulesGenerateRequestEntityTooLarge() *PostWorkforcemanagementBusinessunitWeekSchedulesGenerateRequestEntityTooLarge {
	return &PostWorkforcemanagementBusinessunitWeekSchedulesGenerateRequestEntityTooLarge{}
}

/*PostWorkforcemanagementBusinessunitWeekSchedulesGenerateRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostWorkforcemanagementBusinessunitWeekSchedulesGenerateRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementBusinessunitWeekSchedulesGenerateRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/generate][%d] postWorkforcemanagementBusinessunitWeekSchedulesGenerateRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitWeekSchedulesGenerateRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitWeekSchedulesGenerateRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitWeekSchedulesGenerateUnsupportedMediaType creates a PostWorkforcemanagementBusinessunitWeekSchedulesGenerateUnsupportedMediaType with default headers values
func NewPostWorkforcemanagementBusinessunitWeekSchedulesGenerateUnsupportedMediaType() *PostWorkforcemanagementBusinessunitWeekSchedulesGenerateUnsupportedMediaType {
	return &PostWorkforcemanagementBusinessunitWeekSchedulesGenerateUnsupportedMediaType{}
}

/*PostWorkforcemanagementBusinessunitWeekSchedulesGenerateUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostWorkforcemanagementBusinessunitWeekSchedulesGenerateUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementBusinessunitWeekSchedulesGenerateUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/generate][%d] postWorkforcemanagementBusinessunitWeekSchedulesGenerateUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitWeekSchedulesGenerateUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitWeekSchedulesGenerateUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitWeekSchedulesGenerateTooManyRequests creates a PostWorkforcemanagementBusinessunitWeekSchedulesGenerateTooManyRequests with default headers values
func NewPostWorkforcemanagementBusinessunitWeekSchedulesGenerateTooManyRequests() *PostWorkforcemanagementBusinessunitWeekSchedulesGenerateTooManyRequests {
	return &PostWorkforcemanagementBusinessunitWeekSchedulesGenerateTooManyRequests{}
}

/*PostWorkforcemanagementBusinessunitWeekSchedulesGenerateTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PostWorkforcemanagementBusinessunitWeekSchedulesGenerateTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementBusinessunitWeekSchedulesGenerateTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/generate][%d] postWorkforcemanagementBusinessunitWeekSchedulesGenerateTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitWeekSchedulesGenerateTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitWeekSchedulesGenerateTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitWeekSchedulesGenerateInternalServerError creates a PostWorkforcemanagementBusinessunitWeekSchedulesGenerateInternalServerError with default headers values
func NewPostWorkforcemanagementBusinessunitWeekSchedulesGenerateInternalServerError() *PostWorkforcemanagementBusinessunitWeekSchedulesGenerateInternalServerError {
	return &PostWorkforcemanagementBusinessunitWeekSchedulesGenerateInternalServerError{}
}

/*PostWorkforcemanagementBusinessunitWeekSchedulesGenerateInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostWorkforcemanagementBusinessunitWeekSchedulesGenerateInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementBusinessunitWeekSchedulesGenerateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/generate][%d] postWorkforcemanagementBusinessunitWeekSchedulesGenerateInternalServerError  %+v", 500, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitWeekSchedulesGenerateInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitWeekSchedulesGenerateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitWeekSchedulesGenerateServiceUnavailable creates a PostWorkforcemanagementBusinessunitWeekSchedulesGenerateServiceUnavailable with default headers values
func NewPostWorkforcemanagementBusinessunitWeekSchedulesGenerateServiceUnavailable() *PostWorkforcemanagementBusinessunitWeekSchedulesGenerateServiceUnavailable {
	return &PostWorkforcemanagementBusinessunitWeekSchedulesGenerateServiceUnavailable{}
}

/*PostWorkforcemanagementBusinessunitWeekSchedulesGenerateServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostWorkforcemanagementBusinessunitWeekSchedulesGenerateServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementBusinessunitWeekSchedulesGenerateServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/generate][%d] postWorkforcemanagementBusinessunitWeekSchedulesGenerateServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitWeekSchedulesGenerateServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitWeekSchedulesGenerateServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitWeekSchedulesGenerateGatewayTimeout creates a PostWorkforcemanagementBusinessunitWeekSchedulesGenerateGatewayTimeout with default headers values
func NewPostWorkforcemanagementBusinessunitWeekSchedulesGenerateGatewayTimeout() *PostWorkforcemanagementBusinessunitWeekSchedulesGenerateGatewayTimeout {
	return &PostWorkforcemanagementBusinessunitWeekSchedulesGenerateGatewayTimeout{}
}

/*PostWorkforcemanagementBusinessunitWeekSchedulesGenerateGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostWorkforcemanagementBusinessunitWeekSchedulesGenerateGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementBusinessunitWeekSchedulesGenerateGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/generate][%d] postWorkforcemanagementBusinessunitWeekSchedulesGenerateGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitWeekSchedulesGenerateGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitWeekSchedulesGenerateGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

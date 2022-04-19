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

// PostWorkforcemanagementBusinessunitWeekSchedulesImportReader is a Reader for the PostWorkforcemanagementBusinessunitWeekSchedulesImport structure.
type PostWorkforcemanagementBusinessunitWeekSchedulesImportReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostWorkforcemanagementBusinessunitWeekSchedulesImportReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostWorkforcemanagementBusinessunitWeekSchedulesImportOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewPostWorkforcemanagementBusinessunitWeekSchedulesImportAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostWorkforcemanagementBusinessunitWeekSchedulesImportBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostWorkforcemanagementBusinessunitWeekSchedulesImportUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostWorkforcemanagementBusinessunitWeekSchedulesImportForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostWorkforcemanagementBusinessunitWeekSchedulesImportNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostWorkforcemanagementBusinessunitWeekSchedulesImportRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostWorkforcemanagementBusinessunitWeekSchedulesImportRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostWorkforcemanagementBusinessunitWeekSchedulesImportUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostWorkforcemanagementBusinessunitWeekSchedulesImportTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostWorkforcemanagementBusinessunitWeekSchedulesImportInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostWorkforcemanagementBusinessunitWeekSchedulesImportServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostWorkforcemanagementBusinessunitWeekSchedulesImportGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostWorkforcemanagementBusinessunitWeekSchedulesImportOK creates a PostWorkforcemanagementBusinessunitWeekSchedulesImportOK with default headers values
func NewPostWorkforcemanagementBusinessunitWeekSchedulesImportOK() *PostWorkforcemanagementBusinessunitWeekSchedulesImportOK {
	return &PostWorkforcemanagementBusinessunitWeekSchedulesImportOK{}
}

/*PostWorkforcemanagementBusinessunitWeekSchedulesImportOK handles this case with default header values.

successful operation
*/
type PostWorkforcemanagementBusinessunitWeekSchedulesImportOK struct {
	Payload *models.ScheduleUploadProcessingResponse
}

func (o *PostWorkforcemanagementBusinessunitWeekSchedulesImportOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/import][%d] postWorkforcemanagementBusinessunitWeekSchedulesImportOK  %+v", 200, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitWeekSchedulesImportOK) GetPayload() *models.ScheduleUploadProcessingResponse {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitWeekSchedulesImportOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ScheduleUploadProcessingResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitWeekSchedulesImportAccepted creates a PostWorkforcemanagementBusinessunitWeekSchedulesImportAccepted with default headers values
func NewPostWorkforcemanagementBusinessunitWeekSchedulesImportAccepted() *PostWorkforcemanagementBusinessunitWeekSchedulesImportAccepted {
	return &PostWorkforcemanagementBusinessunitWeekSchedulesImportAccepted{}
}

/*PostWorkforcemanagementBusinessunitWeekSchedulesImportAccepted handles this case with default header values.

The schedule import is processing
*/
type PostWorkforcemanagementBusinessunitWeekSchedulesImportAccepted struct {
	Payload *models.ScheduleUploadProcessingResponse
}

func (o *PostWorkforcemanagementBusinessunitWeekSchedulesImportAccepted) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/import][%d] postWorkforcemanagementBusinessunitWeekSchedulesImportAccepted  %+v", 202, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitWeekSchedulesImportAccepted) GetPayload() *models.ScheduleUploadProcessingResponse {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitWeekSchedulesImportAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ScheduleUploadProcessingResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitWeekSchedulesImportBadRequest creates a PostWorkforcemanagementBusinessunitWeekSchedulesImportBadRequest with default headers values
func NewPostWorkforcemanagementBusinessunitWeekSchedulesImportBadRequest() *PostWorkforcemanagementBusinessunitWeekSchedulesImportBadRequest {
	return &PostWorkforcemanagementBusinessunitWeekSchedulesImportBadRequest{}
}

/*PostWorkforcemanagementBusinessunitWeekSchedulesImportBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostWorkforcemanagementBusinessunitWeekSchedulesImportBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementBusinessunitWeekSchedulesImportBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/import][%d] postWorkforcemanagementBusinessunitWeekSchedulesImportBadRequest  %+v", 400, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitWeekSchedulesImportBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitWeekSchedulesImportBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitWeekSchedulesImportUnauthorized creates a PostWorkforcemanagementBusinessunitWeekSchedulesImportUnauthorized with default headers values
func NewPostWorkforcemanagementBusinessunitWeekSchedulesImportUnauthorized() *PostWorkforcemanagementBusinessunitWeekSchedulesImportUnauthorized {
	return &PostWorkforcemanagementBusinessunitWeekSchedulesImportUnauthorized{}
}

/*PostWorkforcemanagementBusinessunitWeekSchedulesImportUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostWorkforcemanagementBusinessunitWeekSchedulesImportUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementBusinessunitWeekSchedulesImportUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/import][%d] postWorkforcemanagementBusinessunitWeekSchedulesImportUnauthorized  %+v", 401, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitWeekSchedulesImportUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitWeekSchedulesImportUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitWeekSchedulesImportForbidden creates a PostWorkforcemanagementBusinessunitWeekSchedulesImportForbidden with default headers values
func NewPostWorkforcemanagementBusinessunitWeekSchedulesImportForbidden() *PostWorkforcemanagementBusinessunitWeekSchedulesImportForbidden {
	return &PostWorkforcemanagementBusinessunitWeekSchedulesImportForbidden{}
}

/*PostWorkforcemanagementBusinessunitWeekSchedulesImportForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostWorkforcemanagementBusinessunitWeekSchedulesImportForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementBusinessunitWeekSchedulesImportForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/import][%d] postWorkforcemanagementBusinessunitWeekSchedulesImportForbidden  %+v", 403, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitWeekSchedulesImportForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitWeekSchedulesImportForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitWeekSchedulesImportNotFound creates a PostWorkforcemanagementBusinessunitWeekSchedulesImportNotFound with default headers values
func NewPostWorkforcemanagementBusinessunitWeekSchedulesImportNotFound() *PostWorkforcemanagementBusinessunitWeekSchedulesImportNotFound {
	return &PostWorkforcemanagementBusinessunitWeekSchedulesImportNotFound{}
}

/*PostWorkforcemanagementBusinessunitWeekSchedulesImportNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostWorkforcemanagementBusinessunitWeekSchedulesImportNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementBusinessunitWeekSchedulesImportNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/import][%d] postWorkforcemanagementBusinessunitWeekSchedulesImportNotFound  %+v", 404, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitWeekSchedulesImportNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitWeekSchedulesImportNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitWeekSchedulesImportRequestTimeout creates a PostWorkforcemanagementBusinessunitWeekSchedulesImportRequestTimeout with default headers values
func NewPostWorkforcemanagementBusinessunitWeekSchedulesImportRequestTimeout() *PostWorkforcemanagementBusinessunitWeekSchedulesImportRequestTimeout {
	return &PostWorkforcemanagementBusinessunitWeekSchedulesImportRequestTimeout{}
}

/*PostWorkforcemanagementBusinessunitWeekSchedulesImportRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostWorkforcemanagementBusinessunitWeekSchedulesImportRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementBusinessunitWeekSchedulesImportRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/import][%d] postWorkforcemanagementBusinessunitWeekSchedulesImportRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitWeekSchedulesImportRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitWeekSchedulesImportRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitWeekSchedulesImportRequestEntityTooLarge creates a PostWorkforcemanagementBusinessunitWeekSchedulesImportRequestEntityTooLarge with default headers values
func NewPostWorkforcemanagementBusinessunitWeekSchedulesImportRequestEntityTooLarge() *PostWorkforcemanagementBusinessunitWeekSchedulesImportRequestEntityTooLarge {
	return &PostWorkforcemanagementBusinessunitWeekSchedulesImportRequestEntityTooLarge{}
}

/*PostWorkforcemanagementBusinessunitWeekSchedulesImportRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PostWorkforcemanagementBusinessunitWeekSchedulesImportRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementBusinessunitWeekSchedulesImportRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/import][%d] postWorkforcemanagementBusinessunitWeekSchedulesImportRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitWeekSchedulesImportRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitWeekSchedulesImportRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitWeekSchedulesImportUnsupportedMediaType creates a PostWorkforcemanagementBusinessunitWeekSchedulesImportUnsupportedMediaType with default headers values
func NewPostWorkforcemanagementBusinessunitWeekSchedulesImportUnsupportedMediaType() *PostWorkforcemanagementBusinessunitWeekSchedulesImportUnsupportedMediaType {
	return &PostWorkforcemanagementBusinessunitWeekSchedulesImportUnsupportedMediaType{}
}

/*PostWorkforcemanagementBusinessunitWeekSchedulesImportUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostWorkforcemanagementBusinessunitWeekSchedulesImportUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementBusinessunitWeekSchedulesImportUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/import][%d] postWorkforcemanagementBusinessunitWeekSchedulesImportUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitWeekSchedulesImportUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitWeekSchedulesImportUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitWeekSchedulesImportTooManyRequests creates a PostWorkforcemanagementBusinessunitWeekSchedulesImportTooManyRequests with default headers values
func NewPostWorkforcemanagementBusinessunitWeekSchedulesImportTooManyRequests() *PostWorkforcemanagementBusinessunitWeekSchedulesImportTooManyRequests {
	return &PostWorkforcemanagementBusinessunitWeekSchedulesImportTooManyRequests{}
}

/*PostWorkforcemanagementBusinessunitWeekSchedulesImportTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostWorkforcemanagementBusinessunitWeekSchedulesImportTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementBusinessunitWeekSchedulesImportTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/import][%d] postWorkforcemanagementBusinessunitWeekSchedulesImportTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitWeekSchedulesImportTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitWeekSchedulesImportTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitWeekSchedulesImportInternalServerError creates a PostWorkforcemanagementBusinessunitWeekSchedulesImportInternalServerError with default headers values
func NewPostWorkforcemanagementBusinessunitWeekSchedulesImportInternalServerError() *PostWorkforcemanagementBusinessunitWeekSchedulesImportInternalServerError {
	return &PostWorkforcemanagementBusinessunitWeekSchedulesImportInternalServerError{}
}

/*PostWorkforcemanagementBusinessunitWeekSchedulesImportInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostWorkforcemanagementBusinessunitWeekSchedulesImportInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementBusinessunitWeekSchedulesImportInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/import][%d] postWorkforcemanagementBusinessunitWeekSchedulesImportInternalServerError  %+v", 500, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitWeekSchedulesImportInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitWeekSchedulesImportInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitWeekSchedulesImportServiceUnavailable creates a PostWorkforcemanagementBusinessunitWeekSchedulesImportServiceUnavailable with default headers values
func NewPostWorkforcemanagementBusinessunitWeekSchedulesImportServiceUnavailable() *PostWorkforcemanagementBusinessunitWeekSchedulesImportServiceUnavailable {
	return &PostWorkforcemanagementBusinessunitWeekSchedulesImportServiceUnavailable{}
}

/*PostWorkforcemanagementBusinessunitWeekSchedulesImportServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostWorkforcemanagementBusinessunitWeekSchedulesImportServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementBusinessunitWeekSchedulesImportServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/import][%d] postWorkforcemanagementBusinessunitWeekSchedulesImportServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitWeekSchedulesImportServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitWeekSchedulesImportServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitWeekSchedulesImportGatewayTimeout creates a PostWorkforcemanagementBusinessunitWeekSchedulesImportGatewayTimeout with default headers values
func NewPostWorkforcemanagementBusinessunitWeekSchedulesImportGatewayTimeout() *PostWorkforcemanagementBusinessunitWeekSchedulesImportGatewayTimeout {
	return &PostWorkforcemanagementBusinessunitWeekSchedulesImportGatewayTimeout{}
}

/*PostWorkforcemanagementBusinessunitWeekSchedulesImportGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostWorkforcemanagementBusinessunitWeekSchedulesImportGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementBusinessunitWeekSchedulesImportGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/import][%d] postWorkforcemanagementBusinessunitWeekSchedulesImportGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitWeekSchedulesImportGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitWeekSchedulesImportGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

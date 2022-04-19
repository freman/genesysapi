// Code generated by go-swagger; DO NOT EDIT.

package recording

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PatchRecordingsScreensessionReader is a Reader for the PatchRecordingsScreensession structure.
type PatchRecordingsScreensessionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchRecordingsScreensessionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 400:
		result := NewPatchRecordingsScreensessionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchRecordingsScreensessionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchRecordingsScreensessionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchRecordingsScreensessionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPatchRecordingsScreensessionRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPatchRecordingsScreensessionRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPatchRecordingsScreensessionUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPatchRecordingsScreensessionTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchRecordingsScreensessionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPatchRecordingsScreensessionServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPatchRecordingsScreensessionGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPatchRecordingsScreensessionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchRecordingsScreensessionBadRequest creates a PatchRecordingsScreensessionBadRequest with default headers values
func NewPatchRecordingsScreensessionBadRequest() *PatchRecordingsScreensessionBadRequest {
	return &PatchRecordingsScreensessionBadRequest{}
}

/*PatchRecordingsScreensessionBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PatchRecordingsScreensessionBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PatchRecordingsScreensessionBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/recordings/screensessions/{recordingSessionId}][%d] patchRecordingsScreensessionBadRequest  %+v", 400, o.Payload)
}

func (o *PatchRecordingsScreensessionBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRecordingsScreensessionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRecordingsScreensessionUnauthorized creates a PatchRecordingsScreensessionUnauthorized with default headers values
func NewPatchRecordingsScreensessionUnauthorized() *PatchRecordingsScreensessionUnauthorized {
	return &PatchRecordingsScreensessionUnauthorized{}
}

/*PatchRecordingsScreensessionUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PatchRecordingsScreensessionUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PatchRecordingsScreensessionUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/recordings/screensessions/{recordingSessionId}][%d] patchRecordingsScreensessionUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchRecordingsScreensessionUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRecordingsScreensessionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRecordingsScreensessionForbidden creates a PatchRecordingsScreensessionForbidden with default headers values
func NewPatchRecordingsScreensessionForbidden() *PatchRecordingsScreensessionForbidden {
	return &PatchRecordingsScreensessionForbidden{}
}

/*PatchRecordingsScreensessionForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PatchRecordingsScreensessionForbidden struct {
	Payload *models.ErrorBody
}

func (o *PatchRecordingsScreensessionForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/recordings/screensessions/{recordingSessionId}][%d] patchRecordingsScreensessionForbidden  %+v", 403, o.Payload)
}

func (o *PatchRecordingsScreensessionForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRecordingsScreensessionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRecordingsScreensessionNotFound creates a PatchRecordingsScreensessionNotFound with default headers values
func NewPatchRecordingsScreensessionNotFound() *PatchRecordingsScreensessionNotFound {
	return &PatchRecordingsScreensessionNotFound{}
}

/*PatchRecordingsScreensessionNotFound handles this case with default header values.

The requested resource was not found.
*/
type PatchRecordingsScreensessionNotFound struct {
	Payload *models.ErrorBody
}

func (o *PatchRecordingsScreensessionNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/recordings/screensessions/{recordingSessionId}][%d] patchRecordingsScreensessionNotFound  %+v", 404, o.Payload)
}

func (o *PatchRecordingsScreensessionNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRecordingsScreensessionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRecordingsScreensessionRequestTimeout creates a PatchRecordingsScreensessionRequestTimeout with default headers values
func NewPatchRecordingsScreensessionRequestTimeout() *PatchRecordingsScreensessionRequestTimeout {
	return &PatchRecordingsScreensessionRequestTimeout{}
}

/*PatchRecordingsScreensessionRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PatchRecordingsScreensessionRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PatchRecordingsScreensessionRequestTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/recordings/screensessions/{recordingSessionId}][%d] patchRecordingsScreensessionRequestTimeout  %+v", 408, o.Payload)
}

func (o *PatchRecordingsScreensessionRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRecordingsScreensessionRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRecordingsScreensessionRequestEntityTooLarge creates a PatchRecordingsScreensessionRequestEntityTooLarge with default headers values
func NewPatchRecordingsScreensessionRequestEntityTooLarge() *PatchRecordingsScreensessionRequestEntityTooLarge {
	return &PatchRecordingsScreensessionRequestEntityTooLarge{}
}

/*PatchRecordingsScreensessionRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PatchRecordingsScreensessionRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PatchRecordingsScreensessionRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/recordings/screensessions/{recordingSessionId}][%d] patchRecordingsScreensessionRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchRecordingsScreensessionRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRecordingsScreensessionRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRecordingsScreensessionUnsupportedMediaType creates a PatchRecordingsScreensessionUnsupportedMediaType with default headers values
func NewPatchRecordingsScreensessionUnsupportedMediaType() *PatchRecordingsScreensessionUnsupportedMediaType {
	return &PatchRecordingsScreensessionUnsupportedMediaType{}
}

/*PatchRecordingsScreensessionUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PatchRecordingsScreensessionUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PatchRecordingsScreensessionUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/recordings/screensessions/{recordingSessionId}][%d] patchRecordingsScreensessionUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchRecordingsScreensessionUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRecordingsScreensessionUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRecordingsScreensessionTooManyRequests creates a PatchRecordingsScreensessionTooManyRequests with default headers values
func NewPatchRecordingsScreensessionTooManyRequests() *PatchRecordingsScreensessionTooManyRequests {
	return &PatchRecordingsScreensessionTooManyRequests{}
}

/*PatchRecordingsScreensessionTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PatchRecordingsScreensessionTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PatchRecordingsScreensessionTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/recordings/screensessions/{recordingSessionId}][%d] patchRecordingsScreensessionTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchRecordingsScreensessionTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRecordingsScreensessionTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRecordingsScreensessionInternalServerError creates a PatchRecordingsScreensessionInternalServerError with default headers values
func NewPatchRecordingsScreensessionInternalServerError() *PatchRecordingsScreensessionInternalServerError {
	return &PatchRecordingsScreensessionInternalServerError{}
}

/*PatchRecordingsScreensessionInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PatchRecordingsScreensessionInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PatchRecordingsScreensessionInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/recordings/screensessions/{recordingSessionId}][%d] patchRecordingsScreensessionInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchRecordingsScreensessionInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRecordingsScreensessionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRecordingsScreensessionServiceUnavailable creates a PatchRecordingsScreensessionServiceUnavailable with default headers values
func NewPatchRecordingsScreensessionServiceUnavailable() *PatchRecordingsScreensessionServiceUnavailable {
	return &PatchRecordingsScreensessionServiceUnavailable{}
}

/*PatchRecordingsScreensessionServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PatchRecordingsScreensessionServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PatchRecordingsScreensessionServiceUnavailable) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/recordings/screensessions/{recordingSessionId}][%d] patchRecordingsScreensessionServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchRecordingsScreensessionServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRecordingsScreensessionServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRecordingsScreensessionGatewayTimeout creates a PatchRecordingsScreensessionGatewayTimeout with default headers values
func NewPatchRecordingsScreensessionGatewayTimeout() *PatchRecordingsScreensessionGatewayTimeout {
	return &PatchRecordingsScreensessionGatewayTimeout{}
}

/*PatchRecordingsScreensessionGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PatchRecordingsScreensessionGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PatchRecordingsScreensessionGatewayTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/recordings/screensessions/{recordingSessionId}][%d] patchRecordingsScreensessionGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchRecordingsScreensessionGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRecordingsScreensessionGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRecordingsScreensessionDefault creates a PatchRecordingsScreensessionDefault with default headers values
func NewPatchRecordingsScreensessionDefault(code int) *PatchRecordingsScreensessionDefault {
	return &PatchRecordingsScreensessionDefault{
		_statusCode: code,
	}
}

/*PatchRecordingsScreensessionDefault handles this case with default header values.

successful operation
*/
type PatchRecordingsScreensessionDefault struct {
	_statusCode int
}

// Code gets the status code for the patch recordings screensession default response
func (o *PatchRecordingsScreensessionDefault) Code() int {
	return o._statusCode
}

func (o *PatchRecordingsScreensessionDefault) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/recordings/screensessions/{recordingSessionId}][%d] patchRecordingsScreensession default ", o._statusCode)
}

func (o *PatchRecordingsScreensessionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

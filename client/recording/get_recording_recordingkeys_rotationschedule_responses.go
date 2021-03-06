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

// GetRecordingRecordingkeysRotationscheduleReader is a Reader for the GetRecordingRecordingkeysRotationschedule structure.
type GetRecordingRecordingkeysRotationscheduleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRecordingRecordingkeysRotationscheduleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRecordingRecordingkeysRotationscheduleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRecordingRecordingkeysRotationscheduleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetRecordingRecordingkeysRotationscheduleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRecordingRecordingkeysRotationscheduleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRecordingRecordingkeysRotationscheduleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetRecordingRecordingkeysRotationscheduleRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetRecordingRecordingkeysRotationscheduleUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetRecordingRecordingkeysRotationscheduleTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRecordingRecordingkeysRotationscheduleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetRecordingRecordingkeysRotationscheduleServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetRecordingRecordingkeysRotationscheduleGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRecordingRecordingkeysRotationscheduleOK creates a GetRecordingRecordingkeysRotationscheduleOK with default headers values
func NewGetRecordingRecordingkeysRotationscheduleOK() *GetRecordingRecordingkeysRotationscheduleOK {
	return &GetRecordingRecordingkeysRotationscheduleOK{}
}

/*GetRecordingRecordingkeysRotationscheduleOK handles this case with default header values.

successful operation
*/
type GetRecordingRecordingkeysRotationscheduleOK struct {
	Payload *models.KeyRotationSchedule
}

func (o *GetRecordingRecordingkeysRotationscheduleOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/recordingkeys/rotationschedule][%d] getRecordingRecordingkeysRotationscheduleOK  %+v", 200, o.Payload)
}

func (o *GetRecordingRecordingkeysRotationscheduleOK) GetPayload() *models.KeyRotationSchedule {
	return o.Payload
}

func (o *GetRecordingRecordingkeysRotationscheduleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.KeyRotationSchedule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingRecordingkeysRotationscheduleBadRequest creates a GetRecordingRecordingkeysRotationscheduleBadRequest with default headers values
func NewGetRecordingRecordingkeysRotationscheduleBadRequest() *GetRecordingRecordingkeysRotationscheduleBadRequest {
	return &GetRecordingRecordingkeysRotationscheduleBadRequest{}
}

/*GetRecordingRecordingkeysRotationscheduleBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetRecordingRecordingkeysRotationscheduleBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetRecordingRecordingkeysRotationscheduleBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/recordingkeys/rotationschedule][%d] getRecordingRecordingkeysRotationscheduleBadRequest  %+v", 400, o.Payload)
}

func (o *GetRecordingRecordingkeysRotationscheduleBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingRecordingkeysRotationscheduleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingRecordingkeysRotationscheduleUnauthorized creates a GetRecordingRecordingkeysRotationscheduleUnauthorized with default headers values
func NewGetRecordingRecordingkeysRotationscheduleUnauthorized() *GetRecordingRecordingkeysRotationscheduleUnauthorized {
	return &GetRecordingRecordingkeysRotationscheduleUnauthorized{}
}

/*GetRecordingRecordingkeysRotationscheduleUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetRecordingRecordingkeysRotationscheduleUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetRecordingRecordingkeysRotationscheduleUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/recordingkeys/rotationschedule][%d] getRecordingRecordingkeysRotationscheduleUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRecordingRecordingkeysRotationscheduleUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingRecordingkeysRotationscheduleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingRecordingkeysRotationscheduleForbidden creates a GetRecordingRecordingkeysRotationscheduleForbidden with default headers values
func NewGetRecordingRecordingkeysRotationscheduleForbidden() *GetRecordingRecordingkeysRotationscheduleForbidden {
	return &GetRecordingRecordingkeysRotationscheduleForbidden{}
}

/*GetRecordingRecordingkeysRotationscheduleForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetRecordingRecordingkeysRotationscheduleForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetRecordingRecordingkeysRotationscheduleForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/recordingkeys/rotationschedule][%d] getRecordingRecordingkeysRotationscheduleForbidden  %+v", 403, o.Payload)
}

func (o *GetRecordingRecordingkeysRotationscheduleForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingRecordingkeysRotationscheduleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingRecordingkeysRotationscheduleNotFound creates a GetRecordingRecordingkeysRotationscheduleNotFound with default headers values
func NewGetRecordingRecordingkeysRotationscheduleNotFound() *GetRecordingRecordingkeysRotationscheduleNotFound {
	return &GetRecordingRecordingkeysRotationscheduleNotFound{}
}

/*GetRecordingRecordingkeysRotationscheduleNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetRecordingRecordingkeysRotationscheduleNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetRecordingRecordingkeysRotationscheduleNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/recordingkeys/rotationschedule][%d] getRecordingRecordingkeysRotationscheduleNotFound  %+v", 404, o.Payload)
}

func (o *GetRecordingRecordingkeysRotationscheduleNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingRecordingkeysRotationscheduleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingRecordingkeysRotationscheduleRequestEntityTooLarge creates a GetRecordingRecordingkeysRotationscheduleRequestEntityTooLarge with default headers values
func NewGetRecordingRecordingkeysRotationscheduleRequestEntityTooLarge() *GetRecordingRecordingkeysRotationscheduleRequestEntityTooLarge {
	return &GetRecordingRecordingkeysRotationscheduleRequestEntityTooLarge{}
}

/*GetRecordingRecordingkeysRotationscheduleRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetRecordingRecordingkeysRotationscheduleRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetRecordingRecordingkeysRotationscheduleRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/recordingkeys/rotationschedule][%d] getRecordingRecordingkeysRotationscheduleRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetRecordingRecordingkeysRotationscheduleRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingRecordingkeysRotationscheduleRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingRecordingkeysRotationscheduleUnsupportedMediaType creates a GetRecordingRecordingkeysRotationscheduleUnsupportedMediaType with default headers values
func NewGetRecordingRecordingkeysRotationscheduleUnsupportedMediaType() *GetRecordingRecordingkeysRotationscheduleUnsupportedMediaType {
	return &GetRecordingRecordingkeysRotationscheduleUnsupportedMediaType{}
}

/*GetRecordingRecordingkeysRotationscheduleUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetRecordingRecordingkeysRotationscheduleUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetRecordingRecordingkeysRotationscheduleUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/recordingkeys/rotationschedule][%d] getRecordingRecordingkeysRotationscheduleUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetRecordingRecordingkeysRotationscheduleUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingRecordingkeysRotationscheduleUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingRecordingkeysRotationscheduleTooManyRequests creates a GetRecordingRecordingkeysRotationscheduleTooManyRequests with default headers values
func NewGetRecordingRecordingkeysRotationscheduleTooManyRequests() *GetRecordingRecordingkeysRotationscheduleTooManyRequests {
	return &GetRecordingRecordingkeysRotationscheduleTooManyRequests{}
}

/*GetRecordingRecordingkeysRotationscheduleTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetRecordingRecordingkeysRotationscheduleTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetRecordingRecordingkeysRotationscheduleTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/recordingkeys/rotationschedule][%d] getRecordingRecordingkeysRotationscheduleTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetRecordingRecordingkeysRotationscheduleTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingRecordingkeysRotationscheduleTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingRecordingkeysRotationscheduleInternalServerError creates a GetRecordingRecordingkeysRotationscheduleInternalServerError with default headers values
func NewGetRecordingRecordingkeysRotationscheduleInternalServerError() *GetRecordingRecordingkeysRotationscheduleInternalServerError {
	return &GetRecordingRecordingkeysRotationscheduleInternalServerError{}
}

/*GetRecordingRecordingkeysRotationscheduleInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetRecordingRecordingkeysRotationscheduleInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetRecordingRecordingkeysRotationscheduleInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/recordingkeys/rotationschedule][%d] getRecordingRecordingkeysRotationscheduleInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRecordingRecordingkeysRotationscheduleInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingRecordingkeysRotationscheduleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingRecordingkeysRotationscheduleServiceUnavailable creates a GetRecordingRecordingkeysRotationscheduleServiceUnavailable with default headers values
func NewGetRecordingRecordingkeysRotationscheduleServiceUnavailable() *GetRecordingRecordingkeysRotationscheduleServiceUnavailable {
	return &GetRecordingRecordingkeysRotationscheduleServiceUnavailable{}
}

/*GetRecordingRecordingkeysRotationscheduleServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetRecordingRecordingkeysRotationscheduleServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetRecordingRecordingkeysRotationscheduleServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/recordingkeys/rotationschedule][%d] getRecordingRecordingkeysRotationscheduleServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetRecordingRecordingkeysRotationscheduleServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingRecordingkeysRotationscheduleServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingRecordingkeysRotationscheduleGatewayTimeout creates a GetRecordingRecordingkeysRotationscheduleGatewayTimeout with default headers values
func NewGetRecordingRecordingkeysRotationscheduleGatewayTimeout() *GetRecordingRecordingkeysRotationscheduleGatewayTimeout {
	return &GetRecordingRecordingkeysRotationscheduleGatewayTimeout{}
}

/*GetRecordingRecordingkeysRotationscheduleGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetRecordingRecordingkeysRotationscheduleGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetRecordingRecordingkeysRotationscheduleGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/recordingkeys/rotationschedule][%d] getRecordingRecordingkeysRotationscheduleGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetRecordingRecordingkeysRotationscheduleGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingRecordingkeysRotationscheduleGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

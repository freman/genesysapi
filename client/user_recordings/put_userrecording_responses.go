// Code generated by go-swagger; DO NOT EDIT.

package user_recordings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PutUserrecordingReader is a Reader for the PutUserrecording structure.
type PutUserrecordingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutUserrecordingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutUserrecordingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutUserrecordingBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutUserrecordingUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutUserrecordingForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutUserrecordingNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPutUserrecordingRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPutUserrecordingRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPutUserrecordingUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutUserrecordingTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutUserrecordingInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutUserrecordingServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutUserrecordingGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutUserrecordingOK creates a PutUserrecordingOK with default headers values
func NewPutUserrecordingOK() *PutUserrecordingOK {
	return &PutUserrecordingOK{}
}

/*PutUserrecordingOK handles this case with default header values.

successful operation
*/
type PutUserrecordingOK struct {
	Payload *models.UserRecording
}

func (o *PutUserrecordingOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/userrecordings/{recordingId}][%d] putUserrecordingOK  %+v", 200, o.Payload)
}

func (o *PutUserrecordingOK) GetPayload() *models.UserRecording {
	return o.Payload
}

func (o *PutUserrecordingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserRecording)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutUserrecordingBadRequest creates a PutUserrecordingBadRequest with default headers values
func NewPutUserrecordingBadRequest() *PutUserrecordingBadRequest {
	return &PutUserrecordingBadRequest{}
}

/*PutUserrecordingBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutUserrecordingBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PutUserrecordingBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/userrecordings/{recordingId}][%d] putUserrecordingBadRequest  %+v", 400, o.Payload)
}

func (o *PutUserrecordingBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutUserrecordingBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutUserrecordingUnauthorized creates a PutUserrecordingUnauthorized with default headers values
func NewPutUserrecordingUnauthorized() *PutUserrecordingUnauthorized {
	return &PutUserrecordingUnauthorized{}
}

/*PutUserrecordingUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PutUserrecordingUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PutUserrecordingUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/userrecordings/{recordingId}][%d] putUserrecordingUnauthorized  %+v", 401, o.Payload)
}

func (o *PutUserrecordingUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutUserrecordingUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutUserrecordingForbidden creates a PutUserrecordingForbidden with default headers values
func NewPutUserrecordingForbidden() *PutUserrecordingForbidden {
	return &PutUserrecordingForbidden{}
}

/*PutUserrecordingForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PutUserrecordingForbidden struct {
	Payload *models.ErrorBody
}

func (o *PutUserrecordingForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/userrecordings/{recordingId}][%d] putUserrecordingForbidden  %+v", 403, o.Payload)
}

func (o *PutUserrecordingForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutUserrecordingForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutUserrecordingNotFound creates a PutUserrecordingNotFound with default headers values
func NewPutUserrecordingNotFound() *PutUserrecordingNotFound {
	return &PutUserrecordingNotFound{}
}

/*PutUserrecordingNotFound handles this case with default header values.

The requested resource was not found.
*/
type PutUserrecordingNotFound struct {
	Payload *models.ErrorBody
}

func (o *PutUserrecordingNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/userrecordings/{recordingId}][%d] putUserrecordingNotFound  %+v", 404, o.Payload)
}

func (o *PutUserrecordingNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutUserrecordingNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutUserrecordingRequestTimeout creates a PutUserrecordingRequestTimeout with default headers values
func NewPutUserrecordingRequestTimeout() *PutUserrecordingRequestTimeout {
	return &PutUserrecordingRequestTimeout{}
}

/*PutUserrecordingRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PutUserrecordingRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PutUserrecordingRequestTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/userrecordings/{recordingId}][%d] putUserrecordingRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutUserrecordingRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutUserrecordingRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutUserrecordingRequestEntityTooLarge creates a PutUserrecordingRequestEntityTooLarge with default headers values
func NewPutUserrecordingRequestEntityTooLarge() *PutUserrecordingRequestEntityTooLarge {
	return &PutUserrecordingRequestEntityTooLarge{}
}

/*PutUserrecordingRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PutUserrecordingRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PutUserrecordingRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /api/v2/userrecordings/{recordingId}][%d] putUserrecordingRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutUserrecordingRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutUserrecordingRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutUserrecordingUnsupportedMediaType creates a PutUserrecordingUnsupportedMediaType with default headers values
func NewPutUserrecordingUnsupportedMediaType() *PutUserrecordingUnsupportedMediaType {
	return &PutUserrecordingUnsupportedMediaType{}
}

/*PutUserrecordingUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutUserrecordingUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PutUserrecordingUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v2/userrecordings/{recordingId}][%d] putUserrecordingUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutUserrecordingUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutUserrecordingUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutUserrecordingTooManyRequests creates a PutUserrecordingTooManyRequests with default headers values
func NewPutUserrecordingTooManyRequests() *PutUserrecordingTooManyRequests {
	return &PutUserrecordingTooManyRequests{}
}

/*PutUserrecordingTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PutUserrecordingTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PutUserrecordingTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v2/userrecordings/{recordingId}][%d] putUserrecordingTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutUserrecordingTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutUserrecordingTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutUserrecordingInternalServerError creates a PutUserrecordingInternalServerError with default headers values
func NewPutUserrecordingInternalServerError() *PutUserrecordingInternalServerError {
	return &PutUserrecordingInternalServerError{}
}

/*PutUserrecordingInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutUserrecordingInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PutUserrecordingInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/userrecordings/{recordingId}][%d] putUserrecordingInternalServerError  %+v", 500, o.Payload)
}

func (o *PutUserrecordingInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutUserrecordingInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutUserrecordingServiceUnavailable creates a PutUserrecordingServiceUnavailable with default headers values
func NewPutUserrecordingServiceUnavailable() *PutUserrecordingServiceUnavailable {
	return &PutUserrecordingServiceUnavailable{}
}

/*PutUserrecordingServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutUserrecordingServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PutUserrecordingServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/userrecordings/{recordingId}][%d] putUserrecordingServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutUserrecordingServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutUserrecordingServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutUserrecordingGatewayTimeout creates a PutUserrecordingGatewayTimeout with default headers values
func NewPutUserrecordingGatewayTimeout() *PutUserrecordingGatewayTimeout {
	return &PutUserrecordingGatewayTimeout{}
}

/*PutUserrecordingGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PutUserrecordingGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PutUserrecordingGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/userrecordings/{recordingId}][%d] putUserrecordingGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutUserrecordingGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutUserrecordingGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

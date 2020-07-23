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

// DeleteUserrecordingReader is a Reader for the DeleteUserrecording structure.
type DeleteUserrecordingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteUserrecordingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewDeleteUserrecordingAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteUserrecordingBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteUserrecordingUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteUserrecordingForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteUserrecordingNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteUserrecordingRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteUserrecordingUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteUserrecordingTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteUserrecordingInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteUserrecordingServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteUserrecordingGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteUserrecordingAccepted creates a DeleteUserrecordingAccepted with default headers values
func NewDeleteUserrecordingAccepted() *DeleteUserrecordingAccepted {
	return &DeleteUserrecordingAccepted{}
}

/*DeleteUserrecordingAccepted handles this case with default header values.

Accepted - Processing Delete
*/
type DeleteUserrecordingAccepted struct {
}

func (o *DeleteUserrecordingAccepted) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/userrecordings/{recordingId}][%d] deleteUserrecordingAccepted ", 202)
}

func (o *DeleteUserrecordingAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUserrecordingBadRequest creates a DeleteUserrecordingBadRequest with default headers values
func NewDeleteUserrecordingBadRequest() *DeleteUserrecordingBadRequest {
	return &DeleteUserrecordingBadRequest{}
}

/*DeleteUserrecordingBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteUserrecordingBadRequest struct {
	Payload *models.ErrorBody
}

func (o *DeleteUserrecordingBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/userrecordings/{recordingId}][%d] deleteUserrecordingBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteUserrecordingBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteUserrecordingBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserrecordingUnauthorized creates a DeleteUserrecordingUnauthorized with default headers values
func NewDeleteUserrecordingUnauthorized() *DeleteUserrecordingUnauthorized {
	return &DeleteUserrecordingUnauthorized{}
}

/*DeleteUserrecordingUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteUserrecordingUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *DeleteUserrecordingUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/userrecordings/{recordingId}][%d] deleteUserrecordingUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteUserrecordingUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteUserrecordingUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserrecordingForbidden creates a DeleteUserrecordingForbidden with default headers values
func NewDeleteUserrecordingForbidden() *DeleteUserrecordingForbidden {
	return &DeleteUserrecordingForbidden{}
}

/*DeleteUserrecordingForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type DeleteUserrecordingForbidden struct {
	Payload *models.ErrorBody
}

func (o *DeleteUserrecordingForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/userrecordings/{recordingId}][%d] deleteUserrecordingForbidden  %+v", 403, o.Payload)
}

func (o *DeleteUserrecordingForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteUserrecordingForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserrecordingNotFound creates a DeleteUserrecordingNotFound with default headers values
func NewDeleteUserrecordingNotFound() *DeleteUserrecordingNotFound {
	return &DeleteUserrecordingNotFound{}
}

/*DeleteUserrecordingNotFound handles this case with default header values.

The requested resource was not found.
*/
type DeleteUserrecordingNotFound struct {
	Payload *models.ErrorBody
}

func (o *DeleteUserrecordingNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/userrecordings/{recordingId}][%d] deleteUserrecordingNotFound  %+v", 404, o.Payload)
}

func (o *DeleteUserrecordingNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteUserrecordingNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserrecordingRequestEntityTooLarge creates a DeleteUserrecordingRequestEntityTooLarge with default headers values
func NewDeleteUserrecordingRequestEntityTooLarge() *DeleteUserrecordingRequestEntityTooLarge {
	return &DeleteUserrecordingRequestEntityTooLarge{}
}

/*DeleteUserrecordingRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type DeleteUserrecordingRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *DeleteUserrecordingRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/userrecordings/{recordingId}][%d] deleteUserrecordingRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteUserrecordingRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteUserrecordingRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserrecordingUnsupportedMediaType creates a DeleteUserrecordingUnsupportedMediaType with default headers values
func NewDeleteUserrecordingUnsupportedMediaType() *DeleteUserrecordingUnsupportedMediaType {
	return &DeleteUserrecordingUnsupportedMediaType{}
}

/*DeleteUserrecordingUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteUserrecordingUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *DeleteUserrecordingUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/userrecordings/{recordingId}][%d] deleteUserrecordingUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteUserrecordingUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteUserrecordingUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserrecordingTooManyRequests creates a DeleteUserrecordingTooManyRequests with default headers values
func NewDeleteUserrecordingTooManyRequests() *DeleteUserrecordingTooManyRequests {
	return &DeleteUserrecordingTooManyRequests{}
}

/*DeleteUserrecordingTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type DeleteUserrecordingTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *DeleteUserrecordingTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/userrecordings/{recordingId}][%d] deleteUserrecordingTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteUserrecordingTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteUserrecordingTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserrecordingInternalServerError creates a DeleteUserrecordingInternalServerError with default headers values
func NewDeleteUserrecordingInternalServerError() *DeleteUserrecordingInternalServerError {
	return &DeleteUserrecordingInternalServerError{}
}

/*DeleteUserrecordingInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteUserrecordingInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *DeleteUserrecordingInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/userrecordings/{recordingId}][%d] deleteUserrecordingInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteUserrecordingInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteUserrecordingInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserrecordingServiceUnavailable creates a DeleteUserrecordingServiceUnavailable with default headers values
func NewDeleteUserrecordingServiceUnavailable() *DeleteUserrecordingServiceUnavailable {
	return &DeleteUserrecordingServiceUnavailable{}
}

/*DeleteUserrecordingServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteUserrecordingServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *DeleteUserrecordingServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/userrecordings/{recordingId}][%d] deleteUserrecordingServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteUserrecordingServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteUserrecordingServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserrecordingGatewayTimeout creates a DeleteUserrecordingGatewayTimeout with default headers values
func NewDeleteUserrecordingGatewayTimeout() *DeleteUserrecordingGatewayTimeout {
	return &DeleteUserrecordingGatewayTimeout{}
}

/*DeleteUserrecordingGatewayTimeout handles this case with default header values.

The request timed out.
*/
type DeleteUserrecordingGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteUserrecordingGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/userrecordings/{recordingId}][%d] deleteUserrecordingGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteUserrecordingGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteUserrecordingGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

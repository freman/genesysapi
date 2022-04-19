// Code generated by go-swagger; DO NOT EDIT.

package greetings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// DeleteGreetingReader is a Reader for the DeleteGreeting structure.
type DeleteGreetingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteGreetingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 400:
		result := NewDeleteGreetingBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteGreetingUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteGreetingForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteGreetingNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewDeleteGreetingRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteGreetingRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteGreetingUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteGreetingTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteGreetingInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteGreetingServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteGreetingGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteGreetingDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteGreetingBadRequest creates a DeleteGreetingBadRequest with default headers values
func NewDeleteGreetingBadRequest() *DeleteGreetingBadRequest {
	return &DeleteGreetingBadRequest{}
}

/*DeleteGreetingBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteGreetingBadRequest struct {
	Payload *models.ErrorBody
}

func (o *DeleteGreetingBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/greetings/{greetingId}][%d] deleteGreetingBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteGreetingBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteGreetingBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteGreetingUnauthorized creates a DeleteGreetingUnauthorized with default headers values
func NewDeleteGreetingUnauthorized() *DeleteGreetingUnauthorized {
	return &DeleteGreetingUnauthorized{}
}

/*DeleteGreetingUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteGreetingUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *DeleteGreetingUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/greetings/{greetingId}][%d] deleteGreetingUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteGreetingUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteGreetingUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteGreetingForbidden creates a DeleteGreetingForbidden with default headers values
func NewDeleteGreetingForbidden() *DeleteGreetingForbidden {
	return &DeleteGreetingForbidden{}
}

/*DeleteGreetingForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type DeleteGreetingForbidden struct {
	Payload *models.ErrorBody
}

func (o *DeleteGreetingForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/greetings/{greetingId}][%d] deleteGreetingForbidden  %+v", 403, o.Payload)
}

func (o *DeleteGreetingForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteGreetingForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteGreetingNotFound creates a DeleteGreetingNotFound with default headers values
func NewDeleteGreetingNotFound() *DeleteGreetingNotFound {
	return &DeleteGreetingNotFound{}
}

/*DeleteGreetingNotFound handles this case with default header values.

The requested resource was not found.
*/
type DeleteGreetingNotFound struct {
	Payload *models.ErrorBody
}

func (o *DeleteGreetingNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/greetings/{greetingId}][%d] deleteGreetingNotFound  %+v", 404, o.Payload)
}

func (o *DeleteGreetingNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteGreetingNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteGreetingRequestTimeout creates a DeleteGreetingRequestTimeout with default headers values
func NewDeleteGreetingRequestTimeout() *DeleteGreetingRequestTimeout {
	return &DeleteGreetingRequestTimeout{}
}

/*DeleteGreetingRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type DeleteGreetingRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteGreetingRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/greetings/{greetingId}][%d] deleteGreetingRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteGreetingRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteGreetingRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteGreetingRequestEntityTooLarge creates a DeleteGreetingRequestEntityTooLarge with default headers values
func NewDeleteGreetingRequestEntityTooLarge() *DeleteGreetingRequestEntityTooLarge {
	return &DeleteGreetingRequestEntityTooLarge{}
}

/*DeleteGreetingRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type DeleteGreetingRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *DeleteGreetingRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/greetings/{greetingId}][%d] deleteGreetingRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteGreetingRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteGreetingRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteGreetingUnsupportedMediaType creates a DeleteGreetingUnsupportedMediaType with default headers values
func NewDeleteGreetingUnsupportedMediaType() *DeleteGreetingUnsupportedMediaType {
	return &DeleteGreetingUnsupportedMediaType{}
}

/*DeleteGreetingUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteGreetingUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *DeleteGreetingUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/greetings/{greetingId}][%d] deleteGreetingUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteGreetingUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteGreetingUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteGreetingTooManyRequests creates a DeleteGreetingTooManyRequests with default headers values
func NewDeleteGreetingTooManyRequests() *DeleteGreetingTooManyRequests {
	return &DeleteGreetingTooManyRequests{}
}

/*DeleteGreetingTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type DeleteGreetingTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *DeleteGreetingTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/greetings/{greetingId}][%d] deleteGreetingTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteGreetingTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteGreetingTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteGreetingInternalServerError creates a DeleteGreetingInternalServerError with default headers values
func NewDeleteGreetingInternalServerError() *DeleteGreetingInternalServerError {
	return &DeleteGreetingInternalServerError{}
}

/*DeleteGreetingInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteGreetingInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *DeleteGreetingInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/greetings/{greetingId}][%d] deleteGreetingInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteGreetingInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteGreetingInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteGreetingServiceUnavailable creates a DeleteGreetingServiceUnavailable with default headers values
func NewDeleteGreetingServiceUnavailable() *DeleteGreetingServiceUnavailable {
	return &DeleteGreetingServiceUnavailable{}
}

/*DeleteGreetingServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteGreetingServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *DeleteGreetingServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/greetings/{greetingId}][%d] deleteGreetingServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteGreetingServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteGreetingServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteGreetingGatewayTimeout creates a DeleteGreetingGatewayTimeout with default headers values
func NewDeleteGreetingGatewayTimeout() *DeleteGreetingGatewayTimeout {
	return &DeleteGreetingGatewayTimeout{}
}

/*DeleteGreetingGatewayTimeout handles this case with default header values.

The request timed out.
*/
type DeleteGreetingGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteGreetingGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/greetings/{greetingId}][%d] deleteGreetingGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteGreetingGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteGreetingGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteGreetingDefault creates a DeleteGreetingDefault with default headers values
func NewDeleteGreetingDefault(code int) *DeleteGreetingDefault {
	return &DeleteGreetingDefault{
		_statusCode: code,
	}
}

/*DeleteGreetingDefault handles this case with default header values.

successful operation
*/
type DeleteGreetingDefault struct {
	_statusCode int
}

// Code gets the status code for the delete greeting default response
func (o *DeleteGreetingDefault) Code() int {
	return o._statusCode
}

func (o *DeleteGreetingDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/greetings/{greetingId}][%d] deleteGreeting default ", o._statusCode)
}

func (o *DeleteGreetingDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

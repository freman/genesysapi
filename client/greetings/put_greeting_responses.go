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

// PutGreetingReader is a Reader for the PutGreeting structure.
type PutGreetingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutGreetingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutGreetingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutGreetingBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutGreetingUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutGreetingForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutGreetingNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPutGreetingRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPutGreetingUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutGreetingTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutGreetingInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutGreetingServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutGreetingGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutGreetingOK creates a PutGreetingOK with default headers values
func NewPutGreetingOK() *PutGreetingOK {
	return &PutGreetingOK{}
}

/*PutGreetingOK handles this case with default header values.

successful operation
*/
type PutGreetingOK struct {
	Payload *models.Greeting
}

func (o *PutGreetingOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/greetings/{greetingId}][%d] putGreetingOK  %+v", 200, o.Payload)
}

func (o *PutGreetingOK) GetPayload() *models.Greeting {
	return o.Payload
}

func (o *PutGreetingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Greeting)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutGreetingBadRequest creates a PutGreetingBadRequest with default headers values
func NewPutGreetingBadRequest() *PutGreetingBadRequest {
	return &PutGreetingBadRequest{}
}

/*PutGreetingBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutGreetingBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PutGreetingBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/greetings/{greetingId}][%d] putGreetingBadRequest  %+v", 400, o.Payload)
}

func (o *PutGreetingBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutGreetingBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutGreetingUnauthorized creates a PutGreetingUnauthorized with default headers values
func NewPutGreetingUnauthorized() *PutGreetingUnauthorized {
	return &PutGreetingUnauthorized{}
}

/*PutGreetingUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PutGreetingUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PutGreetingUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/greetings/{greetingId}][%d] putGreetingUnauthorized  %+v", 401, o.Payload)
}

func (o *PutGreetingUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutGreetingUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutGreetingForbidden creates a PutGreetingForbidden with default headers values
func NewPutGreetingForbidden() *PutGreetingForbidden {
	return &PutGreetingForbidden{}
}

/*PutGreetingForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PutGreetingForbidden struct {
	Payload *models.ErrorBody
}

func (o *PutGreetingForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/greetings/{greetingId}][%d] putGreetingForbidden  %+v", 403, o.Payload)
}

func (o *PutGreetingForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutGreetingForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutGreetingNotFound creates a PutGreetingNotFound with default headers values
func NewPutGreetingNotFound() *PutGreetingNotFound {
	return &PutGreetingNotFound{}
}

/*PutGreetingNotFound handles this case with default header values.

The requested resource was not found.
*/
type PutGreetingNotFound struct {
	Payload *models.ErrorBody
}

func (o *PutGreetingNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/greetings/{greetingId}][%d] putGreetingNotFound  %+v", 404, o.Payload)
}

func (o *PutGreetingNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutGreetingNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutGreetingRequestEntityTooLarge creates a PutGreetingRequestEntityTooLarge with default headers values
func NewPutGreetingRequestEntityTooLarge() *PutGreetingRequestEntityTooLarge {
	return &PutGreetingRequestEntityTooLarge{}
}

/*PutGreetingRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PutGreetingRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PutGreetingRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /api/v2/greetings/{greetingId}][%d] putGreetingRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutGreetingRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutGreetingRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutGreetingUnsupportedMediaType creates a PutGreetingUnsupportedMediaType with default headers values
func NewPutGreetingUnsupportedMediaType() *PutGreetingUnsupportedMediaType {
	return &PutGreetingUnsupportedMediaType{}
}

/*PutGreetingUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutGreetingUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PutGreetingUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v2/greetings/{greetingId}][%d] putGreetingUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutGreetingUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutGreetingUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutGreetingTooManyRequests creates a PutGreetingTooManyRequests with default headers values
func NewPutGreetingTooManyRequests() *PutGreetingTooManyRequests {
	return &PutGreetingTooManyRequests{}
}

/*PutGreetingTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PutGreetingTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PutGreetingTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v2/greetings/{greetingId}][%d] putGreetingTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutGreetingTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutGreetingTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutGreetingInternalServerError creates a PutGreetingInternalServerError with default headers values
func NewPutGreetingInternalServerError() *PutGreetingInternalServerError {
	return &PutGreetingInternalServerError{}
}

/*PutGreetingInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutGreetingInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PutGreetingInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/greetings/{greetingId}][%d] putGreetingInternalServerError  %+v", 500, o.Payload)
}

func (o *PutGreetingInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutGreetingInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutGreetingServiceUnavailable creates a PutGreetingServiceUnavailable with default headers values
func NewPutGreetingServiceUnavailable() *PutGreetingServiceUnavailable {
	return &PutGreetingServiceUnavailable{}
}

/*PutGreetingServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutGreetingServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PutGreetingServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/greetings/{greetingId}][%d] putGreetingServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutGreetingServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutGreetingServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutGreetingGatewayTimeout creates a PutGreetingGatewayTimeout with default headers values
func NewPutGreetingGatewayTimeout() *PutGreetingGatewayTimeout {
	return &PutGreetingGatewayTimeout{}
}

/*PutGreetingGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PutGreetingGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PutGreetingGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/greetings/{greetingId}][%d] putGreetingGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutGreetingGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutGreetingGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
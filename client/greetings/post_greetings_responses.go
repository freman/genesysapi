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

// PostGreetingsReader is a Reader for the PostGreetings structure.
type PostGreetingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostGreetingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostGreetingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostGreetingsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostGreetingsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostGreetingsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostGreetingsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostGreetingsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostGreetingsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostGreetingsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostGreetingsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostGreetingsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostGreetingsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostGreetingsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostGreetingsOK creates a PostGreetingsOK with default headers values
func NewPostGreetingsOK() *PostGreetingsOK {
	return &PostGreetingsOK{}
}

/*PostGreetingsOK handles this case with default header values.

successful operation
*/
type PostGreetingsOK struct {
	Payload *models.Greeting
}

func (o *PostGreetingsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/greetings][%d] postGreetingsOK  %+v", 200, o.Payload)
}

func (o *PostGreetingsOK) GetPayload() *models.Greeting {
	return o.Payload
}

func (o *PostGreetingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Greeting)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGreetingsBadRequest creates a PostGreetingsBadRequest with default headers values
func NewPostGreetingsBadRequest() *PostGreetingsBadRequest {
	return &PostGreetingsBadRequest{}
}

/*PostGreetingsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostGreetingsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostGreetingsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/greetings][%d] postGreetingsBadRequest  %+v", 400, o.Payload)
}

func (o *PostGreetingsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGreetingsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGreetingsUnauthorized creates a PostGreetingsUnauthorized with default headers values
func NewPostGreetingsUnauthorized() *PostGreetingsUnauthorized {
	return &PostGreetingsUnauthorized{}
}

/*PostGreetingsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostGreetingsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostGreetingsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/greetings][%d] postGreetingsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostGreetingsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGreetingsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGreetingsForbidden creates a PostGreetingsForbidden with default headers values
func NewPostGreetingsForbidden() *PostGreetingsForbidden {
	return &PostGreetingsForbidden{}
}

/*PostGreetingsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostGreetingsForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostGreetingsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/greetings][%d] postGreetingsForbidden  %+v", 403, o.Payload)
}

func (o *PostGreetingsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGreetingsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGreetingsNotFound creates a PostGreetingsNotFound with default headers values
func NewPostGreetingsNotFound() *PostGreetingsNotFound {
	return &PostGreetingsNotFound{}
}

/*PostGreetingsNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostGreetingsNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostGreetingsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/greetings][%d] postGreetingsNotFound  %+v", 404, o.Payload)
}

func (o *PostGreetingsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGreetingsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGreetingsRequestTimeout creates a PostGreetingsRequestTimeout with default headers values
func NewPostGreetingsRequestTimeout() *PostGreetingsRequestTimeout {
	return &PostGreetingsRequestTimeout{}
}

/*PostGreetingsRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostGreetingsRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostGreetingsRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/greetings][%d] postGreetingsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostGreetingsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGreetingsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGreetingsRequestEntityTooLarge creates a PostGreetingsRequestEntityTooLarge with default headers values
func NewPostGreetingsRequestEntityTooLarge() *PostGreetingsRequestEntityTooLarge {
	return &PostGreetingsRequestEntityTooLarge{}
}

/*PostGreetingsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PostGreetingsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostGreetingsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/greetings][%d] postGreetingsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostGreetingsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGreetingsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGreetingsUnsupportedMediaType creates a PostGreetingsUnsupportedMediaType with default headers values
func NewPostGreetingsUnsupportedMediaType() *PostGreetingsUnsupportedMediaType {
	return &PostGreetingsUnsupportedMediaType{}
}

/*PostGreetingsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostGreetingsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostGreetingsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/greetings][%d] postGreetingsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostGreetingsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGreetingsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGreetingsTooManyRequests creates a PostGreetingsTooManyRequests with default headers values
func NewPostGreetingsTooManyRequests() *PostGreetingsTooManyRequests {
	return &PostGreetingsTooManyRequests{}
}

/*PostGreetingsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostGreetingsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostGreetingsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/greetings][%d] postGreetingsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostGreetingsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGreetingsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGreetingsInternalServerError creates a PostGreetingsInternalServerError with default headers values
func NewPostGreetingsInternalServerError() *PostGreetingsInternalServerError {
	return &PostGreetingsInternalServerError{}
}

/*PostGreetingsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostGreetingsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostGreetingsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/greetings][%d] postGreetingsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostGreetingsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGreetingsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGreetingsServiceUnavailable creates a PostGreetingsServiceUnavailable with default headers values
func NewPostGreetingsServiceUnavailable() *PostGreetingsServiceUnavailable {
	return &PostGreetingsServiceUnavailable{}
}

/*PostGreetingsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostGreetingsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostGreetingsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/greetings][%d] postGreetingsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostGreetingsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGreetingsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGreetingsGatewayTimeout creates a PostGreetingsGatewayTimeout with default headers values
func NewPostGreetingsGatewayTimeout() *PostGreetingsGatewayTimeout {
	return &PostGreetingsGatewayTimeout{}
}

/*PostGreetingsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostGreetingsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostGreetingsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/greetings][%d] postGreetingsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostGreetingsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGreetingsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

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

// PostUserGreetingsReader is a Reader for the PostUserGreetings structure.
type PostUserGreetingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostUserGreetingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostUserGreetingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostUserGreetingsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostUserGreetingsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostUserGreetingsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostUserGreetingsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostUserGreetingsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostUserGreetingsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostUserGreetingsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostUserGreetingsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostUserGreetingsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostUserGreetingsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostUserGreetingsOK creates a PostUserGreetingsOK with default headers values
func NewPostUserGreetingsOK() *PostUserGreetingsOK {
	return &PostUserGreetingsOK{}
}

/*PostUserGreetingsOK handles this case with default header values.

successful operation
*/
type PostUserGreetingsOK struct {
	Payload *models.Greeting
}

func (o *PostUserGreetingsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/greetings][%d] postUserGreetingsOK  %+v", 200, o.Payload)
}

func (o *PostUserGreetingsOK) GetPayload() *models.Greeting {
	return o.Payload
}

func (o *PostUserGreetingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Greeting)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUserGreetingsBadRequest creates a PostUserGreetingsBadRequest with default headers values
func NewPostUserGreetingsBadRequest() *PostUserGreetingsBadRequest {
	return &PostUserGreetingsBadRequest{}
}

/*PostUserGreetingsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostUserGreetingsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostUserGreetingsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/greetings][%d] postUserGreetingsBadRequest  %+v", 400, o.Payload)
}

func (o *PostUserGreetingsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUserGreetingsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUserGreetingsUnauthorized creates a PostUserGreetingsUnauthorized with default headers values
func NewPostUserGreetingsUnauthorized() *PostUserGreetingsUnauthorized {
	return &PostUserGreetingsUnauthorized{}
}

/*PostUserGreetingsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostUserGreetingsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostUserGreetingsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/greetings][%d] postUserGreetingsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostUserGreetingsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUserGreetingsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUserGreetingsForbidden creates a PostUserGreetingsForbidden with default headers values
func NewPostUserGreetingsForbidden() *PostUserGreetingsForbidden {
	return &PostUserGreetingsForbidden{}
}

/*PostUserGreetingsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostUserGreetingsForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostUserGreetingsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/greetings][%d] postUserGreetingsForbidden  %+v", 403, o.Payload)
}

func (o *PostUserGreetingsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUserGreetingsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUserGreetingsNotFound creates a PostUserGreetingsNotFound with default headers values
func NewPostUserGreetingsNotFound() *PostUserGreetingsNotFound {
	return &PostUserGreetingsNotFound{}
}

/*PostUserGreetingsNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostUserGreetingsNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostUserGreetingsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/greetings][%d] postUserGreetingsNotFound  %+v", 404, o.Payload)
}

func (o *PostUserGreetingsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUserGreetingsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUserGreetingsRequestEntityTooLarge creates a PostUserGreetingsRequestEntityTooLarge with default headers values
func NewPostUserGreetingsRequestEntityTooLarge() *PostUserGreetingsRequestEntityTooLarge {
	return &PostUserGreetingsRequestEntityTooLarge{}
}

/*PostUserGreetingsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostUserGreetingsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostUserGreetingsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/greetings][%d] postUserGreetingsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostUserGreetingsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUserGreetingsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUserGreetingsUnsupportedMediaType creates a PostUserGreetingsUnsupportedMediaType with default headers values
func NewPostUserGreetingsUnsupportedMediaType() *PostUserGreetingsUnsupportedMediaType {
	return &PostUserGreetingsUnsupportedMediaType{}
}

/*PostUserGreetingsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostUserGreetingsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostUserGreetingsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/greetings][%d] postUserGreetingsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostUserGreetingsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUserGreetingsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUserGreetingsTooManyRequests creates a PostUserGreetingsTooManyRequests with default headers values
func NewPostUserGreetingsTooManyRequests() *PostUserGreetingsTooManyRequests {
	return &PostUserGreetingsTooManyRequests{}
}

/*PostUserGreetingsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PostUserGreetingsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostUserGreetingsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/greetings][%d] postUserGreetingsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostUserGreetingsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUserGreetingsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUserGreetingsInternalServerError creates a PostUserGreetingsInternalServerError with default headers values
func NewPostUserGreetingsInternalServerError() *PostUserGreetingsInternalServerError {
	return &PostUserGreetingsInternalServerError{}
}

/*PostUserGreetingsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostUserGreetingsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostUserGreetingsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/greetings][%d] postUserGreetingsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostUserGreetingsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUserGreetingsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUserGreetingsServiceUnavailable creates a PostUserGreetingsServiceUnavailable with default headers values
func NewPostUserGreetingsServiceUnavailable() *PostUserGreetingsServiceUnavailable {
	return &PostUserGreetingsServiceUnavailable{}
}

/*PostUserGreetingsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostUserGreetingsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostUserGreetingsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/greetings][%d] postUserGreetingsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostUserGreetingsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUserGreetingsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUserGreetingsGatewayTimeout creates a PostUserGreetingsGatewayTimeout with default headers values
func NewPostUserGreetingsGatewayTimeout() *PostUserGreetingsGatewayTimeout {
	return &PostUserGreetingsGatewayTimeout{}
}

/*PostUserGreetingsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostUserGreetingsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostUserGreetingsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/greetings][%d] postUserGreetingsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostUserGreetingsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUserGreetingsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

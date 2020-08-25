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

// PostGroupGreetingsReader is a Reader for the PostGroupGreetings structure.
type PostGroupGreetingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostGroupGreetingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostGroupGreetingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostGroupGreetingsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostGroupGreetingsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostGroupGreetingsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostGroupGreetingsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostGroupGreetingsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostGroupGreetingsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostGroupGreetingsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostGroupGreetingsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostGroupGreetingsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostGroupGreetingsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostGroupGreetingsOK creates a PostGroupGreetingsOK with default headers values
func NewPostGroupGreetingsOK() *PostGroupGreetingsOK {
	return &PostGroupGreetingsOK{}
}

/*PostGroupGreetingsOK handles this case with default header values.

successful operation
*/
type PostGroupGreetingsOK struct {
	Payload *models.Greeting
}

func (o *PostGroupGreetingsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/groups/{groupId}/greetings][%d] postGroupGreetingsOK  %+v", 200, o.Payload)
}

func (o *PostGroupGreetingsOK) GetPayload() *models.Greeting {
	return o.Payload
}

func (o *PostGroupGreetingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Greeting)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGroupGreetingsBadRequest creates a PostGroupGreetingsBadRequest with default headers values
func NewPostGroupGreetingsBadRequest() *PostGroupGreetingsBadRequest {
	return &PostGroupGreetingsBadRequest{}
}

/*PostGroupGreetingsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostGroupGreetingsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostGroupGreetingsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/groups/{groupId}/greetings][%d] postGroupGreetingsBadRequest  %+v", 400, o.Payload)
}

func (o *PostGroupGreetingsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGroupGreetingsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGroupGreetingsUnauthorized creates a PostGroupGreetingsUnauthorized with default headers values
func NewPostGroupGreetingsUnauthorized() *PostGroupGreetingsUnauthorized {
	return &PostGroupGreetingsUnauthorized{}
}

/*PostGroupGreetingsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostGroupGreetingsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostGroupGreetingsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/groups/{groupId}/greetings][%d] postGroupGreetingsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostGroupGreetingsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGroupGreetingsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGroupGreetingsForbidden creates a PostGroupGreetingsForbidden with default headers values
func NewPostGroupGreetingsForbidden() *PostGroupGreetingsForbidden {
	return &PostGroupGreetingsForbidden{}
}

/*PostGroupGreetingsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostGroupGreetingsForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostGroupGreetingsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/groups/{groupId}/greetings][%d] postGroupGreetingsForbidden  %+v", 403, o.Payload)
}

func (o *PostGroupGreetingsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGroupGreetingsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGroupGreetingsNotFound creates a PostGroupGreetingsNotFound with default headers values
func NewPostGroupGreetingsNotFound() *PostGroupGreetingsNotFound {
	return &PostGroupGreetingsNotFound{}
}

/*PostGroupGreetingsNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostGroupGreetingsNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostGroupGreetingsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/groups/{groupId}/greetings][%d] postGroupGreetingsNotFound  %+v", 404, o.Payload)
}

func (o *PostGroupGreetingsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGroupGreetingsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGroupGreetingsRequestEntityTooLarge creates a PostGroupGreetingsRequestEntityTooLarge with default headers values
func NewPostGroupGreetingsRequestEntityTooLarge() *PostGroupGreetingsRequestEntityTooLarge {
	return &PostGroupGreetingsRequestEntityTooLarge{}
}

/*PostGroupGreetingsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostGroupGreetingsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostGroupGreetingsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/groups/{groupId}/greetings][%d] postGroupGreetingsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostGroupGreetingsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGroupGreetingsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGroupGreetingsUnsupportedMediaType creates a PostGroupGreetingsUnsupportedMediaType with default headers values
func NewPostGroupGreetingsUnsupportedMediaType() *PostGroupGreetingsUnsupportedMediaType {
	return &PostGroupGreetingsUnsupportedMediaType{}
}

/*PostGroupGreetingsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostGroupGreetingsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostGroupGreetingsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/groups/{groupId}/greetings][%d] postGroupGreetingsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostGroupGreetingsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGroupGreetingsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGroupGreetingsTooManyRequests creates a PostGroupGreetingsTooManyRequests with default headers values
func NewPostGroupGreetingsTooManyRequests() *PostGroupGreetingsTooManyRequests {
	return &PostGroupGreetingsTooManyRequests{}
}

/*PostGroupGreetingsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PostGroupGreetingsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostGroupGreetingsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/groups/{groupId}/greetings][%d] postGroupGreetingsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostGroupGreetingsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGroupGreetingsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGroupGreetingsInternalServerError creates a PostGroupGreetingsInternalServerError with default headers values
func NewPostGroupGreetingsInternalServerError() *PostGroupGreetingsInternalServerError {
	return &PostGroupGreetingsInternalServerError{}
}

/*PostGroupGreetingsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostGroupGreetingsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostGroupGreetingsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/groups/{groupId}/greetings][%d] postGroupGreetingsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostGroupGreetingsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGroupGreetingsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGroupGreetingsServiceUnavailable creates a PostGroupGreetingsServiceUnavailable with default headers values
func NewPostGroupGreetingsServiceUnavailable() *PostGroupGreetingsServiceUnavailable {
	return &PostGroupGreetingsServiceUnavailable{}
}

/*PostGroupGreetingsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostGroupGreetingsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostGroupGreetingsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/groups/{groupId}/greetings][%d] postGroupGreetingsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostGroupGreetingsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGroupGreetingsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGroupGreetingsGatewayTimeout creates a PostGroupGreetingsGatewayTimeout with default headers values
func NewPostGroupGreetingsGatewayTimeout() *PostGroupGreetingsGatewayTimeout {
	return &PostGroupGreetingsGatewayTimeout{}
}

/*PostGroupGreetingsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostGroupGreetingsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostGroupGreetingsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/groups/{groupId}/greetings][%d] postGroupGreetingsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostGroupGreetingsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGroupGreetingsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
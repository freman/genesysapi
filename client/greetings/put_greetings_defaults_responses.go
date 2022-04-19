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

// PutGreetingsDefaultsReader is a Reader for the PutGreetingsDefaults structure.
type PutGreetingsDefaultsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutGreetingsDefaultsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutGreetingsDefaultsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutGreetingsDefaultsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutGreetingsDefaultsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutGreetingsDefaultsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutGreetingsDefaultsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPutGreetingsDefaultsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPutGreetingsDefaultsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPutGreetingsDefaultsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutGreetingsDefaultsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutGreetingsDefaultsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutGreetingsDefaultsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutGreetingsDefaultsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutGreetingsDefaultsOK creates a PutGreetingsDefaultsOK with default headers values
func NewPutGreetingsDefaultsOK() *PutGreetingsDefaultsOK {
	return &PutGreetingsDefaultsOK{}
}

/*PutGreetingsDefaultsOK handles this case with default header values.

successful operation
*/
type PutGreetingsDefaultsOK struct {
	Payload *models.DefaultGreetingList
}

func (o *PutGreetingsDefaultsOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/greetings/defaults][%d] putGreetingsDefaultsOK  %+v", 200, o.Payload)
}

func (o *PutGreetingsDefaultsOK) GetPayload() *models.DefaultGreetingList {
	return o.Payload
}

func (o *PutGreetingsDefaultsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DefaultGreetingList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutGreetingsDefaultsBadRequest creates a PutGreetingsDefaultsBadRequest with default headers values
func NewPutGreetingsDefaultsBadRequest() *PutGreetingsDefaultsBadRequest {
	return &PutGreetingsDefaultsBadRequest{}
}

/*PutGreetingsDefaultsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutGreetingsDefaultsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PutGreetingsDefaultsBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/greetings/defaults][%d] putGreetingsDefaultsBadRequest  %+v", 400, o.Payload)
}

func (o *PutGreetingsDefaultsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutGreetingsDefaultsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutGreetingsDefaultsUnauthorized creates a PutGreetingsDefaultsUnauthorized with default headers values
func NewPutGreetingsDefaultsUnauthorized() *PutGreetingsDefaultsUnauthorized {
	return &PutGreetingsDefaultsUnauthorized{}
}

/*PutGreetingsDefaultsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PutGreetingsDefaultsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PutGreetingsDefaultsUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/greetings/defaults][%d] putGreetingsDefaultsUnauthorized  %+v", 401, o.Payload)
}

func (o *PutGreetingsDefaultsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutGreetingsDefaultsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutGreetingsDefaultsForbidden creates a PutGreetingsDefaultsForbidden with default headers values
func NewPutGreetingsDefaultsForbidden() *PutGreetingsDefaultsForbidden {
	return &PutGreetingsDefaultsForbidden{}
}

/*PutGreetingsDefaultsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PutGreetingsDefaultsForbidden struct {
	Payload *models.ErrorBody
}

func (o *PutGreetingsDefaultsForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/greetings/defaults][%d] putGreetingsDefaultsForbidden  %+v", 403, o.Payload)
}

func (o *PutGreetingsDefaultsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutGreetingsDefaultsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutGreetingsDefaultsNotFound creates a PutGreetingsDefaultsNotFound with default headers values
func NewPutGreetingsDefaultsNotFound() *PutGreetingsDefaultsNotFound {
	return &PutGreetingsDefaultsNotFound{}
}

/*PutGreetingsDefaultsNotFound handles this case with default header values.

The requested resource was not found.
*/
type PutGreetingsDefaultsNotFound struct {
	Payload *models.ErrorBody
}

func (o *PutGreetingsDefaultsNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/greetings/defaults][%d] putGreetingsDefaultsNotFound  %+v", 404, o.Payload)
}

func (o *PutGreetingsDefaultsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutGreetingsDefaultsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutGreetingsDefaultsRequestTimeout creates a PutGreetingsDefaultsRequestTimeout with default headers values
func NewPutGreetingsDefaultsRequestTimeout() *PutGreetingsDefaultsRequestTimeout {
	return &PutGreetingsDefaultsRequestTimeout{}
}

/*PutGreetingsDefaultsRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PutGreetingsDefaultsRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PutGreetingsDefaultsRequestTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/greetings/defaults][%d] putGreetingsDefaultsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutGreetingsDefaultsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutGreetingsDefaultsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutGreetingsDefaultsRequestEntityTooLarge creates a PutGreetingsDefaultsRequestEntityTooLarge with default headers values
func NewPutGreetingsDefaultsRequestEntityTooLarge() *PutGreetingsDefaultsRequestEntityTooLarge {
	return &PutGreetingsDefaultsRequestEntityTooLarge{}
}

/*PutGreetingsDefaultsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PutGreetingsDefaultsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PutGreetingsDefaultsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /api/v2/greetings/defaults][%d] putGreetingsDefaultsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutGreetingsDefaultsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutGreetingsDefaultsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutGreetingsDefaultsUnsupportedMediaType creates a PutGreetingsDefaultsUnsupportedMediaType with default headers values
func NewPutGreetingsDefaultsUnsupportedMediaType() *PutGreetingsDefaultsUnsupportedMediaType {
	return &PutGreetingsDefaultsUnsupportedMediaType{}
}

/*PutGreetingsDefaultsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutGreetingsDefaultsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PutGreetingsDefaultsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v2/greetings/defaults][%d] putGreetingsDefaultsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutGreetingsDefaultsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutGreetingsDefaultsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutGreetingsDefaultsTooManyRequests creates a PutGreetingsDefaultsTooManyRequests with default headers values
func NewPutGreetingsDefaultsTooManyRequests() *PutGreetingsDefaultsTooManyRequests {
	return &PutGreetingsDefaultsTooManyRequests{}
}

/*PutGreetingsDefaultsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PutGreetingsDefaultsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PutGreetingsDefaultsTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v2/greetings/defaults][%d] putGreetingsDefaultsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutGreetingsDefaultsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutGreetingsDefaultsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutGreetingsDefaultsInternalServerError creates a PutGreetingsDefaultsInternalServerError with default headers values
func NewPutGreetingsDefaultsInternalServerError() *PutGreetingsDefaultsInternalServerError {
	return &PutGreetingsDefaultsInternalServerError{}
}

/*PutGreetingsDefaultsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutGreetingsDefaultsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PutGreetingsDefaultsInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/greetings/defaults][%d] putGreetingsDefaultsInternalServerError  %+v", 500, o.Payload)
}

func (o *PutGreetingsDefaultsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutGreetingsDefaultsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutGreetingsDefaultsServiceUnavailable creates a PutGreetingsDefaultsServiceUnavailable with default headers values
func NewPutGreetingsDefaultsServiceUnavailable() *PutGreetingsDefaultsServiceUnavailable {
	return &PutGreetingsDefaultsServiceUnavailable{}
}

/*PutGreetingsDefaultsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutGreetingsDefaultsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PutGreetingsDefaultsServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/greetings/defaults][%d] putGreetingsDefaultsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutGreetingsDefaultsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutGreetingsDefaultsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutGreetingsDefaultsGatewayTimeout creates a PutGreetingsDefaultsGatewayTimeout with default headers values
func NewPutGreetingsDefaultsGatewayTimeout() *PutGreetingsDefaultsGatewayTimeout {
	return &PutGreetingsDefaultsGatewayTimeout{}
}

/*PutGreetingsDefaultsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PutGreetingsDefaultsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PutGreetingsDefaultsGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/greetings/defaults][%d] putGreetingsDefaultsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutGreetingsDefaultsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutGreetingsDefaultsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

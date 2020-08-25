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

// PutUserGreetingsDefaultsReader is a Reader for the PutUserGreetingsDefaults structure.
type PutUserGreetingsDefaultsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutUserGreetingsDefaultsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutUserGreetingsDefaultsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutUserGreetingsDefaultsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutUserGreetingsDefaultsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutUserGreetingsDefaultsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutUserGreetingsDefaultsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPutUserGreetingsDefaultsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPutUserGreetingsDefaultsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutUserGreetingsDefaultsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutUserGreetingsDefaultsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutUserGreetingsDefaultsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutUserGreetingsDefaultsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutUserGreetingsDefaultsOK creates a PutUserGreetingsDefaultsOK with default headers values
func NewPutUserGreetingsDefaultsOK() *PutUserGreetingsDefaultsOK {
	return &PutUserGreetingsDefaultsOK{}
}

/*PutUserGreetingsDefaultsOK handles this case with default header values.

successful operation
*/
type PutUserGreetingsDefaultsOK struct {
	Payload *models.DefaultGreetingList
}

func (o *PutUserGreetingsDefaultsOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/users/{userId}/greetings/defaults][%d] putUserGreetingsDefaultsOK  %+v", 200, o.Payload)
}

func (o *PutUserGreetingsDefaultsOK) GetPayload() *models.DefaultGreetingList {
	return o.Payload
}

func (o *PutUserGreetingsDefaultsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DefaultGreetingList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutUserGreetingsDefaultsBadRequest creates a PutUserGreetingsDefaultsBadRequest with default headers values
func NewPutUserGreetingsDefaultsBadRequest() *PutUserGreetingsDefaultsBadRequest {
	return &PutUserGreetingsDefaultsBadRequest{}
}

/*PutUserGreetingsDefaultsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutUserGreetingsDefaultsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PutUserGreetingsDefaultsBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/users/{userId}/greetings/defaults][%d] putUserGreetingsDefaultsBadRequest  %+v", 400, o.Payload)
}

func (o *PutUserGreetingsDefaultsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutUserGreetingsDefaultsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutUserGreetingsDefaultsUnauthorized creates a PutUserGreetingsDefaultsUnauthorized with default headers values
func NewPutUserGreetingsDefaultsUnauthorized() *PutUserGreetingsDefaultsUnauthorized {
	return &PutUserGreetingsDefaultsUnauthorized{}
}

/*PutUserGreetingsDefaultsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PutUserGreetingsDefaultsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PutUserGreetingsDefaultsUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/users/{userId}/greetings/defaults][%d] putUserGreetingsDefaultsUnauthorized  %+v", 401, o.Payload)
}

func (o *PutUserGreetingsDefaultsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutUserGreetingsDefaultsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutUserGreetingsDefaultsForbidden creates a PutUserGreetingsDefaultsForbidden with default headers values
func NewPutUserGreetingsDefaultsForbidden() *PutUserGreetingsDefaultsForbidden {
	return &PutUserGreetingsDefaultsForbidden{}
}

/*PutUserGreetingsDefaultsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PutUserGreetingsDefaultsForbidden struct {
	Payload *models.ErrorBody
}

func (o *PutUserGreetingsDefaultsForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/users/{userId}/greetings/defaults][%d] putUserGreetingsDefaultsForbidden  %+v", 403, o.Payload)
}

func (o *PutUserGreetingsDefaultsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutUserGreetingsDefaultsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutUserGreetingsDefaultsNotFound creates a PutUserGreetingsDefaultsNotFound with default headers values
func NewPutUserGreetingsDefaultsNotFound() *PutUserGreetingsDefaultsNotFound {
	return &PutUserGreetingsDefaultsNotFound{}
}

/*PutUserGreetingsDefaultsNotFound handles this case with default header values.

The requested resource was not found.
*/
type PutUserGreetingsDefaultsNotFound struct {
	Payload *models.ErrorBody
}

func (o *PutUserGreetingsDefaultsNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/users/{userId}/greetings/defaults][%d] putUserGreetingsDefaultsNotFound  %+v", 404, o.Payload)
}

func (o *PutUserGreetingsDefaultsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutUserGreetingsDefaultsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutUserGreetingsDefaultsRequestEntityTooLarge creates a PutUserGreetingsDefaultsRequestEntityTooLarge with default headers values
func NewPutUserGreetingsDefaultsRequestEntityTooLarge() *PutUserGreetingsDefaultsRequestEntityTooLarge {
	return &PutUserGreetingsDefaultsRequestEntityTooLarge{}
}

/*PutUserGreetingsDefaultsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PutUserGreetingsDefaultsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PutUserGreetingsDefaultsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /api/v2/users/{userId}/greetings/defaults][%d] putUserGreetingsDefaultsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutUserGreetingsDefaultsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutUserGreetingsDefaultsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutUserGreetingsDefaultsUnsupportedMediaType creates a PutUserGreetingsDefaultsUnsupportedMediaType with default headers values
func NewPutUserGreetingsDefaultsUnsupportedMediaType() *PutUserGreetingsDefaultsUnsupportedMediaType {
	return &PutUserGreetingsDefaultsUnsupportedMediaType{}
}

/*PutUserGreetingsDefaultsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutUserGreetingsDefaultsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PutUserGreetingsDefaultsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v2/users/{userId}/greetings/defaults][%d] putUserGreetingsDefaultsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutUserGreetingsDefaultsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutUserGreetingsDefaultsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutUserGreetingsDefaultsTooManyRequests creates a PutUserGreetingsDefaultsTooManyRequests with default headers values
func NewPutUserGreetingsDefaultsTooManyRequests() *PutUserGreetingsDefaultsTooManyRequests {
	return &PutUserGreetingsDefaultsTooManyRequests{}
}

/*PutUserGreetingsDefaultsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PutUserGreetingsDefaultsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PutUserGreetingsDefaultsTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v2/users/{userId}/greetings/defaults][%d] putUserGreetingsDefaultsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutUserGreetingsDefaultsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutUserGreetingsDefaultsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutUserGreetingsDefaultsInternalServerError creates a PutUserGreetingsDefaultsInternalServerError with default headers values
func NewPutUserGreetingsDefaultsInternalServerError() *PutUserGreetingsDefaultsInternalServerError {
	return &PutUserGreetingsDefaultsInternalServerError{}
}

/*PutUserGreetingsDefaultsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutUserGreetingsDefaultsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PutUserGreetingsDefaultsInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/users/{userId}/greetings/defaults][%d] putUserGreetingsDefaultsInternalServerError  %+v", 500, o.Payload)
}

func (o *PutUserGreetingsDefaultsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutUserGreetingsDefaultsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutUserGreetingsDefaultsServiceUnavailable creates a PutUserGreetingsDefaultsServiceUnavailable with default headers values
func NewPutUserGreetingsDefaultsServiceUnavailable() *PutUserGreetingsDefaultsServiceUnavailable {
	return &PutUserGreetingsDefaultsServiceUnavailable{}
}

/*PutUserGreetingsDefaultsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutUserGreetingsDefaultsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PutUserGreetingsDefaultsServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/users/{userId}/greetings/defaults][%d] putUserGreetingsDefaultsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutUserGreetingsDefaultsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutUserGreetingsDefaultsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutUserGreetingsDefaultsGatewayTimeout creates a PutUserGreetingsDefaultsGatewayTimeout with default headers values
func NewPutUserGreetingsDefaultsGatewayTimeout() *PutUserGreetingsDefaultsGatewayTimeout {
	return &PutUserGreetingsDefaultsGatewayTimeout{}
}

/*PutUserGreetingsDefaultsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PutUserGreetingsDefaultsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PutUserGreetingsDefaultsGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/users/{userId}/greetings/defaults][%d] putUserGreetingsDefaultsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutUserGreetingsDefaultsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutUserGreetingsDefaultsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
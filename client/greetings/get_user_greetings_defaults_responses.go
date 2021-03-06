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

// GetUserGreetingsDefaultsReader is a Reader for the GetUserGreetingsDefaults structure.
type GetUserGreetingsDefaultsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserGreetingsDefaultsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUserGreetingsDefaultsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetUserGreetingsDefaultsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetUserGreetingsDefaultsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetUserGreetingsDefaultsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetUserGreetingsDefaultsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetUserGreetingsDefaultsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetUserGreetingsDefaultsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetUserGreetingsDefaultsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUserGreetingsDefaultsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetUserGreetingsDefaultsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetUserGreetingsDefaultsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUserGreetingsDefaultsOK creates a GetUserGreetingsDefaultsOK with default headers values
func NewGetUserGreetingsDefaultsOK() *GetUserGreetingsDefaultsOK {
	return &GetUserGreetingsDefaultsOK{}
}

/*GetUserGreetingsDefaultsOK handles this case with default header values.

successful operation
*/
type GetUserGreetingsDefaultsOK struct {
	Payload *models.DefaultGreetingList
}

func (o *GetUserGreetingsDefaultsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/greetings/defaults][%d] getUserGreetingsDefaultsOK  %+v", 200, o.Payload)
}

func (o *GetUserGreetingsDefaultsOK) GetPayload() *models.DefaultGreetingList {
	return o.Payload
}

func (o *GetUserGreetingsDefaultsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DefaultGreetingList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserGreetingsDefaultsBadRequest creates a GetUserGreetingsDefaultsBadRequest with default headers values
func NewGetUserGreetingsDefaultsBadRequest() *GetUserGreetingsDefaultsBadRequest {
	return &GetUserGreetingsDefaultsBadRequest{}
}

/*GetUserGreetingsDefaultsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetUserGreetingsDefaultsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetUserGreetingsDefaultsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/greetings/defaults][%d] getUserGreetingsDefaultsBadRequest  %+v", 400, o.Payload)
}

func (o *GetUserGreetingsDefaultsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserGreetingsDefaultsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserGreetingsDefaultsUnauthorized creates a GetUserGreetingsDefaultsUnauthorized with default headers values
func NewGetUserGreetingsDefaultsUnauthorized() *GetUserGreetingsDefaultsUnauthorized {
	return &GetUserGreetingsDefaultsUnauthorized{}
}

/*GetUserGreetingsDefaultsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetUserGreetingsDefaultsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetUserGreetingsDefaultsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/greetings/defaults][%d] getUserGreetingsDefaultsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetUserGreetingsDefaultsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserGreetingsDefaultsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserGreetingsDefaultsForbidden creates a GetUserGreetingsDefaultsForbidden with default headers values
func NewGetUserGreetingsDefaultsForbidden() *GetUserGreetingsDefaultsForbidden {
	return &GetUserGreetingsDefaultsForbidden{}
}

/*GetUserGreetingsDefaultsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetUserGreetingsDefaultsForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetUserGreetingsDefaultsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/greetings/defaults][%d] getUserGreetingsDefaultsForbidden  %+v", 403, o.Payload)
}

func (o *GetUserGreetingsDefaultsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserGreetingsDefaultsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserGreetingsDefaultsNotFound creates a GetUserGreetingsDefaultsNotFound with default headers values
func NewGetUserGreetingsDefaultsNotFound() *GetUserGreetingsDefaultsNotFound {
	return &GetUserGreetingsDefaultsNotFound{}
}

/*GetUserGreetingsDefaultsNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetUserGreetingsDefaultsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetUserGreetingsDefaultsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/greetings/defaults][%d] getUserGreetingsDefaultsNotFound  %+v", 404, o.Payload)
}

func (o *GetUserGreetingsDefaultsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserGreetingsDefaultsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserGreetingsDefaultsRequestEntityTooLarge creates a GetUserGreetingsDefaultsRequestEntityTooLarge with default headers values
func NewGetUserGreetingsDefaultsRequestEntityTooLarge() *GetUserGreetingsDefaultsRequestEntityTooLarge {
	return &GetUserGreetingsDefaultsRequestEntityTooLarge{}
}

/*GetUserGreetingsDefaultsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetUserGreetingsDefaultsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetUserGreetingsDefaultsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/greetings/defaults][%d] getUserGreetingsDefaultsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetUserGreetingsDefaultsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserGreetingsDefaultsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserGreetingsDefaultsUnsupportedMediaType creates a GetUserGreetingsDefaultsUnsupportedMediaType with default headers values
func NewGetUserGreetingsDefaultsUnsupportedMediaType() *GetUserGreetingsDefaultsUnsupportedMediaType {
	return &GetUserGreetingsDefaultsUnsupportedMediaType{}
}

/*GetUserGreetingsDefaultsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetUserGreetingsDefaultsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetUserGreetingsDefaultsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/greetings/defaults][%d] getUserGreetingsDefaultsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetUserGreetingsDefaultsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserGreetingsDefaultsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserGreetingsDefaultsTooManyRequests creates a GetUserGreetingsDefaultsTooManyRequests with default headers values
func NewGetUserGreetingsDefaultsTooManyRequests() *GetUserGreetingsDefaultsTooManyRequests {
	return &GetUserGreetingsDefaultsTooManyRequests{}
}

/*GetUserGreetingsDefaultsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetUserGreetingsDefaultsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetUserGreetingsDefaultsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/greetings/defaults][%d] getUserGreetingsDefaultsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetUserGreetingsDefaultsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserGreetingsDefaultsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserGreetingsDefaultsInternalServerError creates a GetUserGreetingsDefaultsInternalServerError with default headers values
func NewGetUserGreetingsDefaultsInternalServerError() *GetUserGreetingsDefaultsInternalServerError {
	return &GetUserGreetingsDefaultsInternalServerError{}
}

/*GetUserGreetingsDefaultsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetUserGreetingsDefaultsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetUserGreetingsDefaultsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/greetings/defaults][%d] getUserGreetingsDefaultsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUserGreetingsDefaultsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserGreetingsDefaultsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserGreetingsDefaultsServiceUnavailable creates a GetUserGreetingsDefaultsServiceUnavailable with default headers values
func NewGetUserGreetingsDefaultsServiceUnavailable() *GetUserGreetingsDefaultsServiceUnavailable {
	return &GetUserGreetingsDefaultsServiceUnavailable{}
}

/*GetUserGreetingsDefaultsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetUserGreetingsDefaultsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetUserGreetingsDefaultsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/greetings/defaults][%d] getUserGreetingsDefaultsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetUserGreetingsDefaultsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserGreetingsDefaultsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserGreetingsDefaultsGatewayTimeout creates a GetUserGreetingsDefaultsGatewayTimeout with default headers values
func NewGetUserGreetingsDefaultsGatewayTimeout() *GetUserGreetingsDefaultsGatewayTimeout {
	return &GetUserGreetingsDefaultsGatewayTimeout{}
}

/*GetUserGreetingsDefaultsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetUserGreetingsDefaultsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetUserGreetingsDefaultsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/greetings/defaults][%d] getUserGreetingsDefaultsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetUserGreetingsDefaultsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserGreetingsDefaultsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostUserRoutinglanguagesReader is a Reader for the PostUserRoutinglanguages structure.
type PostUserRoutinglanguagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostUserRoutinglanguagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostUserRoutinglanguagesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostUserRoutinglanguagesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostUserRoutinglanguagesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostUserRoutinglanguagesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostUserRoutinglanguagesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostUserRoutinglanguagesRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostUserRoutinglanguagesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostUserRoutinglanguagesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostUserRoutinglanguagesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostUserRoutinglanguagesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostUserRoutinglanguagesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostUserRoutinglanguagesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostUserRoutinglanguagesOK creates a PostUserRoutinglanguagesOK with default headers values
func NewPostUserRoutinglanguagesOK() *PostUserRoutinglanguagesOK {
	return &PostUserRoutinglanguagesOK{}
}

/*PostUserRoutinglanguagesOK handles this case with default header values.

successful operation
*/
type PostUserRoutinglanguagesOK struct {
	Payload *models.UserRoutingLanguage
}

func (o *PostUserRoutinglanguagesOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/routinglanguages][%d] postUserRoutinglanguagesOK  %+v", 200, o.Payload)
}

func (o *PostUserRoutinglanguagesOK) GetPayload() *models.UserRoutingLanguage {
	return o.Payload
}

func (o *PostUserRoutinglanguagesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserRoutingLanguage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUserRoutinglanguagesBadRequest creates a PostUserRoutinglanguagesBadRequest with default headers values
func NewPostUserRoutinglanguagesBadRequest() *PostUserRoutinglanguagesBadRequest {
	return &PostUserRoutinglanguagesBadRequest{}
}

/*PostUserRoutinglanguagesBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostUserRoutinglanguagesBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostUserRoutinglanguagesBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/routinglanguages][%d] postUserRoutinglanguagesBadRequest  %+v", 400, o.Payload)
}

func (o *PostUserRoutinglanguagesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUserRoutinglanguagesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUserRoutinglanguagesUnauthorized creates a PostUserRoutinglanguagesUnauthorized with default headers values
func NewPostUserRoutinglanguagesUnauthorized() *PostUserRoutinglanguagesUnauthorized {
	return &PostUserRoutinglanguagesUnauthorized{}
}

/*PostUserRoutinglanguagesUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostUserRoutinglanguagesUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostUserRoutinglanguagesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/routinglanguages][%d] postUserRoutinglanguagesUnauthorized  %+v", 401, o.Payload)
}

func (o *PostUserRoutinglanguagesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUserRoutinglanguagesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUserRoutinglanguagesForbidden creates a PostUserRoutinglanguagesForbidden with default headers values
func NewPostUserRoutinglanguagesForbidden() *PostUserRoutinglanguagesForbidden {
	return &PostUserRoutinglanguagesForbidden{}
}

/*PostUserRoutinglanguagesForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostUserRoutinglanguagesForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostUserRoutinglanguagesForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/routinglanguages][%d] postUserRoutinglanguagesForbidden  %+v", 403, o.Payload)
}

func (o *PostUserRoutinglanguagesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUserRoutinglanguagesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUserRoutinglanguagesNotFound creates a PostUserRoutinglanguagesNotFound with default headers values
func NewPostUserRoutinglanguagesNotFound() *PostUserRoutinglanguagesNotFound {
	return &PostUserRoutinglanguagesNotFound{}
}

/*PostUserRoutinglanguagesNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostUserRoutinglanguagesNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostUserRoutinglanguagesNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/routinglanguages][%d] postUserRoutinglanguagesNotFound  %+v", 404, o.Payload)
}

func (o *PostUserRoutinglanguagesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUserRoutinglanguagesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUserRoutinglanguagesRequestTimeout creates a PostUserRoutinglanguagesRequestTimeout with default headers values
func NewPostUserRoutinglanguagesRequestTimeout() *PostUserRoutinglanguagesRequestTimeout {
	return &PostUserRoutinglanguagesRequestTimeout{}
}

/*PostUserRoutinglanguagesRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostUserRoutinglanguagesRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostUserRoutinglanguagesRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/routinglanguages][%d] postUserRoutinglanguagesRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostUserRoutinglanguagesRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUserRoutinglanguagesRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUserRoutinglanguagesRequestEntityTooLarge creates a PostUserRoutinglanguagesRequestEntityTooLarge with default headers values
func NewPostUserRoutinglanguagesRequestEntityTooLarge() *PostUserRoutinglanguagesRequestEntityTooLarge {
	return &PostUserRoutinglanguagesRequestEntityTooLarge{}
}

/*PostUserRoutinglanguagesRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PostUserRoutinglanguagesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostUserRoutinglanguagesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/routinglanguages][%d] postUserRoutinglanguagesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostUserRoutinglanguagesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUserRoutinglanguagesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUserRoutinglanguagesUnsupportedMediaType creates a PostUserRoutinglanguagesUnsupportedMediaType with default headers values
func NewPostUserRoutinglanguagesUnsupportedMediaType() *PostUserRoutinglanguagesUnsupportedMediaType {
	return &PostUserRoutinglanguagesUnsupportedMediaType{}
}

/*PostUserRoutinglanguagesUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostUserRoutinglanguagesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostUserRoutinglanguagesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/routinglanguages][%d] postUserRoutinglanguagesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostUserRoutinglanguagesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUserRoutinglanguagesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUserRoutinglanguagesTooManyRequests creates a PostUserRoutinglanguagesTooManyRequests with default headers values
func NewPostUserRoutinglanguagesTooManyRequests() *PostUserRoutinglanguagesTooManyRequests {
	return &PostUserRoutinglanguagesTooManyRequests{}
}

/*PostUserRoutinglanguagesTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostUserRoutinglanguagesTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostUserRoutinglanguagesTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/routinglanguages][%d] postUserRoutinglanguagesTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostUserRoutinglanguagesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUserRoutinglanguagesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUserRoutinglanguagesInternalServerError creates a PostUserRoutinglanguagesInternalServerError with default headers values
func NewPostUserRoutinglanguagesInternalServerError() *PostUserRoutinglanguagesInternalServerError {
	return &PostUserRoutinglanguagesInternalServerError{}
}

/*PostUserRoutinglanguagesInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostUserRoutinglanguagesInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostUserRoutinglanguagesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/routinglanguages][%d] postUserRoutinglanguagesInternalServerError  %+v", 500, o.Payload)
}

func (o *PostUserRoutinglanguagesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUserRoutinglanguagesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUserRoutinglanguagesServiceUnavailable creates a PostUserRoutinglanguagesServiceUnavailable with default headers values
func NewPostUserRoutinglanguagesServiceUnavailable() *PostUserRoutinglanguagesServiceUnavailable {
	return &PostUserRoutinglanguagesServiceUnavailable{}
}

/*PostUserRoutinglanguagesServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostUserRoutinglanguagesServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostUserRoutinglanguagesServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/routinglanguages][%d] postUserRoutinglanguagesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostUserRoutinglanguagesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUserRoutinglanguagesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUserRoutinglanguagesGatewayTimeout creates a PostUserRoutinglanguagesGatewayTimeout with default headers values
func NewPostUserRoutinglanguagesGatewayTimeout() *PostUserRoutinglanguagesGatewayTimeout {
	return &PostUserRoutinglanguagesGatewayTimeout{}
}

/*PostUserRoutinglanguagesGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostUserRoutinglanguagesGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostUserRoutinglanguagesGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/routinglanguages][%d] postUserRoutinglanguagesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostUserRoutinglanguagesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUserRoutinglanguagesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

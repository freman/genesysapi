// Code generated by go-swagger; DO NOT EDIT.

package presence

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetUserPresenceReader is a Reader for the GetUserPresence structure.
type GetUserPresenceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserPresenceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUserPresenceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetUserPresenceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetUserPresenceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetUserPresenceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetUserPresenceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetUserPresenceRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetUserPresenceUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetUserPresenceTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUserPresenceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetUserPresenceServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetUserPresenceGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUserPresenceOK creates a GetUserPresenceOK with default headers values
func NewGetUserPresenceOK() *GetUserPresenceOK {
	return &GetUserPresenceOK{}
}

/*GetUserPresenceOK handles this case with default header values.

successful operation
*/
type GetUserPresenceOK struct {
	Payload *models.UserPresence
}

func (o *GetUserPresenceOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/presences/{sourceId}][%d] getUserPresenceOK  %+v", 200, o.Payload)
}

func (o *GetUserPresenceOK) GetPayload() *models.UserPresence {
	return o.Payload
}

func (o *GetUserPresenceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserPresence)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserPresenceBadRequest creates a GetUserPresenceBadRequest with default headers values
func NewGetUserPresenceBadRequest() *GetUserPresenceBadRequest {
	return &GetUserPresenceBadRequest{}
}

/*GetUserPresenceBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetUserPresenceBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetUserPresenceBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/presences/{sourceId}][%d] getUserPresenceBadRequest  %+v", 400, o.Payload)
}

func (o *GetUserPresenceBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserPresenceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserPresenceUnauthorized creates a GetUserPresenceUnauthorized with default headers values
func NewGetUserPresenceUnauthorized() *GetUserPresenceUnauthorized {
	return &GetUserPresenceUnauthorized{}
}

/*GetUserPresenceUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetUserPresenceUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetUserPresenceUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/presences/{sourceId}][%d] getUserPresenceUnauthorized  %+v", 401, o.Payload)
}

func (o *GetUserPresenceUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserPresenceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserPresenceForbidden creates a GetUserPresenceForbidden with default headers values
func NewGetUserPresenceForbidden() *GetUserPresenceForbidden {
	return &GetUserPresenceForbidden{}
}

/*GetUserPresenceForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetUserPresenceForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetUserPresenceForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/presences/{sourceId}][%d] getUserPresenceForbidden  %+v", 403, o.Payload)
}

func (o *GetUserPresenceForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserPresenceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserPresenceNotFound creates a GetUserPresenceNotFound with default headers values
func NewGetUserPresenceNotFound() *GetUserPresenceNotFound {
	return &GetUserPresenceNotFound{}
}

/*GetUserPresenceNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetUserPresenceNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetUserPresenceNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/presences/{sourceId}][%d] getUserPresenceNotFound  %+v", 404, o.Payload)
}

func (o *GetUserPresenceNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserPresenceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserPresenceRequestEntityTooLarge creates a GetUserPresenceRequestEntityTooLarge with default headers values
func NewGetUserPresenceRequestEntityTooLarge() *GetUserPresenceRequestEntityTooLarge {
	return &GetUserPresenceRequestEntityTooLarge{}
}

/*GetUserPresenceRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetUserPresenceRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetUserPresenceRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/presences/{sourceId}][%d] getUserPresenceRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetUserPresenceRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserPresenceRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserPresenceUnsupportedMediaType creates a GetUserPresenceUnsupportedMediaType with default headers values
func NewGetUserPresenceUnsupportedMediaType() *GetUserPresenceUnsupportedMediaType {
	return &GetUserPresenceUnsupportedMediaType{}
}

/*GetUserPresenceUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetUserPresenceUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetUserPresenceUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/presences/{sourceId}][%d] getUserPresenceUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetUserPresenceUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserPresenceUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserPresenceTooManyRequests creates a GetUserPresenceTooManyRequests with default headers values
func NewGetUserPresenceTooManyRequests() *GetUserPresenceTooManyRequests {
	return &GetUserPresenceTooManyRequests{}
}

/*GetUserPresenceTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetUserPresenceTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetUserPresenceTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/presences/{sourceId}][%d] getUserPresenceTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetUserPresenceTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserPresenceTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserPresenceInternalServerError creates a GetUserPresenceInternalServerError with default headers values
func NewGetUserPresenceInternalServerError() *GetUserPresenceInternalServerError {
	return &GetUserPresenceInternalServerError{}
}

/*GetUserPresenceInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetUserPresenceInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetUserPresenceInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/presences/{sourceId}][%d] getUserPresenceInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUserPresenceInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserPresenceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserPresenceServiceUnavailable creates a GetUserPresenceServiceUnavailable with default headers values
func NewGetUserPresenceServiceUnavailable() *GetUserPresenceServiceUnavailable {
	return &GetUserPresenceServiceUnavailable{}
}

/*GetUserPresenceServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetUserPresenceServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetUserPresenceServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/presences/{sourceId}][%d] getUserPresenceServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetUserPresenceServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserPresenceServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserPresenceGatewayTimeout creates a GetUserPresenceGatewayTimeout with default headers values
func NewGetUserPresenceGatewayTimeout() *GetUserPresenceGatewayTimeout {
	return &GetUserPresenceGatewayTimeout{}
}

/*GetUserPresenceGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetUserPresenceGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetUserPresenceGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/presences/{sourceId}][%d] getUserPresenceGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetUserPresenceGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserPresenceGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
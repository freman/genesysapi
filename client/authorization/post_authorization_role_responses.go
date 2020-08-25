// Code generated by go-swagger; DO NOT EDIT.

package authorization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostAuthorizationRoleReader is a Reader for the PostAuthorizationRole structure.
type PostAuthorizationRoleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAuthorizationRoleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPostAuthorizationRoleNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostAuthorizationRoleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostAuthorizationRoleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostAuthorizationRoleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostAuthorizationRoleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostAuthorizationRoleRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostAuthorizationRoleUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostAuthorizationRoleTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostAuthorizationRoleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostAuthorizationRoleServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostAuthorizationRoleGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostAuthorizationRoleNoContent creates a PostAuthorizationRoleNoContent with default headers values
func NewPostAuthorizationRoleNoContent() *PostAuthorizationRoleNoContent {
	return &PostAuthorizationRoleNoContent{}
}

/*PostAuthorizationRoleNoContent handles this case with default header values.

Bulk Grants Created
*/
type PostAuthorizationRoleNoContent struct {
}

func (o *PostAuthorizationRoleNoContent) Error() string {
	return fmt.Sprintf("[POST /api/v2/authorization/roles/{roleId}][%d] postAuthorizationRoleNoContent ", 204)
}

func (o *PostAuthorizationRoleNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostAuthorizationRoleBadRequest creates a PostAuthorizationRoleBadRequest with default headers values
func NewPostAuthorizationRoleBadRequest() *PostAuthorizationRoleBadRequest {
	return &PostAuthorizationRoleBadRequest{}
}

/*PostAuthorizationRoleBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostAuthorizationRoleBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostAuthorizationRoleBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/authorization/roles/{roleId}][%d] postAuthorizationRoleBadRequest  %+v", 400, o.Payload)
}

func (o *PostAuthorizationRoleBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAuthorizationRoleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAuthorizationRoleUnauthorized creates a PostAuthorizationRoleUnauthorized with default headers values
func NewPostAuthorizationRoleUnauthorized() *PostAuthorizationRoleUnauthorized {
	return &PostAuthorizationRoleUnauthorized{}
}

/*PostAuthorizationRoleUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostAuthorizationRoleUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostAuthorizationRoleUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/authorization/roles/{roleId}][%d] postAuthorizationRoleUnauthorized  %+v", 401, o.Payload)
}

func (o *PostAuthorizationRoleUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAuthorizationRoleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAuthorizationRoleForbidden creates a PostAuthorizationRoleForbidden with default headers values
func NewPostAuthorizationRoleForbidden() *PostAuthorizationRoleForbidden {
	return &PostAuthorizationRoleForbidden{}
}

/*PostAuthorizationRoleForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostAuthorizationRoleForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostAuthorizationRoleForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/authorization/roles/{roleId}][%d] postAuthorizationRoleForbidden  %+v", 403, o.Payload)
}

func (o *PostAuthorizationRoleForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAuthorizationRoleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAuthorizationRoleNotFound creates a PostAuthorizationRoleNotFound with default headers values
func NewPostAuthorizationRoleNotFound() *PostAuthorizationRoleNotFound {
	return &PostAuthorizationRoleNotFound{}
}

/*PostAuthorizationRoleNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostAuthorizationRoleNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostAuthorizationRoleNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/authorization/roles/{roleId}][%d] postAuthorizationRoleNotFound  %+v", 404, o.Payload)
}

func (o *PostAuthorizationRoleNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAuthorizationRoleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAuthorizationRoleRequestEntityTooLarge creates a PostAuthorizationRoleRequestEntityTooLarge with default headers values
func NewPostAuthorizationRoleRequestEntityTooLarge() *PostAuthorizationRoleRequestEntityTooLarge {
	return &PostAuthorizationRoleRequestEntityTooLarge{}
}

/*PostAuthorizationRoleRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostAuthorizationRoleRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostAuthorizationRoleRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/authorization/roles/{roleId}][%d] postAuthorizationRoleRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostAuthorizationRoleRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAuthorizationRoleRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAuthorizationRoleUnsupportedMediaType creates a PostAuthorizationRoleUnsupportedMediaType with default headers values
func NewPostAuthorizationRoleUnsupportedMediaType() *PostAuthorizationRoleUnsupportedMediaType {
	return &PostAuthorizationRoleUnsupportedMediaType{}
}

/*PostAuthorizationRoleUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostAuthorizationRoleUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostAuthorizationRoleUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/authorization/roles/{roleId}][%d] postAuthorizationRoleUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostAuthorizationRoleUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAuthorizationRoleUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAuthorizationRoleTooManyRequests creates a PostAuthorizationRoleTooManyRequests with default headers values
func NewPostAuthorizationRoleTooManyRequests() *PostAuthorizationRoleTooManyRequests {
	return &PostAuthorizationRoleTooManyRequests{}
}

/*PostAuthorizationRoleTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PostAuthorizationRoleTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostAuthorizationRoleTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/authorization/roles/{roleId}][%d] postAuthorizationRoleTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostAuthorizationRoleTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAuthorizationRoleTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAuthorizationRoleInternalServerError creates a PostAuthorizationRoleInternalServerError with default headers values
func NewPostAuthorizationRoleInternalServerError() *PostAuthorizationRoleInternalServerError {
	return &PostAuthorizationRoleInternalServerError{}
}

/*PostAuthorizationRoleInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostAuthorizationRoleInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostAuthorizationRoleInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/authorization/roles/{roleId}][%d] postAuthorizationRoleInternalServerError  %+v", 500, o.Payload)
}

func (o *PostAuthorizationRoleInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAuthorizationRoleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAuthorizationRoleServiceUnavailable creates a PostAuthorizationRoleServiceUnavailable with default headers values
func NewPostAuthorizationRoleServiceUnavailable() *PostAuthorizationRoleServiceUnavailable {
	return &PostAuthorizationRoleServiceUnavailable{}
}

/*PostAuthorizationRoleServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostAuthorizationRoleServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostAuthorizationRoleServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/authorization/roles/{roleId}][%d] postAuthorizationRoleServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostAuthorizationRoleServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAuthorizationRoleServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAuthorizationRoleGatewayTimeout creates a PostAuthorizationRoleGatewayTimeout with default headers values
func NewPostAuthorizationRoleGatewayTimeout() *PostAuthorizationRoleGatewayTimeout {
	return &PostAuthorizationRoleGatewayTimeout{}
}

/*PostAuthorizationRoleGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostAuthorizationRoleGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostAuthorizationRoleGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/authorization/roles/{roleId}][%d] postAuthorizationRoleGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostAuthorizationRoleGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAuthorizationRoleGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
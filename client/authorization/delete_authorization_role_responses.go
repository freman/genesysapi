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

// DeleteAuthorizationRoleReader is a Reader for the DeleteAuthorizationRole structure.
type DeleteAuthorizationRoleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAuthorizationRoleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 400:
		result := NewDeleteAuthorizationRoleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteAuthorizationRoleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteAuthorizationRoleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteAuthorizationRoleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteAuthorizationRoleRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteAuthorizationRoleUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteAuthorizationRoleTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteAuthorizationRoleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteAuthorizationRoleServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteAuthorizationRoleGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteAuthorizationRoleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteAuthorizationRoleBadRequest creates a DeleteAuthorizationRoleBadRequest with default headers values
func NewDeleteAuthorizationRoleBadRequest() *DeleteAuthorizationRoleBadRequest {
	return &DeleteAuthorizationRoleBadRequest{}
}

/*DeleteAuthorizationRoleBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteAuthorizationRoleBadRequest struct {
	Payload *models.ErrorBody
}

func (o *DeleteAuthorizationRoleBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/authorization/roles/{roleId}][%d] deleteAuthorizationRoleBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteAuthorizationRoleBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteAuthorizationRoleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAuthorizationRoleUnauthorized creates a DeleteAuthorizationRoleUnauthorized with default headers values
func NewDeleteAuthorizationRoleUnauthorized() *DeleteAuthorizationRoleUnauthorized {
	return &DeleteAuthorizationRoleUnauthorized{}
}

/*DeleteAuthorizationRoleUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteAuthorizationRoleUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *DeleteAuthorizationRoleUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/authorization/roles/{roleId}][%d] deleteAuthorizationRoleUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteAuthorizationRoleUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteAuthorizationRoleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAuthorizationRoleForbidden creates a DeleteAuthorizationRoleForbidden with default headers values
func NewDeleteAuthorizationRoleForbidden() *DeleteAuthorizationRoleForbidden {
	return &DeleteAuthorizationRoleForbidden{}
}

/*DeleteAuthorizationRoleForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type DeleteAuthorizationRoleForbidden struct {
	Payload *models.ErrorBody
}

func (o *DeleteAuthorizationRoleForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/authorization/roles/{roleId}][%d] deleteAuthorizationRoleForbidden  %+v", 403, o.Payload)
}

func (o *DeleteAuthorizationRoleForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteAuthorizationRoleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAuthorizationRoleNotFound creates a DeleteAuthorizationRoleNotFound with default headers values
func NewDeleteAuthorizationRoleNotFound() *DeleteAuthorizationRoleNotFound {
	return &DeleteAuthorizationRoleNotFound{}
}

/*DeleteAuthorizationRoleNotFound handles this case with default header values.

The requested resource was not found.
*/
type DeleteAuthorizationRoleNotFound struct {
	Payload *models.ErrorBody
}

func (o *DeleteAuthorizationRoleNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/authorization/roles/{roleId}][%d] deleteAuthorizationRoleNotFound  %+v", 404, o.Payload)
}

func (o *DeleteAuthorizationRoleNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteAuthorizationRoleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAuthorizationRoleRequestEntityTooLarge creates a DeleteAuthorizationRoleRequestEntityTooLarge with default headers values
func NewDeleteAuthorizationRoleRequestEntityTooLarge() *DeleteAuthorizationRoleRequestEntityTooLarge {
	return &DeleteAuthorizationRoleRequestEntityTooLarge{}
}

/*DeleteAuthorizationRoleRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type DeleteAuthorizationRoleRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *DeleteAuthorizationRoleRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/authorization/roles/{roleId}][%d] deleteAuthorizationRoleRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteAuthorizationRoleRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteAuthorizationRoleRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAuthorizationRoleUnsupportedMediaType creates a DeleteAuthorizationRoleUnsupportedMediaType with default headers values
func NewDeleteAuthorizationRoleUnsupportedMediaType() *DeleteAuthorizationRoleUnsupportedMediaType {
	return &DeleteAuthorizationRoleUnsupportedMediaType{}
}

/*DeleteAuthorizationRoleUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteAuthorizationRoleUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *DeleteAuthorizationRoleUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/authorization/roles/{roleId}][%d] deleteAuthorizationRoleUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteAuthorizationRoleUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteAuthorizationRoleUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAuthorizationRoleTooManyRequests creates a DeleteAuthorizationRoleTooManyRequests with default headers values
func NewDeleteAuthorizationRoleTooManyRequests() *DeleteAuthorizationRoleTooManyRequests {
	return &DeleteAuthorizationRoleTooManyRequests{}
}

/*DeleteAuthorizationRoleTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type DeleteAuthorizationRoleTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *DeleteAuthorizationRoleTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/authorization/roles/{roleId}][%d] deleteAuthorizationRoleTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteAuthorizationRoleTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteAuthorizationRoleTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAuthorizationRoleInternalServerError creates a DeleteAuthorizationRoleInternalServerError with default headers values
func NewDeleteAuthorizationRoleInternalServerError() *DeleteAuthorizationRoleInternalServerError {
	return &DeleteAuthorizationRoleInternalServerError{}
}

/*DeleteAuthorizationRoleInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteAuthorizationRoleInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *DeleteAuthorizationRoleInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/authorization/roles/{roleId}][%d] deleteAuthorizationRoleInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteAuthorizationRoleInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteAuthorizationRoleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAuthorizationRoleServiceUnavailable creates a DeleteAuthorizationRoleServiceUnavailable with default headers values
func NewDeleteAuthorizationRoleServiceUnavailable() *DeleteAuthorizationRoleServiceUnavailable {
	return &DeleteAuthorizationRoleServiceUnavailable{}
}

/*DeleteAuthorizationRoleServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteAuthorizationRoleServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *DeleteAuthorizationRoleServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/authorization/roles/{roleId}][%d] deleteAuthorizationRoleServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteAuthorizationRoleServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteAuthorizationRoleServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAuthorizationRoleGatewayTimeout creates a DeleteAuthorizationRoleGatewayTimeout with default headers values
func NewDeleteAuthorizationRoleGatewayTimeout() *DeleteAuthorizationRoleGatewayTimeout {
	return &DeleteAuthorizationRoleGatewayTimeout{}
}

/*DeleteAuthorizationRoleGatewayTimeout handles this case with default header values.

The request timed out.
*/
type DeleteAuthorizationRoleGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteAuthorizationRoleGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/authorization/roles/{roleId}][%d] deleteAuthorizationRoleGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteAuthorizationRoleGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteAuthorizationRoleGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAuthorizationRoleDefault creates a DeleteAuthorizationRoleDefault with default headers values
func NewDeleteAuthorizationRoleDefault(code int) *DeleteAuthorizationRoleDefault {
	return &DeleteAuthorizationRoleDefault{
		_statusCode: code,
	}
}

/*DeleteAuthorizationRoleDefault handles this case with default header values.

successful operation
*/
type DeleteAuthorizationRoleDefault struct {
	_statusCode int
}

// Code gets the status code for the delete authorization role default response
func (o *DeleteAuthorizationRoleDefault) Code() int {
	return o._statusCode
}

func (o *DeleteAuthorizationRoleDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/authorization/roles/{roleId}][%d] deleteAuthorizationRole default ", o._statusCode)
}

func (o *DeleteAuthorizationRoleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

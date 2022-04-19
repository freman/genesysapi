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

// PutAuthorizationRoleReader is a Reader for the PutAuthorizationRole structure.
type PutAuthorizationRoleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutAuthorizationRoleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutAuthorizationRoleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutAuthorizationRoleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutAuthorizationRoleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutAuthorizationRoleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutAuthorizationRoleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPutAuthorizationRoleRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPutAuthorizationRoleRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPutAuthorizationRoleUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutAuthorizationRoleTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutAuthorizationRoleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutAuthorizationRoleServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutAuthorizationRoleGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutAuthorizationRoleOK creates a PutAuthorizationRoleOK with default headers values
func NewPutAuthorizationRoleOK() *PutAuthorizationRoleOK {
	return &PutAuthorizationRoleOK{}
}

/*PutAuthorizationRoleOK handles this case with default header values.

successful operation
*/
type PutAuthorizationRoleOK struct {
	Payload *models.DomainOrganizationRole
}

func (o *PutAuthorizationRoleOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/authorization/roles/{roleId}][%d] putAuthorizationRoleOK  %+v", 200, o.Payload)
}

func (o *PutAuthorizationRoleOK) GetPayload() *models.DomainOrganizationRole {
	return o.Payload
}

func (o *PutAuthorizationRoleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DomainOrganizationRole)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAuthorizationRoleBadRequest creates a PutAuthorizationRoleBadRequest with default headers values
func NewPutAuthorizationRoleBadRequest() *PutAuthorizationRoleBadRequest {
	return &PutAuthorizationRoleBadRequest{}
}

/*PutAuthorizationRoleBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutAuthorizationRoleBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PutAuthorizationRoleBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/authorization/roles/{roleId}][%d] putAuthorizationRoleBadRequest  %+v", 400, o.Payload)
}

func (o *PutAuthorizationRoleBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutAuthorizationRoleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAuthorizationRoleUnauthorized creates a PutAuthorizationRoleUnauthorized with default headers values
func NewPutAuthorizationRoleUnauthorized() *PutAuthorizationRoleUnauthorized {
	return &PutAuthorizationRoleUnauthorized{}
}

/*PutAuthorizationRoleUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PutAuthorizationRoleUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PutAuthorizationRoleUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/authorization/roles/{roleId}][%d] putAuthorizationRoleUnauthorized  %+v", 401, o.Payload)
}

func (o *PutAuthorizationRoleUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutAuthorizationRoleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAuthorizationRoleForbidden creates a PutAuthorizationRoleForbidden with default headers values
func NewPutAuthorizationRoleForbidden() *PutAuthorizationRoleForbidden {
	return &PutAuthorizationRoleForbidden{}
}

/*PutAuthorizationRoleForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PutAuthorizationRoleForbidden struct {
	Payload *models.ErrorBody
}

func (o *PutAuthorizationRoleForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/authorization/roles/{roleId}][%d] putAuthorizationRoleForbidden  %+v", 403, o.Payload)
}

func (o *PutAuthorizationRoleForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutAuthorizationRoleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAuthorizationRoleNotFound creates a PutAuthorizationRoleNotFound with default headers values
func NewPutAuthorizationRoleNotFound() *PutAuthorizationRoleNotFound {
	return &PutAuthorizationRoleNotFound{}
}

/*PutAuthorizationRoleNotFound handles this case with default header values.

The requested resource was not found.
*/
type PutAuthorizationRoleNotFound struct {
	Payload *models.ErrorBody
}

func (o *PutAuthorizationRoleNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/authorization/roles/{roleId}][%d] putAuthorizationRoleNotFound  %+v", 404, o.Payload)
}

func (o *PutAuthorizationRoleNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutAuthorizationRoleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAuthorizationRoleRequestTimeout creates a PutAuthorizationRoleRequestTimeout with default headers values
func NewPutAuthorizationRoleRequestTimeout() *PutAuthorizationRoleRequestTimeout {
	return &PutAuthorizationRoleRequestTimeout{}
}

/*PutAuthorizationRoleRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PutAuthorizationRoleRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PutAuthorizationRoleRequestTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/authorization/roles/{roleId}][%d] putAuthorizationRoleRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutAuthorizationRoleRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutAuthorizationRoleRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAuthorizationRoleRequestEntityTooLarge creates a PutAuthorizationRoleRequestEntityTooLarge with default headers values
func NewPutAuthorizationRoleRequestEntityTooLarge() *PutAuthorizationRoleRequestEntityTooLarge {
	return &PutAuthorizationRoleRequestEntityTooLarge{}
}

/*PutAuthorizationRoleRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PutAuthorizationRoleRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PutAuthorizationRoleRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /api/v2/authorization/roles/{roleId}][%d] putAuthorizationRoleRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutAuthorizationRoleRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutAuthorizationRoleRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAuthorizationRoleUnsupportedMediaType creates a PutAuthorizationRoleUnsupportedMediaType with default headers values
func NewPutAuthorizationRoleUnsupportedMediaType() *PutAuthorizationRoleUnsupportedMediaType {
	return &PutAuthorizationRoleUnsupportedMediaType{}
}

/*PutAuthorizationRoleUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutAuthorizationRoleUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PutAuthorizationRoleUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v2/authorization/roles/{roleId}][%d] putAuthorizationRoleUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutAuthorizationRoleUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutAuthorizationRoleUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAuthorizationRoleTooManyRequests creates a PutAuthorizationRoleTooManyRequests with default headers values
func NewPutAuthorizationRoleTooManyRequests() *PutAuthorizationRoleTooManyRequests {
	return &PutAuthorizationRoleTooManyRequests{}
}

/*PutAuthorizationRoleTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PutAuthorizationRoleTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PutAuthorizationRoleTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v2/authorization/roles/{roleId}][%d] putAuthorizationRoleTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutAuthorizationRoleTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutAuthorizationRoleTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAuthorizationRoleInternalServerError creates a PutAuthorizationRoleInternalServerError with default headers values
func NewPutAuthorizationRoleInternalServerError() *PutAuthorizationRoleInternalServerError {
	return &PutAuthorizationRoleInternalServerError{}
}

/*PutAuthorizationRoleInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutAuthorizationRoleInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PutAuthorizationRoleInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/authorization/roles/{roleId}][%d] putAuthorizationRoleInternalServerError  %+v", 500, o.Payload)
}

func (o *PutAuthorizationRoleInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutAuthorizationRoleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAuthorizationRoleServiceUnavailable creates a PutAuthorizationRoleServiceUnavailable with default headers values
func NewPutAuthorizationRoleServiceUnavailable() *PutAuthorizationRoleServiceUnavailable {
	return &PutAuthorizationRoleServiceUnavailable{}
}

/*PutAuthorizationRoleServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutAuthorizationRoleServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PutAuthorizationRoleServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/authorization/roles/{roleId}][%d] putAuthorizationRoleServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutAuthorizationRoleServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutAuthorizationRoleServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAuthorizationRoleGatewayTimeout creates a PutAuthorizationRoleGatewayTimeout with default headers values
func NewPutAuthorizationRoleGatewayTimeout() *PutAuthorizationRoleGatewayTimeout {
	return &PutAuthorizationRoleGatewayTimeout{}
}

/*PutAuthorizationRoleGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PutAuthorizationRoleGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PutAuthorizationRoleGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/authorization/roles/{roleId}][%d] putAuthorizationRoleGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutAuthorizationRoleGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutAuthorizationRoleGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

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

// GetAuthorizationRoleReader is a Reader for the GetAuthorizationRole structure.
type GetAuthorizationRoleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAuthorizationRoleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAuthorizationRoleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAuthorizationRoleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAuthorizationRoleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAuthorizationRoleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAuthorizationRoleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetAuthorizationRoleRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetAuthorizationRoleUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetAuthorizationRoleTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAuthorizationRoleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetAuthorizationRoleServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetAuthorizationRoleGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAuthorizationRoleOK creates a GetAuthorizationRoleOK with default headers values
func NewGetAuthorizationRoleOK() *GetAuthorizationRoleOK {
	return &GetAuthorizationRoleOK{}
}

/*GetAuthorizationRoleOK handles this case with default header values.

successful operation
*/
type GetAuthorizationRoleOK struct {
	Payload *models.DomainOrganizationRole
}

func (o *GetAuthorizationRoleOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/roles/{roleId}][%d] getAuthorizationRoleOK  %+v", 200, o.Payload)
}

func (o *GetAuthorizationRoleOK) GetPayload() *models.DomainOrganizationRole {
	return o.Payload
}

func (o *GetAuthorizationRoleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DomainOrganizationRole)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationRoleBadRequest creates a GetAuthorizationRoleBadRequest with default headers values
func NewGetAuthorizationRoleBadRequest() *GetAuthorizationRoleBadRequest {
	return &GetAuthorizationRoleBadRequest{}
}

/*GetAuthorizationRoleBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetAuthorizationRoleBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetAuthorizationRoleBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/roles/{roleId}][%d] getAuthorizationRoleBadRequest  %+v", 400, o.Payload)
}

func (o *GetAuthorizationRoleBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationRoleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationRoleUnauthorized creates a GetAuthorizationRoleUnauthorized with default headers values
func NewGetAuthorizationRoleUnauthorized() *GetAuthorizationRoleUnauthorized {
	return &GetAuthorizationRoleUnauthorized{}
}

/*GetAuthorizationRoleUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetAuthorizationRoleUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetAuthorizationRoleUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/roles/{roleId}][%d] getAuthorizationRoleUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAuthorizationRoleUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationRoleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationRoleForbidden creates a GetAuthorizationRoleForbidden with default headers values
func NewGetAuthorizationRoleForbidden() *GetAuthorizationRoleForbidden {
	return &GetAuthorizationRoleForbidden{}
}

/*GetAuthorizationRoleForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetAuthorizationRoleForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetAuthorizationRoleForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/roles/{roleId}][%d] getAuthorizationRoleForbidden  %+v", 403, o.Payload)
}

func (o *GetAuthorizationRoleForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationRoleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationRoleNotFound creates a GetAuthorizationRoleNotFound with default headers values
func NewGetAuthorizationRoleNotFound() *GetAuthorizationRoleNotFound {
	return &GetAuthorizationRoleNotFound{}
}

/*GetAuthorizationRoleNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetAuthorizationRoleNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetAuthorizationRoleNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/roles/{roleId}][%d] getAuthorizationRoleNotFound  %+v", 404, o.Payload)
}

func (o *GetAuthorizationRoleNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationRoleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationRoleRequestEntityTooLarge creates a GetAuthorizationRoleRequestEntityTooLarge with default headers values
func NewGetAuthorizationRoleRequestEntityTooLarge() *GetAuthorizationRoleRequestEntityTooLarge {
	return &GetAuthorizationRoleRequestEntityTooLarge{}
}

/*GetAuthorizationRoleRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetAuthorizationRoleRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetAuthorizationRoleRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/roles/{roleId}][%d] getAuthorizationRoleRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetAuthorizationRoleRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationRoleRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationRoleUnsupportedMediaType creates a GetAuthorizationRoleUnsupportedMediaType with default headers values
func NewGetAuthorizationRoleUnsupportedMediaType() *GetAuthorizationRoleUnsupportedMediaType {
	return &GetAuthorizationRoleUnsupportedMediaType{}
}

/*GetAuthorizationRoleUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetAuthorizationRoleUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetAuthorizationRoleUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/roles/{roleId}][%d] getAuthorizationRoleUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetAuthorizationRoleUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationRoleUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationRoleTooManyRequests creates a GetAuthorizationRoleTooManyRequests with default headers values
func NewGetAuthorizationRoleTooManyRequests() *GetAuthorizationRoleTooManyRequests {
	return &GetAuthorizationRoleTooManyRequests{}
}

/*GetAuthorizationRoleTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetAuthorizationRoleTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetAuthorizationRoleTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/roles/{roleId}][%d] getAuthorizationRoleTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetAuthorizationRoleTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationRoleTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationRoleInternalServerError creates a GetAuthorizationRoleInternalServerError with default headers values
func NewGetAuthorizationRoleInternalServerError() *GetAuthorizationRoleInternalServerError {
	return &GetAuthorizationRoleInternalServerError{}
}

/*GetAuthorizationRoleInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetAuthorizationRoleInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetAuthorizationRoleInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/roles/{roleId}][%d] getAuthorizationRoleInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAuthorizationRoleInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationRoleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationRoleServiceUnavailable creates a GetAuthorizationRoleServiceUnavailable with default headers values
func NewGetAuthorizationRoleServiceUnavailable() *GetAuthorizationRoleServiceUnavailable {
	return &GetAuthorizationRoleServiceUnavailable{}
}

/*GetAuthorizationRoleServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetAuthorizationRoleServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetAuthorizationRoleServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/roles/{roleId}][%d] getAuthorizationRoleServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetAuthorizationRoleServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationRoleServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationRoleGatewayTimeout creates a GetAuthorizationRoleGatewayTimeout with default headers values
func NewGetAuthorizationRoleGatewayTimeout() *GetAuthorizationRoleGatewayTimeout {
	return &GetAuthorizationRoleGatewayTimeout{}
}

/*GetAuthorizationRoleGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetAuthorizationRoleGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetAuthorizationRoleGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/roles/{roleId}][%d] getAuthorizationRoleGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetAuthorizationRoleGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationRoleGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
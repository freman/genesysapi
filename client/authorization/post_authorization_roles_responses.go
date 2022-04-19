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

// PostAuthorizationRolesReader is a Reader for the PostAuthorizationRoles structure.
type PostAuthorizationRolesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAuthorizationRolesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostAuthorizationRolesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostAuthorizationRolesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostAuthorizationRolesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostAuthorizationRolesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostAuthorizationRolesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostAuthorizationRolesRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostAuthorizationRolesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostAuthorizationRolesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostAuthorizationRolesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostAuthorizationRolesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostAuthorizationRolesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostAuthorizationRolesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostAuthorizationRolesOK creates a PostAuthorizationRolesOK with default headers values
func NewPostAuthorizationRolesOK() *PostAuthorizationRolesOK {
	return &PostAuthorizationRolesOK{}
}

/*PostAuthorizationRolesOK handles this case with default header values.

successful operation
*/
type PostAuthorizationRolesOK struct {
	Payload *models.DomainOrganizationRole
}

func (o *PostAuthorizationRolesOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/authorization/roles][%d] postAuthorizationRolesOK  %+v", 200, o.Payload)
}

func (o *PostAuthorizationRolesOK) GetPayload() *models.DomainOrganizationRole {
	return o.Payload
}

func (o *PostAuthorizationRolesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DomainOrganizationRole)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAuthorizationRolesBadRequest creates a PostAuthorizationRolesBadRequest with default headers values
func NewPostAuthorizationRolesBadRequest() *PostAuthorizationRolesBadRequest {
	return &PostAuthorizationRolesBadRequest{}
}

/*PostAuthorizationRolesBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostAuthorizationRolesBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostAuthorizationRolesBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/authorization/roles][%d] postAuthorizationRolesBadRequest  %+v", 400, o.Payload)
}

func (o *PostAuthorizationRolesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAuthorizationRolesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAuthorizationRolesUnauthorized creates a PostAuthorizationRolesUnauthorized with default headers values
func NewPostAuthorizationRolesUnauthorized() *PostAuthorizationRolesUnauthorized {
	return &PostAuthorizationRolesUnauthorized{}
}

/*PostAuthorizationRolesUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostAuthorizationRolesUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostAuthorizationRolesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/authorization/roles][%d] postAuthorizationRolesUnauthorized  %+v", 401, o.Payload)
}

func (o *PostAuthorizationRolesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAuthorizationRolesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAuthorizationRolesForbidden creates a PostAuthorizationRolesForbidden with default headers values
func NewPostAuthorizationRolesForbidden() *PostAuthorizationRolesForbidden {
	return &PostAuthorizationRolesForbidden{}
}

/*PostAuthorizationRolesForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostAuthorizationRolesForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostAuthorizationRolesForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/authorization/roles][%d] postAuthorizationRolesForbidden  %+v", 403, o.Payload)
}

func (o *PostAuthorizationRolesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAuthorizationRolesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAuthorizationRolesNotFound creates a PostAuthorizationRolesNotFound with default headers values
func NewPostAuthorizationRolesNotFound() *PostAuthorizationRolesNotFound {
	return &PostAuthorizationRolesNotFound{}
}

/*PostAuthorizationRolesNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostAuthorizationRolesNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostAuthorizationRolesNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/authorization/roles][%d] postAuthorizationRolesNotFound  %+v", 404, o.Payload)
}

func (o *PostAuthorizationRolesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAuthorizationRolesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAuthorizationRolesRequestTimeout creates a PostAuthorizationRolesRequestTimeout with default headers values
func NewPostAuthorizationRolesRequestTimeout() *PostAuthorizationRolesRequestTimeout {
	return &PostAuthorizationRolesRequestTimeout{}
}

/*PostAuthorizationRolesRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostAuthorizationRolesRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostAuthorizationRolesRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/authorization/roles][%d] postAuthorizationRolesRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostAuthorizationRolesRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAuthorizationRolesRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAuthorizationRolesRequestEntityTooLarge creates a PostAuthorizationRolesRequestEntityTooLarge with default headers values
func NewPostAuthorizationRolesRequestEntityTooLarge() *PostAuthorizationRolesRequestEntityTooLarge {
	return &PostAuthorizationRolesRequestEntityTooLarge{}
}

/*PostAuthorizationRolesRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PostAuthorizationRolesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostAuthorizationRolesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/authorization/roles][%d] postAuthorizationRolesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostAuthorizationRolesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAuthorizationRolesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAuthorizationRolesUnsupportedMediaType creates a PostAuthorizationRolesUnsupportedMediaType with default headers values
func NewPostAuthorizationRolesUnsupportedMediaType() *PostAuthorizationRolesUnsupportedMediaType {
	return &PostAuthorizationRolesUnsupportedMediaType{}
}

/*PostAuthorizationRolesUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostAuthorizationRolesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostAuthorizationRolesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/authorization/roles][%d] postAuthorizationRolesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostAuthorizationRolesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAuthorizationRolesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAuthorizationRolesTooManyRequests creates a PostAuthorizationRolesTooManyRequests with default headers values
func NewPostAuthorizationRolesTooManyRequests() *PostAuthorizationRolesTooManyRequests {
	return &PostAuthorizationRolesTooManyRequests{}
}

/*PostAuthorizationRolesTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostAuthorizationRolesTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostAuthorizationRolesTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/authorization/roles][%d] postAuthorizationRolesTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostAuthorizationRolesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAuthorizationRolesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAuthorizationRolesInternalServerError creates a PostAuthorizationRolesInternalServerError with default headers values
func NewPostAuthorizationRolesInternalServerError() *PostAuthorizationRolesInternalServerError {
	return &PostAuthorizationRolesInternalServerError{}
}

/*PostAuthorizationRolesInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostAuthorizationRolesInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostAuthorizationRolesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/authorization/roles][%d] postAuthorizationRolesInternalServerError  %+v", 500, o.Payload)
}

func (o *PostAuthorizationRolesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAuthorizationRolesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAuthorizationRolesServiceUnavailable creates a PostAuthorizationRolesServiceUnavailable with default headers values
func NewPostAuthorizationRolesServiceUnavailable() *PostAuthorizationRolesServiceUnavailable {
	return &PostAuthorizationRolesServiceUnavailable{}
}

/*PostAuthorizationRolesServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostAuthorizationRolesServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostAuthorizationRolesServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/authorization/roles][%d] postAuthorizationRolesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostAuthorizationRolesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAuthorizationRolesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAuthorizationRolesGatewayTimeout creates a PostAuthorizationRolesGatewayTimeout with default headers values
func NewPostAuthorizationRolesGatewayTimeout() *PostAuthorizationRolesGatewayTimeout {
	return &PostAuthorizationRolesGatewayTimeout{}
}

/*PostAuthorizationRolesGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostAuthorizationRolesGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostAuthorizationRolesGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/authorization/roles][%d] postAuthorizationRolesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostAuthorizationRolesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAuthorizationRolesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

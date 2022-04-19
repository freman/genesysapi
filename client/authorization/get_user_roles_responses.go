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

// GetUserRolesReader is a Reader for the GetUserRoles structure.
type GetUserRolesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserRolesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUserRolesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetUserRolesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetUserRolesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetUserRolesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetUserRolesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetUserRolesRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetUserRolesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetUserRolesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetUserRolesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUserRolesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetUserRolesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetUserRolesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUserRolesOK creates a GetUserRolesOK with default headers values
func NewGetUserRolesOK() *GetUserRolesOK {
	return &GetUserRolesOK{}
}

/*GetUserRolesOK handles this case with default header values.

successful operation
*/
type GetUserRolesOK struct {
	Payload *models.UserAuthorization
}

func (o *GetUserRolesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/roles][%d] getUserRolesOK  %+v", 200, o.Payload)
}

func (o *GetUserRolesOK) GetPayload() *models.UserAuthorization {
	return o.Payload
}

func (o *GetUserRolesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserAuthorization)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserRolesBadRequest creates a GetUserRolesBadRequest with default headers values
func NewGetUserRolesBadRequest() *GetUserRolesBadRequest {
	return &GetUserRolesBadRequest{}
}

/*GetUserRolesBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetUserRolesBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetUserRolesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/roles][%d] getUserRolesBadRequest  %+v", 400, o.Payload)
}

func (o *GetUserRolesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserRolesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserRolesUnauthorized creates a GetUserRolesUnauthorized with default headers values
func NewGetUserRolesUnauthorized() *GetUserRolesUnauthorized {
	return &GetUserRolesUnauthorized{}
}

/*GetUserRolesUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetUserRolesUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetUserRolesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/roles][%d] getUserRolesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetUserRolesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserRolesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserRolesForbidden creates a GetUserRolesForbidden with default headers values
func NewGetUserRolesForbidden() *GetUserRolesForbidden {
	return &GetUserRolesForbidden{}
}

/*GetUserRolesForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetUserRolesForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetUserRolesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/roles][%d] getUserRolesForbidden  %+v", 403, o.Payload)
}

func (o *GetUserRolesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserRolesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserRolesNotFound creates a GetUserRolesNotFound with default headers values
func NewGetUserRolesNotFound() *GetUserRolesNotFound {
	return &GetUserRolesNotFound{}
}

/*GetUserRolesNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetUserRolesNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetUserRolesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/roles][%d] getUserRolesNotFound  %+v", 404, o.Payload)
}

func (o *GetUserRolesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserRolesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserRolesRequestTimeout creates a GetUserRolesRequestTimeout with default headers values
func NewGetUserRolesRequestTimeout() *GetUserRolesRequestTimeout {
	return &GetUserRolesRequestTimeout{}
}

/*GetUserRolesRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetUserRolesRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetUserRolesRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/roles][%d] getUserRolesRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetUserRolesRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserRolesRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserRolesRequestEntityTooLarge creates a GetUserRolesRequestEntityTooLarge with default headers values
func NewGetUserRolesRequestEntityTooLarge() *GetUserRolesRequestEntityTooLarge {
	return &GetUserRolesRequestEntityTooLarge{}
}

/*GetUserRolesRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetUserRolesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetUserRolesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/roles][%d] getUserRolesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetUserRolesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserRolesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserRolesUnsupportedMediaType creates a GetUserRolesUnsupportedMediaType with default headers values
func NewGetUserRolesUnsupportedMediaType() *GetUserRolesUnsupportedMediaType {
	return &GetUserRolesUnsupportedMediaType{}
}

/*GetUserRolesUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetUserRolesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetUserRolesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/roles][%d] getUserRolesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetUserRolesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserRolesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserRolesTooManyRequests creates a GetUserRolesTooManyRequests with default headers values
func NewGetUserRolesTooManyRequests() *GetUserRolesTooManyRequests {
	return &GetUserRolesTooManyRequests{}
}

/*GetUserRolesTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetUserRolesTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetUserRolesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/roles][%d] getUserRolesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetUserRolesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserRolesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserRolesInternalServerError creates a GetUserRolesInternalServerError with default headers values
func NewGetUserRolesInternalServerError() *GetUserRolesInternalServerError {
	return &GetUserRolesInternalServerError{}
}

/*GetUserRolesInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetUserRolesInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetUserRolesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/roles][%d] getUserRolesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUserRolesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserRolesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserRolesServiceUnavailable creates a GetUserRolesServiceUnavailable with default headers values
func NewGetUserRolesServiceUnavailable() *GetUserRolesServiceUnavailable {
	return &GetUserRolesServiceUnavailable{}
}

/*GetUserRolesServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetUserRolesServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetUserRolesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/roles][%d] getUserRolesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetUserRolesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserRolesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserRolesGatewayTimeout creates a GetUserRolesGatewayTimeout with default headers values
func NewGetUserRolesGatewayTimeout() *GetUserRolesGatewayTimeout {
	return &GetUserRolesGatewayTimeout{}
}

/*GetUserRolesGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetUserRolesGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetUserRolesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/roles][%d] getUserRolesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetUserRolesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserRolesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

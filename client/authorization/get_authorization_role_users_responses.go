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

// GetAuthorizationRoleUsersReader is a Reader for the GetAuthorizationRoleUsers structure.
type GetAuthorizationRoleUsersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAuthorizationRoleUsersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAuthorizationRoleUsersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAuthorizationRoleUsersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAuthorizationRoleUsersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAuthorizationRoleUsersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAuthorizationRoleUsersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetAuthorizationRoleUsersRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetAuthorizationRoleUsersRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetAuthorizationRoleUsersUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetAuthorizationRoleUsersTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAuthorizationRoleUsersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetAuthorizationRoleUsersServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetAuthorizationRoleUsersGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAuthorizationRoleUsersOK creates a GetAuthorizationRoleUsersOK with default headers values
func NewGetAuthorizationRoleUsersOK() *GetAuthorizationRoleUsersOK {
	return &GetAuthorizationRoleUsersOK{}
}

/*GetAuthorizationRoleUsersOK handles this case with default header values.

successful operation
*/
type GetAuthorizationRoleUsersOK struct {
	Payload *models.UserEntityListing
}

func (o *GetAuthorizationRoleUsersOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/roles/{roleId}/users][%d] getAuthorizationRoleUsersOK  %+v", 200, o.Payload)
}

func (o *GetAuthorizationRoleUsersOK) GetPayload() *models.UserEntityListing {
	return o.Payload
}

func (o *GetAuthorizationRoleUsersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationRoleUsersBadRequest creates a GetAuthorizationRoleUsersBadRequest with default headers values
func NewGetAuthorizationRoleUsersBadRequest() *GetAuthorizationRoleUsersBadRequest {
	return &GetAuthorizationRoleUsersBadRequest{}
}

/*GetAuthorizationRoleUsersBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetAuthorizationRoleUsersBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetAuthorizationRoleUsersBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/roles/{roleId}/users][%d] getAuthorizationRoleUsersBadRequest  %+v", 400, o.Payload)
}

func (o *GetAuthorizationRoleUsersBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationRoleUsersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationRoleUsersUnauthorized creates a GetAuthorizationRoleUsersUnauthorized with default headers values
func NewGetAuthorizationRoleUsersUnauthorized() *GetAuthorizationRoleUsersUnauthorized {
	return &GetAuthorizationRoleUsersUnauthorized{}
}

/*GetAuthorizationRoleUsersUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetAuthorizationRoleUsersUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetAuthorizationRoleUsersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/roles/{roleId}/users][%d] getAuthorizationRoleUsersUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAuthorizationRoleUsersUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationRoleUsersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationRoleUsersForbidden creates a GetAuthorizationRoleUsersForbidden with default headers values
func NewGetAuthorizationRoleUsersForbidden() *GetAuthorizationRoleUsersForbidden {
	return &GetAuthorizationRoleUsersForbidden{}
}

/*GetAuthorizationRoleUsersForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetAuthorizationRoleUsersForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetAuthorizationRoleUsersForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/roles/{roleId}/users][%d] getAuthorizationRoleUsersForbidden  %+v", 403, o.Payload)
}

func (o *GetAuthorizationRoleUsersForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationRoleUsersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationRoleUsersNotFound creates a GetAuthorizationRoleUsersNotFound with default headers values
func NewGetAuthorizationRoleUsersNotFound() *GetAuthorizationRoleUsersNotFound {
	return &GetAuthorizationRoleUsersNotFound{}
}

/*GetAuthorizationRoleUsersNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetAuthorizationRoleUsersNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetAuthorizationRoleUsersNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/roles/{roleId}/users][%d] getAuthorizationRoleUsersNotFound  %+v", 404, o.Payload)
}

func (o *GetAuthorizationRoleUsersNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationRoleUsersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationRoleUsersRequestTimeout creates a GetAuthorizationRoleUsersRequestTimeout with default headers values
func NewGetAuthorizationRoleUsersRequestTimeout() *GetAuthorizationRoleUsersRequestTimeout {
	return &GetAuthorizationRoleUsersRequestTimeout{}
}

/*GetAuthorizationRoleUsersRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetAuthorizationRoleUsersRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetAuthorizationRoleUsersRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/roles/{roleId}/users][%d] getAuthorizationRoleUsersRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetAuthorizationRoleUsersRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationRoleUsersRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationRoleUsersRequestEntityTooLarge creates a GetAuthorizationRoleUsersRequestEntityTooLarge with default headers values
func NewGetAuthorizationRoleUsersRequestEntityTooLarge() *GetAuthorizationRoleUsersRequestEntityTooLarge {
	return &GetAuthorizationRoleUsersRequestEntityTooLarge{}
}

/*GetAuthorizationRoleUsersRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetAuthorizationRoleUsersRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetAuthorizationRoleUsersRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/roles/{roleId}/users][%d] getAuthorizationRoleUsersRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetAuthorizationRoleUsersRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationRoleUsersRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationRoleUsersUnsupportedMediaType creates a GetAuthorizationRoleUsersUnsupportedMediaType with default headers values
func NewGetAuthorizationRoleUsersUnsupportedMediaType() *GetAuthorizationRoleUsersUnsupportedMediaType {
	return &GetAuthorizationRoleUsersUnsupportedMediaType{}
}

/*GetAuthorizationRoleUsersUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetAuthorizationRoleUsersUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetAuthorizationRoleUsersUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/roles/{roleId}/users][%d] getAuthorizationRoleUsersUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetAuthorizationRoleUsersUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationRoleUsersUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationRoleUsersTooManyRequests creates a GetAuthorizationRoleUsersTooManyRequests with default headers values
func NewGetAuthorizationRoleUsersTooManyRequests() *GetAuthorizationRoleUsersTooManyRequests {
	return &GetAuthorizationRoleUsersTooManyRequests{}
}

/*GetAuthorizationRoleUsersTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetAuthorizationRoleUsersTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetAuthorizationRoleUsersTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/roles/{roleId}/users][%d] getAuthorizationRoleUsersTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetAuthorizationRoleUsersTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationRoleUsersTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationRoleUsersInternalServerError creates a GetAuthorizationRoleUsersInternalServerError with default headers values
func NewGetAuthorizationRoleUsersInternalServerError() *GetAuthorizationRoleUsersInternalServerError {
	return &GetAuthorizationRoleUsersInternalServerError{}
}

/*GetAuthorizationRoleUsersInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetAuthorizationRoleUsersInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetAuthorizationRoleUsersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/roles/{roleId}/users][%d] getAuthorizationRoleUsersInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAuthorizationRoleUsersInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationRoleUsersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationRoleUsersServiceUnavailable creates a GetAuthorizationRoleUsersServiceUnavailable with default headers values
func NewGetAuthorizationRoleUsersServiceUnavailable() *GetAuthorizationRoleUsersServiceUnavailable {
	return &GetAuthorizationRoleUsersServiceUnavailable{}
}

/*GetAuthorizationRoleUsersServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetAuthorizationRoleUsersServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetAuthorizationRoleUsersServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/roles/{roleId}/users][%d] getAuthorizationRoleUsersServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetAuthorizationRoleUsersServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationRoleUsersServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationRoleUsersGatewayTimeout creates a GetAuthorizationRoleUsersGatewayTimeout with default headers values
func NewGetAuthorizationRoleUsersGatewayTimeout() *GetAuthorizationRoleUsersGatewayTimeout {
	return &GetAuthorizationRoleUsersGatewayTimeout{}
}

/*GetAuthorizationRoleUsersGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetAuthorizationRoleUsersGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetAuthorizationRoleUsersGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/roles/{roleId}/users][%d] getAuthorizationRoleUsersGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetAuthorizationRoleUsersGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationRoleUsersGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

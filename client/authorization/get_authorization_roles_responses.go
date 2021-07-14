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

// GetAuthorizationRolesReader is a Reader for the GetAuthorizationRoles structure.
type GetAuthorizationRolesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAuthorizationRolesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAuthorizationRolesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAuthorizationRolesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAuthorizationRolesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAuthorizationRolesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAuthorizationRolesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetAuthorizationRolesRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetAuthorizationRolesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetAuthorizationRolesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetAuthorizationRolesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAuthorizationRolesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetAuthorizationRolesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetAuthorizationRolesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAuthorizationRolesOK creates a GetAuthorizationRolesOK with default headers values
func NewGetAuthorizationRolesOK() *GetAuthorizationRolesOK {
	return &GetAuthorizationRolesOK{}
}

/*GetAuthorizationRolesOK handles this case with default header values.

successful operation
*/
type GetAuthorizationRolesOK struct {
	Payload *models.OrganizationRoleEntityListing
}

func (o *GetAuthorizationRolesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/roles][%d] getAuthorizationRolesOK  %+v", 200, o.Payload)
}

func (o *GetAuthorizationRolesOK) GetPayload() *models.OrganizationRoleEntityListing {
	return o.Payload
}

func (o *GetAuthorizationRolesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OrganizationRoleEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationRolesBadRequest creates a GetAuthorizationRolesBadRequest with default headers values
func NewGetAuthorizationRolesBadRequest() *GetAuthorizationRolesBadRequest {
	return &GetAuthorizationRolesBadRequest{}
}

/*GetAuthorizationRolesBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetAuthorizationRolesBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetAuthorizationRolesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/roles][%d] getAuthorizationRolesBadRequest  %+v", 400, o.Payload)
}

func (o *GetAuthorizationRolesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationRolesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationRolesUnauthorized creates a GetAuthorizationRolesUnauthorized with default headers values
func NewGetAuthorizationRolesUnauthorized() *GetAuthorizationRolesUnauthorized {
	return &GetAuthorizationRolesUnauthorized{}
}

/*GetAuthorizationRolesUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetAuthorizationRolesUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetAuthorizationRolesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/roles][%d] getAuthorizationRolesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAuthorizationRolesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationRolesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationRolesForbidden creates a GetAuthorizationRolesForbidden with default headers values
func NewGetAuthorizationRolesForbidden() *GetAuthorizationRolesForbidden {
	return &GetAuthorizationRolesForbidden{}
}

/*GetAuthorizationRolesForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetAuthorizationRolesForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetAuthorizationRolesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/roles][%d] getAuthorizationRolesForbidden  %+v", 403, o.Payload)
}

func (o *GetAuthorizationRolesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationRolesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationRolesNotFound creates a GetAuthorizationRolesNotFound with default headers values
func NewGetAuthorizationRolesNotFound() *GetAuthorizationRolesNotFound {
	return &GetAuthorizationRolesNotFound{}
}

/*GetAuthorizationRolesNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetAuthorizationRolesNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetAuthorizationRolesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/roles][%d] getAuthorizationRolesNotFound  %+v", 404, o.Payload)
}

func (o *GetAuthorizationRolesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationRolesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationRolesRequestTimeout creates a GetAuthorizationRolesRequestTimeout with default headers values
func NewGetAuthorizationRolesRequestTimeout() *GetAuthorizationRolesRequestTimeout {
	return &GetAuthorizationRolesRequestTimeout{}
}

/*GetAuthorizationRolesRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetAuthorizationRolesRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetAuthorizationRolesRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/roles][%d] getAuthorizationRolesRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetAuthorizationRolesRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationRolesRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationRolesRequestEntityTooLarge creates a GetAuthorizationRolesRequestEntityTooLarge with default headers values
func NewGetAuthorizationRolesRequestEntityTooLarge() *GetAuthorizationRolesRequestEntityTooLarge {
	return &GetAuthorizationRolesRequestEntityTooLarge{}
}

/*GetAuthorizationRolesRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetAuthorizationRolesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetAuthorizationRolesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/roles][%d] getAuthorizationRolesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetAuthorizationRolesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationRolesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationRolesUnsupportedMediaType creates a GetAuthorizationRolesUnsupportedMediaType with default headers values
func NewGetAuthorizationRolesUnsupportedMediaType() *GetAuthorizationRolesUnsupportedMediaType {
	return &GetAuthorizationRolesUnsupportedMediaType{}
}

/*GetAuthorizationRolesUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetAuthorizationRolesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetAuthorizationRolesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/roles][%d] getAuthorizationRolesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetAuthorizationRolesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationRolesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationRolesTooManyRequests creates a GetAuthorizationRolesTooManyRequests with default headers values
func NewGetAuthorizationRolesTooManyRequests() *GetAuthorizationRolesTooManyRequests {
	return &GetAuthorizationRolesTooManyRequests{}
}

/*GetAuthorizationRolesTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetAuthorizationRolesTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetAuthorizationRolesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/roles][%d] getAuthorizationRolesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetAuthorizationRolesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationRolesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationRolesInternalServerError creates a GetAuthorizationRolesInternalServerError with default headers values
func NewGetAuthorizationRolesInternalServerError() *GetAuthorizationRolesInternalServerError {
	return &GetAuthorizationRolesInternalServerError{}
}

/*GetAuthorizationRolesInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetAuthorizationRolesInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetAuthorizationRolesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/roles][%d] getAuthorizationRolesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAuthorizationRolesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationRolesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationRolesServiceUnavailable creates a GetAuthorizationRolesServiceUnavailable with default headers values
func NewGetAuthorizationRolesServiceUnavailable() *GetAuthorizationRolesServiceUnavailable {
	return &GetAuthorizationRolesServiceUnavailable{}
}

/*GetAuthorizationRolesServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetAuthorizationRolesServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetAuthorizationRolesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/roles][%d] getAuthorizationRolesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetAuthorizationRolesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationRolesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationRolesGatewayTimeout creates a GetAuthorizationRolesGatewayTimeout with default headers values
func NewGetAuthorizationRolesGatewayTimeout() *GetAuthorizationRolesGatewayTimeout {
	return &GetAuthorizationRolesGatewayTimeout{}
}

/*GetAuthorizationRolesGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetAuthorizationRolesGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetAuthorizationRolesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/roles][%d] getAuthorizationRolesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetAuthorizationRolesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationRolesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

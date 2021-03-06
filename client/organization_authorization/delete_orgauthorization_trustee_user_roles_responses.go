// Code generated by go-swagger; DO NOT EDIT.

package organization_authorization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// DeleteOrgauthorizationTrusteeUserRolesReader is a Reader for the DeleteOrgauthorizationTrusteeUserRoles structure.
type DeleteOrgauthorizationTrusteeUserRolesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteOrgauthorizationTrusteeUserRolesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteOrgauthorizationTrusteeUserRolesNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteOrgauthorizationTrusteeUserRolesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteOrgauthorizationTrusteeUserRolesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteOrgauthorizationTrusteeUserRolesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteOrgauthorizationTrusteeUserRolesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteOrgauthorizationTrusteeUserRolesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteOrgauthorizationTrusteeUserRolesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteOrgauthorizationTrusteeUserRolesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteOrgauthorizationTrusteeUserRolesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteOrgauthorizationTrusteeUserRolesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteOrgauthorizationTrusteeUserRolesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteOrgauthorizationTrusteeUserRolesNoContent creates a DeleteOrgauthorizationTrusteeUserRolesNoContent with default headers values
func NewDeleteOrgauthorizationTrusteeUserRolesNoContent() *DeleteOrgauthorizationTrusteeUserRolesNoContent {
	return &DeleteOrgauthorizationTrusteeUserRolesNoContent{}
}

/*DeleteOrgauthorizationTrusteeUserRolesNoContent handles this case with default header values.

Roles deleted
*/
type DeleteOrgauthorizationTrusteeUserRolesNoContent struct {
}

func (o *DeleteOrgauthorizationTrusteeUserRolesNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/orgauthorization/trustees/{trusteeOrgId}/users/{trusteeUserId}/roles][%d] deleteOrgauthorizationTrusteeUserRolesNoContent ", 204)
}

func (o *DeleteOrgauthorizationTrusteeUserRolesNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteOrgauthorizationTrusteeUserRolesBadRequest creates a DeleteOrgauthorizationTrusteeUserRolesBadRequest with default headers values
func NewDeleteOrgauthorizationTrusteeUserRolesBadRequest() *DeleteOrgauthorizationTrusteeUserRolesBadRequest {
	return &DeleteOrgauthorizationTrusteeUserRolesBadRequest{}
}

/*DeleteOrgauthorizationTrusteeUserRolesBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteOrgauthorizationTrusteeUserRolesBadRequest struct {
	Payload *models.ErrorBody
}

func (o *DeleteOrgauthorizationTrusteeUserRolesBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/orgauthorization/trustees/{trusteeOrgId}/users/{trusteeUserId}/roles][%d] deleteOrgauthorizationTrusteeUserRolesBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteOrgauthorizationTrusteeUserRolesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOrgauthorizationTrusteeUserRolesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOrgauthorizationTrusteeUserRolesUnauthorized creates a DeleteOrgauthorizationTrusteeUserRolesUnauthorized with default headers values
func NewDeleteOrgauthorizationTrusteeUserRolesUnauthorized() *DeleteOrgauthorizationTrusteeUserRolesUnauthorized {
	return &DeleteOrgauthorizationTrusteeUserRolesUnauthorized{}
}

/*DeleteOrgauthorizationTrusteeUserRolesUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteOrgauthorizationTrusteeUserRolesUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *DeleteOrgauthorizationTrusteeUserRolesUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/orgauthorization/trustees/{trusteeOrgId}/users/{trusteeUserId}/roles][%d] deleteOrgauthorizationTrusteeUserRolesUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteOrgauthorizationTrusteeUserRolesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOrgauthorizationTrusteeUserRolesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOrgauthorizationTrusteeUserRolesForbidden creates a DeleteOrgauthorizationTrusteeUserRolesForbidden with default headers values
func NewDeleteOrgauthorizationTrusteeUserRolesForbidden() *DeleteOrgauthorizationTrusteeUserRolesForbidden {
	return &DeleteOrgauthorizationTrusteeUserRolesForbidden{}
}

/*DeleteOrgauthorizationTrusteeUserRolesForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type DeleteOrgauthorizationTrusteeUserRolesForbidden struct {
	Payload *models.ErrorBody
}

func (o *DeleteOrgauthorizationTrusteeUserRolesForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/orgauthorization/trustees/{trusteeOrgId}/users/{trusteeUserId}/roles][%d] deleteOrgauthorizationTrusteeUserRolesForbidden  %+v", 403, o.Payload)
}

func (o *DeleteOrgauthorizationTrusteeUserRolesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOrgauthorizationTrusteeUserRolesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOrgauthorizationTrusteeUserRolesNotFound creates a DeleteOrgauthorizationTrusteeUserRolesNotFound with default headers values
func NewDeleteOrgauthorizationTrusteeUserRolesNotFound() *DeleteOrgauthorizationTrusteeUserRolesNotFound {
	return &DeleteOrgauthorizationTrusteeUserRolesNotFound{}
}

/*DeleteOrgauthorizationTrusteeUserRolesNotFound handles this case with default header values.

The requested resource was not found.
*/
type DeleteOrgauthorizationTrusteeUserRolesNotFound struct {
	Payload *models.ErrorBody
}

func (o *DeleteOrgauthorizationTrusteeUserRolesNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/orgauthorization/trustees/{trusteeOrgId}/users/{trusteeUserId}/roles][%d] deleteOrgauthorizationTrusteeUserRolesNotFound  %+v", 404, o.Payload)
}

func (o *DeleteOrgauthorizationTrusteeUserRolesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOrgauthorizationTrusteeUserRolesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOrgauthorizationTrusteeUserRolesRequestEntityTooLarge creates a DeleteOrgauthorizationTrusteeUserRolesRequestEntityTooLarge with default headers values
func NewDeleteOrgauthorizationTrusteeUserRolesRequestEntityTooLarge() *DeleteOrgauthorizationTrusteeUserRolesRequestEntityTooLarge {
	return &DeleteOrgauthorizationTrusteeUserRolesRequestEntityTooLarge{}
}

/*DeleteOrgauthorizationTrusteeUserRolesRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type DeleteOrgauthorizationTrusteeUserRolesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *DeleteOrgauthorizationTrusteeUserRolesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/orgauthorization/trustees/{trusteeOrgId}/users/{trusteeUserId}/roles][%d] deleteOrgauthorizationTrusteeUserRolesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteOrgauthorizationTrusteeUserRolesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOrgauthorizationTrusteeUserRolesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOrgauthorizationTrusteeUserRolesUnsupportedMediaType creates a DeleteOrgauthorizationTrusteeUserRolesUnsupportedMediaType with default headers values
func NewDeleteOrgauthorizationTrusteeUserRolesUnsupportedMediaType() *DeleteOrgauthorizationTrusteeUserRolesUnsupportedMediaType {
	return &DeleteOrgauthorizationTrusteeUserRolesUnsupportedMediaType{}
}

/*DeleteOrgauthorizationTrusteeUserRolesUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteOrgauthorizationTrusteeUserRolesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *DeleteOrgauthorizationTrusteeUserRolesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/orgauthorization/trustees/{trusteeOrgId}/users/{trusteeUserId}/roles][%d] deleteOrgauthorizationTrusteeUserRolesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteOrgauthorizationTrusteeUserRolesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOrgauthorizationTrusteeUserRolesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOrgauthorizationTrusteeUserRolesTooManyRequests creates a DeleteOrgauthorizationTrusteeUserRolesTooManyRequests with default headers values
func NewDeleteOrgauthorizationTrusteeUserRolesTooManyRequests() *DeleteOrgauthorizationTrusteeUserRolesTooManyRequests {
	return &DeleteOrgauthorizationTrusteeUserRolesTooManyRequests{}
}

/*DeleteOrgauthorizationTrusteeUserRolesTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type DeleteOrgauthorizationTrusteeUserRolesTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *DeleteOrgauthorizationTrusteeUserRolesTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/orgauthorization/trustees/{trusteeOrgId}/users/{trusteeUserId}/roles][%d] deleteOrgauthorizationTrusteeUserRolesTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteOrgauthorizationTrusteeUserRolesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOrgauthorizationTrusteeUserRolesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOrgauthorizationTrusteeUserRolesInternalServerError creates a DeleteOrgauthorizationTrusteeUserRolesInternalServerError with default headers values
func NewDeleteOrgauthorizationTrusteeUserRolesInternalServerError() *DeleteOrgauthorizationTrusteeUserRolesInternalServerError {
	return &DeleteOrgauthorizationTrusteeUserRolesInternalServerError{}
}

/*DeleteOrgauthorizationTrusteeUserRolesInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteOrgauthorizationTrusteeUserRolesInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *DeleteOrgauthorizationTrusteeUserRolesInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/orgauthorization/trustees/{trusteeOrgId}/users/{trusteeUserId}/roles][%d] deleteOrgauthorizationTrusteeUserRolesInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteOrgauthorizationTrusteeUserRolesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOrgauthorizationTrusteeUserRolesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOrgauthorizationTrusteeUserRolesServiceUnavailable creates a DeleteOrgauthorizationTrusteeUserRolesServiceUnavailable with default headers values
func NewDeleteOrgauthorizationTrusteeUserRolesServiceUnavailable() *DeleteOrgauthorizationTrusteeUserRolesServiceUnavailable {
	return &DeleteOrgauthorizationTrusteeUserRolesServiceUnavailable{}
}

/*DeleteOrgauthorizationTrusteeUserRolesServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteOrgauthorizationTrusteeUserRolesServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *DeleteOrgauthorizationTrusteeUserRolesServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/orgauthorization/trustees/{trusteeOrgId}/users/{trusteeUserId}/roles][%d] deleteOrgauthorizationTrusteeUserRolesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteOrgauthorizationTrusteeUserRolesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOrgauthorizationTrusteeUserRolesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOrgauthorizationTrusteeUserRolesGatewayTimeout creates a DeleteOrgauthorizationTrusteeUserRolesGatewayTimeout with default headers values
func NewDeleteOrgauthorizationTrusteeUserRolesGatewayTimeout() *DeleteOrgauthorizationTrusteeUserRolesGatewayTimeout {
	return &DeleteOrgauthorizationTrusteeUserRolesGatewayTimeout{}
}

/*DeleteOrgauthorizationTrusteeUserRolesGatewayTimeout handles this case with default header values.

The request timed out.
*/
type DeleteOrgauthorizationTrusteeUserRolesGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteOrgauthorizationTrusteeUserRolesGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/orgauthorization/trustees/{trusteeOrgId}/users/{trusteeUserId}/roles][%d] deleteOrgauthorizationTrusteeUserRolesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteOrgauthorizationTrusteeUserRolesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOrgauthorizationTrusteeUserRolesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

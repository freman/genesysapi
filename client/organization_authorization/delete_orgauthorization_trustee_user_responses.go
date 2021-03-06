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

// DeleteOrgauthorizationTrusteeUserReader is a Reader for the DeleteOrgauthorizationTrusteeUser structure.
type DeleteOrgauthorizationTrusteeUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteOrgauthorizationTrusteeUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteOrgauthorizationTrusteeUserNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteOrgauthorizationTrusteeUserBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteOrgauthorizationTrusteeUserUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteOrgauthorizationTrusteeUserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteOrgauthorizationTrusteeUserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteOrgauthorizationTrusteeUserRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteOrgauthorizationTrusteeUserUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteOrgauthorizationTrusteeUserTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteOrgauthorizationTrusteeUserInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteOrgauthorizationTrusteeUserServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteOrgauthorizationTrusteeUserGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteOrgauthorizationTrusteeUserNoContent creates a DeleteOrgauthorizationTrusteeUserNoContent with default headers values
func NewDeleteOrgauthorizationTrusteeUserNoContent() *DeleteOrgauthorizationTrusteeUserNoContent {
	return &DeleteOrgauthorizationTrusteeUserNoContent{}
}

/*DeleteOrgauthorizationTrusteeUserNoContent handles this case with default header values.

Trust deleted
*/
type DeleteOrgauthorizationTrusteeUserNoContent struct {
}

func (o *DeleteOrgauthorizationTrusteeUserNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/orgauthorization/trustees/{trusteeOrgId}/users/{trusteeUserId}][%d] deleteOrgauthorizationTrusteeUserNoContent ", 204)
}

func (o *DeleteOrgauthorizationTrusteeUserNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteOrgauthorizationTrusteeUserBadRequest creates a DeleteOrgauthorizationTrusteeUserBadRequest with default headers values
func NewDeleteOrgauthorizationTrusteeUserBadRequest() *DeleteOrgauthorizationTrusteeUserBadRequest {
	return &DeleteOrgauthorizationTrusteeUserBadRequest{}
}

/*DeleteOrgauthorizationTrusteeUserBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteOrgauthorizationTrusteeUserBadRequest struct {
	Payload *models.ErrorBody
}

func (o *DeleteOrgauthorizationTrusteeUserBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/orgauthorization/trustees/{trusteeOrgId}/users/{trusteeUserId}][%d] deleteOrgauthorizationTrusteeUserBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteOrgauthorizationTrusteeUserBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOrgauthorizationTrusteeUserBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOrgauthorizationTrusteeUserUnauthorized creates a DeleteOrgauthorizationTrusteeUserUnauthorized with default headers values
func NewDeleteOrgauthorizationTrusteeUserUnauthorized() *DeleteOrgauthorizationTrusteeUserUnauthorized {
	return &DeleteOrgauthorizationTrusteeUserUnauthorized{}
}

/*DeleteOrgauthorizationTrusteeUserUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteOrgauthorizationTrusteeUserUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *DeleteOrgauthorizationTrusteeUserUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/orgauthorization/trustees/{trusteeOrgId}/users/{trusteeUserId}][%d] deleteOrgauthorizationTrusteeUserUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteOrgauthorizationTrusteeUserUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOrgauthorizationTrusteeUserUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOrgauthorizationTrusteeUserForbidden creates a DeleteOrgauthorizationTrusteeUserForbidden with default headers values
func NewDeleteOrgauthorizationTrusteeUserForbidden() *DeleteOrgauthorizationTrusteeUserForbidden {
	return &DeleteOrgauthorizationTrusteeUserForbidden{}
}

/*DeleteOrgauthorizationTrusteeUserForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type DeleteOrgauthorizationTrusteeUserForbidden struct {
	Payload *models.ErrorBody
}

func (o *DeleteOrgauthorizationTrusteeUserForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/orgauthorization/trustees/{trusteeOrgId}/users/{trusteeUserId}][%d] deleteOrgauthorizationTrusteeUserForbidden  %+v", 403, o.Payload)
}

func (o *DeleteOrgauthorizationTrusteeUserForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOrgauthorizationTrusteeUserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOrgauthorizationTrusteeUserNotFound creates a DeleteOrgauthorizationTrusteeUserNotFound with default headers values
func NewDeleteOrgauthorizationTrusteeUserNotFound() *DeleteOrgauthorizationTrusteeUserNotFound {
	return &DeleteOrgauthorizationTrusteeUserNotFound{}
}

/*DeleteOrgauthorizationTrusteeUserNotFound handles this case with default header values.

The requested resource was not found.
*/
type DeleteOrgauthorizationTrusteeUserNotFound struct {
	Payload *models.ErrorBody
}

func (o *DeleteOrgauthorizationTrusteeUserNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/orgauthorization/trustees/{trusteeOrgId}/users/{trusteeUserId}][%d] deleteOrgauthorizationTrusteeUserNotFound  %+v", 404, o.Payload)
}

func (o *DeleteOrgauthorizationTrusteeUserNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOrgauthorizationTrusteeUserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOrgauthorizationTrusteeUserRequestEntityTooLarge creates a DeleteOrgauthorizationTrusteeUserRequestEntityTooLarge with default headers values
func NewDeleteOrgauthorizationTrusteeUserRequestEntityTooLarge() *DeleteOrgauthorizationTrusteeUserRequestEntityTooLarge {
	return &DeleteOrgauthorizationTrusteeUserRequestEntityTooLarge{}
}

/*DeleteOrgauthorizationTrusteeUserRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type DeleteOrgauthorizationTrusteeUserRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *DeleteOrgauthorizationTrusteeUserRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/orgauthorization/trustees/{trusteeOrgId}/users/{trusteeUserId}][%d] deleteOrgauthorizationTrusteeUserRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteOrgauthorizationTrusteeUserRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOrgauthorizationTrusteeUserRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOrgauthorizationTrusteeUserUnsupportedMediaType creates a DeleteOrgauthorizationTrusteeUserUnsupportedMediaType with default headers values
func NewDeleteOrgauthorizationTrusteeUserUnsupportedMediaType() *DeleteOrgauthorizationTrusteeUserUnsupportedMediaType {
	return &DeleteOrgauthorizationTrusteeUserUnsupportedMediaType{}
}

/*DeleteOrgauthorizationTrusteeUserUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteOrgauthorizationTrusteeUserUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *DeleteOrgauthorizationTrusteeUserUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/orgauthorization/trustees/{trusteeOrgId}/users/{trusteeUserId}][%d] deleteOrgauthorizationTrusteeUserUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteOrgauthorizationTrusteeUserUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOrgauthorizationTrusteeUserUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOrgauthorizationTrusteeUserTooManyRequests creates a DeleteOrgauthorizationTrusteeUserTooManyRequests with default headers values
func NewDeleteOrgauthorizationTrusteeUserTooManyRequests() *DeleteOrgauthorizationTrusteeUserTooManyRequests {
	return &DeleteOrgauthorizationTrusteeUserTooManyRequests{}
}

/*DeleteOrgauthorizationTrusteeUserTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type DeleteOrgauthorizationTrusteeUserTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *DeleteOrgauthorizationTrusteeUserTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/orgauthorization/trustees/{trusteeOrgId}/users/{trusteeUserId}][%d] deleteOrgauthorizationTrusteeUserTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteOrgauthorizationTrusteeUserTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOrgauthorizationTrusteeUserTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOrgauthorizationTrusteeUserInternalServerError creates a DeleteOrgauthorizationTrusteeUserInternalServerError with default headers values
func NewDeleteOrgauthorizationTrusteeUserInternalServerError() *DeleteOrgauthorizationTrusteeUserInternalServerError {
	return &DeleteOrgauthorizationTrusteeUserInternalServerError{}
}

/*DeleteOrgauthorizationTrusteeUserInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteOrgauthorizationTrusteeUserInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *DeleteOrgauthorizationTrusteeUserInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/orgauthorization/trustees/{trusteeOrgId}/users/{trusteeUserId}][%d] deleteOrgauthorizationTrusteeUserInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteOrgauthorizationTrusteeUserInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOrgauthorizationTrusteeUserInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOrgauthorizationTrusteeUserServiceUnavailable creates a DeleteOrgauthorizationTrusteeUserServiceUnavailable with default headers values
func NewDeleteOrgauthorizationTrusteeUserServiceUnavailable() *DeleteOrgauthorizationTrusteeUserServiceUnavailable {
	return &DeleteOrgauthorizationTrusteeUserServiceUnavailable{}
}

/*DeleteOrgauthorizationTrusteeUserServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteOrgauthorizationTrusteeUserServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *DeleteOrgauthorizationTrusteeUserServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/orgauthorization/trustees/{trusteeOrgId}/users/{trusteeUserId}][%d] deleteOrgauthorizationTrusteeUserServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteOrgauthorizationTrusteeUserServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOrgauthorizationTrusteeUserServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOrgauthorizationTrusteeUserGatewayTimeout creates a DeleteOrgauthorizationTrusteeUserGatewayTimeout with default headers values
func NewDeleteOrgauthorizationTrusteeUserGatewayTimeout() *DeleteOrgauthorizationTrusteeUserGatewayTimeout {
	return &DeleteOrgauthorizationTrusteeUserGatewayTimeout{}
}

/*DeleteOrgauthorizationTrusteeUserGatewayTimeout handles this case with default header values.

The request timed out.
*/
type DeleteOrgauthorizationTrusteeUserGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteOrgauthorizationTrusteeUserGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/orgauthorization/trustees/{trusteeOrgId}/users/{trusteeUserId}][%d] deleteOrgauthorizationTrusteeUserGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteOrgauthorizationTrusteeUserGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOrgauthorizationTrusteeUserGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

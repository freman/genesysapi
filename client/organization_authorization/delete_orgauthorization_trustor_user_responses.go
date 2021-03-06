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

// DeleteOrgauthorizationTrustorUserReader is a Reader for the DeleteOrgauthorizationTrustorUser structure.
type DeleteOrgauthorizationTrustorUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteOrgauthorizationTrustorUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteOrgauthorizationTrustorUserNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteOrgauthorizationTrustorUserBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteOrgauthorizationTrustorUserUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteOrgauthorizationTrustorUserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteOrgauthorizationTrustorUserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteOrgauthorizationTrustorUserRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteOrgauthorizationTrustorUserUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteOrgauthorizationTrustorUserTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteOrgauthorizationTrustorUserInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteOrgauthorizationTrustorUserServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteOrgauthorizationTrustorUserGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteOrgauthorizationTrustorUserNoContent creates a DeleteOrgauthorizationTrustorUserNoContent with default headers values
func NewDeleteOrgauthorizationTrustorUserNoContent() *DeleteOrgauthorizationTrustorUserNoContent {
	return &DeleteOrgauthorizationTrustorUserNoContent{}
}

/*DeleteOrgauthorizationTrustorUserNoContent handles this case with default header values.

Trust deleted
*/
type DeleteOrgauthorizationTrustorUserNoContent struct {
}

func (o *DeleteOrgauthorizationTrustorUserNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/orgauthorization/trustors/{trustorOrgId}/users/{trusteeUserId}][%d] deleteOrgauthorizationTrustorUserNoContent ", 204)
}

func (o *DeleteOrgauthorizationTrustorUserNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteOrgauthorizationTrustorUserBadRequest creates a DeleteOrgauthorizationTrustorUserBadRequest with default headers values
func NewDeleteOrgauthorizationTrustorUserBadRequest() *DeleteOrgauthorizationTrustorUserBadRequest {
	return &DeleteOrgauthorizationTrustorUserBadRequest{}
}

/*DeleteOrgauthorizationTrustorUserBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteOrgauthorizationTrustorUserBadRequest struct {
	Payload *models.ErrorBody
}

func (o *DeleteOrgauthorizationTrustorUserBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/orgauthorization/trustors/{trustorOrgId}/users/{trusteeUserId}][%d] deleteOrgauthorizationTrustorUserBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteOrgauthorizationTrustorUserBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOrgauthorizationTrustorUserBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOrgauthorizationTrustorUserUnauthorized creates a DeleteOrgauthorizationTrustorUserUnauthorized with default headers values
func NewDeleteOrgauthorizationTrustorUserUnauthorized() *DeleteOrgauthorizationTrustorUserUnauthorized {
	return &DeleteOrgauthorizationTrustorUserUnauthorized{}
}

/*DeleteOrgauthorizationTrustorUserUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteOrgauthorizationTrustorUserUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *DeleteOrgauthorizationTrustorUserUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/orgauthorization/trustors/{trustorOrgId}/users/{trusteeUserId}][%d] deleteOrgauthorizationTrustorUserUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteOrgauthorizationTrustorUserUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOrgauthorizationTrustorUserUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOrgauthorizationTrustorUserForbidden creates a DeleteOrgauthorizationTrustorUserForbidden with default headers values
func NewDeleteOrgauthorizationTrustorUserForbidden() *DeleteOrgauthorizationTrustorUserForbidden {
	return &DeleteOrgauthorizationTrustorUserForbidden{}
}

/*DeleteOrgauthorizationTrustorUserForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type DeleteOrgauthorizationTrustorUserForbidden struct {
	Payload *models.ErrorBody
}

func (o *DeleteOrgauthorizationTrustorUserForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/orgauthorization/trustors/{trustorOrgId}/users/{trusteeUserId}][%d] deleteOrgauthorizationTrustorUserForbidden  %+v", 403, o.Payload)
}

func (o *DeleteOrgauthorizationTrustorUserForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOrgauthorizationTrustorUserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOrgauthorizationTrustorUserNotFound creates a DeleteOrgauthorizationTrustorUserNotFound with default headers values
func NewDeleteOrgauthorizationTrustorUserNotFound() *DeleteOrgauthorizationTrustorUserNotFound {
	return &DeleteOrgauthorizationTrustorUserNotFound{}
}

/*DeleteOrgauthorizationTrustorUserNotFound handles this case with default header values.

The requested resource was not found.
*/
type DeleteOrgauthorizationTrustorUserNotFound struct {
	Payload *models.ErrorBody
}

func (o *DeleteOrgauthorizationTrustorUserNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/orgauthorization/trustors/{trustorOrgId}/users/{trusteeUserId}][%d] deleteOrgauthorizationTrustorUserNotFound  %+v", 404, o.Payload)
}

func (o *DeleteOrgauthorizationTrustorUserNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOrgauthorizationTrustorUserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOrgauthorizationTrustorUserRequestEntityTooLarge creates a DeleteOrgauthorizationTrustorUserRequestEntityTooLarge with default headers values
func NewDeleteOrgauthorizationTrustorUserRequestEntityTooLarge() *DeleteOrgauthorizationTrustorUserRequestEntityTooLarge {
	return &DeleteOrgauthorizationTrustorUserRequestEntityTooLarge{}
}

/*DeleteOrgauthorizationTrustorUserRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type DeleteOrgauthorizationTrustorUserRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *DeleteOrgauthorizationTrustorUserRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/orgauthorization/trustors/{trustorOrgId}/users/{trusteeUserId}][%d] deleteOrgauthorizationTrustorUserRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteOrgauthorizationTrustorUserRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOrgauthorizationTrustorUserRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOrgauthorizationTrustorUserUnsupportedMediaType creates a DeleteOrgauthorizationTrustorUserUnsupportedMediaType with default headers values
func NewDeleteOrgauthorizationTrustorUserUnsupportedMediaType() *DeleteOrgauthorizationTrustorUserUnsupportedMediaType {
	return &DeleteOrgauthorizationTrustorUserUnsupportedMediaType{}
}

/*DeleteOrgauthorizationTrustorUserUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteOrgauthorizationTrustorUserUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *DeleteOrgauthorizationTrustorUserUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/orgauthorization/trustors/{trustorOrgId}/users/{trusteeUserId}][%d] deleteOrgauthorizationTrustorUserUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteOrgauthorizationTrustorUserUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOrgauthorizationTrustorUserUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOrgauthorizationTrustorUserTooManyRequests creates a DeleteOrgauthorizationTrustorUserTooManyRequests with default headers values
func NewDeleteOrgauthorizationTrustorUserTooManyRequests() *DeleteOrgauthorizationTrustorUserTooManyRequests {
	return &DeleteOrgauthorizationTrustorUserTooManyRequests{}
}

/*DeleteOrgauthorizationTrustorUserTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type DeleteOrgauthorizationTrustorUserTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *DeleteOrgauthorizationTrustorUserTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/orgauthorization/trustors/{trustorOrgId}/users/{trusteeUserId}][%d] deleteOrgauthorizationTrustorUserTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteOrgauthorizationTrustorUserTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOrgauthorizationTrustorUserTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOrgauthorizationTrustorUserInternalServerError creates a DeleteOrgauthorizationTrustorUserInternalServerError with default headers values
func NewDeleteOrgauthorizationTrustorUserInternalServerError() *DeleteOrgauthorizationTrustorUserInternalServerError {
	return &DeleteOrgauthorizationTrustorUserInternalServerError{}
}

/*DeleteOrgauthorizationTrustorUserInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteOrgauthorizationTrustorUserInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *DeleteOrgauthorizationTrustorUserInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/orgauthorization/trustors/{trustorOrgId}/users/{trusteeUserId}][%d] deleteOrgauthorizationTrustorUserInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteOrgauthorizationTrustorUserInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOrgauthorizationTrustorUserInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOrgauthorizationTrustorUserServiceUnavailable creates a DeleteOrgauthorizationTrustorUserServiceUnavailable with default headers values
func NewDeleteOrgauthorizationTrustorUserServiceUnavailable() *DeleteOrgauthorizationTrustorUserServiceUnavailable {
	return &DeleteOrgauthorizationTrustorUserServiceUnavailable{}
}

/*DeleteOrgauthorizationTrustorUserServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteOrgauthorizationTrustorUserServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *DeleteOrgauthorizationTrustorUserServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/orgauthorization/trustors/{trustorOrgId}/users/{trusteeUserId}][%d] deleteOrgauthorizationTrustorUserServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteOrgauthorizationTrustorUserServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOrgauthorizationTrustorUserServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOrgauthorizationTrustorUserGatewayTimeout creates a DeleteOrgauthorizationTrustorUserGatewayTimeout with default headers values
func NewDeleteOrgauthorizationTrustorUserGatewayTimeout() *DeleteOrgauthorizationTrustorUserGatewayTimeout {
	return &DeleteOrgauthorizationTrustorUserGatewayTimeout{}
}

/*DeleteOrgauthorizationTrustorUserGatewayTimeout handles this case with default header values.

The request timed out.
*/
type DeleteOrgauthorizationTrustorUserGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteOrgauthorizationTrustorUserGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/orgauthorization/trustors/{trustorOrgId}/users/{trusteeUserId}][%d] deleteOrgauthorizationTrustorUserGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteOrgauthorizationTrustorUserGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOrgauthorizationTrustorUserGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

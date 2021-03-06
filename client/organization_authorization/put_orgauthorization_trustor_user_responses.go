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

// PutOrgauthorizationTrustorUserReader is a Reader for the PutOrgauthorizationTrustorUser structure.
type PutOrgauthorizationTrustorUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutOrgauthorizationTrustorUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutOrgauthorizationTrustorUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutOrgauthorizationTrustorUserBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutOrgauthorizationTrustorUserUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutOrgauthorizationTrustorUserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutOrgauthorizationTrustorUserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPutOrgauthorizationTrustorUserRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPutOrgauthorizationTrustorUserUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutOrgauthorizationTrustorUserTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutOrgauthorizationTrustorUserInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutOrgauthorizationTrustorUserServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutOrgauthorizationTrustorUserGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutOrgauthorizationTrustorUserOK creates a PutOrgauthorizationTrustorUserOK with default headers values
func NewPutOrgauthorizationTrustorUserOK() *PutOrgauthorizationTrustorUserOK {
	return &PutOrgauthorizationTrustorUserOK{}
}

/*PutOrgauthorizationTrustorUserOK handles this case with default header values.

successful operation
*/
type PutOrgauthorizationTrustorUserOK struct {
	Payload *models.TrustUser
}

func (o *PutOrgauthorizationTrustorUserOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/orgauthorization/trustors/{trustorOrgId}/users/{trusteeUserId}][%d] putOrgauthorizationTrustorUserOK  %+v", 200, o.Payload)
}

func (o *PutOrgauthorizationTrustorUserOK) GetPayload() *models.TrustUser {
	return o.Payload
}

func (o *PutOrgauthorizationTrustorUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TrustUser)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOrgauthorizationTrustorUserBadRequest creates a PutOrgauthorizationTrustorUserBadRequest with default headers values
func NewPutOrgauthorizationTrustorUserBadRequest() *PutOrgauthorizationTrustorUserBadRequest {
	return &PutOrgauthorizationTrustorUserBadRequest{}
}

/*PutOrgauthorizationTrustorUserBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutOrgauthorizationTrustorUserBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PutOrgauthorizationTrustorUserBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/orgauthorization/trustors/{trustorOrgId}/users/{trusteeUserId}][%d] putOrgauthorizationTrustorUserBadRequest  %+v", 400, o.Payload)
}

func (o *PutOrgauthorizationTrustorUserBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOrgauthorizationTrustorUserBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOrgauthorizationTrustorUserUnauthorized creates a PutOrgauthorizationTrustorUserUnauthorized with default headers values
func NewPutOrgauthorizationTrustorUserUnauthorized() *PutOrgauthorizationTrustorUserUnauthorized {
	return &PutOrgauthorizationTrustorUserUnauthorized{}
}

/*PutOrgauthorizationTrustorUserUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PutOrgauthorizationTrustorUserUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PutOrgauthorizationTrustorUserUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/orgauthorization/trustors/{trustorOrgId}/users/{trusteeUserId}][%d] putOrgauthorizationTrustorUserUnauthorized  %+v", 401, o.Payload)
}

func (o *PutOrgauthorizationTrustorUserUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOrgauthorizationTrustorUserUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOrgauthorizationTrustorUserForbidden creates a PutOrgauthorizationTrustorUserForbidden with default headers values
func NewPutOrgauthorizationTrustorUserForbidden() *PutOrgauthorizationTrustorUserForbidden {
	return &PutOrgauthorizationTrustorUserForbidden{}
}

/*PutOrgauthorizationTrustorUserForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PutOrgauthorizationTrustorUserForbidden struct {
	Payload *models.ErrorBody
}

func (o *PutOrgauthorizationTrustorUserForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/orgauthorization/trustors/{trustorOrgId}/users/{trusteeUserId}][%d] putOrgauthorizationTrustorUserForbidden  %+v", 403, o.Payload)
}

func (o *PutOrgauthorizationTrustorUserForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOrgauthorizationTrustorUserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOrgauthorizationTrustorUserNotFound creates a PutOrgauthorizationTrustorUserNotFound with default headers values
func NewPutOrgauthorizationTrustorUserNotFound() *PutOrgauthorizationTrustorUserNotFound {
	return &PutOrgauthorizationTrustorUserNotFound{}
}

/*PutOrgauthorizationTrustorUserNotFound handles this case with default header values.

The requested resource was not found.
*/
type PutOrgauthorizationTrustorUserNotFound struct {
	Payload *models.ErrorBody
}

func (o *PutOrgauthorizationTrustorUserNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/orgauthorization/trustors/{trustorOrgId}/users/{trusteeUserId}][%d] putOrgauthorizationTrustorUserNotFound  %+v", 404, o.Payload)
}

func (o *PutOrgauthorizationTrustorUserNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOrgauthorizationTrustorUserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOrgauthorizationTrustorUserRequestEntityTooLarge creates a PutOrgauthorizationTrustorUserRequestEntityTooLarge with default headers values
func NewPutOrgauthorizationTrustorUserRequestEntityTooLarge() *PutOrgauthorizationTrustorUserRequestEntityTooLarge {
	return &PutOrgauthorizationTrustorUserRequestEntityTooLarge{}
}

/*PutOrgauthorizationTrustorUserRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PutOrgauthorizationTrustorUserRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PutOrgauthorizationTrustorUserRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /api/v2/orgauthorization/trustors/{trustorOrgId}/users/{trusteeUserId}][%d] putOrgauthorizationTrustorUserRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutOrgauthorizationTrustorUserRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOrgauthorizationTrustorUserRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOrgauthorizationTrustorUserUnsupportedMediaType creates a PutOrgauthorizationTrustorUserUnsupportedMediaType with default headers values
func NewPutOrgauthorizationTrustorUserUnsupportedMediaType() *PutOrgauthorizationTrustorUserUnsupportedMediaType {
	return &PutOrgauthorizationTrustorUserUnsupportedMediaType{}
}

/*PutOrgauthorizationTrustorUserUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutOrgauthorizationTrustorUserUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PutOrgauthorizationTrustorUserUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v2/orgauthorization/trustors/{trustorOrgId}/users/{trusteeUserId}][%d] putOrgauthorizationTrustorUserUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutOrgauthorizationTrustorUserUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOrgauthorizationTrustorUserUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOrgauthorizationTrustorUserTooManyRequests creates a PutOrgauthorizationTrustorUserTooManyRequests with default headers values
func NewPutOrgauthorizationTrustorUserTooManyRequests() *PutOrgauthorizationTrustorUserTooManyRequests {
	return &PutOrgauthorizationTrustorUserTooManyRequests{}
}

/*PutOrgauthorizationTrustorUserTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PutOrgauthorizationTrustorUserTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PutOrgauthorizationTrustorUserTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v2/orgauthorization/trustors/{trustorOrgId}/users/{trusteeUserId}][%d] putOrgauthorizationTrustorUserTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutOrgauthorizationTrustorUserTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOrgauthorizationTrustorUserTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOrgauthorizationTrustorUserInternalServerError creates a PutOrgauthorizationTrustorUserInternalServerError with default headers values
func NewPutOrgauthorizationTrustorUserInternalServerError() *PutOrgauthorizationTrustorUserInternalServerError {
	return &PutOrgauthorizationTrustorUserInternalServerError{}
}

/*PutOrgauthorizationTrustorUserInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutOrgauthorizationTrustorUserInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PutOrgauthorizationTrustorUserInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/orgauthorization/trustors/{trustorOrgId}/users/{trusteeUserId}][%d] putOrgauthorizationTrustorUserInternalServerError  %+v", 500, o.Payload)
}

func (o *PutOrgauthorizationTrustorUserInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOrgauthorizationTrustorUserInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOrgauthorizationTrustorUserServiceUnavailable creates a PutOrgauthorizationTrustorUserServiceUnavailable with default headers values
func NewPutOrgauthorizationTrustorUserServiceUnavailable() *PutOrgauthorizationTrustorUserServiceUnavailable {
	return &PutOrgauthorizationTrustorUserServiceUnavailable{}
}

/*PutOrgauthorizationTrustorUserServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutOrgauthorizationTrustorUserServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PutOrgauthorizationTrustorUserServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/orgauthorization/trustors/{trustorOrgId}/users/{trusteeUserId}][%d] putOrgauthorizationTrustorUserServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutOrgauthorizationTrustorUserServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOrgauthorizationTrustorUserServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOrgauthorizationTrustorUserGatewayTimeout creates a PutOrgauthorizationTrustorUserGatewayTimeout with default headers values
func NewPutOrgauthorizationTrustorUserGatewayTimeout() *PutOrgauthorizationTrustorUserGatewayTimeout {
	return &PutOrgauthorizationTrustorUserGatewayTimeout{}
}

/*PutOrgauthorizationTrustorUserGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PutOrgauthorizationTrustorUserGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PutOrgauthorizationTrustorUserGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/orgauthorization/trustors/{trustorOrgId}/users/{trusteeUserId}][%d] putOrgauthorizationTrustorUserGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutOrgauthorizationTrustorUserGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOrgauthorizationTrustorUserGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

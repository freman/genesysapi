// Code generated by go-swagger; DO NOT EDIT.

package s_c_i_m

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PutScimUserReader is a Reader for the PutScimUser structure.
type PutScimUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutScimUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutScimUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutScimUserBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutScimUserUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutScimUserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutScimUserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPutScimUserConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPutScimUserRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPutScimUserUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutScimUserTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutScimUserInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutScimUserServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutScimUserGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutScimUserOK creates a PutScimUserOK with default headers values
func NewPutScimUserOK() *PutScimUserOK {
	return &PutScimUserOK{}
}

/*PutScimUserOK handles this case with default header values.

successful operation
*/
type PutScimUserOK struct {
	Payload *models.ScimV2User
}

func (o *PutScimUserOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/scim/users/{userId}][%d] putScimUserOK  %+v", 200, o.Payload)
}

func (o *PutScimUserOK) GetPayload() *models.ScimV2User {
	return o.Payload
}

func (o *PutScimUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ScimV2User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutScimUserBadRequest creates a PutScimUserBadRequest with default headers values
func NewPutScimUserBadRequest() *PutScimUserBadRequest {
	return &PutScimUserBadRequest{}
}

/*PutScimUserBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutScimUserBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PutScimUserBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/scim/users/{userId}][%d] putScimUserBadRequest  %+v", 400, o.Payload)
}

func (o *PutScimUserBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutScimUserBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutScimUserUnauthorized creates a PutScimUserUnauthorized with default headers values
func NewPutScimUserUnauthorized() *PutScimUserUnauthorized {
	return &PutScimUserUnauthorized{}
}

/*PutScimUserUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PutScimUserUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PutScimUserUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/scim/users/{userId}][%d] putScimUserUnauthorized  %+v", 401, o.Payload)
}

func (o *PutScimUserUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutScimUserUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutScimUserForbidden creates a PutScimUserForbidden with default headers values
func NewPutScimUserForbidden() *PutScimUserForbidden {
	return &PutScimUserForbidden{}
}

/*PutScimUserForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PutScimUserForbidden struct {
	Payload *models.ErrorBody
}

func (o *PutScimUserForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/scim/users/{userId}][%d] putScimUserForbidden  %+v", 403, o.Payload)
}

func (o *PutScimUserForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutScimUserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutScimUserNotFound creates a PutScimUserNotFound with default headers values
func NewPutScimUserNotFound() *PutScimUserNotFound {
	return &PutScimUserNotFound{}
}

/*PutScimUserNotFound handles this case with default header values.

The requested resource was not found.
*/
type PutScimUserNotFound struct {
	Payload *models.ErrorBody
}

func (o *PutScimUserNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/scim/users/{userId}][%d] putScimUserNotFound  %+v", 404, o.Payload)
}

func (o *PutScimUserNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutScimUserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutScimUserConflict creates a PutScimUserConflict with default headers values
func NewPutScimUserConflict() *PutScimUserConflict {
	return &PutScimUserConflict{}
}

/*PutScimUserConflict handles this case with default header values.

Version does not match current version.
*/
type PutScimUserConflict struct {
	Payload *models.ScimError
}

func (o *PutScimUserConflict) Error() string {
	return fmt.Sprintf("[PUT /api/v2/scim/users/{userId}][%d] putScimUserConflict  %+v", 409, o.Payload)
}

func (o *PutScimUserConflict) GetPayload() *models.ScimError {
	return o.Payload
}

func (o *PutScimUserConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ScimError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutScimUserRequestEntityTooLarge creates a PutScimUserRequestEntityTooLarge with default headers values
func NewPutScimUserRequestEntityTooLarge() *PutScimUserRequestEntityTooLarge {
	return &PutScimUserRequestEntityTooLarge{}
}

/*PutScimUserRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PutScimUserRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PutScimUserRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /api/v2/scim/users/{userId}][%d] putScimUserRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutScimUserRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutScimUserRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutScimUserUnsupportedMediaType creates a PutScimUserUnsupportedMediaType with default headers values
func NewPutScimUserUnsupportedMediaType() *PutScimUserUnsupportedMediaType {
	return &PutScimUserUnsupportedMediaType{}
}

/*PutScimUserUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutScimUserUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PutScimUserUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v2/scim/users/{userId}][%d] putScimUserUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutScimUserUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutScimUserUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutScimUserTooManyRequests creates a PutScimUserTooManyRequests with default headers values
func NewPutScimUserTooManyRequests() *PutScimUserTooManyRequests {
	return &PutScimUserTooManyRequests{}
}

/*PutScimUserTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PutScimUserTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PutScimUserTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v2/scim/users/{userId}][%d] putScimUserTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutScimUserTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutScimUserTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutScimUserInternalServerError creates a PutScimUserInternalServerError with default headers values
func NewPutScimUserInternalServerError() *PutScimUserInternalServerError {
	return &PutScimUserInternalServerError{}
}

/*PutScimUserInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutScimUserInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PutScimUserInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/scim/users/{userId}][%d] putScimUserInternalServerError  %+v", 500, o.Payload)
}

func (o *PutScimUserInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutScimUserInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutScimUserServiceUnavailable creates a PutScimUserServiceUnavailable with default headers values
func NewPutScimUserServiceUnavailable() *PutScimUserServiceUnavailable {
	return &PutScimUserServiceUnavailable{}
}

/*PutScimUserServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutScimUserServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PutScimUserServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/scim/users/{userId}][%d] putScimUserServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutScimUserServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutScimUserServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutScimUserGatewayTimeout creates a PutScimUserGatewayTimeout with default headers values
func NewPutScimUserGatewayTimeout() *PutScimUserGatewayTimeout {
	return &PutScimUserGatewayTimeout{}
}

/*PutScimUserGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PutScimUserGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PutScimUserGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/scim/users/{userId}][%d] putScimUserGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutScimUserGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutScimUserGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetUserReader is a Reader for the GetUser structure.
type GetUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetUserBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetUserUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetUserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetUserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetUserRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetUserRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetUserUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetUserTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUserInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetUserServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetUserGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUserOK creates a GetUserOK with default headers values
func NewGetUserOK() *GetUserOK {
	return &GetUserOK{}
}

/*GetUserOK handles this case with default header values.

successful operation
*/
type GetUserOK struct {
	Payload *models.User
}

func (o *GetUserOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}][%d] getUserOK  %+v", 200, o.Payload)
}

func (o *GetUserOK) GetPayload() *models.User {
	return o.Payload
}

func (o *GetUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserBadRequest creates a GetUserBadRequest with default headers values
func NewGetUserBadRequest() *GetUserBadRequest {
	return &GetUserBadRequest{}
}

/*GetUserBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetUserBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetUserBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}][%d] getUserBadRequest  %+v", 400, o.Payload)
}

func (o *GetUserBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserUnauthorized creates a GetUserUnauthorized with default headers values
func NewGetUserUnauthorized() *GetUserUnauthorized {
	return &GetUserUnauthorized{}
}

/*GetUserUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetUserUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetUserUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}][%d] getUserUnauthorized  %+v", 401, o.Payload)
}

func (o *GetUserUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserForbidden creates a GetUserForbidden with default headers values
func NewGetUserForbidden() *GetUserForbidden {
	return &GetUserForbidden{}
}

/*GetUserForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetUserForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetUserForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}][%d] getUserForbidden  %+v", 403, o.Payload)
}

func (o *GetUserForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserNotFound creates a GetUserNotFound with default headers values
func NewGetUserNotFound() *GetUserNotFound {
	return &GetUserNotFound{}
}

/*GetUserNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetUserNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetUserNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}][%d] getUserNotFound  %+v", 404, o.Payload)
}

func (o *GetUserNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserRequestTimeout creates a GetUserRequestTimeout with default headers values
func NewGetUserRequestTimeout() *GetUserRequestTimeout {
	return &GetUserRequestTimeout{}
}

/*GetUserRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetUserRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetUserRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}][%d] getUserRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetUserRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserRequestEntityTooLarge creates a GetUserRequestEntityTooLarge with default headers values
func NewGetUserRequestEntityTooLarge() *GetUserRequestEntityTooLarge {
	return &GetUserRequestEntityTooLarge{}
}

/*GetUserRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetUserRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetUserRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}][%d] getUserRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetUserRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserUnsupportedMediaType creates a GetUserUnsupportedMediaType with default headers values
func NewGetUserUnsupportedMediaType() *GetUserUnsupportedMediaType {
	return &GetUserUnsupportedMediaType{}
}

/*GetUserUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetUserUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetUserUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}][%d] getUserUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetUserUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserTooManyRequests creates a GetUserTooManyRequests with default headers values
func NewGetUserTooManyRequests() *GetUserTooManyRequests {
	return &GetUserTooManyRequests{}
}

/*GetUserTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetUserTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetUserTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}][%d] getUserTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetUserTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserInternalServerError creates a GetUserInternalServerError with default headers values
func NewGetUserInternalServerError() *GetUserInternalServerError {
	return &GetUserInternalServerError{}
}

/*GetUserInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetUserInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetUserInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}][%d] getUserInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUserInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserServiceUnavailable creates a GetUserServiceUnavailable with default headers values
func NewGetUserServiceUnavailable() *GetUserServiceUnavailable {
	return &GetUserServiceUnavailable{}
}

/*GetUserServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetUserServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetUserServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}][%d] getUserServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetUserServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserGatewayTimeout creates a GetUserGatewayTimeout with default headers values
func NewGetUserGatewayTimeout() *GetUserGatewayTimeout {
	return &GetUserGatewayTimeout{}
}

/*GetUserGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetUserGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetUserGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}][%d] getUserGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetUserGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

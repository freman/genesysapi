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

// GetProfilesUsersReader is a Reader for the GetProfilesUsers structure.
type GetProfilesUsersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProfilesUsersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProfilesUsersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetProfilesUsersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetProfilesUsersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetProfilesUsersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetProfilesUsersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetProfilesUsersRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetProfilesUsersUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetProfilesUsersTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetProfilesUsersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetProfilesUsersServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetProfilesUsersGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetProfilesUsersOK creates a GetProfilesUsersOK with default headers values
func NewGetProfilesUsersOK() *GetProfilesUsersOK {
	return &GetProfilesUsersOK{}
}

/*GetProfilesUsersOK handles this case with default header values.

successful operation
*/
type GetProfilesUsersOK struct {
	Payload *models.UserProfileEntityListing
}

func (o *GetProfilesUsersOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/profiles/users][%d] getProfilesUsersOK  %+v", 200, o.Payload)
}

func (o *GetProfilesUsersOK) GetPayload() *models.UserProfileEntityListing {
	return o.Payload
}

func (o *GetProfilesUsersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserProfileEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProfilesUsersBadRequest creates a GetProfilesUsersBadRequest with default headers values
func NewGetProfilesUsersBadRequest() *GetProfilesUsersBadRequest {
	return &GetProfilesUsersBadRequest{}
}

/*GetProfilesUsersBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetProfilesUsersBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetProfilesUsersBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/profiles/users][%d] getProfilesUsersBadRequest  %+v", 400, o.Payload)
}

func (o *GetProfilesUsersBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetProfilesUsersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProfilesUsersUnauthorized creates a GetProfilesUsersUnauthorized with default headers values
func NewGetProfilesUsersUnauthorized() *GetProfilesUsersUnauthorized {
	return &GetProfilesUsersUnauthorized{}
}

/*GetProfilesUsersUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetProfilesUsersUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetProfilesUsersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/profiles/users][%d] getProfilesUsersUnauthorized  %+v", 401, o.Payload)
}

func (o *GetProfilesUsersUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetProfilesUsersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProfilesUsersForbidden creates a GetProfilesUsersForbidden with default headers values
func NewGetProfilesUsersForbidden() *GetProfilesUsersForbidden {
	return &GetProfilesUsersForbidden{}
}

/*GetProfilesUsersForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetProfilesUsersForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetProfilesUsersForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/profiles/users][%d] getProfilesUsersForbidden  %+v", 403, o.Payload)
}

func (o *GetProfilesUsersForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetProfilesUsersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProfilesUsersNotFound creates a GetProfilesUsersNotFound with default headers values
func NewGetProfilesUsersNotFound() *GetProfilesUsersNotFound {
	return &GetProfilesUsersNotFound{}
}

/*GetProfilesUsersNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetProfilesUsersNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetProfilesUsersNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/profiles/users][%d] getProfilesUsersNotFound  %+v", 404, o.Payload)
}

func (o *GetProfilesUsersNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetProfilesUsersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProfilesUsersRequestEntityTooLarge creates a GetProfilesUsersRequestEntityTooLarge with default headers values
func NewGetProfilesUsersRequestEntityTooLarge() *GetProfilesUsersRequestEntityTooLarge {
	return &GetProfilesUsersRequestEntityTooLarge{}
}

/*GetProfilesUsersRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetProfilesUsersRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetProfilesUsersRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/profiles/users][%d] getProfilesUsersRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetProfilesUsersRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetProfilesUsersRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProfilesUsersUnsupportedMediaType creates a GetProfilesUsersUnsupportedMediaType with default headers values
func NewGetProfilesUsersUnsupportedMediaType() *GetProfilesUsersUnsupportedMediaType {
	return &GetProfilesUsersUnsupportedMediaType{}
}

/*GetProfilesUsersUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetProfilesUsersUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetProfilesUsersUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/profiles/users][%d] getProfilesUsersUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetProfilesUsersUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetProfilesUsersUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProfilesUsersTooManyRequests creates a GetProfilesUsersTooManyRequests with default headers values
func NewGetProfilesUsersTooManyRequests() *GetProfilesUsersTooManyRequests {
	return &GetProfilesUsersTooManyRequests{}
}

/*GetProfilesUsersTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetProfilesUsersTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetProfilesUsersTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/profiles/users][%d] getProfilesUsersTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetProfilesUsersTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetProfilesUsersTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProfilesUsersInternalServerError creates a GetProfilesUsersInternalServerError with default headers values
func NewGetProfilesUsersInternalServerError() *GetProfilesUsersInternalServerError {
	return &GetProfilesUsersInternalServerError{}
}

/*GetProfilesUsersInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetProfilesUsersInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetProfilesUsersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/profiles/users][%d] getProfilesUsersInternalServerError  %+v", 500, o.Payload)
}

func (o *GetProfilesUsersInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetProfilesUsersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProfilesUsersServiceUnavailable creates a GetProfilesUsersServiceUnavailable with default headers values
func NewGetProfilesUsersServiceUnavailable() *GetProfilesUsersServiceUnavailable {
	return &GetProfilesUsersServiceUnavailable{}
}

/*GetProfilesUsersServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetProfilesUsersServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetProfilesUsersServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/profiles/users][%d] getProfilesUsersServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetProfilesUsersServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetProfilesUsersServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProfilesUsersGatewayTimeout creates a GetProfilesUsersGatewayTimeout with default headers values
func NewGetProfilesUsersGatewayTimeout() *GetProfilesUsersGatewayTimeout {
	return &GetProfilesUsersGatewayTimeout{}
}

/*GetProfilesUsersGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetProfilesUsersGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetProfilesUsersGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/profiles/users][%d] getProfilesUsersGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetProfilesUsersGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetProfilesUsersGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

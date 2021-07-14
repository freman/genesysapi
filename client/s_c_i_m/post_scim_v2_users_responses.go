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

// PostScimV2UsersReader is a Reader for the PostScimV2Users structure.
type PostScimV2UsersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostScimV2UsersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostScimV2UsersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewPostScimV2UsersCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostScimV2UsersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostScimV2UsersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostScimV2UsersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostScimV2UsersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostScimV2UsersRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPostScimV2UsersConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostScimV2UsersRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostScimV2UsersUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostScimV2UsersTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostScimV2UsersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostScimV2UsersServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostScimV2UsersGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostScimV2UsersOK creates a PostScimV2UsersOK with default headers values
func NewPostScimV2UsersOK() *PostScimV2UsersOK {
	return &PostScimV2UsersOK{}
}

/*PostScimV2UsersOK handles this case with default header values.

successful operation
*/
type PostScimV2UsersOK struct {
	Payload *models.ScimV2User
}

func (o *PostScimV2UsersOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/scim/v2/users][%d] postScimV2UsersOK  %+v", 200, o.Payload)
}

func (o *PostScimV2UsersOK) GetPayload() *models.ScimV2User {
	return o.Payload
}

func (o *PostScimV2UsersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ScimV2User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostScimV2UsersCreated creates a PostScimV2UsersCreated with default headers values
func NewPostScimV2UsersCreated() *PostScimV2UsersCreated {
	return &PostScimV2UsersCreated{}
}

/*PostScimV2UsersCreated handles this case with default header values.

User Created.
*/
type PostScimV2UsersCreated struct {
	Payload *models.ScimV2User
}

func (o *PostScimV2UsersCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/scim/v2/users][%d] postScimV2UsersCreated  %+v", 201, o.Payload)
}

func (o *PostScimV2UsersCreated) GetPayload() *models.ScimV2User {
	return o.Payload
}

func (o *PostScimV2UsersCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ScimV2User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostScimV2UsersBadRequest creates a PostScimV2UsersBadRequest with default headers values
func NewPostScimV2UsersBadRequest() *PostScimV2UsersBadRequest {
	return &PostScimV2UsersBadRequest{}
}

/*PostScimV2UsersBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostScimV2UsersBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostScimV2UsersBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/scim/v2/users][%d] postScimV2UsersBadRequest  %+v", 400, o.Payload)
}

func (o *PostScimV2UsersBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostScimV2UsersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostScimV2UsersUnauthorized creates a PostScimV2UsersUnauthorized with default headers values
func NewPostScimV2UsersUnauthorized() *PostScimV2UsersUnauthorized {
	return &PostScimV2UsersUnauthorized{}
}

/*PostScimV2UsersUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostScimV2UsersUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostScimV2UsersUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/scim/v2/users][%d] postScimV2UsersUnauthorized  %+v", 401, o.Payload)
}

func (o *PostScimV2UsersUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostScimV2UsersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostScimV2UsersForbidden creates a PostScimV2UsersForbidden with default headers values
func NewPostScimV2UsersForbidden() *PostScimV2UsersForbidden {
	return &PostScimV2UsersForbidden{}
}

/*PostScimV2UsersForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostScimV2UsersForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostScimV2UsersForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/scim/v2/users][%d] postScimV2UsersForbidden  %+v", 403, o.Payload)
}

func (o *PostScimV2UsersForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostScimV2UsersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostScimV2UsersNotFound creates a PostScimV2UsersNotFound with default headers values
func NewPostScimV2UsersNotFound() *PostScimV2UsersNotFound {
	return &PostScimV2UsersNotFound{}
}

/*PostScimV2UsersNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostScimV2UsersNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostScimV2UsersNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/scim/v2/users][%d] postScimV2UsersNotFound  %+v", 404, o.Payload)
}

func (o *PostScimV2UsersNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostScimV2UsersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostScimV2UsersRequestTimeout creates a PostScimV2UsersRequestTimeout with default headers values
func NewPostScimV2UsersRequestTimeout() *PostScimV2UsersRequestTimeout {
	return &PostScimV2UsersRequestTimeout{}
}

/*PostScimV2UsersRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostScimV2UsersRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostScimV2UsersRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/scim/v2/users][%d] postScimV2UsersRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostScimV2UsersRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostScimV2UsersRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostScimV2UsersConflict creates a PostScimV2UsersConflict with default headers values
func NewPostScimV2UsersConflict() *PostScimV2UsersConflict {
	return &PostScimV2UsersConflict{}
}

/*PostScimV2UsersConflict handles this case with default header values.

User name already in use by non-deleted user.
*/
type PostScimV2UsersConflict struct {
	Payload *models.ScimError
}

func (o *PostScimV2UsersConflict) Error() string {
	return fmt.Sprintf("[POST /api/v2/scim/v2/users][%d] postScimV2UsersConflict  %+v", 409, o.Payload)
}

func (o *PostScimV2UsersConflict) GetPayload() *models.ScimError {
	return o.Payload
}

func (o *PostScimV2UsersConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ScimError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostScimV2UsersRequestEntityTooLarge creates a PostScimV2UsersRequestEntityTooLarge with default headers values
func NewPostScimV2UsersRequestEntityTooLarge() *PostScimV2UsersRequestEntityTooLarge {
	return &PostScimV2UsersRequestEntityTooLarge{}
}

/*PostScimV2UsersRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostScimV2UsersRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostScimV2UsersRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/scim/v2/users][%d] postScimV2UsersRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostScimV2UsersRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostScimV2UsersRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostScimV2UsersUnsupportedMediaType creates a PostScimV2UsersUnsupportedMediaType with default headers values
func NewPostScimV2UsersUnsupportedMediaType() *PostScimV2UsersUnsupportedMediaType {
	return &PostScimV2UsersUnsupportedMediaType{}
}

/*PostScimV2UsersUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostScimV2UsersUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostScimV2UsersUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/scim/v2/users][%d] postScimV2UsersUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostScimV2UsersUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostScimV2UsersUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostScimV2UsersTooManyRequests creates a PostScimV2UsersTooManyRequests with default headers values
func NewPostScimV2UsersTooManyRequests() *PostScimV2UsersTooManyRequests {
	return &PostScimV2UsersTooManyRequests{}
}

/*PostScimV2UsersTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostScimV2UsersTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostScimV2UsersTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/scim/v2/users][%d] postScimV2UsersTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostScimV2UsersTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostScimV2UsersTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostScimV2UsersInternalServerError creates a PostScimV2UsersInternalServerError with default headers values
func NewPostScimV2UsersInternalServerError() *PostScimV2UsersInternalServerError {
	return &PostScimV2UsersInternalServerError{}
}

/*PostScimV2UsersInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostScimV2UsersInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostScimV2UsersInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/scim/v2/users][%d] postScimV2UsersInternalServerError  %+v", 500, o.Payload)
}

func (o *PostScimV2UsersInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostScimV2UsersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostScimV2UsersServiceUnavailable creates a PostScimV2UsersServiceUnavailable with default headers values
func NewPostScimV2UsersServiceUnavailable() *PostScimV2UsersServiceUnavailable {
	return &PostScimV2UsersServiceUnavailable{}
}

/*PostScimV2UsersServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostScimV2UsersServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostScimV2UsersServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/scim/v2/users][%d] postScimV2UsersServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostScimV2UsersServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostScimV2UsersServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostScimV2UsersGatewayTimeout creates a PostScimV2UsersGatewayTimeout with default headers values
func NewPostScimV2UsersGatewayTimeout() *PostScimV2UsersGatewayTimeout {
	return &PostScimV2UsersGatewayTimeout{}
}

/*PostScimV2UsersGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostScimV2UsersGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostScimV2UsersGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/scim/v2/users][%d] postScimV2UsersGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostScimV2UsersGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostScimV2UsersGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

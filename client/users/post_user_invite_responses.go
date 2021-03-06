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

// PostUserInviteReader is a Reader for the PostUserInvite structure.
type PostUserInviteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostUserInviteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPostUserInviteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostUserInviteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostUserInviteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostUserInviteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostUserInviteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostUserInviteRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostUserInviteUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostUserInviteTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostUserInviteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostUserInviteServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostUserInviteGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostUserInviteNoContent creates a PostUserInviteNoContent with default headers values
func NewPostUserInviteNoContent() *PostUserInviteNoContent {
	return &PostUserInviteNoContent{}
}

/*PostUserInviteNoContent handles this case with default header values.

Invitation Sent
*/
type PostUserInviteNoContent struct {
}

func (o *PostUserInviteNoContent) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/invite][%d] postUserInviteNoContent ", 204)
}

func (o *PostUserInviteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostUserInviteBadRequest creates a PostUserInviteBadRequest with default headers values
func NewPostUserInviteBadRequest() *PostUserInviteBadRequest {
	return &PostUserInviteBadRequest{}
}

/*PostUserInviteBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostUserInviteBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostUserInviteBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/invite][%d] postUserInviteBadRequest  %+v", 400, o.Payload)
}

func (o *PostUserInviteBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUserInviteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUserInviteUnauthorized creates a PostUserInviteUnauthorized with default headers values
func NewPostUserInviteUnauthorized() *PostUserInviteUnauthorized {
	return &PostUserInviteUnauthorized{}
}

/*PostUserInviteUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostUserInviteUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostUserInviteUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/invite][%d] postUserInviteUnauthorized  %+v", 401, o.Payload)
}

func (o *PostUserInviteUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUserInviteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUserInviteForbidden creates a PostUserInviteForbidden with default headers values
func NewPostUserInviteForbidden() *PostUserInviteForbidden {
	return &PostUserInviteForbidden{}
}

/*PostUserInviteForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostUserInviteForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostUserInviteForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/invite][%d] postUserInviteForbidden  %+v", 403, o.Payload)
}

func (o *PostUserInviteForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUserInviteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUserInviteNotFound creates a PostUserInviteNotFound with default headers values
func NewPostUserInviteNotFound() *PostUserInviteNotFound {
	return &PostUserInviteNotFound{}
}

/*PostUserInviteNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostUserInviteNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostUserInviteNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/invite][%d] postUserInviteNotFound  %+v", 404, o.Payload)
}

func (o *PostUserInviteNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUserInviteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUserInviteRequestEntityTooLarge creates a PostUserInviteRequestEntityTooLarge with default headers values
func NewPostUserInviteRequestEntityTooLarge() *PostUserInviteRequestEntityTooLarge {
	return &PostUserInviteRequestEntityTooLarge{}
}

/*PostUserInviteRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostUserInviteRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostUserInviteRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/invite][%d] postUserInviteRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostUserInviteRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUserInviteRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUserInviteUnsupportedMediaType creates a PostUserInviteUnsupportedMediaType with default headers values
func NewPostUserInviteUnsupportedMediaType() *PostUserInviteUnsupportedMediaType {
	return &PostUserInviteUnsupportedMediaType{}
}

/*PostUserInviteUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostUserInviteUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostUserInviteUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/invite][%d] postUserInviteUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostUserInviteUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUserInviteUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUserInviteTooManyRequests creates a PostUserInviteTooManyRequests with default headers values
func NewPostUserInviteTooManyRequests() *PostUserInviteTooManyRequests {
	return &PostUserInviteTooManyRequests{}
}

/*PostUserInviteTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PostUserInviteTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostUserInviteTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/invite][%d] postUserInviteTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostUserInviteTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUserInviteTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUserInviteInternalServerError creates a PostUserInviteInternalServerError with default headers values
func NewPostUserInviteInternalServerError() *PostUserInviteInternalServerError {
	return &PostUserInviteInternalServerError{}
}

/*PostUserInviteInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostUserInviteInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostUserInviteInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/invite][%d] postUserInviteInternalServerError  %+v", 500, o.Payload)
}

func (o *PostUserInviteInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUserInviteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUserInviteServiceUnavailable creates a PostUserInviteServiceUnavailable with default headers values
func NewPostUserInviteServiceUnavailable() *PostUserInviteServiceUnavailable {
	return &PostUserInviteServiceUnavailable{}
}

/*PostUserInviteServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostUserInviteServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostUserInviteServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/invite][%d] postUserInviteServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostUserInviteServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUserInviteServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUserInviteGatewayTimeout creates a PostUserInviteGatewayTimeout with default headers values
func NewPostUserInviteGatewayTimeout() *PostUserInviteGatewayTimeout {
	return &PostUserInviteGatewayTimeout{}
}

/*PostUserInviteGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostUserInviteGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostUserInviteGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/invite][%d] postUserInviteGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostUserInviteGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUserInviteGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

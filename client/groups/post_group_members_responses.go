// Code generated by go-swagger; DO NOT EDIT.

package groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostGroupMembersReader is a Reader for the PostGroupMembers structure.
type PostGroupMembersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostGroupMembersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewPostGroupMembersAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostGroupMembersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostGroupMembersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostGroupMembersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostGroupMembersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPostGroupMembersConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostGroupMembersRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostGroupMembersUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostGroupMembersTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostGroupMembersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostGroupMembersServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostGroupMembersGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostGroupMembersAccepted creates a PostGroupMembersAccepted with default headers values
func NewPostGroupMembersAccepted() *PostGroupMembersAccepted {
	return &PostGroupMembersAccepted{}
}

/*PostGroupMembersAccepted handles this case with default header values.

Success, group membership was updated
*/
type PostGroupMembersAccepted struct {
	Payload models.Empty
}

func (o *PostGroupMembersAccepted) Error() string {
	return fmt.Sprintf("[POST /api/v2/groups/{groupId}/members][%d] postGroupMembersAccepted  %+v", 202, o.Payload)
}

func (o *PostGroupMembersAccepted) GetPayload() models.Empty {
	return o.Payload
}

func (o *PostGroupMembersAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGroupMembersBadRequest creates a PostGroupMembersBadRequest with default headers values
func NewPostGroupMembersBadRequest() *PostGroupMembersBadRequest {
	return &PostGroupMembersBadRequest{}
}

/*PostGroupMembersBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostGroupMembersBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostGroupMembersBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/groups/{groupId}/members][%d] postGroupMembersBadRequest  %+v", 400, o.Payload)
}

func (o *PostGroupMembersBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGroupMembersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGroupMembersUnauthorized creates a PostGroupMembersUnauthorized with default headers values
func NewPostGroupMembersUnauthorized() *PostGroupMembersUnauthorized {
	return &PostGroupMembersUnauthorized{}
}

/*PostGroupMembersUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostGroupMembersUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostGroupMembersUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/groups/{groupId}/members][%d] postGroupMembersUnauthorized  %+v", 401, o.Payload)
}

func (o *PostGroupMembersUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGroupMembersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGroupMembersForbidden creates a PostGroupMembersForbidden with default headers values
func NewPostGroupMembersForbidden() *PostGroupMembersForbidden {
	return &PostGroupMembersForbidden{}
}

/*PostGroupMembersForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostGroupMembersForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostGroupMembersForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/groups/{groupId}/members][%d] postGroupMembersForbidden  %+v", 403, o.Payload)
}

func (o *PostGroupMembersForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGroupMembersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGroupMembersNotFound creates a PostGroupMembersNotFound with default headers values
func NewPostGroupMembersNotFound() *PostGroupMembersNotFound {
	return &PostGroupMembersNotFound{}
}

/*PostGroupMembersNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostGroupMembersNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostGroupMembersNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/groups/{groupId}/members][%d] postGroupMembersNotFound  %+v", 404, o.Payload)
}

func (o *PostGroupMembersNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGroupMembersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGroupMembersConflict creates a PostGroupMembersConflict with default headers values
func NewPostGroupMembersConflict() *PostGroupMembersConflict {
	return &PostGroupMembersConflict{}
}

/*PostGroupMembersConflict handles this case with default header values.

Resource conflict - Unexpected version was provided
*/
type PostGroupMembersConflict struct {
}

func (o *PostGroupMembersConflict) Error() string {
	return fmt.Sprintf("[POST /api/v2/groups/{groupId}/members][%d] postGroupMembersConflict ", 409)
}

func (o *PostGroupMembersConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostGroupMembersRequestEntityTooLarge creates a PostGroupMembersRequestEntityTooLarge with default headers values
func NewPostGroupMembersRequestEntityTooLarge() *PostGroupMembersRequestEntityTooLarge {
	return &PostGroupMembersRequestEntityTooLarge{}
}

/*PostGroupMembersRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostGroupMembersRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostGroupMembersRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/groups/{groupId}/members][%d] postGroupMembersRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostGroupMembersRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGroupMembersRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGroupMembersUnsupportedMediaType creates a PostGroupMembersUnsupportedMediaType with default headers values
func NewPostGroupMembersUnsupportedMediaType() *PostGroupMembersUnsupportedMediaType {
	return &PostGroupMembersUnsupportedMediaType{}
}

/*PostGroupMembersUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostGroupMembersUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostGroupMembersUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/groups/{groupId}/members][%d] postGroupMembersUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostGroupMembersUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGroupMembersUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGroupMembersTooManyRequests creates a PostGroupMembersTooManyRequests with default headers values
func NewPostGroupMembersTooManyRequests() *PostGroupMembersTooManyRequests {
	return &PostGroupMembersTooManyRequests{}
}

/*PostGroupMembersTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PostGroupMembersTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostGroupMembersTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/groups/{groupId}/members][%d] postGroupMembersTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostGroupMembersTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGroupMembersTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGroupMembersInternalServerError creates a PostGroupMembersInternalServerError with default headers values
func NewPostGroupMembersInternalServerError() *PostGroupMembersInternalServerError {
	return &PostGroupMembersInternalServerError{}
}

/*PostGroupMembersInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostGroupMembersInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostGroupMembersInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/groups/{groupId}/members][%d] postGroupMembersInternalServerError  %+v", 500, o.Payload)
}

func (o *PostGroupMembersInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGroupMembersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGroupMembersServiceUnavailable creates a PostGroupMembersServiceUnavailable with default headers values
func NewPostGroupMembersServiceUnavailable() *PostGroupMembersServiceUnavailable {
	return &PostGroupMembersServiceUnavailable{}
}

/*PostGroupMembersServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostGroupMembersServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostGroupMembersServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/groups/{groupId}/members][%d] postGroupMembersServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostGroupMembersServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGroupMembersServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGroupMembersGatewayTimeout creates a PostGroupMembersGatewayTimeout with default headers values
func NewPostGroupMembersGatewayTimeout() *PostGroupMembersGatewayTimeout {
	return &PostGroupMembersGatewayTimeout{}
}

/*PostGroupMembersGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostGroupMembersGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostGroupMembersGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/groups/{groupId}/members][%d] postGroupMembersGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostGroupMembersGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGroupMembersGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
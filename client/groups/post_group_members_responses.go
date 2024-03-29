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
	case 408:
		result := NewPostGroupMembersRequestTimeout()
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

/*
PostGroupMembersAccepted describes a response with status code 202, with default header values.

Success, group membership was updated
*/
type PostGroupMembersAccepted struct {
	Payload models.Empty
}

// IsSuccess returns true when this post group members accepted response has a 2xx status code
func (o *PostGroupMembersAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post group members accepted response has a 3xx status code
func (o *PostGroupMembersAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post group members accepted response has a 4xx status code
func (o *PostGroupMembersAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this post group members accepted response has a 5xx status code
func (o *PostGroupMembersAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this post group members accepted response a status code equal to that given
func (o *PostGroupMembersAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *PostGroupMembersAccepted) Error() string {
	return fmt.Sprintf("[POST /api/v2/groups/{groupId}/members][%d] postGroupMembersAccepted  %+v", 202, o.Payload)
}

func (o *PostGroupMembersAccepted) String() string {
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

/*
PostGroupMembersBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostGroupMembersBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post group members bad request response has a 2xx status code
func (o *PostGroupMembersBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post group members bad request response has a 3xx status code
func (o *PostGroupMembersBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post group members bad request response has a 4xx status code
func (o *PostGroupMembersBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post group members bad request response has a 5xx status code
func (o *PostGroupMembersBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post group members bad request response a status code equal to that given
func (o *PostGroupMembersBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostGroupMembersBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/groups/{groupId}/members][%d] postGroupMembersBadRequest  %+v", 400, o.Payload)
}

func (o *PostGroupMembersBadRequest) String() string {
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

/*
PostGroupMembersUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostGroupMembersUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post group members unauthorized response has a 2xx status code
func (o *PostGroupMembersUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post group members unauthorized response has a 3xx status code
func (o *PostGroupMembersUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post group members unauthorized response has a 4xx status code
func (o *PostGroupMembersUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post group members unauthorized response has a 5xx status code
func (o *PostGroupMembersUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post group members unauthorized response a status code equal to that given
func (o *PostGroupMembersUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostGroupMembersUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/groups/{groupId}/members][%d] postGroupMembersUnauthorized  %+v", 401, o.Payload)
}

func (o *PostGroupMembersUnauthorized) String() string {
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

/*
PostGroupMembersForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostGroupMembersForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post group members forbidden response has a 2xx status code
func (o *PostGroupMembersForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post group members forbidden response has a 3xx status code
func (o *PostGroupMembersForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post group members forbidden response has a 4xx status code
func (o *PostGroupMembersForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post group members forbidden response has a 5xx status code
func (o *PostGroupMembersForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post group members forbidden response a status code equal to that given
func (o *PostGroupMembersForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostGroupMembersForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/groups/{groupId}/members][%d] postGroupMembersForbidden  %+v", 403, o.Payload)
}

func (o *PostGroupMembersForbidden) String() string {
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

/*
PostGroupMembersNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostGroupMembersNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post group members not found response has a 2xx status code
func (o *PostGroupMembersNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post group members not found response has a 3xx status code
func (o *PostGroupMembersNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post group members not found response has a 4xx status code
func (o *PostGroupMembersNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post group members not found response has a 5xx status code
func (o *PostGroupMembersNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post group members not found response a status code equal to that given
func (o *PostGroupMembersNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostGroupMembersNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/groups/{groupId}/members][%d] postGroupMembersNotFound  %+v", 404, o.Payload)
}

func (o *PostGroupMembersNotFound) String() string {
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

// NewPostGroupMembersRequestTimeout creates a PostGroupMembersRequestTimeout with default headers values
func NewPostGroupMembersRequestTimeout() *PostGroupMembersRequestTimeout {
	return &PostGroupMembersRequestTimeout{}
}

/*
PostGroupMembersRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostGroupMembersRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post group members request timeout response has a 2xx status code
func (o *PostGroupMembersRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post group members request timeout response has a 3xx status code
func (o *PostGroupMembersRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post group members request timeout response has a 4xx status code
func (o *PostGroupMembersRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post group members request timeout response has a 5xx status code
func (o *PostGroupMembersRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post group members request timeout response a status code equal to that given
func (o *PostGroupMembersRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostGroupMembersRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/groups/{groupId}/members][%d] postGroupMembersRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostGroupMembersRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/groups/{groupId}/members][%d] postGroupMembersRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostGroupMembersRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGroupMembersRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*
PostGroupMembersConflict describes a response with status code 409, with default header values.

Resource conflict - Unexpected version was provided
*/
type PostGroupMembersConflict struct {
}

// IsSuccess returns true when this post group members conflict response has a 2xx status code
func (o *PostGroupMembersConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post group members conflict response has a 3xx status code
func (o *PostGroupMembersConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post group members conflict response has a 4xx status code
func (o *PostGroupMembersConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this post group members conflict response has a 5xx status code
func (o *PostGroupMembersConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this post group members conflict response a status code equal to that given
func (o *PostGroupMembersConflict) IsCode(code int) bool {
	return code == 409
}

func (o *PostGroupMembersConflict) Error() string {
	return fmt.Sprintf("[POST /api/v2/groups/{groupId}/members][%d] postGroupMembersConflict ", 409)
}

func (o *PostGroupMembersConflict) String() string {
	return fmt.Sprintf("[POST /api/v2/groups/{groupId}/members][%d] postGroupMembersConflict ", 409)
}

func (o *PostGroupMembersConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostGroupMembersRequestEntityTooLarge creates a PostGroupMembersRequestEntityTooLarge with default headers values
func NewPostGroupMembersRequestEntityTooLarge() *PostGroupMembersRequestEntityTooLarge {
	return &PostGroupMembersRequestEntityTooLarge{}
}

/*
PostGroupMembersRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PostGroupMembersRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post group members request entity too large response has a 2xx status code
func (o *PostGroupMembersRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post group members request entity too large response has a 3xx status code
func (o *PostGroupMembersRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post group members request entity too large response has a 4xx status code
func (o *PostGroupMembersRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post group members request entity too large response has a 5xx status code
func (o *PostGroupMembersRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post group members request entity too large response a status code equal to that given
func (o *PostGroupMembersRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostGroupMembersRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/groups/{groupId}/members][%d] postGroupMembersRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostGroupMembersRequestEntityTooLarge) String() string {
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

/*
PostGroupMembersUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostGroupMembersUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post group members unsupported media type response has a 2xx status code
func (o *PostGroupMembersUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post group members unsupported media type response has a 3xx status code
func (o *PostGroupMembersUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post group members unsupported media type response has a 4xx status code
func (o *PostGroupMembersUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post group members unsupported media type response has a 5xx status code
func (o *PostGroupMembersUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post group members unsupported media type response a status code equal to that given
func (o *PostGroupMembersUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostGroupMembersUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/groups/{groupId}/members][%d] postGroupMembersUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostGroupMembersUnsupportedMediaType) String() string {
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

/*
PostGroupMembersTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostGroupMembersTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post group members too many requests response has a 2xx status code
func (o *PostGroupMembersTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post group members too many requests response has a 3xx status code
func (o *PostGroupMembersTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post group members too many requests response has a 4xx status code
func (o *PostGroupMembersTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post group members too many requests response has a 5xx status code
func (o *PostGroupMembersTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post group members too many requests response a status code equal to that given
func (o *PostGroupMembersTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostGroupMembersTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/groups/{groupId}/members][%d] postGroupMembersTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostGroupMembersTooManyRequests) String() string {
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

/*
PostGroupMembersInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostGroupMembersInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post group members internal server error response has a 2xx status code
func (o *PostGroupMembersInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post group members internal server error response has a 3xx status code
func (o *PostGroupMembersInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post group members internal server error response has a 4xx status code
func (o *PostGroupMembersInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post group members internal server error response has a 5xx status code
func (o *PostGroupMembersInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post group members internal server error response a status code equal to that given
func (o *PostGroupMembersInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostGroupMembersInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/groups/{groupId}/members][%d] postGroupMembersInternalServerError  %+v", 500, o.Payload)
}

func (o *PostGroupMembersInternalServerError) String() string {
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

/*
PostGroupMembersServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostGroupMembersServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post group members service unavailable response has a 2xx status code
func (o *PostGroupMembersServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post group members service unavailable response has a 3xx status code
func (o *PostGroupMembersServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post group members service unavailable response has a 4xx status code
func (o *PostGroupMembersServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post group members service unavailable response has a 5xx status code
func (o *PostGroupMembersServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post group members service unavailable response a status code equal to that given
func (o *PostGroupMembersServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostGroupMembersServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/groups/{groupId}/members][%d] postGroupMembersServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostGroupMembersServiceUnavailable) String() string {
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

/*
PostGroupMembersGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostGroupMembersGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post group members gateway timeout response has a 2xx status code
func (o *PostGroupMembersGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post group members gateway timeout response has a 3xx status code
func (o *PostGroupMembersGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post group members gateway timeout response has a 4xx status code
func (o *PostGroupMembersGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post group members gateway timeout response has a 5xx status code
func (o *PostGroupMembersGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post group members gateway timeout response a status code equal to that given
func (o *PostGroupMembersGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostGroupMembersGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/groups/{groupId}/members][%d] postGroupMembersGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostGroupMembersGatewayTimeout) String() string {
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

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

// DeleteGroupMembersReader is a Reader for the DeleteGroupMembers structure.
type DeleteGroupMembersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteGroupMembersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewDeleteGroupMembersAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteGroupMembersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteGroupMembersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteGroupMembersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteGroupMembersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewDeleteGroupMembersRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteGroupMembersRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteGroupMembersUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteGroupMembersTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteGroupMembersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteGroupMembersServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteGroupMembersGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteGroupMembersAccepted creates a DeleteGroupMembersAccepted with default headers values
func NewDeleteGroupMembersAccepted() *DeleteGroupMembersAccepted {
	return &DeleteGroupMembersAccepted{}
}

/*
DeleteGroupMembersAccepted describes a response with status code 202, with default header values.

Success, group membership was updated
*/
type DeleteGroupMembersAccepted struct {
	Payload models.Empty
}

// IsSuccess returns true when this delete group members accepted response has a 2xx status code
func (o *DeleteGroupMembersAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete group members accepted response has a 3xx status code
func (o *DeleteGroupMembersAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete group members accepted response has a 4xx status code
func (o *DeleteGroupMembersAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete group members accepted response has a 5xx status code
func (o *DeleteGroupMembersAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this delete group members accepted response a status code equal to that given
func (o *DeleteGroupMembersAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *DeleteGroupMembersAccepted) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/groups/{groupId}/members][%d] deleteGroupMembersAccepted  %+v", 202, o.Payload)
}

func (o *DeleteGroupMembersAccepted) String() string {
	return fmt.Sprintf("[DELETE /api/v2/groups/{groupId}/members][%d] deleteGroupMembersAccepted  %+v", 202, o.Payload)
}

func (o *DeleteGroupMembersAccepted) GetPayload() models.Empty {
	return o.Payload
}

func (o *DeleteGroupMembersAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteGroupMembersBadRequest creates a DeleteGroupMembersBadRequest with default headers values
func NewDeleteGroupMembersBadRequest() *DeleteGroupMembersBadRequest {
	return &DeleteGroupMembersBadRequest{}
}

/*
DeleteGroupMembersBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteGroupMembersBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete group members bad request response has a 2xx status code
func (o *DeleteGroupMembersBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete group members bad request response has a 3xx status code
func (o *DeleteGroupMembersBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete group members bad request response has a 4xx status code
func (o *DeleteGroupMembersBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete group members bad request response has a 5xx status code
func (o *DeleteGroupMembersBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete group members bad request response a status code equal to that given
func (o *DeleteGroupMembersBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DeleteGroupMembersBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/groups/{groupId}/members][%d] deleteGroupMembersBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteGroupMembersBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api/v2/groups/{groupId}/members][%d] deleteGroupMembersBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteGroupMembersBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteGroupMembersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteGroupMembersUnauthorized creates a DeleteGroupMembersUnauthorized with default headers values
func NewDeleteGroupMembersUnauthorized() *DeleteGroupMembersUnauthorized {
	return &DeleteGroupMembersUnauthorized{}
}

/*
DeleteGroupMembersUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteGroupMembersUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete group members unauthorized response has a 2xx status code
func (o *DeleteGroupMembersUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete group members unauthorized response has a 3xx status code
func (o *DeleteGroupMembersUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete group members unauthorized response has a 4xx status code
func (o *DeleteGroupMembersUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete group members unauthorized response has a 5xx status code
func (o *DeleteGroupMembersUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete group members unauthorized response a status code equal to that given
func (o *DeleteGroupMembersUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteGroupMembersUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/groups/{groupId}/members][%d] deleteGroupMembersUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteGroupMembersUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/v2/groups/{groupId}/members][%d] deleteGroupMembersUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteGroupMembersUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteGroupMembersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteGroupMembersForbidden creates a DeleteGroupMembersForbidden with default headers values
func NewDeleteGroupMembersForbidden() *DeleteGroupMembersForbidden {
	return &DeleteGroupMembersForbidden{}
}

/*
DeleteGroupMembersForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type DeleteGroupMembersForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete group members forbidden response has a 2xx status code
func (o *DeleteGroupMembersForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete group members forbidden response has a 3xx status code
func (o *DeleteGroupMembersForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete group members forbidden response has a 4xx status code
func (o *DeleteGroupMembersForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete group members forbidden response has a 5xx status code
func (o *DeleteGroupMembersForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete group members forbidden response a status code equal to that given
func (o *DeleteGroupMembersForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteGroupMembersForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/groups/{groupId}/members][%d] deleteGroupMembersForbidden  %+v", 403, o.Payload)
}

func (o *DeleteGroupMembersForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v2/groups/{groupId}/members][%d] deleteGroupMembersForbidden  %+v", 403, o.Payload)
}

func (o *DeleteGroupMembersForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteGroupMembersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteGroupMembersNotFound creates a DeleteGroupMembersNotFound with default headers values
func NewDeleteGroupMembersNotFound() *DeleteGroupMembersNotFound {
	return &DeleteGroupMembersNotFound{}
}

/*
DeleteGroupMembersNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type DeleteGroupMembersNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete group members not found response has a 2xx status code
func (o *DeleteGroupMembersNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete group members not found response has a 3xx status code
func (o *DeleteGroupMembersNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete group members not found response has a 4xx status code
func (o *DeleteGroupMembersNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete group members not found response has a 5xx status code
func (o *DeleteGroupMembersNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete group members not found response a status code equal to that given
func (o *DeleteGroupMembersNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteGroupMembersNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/groups/{groupId}/members][%d] deleteGroupMembersNotFound  %+v", 404, o.Payload)
}

func (o *DeleteGroupMembersNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/v2/groups/{groupId}/members][%d] deleteGroupMembersNotFound  %+v", 404, o.Payload)
}

func (o *DeleteGroupMembersNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteGroupMembersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteGroupMembersRequestTimeout creates a DeleteGroupMembersRequestTimeout with default headers values
func NewDeleteGroupMembersRequestTimeout() *DeleteGroupMembersRequestTimeout {
	return &DeleteGroupMembersRequestTimeout{}
}

/*
DeleteGroupMembersRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type DeleteGroupMembersRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete group members request timeout response has a 2xx status code
func (o *DeleteGroupMembersRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete group members request timeout response has a 3xx status code
func (o *DeleteGroupMembersRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete group members request timeout response has a 4xx status code
func (o *DeleteGroupMembersRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete group members request timeout response has a 5xx status code
func (o *DeleteGroupMembersRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this delete group members request timeout response a status code equal to that given
func (o *DeleteGroupMembersRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *DeleteGroupMembersRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/groups/{groupId}/members][%d] deleteGroupMembersRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteGroupMembersRequestTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/groups/{groupId}/members][%d] deleteGroupMembersRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteGroupMembersRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteGroupMembersRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteGroupMembersRequestEntityTooLarge creates a DeleteGroupMembersRequestEntityTooLarge with default headers values
func NewDeleteGroupMembersRequestEntityTooLarge() *DeleteGroupMembersRequestEntityTooLarge {
	return &DeleteGroupMembersRequestEntityTooLarge{}
}

/*
DeleteGroupMembersRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type DeleteGroupMembersRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete group members request entity too large response has a 2xx status code
func (o *DeleteGroupMembersRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete group members request entity too large response has a 3xx status code
func (o *DeleteGroupMembersRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete group members request entity too large response has a 4xx status code
func (o *DeleteGroupMembersRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete group members request entity too large response has a 5xx status code
func (o *DeleteGroupMembersRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this delete group members request entity too large response a status code equal to that given
func (o *DeleteGroupMembersRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *DeleteGroupMembersRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/groups/{groupId}/members][%d] deleteGroupMembersRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteGroupMembersRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[DELETE /api/v2/groups/{groupId}/members][%d] deleteGroupMembersRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteGroupMembersRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteGroupMembersRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteGroupMembersUnsupportedMediaType creates a DeleteGroupMembersUnsupportedMediaType with default headers values
func NewDeleteGroupMembersUnsupportedMediaType() *DeleteGroupMembersUnsupportedMediaType {
	return &DeleteGroupMembersUnsupportedMediaType{}
}

/*
DeleteGroupMembersUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteGroupMembersUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete group members unsupported media type response has a 2xx status code
func (o *DeleteGroupMembersUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete group members unsupported media type response has a 3xx status code
func (o *DeleteGroupMembersUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete group members unsupported media type response has a 4xx status code
func (o *DeleteGroupMembersUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete group members unsupported media type response has a 5xx status code
func (o *DeleteGroupMembersUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this delete group members unsupported media type response a status code equal to that given
func (o *DeleteGroupMembersUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *DeleteGroupMembersUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/groups/{groupId}/members][%d] deleteGroupMembersUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteGroupMembersUnsupportedMediaType) String() string {
	return fmt.Sprintf("[DELETE /api/v2/groups/{groupId}/members][%d] deleteGroupMembersUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteGroupMembersUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteGroupMembersUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteGroupMembersTooManyRequests creates a DeleteGroupMembersTooManyRequests with default headers values
func NewDeleteGroupMembersTooManyRequests() *DeleteGroupMembersTooManyRequests {
	return &DeleteGroupMembersTooManyRequests{}
}

/*
DeleteGroupMembersTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type DeleteGroupMembersTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete group members too many requests response has a 2xx status code
func (o *DeleteGroupMembersTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete group members too many requests response has a 3xx status code
func (o *DeleteGroupMembersTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete group members too many requests response has a 4xx status code
func (o *DeleteGroupMembersTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete group members too many requests response has a 5xx status code
func (o *DeleteGroupMembersTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this delete group members too many requests response a status code equal to that given
func (o *DeleteGroupMembersTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *DeleteGroupMembersTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/groups/{groupId}/members][%d] deleteGroupMembersTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteGroupMembersTooManyRequests) String() string {
	return fmt.Sprintf("[DELETE /api/v2/groups/{groupId}/members][%d] deleteGroupMembersTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteGroupMembersTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteGroupMembersTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteGroupMembersInternalServerError creates a DeleteGroupMembersInternalServerError with default headers values
func NewDeleteGroupMembersInternalServerError() *DeleteGroupMembersInternalServerError {
	return &DeleteGroupMembersInternalServerError{}
}

/*
DeleteGroupMembersInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteGroupMembersInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete group members internal server error response has a 2xx status code
func (o *DeleteGroupMembersInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete group members internal server error response has a 3xx status code
func (o *DeleteGroupMembersInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete group members internal server error response has a 4xx status code
func (o *DeleteGroupMembersInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete group members internal server error response has a 5xx status code
func (o *DeleteGroupMembersInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete group members internal server error response a status code equal to that given
func (o *DeleteGroupMembersInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DeleteGroupMembersInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/groups/{groupId}/members][%d] deleteGroupMembersInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteGroupMembersInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /api/v2/groups/{groupId}/members][%d] deleteGroupMembersInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteGroupMembersInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteGroupMembersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteGroupMembersServiceUnavailable creates a DeleteGroupMembersServiceUnavailable with default headers values
func NewDeleteGroupMembersServiceUnavailable() *DeleteGroupMembersServiceUnavailable {
	return &DeleteGroupMembersServiceUnavailable{}
}

/*
DeleteGroupMembersServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteGroupMembersServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete group members service unavailable response has a 2xx status code
func (o *DeleteGroupMembersServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete group members service unavailable response has a 3xx status code
func (o *DeleteGroupMembersServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete group members service unavailable response has a 4xx status code
func (o *DeleteGroupMembersServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete group members service unavailable response has a 5xx status code
func (o *DeleteGroupMembersServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this delete group members service unavailable response a status code equal to that given
func (o *DeleteGroupMembersServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *DeleteGroupMembersServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/groups/{groupId}/members][%d] deleteGroupMembersServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteGroupMembersServiceUnavailable) String() string {
	return fmt.Sprintf("[DELETE /api/v2/groups/{groupId}/members][%d] deleteGroupMembersServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteGroupMembersServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteGroupMembersServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteGroupMembersGatewayTimeout creates a DeleteGroupMembersGatewayTimeout with default headers values
func NewDeleteGroupMembersGatewayTimeout() *DeleteGroupMembersGatewayTimeout {
	return &DeleteGroupMembersGatewayTimeout{}
}

/*
DeleteGroupMembersGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type DeleteGroupMembersGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete group members gateway timeout response has a 2xx status code
func (o *DeleteGroupMembersGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete group members gateway timeout response has a 3xx status code
func (o *DeleteGroupMembersGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete group members gateway timeout response has a 4xx status code
func (o *DeleteGroupMembersGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete group members gateway timeout response has a 5xx status code
func (o *DeleteGroupMembersGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this delete group members gateway timeout response a status code equal to that given
func (o *DeleteGroupMembersGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *DeleteGroupMembersGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/groups/{groupId}/members][%d] deleteGroupMembersGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteGroupMembersGatewayTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/groups/{groupId}/members][%d] deleteGroupMembersGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteGroupMembersGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteGroupMembersGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

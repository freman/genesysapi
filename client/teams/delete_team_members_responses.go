// Code generated by go-swagger; DO NOT EDIT.

package teams

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// DeleteTeamMembersReader is a Reader for the DeleteTeamMembers structure.
type DeleteTeamMembersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteTeamMembersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteTeamMembersNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteTeamMembersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteTeamMembersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteTeamMembersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteTeamMembersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewDeleteTeamMembersRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteTeamMembersRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteTeamMembersUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteTeamMembersTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteTeamMembersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteTeamMembersServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteTeamMembersGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteTeamMembersNoContent creates a DeleteTeamMembersNoContent with default headers values
func NewDeleteTeamMembersNoContent() *DeleteTeamMembersNoContent {
	return &DeleteTeamMembersNoContent{}
}

/*
DeleteTeamMembersNoContent describes a response with status code 204, with default header values.

Success, team members are removed
*/
type DeleteTeamMembersNoContent struct {
}

// IsSuccess returns true when this delete team members no content response has a 2xx status code
func (o *DeleteTeamMembersNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete team members no content response has a 3xx status code
func (o *DeleteTeamMembersNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete team members no content response has a 4xx status code
func (o *DeleteTeamMembersNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete team members no content response has a 5xx status code
func (o *DeleteTeamMembersNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete team members no content response a status code equal to that given
func (o *DeleteTeamMembersNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DeleteTeamMembersNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/teams/{teamId}/members][%d] deleteTeamMembersNoContent ", 204)
}

func (o *DeleteTeamMembersNoContent) String() string {
	return fmt.Sprintf("[DELETE /api/v2/teams/{teamId}/members][%d] deleteTeamMembersNoContent ", 204)
}

func (o *DeleteTeamMembersNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteTeamMembersBadRequest creates a DeleteTeamMembersBadRequest with default headers values
func NewDeleteTeamMembersBadRequest() *DeleteTeamMembersBadRequest {
	return &DeleteTeamMembersBadRequest{}
}

/*
DeleteTeamMembersBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteTeamMembersBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete team members bad request response has a 2xx status code
func (o *DeleteTeamMembersBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete team members bad request response has a 3xx status code
func (o *DeleteTeamMembersBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete team members bad request response has a 4xx status code
func (o *DeleteTeamMembersBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete team members bad request response has a 5xx status code
func (o *DeleteTeamMembersBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete team members bad request response a status code equal to that given
func (o *DeleteTeamMembersBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DeleteTeamMembersBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/teams/{teamId}/members][%d] deleteTeamMembersBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteTeamMembersBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api/v2/teams/{teamId}/members][%d] deleteTeamMembersBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteTeamMembersBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTeamMembersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTeamMembersUnauthorized creates a DeleteTeamMembersUnauthorized with default headers values
func NewDeleteTeamMembersUnauthorized() *DeleteTeamMembersUnauthorized {
	return &DeleteTeamMembersUnauthorized{}
}

/*
DeleteTeamMembersUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteTeamMembersUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete team members unauthorized response has a 2xx status code
func (o *DeleteTeamMembersUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete team members unauthorized response has a 3xx status code
func (o *DeleteTeamMembersUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete team members unauthorized response has a 4xx status code
func (o *DeleteTeamMembersUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete team members unauthorized response has a 5xx status code
func (o *DeleteTeamMembersUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete team members unauthorized response a status code equal to that given
func (o *DeleteTeamMembersUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteTeamMembersUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/teams/{teamId}/members][%d] deleteTeamMembersUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteTeamMembersUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/v2/teams/{teamId}/members][%d] deleteTeamMembersUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteTeamMembersUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTeamMembersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTeamMembersForbidden creates a DeleteTeamMembersForbidden with default headers values
func NewDeleteTeamMembersForbidden() *DeleteTeamMembersForbidden {
	return &DeleteTeamMembersForbidden{}
}

/*
DeleteTeamMembersForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type DeleteTeamMembersForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete team members forbidden response has a 2xx status code
func (o *DeleteTeamMembersForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete team members forbidden response has a 3xx status code
func (o *DeleteTeamMembersForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete team members forbidden response has a 4xx status code
func (o *DeleteTeamMembersForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete team members forbidden response has a 5xx status code
func (o *DeleteTeamMembersForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete team members forbidden response a status code equal to that given
func (o *DeleteTeamMembersForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteTeamMembersForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/teams/{teamId}/members][%d] deleteTeamMembersForbidden  %+v", 403, o.Payload)
}

func (o *DeleteTeamMembersForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v2/teams/{teamId}/members][%d] deleteTeamMembersForbidden  %+v", 403, o.Payload)
}

func (o *DeleteTeamMembersForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTeamMembersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTeamMembersNotFound creates a DeleteTeamMembersNotFound with default headers values
func NewDeleteTeamMembersNotFound() *DeleteTeamMembersNotFound {
	return &DeleteTeamMembersNotFound{}
}

/*
DeleteTeamMembersNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type DeleteTeamMembersNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete team members not found response has a 2xx status code
func (o *DeleteTeamMembersNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete team members not found response has a 3xx status code
func (o *DeleteTeamMembersNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete team members not found response has a 4xx status code
func (o *DeleteTeamMembersNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete team members not found response has a 5xx status code
func (o *DeleteTeamMembersNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete team members not found response a status code equal to that given
func (o *DeleteTeamMembersNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteTeamMembersNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/teams/{teamId}/members][%d] deleteTeamMembersNotFound  %+v", 404, o.Payload)
}

func (o *DeleteTeamMembersNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/v2/teams/{teamId}/members][%d] deleteTeamMembersNotFound  %+v", 404, o.Payload)
}

func (o *DeleteTeamMembersNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTeamMembersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTeamMembersRequestTimeout creates a DeleteTeamMembersRequestTimeout with default headers values
func NewDeleteTeamMembersRequestTimeout() *DeleteTeamMembersRequestTimeout {
	return &DeleteTeamMembersRequestTimeout{}
}

/*
DeleteTeamMembersRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type DeleteTeamMembersRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete team members request timeout response has a 2xx status code
func (o *DeleteTeamMembersRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete team members request timeout response has a 3xx status code
func (o *DeleteTeamMembersRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete team members request timeout response has a 4xx status code
func (o *DeleteTeamMembersRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete team members request timeout response has a 5xx status code
func (o *DeleteTeamMembersRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this delete team members request timeout response a status code equal to that given
func (o *DeleteTeamMembersRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *DeleteTeamMembersRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/teams/{teamId}/members][%d] deleteTeamMembersRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteTeamMembersRequestTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/teams/{teamId}/members][%d] deleteTeamMembersRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteTeamMembersRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTeamMembersRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTeamMembersRequestEntityTooLarge creates a DeleteTeamMembersRequestEntityTooLarge with default headers values
func NewDeleteTeamMembersRequestEntityTooLarge() *DeleteTeamMembersRequestEntityTooLarge {
	return &DeleteTeamMembersRequestEntityTooLarge{}
}

/*
DeleteTeamMembersRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type DeleteTeamMembersRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete team members request entity too large response has a 2xx status code
func (o *DeleteTeamMembersRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete team members request entity too large response has a 3xx status code
func (o *DeleteTeamMembersRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete team members request entity too large response has a 4xx status code
func (o *DeleteTeamMembersRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete team members request entity too large response has a 5xx status code
func (o *DeleteTeamMembersRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this delete team members request entity too large response a status code equal to that given
func (o *DeleteTeamMembersRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *DeleteTeamMembersRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/teams/{teamId}/members][%d] deleteTeamMembersRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteTeamMembersRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[DELETE /api/v2/teams/{teamId}/members][%d] deleteTeamMembersRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteTeamMembersRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTeamMembersRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTeamMembersUnsupportedMediaType creates a DeleteTeamMembersUnsupportedMediaType with default headers values
func NewDeleteTeamMembersUnsupportedMediaType() *DeleteTeamMembersUnsupportedMediaType {
	return &DeleteTeamMembersUnsupportedMediaType{}
}

/*
DeleteTeamMembersUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteTeamMembersUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete team members unsupported media type response has a 2xx status code
func (o *DeleteTeamMembersUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete team members unsupported media type response has a 3xx status code
func (o *DeleteTeamMembersUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete team members unsupported media type response has a 4xx status code
func (o *DeleteTeamMembersUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete team members unsupported media type response has a 5xx status code
func (o *DeleteTeamMembersUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this delete team members unsupported media type response a status code equal to that given
func (o *DeleteTeamMembersUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *DeleteTeamMembersUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/teams/{teamId}/members][%d] deleteTeamMembersUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteTeamMembersUnsupportedMediaType) String() string {
	return fmt.Sprintf("[DELETE /api/v2/teams/{teamId}/members][%d] deleteTeamMembersUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteTeamMembersUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTeamMembersUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTeamMembersTooManyRequests creates a DeleteTeamMembersTooManyRequests with default headers values
func NewDeleteTeamMembersTooManyRequests() *DeleteTeamMembersTooManyRequests {
	return &DeleteTeamMembersTooManyRequests{}
}

/*
DeleteTeamMembersTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type DeleteTeamMembersTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete team members too many requests response has a 2xx status code
func (o *DeleteTeamMembersTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete team members too many requests response has a 3xx status code
func (o *DeleteTeamMembersTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete team members too many requests response has a 4xx status code
func (o *DeleteTeamMembersTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete team members too many requests response has a 5xx status code
func (o *DeleteTeamMembersTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this delete team members too many requests response a status code equal to that given
func (o *DeleteTeamMembersTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *DeleteTeamMembersTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/teams/{teamId}/members][%d] deleteTeamMembersTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteTeamMembersTooManyRequests) String() string {
	return fmt.Sprintf("[DELETE /api/v2/teams/{teamId}/members][%d] deleteTeamMembersTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteTeamMembersTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTeamMembersTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTeamMembersInternalServerError creates a DeleteTeamMembersInternalServerError with default headers values
func NewDeleteTeamMembersInternalServerError() *DeleteTeamMembersInternalServerError {
	return &DeleteTeamMembersInternalServerError{}
}

/*
DeleteTeamMembersInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteTeamMembersInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete team members internal server error response has a 2xx status code
func (o *DeleteTeamMembersInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete team members internal server error response has a 3xx status code
func (o *DeleteTeamMembersInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete team members internal server error response has a 4xx status code
func (o *DeleteTeamMembersInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete team members internal server error response has a 5xx status code
func (o *DeleteTeamMembersInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete team members internal server error response a status code equal to that given
func (o *DeleteTeamMembersInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DeleteTeamMembersInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/teams/{teamId}/members][%d] deleteTeamMembersInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteTeamMembersInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /api/v2/teams/{teamId}/members][%d] deleteTeamMembersInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteTeamMembersInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTeamMembersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTeamMembersServiceUnavailable creates a DeleteTeamMembersServiceUnavailable with default headers values
func NewDeleteTeamMembersServiceUnavailable() *DeleteTeamMembersServiceUnavailable {
	return &DeleteTeamMembersServiceUnavailable{}
}

/*
DeleteTeamMembersServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteTeamMembersServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete team members service unavailable response has a 2xx status code
func (o *DeleteTeamMembersServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete team members service unavailable response has a 3xx status code
func (o *DeleteTeamMembersServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete team members service unavailable response has a 4xx status code
func (o *DeleteTeamMembersServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete team members service unavailable response has a 5xx status code
func (o *DeleteTeamMembersServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this delete team members service unavailable response a status code equal to that given
func (o *DeleteTeamMembersServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *DeleteTeamMembersServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/teams/{teamId}/members][%d] deleteTeamMembersServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteTeamMembersServiceUnavailable) String() string {
	return fmt.Sprintf("[DELETE /api/v2/teams/{teamId}/members][%d] deleteTeamMembersServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteTeamMembersServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTeamMembersServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTeamMembersGatewayTimeout creates a DeleteTeamMembersGatewayTimeout with default headers values
func NewDeleteTeamMembersGatewayTimeout() *DeleteTeamMembersGatewayTimeout {
	return &DeleteTeamMembersGatewayTimeout{}
}

/*
DeleteTeamMembersGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type DeleteTeamMembersGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete team members gateway timeout response has a 2xx status code
func (o *DeleteTeamMembersGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete team members gateway timeout response has a 3xx status code
func (o *DeleteTeamMembersGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete team members gateway timeout response has a 4xx status code
func (o *DeleteTeamMembersGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete team members gateway timeout response has a 5xx status code
func (o *DeleteTeamMembersGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this delete team members gateway timeout response a status code equal to that given
func (o *DeleteTeamMembersGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *DeleteTeamMembersGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/teams/{teamId}/members][%d] deleteTeamMembersGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteTeamMembersGatewayTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/teams/{teamId}/members][%d] deleteTeamMembersGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteTeamMembersGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTeamMembersGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

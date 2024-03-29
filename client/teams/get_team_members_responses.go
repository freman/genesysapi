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

// GetTeamMembersReader is a Reader for the GetTeamMembers structure.
type GetTeamMembersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTeamMembersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTeamMembersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetTeamMembersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetTeamMembersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetTeamMembersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTeamMembersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetTeamMembersRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetTeamMembersRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetTeamMembersUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetTeamMembersTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetTeamMembersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetTeamMembersServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetTeamMembersGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTeamMembersOK creates a GetTeamMembersOK with default headers values
func NewGetTeamMembersOK() *GetTeamMembersOK {
	return &GetTeamMembersOK{}
}

/*
GetTeamMembersOK describes a response with status code 200, with default header values.

successful operation
*/
type GetTeamMembersOK struct {
	Payload *models.TeamMemberEntityListing
}

// IsSuccess returns true when this get team members o k response has a 2xx status code
func (o *GetTeamMembersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get team members o k response has a 3xx status code
func (o *GetTeamMembersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get team members o k response has a 4xx status code
func (o *GetTeamMembersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get team members o k response has a 5xx status code
func (o *GetTeamMembersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get team members o k response a status code equal to that given
func (o *GetTeamMembersOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetTeamMembersOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/teams/{teamId}/members][%d] getTeamMembersOK  %+v", 200, o.Payload)
}

func (o *GetTeamMembersOK) String() string {
	return fmt.Sprintf("[GET /api/v2/teams/{teamId}/members][%d] getTeamMembersOK  %+v", 200, o.Payload)
}

func (o *GetTeamMembersOK) GetPayload() *models.TeamMemberEntityListing {
	return o.Payload
}

func (o *GetTeamMembersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TeamMemberEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTeamMembersBadRequest creates a GetTeamMembersBadRequest with default headers values
func NewGetTeamMembersBadRequest() *GetTeamMembersBadRequest {
	return &GetTeamMembersBadRequest{}
}

/*
GetTeamMembersBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetTeamMembersBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get team members bad request response has a 2xx status code
func (o *GetTeamMembersBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get team members bad request response has a 3xx status code
func (o *GetTeamMembersBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get team members bad request response has a 4xx status code
func (o *GetTeamMembersBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get team members bad request response has a 5xx status code
func (o *GetTeamMembersBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get team members bad request response a status code equal to that given
func (o *GetTeamMembersBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetTeamMembersBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/teams/{teamId}/members][%d] getTeamMembersBadRequest  %+v", 400, o.Payload)
}

func (o *GetTeamMembersBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/teams/{teamId}/members][%d] getTeamMembersBadRequest  %+v", 400, o.Payload)
}

func (o *GetTeamMembersBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTeamMembersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTeamMembersUnauthorized creates a GetTeamMembersUnauthorized with default headers values
func NewGetTeamMembersUnauthorized() *GetTeamMembersUnauthorized {
	return &GetTeamMembersUnauthorized{}
}

/*
GetTeamMembersUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetTeamMembersUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get team members unauthorized response has a 2xx status code
func (o *GetTeamMembersUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get team members unauthorized response has a 3xx status code
func (o *GetTeamMembersUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get team members unauthorized response has a 4xx status code
func (o *GetTeamMembersUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get team members unauthorized response has a 5xx status code
func (o *GetTeamMembersUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get team members unauthorized response a status code equal to that given
func (o *GetTeamMembersUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetTeamMembersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/teams/{teamId}/members][%d] getTeamMembersUnauthorized  %+v", 401, o.Payload)
}

func (o *GetTeamMembersUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/teams/{teamId}/members][%d] getTeamMembersUnauthorized  %+v", 401, o.Payload)
}

func (o *GetTeamMembersUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTeamMembersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTeamMembersForbidden creates a GetTeamMembersForbidden with default headers values
func NewGetTeamMembersForbidden() *GetTeamMembersForbidden {
	return &GetTeamMembersForbidden{}
}

/*
GetTeamMembersForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetTeamMembersForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get team members forbidden response has a 2xx status code
func (o *GetTeamMembersForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get team members forbidden response has a 3xx status code
func (o *GetTeamMembersForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get team members forbidden response has a 4xx status code
func (o *GetTeamMembersForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get team members forbidden response has a 5xx status code
func (o *GetTeamMembersForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get team members forbidden response a status code equal to that given
func (o *GetTeamMembersForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetTeamMembersForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/teams/{teamId}/members][%d] getTeamMembersForbidden  %+v", 403, o.Payload)
}

func (o *GetTeamMembersForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/teams/{teamId}/members][%d] getTeamMembersForbidden  %+v", 403, o.Payload)
}

func (o *GetTeamMembersForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTeamMembersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTeamMembersNotFound creates a GetTeamMembersNotFound with default headers values
func NewGetTeamMembersNotFound() *GetTeamMembersNotFound {
	return &GetTeamMembersNotFound{}
}

/*
GetTeamMembersNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetTeamMembersNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get team members not found response has a 2xx status code
func (o *GetTeamMembersNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get team members not found response has a 3xx status code
func (o *GetTeamMembersNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get team members not found response has a 4xx status code
func (o *GetTeamMembersNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get team members not found response has a 5xx status code
func (o *GetTeamMembersNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get team members not found response a status code equal to that given
func (o *GetTeamMembersNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetTeamMembersNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/teams/{teamId}/members][%d] getTeamMembersNotFound  %+v", 404, o.Payload)
}

func (o *GetTeamMembersNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/teams/{teamId}/members][%d] getTeamMembersNotFound  %+v", 404, o.Payload)
}

func (o *GetTeamMembersNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTeamMembersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTeamMembersRequestTimeout creates a GetTeamMembersRequestTimeout with default headers values
func NewGetTeamMembersRequestTimeout() *GetTeamMembersRequestTimeout {
	return &GetTeamMembersRequestTimeout{}
}

/*
GetTeamMembersRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetTeamMembersRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get team members request timeout response has a 2xx status code
func (o *GetTeamMembersRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get team members request timeout response has a 3xx status code
func (o *GetTeamMembersRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get team members request timeout response has a 4xx status code
func (o *GetTeamMembersRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get team members request timeout response has a 5xx status code
func (o *GetTeamMembersRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get team members request timeout response a status code equal to that given
func (o *GetTeamMembersRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetTeamMembersRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/teams/{teamId}/members][%d] getTeamMembersRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetTeamMembersRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/teams/{teamId}/members][%d] getTeamMembersRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetTeamMembersRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTeamMembersRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTeamMembersRequestEntityTooLarge creates a GetTeamMembersRequestEntityTooLarge with default headers values
func NewGetTeamMembersRequestEntityTooLarge() *GetTeamMembersRequestEntityTooLarge {
	return &GetTeamMembersRequestEntityTooLarge{}
}

/*
GetTeamMembersRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetTeamMembersRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get team members request entity too large response has a 2xx status code
func (o *GetTeamMembersRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get team members request entity too large response has a 3xx status code
func (o *GetTeamMembersRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get team members request entity too large response has a 4xx status code
func (o *GetTeamMembersRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get team members request entity too large response has a 5xx status code
func (o *GetTeamMembersRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get team members request entity too large response a status code equal to that given
func (o *GetTeamMembersRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetTeamMembersRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/teams/{teamId}/members][%d] getTeamMembersRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetTeamMembersRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/teams/{teamId}/members][%d] getTeamMembersRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetTeamMembersRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTeamMembersRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTeamMembersUnsupportedMediaType creates a GetTeamMembersUnsupportedMediaType with default headers values
func NewGetTeamMembersUnsupportedMediaType() *GetTeamMembersUnsupportedMediaType {
	return &GetTeamMembersUnsupportedMediaType{}
}

/*
GetTeamMembersUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetTeamMembersUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get team members unsupported media type response has a 2xx status code
func (o *GetTeamMembersUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get team members unsupported media type response has a 3xx status code
func (o *GetTeamMembersUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get team members unsupported media type response has a 4xx status code
func (o *GetTeamMembersUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get team members unsupported media type response has a 5xx status code
func (o *GetTeamMembersUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get team members unsupported media type response a status code equal to that given
func (o *GetTeamMembersUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetTeamMembersUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/teams/{teamId}/members][%d] getTeamMembersUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetTeamMembersUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/teams/{teamId}/members][%d] getTeamMembersUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetTeamMembersUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTeamMembersUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTeamMembersTooManyRequests creates a GetTeamMembersTooManyRequests with default headers values
func NewGetTeamMembersTooManyRequests() *GetTeamMembersTooManyRequests {
	return &GetTeamMembersTooManyRequests{}
}

/*
GetTeamMembersTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetTeamMembersTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get team members too many requests response has a 2xx status code
func (o *GetTeamMembersTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get team members too many requests response has a 3xx status code
func (o *GetTeamMembersTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get team members too many requests response has a 4xx status code
func (o *GetTeamMembersTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get team members too many requests response has a 5xx status code
func (o *GetTeamMembersTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get team members too many requests response a status code equal to that given
func (o *GetTeamMembersTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetTeamMembersTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/teams/{teamId}/members][%d] getTeamMembersTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetTeamMembersTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/teams/{teamId}/members][%d] getTeamMembersTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetTeamMembersTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTeamMembersTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTeamMembersInternalServerError creates a GetTeamMembersInternalServerError with default headers values
func NewGetTeamMembersInternalServerError() *GetTeamMembersInternalServerError {
	return &GetTeamMembersInternalServerError{}
}

/*
GetTeamMembersInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetTeamMembersInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get team members internal server error response has a 2xx status code
func (o *GetTeamMembersInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get team members internal server error response has a 3xx status code
func (o *GetTeamMembersInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get team members internal server error response has a 4xx status code
func (o *GetTeamMembersInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get team members internal server error response has a 5xx status code
func (o *GetTeamMembersInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get team members internal server error response a status code equal to that given
func (o *GetTeamMembersInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetTeamMembersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/teams/{teamId}/members][%d] getTeamMembersInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTeamMembersInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/teams/{teamId}/members][%d] getTeamMembersInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTeamMembersInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTeamMembersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTeamMembersServiceUnavailable creates a GetTeamMembersServiceUnavailable with default headers values
func NewGetTeamMembersServiceUnavailable() *GetTeamMembersServiceUnavailable {
	return &GetTeamMembersServiceUnavailable{}
}

/*
GetTeamMembersServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetTeamMembersServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get team members service unavailable response has a 2xx status code
func (o *GetTeamMembersServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get team members service unavailable response has a 3xx status code
func (o *GetTeamMembersServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get team members service unavailable response has a 4xx status code
func (o *GetTeamMembersServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get team members service unavailable response has a 5xx status code
func (o *GetTeamMembersServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get team members service unavailable response a status code equal to that given
func (o *GetTeamMembersServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetTeamMembersServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/teams/{teamId}/members][%d] getTeamMembersServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetTeamMembersServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/teams/{teamId}/members][%d] getTeamMembersServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetTeamMembersServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTeamMembersServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTeamMembersGatewayTimeout creates a GetTeamMembersGatewayTimeout with default headers values
func NewGetTeamMembersGatewayTimeout() *GetTeamMembersGatewayTimeout {
	return &GetTeamMembersGatewayTimeout{}
}

/*
GetTeamMembersGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetTeamMembersGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get team members gateway timeout response has a 2xx status code
func (o *GetTeamMembersGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get team members gateway timeout response has a 3xx status code
func (o *GetTeamMembersGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get team members gateway timeout response has a 4xx status code
func (o *GetTeamMembersGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get team members gateway timeout response has a 5xx status code
func (o *GetTeamMembersGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get team members gateway timeout response a status code equal to that given
func (o *GetTeamMembersGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetTeamMembersGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/teams/{teamId}/members][%d] getTeamMembersGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetTeamMembersGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/teams/{teamId}/members][%d] getTeamMembersGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetTeamMembersGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTeamMembersGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package routing

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetRoutingSkillgroupMembersReader is a Reader for the GetRoutingSkillgroupMembers structure.
type GetRoutingSkillgroupMembersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRoutingSkillgroupMembersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRoutingSkillgroupMembersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRoutingSkillgroupMembersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetRoutingSkillgroupMembersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRoutingSkillgroupMembersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRoutingSkillgroupMembersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetRoutingSkillgroupMembersRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetRoutingSkillgroupMembersRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetRoutingSkillgroupMembersUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetRoutingSkillgroupMembersTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRoutingSkillgroupMembersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetRoutingSkillgroupMembersServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetRoutingSkillgroupMembersGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRoutingSkillgroupMembersOK creates a GetRoutingSkillgroupMembersOK with default headers values
func NewGetRoutingSkillgroupMembersOK() *GetRoutingSkillgroupMembersOK {
	return &GetRoutingSkillgroupMembersOK{}
}

/*
GetRoutingSkillgroupMembersOK describes a response with status code 200, with default header values.

successful operation
*/
type GetRoutingSkillgroupMembersOK struct {
	Payload *models.SkillGroupMemberEntityListing
}

// IsSuccess returns true when this get routing skillgroup members o k response has a 2xx status code
func (o *GetRoutingSkillgroupMembersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get routing skillgroup members o k response has a 3xx status code
func (o *GetRoutingSkillgroupMembersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing skillgroup members o k response has a 4xx status code
func (o *GetRoutingSkillgroupMembersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get routing skillgroup members o k response has a 5xx status code
func (o *GetRoutingSkillgroupMembersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing skillgroup members o k response a status code equal to that given
func (o *GetRoutingSkillgroupMembersOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetRoutingSkillgroupMembersOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/skillgroups/{skillGroupId}/members][%d] getRoutingSkillgroupMembersOK  %+v", 200, o.Payload)
}

func (o *GetRoutingSkillgroupMembersOK) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/skillgroups/{skillGroupId}/members][%d] getRoutingSkillgroupMembersOK  %+v", 200, o.Payload)
}

func (o *GetRoutingSkillgroupMembersOK) GetPayload() *models.SkillGroupMemberEntityListing {
	return o.Payload
}

func (o *GetRoutingSkillgroupMembersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SkillGroupMemberEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSkillgroupMembersBadRequest creates a GetRoutingSkillgroupMembersBadRequest with default headers values
func NewGetRoutingSkillgroupMembersBadRequest() *GetRoutingSkillgroupMembersBadRequest {
	return &GetRoutingSkillgroupMembersBadRequest{}
}

/*
GetRoutingSkillgroupMembersBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetRoutingSkillgroupMembersBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing skillgroup members bad request response has a 2xx status code
func (o *GetRoutingSkillgroupMembersBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing skillgroup members bad request response has a 3xx status code
func (o *GetRoutingSkillgroupMembersBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing skillgroup members bad request response has a 4xx status code
func (o *GetRoutingSkillgroupMembersBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing skillgroup members bad request response has a 5xx status code
func (o *GetRoutingSkillgroupMembersBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing skillgroup members bad request response a status code equal to that given
func (o *GetRoutingSkillgroupMembersBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetRoutingSkillgroupMembersBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/skillgroups/{skillGroupId}/members][%d] getRoutingSkillgroupMembersBadRequest  %+v", 400, o.Payload)
}

func (o *GetRoutingSkillgroupMembersBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/skillgroups/{skillGroupId}/members][%d] getRoutingSkillgroupMembersBadRequest  %+v", 400, o.Payload)
}

func (o *GetRoutingSkillgroupMembersBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSkillgroupMembersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSkillgroupMembersUnauthorized creates a GetRoutingSkillgroupMembersUnauthorized with default headers values
func NewGetRoutingSkillgroupMembersUnauthorized() *GetRoutingSkillgroupMembersUnauthorized {
	return &GetRoutingSkillgroupMembersUnauthorized{}
}

/*
GetRoutingSkillgroupMembersUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetRoutingSkillgroupMembersUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing skillgroup members unauthorized response has a 2xx status code
func (o *GetRoutingSkillgroupMembersUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing skillgroup members unauthorized response has a 3xx status code
func (o *GetRoutingSkillgroupMembersUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing skillgroup members unauthorized response has a 4xx status code
func (o *GetRoutingSkillgroupMembersUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing skillgroup members unauthorized response has a 5xx status code
func (o *GetRoutingSkillgroupMembersUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing skillgroup members unauthorized response a status code equal to that given
func (o *GetRoutingSkillgroupMembersUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetRoutingSkillgroupMembersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/skillgroups/{skillGroupId}/members][%d] getRoutingSkillgroupMembersUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRoutingSkillgroupMembersUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/skillgroups/{skillGroupId}/members][%d] getRoutingSkillgroupMembersUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRoutingSkillgroupMembersUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSkillgroupMembersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSkillgroupMembersForbidden creates a GetRoutingSkillgroupMembersForbidden with default headers values
func NewGetRoutingSkillgroupMembersForbidden() *GetRoutingSkillgroupMembersForbidden {
	return &GetRoutingSkillgroupMembersForbidden{}
}

/*
GetRoutingSkillgroupMembersForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetRoutingSkillgroupMembersForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing skillgroup members forbidden response has a 2xx status code
func (o *GetRoutingSkillgroupMembersForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing skillgroup members forbidden response has a 3xx status code
func (o *GetRoutingSkillgroupMembersForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing skillgroup members forbidden response has a 4xx status code
func (o *GetRoutingSkillgroupMembersForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing skillgroup members forbidden response has a 5xx status code
func (o *GetRoutingSkillgroupMembersForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing skillgroup members forbidden response a status code equal to that given
func (o *GetRoutingSkillgroupMembersForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetRoutingSkillgroupMembersForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/skillgroups/{skillGroupId}/members][%d] getRoutingSkillgroupMembersForbidden  %+v", 403, o.Payload)
}

func (o *GetRoutingSkillgroupMembersForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/skillgroups/{skillGroupId}/members][%d] getRoutingSkillgroupMembersForbidden  %+v", 403, o.Payload)
}

func (o *GetRoutingSkillgroupMembersForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSkillgroupMembersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSkillgroupMembersNotFound creates a GetRoutingSkillgroupMembersNotFound with default headers values
func NewGetRoutingSkillgroupMembersNotFound() *GetRoutingSkillgroupMembersNotFound {
	return &GetRoutingSkillgroupMembersNotFound{}
}

/*
GetRoutingSkillgroupMembersNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetRoutingSkillgroupMembersNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing skillgroup members not found response has a 2xx status code
func (o *GetRoutingSkillgroupMembersNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing skillgroup members not found response has a 3xx status code
func (o *GetRoutingSkillgroupMembersNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing skillgroup members not found response has a 4xx status code
func (o *GetRoutingSkillgroupMembersNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing skillgroup members not found response has a 5xx status code
func (o *GetRoutingSkillgroupMembersNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing skillgroup members not found response a status code equal to that given
func (o *GetRoutingSkillgroupMembersNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetRoutingSkillgroupMembersNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/skillgroups/{skillGroupId}/members][%d] getRoutingSkillgroupMembersNotFound  %+v", 404, o.Payload)
}

func (o *GetRoutingSkillgroupMembersNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/skillgroups/{skillGroupId}/members][%d] getRoutingSkillgroupMembersNotFound  %+v", 404, o.Payload)
}

func (o *GetRoutingSkillgroupMembersNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSkillgroupMembersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSkillgroupMembersRequestTimeout creates a GetRoutingSkillgroupMembersRequestTimeout with default headers values
func NewGetRoutingSkillgroupMembersRequestTimeout() *GetRoutingSkillgroupMembersRequestTimeout {
	return &GetRoutingSkillgroupMembersRequestTimeout{}
}

/*
GetRoutingSkillgroupMembersRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetRoutingSkillgroupMembersRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing skillgroup members request timeout response has a 2xx status code
func (o *GetRoutingSkillgroupMembersRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing skillgroup members request timeout response has a 3xx status code
func (o *GetRoutingSkillgroupMembersRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing skillgroup members request timeout response has a 4xx status code
func (o *GetRoutingSkillgroupMembersRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing skillgroup members request timeout response has a 5xx status code
func (o *GetRoutingSkillgroupMembersRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing skillgroup members request timeout response a status code equal to that given
func (o *GetRoutingSkillgroupMembersRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetRoutingSkillgroupMembersRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/skillgroups/{skillGroupId}/members][%d] getRoutingSkillgroupMembersRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetRoutingSkillgroupMembersRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/skillgroups/{skillGroupId}/members][%d] getRoutingSkillgroupMembersRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetRoutingSkillgroupMembersRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSkillgroupMembersRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSkillgroupMembersRequestEntityTooLarge creates a GetRoutingSkillgroupMembersRequestEntityTooLarge with default headers values
func NewGetRoutingSkillgroupMembersRequestEntityTooLarge() *GetRoutingSkillgroupMembersRequestEntityTooLarge {
	return &GetRoutingSkillgroupMembersRequestEntityTooLarge{}
}

/*
GetRoutingSkillgroupMembersRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetRoutingSkillgroupMembersRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing skillgroup members request entity too large response has a 2xx status code
func (o *GetRoutingSkillgroupMembersRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing skillgroup members request entity too large response has a 3xx status code
func (o *GetRoutingSkillgroupMembersRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing skillgroup members request entity too large response has a 4xx status code
func (o *GetRoutingSkillgroupMembersRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing skillgroup members request entity too large response has a 5xx status code
func (o *GetRoutingSkillgroupMembersRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing skillgroup members request entity too large response a status code equal to that given
func (o *GetRoutingSkillgroupMembersRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetRoutingSkillgroupMembersRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/skillgroups/{skillGroupId}/members][%d] getRoutingSkillgroupMembersRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetRoutingSkillgroupMembersRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/skillgroups/{skillGroupId}/members][%d] getRoutingSkillgroupMembersRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetRoutingSkillgroupMembersRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSkillgroupMembersRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSkillgroupMembersUnsupportedMediaType creates a GetRoutingSkillgroupMembersUnsupportedMediaType with default headers values
func NewGetRoutingSkillgroupMembersUnsupportedMediaType() *GetRoutingSkillgroupMembersUnsupportedMediaType {
	return &GetRoutingSkillgroupMembersUnsupportedMediaType{}
}

/*
GetRoutingSkillgroupMembersUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetRoutingSkillgroupMembersUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing skillgroup members unsupported media type response has a 2xx status code
func (o *GetRoutingSkillgroupMembersUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing skillgroup members unsupported media type response has a 3xx status code
func (o *GetRoutingSkillgroupMembersUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing skillgroup members unsupported media type response has a 4xx status code
func (o *GetRoutingSkillgroupMembersUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing skillgroup members unsupported media type response has a 5xx status code
func (o *GetRoutingSkillgroupMembersUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing skillgroup members unsupported media type response a status code equal to that given
func (o *GetRoutingSkillgroupMembersUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetRoutingSkillgroupMembersUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/skillgroups/{skillGroupId}/members][%d] getRoutingSkillgroupMembersUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetRoutingSkillgroupMembersUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/skillgroups/{skillGroupId}/members][%d] getRoutingSkillgroupMembersUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetRoutingSkillgroupMembersUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSkillgroupMembersUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSkillgroupMembersTooManyRequests creates a GetRoutingSkillgroupMembersTooManyRequests with default headers values
func NewGetRoutingSkillgroupMembersTooManyRequests() *GetRoutingSkillgroupMembersTooManyRequests {
	return &GetRoutingSkillgroupMembersTooManyRequests{}
}

/*
GetRoutingSkillgroupMembersTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetRoutingSkillgroupMembersTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing skillgroup members too many requests response has a 2xx status code
func (o *GetRoutingSkillgroupMembersTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing skillgroup members too many requests response has a 3xx status code
func (o *GetRoutingSkillgroupMembersTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing skillgroup members too many requests response has a 4xx status code
func (o *GetRoutingSkillgroupMembersTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing skillgroup members too many requests response has a 5xx status code
func (o *GetRoutingSkillgroupMembersTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing skillgroup members too many requests response a status code equal to that given
func (o *GetRoutingSkillgroupMembersTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetRoutingSkillgroupMembersTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/skillgroups/{skillGroupId}/members][%d] getRoutingSkillgroupMembersTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetRoutingSkillgroupMembersTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/skillgroups/{skillGroupId}/members][%d] getRoutingSkillgroupMembersTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetRoutingSkillgroupMembersTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSkillgroupMembersTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSkillgroupMembersInternalServerError creates a GetRoutingSkillgroupMembersInternalServerError with default headers values
func NewGetRoutingSkillgroupMembersInternalServerError() *GetRoutingSkillgroupMembersInternalServerError {
	return &GetRoutingSkillgroupMembersInternalServerError{}
}

/*
GetRoutingSkillgroupMembersInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetRoutingSkillgroupMembersInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing skillgroup members internal server error response has a 2xx status code
func (o *GetRoutingSkillgroupMembersInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing skillgroup members internal server error response has a 3xx status code
func (o *GetRoutingSkillgroupMembersInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing skillgroup members internal server error response has a 4xx status code
func (o *GetRoutingSkillgroupMembersInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get routing skillgroup members internal server error response has a 5xx status code
func (o *GetRoutingSkillgroupMembersInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get routing skillgroup members internal server error response a status code equal to that given
func (o *GetRoutingSkillgroupMembersInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetRoutingSkillgroupMembersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/skillgroups/{skillGroupId}/members][%d] getRoutingSkillgroupMembersInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRoutingSkillgroupMembersInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/skillgroups/{skillGroupId}/members][%d] getRoutingSkillgroupMembersInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRoutingSkillgroupMembersInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSkillgroupMembersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSkillgroupMembersServiceUnavailable creates a GetRoutingSkillgroupMembersServiceUnavailable with default headers values
func NewGetRoutingSkillgroupMembersServiceUnavailable() *GetRoutingSkillgroupMembersServiceUnavailable {
	return &GetRoutingSkillgroupMembersServiceUnavailable{}
}

/*
GetRoutingSkillgroupMembersServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetRoutingSkillgroupMembersServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing skillgroup members service unavailable response has a 2xx status code
func (o *GetRoutingSkillgroupMembersServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing skillgroup members service unavailable response has a 3xx status code
func (o *GetRoutingSkillgroupMembersServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing skillgroup members service unavailable response has a 4xx status code
func (o *GetRoutingSkillgroupMembersServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get routing skillgroup members service unavailable response has a 5xx status code
func (o *GetRoutingSkillgroupMembersServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get routing skillgroup members service unavailable response a status code equal to that given
func (o *GetRoutingSkillgroupMembersServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetRoutingSkillgroupMembersServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/skillgroups/{skillGroupId}/members][%d] getRoutingSkillgroupMembersServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetRoutingSkillgroupMembersServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/skillgroups/{skillGroupId}/members][%d] getRoutingSkillgroupMembersServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetRoutingSkillgroupMembersServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSkillgroupMembersServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSkillgroupMembersGatewayTimeout creates a GetRoutingSkillgroupMembersGatewayTimeout with default headers values
func NewGetRoutingSkillgroupMembersGatewayTimeout() *GetRoutingSkillgroupMembersGatewayTimeout {
	return &GetRoutingSkillgroupMembersGatewayTimeout{}
}

/*
GetRoutingSkillgroupMembersGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetRoutingSkillgroupMembersGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing skillgroup members gateway timeout response has a 2xx status code
func (o *GetRoutingSkillgroupMembersGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing skillgroup members gateway timeout response has a 3xx status code
func (o *GetRoutingSkillgroupMembersGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing skillgroup members gateway timeout response has a 4xx status code
func (o *GetRoutingSkillgroupMembersGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get routing skillgroup members gateway timeout response has a 5xx status code
func (o *GetRoutingSkillgroupMembersGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get routing skillgroup members gateway timeout response a status code equal to that given
func (o *GetRoutingSkillgroupMembersGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetRoutingSkillgroupMembersGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/skillgroups/{skillGroupId}/members][%d] getRoutingSkillgroupMembersGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetRoutingSkillgroupMembersGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/skillgroups/{skillGroupId}/members][%d] getRoutingSkillgroupMembersGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetRoutingSkillgroupMembersGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSkillgroupMembersGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

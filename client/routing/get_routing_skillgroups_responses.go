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

// GetRoutingSkillgroupsReader is a Reader for the GetRoutingSkillgroups structure.
type GetRoutingSkillgroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRoutingSkillgroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRoutingSkillgroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRoutingSkillgroupsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetRoutingSkillgroupsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRoutingSkillgroupsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRoutingSkillgroupsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetRoutingSkillgroupsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetRoutingSkillgroupsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetRoutingSkillgroupsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetRoutingSkillgroupsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRoutingSkillgroupsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetRoutingSkillgroupsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetRoutingSkillgroupsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRoutingSkillgroupsOK creates a GetRoutingSkillgroupsOK with default headers values
func NewGetRoutingSkillgroupsOK() *GetRoutingSkillgroupsOK {
	return &GetRoutingSkillgroupsOK{}
}

/*
GetRoutingSkillgroupsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetRoutingSkillgroupsOK struct {
	Payload *models.SkillGroupEntityListing
}

// IsSuccess returns true when this get routing skillgroups o k response has a 2xx status code
func (o *GetRoutingSkillgroupsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get routing skillgroups o k response has a 3xx status code
func (o *GetRoutingSkillgroupsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing skillgroups o k response has a 4xx status code
func (o *GetRoutingSkillgroupsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get routing skillgroups o k response has a 5xx status code
func (o *GetRoutingSkillgroupsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing skillgroups o k response a status code equal to that given
func (o *GetRoutingSkillgroupsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetRoutingSkillgroupsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/skillgroups][%d] getRoutingSkillgroupsOK  %+v", 200, o.Payload)
}

func (o *GetRoutingSkillgroupsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/skillgroups][%d] getRoutingSkillgroupsOK  %+v", 200, o.Payload)
}

func (o *GetRoutingSkillgroupsOK) GetPayload() *models.SkillGroupEntityListing {
	return o.Payload
}

func (o *GetRoutingSkillgroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SkillGroupEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSkillgroupsBadRequest creates a GetRoutingSkillgroupsBadRequest with default headers values
func NewGetRoutingSkillgroupsBadRequest() *GetRoutingSkillgroupsBadRequest {
	return &GetRoutingSkillgroupsBadRequest{}
}

/*
GetRoutingSkillgroupsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetRoutingSkillgroupsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing skillgroups bad request response has a 2xx status code
func (o *GetRoutingSkillgroupsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing skillgroups bad request response has a 3xx status code
func (o *GetRoutingSkillgroupsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing skillgroups bad request response has a 4xx status code
func (o *GetRoutingSkillgroupsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing skillgroups bad request response has a 5xx status code
func (o *GetRoutingSkillgroupsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing skillgroups bad request response a status code equal to that given
func (o *GetRoutingSkillgroupsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetRoutingSkillgroupsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/skillgroups][%d] getRoutingSkillgroupsBadRequest  %+v", 400, o.Payload)
}

func (o *GetRoutingSkillgroupsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/skillgroups][%d] getRoutingSkillgroupsBadRequest  %+v", 400, o.Payload)
}

func (o *GetRoutingSkillgroupsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSkillgroupsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSkillgroupsUnauthorized creates a GetRoutingSkillgroupsUnauthorized with default headers values
func NewGetRoutingSkillgroupsUnauthorized() *GetRoutingSkillgroupsUnauthorized {
	return &GetRoutingSkillgroupsUnauthorized{}
}

/*
GetRoutingSkillgroupsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetRoutingSkillgroupsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing skillgroups unauthorized response has a 2xx status code
func (o *GetRoutingSkillgroupsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing skillgroups unauthorized response has a 3xx status code
func (o *GetRoutingSkillgroupsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing skillgroups unauthorized response has a 4xx status code
func (o *GetRoutingSkillgroupsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing skillgroups unauthorized response has a 5xx status code
func (o *GetRoutingSkillgroupsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing skillgroups unauthorized response a status code equal to that given
func (o *GetRoutingSkillgroupsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetRoutingSkillgroupsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/skillgroups][%d] getRoutingSkillgroupsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRoutingSkillgroupsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/skillgroups][%d] getRoutingSkillgroupsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRoutingSkillgroupsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSkillgroupsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSkillgroupsForbidden creates a GetRoutingSkillgroupsForbidden with default headers values
func NewGetRoutingSkillgroupsForbidden() *GetRoutingSkillgroupsForbidden {
	return &GetRoutingSkillgroupsForbidden{}
}

/*
GetRoutingSkillgroupsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetRoutingSkillgroupsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing skillgroups forbidden response has a 2xx status code
func (o *GetRoutingSkillgroupsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing skillgroups forbidden response has a 3xx status code
func (o *GetRoutingSkillgroupsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing skillgroups forbidden response has a 4xx status code
func (o *GetRoutingSkillgroupsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing skillgroups forbidden response has a 5xx status code
func (o *GetRoutingSkillgroupsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing skillgroups forbidden response a status code equal to that given
func (o *GetRoutingSkillgroupsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetRoutingSkillgroupsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/skillgroups][%d] getRoutingSkillgroupsForbidden  %+v", 403, o.Payload)
}

func (o *GetRoutingSkillgroupsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/skillgroups][%d] getRoutingSkillgroupsForbidden  %+v", 403, o.Payload)
}

func (o *GetRoutingSkillgroupsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSkillgroupsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSkillgroupsNotFound creates a GetRoutingSkillgroupsNotFound with default headers values
func NewGetRoutingSkillgroupsNotFound() *GetRoutingSkillgroupsNotFound {
	return &GetRoutingSkillgroupsNotFound{}
}

/*
GetRoutingSkillgroupsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetRoutingSkillgroupsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing skillgroups not found response has a 2xx status code
func (o *GetRoutingSkillgroupsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing skillgroups not found response has a 3xx status code
func (o *GetRoutingSkillgroupsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing skillgroups not found response has a 4xx status code
func (o *GetRoutingSkillgroupsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing skillgroups not found response has a 5xx status code
func (o *GetRoutingSkillgroupsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing skillgroups not found response a status code equal to that given
func (o *GetRoutingSkillgroupsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetRoutingSkillgroupsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/skillgroups][%d] getRoutingSkillgroupsNotFound  %+v", 404, o.Payload)
}

func (o *GetRoutingSkillgroupsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/skillgroups][%d] getRoutingSkillgroupsNotFound  %+v", 404, o.Payload)
}

func (o *GetRoutingSkillgroupsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSkillgroupsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSkillgroupsRequestTimeout creates a GetRoutingSkillgroupsRequestTimeout with default headers values
func NewGetRoutingSkillgroupsRequestTimeout() *GetRoutingSkillgroupsRequestTimeout {
	return &GetRoutingSkillgroupsRequestTimeout{}
}

/*
GetRoutingSkillgroupsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetRoutingSkillgroupsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing skillgroups request timeout response has a 2xx status code
func (o *GetRoutingSkillgroupsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing skillgroups request timeout response has a 3xx status code
func (o *GetRoutingSkillgroupsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing skillgroups request timeout response has a 4xx status code
func (o *GetRoutingSkillgroupsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing skillgroups request timeout response has a 5xx status code
func (o *GetRoutingSkillgroupsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing skillgroups request timeout response a status code equal to that given
func (o *GetRoutingSkillgroupsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetRoutingSkillgroupsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/skillgroups][%d] getRoutingSkillgroupsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetRoutingSkillgroupsRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/skillgroups][%d] getRoutingSkillgroupsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetRoutingSkillgroupsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSkillgroupsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSkillgroupsRequestEntityTooLarge creates a GetRoutingSkillgroupsRequestEntityTooLarge with default headers values
func NewGetRoutingSkillgroupsRequestEntityTooLarge() *GetRoutingSkillgroupsRequestEntityTooLarge {
	return &GetRoutingSkillgroupsRequestEntityTooLarge{}
}

/*
GetRoutingSkillgroupsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetRoutingSkillgroupsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing skillgroups request entity too large response has a 2xx status code
func (o *GetRoutingSkillgroupsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing skillgroups request entity too large response has a 3xx status code
func (o *GetRoutingSkillgroupsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing skillgroups request entity too large response has a 4xx status code
func (o *GetRoutingSkillgroupsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing skillgroups request entity too large response has a 5xx status code
func (o *GetRoutingSkillgroupsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing skillgroups request entity too large response a status code equal to that given
func (o *GetRoutingSkillgroupsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetRoutingSkillgroupsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/skillgroups][%d] getRoutingSkillgroupsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetRoutingSkillgroupsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/skillgroups][%d] getRoutingSkillgroupsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetRoutingSkillgroupsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSkillgroupsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSkillgroupsUnsupportedMediaType creates a GetRoutingSkillgroupsUnsupportedMediaType with default headers values
func NewGetRoutingSkillgroupsUnsupportedMediaType() *GetRoutingSkillgroupsUnsupportedMediaType {
	return &GetRoutingSkillgroupsUnsupportedMediaType{}
}

/*
GetRoutingSkillgroupsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetRoutingSkillgroupsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing skillgroups unsupported media type response has a 2xx status code
func (o *GetRoutingSkillgroupsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing skillgroups unsupported media type response has a 3xx status code
func (o *GetRoutingSkillgroupsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing skillgroups unsupported media type response has a 4xx status code
func (o *GetRoutingSkillgroupsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing skillgroups unsupported media type response has a 5xx status code
func (o *GetRoutingSkillgroupsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing skillgroups unsupported media type response a status code equal to that given
func (o *GetRoutingSkillgroupsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetRoutingSkillgroupsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/skillgroups][%d] getRoutingSkillgroupsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetRoutingSkillgroupsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/skillgroups][%d] getRoutingSkillgroupsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetRoutingSkillgroupsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSkillgroupsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSkillgroupsTooManyRequests creates a GetRoutingSkillgroupsTooManyRequests with default headers values
func NewGetRoutingSkillgroupsTooManyRequests() *GetRoutingSkillgroupsTooManyRequests {
	return &GetRoutingSkillgroupsTooManyRequests{}
}

/*
GetRoutingSkillgroupsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetRoutingSkillgroupsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing skillgroups too many requests response has a 2xx status code
func (o *GetRoutingSkillgroupsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing skillgroups too many requests response has a 3xx status code
func (o *GetRoutingSkillgroupsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing skillgroups too many requests response has a 4xx status code
func (o *GetRoutingSkillgroupsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing skillgroups too many requests response has a 5xx status code
func (o *GetRoutingSkillgroupsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing skillgroups too many requests response a status code equal to that given
func (o *GetRoutingSkillgroupsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetRoutingSkillgroupsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/skillgroups][%d] getRoutingSkillgroupsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetRoutingSkillgroupsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/skillgroups][%d] getRoutingSkillgroupsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetRoutingSkillgroupsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSkillgroupsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSkillgroupsInternalServerError creates a GetRoutingSkillgroupsInternalServerError with default headers values
func NewGetRoutingSkillgroupsInternalServerError() *GetRoutingSkillgroupsInternalServerError {
	return &GetRoutingSkillgroupsInternalServerError{}
}

/*
GetRoutingSkillgroupsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetRoutingSkillgroupsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing skillgroups internal server error response has a 2xx status code
func (o *GetRoutingSkillgroupsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing skillgroups internal server error response has a 3xx status code
func (o *GetRoutingSkillgroupsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing skillgroups internal server error response has a 4xx status code
func (o *GetRoutingSkillgroupsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get routing skillgroups internal server error response has a 5xx status code
func (o *GetRoutingSkillgroupsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get routing skillgroups internal server error response a status code equal to that given
func (o *GetRoutingSkillgroupsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetRoutingSkillgroupsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/skillgroups][%d] getRoutingSkillgroupsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRoutingSkillgroupsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/skillgroups][%d] getRoutingSkillgroupsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRoutingSkillgroupsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSkillgroupsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSkillgroupsServiceUnavailable creates a GetRoutingSkillgroupsServiceUnavailable with default headers values
func NewGetRoutingSkillgroupsServiceUnavailable() *GetRoutingSkillgroupsServiceUnavailable {
	return &GetRoutingSkillgroupsServiceUnavailable{}
}

/*
GetRoutingSkillgroupsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetRoutingSkillgroupsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing skillgroups service unavailable response has a 2xx status code
func (o *GetRoutingSkillgroupsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing skillgroups service unavailable response has a 3xx status code
func (o *GetRoutingSkillgroupsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing skillgroups service unavailable response has a 4xx status code
func (o *GetRoutingSkillgroupsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get routing skillgroups service unavailable response has a 5xx status code
func (o *GetRoutingSkillgroupsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get routing skillgroups service unavailable response a status code equal to that given
func (o *GetRoutingSkillgroupsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetRoutingSkillgroupsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/skillgroups][%d] getRoutingSkillgroupsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetRoutingSkillgroupsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/skillgroups][%d] getRoutingSkillgroupsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetRoutingSkillgroupsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSkillgroupsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSkillgroupsGatewayTimeout creates a GetRoutingSkillgroupsGatewayTimeout with default headers values
func NewGetRoutingSkillgroupsGatewayTimeout() *GetRoutingSkillgroupsGatewayTimeout {
	return &GetRoutingSkillgroupsGatewayTimeout{}
}

/*
GetRoutingSkillgroupsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetRoutingSkillgroupsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing skillgroups gateway timeout response has a 2xx status code
func (o *GetRoutingSkillgroupsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing skillgroups gateway timeout response has a 3xx status code
func (o *GetRoutingSkillgroupsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing skillgroups gateway timeout response has a 4xx status code
func (o *GetRoutingSkillgroupsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get routing skillgroups gateway timeout response has a 5xx status code
func (o *GetRoutingSkillgroupsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get routing skillgroups gateway timeout response a status code equal to that given
func (o *GetRoutingSkillgroupsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetRoutingSkillgroupsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/skillgroups][%d] getRoutingSkillgroupsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetRoutingSkillgroupsGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/skillgroups][%d] getRoutingSkillgroupsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetRoutingSkillgroupsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSkillgroupsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

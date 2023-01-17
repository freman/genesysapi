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

// DeleteRoutingSkillgroupReader is a Reader for the DeleteRoutingSkillgroup structure.
type DeleteRoutingSkillgroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRoutingSkillgroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteRoutingSkillgroupNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteRoutingSkillgroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteRoutingSkillgroupUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteRoutingSkillgroupForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteRoutingSkillgroupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewDeleteRoutingSkillgroupRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteRoutingSkillgroupRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteRoutingSkillgroupUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteRoutingSkillgroupTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteRoutingSkillgroupInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteRoutingSkillgroupServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteRoutingSkillgroupGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteRoutingSkillgroupNoContent creates a DeleteRoutingSkillgroupNoContent with default headers values
func NewDeleteRoutingSkillgroupNoContent() *DeleteRoutingSkillgroupNoContent {
	return &DeleteRoutingSkillgroupNoContent{}
}

/*
DeleteRoutingSkillgroupNoContent describes a response with status code 204, with default header values.

Success, skill group was removed
*/
type DeleteRoutingSkillgroupNoContent struct {
}

// IsSuccess returns true when this delete routing skillgroup no content response has a 2xx status code
func (o *DeleteRoutingSkillgroupNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete routing skillgroup no content response has a 3xx status code
func (o *DeleteRoutingSkillgroupNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete routing skillgroup no content response has a 4xx status code
func (o *DeleteRoutingSkillgroupNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete routing skillgroup no content response has a 5xx status code
func (o *DeleteRoutingSkillgroupNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete routing skillgroup no content response a status code equal to that given
func (o *DeleteRoutingSkillgroupNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DeleteRoutingSkillgroupNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/skillgroups/{skillGroupId}][%d] deleteRoutingSkillgroupNoContent ", 204)
}

func (o *DeleteRoutingSkillgroupNoContent) String() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/skillgroups/{skillGroupId}][%d] deleteRoutingSkillgroupNoContent ", 204)
}

func (o *DeleteRoutingSkillgroupNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteRoutingSkillgroupBadRequest creates a DeleteRoutingSkillgroupBadRequest with default headers values
func NewDeleteRoutingSkillgroupBadRequest() *DeleteRoutingSkillgroupBadRequest {
	return &DeleteRoutingSkillgroupBadRequest{}
}

/*
DeleteRoutingSkillgroupBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteRoutingSkillgroupBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete routing skillgroup bad request response has a 2xx status code
func (o *DeleteRoutingSkillgroupBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete routing skillgroup bad request response has a 3xx status code
func (o *DeleteRoutingSkillgroupBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete routing skillgroup bad request response has a 4xx status code
func (o *DeleteRoutingSkillgroupBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete routing skillgroup bad request response has a 5xx status code
func (o *DeleteRoutingSkillgroupBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete routing skillgroup bad request response a status code equal to that given
func (o *DeleteRoutingSkillgroupBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DeleteRoutingSkillgroupBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/skillgroups/{skillGroupId}][%d] deleteRoutingSkillgroupBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteRoutingSkillgroupBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/skillgroups/{skillGroupId}][%d] deleteRoutingSkillgroupBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteRoutingSkillgroupBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingSkillgroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingSkillgroupUnauthorized creates a DeleteRoutingSkillgroupUnauthorized with default headers values
func NewDeleteRoutingSkillgroupUnauthorized() *DeleteRoutingSkillgroupUnauthorized {
	return &DeleteRoutingSkillgroupUnauthorized{}
}

/*
DeleteRoutingSkillgroupUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteRoutingSkillgroupUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete routing skillgroup unauthorized response has a 2xx status code
func (o *DeleteRoutingSkillgroupUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete routing skillgroup unauthorized response has a 3xx status code
func (o *DeleteRoutingSkillgroupUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete routing skillgroup unauthorized response has a 4xx status code
func (o *DeleteRoutingSkillgroupUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete routing skillgroup unauthorized response has a 5xx status code
func (o *DeleteRoutingSkillgroupUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete routing skillgroup unauthorized response a status code equal to that given
func (o *DeleteRoutingSkillgroupUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteRoutingSkillgroupUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/skillgroups/{skillGroupId}][%d] deleteRoutingSkillgroupUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteRoutingSkillgroupUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/skillgroups/{skillGroupId}][%d] deleteRoutingSkillgroupUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteRoutingSkillgroupUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingSkillgroupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingSkillgroupForbidden creates a DeleteRoutingSkillgroupForbidden with default headers values
func NewDeleteRoutingSkillgroupForbidden() *DeleteRoutingSkillgroupForbidden {
	return &DeleteRoutingSkillgroupForbidden{}
}

/*
DeleteRoutingSkillgroupForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type DeleteRoutingSkillgroupForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete routing skillgroup forbidden response has a 2xx status code
func (o *DeleteRoutingSkillgroupForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete routing skillgroup forbidden response has a 3xx status code
func (o *DeleteRoutingSkillgroupForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete routing skillgroup forbidden response has a 4xx status code
func (o *DeleteRoutingSkillgroupForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete routing skillgroup forbidden response has a 5xx status code
func (o *DeleteRoutingSkillgroupForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete routing skillgroup forbidden response a status code equal to that given
func (o *DeleteRoutingSkillgroupForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteRoutingSkillgroupForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/skillgroups/{skillGroupId}][%d] deleteRoutingSkillgroupForbidden  %+v", 403, o.Payload)
}

func (o *DeleteRoutingSkillgroupForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/skillgroups/{skillGroupId}][%d] deleteRoutingSkillgroupForbidden  %+v", 403, o.Payload)
}

func (o *DeleteRoutingSkillgroupForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingSkillgroupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingSkillgroupNotFound creates a DeleteRoutingSkillgroupNotFound with default headers values
func NewDeleteRoutingSkillgroupNotFound() *DeleteRoutingSkillgroupNotFound {
	return &DeleteRoutingSkillgroupNotFound{}
}

/*
DeleteRoutingSkillgroupNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type DeleteRoutingSkillgroupNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete routing skillgroup not found response has a 2xx status code
func (o *DeleteRoutingSkillgroupNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete routing skillgroup not found response has a 3xx status code
func (o *DeleteRoutingSkillgroupNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete routing skillgroup not found response has a 4xx status code
func (o *DeleteRoutingSkillgroupNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete routing skillgroup not found response has a 5xx status code
func (o *DeleteRoutingSkillgroupNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete routing skillgroup not found response a status code equal to that given
func (o *DeleteRoutingSkillgroupNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteRoutingSkillgroupNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/skillgroups/{skillGroupId}][%d] deleteRoutingSkillgroupNotFound  %+v", 404, o.Payload)
}

func (o *DeleteRoutingSkillgroupNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/skillgroups/{skillGroupId}][%d] deleteRoutingSkillgroupNotFound  %+v", 404, o.Payload)
}

func (o *DeleteRoutingSkillgroupNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingSkillgroupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingSkillgroupRequestTimeout creates a DeleteRoutingSkillgroupRequestTimeout with default headers values
func NewDeleteRoutingSkillgroupRequestTimeout() *DeleteRoutingSkillgroupRequestTimeout {
	return &DeleteRoutingSkillgroupRequestTimeout{}
}

/*
DeleteRoutingSkillgroupRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type DeleteRoutingSkillgroupRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete routing skillgroup request timeout response has a 2xx status code
func (o *DeleteRoutingSkillgroupRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete routing skillgroup request timeout response has a 3xx status code
func (o *DeleteRoutingSkillgroupRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete routing skillgroup request timeout response has a 4xx status code
func (o *DeleteRoutingSkillgroupRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete routing skillgroup request timeout response has a 5xx status code
func (o *DeleteRoutingSkillgroupRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this delete routing skillgroup request timeout response a status code equal to that given
func (o *DeleteRoutingSkillgroupRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *DeleteRoutingSkillgroupRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/skillgroups/{skillGroupId}][%d] deleteRoutingSkillgroupRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteRoutingSkillgroupRequestTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/skillgroups/{skillGroupId}][%d] deleteRoutingSkillgroupRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteRoutingSkillgroupRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingSkillgroupRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingSkillgroupRequestEntityTooLarge creates a DeleteRoutingSkillgroupRequestEntityTooLarge with default headers values
func NewDeleteRoutingSkillgroupRequestEntityTooLarge() *DeleteRoutingSkillgroupRequestEntityTooLarge {
	return &DeleteRoutingSkillgroupRequestEntityTooLarge{}
}

/*
DeleteRoutingSkillgroupRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type DeleteRoutingSkillgroupRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete routing skillgroup request entity too large response has a 2xx status code
func (o *DeleteRoutingSkillgroupRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete routing skillgroup request entity too large response has a 3xx status code
func (o *DeleteRoutingSkillgroupRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete routing skillgroup request entity too large response has a 4xx status code
func (o *DeleteRoutingSkillgroupRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete routing skillgroup request entity too large response has a 5xx status code
func (o *DeleteRoutingSkillgroupRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this delete routing skillgroup request entity too large response a status code equal to that given
func (o *DeleteRoutingSkillgroupRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *DeleteRoutingSkillgroupRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/skillgroups/{skillGroupId}][%d] deleteRoutingSkillgroupRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteRoutingSkillgroupRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/skillgroups/{skillGroupId}][%d] deleteRoutingSkillgroupRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteRoutingSkillgroupRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingSkillgroupRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingSkillgroupUnsupportedMediaType creates a DeleteRoutingSkillgroupUnsupportedMediaType with default headers values
func NewDeleteRoutingSkillgroupUnsupportedMediaType() *DeleteRoutingSkillgroupUnsupportedMediaType {
	return &DeleteRoutingSkillgroupUnsupportedMediaType{}
}

/*
DeleteRoutingSkillgroupUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteRoutingSkillgroupUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete routing skillgroup unsupported media type response has a 2xx status code
func (o *DeleteRoutingSkillgroupUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete routing skillgroup unsupported media type response has a 3xx status code
func (o *DeleteRoutingSkillgroupUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete routing skillgroup unsupported media type response has a 4xx status code
func (o *DeleteRoutingSkillgroupUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete routing skillgroup unsupported media type response has a 5xx status code
func (o *DeleteRoutingSkillgroupUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this delete routing skillgroup unsupported media type response a status code equal to that given
func (o *DeleteRoutingSkillgroupUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *DeleteRoutingSkillgroupUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/skillgroups/{skillGroupId}][%d] deleteRoutingSkillgroupUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteRoutingSkillgroupUnsupportedMediaType) String() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/skillgroups/{skillGroupId}][%d] deleteRoutingSkillgroupUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteRoutingSkillgroupUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingSkillgroupUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingSkillgroupTooManyRequests creates a DeleteRoutingSkillgroupTooManyRequests with default headers values
func NewDeleteRoutingSkillgroupTooManyRequests() *DeleteRoutingSkillgroupTooManyRequests {
	return &DeleteRoutingSkillgroupTooManyRequests{}
}

/*
DeleteRoutingSkillgroupTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type DeleteRoutingSkillgroupTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete routing skillgroup too many requests response has a 2xx status code
func (o *DeleteRoutingSkillgroupTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete routing skillgroup too many requests response has a 3xx status code
func (o *DeleteRoutingSkillgroupTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete routing skillgroup too many requests response has a 4xx status code
func (o *DeleteRoutingSkillgroupTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete routing skillgroup too many requests response has a 5xx status code
func (o *DeleteRoutingSkillgroupTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this delete routing skillgroup too many requests response a status code equal to that given
func (o *DeleteRoutingSkillgroupTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *DeleteRoutingSkillgroupTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/skillgroups/{skillGroupId}][%d] deleteRoutingSkillgroupTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteRoutingSkillgroupTooManyRequests) String() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/skillgroups/{skillGroupId}][%d] deleteRoutingSkillgroupTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteRoutingSkillgroupTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingSkillgroupTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingSkillgroupInternalServerError creates a DeleteRoutingSkillgroupInternalServerError with default headers values
func NewDeleteRoutingSkillgroupInternalServerError() *DeleteRoutingSkillgroupInternalServerError {
	return &DeleteRoutingSkillgroupInternalServerError{}
}

/*
DeleteRoutingSkillgroupInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteRoutingSkillgroupInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete routing skillgroup internal server error response has a 2xx status code
func (o *DeleteRoutingSkillgroupInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete routing skillgroup internal server error response has a 3xx status code
func (o *DeleteRoutingSkillgroupInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete routing skillgroup internal server error response has a 4xx status code
func (o *DeleteRoutingSkillgroupInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete routing skillgroup internal server error response has a 5xx status code
func (o *DeleteRoutingSkillgroupInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete routing skillgroup internal server error response a status code equal to that given
func (o *DeleteRoutingSkillgroupInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DeleteRoutingSkillgroupInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/skillgroups/{skillGroupId}][%d] deleteRoutingSkillgroupInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteRoutingSkillgroupInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/skillgroups/{skillGroupId}][%d] deleteRoutingSkillgroupInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteRoutingSkillgroupInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingSkillgroupInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingSkillgroupServiceUnavailable creates a DeleteRoutingSkillgroupServiceUnavailable with default headers values
func NewDeleteRoutingSkillgroupServiceUnavailable() *DeleteRoutingSkillgroupServiceUnavailable {
	return &DeleteRoutingSkillgroupServiceUnavailable{}
}

/*
DeleteRoutingSkillgroupServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteRoutingSkillgroupServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete routing skillgroup service unavailable response has a 2xx status code
func (o *DeleteRoutingSkillgroupServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete routing skillgroup service unavailable response has a 3xx status code
func (o *DeleteRoutingSkillgroupServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete routing skillgroup service unavailable response has a 4xx status code
func (o *DeleteRoutingSkillgroupServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete routing skillgroup service unavailable response has a 5xx status code
func (o *DeleteRoutingSkillgroupServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this delete routing skillgroup service unavailable response a status code equal to that given
func (o *DeleteRoutingSkillgroupServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *DeleteRoutingSkillgroupServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/skillgroups/{skillGroupId}][%d] deleteRoutingSkillgroupServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteRoutingSkillgroupServiceUnavailable) String() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/skillgroups/{skillGroupId}][%d] deleteRoutingSkillgroupServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteRoutingSkillgroupServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingSkillgroupServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingSkillgroupGatewayTimeout creates a DeleteRoutingSkillgroupGatewayTimeout with default headers values
func NewDeleteRoutingSkillgroupGatewayTimeout() *DeleteRoutingSkillgroupGatewayTimeout {
	return &DeleteRoutingSkillgroupGatewayTimeout{}
}

/*
DeleteRoutingSkillgroupGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type DeleteRoutingSkillgroupGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete routing skillgroup gateway timeout response has a 2xx status code
func (o *DeleteRoutingSkillgroupGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete routing skillgroup gateway timeout response has a 3xx status code
func (o *DeleteRoutingSkillgroupGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete routing skillgroup gateway timeout response has a 4xx status code
func (o *DeleteRoutingSkillgroupGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete routing skillgroup gateway timeout response has a 5xx status code
func (o *DeleteRoutingSkillgroupGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this delete routing skillgroup gateway timeout response a status code equal to that given
func (o *DeleteRoutingSkillgroupGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *DeleteRoutingSkillgroupGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/skillgroups/{skillGroupId}][%d] deleteRoutingSkillgroupGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteRoutingSkillgroupGatewayTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/skillgroups/{skillGroupId}][%d] deleteRoutingSkillgroupGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteRoutingSkillgroupGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingSkillgroupGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

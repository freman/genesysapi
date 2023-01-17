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

// DeleteRoutingSkillReader is a Reader for the DeleteRoutingSkill structure.
type DeleteRoutingSkillReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRoutingSkillReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteRoutingSkillOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteRoutingSkillBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteRoutingSkillUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteRoutingSkillForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteRoutingSkillNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewDeleteRoutingSkillRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteRoutingSkillRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteRoutingSkillUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteRoutingSkillTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteRoutingSkillInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteRoutingSkillServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteRoutingSkillGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteRoutingSkillOK creates a DeleteRoutingSkillOK with default headers values
func NewDeleteRoutingSkillOK() *DeleteRoutingSkillOK {
	return &DeleteRoutingSkillOK{}
}

/*
DeleteRoutingSkillOK describes a response with status code 200, with default header values.

Operation was successful.
*/
type DeleteRoutingSkillOK struct {
}

// IsSuccess returns true when this delete routing skill o k response has a 2xx status code
func (o *DeleteRoutingSkillOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete routing skill o k response has a 3xx status code
func (o *DeleteRoutingSkillOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete routing skill o k response has a 4xx status code
func (o *DeleteRoutingSkillOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete routing skill o k response has a 5xx status code
func (o *DeleteRoutingSkillOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete routing skill o k response a status code equal to that given
func (o *DeleteRoutingSkillOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteRoutingSkillOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/skills/{skillId}][%d] deleteRoutingSkillOK ", 200)
}

func (o *DeleteRoutingSkillOK) String() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/skills/{skillId}][%d] deleteRoutingSkillOK ", 200)
}

func (o *DeleteRoutingSkillOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteRoutingSkillBadRequest creates a DeleteRoutingSkillBadRequest with default headers values
func NewDeleteRoutingSkillBadRequest() *DeleteRoutingSkillBadRequest {
	return &DeleteRoutingSkillBadRequest{}
}

/*
DeleteRoutingSkillBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteRoutingSkillBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete routing skill bad request response has a 2xx status code
func (o *DeleteRoutingSkillBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete routing skill bad request response has a 3xx status code
func (o *DeleteRoutingSkillBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete routing skill bad request response has a 4xx status code
func (o *DeleteRoutingSkillBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete routing skill bad request response has a 5xx status code
func (o *DeleteRoutingSkillBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete routing skill bad request response a status code equal to that given
func (o *DeleteRoutingSkillBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DeleteRoutingSkillBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/skills/{skillId}][%d] deleteRoutingSkillBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteRoutingSkillBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/skills/{skillId}][%d] deleteRoutingSkillBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteRoutingSkillBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingSkillBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingSkillUnauthorized creates a DeleteRoutingSkillUnauthorized with default headers values
func NewDeleteRoutingSkillUnauthorized() *DeleteRoutingSkillUnauthorized {
	return &DeleteRoutingSkillUnauthorized{}
}

/*
DeleteRoutingSkillUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteRoutingSkillUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete routing skill unauthorized response has a 2xx status code
func (o *DeleteRoutingSkillUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete routing skill unauthorized response has a 3xx status code
func (o *DeleteRoutingSkillUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete routing skill unauthorized response has a 4xx status code
func (o *DeleteRoutingSkillUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete routing skill unauthorized response has a 5xx status code
func (o *DeleteRoutingSkillUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete routing skill unauthorized response a status code equal to that given
func (o *DeleteRoutingSkillUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteRoutingSkillUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/skills/{skillId}][%d] deleteRoutingSkillUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteRoutingSkillUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/skills/{skillId}][%d] deleteRoutingSkillUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteRoutingSkillUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingSkillUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingSkillForbidden creates a DeleteRoutingSkillForbidden with default headers values
func NewDeleteRoutingSkillForbidden() *DeleteRoutingSkillForbidden {
	return &DeleteRoutingSkillForbidden{}
}

/*
DeleteRoutingSkillForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type DeleteRoutingSkillForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete routing skill forbidden response has a 2xx status code
func (o *DeleteRoutingSkillForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete routing skill forbidden response has a 3xx status code
func (o *DeleteRoutingSkillForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete routing skill forbidden response has a 4xx status code
func (o *DeleteRoutingSkillForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete routing skill forbidden response has a 5xx status code
func (o *DeleteRoutingSkillForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete routing skill forbidden response a status code equal to that given
func (o *DeleteRoutingSkillForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteRoutingSkillForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/skills/{skillId}][%d] deleteRoutingSkillForbidden  %+v", 403, o.Payload)
}

func (o *DeleteRoutingSkillForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/skills/{skillId}][%d] deleteRoutingSkillForbidden  %+v", 403, o.Payload)
}

func (o *DeleteRoutingSkillForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingSkillForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingSkillNotFound creates a DeleteRoutingSkillNotFound with default headers values
func NewDeleteRoutingSkillNotFound() *DeleteRoutingSkillNotFound {
	return &DeleteRoutingSkillNotFound{}
}

/*
DeleteRoutingSkillNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type DeleteRoutingSkillNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete routing skill not found response has a 2xx status code
func (o *DeleteRoutingSkillNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete routing skill not found response has a 3xx status code
func (o *DeleteRoutingSkillNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete routing skill not found response has a 4xx status code
func (o *DeleteRoutingSkillNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete routing skill not found response has a 5xx status code
func (o *DeleteRoutingSkillNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete routing skill not found response a status code equal to that given
func (o *DeleteRoutingSkillNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteRoutingSkillNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/skills/{skillId}][%d] deleteRoutingSkillNotFound  %+v", 404, o.Payload)
}

func (o *DeleteRoutingSkillNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/skills/{skillId}][%d] deleteRoutingSkillNotFound  %+v", 404, o.Payload)
}

func (o *DeleteRoutingSkillNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingSkillNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingSkillRequestTimeout creates a DeleteRoutingSkillRequestTimeout with default headers values
func NewDeleteRoutingSkillRequestTimeout() *DeleteRoutingSkillRequestTimeout {
	return &DeleteRoutingSkillRequestTimeout{}
}

/*
DeleteRoutingSkillRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type DeleteRoutingSkillRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete routing skill request timeout response has a 2xx status code
func (o *DeleteRoutingSkillRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete routing skill request timeout response has a 3xx status code
func (o *DeleteRoutingSkillRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete routing skill request timeout response has a 4xx status code
func (o *DeleteRoutingSkillRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete routing skill request timeout response has a 5xx status code
func (o *DeleteRoutingSkillRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this delete routing skill request timeout response a status code equal to that given
func (o *DeleteRoutingSkillRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *DeleteRoutingSkillRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/skills/{skillId}][%d] deleteRoutingSkillRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteRoutingSkillRequestTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/skills/{skillId}][%d] deleteRoutingSkillRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteRoutingSkillRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingSkillRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingSkillRequestEntityTooLarge creates a DeleteRoutingSkillRequestEntityTooLarge with default headers values
func NewDeleteRoutingSkillRequestEntityTooLarge() *DeleteRoutingSkillRequestEntityTooLarge {
	return &DeleteRoutingSkillRequestEntityTooLarge{}
}

/*
DeleteRoutingSkillRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type DeleteRoutingSkillRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete routing skill request entity too large response has a 2xx status code
func (o *DeleteRoutingSkillRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete routing skill request entity too large response has a 3xx status code
func (o *DeleteRoutingSkillRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete routing skill request entity too large response has a 4xx status code
func (o *DeleteRoutingSkillRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete routing skill request entity too large response has a 5xx status code
func (o *DeleteRoutingSkillRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this delete routing skill request entity too large response a status code equal to that given
func (o *DeleteRoutingSkillRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *DeleteRoutingSkillRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/skills/{skillId}][%d] deleteRoutingSkillRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteRoutingSkillRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/skills/{skillId}][%d] deleteRoutingSkillRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteRoutingSkillRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingSkillRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingSkillUnsupportedMediaType creates a DeleteRoutingSkillUnsupportedMediaType with default headers values
func NewDeleteRoutingSkillUnsupportedMediaType() *DeleteRoutingSkillUnsupportedMediaType {
	return &DeleteRoutingSkillUnsupportedMediaType{}
}

/*
DeleteRoutingSkillUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteRoutingSkillUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete routing skill unsupported media type response has a 2xx status code
func (o *DeleteRoutingSkillUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete routing skill unsupported media type response has a 3xx status code
func (o *DeleteRoutingSkillUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete routing skill unsupported media type response has a 4xx status code
func (o *DeleteRoutingSkillUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete routing skill unsupported media type response has a 5xx status code
func (o *DeleteRoutingSkillUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this delete routing skill unsupported media type response a status code equal to that given
func (o *DeleteRoutingSkillUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *DeleteRoutingSkillUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/skills/{skillId}][%d] deleteRoutingSkillUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteRoutingSkillUnsupportedMediaType) String() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/skills/{skillId}][%d] deleteRoutingSkillUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteRoutingSkillUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingSkillUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingSkillTooManyRequests creates a DeleteRoutingSkillTooManyRequests with default headers values
func NewDeleteRoutingSkillTooManyRequests() *DeleteRoutingSkillTooManyRequests {
	return &DeleteRoutingSkillTooManyRequests{}
}

/*
DeleteRoutingSkillTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type DeleteRoutingSkillTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete routing skill too many requests response has a 2xx status code
func (o *DeleteRoutingSkillTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete routing skill too many requests response has a 3xx status code
func (o *DeleteRoutingSkillTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete routing skill too many requests response has a 4xx status code
func (o *DeleteRoutingSkillTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete routing skill too many requests response has a 5xx status code
func (o *DeleteRoutingSkillTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this delete routing skill too many requests response a status code equal to that given
func (o *DeleteRoutingSkillTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *DeleteRoutingSkillTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/skills/{skillId}][%d] deleteRoutingSkillTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteRoutingSkillTooManyRequests) String() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/skills/{skillId}][%d] deleteRoutingSkillTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteRoutingSkillTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingSkillTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingSkillInternalServerError creates a DeleteRoutingSkillInternalServerError with default headers values
func NewDeleteRoutingSkillInternalServerError() *DeleteRoutingSkillInternalServerError {
	return &DeleteRoutingSkillInternalServerError{}
}

/*
DeleteRoutingSkillInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteRoutingSkillInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete routing skill internal server error response has a 2xx status code
func (o *DeleteRoutingSkillInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete routing skill internal server error response has a 3xx status code
func (o *DeleteRoutingSkillInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete routing skill internal server error response has a 4xx status code
func (o *DeleteRoutingSkillInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete routing skill internal server error response has a 5xx status code
func (o *DeleteRoutingSkillInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete routing skill internal server error response a status code equal to that given
func (o *DeleteRoutingSkillInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DeleteRoutingSkillInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/skills/{skillId}][%d] deleteRoutingSkillInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteRoutingSkillInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/skills/{skillId}][%d] deleteRoutingSkillInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteRoutingSkillInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingSkillInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingSkillServiceUnavailable creates a DeleteRoutingSkillServiceUnavailable with default headers values
func NewDeleteRoutingSkillServiceUnavailable() *DeleteRoutingSkillServiceUnavailable {
	return &DeleteRoutingSkillServiceUnavailable{}
}

/*
DeleteRoutingSkillServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteRoutingSkillServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete routing skill service unavailable response has a 2xx status code
func (o *DeleteRoutingSkillServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete routing skill service unavailable response has a 3xx status code
func (o *DeleteRoutingSkillServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete routing skill service unavailable response has a 4xx status code
func (o *DeleteRoutingSkillServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete routing skill service unavailable response has a 5xx status code
func (o *DeleteRoutingSkillServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this delete routing skill service unavailable response a status code equal to that given
func (o *DeleteRoutingSkillServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *DeleteRoutingSkillServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/skills/{skillId}][%d] deleteRoutingSkillServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteRoutingSkillServiceUnavailable) String() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/skills/{skillId}][%d] deleteRoutingSkillServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteRoutingSkillServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingSkillServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingSkillGatewayTimeout creates a DeleteRoutingSkillGatewayTimeout with default headers values
func NewDeleteRoutingSkillGatewayTimeout() *DeleteRoutingSkillGatewayTimeout {
	return &DeleteRoutingSkillGatewayTimeout{}
}

/*
DeleteRoutingSkillGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type DeleteRoutingSkillGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete routing skill gateway timeout response has a 2xx status code
func (o *DeleteRoutingSkillGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete routing skill gateway timeout response has a 3xx status code
func (o *DeleteRoutingSkillGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete routing skill gateway timeout response has a 4xx status code
func (o *DeleteRoutingSkillGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete routing skill gateway timeout response has a 5xx status code
func (o *DeleteRoutingSkillGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this delete routing skill gateway timeout response a status code equal to that given
func (o *DeleteRoutingSkillGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *DeleteRoutingSkillGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/skills/{skillId}][%d] deleteRoutingSkillGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteRoutingSkillGatewayTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/skills/{skillId}][%d] deleteRoutingSkillGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteRoutingSkillGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingSkillGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

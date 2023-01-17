// Code generated by go-swagger; DO NOT EDIT.

package process_automation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// DeleteProcessautomationTriggerReader is a Reader for the DeleteProcessautomationTrigger structure.
type DeleteProcessautomationTriggerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteProcessautomationTriggerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteProcessautomationTriggerNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteProcessautomationTriggerBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteProcessautomationTriggerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteProcessautomationTriggerForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteProcessautomationTriggerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewDeleteProcessautomationTriggerRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteProcessautomationTriggerRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteProcessautomationTriggerUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteProcessautomationTriggerTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteProcessautomationTriggerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteProcessautomationTriggerServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteProcessautomationTriggerGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteProcessautomationTriggerNoContent creates a DeleteProcessautomationTriggerNoContent with default headers values
func NewDeleteProcessautomationTriggerNoContent() *DeleteProcessautomationTriggerNoContent {
	return &DeleteProcessautomationTriggerNoContent{}
}

/*
DeleteProcessautomationTriggerNoContent describes a response with status code 204, with default header values.

Delete was successful
*/
type DeleteProcessautomationTriggerNoContent struct {
}

// IsSuccess returns true when this delete processautomation trigger no content response has a 2xx status code
func (o *DeleteProcessautomationTriggerNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete processautomation trigger no content response has a 3xx status code
func (o *DeleteProcessautomationTriggerNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete processautomation trigger no content response has a 4xx status code
func (o *DeleteProcessautomationTriggerNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete processautomation trigger no content response has a 5xx status code
func (o *DeleteProcessautomationTriggerNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete processautomation trigger no content response a status code equal to that given
func (o *DeleteProcessautomationTriggerNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DeleteProcessautomationTriggerNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/processautomation/triggers/{triggerId}][%d] deleteProcessautomationTriggerNoContent ", 204)
}

func (o *DeleteProcessautomationTriggerNoContent) String() string {
	return fmt.Sprintf("[DELETE /api/v2/processautomation/triggers/{triggerId}][%d] deleteProcessautomationTriggerNoContent ", 204)
}

func (o *DeleteProcessautomationTriggerNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteProcessautomationTriggerBadRequest creates a DeleteProcessautomationTriggerBadRequest with default headers values
func NewDeleteProcessautomationTriggerBadRequest() *DeleteProcessautomationTriggerBadRequest {
	return &DeleteProcessautomationTriggerBadRequest{}
}

/*
DeleteProcessautomationTriggerBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteProcessautomationTriggerBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete processautomation trigger bad request response has a 2xx status code
func (o *DeleteProcessautomationTriggerBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete processautomation trigger bad request response has a 3xx status code
func (o *DeleteProcessautomationTriggerBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete processautomation trigger bad request response has a 4xx status code
func (o *DeleteProcessautomationTriggerBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete processautomation trigger bad request response has a 5xx status code
func (o *DeleteProcessautomationTriggerBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete processautomation trigger bad request response a status code equal to that given
func (o *DeleteProcessautomationTriggerBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DeleteProcessautomationTriggerBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/processautomation/triggers/{triggerId}][%d] deleteProcessautomationTriggerBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteProcessautomationTriggerBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api/v2/processautomation/triggers/{triggerId}][%d] deleteProcessautomationTriggerBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteProcessautomationTriggerBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteProcessautomationTriggerBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteProcessautomationTriggerUnauthorized creates a DeleteProcessautomationTriggerUnauthorized with default headers values
func NewDeleteProcessautomationTriggerUnauthorized() *DeleteProcessautomationTriggerUnauthorized {
	return &DeleteProcessautomationTriggerUnauthorized{}
}

/*
DeleteProcessautomationTriggerUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteProcessautomationTriggerUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete processautomation trigger unauthorized response has a 2xx status code
func (o *DeleteProcessautomationTriggerUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete processautomation trigger unauthorized response has a 3xx status code
func (o *DeleteProcessautomationTriggerUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete processautomation trigger unauthorized response has a 4xx status code
func (o *DeleteProcessautomationTriggerUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete processautomation trigger unauthorized response has a 5xx status code
func (o *DeleteProcessautomationTriggerUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete processautomation trigger unauthorized response a status code equal to that given
func (o *DeleteProcessautomationTriggerUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteProcessautomationTriggerUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/processautomation/triggers/{triggerId}][%d] deleteProcessautomationTriggerUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteProcessautomationTriggerUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/v2/processautomation/triggers/{triggerId}][%d] deleteProcessautomationTriggerUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteProcessautomationTriggerUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteProcessautomationTriggerUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteProcessautomationTriggerForbidden creates a DeleteProcessautomationTriggerForbidden with default headers values
func NewDeleteProcessautomationTriggerForbidden() *DeleteProcessautomationTriggerForbidden {
	return &DeleteProcessautomationTriggerForbidden{}
}

/*
DeleteProcessautomationTriggerForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type DeleteProcessautomationTriggerForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete processautomation trigger forbidden response has a 2xx status code
func (o *DeleteProcessautomationTriggerForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete processautomation trigger forbidden response has a 3xx status code
func (o *DeleteProcessautomationTriggerForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete processautomation trigger forbidden response has a 4xx status code
func (o *DeleteProcessautomationTriggerForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete processautomation trigger forbidden response has a 5xx status code
func (o *DeleteProcessautomationTriggerForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete processautomation trigger forbidden response a status code equal to that given
func (o *DeleteProcessautomationTriggerForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteProcessautomationTriggerForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/processautomation/triggers/{triggerId}][%d] deleteProcessautomationTriggerForbidden  %+v", 403, o.Payload)
}

func (o *DeleteProcessautomationTriggerForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v2/processautomation/triggers/{triggerId}][%d] deleteProcessautomationTriggerForbidden  %+v", 403, o.Payload)
}

func (o *DeleteProcessautomationTriggerForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteProcessautomationTriggerForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteProcessautomationTriggerNotFound creates a DeleteProcessautomationTriggerNotFound with default headers values
func NewDeleteProcessautomationTriggerNotFound() *DeleteProcessautomationTriggerNotFound {
	return &DeleteProcessautomationTriggerNotFound{}
}

/*
DeleteProcessautomationTriggerNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type DeleteProcessautomationTriggerNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete processautomation trigger not found response has a 2xx status code
func (o *DeleteProcessautomationTriggerNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete processautomation trigger not found response has a 3xx status code
func (o *DeleteProcessautomationTriggerNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete processautomation trigger not found response has a 4xx status code
func (o *DeleteProcessautomationTriggerNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete processautomation trigger not found response has a 5xx status code
func (o *DeleteProcessautomationTriggerNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete processautomation trigger not found response a status code equal to that given
func (o *DeleteProcessautomationTriggerNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteProcessautomationTriggerNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/processautomation/triggers/{triggerId}][%d] deleteProcessautomationTriggerNotFound  %+v", 404, o.Payload)
}

func (o *DeleteProcessautomationTriggerNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/v2/processautomation/triggers/{triggerId}][%d] deleteProcessautomationTriggerNotFound  %+v", 404, o.Payload)
}

func (o *DeleteProcessautomationTriggerNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteProcessautomationTriggerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteProcessautomationTriggerRequestTimeout creates a DeleteProcessautomationTriggerRequestTimeout with default headers values
func NewDeleteProcessautomationTriggerRequestTimeout() *DeleteProcessautomationTriggerRequestTimeout {
	return &DeleteProcessautomationTriggerRequestTimeout{}
}

/*
DeleteProcessautomationTriggerRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type DeleteProcessautomationTriggerRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete processautomation trigger request timeout response has a 2xx status code
func (o *DeleteProcessautomationTriggerRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete processautomation trigger request timeout response has a 3xx status code
func (o *DeleteProcessautomationTriggerRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete processautomation trigger request timeout response has a 4xx status code
func (o *DeleteProcessautomationTriggerRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete processautomation trigger request timeout response has a 5xx status code
func (o *DeleteProcessautomationTriggerRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this delete processautomation trigger request timeout response a status code equal to that given
func (o *DeleteProcessautomationTriggerRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *DeleteProcessautomationTriggerRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/processautomation/triggers/{triggerId}][%d] deleteProcessautomationTriggerRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteProcessautomationTriggerRequestTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/processautomation/triggers/{triggerId}][%d] deleteProcessautomationTriggerRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteProcessautomationTriggerRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteProcessautomationTriggerRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteProcessautomationTriggerRequestEntityTooLarge creates a DeleteProcessautomationTriggerRequestEntityTooLarge with default headers values
func NewDeleteProcessautomationTriggerRequestEntityTooLarge() *DeleteProcessautomationTriggerRequestEntityTooLarge {
	return &DeleteProcessautomationTriggerRequestEntityTooLarge{}
}

/*
DeleteProcessautomationTriggerRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type DeleteProcessautomationTriggerRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete processautomation trigger request entity too large response has a 2xx status code
func (o *DeleteProcessautomationTriggerRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete processautomation trigger request entity too large response has a 3xx status code
func (o *DeleteProcessautomationTriggerRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete processautomation trigger request entity too large response has a 4xx status code
func (o *DeleteProcessautomationTriggerRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete processautomation trigger request entity too large response has a 5xx status code
func (o *DeleteProcessautomationTriggerRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this delete processautomation trigger request entity too large response a status code equal to that given
func (o *DeleteProcessautomationTriggerRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *DeleteProcessautomationTriggerRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/processautomation/triggers/{triggerId}][%d] deleteProcessautomationTriggerRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteProcessautomationTriggerRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[DELETE /api/v2/processautomation/triggers/{triggerId}][%d] deleteProcessautomationTriggerRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteProcessautomationTriggerRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteProcessautomationTriggerRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteProcessautomationTriggerUnsupportedMediaType creates a DeleteProcessautomationTriggerUnsupportedMediaType with default headers values
func NewDeleteProcessautomationTriggerUnsupportedMediaType() *DeleteProcessautomationTriggerUnsupportedMediaType {
	return &DeleteProcessautomationTriggerUnsupportedMediaType{}
}

/*
DeleteProcessautomationTriggerUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteProcessautomationTriggerUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete processautomation trigger unsupported media type response has a 2xx status code
func (o *DeleteProcessautomationTriggerUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete processautomation trigger unsupported media type response has a 3xx status code
func (o *DeleteProcessautomationTriggerUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete processautomation trigger unsupported media type response has a 4xx status code
func (o *DeleteProcessautomationTriggerUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete processautomation trigger unsupported media type response has a 5xx status code
func (o *DeleteProcessautomationTriggerUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this delete processautomation trigger unsupported media type response a status code equal to that given
func (o *DeleteProcessautomationTriggerUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *DeleteProcessautomationTriggerUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/processautomation/triggers/{triggerId}][%d] deleteProcessautomationTriggerUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteProcessautomationTriggerUnsupportedMediaType) String() string {
	return fmt.Sprintf("[DELETE /api/v2/processautomation/triggers/{triggerId}][%d] deleteProcessautomationTriggerUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteProcessautomationTriggerUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteProcessautomationTriggerUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteProcessautomationTriggerTooManyRequests creates a DeleteProcessautomationTriggerTooManyRequests with default headers values
func NewDeleteProcessautomationTriggerTooManyRequests() *DeleteProcessautomationTriggerTooManyRequests {
	return &DeleteProcessautomationTriggerTooManyRequests{}
}

/*
DeleteProcessautomationTriggerTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type DeleteProcessautomationTriggerTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete processautomation trigger too many requests response has a 2xx status code
func (o *DeleteProcessautomationTriggerTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete processautomation trigger too many requests response has a 3xx status code
func (o *DeleteProcessautomationTriggerTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete processautomation trigger too many requests response has a 4xx status code
func (o *DeleteProcessautomationTriggerTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete processautomation trigger too many requests response has a 5xx status code
func (o *DeleteProcessautomationTriggerTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this delete processautomation trigger too many requests response a status code equal to that given
func (o *DeleteProcessautomationTriggerTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *DeleteProcessautomationTriggerTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/processautomation/triggers/{triggerId}][%d] deleteProcessautomationTriggerTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteProcessautomationTriggerTooManyRequests) String() string {
	return fmt.Sprintf("[DELETE /api/v2/processautomation/triggers/{triggerId}][%d] deleteProcessautomationTriggerTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteProcessautomationTriggerTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteProcessautomationTriggerTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteProcessautomationTriggerInternalServerError creates a DeleteProcessautomationTriggerInternalServerError with default headers values
func NewDeleteProcessautomationTriggerInternalServerError() *DeleteProcessautomationTriggerInternalServerError {
	return &DeleteProcessautomationTriggerInternalServerError{}
}

/*
DeleteProcessautomationTriggerInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteProcessautomationTriggerInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete processautomation trigger internal server error response has a 2xx status code
func (o *DeleteProcessautomationTriggerInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete processautomation trigger internal server error response has a 3xx status code
func (o *DeleteProcessautomationTriggerInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete processautomation trigger internal server error response has a 4xx status code
func (o *DeleteProcessautomationTriggerInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete processautomation trigger internal server error response has a 5xx status code
func (o *DeleteProcessautomationTriggerInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete processautomation trigger internal server error response a status code equal to that given
func (o *DeleteProcessautomationTriggerInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DeleteProcessautomationTriggerInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/processautomation/triggers/{triggerId}][%d] deleteProcessautomationTriggerInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteProcessautomationTriggerInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /api/v2/processautomation/triggers/{triggerId}][%d] deleteProcessautomationTriggerInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteProcessautomationTriggerInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteProcessautomationTriggerInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteProcessautomationTriggerServiceUnavailable creates a DeleteProcessautomationTriggerServiceUnavailable with default headers values
func NewDeleteProcessautomationTriggerServiceUnavailable() *DeleteProcessautomationTriggerServiceUnavailable {
	return &DeleteProcessautomationTriggerServiceUnavailable{}
}

/*
DeleteProcessautomationTriggerServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteProcessautomationTriggerServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete processautomation trigger service unavailable response has a 2xx status code
func (o *DeleteProcessautomationTriggerServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete processautomation trigger service unavailable response has a 3xx status code
func (o *DeleteProcessautomationTriggerServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete processautomation trigger service unavailable response has a 4xx status code
func (o *DeleteProcessautomationTriggerServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete processautomation trigger service unavailable response has a 5xx status code
func (o *DeleteProcessautomationTriggerServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this delete processautomation trigger service unavailable response a status code equal to that given
func (o *DeleteProcessautomationTriggerServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *DeleteProcessautomationTriggerServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/processautomation/triggers/{triggerId}][%d] deleteProcessautomationTriggerServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteProcessautomationTriggerServiceUnavailable) String() string {
	return fmt.Sprintf("[DELETE /api/v2/processautomation/triggers/{triggerId}][%d] deleteProcessautomationTriggerServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteProcessautomationTriggerServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteProcessautomationTriggerServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteProcessautomationTriggerGatewayTimeout creates a DeleteProcessautomationTriggerGatewayTimeout with default headers values
func NewDeleteProcessautomationTriggerGatewayTimeout() *DeleteProcessautomationTriggerGatewayTimeout {
	return &DeleteProcessautomationTriggerGatewayTimeout{}
}

/*
DeleteProcessautomationTriggerGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type DeleteProcessautomationTriggerGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete processautomation trigger gateway timeout response has a 2xx status code
func (o *DeleteProcessautomationTriggerGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete processautomation trigger gateway timeout response has a 3xx status code
func (o *DeleteProcessautomationTriggerGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete processautomation trigger gateway timeout response has a 4xx status code
func (o *DeleteProcessautomationTriggerGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete processautomation trigger gateway timeout response has a 5xx status code
func (o *DeleteProcessautomationTriggerGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this delete processautomation trigger gateway timeout response a status code equal to that given
func (o *DeleteProcessautomationTriggerGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *DeleteProcessautomationTriggerGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/processautomation/triggers/{triggerId}][%d] deleteProcessautomationTriggerGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteProcessautomationTriggerGatewayTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/processautomation/triggers/{triggerId}][%d] deleteProcessautomationTriggerGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteProcessautomationTriggerGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteProcessautomationTriggerGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
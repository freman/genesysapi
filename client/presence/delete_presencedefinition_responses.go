// Code generated by go-swagger; DO NOT EDIT.

package presence

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// DeletePresencedefinitionReader is a Reader for the DeletePresencedefinition structure.
type DeletePresencedefinitionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeletePresencedefinitionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 400:
		result := NewDeletePresencedefinitionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeletePresencedefinitionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeletePresencedefinitionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeletePresencedefinitionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewDeletePresencedefinitionRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeletePresencedefinitionRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeletePresencedefinitionUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeletePresencedefinitionTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeletePresencedefinitionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeletePresencedefinitionServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeletePresencedefinitionGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeletePresencedefinitionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeletePresencedefinitionBadRequest creates a DeletePresencedefinitionBadRequest with default headers values
func NewDeletePresencedefinitionBadRequest() *DeletePresencedefinitionBadRequest {
	return &DeletePresencedefinitionBadRequest{}
}

/*
DeletePresencedefinitionBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeletePresencedefinitionBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete presencedefinition bad request response has a 2xx status code
func (o *DeletePresencedefinitionBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete presencedefinition bad request response has a 3xx status code
func (o *DeletePresencedefinitionBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete presencedefinition bad request response has a 4xx status code
func (o *DeletePresencedefinitionBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete presencedefinition bad request response has a 5xx status code
func (o *DeletePresencedefinitionBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete presencedefinition bad request response a status code equal to that given
func (o *DeletePresencedefinitionBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DeletePresencedefinitionBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/presencedefinitions/{presenceId}][%d] deletePresencedefinitionBadRequest  %+v", 400, o.Payload)
}

func (o *DeletePresencedefinitionBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api/v2/presencedefinitions/{presenceId}][%d] deletePresencedefinitionBadRequest  %+v", 400, o.Payload)
}

func (o *DeletePresencedefinitionBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeletePresencedefinitionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePresencedefinitionUnauthorized creates a DeletePresencedefinitionUnauthorized with default headers values
func NewDeletePresencedefinitionUnauthorized() *DeletePresencedefinitionUnauthorized {
	return &DeletePresencedefinitionUnauthorized{}
}

/*
DeletePresencedefinitionUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type DeletePresencedefinitionUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete presencedefinition unauthorized response has a 2xx status code
func (o *DeletePresencedefinitionUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete presencedefinition unauthorized response has a 3xx status code
func (o *DeletePresencedefinitionUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete presencedefinition unauthorized response has a 4xx status code
func (o *DeletePresencedefinitionUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete presencedefinition unauthorized response has a 5xx status code
func (o *DeletePresencedefinitionUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete presencedefinition unauthorized response a status code equal to that given
func (o *DeletePresencedefinitionUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeletePresencedefinitionUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/presencedefinitions/{presenceId}][%d] deletePresencedefinitionUnauthorized  %+v", 401, o.Payload)
}

func (o *DeletePresencedefinitionUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/v2/presencedefinitions/{presenceId}][%d] deletePresencedefinitionUnauthorized  %+v", 401, o.Payload)
}

func (o *DeletePresencedefinitionUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeletePresencedefinitionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePresencedefinitionForbidden creates a DeletePresencedefinitionForbidden with default headers values
func NewDeletePresencedefinitionForbidden() *DeletePresencedefinitionForbidden {
	return &DeletePresencedefinitionForbidden{}
}

/*
DeletePresencedefinitionForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type DeletePresencedefinitionForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete presencedefinition forbidden response has a 2xx status code
func (o *DeletePresencedefinitionForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete presencedefinition forbidden response has a 3xx status code
func (o *DeletePresencedefinitionForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete presencedefinition forbidden response has a 4xx status code
func (o *DeletePresencedefinitionForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete presencedefinition forbidden response has a 5xx status code
func (o *DeletePresencedefinitionForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete presencedefinition forbidden response a status code equal to that given
func (o *DeletePresencedefinitionForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeletePresencedefinitionForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/presencedefinitions/{presenceId}][%d] deletePresencedefinitionForbidden  %+v", 403, o.Payload)
}

func (o *DeletePresencedefinitionForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v2/presencedefinitions/{presenceId}][%d] deletePresencedefinitionForbidden  %+v", 403, o.Payload)
}

func (o *DeletePresencedefinitionForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeletePresencedefinitionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePresencedefinitionNotFound creates a DeletePresencedefinitionNotFound with default headers values
func NewDeletePresencedefinitionNotFound() *DeletePresencedefinitionNotFound {
	return &DeletePresencedefinitionNotFound{}
}

/*
DeletePresencedefinitionNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type DeletePresencedefinitionNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete presencedefinition not found response has a 2xx status code
func (o *DeletePresencedefinitionNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete presencedefinition not found response has a 3xx status code
func (o *DeletePresencedefinitionNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete presencedefinition not found response has a 4xx status code
func (o *DeletePresencedefinitionNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete presencedefinition not found response has a 5xx status code
func (o *DeletePresencedefinitionNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete presencedefinition not found response a status code equal to that given
func (o *DeletePresencedefinitionNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeletePresencedefinitionNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/presencedefinitions/{presenceId}][%d] deletePresencedefinitionNotFound  %+v", 404, o.Payload)
}

func (o *DeletePresencedefinitionNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/v2/presencedefinitions/{presenceId}][%d] deletePresencedefinitionNotFound  %+v", 404, o.Payload)
}

func (o *DeletePresencedefinitionNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeletePresencedefinitionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePresencedefinitionRequestTimeout creates a DeletePresencedefinitionRequestTimeout with default headers values
func NewDeletePresencedefinitionRequestTimeout() *DeletePresencedefinitionRequestTimeout {
	return &DeletePresencedefinitionRequestTimeout{}
}

/*
DeletePresencedefinitionRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type DeletePresencedefinitionRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete presencedefinition request timeout response has a 2xx status code
func (o *DeletePresencedefinitionRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete presencedefinition request timeout response has a 3xx status code
func (o *DeletePresencedefinitionRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete presencedefinition request timeout response has a 4xx status code
func (o *DeletePresencedefinitionRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete presencedefinition request timeout response has a 5xx status code
func (o *DeletePresencedefinitionRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this delete presencedefinition request timeout response a status code equal to that given
func (o *DeletePresencedefinitionRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *DeletePresencedefinitionRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/presencedefinitions/{presenceId}][%d] deletePresencedefinitionRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeletePresencedefinitionRequestTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/presencedefinitions/{presenceId}][%d] deletePresencedefinitionRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeletePresencedefinitionRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeletePresencedefinitionRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePresencedefinitionRequestEntityTooLarge creates a DeletePresencedefinitionRequestEntityTooLarge with default headers values
func NewDeletePresencedefinitionRequestEntityTooLarge() *DeletePresencedefinitionRequestEntityTooLarge {
	return &DeletePresencedefinitionRequestEntityTooLarge{}
}

/*
DeletePresencedefinitionRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type DeletePresencedefinitionRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete presencedefinition request entity too large response has a 2xx status code
func (o *DeletePresencedefinitionRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete presencedefinition request entity too large response has a 3xx status code
func (o *DeletePresencedefinitionRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete presencedefinition request entity too large response has a 4xx status code
func (o *DeletePresencedefinitionRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete presencedefinition request entity too large response has a 5xx status code
func (o *DeletePresencedefinitionRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this delete presencedefinition request entity too large response a status code equal to that given
func (o *DeletePresencedefinitionRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *DeletePresencedefinitionRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/presencedefinitions/{presenceId}][%d] deletePresencedefinitionRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeletePresencedefinitionRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[DELETE /api/v2/presencedefinitions/{presenceId}][%d] deletePresencedefinitionRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeletePresencedefinitionRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeletePresencedefinitionRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePresencedefinitionUnsupportedMediaType creates a DeletePresencedefinitionUnsupportedMediaType with default headers values
func NewDeletePresencedefinitionUnsupportedMediaType() *DeletePresencedefinitionUnsupportedMediaType {
	return &DeletePresencedefinitionUnsupportedMediaType{}
}

/*
DeletePresencedefinitionUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeletePresencedefinitionUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete presencedefinition unsupported media type response has a 2xx status code
func (o *DeletePresencedefinitionUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete presencedefinition unsupported media type response has a 3xx status code
func (o *DeletePresencedefinitionUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete presencedefinition unsupported media type response has a 4xx status code
func (o *DeletePresencedefinitionUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete presencedefinition unsupported media type response has a 5xx status code
func (o *DeletePresencedefinitionUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this delete presencedefinition unsupported media type response a status code equal to that given
func (o *DeletePresencedefinitionUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *DeletePresencedefinitionUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/presencedefinitions/{presenceId}][%d] deletePresencedefinitionUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeletePresencedefinitionUnsupportedMediaType) String() string {
	return fmt.Sprintf("[DELETE /api/v2/presencedefinitions/{presenceId}][%d] deletePresencedefinitionUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeletePresencedefinitionUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeletePresencedefinitionUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePresencedefinitionTooManyRequests creates a DeletePresencedefinitionTooManyRequests with default headers values
func NewDeletePresencedefinitionTooManyRequests() *DeletePresencedefinitionTooManyRequests {
	return &DeletePresencedefinitionTooManyRequests{}
}

/*
DeletePresencedefinitionTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type DeletePresencedefinitionTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete presencedefinition too many requests response has a 2xx status code
func (o *DeletePresencedefinitionTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete presencedefinition too many requests response has a 3xx status code
func (o *DeletePresencedefinitionTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete presencedefinition too many requests response has a 4xx status code
func (o *DeletePresencedefinitionTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete presencedefinition too many requests response has a 5xx status code
func (o *DeletePresencedefinitionTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this delete presencedefinition too many requests response a status code equal to that given
func (o *DeletePresencedefinitionTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *DeletePresencedefinitionTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/presencedefinitions/{presenceId}][%d] deletePresencedefinitionTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeletePresencedefinitionTooManyRequests) String() string {
	return fmt.Sprintf("[DELETE /api/v2/presencedefinitions/{presenceId}][%d] deletePresencedefinitionTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeletePresencedefinitionTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeletePresencedefinitionTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePresencedefinitionInternalServerError creates a DeletePresencedefinitionInternalServerError with default headers values
func NewDeletePresencedefinitionInternalServerError() *DeletePresencedefinitionInternalServerError {
	return &DeletePresencedefinitionInternalServerError{}
}

/*
DeletePresencedefinitionInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeletePresencedefinitionInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete presencedefinition internal server error response has a 2xx status code
func (o *DeletePresencedefinitionInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete presencedefinition internal server error response has a 3xx status code
func (o *DeletePresencedefinitionInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete presencedefinition internal server error response has a 4xx status code
func (o *DeletePresencedefinitionInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete presencedefinition internal server error response has a 5xx status code
func (o *DeletePresencedefinitionInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete presencedefinition internal server error response a status code equal to that given
func (o *DeletePresencedefinitionInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DeletePresencedefinitionInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/presencedefinitions/{presenceId}][%d] deletePresencedefinitionInternalServerError  %+v", 500, o.Payload)
}

func (o *DeletePresencedefinitionInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /api/v2/presencedefinitions/{presenceId}][%d] deletePresencedefinitionInternalServerError  %+v", 500, o.Payload)
}

func (o *DeletePresencedefinitionInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeletePresencedefinitionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePresencedefinitionServiceUnavailable creates a DeletePresencedefinitionServiceUnavailable with default headers values
func NewDeletePresencedefinitionServiceUnavailable() *DeletePresencedefinitionServiceUnavailable {
	return &DeletePresencedefinitionServiceUnavailable{}
}

/*
DeletePresencedefinitionServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeletePresencedefinitionServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete presencedefinition service unavailable response has a 2xx status code
func (o *DeletePresencedefinitionServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete presencedefinition service unavailable response has a 3xx status code
func (o *DeletePresencedefinitionServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete presencedefinition service unavailable response has a 4xx status code
func (o *DeletePresencedefinitionServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete presencedefinition service unavailable response has a 5xx status code
func (o *DeletePresencedefinitionServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this delete presencedefinition service unavailable response a status code equal to that given
func (o *DeletePresencedefinitionServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *DeletePresencedefinitionServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/presencedefinitions/{presenceId}][%d] deletePresencedefinitionServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeletePresencedefinitionServiceUnavailable) String() string {
	return fmt.Sprintf("[DELETE /api/v2/presencedefinitions/{presenceId}][%d] deletePresencedefinitionServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeletePresencedefinitionServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeletePresencedefinitionServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePresencedefinitionGatewayTimeout creates a DeletePresencedefinitionGatewayTimeout with default headers values
func NewDeletePresencedefinitionGatewayTimeout() *DeletePresencedefinitionGatewayTimeout {
	return &DeletePresencedefinitionGatewayTimeout{}
}

/*
DeletePresencedefinitionGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type DeletePresencedefinitionGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete presencedefinition gateway timeout response has a 2xx status code
func (o *DeletePresencedefinitionGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete presencedefinition gateway timeout response has a 3xx status code
func (o *DeletePresencedefinitionGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete presencedefinition gateway timeout response has a 4xx status code
func (o *DeletePresencedefinitionGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete presencedefinition gateway timeout response has a 5xx status code
func (o *DeletePresencedefinitionGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this delete presencedefinition gateway timeout response a status code equal to that given
func (o *DeletePresencedefinitionGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *DeletePresencedefinitionGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/presencedefinitions/{presenceId}][%d] deletePresencedefinitionGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeletePresencedefinitionGatewayTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/presencedefinitions/{presenceId}][%d] deletePresencedefinitionGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeletePresencedefinitionGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeletePresencedefinitionGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePresencedefinitionDefault creates a DeletePresencedefinitionDefault with default headers values
func NewDeletePresencedefinitionDefault(code int) *DeletePresencedefinitionDefault {
	return &DeletePresencedefinitionDefault{
		_statusCode: code,
	}
}

/*
DeletePresencedefinitionDefault describes a response with status code -1, with default header values.

successful operation
*/
type DeletePresencedefinitionDefault struct {
	_statusCode int
}

// Code gets the status code for the delete presencedefinition default response
func (o *DeletePresencedefinitionDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this delete presencedefinition default response has a 2xx status code
func (o *DeletePresencedefinitionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete presencedefinition default response has a 3xx status code
func (o *DeletePresencedefinitionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete presencedefinition default response has a 4xx status code
func (o *DeletePresencedefinitionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete presencedefinition default response has a 5xx status code
func (o *DeletePresencedefinitionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete presencedefinition default response a status code equal to that given
func (o *DeletePresencedefinitionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DeletePresencedefinitionDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/presencedefinitions/{presenceId}][%d] deletePresencedefinition default ", o._statusCode)
}

func (o *DeletePresencedefinitionDefault) String() string {
	return fmt.Sprintf("[DELETE /api/v2/presencedefinitions/{presenceId}][%d] deletePresencedefinition default ", o._statusCode)
}

func (o *DeletePresencedefinitionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

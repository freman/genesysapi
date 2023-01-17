// Code generated by go-swagger; DO NOT EDIT.

package web_chat

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// DeleteWebchatDeploymentReader is a Reader for the DeleteWebchatDeployment structure.
type DeleteWebchatDeploymentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteWebchatDeploymentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteWebchatDeploymentNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteWebchatDeploymentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteWebchatDeploymentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteWebchatDeploymentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteWebchatDeploymentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewDeleteWebchatDeploymentRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteWebchatDeploymentRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteWebchatDeploymentUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteWebchatDeploymentTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteWebchatDeploymentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteWebchatDeploymentServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteWebchatDeploymentGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteWebchatDeploymentNoContent creates a DeleteWebchatDeploymentNoContent with default headers values
func NewDeleteWebchatDeploymentNoContent() *DeleteWebchatDeploymentNoContent {
	return &DeleteWebchatDeploymentNoContent{}
}

/*
DeleteWebchatDeploymentNoContent describes a response with status code 204, with default header values.

Deleted
*/
type DeleteWebchatDeploymentNoContent struct {
}

// IsSuccess returns true when this delete webchat deployment no content response has a 2xx status code
func (o *DeleteWebchatDeploymentNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete webchat deployment no content response has a 3xx status code
func (o *DeleteWebchatDeploymentNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete webchat deployment no content response has a 4xx status code
func (o *DeleteWebchatDeploymentNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete webchat deployment no content response has a 5xx status code
func (o *DeleteWebchatDeploymentNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete webchat deployment no content response a status code equal to that given
func (o *DeleteWebchatDeploymentNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DeleteWebchatDeploymentNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/webchat/deployments/{deploymentId}][%d] deleteWebchatDeploymentNoContent ", 204)
}

func (o *DeleteWebchatDeploymentNoContent) String() string {
	return fmt.Sprintf("[DELETE /api/v2/webchat/deployments/{deploymentId}][%d] deleteWebchatDeploymentNoContent ", 204)
}

func (o *DeleteWebchatDeploymentNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteWebchatDeploymentBadRequest creates a DeleteWebchatDeploymentBadRequest with default headers values
func NewDeleteWebchatDeploymentBadRequest() *DeleteWebchatDeploymentBadRequest {
	return &DeleteWebchatDeploymentBadRequest{}
}

/*
DeleteWebchatDeploymentBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteWebchatDeploymentBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete webchat deployment bad request response has a 2xx status code
func (o *DeleteWebchatDeploymentBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete webchat deployment bad request response has a 3xx status code
func (o *DeleteWebchatDeploymentBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete webchat deployment bad request response has a 4xx status code
func (o *DeleteWebchatDeploymentBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete webchat deployment bad request response has a 5xx status code
func (o *DeleteWebchatDeploymentBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete webchat deployment bad request response a status code equal to that given
func (o *DeleteWebchatDeploymentBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DeleteWebchatDeploymentBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/webchat/deployments/{deploymentId}][%d] deleteWebchatDeploymentBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteWebchatDeploymentBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api/v2/webchat/deployments/{deploymentId}][%d] deleteWebchatDeploymentBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteWebchatDeploymentBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWebchatDeploymentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWebchatDeploymentUnauthorized creates a DeleteWebchatDeploymentUnauthorized with default headers values
func NewDeleteWebchatDeploymentUnauthorized() *DeleteWebchatDeploymentUnauthorized {
	return &DeleteWebchatDeploymentUnauthorized{}
}

/*
DeleteWebchatDeploymentUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteWebchatDeploymentUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete webchat deployment unauthorized response has a 2xx status code
func (o *DeleteWebchatDeploymentUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete webchat deployment unauthorized response has a 3xx status code
func (o *DeleteWebchatDeploymentUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete webchat deployment unauthorized response has a 4xx status code
func (o *DeleteWebchatDeploymentUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete webchat deployment unauthorized response has a 5xx status code
func (o *DeleteWebchatDeploymentUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete webchat deployment unauthorized response a status code equal to that given
func (o *DeleteWebchatDeploymentUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteWebchatDeploymentUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/webchat/deployments/{deploymentId}][%d] deleteWebchatDeploymentUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteWebchatDeploymentUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/v2/webchat/deployments/{deploymentId}][%d] deleteWebchatDeploymentUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteWebchatDeploymentUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWebchatDeploymentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWebchatDeploymentForbidden creates a DeleteWebchatDeploymentForbidden with default headers values
func NewDeleteWebchatDeploymentForbidden() *DeleteWebchatDeploymentForbidden {
	return &DeleteWebchatDeploymentForbidden{}
}

/*
DeleteWebchatDeploymentForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type DeleteWebchatDeploymentForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete webchat deployment forbidden response has a 2xx status code
func (o *DeleteWebchatDeploymentForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete webchat deployment forbidden response has a 3xx status code
func (o *DeleteWebchatDeploymentForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete webchat deployment forbidden response has a 4xx status code
func (o *DeleteWebchatDeploymentForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete webchat deployment forbidden response has a 5xx status code
func (o *DeleteWebchatDeploymentForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete webchat deployment forbidden response a status code equal to that given
func (o *DeleteWebchatDeploymentForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteWebchatDeploymentForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/webchat/deployments/{deploymentId}][%d] deleteWebchatDeploymentForbidden  %+v", 403, o.Payload)
}

func (o *DeleteWebchatDeploymentForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v2/webchat/deployments/{deploymentId}][%d] deleteWebchatDeploymentForbidden  %+v", 403, o.Payload)
}

func (o *DeleteWebchatDeploymentForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWebchatDeploymentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWebchatDeploymentNotFound creates a DeleteWebchatDeploymentNotFound with default headers values
func NewDeleteWebchatDeploymentNotFound() *DeleteWebchatDeploymentNotFound {
	return &DeleteWebchatDeploymentNotFound{}
}

/*
DeleteWebchatDeploymentNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type DeleteWebchatDeploymentNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete webchat deployment not found response has a 2xx status code
func (o *DeleteWebchatDeploymentNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete webchat deployment not found response has a 3xx status code
func (o *DeleteWebchatDeploymentNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete webchat deployment not found response has a 4xx status code
func (o *DeleteWebchatDeploymentNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete webchat deployment not found response has a 5xx status code
func (o *DeleteWebchatDeploymentNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete webchat deployment not found response a status code equal to that given
func (o *DeleteWebchatDeploymentNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteWebchatDeploymentNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/webchat/deployments/{deploymentId}][%d] deleteWebchatDeploymentNotFound  %+v", 404, o.Payload)
}

func (o *DeleteWebchatDeploymentNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/v2/webchat/deployments/{deploymentId}][%d] deleteWebchatDeploymentNotFound  %+v", 404, o.Payload)
}

func (o *DeleteWebchatDeploymentNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWebchatDeploymentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWebchatDeploymentRequestTimeout creates a DeleteWebchatDeploymentRequestTimeout with default headers values
func NewDeleteWebchatDeploymentRequestTimeout() *DeleteWebchatDeploymentRequestTimeout {
	return &DeleteWebchatDeploymentRequestTimeout{}
}

/*
DeleteWebchatDeploymentRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type DeleteWebchatDeploymentRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete webchat deployment request timeout response has a 2xx status code
func (o *DeleteWebchatDeploymentRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete webchat deployment request timeout response has a 3xx status code
func (o *DeleteWebchatDeploymentRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete webchat deployment request timeout response has a 4xx status code
func (o *DeleteWebchatDeploymentRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete webchat deployment request timeout response has a 5xx status code
func (o *DeleteWebchatDeploymentRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this delete webchat deployment request timeout response a status code equal to that given
func (o *DeleteWebchatDeploymentRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *DeleteWebchatDeploymentRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/webchat/deployments/{deploymentId}][%d] deleteWebchatDeploymentRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteWebchatDeploymentRequestTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/webchat/deployments/{deploymentId}][%d] deleteWebchatDeploymentRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteWebchatDeploymentRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWebchatDeploymentRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWebchatDeploymentRequestEntityTooLarge creates a DeleteWebchatDeploymentRequestEntityTooLarge with default headers values
func NewDeleteWebchatDeploymentRequestEntityTooLarge() *DeleteWebchatDeploymentRequestEntityTooLarge {
	return &DeleteWebchatDeploymentRequestEntityTooLarge{}
}

/*
DeleteWebchatDeploymentRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type DeleteWebchatDeploymentRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete webchat deployment request entity too large response has a 2xx status code
func (o *DeleteWebchatDeploymentRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete webchat deployment request entity too large response has a 3xx status code
func (o *DeleteWebchatDeploymentRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete webchat deployment request entity too large response has a 4xx status code
func (o *DeleteWebchatDeploymentRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete webchat deployment request entity too large response has a 5xx status code
func (o *DeleteWebchatDeploymentRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this delete webchat deployment request entity too large response a status code equal to that given
func (o *DeleteWebchatDeploymentRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *DeleteWebchatDeploymentRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/webchat/deployments/{deploymentId}][%d] deleteWebchatDeploymentRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteWebchatDeploymentRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[DELETE /api/v2/webchat/deployments/{deploymentId}][%d] deleteWebchatDeploymentRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteWebchatDeploymentRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWebchatDeploymentRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWebchatDeploymentUnsupportedMediaType creates a DeleteWebchatDeploymentUnsupportedMediaType with default headers values
func NewDeleteWebchatDeploymentUnsupportedMediaType() *DeleteWebchatDeploymentUnsupportedMediaType {
	return &DeleteWebchatDeploymentUnsupportedMediaType{}
}

/*
DeleteWebchatDeploymentUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteWebchatDeploymentUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete webchat deployment unsupported media type response has a 2xx status code
func (o *DeleteWebchatDeploymentUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete webchat deployment unsupported media type response has a 3xx status code
func (o *DeleteWebchatDeploymentUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete webchat deployment unsupported media type response has a 4xx status code
func (o *DeleteWebchatDeploymentUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete webchat deployment unsupported media type response has a 5xx status code
func (o *DeleteWebchatDeploymentUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this delete webchat deployment unsupported media type response a status code equal to that given
func (o *DeleteWebchatDeploymentUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *DeleteWebchatDeploymentUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/webchat/deployments/{deploymentId}][%d] deleteWebchatDeploymentUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteWebchatDeploymentUnsupportedMediaType) String() string {
	return fmt.Sprintf("[DELETE /api/v2/webchat/deployments/{deploymentId}][%d] deleteWebchatDeploymentUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteWebchatDeploymentUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWebchatDeploymentUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWebchatDeploymentTooManyRequests creates a DeleteWebchatDeploymentTooManyRequests with default headers values
func NewDeleteWebchatDeploymentTooManyRequests() *DeleteWebchatDeploymentTooManyRequests {
	return &DeleteWebchatDeploymentTooManyRequests{}
}

/*
DeleteWebchatDeploymentTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type DeleteWebchatDeploymentTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete webchat deployment too many requests response has a 2xx status code
func (o *DeleteWebchatDeploymentTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete webchat deployment too many requests response has a 3xx status code
func (o *DeleteWebchatDeploymentTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete webchat deployment too many requests response has a 4xx status code
func (o *DeleteWebchatDeploymentTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete webchat deployment too many requests response has a 5xx status code
func (o *DeleteWebchatDeploymentTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this delete webchat deployment too many requests response a status code equal to that given
func (o *DeleteWebchatDeploymentTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *DeleteWebchatDeploymentTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/webchat/deployments/{deploymentId}][%d] deleteWebchatDeploymentTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteWebchatDeploymentTooManyRequests) String() string {
	return fmt.Sprintf("[DELETE /api/v2/webchat/deployments/{deploymentId}][%d] deleteWebchatDeploymentTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteWebchatDeploymentTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWebchatDeploymentTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWebchatDeploymentInternalServerError creates a DeleteWebchatDeploymentInternalServerError with default headers values
func NewDeleteWebchatDeploymentInternalServerError() *DeleteWebchatDeploymentInternalServerError {
	return &DeleteWebchatDeploymentInternalServerError{}
}

/*
DeleteWebchatDeploymentInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteWebchatDeploymentInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete webchat deployment internal server error response has a 2xx status code
func (o *DeleteWebchatDeploymentInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete webchat deployment internal server error response has a 3xx status code
func (o *DeleteWebchatDeploymentInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete webchat deployment internal server error response has a 4xx status code
func (o *DeleteWebchatDeploymentInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete webchat deployment internal server error response has a 5xx status code
func (o *DeleteWebchatDeploymentInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete webchat deployment internal server error response a status code equal to that given
func (o *DeleteWebchatDeploymentInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DeleteWebchatDeploymentInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/webchat/deployments/{deploymentId}][%d] deleteWebchatDeploymentInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteWebchatDeploymentInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /api/v2/webchat/deployments/{deploymentId}][%d] deleteWebchatDeploymentInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteWebchatDeploymentInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWebchatDeploymentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWebchatDeploymentServiceUnavailable creates a DeleteWebchatDeploymentServiceUnavailable with default headers values
func NewDeleteWebchatDeploymentServiceUnavailable() *DeleteWebchatDeploymentServiceUnavailable {
	return &DeleteWebchatDeploymentServiceUnavailable{}
}

/*
DeleteWebchatDeploymentServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteWebchatDeploymentServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete webchat deployment service unavailable response has a 2xx status code
func (o *DeleteWebchatDeploymentServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete webchat deployment service unavailable response has a 3xx status code
func (o *DeleteWebchatDeploymentServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete webchat deployment service unavailable response has a 4xx status code
func (o *DeleteWebchatDeploymentServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete webchat deployment service unavailable response has a 5xx status code
func (o *DeleteWebchatDeploymentServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this delete webchat deployment service unavailable response a status code equal to that given
func (o *DeleteWebchatDeploymentServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *DeleteWebchatDeploymentServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/webchat/deployments/{deploymentId}][%d] deleteWebchatDeploymentServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteWebchatDeploymentServiceUnavailable) String() string {
	return fmt.Sprintf("[DELETE /api/v2/webchat/deployments/{deploymentId}][%d] deleteWebchatDeploymentServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteWebchatDeploymentServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWebchatDeploymentServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWebchatDeploymentGatewayTimeout creates a DeleteWebchatDeploymentGatewayTimeout with default headers values
func NewDeleteWebchatDeploymentGatewayTimeout() *DeleteWebchatDeploymentGatewayTimeout {
	return &DeleteWebchatDeploymentGatewayTimeout{}
}

/*
DeleteWebchatDeploymentGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type DeleteWebchatDeploymentGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete webchat deployment gateway timeout response has a 2xx status code
func (o *DeleteWebchatDeploymentGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete webchat deployment gateway timeout response has a 3xx status code
func (o *DeleteWebchatDeploymentGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete webchat deployment gateway timeout response has a 4xx status code
func (o *DeleteWebchatDeploymentGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete webchat deployment gateway timeout response has a 5xx status code
func (o *DeleteWebchatDeploymentGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this delete webchat deployment gateway timeout response a status code equal to that given
func (o *DeleteWebchatDeploymentGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *DeleteWebchatDeploymentGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/webchat/deployments/{deploymentId}][%d] deleteWebchatDeploymentGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteWebchatDeploymentGatewayTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/webchat/deployments/{deploymentId}][%d] deleteWebchatDeploymentGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteWebchatDeploymentGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWebchatDeploymentGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

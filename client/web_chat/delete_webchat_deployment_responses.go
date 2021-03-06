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

/*DeleteWebchatDeploymentNoContent handles this case with default header values.

Deleted
*/
type DeleteWebchatDeploymentNoContent struct {
}

func (o *DeleteWebchatDeploymentNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/webchat/deployments/{deploymentId}][%d] deleteWebchatDeploymentNoContent ", 204)
}

func (o *DeleteWebchatDeploymentNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteWebchatDeploymentBadRequest creates a DeleteWebchatDeploymentBadRequest with default headers values
func NewDeleteWebchatDeploymentBadRequest() *DeleteWebchatDeploymentBadRequest {
	return &DeleteWebchatDeploymentBadRequest{}
}

/*DeleteWebchatDeploymentBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteWebchatDeploymentBadRequest struct {
	Payload *models.ErrorBody
}

func (o *DeleteWebchatDeploymentBadRequest) Error() string {
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

/*DeleteWebchatDeploymentUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteWebchatDeploymentUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *DeleteWebchatDeploymentUnauthorized) Error() string {
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

/*DeleteWebchatDeploymentForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type DeleteWebchatDeploymentForbidden struct {
	Payload *models.ErrorBody
}

func (o *DeleteWebchatDeploymentForbidden) Error() string {
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

/*DeleteWebchatDeploymentNotFound handles this case with default header values.

The requested resource was not found.
*/
type DeleteWebchatDeploymentNotFound struct {
	Payload *models.ErrorBody
}

func (o *DeleteWebchatDeploymentNotFound) Error() string {
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

// NewDeleteWebchatDeploymentRequestEntityTooLarge creates a DeleteWebchatDeploymentRequestEntityTooLarge with default headers values
func NewDeleteWebchatDeploymentRequestEntityTooLarge() *DeleteWebchatDeploymentRequestEntityTooLarge {
	return &DeleteWebchatDeploymentRequestEntityTooLarge{}
}

/*DeleteWebchatDeploymentRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type DeleteWebchatDeploymentRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *DeleteWebchatDeploymentRequestEntityTooLarge) Error() string {
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

/*DeleteWebchatDeploymentUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteWebchatDeploymentUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *DeleteWebchatDeploymentUnsupportedMediaType) Error() string {
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

/*DeleteWebchatDeploymentTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type DeleteWebchatDeploymentTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *DeleteWebchatDeploymentTooManyRequests) Error() string {
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

/*DeleteWebchatDeploymentInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteWebchatDeploymentInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *DeleteWebchatDeploymentInternalServerError) Error() string {
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

/*DeleteWebchatDeploymentServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteWebchatDeploymentServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *DeleteWebchatDeploymentServiceUnavailable) Error() string {
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

/*DeleteWebchatDeploymentGatewayTimeout handles this case with default header values.

The request timed out.
*/
type DeleteWebchatDeploymentGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteWebchatDeploymentGatewayTimeout) Error() string {
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

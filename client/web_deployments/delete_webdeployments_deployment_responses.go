// Code generated by go-swagger; DO NOT EDIT.

package web_deployments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// DeleteWebdeploymentsDeploymentReader is a Reader for the DeleteWebdeploymentsDeployment structure.
type DeleteWebdeploymentsDeploymentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteWebdeploymentsDeploymentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteWebdeploymentsDeploymentNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteWebdeploymentsDeploymentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteWebdeploymentsDeploymentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteWebdeploymentsDeploymentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteWebdeploymentsDeploymentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewDeleteWebdeploymentsDeploymentRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteWebdeploymentsDeploymentRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteWebdeploymentsDeploymentUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteWebdeploymentsDeploymentTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteWebdeploymentsDeploymentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteWebdeploymentsDeploymentServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteWebdeploymentsDeploymentGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteWebdeploymentsDeploymentNoContent creates a DeleteWebdeploymentsDeploymentNoContent with default headers values
func NewDeleteWebdeploymentsDeploymentNoContent() *DeleteWebdeploymentsDeploymentNoContent {
	return &DeleteWebdeploymentsDeploymentNoContent{}
}

/*DeleteWebdeploymentsDeploymentNoContent handles this case with default header values.

The deployment was deleted successfully
*/
type DeleteWebdeploymentsDeploymentNoContent struct {
}

func (o *DeleteWebdeploymentsDeploymentNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/webdeployments/deployments/{deploymentId}][%d] deleteWebdeploymentsDeploymentNoContent ", 204)
}

func (o *DeleteWebdeploymentsDeploymentNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteWebdeploymentsDeploymentBadRequest creates a DeleteWebdeploymentsDeploymentBadRequest with default headers values
func NewDeleteWebdeploymentsDeploymentBadRequest() *DeleteWebdeploymentsDeploymentBadRequest {
	return &DeleteWebdeploymentsDeploymentBadRequest{}
}

/*DeleteWebdeploymentsDeploymentBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteWebdeploymentsDeploymentBadRequest struct {
	Payload *models.ErrorBody
}

func (o *DeleteWebdeploymentsDeploymentBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/webdeployments/deployments/{deploymentId}][%d] deleteWebdeploymentsDeploymentBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteWebdeploymentsDeploymentBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWebdeploymentsDeploymentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWebdeploymentsDeploymentUnauthorized creates a DeleteWebdeploymentsDeploymentUnauthorized with default headers values
func NewDeleteWebdeploymentsDeploymentUnauthorized() *DeleteWebdeploymentsDeploymentUnauthorized {
	return &DeleteWebdeploymentsDeploymentUnauthorized{}
}

/*DeleteWebdeploymentsDeploymentUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteWebdeploymentsDeploymentUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *DeleteWebdeploymentsDeploymentUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/webdeployments/deployments/{deploymentId}][%d] deleteWebdeploymentsDeploymentUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteWebdeploymentsDeploymentUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWebdeploymentsDeploymentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWebdeploymentsDeploymentForbidden creates a DeleteWebdeploymentsDeploymentForbidden with default headers values
func NewDeleteWebdeploymentsDeploymentForbidden() *DeleteWebdeploymentsDeploymentForbidden {
	return &DeleteWebdeploymentsDeploymentForbidden{}
}

/*DeleteWebdeploymentsDeploymentForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type DeleteWebdeploymentsDeploymentForbidden struct {
	Payload *models.ErrorBody
}

func (o *DeleteWebdeploymentsDeploymentForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/webdeployments/deployments/{deploymentId}][%d] deleteWebdeploymentsDeploymentForbidden  %+v", 403, o.Payload)
}

func (o *DeleteWebdeploymentsDeploymentForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWebdeploymentsDeploymentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWebdeploymentsDeploymentNotFound creates a DeleteWebdeploymentsDeploymentNotFound with default headers values
func NewDeleteWebdeploymentsDeploymentNotFound() *DeleteWebdeploymentsDeploymentNotFound {
	return &DeleteWebdeploymentsDeploymentNotFound{}
}

/*DeleteWebdeploymentsDeploymentNotFound handles this case with default header values.

The requested resource was not found.
*/
type DeleteWebdeploymentsDeploymentNotFound struct {
	Payload *models.ErrorBody
}

func (o *DeleteWebdeploymentsDeploymentNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/webdeployments/deployments/{deploymentId}][%d] deleteWebdeploymentsDeploymentNotFound  %+v", 404, o.Payload)
}

func (o *DeleteWebdeploymentsDeploymentNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWebdeploymentsDeploymentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWebdeploymentsDeploymentRequestTimeout creates a DeleteWebdeploymentsDeploymentRequestTimeout with default headers values
func NewDeleteWebdeploymentsDeploymentRequestTimeout() *DeleteWebdeploymentsDeploymentRequestTimeout {
	return &DeleteWebdeploymentsDeploymentRequestTimeout{}
}

/*DeleteWebdeploymentsDeploymentRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type DeleteWebdeploymentsDeploymentRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteWebdeploymentsDeploymentRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/webdeployments/deployments/{deploymentId}][%d] deleteWebdeploymentsDeploymentRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteWebdeploymentsDeploymentRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWebdeploymentsDeploymentRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWebdeploymentsDeploymentRequestEntityTooLarge creates a DeleteWebdeploymentsDeploymentRequestEntityTooLarge with default headers values
func NewDeleteWebdeploymentsDeploymentRequestEntityTooLarge() *DeleteWebdeploymentsDeploymentRequestEntityTooLarge {
	return &DeleteWebdeploymentsDeploymentRequestEntityTooLarge{}
}

/*DeleteWebdeploymentsDeploymentRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type DeleteWebdeploymentsDeploymentRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *DeleteWebdeploymentsDeploymentRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/webdeployments/deployments/{deploymentId}][%d] deleteWebdeploymentsDeploymentRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteWebdeploymentsDeploymentRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWebdeploymentsDeploymentRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWebdeploymentsDeploymentUnsupportedMediaType creates a DeleteWebdeploymentsDeploymentUnsupportedMediaType with default headers values
func NewDeleteWebdeploymentsDeploymentUnsupportedMediaType() *DeleteWebdeploymentsDeploymentUnsupportedMediaType {
	return &DeleteWebdeploymentsDeploymentUnsupportedMediaType{}
}

/*DeleteWebdeploymentsDeploymentUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteWebdeploymentsDeploymentUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *DeleteWebdeploymentsDeploymentUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/webdeployments/deployments/{deploymentId}][%d] deleteWebdeploymentsDeploymentUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteWebdeploymentsDeploymentUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWebdeploymentsDeploymentUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWebdeploymentsDeploymentTooManyRequests creates a DeleteWebdeploymentsDeploymentTooManyRequests with default headers values
func NewDeleteWebdeploymentsDeploymentTooManyRequests() *DeleteWebdeploymentsDeploymentTooManyRequests {
	return &DeleteWebdeploymentsDeploymentTooManyRequests{}
}

/*DeleteWebdeploymentsDeploymentTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type DeleteWebdeploymentsDeploymentTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *DeleteWebdeploymentsDeploymentTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/webdeployments/deployments/{deploymentId}][%d] deleteWebdeploymentsDeploymentTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteWebdeploymentsDeploymentTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWebdeploymentsDeploymentTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWebdeploymentsDeploymentInternalServerError creates a DeleteWebdeploymentsDeploymentInternalServerError with default headers values
func NewDeleteWebdeploymentsDeploymentInternalServerError() *DeleteWebdeploymentsDeploymentInternalServerError {
	return &DeleteWebdeploymentsDeploymentInternalServerError{}
}

/*DeleteWebdeploymentsDeploymentInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteWebdeploymentsDeploymentInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *DeleteWebdeploymentsDeploymentInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/webdeployments/deployments/{deploymentId}][%d] deleteWebdeploymentsDeploymentInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteWebdeploymentsDeploymentInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWebdeploymentsDeploymentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWebdeploymentsDeploymentServiceUnavailable creates a DeleteWebdeploymentsDeploymentServiceUnavailable with default headers values
func NewDeleteWebdeploymentsDeploymentServiceUnavailable() *DeleteWebdeploymentsDeploymentServiceUnavailable {
	return &DeleteWebdeploymentsDeploymentServiceUnavailable{}
}

/*DeleteWebdeploymentsDeploymentServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteWebdeploymentsDeploymentServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *DeleteWebdeploymentsDeploymentServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/webdeployments/deployments/{deploymentId}][%d] deleteWebdeploymentsDeploymentServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteWebdeploymentsDeploymentServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWebdeploymentsDeploymentServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWebdeploymentsDeploymentGatewayTimeout creates a DeleteWebdeploymentsDeploymentGatewayTimeout with default headers values
func NewDeleteWebdeploymentsDeploymentGatewayTimeout() *DeleteWebdeploymentsDeploymentGatewayTimeout {
	return &DeleteWebdeploymentsDeploymentGatewayTimeout{}
}

/*DeleteWebdeploymentsDeploymentGatewayTimeout handles this case with default header values.

The request timed out.
*/
type DeleteWebdeploymentsDeploymentGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteWebdeploymentsDeploymentGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/webdeployments/deployments/{deploymentId}][%d] deleteWebdeploymentsDeploymentGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteWebdeploymentsDeploymentGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWebdeploymentsDeploymentGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

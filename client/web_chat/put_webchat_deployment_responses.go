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

// PutWebchatDeploymentReader is a Reader for the PutWebchatDeployment structure.
type PutWebchatDeploymentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutWebchatDeploymentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutWebchatDeploymentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutWebchatDeploymentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutWebchatDeploymentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutWebchatDeploymentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutWebchatDeploymentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPutWebchatDeploymentRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPutWebchatDeploymentUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutWebchatDeploymentTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutWebchatDeploymentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutWebchatDeploymentServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutWebchatDeploymentGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutWebchatDeploymentOK creates a PutWebchatDeploymentOK with default headers values
func NewPutWebchatDeploymentOK() *PutWebchatDeploymentOK {
	return &PutWebchatDeploymentOK{}
}

/*PutWebchatDeploymentOK handles this case with default header values.

successful operation
*/
type PutWebchatDeploymentOK struct {
	Payload *models.WebChatDeployment
}

func (o *PutWebchatDeploymentOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/webchat/deployments/{deploymentId}][%d] putWebchatDeploymentOK  %+v", 200, o.Payload)
}

func (o *PutWebchatDeploymentOK) GetPayload() *models.WebChatDeployment {
	return o.Payload
}

func (o *PutWebchatDeploymentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WebChatDeployment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutWebchatDeploymentBadRequest creates a PutWebchatDeploymentBadRequest with default headers values
func NewPutWebchatDeploymentBadRequest() *PutWebchatDeploymentBadRequest {
	return &PutWebchatDeploymentBadRequest{}
}

/*PutWebchatDeploymentBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutWebchatDeploymentBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PutWebchatDeploymentBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/webchat/deployments/{deploymentId}][%d] putWebchatDeploymentBadRequest  %+v", 400, o.Payload)
}

func (o *PutWebchatDeploymentBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutWebchatDeploymentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutWebchatDeploymentUnauthorized creates a PutWebchatDeploymentUnauthorized with default headers values
func NewPutWebchatDeploymentUnauthorized() *PutWebchatDeploymentUnauthorized {
	return &PutWebchatDeploymentUnauthorized{}
}

/*PutWebchatDeploymentUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PutWebchatDeploymentUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PutWebchatDeploymentUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/webchat/deployments/{deploymentId}][%d] putWebchatDeploymentUnauthorized  %+v", 401, o.Payload)
}

func (o *PutWebchatDeploymentUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutWebchatDeploymentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutWebchatDeploymentForbidden creates a PutWebchatDeploymentForbidden with default headers values
func NewPutWebchatDeploymentForbidden() *PutWebchatDeploymentForbidden {
	return &PutWebchatDeploymentForbidden{}
}

/*PutWebchatDeploymentForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PutWebchatDeploymentForbidden struct {
	Payload *models.ErrorBody
}

func (o *PutWebchatDeploymentForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/webchat/deployments/{deploymentId}][%d] putWebchatDeploymentForbidden  %+v", 403, o.Payload)
}

func (o *PutWebchatDeploymentForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutWebchatDeploymentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutWebchatDeploymentNotFound creates a PutWebchatDeploymentNotFound with default headers values
func NewPutWebchatDeploymentNotFound() *PutWebchatDeploymentNotFound {
	return &PutWebchatDeploymentNotFound{}
}

/*PutWebchatDeploymentNotFound handles this case with default header values.

The requested resource was not found.
*/
type PutWebchatDeploymentNotFound struct {
	Payload *models.ErrorBody
}

func (o *PutWebchatDeploymentNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/webchat/deployments/{deploymentId}][%d] putWebchatDeploymentNotFound  %+v", 404, o.Payload)
}

func (o *PutWebchatDeploymentNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutWebchatDeploymentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutWebchatDeploymentRequestEntityTooLarge creates a PutWebchatDeploymentRequestEntityTooLarge with default headers values
func NewPutWebchatDeploymentRequestEntityTooLarge() *PutWebchatDeploymentRequestEntityTooLarge {
	return &PutWebchatDeploymentRequestEntityTooLarge{}
}

/*PutWebchatDeploymentRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PutWebchatDeploymentRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PutWebchatDeploymentRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /api/v2/webchat/deployments/{deploymentId}][%d] putWebchatDeploymentRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutWebchatDeploymentRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutWebchatDeploymentRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutWebchatDeploymentUnsupportedMediaType creates a PutWebchatDeploymentUnsupportedMediaType with default headers values
func NewPutWebchatDeploymentUnsupportedMediaType() *PutWebchatDeploymentUnsupportedMediaType {
	return &PutWebchatDeploymentUnsupportedMediaType{}
}

/*PutWebchatDeploymentUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutWebchatDeploymentUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PutWebchatDeploymentUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v2/webchat/deployments/{deploymentId}][%d] putWebchatDeploymentUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutWebchatDeploymentUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutWebchatDeploymentUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutWebchatDeploymentTooManyRequests creates a PutWebchatDeploymentTooManyRequests with default headers values
func NewPutWebchatDeploymentTooManyRequests() *PutWebchatDeploymentTooManyRequests {
	return &PutWebchatDeploymentTooManyRequests{}
}

/*PutWebchatDeploymentTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PutWebchatDeploymentTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PutWebchatDeploymentTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v2/webchat/deployments/{deploymentId}][%d] putWebchatDeploymentTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutWebchatDeploymentTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutWebchatDeploymentTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutWebchatDeploymentInternalServerError creates a PutWebchatDeploymentInternalServerError with default headers values
func NewPutWebchatDeploymentInternalServerError() *PutWebchatDeploymentInternalServerError {
	return &PutWebchatDeploymentInternalServerError{}
}

/*PutWebchatDeploymentInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutWebchatDeploymentInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PutWebchatDeploymentInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/webchat/deployments/{deploymentId}][%d] putWebchatDeploymentInternalServerError  %+v", 500, o.Payload)
}

func (o *PutWebchatDeploymentInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutWebchatDeploymentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutWebchatDeploymentServiceUnavailable creates a PutWebchatDeploymentServiceUnavailable with default headers values
func NewPutWebchatDeploymentServiceUnavailable() *PutWebchatDeploymentServiceUnavailable {
	return &PutWebchatDeploymentServiceUnavailable{}
}

/*PutWebchatDeploymentServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutWebchatDeploymentServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PutWebchatDeploymentServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/webchat/deployments/{deploymentId}][%d] putWebchatDeploymentServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutWebchatDeploymentServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutWebchatDeploymentServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutWebchatDeploymentGatewayTimeout creates a PutWebchatDeploymentGatewayTimeout with default headers values
func NewPutWebchatDeploymentGatewayTimeout() *PutWebchatDeploymentGatewayTimeout {
	return &PutWebchatDeploymentGatewayTimeout{}
}

/*PutWebchatDeploymentGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PutWebchatDeploymentGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PutWebchatDeploymentGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/webchat/deployments/{deploymentId}][%d] putWebchatDeploymentGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutWebchatDeploymentGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutWebchatDeploymentGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

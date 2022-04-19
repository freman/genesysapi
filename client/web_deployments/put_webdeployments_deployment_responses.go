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

// PutWebdeploymentsDeploymentReader is a Reader for the PutWebdeploymentsDeployment structure.
type PutWebdeploymentsDeploymentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutWebdeploymentsDeploymentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutWebdeploymentsDeploymentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutWebdeploymentsDeploymentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutWebdeploymentsDeploymentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutWebdeploymentsDeploymentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutWebdeploymentsDeploymentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPutWebdeploymentsDeploymentRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPutWebdeploymentsDeploymentRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPutWebdeploymentsDeploymentUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutWebdeploymentsDeploymentTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutWebdeploymentsDeploymentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutWebdeploymentsDeploymentServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutWebdeploymentsDeploymentGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutWebdeploymentsDeploymentOK creates a PutWebdeploymentsDeploymentOK with default headers values
func NewPutWebdeploymentsDeploymentOK() *PutWebdeploymentsDeploymentOK {
	return &PutWebdeploymentsDeploymentOK{}
}

/*PutWebdeploymentsDeploymentOK handles this case with default header values.

successful operation
*/
type PutWebdeploymentsDeploymentOK struct {
	Payload *models.WebDeployment
}

func (o *PutWebdeploymentsDeploymentOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/webdeployments/deployments/{deploymentId}][%d] putWebdeploymentsDeploymentOK  %+v", 200, o.Payload)
}

func (o *PutWebdeploymentsDeploymentOK) GetPayload() *models.WebDeployment {
	return o.Payload
}

func (o *PutWebdeploymentsDeploymentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WebDeployment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutWebdeploymentsDeploymentBadRequest creates a PutWebdeploymentsDeploymentBadRequest with default headers values
func NewPutWebdeploymentsDeploymentBadRequest() *PutWebdeploymentsDeploymentBadRequest {
	return &PutWebdeploymentsDeploymentBadRequest{}
}

/*PutWebdeploymentsDeploymentBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutWebdeploymentsDeploymentBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PutWebdeploymentsDeploymentBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/webdeployments/deployments/{deploymentId}][%d] putWebdeploymentsDeploymentBadRequest  %+v", 400, o.Payload)
}

func (o *PutWebdeploymentsDeploymentBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutWebdeploymentsDeploymentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutWebdeploymentsDeploymentUnauthorized creates a PutWebdeploymentsDeploymentUnauthorized with default headers values
func NewPutWebdeploymentsDeploymentUnauthorized() *PutWebdeploymentsDeploymentUnauthorized {
	return &PutWebdeploymentsDeploymentUnauthorized{}
}

/*PutWebdeploymentsDeploymentUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PutWebdeploymentsDeploymentUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PutWebdeploymentsDeploymentUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/webdeployments/deployments/{deploymentId}][%d] putWebdeploymentsDeploymentUnauthorized  %+v", 401, o.Payload)
}

func (o *PutWebdeploymentsDeploymentUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutWebdeploymentsDeploymentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutWebdeploymentsDeploymentForbidden creates a PutWebdeploymentsDeploymentForbidden with default headers values
func NewPutWebdeploymentsDeploymentForbidden() *PutWebdeploymentsDeploymentForbidden {
	return &PutWebdeploymentsDeploymentForbidden{}
}

/*PutWebdeploymentsDeploymentForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PutWebdeploymentsDeploymentForbidden struct {
	Payload *models.ErrorBody
}

func (o *PutWebdeploymentsDeploymentForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/webdeployments/deployments/{deploymentId}][%d] putWebdeploymentsDeploymentForbidden  %+v", 403, o.Payload)
}

func (o *PutWebdeploymentsDeploymentForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutWebdeploymentsDeploymentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutWebdeploymentsDeploymentNotFound creates a PutWebdeploymentsDeploymentNotFound with default headers values
func NewPutWebdeploymentsDeploymentNotFound() *PutWebdeploymentsDeploymentNotFound {
	return &PutWebdeploymentsDeploymentNotFound{}
}

/*PutWebdeploymentsDeploymentNotFound handles this case with default header values.

The requested resource was not found.
*/
type PutWebdeploymentsDeploymentNotFound struct {
	Payload *models.ErrorBody
}

func (o *PutWebdeploymentsDeploymentNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/webdeployments/deployments/{deploymentId}][%d] putWebdeploymentsDeploymentNotFound  %+v", 404, o.Payload)
}

func (o *PutWebdeploymentsDeploymentNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutWebdeploymentsDeploymentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutWebdeploymentsDeploymentRequestTimeout creates a PutWebdeploymentsDeploymentRequestTimeout with default headers values
func NewPutWebdeploymentsDeploymentRequestTimeout() *PutWebdeploymentsDeploymentRequestTimeout {
	return &PutWebdeploymentsDeploymentRequestTimeout{}
}

/*PutWebdeploymentsDeploymentRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PutWebdeploymentsDeploymentRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PutWebdeploymentsDeploymentRequestTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/webdeployments/deployments/{deploymentId}][%d] putWebdeploymentsDeploymentRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutWebdeploymentsDeploymentRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutWebdeploymentsDeploymentRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutWebdeploymentsDeploymentRequestEntityTooLarge creates a PutWebdeploymentsDeploymentRequestEntityTooLarge with default headers values
func NewPutWebdeploymentsDeploymentRequestEntityTooLarge() *PutWebdeploymentsDeploymentRequestEntityTooLarge {
	return &PutWebdeploymentsDeploymentRequestEntityTooLarge{}
}

/*PutWebdeploymentsDeploymentRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PutWebdeploymentsDeploymentRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PutWebdeploymentsDeploymentRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /api/v2/webdeployments/deployments/{deploymentId}][%d] putWebdeploymentsDeploymentRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutWebdeploymentsDeploymentRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutWebdeploymentsDeploymentRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutWebdeploymentsDeploymentUnsupportedMediaType creates a PutWebdeploymentsDeploymentUnsupportedMediaType with default headers values
func NewPutWebdeploymentsDeploymentUnsupportedMediaType() *PutWebdeploymentsDeploymentUnsupportedMediaType {
	return &PutWebdeploymentsDeploymentUnsupportedMediaType{}
}

/*PutWebdeploymentsDeploymentUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutWebdeploymentsDeploymentUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PutWebdeploymentsDeploymentUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v2/webdeployments/deployments/{deploymentId}][%d] putWebdeploymentsDeploymentUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutWebdeploymentsDeploymentUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutWebdeploymentsDeploymentUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutWebdeploymentsDeploymentTooManyRequests creates a PutWebdeploymentsDeploymentTooManyRequests with default headers values
func NewPutWebdeploymentsDeploymentTooManyRequests() *PutWebdeploymentsDeploymentTooManyRequests {
	return &PutWebdeploymentsDeploymentTooManyRequests{}
}

/*PutWebdeploymentsDeploymentTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PutWebdeploymentsDeploymentTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PutWebdeploymentsDeploymentTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v2/webdeployments/deployments/{deploymentId}][%d] putWebdeploymentsDeploymentTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutWebdeploymentsDeploymentTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutWebdeploymentsDeploymentTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutWebdeploymentsDeploymentInternalServerError creates a PutWebdeploymentsDeploymentInternalServerError with default headers values
func NewPutWebdeploymentsDeploymentInternalServerError() *PutWebdeploymentsDeploymentInternalServerError {
	return &PutWebdeploymentsDeploymentInternalServerError{}
}

/*PutWebdeploymentsDeploymentInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutWebdeploymentsDeploymentInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PutWebdeploymentsDeploymentInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/webdeployments/deployments/{deploymentId}][%d] putWebdeploymentsDeploymentInternalServerError  %+v", 500, o.Payload)
}

func (o *PutWebdeploymentsDeploymentInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutWebdeploymentsDeploymentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutWebdeploymentsDeploymentServiceUnavailable creates a PutWebdeploymentsDeploymentServiceUnavailable with default headers values
func NewPutWebdeploymentsDeploymentServiceUnavailable() *PutWebdeploymentsDeploymentServiceUnavailable {
	return &PutWebdeploymentsDeploymentServiceUnavailable{}
}

/*PutWebdeploymentsDeploymentServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutWebdeploymentsDeploymentServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PutWebdeploymentsDeploymentServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/webdeployments/deployments/{deploymentId}][%d] putWebdeploymentsDeploymentServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutWebdeploymentsDeploymentServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutWebdeploymentsDeploymentServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutWebdeploymentsDeploymentGatewayTimeout creates a PutWebdeploymentsDeploymentGatewayTimeout with default headers values
func NewPutWebdeploymentsDeploymentGatewayTimeout() *PutWebdeploymentsDeploymentGatewayTimeout {
	return &PutWebdeploymentsDeploymentGatewayTimeout{}
}

/*PutWebdeploymentsDeploymentGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PutWebdeploymentsDeploymentGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PutWebdeploymentsDeploymentGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/webdeployments/deployments/{deploymentId}][%d] putWebdeploymentsDeploymentGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutWebdeploymentsDeploymentGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutWebdeploymentsDeploymentGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

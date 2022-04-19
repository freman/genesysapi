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

// GetWebdeploymentsDeploymentReader is a Reader for the GetWebdeploymentsDeployment structure.
type GetWebdeploymentsDeploymentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWebdeploymentsDeploymentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWebdeploymentsDeploymentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetWebdeploymentsDeploymentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetWebdeploymentsDeploymentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetWebdeploymentsDeploymentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetWebdeploymentsDeploymentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetWebdeploymentsDeploymentRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetWebdeploymentsDeploymentRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetWebdeploymentsDeploymentUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetWebdeploymentsDeploymentTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetWebdeploymentsDeploymentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetWebdeploymentsDeploymentServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetWebdeploymentsDeploymentGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetWebdeploymentsDeploymentOK creates a GetWebdeploymentsDeploymentOK with default headers values
func NewGetWebdeploymentsDeploymentOK() *GetWebdeploymentsDeploymentOK {
	return &GetWebdeploymentsDeploymentOK{}
}

/*GetWebdeploymentsDeploymentOK handles this case with default header values.

successful operation
*/
type GetWebdeploymentsDeploymentOK struct {
	Payload *models.WebDeployment
}

func (o *GetWebdeploymentsDeploymentOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/webdeployments/deployments/{deploymentId}][%d] getWebdeploymentsDeploymentOK  %+v", 200, o.Payload)
}

func (o *GetWebdeploymentsDeploymentOK) GetPayload() *models.WebDeployment {
	return o.Payload
}

func (o *GetWebdeploymentsDeploymentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WebDeployment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWebdeploymentsDeploymentBadRequest creates a GetWebdeploymentsDeploymentBadRequest with default headers values
func NewGetWebdeploymentsDeploymentBadRequest() *GetWebdeploymentsDeploymentBadRequest {
	return &GetWebdeploymentsDeploymentBadRequest{}
}

/*GetWebdeploymentsDeploymentBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetWebdeploymentsDeploymentBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetWebdeploymentsDeploymentBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/webdeployments/deployments/{deploymentId}][%d] getWebdeploymentsDeploymentBadRequest  %+v", 400, o.Payload)
}

func (o *GetWebdeploymentsDeploymentBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWebdeploymentsDeploymentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWebdeploymentsDeploymentUnauthorized creates a GetWebdeploymentsDeploymentUnauthorized with default headers values
func NewGetWebdeploymentsDeploymentUnauthorized() *GetWebdeploymentsDeploymentUnauthorized {
	return &GetWebdeploymentsDeploymentUnauthorized{}
}

/*GetWebdeploymentsDeploymentUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetWebdeploymentsDeploymentUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetWebdeploymentsDeploymentUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/webdeployments/deployments/{deploymentId}][%d] getWebdeploymentsDeploymentUnauthorized  %+v", 401, o.Payload)
}

func (o *GetWebdeploymentsDeploymentUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWebdeploymentsDeploymentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWebdeploymentsDeploymentForbidden creates a GetWebdeploymentsDeploymentForbidden with default headers values
func NewGetWebdeploymentsDeploymentForbidden() *GetWebdeploymentsDeploymentForbidden {
	return &GetWebdeploymentsDeploymentForbidden{}
}

/*GetWebdeploymentsDeploymentForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetWebdeploymentsDeploymentForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetWebdeploymentsDeploymentForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/webdeployments/deployments/{deploymentId}][%d] getWebdeploymentsDeploymentForbidden  %+v", 403, o.Payload)
}

func (o *GetWebdeploymentsDeploymentForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWebdeploymentsDeploymentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWebdeploymentsDeploymentNotFound creates a GetWebdeploymentsDeploymentNotFound with default headers values
func NewGetWebdeploymentsDeploymentNotFound() *GetWebdeploymentsDeploymentNotFound {
	return &GetWebdeploymentsDeploymentNotFound{}
}

/*GetWebdeploymentsDeploymentNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetWebdeploymentsDeploymentNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetWebdeploymentsDeploymentNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/webdeployments/deployments/{deploymentId}][%d] getWebdeploymentsDeploymentNotFound  %+v", 404, o.Payload)
}

func (o *GetWebdeploymentsDeploymentNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWebdeploymentsDeploymentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWebdeploymentsDeploymentRequestTimeout creates a GetWebdeploymentsDeploymentRequestTimeout with default headers values
func NewGetWebdeploymentsDeploymentRequestTimeout() *GetWebdeploymentsDeploymentRequestTimeout {
	return &GetWebdeploymentsDeploymentRequestTimeout{}
}

/*GetWebdeploymentsDeploymentRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetWebdeploymentsDeploymentRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetWebdeploymentsDeploymentRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/webdeployments/deployments/{deploymentId}][%d] getWebdeploymentsDeploymentRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetWebdeploymentsDeploymentRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWebdeploymentsDeploymentRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWebdeploymentsDeploymentRequestEntityTooLarge creates a GetWebdeploymentsDeploymentRequestEntityTooLarge with default headers values
func NewGetWebdeploymentsDeploymentRequestEntityTooLarge() *GetWebdeploymentsDeploymentRequestEntityTooLarge {
	return &GetWebdeploymentsDeploymentRequestEntityTooLarge{}
}

/*GetWebdeploymentsDeploymentRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetWebdeploymentsDeploymentRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetWebdeploymentsDeploymentRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/webdeployments/deployments/{deploymentId}][%d] getWebdeploymentsDeploymentRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetWebdeploymentsDeploymentRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWebdeploymentsDeploymentRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWebdeploymentsDeploymentUnsupportedMediaType creates a GetWebdeploymentsDeploymentUnsupportedMediaType with default headers values
func NewGetWebdeploymentsDeploymentUnsupportedMediaType() *GetWebdeploymentsDeploymentUnsupportedMediaType {
	return &GetWebdeploymentsDeploymentUnsupportedMediaType{}
}

/*GetWebdeploymentsDeploymentUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetWebdeploymentsDeploymentUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetWebdeploymentsDeploymentUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/webdeployments/deployments/{deploymentId}][%d] getWebdeploymentsDeploymentUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetWebdeploymentsDeploymentUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWebdeploymentsDeploymentUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWebdeploymentsDeploymentTooManyRequests creates a GetWebdeploymentsDeploymentTooManyRequests with default headers values
func NewGetWebdeploymentsDeploymentTooManyRequests() *GetWebdeploymentsDeploymentTooManyRequests {
	return &GetWebdeploymentsDeploymentTooManyRequests{}
}

/*GetWebdeploymentsDeploymentTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetWebdeploymentsDeploymentTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetWebdeploymentsDeploymentTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/webdeployments/deployments/{deploymentId}][%d] getWebdeploymentsDeploymentTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetWebdeploymentsDeploymentTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWebdeploymentsDeploymentTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWebdeploymentsDeploymentInternalServerError creates a GetWebdeploymentsDeploymentInternalServerError with default headers values
func NewGetWebdeploymentsDeploymentInternalServerError() *GetWebdeploymentsDeploymentInternalServerError {
	return &GetWebdeploymentsDeploymentInternalServerError{}
}

/*GetWebdeploymentsDeploymentInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetWebdeploymentsDeploymentInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetWebdeploymentsDeploymentInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/webdeployments/deployments/{deploymentId}][%d] getWebdeploymentsDeploymentInternalServerError  %+v", 500, o.Payload)
}

func (o *GetWebdeploymentsDeploymentInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWebdeploymentsDeploymentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWebdeploymentsDeploymentServiceUnavailable creates a GetWebdeploymentsDeploymentServiceUnavailable with default headers values
func NewGetWebdeploymentsDeploymentServiceUnavailable() *GetWebdeploymentsDeploymentServiceUnavailable {
	return &GetWebdeploymentsDeploymentServiceUnavailable{}
}

/*GetWebdeploymentsDeploymentServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetWebdeploymentsDeploymentServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetWebdeploymentsDeploymentServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/webdeployments/deployments/{deploymentId}][%d] getWebdeploymentsDeploymentServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetWebdeploymentsDeploymentServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWebdeploymentsDeploymentServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWebdeploymentsDeploymentGatewayTimeout creates a GetWebdeploymentsDeploymentGatewayTimeout with default headers values
func NewGetWebdeploymentsDeploymentGatewayTimeout() *GetWebdeploymentsDeploymentGatewayTimeout {
	return &GetWebdeploymentsDeploymentGatewayTimeout{}
}

/*GetWebdeploymentsDeploymentGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetWebdeploymentsDeploymentGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetWebdeploymentsDeploymentGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/webdeployments/deployments/{deploymentId}][%d] getWebdeploymentsDeploymentGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetWebdeploymentsDeploymentGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWebdeploymentsDeploymentGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package widgets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetWidgetsDeploymentReader is a Reader for the GetWidgetsDeployment structure.
type GetWidgetsDeploymentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWidgetsDeploymentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWidgetsDeploymentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetWidgetsDeploymentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetWidgetsDeploymentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetWidgetsDeploymentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetWidgetsDeploymentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetWidgetsDeploymentRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetWidgetsDeploymentRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetWidgetsDeploymentUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetWidgetsDeploymentTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetWidgetsDeploymentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetWidgetsDeploymentServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetWidgetsDeploymentGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetWidgetsDeploymentOK creates a GetWidgetsDeploymentOK with default headers values
func NewGetWidgetsDeploymentOK() *GetWidgetsDeploymentOK {
	return &GetWidgetsDeploymentOK{}
}

/*GetWidgetsDeploymentOK handles this case with default header values.

successful operation
*/
type GetWidgetsDeploymentOK struct {
	Payload *models.WidgetDeployment
}

func (o *GetWidgetsDeploymentOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/widgets/deployments/{deploymentId}][%d] getWidgetsDeploymentOK  %+v", 200, o.Payload)
}

func (o *GetWidgetsDeploymentOK) GetPayload() *models.WidgetDeployment {
	return o.Payload
}

func (o *GetWidgetsDeploymentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WidgetDeployment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWidgetsDeploymentBadRequest creates a GetWidgetsDeploymentBadRequest with default headers values
func NewGetWidgetsDeploymentBadRequest() *GetWidgetsDeploymentBadRequest {
	return &GetWidgetsDeploymentBadRequest{}
}

/*GetWidgetsDeploymentBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetWidgetsDeploymentBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetWidgetsDeploymentBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/widgets/deployments/{deploymentId}][%d] getWidgetsDeploymentBadRequest  %+v", 400, o.Payload)
}

func (o *GetWidgetsDeploymentBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWidgetsDeploymentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWidgetsDeploymentUnauthorized creates a GetWidgetsDeploymentUnauthorized with default headers values
func NewGetWidgetsDeploymentUnauthorized() *GetWidgetsDeploymentUnauthorized {
	return &GetWidgetsDeploymentUnauthorized{}
}

/*GetWidgetsDeploymentUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetWidgetsDeploymentUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetWidgetsDeploymentUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/widgets/deployments/{deploymentId}][%d] getWidgetsDeploymentUnauthorized  %+v", 401, o.Payload)
}

func (o *GetWidgetsDeploymentUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWidgetsDeploymentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWidgetsDeploymentForbidden creates a GetWidgetsDeploymentForbidden with default headers values
func NewGetWidgetsDeploymentForbidden() *GetWidgetsDeploymentForbidden {
	return &GetWidgetsDeploymentForbidden{}
}

/*GetWidgetsDeploymentForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetWidgetsDeploymentForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetWidgetsDeploymentForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/widgets/deployments/{deploymentId}][%d] getWidgetsDeploymentForbidden  %+v", 403, o.Payload)
}

func (o *GetWidgetsDeploymentForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWidgetsDeploymentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWidgetsDeploymentNotFound creates a GetWidgetsDeploymentNotFound with default headers values
func NewGetWidgetsDeploymentNotFound() *GetWidgetsDeploymentNotFound {
	return &GetWidgetsDeploymentNotFound{}
}

/*GetWidgetsDeploymentNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetWidgetsDeploymentNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetWidgetsDeploymentNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/widgets/deployments/{deploymentId}][%d] getWidgetsDeploymentNotFound  %+v", 404, o.Payload)
}

func (o *GetWidgetsDeploymentNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWidgetsDeploymentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWidgetsDeploymentRequestTimeout creates a GetWidgetsDeploymentRequestTimeout with default headers values
func NewGetWidgetsDeploymentRequestTimeout() *GetWidgetsDeploymentRequestTimeout {
	return &GetWidgetsDeploymentRequestTimeout{}
}

/*GetWidgetsDeploymentRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetWidgetsDeploymentRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetWidgetsDeploymentRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/widgets/deployments/{deploymentId}][%d] getWidgetsDeploymentRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetWidgetsDeploymentRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWidgetsDeploymentRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWidgetsDeploymentRequestEntityTooLarge creates a GetWidgetsDeploymentRequestEntityTooLarge with default headers values
func NewGetWidgetsDeploymentRequestEntityTooLarge() *GetWidgetsDeploymentRequestEntityTooLarge {
	return &GetWidgetsDeploymentRequestEntityTooLarge{}
}

/*GetWidgetsDeploymentRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetWidgetsDeploymentRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetWidgetsDeploymentRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/widgets/deployments/{deploymentId}][%d] getWidgetsDeploymentRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetWidgetsDeploymentRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWidgetsDeploymentRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWidgetsDeploymentUnsupportedMediaType creates a GetWidgetsDeploymentUnsupportedMediaType with default headers values
func NewGetWidgetsDeploymentUnsupportedMediaType() *GetWidgetsDeploymentUnsupportedMediaType {
	return &GetWidgetsDeploymentUnsupportedMediaType{}
}

/*GetWidgetsDeploymentUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetWidgetsDeploymentUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetWidgetsDeploymentUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/widgets/deployments/{deploymentId}][%d] getWidgetsDeploymentUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetWidgetsDeploymentUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWidgetsDeploymentUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWidgetsDeploymentTooManyRequests creates a GetWidgetsDeploymentTooManyRequests with default headers values
func NewGetWidgetsDeploymentTooManyRequests() *GetWidgetsDeploymentTooManyRequests {
	return &GetWidgetsDeploymentTooManyRequests{}
}

/*GetWidgetsDeploymentTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetWidgetsDeploymentTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetWidgetsDeploymentTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/widgets/deployments/{deploymentId}][%d] getWidgetsDeploymentTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetWidgetsDeploymentTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWidgetsDeploymentTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWidgetsDeploymentInternalServerError creates a GetWidgetsDeploymentInternalServerError with default headers values
func NewGetWidgetsDeploymentInternalServerError() *GetWidgetsDeploymentInternalServerError {
	return &GetWidgetsDeploymentInternalServerError{}
}

/*GetWidgetsDeploymentInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetWidgetsDeploymentInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetWidgetsDeploymentInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/widgets/deployments/{deploymentId}][%d] getWidgetsDeploymentInternalServerError  %+v", 500, o.Payload)
}

func (o *GetWidgetsDeploymentInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWidgetsDeploymentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWidgetsDeploymentServiceUnavailable creates a GetWidgetsDeploymentServiceUnavailable with default headers values
func NewGetWidgetsDeploymentServiceUnavailable() *GetWidgetsDeploymentServiceUnavailable {
	return &GetWidgetsDeploymentServiceUnavailable{}
}

/*GetWidgetsDeploymentServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetWidgetsDeploymentServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetWidgetsDeploymentServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/widgets/deployments/{deploymentId}][%d] getWidgetsDeploymentServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetWidgetsDeploymentServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWidgetsDeploymentServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWidgetsDeploymentGatewayTimeout creates a GetWidgetsDeploymentGatewayTimeout with default headers values
func NewGetWidgetsDeploymentGatewayTimeout() *GetWidgetsDeploymentGatewayTimeout {
	return &GetWidgetsDeploymentGatewayTimeout{}
}

/*GetWidgetsDeploymentGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetWidgetsDeploymentGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetWidgetsDeploymentGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/widgets/deployments/{deploymentId}][%d] getWidgetsDeploymentGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetWidgetsDeploymentGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWidgetsDeploymentGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

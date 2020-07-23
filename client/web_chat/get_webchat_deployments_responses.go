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

// GetWebchatDeploymentsReader is a Reader for the GetWebchatDeployments structure.
type GetWebchatDeploymentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWebchatDeploymentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWebchatDeploymentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetWebchatDeploymentsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetWebchatDeploymentsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetWebchatDeploymentsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetWebchatDeploymentsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetWebchatDeploymentsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetWebchatDeploymentsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetWebchatDeploymentsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetWebchatDeploymentsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetWebchatDeploymentsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetWebchatDeploymentsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetWebchatDeploymentsOK creates a GetWebchatDeploymentsOK with default headers values
func NewGetWebchatDeploymentsOK() *GetWebchatDeploymentsOK {
	return &GetWebchatDeploymentsOK{}
}

/*GetWebchatDeploymentsOK handles this case with default header values.

successful operation
*/
type GetWebchatDeploymentsOK struct {
	Payload *models.WebChatDeploymentEntityListing
}

func (o *GetWebchatDeploymentsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/webchat/deployments][%d] getWebchatDeploymentsOK  %+v", 200, o.Payload)
}

func (o *GetWebchatDeploymentsOK) GetPayload() *models.WebChatDeploymentEntityListing {
	return o.Payload
}

func (o *GetWebchatDeploymentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WebChatDeploymentEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWebchatDeploymentsBadRequest creates a GetWebchatDeploymentsBadRequest with default headers values
func NewGetWebchatDeploymentsBadRequest() *GetWebchatDeploymentsBadRequest {
	return &GetWebchatDeploymentsBadRequest{}
}

/*GetWebchatDeploymentsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetWebchatDeploymentsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetWebchatDeploymentsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/webchat/deployments][%d] getWebchatDeploymentsBadRequest  %+v", 400, o.Payload)
}

func (o *GetWebchatDeploymentsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWebchatDeploymentsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWebchatDeploymentsUnauthorized creates a GetWebchatDeploymentsUnauthorized with default headers values
func NewGetWebchatDeploymentsUnauthorized() *GetWebchatDeploymentsUnauthorized {
	return &GetWebchatDeploymentsUnauthorized{}
}

/*GetWebchatDeploymentsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetWebchatDeploymentsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetWebchatDeploymentsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/webchat/deployments][%d] getWebchatDeploymentsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetWebchatDeploymentsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWebchatDeploymentsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWebchatDeploymentsForbidden creates a GetWebchatDeploymentsForbidden with default headers values
func NewGetWebchatDeploymentsForbidden() *GetWebchatDeploymentsForbidden {
	return &GetWebchatDeploymentsForbidden{}
}

/*GetWebchatDeploymentsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetWebchatDeploymentsForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetWebchatDeploymentsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/webchat/deployments][%d] getWebchatDeploymentsForbidden  %+v", 403, o.Payload)
}

func (o *GetWebchatDeploymentsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWebchatDeploymentsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWebchatDeploymentsNotFound creates a GetWebchatDeploymentsNotFound with default headers values
func NewGetWebchatDeploymentsNotFound() *GetWebchatDeploymentsNotFound {
	return &GetWebchatDeploymentsNotFound{}
}

/*GetWebchatDeploymentsNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetWebchatDeploymentsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetWebchatDeploymentsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/webchat/deployments][%d] getWebchatDeploymentsNotFound  %+v", 404, o.Payload)
}

func (o *GetWebchatDeploymentsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWebchatDeploymentsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWebchatDeploymentsRequestEntityTooLarge creates a GetWebchatDeploymentsRequestEntityTooLarge with default headers values
func NewGetWebchatDeploymentsRequestEntityTooLarge() *GetWebchatDeploymentsRequestEntityTooLarge {
	return &GetWebchatDeploymentsRequestEntityTooLarge{}
}

/*GetWebchatDeploymentsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetWebchatDeploymentsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetWebchatDeploymentsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/webchat/deployments][%d] getWebchatDeploymentsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetWebchatDeploymentsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWebchatDeploymentsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWebchatDeploymentsUnsupportedMediaType creates a GetWebchatDeploymentsUnsupportedMediaType with default headers values
func NewGetWebchatDeploymentsUnsupportedMediaType() *GetWebchatDeploymentsUnsupportedMediaType {
	return &GetWebchatDeploymentsUnsupportedMediaType{}
}

/*GetWebchatDeploymentsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetWebchatDeploymentsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetWebchatDeploymentsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/webchat/deployments][%d] getWebchatDeploymentsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetWebchatDeploymentsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWebchatDeploymentsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWebchatDeploymentsTooManyRequests creates a GetWebchatDeploymentsTooManyRequests with default headers values
func NewGetWebchatDeploymentsTooManyRequests() *GetWebchatDeploymentsTooManyRequests {
	return &GetWebchatDeploymentsTooManyRequests{}
}

/*GetWebchatDeploymentsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetWebchatDeploymentsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetWebchatDeploymentsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/webchat/deployments][%d] getWebchatDeploymentsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetWebchatDeploymentsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWebchatDeploymentsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWebchatDeploymentsInternalServerError creates a GetWebchatDeploymentsInternalServerError with default headers values
func NewGetWebchatDeploymentsInternalServerError() *GetWebchatDeploymentsInternalServerError {
	return &GetWebchatDeploymentsInternalServerError{}
}

/*GetWebchatDeploymentsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetWebchatDeploymentsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetWebchatDeploymentsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/webchat/deployments][%d] getWebchatDeploymentsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetWebchatDeploymentsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWebchatDeploymentsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWebchatDeploymentsServiceUnavailable creates a GetWebchatDeploymentsServiceUnavailable with default headers values
func NewGetWebchatDeploymentsServiceUnavailable() *GetWebchatDeploymentsServiceUnavailable {
	return &GetWebchatDeploymentsServiceUnavailable{}
}

/*GetWebchatDeploymentsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetWebchatDeploymentsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetWebchatDeploymentsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/webchat/deployments][%d] getWebchatDeploymentsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetWebchatDeploymentsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWebchatDeploymentsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWebchatDeploymentsGatewayTimeout creates a GetWebchatDeploymentsGatewayTimeout with default headers values
func NewGetWebchatDeploymentsGatewayTimeout() *GetWebchatDeploymentsGatewayTimeout {
	return &GetWebchatDeploymentsGatewayTimeout{}
}

/*GetWebchatDeploymentsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetWebchatDeploymentsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetWebchatDeploymentsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/webchat/deployments][%d] getWebchatDeploymentsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetWebchatDeploymentsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWebchatDeploymentsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

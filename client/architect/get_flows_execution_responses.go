// Code generated by go-swagger; DO NOT EDIT.

package architect

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetFlowsExecutionReader is a Reader for the GetFlowsExecution structure.
type GetFlowsExecutionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFlowsExecutionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFlowsExecutionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetFlowsExecutionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetFlowsExecutionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetFlowsExecutionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetFlowsExecutionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetFlowsExecutionRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetFlowsExecutionRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetFlowsExecutionUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetFlowsExecutionTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetFlowsExecutionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetFlowsExecutionServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetFlowsExecutionGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetFlowsExecutionOK creates a GetFlowsExecutionOK with default headers values
func NewGetFlowsExecutionOK() *GetFlowsExecutionOK {
	return &GetFlowsExecutionOK{}
}

/*GetFlowsExecutionOK handles this case with default header values.

successful operation
*/
type GetFlowsExecutionOK struct {
	Payload *models.FlowRuntimeExecution
}

func (o *GetFlowsExecutionOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/executions/{flowExecutionId}][%d] getFlowsExecutionOK  %+v", 200, o.Payload)
}

func (o *GetFlowsExecutionOK) GetPayload() *models.FlowRuntimeExecution {
	return o.Payload
}

func (o *GetFlowsExecutionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FlowRuntimeExecution)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsExecutionBadRequest creates a GetFlowsExecutionBadRequest with default headers values
func NewGetFlowsExecutionBadRequest() *GetFlowsExecutionBadRequest {
	return &GetFlowsExecutionBadRequest{}
}

/*GetFlowsExecutionBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetFlowsExecutionBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsExecutionBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/executions/{flowExecutionId}][%d] getFlowsExecutionBadRequest  %+v", 400, o.Payload)
}

func (o *GetFlowsExecutionBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsExecutionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsExecutionUnauthorized creates a GetFlowsExecutionUnauthorized with default headers values
func NewGetFlowsExecutionUnauthorized() *GetFlowsExecutionUnauthorized {
	return &GetFlowsExecutionUnauthorized{}
}

/*GetFlowsExecutionUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetFlowsExecutionUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsExecutionUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/executions/{flowExecutionId}][%d] getFlowsExecutionUnauthorized  %+v", 401, o.Payload)
}

func (o *GetFlowsExecutionUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsExecutionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsExecutionForbidden creates a GetFlowsExecutionForbidden with default headers values
func NewGetFlowsExecutionForbidden() *GetFlowsExecutionForbidden {
	return &GetFlowsExecutionForbidden{}
}

/*GetFlowsExecutionForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetFlowsExecutionForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsExecutionForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/executions/{flowExecutionId}][%d] getFlowsExecutionForbidden  %+v", 403, o.Payload)
}

func (o *GetFlowsExecutionForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsExecutionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsExecutionNotFound creates a GetFlowsExecutionNotFound with default headers values
func NewGetFlowsExecutionNotFound() *GetFlowsExecutionNotFound {
	return &GetFlowsExecutionNotFound{}
}

/*GetFlowsExecutionNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetFlowsExecutionNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsExecutionNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/executions/{flowExecutionId}][%d] getFlowsExecutionNotFound  %+v", 404, o.Payload)
}

func (o *GetFlowsExecutionNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsExecutionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsExecutionRequestTimeout creates a GetFlowsExecutionRequestTimeout with default headers values
func NewGetFlowsExecutionRequestTimeout() *GetFlowsExecutionRequestTimeout {
	return &GetFlowsExecutionRequestTimeout{}
}

/*GetFlowsExecutionRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetFlowsExecutionRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsExecutionRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/executions/{flowExecutionId}][%d] getFlowsExecutionRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetFlowsExecutionRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsExecutionRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsExecutionRequestEntityTooLarge creates a GetFlowsExecutionRequestEntityTooLarge with default headers values
func NewGetFlowsExecutionRequestEntityTooLarge() *GetFlowsExecutionRequestEntityTooLarge {
	return &GetFlowsExecutionRequestEntityTooLarge{}
}

/*GetFlowsExecutionRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetFlowsExecutionRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsExecutionRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/executions/{flowExecutionId}][%d] getFlowsExecutionRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetFlowsExecutionRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsExecutionRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsExecutionUnsupportedMediaType creates a GetFlowsExecutionUnsupportedMediaType with default headers values
func NewGetFlowsExecutionUnsupportedMediaType() *GetFlowsExecutionUnsupportedMediaType {
	return &GetFlowsExecutionUnsupportedMediaType{}
}

/*GetFlowsExecutionUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetFlowsExecutionUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsExecutionUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/executions/{flowExecutionId}][%d] getFlowsExecutionUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetFlowsExecutionUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsExecutionUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsExecutionTooManyRequests creates a GetFlowsExecutionTooManyRequests with default headers values
func NewGetFlowsExecutionTooManyRequests() *GetFlowsExecutionTooManyRequests {
	return &GetFlowsExecutionTooManyRequests{}
}

/*GetFlowsExecutionTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetFlowsExecutionTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsExecutionTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/executions/{flowExecutionId}][%d] getFlowsExecutionTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetFlowsExecutionTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsExecutionTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsExecutionInternalServerError creates a GetFlowsExecutionInternalServerError with default headers values
func NewGetFlowsExecutionInternalServerError() *GetFlowsExecutionInternalServerError {
	return &GetFlowsExecutionInternalServerError{}
}

/*GetFlowsExecutionInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetFlowsExecutionInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsExecutionInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/executions/{flowExecutionId}][%d] getFlowsExecutionInternalServerError  %+v", 500, o.Payload)
}

func (o *GetFlowsExecutionInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsExecutionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsExecutionServiceUnavailable creates a GetFlowsExecutionServiceUnavailable with default headers values
func NewGetFlowsExecutionServiceUnavailable() *GetFlowsExecutionServiceUnavailable {
	return &GetFlowsExecutionServiceUnavailable{}
}

/*GetFlowsExecutionServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetFlowsExecutionServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsExecutionServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/executions/{flowExecutionId}][%d] getFlowsExecutionServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetFlowsExecutionServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsExecutionServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsExecutionGatewayTimeout creates a GetFlowsExecutionGatewayTimeout with default headers values
func NewGetFlowsExecutionGatewayTimeout() *GetFlowsExecutionGatewayTimeout {
	return &GetFlowsExecutionGatewayTimeout{}
}

/*GetFlowsExecutionGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetFlowsExecutionGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsExecutionGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/executions/{flowExecutionId}][%d] getFlowsExecutionGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetFlowsExecutionGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsExecutionGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

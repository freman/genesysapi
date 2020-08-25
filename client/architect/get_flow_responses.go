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

// GetFlowReader is a Reader for the GetFlow structure.
type GetFlowReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFlowReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFlowOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetFlowBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetFlowUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetFlowForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetFlowNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewGetFlowMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewGetFlowGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetFlowRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetFlowUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetFlowTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetFlowInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetFlowServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetFlowGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetFlowOK creates a GetFlowOK with default headers values
func NewGetFlowOK() *GetFlowOK {
	return &GetFlowOK{}
}

/*GetFlowOK handles this case with default header values.

successful operation
*/
type GetFlowOK struct {
	Payload *models.Flow
}

func (o *GetFlowOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/{flowId}][%d] getFlowOK  %+v", 200, o.Payload)
}

func (o *GetFlowOK) GetPayload() *models.Flow {
	return o.Payload
}

func (o *GetFlowOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Flow)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowBadRequest creates a GetFlowBadRequest with default headers values
func NewGetFlowBadRequest() *GetFlowBadRequest {
	return &GetFlowBadRequest{}
}

/*GetFlowBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetFlowBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetFlowBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/{flowId}][%d] getFlowBadRequest  %+v", 400, o.Payload)
}

func (o *GetFlowBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowUnauthorized creates a GetFlowUnauthorized with default headers values
func NewGetFlowUnauthorized() *GetFlowUnauthorized {
	return &GetFlowUnauthorized{}
}

/*GetFlowUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetFlowUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetFlowUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/{flowId}][%d] getFlowUnauthorized  %+v", 401, o.Payload)
}

func (o *GetFlowUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowForbidden creates a GetFlowForbidden with default headers values
func NewGetFlowForbidden() *GetFlowForbidden {
	return &GetFlowForbidden{}
}

/*GetFlowForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetFlowForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetFlowForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/{flowId}][%d] getFlowForbidden  %+v", 403, o.Payload)
}

func (o *GetFlowForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowNotFound creates a GetFlowNotFound with default headers values
func NewGetFlowNotFound() *GetFlowNotFound {
	return &GetFlowNotFound{}
}

/*GetFlowNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetFlowNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetFlowNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/{flowId}][%d] getFlowNotFound  %+v", 404, o.Payload)
}

func (o *GetFlowNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowMethodNotAllowed creates a GetFlowMethodNotAllowed with default headers values
func NewGetFlowMethodNotAllowed() *GetFlowMethodNotAllowed {
	return &GetFlowMethodNotAllowed{}
}

/*GetFlowMethodNotAllowed handles this case with default header values.

Method Not Allowed
*/
type GetFlowMethodNotAllowed struct {
	Payload *models.ErrorBody
}

func (o *GetFlowMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/{flowId}][%d] getFlowMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *GetFlowMethodNotAllowed) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowGone creates a GetFlowGone with default headers values
func NewGetFlowGone() *GetFlowGone {
	return &GetFlowGone{}
}

/*GetFlowGone handles this case with default header values.

Gone
*/
type GetFlowGone struct {
	Payload *models.ErrorBody
}

func (o *GetFlowGone) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/{flowId}][%d] getFlowGone  %+v", 410, o.Payload)
}

func (o *GetFlowGone) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowRequestEntityTooLarge creates a GetFlowRequestEntityTooLarge with default headers values
func NewGetFlowRequestEntityTooLarge() *GetFlowRequestEntityTooLarge {
	return &GetFlowRequestEntityTooLarge{}
}

/*GetFlowRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetFlowRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetFlowRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/{flowId}][%d] getFlowRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetFlowRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowUnsupportedMediaType creates a GetFlowUnsupportedMediaType with default headers values
func NewGetFlowUnsupportedMediaType() *GetFlowUnsupportedMediaType {
	return &GetFlowUnsupportedMediaType{}
}

/*GetFlowUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetFlowUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetFlowUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/{flowId}][%d] getFlowUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetFlowUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowTooManyRequests creates a GetFlowTooManyRequests with default headers values
func NewGetFlowTooManyRequests() *GetFlowTooManyRequests {
	return &GetFlowTooManyRequests{}
}

/*GetFlowTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetFlowTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetFlowTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/{flowId}][%d] getFlowTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetFlowTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowInternalServerError creates a GetFlowInternalServerError with default headers values
func NewGetFlowInternalServerError() *GetFlowInternalServerError {
	return &GetFlowInternalServerError{}
}

/*GetFlowInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetFlowInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetFlowInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/{flowId}][%d] getFlowInternalServerError  %+v", 500, o.Payload)
}

func (o *GetFlowInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowServiceUnavailable creates a GetFlowServiceUnavailable with default headers values
func NewGetFlowServiceUnavailable() *GetFlowServiceUnavailable {
	return &GetFlowServiceUnavailable{}
}

/*GetFlowServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetFlowServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetFlowServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/{flowId}][%d] getFlowServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetFlowServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowGatewayTimeout creates a GetFlowGatewayTimeout with default headers values
func NewGetFlowGatewayTimeout() *GetFlowGatewayTimeout {
	return &GetFlowGatewayTimeout{}
}

/*GetFlowGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetFlowGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetFlowGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/{flowId}][%d] getFlowGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetFlowGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
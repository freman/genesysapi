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

// GetFlowsReader is a Reader for the GetFlows structure.
type GetFlowsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFlowsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFlowsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetFlowsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetFlowsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetFlowsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetFlowsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewGetFlowsMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetFlowsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetFlowsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetFlowsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetFlowsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetFlowsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetFlowsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetFlowsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetFlowsOK creates a GetFlowsOK with default headers values
func NewGetFlowsOK() *GetFlowsOK {
	return &GetFlowsOK{}
}

/*GetFlowsOK handles this case with default header values.

successful operation
*/
type GetFlowsOK struct {
	Payload *models.FlowEntityListing
}

func (o *GetFlowsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows][%d] getFlowsOK  %+v", 200, o.Payload)
}

func (o *GetFlowsOK) GetPayload() *models.FlowEntityListing {
	return o.Payload
}

func (o *GetFlowsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FlowEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsBadRequest creates a GetFlowsBadRequest with default headers values
func NewGetFlowsBadRequest() *GetFlowsBadRequest {
	return &GetFlowsBadRequest{}
}

/*GetFlowsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetFlowsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows][%d] getFlowsBadRequest  %+v", 400, o.Payload)
}

func (o *GetFlowsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsUnauthorized creates a GetFlowsUnauthorized with default headers values
func NewGetFlowsUnauthorized() *GetFlowsUnauthorized {
	return &GetFlowsUnauthorized{}
}

/*GetFlowsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetFlowsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows][%d] getFlowsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetFlowsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsForbidden creates a GetFlowsForbidden with default headers values
func NewGetFlowsForbidden() *GetFlowsForbidden {
	return &GetFlowsForbidden{}
}

/*GetFlowsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetFlowsForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows][%d] getFlowsForbidden  %+v", 403, o.Payload)
}

func (o *GetFlowsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsNotFound creates a GetFlowsNotFound with default headers values
func NewGetFlowsNotFound() *GetFlowsNotFound {
	return &GetFlowsNotFound{}
}

/*GetFlowsNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetFlowsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows][%d] getFlowsNotFound  %+v", 404, o.Payload)
}

func (o *GetFlowsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsMethodNotAllowed creates a GetFlowsMethodNotAllowed with default headers values
func NewGetFlowsMethodNotAllowed() *GetFlowsMethodNotAllowed {
	return &GetFlowsMethodNotAllowed{}
}

/*GetFlowsMethodNotAllowed handles this case with default header values.

Method Not Allowed
*/
type GetFlowsMethodNotAllowed struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows][%d] getFlowsMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *GetFlowsMethodNotAllowed) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsRequestTimeout creates a GetFlowsRequestTimeout with default headers values
func NewGetFlowsRequestTimeout() *GetFlowsRequestTimeout {
	return &GetFlowsRequestTimeout{}
}

/*GetFlowsRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetFlowsRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows][%d] getFlowsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetFlowsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsRequestEntityTooLarge creates a GetFlowsRequestEntityTooLarge with default headers values
func NewGetFlowsRequestEntityTooLarge() *GetFlowsRequestEntityTooLarge {
	return &GetFlowsRequestEntityTooLarge{}
}

/*GetFlowsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetFlowsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows][%d] getFlowsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetFlowsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsUnsupportedMediaType creates a GetFlowsUnsupportedMediaType with default headers values
func NewGetFlowsUnsupportedMediaType() *GetFlowsUnsupportedMediaType {
	return &GetFlowsUnsupportedMediaType{}
}

/*GetFlowsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetFlowsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows][%d] getFlowsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetFlowsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsTooManyRequests creates a GetFlowsTooManyRequests with default headers values
func NewGetFlowsTooManyRequests() *GetFlowsTooManyRequests {
	return &GetFlowsTooManyRequests{}
}

/*GetFlowsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetFlowsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows][%d] getFlowsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetFlowsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsInternalServerError creates a GetFlowsInternalServerError with default headers values
func NewGetFlowsInternalServerError() *GetFlowsInternalServerError {
	return &GetFlowsInternalServerError{}
}

/*GetFlowsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetFlowsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows][%d] getFlowsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetFlowsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsServiceUnavailable creates a GetFlowsServiceUnavailable with default headers values
func NewGetFlowsServiceUnavailable() *GetFlowsServiceUnavailable {
	return &GetFlowsServiceUnavailable{}
}

/*GetFlowsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetFlowsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows][%d] getFlowsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetFlowsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsGatewayTimeout creates a GetFlowsGatewayTimeout with default headers values
func NewGetFlowsGatewayTimeout() *GetFlowsGatewayTimeout {
	return &GetFlowsGatewayTimeout{}
}

/*GetFlowsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetFlowsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows][%d] getFlowsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetFlowsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

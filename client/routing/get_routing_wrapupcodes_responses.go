// Code generated by go-swagger; DO NOT EDIT.

package routing

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetRoutingWrapupcodesReader is a Reader for the GetRoutingWrapupcodes structure.
type GetRoutingWrapupcodesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRoutingWrapupcodesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRoutingWrapupcodesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRoutingWrapupcodesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetRoutingWrapupcodesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRoutingWrapupcodesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRoutingWrapupcodesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetRoutingWrapupcodesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetRoutingWrapupcodesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetRoutingWrapupcodesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRoutingWrapupcodesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetRoutingWrapupcodesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetRoutingWrapupcodesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRoutingWrapupcodesOK creates a GetRoutingWrapupcodesOK with default headers values
func NewGetRoutingWrapupcodesOK() *GetRoutingWrapupcodesOK {
	return &GetRoutingWrapupcodesOK{}
}

/*GetRoutingWrapupcodesOK handles this case with default header values.

successful operation
*/
type GetRoutingWrapupcodesOK struct {
	Payload *models.WrapupCodeEntityListing
}

func (o *GetRoutingWrapupcodesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/wrapupcodes][%d] getRoutingWrapupcodesOK  %+v", 200, o.Payload)
}

func (o *GetRoutingWrapupcodesOK) GetPayload() *models.WrapupCodeEntityListing {
	return o.Payload
}

func (o *GetRoutingWrapupcodesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WrapupCodeEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingWrapupcodesBadRequest creates a GetRoutingWrapupcodesBadRequest with default headers values
func NewGetRoutingWrapupcodesBadRequest() *GetRoutingWrapupcodesBadRequest {
	return &GetRoutingWrapupcodesBadRequest{}
}

/*GetRoutingWrapupcodesBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetRoutingWrapupcodesBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingWrapupcodesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/wrapupcodes][%d] getRoutingWrapupcodesBadRequest  %+v", 400, o.Payload)
}

func (o *GetRoutingWrapupcodesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingWrapupcodesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingWrapupcodesUnauthorized creates a GetRoutingWrapupcodesUnauthorized with default headers values
func NewGetRoutingWrapupcodesUnauthorized() *GetRoutingWrapupcodesUnauthorized {
	return &GetRoutingWrapupcodesUnauthorized{}
}

/*GetRoutingWrapupcodesUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetRoutingWrapupcodesUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingWrapupcodesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/wrapupcodes][%d] getRoutingWrapupcodesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRoutingWrapupcodesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingWrapupcodesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingWrapupcodesForbidden creates a GetRoutingWrapupcodesForbidden with default headers values
func NewGetRoutingWrapupcodesForbidden() *GetRoutingWrapupcodesForbidden {
	return &GetRoutingWrapupcodesForbidden{}
}

/*GetRoutingWrapupcodesForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetRoutingWrapupcodesForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingWrapupcodesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/wrapupcodes][%d] getRoutingWrapupcodesForbidden  %+v", 403, o.Payload)
}

func (o *GetRoutingWrapupcodesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingWrapupcodesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingWrapupcodesNotFound creates a GetRoutingWrapupcodesNotFound with default headers values
func NewGetRoutingWrapupcodesNotFound() *GetRoutingWrapupcodesNotFound {
	return &GetRoutingWrapupcodesNotFound{}
}

/*GetRoutingWrapupcodesNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetRoutingWrapupcodesNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingWrapupcodesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/wrapupcodes][%d] getRoutingWrapupcodesNotFound  %+v", 404, o.Payload)
}

func (o *GetRoutingWrapupcodesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingWrapupcodesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingWrapupcodesRequestEntityTooLarge creates a GetRoutingWrapupcodesRequestEntityTooLarge with default headers values
func NewGetRoutingWrapupcodesRequestEntityTooLarge() *GetRoutingWrapupcodesRequestEntityTooLarge {
	return &GetRoutingWrapupcodesRequestEntityTooLarge{}
}

/*GetRoutingWrapupcodesRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetRoutingWrapupcodesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingWrapupcodesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/wrapupcodes][%d] getRoutingWrapupcodesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetRoutingWrapupcodesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingWrapupcodesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingWrapupcodesUnsupportedMediaType creates a GetRoutingWrapupcodesUnsupportedMediaType with default headers values
func NewGetRoutingWrapupcodesUnsupportedMediaType() *GetRoutingWrapupcodesUnsupportedMediaType {
	return &GetRoutingWrapupcodesUnsupportedMediaType{}
}

/*GetRoutingWrapupcodesUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetRoutingWrapupcodesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingWrapupcodesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/wrapupcodes][%d] getRoutingWrapupcodesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetRoutingWrapupcodesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingWrapupcodesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingWrapupcodesTooManyRequests creates a GetRoutingWrapupcodesTooManyRequests with default headers values
func NewGetRoutingWrapupcodesTooManyRequests() *GetRoutingWrapupcodesTooManyRequests {
	return &GetRoutingWrapupcodesTooManyRequests{}
}

/*GetRoutingWrapupcodesTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetRoutingWrapupcodesTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingWrapupcodesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/wrapupcodes][%d] getRoutingWrapupcodesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetRoutingWrapupcodesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingWrapupcodesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingWrapupcodesInternalServerError creates a GetRoutingWrapupcodesInternalServerError with default headers values
func NewGetRoutingWrapupcodesInternalServerError() *GetRoutingWrapupcodesInternalServerError {
	return &GetRoutingWrapupcodesInternalServerError{}
}

/*GetRoutingWrapupcodesInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetRoutingWrapupcodesInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingWrapupcodesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/wrapupcodes][%d] getRoutingWrapupcodesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRoutingWrapupcodesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingWrapupcodesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingWrapupcodesServiceUnavailable creates a GetRoutingWrapupcodesServiceUnavailable with default headers values
func NewGetRoutingWrapupcodesServiceUnavailable() *GetRoutingWrapupcodesServiceUnavailable {
	return &GetRoutingWrapupcodesServiceUnavailable{}
}

/*GetRoutingWrapupcodesServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetRoutingWrapupcodesServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingWrapupcodesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/wrapupcodes][%d] getRoutingWrapupcodesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetRoutingWrapupcodesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingWrapupcodesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingWrapupcodesGatewayTimeout creates a GetRoutingWrapupcodesGatewayTimeout with default headers values
func NewGetRoutingWrapupcodesGatewayTimeout() *GetRoutingWrapupcodesGatewayTimeout {
	return &GetRoutingWrapupcodesGatewayTimeout{}
}

/*GetRoutingWrapupcodesGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetRoutingWrapupcodesGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingWrapupcodesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/wrapupcodes][%d] getRoutingWrapupcodesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetRoutingWrapupcodesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingWrapupcodesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

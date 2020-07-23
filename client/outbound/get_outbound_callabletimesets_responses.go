// Code generated by go-swagger; DO NOT EDIT.

package outbound

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetOutboundCallabletimesetsReader is a Reader for the GetOutboundCallabletimesets structure.
type GetOutboundCallabletimesetsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOutboundCallabletimesetsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOutboundCallabletimesetsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOutboundCallabletimesetsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetOutboundCallabletimesetsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOutboundCallabletimesetsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOutboundCallabletimesetsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetOutboundCallabletimesetsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetOutboundCallabletimesetsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetOutboundCallabletimesetsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOutboundCallabletimesetsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetOutboundCallabletimesetsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetOutboundCallabletimesetsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOutboundCallabletimesetsOK creates a GetOutboundCallabletimesetsOK with default headers values
func NewGetOutboundCallabletimesetsOK() *GetOutboundCallabletimesetsOK {
	return &GetOutboundCallabletimesetsOK{}
}

/*GetOutboundCallabletimesetsOK handles this case with default header values.

successful operation
*/
type GetOutboundCallabletimesetsOK struct {
	Payload *models.CallableTimeSetEntityListing
}

func (o *GetOutboundCallabletimesetsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/callabletimesets][%d] getOutboundCallabletimesetsOK  %+v", 200, o.Payload)
}

func (o *GetOutboundCallabletimesetsOK) GetPayload() *models.CallableTimeSetEntityListing {
	return o.Payload
}

func (o *GetOutboundCallabletimesetsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CallableTimeSetEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCallabletimesetsBadRequest creates a GetOutboundCallabletimesetsBadRequest with default headers values
func NewGetOutboundCallabletimesetsBadRequest() *GetOutboundCallabletimesetsBadRequest {
	return &GetOutboundCallabletimesetsBadRequest{}
}

/*GetOutboundCallabletimesetsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetOutboundCallabletimesetsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCallabletimesetsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/callabletimesets][%d] getOutboundCallabletimesetsBadRequest  %+v", 400, o.Payload)
}

func (o *GetOutboundCallabletimesetsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCallabletimesetsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCallabletimesetsUnauthorized creates a GetOutboundCallabletimesetsUnauthorized with default headers values
func NewGetOutboundCallabletimesetsUnauthorized() *GetOutboundCallabletimesetsUnauthorized {
	return &GetOutboundCallabletimesetsUnauthorized{}
}

/*GetOutboundCallabletimesetsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetOutboundCallabletimesetsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCallabletimesetsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/callabletimesets][%d] getOutboundCallabletimesetsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOutboundCallabletimesetsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCallabletimesetsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCallabletimesetsForbidden creates a GetOutboundCallabletimesetsForbidden with default headers values
func NewGetOutboundCallabletimesetsForbidden() *GetOutboundCallabletimesetsForbidden {
	return &GetOutboundCallabletimesetsForbidden{}
}

/*GetOutboundCallabletimesetsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetOutboundCallabletimesetsForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCallabletimesetsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/callabletimesets][%d] getOutboundCallabletimesetsForbidden  %+v", 403, o.Payload)
}

func (o *GetOutboundCallabletimesetsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCallabletimesetsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCallabletimesetsNotFound creates a GetOutboundCallabletimesetsNotFound with default headers values
func NewGetOutboundCallabletimesetsNotFound() *GetOutboundCallabletimesetsNotFound {
	return &GetOutboundCallabletimesetsNotFound{}
}

/*GetOutboundCallabletimesetsNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetOutboundCallabletimesetsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCallabletimesetsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/callabletimesets][%d] getOutboundCallabletimesetsNotFound  %+v", 404, o.Payload)
}

func (o *GetOutboundCallabletimesetsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCallabletimesetsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCallabletimesetsRequestEntityTooLarge creates a GetOutboundCallabletimesetsRequestEntityTooLarge with default headers values
func NewGetOutboundCallabletimesetsRequestEntityTooLarge() *GetOutboundCallabletimesetsRequestEntityTooLarge {
	return &GetOutboundCallabletimesetsRequestEntityTooLarge{}
}

/*GetOutboundCallabletimesetsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetOutboundCallabletimesetsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCallabletimesetsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/callabletimesets][%d] getOutboundCallabletimesetsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetOutboundCallabletimesetsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCallabletimesetsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCallabletimesetsUnsupportedMediaType creates a GetOutboundCallabletimesetsUnsupportedMediaType with default headers values
func NewGetOutboundCallabletimesetsUnsupportedMediaType() *GetOutboundCallabletimesetsUnsupportedMediaType {
	return &GetOutboundCallabletimesetsUnsupportedMediaType{}
}

/*GetOutboundCallabletimesetsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetOutboundCallabletimesetsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCallabletimesetsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/callabletimesets][%d] getOutboundCallabletimesetsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetOutboundCallabletimesetsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCallabletimesetsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCallabletimesetsTooManyRequests creates a GetOutboundCallabletimesetsTooManyRequests with default headers values
func NewGetOutboundCallabletimesetsTooManyRequests() *GetOutboundCallabletimesetsTooManyRequests {
	return &GetOutboundCallabletimesetsTooManyRequests{}
}

/*GetOutboundCallabletimesetsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetOutboundCallabletimesetsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCallabletimesetsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/callabletimesets][%d] getOutboundCallabletimesetsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOutboundCallabletimesetsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCallabletimesetsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCallabletimesetsInternalServerError creates a GetOutboundCallabletimesetsInternalServerError with default headers values
func NewGetOutboundCallabletimesetsInternalServerError() *GetOutboundCallabletimesetsInternalServerError {
	return &GetOutboundCallabletimesetsInternalServerError{}
}

/*GetOutboundCallabletimesetsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetOutboundCallabletimesetsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCallabletimesetsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/callabletimesets][%d] getOutboundCallabletimesetsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOutboundCallabletimesetsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCallabletimesetsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCallabletimesetsServiceUnavailable creates a GetOutboundCallabletimesetsServiceUnavailable with default headers values
func NewGetOutboundCallabletimesetsServiceUnavailable() *GetOutboundCallabletimesetsServiceUnavailable {
	return &GetOutboundCallabletimesetsServiceUnavailable{}
}

/*GetOutboundCallabletimesetsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetOutboundCallabletimesetsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCallabletimesetsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/callabletimesets][%d] getOutboundCallabletimesetsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetOutboundCallabletimesetsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCallabletimesetsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCallabletimesetsGatewayTimeout creates a GetOutboundCallabletimesetsGatewayTimeout with default headers values
func NewGetOutboundCallabletimesetsGatewayTimeout() *GetOutboundCallabletimesetsGatewayTimeout {
	return &GetOutboundCallabletimesetsGatewayTimeout{}
}

/*GetOutboundCallabletimesetsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetOutboundCallabletimesetsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCallabletimesetsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/callabletimesets][%d] getOutboundCallabletimesetsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetOutboundCallabletimesetsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCallabletimesetsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

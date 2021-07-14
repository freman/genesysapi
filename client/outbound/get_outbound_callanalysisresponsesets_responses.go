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

// GetOutboundCallanalysisresponsesetsReader is a Reader for the GetOutboundCallanalysisresponsesets structure.
type GetOutboundCallanalysisresponsesetsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOutboundCallanalysisresponsesetsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOutboundCallanalysisresponsesetsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOutboundCallanalysisresponsesetsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetOutboundCallanalysisresponsesetsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOutboundCallanalysisresponsesetsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOutboundCallanalysisresponsesetsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetOutboundCallanalysisresponsesetsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetOutboundCallanalysisresponsesetsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetOutboundCallanalysisresponsesetsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetOutboundCallanalysisresponsesetsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOutboundCallanalysisresponsesetsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetOutboundCallanalysisresponsesetsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetOutboundCallanalysisresponsesetsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOutboundCallanalysisresponsesetsOK creates a GetOutboundCallanalysisresponsesetsOK with default headers values
func NewGetOutboundCallanalysisresponsesetsOK() *GetOutboundCallanalysisresponsesetsOK {
	return &GetOutboundCallanalysisresponsesetsOK{}
}

/*GetOutboundCallanalysisresponsesetsOK handles this case with default header values.

successful operation
*/
type GetOutboundCallanalysisresponsesetsOK struct {
	Payload *models.ResponseSetEntityListing
}

func (o *GetOutboundCallanalysisresponsesetsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/callanalysisresponsesets][%d] getOutboundCallanalysisresponsesetsOK  %+v", 200, o.Payload)
}

func (o *GetOutboundCallanalysisresponsesetsOK) GetPayload() *models.ResponseSetEntityListing {
	return o.Payload
}

func (o *GetOutboundCallanalysisresponsesetsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseSetEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCallanalysisresponsesetsBadRequest creates a GetOutboundCallanalysisresponsesetsBadRequest with default headers values
func NewGetOutboundCallanalysisresponsesetsBadRequest() *GetOutboundCallanalysisresponsesetsBadRequest {
	return &GetOutboundCallanalysisresponsesetsBadRequest{}
}

/*GetOutboundCallanalysisresponsesetsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetOutboundCallanalysisresponsesetsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCallanalysisresponsesetsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/callanalysisresponsesets][%d] getOutboundCallanalysisresponsesetsBadRequest  %+v", 400, o.Payload)
}

func (o *GetOutboundCallanalysisresponsesetsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCallanalysisresponsesetsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCallanalysisresponsesetsUnauthorized creates a GetOutboundCallanalysisresponsesetsUnauthorized with default headers values
func NewGetOutboundCallanalysisresponsesetsUnauthorized() *GetOutboundCallanalysisresponsesetsUnauthorized {
	return &GetOutboundCallanalysisresponsesetsUnauthorized{}
}

/*GetOutboundCallanalysisresponsesetsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetOutboundCallanalysisresponsesetsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCallanalysisresponsesetsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/callanalysisresponsesets][%d] getOutboundCallanalysisresponsesetsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOutboundCallanalysisresponsesetsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCallanalysisresponsesetsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCallanalysisresponsesetsForbidden creates a GetOutboundCallanalysisresponsesetsForbidden with default headers values
func NewGetOutboundCallanalysisresponsesetsForbidden() *GetOutboundCallanalysisresponsesetsForbidden {
	return &GetOutboundCallanalysisresponsesetsForbidden{}
}

/*GetOutboundCallanalysisresponsesetsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetOutboundCallanalysisresponsesetsForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCallanalysisresponsesetsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/callanalysisresponsesets][%d] getOutboundCallanalysisresponsesetsForbidden  %+v", 403, o.Payload)
}

func (o *GetOutboundCallanalysisresponsesetsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCallanalysisresponsesetsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCallanalysisresponsesetsNotFound creates a GetOutboundCallanalysisresponsesetsNotFound with default headers values
func NewGetOutboundCallanalysisresponsesetsNotFound() *GetOutboundCallanalysisresponsesetsNotFound {
	return &GetOutboundCallanalysisresponsesetsNotFound{}
}

/*GetOutboundCallanalysisresponsesetsNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetOutboundCallanalysisresponsesetsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCallanalysisresponsesetsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/callanalysisresponsesets][%d] getOutboundCallanalysisresponsesetsNotFound  %+v", 404, o.Payload)
}

func (o *GetOutboundCallanalysisresponsesetsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCallanalysisresponsesetsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCallanalysisresponsesetsRequestTimeout creates a GetOutboundCallanalysisresponsesetsRequestTimeout with default headers values
func NewGetOutboundCallanalysisresponsesetsRequestTimeout() *GetOutboundCallanalysisresponsesetsRequestTimeout {
	return &GetOutboundCallanalysisresponsesetsRequestTimeout{}
}

/*GetOutboundCallanalysisresponsesetsRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetOutboundCallanalysisresponsesetsRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCallanalysisresponsesetsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/callanalysisresponsesets][%d] getOutboundCallanalysisresponsesetsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetOutboundCallanalysisresponsesetsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCallanalysisresponsesetsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCallanalysisresponsesetsRequestEntityTooLarge creates a GetOutboundCallanalysisresponsesetsRequestEntityTooLarge with default headers values
func NewGetOutboundCallanalysisresponsesetsRequestEntityTooLarge() *GetOutboundCallanalysisresponsesetsRequestEntityTooLarge {
	return &GetOutboundCallanalysisresponsesetsRequestEntityTooLarge{}
}

/*GetOutboundCallanalysisresponsesetsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetOutboundCallanalysisresponsesetsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCallanalysisresponsesetsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/callanalysisresponsesets][%d] getOutboundCallanalysisresponsesetsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetOutboundCallanalysisresponsesetsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCallanalysisresponsesetsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCallanalysisresponsesetsUnsupportedMediaType creates a GetOutboundCallanalysisresponsesetsUnsupportedMediaType with default headers values
func NewGetOutboundCallanalysisresponsesetsUnsupportedMediaType() *GetOutboundCallanalysisresponsesetsUnsupportedMediaType {
	return &GetOutboundCallanalysisresponsesetsUnsupportedMediaType{}
}

/*GetOutboundCallanalysisresponsesetsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetOutboundCallanalysisresponsesetsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCallanalysisresponsesetsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/callanalysisresponsesets][%d] getOutboundCallanalysisresponsesetsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetOutboundCallanalysisresponsesetsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCallanalysisresponsesetsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCallanalysisresponsesetsTooManyRequests creates a GetOutboundCallanalysisresponsesetsTooManyRequests with default headers values
func NewGetOutboundCallanalysisresponsesetsTooManyRequests() *GetOutboundCallanalysisresponsesetsTooManyRequests {
	return &GetOutboundCallanalysisresponsesetsTooManyRequests{}
}

/*GetOutboundCallanalysisresponsesetsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetOutboundCallanalysisresponsesetsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCallanalysisresponsesetsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/callanalysisresponsesets][%d] getOutboundCallanalysisresponsesetsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOutboundCallanalysisresponsesetsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCallanalysisresponsesetsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCallanalysisresponsesetsInternalServerError creates a GetOutboundCallanalysisresponsesetsInternalServerError with default headers values
func NewGetOutboundCallanalysisresponsesetsInternalServerError() *GetOutboundCallanalysisresponsesetsInternalServerError {
	return &GetOutboundCallanalysisresponsesetsInternalServerError{}
}

/*GetOutboundCallanalysisresponsesetsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetOutboundCallanalysisresponsesetsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCallanalysisresponsesetsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/callanalysisresponsesets][%d] getOutboundCallanalysisresponsesetsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOutboundCallanalysisresponsesetsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCallanalysisresponsesetsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCallanalysisresponsesetsServiceUnavailable creates a GetOutboundCallanalysisresponsesetsServiceUnavailable with default headers values
func NewGetOutboundCallanalysisresponsesetsServiceUnavailable() *GetOutboundCallanalysisresponsesetsServiceUnavailable {
	return &GetOutboundCallanalysisresponsesetsServiceUnavailable{}
}

/*GetOutboundCallanalysisresponsesetsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetOutboundCallanalysisresponsesetsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCallanalysisresponsesetsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/callanalysisresponsesets][%d] getOutboundCallanalysisresponsesetsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetOutboundCallanalysisresponsesetsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCallanalysisresponsesetsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCallanalysisresponsesetsGatewayTimeout creates a GetOutboundCallanalysisresponsesetsGatewayTimeout with default headers values
func NewGetOutboundCallanalysisresponsesetsGatewayTimeout() *GetOutboundCallanalysisresponsesetsGatewayTimeout {
	return &GetOutboundCallanalysisresponsesetsGatewayTimeout{}
}

/*GetOutboundCallanalysisresponsesetsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetOutboundCallanalysisresponsesetsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCallanalysisresponsesetsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/callanalysisresponsesets][%d] getOutboundCallanalysisresponsesetsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetOutboundCallanalysisresponsesetsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCallanalysisresponsesetsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

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

// GetRoutingEmailDomainRoutesReader is a Reader for the GetRoutingEmailDomainRoutes structure.
type GetRoutingEmailDomainRoutesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRoutingEmailDomainRoutesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRoutingEmailDomainRoutesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRoutingEmailDomainRoutesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetRoutingEmailDomainRoutesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRoutingEmailDomainRoutesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRoutingEmailDomainRoutesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetRoutingEmailDomainRoutesRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetRoutingEmailDomainRoutesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetRoutingEmailDomainRoutesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetRoutingEmailDomainRoutesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRoutingEmailDomainRoutesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetRoutingEmailDomainRoutesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetRoutingEmailDomainRoutesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRoutingEmailDomainRoutesOK creates a GetRoutingEmailDomainRoutesOK with default headers values
func NewGetRoutingEmailDomainRoutesOK() *GetRoutingEmailDomainRoutesOK {
	return &GetRoutingEmailDomainRoutesOK{}
}

/*GetRoutingEmailDomainRoutesOK handles this case with default header values.

successful operation
*/
type GetRoutingEmailDomainRoutesOK struct {
	Payload *models.InboundRouteEntityListing
}

func (o *GetRoutingEmailDomainRoutesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/domains/{domainName}/routes][%d] getRoutingEmailDomainRoutesOK  %+v", 200, o.Payload)
}

func (o *GetRoutingEmailDomainRoutesOK) GetPayload() *models.InboundRouteEntityListing {
	return o.Payload
}

func (o *GetRoutingEmailDomainRoutesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InboundRouteEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingEmailDomainRoutesBadRequest creates a GetRoutingEmailDomainRoutesBadRequest with default headers values
func NewGetRoutingEmailDomainRoutesBadRequest() *GetRoutingEmailDomainRoutesBadRequest {
	return &GetRoutingEmailDomainRoutesBadRequest{}
}

/*GetRoutingEmailDomainRoutesBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetRoutingEmailDomainRoutesBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingEmailDomainRoutesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/domains/{domainName}/routes][%d] getRoutingEmailDomainRoutesBadRequest  %+v", 400, o.Payload)
}

func (o *GetRoutingEmailDomainRoutesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingEmailDomainRoutesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingEmailDomainRoutesUnauthorized creates a GetRoutingEmailDomainRoutesUnauthorized with default headers values
func NewGetRoutingEmailDomainRoutesUnauthorized() *GetRoutingEmailDomainRoutesUnauthorized {
	return &GetRoutingEmailDomainRoutesUnauthorized{}
}

/*GetRoutingEmailDomainRoutesUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetRoutingEmailDomainRoutesUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingEmailDomainRoutesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/domains/{domainName}/routes][%d] getRoutingEmailDomainRoutesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRoutingEmailDomainRoutesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingEmailDomainRoutesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingEmailDomainRoutesForbidden creates a GetRoutingEmailDomainRoutesForbidden with default headers values
func NewGetRoutingEmailDomainRoutesForbidden() *GetRoutingEmailDomainRoutesForbidden {
	return &GetRoutingEmailDomainRoutesForbidden{}
}

/*GetRoutingEmailDomainRoutesForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetRoutingEmailDomainRoutesForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingEmailDomainRoutesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/domains/{domainName}/routes][%d] getRoutingEmailDomainRoutesForbidden  %+v", 403, o.Payload)
}

func (o *GetRoutingEmailDomainRoutesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingEmailDomainRoutesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingEmailDomainRoutesNotFound creates a GetRoutingEmailDomainRoutesNotFound with default headers values
func NewGetRoutingEmailDomainRoutesNotFound() *GetRoutingEmailDomainRoutesNotFound {
	return &GetRoutingEmailDomainRoutesNotFound{}
}

/*GetRoutingEmailDomainRoutesNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetRoutingEmailDomainRoutesNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingEmailDomainRoutesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/domains/{domainName}/routes][%d] getRoutingEmailDomainRoutesNotFound  %+v", 404, o.Payload)
}

func (o *GetRoutingEmailDomainRoutesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingEmailDomainRoutesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingEmailDomainRoutesRequestTimeout creates a GetRoutingEmailDomainRoutesRequestTimeout with default headers values
func NewGetRoutingEmailDomainRoutesRequestTimeout() *GetRoutingEmailDomainRoutesRequestTimeout {
	return &GetRoutingEmailDomainRoutesRequestTimeout{}
}

/*GetRoutingEmailDomainRoutesRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetRoutingEmailDomainRoutesRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingEmailDomainRoutesRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/domains/{domainName}/routes][%d] getRoutingEmailDomainRoutesRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetRoutingEmailDomainRoutesRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingEmailDomainRoutesRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingEmailDomainRoutesRequestEntityTooLarge creates a GetRoutingEmailDomainRoutesRequestEntityTooLarge with default headers values
func NewGetRoutingEmailDomainRoutesRequestEntityTooLarge() *GetRoutingEmailDomainRoutesRequestEntityTooLarge {
	return &GetRoutingEmailDomainRoutesRequestEntityTooLarge{}
}

/*GetRoutingEmailDomainRoutesRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetRoutingEmailDomainRoutesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingEmailDomainRoutesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/domains/{domainName}/routes][%d] getRoutingEmailDomainRoutesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetRoutingEmailDomainRoutesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingEmailDomainRoutesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingEmailDomainRoutesUnsupportedMediaType creates a GetRoutingEmailDomainRoutesUnsupportedMediaType with default headers values
func NewGetRoutingEmailDomainRoutesUnsupportedMediaType() *GetRoutingEmailDomainRoutesUnsupportedMediaType {
	return &GetRoutingEmailDomainRoutesUnsupportedMediaType{}
}

/*GetRoutingEmailDomainRoutesUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetRoutingEmailDomainRoutesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingEmailDomainRoutesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/domains/{domainName}/routes][%d] getRoutingEmailDomainRoutesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetRoutingEmailDomainRoutesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingEmailDomainRoutesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingEmailDomainRoutesTooManyRequests creates a GetRoutingEmailDomainRoutesTooManyRequests with default headers values
func NewGetRoutingEmailDomainRoutesTooManyRequests() *GetRoutingEmailDomainRoutesTooManyRequests {
	return &GetRoutingEmailDomainRoutesTooManyRequests{}
}

/*GetRoutingEmailDomainRoutesTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetRoutingEmailDomainRoutesTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingEmailDomainRoutesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/domains/{domainName}/routes][%d] getRoutingEmailDomainRoutesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetRoutingEmailDomainRoutesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingEmailDomainRoutesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingEmailDomainRoutesInternalServerError creates a GetRoutingEmailDomainRoutesInternalServerError with default headers values
func NewGetRoutingEmailDomainRoutesInternalServerError() *GetRoutingEmailDomainRoutesInternalServerError {
	return &GetRoutingEmailDomainRoutesInternalServerError{}
}

/*GetRoutingEmailDomainRoutesInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetRoutingEmailDomainRoutesInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingEmailDomainRoutesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/domains/{domainName}/routes][%d] getRoutingEmailDomainRoutesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRoutingEmailDomainRoutesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingEmailDomainRoutesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingEmailDomainRoutesServiceUnavailable creates a GetRoutingEmailDomainRoutesServiceUnavailable with default headers values
func NewGetRoutingEmailDomainRoutesServiceUnavailable() *GetRoutingEmailDomainRoutesServiceUnavailable {
	return &GetRoutingEmailDomainRoutesServiceUnavailable{}
}

/*GetRoutingEmailDomainRoutesServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetRoutingEmailDomainRoutesServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingEmailDomainRoutesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/domains/{domainName}/routes][%d] getRoutingEmailDomainRoutesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetRoutingEmailDomainRoutesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingEmailDomainRoutesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingEmailDomainRoutesGatewayTimeout creates a GetRoutingEmailDomainRoutesGatewayTimeout with default headers values
func NewGetRoutingEmailDomainRoutesGatewayTimeout() *GetRoutingEmailDomainRoutesGatewayTimeout {
	return &GetRoutingEmailDomainRoutesGatewayTimeout{}
}

/*GetRoutingEmailDomainRoutesGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetRoutingEmailDomainRoutesGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingEmailDomainRoutesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/domains/{domainName}/routes][%d] getRoutingEmailDomainRoutesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetRoutingEmailDomainRoutesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingEmailDomainRoutesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

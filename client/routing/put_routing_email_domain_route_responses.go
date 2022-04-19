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

// PutRoutingEmailDomainRouteReader is a Reader for the PutRoutingEmailDomainRoute structure.
type PutRoutingEmailDomainRouteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutRoutingEmailDomainRouteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutRoutingEmailDomainRouteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutRoutingEmailDomainRouteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutRoutingEmailDomainRouteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutRoutingEmailDomainRouteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutRoutingEmailDomainRouteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPutRoutingEmailDomainRouteRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPutRoutingEmailDomainRouteRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPutRoutingEmailDomainRouteUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutRoutingEmailDomainRouteTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutRoutingEmailDomainRouteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutRoutingEmailDomainRouteServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutRoutingEmailDomainRouteGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutRoutingEmailDomainRouteOK creates a PutRoutingEmailDomainRouteOK with default headers values
func NewPutRoutingEmailDomainRouteOK() *PutRoutingEmailDomainRouteOK {
	return &PutRoutingEmailDomainRouteOK{}
}

/*PutRoutingEmailDomainRouteOK handles this case with default header values.

successful operation
*/
type PutRoutingEmailDomainRouteOK struct {
	Payload *models.InboundRoute
}

func (o *PutRoutingEmailDomainRouteOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/routing/email/domains/{domainName}/routes/{routeId}][%d] putRoutingEmailDomainRouteOK  %+v", 200, o.Payload)
}

func (o *PutRoutingEmailDomainRouteOK) GetPayload() *models.InboundRoute {
	return o.Payload
}

func (o *PutRoutingEmailDomainRouteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InboundRoute)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRoutingEmailDomainRouteBadRequest creates a PutRoutingEmailDomainRouteBadRequest with default headers values
func NewPutRoutingEmailDomainRouteBadRequest() *PutRoutingEmailDomainRouteBadRequest {
	return &PutRoutingEmailDomainRouteBadRequest{}
}

/*PutRoutingEmailDomainRouteBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutRoutingEmailDomainRouteBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PutRoutingEmailDomainRouteBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/routing/email/domains/{domainName}/routes/{routeId}][%d] putRoutingEmailDomainRouteBadRequest  %+v", 400, o.Payload)
}

func (o *PutRoutingEmailDomainRouteBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutRoutingEmailDomainRouteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRoutingEmailDomainRouteUnauthorized creates a PutRoutingEmailDomainRouteUnauthorized with default headers values
func NewPutRoutingEmailDomainRouteUnauthorized() *PutRoutingEmailDomainRouteUnauthorized {
	return &PutRoutingEmailDomainRouteUnauthorized{}
}

/*PutRoutingEmailDomainRouteUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PutRoutingEmailDomainRouteUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PutRoutingEmailDomainRouteUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/routing/email/domains/{domainName}/routes/{routeId}][%d] putRoutingEmailDomainRouteUnauthorized  %+v", 401, o.Payload)
}

func (o *PutRoutingEmailDomainRouteUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutRoutingEmailDomainRouteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRoutingEmailDomainRouteForbidden creates a PutRoutingEmailDomainRouteForbidden with default headers values
func NewPutRoutingEmailDomainRouteForbidden() *PutRoutingEmailDomainRouteForbidden {
	return &PutRoutingEmailDomainRouteForbidden{}
}

/*PutRoutingEmailDomainRouteForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PutRoutingEmailDomainRouteForbidden struct {
	Payload *models.ErrorBody
}

func (o *PutRoutingEmailDomainRouteForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/routing/email/domains/{domainName}/routes/{routeId}][%d] putRoutingEmailDomainRouteForbidden  %+v", 403, o.Payload)
}

func (o *PutRoutingEmailDomainRouteForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutRoutingEmailDomainRouteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRoutingEmailDomainRouteNotFound creates a PutRoutingEmailDomainRouteNotFound with default headers values
func NewPutRoutingEmailDomainRouteNotFound() *PutRoutingEmailDomainRouteNotFound {
	return &PutRoutingEmailDomainRouteNotFound{}
}

/*PutRoutingEmailDomainRouteNotFound handles this case with default header values.

The requested resource was not found.
*/
type PutRoutingEmailDomainRouteNotFound struct {
	Payload *models.ErrorBody
}

func (o *PutRoutingEmailDomainRouteNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/routing/email/domains/{domainName}/routes/{routeId}][%d] putRoutingEmailDomainRouteNotFound  %+v", 404, o.Payload)
}

func (o *PutRoutingEmailDomainRouteNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutRoutingEmailDomainRouteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRoutingEmailDomainRouteRequestTimeout creates a PutRoutingEmailDomainRouteRequestTimeout with default headers values
func NewPutRoutingEmailDomainRouteRequestTimeout() *PutRoutingEmailDomainRouteRequestTimeout {
	return &PutRoutingEmailDomainRouteRequestTimeout{}
}

/*PutRoutingEmailDomainRouteRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PutRoutingEmailDomainRouteRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PutRoutingEmailDomainRouteRequestTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/routing/email/domains/{domainName}/routes/{routeId}][%d] putRoutingEmailDomainRouteRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutRoutingEmailDomainRouteRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutRoutingEmailDomainRouteRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRoutingEmailDomainRouteRequestEntityTooLarge creates a PutRoutingEmailDomainRouteRequestEntityTooLarge with default headers values
func NewPutRoutingEmailDomainRouteRequestEntityTooLarge() *PutRoutingEmailDomainRouteRequestEntityTooLarge {
	return &PutRoutingEmailDomainRouteRequestEntityTooLarge{}
}

/*PutRoutingEmailDomainRouteRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PutRoutingEmailDomainRouteRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PutRoutingEmailDomainRouteRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /api/v2/routing/email/domains/{domainName}/routes/{routeId}][%d] putRoutingEmailDomainRouteRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutRoutingEmailDomainRouteRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutRoutingEmailDomainRouteRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRoutingEmailDomainRouteUnsupportedMediaType creates a PutRoutingEmailDomainRouteUnsupportedMediaType with default headers values
func NewPutRoutingEmailDomainRouteUnsupportedMediaType() *PutRoutingEmailDomainRouteUnsupportedMediaType {
	return &PutRoutingEmailDomainRouteUnsupportedMediaType{}
}

/*PutRoutingEmailDomainRouteUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutRoutingEmailDomainRouteUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PutRoutingEmailDomainRouteUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v2/routing/email/domains/{domainName}/routes/{routeId}][%d] putRoutingEmailDomainRouteUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutRoutingEmailDomainRouteUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutRoutingEmailDomainRouteUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRoutingEmailDomainRouteTooManyRequests creates a PutRoutingEmailDomainRouteTooManyRequests with default headers values
func NewPutRoutingEmailDomainRouteTooManyRequests() *PutRoutingEmailDomainRouteTooManyRequests {
	return &PutRoutingEmailDomainRouteTooManyRequests{}
}

/*PutRoutingEmailDomainRouteTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PutRoutingEmailDomainRouteTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PutRoutingEmailDomainRouteTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v2/routing/email/domains/{domainName}/routes/{routeId}][%d] putRoutingEmailDomainRouteTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutRoutingEmailDomainRouteTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutRoutingEmailDomainRouteTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRoutingEmailDomainRouteInternalServerError creates a PutRoutingEmailDomainRouteInternalServerError with default headers values
func NewPutRoutingEmailDomainRouteInternalServerError() *PutRoutingEmailDomainRouteInternalServerError {
	return &PutRoutingEmailDomainRouteInternalServerError{}
}

/*PutRoutingEmailDomainRouteInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutRoutingEmailDomainRouteInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PutRoutingEmailDomainRouteInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/routing/email/domains/{domainName}/routes/{routeId}][%d] putRoutingEmailDomainRouteInternalServerError  %+v", 500, o.Payload)
}

func (o *PutRoutingEmailDomainRouteInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutRoutingEmailDomainRouteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRoutingEmailDomainRouteServiceUnavailable creates a PutRoutingEmailDomainRouteServiceUnavailable with default headers values
func NewPutRoutingEmailDomainRouteServiceUnavailable() *PutRoutingEmailDomainRouteServiceUnavailable {
	return &PutRoutingEmailDomainRouteServiceUnavailable{}
}

/*PutRoutingEmailDomainRouteServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutRoutingEmailDomainRouteServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PutRoutingEmailDomainRouteServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/routing/email/domains/{domainName}/routes/{routeId}][%d] putRoutingEmailDomainRouteServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutRoutingEmailDomainRouteServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutRoutingEmailDomainRouteServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRoutingEmailDomainRouteGatewayTimeout creates a PutRoutingEmailDomainRouteGatewayTimeout with default headers values
func NewPutRoutingEmailDomainRouteGatewayTimeout() *PutRoutingEmailDomainRouteGatewayTimeout {
	return &PutRoutingEmailDomainRouteGatewayTimeout{}
}

/*PutRoutingEmailDomainRouteGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PutRoutingEmailDomainRouteGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PutRoutingEmailDomainRouteGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/routing/email/domains/{domainName}/routes/{routeId}][%d] putRoutingEmailDomainRouteGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutRoutingEmailDomainRouteGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutRoutingEmailDomainRouteGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

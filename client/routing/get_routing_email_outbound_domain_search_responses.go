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

// GetRoutingEmailOutboundDomainSearchReader is a Reader for the GetRoutingEmailOutboundDomainSearch structure.
type GetRoutingEmailOutboundDomainSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRoutingEmailOutboundDomainSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRoutingEmailOutboundDomainSearchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRoutingEmailOutboundDomainSearchBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetRoutingEmailOutboundDomainSearchUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRoutingEmailOutboundDomainSearchForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRoutingEmailOutboundDomainSearchNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetRoutingEmailOutboundDomainSearchRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetRoutingEmailOutboundDomainSearchRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetRoutingEmailOutboundDomainSearchUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetRoutingEmailOutboundDomainSearchTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRoutingEmailOutboundDomainSearchInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetRoutingEmailOutboundDomainSearchServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetRoutingEmailOutboundDomainSearchGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRoutingEmailOutboundDomainSearchOK creates a GetRoutingEmailOutboundDomainSearchOK with default headers values
func NewGetRoutingEmailOutboundDomainSearchOK() *GetRoutingEmailOutboundDomainSearchOK {
	return &GetRoutingEmailOutboundDomainSearchOK{}
}

/*GetRoutingEmailOutboundDomainSearchOK handles this case with default header values.

successful operation
*/
type GetRoutingEmailOutboundDomainSearchOK struct {
	Payload *models.OutboundDomain
}

func (o *GetRoutingEmailOutboundDomainSearchOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/outbound/domains/{domainId}/search][%d] getRoutingEmailOutboundDomainSearchOK  %+v", 200, o.Payload)
}

func (o *GetRoutingEmailOutboundDomainSearchOK) GetPayload() *models.OutboundDomain {
	return o.Payload
}

func (o *GetRoutingEmailOutboundDomainSearchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OutboundDomain)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingEmailOutboundDomainSearchBadRequest creates a GetRoutingEmailOutboundDomainSearchBadRequest with default headers values
func NewGetRoutingEmailOutboundDomainSearchBadRequest() *GetRoutingEmailOutboundDomainSearchBadRequest {
	return &GetRoutingEmailOutboundDomainSearchBadRequest{}
}

/*GetRoutingEmailOutboundDomainSearchBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetRoutingEmailOutboundDomainSearchBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingEmailOutboundDomainSearchBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/outbound/domains/{domainId}/search][%d] getRoutingEmailOutboundDomainSearchBadRequest  %+v", 400, o.Payload)
}

func (o *GetRoutingEmailOutboundDomainSearchBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingEmailOutboundDomainSearchBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingEmailOutboundDomainSearchUnauthorized creates a GetRoutingEmailOutboundDomainSearchUnauthorized with default headers values
func NewGetRoutingEmailOutboundDomainSearchUnauthorized() *GetRoutingEmailOutboundDomainSearchUnauthorized {
	return &GetRoutingEmailOutboundDomainSearchUnauthorized{}
}

/*GetRoutingEmailOutboundDomainSearchUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetRoutingEmailOutboundDomainSearchUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingEmailOutboundDomainSearchUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/outbound/domains/{domainId}/search][%d] getRoutingEmailOutboundDomainSearchUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRoutingEmailOutboundDomainSearchUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingEmailOutboundDomainSearchUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingEmailOutboundDomainSearchForbidden creates a GetRoutingEmailOutboundDomainSearchForbidden with default headers values
func NewGetRoutingEmailOutboundDomainSearchForbidden() *GetRoutingEmailOutboundDomainSearchForbidden {
	return &GetRoutingEmailOutboundDomainSearchForbidden{}
}

/*GetRoutingEmailOutboundDomainSearchForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetRoutingEmailOutboundDomainSearchForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingEmailOutboundDomainSearchForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/outbound/domains/{domainId}/search][%d] getRoutingEmailOutboundDomainSearchForbidden  %+v", 403, o.Payload)
}

func (o *GetRoutingEmailOutboundDomainSearchForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingEmailOutboundDomainSearchForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingEmailOutboundDomainSearchNotFound creates a GetRoutingEmailOutboundDomainSearchNotFound with default headers values
func NewGetRoutingEmailOutboundDomainSearchNotFound() *GetRoutingEmailOutboundDomainSearchNotFound {
	return &GetRoutingEmailOutboundDomainSearchNotFound{}
}

/*GetRoutingEmailOutboundDomainSearchNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetRoutingEmailOutboundDomainSearchNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingEmailOutboundDomainSearchNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/outbound/domains/{domainId}/search][%d] getRoutingEmailOutboundDomainSearchNotFound  %+v", 404, o.Payload)
}

func (o *GetRoutingEmailOutboundDomainSearchNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingEmailOutboundDomainSearchNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingEmailOutboundDomainSearchRequestTimeout creates a GetRoutingEmailOutboundDomainSearchRequestTimeout with default headers values
func NewGetRoutingEmailOutboundDomainSearchRequestTimeout() *GetRoutingEmailOutboundDomainSearchRequestTimeout {
	return &GetRoutingEmailOutboundDomainSearchRequestTimeout{}
}

/*GetRoutingEmailOutboundDomainSearchRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetRoutingEmailOutboundDomainSearchRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingEmailOutboundDomainSearchRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/outbound/domains/{domainId}/search][%d] getRoutingEmailOutboundDomainSearchRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetRoutingEmailOutboundDomainSearchRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingEmailOutboundDomainSearchRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingEmailOutboundDomainSearchRequestEntityTooLarge creates a GetRoutingEmailOutboundDomainSearchRequestEntityTooLarge with default headers values
func NewGetRoutingEmailOutboundDomainSearchRequestEntityTooLarge() *GetRoutingEmailOutboundDomainSearchRequestEntityTooLarge {
	return &GetRoutingEmailOutboundDomainSearchRequestEntityTooLarge{}
}

/*GetRoutingEmailOutboundDomainSearchRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetRoutingEmailOutboundDomainSearchRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingEmailOutboundDomainSearchRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/outbound/domains/{domainId}/search][%d] getRoutingEmailOutboundDomainSearchRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetRoutingEmailOutboundDomainSearchRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingEmailOutboundDomainSearchRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingEmailOutboundDomainSearchUnsupportedMediaType creates a GetRoutingEmailOutboundDomainSearchUnsupportedMediaType with default headers values
func NewGetRoutingEmailOutboundDomainSearchUnsupportedMediaType() *GetRoutingEmailOutboundDomainSearchUnsupportedMediaType {
	return &GetRoutingEmailOutboundDomainSearchUnsupportedMediaType{}
}

/*GetRoutingEmailOutboundDomainSearchUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetRoutingEmailOutboundDomainSearchUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingEmailOutboundDomainSearchUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/outbound/domains/{domainId}/search][%d] getRoutingEmailOutboundDomainSearchUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetRoutingEmailOutboundDomainSearchUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingEmailOutboundDomainSearchUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingEmailOutboundDomainSearchTooManyRequests creates a GetRoutingEmailOutboundDomainSearchTooManyRequests with default headers values
func NewGetRoutingEmailOutboundDomainSearchTooManyRequests() *GetRoutingEmailOutboundDomainSearchTooManyRequests {
	return &GetRoutingEmailOutboundDomainSearchTooManyRequests{}
}

/*GetRoutingEmailOutboundDomainSearchTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetRoutingEmailOutboundDomainSearchTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingEmailOutboundDomainSearchTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/outbound/domains/{domainId}/search][%d] getRoutingEmailOutboundDomainSearchTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetRoutingEmailOutboundDomainSearchTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingEmailOutboundDomainSearchTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingEmailOutboundDomainSearchInternalServerError creates a GetRoutingEmailOutboundDomainSearchInternalServerError with default headers values
func NewGetRoutingEmailOutboundDomainSearchInternalServerError() *GetRoutingEmailOutboundDomainSearchInternalServerError {
	return &GetRoutingEmailOutboundDomainSearchInternalServerError{}
}

/*GetRoutingEmailOutboundDomainSearchInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetRoutingEmailOutboundDomainSearchInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingEmailOutboundDomainSearchInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/outbound/domains/{domainId}/search][%d] getRoutingEmailOutboundDomainSearchInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRoutingEmailOutboundDomainSearchInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingEmailOutboundDomainSearchInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingEmailOutboundDomainSearchServiceUnavailable creates a GetRoutingEmailOutboundDomainSearchServiceUnavailable with default headers values
func NewGetRoutingEmailOutboundDomainSearchServiceUnavailable() *GetRoutingEmailOutboundDomainSearchServiceUnavailable {
	return &GetRoutingEmailOutboundDomainSearchServiceUnavailable{}
}

/*GetRoutingEmailOutboundDomainSearchServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetRoutingEmailOutboundDomainSearchServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingEmailOutboundDomainSearchServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/outbound/domains/{domainId}/search][%d] getRoutingEmailOutboundDomainSearchServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetRoutingEmailOutboundDomainSearchServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingEmailOutboundDomainSearchServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingEmailOutboundDomainSearchGatewayTimeout creates a GetRoutingEmailOutboundDomainSearchGatewayTimeout with default headers values
func NewGetRoutingEmailOutboundDomainSearchGatewayTimeout() *GetRoutingEmailOutboundDomainSearchGatewayTimeout {
	return &GetRoutingEmailOutboundDomainSearchGatewayTimeout{}
}

/*GetRoutingEmailOutboundDomainSearchGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetRoutingEmailOutboundDomainSearchGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingEmailOutboundDomainSearchGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/outbound/domains/{domainId}/search][%d] getRoutingEmailOutboundDomainSearchGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetRoutingEmailOutboundDomainSearchGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingEmailOutboundDomainSearchGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

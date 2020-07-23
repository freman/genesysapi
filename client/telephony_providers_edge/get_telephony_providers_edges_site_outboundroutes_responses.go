// Code generated by go-swagger; DO NOT EDIT.

package telephony_providers_edge

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetTelephonyProvidersEdgesSiteOutboundroutesReader is a Reader for the GetTelephonyProvidersEdgesSiteOutboundroutes structure.
type GetTelephonyProvidersEdgesSiteOutboundroutesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTelephonyProvidersEdgesSiteOutboundroutesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTelephonyProvidersEdgesSiteOutboundroutesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetTelephonyProvidersEdgesSiteOutboundroutesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetTelephonyProvidersEdgesSiteOutboundroutesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetTelephonyProvidersEdgesSiteOutboundroutesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTelephonyProvidersEdgesSiteOutboundroutesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetTelephonyProvidersEdgesSiteOutboundroutesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetTelephonyProvidersEdgesSiteOutboundroutesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetTelephonyProvidersEdgesSiteOutboundroutesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetTelephonyProvidersEdgesSiteOutboundroutesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetTelephonyProvidersEdgesSiteOutboundroutesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetTelephonyProvidersEdgesSiteOutboundroutesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTelephonyProvidersEdgesSiteOutboundroutesOK creates a GetTelephonyProvidersEdgesSiteOutboundroutesOK with default headers values
func NewGetTelephonyProvidersEdgesSiteOutboundroutesOK() *GetTelephonyProvidersEdgesSiteOutboundroutesOK {
	return &GetTelephonyProvidersEdgesSiteOutboundroutesOK{}
}

/*GetTelephonyProvidersEdgesSiteOutboundroutesOK handles this case with default header values.

successful operation
*/
type GetTelephonyProvidersEdgesSiteOutboundroutesOK struct {
	Payload *models.OutboundRouteBaseEntityListing
}

func (o *GetTelephonyProvidersEdgesSiteOutboundroutesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/sites/{siteId}/outboundroutes][%d] getTelephonyProvidersEdgesSiteOutboundroutesOK  %+v", 200, o.Payload)
}

func (o *GetTelephonyProvidersEdgesSiteOutboundroutesOK) GetPayload() *models.OutboundRouteBaseEntityListing {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesSiteOutboundroutesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OutboundRouteBaseEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesSiteOutboundroutesBadRequest creates a GetTelephonyProvidersEdgesSiteOutboundroutesBadRequest with default headers values
func NewGetTelephonyProvidersEdgesSiteOutboundroutesBadRequest() *GetTelephonyProvidersEdgesSiteOutboundroutesBadRequest {
	return &GetTelephonyProvidersEdgesSiteOutboundroutesBadRequest{}
}

/*GetTelephonyProvidersEdgesSiteOutboundroutesBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetTelephonyProvidersEdgesSiteOutboundroutesBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesSiteOutboundroutesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/sites/{siteId}/outboundroutes][%d] getTelephonyProvidersEdgesSiteOutboundroutesBadRequest  %+v", 400, o.Payload)
}

func (o *GetTelephonyProvidersEdgesSiteOutboundroutesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesSiteOutboundroutesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesSiteOutboundroutesUnauthorized creates a GetTelephonyProvidersEdgesSiteOutboundroutesUnauthorized with default headers values
func NewGetTelephonyProvidersEdgesSiteOutboundroutesUnauthorized() *GetTelephonyProvidersEdgesSiteOutboundroutesUnauthorized {
	return &GetTelephonyProvidersEdgesSiteOutboundroutesUnauthorized{}
}

/*GetTelephonyProvidersEdgesSiteOutboundroutesUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetTelephonyProvidersEdgesSiteOutboundroutesUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesSiteOutboundroutesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/sites/{siteId}/outboundroutes][%d] getTelephonyProvidersEdgesSiteOutboundroutesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetTelephonyProvidersEdgesSiteOutboundroutesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesSiteOutboundroutesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesSiteOutboundroutesForbidden creates a GetTelephonyProvidersEdgesSiteOutboundroutesForbidden with default headers values
func NewGetTelephonyProvidersEdgesSiteOutboundroutesForbidden() *GetTelephonyProvidersEdgesSiteOutboundroutesForbidden {
	return &GetTelephonyProvidersEdgesSiteOutboundroutesForbidden{}
}

/*GetTelephonyProvidersEdgesSiteOutboundroutesForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetTelephonyProvidersEdgesSiteOutboundroutesForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesSiteOutboundroutesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/sites/{siteId}/outboundroutes][%d] getTelephonyProvidersEdgesSiteOutboundroutesForbidden  %+v", 403, o.Payload)
}

func (o *GetTelephonyProvidersEdgesSiteOutboundroutesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesSiteOutboundroutesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesSiteOutboundroutesNotFound creates a GetTelephonyProvidersEdgesSiteOutboundroutesNotFound with default headers values
func NewGetTelephonyProvidersEdgesSiteOutboundroutesNotFound() *GetTelephonyProvidersEdgesSiteOutboundroutesNotFound {
	return &GetTelephonyProvidersEdgesSiteOutboundroutesNotFound{}
}

/*GetTelephonyProvidersEdgesSiteOutboundroutesNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetTelephonyProvidersEdgesSiteOutboundroutesNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesSiteOutboundroutesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/sites/{siteId}/outboundroutes][%d] getTelephonyProvidersEdgesSiteOutboundroutesNotFound  %+v", 404, o.Payload)
}

func (o *GetTelephonyProvidersEdgesSiteOutboundroutesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesSiteOutboundroutesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesSiteOutboundroutesRequestEntityTooLarge creates a GetTelephonyProvidersEdgesSiteOutboundroutesRequestEntityTooLarge with default headers values
func NewGetTelephonyProvidersEdgesSiteOutboundroutesRequestEntityTooLarge() *GetTelephonyProvidersEdgesSiteOutboundroutesRequestEntityTooLarge {
	return &GetTelephonyProvidersEdgesSiteOutboundroutesRequestEntityTooLarge{}
}

/*GetTelephonyProvidersEdgesSiteOutboundroutesRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetTelephonyProvidersEdgesSiteOutboundroutesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesSiteOutboundroutesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/sites/{siteId}/outboundroutes][%d] getTelephonyProvidersEdgesSiteOutboundroutesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetTelephonyProvidersEdgesSiteOutboundroutesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesSiteOutboundroutesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesSiteOutboundroutesUnsupportedMediaType creates a GetTelephonyProvidersEdgesSiteOutboundroutesUnsupportedMediaType with default headers values
func NewGetTelephonyProvidersEdgesSiteOutboundroutesUnsupportedMediaType() *GetTelephonyProvidersEdgesSiteOutboundroutesUnsupportedMediaType {
	return &GetTelephonyProvidersEdgesSiteOutboundroutesUnsupportedMediaType{}
}

/*GetTelephonyProvidersEdgesSiteOutboundroutesUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetTelephonyProvidersEdgesSiteOutboundroutesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesSiteOutboundroutesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/sites/{siteId}/outboundroutes][%d] getTelephonyProvidersEdgesSiteOutboundroutesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetTelephonyProvidersEdgesSiteOutboundroutesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesSiteOutboundroutesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesSiteOutboundroutesTooManyRequests creates a GetTelephonyProvidersEdgesSiteOutboundroutesTooManyRequests with default headers values
func NewGetTelephonyProvidersEdgesSiteOutboundroutesTooManyRequests() *GetTelephonyProvidersEdgesSiteOutboundroutesTooManyRequests {
	return &GetTelephonyProvidersEdgesSiteOutboundroutesTooManyRequests{}
}

/*GetTelephonyProvidersEdgesSiteOutboundroutesTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetTelephonyProvidersEdgesSiteOutboundroutesTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesSiteOutboundroutesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/sites/{siteId}/outboundroutes][%d] getTelephonyProvidersEdgesSiteOutboundroutesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetTelephonyProvidersEdgesSiteOutboundroutesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesSiteOutboundroutesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesSiteOutboundroutesInternalServerError creates a GetTelephonyProvidersEdgesSiteOutboundroutesInternalServerError with default headers values
func NewGetTelephonyProvidersEdgesSiteOutboundroutesInternalServerError() *GetTelephonyProvidersEdgesSiteOutboundroutesInternalServerError {
	return &GetTelephonyProvidersEdgesSiteOutboundroutesInternalServerError{}
}

/*GetTelephonyProvidersEdgesSiteOutboundroutesInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetTelephonyProvidersEdgesSiteOutboundroutesInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesSiteOutboundroutesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/sites/{siteId}/outboundroutes][%d] getTelephonyProvidersEdgesSiteOutboundroutesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTelephonyProvidersEdgesSiteOutboundroutesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesSiteOutboundroutesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesSiteOutboundroutesServiceUnavailable creates a GetTelephonyProvidersEdgesSiteOutboundroutesServiceUnavailable with default headers values
func NewGetTelephonyProvidersEdgesSiteOutboundroutesServiceUnavailable() *GetTelephonyProvidersEdgesSiteOutboundroutesServiceUnavailable {
	return &GetTelephonyProvidersEdgesSiteOutboundroutesServiceUnavailable{}
}

/*GetTelephonyProvidersEdgesSiteOutboundroutesServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetTelephonyProvidersEdgesSiteOutboundroutesServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesSiteOutboundroutesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/sites/{siteId}/outboundroutes][%d] getTelephonyProvidersEdgesSiteOutboundroutesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetTelephonyProvidersEdgesSiteOutboundroutesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesSiteOutboundroutesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesSiteOutboundroutesGatewayTimeout creates a GetTelephonyProvidersEdgesSiteOutboundroutesGatewayTimeout with default headers values
func NewGetTelephonyProvidersEdgesSiteOutboundroutesGatewayTimeout() *GetTelephonyProvidersEdgesSiteOutboundroutesGatewayTimeout {
	return &GetTelephonyProvidersEdgesSiteOutboundroutesGatewayTimeout{}
}

/*GetTelephonyProvidersEdgesSiteOutboundroutesGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetTelephonyProvidersEdgesSiteOutboundroutesGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesSiteOutboundroutesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/sites/{siteId}/outboundroutes][%d] getTelephonyProvidersEdgesSiteOutboundroutesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetTelephonyProvidersEdgesSiteOutboundroutesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesSiteOutboundroutesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

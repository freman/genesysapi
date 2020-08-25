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

// GetTelephonyProvidersEdgesSiteOutboundrouteReader is a Reader for the GetTelephonyProvidersEdgesSiteOutboundroute structure.
type GetTelephonyProvidersEdgesSiteOutboundrouteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTelephonyProvidersEdgesSiteOutboundrouteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTelephonyProvidersEdgesSiteOutboundrouteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetTelephonyProvidersEdgesSiteOutboundrouteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetTelephonyProvidersEdgesSiteOutboundrouteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetTelephonyProvidersEdgesSiteOutboundrouteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTelephonyProvidersEdgesSiteOutboundrouteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetTelephonyProvidersEdgesSiteOutboundrouteRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetTelephonyProvidersEdgesSiteOutboundrouteUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetTelephonyProvidersEdgesSiteOutboundrouteTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetTelephonyProvidersEdgesSiteOutboundrouteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetTelephonyProvidersEdgesSiteOutboundrouteServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetTelephonyProvidersEdgesSiteOutboundrouteGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTelephonyProvidersEdgesSiteOutboundrouteOK creates a GetTelephonyProvidersEdgesSiteOutboundrouteOK with default headers values
func NewGetTelephonyProvidersEdgesSiteOutboundrouteOK() *GetTelephonyProvidersEdgesSiteOutboundrouteOK {
	return &GetTelephonyProvidersEdgesSiteOutboundrouteOK{}
}

/*GetTelephonyProvidersEdgesSiteOutboundrouteOK handles this case with default header values.

successful operation
*/
type GetTelephonyProvidersEdgesSiteOutboundrouteOK struct {
	Payload *models.OutboundRouteBase
}

func (o *GetTelephonyProvidersEdgesSiteOutboundrouteOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/sites/{siteId}/outboundroutes/{outboundRouteId}][%d] getTelephonyProvidersEdgesSiteOutboundrouteOK  %+v", 200, o.Payload)
}

func (o *GetTelephonyProvidersEdgesSiteOutboundrouteOK) GetPayload() *models.OutboundRouteBase {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesSiteOutboundrouteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OutboundRouteBase)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesSiteOutboundrouteBadRequest creates a GetTelephonyProvidersEdgesSiteOutboundrouteBadRequest with default headers values
func NewGetTelephonyProvidersEdgesSiteOutboundrouteBadRequest() *GetTelephonyProvidersEdgesSiteOutboundrouteBadRequest {
	return &GetTelephonyProvidersEdgesSiteOutboundrouteBadRequest{}
}

/*GetTelephonyProvidersEdgesSiteOutboundrouteBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetTelephonyProvidersEdgesSiteOutboundrouteBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesSiteOutboundrouteBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/sites/{siteId}/outboundroutes/{outboundRouteId}][%d] getTelephonyProvidersEdgesSiteOutboundrouteBadRequest  %+v", 400, o.Payload)
}

func (o *GetTelephonyProvidersEdgesSiteOutboundrouteBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesSiteOutboundrouteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesSiteOutboundrouteUnauthorized creates a GetTelephonyProvidersEdgesSiteOutboundrouteUnauthorized with default headers values
func NewGetTelephonyProvidersEdgesSiteOutboundrouteUnauthorized() *GetTelephonyProvidersEdgesSiteOutboundrouteUnauthorized {
	return &GetTelephonyProvidersEdgesSiteOutboundrouteUnauthorized{}
}

/*GetTelephonyProvidersEdgesSiteOutboundrouteUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetTelephonyProvidersEdgesSiteOutboundrouteUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesSiteOutboundrouteUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/sites/{siteId}/outboundroutes/{outboundRouteId}][%d] getTelephonyProvidersEdgesSiteOutboundrouteUnauthorized  %+v", 401, o.Payload)
}

func (o *GetTelephonyProvidersEdgesSiteOutboundrouteUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesSiteOutboundrouteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesSiteOutboundrouteForbidden creates a GetTelephonyProvidersEdgesSiteOutboundrouteForbidden with default headers values
func NewGetTelephonyProvidersEdgesSiteOutboundrouteForbidden() *GetTelephonyProvidersEdgesSiteOutboundrouteForbidden {
	return &GetTelephonyProvidersEdgesSiteOutboundrouteForbidden{}
}

/*GetTelephonyProvidersEdgesSiteOutboundrouteForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetTelephonyProvidersEdgesSiteOutboundrouteForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesSiteOutboundrouteForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/sites/{siteId}/outboundroutes/{outboundRouteId}][%d] getTelephonyProvidersEdgesSiteOutboundrouteForbidden  %+v", 403, o.Payload)
}

func (o *GetTelephonyProvidersEdgesSiteOutboundrouteForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesSiteOutboundrouteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesSiteOutboundrouteNotFound creates a GetTelephonyProvidersEdgesSiteOutboundrouteNotFound with default headers values
func NewGetTelephonyProvidersEdgesSiteOutboundrouteNotFound() *GetTelephonyProvidersEdgesSiteOutboundrouteNotFound {
	return &GetTelephonyProvidersEdgesSiteOutboundrouteNotFound{}
}

/*GetTelephonyProvidersEdgesSiteOutboundrouteNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetTelephonyProvidersEdgesSiteOutboundrouteNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesSiteOutboundrouteNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/sites/{siteId}/outboundroutes/{outboundRouteId}][%d] getTelephonyProvidersEdgesSiteOutboundrouteNotFound  %+v", 404, o.Payload)
}

func (o *GetTelephonyProvidersEdgesSiteOutboundrouteNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesSiteOutboundrouteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesSiteOutboundrouteRequestEntityTooLarge creates a GetTelephonyProvidersEdgesSiteOutboundrouteRequestEntityTooLarge with default headers values
func NewGetTelephonyProvidersEdgesSiteOutboundrouteRequestEntityTooLarge() *GetTelephonyProvidersEdgesSiteOutboundrouteRequestEntityTooLarge {
	return &GetTelephonyProvidersEdgesSiteOutboundrouteRequestEntityTooLarge{}
}

/*GetTelephonyProvidersEdgesSiteOutboundrouteRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetTelephonyProvidersEdgesSiteOutboundrouteRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesSiteOutboundrouteRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/sites/{siteId}/outboundroutes/{outboundRouteId}][%d] getTelephonyProvidersEdgesSiteOutboundrouteRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetTelephonyProvidersEdgesSiteOutboundrouteRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesSiteOutboundrouteRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesSiteOutboundrouteUnsupportedMediaType creates a GetTelephonyProvidersEdgesSiteOutboundrouteUnsupportedMediaType with default headers values
func NewGetTelephonyProvidersEdgesSiteOutboundrouteUnsupportedMediaType() *GetTelephonyProvidersEdgesSiteOutboundrouteUnsupportedMediaType {
	return &GetTelephonyProvidersEdgesSiteOutboundrouteUnsupportedMediaType{}
}

/*GetTelephonyProvidersEdgesSiteOutboundrouteUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetTelephonyProvidersEdgesSiteOutboundrouteUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesSiteOutboundrouteUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/sites/{siteId}/outboundroutes/{outboundRouteId}][%d] getTelephonyProvidersEdgesSiteOutboundrouteUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetTelephonyProvidersEdgesSiteOutboundrouteUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesSiteOutboundrouteUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesSiteOutboundrouteTooManyRequests creates a GetTelephonyProvidersEdgesSiteOutboundrouteTooManyRequests with default headers values
func NewGetTelephonyProvidersEdgesSiteOutboundrouteTooManyRequests() *GetTelephonyProvidersEdgesSiteOutboundrouteTooManyRequests {
	return &GetTelephonyProvidersEdgesSiteOutboundrouteTooManyRequests{}
}

/*GetTelephonyProvidersEdgesSiteOutboundrouteTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetTelephonyProvidersEdgesSiteOutboundrouteTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesSiteOutboundrouteTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/sites/{siteId}/outboundroutes/{outboundRouteId}][%d] getTelephonyProvidersEdgesSiteOutboundrouteTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetTelephonyProvidersEdgesSiteOutboundrouteTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesSiteOutboundrouteTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesSiteOutboundrouteInternalServerError creates a GetTelephonyProvidersEdgesSiteOutboundrouteInternalServerError with default headers values
func NewGetTelephonyProvidersEdgesSiteOutboundrouteInternalServerError() *GetTelephonyProvidersEdgesSiteOutboundrouteInternalServerError {
	return &GetTelephonyProvidersEdgesSiteOutboundrouteInternalServerError{}
}

/*GetTelephonyProvidersEdgesSiteOutboundrouteInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetTelephonyProvidersEdgesSiteOutboundrouteInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesSiteOutboundrouteInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/sites/{siteId}/outboundroutes/{outboundRouteId}][%d] getTelephonyProvidersEdgesSiteOutboundrouteInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTelephonyProvidersEdgesSiteOutboundrouteInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesSiteOutboundrouteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesSiteOutboundrouteServiceUnavailable creates a GetTelephonyProvidersEdgesSiteOutboundrouteServiceUnavailable with default headers values
func NewGetTelephonyProvidersEdgesSiteOutboundrouteServiceUnavailable() *GetTelephonyProvidersEdgesSiteOutboundrouteServiceUnavailable {
	return &GetTelephonyProvidersEdgesSiteOutboundrouteServiceUnavailable{}
}

/*GetTelephonyProvidersEdgesSiteOutboundrouteServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetTelephonyProvidersEdgesSiteOutboundrouteServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesSiteOutboundrouteServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/sites/{siteId}/outboundroutes/{outboundRouteId}][%d] getTelephonyProvidersEdgesSiteOutboundrouteServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetTelephonyProvidersEdgesSiteOutboundrouteServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesSiteOutboundrouteServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesSiteOutboundrouteGatewayTimeout creates a GetTelephonyProvidersEdgesSiteOutboundrouteGatewayTimeout with default headers values
func NewGetTelephonyProvidersEdgesSiteOutboundrouteGatewayTimeout() *GetTelephonyProvidersEdgesSiteOutboundrouteGatewayTimeout {
	return &GetTelephonyProvidersEdgesSiteOutboundrouteGatewayTimeout{}
}

/*GetTelephonyProvidersEdgesSiteOutboundrouteGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetTelephonyProvidersEdgesSiteOutboundrouteGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesSiteOutboundrouteGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/sites/{siteId}/outboundroutes/{outboundRouteId}][%d] getTelephonyProvidersEdgesSiteOutboundrouteGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetTelephonyProvidersEdgesSiteOutboundrouteGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesSiteOutboundrouteGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
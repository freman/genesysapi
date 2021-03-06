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

// PostTelephonyProvidersEdgesSiteOutboundroutesReader is a Reader for the PostTelephonyProvidersEdgesSiteOutboundroutes structure.
type PostTelephonyProvidersEdgesSiteOutboundroutesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostTelephonyProvidersEdgesSiteOutboundroutesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostTelephonyProvidersEdgesSiteOutboundroutesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostTelephonyProvidersEdgesSiteOutboundroutesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostTelephonyProvidersEdgesSiteOutboundroutesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostTelephonyProvidersEdgesSiteOutboundroutesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostTelephonyProvidersEdgesSiteOutboundroutesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostTelephonyProvidersEdgesSiteOutboundroutesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostTelephonyProvidersEdgesSiteOutboundroutesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostTelephonyProvidersEdgesSiteOutboundroutesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostTelephonyProvidersEdgesSiteOutboundroutesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostTelephonyProvidersEdgesSiteOutboundroutesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostTelephonyProvidersEdgesSiteOutboundroutesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostTelephonyProvidersEdgesSiteOutboundroutesOK creates a PostTelephonyProvidersEdgesSiteOutboundroutesOK with default headers values
func NewPostTelephonyProvidersEdgesSiteOutboundroutesOK() *PostTelephonyProvidersEdgesSiteOutboundroutesOK {
	return &PostTelephonyProvidersEdgesSiteOutboundroutesOK{}
}

/*PostTelephonyProvidersEdgesSiteOutboundroutesOK handles this case with default header values.

successful operation
*/
type PostTelephonyProvidersEdgesSiteOutboundroutesOK struct {
	Payload *models.OutboundRouteBase
}

func (o *PostTelephonyProvidersEdgesSiteOutboundroutesOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/sites/{siteId}/outboundroutes][%d] postTelephonyProvidersEdgesSiteOutboundroutesOK  %+v", 200, o.Payload)
}

func (o *PostTelephonyProvidersEdgesSiteOutboundroutesOK) GetPayload() *models.OutboundRouteBase {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesSiteOutboundroutesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OutboundRouteBase)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgesSiteOutboundroutesBadRequest creates a PostTelephonyProvidersEdgesSiteOutboundroutesBadRequest with default headers values
func NewPostTelephonyProvidersEdgesSiteOutboundroutesBadRequest() *PostTelephonyProvidersEdgesSiteOutboundroutesBadRequest {
	return &PostTelephonyProvidersEdgesSiteOutboundroutesBadRequest{}
}

/*PostTelephonyProvidersEdgesSiteOutboundroutesBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostTelephonyProvidersEdgesSiteOutboundroutesBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgesSiteOutboundroutesBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/sites/{siteId}/outboundroutes][%d] postTelephonyProvidersEdgesSiteOutboundroutesBadRequest  %+v", 400, o.Payload)
}

func (o *PostTelephonyProvidersEdgesSiteOutboundroutesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesSiteOutboundroutesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgesSiteOutboundroutesUnauthorized creates a PostTelephonyProvidersEdgesSiteOutboundroutesUnauthorized with default headers values
func NewPostTelephonyProvidersEdgesSiteOutboundroutesUnauthorized() *PostTelephonyProvidersEdgesSiteOutboundroutesUnauthorized {
	return &PostTelephonyProvidersEdgesSiteOutboundroutesUnauthorized{}
}

/*PostTelephonyProvidersEdgesSiteOutboundroutesUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostTelephonyProvidersEdgesSiteOutboundroutesUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgesSiteOutboundroutesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/sites/{siteId}/outboundroutes][%d] postTelephonyProvidersEdgesSiteOutboundroutesUnauthorized  %+v", 401, o.Payload)
}

func (o *PostTelephonyProvidersEdgesSiteOutboundroutesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesSiteOutboundroutesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgesSiteOutboundroutesForbidden creates a PostTelephonyProvidersEdgesSiteOutboundroutesForbidden with default headers values
func NewPostTelephonyProvidersEdgesSiteOutboundroutesForbidden() *PostTelephonyProvidersEdgesSiteOutboundroutesForbidden {
	return &PostTelephonyProvidersEdgesSiteOutboundroutesForbidden{}
}

/*PostTelephonyProvidersEdgesSiteOutboundroutesForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostTelephonyProvidersEdgesSiteOutboundroutesForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgesSiteOutboundroutesForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/sites/{siteId}/outboundroutes][%d] postTelephonyProvidersEdgesSiteOutboundroutesForbidden  %+v", 403, o.Payload)
}

func (o *PostTelephonyProvidersEdgesSiteOutboundroutesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesSiteOutboundroutesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgesSiteOutboundroutesNotFound creates a PostTelephonyProvidersEdgesSiteOutboundroutesNotFound with default headers values
func NewPostTelephonyProvidersEdgesSiteOutboundroutesNotFound() *PostTelephonyProvidersEdgesSiteOutboundroutesNotFound {
	return &PostTelephonyProvidersEdgesSiteOutboundroutesNotFound{}
}

/*PostTelephonyProvidersEdgesSiteOutboundroutesNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostTelephonyProvidersEdgesSiteOutboundroutesNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgesSiteOutboundroutesNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/sites/{siteId}/outboundroutes][%d] postTelephonyProvidersEdgesSiteOutboundroutesNotFound  %+v", 404, o.Payload)
}

func (o *PostTelephonyProvidersEdgesSiteOutboundroutesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesSiteOutboundroutesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgesSiteOutboundroutesRequestEntityTooLarge creates a PostTelephonyProvidersEdgesSiteOutboundroutesRequestEntityTooLarge with default headers values
func NewPostTelephonyProvidersEdgesSiteOutboundroutesRequestEntityTooLarge() *PostTelephonyProvidersEdgesSiteOutboundroutesRequestEntityTooLarge {
	return &PostTelephonyProvidersEdgesSiteOutboundroutesRequestEntityTooLarge{}
}

/*PostTelephonyProvidersEdgesSiteOutboundroutesRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostTelephonyProvidersEdgesSiteOutboundroutesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgesSiteOutboundroutesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/sites/{siteId}/outboundroutes][%d] postTelephonyProvidersEdgesSiteOutboundroutesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostTelephonyProvidersEdgesSiteOutboundroutesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesSiteOutboundroutesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgesSiteOutboundroutesUnsupportedMediaType creates a PostTelephonyProvidersEdgesSiteOutboundroutesUnsupportedMediaType with default headers values
func NewPostTelephonyProvidersEdgesSiteOutboundroutesUnsupportedMediaType() *PostTelephonyProvidersEdgesSiteOutboundroutesUnsupportedMediaType {
	return &PostTelephonyProvidersEdgesSiteOutboundroutesUnsupportedMediaType{}
}

/*PostTelephonyProvidersEdgesSiteOutboundroutesUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostTelephonyProvidersEdgesSiteOutboundroutesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgesSiteOutboundroutesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/sites/{siteId}/outboundroutes][%d] postTelephonyProvidersEdgesSiteOutboundroutesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostTelephonyProvidersEdgesSiteOutboundroutesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesSiteOutboundroutesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgesSiteOutboundroutesTooManyRequests creates a PostTelephonyProvidersEdgesSiteOutboundroutesTooManyRequests with default headers values
func NewPostTelephonyProvidersEdgesSiteOutboundroutesTooManyRequests() *PostTelephonyProvidersEdgesSiteOutboundroutesTooManyRequests {
	return &PostTelephonyProvidersEdgesSiteOutboundroutesTooManyRequests{}
}

/*PostTelephonyProvidersEdgesSiteOutboundroutesTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PostTelephonyProvidersEdgesSiteOutboundroutesTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgesSiteOutboundroutesTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/sites/{siteId}/outboundroutes][%d] postTelephonyProvidersEdgesSiteOutboundroutesTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostTelephonyProvidersEdgesSiteOutboundroutesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesSiteOutboundroutesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgesSiteOutboundroutesInternalServerError creates a PostTelephonyProvidersEdgesSiteOutboundroutesInternalServerError with default headers values
func NewPostTelephonyProvidersEdgesSiteOutboundroutesInternalServerError() *PostTelephonyProvidersEdgesSiteOutboundroutesInternalServerError {
	return &PostTelephonyProvidersEdgesSiteOutboundroutesInternalServerError{}
}

/*PostTelephonyProvidersEdgesSiteOutboundroutesInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostTelephonyProvidersEdgesSiteOutboundroutesInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgesSiteOutboundroutesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/sites/{siteId}/outboundroutes][%d] postTelephonyProvidersEdgesSiteOutboundroutesInternalServerError  %+v", 500, o.Payload)
}

func (o *PostTelephonyProvidersEdgesSiteOutboundroutesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesSiteOutboundroutesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgesSiteOutboundroutesServiceUnavailable creates a PostTelephonyProvidersEdgesSiteOutboundroutesServiceUnavailable with default headers values
func NewPostTelephonyProvidersEdgesSiteOutboundroutesServiceUnavailable() *PostTelephonyProvidersEdgesSiteOutboundroutesServiceUnavailable {
	return &PostTelephonyProvidersEdgesSiteOutboundroutesServiceUnavailable{}
}

/*PostTelephonyProvidersEdgesSiteOutboundroutesServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostTelephonyProvidersEdgesSiteOutboundroutesServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgesSiteOutboundroutesServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/sites/{siteId}/outboundroutes][%d] postTelephonyProvidersEdgesSiteOutboundroutesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostTelephonyProvidersEdgesSiteOutboundroutesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesSiteOutboundroutesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgesSiteOutboundroutesGatewayTimeout creates a PostTelephonyProvidersEdgesSiteOutboundroutesGatewayTimeout with default headers values
func NewPostTelephonyProvidersEdgesSiteOutboundroutesGatewayTimeout() *PostTelephonyProvidersEdgesSiteOutboundroutesGatewayTimeout {
	return &PostTelephonyProvidersEdgesSiteOutboundroutesGatewayTimeout{}
}

/*PostTelephonyProvidersEdgesSiteOutboundroutesGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostTelephonyProvidersEdgesSiteOutboundroutesGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgesSiteOutboundroutesGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/sites/{siteId}/outboundroutes][%d] postTelephonyProvidersEdgesSiteOutboundroutesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostTelephonyProvidersEdgesSiteOutboundroutesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesSiteOutboundroutesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

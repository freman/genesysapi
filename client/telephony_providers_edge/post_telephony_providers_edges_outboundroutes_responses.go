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

// PostTelephonyProvidersEdgesOutboundroutesReader is a Reader for the PostTelephonyProvidersEdgesOutboundroutes structure.
type PostTelephonyProvidersEdgesOutboundroutesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostTelephonyProvidersEdgesOutboundroutesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostTelephonyProvidersEdgesOutboundroutesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostTelephonyProvidersEdgesOutboundroutesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostTelephonyProvidersEdgesOutboundroutesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostTelephonyProvidersEdgesOutboundroutesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostTelephonyProvidersEdgesOutboundroutesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostTelephonyProvidersEdgesOutboundroutesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostTelephonyProvidersEdgesOutboundroutesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostTelephonyProvidersEdgesOutboundroutesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostTelephonyProvidersEdgesOutboundroutesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostTelephonyProvidersEdgesOutboundroutesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostTelephonyProvidersEdgesOutboundroutesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostTelephonyProvidersEdgesOutboundroutesOK creates a PostTelephonyProvidersEdgesOutboundroutesOK with default headers values
func NewPostTelephonyProvidersEdgesOutboundroutesOK() *PostTelephonyProvidersEdgesOutboundroutesOK {
	return &PostTelephonyProvidersEdgesOutboundroutesOK{}
}

/*PostTelephonyProvidersEdgesOutboundroutesOK handles this case with default header values.

successful operation
*/
type PostTelephonyProvidersEdgesOutboundroutesOK struct {
	Payload *models.OutboundRoute
}

func (o *PostTelephonyProvidersEdgesOutboundroutesOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/outboundroutes][%d] postTelephonyProvidersEdgesOutboundroutesOK  %+v", 200, o.Payload)
}

func (o *PostTelephonyProvidersEdgesOutboundroutesOK) GetPayload() *models.OutboundRoute {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesOutboundroutesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OutboundRoute)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgesOutboundroutesBadRequest creates a PostTelephonyProvidersEdgesOutboundroutesBadRequest with default headers values
func NewPostTelephonyProvidersEdgesOutboundroutesBadRequest() *PostTelephonyProvidersEdgesOutboundroutesBadRequest {
	return &PostTelephonyProvidersEdgesOutboundroutesBadRequest{}
}

/*PostTelephonyProvidersEdgesOutboundroutesBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostTelephonyProvidersEdgesOutboundroutesBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgesOutboundroutesBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/outboundroutes][%d] postTelephonyProvidersEdgesOutboundroutesBadRequest  %+v", 400, o.Payload)
}

func (o *PostTelephonyProvidersEdgesOutboundroutesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesOutboundroutesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgesOutboundroutesUnauthorized creates a PostTelephonyProvidersEdgesOutboundroutesUnauthorized with default headers values
func NewPostTelephonyProvidersEdgesOutboundroutesUnauthorized() *PostTelephonyProvidersEdgesOutboundroutesUnauthorized {
	return &PostTelephonyProvidersEdgesOutboundroutesUnauthorized{}
}

/*PostTelephonyProvidersEdgesOutboundroutesUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostTelephonyProvidersEdgesOutboundroutesUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgesOutboundroutesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/outboundroutes][%d] postTelephonyProvidersEdgesOutboundroutesUnauthorized  %+v", 401, o.Payload)
}

func (o *PostTelephonyProvidersEdgesOutboundroutesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesOutboundroutesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgesOutboundroutesForbidden creates a PostTelephonyProvidersEdgesOutboundroutesForbidden with default headers values
func NewPostTelephonyProvidersEdgesOutboundroutesForbidden() *PostTelephonyProvidersEdgesOutboundroutesForbidden {
	return &PostTelephonyProvidersEdgesOutboundroutesForbidden{}
}

/*PostTelephonyProvidersEdgesOutboundroutesForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostTelephonyProvidersEdgesOutboundroutesForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgesOutboundroutesForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/outboundroutes][%d] postTelephonyProvidersEdgesOutboundroutesForbidden  %+v", 403, o.Payload)
}

func (o *PostTelephonyProvidersEdgesOutboundroutesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesOutboundroutesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgesOutboundroutesNotFound creates a PostTelephonyProvidersEdgesOutboundroutesNotFound with default headers values
func NewPostTelephonyProvidersEdgesOutboundroutesNotFound() *PostTelephonyProvidersEdgesOutboundroutesNotFound {
	return &PostTelephonyProvidersEdgesOutboundroutesNotFound{}
}

/*PostTelephonyProvidersEdgesOutboundroutesNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostTelephonyProvidersEdgesOutboundroutesNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgesOutboundroutesNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/outboundroutes][%d] postTelephonyProvidersEdgesOutboundroutesNotFound  %+v", 404, o.Payload)
}

func (o *PostTelephonyProvidersEdgesOutboundroutesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesOutboundroutesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgesOutboundroutesRequestEntityTooLarge creates a PostTelephonyProvidersEdgesOutboundroutesRequestEntityTooLarge with default headers values
func NewPostTelephonyProvidersEdgesOutboundroutesRequestEntityTooLarge() *PostTelephonyProvidersEdgesOutboundroutesRequestEntityTooLarge {
	return &PostTelephonyProvidersEdgesOutboundroutesRequestEntityTooLarge{}
}

/*PostTelephonyProvidersEdgesOutboundroutesRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostTelephonyProvidersEdgesOutboundroutesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgesOutboundroutesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/outboundroutes][%d] postTelephonyProvidersEdgesOutboundroutesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostTelephonyProvidersEdgesOutboundroutesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesOutboundroutesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgesOutboundroutesUnsupportedMediaType creates a PostTelephonyProvidersEdgesOutboundroutesUnsupportedMediaType with default headers values
func NewPostTelephonyProvidersEdgesOutboundroutesUnsupportedMediaType() *PostTelephonyProvidersEdgesOutboundroutesUnsupportedMediaType {
	return &PostTelephonyProvidersEdgesOutboundroutesUnsupportedMediaType{}
}

/*PostTelephonyProvidersEdgesOutboundroutesUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostTelephonyProvidersEdgesOutboundroutesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgesOutboundroutesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/outboundroutes][%d] postTelephonyProvidersEdgesOutboundroutesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostTelephonyProvidersEdgesOutboundroutesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesOutboundroutesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgesOutboundroutesTooManyRequests creates a PostTelephonyProvidersEdgesOutboundroutesTooManyRequests with default headers values
func NewPostTelephonyProvidersEdgesOutboundroutesTooManyRequests() *PostTelephonyProvidersEdgesOutboundroutesTooManyRequests {
	return &PostTelephonyProvidersEdgesOutboundroutesTooManyRequests{}
}

/*PostTelephonyProvidersEdgesOutboundroutesTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PostTelephonyProvidersEdgesOutboundroutesTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgesOutboundroutesTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/outboundroutes][%d] postTelephonyProvidersEdgesOutboundroutesTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostTelephonyProvidersEdgesOutboundroutesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesOutboundroutesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgesOutboundroutesInternalServerError creates a PostTelephonyProvidersEdgesOutboundroutesInternalServerError with default headers values
func NewPostTelephonyProvidersEdgesOutboundroutesInternalServerError() *PostTelephonyProvidersEdgesOutboundroutesInternalServerError {
	return &PostTelephonyProvidersEdgesOutboundroutesInternalServerError{}
}

/*PostTelephonyProvidersEdgesOutboundroutesInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostTelephonyProvidersEdgesOutboundroutesInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgesOutboundroutesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/outboundroutes][%d] postTelephonyProvidersEdgesOutboundroutesInternalServerError  %+v", 500, o.Payload)
}

func (o *PostTelephonyProvidersEdgesOutboundroutesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesOutboundroutesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgesOutboundroutesServiceUnavailable creates a PostTelephonyProvidersEdgesOutboundroutesServiceUnavailable with default headers values
func NewPostTelephonyProvidersEdgesOutboundroutesServiceUnavailable() *PostTelephonyProvidersEdgesOutboundroutesServiceUnavailable {
	return &PostTelephonyProvidersEdgesOutboundroutesServiceUnavailable{}
}

/*PostTelephonyProvidersEdgesOutboundroutesServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostTelephonyProvidersEdgesOutboundroutesServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgesOutboundroutesServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/outboundroutes][%d] postTelephonyProvidersEdgesOutboundroutesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostTelephonyProvidersEdgesOutboundroutesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesOutboundroutesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgesOutboundroutesGatewayTimeout creates a PostTelephonyProvidersEdgesOutboundroutesGatewayTimeout with default headers values
func NewPostTelephonyProvidersEdgesOutboundroutesGatewayTimeout() *PostTelephonyProvidersEdgesOutboundroutesGatewayTimeout {
	return &PostTelephonyProvidersEdgesOutboundroutesGatewayTimeout{}
}

/*PostTelephonyProvidersEdgesOutboundroutesGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostTelephonyProvidersEdgesOutboundroutesGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgesOutboundroutesGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/outboundroutes][%d] postTelephonyProvidersEdgesOutboundroutesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostTelephonyProvidersEdgesOutboundroutesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesOutboundroutesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

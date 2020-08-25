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

// PostTelephonyProvidersEdgeDiagnosticRouteReader is a Reader for the PostTelephonyProvidersEdgeDiagnosticRoute structure.
type PostTelephonyProvidersEdgeDiagnosticRouteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostTelephonyProvidersEdgeDiagnosticRouteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewPostTelephonyProvidersEdgeDiagnosticRouteAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostTelephonyProvidersEdgeDiagnosticRouteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostTelephonyProvidersEdgeDiagnosticRouteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostTelephonyProvidersEdgeDiagnosticRouteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostTelephonyProvidersEdgeDiagnosticRouteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostTelephonyProvidersEdgeDiagnosticRouteRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostTelephonyProvidersEdgeDiagnosticRouteUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostTelephonyProvidersEdgeDiagnosticRouteTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostTelephonyProvidersEdgeDiagnosticRouteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostTelephonyProvidersEdgeDiagnosticRouteServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostTelephonyProvidersEdgeDiagnosticRouteGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostTelephonyProvidersEdgeDiagnosticRouteAccepted creates a PostTelephonyProvidersEdgeDiagnosticRouteAccepted with default headers values
func NewPostTelephonyProvidersEdgeDiagnosticRouteAccepted() *PostTelephonyProvidersEdgeDiagnosticRouteAccepted {
	return &PostTelephonyProvidersEdgeDiagnosticRouteAccepted{}
}

/*PostTelephonyProvidersEdgeDiagnosticRouteAccepted handles this case with default header values.

Request to get network diagnostic has been accepted
*/
type PostTelephonyProvidersEdgeDiagnosticRouteAccepted struct {
	Payload *models.EdgeNetworkDiagnostic
}

func (o *PostTelephonyProvidersEdgeDiagnosticRouteAccepted) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/diagnostic/route][%d] postTelephonyProvidersEdgeDiagnosticRouteAccepted  %+v", 202, o.Payload)
}

func (o *PostTelephonyProvidersEdgeDiagnosticRouteAccepted) GetPayload() *models.EdgeNetworkDiagnostic {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeDiagnosticRouteAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EdgeNetworkDiagnostic)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeDiagnosticRouteBadRequest creates a PostTelephonyProvidersEdgeDiagnosticRouteBadRequest with default headers values
func NewPostTelephonyProvidersEdgeDiagnosticRouteBadRequest() *PostTelephonyProvidersEdgeDiagnosticRouteBadRequest {
	return &PostTelephonyProvidersEdgeDiagnosticRouteBadRequest{}
}

/*PostTelephonyProvidersEdgeDiagnosticRouteBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostTelephonyProvidersEdgeDiagnosticRouteBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgeDiagnosticRouteBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/diagnostic/route][%d] postTelephonyProvidersEdgeDiagnosticRouteBadRequest  %+v", 400, o.Payload)
}

func (o *PostTelephonyProvidersEdgeDiagnosticRouteBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeDiagnosticRouteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeDiagnosticRouteUnauthorized creates a PostTelephonyProvidersEdgeDiagnosticRouteUnauthorized with default headers values
func NewPostTelephonyProvidersEdgeDiagnosticRouteUnauthorized() *PostTelephonyProvidersEdgeDiagnosticRouteUnauthorized {
	return &PostTelephonyProvidersEdgeDiagnosticRouteUnauthorized{}
}

/*PostTelephonyProvidersEdgeDiagnosticRouteUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostTelephonyProvidersEdgeDiagnosticRouteUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgeDiagnosticRouteUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/diagnostic/route][%d] postTelephonyProvidersEdgeDiagnosticRouteUnauthorized  %+v", 401, o.Payload)
}

func (o *PostTelephonyProvidersEdgeDiagnosticRouteUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeDiagnosticRouteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeDiagnosticRouteForbidden creates a PostTelephonyProvidersEdgeDiagnosticRouteForbidden with default headers values
func NewPostTelephonyProvidersEdgeDiagnosticRouteForbidden() *PostTelephonyProvidersEdgeDiagnosticRouteForbidden {
	return &PostTelephonyProvidersEdgeDiagnosticRouteForbidden{}
}

/*PostTelephonyProvidersEdgeDiagnosticRouteForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostTelephonyProvidersEdgeDiagnosticRouteForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgeDiagnosticRouteForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/diagnostic/route][%d] postTelephonyProvidersEdgeDiagnosticRouteForbidden  %+v", 403, o.Payload)
}

func (o *PostTelephonyProvidersEdgeDiagnosticRouteForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeDiagnosticRouteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeDiagnosticRouteNotFound creates a PostTelephonyProvidersEdgeDiagnosticRouteNotFound with default headers values
func NewPostTelephonyProvidersEdgeDiagnosticRouteNotFound() *PostTelephonyProvidersEdgeDiagnosticRouteNotFound {
	return &PostTelephonyProvidersEdgeDiagnosticRouteNotFound{}
}

/*PostTelephonyProvidersEdgeDiagnosticRouteNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostTelephonyProvidersEdgeDiagnosticRouteNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgeDiagnosticRouteNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/diagnostic/route][%d] postTelephonyProvidersEdgeDiagnosticRouteNotFound  %+v", 404, o.Payload)
}

func (o *PostTelephonyProvidersEdgeDiagnosticRouteNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeDiagnosticRouteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeDiagnosticRouteRequestEntityTooLarge creates a PostTelephonyProvidersEdgeDiagnosticRouteRequestEntityTooLarge with default headers values
func NewPostTelephonyProvidersEdgeDiagnosticRouteRequestEntityTooLarge() *PostTelephonyProvidersEdgeDiagnosticRouteRequestEntityTooLarge {
	return &PostTelephonyProvidersEdgeDiagnosticRouteRequestEntityTooLarge{}
}

/*PostTelephonyProvidersEdgeDiagnosticRouteRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostTelephonyProvidersEdgeDiagnosticRouteRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgeDiagnosticRouteRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/diagnostic/route][%d] postTelephonyProvidersEdgeDiagnosticRouteRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostTelephonyProvidersEdgeDiagnosticRouteRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeDiagnosticRouteRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeDiagnosticRouteUnsupportedMediaType creates a PostTelephonyProvidersEdgeDiagnosticRouteUnsupportedMediaType with default headers values
func NewPostTelephonyProvidersEdgeDiagnosticRouteUnsupportedMediaType() *PostTelephonyProvidersEdgeDiagnosticRouteUnsupportedMediaType {
	return &PostTelephonyProvidersEdgeDiagnosticRouteUnsupportedMediaType{}
}

/*PostTelephonyProvidersEdgeDiagnosticRouteUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostTelephonyProvidersEdgeDiagnosticRouteUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgeDiagnosticRouteUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/diagnostic/route][%d] postTelephonyProvidersEdgeDiagnosticRouteUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostTelephonyProvidersEdgeDiagnosticRouteUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeDiagnosticRouteUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeDiagnosticRouteTooManyRequests creates a PostTelephonyProvidersEdgeDiagnosticRouteTooManyRequests with default headers values
func NewPostTelephonyProvidersEdgeDiagnosticRouteTooManyRequests() *PostTelephonyProvidersEdgeDiagnosticRouteTooManyRequests {
	return &PostTelephonyProvidersEdgeDiagnosticRouteTooManyRequests{}
}

/*PostTelephonyProvidersEdgeDiagnosticRouteTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PostTelephonyProvidersEdgeDiagnosticRouteTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgeDiagnosticRouteTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/diagnostic/route][%d] postTelephonyProvidersEdgeDiagnosticRouteTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostTelephonyProvidersEdgeDiagnosticRouteTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeDiagnosticRouteTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeDiagnosticRouteInternalServerError creates a PostTelephonyProvidersEdgeDiagnosticRouteInternalServerError with default headers values
func NewPostTelephonyProvidersEdgeDiagnosticRouteInternalServerError() *PostTelephonyProvidersEdgeDiagnosticRouteInternalServerError {
	return &PostTelephonyProvidersEdgeDiagnosticRouteInternalServerError{}
}

/*PostTelephonyProvidersEdgeDiagnosticRouteInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostTelephonyProvidersEdgeDiagnosticRouteInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgeDiagnosticRouteInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/diagnostic/route][%d] postTelephonyProvidersEdgeDiagnosticRouteInternalServerError  %+v", 500, o.Payload)
}

func (o *PostTelephonyProvidersEdgeDiagnosticRouteInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeDiagnosticRouteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeDiagnosticRouteServiceUnavailable creates a PostTelephonyProvidersEdgeDiagnosticRouteServiceUnavailable with default headers values
func NewPostTelephonyProvidersEdgeDiagnosticRouteServiceUnavailable() *PostTelephonyProvidersEdgeDiagnosticRouteServiceUnavailable {
	return &PostTelephonyProvidersEdgeDiagnosticRouteServiceUnavailable{}
}

/*PostTelephonyProvidersEdgeDiagnosticRouteServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostTelephonyProvidersEdgeDiagnosticRouteServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgeDiagnosticRouteServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/diagnostic/route][%d] postTelephonyProvidersEdgeDiagnosticRouteServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostTelephonyProvidersEdgeDiagnosticRouteServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeDiagnosticRouteServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeDiagnosticRouteGatewayTimeout creates a PostTelephonyProvidersEdgeDiagnosticRouteGatewayTimeout with default headers values
func NewPostTelephonyProvidersEdgeDiagnosticRouteGatewayTimeout() *PostTelephonyProvidersEdgeDiagnosticRouteGatewayTimeout {
	return &PostTelephonyProvidersEdgeDiagnosticRouteGatewayTimeout{}
}

/*PostTelephonyProvidersEdgeDiagnosticRouteGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostTelephonyProvidersEdgeDiagnosticRouteGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgeDiagnosticRouteGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/diagnostic/route][%d] postTelephonyProvidersEdgeDiagnosticRouteGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostTelephonyProvidersEdgeDiagnosticRouteGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeDiagnosticRouteGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
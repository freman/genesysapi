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

// GetTelephonyProvidersEdgeReader is a Reader for the GetTelephonyProvidersEdge structure.
type GetTelephonyProvidersEdgeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTelephonyProvidersEdgeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTelephonyProvidersEdgeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetTelephonyProvidersEdgeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetTelephonyProvidersEdgeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetTelephonyProvidersEdgeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTelephonyProvidersEdgeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetTelephonyProvidersEdgeRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetTelephonyProvidersEdgeUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetTelephonyProvidersEdgeTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetTelephonyProvidersEdgeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetTelephonyProvidersEdgeServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetTelephonyProvidersEdgeGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTelephonyProvidersEdgeOK creates a GetTelephonyProvidersEdgeOK with default headers values
func NewGetTelephonyProvidersEdgeOK() *GetTelephonyProvidersEdgeOK {
	return &GetTelephonyProvidersEdgeOK{}
}

/*GetTelephonyProvidersEdgeOK handles this case with default header values.

successful operation
*/
type GetTelephonyProvidersEdgeOK struct {
	Payload *models.Edge
}

func (o *GetTelephonyProvidersEdgeOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}][%d] getTelephonyProvidersEdgeOK  %+v", 200, o.Payload)
}

func (o *GetTelephonyProvidersEdgeOK) GetPayload() *models.Edge {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Edge)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgeBadRequest creates a GetTelephonyProvidersEdgeBadRequest with default headers values
func NewGetTelephonyProvidersEdgeBadRequest() *GetTelephonyProvidersEdgeBadRequest {
	return &GetTelephonyProvidersEdgeBadRequest{}
}

/*GetTelephonyProvidersEdgeBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetTelephonyProvidersEdgeBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgeBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}][%d] getTelephonyProvidersEdgeBadRequest  %+v", 400, o.Payload)
}

func (o *GetTelephonyProvidersEdgeBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgeUnauthorized creates a GetTelephonyProvidersEdgeUnauthorized with default headers values
func NewGetTelephonyProvidersEdgeUnauthorized() *GetTelephonyProvidersEdgeUnauthorized {
	return &GetTelephonyProvidersEdgeUnauthorized{}
}

/*GetTelephonyProvidersEdgeUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetTelephonyProvidersEdgeUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgeUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}][%d] getTelephonyProvidersEdgeUnauthorized  %+v", 401, o.Payload)
}

func (o *GetTelephonyProvidersEdgeUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgeForbidden creates a GetTelephonyProvidersEdgeForbidden with default headers values
func NewGetTelephonyProvidersEdgeForbidden() *GetTelephonyProvidersEdgeForbidden {
	return &GetTelephonyProvidersEdgeForbidden{}
}

/*GetTelephonyProvidersEdgeForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetTelephonyProvidersEdgeForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgeForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}][%d] getTelephonyProvidersEdgeForbidden  %+v", 403, o.Payload)
}

func (o *GetTelephonyProvidersEdgeForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgeNotFound creates a GetTelephonyProvidersEdgeNotFound with default headers values
func NewGetTelephonyProvidersEdgeNotFound() *GetTelephonyProvidersEdgeNotFound {
	return &GetTelephonyProvidersEdgeNotFound{}
}

/*GetTelephonyProvidersEdgeNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetTelephonyProvidersEdgeNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgeNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}][%d] getTelephonyProvidersEdgeNotFound  %+v", 404, o.Payload)
}

func (o *GetTelephonyProvidersEdgeNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgeRequestEntityTooLarge creates a GetTelephonyProvidersEdgeRequestEntityTooLarge with default headers values
func NewGetTelephonyProvidersEdgeRequestEntityTooLarge() *GetTelephonyProvidersEdgeRequestEntityTooLarge {
	return &GetTelephonyProvidersEdgeRequestEntityTooLarge{}
}

/*GetTelephonyProvidersEdgeRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetTelephonyProvidersEdgeRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgeRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}][%d] getTelephonyProvidersEdgeRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetTelephonyProvidersEdgeRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgeRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgeUnsupportedMediaType creates a GetTelephonyProvidersEdgeUnsupportedMediaType with default headers values
func NewGetTelephonyProvidersEdgeUnsupportedMediaType() *GetTelephonyProvidersEdgeUnsupportedMediaType {
	return &GetTelephonyProvidersEdgeUnsupportedMediaType{}
}

/*GetTelephonyProvidersEdgeUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetTelephonyProvidersEdgeUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgeUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}][%d] getTelephonyProvidersEdgeUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetTelephonyProvidersEdgeUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgeUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgeTooManyRequests creates a GetTelephonyProvidersEdgeTooManyRequests with default headers values
func NewGetTelephonyProvidersEdgeTooManyRequests() *GetTelephonyProvidersEdgeTooManyRequests {
	return &GetTelephonyProvidersEdgeTooManyRequests{}
}

/*GetTelephonyProvidersEdgeTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetTelephonyProvidersEdgeTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgeTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}][%d] getTelephonyProvidersEdgeTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetTelephonyProvidersEdgeTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgeTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgeInternalServerError creates a GetTelephonyProvidersEdgeInternalServerError with default headers values
func NewGetTelephonyProvidersEdgeInternalServerError() *GetTelephonyProvidersEdgeInternalServerError {
	return &GetTelephonyProvidersEdgeInternalServerError{}
}

/*GetTelephonyProvidersEdgeInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetTelephonyProvidersEdgeInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgeInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}][%d] getTelephonyProvidersEdgeInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTelephonyProvidersEdgeInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgeServiceUnavailable creates a GetTelephonyProvidersEdgeServiceUnavailable with default headers values
func NewGetTelephonyProvidersEdgeServiceUnavailable() *GetTelephonyProvidersEdgeServiceUnavailable {
	return &GetTelephonyProvidersEdgeServiceUnavailable{}
}

/*GetTelephonyProvidersEdgeServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetTelephonyProvidersEdgeServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgeServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}][%d] getTelephonyProvidersEdgeServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetTelephonyProvidersEdgeServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgeServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgeGatewayTimeout creates a GetTelephonyProvidersEdgeGatewayTimeout with default headers values
func NewGetTelephonyProvidersEdgeGatewayTimeout() *GetTelephonyProvidersEdgeGatewayTimeout {
	return &GetTelephonyProvidersEdgeGatewayTimeout{}
}

/*GetTelephonyProvidersEdgeGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetTelephonyProvidersEdgeGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgeGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}][%d] getTelephonyProvidersEdgeGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetTelephonyProvidersEdgeGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgeGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
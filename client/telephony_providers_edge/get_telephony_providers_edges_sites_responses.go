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

// GetTelephonyProvidersEdgesSitesReader is a Reader for the GetTelephonyProvidersEdgesSites structure.
type GetTelephonyProvidersEdgesSitesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTelephonyProvidersEdgesSitesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTelephonyProvidersEdgesSitesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetTelephonyProvidersEdgesSitesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetTelephonyProvidersEdgesSitesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetTelephonyProvidersEdgesSitesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTelephonyProvidersEdgesSitesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetTelephonyProvidersEdgesSitesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetTelephonyProvidersEdgesSitesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetTelephonyProvidersEdgesSitesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetTelephonyProvidersEdgesSitesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetTelephonyProvidersEdgesSitesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetTelephonyProvidersEdgesSitesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTelephonyProvidersEdgesSitesOK creates a GetTelephonyProvidersEdgesSitesOK with default headers values
func NewGetTelephonyProvidersEdgesSitesOK() *GetTelephonyProvidersEdgesSitesOK {
	return &GetTelephonyProvidersEdgesSitesOK{}
}

/*GetTelephonyProvidersEdgesSitesOK handles this case with default header values.

successful operation
*/
type GetTelephonyProvidersEdgesSitesOK struct {
	Payload *models.SiteEntityListing
}

func (o *GetTelephonyProvidersEdgesSitesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/sites][%d] getTelephonyProvidersEdgesSitesOK  %+v", 200, o.Payload)
}

func (o *GetTelephonyProvidersEdgesSitesOK) GetPayload() *models.SiteEntityListing {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesSitesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SiteEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesSitesBadRequest creates a GetTelephonyProvidersEdgesSitesBadRequest with default headers values
func NewGetTelephonyProvidersEdgesSitesBadRequest() *GetTelephonyProvidersEdgesSitesBadRequest {
	return &GetTelephonyProvidersEdgesSitesBadRequest{}
}

/*GetTelephonyProvidersEdgesSitesBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetTelephonyProvidersEdgesSitesBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesSitesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/sites][%d] getTelephonyProvidersEdgesSitesBadRequest  %+v", 400, o.Payload)
}

func (o *GetTelephonyProvidersEdgesSitesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesSitesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesSitesUnauthorized creates a GetTelephonyProvidersEdgesSitesUnauthorized with default headers values
func NewGetTelephonyProvidersEdgesSitesUnauthorized() *GetTelephonyProvidersEdgesSitesUnauthorized {
	return &GetTelephonyProvidersEdgesSitesUnauthorized{}
}

/*GetTelephonyProvidersEdgesSitesUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetTelephonyProvidersEdgesSitesUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesSitesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/sites][%d] getTelephonyProvidersEdgesSitesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetTelephonyProvidersEdgesSitesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesSitesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesSitesForbidden creates a GetTelephonyProvidersEdgesSitesForbidden with default headers values
func NewGetTelephonyProvidersEdgesSitesForbidden() *GetTelephonyProvidersEdgesSitesForbidden {
	return &GetTelephonyProvidersEdgesSitesForbidden{}
}

/*GetTelephonyProvidersEdgesSitesForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetTelephonyProvidersEdgesSitesForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesSitesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/sites][%d] getTelephonyProvidersEdgesSitesForbidden  %+v", 403, o.Payload)
}

func (o *GetTelephonyProvidersEdgesSitesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesSitesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesSitesNotFound creates a GetTelephonyProvidersEdgesSitesNotFound with default headers values
func NewGetTelephonyProvidersEdgesSitesNotFound() *GetTelephonyProvidersEdgesSitesNotFound {
	return &GetTelephonyProvidersEdgesSitesNotFound{}
}

/*GetTelephonyProvidersEdgesSitesNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetTelephonyProvidersEdgesSitesNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesSitesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/sites][%d] getTelephonyProvidersEdgesSitesNotFound  %+v", 404, o.Payload)
}

func (o *GetTelephonyProvidersEdgesSitesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesSitesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesSitesRequestEntityTooLarge creates a GetTelephonyProvidersEdgesSitesRequestEntityTooLarge with default headers values
func NewGetTelephonyProvidersEdgesSitesRequestEntityTooLarge() *GetTelephonyProvidersEdgesSitesRequestEntityTooLarge {
	return &GetTelephonyProvidersEdgesSitesRequestEntityTooLarge{}
}

/*GetTelephonyProvidersEdgesSitesRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetTelephonyProvidersEdgesSitesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesSitesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/sites][%d] getTelephonyProvidersEdgesSitesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetTelephonyProvidersEdgesSitesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesSitesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesSitesUnsupportedMediaType creates a GetTelephonyProvidersEdgesSitesUnsupportedMediaType with default headers values
func NewGetTelephonyProvidersEdgesSitesUnsupportedMediaType() *GetTelephonyProvidersEdgesSitesUnsupportedMediaType {
	return &GetTelephonyProvidersEdgesSitesUnsupportedMediaType{}
}

/*GetTelephonyProvidersEdgesSitesUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetTelephonyProvidersEdgesSitesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesSitesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/sites][%d] getTelephonyProvidersEdgesSitesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetTelephonyProvidersEdgesSitesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesSitesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesSitesTooManyRequests creates a GetTelephonyProvidersEdgesSitesTooManyRequests with default headers values
func NewGetTelephonyProvidersEdgesSitesTooManyRequests() *GetTelephonyProvidersEdgesSitesTooManyRequests {
	return &GetTelephonyProvidersEdgesSitesTooManyRequests{}
}

/*GetTelephonyProvidersEdgesSitesTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetTelephonyProvidersEdgesSitesTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesSitesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/sites][%d] getTelephonyProvidersEdgesSitesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetTelephonyProvidersEdgesSitesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesSitesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesSitesInternalServerError creates a GetTelephonyProvidersEdgesSitesInternalServerError with default headers values
func NewGetTelephonyProvidersEdgesSitesInternalServerError() *GetTelephonyProvidersEdgesSitesInternalServerError {
	return &GetTelephonyProvidersEdgesSitesInternalServerError{}
}

/*GetTelephonyProvidersEdgesSitesInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetTelephonyProvidersEdgesSitesInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesSitesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/sites][%d] getTelephonyProvidersEdgesSitesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTelephonyProvidersEdgesSitesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesSitesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesSitesServiceUnavailable creates a GetTelephonyProvidersEdgesSitesServiceUnavailable with default headers values
func NewGetTelephonyProvidersEdgesSitesServiceUnavailable() *GetTelephonyProvidersEdgesSitesServiceUnavailable {
	return &GetTelephonyProvidersEdgesSitesServiceUnavailable{}
}

/*GetTelephonyProvidersEdgesSitesServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetTelephonyProvidersEdgesSitesServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesSitesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/sites][%d] getTelephonyProvidersEdgesSitesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetTelephonyProvidersEdgesSitesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesSitesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesSitesGatewayTimeout creates a GetTelephonyProvidersEdgesSitesGatewayTimeout with default headers values
func NewGetTelephonyProvidersEdgesSitesGatewayTimeout() *GetTelephonyProvidersEdgesSitesGatewayTimeout {
	return &GetTelephonyProvidersEdgesSitesGatewayTimeout{}
}

/*GetTelephonyProvidersEdgesSitesGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetTelephonyProvidersEdgesSitesGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesSitesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/sites][%d] getTelephonyProvidersEdgesSitesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetTelephonyProvidersEdgesSitesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesSitesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

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

// GetTelephonyProvidersEdgePhysicalinterfacesReader is a Reader for the GetTelephonyProvidersEdgePhysicalinterfaces structure.
type GetTelephonyProvidersEdgePhysicalinterfacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTelephonyProvidersEdgePhysicalinterfacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTelephonyProvidersEdgePhysicalinterfacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetTelephonyProvidersEdgePhysicalinterfacesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetTelephonyProvidersEdgePhysicalinterfacesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetTelephonyProvidersEdgePhysicalinterfacesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTelephonyProvidersEdgePhysicalinterfacesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetTelephonyProvidersEdgePhysicalinterfacesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetTelephonyProvidersEdgePhysicalinterfacesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetTelephonyProvidersEdgePhysicalinterfacesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetTelephonyProvidersEdgePhysicalinterfacesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetTelephonyProvidersEdgePhysicalinterfacesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetTelephonyProvidersEdgePhysicalinterfacesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTelephonyProvidersEdgePhysicalinterfacesOK creates a GetTelephonyProvidersEdgePhysicalinterfacesOK with default headers values
func NewGetTelephonyProvidersEdgePhysicalinterfacesOK() *GetTelephonyProvidersEdgePhysicalinterfacesOK {
	return &GetTelephonyProvidersEdgePhysicalinterfacesOK{}
}

/*GetTelephonyProvidersEdgePhysicalinterfacesOK handles this case with default header values.

successful operation
*/
type GetTelephonyProvidersEdgePhysicalinterfacesOK struct {
	Payload *models.PhysicalInterfaceEntityListing
}

func (o *GetTelephonyProvidersEdgePhysicalinterfacesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/physicalinterfaces][%d] getTelephonyProvidersEdgePhysicalinterfacesOK  %+v", 200, o.Payload)
}

func (o *GetTelephonyProvidersEdgePhysicalinterfacesOK) GetPayload() *models.PhysicalInterfaceEntityListing {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgePhysicalinterfacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PhysicalInterfaceEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgePhysicalinterfacesBadRequest creates a GetTelephonyProvidersEdgePhysicalinterfacesBadRequest with default headers values
func NewGetTelephonyProvidersEdgePhysicalinterfacesBadRequest() *GetTelephonyProvidersEdgePhysicalinterfacesBadRequest {
	return &GetTelephonyProvidersEdgePhysicalinterfacesBadRequest{}
}

/*GetTelephonyProvidersEdgePhysicalinterfacesBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetTelephonyProvidersEdgePhysicalinterfacesBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgePhysicalinterfacesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/physicalinterfaces][%d] getTelephonyProvidersEdgePhysicalinterfacesBadRequest  %+v", 400, o.Payload)
}

func (o *GetTelephonyProvidersEdgePhysicalinterfacesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgePhysicalinterfacesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgePhysicalinterfacesUnauthorized creates a GetTelephonyProvidersEdgePhysicalinterfacesUnauthorized with default headers values
func NewGetTelephonyProvidersEdgePhysicalinterfacesUnauthorized() *GetTelephonyProvidersEdgePhysicalinterfacesUnauthorized {
	return &GetTelephonyProvidersEdgePhysicalinterfacesUnauthorized{}
}

/*GetTelephonyProvidersEdgePhysicalinterfacesUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetTelephonyProvidersEdgePhysicalinterfacesUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgePhysicalinterfacesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/physicalinterfaces][%d] getTelephonyProvidersEdgePhysicalinterfacesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetTelephonyProvidersEdgePhysicalinterfacesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgePhysicalinterfacesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgePhysicalinterfacesForbidden creates a GetTelephonyProvidersEdgePhysicalinterfacesForbidden with default headers values
func NewGetTelephonyProvidersEdgePhysicalinterfacesForbidden() *GetTelephonyProvidersEdgePhysicalinterfacesForbidden {
	return &GetTelephonyProvidersEdgePhysicalinterfacesForbidden{}
}

/*GetTelephonyProvidersEdgePhysicalinterfacesForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetTelephonyProvidersEdgePhysicalinterfacesForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgePhysicalinterfacesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/physicalinterfaces][%d] getTelephonyProvidersEdgePhysicalinterfacesForbidden  %+v", 403, o.Payload)
}

func (o *GetTelephonyProvidersEdgePhysicalinterfacesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgePhysicalinterfacesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgePhysicalinterfacesNotFound creates a GetTelephonyProvidersEdgePhysicalinterfacesNotFound with default headers values
func NewGetTelephonyProvidersEdgePhysicalinterfacesNotFound() *GetTelephonyProvidersEdgePhysicalinterfacesNotFound {
	return &GetTelephonyProvidersEdgePhysicalinterfacesNotFound{}
}

/*GetTelephonyProvidersEdgePhysicalinterfacesNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetTelephonyProvidersEdgePhysicalinterfacesNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgePhysicalinterfacesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/physicalinterfaces][%d] getTelephonyProvidersEdgePhysicalinterfacesNotFound  %+v", 404, o.Payload)
}

func (o *GetTelephonyProvidersEdgePhysicalinterfacesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgePhysicalinterfacesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgePhysicalinterfacesRequestEntityTooLarge creates a GetTelephonyProvidersEdgePhysicalinterfacesRequestEntityTooLarge with default headers values
func NewGetTelephonyProvidersEdgePhysicalinterfacesRequestEntityTooLarge() *GetTelephonyProvidersEdgePhysicalinterfacesRequestEntityTooLarge {
	return &GetTelephonyProvidersEdgePhysicalinterfacesRequestEntityTooLarge{}
}

/*GetTelephonyProvidersEdgePhysicalinterfacesRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetTelephonyProvidersEdgePhysicalinterfacesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgePhysicalinterfacesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/physicalinterfaces][%d] getTelephonyProvidersEdgePhysicalinterfacesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetTelephonyProvidersEdgePhysicalinterfacesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgePhysicalinterfacesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgePhysicalinterfacesUnsupportedMediaType creates a GetTelephonyProvidersEdgePhysicalinterfacesUnsupportedMediaType with default headers values
func NewGetTelephonyProvidersEdgePhysicalinterfacesUnsupportedMediaType() *GetTelephonyProvidersEdgePhysicalinterfacesUnsupportedMediaType {
	return &GetTelephonyProvidersEdgePhysicalinterfacesUnsupportedMediaType{}
}

/*GetTelephonyProvidersEdgePhysicalinterfacesUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetTelephonyProvidersEdgePhysicalinterfacesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgePhysicalinterfacesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/physicalinterfaces][%d] getTelephonyProvidersEdgePhysicalinterfacesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetTelephonyProvidersEdgePhysicalinterfacesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgePhysicalinterfacesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgePhysicalinterfacesTooManyRequests creates a GetTelephonyProvidersEdgePhysicalinterfacesTooManyRequests with default headers values
func NewGetTelephonyProvidersEdgePhysicalinterfacesTooManyRequests() *GetTelephonyProvidersEdgePhysicalinterfacesTooManyRequests {
	return &GetTelephonyProvidersEdgePhysicalinterfacesTooManyRequests{}
}

/*GetTelephonyProvidersEdgePhysicalinterfacesTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetTelephonyProvidersEdgePhysicalinterfacesTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgePhysicalinterfacesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/physicalinterfaces][%d] getTelephonyProvidersEdgePhysicalinterfacesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetTelephonyProvidersEdgePhysicalinterfacesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgePhysicalinterfacesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgePhysicalinterfacesInternalServerError creates a GetTelephonyProvidersEdgePhysicalinterfacesInternalServerError with default headers values
func NewGetTelephonyProvidersEdgePhysicalinterfacesInternalServerError() *GetTelephonyProvidersEdgePhysicalinterfacesInternalServerError {
	return &GetTelephonyProvidersEdgePhysicalinterfacesInternalServerError{}
}

/*GetTelephonyProvidersEdgePhysicalinterfacesInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetTelephonyProvidersEdgePhysicalinterfacesInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgePhysicalinterfacesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/physicalinterfaces][%d] getTelephonyProvidersEdgePhysicalinterfacesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTelephonyProvidersEdgePhysicalinterfacesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgePhysicalinterfacesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgePhysicalinterfacesServiceUnavailable creates a GetTelephonyProvidersEdgePhysicalinterfacesServiceUnavailable with default headers values
func NewGetTelephonyProvidersEdgePhysicalinterfacesServiceUnavailable() *GetTelephonyProvidersEdgePhysicalinterfacesServiceUnavailable {
	return &GetTelephonyProvidersEdgePhysicalinterfacesServiceUnavailable{}
}

/*GetTelephonyProvidersEdgePhysicalinterfacesServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetTelephonyProvidersEdgePhysicalinterfacesServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgePhysicalinterfacesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/physicalinterfaces][%d] getTelephonyProvidersEdgePhysicalinterfacesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetTelephonyProvidersEdgePhysicalinterfacesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgePhysicalinterfacesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgePhysicalinterfacesGatewayTimeout creates a GetTelephonyProvidersEdgePhysicalinterfacesGatewayTimeout with default headers values
func NewGetTelephonyProvidersEdgePhysicalinterfacesGatewayTimeout() *GetTelephonyProvidersEdgePhysicalinterfacesGatewayTimeout {
	return &GetTelephonyProvidersEdgePhysicalinterfacesGatewayTimeout{}
}

/*GetTelephonyProvidersEdgePhysicalinterfacesGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetTelephonyProvidersEdgePhysicalinterfacesGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgePhysicalinterfacesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/physicalinterfaces][%d] getTelephonyProvidersEdgePhysicalinterfacesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetTelephonyProvidersEdgePhysicalinterfacesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgePhysicalinterfacesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

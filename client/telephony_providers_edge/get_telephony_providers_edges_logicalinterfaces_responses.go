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

// GetTelephonyProvidersEdgesLogicalinterfacesReader is a Reader for the GetTelephonyProvidersEdgesLogicalinterfaces structure.
type GetTelephonyProvidersEdgesLogicalinterfacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTelephonyProvidersEdgesLogicalinterfacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTelephonyProvidersEdgesLogicalinterfacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetTelephonyProvidersEdgesLogicalinterfacesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetTelephonyProvidersEdgesLogicalinterfacesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetTelephonyProvidersEdgesLogicalinterfacesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTelephonyProvidersEdgesLogicalinterfacesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetTelephonyProvidersEdgesLogicalinterfacesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetTelephonyProvidersEdgesLogicalinterfacesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetTelephonyProvidersEdgesLogicalinterfacesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetTelephonyProvidersEdgesLogicalinterfacesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetTelephonyProvidersEdgesLogicalinterfacesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetTelephonyProvidersEdgesLogicalinterfacesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTelephonyProvidersEdgesLogicalinterfacesOK creates a GetTelephonyProvidersEdgesLogicalinterfacesOK with default headers values
func NewGetTelephonyProvidersEdgesLogicalinterfacesOK() *GetTelephonyProvidersEdgesLogicalinterfacesOK {
	return &GetTelephonyProvidersEdgesLogicalinterfacesOK{}
}

/*GetTelephonyProvidersEdgesLogicalinterfacesOK handles this case with default header values.

successful operation
*/
type GetTelephonyProvidersEdgesLogicalinterfacesOK struct {
	Payload *models.LogicalInterfaceEntityListing
}

func (o *GetTelephonyProvidersEdgesLogicalinterfacesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/logicalinterfaces][%d] getTelephonyProvidersEdgesLogicalinterfacesOK  %+v", 200, o.Payload)
}

func (o *GetTelephonyProvidersEdgesLogicalinterfacesOK) GetPayload() *models.LogicalInterfaceEntityListing {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesLogicalinterfacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LogicalInterfaceEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesLogicalinterfacesBadRequest creates a GetTelephonyProvidersEdgesLogicalinterfacesBadRequest with default headers values
func NewGetTelephonyProvidersEdgesLogicalinterfacesBadRequest() *GetTelephonyProvidersEdgesLogicalinterfacesBadRequest {
	return &GetTelephonyProvidersEdgesLogicalinterfacesBadRequest{}
}

/*GetTelephonyProvidersEdgesLogicalinterfacesBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetTelephonyProvidersEdgesLogicalinterfacesBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesLogicalinterfacesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/logicalinterfaces][%d] getTelephonyProvidersEdgesLogicalinterfacesBadRequest  %+v", 400, o.Payload)
}

func (o *GetTelephonyProvidersEdgesLogicalinterfacesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesLogicalinterfacesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesLogicalinterfacesUnauthorized creates a GetTelephonyProvidersEdgesLogicalinterfacesUnauthorized with default headers values
func NewGetTelephonyProvidersEdgesLogicalinterfacesUnauthorized() *GetTelephonyProvidersEdgesLogicalinterfacesUnauthorized {
	return &GetTelephonyProvidersEdgesLogicalinterfacesUnauthorized{}
}

/*GetTelephonyProvidersEdgesLogicalinterfacesUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetTelephonyProvidersEdgesLogicalinterfacesUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesLogicalinterfacesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/logicalinterfaces][%d] getTelephonyProvidersEdgesLogicalinterfacesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetTelephonyProvidersEdgesLogicalinterfacesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesLogicalinterfacesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesLogicalinterfacesForbidden creates a GetTelephonyProvidersEdgesLogicalinterfacesForbidden with default headers values
func NewGetTelephonyProvidersEdgesLogicalinterfacesForbidden() *GetTelephonyProvidersEdgesLogicalinterfacesForbidden {
	return &GetTelephonyProvidersEdgesLogicalinterfacesForbidden{}
}

/*GetTelephonyProvidersEdgesLogicalinterfacesForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetTelephonyProvidersEdgesLogicalinterfacesForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesLogicalinterfacesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/logicalinterfaces][%d] getTelephonyProvidersEdgesLogicalinterfacesForbidden  %+v", 403, o.Payload)
}

func (o *GetTelephonyProvidersEdgesLogicalinterfacesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesLogicalinterfacesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesLogicalinterfacesNotFound creates a GetTelephonyProvidersEdgesLogicalinterfacesNotFound with default headers values
func NewGetTelephonyProvidersEdgesLogicalinterfacesNotFound() *GetTelephonyProvidersEdgesLogicalinterfacesNotFound {
	return &GetTelephonyProvidersEdgesLogicalinterfacesNotFound{}
}

/*GetTelephonyProvidersEdgesLogicalinterfacesNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetTelephonyProvidersEdgesLogicalinterfacesNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesLogicalinterfacesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/logicalinterfaces][%d] getTelephonyProvidersEdgesLogicalinterfacesNotFound  %+v", 404, o.Payload)
}

func (o *GetTelephonyProvidersEdgesLogicalinterfacesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesLogicalinterfacesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesLogicalinterfacesRequestEntityTooLarge creates a GetTelephonyProvidersEdgesLogicalinterfacesRequestEntityTooLarge with default headers values
func NewGetTelephonyProvidersEdgesLogicalinterfacesRequestEntityTooLarge() *GetTelephonyProvidersEdgesLogicalinterfacesRequestEntityTooLarge {
	return &GetTelephonyProvidersEdgesLogicalinterfacesRequestEntityTooLarge{}
}

/*GetTelephonyProvidersEdgesLogicalinterfacesRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetTelephonyProvidersEdgesLogicalinterfacesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesLogicalinterfacesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/logicalinterfaces][%d] getTelephonyProvidersEdgesLogicalinterfacesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetTelephonyProvidersEdgesLogicalinterfacesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesLogicalinterfacesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesLogicalinterfacesUnsupportedMediaType creates a GetTelephonyProvidersEdgesLogicalinterfacesUnsupportedMediaType with default headers values
func NewGetTelephonyProvidersEdgesLogicalinterfacesUnsupportedMediaType() *GetTelephonyProvidersEdgesLogicalinterfacesUnsupportedMediaType {
	return &GetTelephonyProvidersEdgesLogicalinterfacesUnsupportedMediaType{}
}

/*GetTelephonyProvidersEdgesLogicalinterfacesUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetTelephonyProvidersEdgesLogicalinterfacesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesLogicalinterfacesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/logicalinterfaces][%d] getTelephonyProvidersEdgesLogicalinterfacesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetTelephonyProvidersEdgesLogicalinterfacesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesLogicalinterfacesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesLogicalinterfacesTooManyRequests creates a GetTelephonyProvidersEdgesLogicalinterfacesTooManyRequests with default headers values
func NewGetTelephonyProvidersEdgesLogicalinterfacesTooManyRequests() *GetTelephonyProvidersEdgesLogicalinterfacesTooManyRequests {
	return &GetTelephonyProvidersEdgesLogicalinterfacesTooManyRequests{}
}

/*GetTelephonyProvidersEdgesLogicalinterfacesTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetTelephonyProvidersEdgesLogicalinterfacesTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesLogicalinterfacesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/logicalinterfaces][%d] getTelephonyProvidersEdgesLogicalinterfacesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetTelephonyProvidersEdgesLogicalinterfacesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesLogicalinterfacesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesLogicalinterfacesInternalServerError creates a GetTelephonyProvidersEdgesLogicalinterfacesInternalServerError with default headers values
func NewGetTelephonyProvidersEdgesLogicalinterfacesInternalServerError() *GetTelephonyProvidersEdgesLogicalinterfacesInternalServerError {
	return &GetTelephonyProvidersEdgesLogicalinterfacesInternalServerError{}
}

/*GetTelephonyProvidersEdgesLogicalinterfacesInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetTelephonyProvidersEdgesLogicalinterfacesInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesLogicalinterfacesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/logicalinterfaces][%d] getTelephonyProvidersEdgesLogicalinterfacesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTelephonyProvidersEdgesLogicalinterfacesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesLogicalinterfacesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesLogicalinterfacesServiceUnavailable creates a GetTelephonyProvidersEdgesLogicalinterfacesServiceUnavailable with default headers values
func NewGetTelephonyProvidersEdgesLogicalinterfacesServiceUnavailable() *GetTelephonyProvidersEdgesLogicalinterfacesServiceUnavailable {
	return &GetTelephonyProvidersEdgesLogicalinterfacesServiceUnavailable{}
}

/*GetTelephonyProvidersEdgesLogicalinterfacesServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetTelephonyProvidersEdgesLogicalinterfacesServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesLogicalinterfacesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/logicalinterfaces][%d] getTelephonyProvidersEdgesLogicalinterfacesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetTelephonyProvidersEdgesLogicalinterfacesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesLogicalinterfacesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesLogicalinterfacesGatewayTimeout creates a GetTelephonyProvidersEdgesLogicalinterfacesGatewayTimeout with default headers values
func NewGetTelephonyProvidersEdgesLogicalinterfacesGatewayTimeout() *GetTelephonyProvidersEdgesLogicalinterfacesGatewayTimeout {
	return &GetTelephonyProvidersEdgesLogicalinterfacesGatewayTimeout{}
}

/*GetTelephonyProvidersEdgesLogicalinterfacesGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetTelephonyProvidersEdgesLogicalinterfacesGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesLogicalinterfacesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/logicalinterfaces][%d] getTelephonyProvidersEdgesLogicalinterfacesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetTelephonyProvidersEdgesLogicalinterfacesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesLogicalinterfacesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

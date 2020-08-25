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

// GetTelephonyProvidersEdgePhysicalinterfaceReader is a Reader for the GetTelephonyProvidersEdgePhysicalinterface structure.
type GetTelephonyProvidersEdgePhysicalinterfaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTelephonyProvidersEdgePhysicalinterfaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTelephonyProvidersEdgePhysicalinterfaceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetTelephonyProvidersEdgePhysicalinterfaceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetTelephonyProvidersEdgePhysicalinterfaceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetTelephonyProvidersEdgePhysicalinterfaceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTelephonyProvidersEdgePhysicalinterfaceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetTelephonyProvidersEdgePhysicalinterfaceRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetTelephonyProvidersEdgePhysicalinterfaceUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetTelephonyProvidersEdgePhysicalinterfaceTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetTelephonyProvidersEdgePhysicalinterfaceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetTelephonyProvidersEdgePhysicalinterfaceServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetTelephonyProvidersEdgePhysicalinterfaceGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTelephonyProvidersEdgePhysicalinterfaceOK creates a GetTelephonyProvidersEdgePhysicalinterfaceOK with default headers values
func NewGetTelephonyProvidersEdgePhysicalinterfaceOK() *GetTelephonyProvidersEdgePhysicalinterfaceOK {
	return &GetTelephonyProvidersEdgePhysicalinterfaceOK{}
}

/*GetTelephonyProvidersEdgePhysicalinterfaceOK handles this case with default header values.

successful operation
*/
type GetTelephonyProvidersEdgePhysicalinterfaceOK struct {
	Payload *models.DomainPhysicalInterface
}

func (o *GetTelephonyProvidersEdgePhysicalinterfaceOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/physicalinterfaces/{interfaceId}][%d] getTelephonyProvidersEdgePhysicalinterfaceOK  %+v", 200, o.Payload)
}

func (o *GetTelephonyProvidersEdgePhysicalinterfaceOK) GetPayload() *models.DomainPhysicalInterface {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgePhysicalinterfaceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DomainPhysicalInterface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgePhysicalinterfaceBadRequest creates a GetTelephonyProvidersEdgePhysicalinterfaceBadRequest with default headers values
func NewGetTelephonyProvidersEdgePhysicalinterfaceBadRequest() *GetTelephonyProvidersEdgePhysicalinterfaceBadRequest {
	return &GetTelephonyProvidersEdgePhysicalinterfaceBadRequest{}
}

/*GetTelephonyProvidersEdgePhysicalinterfaceBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetTelephonyProvidersEdgePhysicalinterfaceBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgePhysicalinterfaceBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/physicalinterfaces/{interfaceId}][%d] getTelephonyProvidersEdgePhysicalinterfaceBadRequest  %+v", 400, o.Payload)
}

func (o *GetTelephonyProvidersEdgePhysicalinterfaceBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgePhysicalinterfaceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgePhysicalinterfaceUnauthorized creates a GetTelephonyProvidersEdgePhysicalinterfaceUnauthorized with default headers values
func NewGetTelephonyProvidersEdgePhysicalinterfaceUnauthorized() *GetTelephonyProvidersEdgePhysicalinterfaceUnauthorized {
	return &GetTelephonyProvidersEdgePhysicalinterfaceUnauthorized{}
}

/*GetTelephonyProvidersEdgePhysicalinterfaceUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetTelephonyProvidersEdgePhysicalinterfaceUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgePhysicalinterfaceUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/physicalinterfaces/{interfaceId}][%d] getTelephonyProvidersEdgePhysicalinterfaceUnauthorized  %+v", 401, o.Payload)
}

func (o *GetTelephonyProvidersEdgePhysicalinterfaceUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgePhysicalinterfaceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgePhysicalinterfaceForbidden creates a GetTelephonyProvidersEdgePhysicalinterfaceForbidden with default headers values
func NewGetTelephonyProvidersEdgePhysicalinterfaceForbidden() *GetTelephonyProvidersEdgePhysicalinterfaceForbidden {
	return &GetTelephonyProvidersEdgePhysicalinterfaceForbidden{}
}

/*GetTelephonyProvidersEdgePhysicalinterfaceForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetTelephonyProvidersEdgePhysicalinterfaceForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgePhysicalinterfaceForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/physicalinterfaces/{interfaceId}][%d] getTelephonyProvidersEdgePhysicalinterfaceForbidden  %+v", 403, o.Payload)
}

func (o *GetTelephonyProvidersEdgePhysicalinterfaceForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgePhysicalinterfaceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgePhysicalinterfaceNotFound creates a GetTelephonyProvidersEdgePhysicalinterfaceNotFound with default headers values
func NewGetTelephonyProvidersEdgePhysicalinterfaceNotFound() *GetTelephonyProvidersEdgePhysicalinterfaceNotFound {
	return &GetTelephonyProvidersEdgePhysicalinterfaceNotFound{}
}

/*GetTelephonyProvidersEdgePhysicalinterfaceNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetTelephonyProvidersEdgePhysicalinterfaceNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgePhysicalinterfaceNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/physicalinterfaces/{interfaceId}][%d] getTelephonyProvidersEdgePhysicalinterfaceNotFound  %+v", 404, o.Payload)
}

func (o *GetTelephonyProvidersEdgePhysicalinterfaceNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgePhysicalinterfaceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgePhysicalinterfaceRequestEntityTooLarge creates a GetTelephonyProvidersEdgePhysicalinterfaceRequestEntityTooLarge with default headers values
func NewGetTelephonyProvidersEdgePhysicalinterfaceRequestEntityTooLarge() *GetTelephonyProvidersEdgePhysicalinterfaceRequestEntityTooLarge {
	return &GetTelephonyProvidersEdgePhysicalinterfaceRequestEntityTooLarge{}
}

/*GetTelephonyProvidersEdgePhysicalinterfaceRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetTelephonyProvidersEdgePhysicalinterfaceRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgePhysicalinterfaceRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/physicalinterfaces/{interfaceId}][%d] getTelephonyProvidersEdgePhysicalinterfaceRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetTelephonyProvidersEdgePhysicalinterfaceRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgePhysicalinterfaceRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgePhysicalinterfaceUnsupportedMediaType creates a GetTelephonyProvidersEdgePhysicalinterfaceUnsupportedMediaType with default headers values
func NewGetTelephonyProvidersEdgePhysicalinterfaceUnsupportedMediaType() *GetTelephonyProvidersEdgePhysicalinterfaceUnsupportedMediaType {
	return &GetTelephonyProvidersEdgePhysicalinterfaceUnsupportedMediaType{}
}

/*GetTelephonyProvidersEdgePhysicalinterfaceUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetTelephonyProvidersEdgePhysicalinterfaceUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgePhysicalinterfaceUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/physicalinterfaces/{interfaceId}][%d] getTelephonyProvidersEdgePhysicalinterfaceUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetTelephonyProvidersEdgePhysicalinterfaceUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgePhysicalinterfaceUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgePhysicalinterfaceTooManyRequests creates a GetTelephonyProvidersEdgePhysicalinterfaceTooManyRequests with default headers values
func NewGetTelephonyProvidersEdgePhysicalinterfaceTooManyRequests() *GetTelephonyProvidersEdgePhysicalinterfaceTooManyRequests {
	return &GetTelephonyProvidersEdgePhysicalinterfaceTooManyRequests{}
}

/*GetTelephonyProvidersEdgePhysicalinterfaceTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetTelephonyProvidersEdgePhysicalinterfaceTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgePhysicalinterfaceTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/physicalinterfaces/{interfaceId}][%d] getTelephonyProvidersEdgePhysicalinterfaceTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetTelephonyProvidersEdgePhysicalinterfaceTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgePhysicalinterfaceTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgePhysicalinterfaceInternalServerError creates a GetTelephonyProvidersEdgePhysicalinterfaceInternalServerError with default headers values
func NewGetTelephonyProvidersEdgePhysicalinterfaceInternalServerError() *GetTelephonyProvidersEdgePhysicalinterfaceInternalServerError {
	return &GetTelephonyProvidersEdgePhysicalinterfaceInternalServerError{}
}

/*GetTelephonyProvidersEdgePhysicalinterfaceInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetTelephonyProvidersEdgePhysicalinterfaceInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgePhysicalinterfaceInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/physicalinterfaces/{interfaceId}][%d] getTelephonyProvidersEdgePhysicalinterfaceInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTelephonyProvidersEdgePhysicalinterfaceInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgePhysicalinterfaceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgePhysicalinterfaceServiceUnavailable creates a GetTelephonyProvidersEdgePhysicalinterfaceServiceUnavailable with default headers values
func NewGetTelephonyProvidersEdgePhysicalinterfaceServiceUnavailable() *GetTelephonyProvidersEdgePhysicalinterfaceServiceUnavailable {
	return &GetTelephonyProvidersEdgePhysicalinterfaceServiceUnavailable{}
}

/*GetTelephonyProvidersEdgePhysicalinterfaceServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetTelephonyProvidersEdgePhysicalinterfaceServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgePhysicalinterfaceServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/physicalinterfaces/{interfaceId}][%d] getTelephonyProvidersEdgePhysicalinterfaceServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetTelephonyProvidersEdgePhysicalinterfaceServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgePhysicalinterfaceServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgePhysicalinterfaceGatewayTimeout creates a GetTelephonyProvidersEdgePhysicalinterfaceGatewayTimeout with default headers values
func NewGetTelephonyProvidersEdgePhysicalinterfaceGatewayTimeout() *GetTelephonyProvidersEdgePhysicalinterfaceGatewayTimeout {
	return &GetTelephonyProvidersEdgePhysicalinterfaceGatewayTimeout{}
}

/*GetTelephonyProvidersEdgePhysicalinterfaceGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetTelephonyProvidersEdgePhysicalinterfaceGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgePhysicalinterfaceGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/physicalinterfaces/{interfaceId}][%d] getTelephonyProvidersEdgePhysicalinterfaceGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetTelephonyProvidersEdgePhysicalinterfaceGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgePhysicalinterfaceGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
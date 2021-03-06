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

// DeleteTelephonyProvidersEdgeLogicalinterfaceReader is a Reader for the DeleteTelephonyProvidersEdgeLogicalinterface structure.
type DeleteTelephonyProvidersEdgeLogicalinterfaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 400:
		result := NewDeleteTelephonyProvidersEdgeLogicalinterfaceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteTelephonyProvidersEdgeLogicalinterfaceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteTelephonyProvidersEdgeLogicalinterfaceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteTelephonyProvidersEdgeLogicalinterfaceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteTelephonyProvidersEdgeLogicalinterfaceRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteTelephonyProvidersEdgeLogicalinterfaceUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteTelephonyProvidersEdgeLogicalinterfaceTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteTelephonyProvidersEdgeLogicalinterfaceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteTelephonyProvidersEdgeLogicalinterfaceServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteTelephonyProvidersEdgeLogicalinterfaceGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteTelephonyProvidersEdgeLogicalinterfaceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteTelephonyProvidersEdgeLogicalinterfaceBadRequest creates a DeleteTelephonyProvidersEdgeLogicalinterfaceBadRequest with default headers values
func NewDeleteTelephonyProvidersEdgeLogicalinterfaceBadRequest() *DeleteTelephonyProvidersEdgeLogicalinterfaceBadRequest {
	return &DeleteTelephonyProvidersEdgeLogicalinterfaceBadRequest{}
}

/*DeleteTelephonyProvidersEdgeLogicalinterfaceBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteTelephonyProvidersEdgeLogicalinterfaceBadRequest struct {
	Payload *models.ErrorBody
}

func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces/{interfaceId}][%d] deleteTelephonyProvidersEdgeLogicalinterfaceBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTelephonyProvidersEdgeLogicalinterfaceUnauthorized creates a DeleteTelephonyProvidersEdgeLogicalinterfaceUnauthorized with default headers values
func NewDeleteTelephonyProvidersEdgeLogicalinterfaceUnauthorized() *DeleteTelephonyProvidersEdgeLogicalinterfaceUnauthorized {
	return &DeleteTelephonyProvidersEdgeLogicalinterfaceUnauthorized{}
}

/*DeleteTelephonyProvidersEdgeLogicalinterfaceUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteTelephonyProvidersEdgeLogicalinterfaceUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces/{interfaceId}][%d] deleteTelephonyProvidersEdgeLogicalinterfaceUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTelephonyProvidersEdgeLogicalinterfaceForbidden creates a DeleteTelephonyProvidersEdgeLogicalinterfaceForbidden with default headers values
func NewDeleteTelephonyProvidersEdgeLogicalinterfaceForbidden() *DeleteTelephonyProvidersEdgeLogicalinterfaceForbidden {
	return &DeleteTelephonyProvidersEdgeLogicalinterfaceForbidden{}
}

/*DeleteTelephonyProvidersEdgeLogicalinterfaceForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type DeleteTelephonyProvidersEdgeLogicalinterfaceForbidden struct {
	Payload *models.ErrorBody
}

func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces/{interfaceId}][%d] deleteTelephonyProvidersEdgeLogicalinterfaceForbidden  %+v", 403, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTelephonyProvidersEdgeLogicalinterfaceNotFound creates a DeleteTelephonyProvidersEdgeLogicalinterfaceNotFound with default headers values
func NewDeleteTelephonyProvidersEdgeLogicalinterfaceNotFound() *DeleteTelephonyProvidersEdgeLogicalinterfaceNotFound {
	return &DeleteTelephonyProvidersEdgeLogicalinterfaceNotFound{}
}

/*DeleteTelephonyProvidersEdgeLogicalinterfaceNotFound handles this case with default header values.

The requested resource was not found.
*/
type DeleteTelephonyProvidersEdgeLogicalinterfaceNotFound struct {
	Payload *models.ErrorBody
}

func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces/{interfaceId}][%d] deleteTelephonyProvidersEdgeLogicalinterfaceNotFound  %+v", 404, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTelephonyProvidersEdgeLogicalinterfaceRequestEntityTooLarge creates a DeleteTelephonyProvidersEdgeLogicalinterfaceRequestEntityTooLarge with default headers values
func NewDeleteTelephonyProvidersEdgeLogicalinterfaceRequestEntityTooLarge() *DeleteTelephonyProvidersEdgeLogicalinterfaceRequestEntityTooLarge {
	return &DeleteTelephonyProvidersEdgeLogicalinterfaceRequestEntityTooLarge{}
}

/*DeleteTelephonyProvidersEdgeLogicalinterfaceRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type DeleteTelephonyProvidersEdgeLogicalinterfaceRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces/{interfaceId}][%d] deleteTelephonyProvidersEdgeLogicalinterfaceRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTelephonyProvidersEdgeLogicalinterfaceUnsupportedMediaType creates a DeleteTelephonyProvidersEdgeLogicalinterfaceUnsupportedMediaType with default headers values
func NewDeleteTelephonyProvidersEdgeLogicalinterfaceUnsupportedMediaType() *DeleteTelephonyProvidersEdgeLogicalinterfaceUnsupportedMediaType {
	return &DeleteTelephonyProvidersEdgeLogicalinterfaceUnsupportedMediaType{}
}

/*DeleteTelephonyProvidersEdgeLogicalinterfaceUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteTelephonyProvidersEdgeLogicalinterfaceUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces/{interfaceId}][%d] deleteTelephonyProvidersEdgeLogicalinterfaceUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTelephonyProvidersEdgeLogicalinterfaceTooManyRequests creates a DeleteTelephonyProvidersEdgeLogicalinterfaceTooManyRequests with default headers values
func NewDeleteTelephonyProvidersEdgeLogicalinterfaceTooManyRequests() *DeleteTelephonyProvidersEdgeLogicalinterfaceTooManyRequests {
	return &DeleteTelephonyProvidersEdgeLogicalinterfaceTooManyRequests{}
}

/*DeleteTelephonyProvidersEdgeLogicalinterfaceTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type DeleteTelephonyProvidersEdgeLogicalinterfaceTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces/{interfaceId}][%d] deleteTelephonyProvidersEdgeLogicalinterfaceTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTelephonyProvidersEdgeLogicalinterfaceInternalServerError creates a DeleteTelephonyProvidersEdgeLogicalinterfaceInternalServerError with default headers values
func NewDeleteTelephonyProvidersEdgeLogicalinterfaceInternalServerError() *DeleteTelephonyProvidersEdgeLogicalinterfaceInternalServerError {
	return &DeleteTelephonyProvidersEdgeLogicalinterfaceInternalServerError{}
}

/*DeleteTelephonyProvidersEdgeLogicalinterfaceInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteTelephonyProvidersEdgeLogicalinterfaceInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces/{interfaceId}][%d] deleteTelephonyProvidersEdgeLogicalinterfaceInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTelephonyProvidersEdgeLogicalinterfaceServiceUnavailable creates a DeleteTelephonyProvidersEdgeLogicalinterfaceServiceUnavailable with default headers values
func NewDeleteTelephonyProvidersEdgeLogicalinterfaceServiceUnavailable() *DeleteTelephonyProvidersEdgeLogicalinterfaceServiceUnavailable {
	return &DeleteTelephonyProvidersEdgeLogicalinterfaceServiceUnavailable{}
}

/*DeleteTelephonyProvidersEdgeLogicalinterfaceServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteTelephonyProvidersEdgeLogicalinterfaceServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces/{interfaceId}][%d] deleteTelephonyProvidersEdgeLogicalinterfaceServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTelephonyProvidersEdgeLogicalinterfaceGatewayTimeout creates a DeleteTelephonyProvidersEdgeLogicalinterfaceGatewayTimeout with default headers values
func NewDeleteTelephonyProvidersEdgeLogicalinterfaceGatewayTimeout() *DeleteTelephonyProvidersEdgeLogicalinterfaceGatewayTimeout {
	return &DeleteTelephonyProvidersEdgeLogicalinterfaceGatewayTimeout{}
}

/*DeleteTelephonyProvidersEdgeLogicalinterfaceGatewayTimeout handles this case with default header values.

The request timed out.
*/
type DeleteTelephonyProvidersEdgeLogicalinterfaceGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces/{interfaceId}][%d] deleteTelephonyProvidersEdgeLogicalinterfaceGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTelephonyProvidersEdgeLogicalinterfaceDefault creates a DeleteTelephonyProvidersEdgeLogicalinterfaceDefault with default headers values
func NewDeleteTelephonyProvidersEdgeLogicalinterfaceDefault(code int) *DeleteTelephonyProvidersEdgeLogicalinterfaceDefault {
	return &DeleteTelephonyProvidersEdgeLogicalinterfaceDefault{
		_statusCode: code,
	}
}

/*DeleteTelephonyProvidersEdgeLogicalinterfaceDefault handles this case with default header values.

successful operation
*/
type DeleteTelephonyProvidersEdgeLogicalinterfaceDefault struct {
	_statusCode int
}

// Code gets the status code for the delete telephony providers edge logicalinterface default response
func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceDefault) Code() int {
	return o._statusCode
}

func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces/{interfaceId}][%d] deleteTelephonyProvidersEdgeLogicalinterface default ", o._statusCode)
}

func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

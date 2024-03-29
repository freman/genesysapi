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
	case 408:
		result := NewDeleteTelephonyProvidersEdgeLogicalinterfaceRequestTimeout()
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

/*
DeleteTelephonyProvidersEdgeLogicalinterfaceBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteTelephonyProvidersEdgeLogicalinterfaceBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete telephony providers edge logicalinterface bad request response has a 2xx status code
func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete telephony providers edge logicalinterface bad request response has a 3xx status code
func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete telephony providers edge logicalinterface bad request response has a 4xx status code
func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete telephony providers edge logicalinterface bad request response has a 5xx status code
func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete telephony providers edge logicalinterface bad request response a status code equal to that given
func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces/{interfaceId}][%d] deleteTelephonyProvidersEdgeLogicalinterfaceBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceBadRequest) String() string {
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

/*
DeleteTelephonyProvidersEdgeLogicalinterfaceUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteTelephonyProvidersEdgeLogicalinterfaceUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete telephony providers edge logicalinterface unauthorized response has a 2xx status code
func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete telephony providers edge logicalinterface unauthorized response has a 3xx status code
func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete telephony providers edge logicalinterface unauthorized response has a 4xx status code
func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete telephony providers edge logicalinterface unauthorized response has a 5xx status code
func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete telephony providers edge logicalinterface unauthorized response a status code equal to that given
func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces/{interfaceId}][%d] deleteTelephonyProvidersEdgeLogicalinterfaceUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceUnauthorized) String() string {
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

/*
DeleteTelephonyProvidersEdgeLogicalinterfaceForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type DeleteTelephonyProvidersEdgeLogicalinterfaceForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete telephony providers edge logicalinterface forbidden response has a 2xx status code
func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete telephony providers edge logicalinterface forbidden response has a 3xx status code
func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete telephony providers edge logicalinterface forbidden response has a 4xx status code
func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete telephony providers edge logicalinterface forbidden response has a 5xx status code
func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete telephony providers edge logicalinterface forbidden response a status code equal to that given
func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces/{interfaceId}][%d] deleteTelephonyProvidersEdgeLogicalinterfaceForbidden  %+v", 403, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceForbidden) String() string {
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

/*
DeleteTelephonyProvidersEdgeLogicalinterfaceNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type DeleteTelephonyProvidersEdgeLogicalinterfaceNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete telephony providers edge logicalinterface not found response has a 2xx status code
func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete telephony providers edge logicalinterface not found response has a 3xx status code
func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete telephony providers edge logicalinterface not found response has a 4xx status code
func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete telephony providers edge logicalinterface not found response has a 5xx status code
func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete telephony providers edge logicalinterface not found response a status code equal to that given
func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces/{interfaceId}][%d] deleteTelephonyProvidersEdgeLogicalinterfaceNotFound  %+v", 404, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceNotFound) String() string {
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

// NewDeleteTelephonyProvidersEdgeLogicalinterfaceRequestTimeout creates a DeleteTelephonyProvidersEdgeLogicalinterfaceRequestTimeout with default headers values
func NewDeleteTelephonyProvidersEdgeLogicalinterfaceRequestTimeout() *DeleteTelephonyProvidersEdgeLogicalinterfaceRequestTimeout {
	return &DeleteTelephonyProvidersEdgeLogicalinterfaceRequestTimeout{}
}

/*
DeleteTelephonyProvidersEdgeLogicalinterfaceRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type DeleteTelephonyProvidersEdgeLogicalinterfaceRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete telephony providers edge logicalinterface request timeout response has a 2xx status code
func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete telephony providers edge logicalinterface request timeout response has a 3xx status code
func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete telephony providers edge logicalinterface request timeout response has a 4xx status code
func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete telephony providers edge logicalinterface request timeout response has a 5xx status code
func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this delete telephony providers edge logicalinterface request timeout response a status code equal to that given
func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces/{interfaceId}][%d] deleteTelephonyProvidersEdgeLogicalinterfaceRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceRequestTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces/{interfaceId}][%d] deleteTelephonyProvidersEdgeLogicalinterfaceRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*
DeleteTelephonyProvidersEdgeLogicalinterfaceRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type DeleteTelephonyProvidersEdgeLogicalinterfaceRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete telephony providers edge logicalinterface request entity too large response has a 2xx status code
func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete telephony providers edge logicalinterface request entity too large response has a 3xx status code
func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete telephony providers edge logicalinterface request entity too large response has a 4xx status code
func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete telephony providers edge logicalinterface request entity too large response has a 5xx status code
func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this delete telephony providers edge logicalinterface request entity too large response a status code equal to that given
func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces/{interfaceId}][%d] deleteTelephonyProvidersEdgeLogicalinterfaceRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceRequestEntityTooLarge) String() string {
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

/*
DeleteTelephonyProvidersEdgeLogicalinterfaceUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteTelephonyProvidersEdgeLogicalinterfaceUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete telephony providers edge logicalinterface unsupported media type response has a 2xx status code
func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete telephony providers edge logicalinterface unsupported media type response has a 3xx status code
func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete telephony providers edge logicalinterface unsupported media type response has a 4xx status code
func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete telephony providers edge logicalinterface unsupported media type response has a 5xx status code
func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this delete telephony providers edge logicalinterface unsupported media type response a status code equal to that given
func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces/{interfaceId}][%d] deleteTelephonyProvidersEdgeLogicalinterfaceUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceUnsupportedMediaType) String() string {
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

/*
DeleteTelephonyProvidersEdgeLogicalinterfaceTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type DeleteTelephonyProvidersEdgeLogicalinterfaceTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete telephony providers edge logicalinterface too many requests response has a 2xx status code
func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete telephony providers edge logicalinterface too many requests response has a 3xx status code
func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete telephony providers edge logicalinterface too many requests response has a 4xx status code
func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete telephony providers edge logicalinterface too many requests response has a 5xx status code
func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this delete telephony providers edge logicalinterface too many requests response a status code equal to that given
func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces/{interfaceId}][%d] deleteTelephonyProvidersEdgeLogicalinterfaceTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceTooManyRequests) String() string {
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

/*
DeleteTelephonyProvidersEdgeLogicalinterfaceInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteTelephonyProvidersEdgeLogicalinterfaceInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete telephony providers edge logicalinterface internal server error response has a 2xx status code
func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete telephony providers edge logicalinterface internal server error response has a 3xx status code
func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete telephony providers edge logicalinterface internal server error response has a 4xx status code
func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete telephony providers edge logicalinterface internal server error response has a 5xx status code
func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete telephony providers edge logicalinterface internal server error response a status code equal to that given
func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces/{interfaceId}][%d] deleteTelephonyProvidersEdgeLogicalinterfaceInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceInternalServerError) String() string {
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

/*
DeleteTelephonyProvidersEdgeLogicalinterfaceServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteTelephonyProvidersEdgeLogicalinterfaceServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete telephony providers edge logicalinterface service unavailable response has a 2xx status code
func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete telephony providers edge logicalinterface service unavailable response has a 3xx status code
func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete telephony providers edge logicalinterface service unavailable response has a 4xx status code
func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete telephony providers edge logicalinterface service unavailable response has a 5xx status code
func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this delete telephony providers edge logicalinterface service unavailable response a status code equal to that given
func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces/{interfaceId}][%d] deleteTelephonyProvidersEdgeLogicalinterfaceServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceServiceUnavailable) String() string {
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

/*
DeleteTelephonyProvidersEdgeLogicalinterfaceGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type DeleteTelephonyProvidersEdgeLogicalinterfaceGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete telephony providers edge logicalinterface gateway timeout response has a 2xx status code
func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete telephony providers edge logicalinterface gateway timeout response has a 3xx status code
func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete telephony providers edge logicalinterface gateway timeout response has a 4xx status code
func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete telephony providers edge logicalinterface gateway timeout response has a 5xx status code
func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this delete telephony providers edge logicalinterface gateway timeout response a status code equal to that given
func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces/{interfaceId}][%d] deleteTelephonyProvidersEdgeLogicalinterfaceGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceGatewayTimeout) String() string {
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

/*
DeleteTelephonyProvidersEdgeLogicalinterfaceDefault describes a response with status code -1, with default header values.

successful operation
*/
type DeleteTelephonyProvidersEdgeLogicalinterfaceDefault struct {
	_statusCode int
}

// Code gets the status code for the delete telephony providers edge logicalinterface default response
func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this delete telephony providers edge logicalinterface default response has a 2xx status code
func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete telephony providers edge logicalinterface default response has a 3xx status code
func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete telephony providers edge logicalinterface default response has a 4xx status code
func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete telephony providers edge logicalinterface default response has a 5xx status code
func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete telephony providers edge logicalinterface default response a status code equal to that given
func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces/{interfaceId}][%d] deleteTelephonyProvidersEdgeLogicalinterface default ", o._statusCode)
}

func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceDefault) String() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces/{interfaceId}][%d] deleteTelephonyProvidersEdgeLogicalinterface default ", o._statusCode)
}

func (o *DeleteTelephonyProvidersEdgeLogicalinterfaceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

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
	case 408:
		result := NewGetTelephonyProvidersEdgePhysicalinterfaceRequestTimeout()
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

/*
GetTelephonyProvidersEdgePhysicalinterfaceOK describes a response with status code 200, with default header values.

successful operation
*/
type GetTelephonyProvidersEdgePhysicalinterfaceOK struct {
	Payload *models.DomainPhysicalInterface
}

// IsSuccess returns true when this get telephony providers edge physicalinterface o k response has a 2xx status code
func (o *GetTelephonyProvidersEdgePhysicalinterfaceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get telephony providers edge physicalinterface o k response has a 3xx status code
func (o *GetTelephonyProvidersEdgePhysicalinterfaceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edge physicalinterface o k response has a 4xx status code
func (o *GetTelephonyProvidersEdgePhysicalinterfaceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get telephony providers edge physicalinterface o k response has a 5xx status code
func (o *GetTelephonyProvidersEdgePhysicalinterfaceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edge physicalinterface o k response a status code equal to that given
func (o *GetTelephonyProvidersEdgePhysicalinterfaceOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetTelephonyProvidersEdgePhysicalinterfaceOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/physicalinterfaces/{interfaceId}][%d] getTelephonyProvidersEdgePhysicalinterfaceOK  %+v", 200, o.Payload)
}

func (o *GetTelephonyProvidersEdgePhysicalinterfaceOK) String() string {
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

/*
GetTelephonyProvidersEdgePhysicalinterfaceBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetTelephonyProvidersEdgePhysicalinterfaceBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edge physicalinterface bad request response has a 2xx status code
func (o *GetTelephonyProvidersEdgePhysicalinterfaceBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edge physicalinterface bad request response has a 3xx status code
func (o *GetTelephonyProvidersEdgePhysicalinterfaceBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edge physicalinterface bad request response has a 4xx status code
func (o *GetTelephonyProvidersEdgePhysicalinterfaceBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony providers edge physicalinterface bad request response has a 5xx status code
func (o *GetTelephonyProvidersEdgePhysicalinterfaceBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edge physicalinterface bad request response a status code equal to that given
func (o *GetTelephonyProvidersEdgePhysicalinterfaceBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetTelephonyProvidersEdgePhysicalinterfaceBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/physicalinterfaces/{interfaceId}][%d] getTelephonyProvidersEdgePhysicalinterfaceBadRequest  %+v", 400, o.Payload)
}

func (o *GetTelephonyProvidersEdgePhysicalinterfaceBadRequest) String() string {
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

/*
GetTelephonyProvidersEdgePhysicalinterfaceUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetTelephonyProvidersEdgePhysicalinterfaceUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edge physicalinterface unauthorized response has a 2xx status code
func (o *GetTelephonyProvidersEdgePhysicalinterfaceUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edge physicalinterface unauthorized response has a 3xx status code
func (o *GetTelephonyProvidersEdgePhysicalinterfaceUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edge physicalinterface unauthorized response has a 4xx status code
func (o *GetTelephonyProvidersEdgePhysicalinterfaceUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony providers edge physicalinterface unauthorized response has a 5xx status code
func (o *GetTelephonyProvidersEdgePhysicalinterfaceUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edge physicalinterface unauthorized response a status code equal to that given
func (o *GetTelephonyProvidersEdgePhysicalinterfaceUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetTelephonyProvidersEdgePhysicalinterfaceUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/physicalinterfaces/{interfaceId}][%d] getTelephonyProvidersEdgePhysicalinterfaceUnauthorized  %+v", 401, o.Payload)
}

func (o *GetTelephonyProvidersEdgePhysicalinterfaceUnauthorized) String() string {
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

/*
GetTelephonyProvidersEdgePhysicalinterfaceForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetTelephonyProvidersEdgePhysicalinterfaceForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edge physicalinterface forbidden response has a 2xx status code
func (o *GetTelephonyProvidersEdgePhysicalinterfaceForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edge physicalinterface forbidden response has a 3xx status code
func (o *GetTelephonyProvidersEdgePhysicalinterfaceForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edge physicalinterface forbidden response has a 4xx status code
func (o *GetTelephonyProvidersEdgePhysicalinterfaceForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony providers edge physicalinterface forbidden response has a 5xx status code
func (o *GetTelephonyProvidersEdgePhysicalinterfaceForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edge physicalinterface forbidden response a status code equal to that given
func (o *GetTelephonyProvidersEdgePhysicalinterfaceForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetTelephonyProvidersEdgePhysicalinterfaceForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/physicalinterfaces/{interfaceId}][%d] getTelephonyProvidersEdgePhysicalinterfaceForbidden  %+v", 403, o.Payload)
}

func (o *GetTelephonyProvidersEdgePhysicalinterfaceForbidden) String() string {
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

/*
GetTelephonyProvidersEdgePhysicalinterfaceNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetTelephonyProvidersEdgePhysicalinterfaceNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edge physicalinterface not found response has a 2xx status code
func (o *GetTelephonyProvidersEdgePhysicalinterfaceNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edge physicalinterface not found response has a 3xx status code
func (o *GetTelephonyProvidersEdgePhysicalinterfaceNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edge physicalinterface not found response has a 4xx status code
func (o *GetTelephonyProvidersEdgePhysicalinterfaceNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony providers edge physicalinterface not found response has a 5xx status code
func (o *GetTelephonyProvidersEdgePhysicalinterfaceNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edge physicalinterface not found response a status code equal to that given
func (o *GetTelephonyProvidersEdgePhysicalinterfaceNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetTelephonyProvidersEdgePhysicalinterfaceNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/physicalinterfaces/{interfaceId}][%d] getTelephonyProvidersEdgePhysicalinterfaceNotFound  %+v", 404, o.Payload)
}

func (o *GetTelephonyProvidersEdgePhysicalinterfaceNotFound) String() string {
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

// NewGetTelephonyProvidersEdgePhysicalinterfaceRequestTimeout creates a GetTelephonyProvidersEdgePhysicalinterfaceRequestTimeout with default headers values
func NewGetTelephonyProvidersEdgePhysicalinterfaceRequestTimeout() *GetTelephonyProvidersEdgePhysicalinterfaceRequestTimeout {
	return &GetTelephonyProvidersEdgePhysicalinterfaceRequestTimeout{}
}

/*
GetTelephonyProvidersEdgePhysicalinterfaceRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetTelephonyProvidersEdgePhysicalinterfaceRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edge physicalinterface request timeout response has a 2xx status code
func (o *GetTelephonyProvidersEdgePhysicalinterfaceRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edge physicalinterface request timeout response has a 3xx status code
func (o *GetTelephonyProvidersEdgePhysicalinterfaceRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edge physicalinterface request timeout response has a 4xx status code
func (o *GetTelephonyProvidersEdgePhysicalinterfaceRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony providers edge physicalinterface request timeout response has a 5xx status code
func (o *GetTelephonyProvidersEdgePhysicalinterfaceRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edge physicalinterface request timeout response a status code equal to that given
func (o *GetTelephonyProvidersEdgePhysicalinterfaceRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetTelephonyProvidersEdgePhysicalinterfaceRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/physicalinterfaces/{interfaceId}][%d] getTelephonyProvidersEdgePhysicalinterfaceRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetTelephonyProvidersEdgePhysicalinterfaceRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/physicalinterfaces/{interfaceId}][%d] getTelephonyProvidersEdgePhysicalinterfaceRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetTelephonyProvidersEdgePhysicalinterfaceRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgePhysicalinterfaceRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*
GetTelephonyProvidersEdgePhysicalinterfaceRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetTelephonyProvidersEdgePhysicalinterfaceRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edge physicalinterface request entity too large response has a 2xx status code
func (o *GetTelephonyProvidersEdgePhysicalinterfaceRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edge physicalinterface request entity too large response has a 3xx status code
func (o *GetTelephonyProvidersEdgePhysicalinterfaceRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edge physicalinterface request entity too large response has a 4xx status code
func (o *GetTelephonyProvidersEdgePhysicalinterfaceRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony providers edge physicalinterface request entity too large response has a 5xx status code
func (o *GetTelephonyProvidersEdgePhysicalinterfaceRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edge physicalinterface request entity too large response a status code equal to that given
func (o *GetTelephonyProvidersEdgePhysicalinterfaceRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetTelephonyProvidersEdgePhysicalinterfaceRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/physicalinterfaces/{interfaceId}][%d] getTelephonyProvidersEdgePhysicalinterfaceRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetTelephonyProvidersEdgePhysicalinterfaceRequestEntityTooLarge) String() string {
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

/*
GetTelephonyProvidersEdgePhysicalinterfaceUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetTelephonyProvidersEdgePhysicalinterfaceUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edge physicalinterface unsupported media type response has a 2xx status code
func (o *GetTelephonyProvidersEdgePhysicalinterfaceUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edge physicalinterface unsupported media type response has a 3xx status code
func (o *GetTelephonyProvidersEdgePhysicalinterfaceUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edge physicalinterface unsupported media type response has a 4xx status code
func (o *GetTelephonyProvidersEdgePhysicalinterfaceUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony providers edge physicalinterface unsupported media type response has a 5xx status code
func (o *GetTelephonyProvidersEdgePhysicalinterfaceUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edge physicalinterface unsupported media type response a status code equal to that given
func (o *GetTelephonyProvidersEdgePhysicalinterfaceUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetTelephonyProvidersEdgePhysicalinterfaceUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/physicalinterfaces/{interfaceId}][%d] getTelephonyProvidersEdgePhysicalinterfaceUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetTelephonyProvidersEdgePhysicalinterfaceUnsupportedMediaType) String() string {
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

/*
GetTelephonyProvidersEdgePhysicalinterfaceTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetTelephonyProvidersEdgePhysicalinterfaceTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edge physicalinterface too many requests response has a 2xx status code
func (o *GetTelephonyProvidersEdgePhysicalinterfaceTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edge physicalinterface too many requests response has a 3xx status code
func (o *GetTelephonyProvidersEdgePhysicalinterfaceTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edge physicalinterface too many requests response has a 4xx status code
func (o *GetTelephonyProvidersEdgePhysicalinterfaceTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony providers edge physicalinterface too many requests response has a 5xx status code
func (o *GetTelephonyProvidersEdgePhysicalinterfaceTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edge physicalinterface too many requests response a status code equal to that given
func (o *GetTelephonyProvidersEdgePhysicalinterfaceTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetTelephonyProvidersEdgePhysicalinterfaceTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/physicalinterfaces/{interfaceId}][%d] getTelephonyProvidersEdgePhysicalinterfaceTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetTelephonyProvidersEdgePhysicalinterfaceTooManyRequests) String() string {
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

/*
GetTelephonyProvidersEdgePhysicalinterfaceInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetTelephonyProvidersEdgePhysicalinterfaceInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edge physicalinterface internal server error response has a 2xx status code
func (o *GetTelephonyProvidersEdgePhysicalinterfaceInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edge physicalinterface internal server error response has a 3xx status code
func (o *GetTelephonyProvidersEdgePhysicalinterfaceInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edge physicalinterface internal server error response has a 4xx status code
func (o *GetTelephonyProvidersEdgePhysicalinterfaceInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get telephony providers edge physicalinterface internal server error response has a 5xx status code
func (o *GetTelephonyProvidersEdgePhysicalinterfaceInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get telephony providers edge physicalinterface internal server error response a status code equal to that given
func (o *GetTelephonyProvidersEdgePhysicalinterfaceInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetTelephonyProvidersEdgePhysicalinterfaceInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/physicalinterfaces/{interfaceId}][%d] getTelephonyProvidersEdgePhysicalinterfaceInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTelephonyProvidersEdgePhysicalinterfaceInternalServerError) String() string {
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

/*
GetTelephonyProvidersEdgePhysicalinterfaceServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetTelephonyProvidersEdgePhysicalinterfaceServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edge physicalinterface service unavailable response has a 2xx status code
func (o *GetTelephonyProvidersEdgePhysicalinterfaceServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edge physicalinterface service unavailable response has a 3xx status code
func (o *GetTelephonyProvidersEdgePhysicalinterfaceServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edge physicalinterface service unavailable response has a 4xx status code
func (o *GetTelephonyProvidersEdgePhysicalinterfaceServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get telephony providers edge physicalinterface service unavailable response has a 5xx status code
func (o *GetTelephonyProvidersEdgePhysicalinterfaceServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get telephony providers edge physicalinterface service unavailable response a status code equal to that given
func (o *GetTelephonyProvidersEdgePhysicalinterfaceServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetTelephonyProvidersEdgePhysicalinterfaceServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/physicalinterfaces/{interfaceId}][%d] getTelephonyProvidersEdgePhysicalinterfaceServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetTelephonyProvidersEdgePhysicalinterfaceServiceUnavailable) String() string {
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

/*
GetTelephonyProvidersEdgePhysicalinterfaceGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetTelephonyProvidersEdgePhysicalinterfaceGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edge physicalinterface gateway timeout response has a 2xx status code
func (o *GetTelephonyProvidersEdgePhysicalinterfaceGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edge physicalinterface gateway timeout response has a 3xx status code
func (o *GetTelephonyProvidersEdgePhysicalinterfaceGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edge physicalinterface gateway timeout response has a 4xx status code
func (o *GetTelephonyProvidersEdgePhysicalinterfaceGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get telephony providers edge physicalinterface gateway timeout response has a 5xx status code
func (o *GetTelephonyProvidersEdgePhysicalinterfaceGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get telephony providers edge physicalinterface gateway timeout response a status code equal to that given
func (o *GetTelephonyProvidersEdgePhysicalinterfaceGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetTelephonyProvidersEdgePhysicalinterfaceGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/physicalinterfaces/{interfaceId}][%d] getTelephonyProvidersEdgePhysicalinterfaceGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetTelephonyProvidersEdgePhysicalinterfaceGatewayTimeout) String() string {
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

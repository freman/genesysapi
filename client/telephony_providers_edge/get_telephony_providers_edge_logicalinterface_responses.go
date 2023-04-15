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

// GetTelephonyProvidersEdgeLogicalinterfaceReader is a Reader for the GetTelephonyProvidersEdgeLogicalinterface structure.
type GetTelephonyProvidersEdgeLogicalinterfaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTelephonyProvidersEdgeLogicalinterfaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTelephonyProvidersEdgeLogicalinterfaceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetTelephonyProvidersEdgeLogicalinterfaceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetTelephonyProvidersEdgeLogicalinterfaceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetTelephonyProvidersEdgeLogicalinterfaceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTelephonyProvidersEdgeLogicalinterfaceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetTelephonyProvidersEdgeLogicalinterfaceRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetTelephonyProvidersEdgeLogicalinterfaceRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetTelephonyProvidersEdgeLogicalinterfaceUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetTelephonyProvidersEdgeLogicalinterfaceTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetTelephonyProvidersEdgeLogicalinterfaceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetTelephonyProvidersEdgeLogicalinterfaceServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetTelephonyProvidersEdgeLogicalinterfaceGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTelephonyProvidersEdgeLogicalinterfaceOK creates a GetTelephonyProvidersEdgeLogicalinterfaceOK with default headers values
func NewGetTelephonyProvidersEdgeLogicalinterfaceOK() *GetTelephonyProvidersEdgeLogicalinterfaceOK {
	return &GetTelephonyProvidersEdgeLogicalinterfaceOK{}
}

/*
GetTelephonyProvidersEdgeLogicalinterfaceOK describes a response with status code 200, with default header values.

successful operation
*/
type GetTelephonyProvidersEdgeLogicalinterfaceOK struct {
	Payload *models.DomainLogicalInterface
}

// IsSuccess returns true when this get telephony providers edge logicalinterface o k response has a 2xx status code
func (o *GetTelephonyProvidersEdgeLogicalinterfaceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get telephony providers edge logicalinterface o k response has a 3xx status code
func (o *GetTelephonyProvidersEdgeLogicalinterfaceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edge logicalinterface o k response has a 4xx status code
func (o *GetTelephonyProvidersEdgeLogicalinterfaceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get telephony providers edge logicalinterface o k response has a 5xx status code
func (o *GetTelephonyProvidersEdgeLogicalinterfaceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edge logicalinterface o k response a status code equal to that given
func (o *GetTelephonyProvidersEdgeLogicalinterfaceOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetTelephonyProvidersEdgeLogicalinterfaceOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces/{interfaceId}][%d] getTelephonyProvidersEdgeLogicalinterfaceOK  %+v", 200, o.Payload)
}

func (o *GetTelephonyProvidersEdgeLogicalinterfaceOK) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces/{interfaceId}][%d] getTelephonyProvidersEdgeLogicalinterfaceOK  %+v", 200, o.Payload)
}

func (o *GetTelephonyProvidersEdgeLogicalinterfaceOK) GetPayload() *models.DomainLogicalInterface {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgeLogicalinterfaceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DomainLogicalInterface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgeLogicalinterfaceBadRequest creates a GetTelephonyProvidersEdgeLogicalinterfaceBadRequest with default headers values
func NewGetTelephonyProvidersEdgeLogicalinterfaceBadRequest() *GetTelephonyProvidersEdgeLogicalinterfaceBadRequest {
	return &GetTelephonyProvidersEdgeLogicalinterfaceBadRequest{}
}

/*
GetTelephonyProvidersEdgeLogicalinterfaceBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetTelephonyProvidersEdgeLogicalinterfaceBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edge logicalinterface bad request response has a 2xx status code
func (o *GetTelephonyProvidersEdgeLogicalinterfaceBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edge logicalinterface bad request response has a 3xx status code
func (o *GetTelephonyProvidersEdgeLogicalinterfaceBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edge logicalinterface bad request response has a 4xx status code
func (o *GetTelephonyProvidersEdgeLogicalinterfaceBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony providers edge logicalinterface bad request response has a 5xx status code
func (o *GetTelephonyProvidersEdgeLogicalinterfaceBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edge logicalinterface bad request response a status code equal to that given
func (o *GetTelephonyProvidersEdgeLogicalinterfaceBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetTelephonyProvidersEdgeLogicalinterfaceBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces/{interfaceId}][%d] getTelephonyProvidersEdgeLogicalinterfaceBadRequest  %+v", 400, o.Payload)
}

func (o *GetTelephonyProvidersEdgeLogicalinterfaceBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces/{interfaceId}][%d] getTelephonyProvidersEdgeLogicalinterfaceBadRequest  %+v", 400, o.Payload)
}

func (o *GetTelephonyProvidersEdgeLogicalinterfaceBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgeLogicalinterfaceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgeLogicalinterfaceUnauthorized creates a GetTelephonyProvidersEdgeLogicalinterfaceUnauthorized with default headers values
func NewGetTelephonyProvidersEdgeLogicalinterfaceUnauthorized() *GetTelephonyProvidersEdgeLogicalinterfaceUnauthorized {
	return &GetTelephonyProvidersEdgeLogicalinterfaceUnauthorized{}
}

/*
GetTelephonyProvidersEdgeLogicalinterfaceUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetTelephonyProvidersEdgeLogicalinterfaceUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edge logicalinterface unauthorized response has a 2xx status code
func (o *GetTelephonyProvidersEdgeLogicalinterfaceUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edge logicalinterface unauthorized response has a 3xx status code
func (o *GetTelephonyProvidersEdgeLogicalinterfaceUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edge logicalinterface unauthorized response has a 4xx status code
func (o *GetTelephonyProvidersEdgeLogicalinterfaceUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony providers edge logicalinterface unauthorized response has a 5xx status code
func (o *GetTelephonyProvidersEdgeLogicalinterfaceUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edge logicalinterface unauthorized response a status code equal to that given
func (o *GetTelephonyProvidersEdgeLogicalinterfaceUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetTelephonyProvidersEdgeLogicalinterfaceUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces/{interfaceId}][%d] getTelephonyProvidersEdgeLogicalinterfaceUnauthorized  %+v", 401, o.Payload)
}

func (o *GetTelephonyProvidersEdgeLogicalinterfaceUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces/{interfaceId}][%d] getTelephonyProvidersEdgeLogicalinterfaceUnauthorized  %+v", 401, o.Payload)
}

func (o *GetTelephonyProvidersEdgeLogicalinterfaceUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgeLogicalinterfaceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgeLogicalinterfaceForbidden creates a GetTelephonyProvidersEdgeLogicalinterfaceForbidden with default headers values
func NewGetTelephonyProvidersEdgeLogicalinterfaceForbidden() *GetTelephonyProvidersEdgeLogicalinterfaceForbidden {
	return &GetTelephonyProvidersEdgeLogicalinterfaceForbidden{}
}

/*
GetTelephonyProvidersEdgeLogicalinterfaceForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetTelephonyProvidersEdgeLogicalinterfaceForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edge logicalinterface forbidden response has a 2xx status code
func (o *GetTelephonyProvidersEdgeLogicalinterfaceForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edge logicalinterface forbidden response has a 3xx status code
func (o *GetTelephonyProvidersEdgeLogicalinterfaceForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edge logicalinterface forbidden response has a 4xx status code
func (o *GetTelephonyProvidersEdgeLogicalinterfaceForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony providers edge logicalinterface forbidden response has a 5xx status code
func (o *GetTelephonyProvidersEdgeLogicalinterfaceForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edge logicalinterface forbidden response a status code equal to that given
func (o *GetTelephonyProvidersEdgeLogicalinterfaceForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetTelephonyProvidersEdgeLogicalinterfaceForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces/{interfaceId}][%d] getTelephonyProvidersEdgeLogicalinterfaceForbidden  %+v", 403, o.Payload)
}

func (o *GetTelephonyProvidersEdgeLogicalinterfaceForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces/{interfaceId}][%d] getTelephonyProvidersEdgeLogicalinterfaceForbidden  %+v", 403, o.Payload)
}

func (o *GetTelephonyProvidersEdgeLogicalinterfaceForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgeLogicalinterfaceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgeLogicalinterfaceNotFound creates a GetTelephonyProvidersEdgeLogicalinterfaceNotFound with default headers values
func NewGetTelephonyProvidersEdgeLogicalinterfaceNotFound() *GetTelephonyProvidersEdgeLogicalinterfaceNotFound {
	return &GetTelephonyProvidersEdgeLogicalinterfaceNotFound{}
}

/*
GetTelephonyProvidersEdgeLogicalinterfaceNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetTelephonyProvidersEdgeLogicalinterfaceNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edge logicalinterface not found response has a 2xx status code
func (o *GetTelephonyProvidersEdgeLogicalinterfaceNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edge logicalinterface not found response has a 3xx status code
func (o *GetTelephonyProvidersEdgeLogicalinterfaceNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edge logicalinterface not found response has a 4xx status code
func (o *GetTelephonyProvidersEdgeLogicalinterfaceNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony providers edge logicalinterface not found response has a 5xx status code
func (o *GetTelephonyProvidersEdgeLogicalinterfaceNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edge logicalinterface not found response a status code equal to that given
func (o *GetTelephonyProvidersEdgeLogicalinterfaceNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetTelephonyProvidersEdgeLogicalinterfaceNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces/{interfaceId}][%d] getTelephonyProvidersEdgeLogicalinterfaceNotFound  %+v", 404, o.Payload)
}

func (o *GetTelephonyProvidersEdgeLogicalinterfaceNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces/{interfaceId}][%d] getTelephonyProvidersEdgeLogicalinterfaceNotFound  %+v", 404, o.Payload)
}

func (o *GetTelephonyProvidersEdgeLogicalinterfaceNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgeLogicalinterfaceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgeLogicalinterfaceRequestTimeout creates a GetTelephonyProvidersEdgeLogicalinterfaceRequestTimeout with default headers values
func NewGetTelephonyProvidersEdgeLogicalinterfaceRequestTimeout() *GetTelephonyProvidersEdgeLogicalinterfaceRequestTimeout {
	return &GetTelephonyProvidersEdgeLogicalinterfaceRequestTimeout{}
}

/*
GetTelephonyProvidersEdgeLogicalinterfaceRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetTelephonyProvidersEdgeLogicalinterfaceRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edge logicalinterface request timeout response has a 2xx status code
func (o *GetTelephonyProvidersEdgeLogicalinterfaceRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edge logicalinterface request timeout response has a 3xx status code
func (o *GetTelephonyProvidersEdgeLogicalinterfaceRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edge logicalinterface request timeout response has a 4xx status code
func (o *GetTelephonyProvidersEdgeLogicalinterfaceRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony providers edge logicalinterface request timeout response has a 5xx status code
func (o *GetTelephonyProvidersEdgeLogicalinterfaceRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edge logicalinterface request timeout response a status code equal to that given
func (o *GetTelephonyProvidersEdgeLogicalinterfaceRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetTelephonyProvidersEdgeLogicalinterfaceRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces/{interfaceId}][%d] getTelephonyProvidersEdgeLogicalinterfaceRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetTelephonyProvidersEdgeLogicalinterfaceRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces/{interfaceId}][%d] getTelephonyProvidersEdgeLogicalinterfaceRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetTelephonyProvidersEdgeLogicalinterfaceRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgeLogicalinterfaceRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgeLogicalinterfaceRequestEntityTooLarge creates a GetTelephonyProvidersEdgeLogicalinterfaceRequestEntityTooLarge with default headers values
func NewGetTelephonyProvidersEdgeLogicalinterfaceRequestEntityTooLarge() *GetTelephonyProvidersEdgeLogicalinterfaceRequestEntityTooLarge {
	return &GetTelephonyProvidersEdgeLogicalinterfaceRequestEntityTooLarge{}
}

/*
GetTelephonyProvidersEdgeLogicalinterfaceRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetTelephonyProvidersEdgeLogicalinterfaceRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edge logicalinterface request entity too large response has a 2xx status code
func (o *GetTelephonyProvidersEdgeLogicalinterfaceRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edge logicalinterface request entity too large response has a 3xx status code
func (o *GetTelephonyProvidersEdgeLogicalinterfaceRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edge logicalinterface request entity too large response has a 4xx status code
func (o *GetTelephonyProvidersEdgeLogicalinterfaceRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony providers edge logicalinterface request entity too large response has a 5xx status code
func (o *GetTelephonyProvidersEdgeLogicalinterfaceRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edge logicalinterface request entity too large response a status code equal to that given
func (o *GetTelephonyProvidersEdgeLogicalinterfaceRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetTelephonyProvidersEdgeLogicalinterfaceRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces/{interfaceId}][%d] getTelephonyProvidersEdgeLogicalinterfaceRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetTelephonyProvidersEdgeLogicalinterfaceRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces/{interfaceId}][%d] getTelephonyProvidersEdgeLogicalinterfaceRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetTelephonyProvidersEdgeLogicalinterfaceRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgeLogicalinterfaceRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgeLogicalinterfaceUnsupportedMediaType creates a GetTelephonyProvidersEdgeLogicalinterfaceUnsupportedMediaType with default headers values
func NewGetTelephonyProvidersEdgeLogicalinterfaceUnsupportedMediaType() *GetTelephonyProvidersEdgeLogicalinterfaceUnsupportedMediaType {
	return &GetTelephonyProvidersEdgeLogicalinterfaceUnsupportedMediaType{}
}

/*
GetTelephonyProvidersEdgeLogicalinterfaceUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetTelephonyProvidersEdgeLogicalinterfaceUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edge logicalinterface unsupported media type response has a 2xx status code
func (o *GetTelephonyProvidersEdgeLogicalinterfaceUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edge logicalinterface unsupported media type response has a 3xx status code
func (o *GetTelephonyProvidersEdgeLogicalinterfaceUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edge logicalinterface unsupported media type response has a 4xx status code
func (o *GetTelephonyProvidersEdgeLogicalinterfaceUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony providers edge logicalinterface unsupported media type response has a 5xx status code
func (o *GetTelephonyProvidersEdgeLogicalinterfaceUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edge logicalinterface unsupported media type response a status code equal to that given
func (o *GetTelephonyProvidersEdgeLogicalinterfaceUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetTelephonyProvidersEdgeLogicalinterfaceUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces/{interfaceId}][%d] getTelephonyProvidersEdgeLogicalinterfaceUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetTelephonyProvidersEdgeLogicalinterfaceUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces/{interfaceId}][%d] getTelephonyProvidersEdgeLogicalinterfaceUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetTelephonyProvidersEdgeLogicalinterfaceUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgeLogicalinterfaceUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgeLogicalinterfaceTooManyRequests creates a GetTelephonyProvidersEdgeLogicalinterfaceTooManyRequests with default headers values
func NewGetTelephonyProvidersEdgeLogicalinterfaceTooManyRequests() *GetTelephonyProvidersEdgeLogicalinterfaceTooManyRequests {
	return &GetTelephonyProvidersEdgeLogicalinterfaceTooManyRequests{}
}

/*
GetTelephonyProvidersEdgeLogicalinterfaceTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetTelephonyProvidersEdgeLogicalinterfaceTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edge logicalinterface too many requests response has a 2xx status code
func (o *GetTelephonyProvidersEdgeLogicalinterfaceTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edge logicalinterface too many requests response has a 3xx status code
func (o *GetTelephonyProvidersEdgeLogicalinterfaceTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edge logicalinterface too many requests response has a 4xx status code
func (o *GetTelephonyProvidersEdgeLogicalinterfaceTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony providers edge logicalinterface too many requests response has a 5xx status code
func (o *GetTelephonyProvidersEdgeLogicalinterfaceTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edge logicalinterface too many requests response a status code equal to that given
func (o *GetTelephonyProvidersEdgeLogicalinterfaceTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetTelephonyProvidersEdgeLogicalinterfaceTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces/{interfaceId}][%d] getTelephonyProvidersEdgeLogicalinterfaceTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetTelephonyProvidersEdgeLogicalinterfaceTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces/{interfaceId}][%d] getTelephonyProvidersEdgeLogicalinterfaceTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetTelephonyProvidersEdgeLogicalinterfaceTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgeLogicalinterfaceTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgeLogicalinterfaceInternalServerError creates a GetTelephonyProvidersEdgeLogicalinterfaceInternalServerError with default headers values
func NewGetTelephonyProvidersEdgeLogicalinterfaceInternalServerError() *GetTelephonyProvidersEdgeLogicalinterfaceInternalServerError {
	return &GetTelephonyProvidersEdgeLogicalinterfaceInternalServerError{}
}

/*
GetTelephonyProvidersEdgeLogicalinterfaceInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetTelephonyProvidersEdgeLogicalinterfaceInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edge logicalinterface internal server error response has a 2xx status code
func (o *GetTelephonyProvidersEdgeLogicalinterfaceInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edge logicalinterface internal server error response has a 3xx status code
func (o *GetTelephonyProvidersEdgeLogicalinterfaceInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edge logicalinterface internal server error response has a 4xx status code
func (o *GetTelephonyProvidersEdgeLogicalinterfaceInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get telephony providers edge logicalinterface internal server error response has a 5xx status code
func (o *GetTelephonyProvidersEdgeLogicalinterfaceInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get telephony providers edge logicalinterface internal server error response a status code equal to that given
func (o *GetTelephonyProvidersEdgeLogicalinterfaceInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetTelephonyProvidersEdgeLogicalinterfaceInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces/{interfaceId}][%d] getTelephonyProvidersEdgeLogicalinterfaceInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTelephonyProvidersEdgeLogicalinterfaceInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces/{interfaceId}][%d] getTelephonyProvidersEdgeLogicalinterfaceInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTelephonyProvidersEdgeLogicalinterfaceInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgeLogicalinterfaceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgeLogicalinterfaceServiceUnavailable creates a GetTelephonyProvidersEdgeLogicalinterfaceServiceUnavailable with default headers values
func NewGetTelephonyProvidersEdgeLogicalinterfaceServiceUnavailable() *GetTelephonyProvidersEdgeLogicalinterfaceServiceUnavailable {
	return &GetTelephonyProvidersEdgeLogicalinterfaceServiceUnavailable{}
}

/*
GetTelephonyProvidersEdgeLogicalinterfaceServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetTelephonyProvidersEdgeLogicalinterfaceServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edge logicalinterface service unavailable response has a 2xx status code
func (o *GetTelephonyProvidersEdgeLogicalinterfaceServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edge logicalinterface service unavailable response has a 3xx status code
func (o *GetTelephonyProvidersEdgeLogicalinterfaceServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edge logicalinterface service unavailable response has a 4xx status code
func (o *GetTelephonyProvidersEdgeLogicalinterfaceServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get telephony providers edge logicalinterface service unavailable response has a 5xx status code
func (o *GetTelephonyProvidersEdgeLogicalinterfaceServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get telephony providers edge logicalinterface service unavailable response a status code equal to that given
func (o *GetTelephonyProvidersEdgeLogicalinterfaceServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetTelephonyProvidersEdgeLogicalinterfaceServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces/{interfaceId}][%d] getTelephonyProvidersEdgeLogicalinterfaceServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetTelephonyProvidersEdgeLogicalinterfaceServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces/{interfaceId}][%d] getTelephonyProvidersEdgeLogicalinterfaceServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetTelephonyProvidersEdgeLogicalinterfaceServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgeLogicalinterfaceServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgeLogicalinterfaceGatewayTimeout creates a GetTelephonyProvidersEdgeLogicalinterfaceGatewayTimeout with default headers values
func NewGetTelephonyProvidersEdgeLogicalinterfaceGatewayTimeout() *GetTelephonyProvidersEdgeLogicalinterfaceGatewayTimeout {
	return &GetTelephonyProvidersEdgeLogicalinterfaceGatewayTimeout{}
}

/*
GetTelephonyProvidersEdgeLogicalinterfaceGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetTelephonyProvidersEdgeLogicalinterfaceGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edge logicalinterface gateway timeout response has a 2xx status code
func (o *GetTelephonyProvidersEdgeLogicalinterfaceGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edge logicalinterface gateway timeout response has a 3xx status code
func (o *GetTelephonyProvidersEdgeLogicalinterfaceGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edge logicalinterface gateway timeout response has a 4xx status code
func (o *GetTelephonyProvidersEdgeLogicalinterfaceGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get telephony providers edge logicalinterface gateway timeout response has a 5xx status code
func (o *GetTelephonyProvidersEdgeLogicalinterfaceGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get telephony providers edge logicalinterface gateway timeout response a status code equal to that given
func (o *GetTelephonyProvidersEdgeLogicalinterfaceGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetTelephonyProvidersEdgeLogicalinterfaceGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces/{interfaceId}][%d] getTelephonyProvidersEdgeLogicalinterfaceGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetTelephonyProvidersEdgeLogicalinterfaceGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces/{interfaceId}][%d] getTelephonyProvidersEdgeLogicalinterfaceGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetTelephonyProvidersEdgeLogicalinterfaceGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgeLogicalinterfaceGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

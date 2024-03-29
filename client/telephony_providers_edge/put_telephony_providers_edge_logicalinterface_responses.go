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

// PutTelephonyProvidersEdgeLogicalinterfaceReader is a Reader for the PutTelephonyProvidersEdgeLogicalinterface structure.
type PutTelephonyProvidersEdgeLogicalinterfaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutTelephonyProvidersEdgeLogicalinterfaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutTelephonyProvidersEdgeLogicalinterfaceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutTelephonyProvidersEdgeLogicalinterfaceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutTelephonyProvidersEdgeLogicalinterfaceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutTelephonyProvidersEdgeLogicalinterfaceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutTelephonyProvidersEdgeLogicalinterfaceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPutTelephonyProvidersEdgeLogicalinterfaceRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPutTelephonyProvidersEdgeLogicalinterfaceRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPutTelephonyProvidersEdgeLogicalinterfaceUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutTelephonyProvidersEdgeLogicalinterfaceTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutTelephonyProvidersEdgeLogicalinterfaceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutTelephonyProvidersEdgeLogicalinterfaceServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutTelephonyProvidersEdgeLogicalinterfaceGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutTelephonyProvidersEdgeLogicalinterfaceOK creates a PutTelephonyProvidersEdgeLogicalinterfaceOK with default headers values
func NewPutTelephonyProvidersEdgeLogicalinterfaceOK() *PutTelephonyProvidersEdgeLogicalinterfaceOK {
	return &PutTelephonyProvidersEdgeLogicalinterfaceOK{}
}

/*
PutTelephonyProvidersEdgeLogicalinterfaceOK describes a response with status code 200, with default header values.

successful operation
*/
type PutTelephonyProvidersEdgeLogicalinterfaceOK struct {
	Payload *models.DomainLogicalInterface
}

// IsSuccess returns true when this put telephony providers edge logicalinterface o k response has a 2xx status code
func (o *PutTelephonyProvidersEdgeLogicalinterfaceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put telephony providers edge logicalinterface o k response has a 3xx status code
func (o *PutTelephonyProvidersEdgeLogicalinterfaceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put telephony providers edge logicalinterface o k response has a 4xx status code
func (o *PutTelephonyProvidersEdgeLogicalinterfaceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put telephony providers edge logicalinterface o k response has a 5xx status code
func (o *PutTelephonyProvidersEdgeLogicalinterfaceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put telephony providers edge logicalinterface o k response a status code equal to that given
func (o *PutTelephonyProvidersEdgeLogicalinterfaceOK) IsCode(code int) bool {
	return code == 200
}

func (o *PutTelephonyProvidersEdgeLogicalinterfaceOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces/{interfaceId}][%d] putTelephonyProvidersEdgeLogicalinterfaceOK  %+v", 200, o.Payload)
}

func (o *PutTelephonyProvidersEdgeLogicalinterfaceOK) String() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces/{interfaceId}][%d] putTelephonyProvidersEdgeLogicalinterfaceOK  %+v", 200, o.Payload)
}

func (o *PutTelephonyProvidersEdgeLogicalinterfaceOK) GetPayload() *models.DomainLogicalInterface {
	return o.Payload
}

func (o *PutTelephonyProvidersEdgeLogicalinterfaceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DomainLogicalInterface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutTelephonyProvidersEdgeLogicalinterfaceBadRequest creates a PutTelephonyProvidersEdgeLogicalinterfaceBadRequest with default headers values
func NewPutTelephonyProvidersEdgeLogicalinterfaceBadRequest() *PutTelephonyProvidersEdgeLogicalinterfaceBadRequest {
	return &PutTelephonyProvidersEdgeLogicalinterfaceBadRequest{}
}

/*
PutTelephonyProvidersEdgeLogicalinterfaceBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutTelephonyProvidersEdgeLogicalinterfaceBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put telephony providers edge logicalinterface bad request response has a 2xx status code
func (o *PutTelephonyProvidersEdgeLogicalinterfaceBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put telephony providers edge logicalinterface bad request response has a 3xx status code
func (o *PutTelephonyProvidersEdgeLogicalinterfaceBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put telephony providers edge logicalinterface bad request response has a 4xx status code
func (o *PutTelephonyProvidersEdgeLogicalinterfaceBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this put telephony providers edge logicalinterface bad request response has a 5xx status code
func (o *PutTelephonyProvidersEdgeLogicalinterfaceBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this put telephony providers edge logicalinterface bad request response a status code equal to that given
func (o *PutTelephonyProvidersEdgeLogicalinterfaceBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PutTelephonyProvidersEdgeLogicalinterfaceBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces/{interfaceId}][%d] putTelephonyProvidersEdgeLogicalinterfaceBadRequest  %+v", 400, o.Payload)
}

func (o *PutTelephonyProvidersEdgeLogicalinterfaceBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces/{interfaceId}][%d] putTelephonyProvidersEdgeLogicalinterfaceBadRequest  %+v", 400, o.Payload)
}

func (o *PutTelephonyProvidersEdgeLogicalinterfaceBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutTelephonyProvidersEdgeLogicalinterfaceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutTelephonyProvidersEdgeLogicalinterfaceUnauthorized creates a PutTelephonyProvidersEdgeLogicalinterfaceUnauthorized with default headers values
func NewPutTelephonyProvidersEdgeLogicalinterfaceUnauthorized() *PutTelephonyProvidersEdgeLogicalinterfaceUnauthorized {
	return &PutTelephonyProvidersEdgeLogicalinterfaceUnauthorized{}
}

/*
PutTelephonyProvidersEdgeLogicalinterfaceUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PutTelephonyProvidersEdgeLogicalinterfaceUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put telephony providers edge logicalinterface unauthorized response has a 2xx status code
func (o *PutTelephonyProvidersEdgeLogicalinterfaceUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put telephony providers edge logicalinterface unauthorized response has a 3xx status code
func (o *PutTelephonyProvidersEdgeLogicalinterfaceUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put telephony providers edge logicalinterface unauthorized response has a 4xx status code
func (o *PutTelephonyProvidersEdgeLogicalinterfaceUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this put telephony providers edge logicalinterface unauthorized response has a 5xx status code
func (o *PutTelephonyProvidersEdgeLogicalinterfaceUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this put telephony providers edge logicalinterface unauthorized response a status code equal to that given
func (o *PutTelephonyProvidersEdgeLogicalinterfaceUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PutTelephonyProvidersEdgeLogicalinterfaceUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces/{interfaceId}][%d] putTelephonyProvidersEdgeLogicalinterfaceUnauthorized  %+v", 401, o.Payload)
}

func (o *PutTelephonyProvidersEdgeLogicalinterfaceUnauthorized) String() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces/{interfaceId}][%d] putTelephonyProvidersEdgeLogicalinterfaceUnauthorized  %+v", 401, o.Payload)
}

func (o *PutTelephonyProvidersEdgeLogicalinterfaceUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutTelephonyProvidersEdgeLogicalinterfaceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutTelephonyProvidersEdgeLogicalinterfaceForbidden creates a PutTelephonyProvidersEdgeLogicalinterfaceForbidden with default headers values
func NewPutTelephonyProvidersEdgeLogicalinterfaceForbidden() *PutTelephonyProvidersEdgeLogicalinterfaceForbidden {
	return &PutTelephonyProvidersEdgeLogicalinterfaceForbidden{}
}

/*
PutTelephonyProvidersEdgeLogicalinterfaceForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PutTelephonyProvidersEdgeLogicalinterfaceForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put telephony providers edge logicalinterface forbidden response has a 2xx status code
func (o *PutTelephonyProvidersEdgeLogicalinterfaceForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put telephony providers edge logicalinterface forbidden response has a 3xx status code
func (o *PutTelephonyProvidersEdgeLogicalinterfaceForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put telephony providers edge logicalinterface forbidden response has a 4xx status code
func (o *PutTelephonyProvidersEdgeLogicalinterfaceForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this put telephony providers edge logicalinterface forbidden response has a 5xx status code
func (o *PutTelephonyProvidersEdgeLogicalinterfaceForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this put telephony providers edge logicalinterface forbidden response a status code equal to that given
func (o *PutTelephonyProvidersEdgeLogicalinterfaceForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PutTelephonyProvidersEdgeLogicalinterfaceForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces/{interfaceId}][%d] putTelephonyProvidersEdgeLogicalinterfaceForbidden  %+v", 403, o.Payload)
}

func (o *PutTelephonyProvidersEdgeLogicalinterfaceForbidden) String() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces/{interfaceId}][%d] putTelephonyProvidersEdgeLogicalinterfaceForbidden  %+v", 403, o.Payload)
}

func (o *PutTelephonyProvidersEdgeLogicalinterfaceForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutTelephonyProvidersEdgeLogicalinterfaceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutTelephonyProvidersEdgeLogicalinterfaceNotFound creates a PutTelephonyProvidersEdgeLogicalinterfaceNotFound with default headers values
func NewPutTelephonyProvidersEdgeLogicalinterfaceNotFound() *PutTelephonyProvidersEdgeLogicalinterfaceNotFound {
	return &PutTelephonyProvidersEdgeLogicalinterfaceNotFound{}
}

/*
PutTelephonyProvidersEdgeLogicalinterfaceNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PutTelephonyProvidersEdgeLogicalinterfaceNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put telephony providers edge logicalinterface not found response has a 2xx status code
func (o *PutTelephonyProvidersEdgeLogicalinterfaceNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put telephony providers edge logicalinterface not found response has a 3xx status code
func (o *PutTelephonyProvidersEdgeLogicalinterfaceNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put telephony providers edge logicalinterface not found response has a 4xx status code
func (o *PutTelephonyProvidersEdgeLogicalinterfaceNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this put telephony providers edge logicalinterface not found response has a 5xx status code
func (o *PutTelephonyProvidersEdgeLogicalinterfaceNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this put telephony providers edge logicalinterface not found response a status code equal to that given
func (o *PutTelephonyProvidersEdgeLogicalinterfaceNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PutTelephonyProvidersEdgeLogicalinterfaceNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces/{interfaceId}][%d] putTelephonyProvidersEdgeLogicalinterfaceNotFound  %+v", 404, o.Payload)
}

func (o *PutTelephonyProvidersEdgeLogicalinterfaceNotFound) String() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces/{interfaceId}][%d] putTelephonyProvidersEdgeLogicalinterfaceNotFound  %+v", 404, o.Payload)
}

func (o *PutTelephonyProvidersEdgeLogicalinterfaceNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutTelephonyProvidersEdgeLogicalinterfaceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutTelephonyProvidersEdgeLogicalinterfaceRequestTimeout creates a PutTelephonyProvidersEdgeLogicalinterfaceRequestTimeout with default headers values
func NewPutTelephonyProvidersEdgeLogicalinterfaceRequestTimeout() *PutTelephonyProvidersEdgeLogicalinterfaceRequestTimeout {
	return &PutTelephonyProvidersEdgeLogicalinterfaceRequestTimeout{}
}

/*
PutTelephonyProvidersEdgeLogicalinterfaceRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PutTelephonyProvidersEdgeLogicalinterfaceRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put telephony providers edge logicalinterface request timeout response has a 2xx status code
func (o *PutTelephonyProvidersEdgeLogicalinterfaceRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put telephony providers edge logicalinterface request timeout response has a 3xx status code
func (o *PutTelephonyProvidersEdgeLogicalinterfaceRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put telephony providers edge logicalinterface request timeout response has a 4xx status code
func (o *PutTelephonyProvidersEdgeLogicalinterfaceRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this put telephony providers edge logicalinterface request timeout response has a 5xx status code
func (o *PutTelephonyProvidersEdgeLogicalinterfaceRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this put telephony providers edge logicalinterface request timeout response a status code equal to that given
func (o *PutTelephonyProvidersEdgeLogicalinterfaceRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PutTelephonyProvidersEdgeLogicalinterfaceRequestTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces/{interfaceId}][%d] putTelephonyProvidersEdgeLogicalinterfaceRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutTelephonyProvidersEdgeLogicalinterfaceRequestTimeout) String() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces/{interfaceId}][%d] putTelephonyProvidersEdgeLogicalinterfaceRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutTelephonyProvidersEdgeLogicalinterfaceRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutTelephonyProvidersEdgeLogicalinterfaceRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutTelephonyProvidersEdgeLogicalinterfaceRequestEntityTooLarge creates a PutTelephonyProvidersEdgeLogicalinterfaceRequestEntityTooLarge with default headers values
func NewPutTelephonyProvidersEdgeLogicalinterfaceRequestEntityTooLarge() *PutTelephonyProvidersEdgeLogicalinterfaceRequestEntityTooLarge {
	return &PutTelephonyProvidersEdgeLogicalinterfaceRequestEntityTooLarge{}
}

/*
PutTelephonyProvidersEdgeLogicalinterfaceRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PutTelephonyProvidersEdgeLogicalinterfaceRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put telephony providers edge logicalinterface request entity too large response has a 2xx status code
func (o *PutTelephonyProvidersEdgeLogicalinterfaceRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put telephony providers edge logicalinterface request entity too large response has a 3xx status code
func (o *PutTelephonyProvidersEdgeLogicalinterfaceRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put telephony providers edge logicalinterface request entity too large response has a 4xx status code
func (o *PutTelephonyProvidersEdgeLogicalinterfaceRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this put telephony providers edge logicalinterface request entity too large response has a 5xx status code
func (o *PutTelephonyProvidersEdgeLogicalinterfaceRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this put telephony providers edge logicalinterface request entity too large response a status code equal to that given
func (o *PutTelephonyProvidersEdgeLogicalinterfaceRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PutTelephonyProvidersEdgeLogicalinterfaceRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces/{interfaceId}][%d] putTelephonyProvidersEdgeLogicalinterfaceRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutTelephonyProvidersEdgeLogicalinterfaceRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces/{interfaceId}][%d] putTelephonyProvidersEdgeLogicalinterfaceRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutTelephonyProvidersEdgeLogicalinterfaceRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutTelephonyProvidersEdgeLogicalinterfaceRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutTelephonyProvidersEdgeLogicalinterfaceUnsupportedMediaType creates a PutTelephonyProvidersEdgeLogicalinterfaceUnsupportedMediaType with default headers values
func NewPutTelephonyProvidersEdgeLogicalinterfaceUnsupportedMediaType() *PutTelephonyProvidersEdgeLogicalinterfaceUnsupportedMediaType {
	return &PutTelephonyProvidersEdgeLogicalinterfaceUnsupportedMediaType{}
}

/*
PutTelephonyProvidersEdgeLogicalinterfaceUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutTelephonyProvidersEdgeLogicalinterfaceUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put telephony providers edge logicalinterface unsupported media type response has a 2xx status code
func (o *PutTelephonyProvidersEdgeLogicalinterfaceUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put telephony providers edge logicalinterface unsupported media type response has a 3xx status code
func (o *PutTelephonyProvidersEdgeLogicalinterfaceUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put telephony providers edge logicalinterface unsupported media type response has a 4xx status code
func (o *PutTelephonyProvidersEdgeLogicalinterfaceUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this put telephony providers edge logicalinterface unsupported media type response has a 5xx status code
func (o *PutTelephonyProvidersEdgeLogicalinterfaceUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this put telephony providers edge logicalinterface unsupported media type response a status code equal to that given
func (o *PutTelephonyProvidersEdgeLogicalinterfaceUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PutTelephonyProvidersEdgeLogicalinterfaceUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces/{interfaceId}][%d] putTelephonyProvidersEdgeLogicalinterfaceUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutTelephonyProvidersEdgeLogicalinterfaceUnsupportedMediaType) String() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces/{interfaceId}][%d] putTelephonyProvidersEdgeLogicalinterfaceUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutTelephonyProvidersEdgeLogicalinterfaceUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutTelephonyProvidersEdgeLogicalinterfaceUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutTelephonyProvidersEdgeLogicalinterfaceTooManyRequests creates a PutTelephonyProvidersEdgeLogicalinterfaceTooManyRequests with default headers values
func NewPutTelephonyProvidersEdgeLogicalinterfaceTooManyRequests() *PutTelephonyProvidersEdgeLogicalinterfaceTooManyRequests {
	return &PutTelephonyProvidersEdgeLogicalinterfaceTooManyRequests{}
}

/*
PutTelephonyProvidersEdgeLogicalinterfaceTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PutTelephonyProvidersEdgeLogicalinterfaceTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put telephony providers edge logicalinterface too many requests response has a 2xx status code
func (o *PutTelephonyProvidersEdgeLogicalinterfaceTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put telephony providers edge logicalinterface too many requests response has a 3xx status code
func (o *PutTelephonyProvidersEdgeLogicalinterfaceTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put telephony providers edge logicalinterface too many requests response has a 4xx status code
func (o *PutTelephonyProvidersEdgeLogicalinterfaceTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this put telephony providers edge logicalinterface too many requests response has a 5xx status code
func (o *PutTelephonyProvidersEdgeLogicalinterfaceTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this put telephony providers edge logicalinterface too many requests response a status code equal to that given
func (o *PutTelephonyProvidersEdgeLogicalinterfaceTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PutTelephonyProvidersEdgeLogicalinterfaceTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces/{interfaceId}][%d] putTelephonyProvidersEdgeLogicalinterfaceTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutTelephonyProvidersEdgeLogicalinterfaceTooManyRequests) String() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces/{interfaceId}][%d] putTelephonyProvidersEdgeLogicalinterfaceTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutTelephonyProvidersEdgeLogicalinterfaceTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutTelephonyProvidersEdgeLogicalinterfaceTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutTelephonyProvidersEdgeLogicalinterfaceInternalServerError creates a PutTelephonyProvidersEdgeLogicalinterfaceInternalServerError with default headers values
func NewPutTelephonyProvidersEdgeLogicalinterfaceInternalServerError() *PutTelephonyProvidersEdgeLogicalinterfaceInternalServerError {
	return &PutTelephonyProvidersEdgeLogicalinterfaceInternalServerError{}
}

/*
PutTelephonyProvidersEdgeLogicalinterfaceInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutTelephonyProvidersEdgeLogicalinterfaceInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put telephony providers edge logicalinterface internal server error response has a 2xx status code
func (o *PutTelephonyProvidersEdgeLogicalinterfaceInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put telephony providers edge logicalinterface internal server error response has a 3xx status code
func (o *PutTelephonyProvidersEdgeLogicalinterfaceInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put telephony providers edge logicalinterface internal server error response has a 4xx status code
func (o *PutTelephonyProvidersEdgeLogicalinterfaceInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this put telephony providers edge logicalinterface internal server error response has a 5xx status code
func (o *PutTelephonyProvidersEdgeLogicalinterfaceInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this put telephony providers edge logicalinterface internal server error response a status code equal to that given
func (o *PutTelephonyProvidersEdgeLogicalinterfaceInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PutTelephonyProvidersEdgeLogicalinterfaceInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces/{interfaceId}][%d] putTelephonyProvidersEdgeLogicalinterfaceInternalServerError  %+v", 500, o.Payload)
}

func (o *PutTelephonyProvidersEdgeLogicalinterfaceInternalServerError) String() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces/{interfaceId}][%d] putTelephonyProvidersEdgeLogicalinterfaceInternalServerError  %+v", 500, o.Payload)
}

func (o *PutTelephonyProvidersEdgeLogicalinterfaceInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutTelephonyProvidersEdgeLogicalinterfaceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutTelephonyProvidersEdgeLogicalinterfaceServiceUnavailable creates a PutTelephonyProvidersEdgeLogicalinterfaceServiceUnavailable with default headers values
func NewPutTelephonyProvidersEdgeLogicalinterfaceServiceUnavailable() *PutTelephonyProvidersEdgeLogicalinterfaceServiceUnavailable {
	return &PutTelephonyProvidersEdgeLogicalinterfaceServiceUnavailable{}
}

/*
PutTelephonyProvidersEdgeLogicalinterfaceServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutTelephonyProvidersEdgeLogicalinterfaceServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put telephony providers edge logicalinterface service unavailable response has a 2xx status code
func (o *PutTelephonyProvidersEdgeLogicalinterfaceServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put telephony providers edge logicalinterface service unavailable response has a 3xx status code
func (o *PutTelephonyProvidersEdgeLogicalinterfaceServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put telephony providers edge logicalinterface service unavailable response has a 4xx status code
func (o *PutTelephonyProvidersEdgeLogicalinterfaceServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this put telephony providers edge logicalinterface service unavailable response has a 5xx status code
func (o *PutTelephonyProvidersEdgeLogicalinterfaceServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this put telephony providers edge logicalinterface service unavailable response a status code equal to that given
func (o *PutTelephonyProvidersEdgeLogicalinterfaceServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PutTelephonyProvidersEdgeLogicalinterfaceServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces/{interfaceId}][%d] putTelephonyProvidersEdgeLogicalinterfaceServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutTelephonyProvidersEdgeLogicalinterfaceServiceUnavailable) String() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces/{interfaceId}][%d] putTelephonyProvidersEdgeLogicalinterfaceServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutTelephonyProvidersEdgeLogicalinterfaceServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutTelephonyProvidersEdgeLogicalinterfaceServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutTelephonyProvidersEdgeLogicalinterfaceGatewayTimeout creates a PutTelephonyProvidersEdgeLogicalinterfaceGatewayTimeout with default headers values
func NewPutTelephonyProvidersEdgeLogicalinterfaceGatewayTimeout() *PutTelephonyProvidersEdgeLogicalinterfaceGatewayTimeout {
	return &PutTelephonyProvidersEdgeLogicalinterfaceGatewayTimeout{}
}

/*
PutTelephonyProvidersEdgeLogicalinterfaceGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PutTelephonyProvidersEdgeLogicalinterfaceGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put telephony providers edge logicalinterface gateway timeout response has a 2xx status code
func (o *PutTelephonyProvidersEdgeLogicalinterfaceGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put telephony providers edge logicalinterface gateway timeout response has a 3xx status code
func (o *PutTelephonyProvidersEdgeLogicalinterfaceGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put telephony providers edge logicalinterface gateway timeout response has a 4xx status code
func (o *PutTelephonyProvidersEdgeLogicalinterfaceGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this put telephony providers edge logicalinterface gateway timeout response has a 5xx status code
func (o *PutTelephonyProvidersEdgeLogicalinterfaceGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this put telephony providers edge logicalinterface gateway timeout response a status code equal to that given
func (o *PutTelephonyProvidersEdgeLogicalinterfaceGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PutTelephonyProvidersEdgeLogicalinterfaceGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces/{interfaceId}][%d] putTelephonyProvidersEdgeLogicalinterfaceGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutTelephonyProvidersEdgeLogicalinterfaceGatewayTimeout) String() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces/{interfaceId}][%d] putTelephonyProvidersEdgeLogicalinterfaceGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutTelephonyProvidersEdgeLogicalinterfaceGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutTelephonyProvidersEdgeLogicalinterfaceGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

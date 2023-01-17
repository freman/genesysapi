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

// GetTelephonyProvidersEdgesCertificateauthorityReader is a Reader for the GetTelephonyProvidersEdgesCertificateauthority structure.
type GetTelephonyProvidersEdgesCertificateauthorityReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTelephonyProvidersEdgesCertificateauthorityReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTelephonyProvidersEdgesCertificateauthorityOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetTelephonyProvidersEdgesCertificateauthorityBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetTelephonyProvidersEdgesCertificateauthorityUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetTelephonyProvidersEdgesCertificateauthorityForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTelephonyProvidersEdgesCertificateauthorityNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetTelephonyProvidersEdgesCertificateauthorityRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetTelephonyProvidersEdgesCertificateauthorityRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetTelephonyProvidersEdgesCertificateauthorityUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetTelephonyProvidersEdgesCertificateauthorityTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetTelephonyProvidersEdgesCertificateauthorityInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetTelephonyProvidersEdgesCertificateauthorityServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetTelephonyProvidersEdgesCertificateauthorityGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTelephonyProvidersEdgesCertificateauthorityOK creates a GetTelephonyProvidersEdgesCertificateauthorityOK with default headers values
func NewGetTelephonyProvidersEdgesCertificateauthorityOK() *GetTelephonyProvidersEdgesCertificateauthorityOK {
	return &GetTelephonyProvidersEdgesCertificateauthorityOK{}
}

/*
GetTelephonyProvidersEdgesCertificateauthorityOK describes a response with status code 200, with default header values.

successful operation
*/
type GetTelephonyProvidersEdgesCertificateauthorityOK struct {
	Payload *models.DomainCertificateAuthority
}

// IsSuccess returns true when this get telephony providers edges certificateauthority o k response has a 2xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthorityOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get telephony providers edges certificateauthority o k response has a 3xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthorityOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges certificateauthority o k response has a 4xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthorityOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get telephony providers edges certificateauthority o k response has a 5xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthorityOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edges certificateauthority o k response a status code equal to that given
func (o *GetTelephonyProvidersEdgesCertificateauthorityOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetTelephonyProvidersEdgesCertificateauthorityOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/certificateauthorities/{certificateId}][%d] getTelephonyProvidersEdgesCertificateauthorityOK  %+v", 200, o.Payload)
}

func (o *GetTelephonyProvidersEdgesCertificateauthorityOK) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/certificateauthorities/{certificateId}][%d] getTelephonyProvidersEdgesCertificateauthorityOK  %+v", 200, o.Payload)
}

func (o *GetTelephonyProvidersEdgesCertificateauthorityOK) GetPayload() *models.DomainCertificateAuthority {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesCertificateauthorityOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DomainCertificateAuthority)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesCertificateauthorityBadRequest creates a GetTelephonyProvidersEdgesCertificateauthorityBadRequest with default headers values
func NewGetTelephonyProvidersEdgesCertificateauthorityBadRequest() *GetTelephonyProvidersEdgesCertificateauthorityBadRequest {
	return &GetTelephonyProvidersEdgesCertificateauthorityBadRequest{}
}

/*
GetTelephonyProvidersEdgesCertificateauthorityBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetTelephonyProvidersEdgesCertificateauthorityBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edges certificateauthority bad request response has a 2xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthorityBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edges certificateauthority bad request response has a 3xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthorityBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges certificateauthority bad request response has a 4xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthorityBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony providers edges certificateauthority bad request response has a 5xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthorityBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edges certificateauthority bad request response a status code equal to that given
func (o *GetTelephonyProvidersEdgesCertificateauthorityBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetTelephonyProvidersEdgesCertificateauthorityBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/certificateauthorities/{certificateId}][%d] getTelephonyProvidersEdgesCertificateauthorityBadRequest  %+v", 400, o.Payload)
}

func (o *GetTelephonyProvidersEdgesCertificateauthorityBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/certificateauthorities/{certificateId}][%d] getTelephonyProvidersEdgesCertificateauthorityBadRequest  %+v", 400, o.Payload)
}

func (o *GetTelephonyProvidersEdgesCertificateauthorityBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesCertificateauthorityBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesCertificateauthorityUnauthorized creates a GetTelephonyProvidersEdgesCertificateauthorityUnauthorized with default headers values
func NewGetTelephonyProvidersEdgesCertificateauthorityUnauthorized() *GetTelephonyProvidersEdgesCertificateauthorityUnauthorized {
	return &GetTelephonyProvidersEdgesCertificateauthorityUnauthorized{}
}

/*
GetTelephonyProvidersEdgesCertificateauthorityUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetTelephonyProvidersEdgesCertificateauthorityUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edges certificateauthority unauthorized response has a 2xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthorityUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edges certificateauthority unauthorized response has a 3xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthorityUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges certificateauthority unauthorized response has a 4xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthorityUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony providers edges certificateauthority unauthorized response has a 5xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthorityUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edges certificateauthority unauthorized response a status code equal to that given
func (o *GetTelephonyProvidersEdgesCertificateauthorityUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetTelephonyProvidersEdgesCertificateauthorityUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/certificateauthorities/{certificateId}][%d] getTelephonyProvidersEdgesCertificateauthorityUnauthorized  %+v", 401, o.Payload)
}

func (o *GetTelephonyProvidersEdgesCertificateauthorityUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/certificateauthorities/{certificateId}][%d] getTelephonyProvidersEdgesCertificateauthorityUnauthorized  %+v", 401, o.Payload)
}

func (o *GetTelephonyProvidersEdgesCertificateauthorityUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesCertificateauthorityUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesCertificateauthorityForbidden creates a GetTelephonyProvidersEdgesCertificateauthorityForbidden with default headers values
func NewGetTelephonyProvidersEdgesCertificateauthorityForbidden() *GetTelephonyProvidersEdgesCertificateauthorityForbidden {
	return &GetTelephonyProvidersEdgesCertificateauthorityForbidden{}
}

/*
GetTelephonyProvidersEdgesCertificateauthorityForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetTelephonyProvidersEdgesCertificateauthorityForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edges certificateauthority forbidden response has a 2xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthorityForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edges certificateauthority forbidden response has a 3xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthorityForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges certificateauthority forbidden response has a 4xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthorityForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony providers edges certificateauthority forbidden response has a 5xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthorityForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edges certificateauthority forbidden response a status code equal to that given
func (o *GetTelephonyProvidersEdgesCertificateauthorityForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetTelephonyProvidersEdgesCertificateauthorityForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/certificateauthorities/{certificateId}][%d] getTelephonyProvidersEdgesCertificateauthorityForbidden  %+v", 403, o.Payload)
}

func (o *GetTelephonyProvidersEdgesCertificateauthorityForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/certificateauthorities/{certificateId}][%d] getTelephonyProvidersEdgesCertificateauthorityForbidden  %+v", 403, o.Payload)
}

func (o *GetTelephonyProvidersEdgesCertificateauthorityForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesCertificateauthorityForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesCertificateauthorityNotFound creates a GetTelephonyProvidersEdgesCertificateauthorityNotFound with default headers values
func NewGetTelephonyProvidersEdgesCertificateauthorityNotFound() *GetTelephonyProvidersEdgesCertificateauthorityNotFound {
	return &GetTelephonyProvidersEdgesCertificateauthorityNotFound{}
}

/*
GetTelephonyProvidersEdgesCertificateauthorityNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetTelephonyProvidersEdgesCertificateauthorityNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edges certificateauthority not found response has a 2xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthorityNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edges certificateauthority not found response has a 3xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthorityNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges certificateauthority not found response has a 4xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthorityNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony providers edges certificateauthority not found response has a 5xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthorityNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edges certificateauthority not found response a status code equal to that given
func (o *GetTelephonyProvidersEdgesCertificateauthorityNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetTelephonyProvidersEdgesCertificateauthorityNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/certificateauthorities/{certificateId}][%d] getTelephonyProvidersEdgesCertificateauthorityNotFound  %+v", 404, o.Payload)
}

func (o *GetTelephonyProvidersEdgesCertificateauthorityNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/certificateauthorities/{certificateId}][%d] getTelephonyProvidersEdgesCertificateauthorityNotFound  %+v", 404, o.Payload)
}

func (o *GetTelephonyProvidersEdgesCertificateauthorityNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesCertificateauthorityNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesCertificateauthorityRequestTimeout creates a GetTelephonyProvidersEdgesCertificateauthorityRequestTimeout with default headers values
func NewGetTelephonyProvidersEdgesCertificateauthorityRequestTimeout() *GetTelephonyProvidersEdgesCertificateauthorityRequestTimeout {
	return &GetTelephonyProvidersEdgesCertificateauthorityRequestTimeout{}
}

/*
GetTelephonyProvidersEdgesCertificateauthorityRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetTelephonyProvidersEdgesCertificateauthorityRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edges certificateauthority request timeout response has a 2xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthorityRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edges certificateauthority request timeout response has a 3xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthorityRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges certificateauthority request timeout response has a 4xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthorityRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony providers edges certificateauthority request timeout response has a 5xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthorityRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edges certificateauthority request timeout response a status code equal to that given
func (o *GetTelephonyProvidersEdgesCertificateauthorityRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetTelephonyProvidersEdgesCertificateauthorityRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/certificateauthorities/{certificateId}][%d] getTelephonyProvidersEdgesCertificateauthorityRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetTelephonyProvidersEdgesCertificateauthorityRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/certificateauthorities/{certificateId}][%d] getTelephonyProvidersEdgesCertificateauthorityRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetTelephonyProvidersEdgesCertificateauthorityRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesCertificateauthorityRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesCertificateauthorityRequestEntityTooLarge creates a GetTelephonyProvidersEdgesCertificateauthorityRequestEntityTooLarge with default headers values
func NewGetTelephonyProvidersEdgesCertificateauthorityRequestEntityTooLarge() *GetTelephonyProvidersEdgesCertificateauthorityRequestEntityTooLarge {
	return &GetTelephonyProvidersEdgesCertificateauthorityRequestEntityTooLarge{}
}

/*
GetTelephonyProvidersEdgesCertificateauthorityRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetTelephonyProvidersEdgesCertificateauthorityRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edges certificateauthority request entity too large response has a 2xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthorityRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edges certificateauthority request entity too large response has a 3xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthorityRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges certificateauthority request entity too large response has a 4xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthorityRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony providers edges certificateauthority request entity too large response has a 5xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthorityRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edges certificateauthority request entity too large response a status code equal to that given
func (o *GetTelephonyProvidersEdgesCertificateauthorityRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetTelephonyProvidersEdgesCertificateauthorityRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/certificateauthorities/{certificateId}][%d] getTelephonyProvidersEdgesCertificateauthorityRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetTelephonyProvidersEdgesCertificateauthorityRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/certificateauthorities/{certificateId}][%d] getTelephonyProvidersEdgesCertificateauthorityRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetTelephonyProvidersEdgesCertificateauthorityRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesCertificateauthorityRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesCertificateauthorityUnsupportedMediaType creates a GetTelephonyProvidersEdgesCertificateauthorityUnsupportedMediaType with default headers values
func NewGetTelephonyProvidersEdgesCertificateauthorityUnsupportedMediaType() *GetTelephonyProvidersEdgesCertificateauthorityUnsupportedMediaType {
	return &GetTelephonyProvidersEdgesCertificateauthorityUnsupportedMediaType{}
}

/*
GetTelephonyProvidersEdgesCertificateauthorityUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetTelephonyProvidersEdgesCertificateauthorityUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edges certificateauthority unsupported media type response has a 2xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthorityUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edges certificateauthority unsupported media type response has a 3xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthorityUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges certificateauthority unsupported media type response has a 4xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthorityUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony providers edges certificateauthority unsupported media type response has a 5xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthorityUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edges certificateauthority unsupported media type response a status code equal to that given
func (o *GetTelephonyProvidersEdgesCertificateauthorityUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetTelephonyProvidersEdgesCertificateauthorityUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/certificateauthorities/{certificateId}][%d] getTelephonyProvidersEdgesCertificateauthorityUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetTelephonyProvidersEdgesCertificateauthorityUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/certificateauthorities/{certificateId}][%d] getTelephonyProvidersEdgesCertificateauthorityUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetTelephonyProvidersEdgesCertificateauthorityUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesCertificateauthorityUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesCertificateauthorityTooManyRequests creates a GetTelephonyProvidersEdgesCertificateauthorityTooManyRequests with default headers values
func NewGetTelephonyProvidersEdgesCertificateauthorityTooManyRequests() *GetTelephonyProvidersEdgesCertificateauthorityTooManyRequests {
	return &GetTelephonyProvidersEdgesCertificateauthorityTooManyRequests{}
}

/*
GetTelephonyProvidersEdgesCertificateauthorityTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetTelephonyProvidersEdgesCertificateauthorityTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edges certificateauthority too many requests response has a 2xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthorityTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edges certificateauthority too many requests response has a 3xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthorityTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges certificateauthority too many requests response has a 4xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthorityTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony providers edges certificateauthority too many requests response has a 5xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthorityTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edges certificateauthority too many requests response a status code equal to that given
func (o *GetTelephonyProvidersEdgesCertificateauthorityTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetTelephonyProvidersEdgesCertificateauthorityTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/certificateauthorities/{certificateId}][%d] getTelephonyProvidersEdgesCertificateauthorityTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetTelephonyProvidersEdgesCertificateauthorityTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/certificateauthorities/{certificateId}][%d] getTelephonyProvidersEdgesCertificateauthorityTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetTelephonyProvidersEdgesCertificateauthorityTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesCertificateauthorityTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesCertificateauthorityInternalServerError creates a GetTelephonyProvidersEdgesCertificateauthorityInternalServerError with default headers values
func NewGetTelephonyProvidersEdgesCertificateauthorityInternalServerError() *GetTelephonyProvidersEdgesCertificateauthorityInternalServerError {
	return &GetTelephonyProvidersEdgesCertificateauthorityInternalServerError{}
}

/*
GetTelephonyProvidersEdgesCertificateauthorityInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetTelephonyProvidersEdgesCertificateauthorityInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edges certificateauthority internal server error response has a 2xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthorityInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edges certificateauthority internal server error response has a 3xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthorityInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges certificateauthority internal server error response has a 4xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthorityInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get telephony providers edges certificateauthority internal server error response has a 5xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthorityInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get telephony providers edges certificateauthority internal server error response a status code equal to that given
func (o *GetTelephonyProvidersEdgesCertificateauthorityInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetTelephonyProvidersEdgesCertificateauthorityInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/certificateauthorities/{certificateId}][%d] getTelephonyProvidersEdgesCertificateauthorityInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTelephonyProvidersEdgesCertificateauthorityInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/certificateauthorities/{certificateId}][%d] getTelephonyProvidersEdgesCertificateauthorityInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTelephonyProvidersEdgesCertificateauthorityInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesCertificateauthorityInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesCertificateauthorityServiceUnavailable creates a GetTelephonyProvidersEdgesCertificateauthorityServiceUnavailable with default headers values
func NewGetTelephonyProvidersEdgesCertificateauthorityServiceUnavailable() *GetTelephonyProvidersEdgesCertificateauthorityServiceUnavailable {
	return &GetTelephonyProvidersEdgesCertificateauthorityServiceUnavailable{}
}

/*
GetTelephonyProvidersEdgesCertificateauthorityServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetTelephonyProvidersEdgesCertificateauthorityServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edges certificateauthority service unavailable response has a 2xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthorityServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edges certificateauthority service unavailable response has a 3xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthorityServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges certificateauthority service unavailable response has a 4xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthorityServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get telephony providers edges certificateauthority service unavailable response has a 5xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthorityServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get telephony providers edges certificateauthority service unavailable response a status code equal to that given
func (o *GetTelephonyProvidersEdgesCertificateauthorityServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetTelephonyProvidersEdgesCertificateauthorityServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/certificateauthorities/{certificateId}][%d] getTelephonyProvidersEdgesCertificateauthorityServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetTelephonyProvidersEdgesCertificateauthorityServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/certificateauthorities/{certificateId}][%d] getTelephonyProvidersEdgesCertificateauthorityServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetTelephonyProvidersEdgesCertificateauthorityServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesCertificateauthorityServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesCertificateauthorityGatewayTimeout creates a GetTelephonyProvidersEdgesCertificateauthorityGatewayTimeout with default headers values
func NewGetTelephonyProvidersEdgesCertificateauthorityGatewayTimeout() *GetTelephonyProvidersEdgesCertificateauthorityGatewayTimeout {
	return &GetTelephonyProvidersEdgesCertificateauthorityGatewayTimeout{}
}

/*
GetTelephonyProvidersEdgesCertificateauthorityGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetTelephonyProvidersEdgesCertificateauthorityGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edges certificateauthority gateway timeout response has a 2xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthorityGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edges certificateauthority gateway timeout response has a 3xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthorityGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges certificateauthority gateway timeout response has a 4xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthorityGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get telephony providers edges certificateauthority gateway timeout response has a 5xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthorityGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get telephony providers edges certificateauthority gateway timeout response a status code equal to that given
func (o *GetTelephonyProvidersEdgesCertificateauthorityGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetTelephonyProvidersEdgesCertificateauthorityGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/certificateauthorities/{certificateId}][%d] getTelephonyProvidersEdgesCertificateauthorityGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetTelephonyProvidersEdgesCertificateauthorityGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/certificateauthorities/{certificateId}][%d] getTelephonyProvidersEdgesCertificateauthorityGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetTelephonyProvidersEdgesCertificateauthorityGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesCertificateauthorityGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

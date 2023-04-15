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

// GetTelephonyProvidersEdgesCertificateauthoritiesReader is a Reader for the GetTelephonyProvidersEdgesCertificateauthorities structure.
type GetTelephonyProvidersEdgesCertificateauthoritiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTelephonyProvidersEdgesCertificateauthoritiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTelephonyProvidersEdgesCertificateauthoritiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetTelephonyProvidersEdgesCertificateauthoritiesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetTelephonyProvidersEdgesCertificateauthoritiesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetTelephonyProvidersEdgesCertificateauthoritiesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTelephonyProvidersEdgesCertificateauthoritiesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetTelephonyProvidersEdgesCertificateauthoritiesRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetTelephonyProvidersEdgesCertificateauthoritiesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetTelephonyProvidersEdgesCertificateauthoritiesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetTelephonyProvidersEdgesCertificateauthoritiesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetTelephonyProvidersEdgesCertificateauthoritiesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetTelephonyProvidersEdgesCertificateauthoritiesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetTelephonyProvidersEdgesCertificateauthoritiesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTelephonyProvidersEdgesCertificateauthoritiesOK creates a GetTelephonyProvidersEdgesCertificateauthoritiesOK with default headers values
func NewGetTelephonyProvidersEdgesCertificateauthoritiesOK() *GetTelephonyProvidersEdgesCertificateauthoritiesOK {
	return &GetTelephonyProvidersEdgesCertificateauthoritiesOK{}
}

/*
GetTelephonyProvidersEdgesCertificateauthoritiesOK describes a response with status code 200, with default header values.

successful operation
*/
type GetTelephonyProvidersEdgesCertificateauthoritiesOK struct {
	Payload *models.CertificateAuthorityEntityListing
}

// IsSuccess returns true when this get telephony providers edges certificateauthorities o k response has a 2xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthoritiesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get telephony providers edges certificateauthorities o k response has a 3xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthoritiesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges certificateauthorities o k response has a 4xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthoritiesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get telephony providers edges certificateauthorities o k response has a 5xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthoritiesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edges certificateauthorities o k response a status code equal to that given
func (o *GetTelephonyProvidersEdgesCertificateauthoritiesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetTelephonyProvidersEdgesCertificateauthoritiesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/certificateauthorities][%d] getTelephonyProvidersEdgesCertificateauthoritiesOK  %+v", 200, o.Payload)
}

func (o *GetTelephonyProvidersEdgesCertificateauthoritiesOK) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/certificateauthorities][%d] getTelephonyProvidersEdgesCertificateauthoritiesOK  %+v", 200, o.Payload)
}

func (o *GetTelephonyProvidersEdgesCertificateauthoritiesOK) GetPayload() *models.CertificateAuthorityEntityListing {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesCertificateauthoritiesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CertificateAuthorityEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesCertificateauthoritiesBadRequest creates a GetTelephonyProvidersEdgesCertificateauthoritiesBadRequest with default headers values
func NewGetTelephonyProvidersEdgesCertificateauthoritiesBadRequest() *GetTelephonyProvidersEdgesCertificateauthoritiesBadRequest {
	return &GetTelephonyProvidersEdgesCertificateauthoritiesBadRequest{}
}

/*
GetTelephonyProvidersEdgesCertificateauthoritiesBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetTelephonyProvidersEdgesCertificateauthoritiesBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edges certificateauthorities bad request response has a 2xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthoritiesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edges certificateauthorities bad request response has a 3xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthoritiesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges certificateauthorities bad request response has a 4xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthoritiesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony providers edges certificateauthorities bad request response has a 5xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthoritiesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edges certificateauthorities bad request response a status code equal to that given
func (o *GetTelephonyProvidersEdgesCertificateauthoritiesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetTelephonyProvidersEdgesCertificateauthoritiesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/certificateauthorities][%d] getTelephonyProvidersEdgesCertificateauthoritiesBadRequest  %+v", 400, o.Payload)
}

func (o *GetTelephonyProvidersEdgesCertificateauthoritiesBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/certificateauthorities][%d] getTelephonyProvidersEdgesCertificateauthoritiesBadRequest  %+v", 400, o.Payload)
}

func (o *GetTelephonyProvidersEdgesCertificateauthoritiesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesCertificateauthoritiesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesCertificateauthoritiesUnauthorized creates a GetTelephonyProvidersEdgesCertificateauthoritiesUnauthorized with default headers values
func NewGetTelephonyProvidersEdgesCertificateauthoritiesUnauthorized() *GetTelephonyProvidersEdgesCertificateauthoritiesUnauthorized {
	return &GetTelephonyProvidersEdgesCertificateauthoritiesUnauthorized{}
}

/*
GetTelephonyProvidersEdgesCertificateauthoritiesUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetTelephonyProvidersEdgesCertificateauthoritiesUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edges certificateauthorities unauthorized response has a 2xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthoritiesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edges certificateauthorities unauthorized response has a 3xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthoritiesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges certificateauthorities unauthorized response has a 4xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthoritiesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony providers edges certificateauthorities unauthorized response has a 5xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthoritiesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edges certificateauthorities unauthorized response a status code equal to that given
func (o *GetTelephonyProvidersEdgesCertificateauthoritiesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetTelephonyProvidersEdgesCertificateauthoritiesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/certificateauthorities][%d] getTelephonyProvidersEdgesCertificateauthoritiesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetTelephonyProvidersEdgesCertificateauthoritiesUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/certificateauthorities][%d] getTelephonyProvidersEdgesCertificateauthoritiesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetTelephonyProvidersEdgesCertificateauthoritiesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesCertificateauthoritiesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesCertificateauthoritiesForbidden creates a GetTelephonyProvidersEdgesCertificateauthoritiesForbidden with default headers values
func NewGetTelephonyProvidersEdgesCertificateauthoritiesForbidden() *GetTelephonyProvidersEdgesCertificateauthoritiesForbidden {
	return &GetTelephonyProvidersEdgesCertificateauthoritiesForbidden{}
}

/*
GetTelephonyProvidersEdgesCertificateauthoritiesForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetTelephonyProvidersEdgesCertificateauthoritiesForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edges certificateauthorities forbidden response has a 2xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthoritiesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edges certificateauthorities forbidden response has a 3xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthoritiesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges certificateauthorities forbidden response has a 4xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthoritiesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony providers edges certificateauthorities forbidden response has a 5xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthoritiesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edges certificateauthorities forbidden response a status code equal to that given
func (o *GetTelephonyProvidersEdgesCertificateauthoritiesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetTelephonyProvidersEdgesCertificateauthoritiesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/certificateauthorities][%d] getTelephonyProvidersEdgesCertificateauthoritiesForbidden  %+v", 403, o.Payload)
}

func (o *GetTelephonyProvidersEdgesCertificateauthoritiesForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/certificateauthorities][%d] getTelephonyProvidersEdgesCertificateauthoritiesForbidden  %+v", 403, o.Payload)
}

func (o *GetTelephonyProvidersEdgesCertificateauthoritiesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesCertificateauthoritiesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesCertificateauthoritiesNotFound creates a GetTelephonyProvidersEdgesCertificateauthoritiesNotFound with default headers values
func NewGetTelephonyProvidersEdgesCertificateauthoritiesNotFound() *GetTelephonyProvidersEdgesCertificateauthoritiesNotFound {
	return &GetTelephonyProvidersEdgesCertificateauthoritiesNotFound{}
}

/*
GetTelephonyProvidersEdgesCertificateauthoritiesNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetTelephonyProvidersEdgesCertificateauthoritiesNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edges certificateauthorities not found response has a 2xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthoritiesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edges certificateauthorities not found response has a 3xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthoritiesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges certificateauthorities not found response has a 4xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthoritiesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony providers edges certificateauthorities not found response has a 5xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthoritiesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edges certificateauthorities not found response a status code equal to that given
func (o *GetTelephonyProvidersEdgesCertificateauthoritiesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetTelephonyProvidersEdgesCertificateauthoritiesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/certificateauthorities][%d] getTelephonyProvidersEdgesCertificateauthoritiesNotFound  %+v", 404, o.Payload)
}

func (o *GetTelephonyProvidersEdgesCertificateauthoritiesNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/certificateauthorities][%d] getTelephonyProvidersEdgesCertificateauthoritiesNotFound  %+v", 404, o.Payload)
}

func (o *GetTelephonyProvidersEdgesCertificateauthoritiesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesCertificateauthoritiesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesCertificateauthoritiesRequestTimeout creates a GetTelephonyProvidersEdgesCertificateauthoritiesRequestTimeout with default headers values
func NewGetTelephonyProvidersEdgesCertificateauthoritiesRequestTimeout() *GetTelephonyProvidersEdgesCertificateauthoritiesRequestTimeout {
	return &GetTelephonyProvidersEdgesCertificateauthoritiesRequestTimeout{}
}

/*
GetTelephonyProvidersEdgesCertificateauthoritiesRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetTelephonyProvidersEdgesCertificateauthoritiesRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edges certificateauthorities request timeout response has a 2xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthoritiesRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edges certificateauthorities request timeout response has a 3xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthoritiesRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges certificateauthorities request timeout response has a 4xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthoritiesRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony providers edges certificateauthorities request timeout response has a 5xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthoritiesRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edges certificateauthorities request timeout response a status code equal to that given
func (o *GetTelephonyProvidersEdgesCertificateauthoritiesRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetTelephonyProvidersEdgesCertificateauthoritiesRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/certificateauthorities][%d] getTelephonyProvidersEdgesCertificateauthoritiesRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetTelephonyProvidersEdgesCertificateauthoritiesRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/certificateauthorities][%d] getTelephonyProvidersEdgesCertificateauthoritiesRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetTelephonyProvidersEdgesCertificateauthoritiesRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesCertificateauthoritiesRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesCertificateauthoritiesRequestEntityTooLarge creates a GetTelephonyProvidersEdgesCertificateauthoritiesRequestEntityTooLarge with default headers values
func NewGetTelephonyProvidersEdgesCertificateauthoritiesRequestEntityTooLarge() *GetTelephonyProvidersEdgesCertificateauthoritiesRequestEntityTooLarge {
	return &GetTelephonyProvidersEdgesCertificateauthoritiesRequestEntityTooLarge{}
}

/*
GetTelephonyProvidersEdgesCertificateauthoritiesRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetTelephonyProvidersEdgesCertificateauthoritiesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edges certificateauthorities request entity too large response has a 2xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthoritiesRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edges certificateauthorities request entity too large response has a 3xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthoritiesRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges certificateauthorities request entity too large response has a 4xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthoritiesRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony providers edges certificateauthorities request entity too large response has a 5xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthoritiesRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edges certificateauthorities request entity too large response a status code equal to that given
func (o *GetTelephonyProvidersEdgesCertificateauthoritiesRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetTelephonyProvidersEdgesCertificateauthoritiesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/certificateauthorities][%d] getTelephonyProvidersEdgesCertificateauthoritiesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetTelephonyProvidersEdgesCertificateauthoritiesRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/certificateauthorities][%d] getTelephonyProvidersEdgesCertificateauthoritiesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetTelephonyProvidersEdgesCertificateauthoritiesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesCertificateauthoritiesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesCertificateauthoritiesUnsupportedMediaType creates a GetTelephonyProvidersEdgesCertificateauthoritiesUnsupportedMediaType with default headers values
func NewGetTelephonyProvidersEdgesCertificateauthoritiesUnsupportedMediaType() *GetTelephonyProvidersEdgesCertificateauthoritiesUnsupportedMediaType {
	return &GetTelephonyProvidersEdgesCertificateauthoritiesUnsupportedMediaType{}
}

/*
GetTelephonyProvidersEdgesCertificateauthoritiesUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetTelephonyProvidersEdgesCertificateauthoritiesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edges certificateauthorities unsupported media type response has a 2xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthoritiesUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edges certificateauthorities unsupported media type response has a 3xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthoritiesUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges certificateauthorities unsupported media type response has a 4xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthoritiesUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony providers edges certificateauthorities unsupported media type response has a 5xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthoritiesUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edges certificateauthorities unsupported media type response a status code equal to that given
func (o *GetTelephonyProvidersEdgesCertificateauthoritiesUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetTelephonyProvidersEdgesCertificateauthoritiesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/certificateauthorities][%d] getTelephonyProvidersEdgesCertificateauthoritiesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetTelephonyProvidersEdgesCertificateauthoritiesUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/certificateauthorities][%d] getTelephonyProvidersEdgesCertificateauthoritiesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetTelephonyProvidersEdgesCertificateauthoritiesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesCertificateauthoritiesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesCertificateauthoritiesTooManyRequests creates a GetTelephonyProvidersEdgesCertificateauthoritiesTooManyRequests with default headers values
func NewGetTelephonyProvidersEdgesCertificateauthoritiesTooManyRequests() *GetTelephonyProvidersEdgesCertificateauthoritiesTooManyRequests {
	return &GetTelephonyProvidersEdgesCertificateauthoritiesTooManyRequests{}
}

/*
GetTelephonyProvidersEdgesCertificateauthoritiesTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetTelephonyProvidersEdgesCertificateauthoritiesTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edges certificateauthorities too many requests response has a 2xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthoritiesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edges certificateauthorities too many requests response has a 3xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthoritiesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges certificateauthorities too many requests response has a 4xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthoritiesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony providers edges certificateauthorities too many requests response has a 5xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthoritiesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edges certificateauthorities too many requests response a status code equal to that given
func (o *GetTelephonyProvidersEdgesCertificateauthoritiesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetTelephonyProvidersEdgesCertificateauthoritiesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/certificateauthorities][%d] getTelephonyProvidersEdgesCertificateauthoritiesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetTelephonyProvidersEdgesCertificateauthoritiesTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/certificateauthorities][%d] getTelephonyProvidersEdgesCertificateauthoritiesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetTelephonyProvidersEdgesCertificateauthoritiesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesCertificateauthoritiesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesCertificateauthoritiesInternalServerError creates a GetTelephonyProvidersEdgesCertificateauthoritiesInternalServerError with default headers values
func NewGetTelephonyProvidersEdgesCertificateauthoritiesInternalServerError() *GetTelephonyProvidersEdgesCertificateauthoritiesInternalServerError {
	return &GetTelephonyProvidersEdgesCertificateauthoritiesInternalServerError{}
}

/*
GetTelephonyProvidersEdgesCertificateauthoritiesInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetTelephonyProvidersEdgesCertificateauthoritiesInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edges certificateauthorities internal server error response has a 2xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthoritiesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edges certificateauthorities internal server error response has a 3xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthoritiesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges certificateauthorities internal server error response has a 4xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthoritiesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get telephony providers edges certificateauthorities internal server error response has a 5xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthoritiesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get telephony providers edges certificateauthorities internal server error response a status code equal to that given
func (o *GetTelephonyProvidersEdgesCertificateauthoritiesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetTelephonyProvidersEdgesCertificateauthoritiesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/certificateauthorities][%d] getTelephonyProvidersEdgesCertificateauthoritiesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTelephonyProvidersEdgesCertificateauthoritiesInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/certificateauthorities][%d] getTelephonyProvidersEdgesCertificateauthoritiesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTelephonyProvidersEdgesCertificateauthoritiesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesCertificateauthoritiesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesCertificateauthoritiesServiceUnavailable creates a GetTelephonyProvidersEdgesCertificateauthoritiesServiceUnavailable with default headers values
func NewGetTelephonyProvidersEdgesCertificateauthoritiesServiceUnavailable() *GetTelephonyProvidersEdgesCertificateauthoritiesServiceUnavailable {
	return &GetTelephonyProvidersEdgesCertificateauthoritiesServiceUnavailable{}
}

/*
GetTelephonyProvidersEdgesCertificateauthoritiesServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetTelephonyProvidersEdgesCertificateauthoritiesServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edges certificateauthorities service unavailable response has a 2xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthoritiesServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edges certificateauthorities service unavailable response has a 3xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthoritiesServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges certificateauthorities service unavailable response has a 4xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthoritiesServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get telephony providers edges certificateauthorities service unavailable response has a 5xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthoritiesServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get telephony providers edges certificateauthorities service unavailable response a status code equal to that given
func (o *GetTelephonyProvidersEdgesCertificateauthoritiesServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetTelephonyProvidersEdgesCertificateauthoritiesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/certificateauthorities][%d] getTelephonyProvidersEdgesCertificateauthoritiesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetTelephonyProvidersEdgesCertificateauthoritiesServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/certificateauthorities][%d] getTelephonyProvidersEdgesCertificateauthoritiesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetTelephonyProvidersEdgesCertificateauthoritiesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesCertificateauthoritiesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesCertificateauthoritiesGatewayTimeout creates a GetTelephonyProvidersEdgesCertificateauthoritiesGatewayTimeout with default headers values
func NewGetTelephonyProvidersEdgesCertificateauthoritiesGatewayTimeout() *GetTelephonyProvidersEdgesCertificateauthoritiesGatewayTimeout {
	return &GetTelephonyProvidersEdgesCertificateauthoritiesGatewayTimeout{}
}

/*
GetTelephonyProvidersEdgesCertificateauthoritiesGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetTelephonyProvidersEdgesCertificateauthoritiesGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edges certificateauthorities gateway timeout response has a 2xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthoritiesGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edges certificateauthorities gateway timeout response has a 3xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthoritiesGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges certificateauthorities gateway timeout response has a 4xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthoritiesGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get telephony providers edges certificateauthorities gateway timeout response has a 5xx status code
func (o *GetTelephonyProvidersEdgesCertificateauthoritiesGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get telephony providers edges certificateauthorities gateway timeout response a status code equal to that given
func (o *GetTelephonyProvidersEdgesCertificateauthoritiesGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetTelephonyProvidersEdgesCertificateauthoritiesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/certificateauthorities][%d] getTelephonyProvidersEdgesCertificateauthoritiesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetTelephonyProvidersEdgesCertificateauthoritiesGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/certificateauthorities][%d] getTelephonyProvidersEdgesCertificateauthoritiesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetTelephonyProvidersEdgesCertificateauthoritiesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesCertificateauthoritiesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

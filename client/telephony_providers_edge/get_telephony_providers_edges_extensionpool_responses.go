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

// GetTelephonyProvidersEdgesExtensionpoolReader is a Reader for the GetTelephonyProvidersEdgesExtensionpool structure.
type GetTelephonyProvidersEdgesExtensionpoolReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTelephonyProvidersEdgesExtensionpoolReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTelephonyProvidersEdgesExtensionpoolOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetTelephonyProvidersEdgesExtensionpoolBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetTelephonyProvidersEdgesExtensionpoolUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetTelephonyProvidersEdgesExtensionpoolForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTelephonyProvidersEdgesExtensionpoolNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetTelephonyProvidersEdgesExtensionpoolRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetTelephonyProvidersEdgesExtensionpoolRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetTelephonyProvidersEdgesExtensionpoolUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetTelephonyProvidersEdgesExtensionpoolTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetTelephonyProvidersEdgesExtensionpoolInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetTelephonyProvidersEdgesExtensionpoolServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetTelephonyProvidersEdgesExtensionpoolGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTelephonyProvidersEdgesExtensionpoolOK creates a GetTelephonyProvidersEdgesExtensionpoolOK with default headers values
func NewGetTelephonyProvidersEdgesExtensionpoolOK() *GetTelephonyProvidersEdgesExtensionpoolOK {
	return &GetTelephonyProvidersEdgesExtensionpoolOK{}
}

/*
GetTelephonyProvidersEdgesExtensionpoolOK describes a response with status code 200, with default header values.

successful operation
*/
type GetTelephonyProvidersEdgesExtensionpoolOK struct {
	Payload *models.ExtensionPool
}

// IsSuccess returns true when this get telephony providers edges extensionpool o k response has a 2xx status code
func (o *GetTelephonyProvidersEdgesExtensionpoolOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get telephony providers edges extensionpool o k response has a 3xx status code
func (o *GetTelephonyProvidersEdgesExtensionpoolOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges extensionpool o k response has a 4xx status code
func (o *GetTelephonyProvidersEdgesExtensionpoolOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get telephony providers edges extensionpool o k response has a 5xx status code
func (o *GetTelephonyProvidersEdgesExtensionpoolOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edges extensionpool o k response a status code equal to that given
func (o *GetTelephonyProvidersEdgesExtensionpoolOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetTelephonyProvidersEdgesExtensionpoolOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/extensionpools/{extensionPoolId}][%d] getTelephonyProvidersEdgesExtensionpoolOK  %+v", 200, o.Payload)
}

func (o *GetTelephonyProvidersEdgesExtensionpoolOK) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/extensionpools/{extensionPoolId}][%d] getTelephonyProvidersEdgesExtensionpoolOK  %+v", 200, o.Payload)
}

func (o *GetTelephonyProvidersEdgesExtensionpoolOK) GetPayload() *models.ExtensionPool {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesExtensionpoolOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ExtensionPool)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesExtensionpoolBadRequest creates a GetTelephonyProvidersEdgesExtensionpoolBadRequest with default headers values
func NewGetTelephonyProvidersEdgesExtensionpoolBadRequest() *GetTelephonyProvidersEdgesExtensionpoolBadRequest {
	return &GetTelephonyProvidersEdgesExtensionpoolBadRequest{}
}

/*
GetTelephonyProvidersEdgesExtensionpoolBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetTelephonyProvidersEdgesExtensionpoolBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edges extensionpool bad request response has a 2xx status code
func (o *GetTelephonyProvidersEdgesExtensionpoolBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edges extensionpool bad request response has a 3xx status code
func (o *GetTelephonyProvidersEdgesExtensionpoolBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges extensionpool bad request response has a 4xx status code
func (o *GetTelephonyProvidersEdgesExtensionpoolBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony providers edges extensionpool bad request response has a 5xx status code
func (o *GetTelephonyProvidersEdgesExtensionpoolBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edges extensionpool bad request response a status code equal to that given
func (o *GetTelephonyProvidersEdgesExtensionpoolBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetTelephonyProvidersEdgesExtensionpoolBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/extensionpools/{extensionPoolId}][%d] getTelephonyProvidersEdgesExtensionpoolBadRequest  %+v", 400, o.Payload)
}

func (o *GetTelephonyProvidersEdgesExtensionpoolBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/extensionpools/{extensionPoolId}][%d] getTelephonyProvidersEdgesExtensionpoolBadRequest  %+v", 400, o.Payload)
}

func (o *GetTelephonyProvidersEdgesExtensionpoolBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesExtensionpoolBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesExtensionpoolUnauthorized creates a GetTelephonyProvidersEdgesExtensionpoolUnauthorized with default headers values
func NewGetTelephonyProvidersEdgesExtensionpoolUnauthorized() *GetTelephonyProvidersEdgesExtensionpoolUnauthorized {
	return &GetTelephonyProvidersEdgesExtensionpoolUnauthorized{}
}

/*
GetTelephonyProvidersEdgesExtensionpoolUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetTelephonyProvidersEdgesExtensionpoolUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edges extensionpool unauthorized response has a 2xx status code
func (o *GetTelephonyProvidersEdgesExtensionpoolUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edges extensionpool unauthorized response has a 3xx status code
func (o *GetTelephonyProvidersEdgesExtensionpoolUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges extensionpool unauthorized response has a 4xx status code
func (o *GetTelephonyProvidersEdgesExtensionpoolUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony providers edges extensionpool unauthorized response has a 5xx status code
func (o *GetTelephonyProvidersEdgesExtensionpoolUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edges extensionpool unauthorized response a status code equal to that given
func (o *GetTelephonyProvidersEdgesExtensionpoolUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetTelephonyProvidersEdgesExtensionpoolUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/extensionpools/{extensionPoolId}][%d] getTelephonyProvidersEdgesExtensionpoolUnauthorized  %+v", 401, o.Payload)
}

func (o *GetTelephonyProvidersEdgesExtensionpoolUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/extensionpools/{extensionPoolId}][%d] getTelephonyProvidersEdgesExtensionpoolUnauthorized  %+v", 401, o.Payload)
}

func (o *GetTelephonyProvidersEdgesExtensionpoolUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesExtensionpoolUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesExtensionpoolForbidden creates a GetTelephonyProvidersEdgesExtensionpoolForbidden with default headers values
func NewGetTelephonyProvidersEdgesExtensionpoolForbidden() *GetTelephonyProvidersEdgesExtensionpoolForbidden {
	return &GetTelephonyProvidersEdgesExtensionpoolForbidden{}
}

/*
GetTelephonyProvidersEdgesExtensionpoolForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetTelephonyProvidersEdgesExtensionpoolForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edges extensionpool forbidden response has a 2xx status code
func (o *GetTelephonyProvidersEdgesExtensionpoolForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edges extensionpool forbidden response has a 3xx status code
func (o *GetTelephonyProvidersEdgesExtensionpoolForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges extensionpool forbidden response has a 4xx status code
func (o *GetTelephonyProvidersEdgesExtensionpoolForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony providers edges extensionpool forbidden response has a 5xx status code
func (o *GetTelephonyProvidersEdgesExtensionpoolForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edges extensionpool forbidden response a status code equal to that given
func (o *GetTelephonyProvidersEdgesExtensionpoolForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetTelephonyProvidersEdgesExtensionpoolForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/extensionpools/{extensionPoolId}][%d] getTelephonyProvidersEdgesExtensionpoolForbidden  %+v", 403, o.Payload)
}

func (o *GetTelephonyProvidersEdgesExtensionpoolForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/extensionpools/{extensionPoolId}][%d] getTelephonyProvidersEdgesExtensionpoolForbidden  %+v", 403, o.Payload)
}

func (o *GetTelephonyProvidersEdgesExtensionpoolForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesExtensionpoolForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesExtensionpoolNotFound creates a GetTelephonyProvidersEdgesExtensionpoolNotFound with default headers values
func NewGetTelephonyProvidersEdgesExtensionpoolNotFound() *GetTelephonyProvidersEdgesExtensionpoolNotFound {
	return &GetTelephonyProvidersEdgesExtensionpoolNotFound{}
}

/*
GetTelephonyProvidersEdgesExtensionpoolNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetTelephonyProvidersEdgesExtensionpoolNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edges extensionpool not found response has a 2xx status code
func (o *GetTelephonyProvidersEdgesExtensionpoolNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edges extensionpool not found response has a 3xx status code
func (o *GetTelephonyProvidersEdgesExtensionpoolNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges extensionpool not found response has a 4xx status code
func (o *GetTelephonyProvidersEdgesExtensionpoolNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony providers edges extensionpool not found response has a 5xx status code
func (o *GetTelephonyProvidersEdgesExtensionpoolNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edges extensionpool not found response a status code equal to that given
func (o *GetTelephonyProvidersEdgesExtensionpoolNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetTelephonyProvidersEdgesExtensionpoolNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/extensionpools/{extensionPoolId}][%d] getTelephonyProvidersEdgesExtensionpoolNotFound  %+v", 404, o.Payload)
}

func (o *GetTelephonyProvidersEdgesExtensionpoolNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/extensionpools/{extensionPoolId}][%d] getTelephonyProvidersEdgesExtensionpoolNotFound  %+v", 404, o.Payload)
}

func (o *GetTelephonyProvidersEdgesExtensionpoolNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesExtensionpoolNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesExtensionpoolRequestTimeout creates a GetTelephonyProvidersEdgesExtensionpoolRequestTimeout with default headers values
func NewGetTelephonyProvidersEdgesExtensionpoolRequestTimeout() *GetTelephonyProvidersEdgesExtensionpoolRequestTimeout {
	return &GetTelephonyProvidersEdgesExtensionpoolRequestTimeout{}
}

/*
GetTelephonyProvidersEdgesExtensionpoolRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetTelephonyProvidersEdgesExtensionpoolRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edges extensionpool request timeout response has a 2xx status code
func (o *GetTelephonyProvidersEdgesExtensionpoolRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edges extensionpool request timeout response has a 3xx status code
func (o *GetTelephonyProvidersEdgesExtensionpoolRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges extensionpool request timeout response has a 4xx status code
func (o *GetTelephonyProvidersEdgesExtensionpoolRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony providers edges extensionpool request timeout response has a 5xx status code
func (o *GetTelephonyProvidersEdgesExtensionpoolRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edges extensionpool request timeout response a status code equal to that given
func (o *GetTelephonyProvidersEdgesExtensionpoolRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetTelephonyProvidersEdgesExtensionpoolRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/extensionpools/{extensionPoolId}][%d] getTelephonyProvidersEdgesExtensionpoolRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetTelephonyProvidersEdgesExtensionpoolRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/extensionpools/{extensionPoolId}][%d] getTelephonyProvidersEdgesExtensionpoolRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetTelephonyProvidersEdgesExtensionpoolRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesExtensionpoolRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesExtensionpoolRequestEntityTooLarge creates a GetTelephonyProvidersEdgesExtensionpoolRequestEntityTooLarge with default headers values
func NewGetTelephonyProvidersEdgesExtensionpoolRequestEntityTooLarge() *GetTelephonyProvidersEdgesExtensionpoolRequestEntityTooLarge {
	return &GetTelephonyProvidersEdgesExtensionpoolRequestEntityTooLarge{}
}

/*
GetTelephonyProvidersEdgesExtensionpoolRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetTelephonyProvidersEdgesExtensionpoolRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edges extensionpool request entity too large response has a 2xx status code
func (o *GetTelephonyProvidersEdgesExtensionpoolRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edges extensionpool request entity too large response has a 3xx status code
func (o *GetTelephonyProvidersEdgesExtensionpoolRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges extensionpool request entity too large response has a 4xx status code
func (o *GetTelephonyProvidersEdgesExtensionpoolRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony providers edges extensionpool request entity too large response has a 5xx status code
func (o *GetTelephonyProvidersEdgesExtensionpoolRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edges extensionpool request entity too large response a status code equal to that given
func (o *GetTelephonyProvidersEdgesExtensionpoolRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetTelephonyProvidersEdgesExtensionpoolRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/extensionpools/{extensionPoolId}][%d] getTelephonyProvidersEdgesExtensionpoolRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetTelephonyProvidersEdgesExtensionpoolRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/extensionpools/{extensionPoolId}][%d] getTelephonyProvidersEdgesExtensionpoolRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetTelephonyProvidersEdgesExtensionpoolRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesExtensionpoolRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesExtensionpoolUnsupportedMediaType creates a GetTelephonyProvidersEdgesExtensionpoolUnsupportedMediaType with default headers values
func NewGetTelephonyProvidersEdgesExtensionpoolUnsupportedMediaType() *GetTelephonyProvidersEdgesExtensionpoolUnsupportedMediaType {
	return &GetTelephonyProvidersEdgesExtensionpoolUnsupportedMediaType{}
}

/*
GetTelephonyProvidersEdgesExtensionpoolUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetTelephonyProvidersEdgesExtensionpoolUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edges extensionpool unsupported media type response has a 2xx status code
func (o *GetTelephonyProvidersEdgesExtensionpoolUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edges extensionpool unsupported media type response has a 3xx status code
func (o *GetTelephonyProvidersEdgesExtensionpoolUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges extensionpool unsupported media type response has a 4xx status code
func (o *GetTelephonyProvidersEdgesExtensionpoolUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony providers edges extensionpool unsupported media type response has a 5xx status code
func (o *GetTelephonyProvidersEdgesExtensionpoolUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edges extensionpool unsupported media type response a status code equal to that given
func (o *GetTelephonyProvidersEdgesExtensionpoolUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetTelephonyProvidersEdgesExtensionpoolUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/extensionpools/{extensionPoolId}][%d] getTelephonyProvidersEdgesExtensionpoolUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetTelephonyProvidersEdgesExtensionpoolUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/extensionpools/{extensionPoolId}][%d] getTelephonyProvidersEdgesExtensionpoolUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetTelephonyProvidersEdgesExtensionpoolUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesExtensionpoolUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesExtensionpoolTooManyRequests creates a GetTelephonyProvidersEdgesExtensionpoolTooManyRequests with default headers values
func NewGetTelephonyProvidersEdgesExtensionpoolTooManyRequests() *GetTelephonyProvidersEdgesExtensionpoolTooManyRequests {
	return &GetTelephonyProvidersEdgesExtensionpoolTooManyRequests{}
}

/*
GetTelephonyProvidersEdgesExtensionpoolTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetTelephonyProvidersEdgesExtensionpoolTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edges extensionpool too many requests response has a 2xx status code
func (o *GetTelephonyProvidersEdgesExtensionpoolTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edges extensionpool too many requests response has a 3xx status code
func (o *GetTelephonyProvidersEdgesExtensionpoolTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges extensionpool too many requests response has a 4xx status code
func (o *GetTelephonyProvidersEdgesExtensionpoolTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony providers edges extensionpool too many requests response has a 5xx status code
func (o *GetTelephonyProvidersEdgesExtensionpoolTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edges extensionpool too many requests response a status code equal to that given
func (o *GetTelephonyProvidersEdgesExtensionpoolTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetTelephonyProvidersEdgesExtensionpoolTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/extensionpools/{extensionPoolId}][%d] getTelephonyProvidersEdgesExtensionpoolTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetTelephonyProvidersEdgesExtensionpoolTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/extensionpools/{extensionPoolId}][%d] getTelephonyProvidersEdgesExtensionpoolTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetTelephonyProvidersEdgesExtensionpoolTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesExtensionpoolTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesExtensionpoolInternalServerError creates a GetTelephonyProvidersEdgesExtensionpoolInternalServerError with default headers values
func NewGetTelephonyProvidersEdgesExtensionpoolInternalServerError() *GetTelephonyProvidersEdgesExtensionpoolInternalServerError {
	return &GetTelephonyProvidersEdgesExtensionpoolInternalServerError{}
}

/*
GetTelephonyProvidersEdgesExtensionpoolInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetTelephonyProvidersEdgesExtensionpoolInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edges extensionpool internal server error response has a 2xx status code
func (o *GetTelephonyProvidersEdgesExtensionpoolInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edges extensionpool internal server error response has a 3xx status code
func (o *GetTelephonyProvidersEdgesExtensionpoolInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges extensionpool internal server error response has a 4xx status code
func (o *GetTelephonyProvidersEdgesExtensionpoolInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get telephony providers edges extensionpool internal server error response has a 5xx status code
func (o *GetTelephonyProvidersEdgesExtensionpoolInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get telephony providers edges extensionpool internal server error response a status code equal to that given
func (o *GetTelephonyProvidersEdgesExtensionpoolInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetTelephonyProvidersEdgesExtensionpoolInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/extensionpools/{extensionPoolId}][%d] getTelephonyProvidersEdgesExtensionpoolInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTelephonyProvidersEdgesExtensionpoolInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/extensionpools/{extensionPoolId}][%d] getTelephonyProvidersEdgesExtensionpoolInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTelephonyProvidersEdgesExtensionpoolInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesExtensionpoolInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesExtensionpoolServiceUnavailable creates a GetTelephonyProvidersEdgesExtensionpoolServiceUnavailable with default headers values
func NewGetTelephonyProvidersEdgesExtensionpoolServiceUnavailable() *GetTelephonyProvidersEdgesExtensionpoolServiceUnavailable {
	return &GetTelephonyProvidersEdgesExtensionpoolServiceUnavailable{}
}

/*
GetTelephonyProvidersEdgesExtensionpoolServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetTelephonyProvidersEdgesExtensionpoolServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edges extensionpool service unavailable response has a 2xx status code
func (o *GetTelephonyProvidersEdgesExtensionpoolServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edges extensionpool service unavailable response has a 3xx status code
func (o *GetTelephonyProvidersEdgesExtensionpoolServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges extensionpool service unavailable response has a 4xx status code
func (o *GetTelephonyProvidersEdgesExtensionpoolServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get telephony providers edges extensionpool service unavailable response has a 5xx status code
func (o *GetTelephonyProvidersEdgesExtensionpoolServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get telephony providers edges extensionpool service unavailable response a status code equal to that given
func (o *GetTelephonyProvidersEdgesExtensionpoolServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetTelephonyProvidersEdgesExtensionpoolServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/extensionpools/{extensionPoolId}][%d] getTelephonyProvidersEdgesExtensionpoolServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetTelephonyProvidersEdgesExtensionpoolServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/extensionpools/{extensionPoolId}][%d] getTelephonyProvidersEdgesExtensionpoolServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetTelephonyProvidersEdgesExtensionpoolServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesExtensionpoolServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesExtensionpoolGatewayTimeout creates a GetTelephonyProvidersEdgesExtensionpoolGatewayTimeout with default headers values
func NewGetTelephonyProvidersEdgesExtensionpoolGatewayTimeout() *GetTelephonyProvidersEdgesExtensionpoolGatewayTimeout {
	return &GetTelephonyProvidersEdgesExtensionpoolGatewayTimeout{}
}

/*
GetTelephonyProvidersEdgesExtensionpoolGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetTelephonyProvidersEdgesExtensionpoolGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edges extensionpool gateway timeout response has a 2xx status code
func (o *GetTelephonyProvidersEdgesExtensionpoolGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edges extensionpool gateway timeout response has a 3xx status code
func (o *GetTelephonyProvidersEdgesExtensionpoolGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges extensionpool gateway timeout response has a 4xx status code
func (o *GetTelephonyProvidersEdgesExtensionpoolGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get telephony providers edges extensionpool gateway timeout response has a 5xx status code
func (o *GetTelephonyProvidersEdgesExtensionpoolGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get telephony providers edges extensionpool gateway timeout response a status code equal to that given
func (o *GetTelephonyProvidersEdgesExtensionpoolGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetTelephonyProvidersEdgesExtensionpoolGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/extensionpools/{extensionPoolId}][%d] getTelephonyProvidersEdgesExtensionpoolGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetTelephonyProvidersEdgesExtensionpoolGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/extensionpools/{extensionPoolId}][%d] getTelephonyProvidersEdgesExtensionpoolGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetTelephonyProvidersEdgesExtensionpoolGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesExtensionpoolGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

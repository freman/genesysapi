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

// GetTelephonyProvidersEdgesSitesReader is a Reader for the GetTelephonyProvidersEdgesSites structure.
type GetTelephonyProvidersEdgesSitesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTelephonyProvidersEdgesSitesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTelephonyProvidersEdgesSitesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetTelephonyProvidersEdgesSitesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetTelephonyProvidersEdgesSitesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetTelephonyProvidersEdgesSitesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTelephonyProvidersEdgesSitesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetTelephonyProvidersEdgesSitesRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetTelephonyProvidersEdgesSitesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetTelephonyProvidersEdgesSitesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetTelephonyProvidersEdgesSitesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetTelephonyProvidersEdgesSitesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetTelephonyProvidersEdgesSitesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetTelephonyProvidersEdgesSitesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTelephonyProvidersEdgesSitesOK creates a GetTelephonyProvidersEdgesSitesOK with default headers values
func NewGetTelephonyProvidersEdgesSitesOK() *GetTelephonyProvidersEdgesSitesOK {
	return &GetTelephonyProvidersEdgesSitesOK{}
}

/*
GetTelephonyProvidersEdgesSitesOK describes a response with status code 200, with default header values.

successful operation
*/
type GetTelephonyProvidersEdgesSitesOK struct {
	Payload *models.SiteEntityListing
}

// IsSuccess returns true when this get telephony providers edges sites o k response has a 2xx status code
func (o *GetTelephonyProvidersEdgesSitesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get telephony providers edges sites o k response has a 3xx status code
func (o *GetTelephonyProvidersEdgesSitesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges sites o k response has a 4xx status code
func (o *GetTelephonyProvidersEdgesSitesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get telephony providers edges sites o k response has a 5xx status code
func (o *GetTelephonyProvidersEdgesSitesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edges sites o k response a status code equal to that given
func (o *GetTelephonyProvidersEdgesSitesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetTelephonyProvidersEdgesSitesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/sites][%d] getTelephonyProvidersEdgesSitesOK  %+v", 200, o.Payload)
}

func (o *GetTelephonyProvidersEdgesSitesOK) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/sites][%d] getTelephonyProvidersEdgesSitesOK  %+v", 200, o.Payload)
}

func (o *GetTelephonyProvidersEdgesSitesOK) GetPayload() *models.SiteEntityListing {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesSitesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SiteEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesSitesBadRequest creates a GetTelephonyProvidersEdgesSitesBadRequest with default headers values
func NewGetTelephonyProvidersEdgesSitesBadRequest() *GetTelephonyProvidersEdgesSitesBadRequest {
	return &GetTelephonyProvidersEdgesSitesBadRequest{}
}

/*
GetTelephonyProvidersEdgesSitesBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetTelephonyProvidersEdgesSitesBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edges sites bad request response has a 2xx status code
func (o *GetTelephonyProvidersEdgesSitesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edges sites bad request response has a 3xx status code
func (o *GetTelephonyProvidersEdgesSitesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges sites bad request response has a 4xx status code
func (o *GetTelephonyProvidersEdgesSitesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony providers edges sites bad request response has a 5xx status code
func (o *GetTelephonyProvidersEdgesSitesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edges sites bad request response a status code equal to that given
func (o *GetTelephonyProvidersEdgesSitesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetTelephonyProvidersEdgesSitesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/sites][%d] getTelephonyProvidersEdgesSitesBadRequest  %+v", 400, o.Payload)
}

func (o *GetTelephonyProvidersEdgesSitesBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/sites][%d] getTelephonyProvidersEdgesSitesBadRequest  %+v", 400, o.Payload)
}

func (o *GetTelephonyProvidersEdgesSitesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesSitesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesSitesUnauthorized creates a GetTelephonyProvidersEdgesSitesUnauthorized with default headers values
func NewGetTelephonyProvidersEdgesSitesUnauthorized() *GetTelephonyProvidersEdgesSitesUnauthorized {
	return &GetTelephonyProvidersEdgesSitesUnauthorized{}
}

/*
GetTelephonyProvidersEdgesSitesUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetTelephonyProvidersEdgesSitesUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edges sites unauthorized response has a 2xx status code
func (o *GetTelephonyProvidersEdgesSitesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edges sites unauthorized response has a 3xx status code
func (o *GetTelephonyProvidersEdgesSitesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges sites unauthorized response has a 4xx status code
func (o *GetTelephonyProvidersEdgesSitesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony providers edges sites unauthorized response has a 5xx status code
func (o *GetTelephonyProvidersEdgesSitesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edges sites unauthorized response a status code equal to that given
func (o *GetTelephonyProvidersEdgesSitesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetTelephonyProvidersEdgesSitesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/sites][%d] getTelephonyProvidersEdgesSitesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetTelephonyProvidersEdgesSitesUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/sites][%d] getTelephonyProvidersEdgesSitesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetTelephonyProvidersEdgesSitesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesSitesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesSitesForbidden creates a GetTelephonyProvidersEdgesSitesForbidden with default headers values
func NewGetTelephonyProvidersEdgesSitesForbidden() *GetTelephonyProvidersEdgesSitesForbidden {
	return &GetTelephonyProvidersEdgesSitesForbidden{}
}

/*
GetTelephonyProvidersEdgesSitesForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetTelephonyProvidersEdgesSitesForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edges sites forbidden response has a 2xx status code
func (o *GetTelephonyProvidersEdgesSitesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edges sites forbidden response has a 3xx status code
func (o *GetTelephonyProvidersEdgesSitesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges sites forbidden response has a 4xx status code
func (o *GetTelephonyProvidersEdgesSitesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony providers edges sites forbidden response has a 5xx status code
func (o *GetTelephonyProvidersEdgesSitesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edges sites forbidden response a status code equal to that given
func (o *GetTelephonyProvidersEdgesSitesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetTelephonyProvidersEdgesSitesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/sites][%d] getTelephonyProvidersEdgesSitesForbidden  %+v", 403, o.Payload)
}

func (o *GetTelephonyProvidersEdgesSitesForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/sites][%d] getTelephonyProvidersEdgesSitesForbidden  %+v", 403, o.Payload)
}

func (o *GetTelephonyProvidersEdgesSitesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesSitesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesSitesNotFound creates a GetTelephonyProvidersEdgesSitesNotFound with default headers values
func NewGetTelephonyProvidersEdgesSitesNotFound() *GetTelephonyProvidersEdgesSitesNotFound {
	return &GetTelephonyProvidersEdgesSitesNotFound{}
}

/*
GetTelephonyProvidersEdgesSitesNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetTelephonyProvidersEdgesSitesNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edges sites not found response has a 2xx status code
func (o *GetTelephonyProvidersEdgesSitesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edges sites not found response has a 3xx status code
func (o *GetTelephonyProvidersEdgesSitesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges sites not found response has a 4xx status code
func (o *GetTelephonyProvidersEdgesSitesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony providers edges sites not found response has a 5xx status code
func (o *GetTelephonyProvidersEdgesSitesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edges sites not found response a status code equal to that given
func (o *GetTelephonyProvidersEdgesSitesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetTelephonyProvidersEdgesSitesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/sites][%d] getTelephonyProvidersEdgesSitesNotFound  %+v", 404, o.Payload)
}

func (o *GetTelephonyProvidersEdgesSitesNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/sites][%d] getTelephonyProvidersEdgesSitesNotFound  %+v", 404, o.Payload)
}

func (o *GetTelephonyProvidersEdgesSitesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesSitesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesSitesRequestTimeout creates a GetTelephonyProvidersEdgesSitesRequestTimeout with default headers values
func NewGetTelephonyProvidersEdgesSitesRequestTimeout() *GetTelephonyProvidersEdgesSitesRequestTimeout {
	return &GetTelephonyProvidersEdgesSitesRequestTimeout{}
}

/*
GetTelephonyProvidersEdgesSitesRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetTelephonyProvidersEdgesSitesRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edges sites request timeout response has a 2xx status code
func (o *GetTelephonyProvidersEdgesSitesRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edges sites request timeout response has a 3xx status code
func (o *GetTelephonyProvidersEdgesSitesRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges sites request timeout response has a 4xx status code
func (o *GetTelephonyProvidersEdgesSitesRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony providers edges sites request timeout response has a 5xx status code
func (o *GetTelephonyProvidersEdgesSitesRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edges sites request timeout response a status code equal to that given
func (o *GetTelephonyProvidersEdgesSitesRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetTelephonyProvidersEdgesSitesRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/sites][%d] getTelephonyProvidersEdgesSitesRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetTelephonyProvidersEdgesSitesRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/sites][%d] getTelephonyProvidersEdgesSitesRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetTelephonyProvidersEdgesSitesRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesSitesRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesSitesRequestEntityTooLarge creates a GetTelephonyProvidersEdgesSitesRequestEntityTooLarge with default headers values
func NewGetTelephonyProvidersEdgesSitesRequestEntityTooLarge() *GetTelephonyProvidersEdgesSitesRequestEntityTooLarge {
	return &GetTelephonyProvidersEdgesSitesRequestEntityTooLarge{}
}

/*
GetTelephonyProvidersEdgesSitesRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetTelephonyProvidersEdgesSitesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edges sites request entity too large response has a 2xx status code
func (o *GetTelephonyProvidersEdgesSitesRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edges sites request entity too large response has a 3xx status code
func (o *GetTelephonyProvidersEdgesSitesRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges sites request entity too large response has a 4xx status code
func (o *GetTelephonyProvidersEdgesSitesRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony providers edges sites request entity too large response has a 5xx status code
func (o *GetTelephonyProvidersEdgesSitesRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edges sites request entity too large response a status code equal to that given
func (o *GetTelephonyProvidersEdgesSitesRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetTelephonyProvidersEdgesSitesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/sites][%d] getTelephonyProvidersEdgesSitesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetTelephonyProvidersEdgesSitesRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/sites][%d] getTelephonyProvidersEdgesSitesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetTelephonyProvidersEdgesSitesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesSitesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesSitesUnsupportedMediaType creates a GetTelephonyProvidersEdgesSitesUnsupportedMediaType with default headers values
func NewGetTelephonyProvidersEdgesSitesUnsupportedMediaType() *GetTelephonyProvidersEdgesSitesUnsupportedMediaType {
	return &GetTelephonyProvidersEdgesSitesUnsupportedMediaType{}
}

/*
GetTelephonyProvidersEdgesSitesUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetTelephonyProvidersEdgesSitesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edges sites unsupported media type response has a 2xx status code
func (o *GetTelephonyProvidersEdgesSitesUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edges sites unsupported media type response has a 3xx status code
func (o *GetTelephonyProvidersEdgesSitesUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges sites unsupported media type response has a 4xx status code
func (o *GetTelephonyProvidersEdgesSitesUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony providers edges sites unsupported media type response has a 5xx status code
func (o *GetTelephonyProvidersEdgesSitesUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edges sites unsupported media type response a status code equal to that given
func (o *GetTelephonyProvidersEdgesSitesUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetTelephonyProvidersEdgesSitesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/sites][%d] getTelephonyProvidersEdgesSitesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetTelephonyProvidersEdgesSitesUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/sites][%d] getTelephonyProvidersEdgesSitesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetTelephonyProvidersEdgesSitesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesSitesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesSitesTooManyRequests creates a GetTelephonyProvidersEdgesSitesTooManyRequests with default headers values
func NewGetTelephonyProvidersEdgesSitesTooManyRequests() *GetTelephonyProvidersEdgesSitesTooManyRequests {
	return &GetTelephonyProvidersEdgesSitesTooManyRequests{}
}

/*
GetTelephonyProvidersEdgesSitesTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetTelephonyProvidersEdgesSitesTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edges sites too many requests response has a 2xx status code
func (o *GetTelephonyProvidersEdgesSitesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edges sites too many requests response has a 3xx status code
func (o *GetTelephonyProvidersEdgesSitesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges sites too many requests response has a 4xx status code
func (o *GetTelephonyProvidersEdgesSitesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony providers edges sites too many requests response has a 5xx status code
func (o *GetTelephonyProvidersEdgesSitesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edges sites too many requests response a status code equal to that given
func (o *GetTelephonyProvidersEdgesSitesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetTelephonyProvidersEdgesSitesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/sites][%d] getTelephonyProvidersEdgesSitesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetTelephonyProvidersEdgesSitesTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/sites][%d] getTelephonyProvidersEdgesSitesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetTelephonyProvidersEdgesSitesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesSitesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesSitesInternalServerError creates a GetTelephonyProvidersEdgesSitesInternalServerError with default headers values
func NewGetTelephonyProvidersEdgesSitesInternalServerError() *GetTelephonyProvidersEdgesSitesInternalServerError {
	return &GetTelephonyProvidersEdgesSitesInternalServerError{}
}

/*
GetTelephonyProvidersEdgesSitesInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetTelephonyProvidersEdgesSitesInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edges sites internal server error response has a 2xx status code
func (o *GetTelephonyProvidersEdgesSitesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edges sites internal server error response has a 3xx status code
func (o *GetTelephonyProvidersEdgesSitesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges sites internal server error response has a 4xx status code
func (o *GetTelephonyProvidersEdgesSitesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get telephony providers edges sites internal server error response has a 5xx status code
func (o *GetTelephonyProvidersEdgesSitesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get telephony providers edges sites internal server error response a status code equal to that given
func (o *GetTelephonyProvidersEdgesSitesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetTelephonyProvidersEdgesSitesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/sites][%d] getTelephonyProvidersEdgesSitesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTelephonyProvidersEdgesSitesInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/sites][%d] getTelephonyProvidersEdgesSitesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTelephonyProvidersEdgesSitesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesSitesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesSitesServiceUnavailable creates a GetTelephonyProvidersEdgesSitesServiceUnavailable with default headers values
func NewGetTelephonyProvidersEdgesSitesServiceUnavailable() *GetTelephonyProvidersEdgesSitesServiceUnavailable {
	return &GetTelephonyProvidersEdgesSitesServiceUnavailable{}
}

/*
GetTelephonyProvidersEdgesSitesServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetTelephonyProvidersEdgesSitesServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edges sites service unavailable response has a 2xx status code
func (o *GetTelephonyProvidersEdgesSitesServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edges sites service unavailable response has a 3xx status code
func (o *GetTelephonyProvidersEdgesSitesServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges sites service unavailable response has a 4xx status code
func (o *GetTelephonyProvidersEdgesSitesServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get telephony providers edges sites service unavailable response has a 5xx status code
func (o *GetTelephonyProvidersEdgesSitesServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get telephony providers edges sites service unavailable response a status code equal to that given
func (o *GetTelephonyProvidersEdgesSitesServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetTelephonyProvidersEdgesSitesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/sites][%d] getTelephonyProvidersEdgesSitesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetTelephonyProvidersEdgesSitesServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/sites][%d] getTelephonyProvidersEdgesSitesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetTelephonyProvidersEdgesSitesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesSitesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesSitesGatewayTimeout creates a GetTelephonyProvidersEdgesSitesGatewayTimeout with default headers values
func NewGetTelephonyProvidersEdgesSitesGatewayTimeout() *GetTelephonyProvidersEdgesSitesGatewayTimeout {
	return &GetTelephonyProvidersEdgesSitesGatewayTimeout{}
}

/*
GetTelephonyProvidersEdgesSitesGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetTelephonyProvidersEdgesSitesGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edges sites gateway timeout response has a 2xx status code
func (o *GetTelephonyProvidersEdgesSitesGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edges sites gateway timeout response has a 3xx status code
func (o *GetTelephonyProvidersEdgesSitesGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges sites gateway timeout response has a 4xx status code
func (o *GetTelephonyProvidersEdgesSitesGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get telephony providers edges sites gateway timeout response has a 5xx status code
func (o *GetTelephonyProvidersEdgesSitesGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get telephony providers edges sites gateway timeout response a status code equal to that given
func (o *GetTelephonyProvidersEdgesSitesGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetTelephonyProvidersEdgesSitesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/sites][%d] getTelephonyProvidersEdgesSitesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetTelephonyProvidersEdgesSitesGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/sites][%d] getTelephonyProvidersEdgesSitesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetTelephonyProvidersEdgesSitesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesSitesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package telephony

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetTelephonyMediaregionsReader is a Reader for the GetTelephonyMediaregions structure.
type GetTelephonyMediaregionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTelephonyMediaregionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTelephonyMediaregionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetTelephonyMediaregionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetTelephonyMediaregionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetTelephonyMediaregionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTelephonyMediaregionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetTelephonyMediaregionsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetTelephonyMediaregionsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetTelephonyMediaregionsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetTelephonyMediaregionsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetTelephonyMediaregionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetTelephonyMediaregionsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetTelephonyMediaregionsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTelephonyMediaregionsOK creates a GetTelephonyMediaregionsOK with default headers values
func NewGetTelephonyMediaregionsOK() *GetTelephonyMediaregionsOK {
	return &GetTelephonyMediaregionsOK{}
}

/*
GetTelephonyMediaregionsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetTelephonyMediaregionsOK struct {
	Payload *models.MediaRegions
}

// IsSuccess returns true when this get telephony mediaregions o k response has a 2xx status code
func (o *GetTelephonyMediaregionsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get telephony mediaregions o k response has a 3xx status code
func (o *GetTelephonyMediaregionsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony mediaregions o k response has a 4xx status code
func (o *GetTelephonyMediaregionsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get telephony mediaregions o k response has a 5xx status code
func (o *GetTelephonyMediaregionsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony mediaregions o k response a status code equal to that given
func (o *GetTelephonyMediaregionsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetTelephonyMediaregionsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/mediaregions][%d] getTelephonyMediaregionsOK  %+v", 200, o.Payload)
}

func (o *GetTelephonyMediaregionsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/mediaregions][%d] getTelephonyMediaregionsOK  %+v", 200, o.Payload)
}

func (o *GetTelephonyMediaregionsOK) GetPayload() *models.MediaRegions {
	return o.Payload
}

func (o *GetTelephonyMediaregionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MediaRegions)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyMediaregionsBadRequest creates a GetTelephonyMediaregionsBadRequest with default headers values
func NewGetTelephonyMediaregionsBadRequest() *GetTelephonyMediaregionsBadRequest {
	return &GetTelephonyMediaregionsBadRequest{}
}

/*
GetTelephonyMediaregionsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetTelephonyMediaregionsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony mediaregions bad request response has a 2xx status code
func (o *GetTelephonyMediaregionsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony mediaregions bad request response has a 3xx status code
func (o *GetTelephonyMediaregionsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony mediaregions bad request response has a 4xx status code
func (o *GetTelephonyMediaregionsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony mediaregions bad request response has a 5xx status code
func (o *GetTelephonyMediaregionsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony mediaregions bad request response a status code equal to that given
func (o *GetTelephonyMediaregionsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetTelephonyMediaregionsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/mediaregions][%d] getTelephonyMediaregionsBadRequest  %+v", 400, o.Payload)
}

func (o *GetTelephonyMediaregionsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/mediaregions][%d] getTelephonyMediaregionsBadRequest  %+v", 400, o.Payload)
}

func (o *GetTelephonyMediaregionsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyMediaregionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyMediaregionsUnauthorized creates a GetTelephonyMediaregionsUnauthorized with default headers values
func NewGetTelephonyMediaregionsUnauthorized() *GetTelephonyMediaregionsUnauthorized {
	return &GetTelephonyMediaregionsUnauthorized{}
}

/*
GetTelephonyMediaregionsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetTelephonyMediaregionsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony mediaregions unauthorized response has a 2xx status code
func (o *GetTelephonyMediaregionsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony mediaregions unauthorized response has a 3xx status code
func (o *GetTelephonyMediaregionsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony mediaregions unauthorized response has a 4xx status code
func (o *GetTelephonyMediaregionsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony mediaregions unauthorized response has a 5xx status code
func (o *GetTelephonyMediaregionsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony mediaregions unauthorized response a status code equal to that given
func (o *GetTelephonyMediaregionsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetTelephonyMediaregionsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/mediaregions][%d] getTelephonyMediaregionsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetTelephonyMediaregionsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/mediaregions][%d] getTelephonyMediaregionsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetTelephonyMediaregionsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyMediaregionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyMediaregionsForbidden creates a GetTelephonyMediaregionsForbidden with default headers values
func NewGetTelephonyMediaregionsForbidden() *GetTelephonyMediaregionsForbidden {
	return &GetTelephonyMediaregionsForbidden{}
}

/*
GetTelephonyMediaregionsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetTelephonyMediaregionsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony mediaregions forbidden response has a 2xx status code
func (o *GetTelephonyMediaregionsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony mediaregions forbidden response has a 3xx status code
func (o *GetTelephonyMediaregionsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony mediaregions forbidden response has a 4xx status code
func (o *GetTelephonyMediaregionsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony mediaregions forbidden response has a 5xx status code
func (o *GetTelephonyMediaregionsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony mediaregions forbidden response a status code equal to that given
func (o *GetTelephonyMediaregionsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetTelephonyMediaregionsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/mediaregions][%d] getTelephonyMediaregionsForbidden  %+v", 403, o.Payload)
}

func (o *GetTelephonyMediaregionsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/mediaregions][%d] getTelephonyMediaregionsForbidden  %+v", 403, o.Payload)
}

func (o *GetTelephonyMediaregionsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyMediaregionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyMediaregionsNotFound creates a GetTelephonyMediaregionsNotFound with default headers values
func NewGetTelephonyMediaregionsNotFound() *GetTelephonyMediaregionsNotFound {
	return &GetTelephonyMediaregionsNotFound{}
}

/*
GetTelephonyMediaregionsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetTelephonyMediaregionsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony mediaregions not found response has a 2xx status code
func (o *GetTelephonyMediaregionsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony mediaregions not found response has a 3xx status code
func (o *GetTelephonyMediaregionsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony mediaregions not found response has a 4xx status code
func (o *GetTelephonyMediaregionsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony mediaregions not found response has a 5xx status code
func (o *GetTelephonyMediaregionsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony mediaregions not found response a status code equal to that given
func (o *GetTelephonyMediaregionsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetTelephonyMediaregionsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/mediaregions][%d] getTelephonyMediaregionsNotFound  %+v", 404, o.Payload)
}

func (o *GetTelephonyMediaregionsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/mediaregions][%d] getTelephonyMediaregionsNotFound  %+v", 404, o.Payload)
}

func (o *GetTelephonyMediaregionsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyMediaregionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyMediaregionsRequestTimeout creates a GetTelephonyMediaregionsRequestTimeout with default headers values
func NewGetTelephonyMediaregionsRequestTimeout() *GetTelephonyMediaregionsRequestTimeout {
	return &GetTelephonyMediaregionsRequestTimeout{}
}

/*
GetTelephonyMediaregionsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetTelephonyMediaregionsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony mediaregions request timeout response has a 2xx status code
func (o *GetTelephonyMediaregionsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony mediaregions request timeout response has a 3xx status code
func (o *GetTelephonyMediaregionsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony mediaregions request timeout response has a 4xx status code
func (o *GetTelephonyMediaregionsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony mediaregions request timeout response has a 5xx status code
func (o *GetTelephonyMediaregionsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony mediaregions request timeout response a status code equal to that given
func (o *GetTelephonyMediaregionsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetTelephonyMediaregionsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/mediaregions][%d] getTelephonyMediaregionsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetTelephonyMediaregionsRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/mediaregions][%d] getTelephonyMediaregionsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetTelephonyMediaregionsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyMediaregionsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyMediaregionsRequestEntityTooLarge creates a GetTelephonyMediaregionsRequestEntityTooLarge with default headers values
func NewGetTelephonyMediaregionsRequestEntityTooLarge() *GetTelephonyMediaregionsRequestEntityTooLarge {
	return &GetTelephonyMediaregionsRequestEntityTooLarge{}
}

/*
GetTelephonyMediaregionsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetTelephonyMediaregionsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony mediaregions request entity too large response has a 2xx status code
func (o *GetTelephonyMediaregionsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony mediaregions request entity too large response has a 3xx status code
func (o *GetTelephonyMediaregionsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony mediaregions request entity too large response has a 4xx status code
func (o *GetTelephonyMediaregionsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony mediaregions request entity too large response has a 5xx status code
func (o *GetTelephonyMediaregionsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony mediaregions request entity too large response a status code equal to that given
func (o *GetTelephonyMediaregionsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetTelephonyMediaregionsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/mediaregions][%d] getTelephonyMediaregionsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetTelephonyMediaregionsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/mediaregions][%d] getTelephonyMediaregionsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetTelephonyMediaregionsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyMediaregionsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyMediaregionsUnsupportedMediaType creates a GetTelephonyMediaregionsUnsupportedMediaType with default headers values
func NewGetTelephonyMediaregionsUnsupportedMediaType() *GetTelephonyMediaregionsUnsupportedMediaType {
	return &GetTelephonyMediaregionsUnsupportedMediaType{}
}

/*
GetTelephonyMediaregionsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetTelephonyMediaregionsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony mediaregions unsupported media type response has a 2xx status code
func (o *GetTelephonyMediaregionsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony mediaregions unsupported media type response has a 3xx status code
func (o *GetTelephonyMediaregionsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony mediaregions unsupported media type response has a 4xx status code
func (o *GetTelephonyMediaregionsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony mediaregions unsupported media type response has a 5xx status code
func (o *GetTelephonyMediaregionsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony mediaregions unsupported media type response a status code equal to that given
func (o *GetTelephonyMediaregionsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetTelephonyMediaregionsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/mediaregions][%d] getTelephonyMediaregionsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetTelephonyMediaregionsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/mediaregions][%d] getTelephonyMediaregionsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetTelephonyMediaregionsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyMediaregionsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyMediaregionsTooManyRequests creates a GetTelephonyMediaregionsTooManyRequests with default headers values
func NewGetTelephonyMediaregionsTooManyRequests() *GetTelephonyMediaregionsTooManyRequests {
	return &GetTelephonyMediaregionsTooManyRequests{}
}

/*
GetTelephonyMediaregionsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetTelephonyMediaregionsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony mediaregions too many requests response has a 2xx status code
func (o *GetTelephonyMediaregionsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony mediaregions too many requests response has a 3xx status code
func (o *GetTelephonyMediaregionsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony mediaregions too many requests response has a 4xx status code
func (o *GetTelephonyMediaregionsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony mediaregions too many requests response has a 5xx status code
func (o *GetTelephonyMediaregionsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony mediaregions too many requests response a status code equal to that given
func (o *GetTelephonyMediaregionsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetTelephonyMediaregionsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/mediaregions][%d] getTelephonyMediaregionsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetTelephonyMediaregionsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/mediaregions][%d] getTelephonyMediaregionsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetTelephonyMediaregionsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyMediaregionsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyMediaregionsInternalServerError creates a GetTelephonyMediaregionsInternalServerError with default headers values
func NewGetTelephonyMediaregionsInternalServerError() *GetTelephonyMediaregionsInternalServerError {
	return &GetTelephonyMediaregionsInternalServerError{}
}

/*
GetTelephonyMediaregionsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetTelephonyMediaregionsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony mediaregions internal server error response has a 2xx status code
func (o *GetTelephonyMediaregionsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony mediaregions internal server error response has a 3xx status code
func (o *GetTelephonyMediaregionsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony mediaregions internal server error response has a 4xx status code
func (o *GetTelephonyMediaregionsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get telephony mediaregions internal server error response has a 5xx status code
func (o *GetTelephonyMediaregionsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get telephony mediaregions internal server error response a status code equal to that given
func (o *GetTelephonyMediaregionsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetTelephonyMediaregionsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/mediaregions][%d] getTelephonyMediaregionsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTelephonyMediaregionsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/mediaregions][%d] getTelephonyMediaregionsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTelephonyMediaregionsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyMediaregionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyMediaregionsServiceUnavailable creates a GetTelephonyMediaregionsServiceUnavailable with default headers values
func NewGetTelephonyMediaregionsServiceUnavailable() *GetTelephonyMediaregionsServiceUnavailable {
	return &GetTelephonyMediaregionsServiceUnavailable{}
}

/*
GetTelephonyMediaregionsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetTelephonyMediaregionsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony mediaregions service unavailable response has a 2xx status code
func (o *GetTelephonyMediaregionsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony mediaregions service unavailable response has a 3xx status code
func (o *GetTelephonyMediaregionsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony mediaregions service unavailable response has a 4xx status code
func (o *GetTelephonyMediaregionsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get telephony mediaregions service unavailable response has a 5xx status code
func (o *GetTelephonyMediaregionsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get telephony mediaregions service unavailable response a status code equal to that given
func (o *GetTelephonyMediaregionsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetTelephonyMediaregionsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/mediaregions][%d] getTelephonyMediaregionsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetTelephonyMediaregionsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/mediaregions][%d] getTelephonyMediaregionsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetTelephonyMediaregionsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyMediaregionsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyMediaregionsGatewayTimeout creates a GetTelephonyMediaregionsGatewayTimeout with default headers values
func NewGetTelephonyMediaregionsGatewayTimeout() *GetTelephonyMediaregionsGatewayTimeout {
	return &GetTelephonyMediaregionsGatewayTimeout{}
}

/*
GetTelephonyMediaregionsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetTelephonyMediaregionsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony mediaregions gateway timeout response has a 2xx status code
func (o *GetTelephonyMediaregionsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony mediaregions gateway timeout response has a 3xx status code
func (o *GetTelephonyMediaregionsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony mediaregions gateway timeout response has a 4xx status code
func (o *GetTelephonyMediaregionsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get telephony mediaregions gateway timeout response has a 5xx status code
func (o *GetTelephonyMediaregionsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get telephony mediaregions gateway timeout response a status code equal to that given
func (o *GetTelephonyMediaregionsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetTelephonyMediaregionsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/mediaregions][%d] getTelephonyMediaregionsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetTelephonyMediaregionsGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/mediaregions][%d] getTelephonyMediaregionsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetTelephonyMediaregionsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyMediaregionsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

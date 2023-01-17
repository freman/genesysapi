// Code generated by go-swagger; DO NOT EDIT.

package outbound

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetOutboundAttemptlimitReader is a Reader for the GetOutboundAttemptlimit structure.
type GetOutboundAttemptlimitReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOutboundAttemptlimitReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOutboundAttemptlimitOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOutboundAttemptlimitBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetOutboundAttemptlimitUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOutboundAttemptlimitForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOutboundAttemptlimitNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetOutboundAttemptlimitRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetOutboundAttemptlimitRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetOutboundAttemptlimitUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetOutboundAttemptlimitTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOutboundAttemptlimitInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetOutboundAttemptlimitServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetOutboundAttemptlimitGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOutboundAttemptlimitOK creates a GetOutboundAttemptlimitOK with default headers values
func NewGetOutboundAttemptlimitOK() *GetOutboundAttemptlimitOK {
	return &GetOutboundAttemptlimitOK{}
}

/*
GetOutboundAttemptlimitOK describes a response with status code 200, with default header values.

successful operation
*/
type GetOutboundAttemptlimitOK struct {
	Payload *models.AttemptLimits
}

// IsSuccess returns true when this get outbound attemptlimit o k response has a 2xx status code
func (o *GetOutboundAttemptlimitOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get outbound attemptlimit o k response has a 3xx status code
func (o *GetOutboundAttemptlimitOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound attemptlimit o k response has a 4xx status code
func (o *GetOutboundAttemptlimitOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get outbound attemptlimit o k response has a 5xx status code
func (o *GetOutboundAttemptlimitOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound attemptlimit o k response a status code equal to that given
func (o *GetOutboundAttemptlimitOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetOutboundAttemptlimitOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/attemptlimits/{attemptLimitsId}][%d] getOutboundAttemptlimitOK  %+v", 200, o.Payload)
}

func (o *GetOutboundAttemptlimitOK) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/attemptlimits/{attemptLimitsId}][%d] getOutboundAttemptlimitOK  %+v", 200, o.Payload)
}

func (o *GetOutboundAttemptlimitOK) GetPayload() *models.AttemptLimits {
	return o.Payload
}

func (o *GetOutboundAttemptlimitOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AttemptLimits)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundAttemptlimitBadRequest creates a GetOutboundAttemptlimitBadRequest with default headers values
func NewGetOutboundAttemptlimitBadRequest() *GetOutboundAttemptlimitBadRequest {
	return &GetOutboundAttemptlimitBadRequest{}
}

/*
GetOutboundAttemptlimitBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetOutboundAttemptlimitBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound attemptlimit bad request response has a 2xx status code
func (o *GetOutboundAttemptlimitBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound attemptlimit bad request response has a 3xx status code
func (o *GetOutboundAttemptlimitBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound attemptlimit bad request response has a 4xx status code
func (o *GetOutboundAttemptlimitBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound attemptlimit bad request response has a 5xx status code
func (o *GetOutboundAttemptlimitBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound attemptlimit bad request response a status code equal to that given
func (o *GetOutboundAttemptlimitBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetOutboundAttemptlimitBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/attemptlimits/{attemptLimitsId}][%d] getOutboundAttemptlimitBadRequest  %+v", 400, o.Payload)
}

func (o *GetOutboundAttemptlimitBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/attemptlimits/{attemptLimitsId}][%d] getOutboundAttemptlimitBadRequest  %+v", 400, o.Payload)
}

func (o *GetOutboundAttemptlimitBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundAttemptlimitBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundAttemptlimitUnauthorized creates a GetOutboundAttemptlimitUnauthorized with default headers values
func NewGetOutboundAttemptlimitUnauthorized() *GetOutboundAttemptlimitUnauthorized {
	return &GetOutboundAttemptlimitUnauthorized{}
}

/*
GetOutboundAttemptlimitUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetOutboundAttemptlimitUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound attemptlimit unauthorized response has a 2xx status code
func (o *GetOutboundAttemptlimitUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound attemptlimit unauthorized response has a 3xx status code
func (o *GetOutboundAttemptlimitUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound attemptlimit unauthorized response has a 4xx status code
func (o *GetOutboundAttemptlimitUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound attemptlimit unauthorized response has a 5xx status code
func (o *GetOutboundAttemptlimitUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound attemptlimit unauthorized response a status code equal to that given
func (o *GetOutboundAttemptlimitUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetOutboundAttemptlimitUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/attemptlimits/{attemptLimitsId}][%d] getOutboundAttemptlimitUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOutboundAttemptlimitUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/attemptlimits/{attemptLimitsId}][%d] getOutboundAttemptlimitUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOutboundAttemptlimitUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundAttemptlimitUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundAttemptlimitForbidden creates a GetOutboundAttemptlimitForbidden with default headers values
func NewGetOutboundAttemptlimitForbidden() *GetOutboundAttemptlimitForbidden {
	return &GetOutboundAttemptlimitForbidden{}
}

/*
GetOutboundAttemptlimitForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetOutboundAttemptlimitForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound attemptlimit forbidden response has a 2xx status code
func (o *GetOutboundAttemptlimitForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound attemptlimit forbidden response has a 3xx status code
func (o *GetOutboundAttemptlimitForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound attemptlimit forbidden response has a 4xx status code
func (o *GetOutboundAttemptlimitForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound attemptlimit forbidden response has a 5xx status code
func (o *GetOutboundAttemptlimitForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound attemptlimit forbidden response a status code equal to that given
func (o *GetOutboundAttemptlimitForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetOutboundAttemptlimitForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/attemptlimits/{attemptLimitsId}][%d] getOutboundAttemptlimitForbidden  %+v", 403, o.Payload)
}

func (o *GetOutboundAttemptlimitForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/attemptlimits/{attemptLimitsId}][%d] getOutboundAttemptlimitForbidden  %+v", 403, o.Payload)
}

func (o *GetOutboundAttemptlimitForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundAttemptlimitForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundAttemptlimitNotFound creates a GetOutboundAttemptlimitNotFound with default headers values
func NewGetOutboundAttemptlimitNotFound() *GetOutboundAttemptlimitNotFound {
	return &GetOutboundAttemptlimitNotFound{}
}

/*
GetOutboundAttemptlimitNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetOutboundAttemptlimitNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound attemptlimit not found response has a 2xx status code
func (o *GetOutboundAttemptlimitNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound attemptlimit not found response has a 3xx status code
func (o *GetOutboundAttemptlimitNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound attemptlimit not found response has a 4xx status code
func (o *GetOutboundAttemptlimitNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound attemptlimit not found response has a 5xx status code
func (o *GetOutboundAttemptlimitNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound attemptlimit not found response a status code equal to that given
func (o *GetOutboundAttemptlimitNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetOutboundAttemptlimitNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/attemptlimits/{attemptLimitsId}][%d] getOutboundAttemptlimitNotFound  %+v", 404, o.Payload)
}

func (o *GetOutboundAttemptlimitNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/attemptlimits/{attemptLimitsId}][%d] getOutboundAttemptlimitNotFound  %+v", 404, o.Payload)
}

func (o *GetOutboundAttemptlimitNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundAttemptlimitNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundAttemptlimitRequestTimeout creates a GetOutboundAttemptlimitRequestTimeout with default headers values
func NewGetOutboundAttemptlimitRequestTimeout() *GetOutboundAttemptlimitRequestTimeout {
	return &GetOutboundAttemptlimitRequestTimeout{}
}

/*
GetOutboundAttemptlimitRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetOutboundAttemptlimitRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound attemptlimit request timeout response has a 2xx status code
func (o *GetOutboundAttemptlimitRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound attemptlimit request timeout response has a 3xx status code
func (o *GetOutboundAttemptlimitRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound attemptlimit request timeout response has a 4xx status code
func (o *GetOutboundAttemptlimitRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound attemptlimit request timeout response has a 5xx status code
func (o *GetOutboundAttemptlimitRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound attemptlimit request timeout response a status code equal to that given
func (o *GetOutboundAttemptlimitRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetOutboundAttemptlimitRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/attemptlimits/{attemptLimitsId}][%d] getOutboundAttemptlimitRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetOutboundAttemptlimitRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/attemptlimits/{attemptLimitsId}][%d] getOutboundAttemptlimitRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetOutboundAttemptlimitRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundAttemptlimitRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundAttemptlimitRequestEntityTooLarge creates a GetOutboundAttemptlimitRequestEntityTooLarge with default headers values
func NewGetOutboundAttemptlimitRequestEntityTooLarge() *GetOutboundAttemptlimitRequestEntityTooLarge {
	return &GetOutboundAttemptlimitRequestEntityTooLarge{}
}

/*
GetOutboundAttemptlimitRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetOutboundAttemptlimitRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound attemptlimit request entity too large response has a 2xx status code
func (o *GetOutboundAttemptlimitRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound attemptlimit request entity too large response has a 3xx status code
func (o *GetOutboundAttemptlimitRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound attemptlimit request entity too large response has a 4xx status code
func (o *GetOutboundAttemptlimitRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound attemptlimit request entity too large response has a 5xx status code
func (o *GetOutboundAttemptlimitRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound attemptlimit request entity too large response a status code equal to that given
func (o *GetOutboundAttemptlimitRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetOutboundAttemptlimitRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/attemptlimits/{attemptLimitsId}][%d] getOutboundAttemptlimitRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetOutboundAttemptlimitRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/attemptlimits/{attemptLimitsId}][%d] getOutboundAttemptlimitRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetOutboundAttemptlimitRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundAttemptlimitRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundAttemptlimitUnsupportedMediaType creates a GetOutboundAttemptlimitUnsupportedMediaType with default headers values
func NewGetOutboundAttemptlimitUnsupportedMediaType() *GetOutboundAttemptlimitUnsupportedMediaType {
	return &GetOutboundAttemptlimitUnsupportedMediaType{}
}

/*
GetOutboundAttemptlimitUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetOutboundAttemptlimitUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound attemptlimit unsupported media type response has a 2xx status code
func (o *GetOutboundAttemptlimitUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound attemptlimit unsupported media type response has a 3xx status code
func (o *GetOutboundAttemptlimitUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound attemptlimit unsupported media type response has a 4xx status code
func (o *GetOutboundAttemptlimitUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound attemptlimit unsupported media type response has a 5xx status code
func (o *GetOutboundAttemptlimitUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound attemptlimit unsupported media type response a status code equal to that given
func (o *GetOutboundAttemptlimitUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetOutboundAttemptlimitUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/attemptlimits/{attemptLimitsId}][%d] getOutboundAttemptlimitUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetOutboundAttemptlimitUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/attemptlimits/{attemptLimitsId}][%d] getOutboundAttemptlimitUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetOutboundAttemptlimitUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundAttemptlimitUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundAttemptlimitTooManyRequests creates a GetOutboundAttemptlimitTooManyRequests with default headers values
func NewGetOutboundAttemptlimitTooManyRequests() *GetOutboundAttemptlimitTooManyRequests {
	return &GetOutboundAttemptlimitTooManyRequests{}
}

/*
GetOutboundAttemptlimitTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetOutboundAttemptlimitTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound attemptlimit too many requests response has a 2xx status code
func (o *GetOutboundAttemptlimitTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound attemptlimit too many requests response has a 3xx status code
func (o *GetOutboundAttemptlimitTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound attemptlimit too many requests response has a 4xx status code
func (o *GetOutboundAttemptlimitTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound attemptlimit too many requests response has a 5xx status code
func (o *GetOutboundAttemptlimitTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound attemptlimit too many requests response a status code equal to that given
func (o *GetOutboundAttemptlimitTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetOutboundAttemptlimitTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/attemptlimits/{attemptLimitsId}][%d] getOutboundAttemptlimitTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOutboundAttemptlimitTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/attemptlimits/{attemptLimitsId}][%d] getOutboundAttemptlimitTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOutboundAttemptlimitTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundAttemptlimitTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundAttemptlimitInternalServerError creates a GetOutboundAttemptlimitInternalServerError with default headers values
func NewGetOutboundAttemptlimitInternalServerError() *GetOutboundAttemptlimitInternalServerError {
	return &GetOutboundAttemptlimitInternalServerError{}
}

/*
GetOutboundAttemptlimitInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetOutboundAttemptlimitInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound attemptlimit internal server error response has a 2xx status code
func (o *GetOutboundAttemptlimitInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound attemptlimit internal server error response has a 3xx status code
func (o *GetOutboundAttemptlimitInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound attemptlimit internal server error response has a 4xx status code
func (o *GetOutboundAttemptlimitInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get outbound attemptlimit internal server error response has a 5xx status code
func (o *GetOutboundAttemptlimitInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get outbound attemptlimit internal server error response a status code equal to that given
func (o *GetOutboundAttemptlimitInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetOutboundAttemptlimitInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/attemptlimits/{attemptLimitsId}][%d] getOutboundAttemptlimitInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOutboundAttemptlimitInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/attemptlimits/{attemptLimitsId}][%d] getOutboundAttemptlimitInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOutboundAttemptlimitInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundAttemptlimitInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundAttemptlimitServiceUnavailable creates a GetOutboundAttemptlimitServiceUnavailable with default headers values
func NewGetOutboundAttemptlimitServiceUnavailable() *GetOutboundAttemptlimitServiceUnavailable {
	return &GetOutboundAttemptlimitServiceUnavailable{}
}

/*
GetOutboundAttemptlimitServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetOutboundAttemptlimitServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound attemptlimit service unavailable response has a 2xx status code
func (o *GetOutboundAttemptlimitServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound attemptlimit service unavailable response has a 3xx status code
func (o *GetOutboundAttemptlimitServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound attemptlimit service unavailable response has a 4xx status code
func (o *GetOutboundAttemptlimitServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get outbound attemptlimit service unavailable response has a 5xx status code
func (o *GetOutboundAttemptlimitServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get outbound attemptlimit service unavailable response a status code equal to that given
func (o *GetOutboundAttemptlimitServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetOutboundAttemptlimitServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/attemptlimits/{attemptLimitsId}][%d] getOutboundAttemptlimitServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetOutboundAttemptlimitServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/attemptlimits/{attemptLimitsId}][%d] getOutboundAttemptlimitServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetOutboundAttemptlimitServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundAttemptlimitServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundAttemptlimitGatewayTimeout creates a GetOutboundAttemptlimitGatewayTimeout with default headers values
func NewGetOutboundAttemptlimitGatewayTimeout() *GetOutboundAttemptlimitGatewayTimeout {
	return &GetOutboundAttemptlimitGatewayTimeout{}
}

/*
GetOutboundAttemptlimitGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetOutboundAttemptlimitGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound attemptlimit gateway timeout response has a 2xx status code
func (o *GetOutboundAttemptlimitGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound attemptlimit gateway timeout response has a 3xx status code
func (o *GetOutboundAttemptlimitGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound attemptlimit gateway timeout response has a 4xx status code
func (o *GetOutboundAttemptlimitGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get outbound attemptlimit gateway timeout response has a 5xx status code
func (o *GetOutboundAttemptlimitGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get outbound attemptlimit gateway timeout response a status code equal to that given
func (o *GetOutboundAttemptlimitGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetOutboundAttemptlimitGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/attemptlimits/{attemptLimitsId}][%d] getOutboundAttemptlimitGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetOutboundAttemptlimitGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/attemptlimits/{attemptLimitsId}][%d] getOutboundAttemptlimitGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetOutboundAttemptlimitGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundAttemptlimitGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

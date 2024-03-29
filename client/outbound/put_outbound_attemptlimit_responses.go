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

// PutOutboundAttemptlimitReader is a Reader for the PutOutboundAttemptlimit structure.
type PutOutboundAttemptlimitReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutOutboundAttemptlimitReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutOutboundAttemptlimitOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutOutboundAttemptlimitBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutOutboundAttemptlimitUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutOutboundAttemptlimitForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutOutboundAttemptlimitNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPutOutboundAttemptlimitRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPutOutboundAttemptlimitRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPutOutboundAttemptlimitUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutOutboundAttemptlimitTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutOutboundAttemptlimitInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutOutboundAttemptlimitServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutOutboundAttemptlimitGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutOutboundAttemptlimitOK creates a PutOutboundAttemptlimitOK with default headers values
func NewPutOutboundAttemptlimitOK() *PutOutboundAttemptlimitOK {
	return &PutOutboundAttemptlimitOK{}
}

/*
PutOutboundAttemptlimitOK describes a response with status code 200, with default header values.

successful operation
*/
type PutOutboundAttemptlimitOK struct {
	Payload *models.AttemptLimits
}

// IsSuccess returns true when this put outbound attemptlimit o k response has a 2xx status code
func (o *PutOutboundAttemptlimitOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put outbound attemptlimit o k response has a 3xx status code
func (o *PutOutboundAttemptlimitOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put outbound attemptlimit o k response has a 4xx status code
func (o *PutOutboundAttemptlimitOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put outbound attemptlimit o k response has a 5xx status code
func (o *PutOutboundAttemptlimitOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put outbound attemptlimit o k response a status code equal to that given
func (o *PutOutboundAttemptlimitOK) IsCode(code int) bool {
	return code == 200
}

func (o *PutOutboundAttemptlimitOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/attemptlimits/{attemptLimitsId}][%d] putOutboundAttemptlimitOK  %+v", 200, o.Payload)
}

func (o *PutOutboundAttemptlimitOK) String() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/attemptlimits/{attemptLimitsId}][%d] putOutboundAttemptlimitOK  %+v", 200, o.Payload)
}

func (o *PutOutboundAttemptlimitOK) GetPayload() *models.AttemptLimits {
	return o.Payload
}

func (o *PutOutboundAttemptlimitOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AttemptLimits)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundAttemptlimitBadRequest creates a PutOutboundAttemptlimitBadRequest with default headers values
func NewPutOutboundAttemptlimitBadRequest() *PutOutboundAttemptlimitBadRequest {
	return &PutOutboundAttemptlimitBadRequest{}
}

/*
PutOutboundAttemptlimitBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutOutboundAttemptlimitBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put outbound attemptlimit bad request response has a 2xx status code
func (o *PutOutboundAttemptlimitBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put outbound attemptlimit bad request response has a 3xx status code
func (o *PutOutboundAttemptlimitBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put outbound attemptlimit bad request response has a 4xx status code
func (o *PutOutboundAttemptlimitBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this put outbound attemptlimit bad request response has a 5xx status code
func (o *PutOutboundAttemptlimitBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this put outbound attemptlimit bad request response a status code equal to that given
func (o *PutOutboundAttemptlimitBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PutOutboundAttemptlimitBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/attemptlimits/{attemptLimitsId}][%d] putOutboundAttemptlimitBadRequest  %+v", 400, o.Payload)
}

func (o *PutOutboundAttemptlimitBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/attemptlimits/{attemptLimitsId}][%d] putOutboundAttemptlimitBadRequest  %+v", 400, o.Payload)
}

func (o *PutOutboundAttemptlimitBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundAttemptlimitBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundAttemptlimitUnauthorized creates a PutOutboundAttemptlimitUnauthorized with default headers values
func NewPutOutboundAttemptlimitUnauthorized() *PutOutboundAttemptlimitUnauthorized {
	return &PutOutboundAttemptlimitUnauthorized{}
}

/*
PutOutboundAttemptlimitUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PutOutboundAttemptlimitUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put outbound attemptlimit unauthorized response has a 2xx status code
func (o *PutOutboundAttemptlimitUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put outbound attemptlimit unauthorized response has a 3xx status code
func (o *PutOutboundAttemptlimitUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put outbound attemptlimit unauthorized response has a 4xx status code
func (o *PutOutboundAttemptlimitUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this put outbound attemptlimit unauthorized response has a 5xx status code
func (o *PutOutboundAttemptlimitUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this put outbound attemptlimit unauthorized response a status code equal to that given
func (o *PutOutboundAttemptlimitUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PutOutboundAttemptlimitUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/attemptlimits/{attemptLimitsId}][%d] putOutboundAttemptlimitUnauthorized  %+v", 401, o.Payload)
}

func (o *PutOutboundAttemptlimitUnauthorized) String() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/attemptlimits/{attemptLimitsId}][%d] putOutboundAttemptlimitUnauthorized  %+v", 401, o.Payload)
}

func (o *PutOutboundAttemptlimitUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundAttemptlimitUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundAttemptlimitForbidden creates a PutOutboundAttemptlimitForbidden with default headers values
func NewPutOutboundAttemptlimitForbidden() *PutOutboundAttemptlimitForbidden {
	return &PutOutboundAttemptlimitForbidden{}
}

/*
PutOutboundAttemptlimitForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PutOutboundAttemptlimitForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put outbound attemptlimit forbidden response has a 2xx status code
func (o *PutOutboundAttemptlimitForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put outbound attemptlimit forbidden response has a 3xx status code
func (o *PutOutboundAttemptlimitForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put outbound attemptlimit forbidden response has a 4xx status code
func (o *PutOutboundAttemptlimitForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this put outbound attemptlimit forbidden response has a 5xx status code
func (o *PutOutboundAttemptlimitForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this put outbound attemptlimit forbidden response a status code equal to that given
func (o *PutOutboundAttemptlimitForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PutOutboundAttemptlimitForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/attemptlimits/{attemptLimitsId}][%d] putOutboundAttemptlimitForbidden  %+v", 403, o.Payload)
}

func (o *PutOutboundAttemptlimitForbidden) String() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/attemptlimits/{attemptLimitsId}][%d] putOutboundAttemptlimitForbidden  %+v", 403, o.Payload)
}

func (o *PutOutboundAttemptlimitForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundAttemptlimitForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundAttemptlimitNotFound creates a PutOutboundAttemptlimitNotFound with default headers values
func NewPutOutboundAttemptlimitNotFound() *PutOutboundAttemptlimitNotFound {
	return &PutOutboundAttemptlimitNotFound{}
}

/*
PutOutboundAttemptlimitNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PutOutboundAttemptlimitNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put outbound attemptlimit not found response has a 2xx status code
func (o *PutOutboundAttemptlimitNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put outbound attemptlimit not found response has a 3xx status code
func (o *PutOutboundAttemptlimitNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put outbound attemptlimit not found response has a 4xx status code
func (o *PutOutboundAttemptlimitNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this put outbound attemptlimit not found response has a 5xx status code
func (o *PutOutboundAttemptlimitNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this put outbound attemptlimit not found response a status code equal to that given
func (o *PutOutboundAttemptlimitNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PutOutboundAttemptlimitNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/attemptlimits/{attemptLimitsId}][%d] putOutboundAttemptlimitNotFound  %+v", 404, o.Payload)
}

func (o *PutOutboundAttemptlimitNotFound) String() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/attemptlimits/{attemptLimitsId}][%d] putOutboundAttemptlimitNotFound  %+v", 404, o.Payload)
}

func (o *PutOutboundAttemptlimitNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundAttemptlimitNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundAttemptlimitRequestTimeout creates a PutOutboundAttemptlimitRequestTimeout with default headers values
func NewPutOutboundAttemptlimitRequestTimeout() *PutOutboundAttemptlimitRequestTimeout {
	return &PutOutboundAttemptlimitRequestTimeout{}
}

/*
PutOutboundAttemptlimitRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PutOutboundAttemptlimitRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put outbound attemptlimit request timeout response has a 2xx status code
func (o *PutOutboundAttemptlimitRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put outbound attemptlimit request timeout response has a 3xx status code
func (o *PutOutboundAttemptlimitRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put outbound attemptlimit request timeout response has a 4xx status code
func (o *PutOutboundAttemptlimitRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this put outbound attemptlimit request timeout response has a 5xx status code
func (o *PutOutboundAttemptlimitRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this put outbound attemptlimit request timeout response a status code equal to that given
func (o *PutOutboundAttemptlimitRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PutOutboundAttemptlimitRequestTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/attemptlimits/{attemptLimitsId}][%d] putOutboundAttemptlimitRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutOutboundAttemptlimitRequestTimeout) String() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/attemptlimits/{attemptLimitsId}][%d] putOutboundAttemptlimitRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutOutboundAttemptlimitRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundAttemptlimitRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundAttemptlimitRequestEntityTooLarge creates a PutOutboundAttemptlimitRequestEntityTooLarge with default headers values
func NewPutOutboundAttemptlimitRequestEntityTooLarge() *PutOutboundAttemptlimitRequestEntityTooLarge {
	return &PutOutboundAttemptlimitRequestEntityTooLarge{}
}

/*
PutOutboundAttemptlimitRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PutOutboundAttemptlimitRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put outbound attemptlimit request entity too large response has a 2xx status code
func (o *PutOutboundAttemptlimitRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put outbound attemptlimit request entity too large response has a 3xx status code
func (o *PutOutboundAttemptlimitRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put outbound attemptlimit request entity too large response has a 4xx status code
func (o *PutOutboundAttemptlimitRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this put outbound attemptlimit request entity too large response has a 5xx status code
func (o *PutOutboundAttemptlimitRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this put outbound attemptlimit request entity too large response a status code equal to that given
func (o *PutOutboundAttemptlimitRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PutOutboundAttemptlimitRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/attemptlimits/{attemptLimitsId}][%d] putOutboundAttemptlimitRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutOutboundAttemptlimitRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/attemptlimits/{attemptLimitsId}][%d] putOutboundAttemptlimitRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutOutboundAttemptlimitRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundAttemptlimitRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundAttemptlimitUnsupportedMediaType creates a PutOutboundAttemptlimitUnsupportedMediaType with default headers values
func NewPutOutboundAttemptlimitUnsupportedMediaType() *PutOutboundAttemptlimitUnsupportedMediaType {
	return &PutOutboundAttemptlimitUnsupportedMediaType{}
}

/*
PutOutboundAttemptlimitUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutOutboundAttemptlimitUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put outbound attemptlimit unsupported media type response has a 2xx status code
func (o *PutOutboundAttemptlimitUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put outbound attemptlimit unsupported media type response has a 3xx status code
func (o *PutOutboundAttemptlimitUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put outbound attemptlimit unsupported media type response has a 4xx status code
func (o *PutOutboundAttemptlimitUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this put outbound attemptlimit unsupported media type response has a 5xx status code
func (o *PutOutboundAttemptlimitUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this put outbound attemptlimit unsupported media type response a status code equal to that given
func (o *PutOutboundAttemptlimitUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PutOutboundAttemptlimitUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/attemptlimits/{attemptLimitsId}][%d] putOutboundAttemptlimitUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutOutboundAttemptlimitUnsupportedMediaType) String() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/attemptlimits/{attemptLimitsId}][%d] putOutboundAttemptlimitUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutOutboundAttemptlimitUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundAttemptlimitUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundAttemptlimitTooManyRequests creates a PutOutboundAttemptlimitTooManyRequests with default headers values
func NewPutOutboundAttemptlimitTooManyRequests() *PutOutboundAttemptlimitTooManyRequests {
	return &PutOutboundAttemptlimitTooManyRequests{}
}

/*
PutOutboundAttemptlimitTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PutOutboundAttemptlimitTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put outbound attemptlimit too many requests response has a 2xx status code
func (o *PutOutboundAttemptlimitTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put outbound attemptlimit too many requests response has a 3xx status code
func (o *PutOutboundAttemptlimitTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put outbound attemptlimit too many requests response has a 4xx status code
func (o *PutOutboundAttemptlimitTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this put outbound attemptlimit too many requests response has a 5xx status code
func (o *PutOutboundAttemptlimitTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this put outbound attemptlimit too many requests response a status code equal to that given
func (o *PutOutboundAttemptlimitTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PutOutboundAttemptlimitTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/attemptlimits/{attemptLimitsId}][%d] putOutboundAttemptlimitTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutOutboundAttemptlimitTooManyRequests) String() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/attemptlimits/{attemptLimitsId}][%d] putOutboundAttemptlimitTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutOutboundAttemptlimitTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundAttemptlimitTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundAttemptlimitInternalServerError creates a PutOutboundAttemptlimitInternalServerError with default headers values
func NewPutOutboundAttemptlimitInternalServerError() *PutOutboundAttemptlimitInternalServerError {
	return &PutOutboundAttemptlimitInternalServerError{}
}

/*
PutOutboundAttemptlimitInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutOutboundAttemptlimitInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put outbound attemptlimit internal server error response has a 2xx status code
func (o *PutOutboundAttemptlimitInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put outbound attemptlimit internal server error response has a 3xx status code
func (o *PutOutboundAttemptlimitInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put outbound attemptlimit internal server error response has a 4xx status code
func (o *PutOutboundAttemptlimitInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this put outbound attemptlimit internal server error response has a 5xx status code
func (o *PutOutboundAttemptlimitInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this put outbound attemptlimit internal server error response a status code equal to that given
func (o *PutOutboundAttemptlimitInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PutOutboundAttemptlimitInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/attemptlimits/{attemptLimitsId}][%d] putOutboundAttemptlimitInternalServerError  %+v", 500, o.Payload)
}

func (o *PutOutboundAttemptlimitInternalServerError) String() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/attemptlimits/{attemptLimitsId}][%d] putOutboundAttemptlimitInternalServerError  %+v", 500, o.Payload)
}

func (o *PutOutboundAttemptlimitInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundAttemptlimitInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundAttemptlimitServiceUnavailable creates a PutOutboundAttemptlimitServiceUnavailable with default headers values
func NewPutOutboundAttemptlimitServiceUnavailable() *PutOutboundAttemptlimitServiceUnavailable {
	return &PutOutboundAttemptlimitServiceUnavailable{}
}

/*
PutOutboundAttemptlimitServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutOutboundAttemptlimitServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put outbound attemptlimit service unavailable response has a 2xx status code
func (o *PutOutboundAttemptlimitServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put outbound attemptlimit service unavailable response has a 3xx status code
func (o *PutOutboundAttemptlimitServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put outbound attemptlimit service unavailable response has a 4xx status code
func (o *PutOutboundAttemptlimitServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this put outbound attemptlimit service unavailable response has a 5xx status code
func (o *PutOutboundAttemptlimitServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this put outbound attemptlimit service unavailable response a status code equal to that given
func (o *PutOutboundAttemptlimitServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PutOutboundAttemptlimitServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/attemptlimits/{attemptLimitsId}][%d] putOutboundAttemptlimitServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutOutboundAttemptlimitServiceUnavailable) String() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/attemptlimits/{attemptLimitsId}][%d] putOutboundAttemptlimitServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutOutboundAttemptlimitServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundAttemptlimitServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundAttemptlimitGatewayTimeout creates a PutOutboundAttemptlimitGatewayTimeout with default headers values
func NewPutOutboundAttemptlimitGatewayTimeout() *PutOutboundAttemptlimitGatewayTimeout {
	return &PutOutboundAttemptlimitGatewayTimeout{}
}

/*
PutOutboundAttemptlimitGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PutOutboundAttemptlimitGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put outbound attemptlimit gateway timeout response has a 2xx status code
func (o *PutOutboundAttemptlimitGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put outbound attemptlimit gateway timeout response has a 3xx status code
func (o *PutOutboundAttemptlimitGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put outbound attemptlimit gateway timeout response has a 4xx status code
func (o *PutOutboundAttemptlimitGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this put outbound attemptlimit gateway timeout response has a 5xx status code
func (o *PutOutboundAttemptlimitGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this put outbound attemptlimit gateway timeout response a status code equal to that given
func (o *PutOutboundAttemptlimitGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PutOutboundAttemptlimitGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/attemptlimits/{attemptLimitsId}][%d] putOutboundAttemptlimitGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutOutboundAttemptlimitGatewayTimeout) String() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/attemptlimits/{attemptLimitsId}][%d] putOutboundAttemptlimitGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutOutboundAttemptlimitGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundAttemptlimitGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

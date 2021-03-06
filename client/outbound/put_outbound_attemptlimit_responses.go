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

/*PutOutboundAttemptlimitOK handles this case with default header values.

successful operation
*/
type PutOutboundAttemptlimitOK struct {
	Payload *models.AttemptLimits
}

func (o *PutOutboundAttemptlimitOK) Error() string {
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

/*PutOutboundAttemptlimitBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutOutboundAttemptlimitBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PutOutboundAttemptlimitBadRequest) Error() string {
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

/*PutOutboundAttemptlimitUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PutOutboundAttemptlimitUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PutOutboundAttemptlimitUnauthorized) Error() string {
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

/*PutOutboundAttemptlimitForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PutOutboundAttemptlimitForbidden struct {
	Payload *models.ErrorBody
}

func (o *PutOutboundAttemptlimitForbidden) Error() string {
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

/*PutOutboundAttemptlimitNotFound handles this case with default header values.

The requested resource was not found.
*/
type PutOutboundAttemptlimitNotFound struct {
	Payload *models.ErrorBody
}

func (o *PutOutboundAttemptlimitNotFound) Error() string {
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

// NewPutOutboundAttemptlimitRequestEntityTooLarge creates a PutOutboundAttemptlimitRequestEntityTooLarge with default headers values
func NewPutOutboundAttemptlimitRequestEntityTooLarge() *PutOutboundAttemptlimitRequestEntityTooLarge {
	return &PutOutboundAttemptlimitRequestEntityTooLarge{}
}

/*PutOutboundAttemptlimitRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PutOutboundAttemptlimitRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PutOutboundAttemptlimitRequestEntityTooLarge) Error() string {
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

/*PutOutboundAttemptlimitUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutOutboundAttemptlimitUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PutOutboundAttemptlimitUnsupportedMediaType) Error() string {
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

/*PutOutboundAttemptlimitTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PutOutboundAttemptlimitTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PutOutboundAttemptlimitTooManyRequests) Error() string {
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

/*PutOutboundAttemptlimitInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutOutboundAttemptlimitInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PutOutboundAttemptlimitInternalServerError) Error() string {
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

/*PutOutboundAttemptlimitServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutOutboundAttemptlimitServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PutOutboundAttemptlimitServiceUnavailable) Error() string {
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

/*PutOutboundAttemptlimitGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PutOutboundAttemptlimitGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PutOutboundAttemptlimitGatewayTimeout) Error() string {
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

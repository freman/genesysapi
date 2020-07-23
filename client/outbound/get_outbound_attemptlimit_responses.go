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

/*GetOutboundAttemptlimitOK handles this case with default header values.

successful operation
*/
type GetOutboundAttemptlimitOK struct {
	Payload *models.AttemptLimits
}

func (o *GetOutboundAttemptlimitOK) Error() string {
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

/*GetOutboundAttemptlimitBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetOutboundAttemptlimitBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundAttemptlimitBadRequest) Error() string {
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

/*GetOutboundAttemptlimitUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetOutboundAttemptlimitUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundAttemptlimitUnauthorized) Error() string {
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

/*GetOutboundAttemptlimitForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetOutboundAttemptlimitForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundAttemptlimitForbidden) Error() string {
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

/*GetOutboundAttemptlimitNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetOutboundAttemptlimitNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundAttemptlimitNotFound) Error() string {
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

// NewGetOutboundAttemptlimitRequestEntityTooLarge creates a GetOutboundAttemptlimitRequestEntityTooLarge with default headers values
func NewGetOutboundAttemptlimitRequestEntityTooLarge() *GetOutboundAttemptlimitRequestEntityTooLarge {
	return &GetOutboundAttemptlimitRequestEntityTooLarge{}
}

/*GetOutboundAttemptlimitRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetOutboundAttemptlimitRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundAttemptlimitRequestEntityTooLarge) Error() string {
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

/*GetOutboundAttemptlimitUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetOutboundAttemptlimitUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundAttemptlimitUnsupportedMediaType) Error() string {
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

/*GetOutboundAttemptlimitTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetOutboundAttemptlimitTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundAttemptlimitTooManyRequests) Error() string {
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

/*GetOutboundAttemptlimitInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetOutboundAttemptlimitInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundAttemptlimitInternalServerError) Error() string {
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

/*GetOutboundAttemptlimitServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetOutboundAttemptlimitServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundAttemptlimitServiceUnavailable) Error() string {
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

/*GetOutboundAttemptlimitGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetOutboundAttemptlimitGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundAttemptlimitGatewayTimeout) Error() string {
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

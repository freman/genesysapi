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

// DeleteOutboundAttemptlimitReader is a Reader for the DeleteOutboundAttemptlimit structure.
type DeleteOutboundAttemptlimitReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteOutboundAttemptlimitReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteOutboundAttemptlimitOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteOutboundAttemptlimitBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteOutboundAttemptlimitUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteOutboundAttemptlimitForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteOutboundAttemptlimitNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteOutboundAttemptlimitRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteOutboundAttemptlimitUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteOutboundAttemptlimitTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteOutboundAttemptlimitInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteOutboundAttemptlimitServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteOutboundAttemptlimitGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteOutboundAttemptlimitOK creates a DeleteOutboundAttemptlimitOK with default headers values
func NewDeleteOutboundAttemptlimitOK() *DeleteOutboundAttemptlimitOK {
	return &DeleteOutboundAttemptlimitOK{}
}

/*DeleteOutboundAttemptlimitOK handles this case with default header values.

Operation was successful.
*/
type DeleteOutboundAttemptlimitOK struct {
}

func (o *DeleteOutboundAttemptlimitOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/attemptlimits/{attemptLimitsId}][%d] deleteOutboundAttemptlimitOK ", 200)
}

func (o *DeleteOutboundAttemptlimitOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteOutboundAttemptlimitBadRequest creates a DeleteOutboundAttemptlimitBadRequest with default headers values
func NewDeleteOutboundAttemptlimitBadRequest() *DeleteOutboundAttemptlimitBadRequest {
	return &DeleteOutboundAttemptlimitBadRequest{}
}

/*DeleteOutboundAttemptlimitBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteOutboundAttemptlimitBadRequest struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundAttemptlimitBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/attemptlimits/{attemptLimitsId}][%d] deleteOutboundAttemptlimitBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteOutboundAttemptlimitBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundAttemptlimitBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundAttemptlimitUnauthorized creates a DeleteOutboundAttemptlimitUnauthorized with default headers values
func NewDeleteOutboundAttemptlimitUnauthorized() *DeleteOutboundAttemptlimitUnauthorized {
	return &DeleteOutboundAttemptlimitUnauthorized{}
}

/*DeleteOutboundAttemptlimitUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteOutboundAttemptlimitUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundAttemptlimitUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/attemptlimits/{attemptLimitsId}][%d] deleteOutboundAttemptlimitUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteOutboundAttemptlimitUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundAttemptlimitUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundAttemptlimitForbidden creates a DeleteOutboundAttemptlimitForbidden with default headers values
func NewDeleteOutboundAttemptlimitForbidden() *DeleteOutboundAttemptlimitForbidden {
	return &DeleteOutboundAttemptlimitForbidden{}
}

/*DeleteOutboundAttemptlimitForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type DeleteOutboundAttemptlimitForbidden struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundAttemptlimitForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/attemptlimits/{attemptLimitsId}][%d] deleteOutboundAttemptlimitForbidden  %+v", 403, o.Payload)
}

func (o *DeleteOutboundAttemptlimitForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundAttemptlimitForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundAttemptlimitNotFound creates a DeleteOutboundAttemptlimitNotFound with default headers values
func NewDeleteOutboundAttemptlimitNotFound() *DeleteOutboundAttemptlimitNotFound {
	return &DeleteOutboundAttemptlimitNotFound{}
}

/*DeleteOutboundAttemptlimitNotFound handles this case with default header values.

The requested resource was not found.
*/
type DeleteOutboundAttemptlimitNotFound struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundAttemptlimitNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/attemptlimits/{attemptLimitsId}][%d] deleteOutboundAttemptlimitNotFound  %+v", 404, o.Payload)
}

func (o *DeleteOutboundAttemptlimitNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundAttemptlimitNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundAttemptlimitRequestEntityTooLarge creates a DeleteOutboundAttemptlimitRequestEntityTooLarge with default headers values
func NewDeleteOutboundAttemptlimitRequestEntityTooLarge() *DeleteOutboundAttemptlimitRequestEntityTooLarge {
	return &DeleteOutboundAttemptlimitRequestEntityTooLarge{}
}

/*DeleteOutboundAttemptlimitRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type DeleteOutboundAttemptlimitRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundAttemptlimitRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/attemptlimits/{attemptLimitsId}][%d] deleteOutboundAttemptlimitRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteOutboundAttemptlimitRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundAttemptlimitRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundAttemptlimitUnsupportedMediaType creates a DeleteOutboundAttemptlimitUnsupportedMediaType with default headers values
func NewDeleteOutboundAttemptlimitUnsupportedMediaType() *DeleteOutboundAttemptlimitUnsupportedMediaType {
	return &DeleteOutboundAttemptlimitUnsupportedMediaType{}
}

/*DeleteOutboundAttemptlimitUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteOutboundAttemptlimitUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundAttemptlimitUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/attemptlimits/{attemptLimitsId}][%d] deleteOutboundAttemptlimitUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteOutboundAttemptlimitUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundAttemptlimitUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundAttemptlimitTooManyRequests creates a DeleteOutboundAttemptlimitTooManyRequests with default headers values
func NewDeleteOutboundAttemptlimitTooManyRequests() *DeleteOutboundAttemptlimitTooManyRequests {
	return &DeleteOutboundAttemptlimitTooManyRequests{}
}

/*DeleteOutboundAttemptlimitTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type DeleteOutboundAttemptlimitTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundAttemptlimitTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/attemptlimits/{attemptLimitsId}][%d] deleteOutboundAttemptlimitTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteOutboundAttemptlimitTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundAttemptlimitTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundAttemptlimitInternalServerError creates a DeleteOutboundAttemptlimitInternalServerError with default headers values
func NewDeleteOutboundAttemptlimitInternalServerError() *DeleteOutboundAttemptlimitInternalServerError {
	return &DeleteOutboundAttemptlimitInternalServerError{}
}

/*DeleteOutboundAttemptlimitInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteOutboundAttemptlimitInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundAttemptlimitInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/attemptlimits/{attemptLimitsId}][%d] deleteOutboundAttemptlimitInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteOutboundAttemptlimitInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundAttemptlimitInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundAttemptlimitServiceUnavailable creates a DeleteOutboundAttemptlimitServiceUnavailable with default headers values
func NewDeleteOutboundAttemptlimitServiceUnavailable() *DeleteOutboundAttemptlimitServiceUnavailable {
	return &DeleteOutboundAttemptlimitServiceUnavailable{}
}

/*DeleteOutboundAttemptlimitServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteOutboundAttemptlimitServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundAttemptlimitServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/attemptlimits/{attemptLimitsId}][%d] deleteOutboundAttemptlimitServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteOutboundAttemptlimitServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundAttemptlimitServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundAttemptlimitGatewayTimeout creates a DeleteOutboundAttemptlimitGatewayTimeout with default headers values
func NewDeleteOutboundAttemptlimitGatewayTimeout() *DeleteOutboundAttemptlimitGatewayTimeout {
	return &DeleteOutboundAttemptlimitGatewayTimeout{}
}

/*DeleteOutboundAttemptlimitGatewayTimeout handles this case with default header values.

The request timed out.
*/
type DeleteOutboundAttemptlimitGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundAttemptlimitGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/attemptlimits/{attemptLimitsId}][%d] deleteOutboundAttemptlimitGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteOutboundAttemptlimitGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundAttemptlimitGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package routing

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PutRoutingMessageRecipientReader is a Reader for the PutRoutingMessageRecipient structure.
type PutRoutingMessageRecipientReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutRoutingMessageRecipientReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutRoutingMessageRecipientOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutRoutingMessageRecipientBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutRoutingMessageRecipientUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutRoutingMessageRecipientForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutRoutingMessageRecipientNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPutRoutingMessageRecipientRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPutRoutingMessageRecipientUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutRoutingMessageRecipientTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutRoutingMessageRecipientInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutRoutingMessageRecipientServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutRoutingMessageRecipientGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutRoutingMessageRecipientOK creates a PutRoutingMessageRecipientOK with default headers values
func NewPutRoutingMessageRecipientOK() *PutRoutingMessageRecipientOK {
	return &PutRoutingMessageRecipientOK{}
}

/*PutRoutingMessageRecipientOK handles this case with default header values.

successful operation
*/
type PutRoutingMessageRecipientOK struct {
	Payload *models.Recipient
}

func (o *PutRoutingMessageRecipientOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/routing/message/recipients/{recipientId}][%d] putRoutingMessageRecipientOK  %+v", 200, o.Payload)
}

func (o *PutRoutingMessageRecipientOK) GetPayload() *models.Recipient {
	return o.Payload
}

func (o *PutRoutingMessageRecipientOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Recipient)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRoutingMessageRecipientBadRequest creates a PutRoutingMessageRecipientBadRequest with default headers values
func NewPutRoutingMessageRecipientBadRequest() *PutRoutingMessageRecipientBadRequest {
	return &PutRoutingMessageRecipientBadRequest{}
}

/*PutRoutingMessageRecipientBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutRoutingMessageRecipientBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PutRoutingMessageRecipientBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/routing/message/recipients/{recipientId}][%d] putRoutingMessageRecipientBadRequest  %+v", 400, o.Payload)
}

func (o *PutRoutingMessageRecipientBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutRoutingMessageRecipientBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRoutingMessageRecipientUnauthorized creates a PutRoutingMessageRecipientUnauthorized with default headers values
func NewPutRoutingMessageRecipientUnauthorized() *PutRoutingMessageRecipientUnauthorized {
	return &PutRoutingMessageRecipientUnauthorized{}
}

/*PutRoutingMessageRecipientUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PutRoutingMessageRecipientUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PutRoutingMessageRecipientUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/routing/message/recipients/{recipientId}][%d] putRoutingMessageRecipientUnauthorized  %+v", 401, o.Payload)
}

func (o *PutRoutingMessageRecipientUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutRoutingMessageRecipientUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRoutingMessageRecipientForbidden creates a PutRoutingMessageRecipientForbidden with default headers values
func NewPutRoutingMessageRecipientForbidden() *PutRoutingMessageRecipientForbidden {
	return &PutRoutingMessageRecipientForbidden{}
}

/*PutRoutingMessageRecipientForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PutRoutingMessageRecipientForbidden struct {
	Payload *models.ErrorBody
}

func (o *PutRoutingMessageRecipientForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/routing/message/recipients/{recipientId}][%d] putRoutingMessageRecipientForbidden  %+v", 403, o.Payload)
}

func (o *PutRoutingMessageRecipientForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutRoutingMessageRecipientForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRoutingMessageRecipientNotFound creates a PutRoutingMessageRecipientNotFound with default headers values
func NewPutRoutingMessageRecipientNotFound() *PutRoutingMessageRecipientNotFound {
	return &PutRoutingMessageRecipientNotFound{}
}

/*PutRoutingMessageRecipientNotFound handles this case with default header values.

The requested resource was not found.
*/
type PutRoutingMessageRecipientNotFound struct {
	Payload *models.ErrorBody
}

func (o *PutRoutingMessageRecipientNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/routing/message/recipients/{recipientId}][%d] putRoutingMessageRecipientNotFound  %+v", 404, o.Payload)
}

func (o *PutRoutingMessageRecipientNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutRoutingMessageRecipientNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRoutingMessageRecipientRequestEntityTooLarge creates a PutRoutingMessageRecipientRequestEntityTooLarge with default headers values
func NewPutRoutingMessageRecipientRequestEntityTooLarge() *PutRoutingMessageRecipientRequestEntityTooLarge {
	return &PutRoutingMessageRecipientRequestEntityTooLarge{}
}

/*PutRoutingMessageRecipientRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PutRoutingMessageRecipientRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PutRoutingMessageRecipientRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /api/v2/routing/message/recipients/{recipientId}][%d] putRoutingMessageRecipientRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutRoutingMessageRecipientRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutRoutingMessageRecipientRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRoutingMessageRecipientUnsupportedMediaType creates a PutRoutingMessageRecipientUnsupportedMediaType with default headers values
func NewPutRoutingMessageRecipientUnsupportedMediaType() *PutRoutingMessageRecipientUnsupportedMediaType {
	return &PutRoutingMessageRecipientUnsupportedMediaType{}
}

/*PutRoutingMessageRecipientUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutRoutingMessageRecipientUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PutRoutingMessageRecipientUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v2/routing/message/recipients/{recipientId}][%d] putRoutingMessageRecipientUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutRoutingMessageRecipientUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutRoutingMessageRecipientUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRoutingMessageRecipientTooManyRequests creates a PutRoutingMessageRecipientTooManyRequests with default headers values
func NewPutRoutingMessageRecipientTooManyRequests() *PutRoutingMessageRecipientTooManyRequests {
	return &PutRoutingMessageRecipientTooManyRequests{}
}

/*PutRoutingMessageRecipientTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PutRoutingMessageRecipientTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PutRoutingMessageRecipientTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v2/routing/message/recipients/{recipientId}][%d] putRoutingMessageRecipientTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutRoutingMessageRecipientTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutRoutingMessageRecipientTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRoutingMessageRecipientInternalServerError creates a PutRoutingMessageRecipientInternalServerError with default headers values
func NewPutRoutingMessageRecipientInternalServerError() *PutRoutingMessageRecipientInternalServerError {
	return &PutRoutingMessageRecipientInternalServerError{}
}

/*PutRoutingMessageRecipientInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutRoutingMessageRecipientInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PutRoutingMessageRecipientInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/routing/message/recipients/{recipientId}][%d] putRoutingMessageRecipientInternalServerError  %+v", 500, o.Payload)
}

func (o *PutRoutingMessageRecipientInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutRoutingMessageRecipientInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRoutingMessageRecipientServiceUnavailable creates a PutRoutingMessageRecipientServiceUnavailable with default headers values
func NewPutRoutingMessageRecipientServiceUnavailable() *PutRoutingMessageRecipientServiceUnavailable {
	return &PutRoutingMessageRecipientServiceUnavailable{}
}

/*PutRoutingMessageRecipientServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutRoutingMessageRecipientServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PutRoutingMessageRecipientServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/routing/message/recipients/{recipientId}][%d] putRoutingMessageRecipientServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutRoutingMessageRecipientServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutRoutingMessageRecipientServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRoutingMessageRecipientGatewayTimeout creates a PutRoutingMessageRecipientGatewayTimeout with default headers values
func NewPutRoutingMessageRecipientGatewayTimeout() *PutRoutingMessageRecipientGatewayTimeout {
	return &PutRoutingMessageRecipientGatewayTimeout{}
}

/*PutRoutingMessageRecipientGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PutRoutingMessageRecipientGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PutRoutingMessageRecipientGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/routing/message/recipients/{recipientId}][%d] putRoutingMessageRecipientGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutRoutingMessageRecipientGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutRoutingMessageRecipientGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
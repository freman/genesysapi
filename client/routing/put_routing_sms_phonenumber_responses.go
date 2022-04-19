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

// PutRoutingSmsPhonenumberReader is a Reader for the PutRoutingSmsPhonenumber structure.
type PutRoutingSmsPhonenumberReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutRoutingSmsPhonenumberReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutRoutingSmsPhonenumberOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewPutRoutingSmsPhonenumberAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutRoutingSmsPhonenumberBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutRoutingSmsPhonenumberUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutRoutingSmsPhonenumberForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutRoutingSmsPhonenumberNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPutRoutingSmsPhonenumberRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPutRoutingSmsPhonenumberConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPutRoutingSmsPhonenumberRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPutRoutingSmsPhonenumberUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutRoutingSmsPhonenumberTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutRoutingSmsPhonenumberInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 501:
		result := NewPutRoutingSmsPhonenumberNotImplemented()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutRoutingSmsPhonenumberServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutRoutingSmsPhonenumberGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutRoutingSmsPhonenumberOK creates a PutRoutingSmsPhonenumberOK with default headers values
func NewPutRoutingSmsPhonenumberOK() *PutRoutingSmsPhonenumberOK {
	return &PutRoutingSmsPhonenumberOK{}
}

/*PutRoutingSmsPhonenumberOK handles this case with default header values.

successful operation
*/
type PutRoutingSmsPhonenumberOK struct {
	Payload *models.SmsPhoneNumber
}

func (o *PutRoutingSmsPhonenumberOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/routing/sms/phonenumbers/{addressId}][%d] putRoutingSmsPhonenumberOK  %+v", 200, o.Payload)
}

func (o *PutRoutingSmsPhonenumberOK) GetPayload() *models.SmsPhoneNumber {
	return o.Payload
}

func (o *PutRoutingSmsPhonenumberOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SmsPhoneNumber)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRoutingSmsPhonenumberAccepted creates a PutRoutingSmsPhonenumberAccepted with default headers values
func NewPutRoutingSmsPhonenumberAccepted() *PutRoutingSmsPhonenumberAccepted {
	return &PutRoutingSmsPhonenumberAccepted{}
}

/*PutRoutingSmsPhonenumberAccepted handles this case with default header values.

Accepted - If async is true, the phone number update is in progress.
*/
type PutRoutingSmsPhonenumberAccepted struct {
	Payload *models.SmsPhoneNumber
}

func (o *PutRoutingSmsPhonenumberAccepted) Error() string {
	return fmt.Sprintf("[PUT /api/v2/routing/sms/phonenumbers/{addressId}][%d] putRoutingSmsPhonenumberAccepted  %+v", 202, o.Payload)
}

func (o *PutRoutingSmsPhonenumberAccepted) GetPayload() *models.SmsPhoneNumber {
	return o.Payload
}

func (o *PutRoutingSmsPhonenumberAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SmsPhoneNumber)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRoutingSmsPhonenumberBadRequest creates a PutRoutingSmsPhonenumberBadRequest with default headers values
func NewPutRoutingSmsPhonenumberBadRequest() *PutRoutingSmsPhonenumberBadRequest {
	return &PutRoutingSmsPhonenumberBadRequest{}
}

/*PutRoutingSmsPhonenumberBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutRoutingSmsPhonenumberBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PutRoutingSmsPhonenumberBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/routing/sms/phonenumbers/{addressId}][%d] putRoutingSmsPhonenumberBadRequest  %+v", 400, o.Payload)
}

func (o *PutRoutingSmsPhonenumberBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutRoutingSmsPhonenumberBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRoutingSmsPhonenumberUnauthorized creates a PutRoutingSmsPhonenumberUnauthorized with default headers values
func NewPutRoutingSmsPhonenumberUnauthorized() *PutRoutingSmsPhonenumberUnauthorized {
	return &PutRoutingSmsPhonenumberUnauthorized{}
}

/*PutRoutingSmsPhonenumberUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PutRoutingSmsPhonenumberUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PutRoutingSmsPhonenumberUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/routing/sms/phonenumbers/{addressId}][%d] putRoutingSmsPhonenumberUnauthorized  %+v", 401, o.Payload)
}

func (o *PutRoutingSmsPhonenumberUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutRoutingSmsPhonenumberUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRoutingSmsPhonenumberForbidden creates a PutRoutingSmsPhonenumberForbidden with default headers values
func NewPutRoutingSmsPhonenumberForbidden() *PutRoutingSmsPhonenumberForbidden {
	return &PutRoutingSmsPhonenumberForbidden{}
}

/*PutRoutingSmsPhonenumberForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PutRoutingSmsPhonenumberForbidden struct {
	Payload *models.ErrorBody
}

func (o *PutRoutingSmsPhonenumberForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/routing/sms/phonenumbers/{addressId}][%d] putRoutingSmsPhonenumberForbidden  %+v", 403, o.Payload)
}

func (o *PutRoutingSmsPhonenumberForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutRoutingSmsPhonenumberForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRoutingSmsPhonenumberNotFound creates a PutRoutingSmsPhonenumberNotFound with default headers values
func NewPutRoutingSmsPhonenumberNotFound() *PutRoutingSmsPhonenumberNotFound {
	return &PutRoutingSmsPhonenumberNotFound{}
}

/*PutRoutingSmsPhonenumberNotFound handles this case with default header values.

The requested resource was not found.
*/
type PutRoutingSmsPhonenumberNotFound struct {
	Payload *models.ErrorBody
}

func (o *PutRoutingSmsPhonenumberNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/routing/sms/phonenumbers/{addressId}][%d] putRoutingSmsPhonenumberNotFound  %+v", 404, o.Payload)
}

func (o *PutRoutingSmsPhonenumberNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutRoutingSmsPhonenumberNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRoutingSmsPhonenumberRequestTimeout creates a PutRoutingSmsPhonenumberRequestTimeout with default headers values
func NewPutRoutingSmsPhonenumberRequestTimeout() *PutRoutingSmsPhonenumberRequestTimeout {
	return &PutRoutingSmsPhonenumberRequestTimeout{}
}

/*PutRoutingSmsPhonenumberRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PutRoutingSmsPhonenumberRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PutRoutingSmsPhonenumberRequestTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/routing/sms/phonenumbers/{addressId}][%d] putRoutingSmsPhonenumberRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutRoutingSmsPhonenumberRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutRoutingSmsPhonenumberRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRoutingSmsPhonenumberConflict creates a PutRoutingSmsPhonenumberConflict with default headers values
func NewPutRoutingSmsPhonenumberConflict() *PutRoutingSmsPhonenumberConflict {
	return &PutRoutingSmsPhonenumberConflict{}
}

/*PutRoutingSmsPhonenumberConflict handles this case with default header values.

Conflict
*/
type PutRoutingSmsPhonenumberConflict struct {
	Payload *models.ErrorBody
}

func (o *PutRoutingSmsPhonenumberConflict) Error() string {
	return fmt.Sprintf("[PUT /api/v2/routing/sms/phonenumbers/{addressId}][%d] putRoutingSmsPhonenumberConflict  %+v", 409, o.Payload)
}

func (o *PutRoutingSmsPhonenumberConflict) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutRoutingSmsPhonenumberConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRoutingSmsPhonenumberRequestEntityTooLarge creates a PutRoutingSmsPhonenumberRequestEntityTooLarge with default headers values
func NewPutRoutingSmsPhonenumberRequestEntityTooLarge() *PutRoutingSmsPhonenumberRequestEntityTooLarge {
	return &PutRoutingSmsPhonenumberRequestEntityTooLarge{}
}

/*PutRoutingSmsPhonenumberRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PutRoutingSmsPhonenumberRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PutRoutingSmsPhonenumberRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /api/v2/routing/sms/phonenumbers/{addressId}][%d] putRoutingSmsPhonenumberRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutRoutingSmsPhonenumberRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutRoutingSmsPhonenumberRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRoutingSmsPhonenumberUnsupportedMediaType creates a PutRoutingSmsPhonenumberUnsupportedMediaType with default headers values
func NewPutRoutingSmsPhonenumberUnsupportedMediaType() *PutRoutingSmsPhonenumberUnsupportedMediaType {
	return &PutRoutingSmsPhonenumberUnsupportedMediaType{}
}

/*PutRoutingSmsPhonenumberUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutRoutingSmsPhonenumberUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PutRoutingSmsPhonenumberUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v2/routing/sms/phonenumbers/{addressId}][%d] putRoutingSmsPhonenumberUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutRoutingSmsPhonenumberUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutRoutingSmsPhonenumberUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRoutingSmsPhonenumberTooManyRequests creates a PutRoutingSmsPhonenumberTooManyRequests with default headers values
func NewPutRoutingSmsPhonenumberTooManyRequests() *PutRoutingSmsPhonenumberTooManyRequests {
	return &PutRoutingSmsPhonenumberTooManyRequests{}
}

/*PutRoutingSmsPhonenumberTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PutRoutingSmsPhonenumberTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PutRoutingSmsPhonenumberTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v2/routing/sms/phonenumbers/{addressId}][%d] putRoutingSmsPhonenumberTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutRoutingSmsPhonenumberTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutRoutingSmsPhonenumberTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRoutingSmsPhonenumberInternalServerError creates a PutRoutingSmsPhonenumberInternalServerError with default headers values
func NewPutRoutingSmsPhonenumberInternalServerError() *PutRoutingSmsPhonenumberInternalServerError {
	return &PutRoutingSmsPhonenumberInternalServerError{}
}

/*PutRoutingSmsPhonenumberInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutRoutingSmsPhonenumberInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PutRoutingSmsPhonenumberInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/routing/sms/phonenumbers/{addressId}][%d] putRoutingSmsPhonenumberInternalServerError  %+v", 500, o.Payload)
}

func (o *PutRoutingSmsPhonenumberInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutRoutingSmsPhonenumberInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRoutingSmsPhonenumberNotImplemented creates a PutRoutingSmsPhonenumberNotImplemented with default headers values
func NewPutRoutingSmsPhonenumberNotImplemented() *PutRoutingSmsPhonenumberNotImplemented {
	return &PutRoutingSmsPhonenumberNotImplemented{}
}

/*PutRoutingSmsPhonenumberNotImplemented handles this case with default header values.

Not Implemented
*/
type PutRoutingSmsPhonenumberNotImplemented struct {
	Payload *models.ErrorBody
}

func (o *PutRoutingSmsPhonenumberNotImplemented) Error() string {
	return fmt.Sprintf("[PUT /api/v2/routing/sms/phonenumbers/{addressId}][%d] putRoutingSmsPhonenumberNotImplemented  %+v", 501, o.Payload)
}

func (o *PutRoutingSmsPhonenumberNotImplemented) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutRoutingSmsPhonenumberNotImplemented) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRoutingSmsPhonenumberServiceUnavailable creates a PutRoutingSmsPhonenumberServiceUnavailable with default headers values
func NewPutRoutingSmsPhonenumberServiceUnavailable() *PutRoutingSmsPhonenumberServiceUnavailable {
	return &PutRoutingSmsPhonenumberServiceUnavailable{}
}

/*PutRoutingSmsPhonenumberServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutRoutingSmsPhonenumberServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PutRoutingSmsPhonenumberServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/routing/sms/phonenumbers/{addressId}][%d] putRoutingSmsPhonenumberServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutRoutingSmsPhonenumberServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutRoutingSmsPhonenumberServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRoutingSmsPhonenumberGatewayTimeout creates a PutRoutingSmsPhonenumberGatewayTimeout with default headers values
func NewPutRoutingSmsPhonenumberGatewayTimeout() *PutRoutingSmsPhonenumberGatewayTimeout {
	return &PutRoutingSmsPhonenumberGatewayTimeout{}
}

/*PutRoutingSmsPhonenumberGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PutRoutingSmsPhonenumberGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PutRoutingSmsPhonenumberGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/routing/sms/phonenumbers/{addressId}][%d] putRoutingSmsPhonenumberGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutRoutingSmsPhonenumberGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutRoutingSmsPhonenumberGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

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

// GetRoutingSmsPhonenumberReader is a Reader for the GetRoutingSmsPhonenumber structure.
type GetRoutingSmsPhonenumberReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRoutingSmsPhonenumberReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRoutingSmsPhonenumberOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRoutingSmsPhonenumberBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetRoutingSmsPhonenumberUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRoutingSmsPhonenumberForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRoutingSmsPhonenumberNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetRoutingSmsPhonenumberRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetRoutingSmsPhonenumberUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetRoutingSmsPhonenumberTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRoutingSmsPhonenumberInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetRoutingSmsPhonenumberServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetRoutingSmsPhonenumberGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRoutingSmsPhonenumberOK creates a GetRoutingSmsPhonenumberOK with default headers values
func NewGetRoutingSmsPhonenumberOK() *GetRoutingSmsPhonenumberOK {
	return &GetRoutingSmsPhonenumberOK{}
}

/*GetRoutingSmsPhonenumberOK handles this case with default header values.

successful operation
*/
type GetRoutingSmsPhonenumberOK struct {
	Payload *models.SmsPhoneNumber
}

func (o *GetRoutingSmsPhonenumberOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/sms/phonenumbers/{addressId}][%d] getRoutingSmsPhonenumberOK  %+v", 200, o.Payload)
}

func (o *GetRoutingSmsPhonenumberOK) GetPayload() *models.SmsPhoneNumber {
	return o.Payload
}

func (o *GetRoutingSmsPhonenumberOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SmsPhoneNumber)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSmsPhonenumberBadRequest creates a GetRoutingSmsPhonenumberBadRequest with default headers values
func NewGetRoutingSmsPhonenumberBadRequest() *GetRoutingSmsPhonenumberBadRequest {
	return &GetRoutingSmsPhonenumberBadRequest{}
}

/*GetRoutingSmsPhonenumberBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetRoutingSmsPhonenumberBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingSmsPhonenumberBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/sms/phonenumbers/{addressId}][%d] getRoutingSmsPhonenumberBadRequest  %+v", 400, o.Payload)
}

func (o *GetRoutingSmsPhonenumberBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSmsPhonenumberBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSmsPhonenumberUnauthorized creates a GetRoutingSmsPhonenumberUnauthorized with default headers values
func NewGetRoutingSmsPhonenumberUnauthorized() *GetRoutingSmsPhonenumberUnauthorized {
	return &GetRoutingSmsPhonenumberUnauthorized{}
}

/*GetRoutingSmsPhonenumberUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetRoutingSmsPhonenumberUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingSmsPhonenumberUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/sms/phonenumbers/{addressId}][%d] getRoutingSmsPhonenumberUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRoutingSmsPhonenumberUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSmsPhonenumberUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSmsPhonenumberForbidden creates a GetRoutingSmsPhonenumberForbidden with default headers values
func NewGetRoutingSmsPhonenumberForbidden() *GetRoutingSmsPhonenumberForbidden {
	return &GetRoutingSmsPhonenumberForbidden{}
}

/*GetRoutingSmsPhonenumberForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetRoutingSmsPhonenumberForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingSmsPhonenumberForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/sms/phonenumbers/{addressId}][%d] getRoutingSmsPhonenumberForbidden  %+v", 403, o.Payload)
}

func (o *GetRoutingSmsPhonenumberForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSmsPhonenumberForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSmsPhonenumberNotFound creates a GetRoutingSmsPhonenumberNotFound with default headers values
func NewGetRoutingSmsPhonenumberNotFound() *GetRoutingSmsPhonenumberNotFound {
	return &GetRoutingSmsPhonenumberNotFound{}
}

/*GetRoutingSmsPhonenumberNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetRoutingSmsPhonenumberNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingSmsPhonenumberNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/sms/phonenumbers/{addressId}][%d] getRoutingSmsPhonenumberNotFound  %+v", 404, o.Payload)
}

func (o *GetRoutingSmsPhonenumberNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSmsPhonenumberNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSmsPhonenumberRequestEntityTooLarge creates a GetRoutingSmsPhonenumberRequestEntityTooLarge with default headers values
func NewGetRoutingSmsPhonenumberRequestEntityTooLarge() *GetRoutingSmsPhonenumberRequestEntityTooLarge {
	return &GetRoutingSmsPhonenumberRequestEntityTooLarge{}
}

/*GetRoutingSmsPhonenumberRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetRoutingSmsPhonenumberRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingSmsPhonenumberRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/sms/phonenumbers/{addressId}][%d] getRoutingSmsPhonenumberRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetRoutingSmsPhonenumberRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSmsPhonenumberRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSmsPhonenumberUnsupportedMediaType creates a GetRoutingSmsPhonenumberUnsupportedMediaType with default headers values
func NewGetRoutingSmsPhonenumberUnsupportedMediaType() *GetRoutingSmsPhonenumberUnsupportedMediaType {
	return &GetRoutingSmsPhonenumberUnsupportedMediaType{}
}

/*GetRoutingSmsPhonenumberUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetRoutingSmsPhonenumberUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingSmsPhonenumberUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/sms/phonenumbers/{addressId}][%d] getRoutingSmsPhonenumberUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetRoutingSmsPhonenumberUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSmsPhonenumberUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSmsPhonenumberTooManyRequests creates a GetRoutingSmsPhonenumberTooManyRequests with default headers values
func NewGetRoutingSmsPhonenumberTooManyRequests() *GetRoutingSmsPhonenumberTooManyRequests {
	return &GetRoutingSmsPhonenumberTooManyRequests{}
}

/*GetRoutingSmsPhonenumberTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetRoutingSmsPhonenumberTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingSmsPhonenumberTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/sms/phonenumbers/{addressId}][%d] getRoutingSmsPhonenumberTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetRoutingSmsPhonenumberTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSmsPhonenumberTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSmsPhonenumberInternalServerError creates a GetRoutingSmsPhonenumberInternalServerError with default headers values
func NewGetRoutingSmsPhonenumberInternalServerError() *GetRoutingSmsPhonenumberInternalServerError {
	return &GetRoutingSmsPhonenumberInternalServerError{}
}

/*GetRoutingSmsPhonenumberInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetRoutingSmsPhonenumberInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingSmsPhonenumberInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/sms/phonenumbers/{addressId}][%d] getRoutingSmsPhonenumberInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRoutingSmsPhonenumberInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSmsPhonenumberInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSmsPhonenumberServiceUnavailable creates a GetRoutingSmsPhonenumberServiceUnavailable with default headers values
func NewGetRoutingSmsPhonenumberServiceUnavailable() *GetRoutingSmsPhonenumberServiceUnavailable {
	return &GetRoutingSmsPhonenumberServiceUnavailable{}
}

/*GetRoutingSmsPhonenumberServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetRoutingSmsPhonenumberServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingSmsPhonenumberServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/sms/phonenumbers/{addressId}][%d] getRoutingSmsPhonenumberServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetRoutingSmsPhonenumberServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSmsPhonenumberServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSmsPhonenumberGatewayTimeout creates a GetRoutingSmsPhonenumberGatewayTimeout with default headers values
func NewGetRoutingSmsPhonenumberGatewayTimeout() *GetRoutingSmsPhonenumberGatewayTimeout {
	return &GetRoutingSmsPhonenumberGatewayTimeout{}
}

/*GetRoutingSmsPhonenumberGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetRoutingSmsPhonenumberGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingSmsPhonenumberGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/sms/phonenumbers/{addressId}][%d] getRoutingSmsPhonenumberGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetRoutingSmsPhonenumberGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSmsPhonenumberGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

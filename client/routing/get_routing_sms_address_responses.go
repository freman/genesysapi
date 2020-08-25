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

// GetRoutingSmsAddressReader is a Reader for the GetRoutingSmsAddress structure.
type GetRoutingSmsAddressReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRoutingSmsAddressReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRoutingSmsAddressOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRoutingSmsAddressBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetRoutingSmsAddressUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRoutingSmsAddressForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRoutingSmsAddressNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetRoutingSmsAddressRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetRoutingSmsAddressUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetRoutingSmsAddressTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRoutingSmsAddressInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetRoutingSmsAddressServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetRoutingSmsAddressGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRoutingSmsAddressOK creates a GetRoutingSmsAddressOK with default headers values
func NewGetRoutingSmsAddressOK() *GetRoutingSmsAddressOK {
	return &GetRoutingSmsAddressOK{}
}

/*GetRoutingSmsAddressOK handles this case with default header values.

successful operation
*/
type GetRoutingSmsAddressOK struct {
	Payload *models.SmsAddress
}

func (o *GetRoutingSmsAddressOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/sms/addresses/{addressId}][%d] getRoutingSmsAddressOK  %+v", 200, o.Payload)
}

func (o *GetRoutingSmsAddressOK) GetPayload() *models.SmsAddress {
	return o.Payload
}

func (o *GetRoutingSmsAddressOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SmsAddress)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSmsAddressBadRequest creates a GetRoutingSmsAddressBadRequest with default headers values
func NewGetRoutingSmsAddressBadRequest() *GetRoutingSmsAddressBadRequest {
	return &GetRoutingSmsAddressBadRequest{}
}

/*GetRoutingSmsAddressBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetRoutingSmsAddressBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingSmsAddressBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/sms/addresses/{addressId}][%d] getRoutingSmsAddressBadRequest  %+v", 400, o.Payload)
}

func (o *GetRoutingSmsAddressBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSmsAddressBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSmsAddressUnauthorized creates a GetRoutingSmsAddressUnauthorized with default headers values
func NewGetRoutingSmsAddressUnauthorized() *GetRoutingSmsAddressUnauthorized {
	return &GetRoutingSmsAddressUnauthorized{}
}

/*GetRoutingSmsAddressUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetRoutingSmsAddressUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingSmsAddressUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/sms/addresses/{addressId}][%d] getRoutingSmsAddressUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRoutingSmsAddressUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSmsAddressUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSmsAddressForbidden creates a GetRoutingSmsAddressForbidden with default headers values
func NewGetRoutingSmsAddressForbidden() *GetRoutingSmsAddressForbidden {
	return &GetRoutingSmsAddressForbidden{}
}

/*GetRoutingSmsAddressForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetRoutingSmsAddressForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingSmsAddressForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/sms/addresses/{addressId}][%d] getRoutingSmsAddressForbidden  %+v", 403, o.Payload)
}

func (o *GetRoutingSmsAddressForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSmsAddressForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSmsAddressNotFound creates a GetRoutingSmsAddressNotFound with default headers values
func NewGetRoutingSmsAddressNotFound() *GetRoutingSmsAddressNotFound {
	return &GetRoutingSmsAddressNotFound{}
}

/*GetRoutingSmsAddressNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetRoutingSmsAddressNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingSmsAddressNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/sms/addresses/{addressId}][%d] getRoutingSmsAddressNotFound  %+v", 404, o.Payload)
}

func (o *GetRoutingSmsAddressNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSmsAddressNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSmsAddressRequestEntityTooLarge creates a GetRoutingSmsAddressRequestEntityTooLarge with default headers values
func NewGetRoutingSmsAddressRequestEntityTooLarge() *GetRoutingSmsAddressRequestEntityTooLarge {
	return &GetRoutingSmsAddressRequestEntityTooLarge{}
}

/*GetRoutingSmsAddressRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetRoutingSmsAddressRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingSmsAddressRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/sms/addresses/{addressId}][%d] getRoutingSmsAddressRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetRoutingSmsAddressRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSmsAddressRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSmsAddressUnsupportedMediaType creates a GetRoutingSmsAddressUnsupportedMediaType with default headers values
func NewGetRoutingSmsAddressUnsupportedMediaType() *GetRoutingSmsAddressUnsupportedMediaType {
	return &GetRoutingSmsAddressUnsupportedMediaType{}
}

/*GetRoutingSmsAddressUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetRoutingSmsAddressUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingSmsAddressUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/sms/addresses/{addressId}][%d] getRoutingSmsAddressUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetRoutingSmsAddressUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSmsAddressUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSmsAddressTooManyRequests creates a GetRoutingSmsAddressTooManyRequests with default headers values
func NewGetRoutingSmsAddressTooManyRequests() *GetRoutingSmsAddressTooManyRequests {
	return &GetRoutingSmsAddressTooManyRequests{}
}

/*GetRoutingSmsAddressTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetRoutingSmsAddressTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingSmsAddressTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/sms/addresses/{addressId}][%d] getRoutingSmsAddressTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetRoutingSmsAddressTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSmsAddressTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSmsAddressInternalServerError creates a GetRoutingSmsAddressInternalServerError with default headers values
func NewGetRoutingSmsAddressInternalServerError() *GetRoutingSmsAddressInternalServerError {
	return &GetRoutingSmsAddressInternalServerError{}
}

/*GetRoutingSmsAddressInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetRoutingSmsAddressInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingSmsAddressInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/sms/addresses/{addressId}][%d] getRoutingSmsAddressInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRoutingSmsAddressInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSmsAddressInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSmsAddressServiceUnavailable creates a GetRoutingSmsAddressServiceUnavailable with default headers values
func NewGetRoutingSmsAddressServiceUnavailable() *GetRoutingSmsAddressServiceUnavailable {
	return &GetRoutingSmsAddressServiceUnavailable{}
}

/*GetRoutingSmsAddressServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetRoutingSmsAddressServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingSmsAddressServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/sms/addresses/{addressId}][%d] getRoutingSmsAddressServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetRoutingSmsAddressServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSmsAddressServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSmsAddressGatewayTimeout creates a GetRoutingSmsAddressGatewayTimeout with default headers values
func NewGetRoutingSmsAddressGatewayTimeout() *GetRoutingSmsAddressGatewayTimeout {
	return &GetRoutingSmsAddressGatewayTimeout{}
}

/*GetRoutingSmsAddressGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetRoutingSmsAddressGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingSmsAddressGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/sms/addresses/{addressId}][%d] getRoutingSmsAddressGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetRoutingSmsAddressGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSmsAddressGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
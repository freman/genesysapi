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

// GetRoutingSmsAvailablephonenumbersReader is a Reader for the GetRoutingSmsAvailablephonenumbers structure.
type GetRoutingSmsAvailablephonenumbersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRoutingSmsAvailablephonenumbersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRoutingSmsAvailablephonenumbersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRoutingSmsAvailablephonenumbersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetRoutingSmsAvailablephonenumbersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRoutingSmsAvailablephonenumbersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRoutingSmsAvailablephonenumbersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetRoutingSmsAvailablephonenumbersRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetRoutingSmsAvailablephonenumbersUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetRoutingSmsAvailablephonenumbersTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRoutingSmsAvailablephonenumbersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetRoutingSmsAvailablephonenumbersServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetRoutingSmsAvailablephonenumbersGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRoutingSmsAvailablephonenumbersOK creates a GetRoutingSmsAvailablephonenumbersOK with default headers values
func NewGetRoutingSmsAvailablephonenumbersOK() *GetRoutingSmsAvailablephonenumbersOK {
	return &GetRoutingSmsAvailablephonenumbersOK{}
}

/*GetRoutingSmsAvailablephonenumbersOK handles this case with default header values.

successful operation
*/
type GetRoutingSmsAvailablephonenumbersOK struct {
	Payload *models.SMSAvailablePhoneNumberEntityListing
}

func (o *GetRoutingSmsAvailablephonenumbersOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/sms/availablephonenumbers][%d] getRoutingSmsAvailablephonenumbersOK  %+v", 200, o.Payload)
}

func (o *GetRoutingSmsAvailablephonenumbersOK) GetPayload() *models.SMSAvailablePhoneNumberEntityListing {
	return o.Payload
}

func (o *GetRoutingSmsAvailablephonenumbersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SMSAvailablePhoneNumberEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSmsAvailablephonenumbersBadRequest creates a GetRoutingSmsAvailablephonenumbersBadRequest with default headers values
func NewGetRoutingSmsAvailablephonenumbersBadRequest() *GetRoutingSmsAvailablephonenumbersBadRequest {
	return &GetRoutingSmsAvailablephonenumbersBadRequest{}
}

/*GetRoutingSmsAvailablephonenumbersBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetRoutingSmsAvailablephonenumbersBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingSmsAvailablephonenumbersBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/sms/availablephonenumbers][%d] getRoutingSmsAvailablephonenumbersBadRequest  %+v", 400, o.Payload)
}

func (o *GetRoutingSmsAvailablephonenumbersBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSmsAvailablephonenumbersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSmsAvailablephonenumbersUnauthorized creates a GetRoutingSmsAvailablephonenumbersUnauthorized with default headers values
func NewGetRoutingSmsAvailablephonenumbersUnauthorized() *GetRoutingSmsAvailablephonenumbersUnauthorized {
	return &GetRoutingSmsAvailablephonenumbersUnauthorized{}
}

/*GetRoutingSmsAvailablephonenumbersUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetRoutingSmsAvailablephonenumbersUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingSmsAvailablephonenumbersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/sms/availablephonenumbers][%d] getRoutingSmsAvailablephonenumbersUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRoutingSmsAvailablephonenumbersUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSmsAvailablephonenumbersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSmsAvailablephonenumbersForbidden creates a GetRoutingSmsAvailablephonenumbersForbidden with default headers values
func NewGetRoutingSmsAvailablephonenumbersForbidden() *GetRoutingSmsAvailablephonenumbersForbidden {
	return &GetRoutingSmsAvailablephonenumbersForbidden{}
}

/*GetRoutingSmsAvailablephonenumbersForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetRoutingSmsAvailablephonenumbersForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingSmsAvailablephonenumbersForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/sms/availablephonenumbers][%d] getRoutingSmsAvailablephonenumbersForbidden  %+v", 403, o.Payload)
}

func (o *GetRoutingSmsAvailablephonenumbersForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSmsAvailablephonenumbersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSmsAvailablephonenumbersNotFound creates a GetRoutingSmsAvailablephonenumbersNotFound with default headers values
func NewGetRoutingSmsAvailablephonenumbersNotFound() *GetRoutingSmsAvailablephonenumbersNotFound {
	return &GetRoutingSmsAvailablephonenumbersNotFound{}
}

/*GetRoutingSmsAvailablephonenumbersNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetRoutingSmsAvailablephonenumbersNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingSmsAvailablephonenumbersNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/sms/availablephonenumbers][%d] getRoutingSmsAvailablephonenumbersNotFound  %+v", 404, o.Payload)
}

func (o *GetRoutingSmsAvailablephonenumbersNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSmsAvailablephonenumbersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSmsAvailablephonenumbersRequestEntityTooLarge creates a GetRoutingSmsAvailablephonenumbersRequestEntityTooLarge with default headers values
func NewGetRoutingSmsAvailablephonenumbersRequestEntityTooLarge() *GetRoutingSmsAvailablephonenumbersRequestEntityTooLarge {
	return &GetRoutingSmsAvailablephonenumbersRequestEntityTooLarge{}
}

/*GetRoutingSmsAvailablephonenumbersRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetRoutingSmsAvailablephonenumbersRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingSmsAvailablephonenumbersRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/sms/availablephonenumbers][%d] getRoutingSmsAvailablephonenumbersRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetRoutingSmsAvailablephonenumbersRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSmsAvailablephonenumbersRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSmsAvailablephonenumbersUnsupportedMediaType creates a GetRoutingSmsAvailablephonenumbersUnsupportedMediaType with default headers values
func NewGetRoutingSmsAvailablephonenumbersUnsupportedMediaType() *GetRoutingSmsAvailablephonenumbersUnsupportedMediaType {
	return &GetRoutingSmsAvailablephonenumbersUnsupportedMediaType{}
}

/*GetRoutingSmsAvailablephonenumbersUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetRoutingSmsAvailablephonenumbersUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingSmsAvailablephonenumbersUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/sms/availablephonenumbers][%d] getRoutingSmsAvailablephonenumbersUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetRoutingSmsAvailablephonenumbersUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSmsAvailablephonenumbersUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSmsAvailablephonenumbersTooManyRequests creates a GetRoutingSmsAvailablephonenumbersTooManyRequests with default headers values
func NewGetRoutingSmsAvailablephonenumbersTooManyRequests() *GetRoutingSmsAvailablephonenumbersTooManyRequests {
	return &GetRoutingSmsAvailablephonenumbersTooManyRequests{}
}

/*GetRoutingSmsAvailablephonenumbersTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetRoutingSmsAvailablephonenumbersTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingSmsAvailablephonenumbersTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/sms/availablephonenumbers][%d] getRoutingSmsAvailablephonenumbersTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetRoutingSmsAvailablephonenumbersTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSmsAvailablephonenumbersTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSmsAvailablephonenumbersInternalServerError creates a GetRoutingSmsAvailablephonenumbersInternalServerError with default headers values
func NewGetRoutingSmsAvailablephonenumbersInternalServerError() *GetRoutingSmsAvailablephonenumbersInternalServerError {
	return &GetRoutingSmsAvailablephonenumbersInternalServerError{}
}

/*GetRoutingSmsAvailablephonenumbersInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetRoutingSmsAvailablephonenumbersInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingSmsAvailablephonenumbersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/sms/availablephonenumbers][%d] getRoutingSmsAvailablephonenumbersInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRoutingSmsAvailablephonenumbersInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSmsAvailablephonenumbersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSmsAvailablephonenumbersServiceUnavailable creates a GetRoutingSmsAvailablephonenumbersServiceUnavailable with default headers values
func NewGetRoutingSmsAvailablephonenumbersServiceUnavailable() *GetRoutingSmsAvailablephonenumbersServiceUnavailable {
	return &GetRoutingSmsAvailablephonenumbersServiceUnavailable{}
}

/*GetRoutingSmsAvailablephonenumbersServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetRoutingSmsAvailablephonenumbersServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingSmsAvailablephonenumbersServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/sms/availablephonenumbers][%d] getRoutingSmsAvailablephonenumbersServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetRoutingSmsAvailablephonenumbersServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSmsAvailablephonenumbersServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSmsAvailablephonenumbersGatewayTimeout creates a GetRoutingSmsAvailablephonenumbersGatewayTimeout with default headers values
func NewGetRoutingSmsAvailablephonenumbersGatewayTimeout() *GetRoutingSmsAvailablephonenumbersGatewayTimeout {
	return &GetRoutingSmsAvailablephonenumbersGatewayTimeout{}
}

/*GetRoutingSmsAvailablephonenumbersGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetRoutingSmsAvailablephonenumbersGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingSmsAvailablephonenumbersGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/sms/availablephonenumbers][%d] getRoutingSmsAvailablephonenumbersGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetRoutingSmsAvailablephonenumbersGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSmsAvailablephonenumbersGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

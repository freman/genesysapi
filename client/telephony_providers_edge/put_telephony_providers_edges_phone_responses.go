// Code generated by go-swagger; DO NOT EDIT.

package telephony_providers_edge

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PutTelephonyProvidersEdgesPhoneReader is a Reader for the PutTelephonyProvidersEdgesPhone structure.
type PutTelephonyProvidersEdgesPhoneReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutTelephonyProvidersEdgesPhoneReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutTelephonyProvidersEdgesPhoneOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutTelephonyProvidersEdgesPhoneBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutTelephonyProvidersEdgesPhoneUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutTelephonyProvidersEdgesPhoneForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutTelephonyProvidersEdgesPhoneNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPutTelephonyProvidersEdgesPhoneRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPutTelephonyProvidersEdgesPhoneUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutTelephonyProvidersEdgesPhoneTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutTelephonyProvidersEdgesPhoneInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutTelephonyProvidersEdgesPhoneServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutTelephonyProvidersEdgesPhoneGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutTelephonyProvidersEdgesPhoneOK creates a PutTelephonyProvidersEdgesPhoneOK with default headers values
func NewPutTelephonyProvidersEdgesPhoneOK() *PutTelephonyProvidersEdgesPhoneOK {
	return &PutTelephonyProvidersEdgesPhoneOK{}
}

/*PutTelephonyProvidersEdgesPhoneOK handles this case with default header values.

successful operation
*/
type PutTelephonyProvidersEdgesPhoneOK struct {
	Payload *models.Phone
}

func (o *PutTelephonyProvidersEdgesPhoneOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/phones/{phoneId}][%d] putTelephonyProvidersEdgesPhoneOK  %+v", 200, o.Payload)
}

func (o *PutTelephonyProvidersEdgesPhoneOK) GetPayload() *models.Phone {
	return o.Payload
}

func (o *PutTelephonyProvidersEdgesPhoneOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Phone)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutTelephonyProvidersEdgesPhoneBadRequest creates a PutTelephonyProvidersEdgesPhoneBadRequest with default headers values
func NewPutTelephonyProvidersEdgesPhoneBadRequest() *PutTelephonyProvidersEdgesPhoneBadRequest {
	return &PutTelephonyProvidersEdgesPhoneBadRequest{}
}

/*PutTelephonyProvidersEdgesPhoneBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutTelephonyProvidersEdgesPhoneBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PutTelephonyProvidersEdgesPhoneBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/phones/{phoneId}][%d] putTelephonyProvidersEdgesPhoneBadRequest  %+v", 400, o.Payload)
}

func (o *PutTelephonyProvidersEdgesPhoneBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutTelephonyProvidersEdgesPhoneBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutTelephonyProvidersEdgesPhoneUnauthorized creates a PutTelephonyProvidersEdgesPhoneUnauthorized with default headers values
func NewPutTelephonyProvidersEdgesPhoneUnauthorized() *PutTelephonyProvidersEdgesPhoneUnauthorized {
	return &PutTelephonyProvidersEdgesPhoneUnauthorized{}
}

/*PutTelephonyProvidersEdgesPhoneUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PutTelephonyProvidersEdgesPhoneUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PutTelephonyProvidersEdgesPhoneUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/phones/{phoneId}][%d] putTelephonyProvidersEdgesPhoneUnauthorized  %+v", 401, o.Payload)
}

func (o *PutTelephonyProvidersEdgesPhoneUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutTelephonyProvidersEdgesPhoneUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutTelephonyProvidersEdgesPhoneForbidden creates a PutTelephonyProvidersEdgesPhoneForbidden with default headers values
func NewPutTelephonyProvidersEdgesPhoneForbidden() *PutTelephonyProvidersEdgesPhoneForbidden {
	return &PutTelephonyProvidersEdgesPhoneForbidden{}
}

/*PutTelephonyProvidersEdgesPhoneForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PutTelephonyProvidersEdgesPhoneForbidden struct {
	Payload *models.ErrorBody
}

func (o *PutTelephonyProvidersEdgesPhoneForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/phones/{phoneId}][%d] putTelephonyProvidersEdgesPhoneForbidden  %+v", 403, o.Payload)
}

func (o *PutTelephonyProvidersEdgesPhoneForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutTelephonyProvidersEdgesPhoneForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutTelephonyProvidersEdgesPhoneNotFound creates a PutTelephonyProvidersEdgesPhoneNotFound with default headers values
func NewPutTelephonyProvidersEdgesPhoneNotFound() *PutTelephonyProvidersEdgesPhoneNotFound {
	return &PutTelephonyProvidersEdgesPhoneNotFound{}
}

/*PutTelephonyProvidersEdgesPhoneNotFound handles this case with default header values.

The requested resource was not found.
*/
type PutTelephonyProvidersEdgesPhoneNotFound struct {
	Payload *models.ErrorBody
}

func (o *PutTelephonyProvidersEdgesPhoneNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/phones/{phoneId}][%d] putTelephonyProvidersEdgesPhoneNotFound  %+v", 404, o.Payload)
}

func (o *PutTelephonyProvidersEdgesPhoneNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutTelephonyProvidersEdgesPhoneNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutTelephonyProvidersEdgesPhoneRequestEntityTooLarge creates a PutTelephonyProvidersEdgesPhoneRequestEntityTooLarge with default headers values
func NewPutTelephonyProvidersEdgesPhoneRequestEntityTooLarge() *PutTelephonyProvidersEdgesPhoneRequestEntityTooLarge {
	return &PutTelephonyProvidersEdgesPhoneRequestEntityTooLarge{}
}

/*PutTelephonyProvidersEdgesPhoneRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PutTelephonyProvidersEdgesPhoneRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PutTelephonyProvidersEdgesPhoneRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/phones/{phoneId}][%d] putTelephonyProvidersEdgesPhoneRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutTelephonyProvidersEdgesPhoneRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutTelephonyProvidersEdgesPhoneRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutTelephonyProvidersEdgesPhoneUnsupportedMediaType creates a PutTelephonyProvidersEdgesPhoneUnsupportedMediaType with default headers values
func NewPutTelephonyProvidersEdgesPhoneUnsupportedMediaType() *PutTelephonyProvidersEdgesPhoneUnsupportedMediaType {
	return &PutTelephonyProvidersEdgesPhoneUnsupportedMediaType{}
}

/*PutTelephonyProvidersEdgesPhoneUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutTelephonyProvidersEdgesPhoneUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PutTelephonyProvidersEdgesPhoneUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/phones/{phoneId}][%d] putTelephonyProvidersEdgesPhoneUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutTelephonyProvidersEdgesPhoneUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutTelephonyProvidersEdgesPhoneUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutTelephonyProvidersEdgesPhoneTooManyRequests creates a PutTelephonyProvidersEdgesPhoneTooManyRequests with default headers values
func NewPutTelephonyProvidersEdgesPhoneTooManyRequests() *PutTelephonyProvidersEdgesPhoneTooManyRequests {
	return &PutTelephonyProvidersEdgesPhoneTooManyRequests{}
}

/*PutTelephonyProvidersEdgesPhoneTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PutTelephonyProvidersEdgesPhoneTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PutTelephonyProvidersEdgesPhoneTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/phones/{phoneId}][%d] putTelephonyProvidersEdgesPhoneTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutTelephonyProvidersEdgesPhoneTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutTelephonyProvidersEdgesPhoneTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutTelephonyProvidersEdgesPhoneInternalServerError creates a PutTelephonyProvidersEdgesPhoneInternalServerError with default headers values
func NewPutTelephonyProvidersEdgesPhoneInternalServerError() *PutTelephonyProvidersEdgesPhoneInternalServerError {
	return &PutTelephonyProvidersEdgesPhoneInternalServerError{}
}

/*PutTelephonyProvidersEdgesPhoneInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutTelephonyProvidersEdgesPhoneInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PutTelephonyProvidersEdgesPhoneInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/phones/{phoneId}][%d] putTelephonyProvidersEdgesPhoneInternalServerError  %+v", 500, o.Payload)
}

func (o *PutTelephonyProvidersEdgesPhoneInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutTelephonyProvidersEdgesPhoneInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutTelephonyProvidersEdgesPhoneServiceUnavailable creates a PutTelephonyProvidersEdgesPhoneServiceUnavailable with default headers values
func NewPutTelephonyProvidersEdgesPhoneServiceUnavailable() *PutTelephonyProvidersEdgesPhoneServiceUnavailable {
	return &PutTelephonyProvidersEdgesPhoneServiceUnavailable{}
}

/*PutTelephonyProvidersEdgesPhoneServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutTelephonyProvidersEdgesPhoneServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PutTelephonyProvidersEdgesPhoneServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/phones/{phoneId}][%d] putTelephonyProvidersEdgesPhoneServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutTelephonyProvidersEdgesPhoneServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutTelephonyProvidersEdgesPhoneServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutTelephonyProvidersEdgesPhoneGatewayTimeout creates a PutTelephonyProvidersEdgesPhoneGatewayTimeout with default headers values
func NewPutTelephonyProvidersEdgesPhoneGatewayTimeout() *PutTelephonyProvidersEdgesPhoneGatewayTimeout {
	return &PutTelephonyProvidersEdgesPhoneGatewayTimeout{}
}

/*PutTelephonyProvidersEdgesPhoneGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PutTelephonyProvidersEdgesPhoneGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PutTelephonyProvidersEdgesPhoneGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/phones/{phoneId}][%d] putTelephonyProvidersEdgesPhoneGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutTelephonyProvidersEdgesPhoneGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutTelephonyProvidersEdgesPhoneGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
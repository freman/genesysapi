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

// GetTelephonyProvidersEdgesPhoneReader is a Reader for the GetTelephonyProvidersEdgesPhone structure.
type GetTelephonyProvidersEdgesPhoneReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTelephonyProvidersEdgesPhoneReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTelephonyProvidersEdgesPhoneOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetTelephonyProvidersEdgesPhoneBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetTelephonyProvidersEdgesPhoneUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetTelephonyProvidersEdgesPhoneForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTelephonyProvidersEdgesPhoneNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetTelephonyProvidersEdgesPhoneRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetTelephonyProvidersEdgesPhoneRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetTelephonyProvidersEdgesPhoneUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetTelephonyProvidersEdgesPhoneTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetTelephonyProvidersEdgesPhoneInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetTelephonyProvidersEdgesPhoneServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetTelephonyProvidersEdgesPhoneGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTelephonyProvidersEdgesPhoneOK creates a GetTelephonyProvidersEdgesPhoneOK with default headers values
func NewGetTelephonyProvidersEdgesPhoneOK() *GetTelephonyProvidersEdgesPhoneOK {
	return &GetTelephonyProvidersEdgesPhoneOK{}
}

/*GetTelephonyProvidersEdgesPhoneOK handles this case with default header values.

successful operation
*/
type GetTelephonyProvidersEdgesPhoneOK struct {
	Payload *models.Phone
}

func (o *GetTelephonyProvidersEdgesPhoneOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/phones/{phoneId}][%d] getTelephonyProvidersEdgesPhoneOK  %+v", 200, o.Payload)
}

func (o *GetTelephonyProvidersEdgesPhoneOK) GetPayload() *models.Phone {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesPhoneOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Phone)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesPhoneBadRequest creates a GetTelephonyProvidersEdgesPhoneBadRequest with default headers values
func NewGetTelephonyProvidersEdgesPhoneBadRequest() *GetTelephonyProvidersEdgesPhoneBadRequest {
	return &GetTelephonyProvidersEdgesPhoneBadRequest{}
}

/*GetTelephonyProvidersEdgesPhoneBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetTelephonyProvidersEdgesPhoneBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesPhoneBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/phones/{phoneId}][%d] getTelephonyProvidersEdgesPhoneBadRequest  %+v", 400, o.Payload)
}

func (o *GetTelephonyProvidersEdgesPhoneBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesPhoneBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesPhoneUnauthorized creates a GetTelephonyProvidersEdgesPhoneUnauthorized with default headers values
func NewGetTelephonyProvidersEdgesPhoneUnauthorized() *GetTelephonyProvidersEdgesPhoneUnauthorized {
	return &GetTelephonyProvidersEdgesPhoneUnauthorized{}
}

/*GetTelephonyProvidersEdgesPhoneUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetTelephonyProvidersEdgesPhoneUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesPhoneUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/phones/{phoneId}][%d] getTelephonyProvidersEdgesPhoneUnauthorized  %+v", 401, o.Payload)
}

func (o *GetTelephonyProvidersEdgesPhoneUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesPhoneUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesPhoneForbidden creates a GetTelephonyProvidersEdgesPhoneForbidden with default headers values
func NewGetTelephonyProvidersEdgesPhoneForbidden() *GetTelephonyProvidersEdgesPhoneForbidden {
	return &GetTelephonyProvidersEdgesPhoneForbidden{}
}

/*GetTelephonyProvidersEdgesPhoneForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetTelephonyProvidersEdgesPhoneForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesPhoneForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/phones/{phoneId}][%d] getTelephonyProvidersEdgesPhoneForbidden  %+v", 403, o.Payload)
}

func (o *GetTelephonyProvidersEdgesPhoneForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesPhoneForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesPhoneNotFound creates a GetTelephonyProvidersEdgesPhoneNotFound with default headers values
func NewGetTelephonyProvidersEdgesPhoneNotFound() *GetTelephonyProvidersEdgesPhoneNotFound {
	return &GetTelephonyProvidersEdgesPhoneNotFound{}
}

/*GetTelephonyProvidersEdgesPhoneNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetTelephonyProvidersEdgesPhoneNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesPhoneNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/phones/{phoneId}][%d] getTelephonyProvidersEdgesPhoneNotFound  %+v", 404, o.Payload)
}

func (o *GetTelephonyProvidersEdgesPhoneNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesPhoneNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesPhoneRequestTimeout creates a GetTelephonyProvidersEdgesPhoneRequestTimeout with default headers values
func NewGetTelephonyProvidersEdgesPhoneRequestTimeout() *GetTelephonyProvidersEdgesPhoneRequestTimeout {
	return &GetTelephonyProvidersEdgesPhoneRequestTimeout{}
}

/*GetTelephonyProvidersEdgesPhoneRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetTelephonyProvidersEdgesPhoneRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesPhoneRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/phones/{phoneId}][%d] getTelephonyProvidersEdgesPhoneRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetTelephonyProvidersEdgesPhoneRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesPhoneRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesPhoneRequestEntityTooLarge creates a GetTelephonyProvidersEdgesPhoneRequestEntityTooLarge with default headers values
func NewGetTelephonyProvidersEdgesPhoneRequestEntityTooLarge() *GetTelephonyProvidersEdgesPhoneRequestEntityTooLarge {
	return &GetTelephonyProvidersEdgesPhoneRequestEntityTooLarge{}
}

/*GetTelephonyProvidersEdgesPhoneRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetTelephonyProvidersEdgesPhoneRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesPhoneRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/phones/{phoneId}][%d] getTelephonyProvidersEdgesPhoneRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetTelephonyProvidersEdgesPhoneRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesPhoneRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesPhoneUnsupportedMediaType creates a GetTelephonyProvidersEdgesPhoneUnsupportedMediaType with default headers values
func NewGetTelephonyProvidersEdgesPhoneUnsupportedMediaType() *GetTelephonyProvidersEdgesPhoneUnsupportedMediaType {
	return &GetTelephonyProvidersEdgesPhoneUnsupportedMediaType{}
}

/*GetTelephonyProvidersEdgesPhoneUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetTelephonyProvidersEdgesPhoneUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesPhoneUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/phones/{phoneId}][%d] getTelephonyProvidersEdgesPhoneUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetTelephonyProvidersEdgesPhoneUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesPhoneUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesPhoneTooManyRequests creates a GetTelephonyProvidersEdgesPhoneTooManyRequests with default headers values
func NewGetTelephonyProvidersEdgesPhoneTooManyRequests() *GetTelephonyProvidersEdgesPhoneTooManyRequests {
	return &GetTelephonyProvidersEdgesPhoneTooManyRequests{}
}

/*GetTelephonyProvidersEdgesPhoneTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetTelephonyProvidersEdgesPhoneTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesPhoneTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/phones/{phoneId}][%d] getTelephonyProvidersEdgesPhoneTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetTelephonyProvidersEdgesPhoneTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesPhoneTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesPhoneInternalServerError creates a GetTelephonyProvidersEdgesPhoneInternalServerError with default headers values
func NewGetTelephonyProvidersEdgesPhoneInternalServerError() *GetTelephonyProvidersEdgesPhoneInternalServerError {
	return &GetTelephonyProvidersEdgesPhoneInternalServerError{}
}

/*GetTelephonyProvidersEdgesPhoneInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetTelephonyProvidersEdgesPhoneInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesPhoneInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/phones/{phoneId}][%d] getTelephonyProvidersEdgesPhoneInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTelephonyProvidersEdgesPhoneInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesPhoneInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesPhoneServiceUnavailable creates a GetTelephonyProvidersEdgesPhoneServiceUnavailable with default headers values
func NewGetTelephonyProvidersEdgesPhoneServiceUnavailable() *GetTelephonyProvidersEdgesPhoneServiceUnavailable {
	return &GetTelephonyProvidersEdgesPhoneServiceUnavailable{}
}

/*GetTelephonyProvidersEdgesPhoneServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetTelephonyProvidersEdgesPhoneServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesPhoneServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/phones/{phoneId}][%d] getTelephonyProvidersEdgesPhoneServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetTelephonyProvidersEdgesPhoneServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesPhoneServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesPhoneGatewayTimeout creates a GetTelephonyProvidersEdgesPhoneGatewayTimeout with default headers values
func NewGetTelephonyProvidersEdgesPhoneGatewayTimeout() *GetTelephonyProvidersEdgesPhoneGatewayTimeout {
	return &GetTelephonyProvidersEdgesPhoneGatewayTimeout{}
}

/*GetTelephonyProvidersEdgesPhoneGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetTelephonyProvidersEdgesPhoneGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesPhoneGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/phones/{phoneId}][%d] getTelephonyProvidersEdgesPhoneGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetTelephonyProvidersEdgesPhoneGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesPhoneGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

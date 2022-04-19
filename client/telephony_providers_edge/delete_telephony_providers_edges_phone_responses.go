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

// DeleteTelephonyProvidersEdgesPhoneReader is a Reader for the DeleteTelephonyProvidersEdgesPhone structure.
type DeleteTelephonyProvidersEdgesPhoneReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteTelephonyProvidersEdgesPhoneReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteTelephonyProvidersEdgesPhoneOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteTelephonyProvidersEdgesPhoneBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteTelephonyProvidersEdgesPhoneUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteTelephonyProvidersEdgesPhoneForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteTelephonyProvidersEdgesPhoneNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewDeleteTelephonyProvidersEdgesPhoneRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteTelephonyProvidersEdgesPhoneRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteTelephonyProvidersEdgesPhoneUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteTelephonyProvidersEdgesPhoneTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteTelephonyProvidersEdgesPhoneInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteTelephonyProvidersEdgesPhoneServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteTelephonyProvidersEdgesPhoneGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteTelephonyProvidersEdgesPhoneOK creates a DeleteTelephonyProvidersEdgesPhoneOK with default headers values
func NewDeleteTelephonyProvidersEdgesPhoneOK() *DeleteTelephonyProvidersEdgesPhoneOK {
	return &DeleteTelephonyProvidersEdgesPhoneOK{}
}

/*DeleteTelephonyProvidersEdgesPhoneOK handles this case with default header values.

Operation was successful.
*/
type DeleteTelephonyProvidersEdgesPhoneOK struct {
}

func (o *DeleteTelephonyProvidersEdgesPhoneOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/phones/{phoneId}][%d] deleteTelephonyProvidersEdgesPhoneOK ", 200)
}

func (o *DeleteTelephonyProvidersEdgesPhoneOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteTelephonyProvidersEdgesPhoneBadRequest creates a DeleteTelephonyProvidersEdgesPhoneBadRequest with default headers values
func NewDeleteTelephonyProvidersEdgesPhoneBadRequest() *DeleteTelephonyProvidersEdgesPhoneBadRequest {
	return &DeleteTelephonyProvidersEdgesPhoneBadRequest{}
}

/*DeleteTelephonyProvidersEdgesPhoneBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteTelephonyProvidersEdgesPhoneBadRequest struct {
	Payload *models.ErrorBody
}

func (o *DeleteTelephonyProvidersEdgesPhoneBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/phones/{phoneId}][%d] deleteTelephonyProvidersEdgesPhoneBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesPhoneBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTelephonyProvidersEdgesPhoneBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTelephonyProvidersEdgesPhoneUnauthorized creates a DeleteTelephonyProvidersEdgesPhoneUnauthorized with default headers values
func NewDeleteTelephonyProvidersEdgesPhoneUnauthorized() *DeleteTelephonyProvidersEdgesPhoneUnauthorized {
	return &DeleteTelephonyProvidersEdgesPhoneUnauthorized{}
}

/*DeleteTelephonyProvidersEdgesPhoneUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteTelephonyProvidersEdgesPhoneUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *DeleteTelephonyProvidersEdgesPhoneUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/phones/{phoneId}][%d] deleteTelephonyProvidersEdgesPhoneUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesPhoneUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTelephonyProvidersEdgesPhoneUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTelephonyProvidersEdgesPhoneForbidden creates a DeleteTelephonyProvidersEdgesPhoneForbidden with default headers values
func NewDeleteTelephonyProvidersEdgesPhoneForbidden() *DeleteTelephonyProvidersEdgesPhoneForbidden {
	return &DeleteTelephonyProvidersEdgesPhoneForbidden{}
}

/*DeleteTelephonyProvidersEdgesPhoneForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type DeleteTelephonyProvidersEdgesPhoneForbidden struct {
	Payload *models.ErrorBody
}

func (o *DeleteTelephonyProvidersEdgesPhoneForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/phones/{phoneId}][%d] deleteTelephonyProvidersEdgesPhoneForbidden  %+v", 403, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesPhoneForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTelephonyProvidersEdgesPhoneForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTelephonyProvidersEdgesPhoneNotFound creates a DeleteTelephonyProvidersEdgesPhoneNotFound with default headers values
func NewDeleteTelephonyProvidersEdgesPhoneNotFound() *DeleteTelephonyProvidersEdgesPhoneNotFound {
	return &DeleteTelephonyProvidersEdgesPhoneNotFound{}
}

/*DeleteTelephonyProvidersEdgesPhoneNotFound handles this case with default header values.

The requested resource was not found.
*/
type DeleteTelephonyProvidersEdgesPhoneNotFound struct {
	Payload *models.ErrorBody
}

func (o *DeleteTelephonyProvidersEdgesPhoneNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/phones/{phoneId}][%d] deleteTelephonyProvidersEdgesPhoneNotFound  %+v", 404, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesPhoneNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTelephonyProvidersEdgesPhoneNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTelephonyProvidersEdgesPhoneRequestTimeout creates a DeleteTelephonyProvidersEdgesPhoneRequestTimeout with default headers values
func NewDeleteTelephonyProvidersEdgesPhoneRequestTimeout() *DeleteTelephonyProvidersEdgesPhoneRequestTimeout {
	return &DeleteTelephonyProvidersEdgesPhoneRequestTimeout{}
}

/*DeleteTelephonyProvidersEdgesPhoneRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type DeleteTelephonyProvidersEdgesPhoneRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteTelephonyProvidersEdgesPhoneRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/phones/{phoneId}][%d] deleteTelephonyProvidersEdgesPhoneRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesPhoneRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTelephonyProvidersEdgesPhoneRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTelephonyProvidersEdgesPhoneRequestEntityTooLarge creates a DeleteTelephonyProvidersEdgesPhoneRequestEntityTooLarge with default headers values
func NewDeleteTelephonyProvidersEdgesPhoneRequestEntityTooLarge() *DeleteTelephonyProvidersEdgesPhoneRequestEntityTooLarge {
	return &DeleteTelephonyProvidersEdgesPhoneRequestEntityTooLarge{}
}

/*DeleteTelephonyProvidersEdgesPhoneRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type DeleteTelephonyProvidersEdgesPhoneRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *DeleteTelephonyProvidersEdgesPhoneRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/phones/{phoneId}][%d] deleteTelephonyProvidersEdgesPhoneRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesPhoneRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTelephonyProvidersEdgesPhoneRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTelephonyProvidersEdgesPhoneUnsupportedMediaType creates a DeleteTelephonyProvidersEdgesPhoneUnsupportedMediaType with default headers values
func NewDeleteTelephonyProvidersEdgesPhoneUnsupportedMediaType() *DeleteTelephonyProvidersEdgesPhoneUnsupportedMediaType {
	return &DeleteTelephonyProvidersEdgesPhoneUnsupportedMediaType{}
}

/*DeleteTelephonyProvidersEdgesPhoneUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteTelephonyProvidersEdgesPhoneUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *DeleteTelephonyProvidersEdgesPhoneUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/phones/{phoneId}][%d] deleteTelephonyProvidersEdgesPhoneUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesPhoneUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTelephonyProvidersEdgesPhoneUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTelephonyProvidersEdgesPhoneTooManyRequests creates a DeleteTelephonyProvidersEdgesPhoneTooManyRequests with default headers values
func NewDeleteTelephonyProvidersEdgesPhoneTooManyRequests() *DeleteTelephonyProvidersEdgesPhoneTooManyRequests {
	return &DeleteTelephonyProvidersEdgesPhoneTooManyRequests{}
}

/*DeleteTelephonyProvidersEdgesPhoneTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type DeleteTelephonyProvidersEdgesPhoneTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *DeleteTelephonyProvidersEdgesPhoneTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/phones/{phoneId}][%d] deleteTelephonyProvidersEdgesPhoneTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesPhoneTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTelephonyProvidersEdgesPhoneTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTelephonyProvidersEdgesPhoneInternalServerError creates a DeleteTelephonyProvidersEdgesPhoneInternalServerError with default headers values
func NewDeleteTelephonyProvidersEdgesPhoneInternalServerError() *DeleteTelephonyProvidersEdgesPhoneInternalServerError {
	return &DeleteTelephonyProvidersEdgesPhoneInternalServerError{}
}

/*DeleteTelephonyProvidersEdgesPhoneInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteTelephonyProvidersEdgesPhoneInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *DeleteTelephonyProvidersEdgesPhoneInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/phones/{phoneId}][%d] deleteTelephonyProvidersEdgesPhoneInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesPhoneInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTelephonyProvidersEdgesPhoneInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTelephonyProvidersEdgesPhoneServiceUnavailable creates a DeleteTelephonyProvidersEdgesPhoneServiceUnavailable with default headers values
func NewDeleteTelephonyProvidersEdgesPhoneServiceUnavailable() *DeleteTelephonyProvidersEdgesPhoneServiceUnavailable {
	return &DeleteTelephonyProvidersEdgesPhoneServiceUnavailable{}
}

/*DeleteTelephonyProvidersEdgesPhoneServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteTelephonyProvidersEdgesPhoneServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *DeleteTelephonyProvidersEdgesPhoneServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/phones/{phoneId}][%d] deleteTelephonyProvidersEdgesPhoneServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesPhoneServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTelephonyProvidersEdgesPhoneServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTelephonyProvidersEdgesPhoneGatewayTimeout creates a DeleteTelephonyProvidersEdgesPhoneGatewayTimeout with default headers values
func NewDeleteTelephonyProvidersEdgesPhoneGatewayTimeout() *DeleteTelephonyProvidersEdgesPhoneGatewayTimeout {
	return &DeleteTelephonyProvidersEdgesPhoneGatewayTimeout{}
}

/*DeleteTelephonyProvidersEdgesPhoneGatewayTimeout handles this case with default header values.

The request timed out.
*/
type DeleteTelephonyProvidersEdgesPhoneGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteTelephonyProvidersEdgesPhoneGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/phones/{phoneId}][%d] deleteTelephonyProvidersEdgesPhoneGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesPhoneGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTelephonyProvidersEdgesPhoneGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

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

// PostTelephonyProvidersEdgesPhonesRebootReader is a Reader for the PostTelephonyProvidersEdgesPhonesReboot structure.
type PostTelephonyProvidersEdgesPhonesRebootReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostTelephonyProvidersEdgesPhonesRebootReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 400:
		result := NewPostTelephonyProvidersEdgesPhonesRebootBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostTelephonyProvidersEdgesPhonesRebootUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostTelephonyProvidersEdgesPhonesRebootForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostTelephonyProvidersEdgesPhonesRebootNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostTelephonyProvidersEdgesPhonesRebootRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostTelephonyProvidersEdgesPhonesRebootRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostTelephonyProvidersEdgesPhonesRebootUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostTelephonyProvidersEdgesPhonesRebootTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostTelephonyProvidersEdgesPhonesRebootInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostTelephonyProvidersEdgesPhonesRebootServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostTelephonyProvidersEdgesPhonesRebootGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPostTelephonyProvidersEdgesPhonesRebootDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostTelephonyProvidersEdgesPhonesRebootBadRequest creates a PostTelephonyProvidersEdgesPhonesRebootBadRequest with default headers values
func NewPostTelephonyProvidersEdgesPhonesRebootBadRequest() *PostTelephonyProvidersEdgesPhonesRebootBadRequest {
	return &PostTelephonyProvidersEdgesPhonesRebootBadRequest{}
}

/*PostTelephonyProvidersEdgesPhonesRebootBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostTelephonyProvidersEdgesPhonesRebootBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgesPhonesRebootBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/phones/reboot][%d] postTelephonyProvidersEdgesPhonesRebootBadRequest  %+v", 400, o.Payload)
}

func (o *PostTelephonyProvidersEdgesPhonesRebootBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesPhonesRebootBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgesPhonesRebootUnauthorized creates a PostTelephonyProvidersEdgesPhonesRebootUnauthorized with default headers values
func NewPostTelephonyProvidersEdgesPhonesRebootUnauthorized() *PostTelephonyProvidersEdgesPhonesRebootUnauthorized {
	return &PostTelephonyProvidersEdgesPhonesRebootUnauthorized{}
}

/*PostTelephonyProvidersEdgesPhonesRebootUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostTelephonyProvidersEdgesPhonesRebootUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgesPhonesRebootUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/phones/reboot][%d] postTelephonyProvidersEdgesPhonesRebootUnauthorized  %+v", 401, o.Payload)
}

func (o *PostTelephonyProvidersEdgesPhonesRebootUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesPhonesRebootUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgesPhonesRebootForbidden creates a PostTelephonyProvidersEdgesPhonesRebootForbidden with default headers values
func NewPostTelephonyProvidersEdgesPhonesRebootForbidden() *PostTelephonyProvidersEdgesPhonesRebootForbidden {
	return &PostTelephonyProvidersEdgesPhonesRebootForbidden{}
}

/*PostTelephonyProvidersEdgesPhonesRebootForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostTelephonyProvidersEdgesPhonesRebootForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgesPhonesRebootForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/phones/reboot][%d] postTelephonyProvidersEdgesPhonesRebootForbidden  %+v", 403, o.Payload)
}

func (o *PostTelephonyProvidersEdgesPhonesRebootForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesPhonesRebootForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgesPhonesRebootNotFound creates a PostTelephonyProvidersEdgesPhonesRebootNotFound with default headers values
func NewPostTelephonyProvidersEdgesPhonesRebootNotFound() *PostTelephonyProvidersEdgesPhonesRebootNotFound {
	return &PostTelephonyProvidersEdgesPhonesRebootNotFound{}
}

/*PostTelephonyProvidersEdgesPhonesRebootNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostTelephonyProvidersEdgesPhonesRebootNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgesPhonesRebootNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/phones/reboot][%d] postTelephonyProvidersEdgesPhonesRebootNotFound  %+v", 404, o.Payload)
}

func (o *PostTelephonyProvidersEdgesPhonesRebootNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesPhonesRebootNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgesPhonesRebootRequestTimeout creates a PostTelephonyProvidersEdgesPhonesRebootRequestTimeout with default headers values
func NewPostTelephonyProvidersEdgesPhonesRebootRequestTimeout() *PostTelephonyProvidersEdgesPhonesRebootRequestTimeout {
	return &PostTelephonyProvidersEdgesPhonesRebootRequestTimeout{}
}

/*PostTelephonyProvidersEdgesPhonesRebootRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostTelephonyProvidersEdgesPhonesRebootRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgesPhonesRebootRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/phones/reboot][%d] postTelephonyProvidersEdgesPhonesRebootRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostTelephonyProvidersEdgesPhonesRebootRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesPhonesRebootRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgesPhonesRebootRequestEntityTooLarge creates a PostTelephonyProvidersEdgesPhonesRebootRequestEntityTooLarge with default headers values
func NewPostTelephonyProvidersEdgesPhonesRebootRequestEntityTooLarge() *PostTelephonyProvidersEdgesPhonesRebootRequestEntityTooLarge {
	return &PostTelephonyProvidersEdgesPhonesRebootRequestEntityTooLarge{}
}

/*PostTelephonyProvidersEdgesPhonesRebootRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostTelephonyProvidersEdgesPhonesRebootRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgesPhonesRebootRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/phones/reboot][%d] postTelephonyProvidersEdgesPhonesRebootRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostTelephonyProvidersEdgesPhonesRebootRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesPhonesRebootRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgesPhonesRebootUnsupportedMediaType creates a PostTelephonyProvidersEdgesPhonesRebootUnsupportedMediaType with default headers values
func NewPostTelephonyProvidersEdgesPhonesRebootUnsupportedMediaType() *PostTelephonyProvidersEdgesPhonesRebootUnsupportedMediaType {
	return &PostTelephonyProvidersEdgesPhonesRebootUnsupportedMediaType{}
}

/*PostTelephonyProvidersEdgesPhonesRebootUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostTelephonyProvidersEdgesPhonesRebootUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgesPhonesRebootUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/phones/reboot][%d] postTelephonyProvidersEdgesPhonesRebootUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostTelephonyProvidersEdgesPhonesRebootUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesPhonesRebootUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgesPhonesRebootTooManyRequests creates a PostTelephonyProvidersEdgesPhonesRebootTooManyRequests with default headers values
func NewPostTelephonyProvidersEdgesPhonesRebootTooManyRequests() *PostTelephonyProvidersEdgesPhonesRebootTooManyRequests {
	return &PostTelephonyProvidersEdgesPhonesRebootTooManyRequests{}
}

/*PostTelephonyProvidersEdgesPhonesRebootTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostTelephonyProvidersEdgesPhonesRebootTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgesPhonesRebootTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/phones/reboot][%d] postTelephonyProvidersEdgesPhonesRebootTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostTelephonyProvidersEdgesPhonesRebootTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesPhonesRebootTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgesPhonesRebootInternalServerError creates a PostTelephonyProvidersEdgesPhonesRebootInternalServerError with default headers values
func NewPostTelephonyProvidersEdgesPhonesRebootInternalServerError() *PostTelephonyProvidersEdgesPhonesRebootInternalServerError {
	return &PostTelephonyProvidersEdgesPhonesRebootInternalServerError{}
}

/*PostTelephonyProvidersEdgesPhonesRebootInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostTelephonyProvidersEdgesPhonesRebootInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgesPhonesRebootInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/phones/reboot][%d] postTelephonyProvidersEdgesPhonesRebootInternalServerError  %+v", 500, o.Payload)
}

func (o *PostTelephonyProvidersEdgesPhonesRebootInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesPhonesRebootInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgesPhonesRebootServiceUnavailable creates a PostTelephonyProvidersEdgesPhonesRebootServiceUnavailable with default headers values
func NewPostTelephonyProvidersEdgesPhonesRebootServiceUnavailable() *PostTelephonyProvidersEdgesPhonesRebootServiceUnavailable {
	return &PostTelephonyProvidersEdgesPhonesRebootServiceUnavailable{}
}

/*PostTelephonyProvidersEdgesPhonesRebootServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostTelephonyProvidersEdgesPhonesRebootServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgesPhonesRebootServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/phones/reboot][%d] postTelephonyProvidersEdgesPhonesRebootServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostTelephonyProvidersEdgesPhonesRebootServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesPhonesRebootServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgesPhonesRebootGatewayTimeout creates a PostTelephonyProvidersEdgesPhonesRebootGatewayTimeout with default headers values
func NewPostTelephonyProvidersEdgesPhonesRebootGatewayTimeout() *PostTelephonyProvidersEdgesPhonesRebootGatewayTimeout {
	return &PostTelephonyProvidersEdgesPhonesRebootGatewayTimeout{}
}

/*PostTelephonyProvidersEdgesPhonesRebootGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostTelephonyProvidersEdgesPhonesRebootGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgesPhonesRebootGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/phones/reboot][%d] postTelephonyProvidersEdgesPhonesRebootGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostTelephonyProvidersEdgesPhonesRebootGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesPhonesRebootGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgesPhonesRebootDefault creates a PostTelephonyProvidersEdgesPhonesRebootDefault with default headers values
func NewPostTelephonyProvidersEdgesPhonesRebootDefault(code int) *PostTelephonyProvidersEdgesPhonesRebootDefault {
	return &PostTelephonyProvidersEdgesPhonesRebootDefault{
		_statusCode: code,
	}
}

/*PostTelephonyProvidersEdgesPhonesRebootDefault handles this case with default header values.

successful operation
*/
type PostTelephonyProvidersEdgesPhonesRebootDefault struct {
	_statusCode int
}

// Code gets the status code for the post telephony providers edges phones reboot default response
func (o *PostTelephonyProvidersEdgesPhonesRebootDefault) Code() int {
	return o._statusCode
}

func (o *PostTelephonyProvidersEdgesPhonesRebootDefault) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/phones/reboot][%d] postTelephonyProvidersEdgesPhonesReboot default ", o._statusCode)
}

func (o *PostTelephonyProvidersEdgesPhonesRebootDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

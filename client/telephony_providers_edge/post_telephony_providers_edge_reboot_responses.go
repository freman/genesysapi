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

// PostTelephonyProvidersEdgeRebootReader is a Reader for the PostTelephonyProvidersEdgeReboot structure.
type PostTelephonyProvidersEdgeRebootReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostTelephonyProvidersEdgeRebootReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostTelephonyProvidersEdgeRebootOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostTelephonyProvidersEdgeRebootBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostTelephonyProvidersEdgeRebootUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostTelephonyProvidersEdgeRebootForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostTelephonyProvidersEdgeRebootNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostTelephonyProvidersEdgeRebootRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostTelephonyProvidersEdgeRebootRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostTelephonyProvidersEdgeRebootUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostTelephonyProvidersEdgeRebootTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostTelephonyProvidersEdgeRebootInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostTelephonyProvidersEdgeRebootServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostTelephonyProvidersEdgeRebootGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostTelephonyProvidersEdgeRebootOK creates a PostTelephonyProvidersEdgeRebootOK with default headers values
func NewPostTelephonyProvidersEdgeRebootOK() *PostTelephonyProvidersEdgeRebootOK {
	return &PostTelephonyProvidersEdgeRebootOK{}
}

/*PostTelephonyProvidersEdgeRebootOK handles this case with default header values.

successful operation
*/
type PostTelephonyProvidersEdgeRebootOK struct {
	Payload string
}

func (o *PostTelephonyProvidersEdgeRebootOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/reboot][%d] postTelephonyProvidersEdgeRebootOK  %+v", 200, o.Payload)
}

func (o *PostTelephonyProvidersEdgeRebootOK) GetPayload() string {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeRebootOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeRebootBadRequest creates a PostTelephonyProvidersEdgeRebootBadRequest with default headers values
func NewPostTelephonyProvidersEdgeRebootBadRequest() *PostTelephonyProvidersEdgeRebootBadRequest {
	return &PostTelephonyProvidersEdgeRebootBadRequest{}
}

/*PostTelephonyProvidersEdgeRebootBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostTelephonyProvidersEdgeRebootBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgeRebootBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/reboot][%d] postTelephonyProvidersEdgeRebootBadRequest  %+v", 400, o.Payload)
}

func (o *PostTelephonyProvidersEdgeRebootBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeRebootBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeRebootUnauthorized creates a PostTelephonyProvidersEdgeRebootUnauthorized with default headers values
func NewPostTelephonyProvidersEdgeRebootUnauthorized() *PostTelephonyProvidersEdgeRebootUnauthorized {
	return &PostTelephonyProvidersEdgeRebootUnauthorized{}
}

/*PostTelephonyProvidersEdgeRebootUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostTelephonyProvidersEdgeRebootUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgeRebootUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/reboot][%d] postTelephonyProvidersEdgeRebootUnauthorized  %+v", 401, o.Payload)
}

func (o *PostTelephonyProvidersEdgeRebootUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeRebootUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeRebootForbidden creates a PostTelephonyProvidersEdgeRebootForbidden with default headers values
func NewPostTelephonyProvidersEdgeRebootForbidden() *PostTelephonyProvidersEdgeRebootForbidden {
	return &PostTelephonyProvidersEdgeRebootForbidden{}
}

/*PostTelephonyProvidersEdgeRebootForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostTelephonyProvidersEdgeRebootForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgeRebootForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/reboot][%d] postTelephonyProvidersEdgeRebootForbidden  %+v", 403, o.Payload)
}

func (o *PostTelephonyProvidersEdgeRebootForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeRebootForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeRebootNotFound creates a PostTelephonyProvidersEdgeRebootNotFound with default headers values
func NewPostTelephonyProvidersEdgeRebootNotFound() *PostTelephonyProvidersEdgeRebootNotFound {
	return &PostTelephonyProvidersEdgeRebootNotFound{}
}

/*PostTelephonyProvidersEdgeRebootNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostTelephonyProvidersEdgeRebootNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgeRebootNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/reboot][%d] postTelephonyProvidersEdgeRebootNotFound  %+v", 404, o.Payload)
}

func (o *PostTelephonyProvidersEdgeRebootNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeRebootNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeRebootRequestTimeout creates a PostTelephonyProvidersEdgeRebootRequestTimeout with default headers values
func NewPostTelephonyProvidersEdgeRebootRequestTimeout() *PostTelephonyProvidersEdgeRebootRequestTimeout {
	return &PostTelephonyProvidersEdgeRebootRequestTimeout{}
}

/*PostTelephonyProvidersEdgeRebootRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostTelephonyProvidersEdgeRebootRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgeRebootRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/reboot][%d] postTelephonyProvidersEdgeRebootRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostTelephonyProvidersEdgeRebootRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeRebootRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeRebootRequestEntityTooLarge creates a PostTelephonyProvidersEdgeRebootRequestEntityTooLarge with default headers values
func NewPostTelephonyProvidersEdgeRebootRequestEntityTooLarge() *PostTelephonyProvidersEdgeRebootRequestEntityTooLarge {
	return &PostTelephonyProvidersEdgeRebootRequestEntityTooLarge{}
}

/*PostTelephonyProvidersEdgeRebootRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PostTelephonyProvidersEdgeRebootRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgeRebootRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/reboot][%d] postTelephonyProvidersEdgeRebootRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostTelephonyProvidersEdgeRebootRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeRebootRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeRebootUnsupportedMediaType creates a PostTelephonyProvidersEdgeRebootUnsupportedMediaType with default headers values
func NewPostTelephonyProvidersEdgeRebootUnsupportedMediaType() *PostTelephonyProvidersEdgeRebootUnsupportedMediaType {
	return &PostTelephonyProvidersEdgeRebootUnsupportedMediaType{}
}

/*PostTelephonyProvidersEdgeRebootUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostTelephonyProvidersEdgeRebootUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgeRebootUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/reboot][%d] postTelephonyProvidersEdgeRebootUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostTelephonyProvidersEdgeRebootUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeRebootUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeRebootTooManyRequests creates a PostTelephonyProvidersEdgeRebootTooManyRequests with default headers values
func NewPostTelephonyProvidersEdgeRebootTooManyRequests() *PostTelephonyProvidersEdgeRebootTooManyRequests {
	return &PostTelephonyProvidersEdgeRebootTooManyRequests{}
}

/*PostTelephonyProvidersEdgeRebootTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostTelephonyProvidersEdgeRebootTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgeRebootTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/reboot][%d] postTelephonyProvidersEdgeRebootTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostTelephonyProvidersEdgeRebootTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeRebootTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeRebootInternalServerError creates a PostTelephonyProvidersEdgeRebootInternalServerError with default headers values
func NewPostTelephonyProvidersEdgeRebootInternalServerError() *PostTelephonyProvidersEdgeRebootInternalServerError {
	return &PostTelephonyProvidersEdgeRebootInternalServerError{}
}

/*PostTelephonyProvidersEdgeRebootInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostTelephonyProvidersEdgeRebootInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgeRebootInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/reboot][%d] postTelephonyProvidersEdgeRebootInternalServerError  %+v", 500, o.Payload)
}

func (o *PostTelephonyProvidersEdgeRebootInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeRebootInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeRebootServiceUnavailable creates a PostTelephonyProvidersEdgeRebootServiceUnavailable with default headers values
func NewPostTelephonyProvidersEdgeRebootServiceUnavailable() *PostTelephonyProvidersEdgeRebootServiceUnavailable {
	return &PostTelephonyProvidersEdgeRebootServiceUnavailable{}
}

/*PostTelephonyProvidersEdgeRebootServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostTelephonyProvidersEdgeRebootServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgeRebootServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/reboot][%d] postTelephonyProvidersEdgeRebootServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostTelephonyProvidersEdgeRebootServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeRebootServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeRebootGatewayTimeout creates a PostTelephonyProvidersEdgeRebootGatewayTimeout with default headers values
func NewPostTelephonyProvidersEdgeRebootGatewayTimeout() *PostTelephonyProvidersEdgeRebootGatewayTimeout {
	return &PostTelephonyProvidersEdgeRebootGatewayTimeout{}
}

/*PostTelephonyProvidersEdgeRebootGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostTelephonyProvidersEdgeRebootGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgeRebootGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/reboot][%d] postTelephonyProvidersEdgeRebootGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostTelephonyProvidersEdgeRebootGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeRebootGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

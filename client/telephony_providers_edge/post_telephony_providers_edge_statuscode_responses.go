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

// PostTelephonyProvidersEdgeStatuscodeReader is a Reader for the PostTelephonyProvidersEdgeStatuscode structure.
type PostTelephonyProvidersEdgeStatuscodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostTelephonyProvidersEdgeStatuscodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostTelephonyProvidersEdgeStatuscodeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostTelephonyProvidersEdgeStatuscodeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostTelephonyProvidersEdgeStatuscodeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostTelephonyProvidersEdgeStatuscodeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostTelephonyProvidersEdgeStatuscodeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostTelephonyProvidersEdgeStatuscodeRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostTelephonyProvidersEdgeStatuscodeUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostTelephonyProvidersEdgeStatuscodeTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostTelephonyProvidersEdgeStatuscodeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostTelephonyProvidersEdgeStatuscodeServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostTelephonyProvidersEdgeStatuscodeGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostTelephonyProvidersEdgeStatuscodeOK creates a PostTelephonyProvidersEdgeStatuscodeOK with default headers values
func NewPostTelephonyProvidersEdgeStatuscodeOK() *PostTelephonyProvidersEdgeStatuscodeOK {
	return &PostTelephonyProvidersEdgeStatuscodeOK{}
}

/*PostTelephonyProvidersEdgeStatuscodeOK handles this case with default header values.

successful operation
*/
type PostTelephonyProvidersEdgeStatuscodeOK struct {
	Payload string
}

func (o *PostTelephonyProvidersEdgeStatuscodeOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/statuscode][%d] postTelephonyProvidersEdgeStatuscodeOK  %+v", 200, o.Payload)
}

func (o *PostTelephonyProvidersEdgeStatuscodeOK) GetPayload() string {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeStatuscodeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeStatuscodeBadRequest creates a PostTelephonyProvidersEdgeStatuscodeBadRequest with default headers values
func NewPostTelephonyProvidersEdgeStatuscodeBadRequest() *PostTelephonyProvidersEdgeStatuscodeBadRequest {
	return &PostTelephonyProvidersEdgeStatuscodeBadRequest{}
}

/*PostTelephonyProvidersEdgeStatuscodeBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostTelephonyProvidersEdgeStatuscodeBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgeStatuscodeBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/statuscode][%d] postTelephonyProvidersEdgeStatuscodeBadRequest  %+v", 400, o.Payload)
}

func (o *PostTelephonyProvidersEdgeStatuscodeBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeStatuscodeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeStatuscodeUnauthorized creates a PostTelephonyProvidersEdgeStatuscodeUnauthorized with default headers values
func NewPostTelephonyProvidersEdgeStatuscodeUnauthorized() *PostTelephonyProvidersEdgeStatuscodeUnauthorized {
	return &PostTelephonyProvidersEdgeStatuscodeUnauthorized{}
}

/*PostTelephonyProvidersEdgeStatuscodeUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostTelephonyProvidersEdgeStatuscodeUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgeStatuscodeUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/statuscode][%d] postTelephonyProvidersEdgeStatuscodeUnauthorized  %+v", 401, o.Payload)
}

func (o *PostTelephonyProvidersEdgeStatuscodeUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeStatuscodeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeStatuscodeForbidden creates a PostTelephonyProvidersEdgeStatuscodeForbidden with default headers values
func NewPostTelephonyProvidersEdgeStatuscodeForbidden() *PostTelephonyProvidersEdgeStatuscodeForbidden {
	return &PostTelephonyProvidersEdgeStatuscodeForbidden{}
}

/*PostTelephonyProvidersEdgeStatuscodeForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostTelephonyProvidersEdgeStatuscodeForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgeStatuscodeForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/statuscode][%d] postTelephonyProvidersEdgeStatuscodeForbidden  %+v", 403, o.Payload)
}

func (o *PostTelephonyProvidersEdgeStatuscodeForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeStatuscodeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeStatuscodeNotFound creates a PostTelephonyProvidersEdgeStatuscodeNotFound with default headers values
func NewPostTelephonyProvidersEdgeStatuscodeNotFound() *PostTelephonyProvidersEdgeStatuscodeNotFound {
	return &PostTelephonyProvidersEdgeStatuscodeNotFound{}
}

/*PostTelephonyProvidersEdgeStatuscodeNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostTelephonyProvidersEdgeStatuscodeNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgeStatuscodeNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/statuscode][%d] postTelephonyProvidersEdgeStatuscodeNotFound  %+v", 404, o.Payload)
}

func (o *PostTelephonyProvidersEdgeStatuscodeNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeStatuscodeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeStatuscodeRequestEntityTooLarge creates a PostTelephonyProvidersEdgeStatuscodeRequestEntityTooLarge with default headers values
func NewPostTelephonyProvidersEdgeStatuscodeRequestEntityTooLarge() *PostTelephonyProvidersEdgeStatuscodeRequestEntityTooLarge {
	return &PostTelephonyProvidersEdgeStatuscodeRequestEntityTooLarge{}
}

/*PostTelephonyProvidersEdgeStatuscodeRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostTelephonyProvidersEdgeStatuscodeRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgeStatuscodeRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/statuscode][%d] postTelephonyProvidersEdgeStatuscodeRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostTelephonyProvidersEdgeStatuscodeRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeStatuscodeRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeStatuscodeUnsupportedMediaType creates a PostTelephonyProvidersEdgeStatuscodeUnsupportedMediaType with default headers values
func NewPostTelephonyProvidersEdgeStatuscodeUnsupportedMediaType() *PostTelephonyProvidersEdgeStatuscodeUnsupportedMediaType {
	return &PostTelephonyProvidersEdgeStatuscodeUnsupportedMediaType{}
}

/*PostTelephonyProvidersEdgeStatuscodeUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostTelephonyProvidersEdgeStatuscodeUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgeStatuscodeUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/statuscode][%d] postTelephonyProvidersEdgeStatuscodeUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostTelephonyProvidersEdgeStatuscodeUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeStatuscodeUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeStatuscodeTooManyRequests creates a PostTelephonyProvidersEdgeStatuscodeTooManyRequests with default headers values
func NewPostTelephonyProvidersEdgeStatuscodeTooManyRequests() *PostTelephonyProvidersEdgeStatuscodeTooManyRequests {
	return &PostTelephonyProvidersEdgeStatuscodeTooManyRequests{}
}

/*PostTelephonyProvidersEdgeStatuscodeTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PostTelephonyProvidersEdgeStatuscodeTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgeStatuscodeTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/statuscode][%d] postTelephonyProvidersEdgeStatuscodeTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostTelephonyProvidersEdgeStatuscodeTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeStatuscodeTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeStatuscodeInternalServerError creates a PostTelephonyProvidersEdgeStatuscodeInternalServerError with default headers values
func NewPostTelephonyProvidersEdgeStatuscodeInternalServerError() *PostTelephonyProvidersEdgeStatuscodeInternalServerError {
	return &PostTelephonyProvidersEdgeStatuscodeInternalServerError{}
}

/*PostTelephonyProvidersEdgeStatuscodeInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostTelephonyProvidersEdgeStatuscodeInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgeStatuscodeInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/statuscode][%d] postTelephonyProvidersEdgeStatuscodeInternalServerError  %+v", 500, o.Payload)
}

func (o *PostTelephonyProvidersEdgeStatuscodeInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeStatuscodeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeStatuscodeServiceUnavailable creates a PostTelephonyProvidersEdgeStatuscodeServiceUnavailable with default headers values
func NewPostTelephonyProvidersEdgeStatuscodeServiceUnavailable() *PostTelephonyProvidersEdgeStatuscodeServiceUnavailable {
	return &PostTelephonyProvidersEdgeStatuscodeServiceUnavailable{}
}

/*PostTelephonyProvidersEdgeStatuscodeServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostTelephonyProvidersEdgeStatuscodeServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgeStatuscodeServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/statuscode][%d] postTelephonyProvidersEdgeStatuscodeServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostTelephonyProvidersEdgeStatuscodeServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeStatuscodeServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeStatuscodeGatewayTimeout creates a PostTelephonyProvidersEdgeStatuscodeGatewayTimeout with default headers values
func NewPostTelephonyProvidersEdgeStatuscodeGatewayTimeout() *PostTelephonyProvidersEdgeStatuscodeGatewayTimeout {
	return &PostTelephonyProvidersEdgeStatuscodeGatewayTimeout{}
}

/*PostTelephonyProvidersEdgeStatuscodeGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostTelephonyProvidersEdgeStatuscodeGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgeStatuscodeGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/statuscode][%d] postTelephonyProvidersEdgeStatuscodeGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostTelephonyProvidersEdgeStatuscodeGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeStatuscodeGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
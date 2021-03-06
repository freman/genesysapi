// Code generated by go-swagger; DO NOT EDIT.

package greetings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetGreetingMediaReader is a Reader for the GetGreetingMedia structure.
type GetGreetingMediaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGreetingMediaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGreetingMediaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetGreetingMediaBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetGreetingMediaUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetGreetingMediaForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetGreetingMediaNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetGreetingMediaRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetGreetingMediaUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetGreetingMediaTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetGreetingMediaInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetGreetingMediaServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetGreetingMediaGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetGreetingMediaOK creates a GetGreetingMediaOK with default headers values
func NewGetGreetingMediaOK() *GetGreetingMediaOK {
	return &GetGreetingMediaOK{}
}

/*GetGreetingMediaOK handles this case with default header values.

successful operation
*/
type GetGreetingMediaOK struct {
	Payload *models.GreetingMediaInfo
}

func (o *GetGreetingMediaOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/greetings/{greetingId}/media][%d] getGreetingMediaOK  %+v", 200, o.Payload)
}

func (o *GetGreetingMediaOK) GetPayload() *models.GreetingMediaInfo {
	return o.Payload
}

func (o *GetGreetingMediaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GreetingMediaInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGreetingMediaBadRequest creates a GetGreetingMediaBadRequest with default headers values
func NewGetGreetingMediaBadRequest() *GetGreetingMediaBadRequest {
	return &GetGreetingMediaBadRequest{}
}

/*GetGreetingMediaBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetGreetingMediaBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetGreetingMediaBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/greetings/{greetingId}/media][%d] getGreetingMediaBadRequest  %+v", 400, o.Payload)
}

func (o *GetGreetingMediaBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGreetingMediaBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGreetingMediaUnauthorized creates a GetGreetingMediaUnauthorized with default headers values
func NewGetGreetingMediaUnauthorized() *GetGreetingMediaUnauthorized {
	return &GetGreetingMediaUnauthorized{}
}

/*GetGreetingMediaUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetGreetingMediaUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetGreetingMediaUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/greetings/{greetingId}/media][%d] getGreetingMediaUnauthorized  %+v", 401, o.Payload)
}

func (o *GetGreetingMediaUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGreetingMediaUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGreetingMediaForbidden creates a GetGreetingMediaForbidden with default headers values
func NewGetGreetingMediaForbidden() *GetGreetingMediaForbidden {
	return &GetGreetingMediaForbidden{}
}

/*GetGreetingMediaForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetGreetingMediaForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetGreetingMediaForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/greetings/{greetingId}/media][%d] getGreetingMediaForbidden  %+v", 403, o.Payload)
}

func (o *GetGreetingMediaForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGreetingMediaForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGreetingMediaNotFound creates a GetGreetingMediaNotFound with default headers values
func NewGetGreetingMediaNotFound() *GetGreetingMediaNotFound {
	return &GetGreetingMediaNotFound{}
}

/*GetGreetingMediaNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetGreetingMediaNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetGreetingMediaNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/greetings/{greetingId}/media][%d] getGreetingMediaNotFound  %+v", 404, o.Payload)
}

func (o *GetGreetingMediaNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGreetingMediaNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGreetingMediaRequestEntityTooLarge creates a GetGreetingMediaRequestEntityTooLarge with default headers values
func NewGetGreetingMediaRequestEntityTooLarge() *GetGreetingMediaRequestEntityTooLarge {
	return &GetGreetingMediaRequestEntityTooLarge{}
}

/*GetGreetingMediaRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetGreetingMediaRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetGreetingMediaRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/greetings/{greetingId}/media][%d] getGreetingMediaRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetGreetingMediaRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGreetingMediaRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGreetingMediaUnsupportedMediaType creates a GetGreetingMediaUnsupportedMediaType with default headers values
func NewGetGreetingMediaUnsupportedMediaType() *GetGreetingMediaUnsupportedMediaType {
	return &GetGreetingMediaUnsupportedMediaType{}
}

/*GetGreetingMediaUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetGreetingMediaUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetGreetingMediaUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/greetings/{greetingId}/media][%d] getGreetingMediaUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetGreetingMediaUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGreetingMediaUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGreetingMediaTooManyRequests creates a GetGreetingMediaTooManyRequests with default headers values
func NewGetGreetingMediaTooManyRequests() *GetGreetingMediaTooManyRequests {
	return &GetGreetingMediaTooManyRequests{}
}

/*GetGreetingMediaTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetGreetingMediaTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetGreetingMediaTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/greetings/{greetingId}/media][%d] getGreetingMediaTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetGreetingMediaTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGreetingMediaTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGreetingMediaInternalServerError creates a GetGreetingMediaInternalServerError with default headers values
func NewGetGreetingMediaInternalServerError() *GetGreetingMediaInternalServerError {
	return &GetGreetingMediaInternalServerError{}
}

/*GetGreetingMediaInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetGreetingMediaInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetGreetingMediaInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/greetings/{greetingId}/media][%d] getGreetingMediaInternalServerError  %+v", 500, o.Payload)
}

func (o *GetGreetingMediaInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGreetingMediaInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGreetingMediaServiceUnavailable creates a GetGreetingMediaServiceUnavailable with default headers values
func NewGetGreetingMediaServiceUnavailable() *GetGreetingMediaServiceUnavailable {
	return &GetGreetingMediaServiceUnavailable{}
}

/*GetGreetingMediaServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetGreetingMediaServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetGreetingMediaServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/greetings/{greetingId}/media][%d] getGreetingMediaServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetGreetingMediaServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGreetingMediaServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGreetingMediaGatewayTimeout creates a GetGreetingMediaGatewayTimeout with default headers values
func NewGetGreetingMediaGatewayTimeout() *GetGreetingMediaGatewayTimeout {
	return &GetGreetingMediaGatewayTimeout{}
}

/*GetGreetingMediaGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetGreetingMediaGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetGreetingMediaGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/greetings/{greetingId}/media][%d] getGreetingMediaGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetGreetingMediaGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGreetingMediaGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

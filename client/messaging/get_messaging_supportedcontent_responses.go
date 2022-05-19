// Code generated by go-swagger; DO NOT EDIT.

package messaging

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetMessagingSupportedcontentReader is a Reader for the GetMessagingSupportedcontent structure.
type GetMessagingSupportedcontentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMessagingSupportedcontentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMessagingSupportedcontentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetMessagingSupportedcontentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetMessagingSupportedcontentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetMessagingSupportedcontentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetMessagingSupportedcontentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetMessagingSupportedcontentRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetMessagingSupportedcontentRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetMessagingSupportedcontentUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetMessagingSupportedcontentTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetMessagingSupportedcontentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetMessagingSupportedcontentServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetMessagingSupportedcontentGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetMessagingSupportedcontentOK creates a GetMessagingSupportedcontentOK with default headers values
func NewGetMessagingSupportedcontentOK() *GetMessagingSupportedcontentOK {
	return &GetMessagingSupportedcontentOK{}
}

/*GetMessagingSupportedcontentOK handles this case with default header values.

successful operation
*/
type GetMessagingSupportedcontentOK struct {
	Payload *models.SupportedContentListing
}

func (o *GetMessagingSupportedcontentOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/messaging/supportedcontent][%d] getMessagingSupportedcontentOK  %+v", 200, o.Payload)
}

func (o *GetMessagingSupportedcontentOK) GetPayload() *models.SupportedContentListing {
	return o.Payload
}

func (o *GetMessagingSupportedcontentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SupportedContentListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMessagingSupportedcontentBadRequest creates a GetMessagingSupportedcontentBadRequest with default headers values
func NewGetMessagingSupportedcontentBadRequest() *GetMessagingSupportedcontentBadRequest {
	return &GetMessagingSupportedcontentBadRequest{}
}

/*GetMessagingSupportedcontentBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetMessagingSupportedcontentBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetMessagingSupportedcontentBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/messaging/supportedcontent][%d] getMessagingSupportedcontentBadRequest  %+v", 400, o.Payload)
}

func (o *GetMessagingSupportedcontentBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetMessagingSupportedcontentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMessagingSupportedcontentUnauthorized creates a GetMessagingSupportedcontentUnauthorized with default headers values
func NewGetMessagingSupportedcontentUnauthorized() *GetMessagingSupportedcontentUnauthorized {
	return &GetMessagingSupportedcontentUnauthorized{}
}

/*GetMessagingSupportedcontentUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetMessagingSupportedcontentUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetMessagingSupportedcontentUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/messaging/supportedcontent][%d] getMessagingSupportedcontentUnauthorized  %+v", 401, o.Payload)
}

func (o *GetMessagingSupportedcontentUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetMessagingSupportedcontentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMessagingSupportedcontentForbidden creates a GetMessagingSupportedcontentForbidden with default headers values
func NewGetMessagingSupportedcontentForbidden() *GetMessagingSupportedcontentForbidden {
	return &GetMessagingSupportedcontentForbidden{}
}

/*GetMessagingSupportedcontentForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetMessagingSupportedcontentForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetMessagingSupportedcontentForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/messaging/supportedcontent][%d] getMessagingSupportedcontentForbidden  %+v", 403, o.Payload)
}

func (o *GetMessagingSupportedcontentForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetMessagingSupportedcontentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMessagingSupportedcontentNotFound creates a GetMessagingSupportedcontentNotFound with default headers values
func NewGetMessagingSupportedcontentNotFound() *GetMessagingSupportedcontentNotFound {
	return &GetMessagingSupportedcontentNotFound{}
}

/*GetMessagingSupportedcontentNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetMessagingSupportedcontentNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetMessagingSupportedcontentNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/messaging/supportedcontent][%d] getMessagingSupportedcontentNotFound  %+v", 404, o.Payload)
}

func (o *GetMessagingSupportedcontentNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetMessagingSupportedcontentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMessagingSupportedcontentRequestTimeout creates a GetMessagingSupportedcontentRequestTimeout with default headers values
func NewGetMessagingSupportedcontentRequestTimeout() *GetMessagingSupportedcontentRequestTimeout {
	return &GetMessagingSupportedcontentRequestTimeout{}
}

/*GetMessagingSupportedcontentRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetMessagingSupportedcontentRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetMessagingSupportedcontentRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/messaging/supportedcontent][%d] getMessagingSupportedcontentRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetMessagingSupportedcontentRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetMessagingSupportedcontentRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMessagingSupportedcontentRequestEntityTooLarge creates a GetMessagingSupportedcontentRequestEntityTooLarge with default headers values
func NewGetMessagingSupportedcontentRequestEntityTooLarge() *GetMessagingSupportedcontentRequestEntityTooLarge {
	return &GetMessagingSupportedcontentRequestEntityTooLarge{}
}

/*GetMessagingSupportedcontentRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetMessagingSupportedcontentRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetMessagingSupportedcontentRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/messaging/supportedcontent][%d] getMessagingSupportedcontentRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetMessagingSupportedcontentRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetMessagingSupportedcontentRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMessagingSupportedcontentUnsupportedMediaType creates a GetMessagingSupportedcontentUnsupportedMediaType with default headers values
func NewGetMessagingSupportedcontentUnsupportedMediaType() *GetMessagingSupportedcontentUnsupportedMediaType {
	return &GetMessagingSupportedcontentUnsupportedMediaType{}
}

/*GetMessagingSupportedcontentUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetMessagingSupportedcontentUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetMessagingSupportedcontentUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/messaging/supportedcontent][%d] getMessagingSupportedcontentUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetMessagingSupportedcontentUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetMessagingSupportedcontentUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMessagingSupportedcontentTooManyRequests creates a GetMessagingSupportedcontentTooManyRequests with default headers values
func NewGetMessagingSupportedcontentTooManyRequests() *GetMessagingSupportedcontentTooManyRequests {
	return &GetMessagingSupportedcontentTooManyRequests{}
}

/*GetMessagingSupportedcontentTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetMessagingSupportedcontentTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetMessagingSupportedcontentTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/messaging/supportedcontent][%d] getMessagingSupportedcontentTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetMessagingSupportedcontentTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetMessagingSupportedcontentTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMessagingSupportedcontentInternalServerError creates a GetMessagingSupportedcontentInternalServerError with default headers values
func NewGetMessagingSupportedcontentInternalServerError() *GetMessagingSupportedcontentInternalServerError {
	return &GetMessagingSupportedcontentInternalServerError{}
}

/*GetMessagingSupportedcontentInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetMessagingSupportedcontentInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetMessagingSupportedcontentInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/messaging/supportedcontent][%d] getMessagingSupportedcontentInternalServerError  %+v", 500, o.Payload)
}

func (o *GetMessagingSupportedcontentInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetMessagingSupportedcontentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMessagingSupportedcontentServiceUnavailable creates a GetMessagingSupportedcontentServiceUnavailable with default headers values
func NewGetMessagingSupportedcontentServiceUnavailable() *GetMessagingSupportedcontentServiceUnavailable {
	return &GetMessagingSupportedcontentServiceUnavailable{}
}

/*GetMessagingSupportedcontentServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetMessagingSupportedcontentServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetMessagingSupportedcontentServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/messaging/supportedcontent][%d] getMessagingSupportedcontentServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetMessagingSupportedcontentServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetMessagingSupportedcontentServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMessagingSupportedcontentGatewayTimeout creates a GetMessagingSupportedcontentGatewayTimeout with default headers values
func NewGetMessagingSupportedcontentGatewayTimeout() *GetMessagingSupportedcontentGatewayTimeout {
	return &GetMessagingSupportedcontentGatewayTimeout{}
}

/*GetMessagingSupportedcontentGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetMessagingSupportedcontentGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetMessagingSupportedcontentGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/messaging/supportedcontent][%d] getMessagingSupportedcontentGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetMessagingSupportedcontentGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetMessagingSupportedcontentGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
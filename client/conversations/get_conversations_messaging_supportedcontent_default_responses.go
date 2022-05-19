// Code generated by go-swagger; DO NOT EDIT.

package conversations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetConversationsMessagingSupportedcontentDefaultReader is a Reader for the GetConversationsMessagingSupportedcontentDefault structure.
type GetConversationsMessagingSupportedcontentDefaultReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConversationsMessagingSupportedcontentDefaultReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetConversationsMessagingSupportedcontentDefaultOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetConversationsMessagingSupportedcontentDefaultBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetConversationsMessagingSupportedcontentDefaultUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetConversationsMessagingSupportedcontentDefaultForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetConversationsMessagingSupportedcontentDefaultNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetConversationsMessagingSupportedcontentDefaultRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetConversationsMessagingSupportedcontentDefaultRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetConversationsMessagingSupportedcontentDefaultUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetConversationsMessagingSupportedcontentDefaultTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetConversationsMessagingSupportedcontentDefaultInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetConversationsMessagingSupportedcontentDefaultServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetConversationsMessagingSupportedcontentDefaultGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetConversationsMessagingSupportedcontentDefaultOK creates a GetConversationsMessagingSupportedcontentDefaultOK with default headers values
func NewGetConversationsMessagingSupportedcontentDefaultOK() *GetConversationsMessagingSupportedcontentDefaultOK {
	return &GetConversationsMessagingSupportedcontentDefaultOK{}
}

/*GetConversationsMessagingSupportedcontentDefaultOK handles this case with default header values.

successful operation
*/
type GetConversationsMessagingSupportedcontentDefaultOK struct {
	Payload *models.SupportedContent
}

func (o *GetConversationsMessagingSupportedcontentDefaultOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/supportedcontent/default][%d] getConversationsMessagingSupportedcontentDefaultOK  %+v", 200, o.Payload)
}

func (o *GetConversationsMessagingSupportedcontentDefaultOK) GetPayload() *models.SupportedContent {
	return o.Payload
}

func (o *GetConversationsMessagingSupportedcontentDefaultOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SupportedContent)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingSupportedcontentDefaultBadRequest creates a GetConversationsMessagingSupportedcontentDefaultBadRequest with default headers values
func NewGetConversationsMessagingSupportedcontentDefaultBadRequest() *GetConversationsMessagingSupportedcontentDefaultBadRequest {
	return &GetConversationsMessagingSupportedcontentDefaultBadRequest{}
}

/*GetConversationsMessagingSupportedcontentDefaultBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetConversationsMessagingSupportedcontentDefaultBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessagingSupportedcontentDefaultBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/supportedcontent/default][%d] getConversationsMessagingSupportedcontentDefaultBadRequest  %+v", 400, o.Payload)
}

func (o *GetConversationsMessagingSupportedcontentDefaultBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingSupportedcontentDefaultBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingSupportedcontentDefaultUnauthorized creates a GetConversationsMessagingSupportedcontentDefaultUnauthorized with default headers values
func NewGetConversationsMessagingSupportedcontentDefaultUnauthorized() *GetConversationsMessagingSupportedcontentDefaultUnauthorized {
	return &GetConversationsMessagingSupportedcontentDefaultUnauthorized{}
}

/*GetConversationsMessagingSupportedcontentDefaultUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetConversationsMessagingSupportedcontentDefaultUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessagingSupportedcontentDefaultUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/supportedcontent/default][%d] getConversationsMessagingSupportedcontentDefaultUnauthorized  %+v", 401, o.Payload)
}

func (o *GetConversationsMessagingSupportedcontentDefaultUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingSupportedcontentDefaultUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingSupportedcontentDefaultForbidden creates a GetConversationsMessagingSupportedcontentDefaultForbidden with default headers values
func NewGetConversationsMessagingSupportedcontentDefaultForbidden() *GetConversationsMessagingSupportedcontentDefaultForbidden {
	return &GetConversationsMessagingSupportedcontentDefaultForbidden{}
}

/*GetConversationsMessagingSupportedcontentDefaultForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetConversationsMessagingSupportedcontentDefaultForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessagingSupportedcontentDefaultForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/supportedcontent/default][%d] getConversationsMessagingSupportedcontentDefaultForbidden  %+v", 403, o.Payload)
}

func (o *GetConversationsMessagingSupportedcontentDefaultForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingSupportedcontentDefaultForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingSupportedcontentDefaultNotFound creates a GetConversationsMessagingSupportedcontentDefaultNotFound with default headers values
func NewGetConversationsMessagingSupportedcontentDefaultNotFound() *GetConversationsMessagingSupportedcontentDefaultNotFound {
	return &GetConversationsMessagingSupportedcontentDefaultNotFound{}
}

/*GetConversationsMessagingSupportedcontentDefaultNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetConversationsMessagingSupportedcontentDefaultNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessagingSupportedcontentDefaultNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/supportedcontent/default][%d] getConversationsMessagingSupportedcontentDefaultNotFound  %+v", 404, o.Payload)
}

func (o *GetConversationsMessagingSupportedcontentDefaultNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingSupportedcontentDefaultNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingSupportedcontentDefaultRequestTimeout creates a GetConversationsMessagingSupportedcontentDefaultRequestTimeout with default headers values
func NewGetConversationsMessagingSupportedcontentDefaultRequestTimeout() *GetConversationsMessagingSupportedcontentDefaultRequestTimeout {
	return &GetConversationsMessagingSupportedcontentDefaultRequestTimeout{}
}

/*GetConversationsMessagingSupportedcontentDefaultRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetConversationsMessagingSupportedcontentDefaultRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessagingSupportedcontentDefaultRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/supportedcontent/default][%d] getConversationsMessagingSupportedcontentDefaultRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetConversationsMessagingSupportedcontentDefaultRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingSupportedcontentDefaultRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingSupportedcontentDefaultRequestEntityTooLarge creates a GetConversationsMessagingSupportedcontentDefaultRequestEntityTooLarge with default headers values
func NewGetConversationsMessagingSupportedcontentDefaultRequestEntityTooLarge() *GetConversationsMessagingSupportedcontentDefaultRequestEntityTooLarge {
	return &GetConversationsMessagingSupportedcontentDefaultRequestEntityTooLarge{}
}

/*GetConversationsMessagingSupportedcontentDefaultRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetConversationsMessagingSupportedcontentDefaultRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessagingSupportedcontentDefaultRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/supportedcontent/default][%d] getConversationsMessagingSupportedcontentDefaultRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetConversationsMessagingSupportedcontentDefaultRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingSupportedcontentDefaultRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingSupportedcontentDefaultUnsupportedMediaType creates a GetConversationsMessagingSupportedcontentDefaultUnsupportedMediaType with default headers values
func NewGetConversationsMessagingSupportedcontentDefaultUnsupportedMediaType() *GetConversationsMessagingSupportedcontentDefaultUnsupportedMediaType {
	return &GetConversationsMessagingSupportedcontentDefaultUnsupportedMediaType{}
}

/*GetConversationsMessagingSupportedcontentDefaultUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetConversationsMessagingSupportedcontentDefaultUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessagingSupportedcontentDefaultUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/supportedcontent/default][%d] getConversationsMessagingSupportedcontentDefaultUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetConversationsMessagingSupportedcontentDefaultUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingSupportedcontentDefaultUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingSupportedcontentDefaultTooManyRequests creates a GetConversationsMessagingSupportedcontentDefaultTooManyRequests with default headers values
func NewGetConversationsMessagingSupportedcontentDefaultTooManyRequests() *GetConversationsMessagingSupportedcontentDefaultTooManyRequests {
	return &GetConversationsMessagingSupportedcontentDefaultTooManyRequests{}
}

/*GetConversationsMessagingSupportedcontentDefaultTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetConversationsMessagingSupportedcontentDefaultTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessagingSupportedcontentDefaultTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/supportedcontent/default][%d] getConversationsMessagingSupportedcontentDefaultTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetConversationsMessagingSupportedcontentDefaultTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingSupportedcontentDefaultTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingSupportedcontentDefaultInternalServerError creates a GetConversationsMessagingSupportedcontentDefaultInternalServerError with default headers values
func NewGetConversationsMessagingSupportedcontentDefaultInternalServerError() *GetConversationsMessagingSupportedcontentDefaultInternalServerError {
	return &GetConversationsMessagingSupportedcontentDefaultInternalServerError{}
}

/*GetConversationsMessagingSupportedcontentDefaultInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetConversationsMessagingSupportedcontentDefaultInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessagingSupportedcontentDefaultInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/supportedcontent/default][%d] getConversationsMessagingSupportedcontentDefaultInternalServerError  %+v", 500, o.Payload)
}

func (o *GetConversationsMessagingSupportedcontentDefaultInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingSupportedcontentDefaultInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingSupportedcontentDefaultServiceUnavailable creates a GetConversationsMessagingSupportedcontentDefaultServiceUnavailable with default headers values
func NewGetConversationsMessagingSupportedcontentDefaultServiceUnavailable() *GetConversationsMessagingSupportedcontentDefaultServiceUnavailable {
	return &GetConversationsMessagingSupportedcontentDefaultServiceUnavailable{}
}

/*GetConversationsMessagingSupportedcontentDefaultServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetConversationsMessagingSupportedcontentDefaultServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessagingSupportedcontentDefaultServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/supportedcontent/default][%d] getConversationsMessagingSupportedcontentDefaultServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetConversationsMessagingSupportedcontentDefaultServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingSupportedcontentDefaultServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingSupportedcontentDefaultGatewayTimeout creates a GetConversationsMessagingSupportedcontentDefaultGatewayTimeout with default headers values
func NewGetConversationsMessagingSupportedcontentDefaultGatewayTimeout() *GetConversationsMessagingSupportedcontentDefaultGatewayTimeout {
	return &GetConversationsMessagingSupportedcontentDefaultGatewayTimeout{}
}

/*GetConversationsMessagingSupportedcontentDefaultGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetConversationsMessagingSupportedcontentDefaultGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessagingSupportedcontentDefaultGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/supportedcontent/default][%d] getConversationsMessagingSupportedcontentDefaultGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetConversationsMessagingSupportedcontentDefaultGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingSupportedcontentDefaultGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
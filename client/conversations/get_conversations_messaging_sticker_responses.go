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

// GetConversationsMessagingStickerReader is a Reader for the GetConversationsMessagingSticker structure.
type GetConversationsMessagingStickerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConversationsMessagingStickerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetConversationsMessagingStickerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetConversationsMessagingStickerBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetConversationsMessagingStickerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetConversationsMessagingStickerForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetConversationsMessagingStickerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetConversationsMessagingStickerRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetConversationsMessagingStickerRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetConversationsMessagingStickerUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetConversationsMessagingStickerTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetConversationsMessagingStickerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetConversationsMessagingStickerServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetConversationsMessagingStickerGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetConversationsMessagingStickerOK creates a GetConversationsMessagingStickerOK with default headers values
func NewGetConversationsMessagingStickerOK() *GetConversationsMessagingStickerOK {
	return &GetConversationsMessagingStickerOK{}
}

/*GetConversationsMessagingStickerOK handles this case with default header values.

successful operation
*/
type GetConversationsMessagingStickerOK struct {
	Payload *models.MessagingStickerEntityListing
}

func (o *GetConversationsMessagingStickerOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/stickers/{messengerType}][%d] getConversationsMessagingStickerOK  %+v", 200, o.Payload)
}

func (o *GetConversationsMessagingStickerOK) GetPayload() *models.MessagingStickerEntityListing {
	return o.Payload
}

func (o *GetConversationsMessagingStickerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MessagingStickerEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingStickerBadRequest creates a GetConversationsMessagingStickerBadRequest with default headers values
func NewGetConversationsMessagingStickerBadRequest() *GetConversationsMessagingStickerBadRequest {
	return &GetConversationsMessagingStickerBadRequest{}
}

/*GetConversationsMessagingStickerBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetConversationsMessagingStickerBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessagingStickerBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/stickers/{messengerType}][%d] getConversationsMessagingStickerBadRequest  %+v", 400, o.Payload)
}

func (o *GetConversationsMessagingStickerBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingStickerBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingStickerUnauthorized creates a GetConversationsMessagingStickerUnauthorized with default headers values
func NewGetConversationsMessagingStickerUnauthorized() *GetConversationsMessagingStickerUnauthorized {
	return &GetConversationsMessagingStickerUnauthorized{}
}

/*GetConversationsMessagingStickerUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetConversationsMessagingStickerUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessagingStickerUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/stickers/{messengerType}][%d] getConversationsMessagingStickerUnauthorized  %+v", 401, o.Payload)
}

func (o *GetConversationsMessagingStickerUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingStickerUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingStickerForbidden creates a GetConversationsMessagingStickerForbidden with default headers values
func NewGetConversationsMessagingStickerForbidden() *GetConversationsMessagingStickerForbidden {
	return &GetConversationsMessagingStickerForbidden{}
}

/*GetConversationsMessagingStickerForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetConversationsMessagingStickerForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessagingStickerForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/stickers/{messengerType}][%d] getConversationsMessagingStickerForbidden  %+v", 403, o.Payload)
}

func (o *GetConversationsMessagingStickerForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingStickerForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingStickerNotFound creates a GetConversationsMessagingStickerNotFound with default headers values
func NewGetConversationsMessagingStickerNotFound() *GetConversationsMessagingStickerNotFound {
	return &GetConversationsMessagingStickerNotFound{}
}

/*GetConversationsMessagingStickerNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetConversationsMessagingStickerNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessagingStickerNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/stickers/{messengerType}][%d] getConversationsMessagingStickerNotFound  %+v", 404, o.Payload)
}

func (o *GetConversationsMessagingStickerNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingStickerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingStickerRequestTimeout creates a GetConversationsMessagingStickerRequestTimeout with default headers values
func NewGetConversationsMessagingStickerRequestTimeout() *GetConversationsMessagingStickerRequestTimeout {
	return &GetConversationsMessagingStickerRequestTimeout{}
}

/*GetConversationsMessagingStickerRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetConversationsMessagingStickerRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessagingStickerRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/stickers/{messengerType}][%d] getConversationsMessagingStickerRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetConversationsMessagingStickerRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingStickerRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingStickerRequestEntityTooLarge creates a GetConversationsMessagingStickerRequestEntityTooLarge with default headers values
func NewGetConversationsMessagingStickerRequestEntityTooLarge() *GetConversationsMessagingStickerRequestEntityTooLarge {
	return &GetConversationsMessagingStickerRequestEntityTooLarge{}
}

/*GetConversationsMessagingStickerRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetConversationsMessagingStickerRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessagingStickerRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/stickers/{messengerType}][%d] getConversationsMessagingStickerRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetConversationsMessagingStickerRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingStickerRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingStickerUnsupportedMediaType creates a GetConversationsMessagingStickerUnsupportedMediaType with default headers values
func NewGetConversationsMessagingStickerUnsupportedMediaType() *GetConversationsMessagingStickerUnsupportedMediaType {
	return &GetConversationsMessagingStickerUnsupportedMediaType{}
}

/*GetConversationsMessagingStickerUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetConversationsMessagingStickerUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessagingStickerUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/stickers/{messengerType}][%d] getConversationsMessagingStickerUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetConversationsMessagingStickerUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingStickerUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingStickerTooManyRequests creates a GetConversationsMessagingStickerTooManyRequests with default headers values
func NewGetConversationsMessagingStickerTooManyRequests() *GetConversationsMessagingStickerTooManyRequests {
	return &GetConversationsMessagingStickerTooManyRequests{}
}

/*GetConversationsMessagingStickerTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetConversationsMessagingStickerTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessagingStickerTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/stickers/{messengerType}][%d] getConversationsMessagingStickerTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetConversationsMessagingStickerTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingStickerTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingStickerInternalServerError creates a GetConversationsMessagingStickerInternalServerError with default headers values
func NewGetConversationsMessagingStickerInternalServerError() *GetConversationsMessagingStickerInternalServerError {
	return &GetConversationsMessagingStickerInternalServerError{}
}

/*GetConversationsMessagingStickerInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetConversationsMessagingStickerInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessagingStickerInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/stickers/{messengerType}][%d] getConversationsMessagingStickerInternalServerError  %+v", 500, o.Payload)
}

func (o *GetConversationsMessagingStickerInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingStickerInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingStickerServiceUnavailable creates a GetConversationsMessagingStickerServiceUnavailable with default headers values
func NewGetConversationsMessagingStickerServiceUnavailable() *GetConversationsMessagingStickerServiceUnavailable {
	return &GetConversationsMessagingStickerServiceUnavailable{}
}

/*GetConversationsMessagingStickerServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetConversationsMessagingStickerServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessagingStickerServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/stickers/{messengerType}][%d] getConversationsMessagingStickerServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetConversationsMessagingStickerServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingStickerServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingStickerGatewayTimeout creates a GetConversationsMessagingStickerGatewayTimeout with default headers values
func NewGetConversationsMessagingStickerGatewayTimeout() *GetConversationsMessagingStickerGatewayTimeout {
	return &GetConversationsMessagingStickerGatewayTimeout{}
}

/*GetConversationsMessagingStickerGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetConversationsMessagingStickerGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessagingStickerGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/stickers/{messengerType}][%d] getConversationsMessagingStickerGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetConversationsMessagingStickerGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingStickerGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

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

// GetConversationsChatReader is a Reader for the GetConversationsChat structure.
type GetConversationsChatReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConversationsChatReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetConversationsChatOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetConversationsChatBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetConversationsChatUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetConversationsChatForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetConversationsChatNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetConversationsChatRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetConversationsChatUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetConversationsChatTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetConversationsChatInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetConversationsChatServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetConversationsChatGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetConversationsChatOK creates a GetConversationsChatOK with default headers values
func NewGetConversationsChatOK() *GetConversationsChatOK {
	return &GetConversationsChatOK{}
}

/*GetConversationsChatOK handles this case with default header values.

successful operation
*/
type GetConversationsChatOK struct {
	Payload *models.ChatConversation
}

func (o *GetConversationsChatOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}][%d] getConversationsChatOK  %+v", 200, o.Payload)
}

func (o *GetConversationsChatOK) GetPayload() *models.ChatConversation {
	return o.Payload
}

func (o *GetConversationsChatOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ChatConversation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsChatBadRequest creates a GetConversationsChatBadRequest with default headers values
func NewGetConversationsChatBadRequest() *GetConversationsChatBadRequest {
	return &GetConversationsChatBadRequest{}
}

/*GetConversationsChatBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetConversationsChatBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsChatBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}][%d] getConversationsChatBadRequest  %+v", 400, o.Payload)
}

func (o *GetConversationsChatBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsChatBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsChatUnauthorized creates a GetConversationsChatUnauthorized with default headers values
func NewGetConversationsChatUnauthorized() *GetConversationsChatUnauthorized {
	return &GetConversationsChatUnauthorized{}
}

/*GetConversationsChatUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetConversationsChatUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsChatUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}][%d] getConversationsChatUnauthorized  %+v", 401, o.Payload)
}

func (o *GetConversationsChatUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsChatUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsChatForbidden creates a GetConversationsChatForbidden with default headers values
func NewGetConversationsChatForbidden() *GetConversationsChatForbidden {
	return &GetConversationsChatForbidden{}
}

/*GetConversationsChatForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetConversationsChatForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsChatForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}][%d] getConversationsChatForbidden  %+v", 403, o.Payload)
}

func (o *GetConversationsChatForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsChatForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsChatNotFound creates a GetConversationsChatNotFound with default headers values
func NewGetConversationsChatNotFound() *GetConversationsChatNotFound {
	return &GetConversationsChatNotFound{}
}

/*GetConversationsChatNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetConversationsChatNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsChatNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}][%d] getConversationsChatNotFound  %+v", 404, o.Payload)
}

func (o *GetConversationsChatNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsChatNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsChatRequestEntityTooLarge creates a GetConversationsChatRequestEntityTooLarge with default headers values
func NewGetConversationsChatRequestEntityTooLarge() *GetConversationsChatRequestEntityTooLarge {
	return &GetConversationsChatRequestEntityTooLarge{}
}

/*GetConversationsChatRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetConversationsChatRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsChatRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}][%d] getConversationsChatRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetConversationsChatRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsChatRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsChatUnsupportedMediaType creates a GetConversationsChatUnsupportedMediaType with default headers values
func NewGetConversationsChatUnsupportedMediaType() *GetConversationsChatUnsupportedMediaType {
	return &GetConversationsChatUnsupportedMediaType{}
}

/*GetConversationsChatUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetConversationsChatUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsChatUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}][%d] getConversationsChatUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetConversationsChatUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsChatUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsChatTooManyRequests creates a GetConversationsChatTooManyRequests with default headers values
func NewGetConversationsChatTooManyRequests() *GetConversationsChatTooManyRequests {
	return &GetConversationsChatTooManyRequests{}
}

/*GetConversationsChatTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetConversationsChatTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsChatTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}][%d] getConversationsChatTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetConversationsChatTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsChatTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsChatInternalServerError creates a GetConversationsChatInternalServerError with default headers values
func NewGetConversationsChatInternalServerError() *GetConversationsChatInternalServerError {
	return &GetConversationsChatInternalServerError{}
}

/*GetConversationsChatInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetConversationsChatInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsChatInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}][%d] getConversationsChatInternalServerError  %+v", 500, o.Payload)
}

func (o *GetConversationsChatInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsChatInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsChatServiceUnavailable creates a GetConversationsChatServiceUnavailable with default headers values
func NewGetConversationsChatServiceUnavailable() *GetConversationsChatServiceUnavailable {
	return &GetConversationsChatServiceUnavailable{}
}

/*GetConversationsChatServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetConversationsChatServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsChatServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}][%d] getConversationsChatServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetConversationsChatServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsChatServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsChatGatewayTimeout creates a GetConversationsChatGatewayTimeout with default headers values
func NewGetConversationsChatGatewayTimeout() *GetConversationsChatGatewayTimeout {
	return &GetConversationsChatGatewayTimeout{}
}

/*GetConversationsChatGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetConversationsChatGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsChatGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}][%d] getConversationsChatGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetConversationsChatGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsChatGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
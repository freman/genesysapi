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

// GetConversationsChatMessagesReader is a Reader for the GetConversationsChatMessages structure.
type GetConversationsChatMessagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConversationsChatMessagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetConversationsChatMessagesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetConversationsChatMessagesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetConversationsChatMessagesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetConversationsChatMessagesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetConversationsChatMessagesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetConversationsChatMessagesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetConversationsChatMessagesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetConversationsChatMessagesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetConversationsChatMessagesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetConversationsChatMessagesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetConversationsChatMessagesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetConversationsChatMessagesOK creates a GetConversationsChatMessagesOK with default headers values
func NewGetConversationsChatMessagesOK() *GetConversationsChatMessagesOK {
	return &GetConversationsChatMessagesOK{}
}

/*GetConversationsChatMessagesOK handles this case with default header values.

successful operation
*/
type GetConversationsChatMessagesOK struct {
	Payload *models.WebChatMessageEntityList
}

func (o *GetConversationsChatMessagesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}/messages][%d] getConversationsChatMessagesOK  %+v", 200, o.Payload)
}

func (o *GetConversationsChatMessagesOK) GetPayload() *models.WebChatMessageEntityList {
	return o.Payload
}

func (o *GetConversationsChatMessagesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WebChatMessageEntityList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsChatMessagesBadRequest creates a GetConversationsChatMessagesBadRequest with default headers values
func NewGetConversationsChatMessagesBadRequest() *GetConversationsChatMessagesBadRequest {
	return &GetConversationsChatMessagesBadRequest{}
}

/*GetConversationsChatMessagesBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetConversationsChatMessagesBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsChatMessagesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}/messages][%d] getConversationsChatMessagesBadRequest  %+v", 400, o.Payload)
}

func (o *GetConversationsChatMessagesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsChatMessagesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsChatMessagesUnauthorized creates a GetConversationsChatMessagesUnauthorized with default headers values
func NewGetConversationsChatMessagesUnauthorized() *GetConversationsChatMessagesUnauthorized {
	return &GetConversationsChatMessagesUnauthorized{}
}

/*GetConversationsChatMessagesUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetConversationsChatMessagesUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsChatMessagesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}/messages][%d] getConversationsChatMessagesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetConversationsChatMessagesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsChatMessagesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsChatMessagesForbidden creates a GetConversationsChatMessagesForbidden with default headers values
func NewGetConversationsChatMessagesForbidden() *GetConversationsChatMessagesForbidden {
	return &GetConversationsChatMessagesForbidden{}
}

/*GetConversationsChatMessagesForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetConversationsChatMessagesForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsChatMessagesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}/messages][%d] getConversationsChatMessagesForbidden  %+v", 403, o.Payload)
}

func (o *GetConversationsChatMessagesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsChatMessagesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsChatMessagesNotFound creates a GetConversationsChatMessagesNotFound with default headers values
func NewGetConversationsChatMessagesNotFound() *GetConversationsChatMessagesNotFound {
	return &GetConversationsChatMessagesNotFound{}
}

/*GetConversationsChatMessagesNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetConversationsChatMessagesNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsChatMessagesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}/messages][%d] getConversationsChatMessagesNotFound  %+v", 404, o.Payload)
}

func (o *GetConversationsChatMessagesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsChatMessagesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsChatMessagesRequestEntityTooLarge creates a GetConversationsChatMessagesRequestEntityTooLarge with default headers values
func NewGetConversationsChatMessagesRequestEntityTooLarge() *GetConversationsChatMessagesRequestEntityTooLarge {
	return &GetConversationsChatMessagesRequestEntityTooLarge{}
}

/*GetConversationsChatMessagesRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetConversationsChatMessagesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsChatMessagesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}/messages][%d] getConversationsChatMessagesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetConversationsChatMessagesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsChatMessagesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsChatMessagesUnsupportedMediaType creates a GetConversationsChatMessagesUnsupportedMediaType with default headers values
func NewGetConversationsChatMessagesUnsupportedMediaType() *GetConversationsChatMessagesUnsupportedMediaType {
	return &GetConversationsChatMessagesUnsupportedMediaType{}
}

/*GetConversationsChatMessagesUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetConversationsChatMessagesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsChatMessagesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}/messages][%d] getConversationsChatMessagesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetConversationsChatMessagesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsChatMessagesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsChatMessagesTooManyRequests creates a GetConversationsChatMessagesTooManyRequests with default headers values
func NewGetConversationsChatMessagesTooManyRequests() *GetConversationsChatMessagesTooManyRequests {
	return &GetConversationsChatMessagesTooManyRequests{}
}

/*GetConversationsChatMessagesTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetConversationsChatMessagesTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsChatMessagesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}/messages][%d] getConversationsChatMessagesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetConversationsChatMessagesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsChatMessagesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsChatMessagesInternalServerError creates a GetConversationsChatMessagesInternalServerError with default headers values
func NewGetConversationsChatMessagesInternalServerError() *GetConversationsChatMessagesInternalServerError {
	return &GetConversationsChatMessagesInternalServerError{}
}

/*GetConversationsChatMessagesInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetConversationsChatMessagesInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsChatMessagesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}/messages][%d] getConversationsChatMessagesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetConversationsChatMessagesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsChatMessagesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsChatMessagesServiceUnavailable creates a GetConversationsChatMessagesServiceUnavailable with default headers values
func NewGetConversationsChatMessagesServiceUnavailable() *GetConversationsChatMessagesServiceUnavailable {
	return &GetConversationsChatMessagesServiceUnavailable{}
}

/*GetConversationsChatMessagesServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetConversationsChatMessagesServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsChatMessagesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}/messages][%d] getConversationsChatMessagesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetConversationsChatMessagesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsChatMessagesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsChatMessagesGatewayTimeout creates a GetConversationsChatMessagesGatewayTimeout with default headers values
func NewGetConversationsChatMessagesGatewayTimeout() *GetConversationsChatMessagesGatewayTimeout {
	return &GetConversationsChatMessagesGatewayTimeout{}
}

/*GetConversationsChatMessagesGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetConversationsChatMessagesGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsChatMessagesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}/messages][%d] getConversationsChatMessagesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetConversationsChatMessagesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsChatMessagesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
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

// GetConversationsChatsReader is a Reader for the GetConversationsChats structure.
type GetConversationsChatsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConversationsChatsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetConversationsChatsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetConversationsChatsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetConversationsChatsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetConversationsChatsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetConversationsChatsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetConversationsChatsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetConversationsChatsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetConversationsChatsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetConversationsChatsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetConversationsChatsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetConversationsChatsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetConversationsChatsOK creates a GetConversationsChatsOK with default headers values
func NewGetConversationsChatsOK() *GetConversationsChatsOK {
	return &GetConversationsChatsOK{}
}

/*GetConversationsChatsOK handles this case with default header values.

successful operation
*/
type GetConversationsChatsOK struct {
	Payload *models.ChatConversationEntityListing
}

func (o *GetConversationsChatsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats][%d] getConversationsChatsOK  %+v", 200, o.Payload)
}

func (o *GetConversationsChatsOK) GetPayload() *models.ChatConversationEntityListing {
	return o.Payload
}

func (o *GetConversationsChatsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ChatConversationEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsChatsBadRequest creates a GetConversationsChatsBadRequest with default headers values
func NewGetConversationsChatsBadRequest() *GetConversationsChatsBadRequest {
	return &GetConversationsChatsBadRequest{}
}

/*GetConversationsChatsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetConversationsChatsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsChatsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats][%d] getConversationsChatsBadRequest  %+v", 400, o.Payload)
}

func (o *GetConversationsChatsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsChatsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsChatsUnauthorized creates a GetConversationsChatsUnauthorized with default headers values
func NewGetConversationsChatsUnauthorized() *GetConversationsChatsUnauthorized {
	return &GetConversationsChatsUnauthorized{}
}

/*GetConversationsChatsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetConversationsChatsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsChatsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats][%d] getConversationsChatsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetConversationsChatsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsChatsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsChatsForbidden creates a GetConversationsChatsForbidden with default headers values
func NewGetConversationsChatsForbidden() *GetConversationsChatsForbidden {
	return &GetConversationsChatsForbidden{}
}

/*GetConversationsChatsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetConversationsChatsForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsChatsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats][%d] getConversationsChatsForbidden  %+v", 403, o.Payload)
}

func (o *GetConversationsChatsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsChatsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsChatsNotFound creates a GetConversationsChatsNotFound with default headers values
func NewGetConversationsChatsNotFound() *GetConversationsChatsNotFound {
	return &GetConversationsChatsNotFound{}
}

/*GetConversationsChatsNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetConversationsChatsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsChatsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats][%d] getConversationsChatsNotFound  %+v", 404, o.Payload)
}

func (o *GetConversationsChatsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsChatsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsChatsRequestEntityTooLarge creates a GetConversationsChatsRequestEntityTooLarge with default headers values
func NewGetConversationsChatsRequestEntityTooLarge() *GetConversationsChatsRequestEntityTooLarge {
	return &GetConversationsChatsRequestEntityTooLarge{}
}

/*GetConversationsChatsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetConversationsChatsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsChatsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats][%d] getConversationsChatsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetConversationsChatsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsChatsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsChatsUnsupportedMediaType creates a GetConversationsChatsUnsupportedMediaType with default headers values
func NewGetConversationsChatsUnsupportedMediaType() *GetConversationsChatsUnsupportedMediaType {
	return &GetConversationsChatsUnsupportedMediaType{}
}

/*GetConversationsChatsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetConversationsChatsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsChatsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats][%d] getConversationsChatsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetConversationsChatsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsChatsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsChatsTooManyRequests creates a GetConversationsChatsTooManyRequests with default headers values
func NewGetConversationsChatsTooManyRequests() *GetConversationsChatsTooManyRequests {
	return &GetConversationsChatsTooManyRequests{}
}

/*GetConversationsChatsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetConversationsChatsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsChatsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats][%d] getConversationsChatsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetConversationsChatsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsChatsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsChatsInternalServerError creates a GetConversationsChatsInternalServerError with default headers values
func NewGetConversationsChatsInternalServerError() *GetConversationsChatsInternalServerError {
	return &GetConversationsChatsInternalServerError{}
}

/*GetConversationsChatsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetConversationsChatsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsChatsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats][%d] getConversationsChatsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetConversationsChatsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsChatsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsChatsServiceUnavailable creates a GetConversationsChatsServiceUnavailable with default headers values
func NewGetConversationsChatsServiceUnavailable() *GetConversationsChatsServiceUnavailable {
	return &GetConversationsChatsServiceUnavailable{}
}

/*GetConversationsChatsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetConversationsChatsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsChatsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats][%d] getConversationsChatsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetConversationsChatsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsChatsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsChatsGatewayTimeout creates a GetConversationsChatsGatewayTimeout with default headers values
func NewGetConversationsChatsGatewayTimeout() *GetConversationsChatsGatewayTimeout {
	return &GetConversationsChatsGatewayTimeout{}
}

/*GetConversationsChatsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetConversationsChatsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsChatsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats][%d] getConversationsChatsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetConversationsChatsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsChatsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
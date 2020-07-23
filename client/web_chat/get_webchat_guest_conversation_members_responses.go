// Code generated by go-swagger; DO NOT EDIT.

package web_chat

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetWebchatGuestConversationMembersReader is a Reader for the GetWebchatGuestConversationMembers structure.
type GetWebchatGuestConversationMembersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWebchatGuestConversationMembersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWebchatGuestConversationMembersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetWebchatGuestConversationMembersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetWebchatGuestConversationMembersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetWebchatGuestConversationMembersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetWebchatGuestConversationMembersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetWebchatGuestConversationMembersRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetWebchatGuestConversationMembersUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetWebchatGuestConversationMembersTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetWebchatGuestConversationMembersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetWebchatGuestConversationMembersServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetWebchatGuestConversationMembersGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetWebchatGuestConversationMembersOK creates a GetWebchatGuestConversationMembersOK with default headers values
func NewGetWebchatGuestConversationMembersOK() *GetWebchatGuestConversationMembersOK {
	return &GetWebchatGuestConversationMembersOK{}
}

/*GetWebchatGuestConversationMembersOK handles this case with default header values.

successful operation
*/
type GetWebchatGuestConversationMembersOK struct {
	Payload *models.WebChatMemberInfoEntityList
}

func (o *GetWebchatGuestConversationMembersOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/webchat/guest/conversations/{conversationId}/members][%d] getWebchatGuestConversationMembersOK  %+v", 200, o.Payload)
}

func (o *GetWebchatGuestConversationMembersOK) GetPayload() *models.WebChatMemberInfoEntityList {
	return o.Payload
}

func (o *GetWebchatGuestConversationMembersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WebChatMemberInfoEntityList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWebchatGuestConversationMembersBadRequest creates a GetWebchatGuestConversationMembersBadRequest with default headers values
func NewGetWebchatGuestConversationMembersBadRequest() *GetWebchatGuestConversationMembersBadRequest {
	return &GetWebchatGuestConversationMembersBadRequest{}
}

/*GetWebchatGuestConversationMembersBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetWebchatGuestConversationMembersBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetWebchatGuestConversationMembersBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/webchat/guest/conversations/{conversationId}/members][%d] getWebchatGuestConversationMembersBadRequest  %+v", 400, o.Payload)
}

func (o *GetWebchatGuestConversationMembersBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWebchatGuestConversationMembersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWebchatGuestConversationMembersUnauthorized creates a GetWebchatGuestConversationMembersUnauthorized with default headers values
func NewGetWebchatGuestConversationMembersUnauthorized() *GetWebchatGuestConversationMembersUnauthorized {
	return &GetWebchatGuestConversationMembersUnauthorized{}
}

/*GetWebchatGuestConversationMembersUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetWebchatGuestConversationMembersUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetWebchatGuestConversationMembersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/webchat/guest/conversations/{conversationId}/members][%d] getWebchatGuestConversationMembersUnauthorized  %+v", 401, o.Payload)
}

func (o *GetWebchatGuestConversationMembersUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWebchatGuestConversationMembersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWebchatGuestConversationMembersForbidden creates a GetWebchatGuestConversationMembersForbidden with default headers values
func NewGetWebchatGuestConversationMembersForbidden() *GetWebchatGuestConversationMembersForbidden {
	return &GetWebchatGuestConversationMembersForbidden{}
}

/*GetWebchatGuestConversationMembersForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetWebchatGuestConversationMembersForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetWebchatGuestConversationMembersForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/webchat/guest/conversations/{conversationId}/members][%d] getWebchatGuestConversationMembersForbidden  %+v", 403, o.Payload)
}

func (o *GetWebchatGuestConversationMembersForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWebchatGuestConversationMembersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWebchatGuestConversationMembersNotFound creates a GetWebchatGuestConversationMembersNotFound with default headers values
func NewGetWebchatGuestConversationMembersNotFound() *GetWebchatGuestConversationMembersNotFound {
	return &GetWebchatGuestConversationMembersNotFound{}
}

/*GetWebchatGuestConversationMembersNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetWebchatGuestConversationMembersNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetWebchatGuestConversationMembersNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/webchat/guest/conversations/{conversationId}/members][%d] getWebchatGuestConversationMembersNotFound  %+v", 404, o.Payload)
}

func (o *GetWebchatGuestConversationMembersNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWebchatGuestConversationMembersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWebchatGuestConversationMembersRequestEntityTooLarge creates a GetWebchatGuestConversationMembersRequestEntityTooLarge with default headers values
func NewGetWebchatGuestConversationMembersRequestEntityTooLarge() *GetWebchatGuestConversationMembersRequestEntityTooLarge {
	return &GetWebchatGuestConversationMembersRequestEntityTooLarge{}
}

/*GetWebchatGuestConversationMembersRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetWebchatGuestConversationMembersRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetWebchatGuestConversationMembersRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/webchat/guest/conversations/{conversationId}/members][%d] getWebchatGuestConversationMembersRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetWebchatGuestConversationMembersRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWebchatGuestConversationMembersRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWebchatGuestConversationMembersUnsupportedMediaType creates a GetWebchatGuestConversationMembersUnsupportedMediaType with default headers values
func NewGetWebchatGuestConversationMembersUnsupportedMediaType() *GetWebchatGuestConversationMembersUnsupportedMediaType {
	return &GetWebchatGuestConversationMembersUnsupportedMediaType{}
}

/*GetWebchatGuestConversationMembersUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetWebchatGuestConversationMembersUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetWebchatGuestConversationMembersUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/webchat/guest/conversations/{conversationId}/members][%d] getWebchatGuestConversationMembersUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetWebchatGuestConversationMembersUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWebchatGuestConversationMembersUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWebchatGuestConversationMembersTooManyRequests creates a GetWebchatGuestConversationMembersTooManyRequests with default headers values
func NewGetWebchatGuestConversationMembersTooManyRequests() *GetWebchatGuestConversationMembersTooManyRequests {
	return &GetWebchatGuestConversationMembersTooManyRequests{}
}

/*GetWebchatGuestConversationMembersTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetWebchatGuestConversationMembersTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetWebchatGuestConversationMembersTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/webchat/guest/conversations/{conversationId}/members][%d] getWebchatGuestConversationMembersTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetWebchatGuestConversationMembersTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWebchatGuestConversationMembersTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWebchatGuestConversationMembersInternalServerError creates a GetWebchatGuestConversationMembersInternalServerError with default headers values
func NewGetWebchatGuestConversationMembersInternalServerError() *GetWebchatGuestConversationMembersInternalServerError {
	return &GetWebchatGuestConversationMembersInternalServerError{}
}

/*GetWebchatGuestConversationMembersInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetWebchatGuestConversationMembersInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetWebchatGuestConversationMembersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/webchat/guest/conversations/{conversationId}/members][%d] getWebchatGuestConversationMembersInternalServerError  %+v", 500, o.Payload)
}

func (o *GetWebchatGuestConversationMembersInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWebchatGuestConversationMembersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWebchatGuestConversationMembersServiceUnavailable creates a GetWebchatGuestConversationMembersServiceUnavailable with default headers values
func NewGetWebchatGuestConversationMembersServiceUnavailable() *GetWebchatGuestConversationMembersServiceUnavailable {
	return &GetWebchatGuestConversationMembersServiceUnavailable{}
}

/*GetWebchatGuestConversationMembersServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetWebchatGuestConversationMembersServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetWebchatGuestConversationMembersServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/webchat/guest/conversations/{conversationId}/members][%d] getWebchatGuestConversationMembersServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetWebchatGuestConversationMembersServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWebchatGuestConversationMembersServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWebchatGuestConversationMembersGatewayTimeout creates a GetWebchatGuestConversationMembersGatewayTimeout with default headers values
func NewGetWebchatGuestConversationMembersGatewayTimeout() *GetWebchatGuestConversationMembersGatewayTimeout {
	return &GetWebchatGuestConversationMembersGatewayTimeout{}
}

/*GetWebchatGuestConversationMembersGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetWebchatGuestConversationMembersGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetWebchatGuestConversationMembersGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/webchat/guest/conversations/{conversationId}/members][%d] getWebchatGuestConversationMembersGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetWebchatGuestConversationMembersGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWebchatGuestConversationMembersGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

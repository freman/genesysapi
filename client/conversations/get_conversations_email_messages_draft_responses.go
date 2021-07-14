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

// GetConversationsEmailMessagesDraftReader is a Reader for the GetConversationsEmailMessagesDraft structure.
type GetConversationsEmailMessagesDraftReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConversationsEmailMessagesDraftReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetConversationsEmailMessagesDraftOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetConversationsEmailMessagesDraftBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetConversationsEmailMessagesDraftUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetConversationsEmailMessagesDraftForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetConversationsEmailMessagesDraftNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetConversationsEmailMessagesDraftRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetConversationsEmailMessagesDraftRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetConversationsEmailMessagesDraftUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetConversationsEmailMessagesDraftTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetConversationsEmailMessagesDraftInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetConversationsEmailMessagesDraftServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetConversationsEmailMessagesDraftGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetConversationsEmailMessagesDraftOK creates a GetConversationsEmailMessagesDraftOK with default headers values
func NewGetConversationsEmailMessagesDraftOK() *GetConversationsEmailMessagesDraftOK {
	return &GetConversationsEmailMessagesDraftOK{}
}

/*GetConversationsEmailMessagesDraftOK handles this case with default header values.

successful operation
*/
type GetConversationsEmailMessagesDraftOK struct {
	Payload *models.EmailMessage
}

func (o *GetConversationsEmailMessagesDraftOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/emails/{conversationId}/messages/draft][%d] getConversationsEmailMessagesDraftOK  %+v", 200, o.Payload)
}

func (o *GetConversationsEmailMessagesDraftOK) GetPayload() *models.EmailMessage {
	return o.Payload
}

func (o *GetConversationsEmailMessagesDraftOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EmailMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsEmailMessagesDraftBadRequest creates a GetConversationsEmailMessagesDraftBadRequest with default headers values
func NewGetConversationsEmailMessagesDraftBadRequest() *GetConversationsEmailMessagesDraftBadRequest {
	return &GetConversationsEmailMessagesDraftBadRequest{}
}

/*GetConversationsEmailMessagesDraftBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetConversationsEmailMessagesDraftBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsEmailMessagesDraftBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/emails/{conversationId}/messages/draft][%d] getConversationsEmailMessagesDraftBadRequest  %+v", 400, o.Payload)
}

func (o *GetConversationsEmailMessagesDraftBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsEmailMessagesDraftBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsEmailMessagesDraftUnauthorized creates a GetConversationsEmailMessagesDraftUnauthorized with default headers values
func NewGetConversationsEmailMessagesDraftUnauthorized() *GetConversationsEmailMessagesDraftUnauthorized {
	return &GetConversationsEmailMessagesDraftUnauthorized{}
}

/*GetConversationsEmailMessagesDraftUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetConversationsEmailMessagesDraftUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsEmailMessagesDraftUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/emails/{conversationId}/messages/draft][%d] getConversationsEmailMessagesDraftUnauthorized  %+v", 401, o.Payload)
}

func (o *GetConversationsEmailMessagesDraftUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsEmailMessagesDraftUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsEmailMessagesDraftForbidden creates a GetConversationsEmailMessagesDraftForbidden with default headers values
func NewGetConversationsEmailMessagesDraftForbidden() *GetConversationsEmailMessagesDraftForbidden {
	return &GetConversationsEmailMessagesDraftForbidden{}
}

/*GetConversationsEmailMessagesDraftForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetConversationsEmailMessagesDraftForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsEmailMessagesDraftForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/emails/{conversationId}/messages/draft][%d] getConversationsEmailMessagesDraftForbidden  %+v", 403, o.Payload)
}

func (o *GetConversationsEmailMessagesDraftForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsEmailMessagesDraftForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsEmailMessagesDraftNotFound creates a GetConversationsEmailMessagesDraftNotFound with default headers values
func NewGetConversationsEmailMessagesDraftNotFound() *GetConversationsEmailMessagesDraftNotFound {
	return &GetConversationsEmailMessagesDraftNotFound{}
}

/*GetConversationsEmailMessagesDraftNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetConversationsEmailMessagesDraftNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsEmailMessagesDraftNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/emails/{conversationId}/messages/draft][%d] getConversationsEmailMessagesDraftNotFound  %+v", 404, o.Payload)
}

func (o *GetConversationsEmailMessagesDraftNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsEmailMessagesDraftNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsEmailMessagesDraftRequestTimeout creates a GetConversationsEmailMessagesDraftRequestTimeout with default headers values
func NewGetConversationsEmailMessagesDraftRequestTimeout() *GetConversationsEmailMessagesDraftRequestTimeout {
	return &GetConversationsEmailMessagesDraftRequestTimeout{}
}

/*GetConversationsEmailMessagesDraftRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetConversationsEmailMessagesDraftRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsEmailMessagesDraftRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/emails/{conversationId}/messages/draft][%d] getConversationsEmailMessagesDraftRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetConversationsEmailMessagesDraftRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsEmailMessagesDraftRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsEmailMessagesDraftRequestEntityTooLarge creates a GetConversationsEmailMessagesDraftRequestEntityTooLarge with default headers values
func NewGetConversationsEmailMessagesDraftRequestEntityTooLarge() *GetConversationsEmailMessagesDraftRequestEntityTooLarge {
	return &GetConversationsEmailMessagesDraftRequestEntityTooLarge{}
}

/*GetConversationsEmailMessagesDraftRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetConversationsEmailMessagesDraftRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsEmailMessagesDraftRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/emails/{conversationId}/messages/draft][%d] getConversationsEmailMessagesDraftRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetConversationsEmailMessagesDraftRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsEmailMessagesDraftRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsEmailMessagesDraftUnsupportedMediaType creates a GetConversationsEmailMessagesDraftUnsupportedMediaType with default headers values
func NewGetConversationsEmailMessagesDraftUnsupportedMediaType() *GetConversationsEmailMessagesDraftUnsupportedMediaType {
	return &GetConversationsEmailMessagesDraftUnsupportedMediaType{}
}

/*GetConversationsEmailMessagesDraftUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetConversationsEmailMessagesDraftUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsEmailMessagesDraftUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/emails/{conversationId}/messages/draft][%d] getConversationsEmailMessagesDraftUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetConversationsEmailMessagesDraftUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsEmailMessagesDraftUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsEmailMessagesDraftTooManyRequests creates a GetConversationsEmailMessagesDraftTooManyRequests with default headers values
func NewGetConversationsEmailMessagesDraftTooManyRequests() *GetConversationsEmailMessagesDraftTooManyRequests {
	return &GetConversationsEmailMessagesDraftTooManyRequests{}
}

/*GetConversationsEmailMessagesDraftTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetConversationsEmailMessagesDraftTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsEmailMessagesDraftTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/emails/{conversationId}/messages/draft][%d] getConversationsEmailMessagesDraftTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetConversationsEmailMessagesDraftTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsEmailMessagesDraftTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsEmailMessagesDraftInternalServerError creates a GetConversationsEmailMessagesDraftInternalServerError with default headers values
func NewGetConversationsEmailMessagesDraftInternalServerError() *GetConversationsEmailMessagesDraftInternalServerError {
	return &GetConversationsEmailMessagesDraftInternalServerError{}
}

/*GetConversationsEmailMessagesDraftInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetConversationsEmailMessagesDraftInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsEmailMessagesDraftInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/emails/{conversationId}/messages/draft][%d] getConversationsEmailMessagesDraftInternalServerError  %+v", 500, o.Payload)
}

func (o *GetConversationsEmailMessagesDraftInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsEmailMessagesDraftInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsEmailMessagesDraftServiceUnavailable creates a GetConversationsEmailMessagesDraftServiceUnavailable with default headers values
func NewGetConversationsEmailMessagesDraftServiceUnavailable() *GetConversationsEmailMessagesDraftServiceUnavailable {
	return &GetConversationsEmailMessagesDraftServiceUnavailable{}
}

/*GetConversationsEmailMessagesDraftServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetConversationsEmailMessagesDraftServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsEmailMessagesDraftServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/emails/{conversationId}/messages/draft][%d] getConversationsEmailMessagesDraftServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetConversationsEmailMessagesDraftServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsEmailMessagesDraftServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsEmailMessagesDraftGatewayTimeout creates a GetConversationsEmailMessagesDraftGatewayTimeout with default headers values
func NewGetConversationsEmailMessagesDraftGatewayTimeout() *GetConversationsEmailMessagesDraftGatewayTimeout {
	return &GetConversationsEmailMessagesDraftGatewayTimeout{}
}

/*GetConversationsEmailMessagesDraftGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetConversationsEmailMessagesDraftGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsEmailMessagesDraftGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/emails/{conversationId}/messages/draft][%d] getConversationsEmailMessagesDraftGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetConversationsEmailMessagesDraftGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsEmailMessagesDraftGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

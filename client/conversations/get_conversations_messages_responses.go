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

// GetConversationsMessagesReader is a Reader for the GetConversationsMessages structure.
type GetConversationsMessagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConversationsMessagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetConversationsMessagesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetConversationsMessagesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetConversationsMessagesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetConversationsMessagesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetConversationsMessagesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetConversationsMessagesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetConversationsMessagesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetConversationsMessagesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetConversationsMessagesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetConversationsMessagesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetConversationsMessagesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetConversationsMessagesOK creates a GetConversationsMessagesOK with default headers values
func NewGetConversationsMessagesOK() *GetConversationsMessagesOK {
	return &GetConversationsMessagesOK{}
}

/*GetConversationsMessagesOK handles this case with default header values.

successful operation
*/
type GetConversationsMessagesOK struct {
	Payload *models.MessageConversationEntityListing
}

func (o *GetConversationsMessagesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages][%d] getConversationsMessagesOK  %+v", 200, o.Payload)
}

func (o *GetConversationsMessagesOK) GetPayload() *models.MessageConversationEntityListing {
	return o.Payload
}

func (o *GetConversationsMessagesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MessageConversationEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagesBadRequest creates a GetConversationsMessagesBadRequest with default headers values
func NewGetConversationsMessagesBadRequest() *GetConversationsMessagesBadRequest {
	return &GetConversationsMessagesBadRequest{}
}

/*GetConversationsMessagesBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetConversationsMessagesBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessagesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages][%d] getConversationsMessagesBadRequest  %+v", 400, o.Payload)
}

func (o *GetConversationsMessagesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagesUnauthorized creates a GetConversationsMessagesUnauthorized with default headers values
func NewGetConversationsMessagesUnauthorized() *GetConversationsMessagesUnauthorized {
	return &GetConversationsMessagesUnauthorized{}
}

/*GetConversationsMessagesUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetConversationsMessagesUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessagesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages][%d] getConversationsMessagesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetConversationsMessagesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagesForbidden creates a GetConversationsMessagesForbidden with default headers values
func NewGetConversationsMessagesForbidden() *GetConversationsMessagesForbidden {
	return &GetConversationsMessagesForbidden{}
}

/*GetConversationsMessagesForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetConversationsMessagesForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessagesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages][%d] getConversationsMessagesForbidden  %+v", 403, o.Payload)
}

func (o *GetConversationsMessagesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagesNotFound creates a GetConversationsMessagesNotFound with default headers values
func NewGetConversationsMessagesNotFound() *GetConversationsMessagesNotFound {
	return &GetConversationsMessagesNotFound{}
}

/*GetConversationsMessagesNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetConversationsMessagesNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessagesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages][%d] getConversationsMessagesNotFound  %+v", 404, o.Payload)
}

func (o *GetConversationsMessagesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagesRequestEntityTooLarge creates a GetConversationsMessagesRequestEntityTooLarge with default headers values
func NewGetConversationsMessagesRequestEntityTooLarge() *GetConversationsMessagesRequestEntityTooLarge {
	return &GetConversationsMessagesRequestEntityTooLarge{}
}

/*GetConversationsMessagesRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetConversationsMessagesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessagesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages][%d] getConversationsMessagesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetConversationsMessagesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagesUnsupportedMediaType creates a GetConversationsMessagesUnsupportedMediaType with default headers values
func NewGetConversationsMessagesUnsupportedMediaType() *GetConversationsMessagesUnsupportedMediaType {
	return &GetConversationsMessagesUnsupportedMediaType{}
}

/*GetConversationsMessagesUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetConversationsMessagesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessagesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages][%d] getConversationsMessagesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetConversationsMessagesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagesTooManyRequests creates a GetConversationsMessagesTooManyRequests with default headers values
func NewGetConversationsMessagesTooManyRequests() *GetConversationsMessagesTooManyRequests {
	return &GetConversationsMessagesTooManyRequests{}
}

/*GetConversationsMessagesTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetConversationsMessagesTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessagesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages][%d] getConversationsMessagesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetConversationsMessagesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagesInternalServerError creates a GetConversationsMessagesInternalServerError with default headers values
func NewGetConversationsMessagesInternalServerError() *GetConversationsMessagesInternalServerError {
	return &GetConversationsMessagesInternalServerError{}
}

/*GetConversationsMessagesInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetConversationsMessagesInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessagesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages][%d] getConversationsMessagesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetConversationsMessagesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagesServiceUnavailable creates a GetConversationsMessagesServiceUnavailable with default headers values
func NewGetConversationsMessagesServiceUnavailable() *GetConversationsMessagesServiceUnavailable {
	return &GetConversationsMessagesServiceUnavailable{}
}

/*GetConversationsMessagesServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetConversationsMessagesServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessagesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages][%d] getConversationsMessagesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetConversationsMessagesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagesGatewayTimeout creates a GetConversationsMessagesGatewayTimeout with default headers values
func NewGetConversationsMessagesGatewayTimeout() *GetConversationsMessagesGatewayTimeout {
	return &GetConversationsMessagesGatewayTimeout{}
}

/*GetConversationsMessagesGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetConversationsMessagesGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessagesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages][%d] getConversationsMessagesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetConversationsMessagesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
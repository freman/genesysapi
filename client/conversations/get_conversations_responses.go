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

// GetConversationsReader is a Reader for the GetConversations structure.
type GetConversationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConversationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetConversationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetConversationsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetConversationsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetConversationsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetConversationsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetConversationsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetConversationsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetConversationsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetConversationsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetConversationsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetConversationsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetConversationsOK creates a GetConversationsOK with default headers values
func NewGetConversationsOK() *GetConversationsOK {
	return &GetConversationsOK{}
}

/*GetConversationsOK handles this case with default header values.

successful operation
*/
type GetConversationsOK struct {
	Payload *models.ConversationEntityListing
}

func (o *GetConversationsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations][%d] getConversationsOK  %+v", 200, o.Payload)
}

func (o *GetConversationsOK) GetPayload() *models.ConversationEntityListing {
	return o.Payload
}

func (o *GetConversationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConversationEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsBadRequest creates a GetConversationsBadRequest with default headers values
func NewGetConversationsBadRequest() *GetConversationsBadRequest {
	return &GetConversationsBadRequest{}
}

/*GetConversationsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetConversationsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations][%d] getConversationsBadRequest  %+v", 400, o.Payload)
}

func (o *GetConversationsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsUnauthorized creates a GetConversationsUnauthorized with default headers values
func NewGetConversationsUnauthorized() *GetConversationsUnauthorized {
	return &GetConversationsUnauthorized{}
}

/*GetConversationsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetConversationsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations][%d] getConversationsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetConversationsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsForbidden creates a GetConversationsForbidden with default headers values
func NewGetConversationsForbidden() *GetConversationsForbidden {
	return &GetConversationsForbidden{}
}

/*GetConversationsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetConversationsForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations][%d] getConversationsForbidden  %+v", 403, o.Payload)
}

func (o *GetConversationsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsNotFound creates a GetConversationsNotFound with default headers values
func NewGetConversationsNotFound() *GetConversationsNotFound {
	return &GetConversationsNotFound{}
}

/*GetConversationsNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetConversationsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations][%d] getConversationsNotFound  %+v", 404, o.Payload)
}

func (o *GetConversationsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsRequestEntityTooLarge creates a GetConversationsRequestEntityTooLarge with default headers values
func NewGetConversationsRequestEntityTooLarge() *GetConversationsRequestEntityTooLarge {
	return &GetConversationsRequestEntityTooLarge{}
}

/*GetConversationsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetConversationsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations][%d] getConversationsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetConversationsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsUnsupportedMediaType creates a GetConversationsUnsupportedMediaType with default headers values
func NewGetConversationsUnsupportedMediaType() *GetConversationsUnsupportedMediaType {
	return &GetConversationsUnsupportedMediaType{}
}

/*GetConversationsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetConversationsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations][%d] getConversationsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetConversationsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsTooManyRequests creates a GetConversationsTooManyRequests with default headers values
func NewGetConversationsTooManyRequests() *GetConversationsTooManyRequests {
	return &GetConversationsTooManyRequests{}
}

/*GetConversationsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetConversationsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations][%d] getConversationsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetConversationsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsInternalServerError creates a GetConversationsInternalServerError with default headers values
func NewGetConversationsInternalServerError() *GetConversationsInternalServerError {
	return &GetConversationsInternalServerError{}
}

/*GetConversationsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetConversationsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations][%d] getConversationsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetConversationsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsServiceUnavailable creates a GetConversationsServiceUnavailable with default headers values
func NewGetConversationsServiceUnavailable() *GetConversationsServiceUnavailable {
	return &GetConversationsServiceUnavailable{}
}

/*GetConversationsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetConversationsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations][%d] getConversationsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetConversationsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsGatewayTimeout creates a GetConversationsGatewayTimeout with default headers values
func NewGetConversationsGatewayTimeout() *GetConversationsGatewayTimeout {
	return &GetConversationsGatewayTimeout{}
}

/*GetConversationsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetConversationsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations][%d] getConversationsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetConversationsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
